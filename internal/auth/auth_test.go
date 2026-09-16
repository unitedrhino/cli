// auth_test.go — CLI 三种认证与历史账号密码兼容回归测试。
package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gitee.com/unitedrhino/cli/internal/config"
)

// TestDoPasswordLoginRawHashesPasswordExactlyOnce 验证 CLI 只对原始密码做一次 SHA-256。
func TestDoPasswordLoginRawHashesPasswordExactlyOnce(t *testing.T) {
	var gotPassword, gotAppID, gotTenantCode string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAppID = r.Header.Get("app-id")
		gotTenantCode = r.Header.Get("tenant-code")
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("Decode request: %v", err)
		}
		gotPassword, _ = body["password"].(string)
		_, _ = w.Write([]byte(`{"code":200,"msg":"","data":{"token":{"accessToken":"session-token","accessExpire":"4102444800"}}}`))
	}))
	defer server.Close()

	result, err := DoPasswordLoginRaw(context.Background(), server.URL, "300", "tenant-a", "admin", "raw-password")
	if err != nil {
		t.Fatalf("DoPasswordLoginRaw error: %v", err)
	}
	wantHash := sha256.Sum256([]byte("raw-password"))
	if gotPassword != hex.EncodeToString(wantHash[:]) {
		t.Fatalf("password = %q, want single SHA-256", gotPassword)
	}
	if result.Token != "session-token" || gotAppID != "300" || gotTenantCode != "tenant-a" {
		t.Fatalf("unexpected login result/header: %#v app=%q tenant=%q", result, gotAppID, gotTenantCode)
	}
}

// TestResolveTokenUsesEnvironmentPasswordWithoutWritingProfile 验证 Sandbox 账号密码无需配置文件且不落盘。
func TestResolveTokenUsesEnvironmentPasswordWithoutWritingProfile(t *testing.T) {
	home := t.TempDir()
	clearRuntimeEnv(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":200,"msg":"","data":{"token":{"accessToken":"runtime-token","accessExpire":"4102444800"}}}`))
	}))
	defer server.Close()
	t.Setenv("HOME", home)
	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "300")
	t.Setenv("UR_TENANT_CODE", "tenant-a")
	t.Setenv("UR_ACCOUNT", "admin")
	t.Setenv("UR_PASSWORD", "raw-password")

	token, err := ResolveToken(context.Background())
	if err != nil {
		t.Fatalf("ResolveToken error: %v", err)
	}
	if token != "runtime-token" {
		t.Fatalf("token = %q", token)
	}
	if _, err := os.Stat(filepath.Join(home, ".ur", "config.json")); !os.IsNotExist(err) {
		t.Fatalf("runtime credentials were persisted: %v", err)
	}
}

// TestResolveTokenPersistsLegacyPasswordSession 验证旧账号密码 profile 首次调用会补充 Token。
func TestResolveTokenPersistsLegacyPasswordSession(t *testing.T) {
	clearRuntimeEnv(t)
	home := t.TempDir()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":200,"msg":"","data":{"token":{"accessToken":"legacy-token","accessExpire":"4102444800"}}}`))
	}))
	defer server.Close()
	t.Setenv("HOME", home)
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "legacy",
		Profiles: map[string]config.Profile{
			"legacy": {
				BaseURL: server.URL, AppID: "300", TenantCode: "tenant-a",
				Account: "admin", Password: "raw-password",
			},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}

	token, err := ResolveToken(context.Background())
	if err != nil {
		t.Fatalf("ResolveToken error: %v", err)
	}
	if token != "legacy-token" {
		t.Fatalf("token = %q", token)
	}
	cfg, err := config.ReadConfig()
	if err != nil {
		t.Fatalf("ReadConfig error: %v", err)
	}
	got := cfg.Profiles["legacy"]
	if got.Token != "legacy-token" || got.Account != "admin" || got.Password != "raw-password" {
		t.Fatalf("legacy profile was not updated safely: %#v", got)
	}
}

// TestResolveAuthHeadersFallsBackToLegacyAccessKey 验证旧混合 profile 的密码被拒绝时仍可使用 AK/SK。
func TestResolveAuthHeadersFallsBackToLegacyAccessKey(t *testing.T) {
	clearRuntimeEnv(t)
	t.Setenv("HOME", t.TempDir())
	loginCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginCalls++
		_, _ = w.Write([]byte(`{"code":401,"msg":"账号或密码错误","data":null}`))
	}))
	defer server.Close()
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "legacy",
		Profiles: map[string]config.Profile{
			"legacy": {
				BaseURL: server.URL, AppID: "300", TenantCode: "tenant-a",
				Account: "admin", Password: "bad-password", AccessKey: "ak_test", AccessSecret: "sk_test",
			},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}

	headers, err := ResolveAuthHeaders(context.Background())
	if err != nil {
		t.Fatalf("ResolveAuthHeaders error: %v", err)
	}
	if headers["Authorization"] == "" || headers["token"] != "" {
		t.Fatalf("unexpected fallback headers: %#v", headers)
	}
	if loginCalls != 1 {
		t.Fatalf("password login calls = %d, want 1 before AK/SK fallback", loginCalls)
	}
}

// TestResolveAuthHeadersDoesNotFallbackOnLoginBusinessError 验证登录业务错误不会误切换到遗留 AK/SK。
func TestResolveAuthHeadersDoesNotFallbackOnLoginBusinessError(t *testing.T) {
	clearRuntimeEnv(t)
	t.Setenv("HOME", t.TempDir())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":1001001,"msg":"该企业未绑定该应用，无法登录","data":null}`))
	}))
	defer server.Close()
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "legacy",
		Profiles: map[string]config.Profile{
			"legacy": {
				BaseURL: server.URL, AppID: "300", TenantCode: "tenant-a",
				Account: "admin", Password: "raw-password", AccessKey: "ak_test", AccessSecret: "sk_test",
			},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}

	headers, err := ResolveAuthHeaders(context.Background())
	if err == nil || headers != nil || !strings.Contains(err.Error(), "该企业未绑定该应用") {
		t.Fatalf("headers=%#v error=%v", headers, err)
	}
}

// clearRuntimeEnv 清理可能影响认证来源选择的环境变量。
func clearRuntimeEnv(t *testing.T) {
	t.Helper()
	for _, key := range []string{
		"UR_BASE_URL", "UR_APP_ID", "UR_TENANT_CODE", "UR_TOKEN", "UR_ACCOUNT", "UR_PASSWORD",
		"UR_USER_ID", "UR_ACCESS_KEY", "UR_ACCESS_SECRET", "UR_PROFILE",
	} {
		t.Setenv(key, "")
	}
}
