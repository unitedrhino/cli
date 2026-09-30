// api.go 实现通用 API 调用、项目上下文选择及响应输出。
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/response"
	"github.com/spf13/cobra"
)

var apiOpts struct {
	body     string
	bodyFile string
	headers  []string
	// projectID 保持项目标识为字符串，避免大整数精度损失。
	projectID string
	fields    string
	summarize bool
	format    string
	transform string
	output    string
	debug     bool
	stream    bool
}

var apiCmd = &cobra.Command{
	Use:   "api <path>",
	Short: "调用平台 API",
	Long:  `通过 CLI 调用联犀平台任意 API 端点，支持自定义 body、header、输出格式与 SSE 实时流。`,
	Example: `  ur api /api/v1/system/user/self/get-one
  ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'
  ur api /api/v1/things/device/info/get-one --fields data.userID,data.userName
  ur api /api/v1/system/client-debug/stream --stream --body '{"userID":"42","clientInstanceID":"uni-..."}'`,
	Args: cobra.ExactArgs(1),
	RunE: runAPI,
}

func init() {
	apiCmd.Flags().StringVar(&apiOpts.body, "body", "", "JSON 请求体")
	apiCmd.Flags().StringVar(&apiOpts.bodyFile, "body-file", "", "从文件读取请求体")
	apiCmd.Flags().StringArrayVarP(&apiOpts.headers, "header", "H", nil, "自定义请求头 (KEY:VALUE)")
	apiCmd.Flags().StringVar(&apiOpts.projectID, "project-id", "", "项目 ID 字符串（默认 UR_PROJECT_ID，与显式项目头冲突时报错）")
	apiCmd.Flags().StringVar(&apiOpts.fields, "fields", "", "字段过滤（逗号分隔）")
	apiCmd.Flags().BoolVar(&apiOpts.summarize, "summarize", false, "摘要模式输出")
	apiCmd.Flags().StringVar(&apiOpts.format, "format", "", "输出格式 (json, table, csv)")
	apiCmd.Flags().StringVar(&apiOpts.transform, "transform", "", "JSON 路径转换")
	apiCmd.Flags().StringVarP(&apiOpts.output, "output", "o", "", "输出到文件")
	apiCmd.Flags().BoolVar(&apiOpts.debug, "debug", false, "调试模式（显示请求详情）")
	apiCmd.Flags().BoolVar(&apiOpts.stream, "stream", false, "持续读取 SSE，逐行输出 event/data JSON")

	RootCmd.AddCommand(apiCmd)
}

func runAPI(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	path := args[0]

	body, err := resolveBody(apiOpts.body, apiOpts.bodyFile)
	if err != nil {
		return err
	}

	headers, err := parseHeaders(apiOpts.headers)
	if err != nil {
		return err
	}

	if err := client.ApplyProjectID(headers, apiOpts.projectID, cmd.Flags().Changed("project-id"), os.Getenv("UR_PROJECT_ID")); err != nil {
		return err
	}

	req := client.APIRequest{
		Path:    path,
		Body:    body,
		Headers: headers,
		Debug:   apiOpts.debug,
	}
	if apiOpts.stream {
		if apiOpts.fields != "" || apiOpts.summarize || apiOpts.format != "" || apiOpts.transform != "" || apiOpts.output != "" {
			return fmt.Errorf("--stream 不支持字段筛选、摘要、格式转换或输出文件")
		}
		return streamAPI(cmd, req)
	}
	resp, err := client.DoAPI(ctx, req)
	if err != nil {
		return err
	}

	return outputResponse(cmd, resp)
}

// streamAPI 持续输出通用 SSE 事件，每行是可由 AI 解析的 JSON 对象。
func streamAPI(cmd *cobra.Command, req client.APIRequest) error {
	resp, err := client.OpenAPIStream(cmd.Context(), req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	eventName := "message"
	var dataLines []string
	emit := func() {
		if len(dataLines) == 0 {
			return
		}
		data := json.RawMessage(strings.Join(dataLines, "\n"))
		if !json.Valid(data) {
			encoded, _ := json.Marshal(string(data))
			data = encoded
		}
		encoded, _ := json.Marshal(map[string]any{"event": eventName, "data": data})
		cmd.Println(string(encoded))
		eventName = "message"
		dataLines = nil
	}
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\r")
		switch {
		case line == "":
			emit()
		case strings.HasPrefix(line, "event:"):
			eventName = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
		case strings.HasPrefix(line, "data:"):
			dataLines = append(dataLines, strings.TrimPrefix(strings.TrimPrefix(line, "data:"), " "))
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	emit()
	return nil
}

func resolveBody(body, bodyFile string) (map[string]any, error) {
	var raw string
	if bodyFile != "" {
		data, err := os.ReadFile(bodyFile)
		if err != nil {
			return nil, fmt.Errorf("读取 body 文件: %w", err)
		}
		raw = string(data)
	} else {
		raw = body
	}

	if strings.TrimSpace(raw) == "" {
		return map[string]any{}, nil
	}

	var out map[string]any
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		return nil, fmt.Errorf("body 必须是 JSON 对象: %w", err)
	}
	return out, nil
}

func parseHeaders(headers []string) (map[string]string, error) {
	result := make(map[string]string)
	for _, h := range headers {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("无效 header 格式: %q", h)
		}
		// 项目头即使重复使用相同大小写，也不能静默覆盖不同项目。
		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		if strings.EqualFold(key, "project-id") {
			key = "project-id"
			if previous, exists := result[key]; exists && previous != value {
				return nil, fmt.Errorf("项目上下文冲突：重复 project-id 请求头必须一致")
			}
		}
		result[key] = value
	}
	return result, nil
}

func outputResponse(cmd *cobra.Command, resp client.APIResponse) error {
	code := normalizeCode(resp.Code)
	if code != 200 {
		cmd.PrintErrf("[错误] 业务返回 code=%d: %s\n", resp.Code, resp.Msg)
	}

	var out any = resp
	if apiOpts.fields != "" {
		respMap, err := toMapAny(resp)
		if err != nil {
			return err
		}
		selectors := strings.Split(apiOpts.fields, ",")
		filtered, err := response.FilterFields(respMap, selectors)
		if err != nil {
			return err
		}
		out = filtered
	} else if apiOpts.summarize {
		respMap, err := toMapAny(resp)
		if err != nil {
			return err
		}
		out = response.Summarize(respMap)
	}

	raw, err := response.FormatOutput(out, response.FormatOptions{
		Format:    apiOpts.format,
		Transform: apiOpts.transform,
	})
	if err != nil {
		return err
	}

	if apiOpts.output != "" {
		if err := os.WriteFile(apiOpts.output, raw, 0644); err != nil {
			return fmt.Errorf("写入输出文件: %w", err)
		}
		cmd.Printf("输出已保存: %s\n", apiOpts.output)
	} else {
		cmd.Println(string(raw))
	}

	if code != 200 {
		return &CLIError{Message: resp.Msg, ExitCode: 1}
	}
	return nil
}

func normalizeCode(code int) int {
	if code == 0 {
		return 200
	}
	return code
}

func toMapAny(v any) (map[string]any, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	return m, nil
}
