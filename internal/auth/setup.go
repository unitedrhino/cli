package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	setupCodeChars    = "ABCDEFGHJKMNPQRSTUVWXYZ23456789" // 排除易混淆字符 0/O, 1/l/I
	setupCodeLength   = 8
	pollInterval      = 5 * time.Second
	maxPollCount      = 120 // 10 分钟
	openclawSetupPath = "/api/v1/system/user/self/openclaw/setup-check"
)

// SetupResult CLI 绑定结果
type SetupResult struct {
	AccessKey    string
	AccessSecret string
	TenantCode   string
}

// GenerateSetupCode 生成随机绑定码
func GenerateSetupCode() string {
	b := make([]byte, setupCodeLength)
	for i := range b {
		b[i] = setupCodeChars[rand.Intn(len(setupCodeChars))]
	}
	return string(b)
}

// BuildConsoleURL 构造控制台访问令牌页面 URL
func BuildConsoleURL(baseURL, setupCode string) string {
	// baseURL 是 API 地址，需要转换为控制台地址
	// 简单替换：api.xxx.com -> console.xxx.com，或直接使用 baseURL
	consoleURL := strings.TrimRight(baseURL, "/")
	// 如果 baseURL 包含 /api/ 路径，去掉它（注意用 /api/ 避免匹配主机名中的 api）
	if idx := strings.LastIndex(consoleURL, "/api/"); idx > 0 {
		consoleURL = consoleURL[:idx]
	}
	return fmt.Sprintf("%s/#/user/settings?tab=access-tokens&setup=%s&redirect=openclaw", consoleURL, setupCode)
}

// PollSetupCheck 轮询绑定状态
func PollSetupCheck(ctx context.Context, baseURL, setupCode string) (SetupResult, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := strings.TrimRight(baseURL, "/") + openclawSetupPath

	for i := 0; i < maxPollCount; i++ {
		select {
		case <-ctx.Done():
			return SetupResult{}, ctx.Err()
		default:
		}

		body, _ := json.Marshal(map[string]string{"setupCode": setupCode})
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			return SetupResult{}, fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			// 网络错误时继续轮询
			time.Sleep(pollInterval)
			continue
		}

		var result struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
			Data struct {
				Status       string `json:"status"`
				AccessKey    string `json:"accessKey"`
				AccessSecret string `json:"accessSecret"`
				TenantCode   string `json:"tenantCode"`
			} `json:"data"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			time.Sleep(pollInterval)
			continue
		}
		resp.Body.Close()

		if result.Code != 200 {
			return SetupResult{}, fmt.Errorf("setup-check failed: %s", result.Msg)
		}

		switch result.Data.Status {
		case "completed":
			return SetupResult{
				AccessKey:    result.Data.AccessKey,
				AccessSecret: result.Data.AccessSecret,
				TenantCode:   result.Data.TenantCode,
			}, nil
		case "expired":
			return SetupResult{}, fmt.Errorf("绑定码已过期，请重新运行 ur-api login")
		case "denied":
			return SetupResult{}, fmt.Errorf("绑定被拒绝")
		}

		// pending，继续轮询
		time.Sleep(pollInterval)
	}

	return SetupResult{}, fmt.Errorf("绑定超时（%d 分钟），请重新运行 ur-api login", maxPollCount*5/60)
}
