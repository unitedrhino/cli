package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/response"
)

var apiOpts struct {
	body       string
	bodyFile   string
	headers    []string
	fields     string
	summarize  bool
	format     string
	transform  string
	output     string
	debug      bool
}

var apiCmd = &cobra.Command{
	Use:   "api <path>",
	Short: "调用平台 API",
	Long:  `通过 CLI 调用联犀平台任意 API 端点，支持自定义 body、header、输出格式。`,
	Example: `  ur api /api/v1/system/user/self/get-one
  ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'
  ur api /api/v1/things/device/info/get-one --fields data.userID,data.userName`,
	Args: cobra.ExactArgs(1),
	RunE: runAPI,
}

func init() {
	apiCmd.Flags().StringVar(&apiOpts.body, "body", "", "JSON 请求体")
	apiCmd.Flags().StringVar(&apiOpts.bodyFile, "body-file", "", "从文件读取请求体")
	apiCmd.Flags().StringArrayVarP(&apiOpts.headers, "header", "H", nil, "自定义请求头 (KEY:VALUE)")
	apiCmd.Flags().StringVar(&apiOpts.fields, "fields", "", "字段过滤（逗号分隔）")
	apiCmd.Flags().BoolVar(&apiOpts.summarize, "summarize", false, "摘要模式输出")
	apiCmd.Flags().StringVar(&apiOpts.format, "format", "", "输出格式 (json, table, csv)")
	apiCmd.Flags().StringVar(&apiOpts.transform, "transform", "", "JSON 路径转换")
	apiCmd.Flags().StringVarP(&apiOpts.output, "output", "o", "", "输出到文件")
	apiCmd.Flags().BoolVar(&apiOpts.debug, "debug", false, "调试模式（显示请求详情）")

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

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path:    path,
		Body:    body,
		Headers: headers,
		Debug:   apiOpts.debug,
	})
	if err != nil {
		return err
	}

	return outputResponse(cmd, resp)
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
		result[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
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
