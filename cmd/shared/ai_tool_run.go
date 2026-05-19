package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"gitee.com/unitedrhino/cli/internal/client"
)

func runAiToolRun(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	var id int64
	inputs := "{}"
	timeout := 60

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
		case "--inputs":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--inputs 需要 JSON 参数")
				return 2
			}
			inputs = args[i+1]
			i++
		case "--timeout":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--timeout 需要参数（秒）")
				return 2
			}
			var err error
			timeout, err = strconv.Atoi(args[i+1])
			if err != nil {
				fmt.Fprintf(stderr, "--timeout 格式无效: %v\n", err)
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

	// 验证 inputs 是合法 JSON
	var inputsMap map[string]any
	if err := json.Unmarshal([]byte(inputs), &inputsMap); err != nil {
		fmt.Fprintf(stderr, "--inputs JSON 格式无效: %v\n", err)
		return 2
	}

	// 1. 启动运行
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/ai/tool/run",
		Body: map[string]any{
			"toolID": strconv.FormatInt(id, 10),
			"inputs": inputs,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "启动运行失败: %v\n", err)
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
	runID, ok := dataMap["runID"].(string)
	if !ok {
		fmt.Fprintln(stderr, "响应中缺少 runID")
		return 1
	}

	fmt.Fprintf(stdout, "运行已启动, runID: %s\n", runID)

	// 2. 轮询状态
	pollInterval := 2 * time.Second
	deadline := time.After(time.Duration(timeout) * time.Second)
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(stderr, "运行被取消")
			return 1
		case <-deadline:
			fmt.Fprintf(stderr, "运行超时 (%ds)\n", timeout)
			return 1
		case <-ticker.C:
			statusResp, err := client.DoAPI(ctx, client.APIRequest{
				Path: "/api/v1/ai/tool/run-status",
				Body: map[string]any{"runID": runID},
			})
			if err != nil {
				fmt.Fprintf(stderr, "查询状态失败: %v\n", err)
				continue
			}
			if statusResp.Code != 200 {
				fmt.Fprintf(stderr, "查询状态 API 错误 code=%d: %s\n", statusResp.Code, statusResp.Msg)
				continue
			}

			statusData, ok := statusResp.Data.(map[string]any)
			if !ok {
				continue
			}

			status, _ := statusData["status"].(string)
			state, _ := statusData["state"].(string)
			logs, _ := statusData["logs"].(string)
			durationMs, _ := statusData["durationMs"].(string)

			switch status {
			case "success":
				fmt.Fprintln(stdout, "\n✅ 运行成功")
				if state != "" {
					var stateObj any
					if err := json.Unmarshal([]byte(state), &stateObj); err == nil {
						raw, _ := json.MarshalIndent(stateObj, "", "  ")
						fmt.Fprintf(stdout, "状态:\n%s\n", string(raw))
					} else {
						fmt.Fprintf(stdout, "状态: %s\n", state)
					}
				}
				if logs != "" {
					fmt.Fprintf(stdout, "日志:\n%s\n", logs)
				}
				if durationMs != "" {
					fmt.Fprintf(stdout, "耗时: %sms\n", durationMs)
				}
				return 0
			case "failed":
				errMsg, _ := statusData["errorMsg"].(string)
				fmt.Fprintf(stdout, "\n❌ 运行失败\n")
				if errMsg != "" {
					fmt.Fprintf(stdout, "错误: %s\n", errMsg)
				}
				if logs != "" {
					fmt.Fprintf(stdout, "日志:\n%s\n", logs)
				}
				return 1
			default:
				// running — 打印等待提示
				fmt.Fprintf(stdout, ".")
			}
		}
	}
}
