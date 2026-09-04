package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/auth"
	"gitee.com/unitedrhino/cli/internal/config"
)

func runLogin(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	// 解析参数
	var (
		noWait     bool
		setupCode  string
		jsonMode   bool
		baseURLArg string
	)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--no-wait":
			noWait = true
		case "--setup-code":
			if i+1 < len(args) {
				setupCode = args[i+1]
				i++
			}
		case "--json":
			jsonMode = true
		case "--base-url":
			if i+1 < len(args) {
				baseURLArg = strings.TrimRight(args[i+1], "/")
				i++
			}
		}
	}

	// 1. 确定 baseURL
	baseURL, err := resolveBaseURL(baseURLArg, stdout, stderr)
	if err != nil {
		return outputLoginError(stderr, jsonMode, err)
	}

	// --setup-code 模式：直接轮询完成授权
	if setupCode != "" {
		return runLoginPoll(ctx, baseURL, setupCode, jsonMode, stdout, stderr)
	}

	// 2. 生成 setupCode
	setupCode = auth.GenerateSetupCode()
	consoleURL := auth.BuildConsoleURL(baseURL, setupCode)

	// --no-wait 模式：输出信息，立即返回（AI 友好）
	if noWait {
		if jsonMode {
			b, _ := json.Marshal(map[string]interface{}{
				"status":           "authorization_required",
				"verification_url": consoleURL,
				"setup_code":       setupCode,
				"expires_in":       600,
				"next_command":     fmt.Sprintf("ur login --setup-code %s", setupCode),
			})
			fmt.Fprintln(stdout, string(b))
		} else {
			fmt.Fprintf(stdout, "请在浏览器中完成授权：%s\n", consoleURL)
			fmt.Fprintf(stdout, "授权完成后执行: ur login --setup-code %s\n", setupCode)
		}
		return 0
	}

	// 3. 默认模式：显示 URL + 阻塞轮询（兼容现有行为）
	fmt.Fprintln(stdout, "请在浏览器中完成授权：")
	fmt.Fprintln(stdout, consoleURL)
	fmt.Fprintln(stdout, "\n步骤：")
	fmt.Fprintln(stdout, "  1. 点击链接进入控制台「访问令牌」页面")
	fmt.Fprintln(stdout, "  2. 创建或选择一个访问令牌")
	fmt.Fprintln(stdout, "  3. 点击「完成第三方客户端绑定」")
	fmt.Fprintln(stdout, "\n正在等待授权...（每5秒检测一次，最多10分钟）")

	return runLoginPoll(ctx, baseURL, setupCode, jsonMode, stdout, stderr)
}

// runLoginPoll 执行轮询 + 保存配置 + 验证连接
func runLoginPoll(ctx context.Context, baseURL, setupCode string, jsonMode bool, stdout, stderr io.Writer) int {
	result, err := auth.PollSetupCheck(ctx, baseURL, setupCode, func(current, total int, done bool, pollErr error) {
		if jsonMode {
			return
		}
		if done {
			fmt.Fprintf(stdout, "\r✓ 检测到授权完成（第 %d/%d 次检测）\n", current+1, total)
			return
		}
		if pollErr != nil {
			// 非 404 的网络错误只显示一次提示
			if current == 0 {
				fmt.Fprintf(stdout, "\r  正在检测...（第 %d/%d 次）网络暂时不稳定，继续等待...\n", current+1, total)
			}
			return
		}
		// 每 6 次（30 秒）更新一次进度
		if current%6 == 0 && current > 0 {
			remainingMin := (total - current) * 5 / 60
			fmt.Fprintf(stdout, "\r  正在等待授权...（第 %d/%d 次检测，剩余约 %d 分钟）", current+1, total, remainingMin)
		}
	})
	if err != nil {
		return outputLoginError(stderr, jsonMode, err)
	}

	// 验证连接并获取用户信息
	userID, err := verifyConnection(ctx, baseURL, result.AccessKey, result.AccessSecret, result.TenantCode)
	if err != nil {
		return outputLoginError(stderr, jsonMode, fmt.Errorf("验证连接失败: %w", err))
	}

	appID, _ := config.GetAppID()

	// 保存配置
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
		return outputLoginError(stderr, jsonMode, fmt.Errorf("保存配置失败: %w", err))
	}

	if jsonMode {
		b, _ := json.Marshal(map[string]interface{}{
			"event":         "authorization_complete",
			"tenant_code":   result.TenantCode,
			"access_key":    result.AccessKey,
			"access_secret": result.AccessSecret,
			"user_id":       userID,
		})
		fmt.Fprintln(stdout, string(b))
	} else {
		fmt.Fprintf(stdout, "\n✓ 授权成功！\n")
		fmt.Fprintf(stdout, "企业:       %s\n", result.TenantCode)
		fmt.Fprintf(stdout, "AccessKey:  %s...（已保存）\n", result.AccessKey[:6])
		fmt.Fprintln(stdout, "✓ 连接验证成功！")
		fmt.Fprintln(stdout, "\n初始化完成，您现在可以使用 ur-api 命令了。")
	}
	return 0
}

// resolveBaseURL 确定 baseURL：参数 > 环境变量 > 交互式
func resolveBaseURL(baseURLArg string, stdout, stderr io.Writer) (string, error) {
	if baseURLArg != "" {
		return baseURLArg, nil
	}
	if envURL := os.Getenv("UR_BASE_URL"); envURL != "" {
		return strings.TrimRight(envURL, "/"), nil
	}
	return selectBaseURLInteractive(stdout)
}

// selectBaseURLInteractive 交互式选择平台地址
func selectBaseURLInteractive(stdout io.Writer) (string, error) {
	fmt.Fprintln(stdout, "请选择联犀平台地址：")
	fmt.Fprintln(stdout, "[1] 联犀 SaaS (https://saas.unitedrhino.com)")
	fmt.Fprintln(stdout, "[2] 私有化部署（自定义）")
	fmt.Fprint(stdout, "> ")

	var choice string
	if _, err := fmt.Fscanln(os.Stdin, &choice); err != nil {
		return "", fmt.Errorf("读取输入失败: %w", err)
	}

	switch choice {
	case "1":
		return "https://saas.unitedrhino.com", nil
	case "2":
		fmt.Fprint(stdout, "请输入私有化地址: ")
		var customURL string
		if _, err := fmt.Fscanln(os.Stdin, &customURL); err != nil {
			return "", fmt.Errorf("读取输入失败: %w", err)
		}
		return strings.TrimRight(customURL, "/"), nil
	default:
		return "", fmt.Errorf("无效选择: %s", choice)
	}
}

// outputLoginError 统一错误输出（支持 JSON 模式）
func outputLoginError(stderr io.Writer, jsonMode bool, err error) int {
	if jsonMode {
		b, _ := json.Marshal(map[string]interface{}{
			"event": "authorization_failed",
			"error": err.Error(),
		})
		fmt.Fprintln(stderr, string(b))
	} else {
		fmt.Fprintf(stderr, "\n✗ %v\n", err)
	}
	return 1
}

func verifyConnection(ctx context.Context, baseURL, accessKey, accessSecret, tenantCode string) (string, error) {
	// Device Flow 首次登录时配置文件尚未保存，临时设置环境变量
	if tenantCode != "" {
		os.Setenv("UR_TENANT_CODE", tenantCode)
	}
	jwt, err := auth.GenerateJWT("0", accessKey, accessSecret)
	if err != nil {
		return "", fmt.Errorf("生成 JWT: %w", err)
	}

	appID, err := config.GetAppID()
	if err != nil {
		// Device Flow 首次登录时配置文件尚未保存，使用默认值
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
