package response

import "fmt"

// Summarize 对 API 响应进行智能摘要，避免 prompt 膨胀。
// 规则：
//   - 业务错误（code != 200 且 code != 0）完整保留
//   - 分页结构保留 total，list 只保留前 5 条摘要
//   - 普通对象保留前 5 个字段
//   - 超长数组截断并加 _note
func Summarize(resp map[string]any) map[string]any {
	code, _ := resp["code"].(float64)
	if code == 0 {
		code = 200
	}
	if int(code) != 200 && int(code) != 0 {
		return resp
	}

	summary := map[string]any{
		"code": resp["code"],
		"msg":  resp["msg"],
	}

	data := resp["data"]
	if data == nil {
		return summary
	}

	d, ok := data.(map[string]any)
	if !ok {
		summary["data"] = data
		return summary
	}

	// 分页结构
	if rawList, ok := d["list"].([]any); ok {
		pageSummary := map[string]any{}
		if total, ok := d["total"].(float64); ok {
			pageSummary["total"] = int64(total)
		}
		pageSummary["list"] = summarizeList(rawList)
		if len(rawList) > 5 {
			pageSummary["_note"] = fmt.Sprintf("共 %d 条，仅展示前 5 条摘要；如需完整数据请使用 --detail", len(rawList))
		}
		summary["data"] = pageSummary
		return summary
	}

	// 普通对象
	summary["data"] = summarizeObject(d)
	return summary
}

var pickKeys = []string{
	"id", "name", "code", "deviceName", "productID",
	"identifier", "tenantCode", "status", "isOnline",
}

func summarizeList(list []any) []any {
	out := make([]any, 0, 5)
	for i, item := range list {
		if i >= 5 {
			break
		}
		out = append(out, summarizeItem(item))
	}
	return out
}

func summarizeItem(item any) any {
	if item == nil {
		return nil
	}
	it, ok := item.(map[string]any)
	if !ok {
		return item
	}

	picked := make(map[string]any)
	for _, k := range pickKeys {
		if v, ok := it[k]; ok {
			picked[k] = v
		}
	}
	if len(picked) == 0 {
		// 没有常见标识字段，保留前 3 个字段
		count := 0
		for k, v := range it {
			if count >= 3 {
				break
			}
			picked[k] = v
			count++
		}
	}
	return picked
}

func summarizeObject(d map[string]any) map[string]any {
	out := make(map[string]any)
	count := 0
	for k, v := range d {
		if count >= 5 {
			out["_note"] = fmt.Sprintf("共 %d 个字段，仅展示前 5 个；如需完整数据请使用 --detail", len(d))
			break
		}
		if arr, ok := v.([]any); ok && len(arr) > 5 {
			out[k] = arr[:5]
			out[k+"_note"] = fmt.Sprintf("共 %d 条，仅展示前 5 条；如需完整数据请使用 --detail", len(arr))
		} else {
			out[k] = v
		}
		count++
	}
	return out
}
