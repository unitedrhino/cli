package shared

import (
	"context"
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runTenant 执行租户用户管理命令
func runTenant(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printTenantHelp(stdout)
		return 0
	}

	switch args[0] {
	case "user":
		return runTenantUser(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printTenantHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown tenant subcommand: %s\n", args[0])
		printTenantHelp(stderr)
		return 2
	}
}

// printTenantHelp 打印租户管理帮助信息
func printTenantHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur tenant <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Tenant user management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  user       Tenant user management (get-list, get-one, batch-create, update, delete, invite, role)")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runTenantUser 执行租户用户管理命令
func runTenantUser(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printTenantUserHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runTenantUserGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runTenantUserGetOne(ctx, args[1:], stdout, stderr)
	case "batch-create":
		return runTenantUserBatchCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runTenantUserUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runTenantUserDelete(ctx, args[1:], stdout, stderr)
	case "invite":
		return runTenantUserInvite(ctx, args[1:], stdout, stderr)
	case "invite-send":
		return runTenantUserInviteSend(ctx, args[1:], stdout, stderr)
	case "invite-code":
		return runTenantUserInviteCode(ctx, args[1:], stdout, stderr)
	case "invite-pending":
		return runTenantUserInvitePending(ctx, args[1:], stdout, stderr)
	case "role":
		return runTenantUserRole(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printTenantUserHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown tenant user subcommand: %s\n", args[0])
		printTenantUserHelp(stderr)
		return 2
	}
}

// printTenantUserHelp 打印租户用户帮助信息
func printTenantUserHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur tenant user <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Tenant user management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list         Query tenant user list")
	fmt.Fprintln(w, "  get-one          Query tenant user detail")
	fmt.Fprintln(w, "  batch-create     Batch create tenant users")
	fmt.Fprintln(w, "  update           Update tenant user")
	fmt.Fprintln(w, "  delete           Delete tenant user")
	fmt.Fprintln(w, "  invite           Invite user")
	fmt.Fprintln(w, "  invite-send      Send invite")
	fmt.Fprintln(w, "  invite-code      Invite code management")
	fmt.Fprintln(w, "  invite-pending   Pending invite management")
	fmt.Fprintln(w, "  role             User role management")
}

// runTenantUserGetList 执行查询租户用户列表命令
func runTenantUserGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--account":
			if i+1 < len(remaining) {
				reqBody["account"] = remaining[i+1]
				i++
			}
		case "--status":
			if i+1 < len(remaining) {
				reqBody["status"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/tenant/user/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserGetOne 执行查询租户用户详情命令
func runTenantUserGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	userID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				userID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if userID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/tenant/user/get-one",
		Body: map[string]any{"id": userID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserBatchCreate 执行批量创建租户用户命令
func runTenantUserBatchCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/tenant/user/batch-create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserUpdate 执行更新租户用户命令
func runTenantUserUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/tenant/user/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserDelete 执行删除租户用户命令
func runTenantUserDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	userID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				userID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if userID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/tenant/user/delete",
		Body: map[string]any{"id": userID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserInvite 执行邀请用户命令
func runTenantUserInvite(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/tenant/user/invite",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserInviteSend 执行发送邀请命令
func runTenantUserInviteSend(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/tenant/user/invite-send",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserInviteCode 执行邀请码管理命令
func runTenantUserInviteCode(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printTenantUserInviteCodeHelp(stdout)
		return 0
	}

	switch args[0] {
	case "gen":
		return runTenantUserInviteCodeGen(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runTenantUserInviteCodeGetOne(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printTenantUserInviteCodeHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown invite code subcommand: %s\n", args[0])
		printTenantUserInviteCodeHelp(stderr)
		return 2
	}
}

// printTenantUserInviteCodeHelp 打印邀请码帮助信息
func printTenantUserInviteCodeHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur tenant user invite-code <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Invite code management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  gen        Generate invite code")
	fmt.Fprintln(w, "  get-one    Query invite code detail")
}

// runTenantUserInviteCodeGen 执行生成邀请码命令
func runTenantUserInviteCodeGen(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/tenant/user/invite-code/gen",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserInviteCodeGetOne 执行查询邀请码详情命令
func runTenantUserInviteCodeGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	codeID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				codeID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if codeID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/tenant/user/invite-code/get-one",
		Body: map[string]any{"id": codeID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserInvitePending 执行待处理邀请管理命令
func runTenantUserInvitePending(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printTenantUserInvitePendingHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runTenantUserInvitePendingGetList(ctx, args[1:], stdout, stderr)
	case "delete":
		return runTenantUserInvitePendingDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printTenantUserInvitePendingHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown invite pending subcommand: %s\n", args[0])
		printTenantUserInvitePendingHelp(stderr)
		return 2
	}
}

// printTenantUserInvitePendingHelp 打印待处理邀请帮助信息
func printTenantUserInvitePendingHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur tenant user invite-pending <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Pending invite management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query pending invite list")
	fmt.Fprintln(w, "  delete     Delete pending invite")
}

// runTenantUserInvitePendingGetList 执行查询待处理邀请列表命令
func runTenantUserInvitePendingGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/tenant/user/invite-pending/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserInvitePendingDelete 执行删除待处理邀请命令
func runTenantUserInvitePendingDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	inviteID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				inviteID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if inviteID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/tenant/user/invite-pending/delete",
		Body: map[string]any{"id": inviteID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserRole 执行用户角色管理命令
func runTenantUserRole(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printTenantUserRoleHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runTenantUserRoleGetList(ctx, args[1:], stdout, stderr)
	case "batch-update":
		return runTenantUserRoleBatchUpdate(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printTenantUserRoleHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown role subcommand: %s\n", args[0])
		printTenantUserRoleHelp(stderr)
		return 2
	}
}

// printTenantUserRoleHelp 打印用户角色帮助信息
func printTenantUserRoleHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur tenant user role <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "User role management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list       Query user role list")
	fmt.Fprintln(w, "  batch-update   Batch update user roles")
}

// runTenantUserRoleGetList 执行查询用户角色列表命令
func runTenantUserRoleGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--user-id":
			if i+1 < len(remaining) {
				reqBody["userID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/tenant/user/role/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runTenantUserRoleBatchUpdate 执行批量更新用户角色命令
func runTenantUserRoleBatchUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/tenant/user/role/batch-update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
