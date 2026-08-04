package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
)

func runAiToolEdit(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	var id int64
	instruction := ""

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
		case "--instruction":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--instruction 需要参数")
				return 2
			}
			instruction = args[i+1]
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
	if instruction == "" {
		fmt.Fprintln(stderr, "必须提供 --instruction")
		return 2
	}

	// 1. 获取当前三件套
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/ai/tool/get-one",
		Body: map[string]any{"id": strconv.FormatInt(id, 10)},
	})
	if err != nil {
		fmt.Fprintf(stderr, "获取三件套失败: %v\n", err)
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
	skillMd, _ := artifact["skillMd"].(string)
	manifestJson, _ := artifact["manifestJson"].(string)

	// 2. 构建 LLM 消息
	// 注意: chat/completions 接口的实际参数由后端定义，这里给出推荐结构
	chatBody := map[string]any{
		"sessionID": fmt.Sprintf("cli-ai-tool-edit-%d", id),
		"messages": []map[string]any{
			{
				"role":    "system",
				"content": buildEditSystemPrompt(),
			},
			{
				"role": "user",
				"content": fmt.Sprintf(
					"## 当前三件套内容\n\n### executor.js\n```javascript\n%s\n```\n\n### skill.md\n```markdown\n%s\n```\n\n### manifest.json\n```json\n%s\n```\n\n## 修改指令\n\n%s",
					executorJs, skillMd, manifestJson, instruction,
				),
			},
		},
	}

	// 3. 调用 AI chat completions
	chatResp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/ai/chat/completions",
		Body: chatBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "AI 调用失败: %v\n", err)
		return 1
	}
	if chatResp.Code != 200 {
		fmt.Fprintf(stderr, "AI API 错误 code=%d: %s\n", chatResp.Code, chatResp.Msg)
		return 1
	}

	chatData, ok := chatResp.Data.(map[string]any)
	if !ok {
		fmt.Fprintln(stderr, "AI 响应 data 格式异常")
		return 1
	}

	content, _ := chatData["content"].(string)
	if content == "" {
		fmt.Fprintln(stderr, "AI 返回内容为空")
		return 1
	}

	// 4. 解析 LLM 返回的 JSON（期望包含 executorJs/skillMd/manifestJson）
	newExecutorJs, newDocumentMd, newManifestJson, parseErr := parseAIResponse(content, executorJs, skillMd, manifestJson)
	if parseErr != nil {
		fmt.Fprintf(stderr, "解析 AI 响应失败: %v\n", parseErr)
		fmt.Fprintf(stderr, "原始响应:\n%s\n", content)
		return 1
	}

	// 5. 保存修改后的三件套
	saveBody := map[string]any{
		"id":         strconv.FormatInt(id, 10),
		"executorJs": newExecutorJs,
	}
	if newDocumentMd != "" {
		saveBody["skillMd"] = newDocumentMd
	}
	if newManifestJson != "" {
		saveBody["manifestJson"] = newManifestJson
	}

	saveResp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/ai/tool/save-artifact",
		Body: saveBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "保存失败: %v\n", err)
		return 1
	}
	if saveResp.Code != 200 {
		fmt.Fprintf(stderr, "保存 API 错误 code=%d: %s\n", saveResp.Code, saveResp.Msg)
		return 1
	}

	fmt.Fprintln(stdout, "AI 编辑完成，已保存")
	return 0
}

func buildEditSystemPrompt() string {
	return `你是 AI 工具开发助手。你的任务是修改 ai-tool 的三件套文件。

## 文件说明
- **executor.js**: QuickJS 执行逻辑，使用 runtime.set/patch 更新状态，tier0.query 查询数据
- **skill.md**: 用户可见文档，Markdown + 组件标签，{{var}} 变量绑定
- **manifest.json**: 元信息（title, runtime, inputs[], output, permissions）

## 组件标签（可在 skill.md 中使用）
| 标签 | 用途 | 属性 |
|------|------|------|
| <chart> | 图表 | type, data, x, y, title, height |
| <metric> | 指标卡片 | title, value, unit, trend |
| <table-cpt> | 数据表格 | data, columns, maxHeight |
| <steps> | 执行步骤 | data[]: {id, title, status, summary} |
| <status> | 状态指示 | value |
| <alert> | 告警 | type, message |
| <mermaid-diagram> | 流程图 | chart |
| <json-view> | JSON查看 | data |

## 变量绑定
- skill.md 中用 {{varName}} 引用变量
- executor.js 中用 runtime.set("varName", value) 写入变量

## 响应格式
返回 JSON:
` + "```json" + `
{
  "executorJs": "...修改后的完整 executor.js...",
  "skillMd": "...修改后的完整 skill.md...",
  "manifestJson": "...修改后的完整 manifest.json..."
}
` + "```" + `

要求:
1. 只修改与指令相关的部分，保持其他内容不变
2. 所有文件内容必须完整，不要省略
3. executor.js 代码要完整可运行`
}

// parseAIResponse 从 LLM 文本响应中提取 JSON
func parseAIResponse(content string, origJs, origMd, origManifest string) (string, string, string, error) {
	// 尝试找到 JSON 代码块
	jsonStart := strings.Index(content, "```json")
	if jsonStart >= 0 {
		jsonStart += 7
		jsonEnd := strings.Index(content[jsonStart:], "```")
		if jsonEnd >= 0 {
			content = content[jsonStart : jsonStart+jsonEnd]
		}
	} else {
		// 尝试直接找到 JSON 对象
		jsonStart = strings.Index(content, "{")
		jsonEnd := strings.LastIndex(content, "}")
		if jsonStart >= 0 && jsonEnd > jsonStart {
			content = content[jsonStart : jsonEnd+1]
		}
	}

	var parsed map[string]string
	if err := json.Unmarshal([]byte(content), &parsed); err != nil {
		// 回退：返回原始内容
		return "", "", "", fmt.Errorf("JSON 解析失败: %w", err)
	}

	executorJs := origJs
	if v, ok := parsed["executorJs"]; ok && v != "" {
		executorJs = v
	}
	skillMd := origMd
	if v, ok := parsed["skillMd"]; ok {
		skillMd = v
	}
	manifestJson := origManifest
	if v, ok := parsed["manifestJson"]; ok {
		manifestJson = v
	}

	if executorJs == origJs && skillMd == origMd && manifestJson == origManifest {
		return "", "", "", fmt.Errorf("AI 未返回任何修改")
	}

	return executorJs, skillMd, manifestJson, nil
}
