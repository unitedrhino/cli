package shared

import (
	"context"
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runProject 执行项目管理命令
func runProject(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printProjectHelp(stdout)
		return 0
	}

	switch args[0] {
	case "info":
		return runProjectInfo(ctx, args[1:], stdout, stderr)
	case "profile":
		return runProjectProfile(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printProjectHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown project subcommand: %s\n", args[0])
		printProjectHelp(stderr)
		return 2
	}
}

// printProjectHelp 打印项目管理帮助信息
func printProjectHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur project <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Project management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  info       Project info management (get-list, get-one, create, update, delete)")
	fmt.Fprintln(w, "  profile    Project profile management (get-list, get-one, update)")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runProjectInfo 执行项目信息管理命令
func runProjectInfo(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printProjectInfoHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runProjectInfoGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runProjectInfoGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runProjectInfoCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runProjectInfoUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runProjectInfoDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printProjectInfoHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown project info subcommand: %s\n", args[0])
		printProjectInfoHelp(stderr)
		return 2
	}
}

// printProjectInfoHelp 打印项目信息帮助信息
func printProjectInfoHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur project info <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Project info management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query project list")
	fmt.Fprintln(w, "  get-one    Query project detail")
	fmt.Fprintln(w, "  create     Create project")
	fmt.Fprintln(w, "  update     Update project")
	fmt.Fprintln(w, "  delete     Delete project")
}

// runProjectInfoGetList 执行查询项目列表命令
func runProjectInfoGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		case "--area-id":
			if i+1 < len(remaining) {
				reqBody["areaID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/project/info/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runProjectInfoGetOne 执行查询项目详情命令
func runProjectInfoGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	projectID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if projectID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/project/info/get-one",
		Body: map[string]any{"id": projectID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runProjectInfoCreate 执行创建项目命令
func runProjectInfoCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/project/info/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runProjectInfoUpdate 执行更新项目命令
func runProjectInfoUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/project/info/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runProjectInfoDelete 执行删除项目命令
func runProjectInfoDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	projectID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if projectID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/project/info/delete",
		Body: map[string]any{"id": projectID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runProjectProfile 执行项目配置管理命令
func runProjectProfile(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printProjectProfileHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runProjectProfileGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runProjectProfileGetOne(ctx, args[1:], stdout, stderr)
	case "update":
		return runProjectProfileUpdate(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printProjectProfileHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown project profile subcommand: %s\n", args[0])
		printProjectProfileHelp(stderr)
		return 2
	}
}

// printProjectProfileHelp 打印项目配置帮助信息
func printProjectProfileHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur project profile <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Project profile management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query project profile list")
	fmt.Fprintln(w, "  get-one    Query project profile detail")
	fmt.Fprintln(w, "  update     Update project profile")
}

// runProjectProfileGetList 执行查询项目配置列表命令
func runProjectProfileGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/project/profile/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runProjectProfileGetOne 执行查询项目配置详情命令
func runProjectProfileGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	projectID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if projectID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/things/project/profile/get-one",
		Body: map[string]any{"id": projectID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runProjectProfileUpdate 执行更新项目配置命令
func runProjectProfileUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/things/project/profile/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
