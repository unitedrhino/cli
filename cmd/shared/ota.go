package shared

import (
	"context"
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runOta 执行 OTA 管理命令
func runOta(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printOtaHelp(stdout)
		return 0
	}

	switch args[0] {
	case "firmware":
		return runOtaFirmware(ctx, args[1:], stdout, stderr)
	case "job":
		return runOtaJob(ctx, args[1:], stdout, stderr)
	case "module":
		return runOtaModule(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printOtaHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown ota subcommand: %s\n", args[0])
		printOtaHelp(stderr)
		return 2
	}
}

// printOtaHelp 打印 OTA 管理帮助信息
func printOtaHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur ota <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "OTA (Over-The-Air) management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  firmware   Firmware management (get-list, get-one, create, update, delete)")
	fmt.Fprintln(w, "  job        OTA job management (get-list, get-one, create, update)")
	fmt.Fprintln(w, "  module     Module management (get-list, get-one, create, update, delete)")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runOtaFirmware 执行固件管理命令
func runOtaFirmware(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printOtaFirmwareHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runOtaFirmwareGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runOtaFirmwareGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runOtaFirmwareCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runOtaFirmwareUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runOtaFirmwareDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printOtaFirmwareHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown ota firmware subcommand: %s\n", args[0])
		printOtaFirmwareHelp(stderr)
		return 2
	}
}

// printOtaFirmwareHelp 打印固件管理帮助信息
func printOtaFirmwareHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur ota firmware <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Firmware management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query firmware list")
	fmt.Fprintln(w, "  get-one    Query firmware detail")
	fmt.Fprintln(w, "  create     Create firmware")
	fmt.Fprintln(w, "  update     Update firmware")
	fmt.Fprintln(w, "  delete     Delete firmware")
}

// runOtaFirmwareGetList 执行查询固件列表命令
func runOtaFirmwareGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		case "--version":
			if i+1 < len(remaining) {
				reqBody["version"] = remaining[i+1]
				i++
			}
		case "--module":
			if i+1 < len(remaining) {
				reqBody["module"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/ota/firmware/info/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaFirmwareGetOne 执行查询固件详情命令
func runOtaFirmwareGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	firmwareID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				firmwareID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if firmwareID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/ota/firmware/info/get-one",
		Body: map[string]any{"id": firmwareID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaFirmwareCreate 执行创建固件命令
func runOtaFirmwareCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/ota/firmware/info/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaFirmwareUpdate 执行更新固件命令
func runOtaFirmwareUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/ota/firmware/info/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaFirmwareDelete 执行删除固件命令
func runOtaFirmwareDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	firmwareID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				firmwareID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if firmwareID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/ota/firmware/info/delete",
		Body: map[string]any{"id": firmwareID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaJob 执行 OTA 任务管理命令
func runOtaJob(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printOtaJobHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runOtaJobGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runOtaJobGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runOtaJobCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runOtaJobUpdate(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printOtaJobHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown ota job subcommand: %s\n", args[0])
		printOtaJobHelp(stderr)
		return 2
	}
}

// printOtaJobHelp 打印 OTA 任务帮助信息
func printOtaJobHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur ota job <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "OTA job management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query OTA job list")
	fmt.Fprintln(w, "  get-one    Query OTA job detail")
	fmt.Fprintln(w, "  create     Create OTA job")
	fmt.Fprintln(w, "  update     Update OTA job")
}

// runOtaJobGetList 执行查询 OTA 任务列表命令
func runOtaJobGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		case "--firmware-id":
			if i+1 < len(remaining) {
				reqBody["firmwareID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/ota/firmware/job/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaJobGetOne 执行查询 OTA 任务详情命令
func runOtaJobGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	jobID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				jobID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if jobID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/ota/firmware/job/get-one",
		Body: map[string]any{"id": jobID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaJobCreate 执行创建 OTA 任务命令
func runOtaJobCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/ota/firmware/job/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaJobUpdate 执行更新 OTA 任务命令
func runOtaJobUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/ota/firmware/job/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaModule 执行模块管理命令
func runOtaModule(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printOtaModuleHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runOtaModuleGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runOtaModuleGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runOtaModuleCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runOtaModuleUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runOtaModuleDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printOtaModuleHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown ota module subcommand: %s\n", args[0])
		printOtaModuleHelp(stderr)
		return 2
	}
}

// printOtaModuleHelp 打印模块管理帮助信息
func printOtaModuleHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur ota module <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Module management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query module list")
	fmt.Fprintln(w, "  get-one    Query module detail")
	fmt.Fprintln(w, "  create     Create module")
	fmt.Fprintln(w, "  update     Update module")
	fmt.Fprintln(w, "  delete     Delete module")
}

// runOtaModuleGetList 执行查询模块列表命令
func runOtaModuleGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/ota/module/info/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaModuleGetOne 执行查询模块详情命令
func runOtaModuleGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	moduleID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				moduleID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if moduleID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/ota/module/info/get-one",
		Body: map[string]any{"id": moduleID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaModuleCreate 执行创建模块命令
func runOtaModuleCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/ota/module/info/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaModuleUpdate 执行更新模块命令
func runOtaModuleUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/ota/module/info/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runOtaModuleDelete 执行删除模块命令
func runOtaModuleDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	moduleID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				moduleID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if moduleID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/ota/module/info/delete",
		Body: map[string]any{"id": moduleID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
