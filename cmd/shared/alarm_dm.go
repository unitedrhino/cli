package shared

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
)

// 新版告警模块（dm alarm v2）命令组，对应 /api/v1/things/alarm/* 接口：
// 规则（rule）/事件（event）/通知记录（notify-record）/通知模板（notify-template）/
// 触发条件模板（condition-template）。
// 旧版规则引擎告警（/api/v1/things/rule/alarm/*）的后端接口已删除，
// 原 info/record/scene 子命令已一并移除。
// 请求体字段以 backend/.swagger/things-api.json 为准。

// alarmDmHasHelpFlag 判断参数中是否带 help/--help/-h，用于子命令级帮助穿透
func alarmDmHasHelpFlag(args []string) bool {
	for _, a := range args {
		if a == "help" || a == "--help" || a == "-h" {
			return true
		}
	}
	return false
}

// runAlarmDmRule 执行新版告警规则命令（/api/v1/things/alarm/info/*）
func runAlarmDmRule(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || alarmDmHasHelpFlag(args) {
		printAlarmDmRuleHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runAlarmDmRuleGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runAlarmDmWithID(ctx, args[1:], "/api/v1/things/alarm/info/get-one", stdout, stderr)
	case "create":
		return runAlarmDmWithBody(ctx, args[1:], "/api/v1/things/alarm/info/create", stdout, stderr)
	case "update":
		return runAlarmDmWithBody(ctx, args[1:], "/api/v1/things/alarm/info/update", stdout, stderr)
	case "delete":
		return runAlarmDmWithID(ctx, args[1:], "/api/v1/things/alarm/info/delete", stdout, stderr)
	case "status-update":
		return runAlarmDmRuleStatusUpdate(ctx, args[1:], stdout, stderr)
	case "evaluate-trigger":
		return runAlarmDmRuleEvaluateTrigger(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAlarmDmRuleHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm rule subcommand: %s\n", args[0])
		printAlarmDmRuleHelp(stderr)
		return 2
	}
}

// printAlarmDmRuleHelp 打印新版告警规则命令帮助
func printAlarmDmRuleHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm rule <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "New alarm rule management (/api/v1/things/alarm/info/*)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list         Query alarm rule list (--name --product-id --status --level --start --end)")
	fmt.Fprintln(w, "  get-one          Query alarm rule detail (--id)")
	fmt.Fprintln(w, "  create           Create alarm rule (--body '<json>')")
	fmt.Fprintln(w, "  update           Update alarm rule (--body '<json>')")
	fmt.Fprintln(w, "  delete           Delete alarm rule (--id)")
	fmt.Fprintln(w, "  status-update    Enable/disable alarm rule (--id --status <0|1>)")
	fmt.Fprintln(w, "  evaluate-trigger Evaluate alarm rules now (--id <alarmID> [--id ...])")
}

// runAlarmDmRuleGetList 执行新版告警规则列表查询
func runAlarmDmRuleGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}
	headers := map[string]string{}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--name":
			if i+1 < len(remaining) {
				reqBody["keyword"] = remaining[i+1]
				i++
			}
		case "--product-id":
			if i+1 < len(remaining) {
				reqBody["productID"] = remaining[i+1]
				i++
			}
		case "--status":
			if i+1 < len(remaining) {
				if v, err := strconv.Atoi(remaining[i+1]); err == nil {
					reqBody["status"] = v
				}
				i++
			}
		case "--level":
			if i+1 < len(remaining) {
				reqBody["levels"] = splitCSV(remaining[i+1])
				i++
			}
		case "--start":
			if i+1 < len(remaining) {
				alarmDmSetTimeRange(reqBody, "createdTimeRange", remaining[i+1], false)
				i++
			}
		case "--end":
			if i+1 < len(remaining) {
				alarmDmSetTimeRange(reqBody, "createdTimeRange", remaining[i+1], true)
				i++
			}
		case "--project-id":
			if i+1 < len(remaining) {
				headers["project-id"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path:    "/api/v1/things/alarm/info/get-list",
		Body:    reqBody,
		Headers: headers,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmRuleStatusUpdate 执行告警规则启用/停用（status: 1=启用 0=停用，以 swagger 为准）
func runAlarmDmRuleStatusUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	id := ""
	status := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				id = args[i+1]
				i++
			}
		case "--status":
			if i+1 < len(args) {
				status = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}
	if id == "" || status == "" {
		fmt.Fprintln(stderr, "--id and --status are required")
		return 2
	}
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		fmt.Fprintf(stderr, "Error: --status must be an integer: %v\n", err)
		return 2
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/alarm/info/status-update",
		Body: map[string]any{"id": id, "status": statusInt},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmRuleEvaluateTrigger 执行告警规则手动评估（对应 alarmIDs 整型数组，--id 可重复）
func runAlarmDmRuleEvaluateTrigger(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	var alarmIDs []int64
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				for _, part := range strings.Split(args[i+1], ",") {
					if v, err := strconv.ParseInt(strings.TrimSpace(part), 10, 64); err == nil {
						alarmIDs = append(alarmIDs, v)
					}
				}
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}
	if len(alarmIDs) == 0 {
		fmt.Fprintln(stderr, "--id is required (one or more alarm rule IDs)")
		return 2
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/alarm/info/evaluate-trigger",
		Body: map[string]any{"alarmIDs": alarmIDs},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmEvent 执行新版告警事件命令（/api/v1/things/alarm/event/*）
func runAlarmDmEvent(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || alarmDmHasHelpFlag(args) {
		printAlarmDmEventHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runAlarmDmEventGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runAlarmDmWithID(ctx, args[1:], "/api/v1/things/alarm/event/get-one", stdout, stderr)
	case "deal":
		return runAlarmDmEventDeal(ctx, args[1:], stdout, stderr)
	case "false-alarm":
		return runAlarmDmEventFalseAlarm(ctx, args[1:], stdout, stderr)
	case "stat":
		return runAlarmDmEventStat(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAlarmDmEventHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm event subcommand: %s\n", args[0])
		printAlarmDmEventHelp(stderr)
		return 2
	}
}

// printAlarmDmEventHelp 打印新版告警事件命令帮助
func printAlarmDmEventHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm event <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "New alarm event management (/api/v1/things/alarm/event/*)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list      Query alarm event list (--rule-id --product-id --device-name --status --level --keyword --is-recovered --is-false-alarm --start --end)")
	fmt.Fprintln(w, "  get-one       Query alarm event detail (--id)")
	fmt.Fprintln(w, "  deal          Deal with alarm event (--id --action <acked|recovered...> [--remark])")
	fmt.Fprintln(w, "  false-alarm   Mark alarm event as false alarm (--id --reason [--reason-type])")
	fmt.Fprintln(w, "  stat          Alarm event statistics (--group-by --status --level --start --end)")
}

// runAlarmDmEventGetList 执行新版告警事件列表查询
func runAlarmDmEventGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}
	headers := map[string]string{}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--rule-id":
			if i+1 < len(remaining) {
				reqBody["alarmID"] = remaining[i+1]
				i++
			}
		case "--product-id":
			if i+1 < len(remaining) {
				reqBody["productID"] = remaining[i+1]
				i++
			}
		case "--device-name":
			if i+1 < len(remaining) {
				reqBody["deviceName"] = remaining[i+1]
				i++
			}
		case "--keyword":
			if i+1 < len(remaining) {
				reqBody["keyword"] = remaining[i+1]
				i++
			}
		case "--status":
			if i+1 < len(remaining) {
				reqBody["status"] = splitCSV(remaining[i+1])
				i++
			}
		case "--level":
			if i+1 < len(remaining) {
				reqBody["level"] = splitCSV(remaining[i+1])
				i++
			}
		case "--is-recovered":
			reqBody["isRecovered"] = true
		case "--is-false-alarm":
			reqBody["isFalseAlarm"] = true
		case "--start":
			if i+1 < len(remaining) {
				alarmDmSetTimeRange(reqBody, "triggerTimeRange", remaining[i+1], false)
				i++
			}
		case "--end":
			if i+1 < len(remaining) {
				alarmDmSetTimeRange(reqBody, "triggerTimeRange", remaining[i+1], true)
				i++
			}
		case "--project-id":
			if i+1 < len(remaining) {
				headers["project-id"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path:    "/api/v1/things/alarm/event/get-list",
		Body:    reqBody,
		Headers: headers,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmEventDeal 执行告警事件处理（action 取值以 swagger/后端为准，如 acked）
func runAlarmDmEventDeal(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	id := ""
	action := ""
	remark := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				id = args[i+1]
				i++
			}
		case "--action":
			if i+1 < len(args) {
				action = args[i+1]
				i++
			}
		case "--remark":
			if i+1 < len(args) {
				remark = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}
	if id == "" || action == "" {
		fmt.Fprintln(stderr, "--id and --action are required")
		return 2
	}
	reqBody := map[string]any{"eventID": id, "action": action}
	if remark != "" {
		reqBody["remark"] = remark
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/alarm/event/deal",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmEventFalseAlarm 执行告警事件误报标记
func runAlarmDmEventFalseAlarm(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	id := ""
	reason := ""
	reasonType := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				id = args[i+1]
				i++
			}
		case "--reason":
			if i+1 < len(args) {
				reason = args[i+1]
				i++
			}
		case "--reason-type":
			if i+1 < len(args) {
				reasonType = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}
	if id == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}
	reqBody := map[string]any{"eventID": id}
	if reason != "" {
		reqBody["reason"] = reason
	}
	if reasonType != "" {
		reqBody["reasonType"] = reasonType
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/alarm/event/false-alarm",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmEventStat 执行告警事件统计
func runAlarmDmEventStat(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	reqBody := map[string]any{}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--group-by":
			if i+1 < len(args) {
				reqBody["groupBy"] = args[i+1]
				i++
			}
		case "--status":
			if i+1 < len(args) {
				reqBody["status"] = splitCSV(args[i+1])
				i++
			}
		case "--level":
			if i+1 < len(args) {
				reqBody["level"] = splitCSV(args[i+1])
				i++
			}
		case "--start":
			if i+1 < len(args) {
				alarmDmSetTimeRange(reqBody, "triggerTimeRange", args[i+1], false)
				i++
			}
		case "--end":
			if i+1 < len(args) {
				alarmDmSetTimeRange(reqBody, "triggerTimeRange", args[i+1], true)
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/alarm/event/stat",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmNotifyRecord 执行新版告警通知记录命令（/api/v1/things/alarm/notify-record/*）
func runAlarmDmNotifyRecord(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || alarmDmHasHelpFlag(args) {
		printAlarmDmNotifyRecordHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runAlarmDmNotifyRecordGetList(ctx, args[1:], stdout, stderr)
	case "resend":
		return runAlarmDmWithID(ctx, args[1:], "/api/v1/things/alarm/notify-record/resend", stdout, stderr)
	case "help", "--help", "-h":
		printAlarmDmNotifyRecordHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm notify-record subcommand: %s\n", args[0])
		printAlarmDmNotifyRecordHelp(stderr)
		return 2
	}
}

// printAlarmDmNotifyRecordHelp 打印告警通知记录命令帮助
func printAlarmDmNotifyRecordHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm notify-record <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "New alarm notify record management (/api/v1/things/alarm/notify-record/*)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query notify record list (--event-id --alarm-id --alarm-name --channel --status --target-name --timing --start --end)")
	fmt.Fprintln(w, "  resend     Resend a notify record (--id)")
}

// runAlarmDmNotifyRecordGetList 执行告警通知记录列表查询
func runAlarmDmNotifyRecordGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}
	headers := map[string]string{}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--event-id":
			if i+1 < len(remaining) {
				reqBody["eventID"] = remaining[i+1]
				i++
			}
		case "--alarm-id":
			if i+1 < len(remaining) {
				reqBody["alarmID"] = remaining[i+1]
				i++
			}
		case "--alarm-name":
			if i+1 < len(remaining) {
				reqBody["alarmName"] = remaining[i+1]
				i++
			}
		case "--channel":
			if i+1 < len(remaining) {
				reqBody["channel"] = splitCSV(remaining[i+1])
				i++
			}
		case "--status":
			if i+1 < len(remaining) {
				reqBody["status"] = splitCSV(remaining[i+1])
				i++
			}
		case "--target-name":
			if i+1 < len(remaining) {
				reqBody["targetName"] = remaining[i+1]
				i++
			}
		case "--timing":
			if i+1 < len(remaining) {
				reqBody["timing"] = remaining[i+1]
				i++
			}
		case "--start":
			if i+1 < len(remaining) {
				alarmDmSetTimeRange(reqBody, "createdTimeRange", remaining[i+1], false)
				i++
			}
		case "--end":
			if i+1 < len(remaining) {
				alarmDmSetTimeRange(reqBody, "createdTimeRange", remaining[i+1], true)
				i++
			}
		case "--project-id":
			if i+1 < len(remaining) {
				headers["project-id"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path:    "/api/v1/things/alarm/notify-record/get-list",
		Body:    reqBody,
		Headers: headers,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmNotifyTemplate 执行新版告警通知模板命令（/api/v1/things/alarm/notify-template/*）
func runAlarmDmNotifyTemplate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || alarmDmHasHelpFlag(args) {
		printAlarmDmNotifyTemplateHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runAlarmDmNotifyTemplateGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runAlarmDmWithID(ctx, args[1:], "/api/v1/things/alarm/notify-template/get-one", stdout, stderr)
	case "create":
		return runAlarmDmWithBody(ctx, args[1:], "/api/v1/things/alarm/notify-template/create", stdout, stderr)
	case "update":
		return runAlarmDmWithBody(ctx, args[1:], "/api/v1/things/alarm/notify-template/update", stdout, stderr)
	case "delete":
		return runAlarmDmWithID(ctx, args[1:], "/api/v1/things/alarm/notify-template/delete", stdout, stderr)
	case "test-send":
		return runAlarmDmNotifyTemplateTestSend(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAlarmDmNotifyTemplateHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm notify-template subcommand: %s\n", args[0])
		printAlarmDmNotifyTemplateHelp(stderr)
		return 2
	}
}

// printAlarmDmNotifyTemplateHelp 打印告警通知模板命令帮助
func printAlarmDmNotifyTemplateHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm notify-template <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "New alarm notify template management (/api/v1/things/alarm/notify-template/*)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query notify template list (--channel --keyword --status)")
	fmt.Fprintln(w, "  get-one    Query notify template detail (--id)")
	fmt.Fprintln(w, "  create     Create notify template (--body '<json>')")
	fmt.Fprintln(w, "  update     Update notify template (--body '<json>')")
	fmt.Fprintln(w, "  delete     Delete notify template (--id)")
	fmt.Fprintln(w, "  test-send  Send a test notification (--id [--user-ids --role-ids --group-ids --timing] or --body)")
}

// runAlarmDmNotifyTemplateGetList 执行告警通知模板列表查询
func runAlarmDmNotifyTemplateGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--channel":
			if i+1 < len(remaining) {
				reqBody["channel"] = remaining[i+1]
				i++
			}
		case "--keyword":
			if i+1 < len(remaining) {
				reqBody["keyword"] = remaining[i+1]
				i++
			}
		case "--status":
			if i+1 < len(remaining) {
				if v, err := strconv.Atoi(remaining[i+1]); err == nil {
					reqBody["status"] = v
				}
				i++
			}
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/alarm/notify-template/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmNotifyTemplateTestSend 执行通知模板测试发送
// targets 结构为 {userIDs,roleIDs,groupIDs}，各 ID 支持逗号分隔或重复传参
func runAlarmDmNotifyTemplateTestSend(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	bodyJSON := ""
	id := ""
	var userIDs, roleIDs, groupIDs []string
	timing := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				id = args[i+1]
				i++
			}
		case "--user-ids":
			if i+1 < len(args) {
				userIDs = append(userIDs, splitCSV(args[i+1])...)
				i++
			}
		case "--role-ids":
			if i+1 < len(args) {
				roleIDs = append(roleIDs, splitCSV(args[i+1])...)
				i++
			}
		case "--group-ids":
			if i+1 < len(args) {
				groupIDs = append(groupIDs, splitCSV(args[i+1])...)
				i++
			}
		case "--timing":
			if i+1 < len(args) {
				timing = args[i+1]
				i++
			}
		case "--body":
			if i+1 < len(args) {
				bodyJSON = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}
	var reqBody map[string]any
	if bodyJSON != "" {
		var err error
		reqBody, err = parseBodyArg(bodyJSON)
		if err != nil {
			fmt.Fprintf(stderr, "Error: %v\n", err)
			return 2
		}
	} else {
		if id == "" {
			fmt.Fprintln(stderr, "--id is required (or use --body)")
			return 2
		}
		reqBody = map[string]any{"id": id}
		targets := map[string]any{}
		if len(userIDs) > 0 {
			targets["userIDs"] = userIDs
		}
		if len(roleIDs) > 0 {
			targets["roleIDs"] = roleIDs
		}
		if len(groupIDs) > 0 {
			targets["groupIDs"] = groupIDs
		}
		if len(targets) > 0 {
			reqBody["targets"] = targets
		}
		if timing != "" {
			reqBody["timing"] = timing
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/alarm/notify-template/test-send",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmConditionTemplate 执行新版触发条件模板命令（/api/v1/things/alarm/condition-template/*）
func runAlarmDmConditionTemplate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || alarmDmHasHelpFlag(args) {
		printAlarmDmConditionTemplateHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runAlarmDmConditionTemplateGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runAlarmDmWithID(ctx, args[1:], "/api/v1/things/alarm/condition-template/get-one", stdout, stderr)
	case "create":
		return runAlarmDmWithBody(ctx, args[1:], "/api/v1/things/alarm/condition-template/create", stdout, stderr)
	case "update":
		return runAlarmDmWithBody(ctx, args[1:], "/api/v1/things/alarm/condition-template/update", stdout, stderr)
	case "delete":
		return runAlarmDmWithID(ctx, args[1:], "/api/v1/things/alarm/condition-template/delete", stdout, stderr)
	case "help", "--help", "-h":
		printAlarmDmConditionTemplateHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm condition-template subcommand: %s\n", args[0])
		printAlarmDmConditionTemplateHelp(stderr)
		return 2
	}
}

// printAlarmDmConditionTemplateHelp 打印触发条件模板命令帮助
func printAlarmDmConditionTemplateHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm condition-template <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "New alarm condition template management (/api/v1/things/alarm/condition-template/*)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query condition template list (--keyword --product-id --product-category-id --status)")
	fmt.Fprintln(w, "  get-one    Query condition template detail (--id)")
	fmt.Fprintln(w, "  create     Create condition template (--body '<json>')")
	fmt.Fprintln(w, "  update     Update condition template (--body '<json>')")
	fmt.Fprintln(w, "  delete     Delete condition template (--id)")
}

// runAlarmDmConditionTemplateGetList 执行触发条件模板列表查询
func runAlarmDmConditionTemplateGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--keyword":
			if i+1 < len(remaining) {
				reqBody["keyword"] = remaining[i+1]
				i++
			}
		case "--product-id":
			if i+1 < len(remaining) {
				reqBody["productID"] = remaining[i+1]
				i++
			}
		case "--product-category-id":
			if i+1 < len(remaining) {
				reqBody["productCategoryID"] = remaining[i+1]
				i++
			}
		case "--status":
			if i+1 < len(remaining) {
				if v, err := strconv.Atoi(remaining[i+1]); err == nil {
					reqBody["status"] = v
				}
				i++
			}
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/alarm/condition-template/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmWithID 通用 get-one/delete/resend 单 ID 子命令（body 只含 id）
func runAlarmDmWithID(ctx context.Context, args []string, path string, stdout, stderr io.Writer) int {
	jsonOutput := false
	id := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				id = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}
	if id == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: path,
		Body: map[string]any{"id": id},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmDmWithBody 通用 create/update 子命令（--body 传完整 JSON）
func runAlarmDmWithBody(ctx context.Context, args []string, path string, stdout, stderr io.Writer) int {
	jsonOutput := false
	bodyJSON := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--body":
			if i+1 < len(args) {
				bodyJSON = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}
	if bodyJSON == "" {
		fmt.Fprintln(stderr, "--body is required")
		return 2
	}
	reqBody, err := parseBodyArg(bodyJSON)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: path,
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// splitCSV 将逗号分隔的字符串切分为字符串数组（过滤空白项）
func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if trimmed := strings.TrimSpace(p); trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

// alarmDmSetTimeRange 组装 {start,end} 毫秒时间范围字段（start/end 至少一个存在才写入）
// isEnd=true 表示本次设置的是 end 值
func alarmDmSetTimeRange(reqBody map[string]any, field, value string, isEnd bool) {
	tr, _ := reqBody[field].(map[string]any)
	if tr == nil {
		tr = map[string]any{}
	}
	if isEnd {
		tr["end"] = value
	} else {
		tr["start"] = value
	}
	reqBody[field] = tr
}
