// chat.go — ur ai chat 命令：裸 LLM 调用（agentID=0，无上下文注入）
package ai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/internal/auth"
	"gitee.com/unitedrhino/cli/internal/config"
)

var chatOpts struct {
	message   string
	modelType string
	stream    bool
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "AI 裸 LLM 对话",
	Long: `调用平台 AI 进行裸 LLM 对话（agentID=0），不注入任何上下文（无 Agent 提示词、无 Skill、无知识库、无记忆、无会话记录）。

适用场景：urcli 命令辅助、外部代码调用、调试。

参与企业 token 计费。`,
	Example: `  ur ai chat -m "帮我写一个 Go 的 hello world"
  ur ai chat -m "解释 IoT 协议" --model-type large
  ur ai chat -m "写一首诗" --stream
  echo "总结以下内容" | ur ai chat`,
	RunE: runChat,
}

func init() {
	chatCmd.Flags().StringVarP(&chatOpts.message, "message", "m", "", "消息文本（必填，也可通过 stdin 管道输入）")
	chatCmd.Flags().StringVar(&chatOpts.modelType, "model-type", "", "LLM 模型类型（small/medium/large/xlarge，默认自动选择）")
	chatCmd.Flags().BoolVarP(&chatOpts.stream, "stream", "s", false, "流式输出（逐 token 显示）")

	AICmd.AddCommand(chatCmd)
}

func runChat(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// 获取消息文本：优先 --message，其次 stdin
	message := chatOpts.message
	if message == "" {
		if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
			// 有 stdin 输入
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("读取 stdin 失败: %w", err)
			}
			message = strings.TrimSpace(string(data))
		}
	}
	if message == "" {
		return fmt.Errorf("请通过 -m 指定消息，或通过管道输入")
	}

	// 构造请求体
	reqBody := map[string]any{
		"agentID": 0,
		"messages": []any{
			map[string]any{
				"role": "user",
				"contents": []any{
					map[string]any{
						"type": "text",
						"text": message,
					},
				},
			},
		},
		"modelType": chatOpts.modelType,
		"stream":    chatOpts.stream,
	}

	if chatOpts.stream {
		return runChatStream(ctx, cmd, reqBody)
	}
	return runChatSync(ctx, cmd, reqBody)
}

// runChatSync 非流式调用
func runChatSync(ctx interface{}, cmd *cobra.Command, reqBody map[string]any) error {
	resp, err := doChatRequest(reqBody, false)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("解析响应失败: %w", err)
	}
	if result.Code != 200 {
		return fmt.Errorf("业务错误 code=%d: %s", result.Code, result.Msg)
	}

	cmd.Println(result.Data.Content)
	return nil
}

// runChatStream 流式调用（SSE）
func runChatStream(ctx interface{}, cmd *cobra.Command, reqBody map[string]any) error {
	resp, err := doChatRequest(reqBody, true)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/event-stream") {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("预期 SSE 响应，实际 Content-Type: %s\n%s", contentType, string(body))
	}

	// 解析 SSE 事件流
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		payload := strings.TrimPrefix(line, "data: ")

		var evt struct {
			Token string `json:"token"`
			Done  bool   `json:"done"`
		}
		if err := json.Unmarshal([]byte(payload), &evt); err != nil {
			continue
		}
		if evt.Token != "" {
			cmd.Print(evt.Token)
		}
		if evt.Done {
			cmd.Println() // 结束换行
			return nil
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取 SSE 流失败: %w", err)
	}
	cmd.Println() // 流结束换行
	return nil
}

// doChatRequest 发送 chat/completions 请求
func doChatRequest(body map[string]any, stream bool) (*http.Response, error) {
	baseURL, err := config.GetBaseURL()
	if err != nil {
		return nil, err
	}
	appID, err := config.GetAppID()
	if err != nil {
		return nil, err
	}
	token, err := auth.ResolveToken(context.Background())
	if err != nil {
		return nil, fmt.Errorf("未登录，请先执行 ur setup 或 ur login")
	}
	tenantCode, _ := config.GetTenantCode()

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %w", err)
	}

	url := strings.TrimRight(baseURL, "/") + "/api/v1/ai/chat/completions"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("app-id", appID)
	req.Header.Set("token", token)
	if tenantCode != "" {
		req.Header.Set("tenant-code", tenantCode)
	}

	client := &http.Client{}
	if stream {
		// 流式请求不设超时
		client.Timeout = 0
	} else {
		client.Timeout = 0 // 由服务端控制
	}

	return client.Do(req)
}
