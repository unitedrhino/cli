package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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
	if err == nil && authCtx.Token != "" {
		return authCtx.Token, nil
	}
	if err == nil && authCtx.UserID != "" && authCtx.AccessKey != "" && authCtx.AccessSecret != "" {
		return GenerateJWT(authCtx.UserID, authCtx.AccessKey, authCtx.AccessSecret)
	}
	profile, err := config.CurrentProfile()
	if err != nil {
		return "", err
	}
	if profile.Account == "" || profile.Password == "" {
		return "", fmt.Errorf("missing session token and password login config")
	}
	baseURL, _ := config.GetBaseURL()
	appID, _ := config.GetAppID()
	tenantCode, _ := config.GetTenantCode()
	result, err := DoPasswordLoginRaw(ctx, baseURL, appID, tenantCode, profile.Account, profile.Password)
	if err != nil {
		return "", err
	}
	return result.Token, nil
}

func ResolveAuthHeaders(ctx context.Context) (map[string]string, error) {
	authCtx, err := config.ResolveAuthContext()
	if err == nil {
		if authCtx.Token != "" {
			return map[string]string{"token": authCtx.Token}, nil
		}
		if authCtx.UserID != "" && authCtx.AccessKey != "" && authCtx.AccessSecret != "" {
			jwt, err := GenerateJWT(authCtx.UserID, authCtx.AccessKey, authCtx.AccessSecret)
			if err != nil {
				return nil, err
			}
			return map[string]string{"Authorization": "Bearer " + jwt}, nil
		}
	}
	token, err := ResolveToken(ctx)
	if err != nil {
		return nil, err
	}
	return map[string]string{"token": token}, nil
}

func GenerateJWT(userID, accessKey, accessSecret string) (string, error) {
	tenantCode, err := config.GetTenantCode()
	if err != nil {
		return "", err
	}
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
