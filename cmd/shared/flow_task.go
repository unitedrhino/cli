// flow_task.go — 流程审批中心命令：审批动线（发起/办理/实例/抄送）。
// 发起、待办办理、我的申请、抄送为全员接口；实例管控（监控/终止/恢复/作废）
// 为管理员接口，后端 ctxs.IsAdmin 兜底校验。
package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runFlowProcess 流程发起命令分发（全员）
func runFlowProcess(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printFlowProcessHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-launch-list":
		return runFlowProcessGetLaunchList(ctx, args[1:], stdout, stderr)
	case "launch":
		return runFlowProcessLaunch(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printFlowProcessHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown process subcommand: %s\n", args[0])
		printFlowProcessHelp(stderr)
		return 2
	}
}

// printFlowProcessHelp 打印流程发起帮助
func printFlowProcessHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur flow process <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-launch-list  可发起流程列表（含分类分组）[--name] [--category] [--page] [--size]")
	fmt.Fprintln(w, "  launch           发起流程 --code <流程编码> [--variable <JSON或@file>] [--title] [--business-key]")
}

// runFlowProcessGetLaunchList 查询可发起流程列表
func runFlowProcessGetLaunchList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{"page": map[string]any{"page": page, "size": size}}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--category":
			reqBody["category"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/process/get-launch-list", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowProcessLaunch 发起流程实例
func runFlowProcessLaunch(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--code":
			reqBody["processCode"] = nextArg(remaining, &i)
		case "--variable":
			v, err := readJSONObjectArg(nextArg(remaining, &i))
			if err != nil {
				fmt.Fprintf(stderr, "Error: --variable %v\n", err)
				return 1
			}
			reqBody["variable"] = v
		case "--title":
			reqBody["title"] = nextArg(remaining, &i)
		case "--business-key":
			reqBody["businessKey"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/process/launch", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// ─── 审批任务办理（全员） ────────────────────────────────────────────────

// runFlowTask 审批任务命令分发
func runFlowTask(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printFlowTaskHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-pending-list":
		return runFlowTaskGetPendingList(ctx, args[1:], stdout, stderr)
	case "get-approved-list":
		return runFlowTaskGetApprovedList(ctx, args[1:], stdout, stderr)
	case "get-detail":
		return runFlowTaskGetDetail(ctx, args[1:], stdout, stderr)
	case "consent":
		return runFlowTaskConsent(ctx, args[1:], stdout, stderr)
	case "reject":
		return runFlowTaskReject(ctx, args[1:], stdout, stderr)
	case "transfer":
		return runFlowTaskTransfer(ctx, args[1:], stdout, stderr)
	case "add-sign":
		return runFlowTaskSign(ctx, "/task/add-sign", args[1:], stdout, stderr)
	case "remove-sign":
		return runFlowTaskSign(ctx, "/task/remove-sign", args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printFlowTaskHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown task subcommand: %s\n", args[0])
		printFlowTaskHelp(stderr)
		return 2
	}
}

// printFlowTaskHelp 打印审批任务帮助
func printFlowTaskHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur flow task <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-pending-list   我的待办 [--task-name] [--process-name] [--page] [--size]")
	fmt.Fprintln(w, "  get-approved-list  我的已办 [--task-name] [--process-name] [--page] [--size]")
	fmt.Fprintln(w, "  get-detail         审批详情（表单+操作权限+记录）--task-id <id>")
	fmt.Fprintln(w, "  consent            同意 --task-id <id> [--comment] [--variable <JSON或@file>]")
	fmt.Fprintln(w, "  reject             驳回 --task-id <id> --strategy <toInitiator|toPreviousNode|toSpecifiedNode|terminateApproval> [--target-node-key] [--comment]")
	fmt.Fprintln(w, "  transfer           转办 --task-id <id> --to-user-id <userID> [--comment]")
	fmt.Fprintln(w, "  add-sign           加签 --task-id <id> --user-ids <id1,id2> [--comment]")
	fmt.Fprintln(w, "  remove-sign        减签 --task-id <id> --user-ids <id1,id2> [--comment]")
}

// flowTaskListTask 查询待办/已办列表（接口契约一致，仅路径不同）
func flowTaskListTask(ctx context.Context, path string, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{"page": map[string]any{"page": page, "size": size}}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--task-name":
			reqBody["taskName"] = nextArg(remaining, &i)
		case "--process-name":
			reqBody["processName"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + path, Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowTaskGetPendingList 我的待办
func runFlowTaskGetPendingList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowTaskListTask(ctx, "/task/get-pending-list", args, stdout, stderr)
}

// runFlowTaskGetApprovedList 我的已办
func runFlowTaskGetApprovedList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowTaskListTask(ctx, "/task/get-approved-list", args, stdout, stderr)
}

// runFlowTaskGetDetail 审批详情
func runFlowTaskGetDetail(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		if remaining[i] == "--task-id" {
			reqBody["taskId"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/task/get-detail", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowTaskConsent 同意
func runFlowTaskConsent(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--task-id":
			reqBody["taskId"] = nextArg(remaining, &i)
		case "--comment":
			reqBody["comment"] = nextArg(remaining, &i)
		case "--variable":
			v, err := readJSONObjectArg(nextArg(remaining, &i))
			if err != nil {
				fmt.Fprintf(stderr, "Error: --variable %v\n", err)
				return 1
			}
			reqBody["variable"] = v
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/task/consent", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowTaskReject 驳回（带策略与目标节点）
func runFlowTaskReject(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--task-id":
			reqBody["taskId"] = nextArg(remaining, &i)
		case "--strategy":
			reqBody["strategy"] = nextArg(remaining, &i)
		case "--target-node-key":
			reqBody["targetNodeKey"] = nextArg(remaining, &i)
		case "--comment":
			reqBody["comment"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/task/reject", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowTaskTransfer 转办
func runFlowTaskTransfer(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--task-id":
			reqBody["taskId"] = nextArg(remaining, &i)
		case "--to-user-id":
			reqBody["toUserId"] = nextArg(remaining, &i)
		case "--comment":
			reqBody["comment"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/task/transfer", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowTaskSign 加签/减签（路径区分，契约一致：taskId + userIds + comment）
func runFlowTaskSign(ctx context.Context, path string, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--task-id":
			reqBody["taskId"] = nextArg(remaining, &i)
		case "--user-ids":
			reqBody["userIds"] = splitCSV(nextArg(remaining, &i))
		case "--comment":
			reqBody["comment"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + path, Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// ─── 流程实例 ────────────────────────────────────────────────────────────

// runFlowInstance 流程实例命令分发
func runFlowInstance(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printFlowInstanceHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-my-list":
		return runFlowInstanceGetMyList(ctx, args[1:], stdout, stderr)
	case "get-record":
		return runFlowInstanceGetRecord(ctx, args[1:], stdout, stderr)
	case "revoke":
		return runFlowInstanceRevoke(ctx, args[1:], stdout, stderr)
	case "get-monitor-list":
		return runFlowInstanceGetMonitorList(ctx, args[1:], stdout, stderr)
	case "terminate":
		return runFlowInstanceOperate(ctx, "/instance/terminate", args[1:], stdout, stderr)
	case "resume":
		return runFlowInstanceOperate(ctx, "/instance/resume", args[1:], stdout, stderr)
	case "destroy":
		return runFlowInstanceOperate(ctx, "/instance/destroy", args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printFlowInstanceHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown instance subcommand: %s\n", args[0])
		printFlowInstanceHelp(stderr)
		return 2
	}
}

// printFlowInstanceHelp 打印流程实例帮助
func printFlowInstanceHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur flow instance <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-my-list        我的申请（全员）[--title] [--page] [--size]")
	fmt.Fprintln(w, "  get-record         审批记录时间线（全员）--instance-id <id>")
	fmt.Fprintln(w, "  revoke             撤销我发起的实例（全员）--instance-id <id> [--comment]")
	fmt.Fprintln(w, "  get-monitor-list   实例监控列表（仅管理员）[--title] [--page] [--size]")
	fmt.Fprintln(w, "  terminate          终止实例（仅管理员）--instance-id <id> [--comment]")
	fmt.Fprintln(w, "  resume             恢复实例（仅管理员）--instance-id <id> [--comment]")
	fmt.Fprintln(w, "  destroy            作废实例（仅管理员）--instance-id <id> [--comment]")
}

// runFlowInstanceGetMyList 我的申请
func runFlowInstanceGetMyList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowInstanceList(ctx, "/instance/get-my-list", args, stdout, stderr)
}

// runFlowInstanceGetMonitorList 实例监控列表（仅管理员）
func runFlowInstanceGetMonitorList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowInstanceList(ctx, "/instance/get-monitor-list", args, stdout, stderr)
}

// flowInstanceList 实例列表查询（我的申请/监控共用，契约一致）
func flowInstanceList(ctx context.Context, path string, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{"page": map[string]any{"page": page, "size": size}}
	for i := 0; i < len(remaining); i++ {
		if remaining[i] == "--title" {
			reqBody["title"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + path, Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowInstanceGetRecord 审批记录时间线
func runFlowInstanceGetRecord(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		if remaining[i] == "--instance-id" {
			reqBody["instanceId"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/instance/get-record", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowInstanceRevoke 撤销实例
func runFlowInstanceRevoke(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowInstanceOperate(ctx, "/instance/revoke", args, stdout, stderr)
}

// runFlowInstanceOperate 实例管控操作（终止/恢复/作废，仅管理员）
func runFlowInstanceOperate(ctx context.Context, path string, args []string, stdout, stderr io.Writer) int {
	return flowInstanceOperate(ctx, path, args, stdout, stderr)
}

// flowInstanceOperate instanceId + comment 契约的实例操作
func flowInstanceOperate(ctx context.Context, path string, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--instance-id":
			reqBody["instanceId"] = nextArg(remaining, &i)
		case "--comment":
			reqBody["comment"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + path, Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// ─── 抄送（全员） ────────────────────────────────────────────────────────

// runFlowCc 我的抄送命令分发
func runFlowCc(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printFlowCcHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runFlowCcGetList(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printFlowCcHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown cc subcommand: %s\n", args[0])
		printFlowCcHelp(stderr)
		return 2
	}
}

// printFlowCcHelp 打印抄送帮助
func printFlowCcHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur flow cc get-list [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "我的抄送列表 [--title] [--page] [--size]")
}

// runFlowCcGetList 我的抄送列表
func runFlowCcGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{"page": map[string]any{"page": page, "size": size}}
	for i := 0; i < len(remaining); i++ {
		if remaining[i] == "--title" {
			reqBody["title"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/cc/get-list", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// ─── 工具 ────────────────────────────────────────────────────────────────

// readJSONObjectArg 解析流程变量参数：@开头读文件，否则视为内联 JSON。
// 仅校验 JSON 合法性（流程变量值可为任意 JSON），返回原文供后端按字符串解析。
func readJSONObjectArg(v string) (string, error) {
	var raw []byte
	if len(v) > 0 && v[0] == '@' {
		var err error
		raw, err = os.ReadFile(v[1:])
		if err != nil {
			return "", fmt.Errorf("读取文件失败: %w", err)
		}
	} else {
		raw = []byte(v)
	}
	var value any
	if err := json.Unmarshal(raw, &value); err != nil {
		return "", fmt.Errorf("不是合法的JSON: %w", err)
	}
	return string(raw), nil
}
