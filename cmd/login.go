package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/internal/auth"
	"gitee.com/unitedrhino/cli/internal/config"
)

var loginOpts struct {
	noWait    bool
	setupCode string
	jsonMode  bool
	baseURL   string
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "通过 OAuth 设备流登录",
	Long:  `授权 CLI 访问联犀平台。支持交互式授权和 --no-wait 模式（AI 友好）。`,
	RunE:  runLogin,
}

func init() {
	loginCmd.Flags().BoolVar(&loginOpts.noWait, "no-wait", false, "不阻塞，输出 URL 后返回")
	loginCmd.Flags().StringVar(&loginOpts.setupCode, "setup-code", "", "使用已有的 setup code 完成授权")
	loginCmd.Flags().BoolVar(&loginOpts.jsonMode, "json", false, "JSON 模式输出")
	loginCmd.Flags().StringVar(&loginOpts.baseURL, "base-url", "", "平台地址（覆盖配置文件）")
	RootCmd.AddCommand(loginCmd)
}

func runLogin(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	baseURL, err := resolveLoginBaseURL(loginOpts.baseURL)
	if err != nil {
		return outputLoginError(cmd, err)
	}

	if loginOpts.setupCode != "" {
		return runLoginPoll(cmd, ctx, baseURL, loginOpts.setupCode)
	}

	setupCode := auth.GenerateSetupCode()
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
	cmd.Println("  3. 点击「完成 CLI 绑定」")
	cmd.Println("\n正在等待授权...（每5秒检测一次，最多10分钟）")

	return runLoginPoll(cmd, ctx, baseURL, setupCode)
}

func runLoginPoll(cmd *cobra.Command, ctx context.Context, baseURL, setupCode string) error {
	result, err := auth.PollSetupCheck(ctx, baseURL, setupCode, nil)
	if err != nil {
		return outputLoginError(cmd, err)
	}

	userID, err := verifyLoginConnection(ctx, baseURL, result.AccessKey, result.AccessSecret, result.TenantCode)
	if err != nil {
		return outputLoginError(cmd, fmt.Errorf("验证连接失败: %w", err))
	}

	appID, _ := config.GetAppID()

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
	if err := config.SaveProfile(profile); err != nil {
		return outputLoginError(cmd, fmt.Errorf("保存配置失败: %w", err))
	}

	if loginOpts.jsonMode {
		cmd.Printf(`{"event":"authorization_complete","tenant_code":"%s","access_key":"%s"}`+"\n", result.TenantCode, result.AccessKey)
	} else {
		cmd.Printf("\n✓ 授权成功！\n")
		cmd.Printf("租户:       %s\n", result.TenantCode)
		cmd.Printf("AccessKey:  %s...（已保存）\n", result.AccessKey[:6])
		cmd.Println("✓ 连接验证成功！")
	}
	return nil
}

func resolveLoginBaseURL(baseURL string) (string, error) {
	if baseURL != "" {
		return strings.TrimRight(baseURL, "/"), nil
	}
	if envURL := os.Getenv("UR_BASE_URL"); envURL != "" {
		return strings.TrimRight(envURL, "/"), nil
	}
	return "", fmt.Errorf("请指定 --base-url 或设置 UR_BASE_URL 环境变量")
}

func outputLoginError(cmd *cobra.Command, err error) error {
	if loginOpts.jsonMode {
		return &CLIError{Message: fmt.Sprintf(`{"event":"authorization_failed","error":"%s"}`, err.Error()), ExitCode: 1}
	}
	return &CLIError{Message: err.Error(), ExitCode: 1}
}

func verifyLoginConnection(ctx context.Context, baseURL, accessKey, accessSecret, tenantCode string) (string, error) {
	if tenantCode != "" {
		os.Setenv("UR_TENANT_CODE", tenantCode)
	}
	jwt, err := auth.GenerateJWT("0", accessKey, accessSecret)
	if err != nil {
		return "", fmt.Errorf("生成 JWT: %w", err)
	}

	appID, err := config.GetAppID()
	if err != nil {
		appID = "77"
	}

	reqBody, _ := json.Marshal(map[string]any{})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/api/v1/system/user/self/get-one", bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("构建请求: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("app-id", appID)

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
