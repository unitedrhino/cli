package shared

import (
	"context"
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runUser 执行用户管理命令
func runUser(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserHelp(stdout)
		return 0
	}

	switch args[0] {
	case "info":
		return runUserInfo(ctx, args[1:], stdout, stderr)
	case "self":
		return runUserSelf(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown user subcommand: %s\n", args[0])
		printUserHelp(stderr)
		return 2
	}
}

// printUserHelp 打印用户管理帮助信息
func printUserHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "User management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  info       User info management (get-list, get-one, create, update, delete)")
	fmt.Fprintln(w, "  self       User self management (login, logout, register, profile, etc.)")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runUserInfo 执行用户信息管理命令
func runUserInfo(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserInfoHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserInfoGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runUserInfoGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runUserInfoCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runUserInfoUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runUserInfoDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserInfoHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown user info subcommand: %s\n", args[0])
		printUserInfoHelp(stderr)
		return 2
	}
}

// printUserInfoHelp 打印用户信息帮助信息
func printUserInfoHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user info <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "User info management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query user list")
	fmt.Fprintln(w, "  get-one    Query user detail")
	fmt.Fprintln(w, "  create     Create user")
	fmt.Fprintln(w, "  update     Update user")
	fmt.Fprintln(w, "  delete     Delete user")
}

// runUserInfoGetList 执行查询用户列表命令
func runUserInfoGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/info/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserInfoGetOne 执行查询用户详情命令
func runUserInfoGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/info/get-one",
		Body: map[string]any{"id": userID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserInfoCreate 执行创建用户命令
func runUserInfoCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/info/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserInfoUpdate 执行更新用户命令
func runUserInfoUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/info/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserInfoDelete 执行删除用户命令
func runUserInfoDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/info/delete",
		Body: map[string]any{"id": userID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelf 执行用户自助管理命令
func runUserSelf(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfHelp(stdout)
		return 0
	}

	switch args[0] {
	case "login":
		return runUserSelfLogin(ctx, args[1:], stdout, stderr)
	case "logout":
		return runUserSelfLogout(ctx, args[1:], stdout, stderr)
	case "register":
		return runUserSelfRegister(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runUserSelfGetOne(ctx, args[1:], stdout, stderr)
	case "update":
		return runUserSelfUpdate(ctx, args[1:], stdout, stderr)
	case "change-pwd":
		return runUserSelfChangePwd(ctx, args[1:], stdout, stderr)
	case "forget-pwd":
		return runUserSelfForgetPwd(ctx, args[1:], stdout, stderr)
	case "captcha":
		return runUserSelfCaptcha(ctx, args[1:], stdout, stderr)
	case "access-token":
		return runUserSelfAccessToken(ctx, args[1:], stdout, stderr)
	case "message":
		return runUserSelfMessage(ctx, args[1:], stdout, stderr)
	case "tenant":
		return runUserSelfTenant(ctx, args[1:], stdout, stderr)
	case "profile":
		return runUserSelfProfile(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown user self subcommand: %s\n", args[0])
		printUserSelfHelp(stderr)
		return 2
	}
}

// printUserSelfHelp 打印用户自助帮助信息
func printUserSelfHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "User self management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  login          User login")
	fmt.Fprintln(w, "  logout         User logout")
	fmt.Fprintln(w, "  register       User register")
	fmt.Fprintln(w, "  get-one        Query current user info")
	fmt.Fprintln(w, "  update         Update current user info")
	fmt.Fprintln(w, "  change-pwd     Change password")
	fmt.Fprintln(w, "  forget-pwd     Forget password")
	fmt.Fprintln(w, "  captcha        Get captcha")
	fmt.Fprintln(w, "  access-token   Access token management")
	fmt.Fprintln(w, "  message        Message management")
	fmt.Fprintln(w, "  tenant         Tenant management")
	fmt.Fprintln(w, "  profile        Profile management")
	fmt.Fprintln(w, "  help           Show this help message")
}

// runUserSelfLogin 执行用户登录命令
func runUserSelfLogin(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/login",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfLogout 执行用户登出命令
func runUserSelfLogout(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/logout",
		Body: map[string]any{},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfRegister 执行用户注册命令
func runUserSelfRegister(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/register",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfGetOne 执行查询当前用户信息命令
func runUserSelfGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/get-one",
		Body: map[string]any{},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfUpdate 执行更新当前用户信息命令
func runUserSelfUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfChangePwd 执行修改密码命令
func runUserSelfChangePwd(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/change-pwd",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfForgetPwd 执行忘记密码命令
func runUserSelfForgetPwd(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/forget-pwd",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfCaptcha 执行获取验证码命令
func runUserSelfCaptcha(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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

	reqBody := map[string]any{}
	if bodyJSON != "" {
		var err error
		reqBody, err = parseBodyArg(bodyJSON)
		if err != nil {
			fmt.Fprintf(stderr, "Error: %v\n", err)
			return 2
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/captcha",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfAccessToken 执行访问令牌管理命令
func runUserSelfAccessToken(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfAccessTokenHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserSelfAccessTokenGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runUserSelfAccessTokenGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runUserSelfAccessTokenCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runUserSelfAccessTokenUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runUserSelfAccessTokenDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfAccessTokenHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown access token subcommand: %s\n", args[0])
		printUserSelfAccessTokenHelp(stderr)
		return 2
	}
}

// printUserSelfAccessTokenHelp 打印访问令牌帮助信息
func printUserSelfAccessTokenHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self access-token <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Access token management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query access token list")
	fmt.Fprintln(w, "  get-one    Query access token detail")
	fmt.Fprintln(w, "  create     Create access token")
	fmt.Fprintln(w, "  update     Update access token")
	fmt.Fprintln(w, "  delete     Delete access token")
}

// runUserSelfAccessTokenGetList 执行查询访问令牌列表命令
func runUserSelfAccessTokenGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/access-token/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfAccessTokenGetOne 执行查询访问令牌详情命令
func runUserSelfAccessTokenGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	tokenID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				tokenID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if tokenID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/access-token/get-one",
		Body: map[string]any{"id": tokenID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfAccessTokenCreate 执行创建访问令牌命令
func runUserSelfAccessTokenCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/access-token/create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfAccessTokenUpdate 执行更新访问令牌命令
func runUserSelfAccessTokenUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/access-token/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfAccessTokenDelete 执行删除访问令牌命令
func runUserSelfAccessTokenDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	tokenID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				tokenID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if tokenID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/access-token/delete",
		Body: map[string]any{"id": tokenID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfMessage 执行消息管理命令
func runUserSelfMessage(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfMessageHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserSelfMessageGetList(ctx, args[1:], stdout, stderr)
	case "get-pending":
		return runUserSelfMessageGetPending(ctx, args[1:], stdout, stderr)
	case "handle":
		return runUserSelfMessageHandle(ctx, args[1:], stdout, stderr)
	case "mark-all-read":
		return runUserSelfMessageMarkAllRead(ctx, args[1:], stdout, stderr)
	case "multi-delete":
		return runUserSelfMessageMultiDelete(ctx, args[1:], stdout, stderr)
	case "multi-is-read":
		return runUserSelfMessageMultiIsRead(ctx, args[1:], stdout, stderr)
	case "statistics":
		return runUserSelfMessageStatistics(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfMessageHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown message subcommand: %s\n", args[0])
		printUserSelfMessageHelp(stderr)
		return 2
	}
}

// printUserSelfMessageHelp 打印消息管理帮助信息
func printUserSelfMessageHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self message <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Message management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list       Query message list")
	fmt.Fprintln(w, "  get-pending    Query pending messages")
	fmt.Fprintln(w, "  handle         Handle message")
	fmt.Fprintln(w, "  mark-all-read  Mark all messages as read")
	fmt.Fprintln(w, "  multi-delete   Batch delete messages")
	fmt.Fprintln(w, "  multi-is-read  Batch mark messages as read")
	fmt.Fprintln(w, "  statistics     Message statistics")
}

// runUserSelfMessageGetList 执行查询消息列表命令
func runUserSelfMessageGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/message/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfMessageGetPending 执行查询待处理消息命令
func runUserSelfMessageGetPending(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/message/get-pending",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfMessageHandle 执行处理消息命令
func runUserSelfMessageHandle(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/message/handle",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfMessageMarkAllRead 执行标记所有消息已读命令
func runUserSelfMessageMarkAllRead(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/message/mark-all-read",
		Body: map[string]any{},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfMessageMultiDelete 执行批量删除消息命令
func runUserSelfMessageMultiDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/message/multi-delete",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfMessageMultiIsRead 执行批量标记消息已读命令
func runUserSelfMessageMultiIsRead(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/message/multi-is-read",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfMessageStatistics 执行消息统计命令
func runUserSelfMessageStatistics(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/message/statistics",
		Body: map[string]any{},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfTenant 执行租户管理命令
func runUserSelfTenant(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfTenantHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserSelfTenantGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runUserSelfTenantGetOne(ctx, args[1:], stdout, stderr)
	case "update":
		return runUserSelfTenantUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runUserSelfTenantDelete(ctx, args[1:], stdout, stderr)
	case "join":
		return runUserSelfTenantJoin(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfTenantHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown tenant subcommand: %s\n", args[0])
		printUserSelfTenantHelp(stderr)
		return 2
	}
}

// printUserSelfTenantHelp 打印租户管理帮助信息
func printUserSelfTenantHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self tenant <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Tenant management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query tenant list")
	fmt.Fprintln(w, "  get-one    Query tenant detail")
	fmt.Fprintln(w, "  update     Update tenant")
	fmt.Fprintln(w, "  delete     Delete tenant")
	fmt.Fprintln(w, "  join       Join tenant")
}

// runUserSelfTenantGetList 执行查询租户列表命令
func runUserSelfTenantGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/tenant/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfTenantGetOne 执行查询租户详情命令
func runUserSelfTenantGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	tenantID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				tenantID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if tenantID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/tenant/get-one",
		Body: map[string]any{"id": tenantID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfTenantUpdate 执行更新租户命令
func runUserSelfTenantUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/tenant/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfTenantDelete 执行删除租户命令
func runUserSelfTenantDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	tenantID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				tenantID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if tenantID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/tenant/delete",
		Body: map[string]any{"id": tenantID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfTenantJoin 执行加入租户命令
func runUserSelfTenantJoin(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/tenant/join",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfProfile 执行配置管理命令
func runUserSelfProfile(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfProfileHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserSelfProfileGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runUserSelfProfileGetOne(ctx, args[1:], stdout, stderr)
	case "update":
		return runUserSelfProfileUpdate(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfProfileHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown profile subcommand: %s\n", args[0])
		printUserSelfProfileHelp(stderr)
		return 2
	}
}

// printUserSelfProfileHelp 打印配置管理帮助信息
func printUserSelfProfileHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self profile <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Profile management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query profile list")
	fmt.Fprintln(w, "  get-one    Query profile detail")
	fmt.Fprintln(w, "  update     Update profile")
}

// runUserSelfProfileGetList 执行查询配置列表命令
func runUserSelfProfileGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/profile/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfProfileGetOne 执行查询配置详情命令
func runUserSelfProfileGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	profileID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				profileID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if profileID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/profile/get-one",
		Body: map[string]any{"id": profileID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfProfileUpdate 执行更新配置命令
func runUserSelfProfileUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/profile/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
