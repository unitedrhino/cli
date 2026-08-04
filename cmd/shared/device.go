package shared

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/config"
)

// runDevice 执行设备调试命令
func runDevice(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeviceHelp(stdout)
		return 0
	}

	switch args[0] {
	case "log":
		return runDeviceLog(ctx, args[1:], stdout, stderr)
	case "control":
		return runDeviceControl(ctx, args[1:], stdout, stderr)
	case "action":
		return runDeviceAction(ctx, args[1:], stdout, stderr)
	case "mock":
		return runDeviceMock(ctx, args[1:], stdout, stderr)
	case "report":
		return runDeviceReport(ctx, args[1:], stdout, stderr)
	case "info":
		return runDeviceInfo(ctx, args[1:], stdout, stderr)
	case "gateway":
		return runDeviceGateway(ctx, args[1:], stdout, stderr)
	case "group":
		return runDeviceGroup(ctx, args[1:], stdout, stderr)
	case "profile":
		return runDeviceProfile(ctx, args[1:], stdout, stderr)
	case "upload":
		return runDeviceUpload(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeviceHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown device subcommand: %s\n", args[0])
		printDeviceHelp(stderr)
		return 2
	}
}

// runDeviceLog 执行设备日志查询命令
func runDeviceLog(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeviceLogHelp(stdout)
		return 0
	}

	switch args[0] {
	case "property":
		return runDeviceLogProperty(ctx, args[1:], stdout, stderr)
	case "event":
		return runDeviceLogEvent(ctx, args[1:], stdout, stderr)
	case "send":
		return runDeviceLogSend(ctx, args[1:], stdout, stderr)
	case "status":
		return runDeviceLogStatus(ctx, args[1:], stdout, stderr)
	case "hub":
		return runDeviceLogHub(ctx, args[1:], stdout, stderr)
	case "abnormal":
		return runDeviceLogAbnormal(ctx, args[1:], stdout, stderr)
	case "sdk":
		return runDeviceLogSDK(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeviceLogHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown device log subcommand: %s\n", args[0])
		printDeviceLogHelp(stderr)
		return 2
	}
}

// printDeviceLogHelp 打印设备日志查询帮助信息
func printDeviceLogHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device log <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device logs")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  property   Query property logs (latest value, history)")
	fmt.Fprintln(w, "  event      Query event logs")
	fmt.Fprintln(w, "  send       Query command send logs")
	fmt.Fprintln(w, "  status     Query online/offline status logs")
	fmt.Fprintln(w, "  hub        Query diagnostic logs (MQTT communication)")
	fmt.Fprintln(w, "  abnormal   Query abnormal logs")
	fmt.Fprintln(w, "  sdk        Query SDK logs")
	fmt.Fprintln(w, "  help       Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Common options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --time-start string        Start timestamp (milliseconds)")
	fmt.Fprintln(w, "  --time-end string          End timestamp (milliseconds)")
	fmt.Fprintln(w, "  --page int                 Page number (default: 1)")
	fmt.Fprintln(w, "  --size int                 Page size (default: 20)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device latest property values")
	fmt.Fprintln(w, "  ur device log property -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query device temperature history")
	fmt.Fprintln(w, "  ur device log property -p p_smartswitch_001 -d switch-001 --data-id Temperature --arg-func avg")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query device event logs")
	fmt.Fprintln(w, "  ur device log event -p p_smartswitch_001 -d switch-001 --types info,alert")
}

// parseLogParams 解析日志查询通用参数
func parseLogParams(args []string) (productID, deviceName string, timeStart, timeEnd string, page, size int, jsonOutput bool, remaining []string, err error) {
	remaining = make([]string, 0, len(args))
	page = 1
	size = 20
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--product-id", "-p":
			if i+1 >= len(args) {
				return "", "", "", "", 0, 0, false, nil, fmt.Errorf("--product-id requires value")
			}
			productID = args[i+1]
			i++
		case "--device-name", "-d":
			if i+1 >= len(args) {
				return "", "", "", "", 0, 0, false, nil, fmt.Errorf("--device-name requires value")
			}
			deviceName = args[i+1]
			i++
		case "--time-start":
			if i+1 >= len(args) {
				return "", "", "", "", 0, 0, false, nil, fmt.Errorf("--time-start requires value")
			}
			timeStart = args[i+1]
			i++
		case "--time-end":
			if i+1 >= len(args) {
				return "", "", "", "", 0, 0, false, nil, fmt.Errorf("--time-end requires value")
			}
			timeEnd = args[i+1]
			i++
		case "--page":
			if i+1 >= len(args) {
				return "", "", "", "", 0, 0, false, nil, fmt.Errorf("--page requires value")
			}
			fmt.Sscanf(args[i+1], "%d", &page)
			i++
		case "--size":
			if i+1 >= len(args) {
				return "", "", "", "", 0, 0, false, nil, fmt.Errorf("--size requires value")
			}
			fmt.Sscanf(args[i+1], "%d", &size)
			i++
		case "--json", "-j":
			jsonOutput = true
		default:
			remaining = append(remaining, args[i])
		}
	}
	if productID == "" {
		return "", "", "", "", 0, 0, false, nil, fmt.Errorf("--product-id is required")
	}
	if deviceName == "" {
		return "", "", "", "", 0, 0, false, nil, fmt.Errorf("--device-name is required")
	}
	return productID, deviceName, timeStart, timeEnd, page, size, jsonOutput, remaining, nil
}

// runDeviceLogProperty 执行属性日志查询命令
func runDeviceLogProperty(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, timeStart, timeEnd, page, size, jsonOutput, remaining, err := parseLogParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceLogPropertyHelp(stderr)
		return 2
	}

	// 解析属性特定参数
	dataID := ""
	dataIDs := ""
	ignoreEmpty := false
	argFunc := ""
	interval := 0
	intervalUnit := ""
	order := 0

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--data-id":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data-id requires value")
				return 2
			}
			dataID = remaining[i+1]
			i++
		case "--data-ids":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data-ids requires value")
				return 2
			}
			dataIDs = remaining[i+1]
			i++
		case "--ignore-empty":
			ignoreEmpty = true
		case "--arg-func":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--arg-func requires value")
				return 2
			}
			argFunc = remaining[i+1]
			i++
		case "--interval":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--interval requires value")
				return 2
			}
			fmt.Sscanf(remaining[i+1], "%d", &interval)
			i++
		case "--interval-unit":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--interval-unit requires value")
				return 2
			}
			intervalUnit = remaining[i+1]
			i++
		case "--order":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--order requires value")
				return 2
			}
			fmt.Sscanf(remaining[i+1], "%d", &order)
			i++
		default:
			fmt.Fprintf(stderr, "unknown property option: %s\n", remaining[i])
			printDeviceLogPropertyHelp(stderr)
			return 2
		}
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"page":       map[string]any{"page": page, "size": size},
	}

	if timeStart != "" {
		reqBody["timeStart"] = timeStart
	}
	if timeEnd != "" {
		reqBody["timeEnd"] = timeEnd
	}

	// 根据参数选择 API
	apiPath := ""
	if dataID != "" || argFunc != "" {
		// 查询属性历史记录
		apiPath = "/api/v1/things/device/msg/property-log/get-list"
		if dataID != "" {
			reqBody["dataID"] = dataID
		}
		if argFunc != "" {
			reqBody["argFunc"] = argFunc
		}
		if interval > 0 {
			reqBody["interval"] = interval
		}
		if intervalUnit != "" {
			reqBody["intervalUnit"] = intervalUnit
		}
		if order > 0 {
			reqBody["order"] = order
		}
	} else {
		// 查询最新属性值
		apiPath = "/api/v1/things/device/msg/property-latest/get-list"
		if dataIDs != "" {
			reqBody["dataIDs"] = strings.Split(dataIDs, ",")
		}
		if ignoreEmpty {
			reqBody["ignoreEmpty"] = true
		}
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: apiPath,
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		formatDeviceResult(resp, stdout, stderr)
	}

	return 0
}

// runDeviceLogEvent 执行事件日志查询命令
func runDeviceLogEvent(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, timeStart, timeEnd, page, size, jsonOutput, remaining, err := parseLogParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceLogEventHelp(stderr)
		return 2
	}

	// 解析事件特定参数
	types := ""
	dataID := ""

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--types":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--types requires value")
				return 2
			}
			types = remaining[i+1]
			i++
		case "--data-id":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data-id requires value")
				return 2
			}
			dataID = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown event option: %s\n", remaining[i])
			printDeviceLogEventHelp(stderr)
			return 2
		}
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"page":       map[string]any{"page": page, "size": size},
	}

	if timeStart != "" {
		reqBody["timeStart"] = timeStart
	}
	if timeEnd != "" {
		reqBody["timeEnd"] = timeEnd
	}
	if types != "" {
		reqBody["types"] = strings.Split(types, ",")
	}
	if dataID != "" {
		reqBody["dataID"] = dataID
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/msg/event-log/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		formatDeviceResult(resp, stdout, stderr)
	}

	return 0
}

// runDeviceLogSend 执行命令日志查询命令
func runDeviceLogSend(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, timeStart, timeEnd, page, size, jsonOutput, remaining, err := parseLogParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceLogSendHelp(stderr)
		return 2
	}

	// 解析命令日志特定参数
	actions := ""
	resultCode := 0
	dataID := ""
	withUser := false

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--actions":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--actions requires value")
				return 2
			}
			actions = remaining[i+1]
			i++
		case "--result-code":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--result-code requires value")
				return 2
			}
			fmt.Sscanf(remaining[i+1], "%d", &resultCode)
			i++
		case "--data-id":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data-id requires value")
				return 2
			}
			dataID = remaining[i+1]
			i++
		case "--with-user":
			withUser = true
		default:
			fmt.Fprintf(stderr, "unknown send option: %s\n", remaining[i])
			printDeviceLogSendHelp(stderr)
			return 2
		}
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"page":       map[string]any{"page": page, "size": size},
	}

	if timeStart != "" {
		reqBody["timeStart"] = timeStart
	}
	if timeEnd != "" {
		reqBody["timeEnd"] = timeEnd
	}
	if actions != "" {
		reqBody["actions"] = strings.Split(actions, ",")
	}
	if resultCode > 0 {
		reqBody["resultCode"] = resultCode
	}
	if dataID != "" {
		reqBody["dataID"] = dataID
	}
	if withUser {
		reqBody["withUser"] = true
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/msg/send-log/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		formatDeviceResult(resp, stdout, stderr)
	}

	return 0
}

// runDeviceLogStatus 执行上下线日志查询命令
func runDeviceLogStatus(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, timeStart, timeEnd, page, size, jsonOutput, remaining, err := parseLogParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceLogStatusHelp(stderr)
		return 2
	}

	// 解析上下线日志特定参数
	status := 0

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--status":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--status requires value")
				return 2
			}
			fmt.Sscanf(remaining[i+1], "%d", &status)
			i++
		default:
			fmt.Fprintf(stderr, "unknown status option: %s\n", remaining[i])
			printDeviceLogStatusHelp(stderr)
			return 2
		}
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"page":       map[string]any{"page": page, "size": size},
	}

	if timeStart != "" {
		reqBody["timeStart"] = timeStart
	}
	if timeEnd != "" {
		reqBody["timeEnd"] = timeEnd
	}
	if status > 0 {
		reqBody["status"] = status
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/msg/status-log/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		formatDeviceResult(resp, stdout, stderr)
	}

	return 0
}

// runDeviceLogHub 执行诊断日志查询命令
func runDeviceLogHub(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, timeStart, timeEnd, page, size, jsonOutput, remaining, err := parseLogParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceLogHubHelp(stderr)
		return 2
	}

	// 解析诊断日志特定参数
	actions := ""
	topics := ""
	content := ""
	requestID := ""

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--actions":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--actions requires value")
				return 2
			}
			actions = remaining[i+1]
			i++
		case "--topics":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--topics requires value")
				return 2
			}
			topics = remaining[i+1]
			i++
		case "--content":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--content requires value")
				return 2
			}
			content = remaining[i+1]
			i++
		case "--request-id":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--request-id requires value")
				return 2
			}
			requestID = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown hub option: %s\n", remaining[i])
			printDeviceLogHubHelp(stderr)
			return 2
		}
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"page":       map[string]any{"page": page, "size": size},
	}

	if timeStart != "" {
		reqBody["timeStart"] = timeStart
	}
	if timeEnd != "" {
		reqBody["timeEnd"] = timeEnd
	}
	if actions != "" {
		reqBody["actions"] = strings.Split(actions, ",")
	}
	if topics != "" {
		reqBody["topics"] = strings.Split(topics, ",")
	}
	if content != "" {
		reqBody["content"] = content
	}
	if requestID != "" {
		reqBody["requestID"] = requestID
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/msg/hub-log/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		formatDeviceResult(resp, stdout, stderr)
	}

	return 0
}

// runDeviceLogAbnormal 执行异常日志查询命令
func runDeviceLogAbnormal(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, timeStart, timeEnd, page, size, jsonOutput, remaining, err := parseLogParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceLogAbnormalHelp(stderr)
		return 2
	}

	// 解析异常日志特定参数
	action := 0
	abnormalType := ""

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--action":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--action requires value")
				return 2
			}
			fmt.Sscanf(remaining[i+1], "%d", &action)
			i++
		case "--type":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--type requires value")
				return 2
			}
			abnormalType = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown abnormal option: %s\n", remaining[i])
			printDeviceLogAbnormalHelp(stderr)
			return 2
		}
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"page":       map[string]any{"page": page, "size": size},
	}

	if timeStart != "" {
		reqBody["timeStart"] = timeStart
	}
	if timeEnd != "" {
		reqBody["timeEnd"] = timeEnd
	}
	if action > 0 {
		reqBody["action"] = action
	}
	if abnormalType != "" {
		reqBody["type"] = abnormalType
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/msg/abnormal-log/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		formatDeviceResult(resp, stdout, stderr)
	}

	return 0
}

// runDeviceLogSDK 执行 SDK 日志查询命令
func runDeviceLogSDK(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, timeStart, timeEnd, page, size, jsonOutput, remaining, err := parseLogParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceLogSDKHelp(stderr)
		return 2
	}

	// 解析 SDK 日志特定参数
	logLevel := 0

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--log-level":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--log-level requires value")
				return 2
			}
			fmt.Sscanf(remaining[i+1], "%d", &logLevel)
			i++
		default:
			fmt.Fprintf(stderr, "unknown sdk option: %s\n", remaining[i])
			printDeviceLogSDKHelp(stderr)
			return 2
		}
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"page":       map[string]any{"page": page, "size": size},
	}

	if timeStart != "" {
		reqBody["timeStart"] = timeStart
	}
	if timeEnd != "" {
		reqBody["timeEnd"] = timeEnd
	}
	if logLevel > 0 {
		reqBody["logLevel"] = logLevel
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/msg/sdk-log/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		formatDeviceResult(resp, stdout, stderr)
	}

	return 0
}

// runDeviceControl 执行设备属性控制命令
func runDeviceControl(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceControlHelp(stderr)
		return 2
	}

	// 解析属性控制特定参数
	data := ""

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--data":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data requires value")
				return 2
			}
			data = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown control option: %s\n", remaining[i])
			printDeviceControlHelp(stderr)
			return 2
		}
	}

	if data == "" {
		fmt.Fprintln(stderr, "--data is required")
		printDeviceControlHelp(stderr)
		return 2
	}

	// 解析 data JSON
	var dataMap map[string]any
	if err := json.Unmarshal([]byte(data), &dataMap); err != nil {
		fmt.Fprintf(stderr, "Error: --data must be JSON object: %v\n", err)
		return 2
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"data":       dataMap,
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/interact/property-control-send",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		if resp.Code != 200 {
			fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
			return 1
		}
		fmt.Fprintln(stdout, "Property control command sent successfully")
	}

	return 0
}

// runDeviceAction 执行设备行为调用命令
func runDeviceAction(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeviceActionHelp(stdout)
		return 0
	}

	switch args[0] {
	case "send":
		return runDeviceActionSend(ctx, args[1:], stdout, stderr)
	case "get":
		return runDeviceActionGet(ctx, args[1:], stdout, stderr)
	case "resp":
		return runDeviceActionResp(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeviceActionHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown device action subcommand: %s\n", args[0])
		printDeviceActionHelp(stderr)
		return 2
	}
}

// printDeviceActionHelp 打印设备行为调用帮助信息
func printDeviceActionHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device action <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Call device action (send, get, resp)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  send       Send action call to device")
	fmt.Fprintln(w, "  get        Get action execution result")
	fmt.Fprintln(w, "  resp       Reply to device action call (upstream)")
	fmt.Fprintln(w, "  help       Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Call device action")
	fmt.Fprintln(w, "  ur device action send -p p_smartswitch_001 -d switch-001 --data-id OpenValve --input '{\"Duration\": 30}'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Get action execution result")
	fmt.Fprintln(w, "  ur device action get -p p_smartswitch_001 -d switch-001 --data-id OpenValve")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Reply to device action call")
	fmt.Fprintln(w, "  ur device action resp -p p_smartswitch_001 -d switch-001 --data-id ReadMeter --output '{\"EP\": \"1234.56\"}'")
}

// runDeviceActionSend 执行发送行为调用命令
func runDeviceActionSend(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceActionSendHelp(stderr)
		return 2
	}

	// 解析行为调用特定参数
	dataID := ""
	input := ""

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--data-id":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data-id requires value")
				return 2
			}
			dataID = remaining[i+1]
			i++
		case "--input":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--input requires value")
				return 2
			}
			input = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown action send option: %s\n", remaining[i])
			printDeviceActionSendHelp(stderr)
			return 2
		}
	}

	if dataID == "" {
		fmt.Fprintln(stderr, "--data-id is required")
		printDeviceActionSendHelp(stderr)
		return 2
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"dataID":     dataID,
	}

	if input != "" {
		var inputMap map[string]any
		if err := json.Unmarshal([]byte(input), &inputMap); err != nil {
			fmt.Fprintf(stderr, "Error: --input must be JSON object: %v\n", err)
			return 2
		}
		reqBody["input"] = inputMap
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/interact/action-send",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		if resp.Code != 200 {
			fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
			return 1
		}
		fmt.Fprintln(stdout, "Action call sent successfully")
	}

	return 0
}

// runDeviceActionGet 执行获取行为执行结果命令
func runDeviceActionGet(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceActionGetHelp(stderr)
		return 2
	}

	// 解析行为获取特定参数
	dataID := ""

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--data-id":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data-id requires value")
				return 2
			}
			dataID = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown action get option: %s\n", remaining[i])
			printDeviceActionGetHelp(stderr)
			return 2
		}
	}

	if dataID == "" {
		fmt.Fprintln(stderr, "--data-id is required")
		printDeviceActionGetHelp(stderr)
		return 2
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"dataID":     dataID,
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/interact/action-get-one",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		if resp.Code != 200 {
			fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
			return 1
		}
		fmt.Fprintln(stdout, "Action execution result retrieved successfully")
	}

	return 0
}

// printDeviceActionGetHelp 打印获取行为执行结果帮助信息
func printDeviceActionGetHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device action get [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Get action execution result")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --data-id string           Action identifier (required)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Get action execution result")
	fmt.Fprintln(w, "  ur device action get -p p_smartswitch_001 -d switch-001 --data-id OpenValve")
}

// runDeviceActionResp 执行回复设备行为调用命令
func runDeviceActionResp(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceActionRespHelp(stderr)
		return 2
	}

	// 解析行为回复特定参数
	dataID := ""
	output := ""

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--data-id":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data-id requires value")
				return 2
			}
			dataID = remaining[i+1]
			i++
		case "--output":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--output requires value")
				return 2
			}
			output = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown action resp option: %s\n", remaining[i])
			printDeviceActionRespHelp(stderr)
			return 2
		}
	}

	if dataID == "" {
		fmt.Fprintln(stderr, "--data-id is required")
		printDeviceActionRespHelp(stderr)
		return 2
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"dataID":     dataID,
	}

	if output != "" {
		var outputMap map[string]any
		if err := json.Unmarshal([]byte(output), &outputMap); err != nil {
			fmt.Fprintf(stderr, "Error: --output must be JSON object: %v\n", err)
			return 2
		}
		reqBody["output"] = outputMap
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/interact/action-resp",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		if resp.Code != 200 {
			fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
			return 1
		}
		fmt.Fprintln(stdout, "Action response sent successfully")
	}

	return 0
}

// runDeviceMock 执行生成 Mock 数据命令
func runDeviceMock(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceMockHelp(stderr)
		return 2
	}

	// 解析 Mock 数据特定参数
	dataID := ""
	num := 1

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--data-id":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data-id requires value")
				return 2
			}
			dataID = remaining[i+1]
			i++
		case "--num":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--num requires value")
				return 2
			}
			fmt.Sscanf(remaining[i+1], "%d", &num)
			i++
		default:
			fmt.Fprintf(stderr, "unknown mock option: %s\n", remaining[i])
			printDeviceMockHelp(stderr)
			return 2
		}
	}

	if dataID == "" {
		fmt.Fprintln(stderr, "--data-id is required")
		printDeviceMockHelp(stderr)
		return 2
	}

	// 构建请求体
	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
		"dataID":     dataID,
		"num":        num,
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/interact/schema-mock-gen",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		if resp.Code != 200 {
			fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
			return 1
		}
		fmt.Fprintln(stdout, "Mock data generated successfully")
	}

	return 0
}

// runDeviceReport 执行模拟设备上报消息命令
// 内部自动获取设备密钥并生成 MQTT 认证凭据，无需手动提供 username/password
func runDeviceReport(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	// 检查帮助标志
	for _, arg := range args {
		if arg == "help" || arg == "--help" || arg == "-h" {
			printDeviceReportHelp(stdout)
			return 0
		}
	}

	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceReportHelp(stderr)
		return 2
	}

	// 解析设备上报特定参数
	handle := "thing"
	msgType := "property"
	method := "report"
	params := ""
	msgToken := ""

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--handle":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--handle requires value")
				return 2
			}
			handle = remaining[i+1]
			i++
		case "--type":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--type requires value")
				return 2
			}
			msgType = remaining[i+1]
			i++
		case "--method":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--method requires value")
				return 2
			}
			method = remaining[i+1]
			i++
		case "--params":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--params requires value")
				return 2
			}
			params = remaining[i+1]
			i++
		case "--msg-token":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--msg-token requires value")
				return 2
			}
			msgToken = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown report option: %s\n", remaining[i])
			printDeviceReportHelp(stderr)
			return 2
		}
	}

	if params == "" {
		fmt.Fprintln(stderr, "--params is required")
		printDeviceReportHelp(stderr)
		return 2
	}

	// 解析 params JSON
	var paramsMap map[string]any
	if err := json.Unmarshal([]byte(params), &paramsMap); err != nil {
		fmt.Fprintf(stderr, "Error: --params must be JSON object: %v\n", err)
		return 2
	}

	// 第一步：调用 API 获取设备密钥
	fmt.Fprintf(stderr, "Fetching device secret for %s/%s...\n", productID, deviceName)
	deviceInfoResp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/get-one",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "Error fetching device info: %v\n", err)
		return 1
	}
	if deviceInfoResp.Code != 200 {
		fmt.Fprintf(stderr, "Error fetching device info: %s\n", deviceInfoResp.Msg)
		return 1
	}

	// 从响应中提取 deviceSecret
	deviceInfo, ok := deviceInfoResp.Data.(map[string]any)
	if !ok {
		fmt.Fprintf(stderr, "Error: invalid device info response format\n")
		return 1
	}
	deviceSecret, _ := deviceInfo["deviceSecret"].(string)
	if deviceSecret == "" {
		fmt.Fprintf(stderr, "Error: device secret is empty, device may not be registered\n")
		return 1
	}

	// 第二步：生成 MQTT 认证凭据
	clientID, userName, password := genDeviceAuth(productID, deviceName, deviceSecret)
	fmt.Fprintf(stderr, "MQTT credentials generated (clientID: %s)\n", clientID)

	// 构建请求体
	reqBody := map[string]any{
		"method": method,
		"params": paramsMap,
	}

	if msgToken != "" {
		reqBody["msgToken"] = msgToken
	}

	// 构建 API 路径
	apiPath := fmt.Sprintf("/api/v1/things/device/edge/send/%s/%s", handle, msgType)

	// 构建请求头（使用自动生成的 MQTT 凭据进行 Basic Auth）
	authValue := base64.StdEncoding.EncodeToString([]byte(userName + ":" + password))
	headers := map[string]string{
		"productID":    productID,
		"deviceName":   deviceName,
		"Authorization": "Basic " + authValue,
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path:    apiPath,
		Body:    reqBody,
		Headers: headers,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	// 输出结果
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	} else {
		if resp.Code != 200 {
			fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
			return 1
		}
		fmt.Fprintln(stdout, "Device report message sent successfully")
	}

	return 0
}

// runDeviceUpload 执行设备文件上传命令
func runDeviceUpload(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	// 检查帮助标志
	for _, arg := range args {
		if arg == "help" || arg == "--help" || arg == "-h" {
			printDeviceUploadHelp(stdout)
			return 0
		}
	}

	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		printDeviceUploadHelp(stderr)
		return 2
	}

	// 解析上传特定参数
	filePath := ""
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--file", "-f":
			if i+1 < len(remaining) {
				filePath = remaining[i+1]
				i++
			}
		default:
			fmt.Fprintf(stderr, "unknown upload option: %s\n", remaining[i])
			return 2
		}
	}

	if filePath == "" {
		fmt.Fprintln(stderr, "--file is required")
		printDeviceUploadHelp(stderr)
		return 2
	}

	// 获取设备密钥
	fmt.Fprintf(stderr, "Fetching device secret for %s/%s...\n", productID, deviceName)
	deviceInfoResp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/get-one",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "Error fetching device info: %v\n", err)
		return 1
	}
	if deviceInfoResp.Code != 200 {
		fmt.Fprintf(stderr, "Error fetching device info: %s\n", deviceInfoResp.Msg)
		return 1
	}

	deviceInfo, ok := deviceInfoResp.Data.(map[string]any)
	if !ok {
		fmt.Fprintf(stderr, "Error: invalid device info response format\n")
		return 1
	}
	deviceSecret, _ := deviceInfo["deviceSecret"].(string)
	if deviceSecret == "" {
		fmt.Fprintf(stderr, "Error: device secret is empty\n")
		return 1
	}

	// 生成 MQTT 认证凭据
	clientID, userName, password := genDeviceAuth(productID, deviceName, deviceSecret)
	fmt.Fprintf(stderr, "MQTT credentials generated (clientID: %s)\n", clientID)

	// 上传文件需要直接调用 HTTP API，不走 client.DoAPI
	// 因为需要 multipart/form-data 格式
	baseURL, _ := config.GetBaseURL()
	authValue := base64.StdEncoding.EncodeToString([]byte(userName + ":" + password))

	fmt.Fprintf(stderr, "Uploading file: %s\n", filePath)
	fmt.Fprintf(stderr, "Use curl to upload:\n")
	fmt.Fprintf(stdout, "curl -X POST \"%s/api/v1/things/device/edge/upload-file\" -H \"Authorization: Basic %s\" -F \"file=@%s\"\n",
		strings.TrimRight(baseURL, "/"), authValue, filePath)

	if jsonOutput {
		raw, _ := json.MarshalIndent(map[string]any{
			"clientID":  clientID,
			"userName":  userName,
			"password":  password,
			"authBasic": "Basic " + authValue,
			"filePath":  filePath,
		}, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	}

	return 0
}

// printDeviceUploadHelp 打印设备文件上传帮助信息
func printDeviceUploadHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device upload [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Upload file from device (auto-generates MQTT credentials)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  -f, --file string          File path to upload (required)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
}

// printDeviceReportHelp 打印模拟设备上报消息帮助信息
func printDeviceReportHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device report [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Simulate device report message via HTTP")
	fmt.Fprintln(w, "MQTT credentials are auto-generated from device secret (no manual username/password needed)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --handle string            Handle type: thing/ota/config (default: thing)")
	fmt.Fprintln(w, "  --type string              Message type: property/event/action (default: property)")
	fmt.Fprintln(w, "  --method string            Method: report (default: report)")
	fmt.Fprintln(w, "  --params string            Report parameters JSON (required)")
	fmt.Fprintln(w, "  --msg-token string         Message token")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Report device properties")
	fmt.Fprintln(w, "  ur device report -p p_smartswitch_001 -d switch-001 --params '{\"Temperature\": 25.3, \"Humidity\": 60}'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Report device event")
	fmt.Fprintln(w, "  ur device report -p p_smartswitch_001 -d switch-001 --type event --params '{\"PowerAlarm\": {\"Voltage\": 220}}'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Report device action")
	fmt.Fprintln(w, "  ur device report -p p_smartswitch_001 -d switch-001 --type action --params '{\"ReadMeter\": {\"EP\": \"1234.56\"}}'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Note:")
	fmt.Fprintln(w, "  - Device secret is fetched automatically via API")
	fmt.Fprintln(w, "  - MQTT credentials (username/password) are generated using HMAC-SHA256")
	fmt.Fprintln(w, "  - Uses Basic Auth to authenticate HTTP requests")
	fmt.Fprintln(w, "  - See: https://unitedrhino.apifox.cn/api-233401356 for API details")
}

// printDeviceMockHelp 打印生成 Mock 数据帮助信息
func printDeviceMockHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device mock [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Generate mock data based on thing model")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --data-id string           Property/action/event identifier (required)")
	fmt.Fprintln(w, "  --num int                  Number of mock data to generate (default: 1)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Generate 5 temperature mock data")
	fmt.Fprintln(w, "  ur device mock -p p_smartswitch_001 -d switch-001 --data-id Temperature --num 5")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Generate 1 power switch mock data")
	fmt.Fprintln(w, "  ur device mock -p p_smartswitch_001 -d switch-001 --data-id PowerSwitch")
}

// printDeviceActionRespHelp 打印回复设备行为调用帮助信息
func printDeviceActionRespHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device action resp [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Reply to device action call (upstream)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --data-id string           Action identifier (required)")
	fmt.Fprintln(w, "  --output string            Action output parameters JSON")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Reply to device action call")
	fmt.Fprintln(w, "  ur device action resp -p p_smartswitch_001 -d switch-001 --data-id ReadMeter --output '{\"EP\": \"1234.56\"}'")
}

// printDeviceActionSendHelp 打印发送行为调用帮助信息
func printDeviceActionSendHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device action send [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Send action call to device")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --data-id string           Action identifier (required)")
	fmt.Fprintln(w, "  --input string             Action input parameters JSON")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Call device action")
	fmt.Fprintln(w, "  ur device action send -p p_smartswitch_001 -d switch-001 --data-id OpenValve --input '{\"Duration\": 30}'")
}

// printDeviceControlHelp 打印设备属性控制帮助信息
func printDeviceControlHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device control [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Send property control command to device")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --data string              Property key-value pairs JSON (required)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Control device power switch")
	fmt.Fprintln(w, "  ur device control -p p_smartswitch_001 -d switch-001 --data '{\"PowerSwitch\": 1}'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Control device brightness")
	fmt.Fprintln(w, "  ur device control -p p_smartswitch_001 -d switch-001 --data '{\"Brightness\": 80}'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Control multiple properties")
	fmt.Fprintln(w, "  ur device control -p p_smartswitch_001 -d switch-001 --data '{\"PowerSwitch\": 1, \"Brightness\": 80}'")
}

// printDeviceLogSDKHelp 打印 SDK 日志查询帮助信息
func printDeviceLogSDKHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device log sdk [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device SDK logs")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --log-level int            Log level: 1=off, 2=error, 3=warn, 4=info, 5=debug")
	fmt.Fprintln(w, "  --time-start string        Start timestamp (milliseconds)")
	fmt.Fprintln(w, "  --time-end string          End timestamp (milliseconds)")
	fmt.Fprintln(w, "  --page int                 Page number (default: 1)")
	fmt.Fprintln(w, "  --size int                 Page size (default: 20)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device SDK logs")
	fmt.Fprintln(w, "  ur device log sdk -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query device error logs")
	fmt.Fprintln(w, "  ur device log sdk -p p_smartswitch_001 -d switch-001 --log-level 2")
}

// printDeviceLogAbnormalHelp 打印异常日志查询帮助信息
func printDeviceLogAbnormalHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device log abnormal [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device abnormal logs")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --action int               1=trigger, 2=recover")
	fmt.Fprintln(w, "  --type string              Abnormal type")
	fmt.Fprintln(w, "  --time-start string        Start timestamp (milliseconds)")
	fmt.Fprintln(w, "  --time-end string          End timestamp (milliseconds)")
	fmt.Fprintln(w, "  --page int                 Page number (default: 1)")
	fmt.Fprintln(w, "  --size int                 Page size (default: 20)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device abnormal logs")
	fmt.Fprintln(w, "  ur device log abnormal -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query trigger abnormal logs")
	fmt.Fprintln(w, "  ur device log abnormal -p p_smartswitch_001 -d switch-001 --action 1")
}

// printDeviceLogHubHelp 打印诊断日志查询帮助信息
func printDeviceLogHubHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device log hub [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device diagnostic logs (MQTT communication)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --actions string           MQTT actions, comma separated: connected/disconnected/property/event/action/thing")
	fmt.Fprintln(w, "  --topics string            MQTT topics, comma separated")
	fmt.Fprintln(w, "  --content string           Content fuzzy match")
	fmt.Fprintln(w, "  --request-id string        Request token exact match")
	fmt.Fprintln(w, "  --time-start string        Start timestamp (milliseconds)")
	fmt.Fprintln(w, "  --time-end string          End timestamp (milliseconds)")
	fmt.Fprintln(w, "  --page int                 Page number (default: 1)")
	fmt.Fprintln(w, "  --size int                 Page size (default: 20)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device diagnostic logs")
	fmt.Fprintln(w, "  ur device log hub -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query property MQTT messages")
	fmt.Fprintln(w, "  ur device log hub -p p_smartswitch_001 -d switch-001 --actions property --content report")
}

// printDeviceLogStatusHelp 打印上下线日志查询帮助信息
func printDeviceLogStatusHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device log status [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device online/offline status logs")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --status int               Status filter: 1=online, 2=offline")
	fmt.Fprintln(w, "  --time-start string        Start timestamp (milliseconds)")
	fmt.Fprintln(w, "  --time-end string          End timestamp (milliseconds)")
	fmt.Fprintln(w, "  --page int                 Page number (default: 1)")
	fmt.Fprintln(w, "  --size int                 Page size (default: 20)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device status logs")
	fmt.Fprintln(w, "  ur device log status -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query device online logs")
	fmt.Fprintln(w, "  ur device log status -p p_smartswitch_001 -d switch-001 --status 1")
}

// printDeviceLogSendHelp 打印命令日志查询帮助信息
func printDeviceLogSendHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device log send [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device command send logs")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --actions string           Command types, comma separated: propertyControlSend/propertyGetReportSend/actionSend")
	fmt.Fprintln(w, "  --result-code int          Result code filter: 200=success")
	fmt.Fprintln(w, "  --data-id string           Property/action identifier filter")
	fmt.Fprintln(w, "  --with-user                Include operator information")
	fmt.Fprintln(w, "  --time-start string        Start timestamp (milliseconds)")
	fmt.Fprintln(w, "  --time-end string          End timestamp (milliseconds)")
	fmt.Fprintln(w, "  --page int                 Page number (default: 1)")
	fmt.Fprintln(w, "  --size int                 Page size (default: 20)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device command send logs")
	fmt.Fprintln(w, "  ur device log send -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query property control commands with operator info")
	fmt.Fprintln(w, "  ur device log send -p p_smartswitch_001 -d switch-001 --actions propertyControlSend --with-user")
}

// printDeviceLogEventHelp 打印事件日志查询帮助信息
func printDeviceLogEventHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device log event [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device event logs")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --types string             Event types, comma separated: info/alert/fault")
	fmt.Fprintln(w, "  --data-id string           Specific event identifier")
	fmt.Fprintln(w, "  --time-start string        Start timestamp (milliseconds)")
	fmt.Fprintln(w, "  --time-end string          End timestamp (milliseconds)")
	fmt.Fprintln(w, "  --page int                 Page number (default: 1)")
	fmt.Fprintln(w, "  --size int                 Page size (default: 20)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device event logs")
	fmt.Fprintln(w, "  ur device log event -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query alert and fault events")
	fmt.Fprintln(w, "  ur device log event -p p_smartswitch_001 -d switch-001 --types alert,fault")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query specific event")
	fmt.Fprintln(w, "  ur device log event -p p_smartswitch_001 -d switch-001 --data-id PowerAlarm")
}

// printDeviceLogPropertyHelp 打印属性日志查询帮助信息
func printDeviceLogPropertyHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device log property [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device property logs (latest value, history)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (required)")
	fmt.Fprintln(w, "  --data-id string           Property identifier (for history query)")
	fmt.Fprintln(w, "  --data-ids string          Property identifiers, comma separated (for latest query)")
	fmt.Fprintln(w, "  --ignore-empty             Skip empty values")
	fmt.Fprintln(w, "  --arg-func string          Aggregation function: avg/first/last/count/twa")
	fmt.Fprintln(w, "  --interval int             Time window interval")
	fmt.Fprintln(w, "  --interval-unit string     Time window unit: s/m/h/d")
	fmt.Fprintln(w, "  --order int                Sort order: 1=ascending, 2=descending")
	fmt.Fprintln(w, "  --time-start string        Start timestamp (milliseconds)")
	fmt.Fprintln(w, "  --time-end string          End timestamp (milliseconds)")
	fmt.Fprintln(w, "  --page int                 Page number (default: 1)")
	fmt.Fprintln(w, "  --size int                 Page size (default: 20)")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device latest property values")
	fmt.Fprintln(w, "  ur device log property -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query device temperature history with average aggregation")
	fmt.Fprintln(w, "  ur device log property -p p_smartswitch_001 -d switch-001 --data-id Temperature --arg-func avg --interval 1 --interval-unit h")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query multiple properties latest values")
	fmt.Fprintln(w, "  ur device log property -p p_smartswitch_001 -d switch-001 --data-ids Temperature,Humidity")
}

// runDeviceInfo 执行设备信息管理命令
func runDeviceInfo(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeviceInfoHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runDeviceInfoGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runDeviceInfoGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runDeviceInfoCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runDeviceInfoUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runDeviceInfoDelete(ctx, args[1:], stdout, stderr)
	case "bind":
		return runDeviceInfoBind(ctx, args[1:], stdout, stderr)
	case "unbind":
		return runDeviceInfoUnbind(ctx, args[1:], stdout, stderr)
	case "count":
		return runDeviceInfoCount(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeviceInfoHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown device info subcommand: %s\n", args[0])
		printDeviceInfoHelp(stderr)
		return 2
	}
}

// printDeviceInfoHelp 打印设备信息管理帮助信息
func printDeviceInfoHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device info <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Device information management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query device list")
	fmt.Fprintln(w, "  get-one    Query device detail")
	fmt.Fprintln(w, "  create     Create device")
	fmt.Fprintln(w, "  update     Update device")
	fmt.Fprintln(w, "  delete     Delete device")
	fmt.Fprintln(w, "  bind       Bind device to gateway")
	fmt.Fprintln(w, "  unbind     Unbind device from gateway")
	fmt.Fprintln(w, "  count      Count devices")
	fmt.Fprintln(w, "  help       Show this help message")
}

// parseInfoListParams 解析设备列表查询通用参数
func parseInfoListParams(args []string) (jsonOutput bool, page, size int, remaining []string) {
	page = 1
	size = 20
	remaining = make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--page":
			if i+1 < len(args) {
				fmt.Sscanf(args[i+1], "%d", &page)
				i++
			}
		case "--size":
			if i+1 < len(args) {
				fmt.Sscanf(args[i+1], "%d", &size)
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			remaining = append(remaining, args[i])
		}
	}
	return
}

// runDeviceInfoGetList 执行查询设备列表命令
func runDeviceInfoGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)

	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	// 解析可选过滤参数
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--product-id", "-p":
			if i+1 < len(remaining) {
				reqBody["productID"] = remaining[i+1]
				i++
			}
		case "--device-name", "-d":
			if i+1 < len(remaining) {
				reqBody["deviceName"] = remaining[i+1]
				i++
			}
		case "--group-id":
			if i+1 < len(remaining) {
				reqBody["groupID"] = remaining[i+1]
				i++
			}
		case "--tag":
			if i+1 < len(remaining) {
				reqBody["tag"] = remaining[i+1]
				i++
			}
		case "--area-id":
			if i+1 < len(remaining) {
				reqBody["areaID"] = remaining[i+1]
				i++
			}
		case "--project-id":
			if i+1 < len(remaining) {
				reqBody["projectID"] = remaining[i+1]
				i++
			}
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", remaining[i])
			return 2
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceInfoGetOne 执行查询设备详情命令
func runDeviceInfoGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, _, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/get-one",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceInfoCreate 执行创建设备命令
func runDeviceInfoCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--alias":
			if i+1 < len(remaining) {
				reqBody["aliasName"] = remaining[i+1]
				i++
			}
		case "--tags":
			if i+1 < len(remaining) {
				var tags []any
				if err := json.Unmarshal([]byte(remaining[i+1]), &tags); err == nil {
					reqBody["tags"] = tags
				}
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceInfoUpdate 执行更新设备命令
func runDeviceInfoUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--alias":
			if i+1 < len(remaining) {
				reqBody["aliasName"] = remaining[i+1]
				i++
			}
		case "--tags":
			if i+1 < len(remaining) {
				var tags []any
				if err := json.Unmarshal([]byte(remaining[i+1]), &tags); err == nil {
					reqBody["tags"] = tags
				}
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceInfoDelete 执行删除设备命令
func runDeviceInfoDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, _, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/delete",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceInfoBind 执行绑定设备到网关命令
func runDeviceInfoBind(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--gateway-id":
			if i+1 < len(remaining) {
				reqBody["gatewayProductID"] = remaining[i+1]
				i++
			}
		case "--gateway-name":
			if i+1 < len(remaining) {
				reqBody["gatewayDeviceName"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/bind",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceInfoUnbind 执行解绑设备命令
func runDeviceInfoUnbind(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, _, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/unbind",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceInfoCount 执行统计设备数量命令
func runDeviceInfoCount(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	remaining := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		default:
			remaining = append(remaining, args[i])
		}
	}

	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--product-id", "-p":
			if i+1 < len(remaining) {
				reqBody["productID"] = remaining[i+1]
				i++
			}
		case "--project-id":
			if i+1 < len(remaining) {
				reqBody["projectID"] = remaining[i+1]
				i++
			}
		case "--area-id":
			if i+1 < len(remaining) {
				reqBody["areaID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/info/count",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// outputResult 统一输出 API 响应结果
func outputResult(resp client.APIResponse, jsonOutput bool, stdout, stderr io.Writer) int {
	if jsonOutput {
		raw, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Fprintln(stdout, string(raw))
		// --json 模式下业务失败同样返回非 0 退出码，保证脚本可判断成败
		if resp.Code != 200 {
			return 1
		}
		return 0
	}
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
		return 1
	}
	data, err := json.MarshalIndent(resp.Data, "", "  ")
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 1
	}
	fmt.Fprintln(stdout, string(data))
	return 0
}

// printDeviceHelp 打印设备命令的帮助信息
func printDeviceHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Device debugging and log query commands")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  log        Query device logs (property, event, send, status, hub, abnormal, sdk)")
	fmt.Fprintln(w, "  info       Device information management (get-list, get-one, create, update, delete, bind, unbind, count)")
	fmt.Fprintln(w, "  gateway    Gateway sub-device management (get-list, batch-create, batch-delete)")
	fmt.Fprintln(w, "  group      Device group management (get-list, create, delete, batch-create-device, batch-delete-device)")
	fmt.Fprintln(w, "  profile    Device profile management (get-list, get-one, update, delete)")
	fmt.Fprintln(w, "  control    Send property control command to device")
	fmt.Fprintln(w, "  action     Call device action (send, get, resp)")
	fmt.Fprintln(w, "  mock       Generate mock data based on thing model")
	fmt.Fprintln(w, "  report     Simulate device report message via HTTP")
	fmt.Fprintln(w, "  upload     Upload file from device")
	fmt.Fprintln(w, "  help       Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device latest property values")
	fmt.Fprintln(w, "  ur device log property -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query device temperature history")
	fmt.Fprintln(w, "  ur device log property -p p_smartswitch_001 -d switch-001 --data-id Temperature --arg-func avg")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Control device property")
	fmt.Fprintln(w, "  ur device control -p p_smartswitch_001 -d switch-001 --data '{\"PowerSwitch\": 1}'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Call device action")
	fmt.Fprintln(w, "  ur device action send -p p_smartswitch_001 -d switch-001 --data-id OpenValve --input '{\"Duration\": 30}'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Generate mock data")
	fmt.Fprintln(w, "  ur device mock -p p_smartswitch_001 -d switch-001 --data-id Temperature --num 5")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Simulate device report")
	fmt.Fprintln(w, "  ur device report -p p_smartswitch_001 -d switch-001 --params '{\"Temperature\": 25.3}'")
}

// parseDeviceParams 解析设备通用参数
func parseDeviceParams(args []string) (productID, deviceName string, jsonOutput bool, remaining []string, err error) {
	remaining = make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--product-id", "-p":
			if i+1 >= len(args) {
				return "", "", false, nil, fmt.Errorf("--product-id requires value")
			}
			productID = args[i+1]
			i++
		case "--device-name", "-d":
			if i+1 >= len(args) {
				return "", "", false, nil, fmt.Errorf("--device-name requires value")
			}
			deviceName = args[i+1]
			i++
		case "--json", "-j":
			jsonOutput = true
		default:
			remaining = append(remaining, args[i])
		}
	}
	if productID == "" {
		return "", "", false, nil, fmt.Errorf("--product-id is required")
	}
	if deviceName == "" {
		return "", "", false, nil, fmt.Errorf("--device-name is required")
	}
	return productID, deviceName, jsonOutput, remaining, nil
}

// genDeviceAuth 根据设备三元组生成 MQTT 认证凭据
// clientID: productID&deviceName
// userName: clientID;12010126;connID;expiry
// password: hmac_sha256(base64_decode(deviceSecret), userName) + ";hmacsha256"
func genDeviceAuth(productID, deviceName, deviceSecret string) (clientID, userName, password string) {
	connID := randomString(5)
	expiry := time.Now().AddDate(0, 0, 10).Unix()
	clientID = fmt.Sprintf("%s&%s", productID, deviceName)
	userName = fmt.Sprintf("%s;12010126;%s;%d", clientID, connID, expiry)
	decodedSecret, _ := base64.StdEncoding.DecodeString(deviceSecret)
	h := hmac.New(sha256.New, decodedSecret)
	h.Write([]byte(userName))
	token := hex.EncodeToString(h.Sum(nil))
	password = fmt.Sprintf("%s;hmacsha256", token)
	return
}

// randomString 生成指定长度的随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// runDeviceGateway 执行网关子设备管理命令
func runDeviceGateway(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeviceGatewayHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runDeviceGatewayGetList(ctx, args[1:], stdout, stderr)
	case "batch-create":
		return runDeviceGatewayBatchCreate(ctx, args[1:], stdout, stderr)
	case "batch-delete":
		return runDeviceGatewayBatchDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeviceGatewayHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown device gateway subcommand: %s\n", args[0])
		printDeviceGatewayHelp(stderr)
		return 2
	}
}

// printDeviceGatewayHelp 打印网关子设备管理帮助信息
func printDeviceGatewayHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device gateway <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Gateway sub-device management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list       Query gateway sub-device list")
	fmt.Fprintln(w, "  batch-create   Batch add sub-devices to gateway")
	fmt.Fprintln(w, "  batch-delete   Batch remove sub-devices from gateway")
	fmt.Fprintln(w, "  help           Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query gateway sub-devices")
	fmt.Fprintln(w, "  ur device gateway get-list -p gateway_pid -d gateway_001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Batch add sub-devices")
	fmt.Fprintln(w, "  ur device gateway batch-create -p gateway_pid -d gateway_001 --devices '[{\"productID\":\"sub_pid\",\"deviceName\":\"sub_001\"}]'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Batch delete sub-devices")
	fmt.Fprintln(w, "  ur device gateway batch-delete -p gateway_pid -d gateway_001 --devices '[{\"productID\":\"sub_pid\",\"deviceName\":\"sub_001\"}]'")
}

// runDeviceGatewayGetList 执行查询网关子设备列表命令
func runDeviceGatewayGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, _, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/gateway/get-list",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceGatewayBatchCreate 执行批量添加网关子设备命令
func runDeviceGatewayBatchCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	devices := ""
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--devices":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--devices requires value")
				return 2
			}
			devices = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", remaining[i])
			return 2
		}
	}

	if devices == "" {
		fmt.Fprintln(stderr, "--devices is required")
		return 2
	}

	var deviceList []map[string]any
	if err := json.Unmarshal([]byte(devices), &deviceList); err != nil {
		fmt.Fprintf(stderr, "Error: --devices must be JSON array: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/gateway/batch-create",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
			"devices":    deviceList,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceGatewayBatchDelete 执行批量删除网关子设备命令
func runDeviceGatewayBatchDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	devices := ""
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--devices":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--devices requires value")
				return 2
			}
			devices = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", remaining[i])
			return 2
		}
	}

	if devices == "" {
		fmt.Fprintln(stderr, "--devices is required")
		return 2
	}

	var deviceList []map[string]any
	if err := json.Unmarshal([]byte(devices), &deviceList); err != nil {
		fmt.Fprintf(stderr, "Error: --devices must be JSON array: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/gateway/batch-delete",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
			"devices":    deviceList,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceGroup 执行设备分组管理命令
func runDeviceGroup(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeviceGroupHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runDeviceGroupGetList(ctx, args[1:], stdout, stderr)
	case "create":
		return runDeviceGroupCreate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runDeviceGroupDelete(ctx, args[1:], stdout, stderr)
	case "batch-create-device":
		return runDeviceGroupBatchCreateDevice(ctx, args[1:], stdout, stderr)
	case "batch-delete-device":
		return runDeviceGroupBatchDeleteDevice(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeviceGroupHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown device group subcommand: %s\n", args[0])
		printDeviceGroupHelp(stderr)
		return 2
	}
}

// printDeviceGroupHelp 打印设备分组管理帮助信息
func printDeviceGroupHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device group <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Device group management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list              Query device group list")
	fmt.Fprintln(w, "  create                Create device group")
	fmt.Fprintln(w, "  delete                Delete device group")
	fmt.Fprintln(w, "  batch-create-device   Batch add devices to group")
	fmt.Fprintln(w, "  batch-delete-device   Batch remove devices from group")
	fmt.Fprintln(w, "  help                  Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device groups")
	fmt.Fprintln(w, "  ur device group get-list")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Create device group")
	fmt.Fprintln(w, "  ur device group create --name \"My Group\" --desc \"Description\"")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Batch add devices to group")
	fmt.Fprintln(w, "  ur device group batch-create-device --group-id 123 --devices '[{\"productID\":\"pid\",\"deviceName\":\"dev1\"}]'")
}

// runDeviceGroupGetList 执行查询设备分组列表命令
func runDeviceGroupGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--name":
			if i+1 < len(remaining) {
				reqBody["name"] = remaining[i+1]
				i++
			}
		case "--group-id":
			if i+1 < len(remaining) {
				reqBody["groupID"] = remaining[i+1]
				i++
			}
		case "--parent-id":
			if i+1 < len(remaining) {
				reqBody["parentID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/group/info/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceGroupCreate 执行创建设备分组命令
func runDeviceGroupCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	name := ""
	desc := ""
	parentID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--name":
			if i+1 < len(args) {
				name = args[i+1]
				i++
			}
		case "--desc":
			if i+1 < len(args) {
				desc = args[i+1]
				i++
			}
		case "--parent-id":
			if i+1 < len(args) {
				parentID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if name == "" {
		fmt.Fprintln(stderr, "--name is required")
		return 2
	}

	reqBody := map[string]any{
		"name": name,
	}
	if desc != "" {
		reqBody["desc"] = desc
	}
	if parentID != "" {
		reqBody["parentID"] = parentID
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/group/info/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceGroupDelete 执行删除设备分组命令
func runDeviceGroupDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	groupID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--group-id":
			if i+1 < len(args) {
				groupID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if groupID == "" {
		fmt.Fprintln(stderr, "--group-id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/group/info/delete",
		Body: map[string]any{
			"groupID": groupID,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceGroupBatchCreateDevice 执行批量添加设备到分组命令
func runDeviceGroupBatchCreateDevice(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	groupID := ""
	devices := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--group-id":
			if i+1 < len(args) {
				groupID = args[i+1]
				i++
			}
		case "--devices":
			if i+1 < len(args) {
				devices = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if groupID == "" {
		fmt.Fprintln(stderr, "--group-id is required")
		return 2
	}
	if devices == "" {
		fmt.Fprintln(stderr, "--devices is required")
		return 2
	}

	var deviceList []map[string]any
	if err := json.Unmarshal([]byte(devices), &deviceList); err != nil {
		fmt.Fprintf(stderr, "Error: --devices must be JSON array: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/group/device/batch-create",
		Body: map[string]any{
			"groupID": groupID,
			"devices": deviceList,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceGroupBatchDeleteDevice 执行批量从分组删除设备命令
func runDeviceGroupBatchDeleteDevice(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	groupID := ""
	devices := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--group-id":
			if i+1 < len(args) {
				groupID = args[i+1]
				i++
			}
		case "--devices":
			if i+1 < len(args) {
				devices = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if groupID == "" {
		fmt.Fprintln(stderr, "--group-id is required")
		return 2
	}
	if devices == "" {
		fmt.Fprintln(stderr, "--devices is required")
		return 2
	}

	var deviceList []map[string]any
	if err := json.Unmarshal([]byte(devices), &deviceList); err != nil {
		fmt.Fprintf(stderr, "Error: --devices must be JSON array: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/group/device/batch-delete",
		Body: map[string]any{
			"groupID": groupID,
			"devices": deviceList,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceProfile 执行设备配置管理命令
func runDeviceProfile(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeviceProfileHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runDeviceProfileGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runDeviceProfileGetOne(ctx, args[1:], stdout, stderr)
	case "update":
		return runDeviceProfileUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runDeviceProfileDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeviceProfileHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown device profile subcommand: %s\n", args[0])
		printDeviceProfileHelp(stderr)
		return 2
	}
}

// printDeviceProfileHelp 打印设备配置管理帮助信息
func printDeviceProfileHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur device profile <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Device profile management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query device profile list")
	fmt.Fprintln(w, "  get-one    Query device profile detail")
	fmt.Fprintln(w, "  update     Update device profile")
	fmt.Fprintln(w, "  delete     Delete device profile")
	fmt.Fprintln(w, "  help       Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query device profile list")
	fmt.Fprintln(w, "  ur device profile get-list -p product_id")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query device profile detail")
	fmt.Fprintln(w, "  ur device profile get-one -p product_id -d device_name")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Update device profile")
	fmt.Fprintln(w, "  ur device profile update -p product_id -d device_name --data '{\"key\":\"value\"}'")
}

// runDeviceProfileGetList 执行查询设备配置列表命令
func runDeviceProfileGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		case "--device-name", "-d":
			if i+1 < len(remaining) {
				reqBody["deviceName"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/profile/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceProfileGetOne 执行查询设备配置详情命令
func runDeviceProfileGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, _, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/profile/get-one",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceProfileUpdate 执行更新设备配置命令
func runDeviceProfileUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	data := ""
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--data":
			if i+1 >= len(remaining) {
				fmt.Fprintln(stderr, "--data requires value")
				return 2
			}
			data = remaining[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", remaining[i])
			return 2
		}
	}

	if data == "" {
		fmt.Fprintln(stderr, "--data is required")
		return 2
	}

	var dataMap map[string]any
	if err := json.Unmarshal([]byte(data), &dataMap); err != nil {
		fmt.Fprintf(stderr, "Error: --data must be JSON object: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID":  productID,
		"deviceName": deviceName,
	}
	for k, v := range dataMap {
		reqBody[k] = v
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/profile/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeviceProfileDelete 执行删除设备配置命令
func runDeviceProfileDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, _, err := parseDeviceParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/profile/delete",
		Body: map[string]any{
			"productID":  productID,
			"deviceName": deviceName,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// formatDeviceResult 格式化设备命令结果
func formatDeviceResult(resp client.APIResponse, stdout, stderr io.Writer) {
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
		return
	}

	data, ok := resp.Data.(map[string]any)
	if !ok {
		fmt.Fprintln(stdout, "No data")
		return
	}

	list, ok := data["list"].([]any)
	if !ok || len(list) == 0 {
		fmt.Fprintln(stdout, "No data")
		return
	}

	for _, item := range list {
		m, ok := item.(map[string]any)
		if !ok {
			continue
		}
		// 格式化输出各个字段
		if timestamp, ok := m["timestamp"]; ok {
			fmt.Fprintf(stdout, "Time: %v\n", timestamp)
		}
		if dataID, ok := m["dataID"]; ok {
			fmt.Fprintf(stdout, "DataID: %v\n", dataID)
		}
		if dataName, ok := m["dataName"]; ok {
			fmt.Fprintf(stdout, "Name: %v\n", dataName)
		}
		if value, ok := m["value"]; ok {
			fmt.Fprintf(stdout, "Value: %v\n", value)
		}
		if status, ok := m["status"]; ok {
			fmt.Fprintf(stdout, "Status: %v\n", status)
		}
		if action, ok := m["action"]; ok {
			fmt.Fprintf(stdout, "Action: %v\n", action)
		}
		if account, ok := m["account"]; ok {
			fmt.Fprintf(stdout, "Account: %v\n", account)
		}
		if resultCode, ok := m["resultCode"]; ok {
			fmt.Fprintf(stdout, "ResultCode: %v\n", resultCode)
		}
		if content, ok := m["content"]; ok {
			fmt.Fprintf(stdout, "Content: %v\n", content)
		}
		if topic, ok := m["topic"]; ok {
			fmt.Fprintf(stdout, "Topic: %v\n", topic)
		}
		if traceID, ok := m["traceID"]; ok {
			fmt.Fprintf(stdout, "TraceID: %v\n", traceID)
		}
		if requestID, ok := m["requestID"]; ok {
			fmt.Fprintf(stdout, "RequestID: %v\n", requestID)
		}
		fmt.Fprintln(stdout, "---")
	}
}
