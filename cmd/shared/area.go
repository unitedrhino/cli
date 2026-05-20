package shared

import (
	"context"
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runArea 执行区域管理命令
func runArea(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAreaHelp(stdout)
		return 0
	}

	switch args[0] {
	case "info":
		return runAreaInfo(ctx, args[1:], stdout, stderr)
	case "profile":
		return runAreaProfile(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAreaHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown area subcommand: %s\n", args[0])
		printAreaHelp(stderr)
		return 2
	}
}

// printAreaHelp 打印区域管理帮助信息
func printAreaHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur area <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Area management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  info       Area info management (get-list, get-one, create, update, delete)")
	fmt.Fprintln(w, "  profile    Area profile management (get-list, get-one, update)")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runAreaInfo 执行区域信息管理命令
func runAreaInfo(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAreaInfoHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runAreaInfoGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runAreaInfoGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runAreaInfoCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runAreaInfoUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runAreaInfoDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAreaInfoHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown area info subcommand: %s\n", args[0])
		printAreaInfoHelp(stderr)
		return 2
	}
}

// printAreaInfoHelp 打印区域信息帮助信息
func printAreaInfoHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur area info <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Area info management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query area list")
	fmt.Fprintln(w, "  get-one    Query area detail")
	fmt.Fprintln(w, "  create     Create area")
	fmt.Fprintln(w, "  update     Update area")
	fmt.Fprintln(w, "  delete     Delete area")
}

// runAreaInfoGetList 执行查询区域列表命令
func runAreaInfoGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		case "--parent-id":
			if i+1 < len(remaining) {
				reqBody["parentID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/area/info/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAreaInfoGetOne 执行查询区域详情命令
func runAreaInfoGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	areaID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				areaID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if areaID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/area/info/get-one",
		Body: map[string]any{"id": areaID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAreaInfoCreate 执行创建区域命令
func runAreaInfoCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/area/info/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAreaInfoUpdate 执行更新区域命令
func runAreaInfoUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/area/info/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAreaInfoDelete 执行删除区域命令
func runAreaInfoDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	areaID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				areaID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if areaID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/area/info/delete",
		Body: map[string]any{"id": areaID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAreaProfile 执行区域配置管理命令
func runAreaProfile(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAreaProfileHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runAreaProfileGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runAreaProfileGetOne(ctx, args[1:], stdout, stderr)
	case "update":
		return runAreaProfileUpdate(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAreaProfileHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown area profile subcommand: %s\n", args[0])
		printAreaProfileHelp(stderr)
		return 2
	}
}

// printAreaProfileHelp 打印区域配置帮助信息
func printAreaProfileHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur area profile <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Area profile management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query area profile list")
	fmt.Fprintln(w, "  get-one    Query area profile detail")
	fmt.Fprintln(w, "  update     Update area profile")
}

// runAreaProfileGetList 执行查询区域配置列表命令
func runAreaProfileGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/area/profile/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAreaProfileGetOne 执行查询区域配置详情命令
func runAreaProfileGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	areaID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				areaID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if areaID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/area/profile/get-one",
		Body: map[string]any{"id": areaID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runAreaProfileUpdate 执行更新区域配置命令
func runAreaProfileUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/area/profile/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
