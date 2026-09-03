package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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
	// openclawSetupInitPath 登记绑定码接口：在生成绑定码时先调用，
	// 让后端从登记时刻开始计算 10 分钟有效期
	openclawSetupInitPath = "/api/v1/system/user/self/openclaw/setup-init"
)

// ErrSetupInitUnsupported 表示后端尚未部署 setup-init 接口（旧版后端，接口返回 404）。
// 调用方应据此降级继续登录流程（绑定码有效期回退为从首次轮询起算），不阻断登录。
var ErrSetupInitUnsupported = errors.New("后端尚未支持 setup-init 绑定码登记（旧版后端），绑定码有效期改为从首次轮询起算")

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

// InitSetup 向后端登记绑定码，使「10 分钟有效期」从生成起算：
// 后端在 setup-init 成功后即开始计时，避免 AI 拿到授权 URL 后用户迟迟未操作导致时间窗口错位。
// 返回 ErrSetupInitUnsupported 表示旧版后端未部署该接口，调用方可降级继续（不阻断登录）。
func InitSetup(ctx context.Context, baseURL, setupCode string) error {
	client := &http.Client{Timeout: 10 * time.Second}
	url := strings.TrimRight(baseURL, "/") + openclawSetupInitPath

	body, _ := json.Marshal(map[string]string{"setupCode": setupCode})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("构建 setup-init 请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("登记绑定码请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 404：后端接口未部署（旧版后端），返回可识别错误让调用方降级
	if resp.StatusCode == http.StatusNotFound {
		return ErrSetupInitUnsupported
	}

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Success bool `json:"success"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("解析 setup-init 响应失败: %w", err)
	}
	if result.Code != 200 {
		return fmt.Errorf("setup-init failed: %s", result.Msg)
	}
	return nil
}

// PollSetupCheck 轮询绑定状态
// onPoll 回调：每次轮询时调用，参数为 (当前次数, 总次数, 是否完成, 错误)
func PollSetupCheck(ctx context.Context, baseURL, setupCode string, onPoll func(current, total int, done bool, err error)) (SetupResult, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := strings.TrimRight(baseURL, "/") + openclawSetupPath

	for i := 0; i < maxPollCount; i++ {
		select {
		case <-ctx.Done():
			if onPoll != nil {
				onPoll(i, maxPollCount, false, ctx.Err())
			}
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
			if onPoll != nil {
				onPoll(i, maxPollCount, false, fmt.Errorf("网络错误: %w", err))
			}
			time.Sleep(pollInterval)
			continue
		}

		// 检查 404：后端接口未部署
		if resp.StatusCode == http.StatusNotFound {
			resp.Body.Close()
			err := fmt.Errorf("后端尚未支持 CLI 绑定功能（接口 %s 返回 404），请联系管理员升级后端版本", openclawSetupPath)
			if onPoll != nil {
				onPoll(i, maxPollCount, false, err)
			}
			return SetupResult{}, err
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
			if onPoll != nil {
				onPoll(i, maxPollCount, false, fmt.Errorf("解析响应失败: %w", err))
			}
			time.Sleep(pollInterval)
			continue
		}
		resp.Body.Close()

		if result.Code != 200 {
			err := fmt.Errorf("setup-check failed: %s", result.Msg)
			if onPoll != nil {
				onPoll(i, maxPollCount, false, err)
			}
			return SetupResult{}, err
		}

		switch result.Data.Status {
		case "completed":
			if onPoll != nil {
				onPoll(i, maxPollCount, true, nil)
			}
			return SetupResult{
				AccessKey:    result.Data.AccessKey,
				AccessSecret: result.Data.AccessSecret,
				TenantCode:   result.Data.TenantCode,
			}, nil
		case "expired":
			err := fmt.Errorf("绑定码已过期，请重新运行 ur-iot login")
			if onPoll != nil {
				onPoll(i, maxPollCount, false, err)
			}
			return SetupResult{}, err
		case "denied":
			err := fmt.Errorf("绑定被拒绝")
			if onPoll != nil {
				onPoll(i, maxPollCount, false, err)
			}
			return SetupResult{}, err
		}

		// pending，继续轮询
		if onPoll != nil {
			onPoll(i, maxPollCount, false, nil)
		}
		time.Sleep(pollInterval)
	}

	err := fmt.Errorf("绑定超时（%d 分钟），请重新运行 ur-iot login", maxPollCount*5/60)
	if onPoll != nil {
		onPoll(maxPollCount, maxPollCount, false, err)
	}
	return SetupResult{}, err
}
