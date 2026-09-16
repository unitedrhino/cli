package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/auth"
	"gitee.com/unitedrhino/cli/internal/config"
	"github.com/spf13/cobra"
)

type loginOptions struct {
	noWait            bool
	setupCode         string
	jsonMode          bool
	baseURL           string
	method            string
	tenantCode        string
	account           string
	password          string
	passwordStdin     bool
	accessKey         string
	accessSecret      string
	accessSecretStdin bool
}

var loginOpts loginOptions

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登录联犀平台",
	Long:  `支持 Device Flow、账号密码和 AK/SK 三种认证方式；未指定 --method 时默认 Device Flow。`,
	RunE:  runLogin,
}

func init() {
	loginCmd.Flags().StringVar(&loginOpts.method, "method", "device", "认证方式 (device, password, aksk)")
	loginCmd.Flags().BoolVar(&loginOpts.noWait, "no-wait", false, "不阻塞，输出 URL 后返回")
	loginCmd.Flags().StringVar(&loginOpts.setupCode, "setup-code", "", "使用已有的 setup code 完成授权")
	loginCmd.Flags().BoolVar(&loginOpts.jsonMode, "json", false, "JSON 模式输出")
	loginCmd.Flags().StringVar(&loginOpts.baseURL, "base-url", "", "平台地址（覆盖配置文件）")
	loginCmd.Flags().StringVar(&loginOpts.tenantCode, "tenant-code", "", "企业编码")
	loginCmd.Flags().StringVar(&loginOpts.account, "account", "", "登录账号")
	loginCmd.Flags().StringVar(&loginOpts.password, "password", "", "原始密码（建议改用 UR_PASSWORD 或 --password-stdin）")
	loginCmd.Flags().BoolVar(&loginOpts.passwordStdin, "password-stdin", false, "从标准输入读取原始密码")
	loginCmd.Flags().StringVar(&loginOpts.accessKey, "access-key", "", "访问令牌 AccessKey")
	loginCmd.Flags().StringVar(&loginOpts.accessSecret, "access-secret", "", "访问令牌 AccessSecret（建议改用 UR_ACCESS_SECRET 或 --access-secret-stdin）")
	loginCmd.Flags().BoolVar(&loginOpts.accessSecretStdin, "access-secret-stdin", false, "从标准输入读取 AccessSecret")
	RootCmd.AddCommand(loginCmd)
}

func runLogin(cmd *cobra.Command, args []string) error {
	method := strings.ToLower(strings.TrimSpace(loginOpts.method))
	if method == "" {
		method = "device"
	}
	switch method {
	case "device":
		return runDeviceLogin(cmd)
	case "password":
		return runPasswordLogin(cmd)
	case "aksk":
		return runAccessKeyLogin(cmd)
	default:
		return outputLoginError(cmd, fmt.Errorf("unsupported login method %q", method))
	}
}

// runDeviceLogin 执行默认 Device Flow 登录流程。
func runDeviceLogin(cmd *cobra.Command) error {
	ctx := cmd.Context()

	baseURL, err := resolveLoginBaseURL(loginOpts.baseURL)
	if err != nil {
		return outputLoginError(cmd, err)
	}

	if loginOpts.setupCode != "" {
		return runLoginPoll(cmd, ctx, baseURL, loginOpts.setupCode)
	}

	setupCode := auth.GenerateSetupCode()

	// 先向后端登记绑定码，使 10 分钟有效期从生成起算；
	// 旧版后端未部署 setup-init（404）时降级继续，不阻断登录流程。
	// --setup-code 恢复轮询分支（上方提前 return）不登记，保持对既有流程无影响。
	if err := auth.InitSetup(ctx, baseURL, setupCode); err != nil {
		if !errors.Is(err, auth.ErrSetupInitUnsupported) {
			return outputLoginError(cmd, fmt.Errorf("登记绑定码失败: %w", err))
		}
		if loginOpts.jsonMode {
			cmd.Printf(`{"event":"setup_init_skipped","reason":%q}`+"\n", err.Error())
		} else {
			cmd.Printf("提示: %v\n", err)
		}
	}

	consoleURL := auth.BuildConsoleURL(baseURL, setupCode)

	if loginOpts.noWait {
		if loginOpts.jsonMode {
			cmd.Println(fmt.Sprintf(`{"status":"authorization_required","verification_url":"%s","setup_code":"%s","expires_in":600}`, consoleURL, setupCode))
		} else {
			cmd.Printf("请在浏览器中完成授权：%s\n", consoleURL)
			cmd.Printf("授权完成后执行: ur login --setup-code %s\n", setupCode)
		}
		return nil
	}

	cmd.Println("请在浏览器中完成授权：")
	cmd.Println(consoleURL)
	cmd.Println("\n步骤：")
	cmd.Println("  1. 点击链接进入控制台「访问令牌」页面")
	cmd.Println("  2. 创建或选择一个访问令牌")
	cmd.Println("  3. 点击「完成第三方客户端绑定」")
	cmd.Println("\n正在等待授权...（每5秒检测一次，最多10分钟）")

	return runLoginPoll(cmd, ctx, baseURL, setupCode)
}

// runPasswordLogin 使用账号和原始密码立即换取 Session Token。
func runPasswordLogin(cmd *cobra.Command) error {
	baseURL, err := resolveLoginBaseURL(loginOpts.baseURL)
	if err != nil {
		return outputLoginError(cmd, err)
	}
	appID, tenantCode, err := resolveDirectLoginContext(loginOpts.tenantCode)
	if err != nil {
		return outputLoginError(cmd, err)
	}
	account := strings.TrimSpace(loginOpts.account)
	if account == "" {
		account = strings.TrimSpace(os.Getenv("UR_ACCOUNT"))
	}
	if account == "" {
		return outputLoginError(cmd, errors.New("account is required for password login"))
	}
	password, err := resolveLoginSecret(cmd.InOrStdin(), loginOpts.password, loginOpts.passwordStdin, "UR_PASSWORD", "password")
	if err != nil {
		return outputLoginError(cmd, err)
	}
	if password == "" {
		return outputLoginError(cmd, errors.New("password is required for password login"))
	}
	if loginOpts.password != "" && !loginOpts.jsonMode {
		cmd.PrintErrln("警告: --password 可能暴露在 shell 历史或进程列表中，建议使用 UR_PASSWORD 或 --password-stdin")
	}
	result, err := auth.DoPasswordLoginRaw(cmd.Context(), baseURL, appID, tenantCode, account, password)
	if err != nil {
		return outputLoginError(cmd, err)
	}
	profile := config.Profile{
		BaseURL: baseURL, AppID: appID, TenantCode: tenantCode, Role: string(resolveAppFromContext()),
		Token: result.Token, Account: account, Password: password,
	}
	if err := config.ReplaceProfileAuth(profile, config.AuthMethodPassword); err != nil {
		return outputLoginError(cmd, fmt.Errorf("保存配置失败: %w", err))
	}
	return outputLoginSuccess(cmd, config.AuthMethodPassword, tenantCode)
}

// runAccessKeyLogin 使用 AK/SK 生成 JWT、验证连接并保存认证配置。
func runAccessKeyLogin(cmd *cobra.Command) error {
	baseURL, err := resolveLoginBaseURL(loginOpts.baseURL)
	if err != nil {
		return outputLoginError(cmd, err)
	}
	appID, tenantCode, err := resolveDirectLoginContext(loginOpts.tenantCode)
	if err != nil {
		return outputLoginError(cmd, err)
	}
	accessKey := strings.TrimSpace(loginOpts.accessKey)
	if accessKey == "" {
		accessKey = strings.TrimSpace(os.Getenv("UR_ACCESS_KEY"))
	}
	if accessKey == "" {
		return outputLoginError(cmd, errors.New("access-key is required for AK/SK login"))
	}
	accessSecret, err := resolveLoginSecret(cmd.InOrStdin(), loginOpts.accessSecret, loginOpts.accessSecretStdin, "UR_ACCESS_SECRET", "access-secret")
	if err != nil {
		return outputLoginError(cmd, err)
	}
	if accessSecret == "" {
		return outputLoginError(cmd, errors.New("access-secret is required for AK/SK login"))
	}
	if loginOpts.accessSecret != "" && !loginOpts.jsonMode {
		cmd.PrintErrln("警告: --access-secret 可能暴露在 shell 历史或进程列表中，建议使用 UR_ACCESS_SECRET 或 --access-secret-stdin")
	}
	userID, err := verifyLoginConnection(cmd.Context(), baseURL, appID, accessKey, accessSecret, tenantCode)
	if err != nil {
		return outputLoginError(cmd, fmt.Errorf("验证连接失败: %w", err))
	}
	profile := config.Profile{
		BaseURL: baseURL, AppID: appID, TenantCode: tenantCode, Role: string(resolveAppFromContext()),
		AccessKey: accessKey, AccessSecret: accessSecret,
	}
	if userID != "" {
		if uid, parseErr := strconv.ParseInt(userID, 10, 64); parseErr == nil {
			profile.UserID = uid
		}
	}
	if err := config.ReplaceProfileAuth(profile, config.AuthMethodAccessKey); err != nil {
		return outputLoginError(cmd, fmt.Errorf("保存配置失败: %w", err))
	}
	return outputLoginSuccess(cmd, config.AuthMethodAccessKey, tenantCode)
}

// resolveLoginSecret 按显式参数、stdin、环境变量的顺序解析敏感值。
func resolveLoginSecret(reader io.Reader, explicit string, fromStdin bool, envKey, label string) (string, error) {
	if explicit != "" {
		return explicit, nil
	}
	if fromStdin {
		raw, err := io.ReadAll(reader)
		if err != nil {
			return "", fmt.Errorf("read %s from stdin: %w", label, err)
		}
		return strings.TrimRight(string(raw), "\r\n"), nil
	}
	return os.Getenv(envKey), nil
}

// resolveDirectLoginContext 解析账号密码和 AK/SK 登录共用的应用及企业上下文。
func resolveDirectLoginContext(explicitTenantCode string) (string, string, error) {
	appID, err := config.GetAppID()
	if err != nil || strings.TrimSpace(appID) == "" {
		appID = resolveAppFromContext().AppID()
	}
	tenantCode := strings.TrimSpace(explicitTenantCode)
	if tenantCode == "" {
		tenantCode = strings.TrimSpace(os.Getenv("UR_TENANT_CODE"))
	}
	if tenantCode == "" {
		if profile, profileErr := config.CurrentProfile(); profileErr == nil {
			tenantCode = strings.TrimSpace(profile.TenantCode)
		}
	}
	if tenantCode == "" {
		return "", "", errors.New("tenant-code is required for this application")
	}
	return appID, tenantCode, nil
}

// outputLoginSuccess 输出不包含秘密值的稳定登录结果。
func outputLoginSuccess(cmd *cobra.Command, method config.AuthMethod, tenantCode string) error {
	if loginOpts.jsonMode {
		payload, _ := json.Marshal(map[string]any{"event": "authorization_complete", "status": "ok", "method": method, "tenant_code": tenantCode})
		cmd.Println(string(payload))
		return nil
	}
	cmd.Printf("✓ 登录成功（方式: %s，企业: %s）\n", method, tenantCode)
	return nil
}

func runLoginPoll(cmd *cobra.Command, ctx context.Context, baseURL, setupCode string) error {
	result, err := auth.PollSetupCheck(ctx, baseURL, setupCode, nil)
	if err != nil {
		return outputLoginError(cmd, err)
	}

	appID, appErr := config.GetAppID()
	if appErr != nil {
		appID = resolveAppFromContext().AppID()
	}
	userID, err := verifyLoginConnection(ctx, baseURL, appID, result.AccessKey, result.AccessSecret, result.TenantCode)
	if err != nil {
		return outputLoginError(cmd, fmt.Errorf("验证连接失败: %w", err))
	}

	profile := config.Profile{
		BaseURL:      baseURL,
		AppID:        appID,
		TenantCode:   result.TenantCode,
		AccessKey:    result.AccessKey,
		AccessSecret: result.AccessSecret,
	}
	if userID != "" {
		if uid, err := strconv.ParseInt(userID, 10, 64); err == nil {
			profile.UserID = uid
		}
	}
	if err := config.ReplaceProfileAuth(profile, config.AuthMethodAccessKey); err != nil {
		return outputLoginError(cmd, fmt.Errorf("保存配置失败: %w", err))
	}

	if loginOpts.jsonMode {
		cmd.Printf(`{"event":"authorization_complete","status":"ok","method":"device","tenant_code":"%s","access_key":"%s"}`+"\n", result.TenantCode, result.AccessKey)
	} else {
		cmd.Printf("\n✓ 授权成功！\n")
		cmd.Printf("企业:       %s\n", result.TenantCode)
		cmd.Printf("AccessKey:  %s...（已保存）\n", result.AccessKey[:6])
		cmd.Println("✓ 连接验证成功！")
	}
	return nil
}

// defaultLoginBaseURL 未指定平台地址时使用的联犀 SaaS 默认地址，
// 使 `ur login` 不传 --base-url 也能开箱即用（文档示例通常省略该参数）
const defaultLoginBaseURL = "https://saas.unitedrhino.com"

// resolveLoginBaseURL 确定登录平台地址：--base-url > UR_BASE_URL > 当前 profile > 默认 SaaS 地址。
func resolveLoginBaseURL(baseURL string) (string, error) {
	if baseURL != "" {
		return strings.TrimRight(baseURL, "/"), nil
	}
	if envURL := os.Getenv("UR_BASE_URL"); envURL != "" {
		return strings.TrimRight(envURL, "/"), nil
	}
	if profile, err := config.CurrentProfile(); err == nil && strings.TrimSpace(profile.BaseURL) != "" {
		return strings.TrimRight(strings.TrimSpace(profile.BaseURL), "/"), nil
	}
	return defaultLoginBaseURL, nil
}

func outputLoginError(cmd *cobra.Command, err error) error {
	if loginOpts.jsonMode {
		payload, _ := json.Marshal(map[string]any{
			"event":  "authorization_failed",
			"status": "error",
			"method": strings.ToLower(strings.TrimSpace(loginOpts.method)),
			"error":  err.Error(),
		})
		return &CLIError{Message: string(payload), ExitCode: 1}
	}
	return &CLIError{Message: err.Error(), ExitCode: 1}
}

func verifyLoginConnection(ctx context.Context, baseURL, appID, accessKey, accessSecret, tenantCode string) (string, error) {
	jwt, err := auth.GenerateJWTForTenant("0", accessKey, accessSecret, tenantCode)
	if err != nil {
		return "", fmt.Errorf("生成 JWT: %w", err)
	}

	reqBody, _ := json.Marshal(map[string]any{})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/api/v1/system/user/self/get-one", bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("构建请求: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("app-id", appID)
	req.Header.Set("tenant-code", tenantCode)

	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return "", fmt.Errorf("发送请求: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			UserID string `json:"userID"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("解析响应: %w", err)
	}
	if result.Code != 200 {
		return "", fmt.Errorf("API 错误: %s", result.Msg)
	}
	return result.Data.UserID, nil
}
