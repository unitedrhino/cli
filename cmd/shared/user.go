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
	case "data":
		return runUserData(ctx, args[1:], stdout, stderr)
	case "dept":
		return runUserDept(ctx, args[1:], stdout, stderr)
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
	fmt.Fprintln(w, "  data       User data permissions (project, area)")
	fmt.Fprintln(w, "  dept       User department management (batch-create, batch-delete)")
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
	case "cancel":
		return runUserSelfCancel(ctx, args[1:], stdout, stderr)
	case "bind-account":
		return runUserSelfBindAccount(ctx, args[1:], stdout, stderr)
	case "user-search":
		return runUserSelfUserSearch(ctx, args[1:], stdout, stderr)
	case "third-auth":
		return runUserSelfThirdAuth(ctx, args[1:], stdout, stderr)
	case "third-login":
		return runUserSelfThirdLogin(ctx, args[1:], stdout, stderr)
	case "third-register":
		return runUserSelfThirdRegister(ctx, args[1:], stdout, stderr)
	case "app":
		return runUserSelfApp(ctx, args[1:], stdout, stderr)
	case "menu":
		return runUserSelfMenu(ctx, args[1:], stdout, stderr)
	case "notify-preference":
		return runUserSelfNotifyPreference(ctx, args[1:], stdout, stderr)
	case "openclaw":
		return runUserSelfOpenclaw(ctx, args[1:], stdout, stderr)
	case "resource":
		return runUserSelfResource(ctx, args[1:], stdout, stderr)
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
	fmt.Fprintln(w, "  login\t\tUser login")
	fmt.Fprintln(w, "  logout\t\tUser logout")
	fmt.Fprintln(w, "  register\t\tUser register")
	fmt.Fprintln(w, "  get-one\t\tQuery current user info")
	fmt.Fprintln(w, "  update\t\tUpdate current user info")
	fmt.Fprintln(w, "  change-pwd\t\tChange password")
	fmt.Fprintln(w, "  forget-pwd\t\tForget password")
	fmt.Fprintln(w, "  captcha\t\tGet captcha")
	fmt.Fprintln(w, "  access-token\t\tAccess token management")
	fmt.Fprintln(w, "  message\t\tMessage management")
	fmt.Fprintln(w, "  tenant\t\tTenant management")
	fmt.Fprintln(w, "  profile\t\tProfile management")
	fmt.Fprintln(w, "  cancel\t\tCancel account")
	fmt.Fprintln(w, "  bind-account\t\tBind account")
	fmt.Fprintln(w, "  user-search\t\tSearch user by account")
	fmt.Fprintln(w, "  third-auth\t\tThird-party auth (start)")
	fmt.Fprintln(w, "  third-login\t\tThird-party login")
	fmt.Fprintln(w, "  third-register\t\tThird-party register")
	fmt.Fprintln(w, "  app\t\tApp management (get-list, get-one)")
	fmt.Fprintln(w, "  menu\t\tMenu management (get-list)")
	fmt.Fprintln(w, "  notify-preference\tNotification preference (read, update)")
	fmt.Fprintln(w, "  openclaw\t\tOpenClaw setup (setup-check, setup-complete)")
	fmt.Fprintln(w, "  resource\t\tResource/action permissions (action)")
	fmt.Fprintln(w, "  help\t\tShow this help message")
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

// runUserSelfCancel 执行注销用户命令
func runUserSelfCancel(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/cancel",
		Body: map[string]any{},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfBindAccount 执行绑定账号命令
func runUserSelfBindAccount(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/bind-account",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfUserSearch 执行精准搜索用户命令
func runUserSelfUserSearch(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	account := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--account":
			if i+1 < len(args) {
				account = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		}
	}

	if account == "" {
		fmt.Fprintln(stderr, "--account is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/user/search",
		Body: map[string]any{"account": account},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfThirdAuth 执行第三方认证命令
func runUserSelfThirdAuth(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfThirdAuthHelp(stdout)
		return 0
	}

	switch args[0] {
	case "start":
		return runUserSelfThirdAuthStart(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfThirdAuthHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown third-auth subcommand: %s\n", args[0])
		printUserSelfThirdAuthHelp(stderr)
		return 2
	}
}

func printUserSelfThirdAuthHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self third-auth <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Third-party auth management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  start   Start third-party auth flow")
}

func runUserSelfThirdAuthStart(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/third-auth/start",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfThirdLogin 执行第三方登录命令
func runUserSelfThirdLogin(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/third-login",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfThirdRegister 执行第三方注册命令
func runUserSelfThirdRegister(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/third-register",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfApp 执行应用管理命令
func runUserSelfApp(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfAppHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserSelfAppGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runUserSelfAppGetOne(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfAppHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown app subcommand: %s\n", args[0])
		printUserSelfAppHelp(stderr)
		return 2
	}
}

func printUserSelfAppHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self app <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "App management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query app list")
	fmt.Fprintln(w, "  get-one    Query app detail")
}

func runUserSelfAppGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/app/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

func runUserSelfAppGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/app/get-one",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfMenu 执行菜单管理命令
func runUserSelfMenu(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfMenuHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserSelfMenuGetList(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfMenuHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown menu subcommand: %s\n", args[0])
		printUserSelfMenuHelp(stderr)
		return 2
	}
}

func printUserSelfMenuHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self menu <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Menu management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query menu list")
}

func runUserSelfMenuGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--app-id":
			if i+1 < len(remaining) {
				reqBody["appID"] = remaining[i+1]
				i++
			}
		case "--is-common":
			if i+1 < len(remaining) {
				reqBody["isCommon"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/menu/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfNotifyPreference 执行通知偏好管理命令
func runUserSelfNotifyPreference(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfNotifyPreferenceHelp(stdout)
		return 0
	}

	switch args[0] {
	case "read":
		return runUserSelfNotifyPreferenceRead(ctx, args[1:], stdout, stderr)
	case "update":
		return runUserSelfNotifyPreferenceUpdate(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfNotifyPreferenceHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown notify-preference subcommand: %s\n", args[0])
		printUserSelfNotifyPreferenceHelp(stderr)
		return 2
	}
}

func printUserSelfNotifyPreferenceHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self notify-preference <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Notification preference management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  read    Read notification preference")
	fmt.Fprintln(w, "  update  Update notification preference")
}

func runUserSelfNotifyPreferenceRead(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/notify-preference/read",
		Body: map[string]any{},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

func runUserSelfNotifyPreferenceUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/notify-preference/update",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfOpenclaw 执行 OpenClaw 管理命令
func runUserSelfOpenclaw(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfOpenclawHelp(stdout)
		return 0
	}

	switch args[0] {
	case "setup-check":
		return runUserSelfOpenclawSetupCheck(ctx, args[1:], stdout, stderr)
	case "setup-complete":
		return runUserSelfOpenclawSetupComplete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfOpenclawHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown openclaw subcommand: %s\n", args[0])
		printUserSelfOpenclawHelp(stderr)
		return 2
	}
}

func printUserSelfOpenclawHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self openclaw <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "OpenClaw setup management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  setup-check     Check CLI binding status")
	fmt.Fprintln(w, "  setup-complete  Complete CLI binding")
}

func runUserSelfOpenclawSetupCheck(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/openclaw/setup-check",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

func runUserSelfOpenclawSetupComplete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/self/openclaw/setup-complete",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserSelfResource 执行资源权限管理命令
func runUserSelfResource(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfResourceHelp(stdout)
		return 0
	}

	switch args[0] {
	case "action":
		return runUserSelfResourceAction(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfResourceHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown resource subcommand: %s\n", args[0])
		printUserSelfResourceHelp(stderr)
		return 2
	}
}

func printUserSelfResourceHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self resource <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Resource permission management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  action   Resource action permissions (get-list)")
}

func runUserSelfResourceAction(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserSelfResourceActionHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserSelfResourceActionGetList(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserSelfResourceActionHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown resource action subcommand: %s\n", args[0])
		return 2
	}
}

func printUserSelfResourceActionHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user self resource action <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Resource action permissions")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query resource action permission list")
}

func runUserSelfResourceActionGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/resource/action/get-list",
		Body: map[string]any{},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserData 执行用户数据权限管理命令
func runUserData(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserDataHelp(stdout)
		return 0
	}

	switch args[0] {
	case "project":
		return runUserDataProject(ctx, args[1:], stdout, stderr)
	case "area":
		return runUserDataArea(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserDataHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown user data subcommand: %s\n", args[0])
		printUserDataHelp(stderr)
		return 2
	}
}

func printUserDataHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user data <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "User data permission management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  project   Project data permissions (get-list)")
	fmt.Fprintln(w, "  area      Area data permissions (get-list)")
}

func runUserDataProject(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stdout, "Usage: ur user data project <subcommand> [options]")
		fmt.Fprintln(stdout, "")
		fmt.Fprintln(stdout, "Subcommands:")
		fmt.Fprintln(stdout, "  get-list   Query project permission list")
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserDataProjectGetList(ctx, args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown user data project subcommand: %s\n", args[0])
		return 2
	}
}

func runUserDataProjectGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/data/project/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

func runUserDataArea(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stdout, "Usage: ur user data area <subcommand> [options]")
		fmt.Fprintln(stdout, "")
		fmt.Fprintln(stdout, "Subcommands:")
		fmt.Fprintln(stdout, "  get-list   Query area permission list")
		return 0
	}

	switch args[0] {
	case "get-list":
		return runUserDataAreaGetList(ctx, args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown user data area subcommand: %s\n", args[0])
		return 2
	}
}

func runUserDataAreaGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/data/area/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runUserDept 执行用户部门管理命令
func runUserDept(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUserDeptHelp(stdout)
		return 0
	}

	switch args[0] {
	case "batch-create":
		return runUserDeptBatchCreate(ctx, args[1:], stdout, stderr)
	case "batch-delete":
		return runUserDeptBatchDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printUserDeptHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown user dept subcommand: %s\n", args[0])
		printUserDeptHelp(stderr)
		return 2
	}
}

func printUserDeptHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur user dept <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "User department management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  batch-create   Batch create user departments")
	fmt.Fprintln(w, "  batch-delete   Batch delete user departments")
}

func runUserDeptBatchCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/dept/batch-create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

func runUserDeptBatchDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		Path: "/api/v1/system/user/dept/batch-delete",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
