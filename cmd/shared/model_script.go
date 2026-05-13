package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func runModelGenerateScript(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: model generate-script <model-file> [--mode up-before|up-after|down-before|down-after] [--output file]")
		return 2
	}

	modelFile := args[0]
	mode := "up-before"
	outputPath := ""

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--mode":
			if i+1 < len(args) {
				mode = args[i+1]
				i++
			}
		case "--output":
			if i+1 < len(args) {
				outputPath = args[i+1]
				i++
			}
		}
	}

	content, err := os.ReadFile(modelFile)
	if err != nil {
		fmt.Fprintf(stderr, "读取物模型文件失败: %v\n", err)
		return 1
	}

	var model map[string]any
	if err := json.Unmarshal(content, &model); err != nil {
		fmt.Fprintf(stderr, "JSON 解析错误: %v\n", err)
		return 1
	}

	script := generateScriptFromModel(model, mode)

	if outputPath != "" {
		if err := os.WriteFile(outputPath, []byte(script), 0644); err != nil {
			fmt.Fprintf(stderr, "写入文件失败: %v\n", err)
			return 1
		}
		fmt.Fprintf(stdout, "脚本已保存: %s\n", outputPath)
	} else {
		fmt.Fprintln(stdout, script)
	}
	return 0
}

func generateScriptFromModel(model map[string]any, mode string) string {
	var props, events, actions []string

	if properties, ok := model["properties"].([]any); ok {
		for _, p := range properties {
			if prop, ok := p.(map[string]any); ok {
				if id, ok := prop["identifier"].(string); ok && id != "" {
					props = append(props, id)
				}
			}
		}
	}

	if evts, ok := model["events"].([]any); ok {
		for _, e := range evts {
			if evt, ok := e.(map[string]any); ok {
				if id, ok := evt["identifier"].(string); ok && id != "" {
					events = append(events, id)
				}
			}
		}
	}

	if acts, ok := model["actions"].([]any); ok {
		for _, a := range acts {
			if act, ok := a.(map[string]any); ok {
				if id, ok := act["identifier"].(string); ok && id != "" {
					actions = append(actions, id)
				}
			}
		}
	}

	var handleSig string
	var returnStmt string
	var comment string

	switch mode {
	case "up-before", "down-before":
		handleSig = "func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg"
		returnStmt = "    return req"
		if mode == "up-before" {
			comment = "// 上行前处理 — 拦截设备上报消息，可修改或丢弃\n// 返回 nil 表示丢弃消息"
		} else {
			comment = "// 下行前处理 — 拦截平台下发指令，可修改或丢弃\n// 返回 nil 表示丢弃消息"
		}
	case "up-after":
		handleSig = "func Handle(ctx context.Context, req *deviceMsg.PublishMsg, resp *deviceMsg.PublishMsg)"
		returnStmt = "    return"
		comment = "// 上行后处理 — 设备消息处理完成后执行，无返回值"
	case "down-after":
		handleSig = "func Handle(ctx context.Context, req *deviceMsg.PublishMsg)"
		returnStmt = "    return"
		comment = "// 下行后处理 — 指令下发后执行，无返回值"
	default:
		handleSig = "func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg"
		returnStmt = "    return req"
		comment = "// 协议脚本 — 请根据物模型填写解析逻辑"
	}

	var sb strings.Builder
	sb.WriteString(comment)
	sb.WriteString("\n\n")
	sb.WriteString("import \"log\"\n")
	sb.WriteString("import \"context\"\n")
	sb.WriteString("import \"deviceMsg\"\n")
	sb.WriteString("import \"json\"\n")
	sb.WriteString("\n")
	sb.WriteString(handleSig)
	sb.WriteString(" {\n")
	sb.WriteString("    log.Printf(\"收到消息: %s\", string(req.Payload))\n")
	sb.WriteString("\n")
	sb.WriteString("    var data map[string]any\n")
	sb.WriteString("    err := json.Unmarshal(req.Payload, &data)\n")
	sb.WriteString("    if err != nil {\n")
	sb.WriteString(returnStmt + "\n")
	sb.WriteString("    }\n")
	sb.WriteString("\n")
	sb.WriteString("    // TODO: 根据物模型映射字段\n")

	if len(props) > 0 {
		sb.WriteString(fmt.Sprintf("    // properties: %s\n", strings.Join(props, ", ")))
	}
	if len(events) > 0 {
		sb.WriteString(fmt.Sprintf("    // events: %s\n", strings.Join(events, ", ")))
	}
	if len(actions) > 0 {
		sb.WriteString(fmt.Sprintf("    // actions: %s\n", strings.Join(actions, ", ")))
	}

	if len(props)+len(events)+len(actions) > 0 {
		sb.WriteString("\n")
		sb.WriteString("    params, _ := data[\"params\"].(map[string]any)\n")
		sb.WriteString("    _ = params\n")
		sb.WriteString("\n")
	}

	sb.WriteString("    // 示例: 字段映射\n")
	sb.WriteString("    // if rawTemp, ok := params[\"temp\"]; ok {\n")
	sb.WriteString("    //     params[\"temperature\"] = rawTemp\n")
	sb.WriteString("    // }\n")
	sb.WriteString("\n")
	sb.WriteString("    req.Payload, _ = json.Marshal(data)\n")
	sb.WriteString(returnStmt + "\n")
	sb.WriteString("}\n")

	return sb.String()
}
