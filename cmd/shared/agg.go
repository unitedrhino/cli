package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runAgg 执行聚合查询命令
func runAgg(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	// 解析参数
	productID := ""
	deviceName := ""
	dataID := ""
	funcs := ""
	fill := ""
	noFirstTs := false
	jsonOutput := false

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--product-id", "-p":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--product-id requires value")
				return 2
			}
			productID = args[i+1]
			i++
		case "--device-name", "-d":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--device-name requires value")
				return 2
			}
			deviceName = args[i+1]
			i++
		case "--data-id", "-i":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--data-id requires value")
				return 2
			}
			dataID = args[i+1]
			i++
		case "--funcs", "-f":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--funcs requires value")
				return 2
			}
			funcs = args[i+1]
			i++
		case "--fill":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--fill requires value")
				return 2
			}
			fill = args[i+1]
			i++
		case "--no-first-ts":
			noFirstTs = true
		case "--json", "-j":
			jsonOutput = true
		case "--help", "-h":
			printAggHelp(stdout)
			return 0
		default:
			fmt.Fprintf(stderr, "unknown agg option: %s\n", args[i])
			printAggHelp(stderr)
			return 2
		}
	}

	// 验证必需参数
	if productID == "" {
		fmt.Fprintln(stderr, "--product-id is required")
		printAggHelp(stderr)
		return 2
	}
	if dataID == "" {
		fmt.Fprintln(stderr, "--data-id is required")
		printAggHelp(stderr)
		return 2
	}
	if funcs == "" {
		fmt.Fprintln(stderr, "--funcs is required")
		printAggHelp(stderr)
		return 2
	}

	// 解析聚合函数
	funcList := strings.Split(funcs, ",")
	for i, f := range funcList {
		funcList[i] = strings.TrimSpace(f)
	}

	// 构建请求体
	agg := map[string]any{
		"dataID":   dataID,
		"argFuncs": funcList,
	}
	if fill != "" {
		agg["fill"] = fill
	}
	if noFirstTs {
		agg["noFirstTs"] = true
	}

	reqBody := map[string]any{
		"productID": productID,
		"aggs":      []map[string]any{agg},
	}

	if deviceName != "" {
		reqBody["deviceName"] = deviceName
	}

	// 调用 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/device/msg/property-latest-agg/get-list",
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
		formatAggResult(resp, stdout, stderr)
	}

	return 0
}

// formatAggResult 格式化聚合结果
func formatAggResult(resp client.APIResponse, stdout, stderr io.Writer) {
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
		fmt.Fprintf(stdout, "DataID: %v\n", m["dataID"])
		if value, ok := m["value"]; ok {
			fmt.Fprintf(stdout, "Value: %v\n", value)
		}
		if timestamp, ok := m["timestamp"]; ok {
			fmt.Fprintf(stdout, "Time: %v\n", timestamp)
		}
		fmt.Fprintln(stdout, "---")
	}
}

// printAggHelp 打印聚合查询命令的帮助信息
func printAggHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur agg [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Query device property aggregation values (e.g., average, max, min)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -p, --product-id string    Product ID (required)")
	fmt.Fprintln(w, "  -d, --device-name string   Device name (optional, query all devices if not specified)")
	fmt.Fprintln(w, "  -i, --data-id string       Property identifier (required, get from thing model)")
	fmt.Fprintln(w, "  -f, --funcs string         Aggregation functions, comma separated (required)")
	fmt.Fprintln(w, "                             Supported: avg, first, last, count, twa, max, min, sum")
	fmt.Fprintln(w, "      --fill string          Fill mode for missing data")
	fmt.Fprintln(w, "      --no-first-ts          Do not fill earliest timestamp")
	fmt.Fprintln(w, "  -j, --json                 Output in JSON format")
	fmt.Fprintln(w, "  -h, --help                 Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query average CPU usage for a device")
	fmt.Fprintln(w, "  ur agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query max and min temperature for a device")
	fmt.Fprintln(w, "  ur agg -p p_smartswitch_001 -d switch-001 -i Temperature -f max,min")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query average temperature for all devices under a product")
	fmt.Fprintln(w, "  ur agg -p p_smartswitch_001 -i Temperature -f avg")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # JSON output")
	fmt.Fprintln(w, "  ur agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg -j")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Note:")
	fmt.Fprintln(w, "  - Property identifier (data-id) must be obtained from thing model")
	fmt.Fprintln(w, "  - Use 'ur api /api/v1/things/product/schema/get-list' to query thing model")
	fmt.Fprintln(w, "  - Use 'ur api /api/v1/things/device/schema/get-list' to query device thing model")
}
