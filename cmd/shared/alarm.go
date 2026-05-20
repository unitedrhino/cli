package shared

import (
	"context"
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runAlarm 执行告警管理命令
func runAlarm(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAlarmHelp(stdout)
		return 0
	}

	switch args[0] {
	case "info":
		return runAlarmInfo(ctx, args[1:], stdout, stderr)
	case "record":
		return runAlarmRecord(ctx, args[1:], stdout, stderr)
	case "scene":
		return runAlarmScene(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAlarmHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm subcommand: %s\n", args[0])
		printAlarmHelp(stderr)
		return 2
	}
}

// printAlarmHelp 打印告警管理帮助信息
func printAlarmHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Alarm management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  info       Alarm info management (get-list, get-one, create, update, delete)")
	fmt.Fprintln(w, "  record     Alarm record management (get-list, deal)")
	fmt.Fprintln(w, "  scene      Alarm scene management (get-list, batch-create, delete)")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runAlarmInfo 执行告警信息管理命令
func runAlarmInfo(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAlarmInfoHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runAlarmInfoGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runAlarmInfoGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runAlarmInfoCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runAlarmInfoUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runAlarmInfoDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAlarmInfoHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm info subcommand: %s\n", args[0])
		printAlarmInfoHelp(stderr)
		return 2
	}
}

// printAlarmInfoHelp 打印告警信息帮助信息
func printAlarmInfoHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm info <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Alarm info management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query alarm list")
	fmt.Fprintln(w, "  get-one    Query alarm detail")
	fmt.Fprintln(w, "  create     Create alarm")
	fmt.Fprintln(w, "  update     Update alarm")
	fmt.Fprintln(w, "  delete     Delete alarm")
}

// runAlarmInfoGetList 执行查询告警列表命令
func runAlarmInfoGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--product-id", "-p":
			if i+1 < len(remaining) {
				reqBody["productID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/rule/alarm/info/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmInfoGetOne 执行查询告警详情命令
func runAlarmInfoGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	alarmID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				alarmID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if alarmID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/rule/alarm/info/get-one",
		Body: map[string]any{"id": alarmID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmInfoCreate 执行创建告警命令
func runAlarmInfoCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/rule/alarm/info/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmInfoUpdate 执行更新告警命令
func runAlarmInfoUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/rule/alarm/info/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmInfoDelete 执行删除告警命令
func runAlarmInfoDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	alarmID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				alarmID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if alarmID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/rule/alarm/info/delete",
		Body: map[string]any{"id": alarmID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmRecord 执行告警记录管理命令
func runAlarmRecord(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAlarmRecordHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runAlarmRecordGetList(ctx, args[1:], stdout, stderr)
	case "deal":
		return runAlarmRecordDeal(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAlarmRecordHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm record subcommand: %s\n", args[0])
		printAlarmRecordHelp(stderr)
		return 2
	}
}

// printAlarmRecordHelp 打印告警记录帮助信息
func printAlarmRecordHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm record <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Alarm record management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query alarm record list")
	fmt.Fprintln(w, "  deal       Deal with alarm record")
}

// runAlarmRecordGetList 执行查询告警记录列表命令
func runAlarmRecordGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--alarm-id":
			if i+1 < len(remaining) {
				reqBody["alarmID"] = remaining[i+1]
				i++
			}
		case "--product-id", "-p":
			if i+1 < len(remaining) {
				reqBody["productID"] = remaining[i+1]
				i++
			}
		case "--level":
			if i+1 < len(remaining) {
				reqBody["level"] = remaining[i+1]
				i++
			}
		case "--deal-state":
			if i+1 < len(remaining) {
				reqBody["dealState"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/rule/alarm/record/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmRecordDeal 执行处理告警记录命令
func runAlarmRecordDeal(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/rule/alarm/record/deal",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmScene 执行告警场景管理命令
func runAlarmScene(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAlarmSceneHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runAlarmSceneGetList(ctx, args[1:], stdout, stderr)
	case "batch-create":
		return runAlarmSceneBatchCreate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runAlarmSceneDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAlarmSceneHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm scene subcommand: %s\n", args[0])
		printAlarmSceneHelp(stderr)
		return 2
	}
}

// printAlarmSceneHelp 打印告警场景帮助信息
func printAlarmSceneHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm scene <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Alarm scene management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list      Query alarm scene list")
	fmt.Fprintln(w, "  batch-create  Batch create alarm scenes")
	fmt.Fprintln(w, "  delete        Delete alarm scene")
}

// runAlarmSceneGetList 执行查询告警场景列表命令
func runAlarmSceneGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--alarm-id":
			if i+1 < len(remaining) {
				reqBody["alarmID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/rule/alarm/scene/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmSceneBatchCreate 执行批量创建告警场景命令
func runAlarmSceneBatchCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/rule/alarm/scene/batch-create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAlarmSceneDelete 执行删除告警场景命令
func runAlarmSceneDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	sceneID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				sceneID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if sceneID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/rule/alarm/scene/delete",
		Body: map[string]any{"id": sceneID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
