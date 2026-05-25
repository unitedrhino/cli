package shared

import (
	"context"
	"encoding/json"
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
	case "group":
		return runProjectGroup(ctx, args[1:], stdout, stderr)
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
	fmt.Fprintln(w, "  group      Device group management (info get-list/create/delete, device batch-create)")
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

// runProjectGroup 执行项目设备分组管理命令
func runProjectGroup(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printProjectGroupHelp(stdout)
		return 0
	}

	switch args[0] {
	case "info":
		return runProjectGroupInfo(ctx, args[1:], stdout, stderr)
	case "device":
		return runProjectGroupDevice(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printProjectGroupHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown project group subcommand: %s\n", args[0])
		printProjectGroupHelp(stderr)
		return 2
	}
}

// printProjectGroupHelp 打印项目设备分组管理帮助信息
func printProjectGroupHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur project group <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Project device group management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  info       Group info management (get-list, create, delete)")
	fmt.Fprintln(w, "  device     Group device management (batch-create)")
	fmt.Fprintln(w, "  help       Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Query project groups")
	fmt.Fprintln(w, "  ur project group info get-list")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Create project group")
	fmt.Fprintln(w, "  ur project group info create --name \"My Group\" --desc \"Description\"")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Batch add devices to group")
	fmt.Fprintln(w, "  ur project group device batch-create --group-id 123 --devices '[{\"productID\":\"pid\",\"deviceName\":\"dev1\"}]'")
}

// runProjectGroupInfo 执行项目分组信息管理命令
func runProjectGroupInfo(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printProjectGroupInfoHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runProjectGroupInfoGetList(ctx, args[1:], stdout, stderr)
	case "create":
		return runProjectGroupInfoCreate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runProjectGroupInfoDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printProjectGroupInfoHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown project group info subcommand: %s\n", args[0])
		printProjectGroupInfoHelp(stderr)
		return 2
	}
}

// printProjectGroupInfoHelp 打印项目分组信息管理帮助信息
func printProjectGroupInfoHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur project group info <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Project group info management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query project group list")
	fmt.Fprintln(w, "  create     Create project group")
	fmt.Fprintln(w, "  delete     Delete project group")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runProjectGroupInfoGetList 执行查询项目分组列表命令
func runProjectGroupInfoGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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

// runProjectGroupInfoCreate 执行创建项目分组命令
func runProjectGroupInfoCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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

// runProjectGroupInfoDelete 执行删除项目分组命令
func runProjectGroupInfoDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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

// runProjectGroupDevice 执行项目分组设备管理命令
func runProjectGroupDevice(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printProjectGroupDeviceHelp(stdout)
		return 0
	}

	switch args[0] {
	case "batch-create":
		return runProjectGroupDeviceBatchCreate(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printProjectGroupDeviceHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown project group device subcommand: %s\n", args[0])
		printProjectGroupDeviceHelp(stderr)
		return 2
	}
}

// printProjectGroupDeviceHelp 打印项目分组设备管理帮助信息
func printProjectGroupDeviceHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur project group device <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Project group device management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  batch-create   Batch add devices to group")
	fmt.Fprintln(w, "  help           Show this help message")
}

// runProjectGroupDeviceBatchCreate 执行批量添加设备到项目分组命令
func runProjectGroupDeviceBatchCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
