package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
)

// 组件白名单
var allowedComponents = map[string]bool{
	"chart":            true,
	"metric":           true,
	"table-cpt":        true,
	"steps":            true,
	"status":           true,
	"alert":            true,
	"mermaid-diagram":  true,
	"json-view":        true,
}

// JS 禁止的表达式
var jsBlockedPatterns = []string{
	"eval(",
	"new Function(",
	"require(",
	"import(",
}

func runAiToolValidate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	var id int64

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--id 需要参数")
				return 2
			}
			var err error
			id, err = strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				fmt.Fprintf(stderr, "--id 格式无效: %v\n", err)
				return 2
			}
			i++
		default:
			fmt.Fprintf(stderr, "未知选项: %s\n", args[i])
			return 2
		}
	}

	if id == 0 {
		fmt.Fprintln(stderr, "必须提供 --id")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/ai/tool/get-one",
		Body: map[string]any{"id": strconv.FormatInt(id, 10)},
	})
	if err != nil {
		fmt.Fprintf(stderr, "请求失败: %v\n", err)
		return 1
	}
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "API 错误 code=%d: %s\n", resp.Code, resp.Msg)
		return 1
	}

	dataMap, ok := resp.Data.(map[string]any)
	if !ok {
		fmt.Fprintln(stderr, "响应 data 格式异常")
		return 1
	}

	artifact, ok := dataMap["artifact"].(map[string]any)
	if !ok {
		fmt.Fprintln(stderr, "响应中没有 artifact")
		return 1
	}

	executorJs, _ := artifact["executorJs"].(string)
	documentMd, _ := artifact["documentMd"].(string)
	manifestJson, _ := artifact["manifestJson"].(string)

	if executorJs == "" {
		fmt.Fprintln(stderr, "executor.js 为空")
		return 1
	}

	var errors []string
	var warnings []string
	passes := 0
	total := 5

	// 1. manifest.json 格式校验
	if manifestJson != "" {
		var m map[string]any
		if err := json.Unmarshal([]byte(manifestJson), &m); err != nil {
			errors = append(errors, fmt.Sprintf("manifest.json JSON 解析失败: %v", err))
		} else {
			passes++
			if _, ok := m["title"]; !ok {
				warnings = append(warnings, "manifest.json 缺少 title 字段")
			}
		}
	} else {
		warnings = append(warnings, "manifest.json 为空，跳过 JSON 格式校验")
		passes++
	}

	// 2. JS 安全扫描
	jsErrors := validateJSSecurity(executorJs)
	if len(jsErrors) > 0 {
		for _, e := range jsErrors {
			errors = append(errors, e)
		}
	} else {
		passes++
	}

	// 3. 组件标签白名单
	tagErrors, tagWarns := validateComponentTags(documentMd)
	errors = append(errors, tagErrors...)
	warnings = append(warnings, tagWarns...)
	if len(tagErrors) == 0 {
		passes++
	}

	// 4. 变量一致性：document.md 中 {{var}} vs executor.js 中 runtime.set
	varErrors, varWarns := validateVariableConsistency(executorJs, documentMd)
	errors = append(errors, varErrors...)
	warnings = append(warnings, varWarns...)
	if len(varErrors) == 0 {
		passes++
	}

	// 5. executor.js 基本结构检查
	if !strings.Contains(executorJs, "runtime.set(") && !strings.Contains(executorJs, "runtime.patch(") {
		warnings = append(warnings, "executor.js 未使用 runtime.set/patch，工具不会更新状态")
	}
	passes++

	// 输出结果
	fmt.Fprintf(stdout, "校验结果: %d/%d 通过\n", passes, total)

	if len(errors) > 0 {
		fmt.Fprintf(stdout, "\n❌ 错误 (%d):\n", len(errors))
		for _, e := range errors {
			fmt.Fprintf(stdout, "  - %s\n", e)
		}
	}
	if len(warnings) > 0 {
		fmt.Fprintf(stdout, "\n⚠️  警告 (%d):\n", len(warnings))
		for _, w := range warnings {
			fmt.Fprintf(stdout, "  - %s\n", w)
		}
	}

	if len(errors) > 0 {
		return 1
	}
	fmt.Fprintln(stdout, "\n✅ 校验通过")
	return 0
}

func validateJSSecurity(js string) []string {
	var errors []string
	lower := strings.ToLower(js)
	for _, pattern := range jsBlockedPatterns {
		if strings.Contains(lower, strings.ToLower(pattern)) {
			errors = append(errors, fmt.Sprintf("executor.js 包含禁止表达式: %s", pattern))
		}
	}
	return errors
}

func validateComponentTags(md string) (errors []string, warnings []string) {
	// 匹配自定义标签: <tag-name ... />
	re := regexp.MustCompile(`<(/)?([a-z][a-z0-9-]*)(\s[^>]*)?/?>`)

	for _, match := range re.FindAllStringSubmatch(md, -1) {
		if match[1] != "" {
			continue // 闭合标签跳过
		}
		tag := strings.ToLower(match[2])
		// 跳过标准 HTML 标签
		if isStandardHTMLTag(tag) {
			continue
		}
		if !allowedComponents[tag] {
			errors = append(errors, fmt.Sprintf("document.md 包含未注册组件: <%s>（白名单: %s）",
				tag, strings.Join(allowedComponentNames(), ", ")))
		}
	}
	return errors, warnings
}

func isStandardHTMLTag(tag string) bool {
	standard := map[string]bool{
		"div": true, "span": true, "p": true, "a": true, "img": true,
		"h1": true, "h2": true, "h3": true, "h4": true, "h5": true, "h6": true,
		"ul": true, "ol": true, "li": true, "table": true, "tr": true, "td": true,
		"th": true, "thead": true, "tbody": true, "br": true, "hr": true,
		"strong": true, "em": true, "code": true, "pre": true, "blockquote": true,
		"section": true, "header": true, "footer": true, "nav": true, "main": true,
		"input": true, "button": true, "form": true, "label": true, "select": true,
		"option": true, "textarea": true, "style": true, "script": true, "link": true,
		"meta": true, "title": true, "head": true, "body": true, "html": true,
	}
	return standard[tag]
}

func allowedComponentNames() []string {
	names := make([]string, 0, len(allowedComponents))
	for k := range allowedComponents {
		names = append(names, k)
	}
	return names
}

func validateVariableConsistency(js string, md string) (errors []string, warnings []string) {
	// 从 document.md 提取 {{var}} 引用的变量
	varRe := regexp.MustCompile(`\{\{(\w+)\}\}`)
	mdVars := make(map[string]bool)
	for _, match := range varRe.FindAllStringSubmatch(md, -1) {
		mdVars[match[1]] = true
	}

	// 从 executor.js 提取 runtime.set("key",...) 设置的变量
	setRe := regexp.MustCompile(`runtime\.set\s*\(\s*["'](\w+)["']`)
	jsVars := make(map[string]bool)
	for _, match := range setRe.FindAllStringSubmatch(js, -1) {
		jsVars[match[1]] = true
	}

	// 检查 document.md 中引用的变量是否在 executor.js 中设置
	for v := range mdVars {
		if !jsVars[v] {
			warnings = append(warnings, fmt.Sprintf("document.md 引用了变量 {{%s}}，但 executor.js 中未找到 runtime.set(\"%s\",...)", v, v))
		}
	}

	// 检查 executor.js 中 set 的变量是否在 document.md 中使用
	for v := range jsVars {
		if !mdVars[v] {
			warnings = append(warnings, fmt.Sprintf("executor.js 设置了 runtime.set(\"%s\",...)，但 document.md 未引用 {{%s}}", v, v))
		}
	}

	return errors, warnings
}
