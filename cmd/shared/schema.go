package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/config"
	"gitee.com/unitedrhino/cli/internal/swagger"
)

func runSchema(app config.CLIApp, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printSchemaHelp(stdout)
		return 0
	}

	// 检查是否是 thing model 管理子命令
	switch args[0] {
	case "get-list", "create", "update", "delete", "tsl-import", "tsl-read":
		return runSchemaModel(args[0], args[1:], stdout, stderr)
	case "browse":
		return runSchemaBrowse(app, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printSchemaHelp(stdout)
		return 0
	default:
		// 默认行为：swagger 浏览（兼容旧行为）
		return runSchemaBrowse(app, args, stdout, stderr)
	}
}

// printSchemaHelp 打印物模型管理帮助信息
func printSchemaHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur schema <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Thing model (schema) management and API browsing")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query schema list")
	fmt.Fprintln(w, "  create     Create schema")
	fmt.Fprintln(w, "  update     Update schema")
	fmt.Fprintln(w, "  delete     Delete schema")
	fmt.Fprintln(w, "  tsl-import Import TSL (Thing Specification Language)")
	fmt.Fprintln(w, "  tsl-read   Read TSL")
	fmt.Fprintln(w, "  browse     Browse API endpoints (default)")
	fmt.Fprintln(w, "  help       Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query product schema")
	fmt.Fprintln(w, "  ur schema get-list -p p_smartswitch_001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Query device schema")
	fmt.Fprintln(w, "  ur schema get-list -p p_smartswitch_001 -d switch-001")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Browse API endpoints")
	fmt.Fprintln(w, "  ur schema browse /api/v1/things/device")
}

// runSchemaBrowse 执行 swagger 浏览命令
func runSchemaBrowse(app config.CLIApp, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	authType := ""
	targetPath := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		case "--auth-type":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--auth-type requires value")
				return 2
			}
			authType = args[i+1]
			i++
		default:
			if strings.HasPrefix(args[i], "-") {
				fmt.Fprintf(stderr, "unknown schema browse option: %s\n", args[i])
				return 2
			}
			targetPath = args[i]
		}
	}
	endpoints, err := swagger.LoadEndpoints()
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}

	// 按路径过滤
	if targetPath != "" {
		endpoints = swagger.FilterEndpoints(endpoints, targetPath, "")
	}

	// 按 authType 过滤：显式指定则用显式值，否则按应用默认
	if authType != "" {
		endpoints = swagger.FilterEndpoints(endpoints, "", authType)
	} else {
		endpoints = swagger.FilterEndpointsByApp(endpoints, app.AllowedAuthTypes())
	}

	if jsonOutput {
		raw, _ := json.MarshalIndent(endpoints, "", "  ")
		_, _ = fmt.Fprintln(stdout, string(raw))
		return 0
	}
	for _, item := range endpoints {
		summary := item.Summary
		if summary == "" {
			summary = item.Description
		}
		fmt.Fprintf(stdout, "%s %-6s [%s] %s\n", item.Path, item.Method, item.AuthType, summary)
	}
	return 0
}

// parseSchemaModelParams 解析物模型管理通用参数
func parseSchemaModelParams(args []string) (productID, deviceName string, jsonOutput bool, remaining []string, err error) {
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
	return
}

// runSchemaModel 执行物模型管理命令
func runSchemaModel(subCmd string, args []string, stdout, stderr io.Writer) int {
	ctx := context.Background()
	switch subCmd {
	case "get-list":
		return runSchemaModelGetList(ctx, args, stdout, stderr)
	case "create":
		return runSchemaModelCreate(ctx, args, stdout, stderr)
	case "update":
		return runSchemaModelUpdate(ctx, args, stdout, stderr)
	case "delete":
		return runSchemaModelDelete(ctx, args, stdout, stderr)
	case "tsl-import":
		return runSchemaModelTslImport(ctx, args, stdout, stderr)
	case "tsl-read":
		return runSchemaModelTslRead(ctx, args, stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown schema model subcommand: %s\n", subCmd)
		return 2
	}
}

// runSchemaModelGetList 执行查询物模型列表命令
func runSchemaModelGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, _, err := parseSchemaModelParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID": productID,
	}
	if deviceName != "" {
		reqBody["deviceName"] = deviceName
	}

	// 根据是否指定设备选择 API
	apiPath := "/api/v1/things/product/schema/get-list"
	if deviceName != "" {
		apiPath = "/api/v1/things/device/schema/get-list"
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: apiPath,
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runSchemaModelCreate 执行创建物模型命令
func runSchemaModelCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseSchemaModelParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	schemaJSON := ""
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--schema":
			if i+1 < len(remaining) {
				schemaJSON = remaining[i+1]
				i++
			}
		}
	}

	if schemaJSON == "" {
		fmt.Fprintln(stderr, "--schema is required")
		return 2
	}

	var schemaData any
	if err := json.Unmarshal([]byte(schemaJSON), &schemaData); err != nil {
		fmt.Fprintf(stderr, "Error: --schema must be JSON: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID": productID,
		"schema":    schemaData,
	}
	if deviceName != "" {
		reqBody["deviceName"] = deviceName
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/product/schema/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runSchemaModelUpdate 执行更新物模型命令
func runSchemaModelUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, remaining, err := parseSchemaModelParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	schemaJSON := ""
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--schema":
			if i+1 < len(remaining) {
				schemaJSON = remaining[i+1]
				i++
			}
		}
	}

	if schemaJSON == "" {
		fmt.Fprintln(stderr, "--schema is required")
		return 2
	}

	var schemaData any
	if err := json.Unmarshal([]byte(schemaJSON), &schemaData); err != nil {
		fmt.Fprintf(stderr, "Error: --schema must be JSON: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID": productID,
		"schema":    schemaData,
	}
	if deviceName != "" {
		reqBody["deviceName"] = deviceName
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/product/schema/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runSchemaModelDelete 执行删除物模型命令
func runSchemaModelDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, _, jsonOutput, remaining, err := parseSchemaModelParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	dataID := ""
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--data-id":
			if i+1 < len(remaining) {
				dataID = remaining[i+1]
				i++
			}
		}
	}

	if dataID == "" {
		fmt.Fprintln(stderr, "--data-id is required")
		return 2
	}

	reqBody := map[string]any{
		"productID": productID,
		"dataID":    dataID,
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/product/schema/delete",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runSchemaModelTslImport 执行导入 TSL 命令
func runSchemaModelTslImport(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, _, jsonOutput, remaining, err := parseSchemaModelParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	tslJSON := ""
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--tsl":
			if i+1 < len(remaining) {
				tslJSON = remaining[i+1]
				i++
			}
		}
	}

	if tslJSON == "" {
		fmt.Fprintln(stderr, "--tsl is required")
		return 2
	}

	var tslData any
	if err := json.Unmarshal([]byte(tslJSON), &tslData); err != nil {
		fmt.Fprintf(stderr, "Error: --tsl must be JSON: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID": productID,
		"tsl":       tslData,
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/product/schema/tsl-import",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runSchemaModelTslRead 执行读取 TSL 命令
func runSchemaModelTslRead(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	productID, deviceName, jsonOutput, _, err := parseSchemaModelParams(args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	reqBody := map[string]any{
		"productID": productID,
	}
	if deviceName != "" {
		reqBody["deviceName"] = deviceName
	}

	// 根据是否指定设备选择 API
	apiPath := "/api/v1/things/product/schema/tsl-read"
	if deviceName != "" {
		apiPath = "/api/v1/things/device/schema/tsl-read"
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: apiPath,
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
