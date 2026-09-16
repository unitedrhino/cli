package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/config"
)

type LoginResult struct {
	Token     string
	ExpireSec int64
}

// CredentialRejectedError 表示后端明确拒绝账号密码，可用于历史 profile 的兼容回退。
type CredentialRejectedError struct {
	Message string
}

// Error 返回后端认证失败信息。
func (e *CredentialRejectedError) Error() string {
	return "login failed: " + e.Message
}

func DoPasswordLoginRaw(ctx context.Context, baseURL, appID, tenantCode, account, password string) (LoginResult, error) {
	body := map[string]any{
		"account":   account,
		"password":  sha256Hex(password),
		"loginType": "pwd",
		"pwdType":   1,
	}
	rawBody, err := json.Marshal(body)
	if err != nil {
		return LoginResult{}, fmt.Errorf("marshal login body: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/api/v1/system/user/self/login", strings.NewReader(string(rawBody)))
	if err != nil {
		return LoginResult{}, fmt.Errorf("build login request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("app-id", appID)
	if tenantCode != "" {
		req.Header.Set("tenant-code", tenantCode)
	}
	resp, err := (&http.Client{Timeout: 30 * time.Second}).Do(req)
	if err != nil {
		return LoginResult{}, fmt.Errorf("do login request: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Token struct {
				AccessToken  string `json:"accessToken"`
				AccessExpire string `json:"accessExpire"`
			} `json:"token"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return LoginResult{}, fmt.Errorf("decode login response: %w", err)
	}
	if data.Code != 200 {
		if isPasswordCredentialRejected(data.Code, data.Msg) {
			return LoginResult{}, &CredentialRejectedError{Message: data.Msg}
		}
		return LoginResult{}, fmt.Errorf("login failed: %s", data.Msg)
	}
	expireAt, _ := strconv.ParseInt(data.Data.Token.AccessExpire, 10, 64)
	expireSec := expireAt - time.Now().Unix()
	if expireSec <= 0 {
		expireSec = 3600
	}
	return LoginResult{Token: data.Data.Token.AccessToken, ExpireSec: expireSec}, nil
}

func ResolveToken(ctx context.Context) (string, error) {
	authCtx, err := config.ResolveAuthContext()
	if err != nil {
		return "", err
	}
	switch authCtx.Method {
	case config.AuthMethodToken:
		return authCtx.Token, nil
	case config.AuthMethodPassword:
		result, loginErr := passwordLoginFromContext(ctx, authCtx)
		if loginErr == nil {
			return result.Token, nil
		}
		if !isCredentialRejected(loginErr) || authCtx.Source != config.AuthSourceProfile || authCtx.AccessKey == "" || authCtx.AccessSecret == "" {
			return "", loginErr
		}
		return generateAccessKeyJWT(authCtx)
	case config.AuthMethodAccessKey:
		return generateAccessKeyJWT(authCtx)
	default:
		return "", fmt.Errorf("unsupported auth method %q", authCtx.Method)
	}
}

// generateAccessKeyJWT 使用 AK/SK 生成 Bearer JWT，userID 缺失时兼容填入 0。
func generateAccessKeyJWT(authCtx config.AuthContext) (string, error) {
	if authCtx.AccessKey != "" && authCtx.AccessSecret != "" {
		userID := authCtx.UserID
		if userID == "" {
			userID = "0"
		}
		return GenerateJWT(userID, authCtx.AccessKey, authCtx.AccessSecret)
	}
	return "", fmt.Errorf("missing access key credentials")
}

// passwordLoginFromContext 使用解析后的账号密码登录，并仅为磁盘 profile 更新 Session Token。
func passwordLoginFromContext(ctx context.Context, authCtx config.AuthContext) (LoginResult, error) {
	if authCtx.Account == "" || authCtx.Password == "" {
		return LoginResult{}, fmt.Errorf("missing account/password login config")
	}
	baseURL, err := config.GetBaseURL()
	if err != nil {
		return LoginResult{}, err
	}
	appID, err := config.GetAppID()
	if err != nil {
		return LoginResult{}, err
	}
	tenantCode, err := config.GetTenantCode()
	if err != nil {
		return LoginResult{}, err
	}
	result, err := DoPasswordLoginRaw(ctx, baseURL, appID, tenantCode, authCtx.Account, authCtx.Password)
	if err != nil {
		return LoginResult{}, err
	}
	if authCtx.Source == config.AuthSourceProfile {
		if err := config.UpdateSessionToken(result.Token); err != nil {
			return LoginResult{}, fmt.Errorf("save session token: %w", err)
		}
	}
	return result, nil
}

// isCredentialRejected 判断错误是否为后端明确的凭据拒绝，而非网络或解析错误。
func isCredentialRejected(err error) bool {
	var rejected *CredentialRejectedError
	return errors.As(err, &rejected)
}

// isPasswordCredentialRejected 仅识别后端明确的账号密码拒绝，避免把企业校验等业务错误当成凭据失效。
func isPasswordCredentialRejected(code int, message string) bool {
	if code == http.StatusUnauthorized {
		return true
	}
	msg := strings.ToLower(message)
	for _, keyword := range []string{"账号或密码", "账户或密码", "密码错误", "invalid credential", "invalid password", "account or password"} {
		if strings.Contains(msg, keyword) {
			return true
		}
	}
	return false
}

func ResolveAuthHeaders(ctx context.Context) (map[string]string, error) {
	authCtx, err := config.ResolveAuthContext()
	if err != nil {
		return nil, err
	}
	switch authCtx.Method {
	case config.AuthMethodToken:
		return map[string]string{"token": authCtx.Token}, nil
	case config.AuthMethodPassword:
		result, loginErr := passwordLoginFromContext(ctx, authCtx)
		if loginErr == nil {
			return map[string]string{"token": result.Token}, nil
		}
		if !isCredentialRejected(loginErr) || authCtx.Source != config.AuthSourceProfile || authCtx.AccessKey == "" || authCtx.AccessSecret == "" {
			return nil, loginErr
		}
		jwt, jwtErr := generateAccessKeyJWT(authCtx)
		if jwtErr != nil {
			return nil, jwtErr
		}
		return map[string]string{"Authorization": "Bearer " + jwt}, nil
	case config.AuthMethodAccessKey:
		jwt, jwtErr := generateAccessKeyJWT(authCtx)
		if jwtErr != nil {
			return nil, jwtErr
		}
		return map[string]string{"Authorization": "Bearer " + jwt}, nil
	default:
		return nil, fmt.Errorf("unsupported auth method %q", authCtx.Method)
	}
}

// ResolveFallbackAuthHeaders 在接口明确返回认证失败后，按历史 profile 的剩余候选凭据生成重试头。
// 环境变量认证不会读取或回退到磁盘 profile，网络与响应解析错误也不会触发候选切换。
func ResolveFallbackAuthHeaders(ctx context.Context) (map[string]string, error) {
	authCtx, err := config.ResolveAuthContext()
	if err != nil {
		return nil, err
	}
	if authCtx.Source != config.AuthSourceProfile {
		return nil, errors.New("environment authentication has no profile fallback")
	}

	if authCtx.Method == config.AuthMethodToken && authCtx.Account != "" && authCtx.Password != "" {
		result, loginErr := passwordLoginFromContext(ctx, authCtx)
		if loginErr == nil {
			return map[string]string{"token": result.Token}, nil
		}
		if !isCredentialRejected(loginErr) {
			return nil, loginErr
		}
	}

	if authCtx.Method != config.AuthMethodAccessKey && authCtx.AccessKey != "" && authCtx.AccessSecret != "" {
		jwt, jwtErr := generateAccessKeyJWT(authCtx)
		if jwtErr != nil {
			return nil, jwtErr
		}
		return map[string]string{"Authorization": "Bearer " + jwt}, nil
	}
	return nil, errors.New("no fallback authentication candidate in profile")
}

func GenerateJWT(userID, accessKey, accessSecret string) (string, error) {
	tenantCode, err := config.GetTenantCode()
	if err != nil {
		return "", err
	}
	return GenerateJWTForTenant(userID, accessKey, accessSecret, tenantCode)
}

// GenerateJWTForTenant 使用明确的企业编码生成 AK/SK Bearer JWT，不依赖或修改进程环境变量。
func GenerateJWTForTenant(userID, accessKey, accessSecret, tenantCode string) (string, error) {
	now := time.Now().Unix()
	header := map[string]any{"alg": "HS256", "typ": "JWT"}
	payload := map[string]any{
		"userID":     userID,
		"tenantCode": tenantCode,
		"accessKey":  accessKey,
		"iat":        now,
		"exp":        now + 3600,
	}
	encodedHeader, err := marshalSegment(header)
	if err != nil {
		return "", err
	}
	encodedPayload, err := marshalSegment(payload)
	if err != nil {
		return "", err
	}
	unsigned := encodedHeader + "." + encodedPayload
	mac := hmac.New(sha256.New, []byte(accessSecret))
	_, _ = mac.Write([]byte(unsigned))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return unsigned + "." + signature, nil
}

func marshalSegment(v any) (string, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(raw), nil
}

func sha256Hex(value string) string {
	sum := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", sum[:])
}

// RefreshToken 尝试用保存的账号密码重新登录获取新 token。
// 成功后会将新 token 写回配置文件。
func RefreshToken(ctx context.Context) (string, error) {
	authCtx, err := config.ResolveAuthContext()
	if err != nil {
		return "", fmt.Errorf("resolve auth context: %w", err)
	}
	if authCtx.Account == "" || authCtx.Password == "" {
		return "", fmt.Errorf("no account/password in profile, cannot refresh token")
	}
	result, err := passwordLoginFromContext(ctx, authCtx)
	if err != nil {
		return "", fmt.Errorf("refresh login failed: %w", err)
	}
	return result.Token, nil
}
