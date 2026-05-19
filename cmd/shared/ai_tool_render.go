package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
)

func runAiToolRender(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	var id int64
	outputPath := ""

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
		case "--output":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--output 需要参数")
				return 2
			}
			outputPath = args[i+1]
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

	documentMd, _ := artifact["documentMd"].(string)
	manifestJson, _ := artifact["manifestJson"].(string)

	if documentMd == "" {
		fmt.Fprintln(stderr, "document.md 为空")
		return 1
	}

	// 提取组件标签和变量列表
	components := extractComponents(documentMd)
	variables := extractVariables(documentMd)

	// 构建渲染结果 JSON（前端 Vue 组件据此替换占位元素）
	result := map[string]any{
		"markdown":   documentMd,
		"components": components,
		"variables":  variables,
	}

	// 如果提供了 manifest，附加 manifest 信息
	if manifestJson != "" {
		var m map[string]any
		if json.Unmarshal([]byte(manifestJson), &m) == nil {
			result["title"] = m["title"]
			result["output"] = m["output"]
		}
	}

	raw, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(stderr, "序列化失败: %v\n", err)
		return 1
	}

	if outputPath != "" {
		if err := os.WriteFile(outputPath, raw, 0644); err != nil {
			fmt.Fprintf(stderr, "写入文件失败: %v\n", err)
			return 1
		}
		fmt.Fprintf(stdout, "已保存到: %s\n", outputPath)
	} else {
		fmt.Fprintln(stdout, string(raw))
	}
	return 0
}

// extractComponents 从 markdown 中提取组件标签及其属性
func extractComponents(md string) []map[string]string {
	var result []map[string]string

	// 匹配自闭合组件标签: <tag-name attr="value" ... />
	re := regexp.MustCompile(`<([a-z][a-z0-9-]*)\s+([^>]*?)\s*/>`)

	for _, match := range re.FindAllStringSubmatch(md, -1) {
		tag := match[1]
		if !allowedComponents[tag] {
			continue
		}
		props := parseTagAttrs(match[2])
		props["_component"] = tag
		result = append(result, props)
	}

	return result
}

// extractVariables 从 markdown 中提取 {{var}} 变量名
func extractVariables(md string) []string {
	re := regexp.MustCompile(`\{\{(\w+)\}\}`)
	seen := make(map[string]bool)
	var result []string
	for _, match := range re.FindAllStringSubmatch(md, -1) {
		name := match[1]
		if !seen[name] {
			seen[name] = true
			result = append(result, name)
		}
	}
	return result
}

// parseTagAttrs 解析标签属性 key="value"
func parseTagAttrs(raw string) map[string]string {
	props := make(map[string]string)
	re := regexp.MustCompile(`(\w+)=["']([^"']*)["']`)
	for _, match := range re.FindAllStringSubmatch(raw, -1) {
		props[match[1]] = match[2]
	}
	// 也支持无值属性
	flagRe := regexp.MustCompile(`(\w+)(?:\s|$)`)
	for _, match := range flagRe.FindAllStringSubmatch(raw, -1) {
		if _, exists := props[match[1]]; !exists {
			// 检查是否不是 key="value" 的 key 部分
			if !strings.Contains(raw, match[1]+"=") {
				props[match[1]] = "true"
			}
		}
	}
	return props
}
