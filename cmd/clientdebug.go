// 文件说明：客户端实时调试的 Agent 命令入口，使用平台仅有的 stream 与 message 两条接口。
package cmd

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/auth"
	"gitee.com/unitedrhino/cli/internal/config"
	"github.com/spf13/cobra"
)

// clientDebugCmd 是客户端调试命令组。
var clientDebugCmd = &cobra.Command{Use: "client-debug", Short: "客户端实时日志与逐条确认的诊断动作"}

// clientDebugWatchOpts 保存 SSE 会话的目标和过滤条件。
var clientDebugWatchOpts struct {
	userID     string
	instanceID string
	levels     []string
	names      []string
}

// clientDebugMessageOpts 保存 Agent 发送的操控请求或动作参数。
var clientDebugMessageOpts struct {
	sessionID string
	action    string
	path      string
}

// init 注册 watch、control 与 command，保持 CLI 只有两个后端端点。
func init() {
	watch := &cobra.Command{Use: "watch", Short: "订阅指定客户端实例的实时日志", RunE: runClientDebugWatch}
	watch.Flags().StringVar(&clientDebugWatchOpts.userID, "user-id", "", "目标用户 ID")
	watch.Flags().StringVar(&clientDebugWatchOpts.instanceID, "instance-id", "", "客户端设置页展示的实例 ID")
	watch.Flags().StringSliceVar(&clientDebugWatchOpts.levels, "level", []string{"warn", "error"}, "日志级别，可重复传入")
	watch.Flags().StringSliceVar(&clientDebugWatchOpts.names, "name", []string{"*"}, "日志名称，可重复传入；BLE 原始包需显式选中")
	_ = watch.MarkFlagRequired("user-id")
	_ = watch.MarkFlagRequired("instance-id")
	clientDebugCmd.AddCommand(watch)

	control := &cobra.Command{Use: "control", Short: "请求客户端用户授权诊断动作", RunE: runClientDebugControl}
	control.Flags().StringVar(&clientDebugMessageOpts.sessionID, "session-id", "", "watch 输出的会话 ID")
	_ = control.MarkFlagRequired("session-id")
	clientDebugCmd.AddCommand(control)

	command := &cobra.Command{Use: "command", Short: "提交一条需要客户端用户确认的诊断动作", RunE: runClientDebugCommand}
	command.Flags().StringVar(&clientDebugMessageOpts.sessionID, "session-id", "", "watch 输出的会话 ID")
	command.Flags().StringVar(&clientDebugMessageOpts.action, "action", "", "app.snapshot、app.navigate 或 device-list.refresh")
	command.Flags().StringVar(&clientDebugMessageOpts.path, "path", "", "app.navigate 的白名单页面路径")
	_ = command.MarkFlagRequired("session-id")
	_ = command.MarkFlagRequired("action")
	clientDebugCmd.AddCommand(command)
	RootCmd.AddCommand(clientDebugCmd)
}

// runClientDebugWatch 建立 SSE，逐条输出事件供 Agent 分析，退出时关闭会话。
func runClientDebugWatch(cmd *cobra.Command, _ []string) error {
	body := map[string]any{
		"userID":           clientDebugWatchOpts.userID,
		"clientInstanceID": clientDebugWatchOpts.instanceID,
		"filter":           map[string]any{"levels": clientDebugWatchOpts.levels, "names": clientDebugWatchOpts.names},
	}
	resp, err := clientDebugPost(cmd.Context(), "stream", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return debugHTTPError(resp)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/event-stream") {
		return debugHTTPError(resp)
	}
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	event := "message"
	var lastSeq int64
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "event:") {
			event = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
			continue
		}
		if strings.HasPrefix(line, "data:") {
			payload := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
			if event == "log" {
				var item struct {
					Log struct {
						Seq int64 `json:"seq"`
					} `json:"log"`
				}
				if json.Unmarshal([]byte(payload), &item) == nil && item.Log.Seq > 0 {
					if lastSeq > 0 && item.Log.Seq > lastSeq+1 {
						cmd.Printf("gap {\"from\":%d,\"to\":%d}\n", lastSeq+1, item.Log.Seq-1)
					}
					lastSeq = item.Log.Seq
				}
			}
			cmd.Printf("%s %s\n", event, payload)
		}
	}
	return scanner.Err()
}

// runClientDebugControl 请求客户端用户授权，结果从运行中的 watch 流读取。
func runClientDebugControl(cmd *cobra.Command, _ []string) error {
	return clientDebugSend(cmd, map[string]any{"sessionID": clientDebugMessageOpts.sessionID, "kind": "control.request"})
}

// runClientDebugCommand 仅允许服务端已定义的三个动作，动作结果从 watch 流读取。
func runClientDebugCommand(cmd *cobra.Command, _ []string) error {
	action := clientDebugMessageOpts.action
	if action != "app.snapshot" && action != "app.navigate" && action != "device-list.refresh" {
		return errors.New("不支持的诊断动作")
	}
	if action == "app.navigate" && clientDebugMessageOpts.path == "" {
		return errors.New("app.navigate 需要 --path")
	}
	commandID := fmt.Sprintf("cmd-%d", time.Now().UnixNano())
	body := map[string]any{"sessionID": clientDebugMessageOpts.sessionID, "kind": "command", "commandID": commandID, "action": action}
	if action == "app.navigate" {
		body["args"] = map[string]any{"path": clientDebugMessageOpts.path}
	}
	if err := clientDebugSend(cmd, body); err != nil {
		return err
	}
	cmd.Printf("commandID=%s，等待客户端用户确认；结果见 watch 流\n", commandID)
	return nil
}

// clientDebugSend 将单条控制消息送往共用 message 接口。
func clientDebugSend(cmd *cobra.Command, body map[string]any) error {
	resp, err := clientDebugPost(cmd.Context(), "message", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return debugHTTPError(resp)
	}
	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	if result.Code != 0 && result.Code != 200 {
		return fmt.Errorf("平台拒绝调试消息: %d %s", result.Code, result.Msg)
	}
	return nil
}

// clientDebugPost 按当前 CLI 配置与登录凭据提交调试请求。
func clientDebugPost(ctx context.Context, endpoint string, body map[string]any) (*http.Response, error) {
	baseURL, err := config.GetBaseURL()
	if err != nil {
		return nil, err
	}
	appID, err := config.GetAppID()
	if err != nil {
		return nil, err
	}
	token, err := auth.ResolveToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("请先执行 ur login: %w", err)
	}
	tenantCode, _ := config.GetTenantCode()
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/api/v1/system/client-debug/"+endpoint, bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream, application/json")
	req.Header.Set("app-id", appID)
	req.Header.Set("token", token)
	if tenantCode != "" {
		req.Header.Set("ithings-tenant-code", tenantCode)
	}
	return (&http.Client{}).Do(req)
}

// debugHTTPError 返回服务器的完整错误体，便于 Agent 识别会话过期或权限错误。
func debugHTTPError(resp *http.Response) error {
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 8192))
	return fmt.Errorf("调试接口 HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
}
