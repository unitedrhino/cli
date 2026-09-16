// login_test.go — ur login 三种认证入口的命令级回归测试。
package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitee.com/unitedrhino/cli/internal/config"
	"github.com/spf13/cobra"
)

// TestRunLoginPasswordReplacesConflictingCredentials 验证账号密码登录立即换取 Token 并清理旧 AK/SK。
func TestRunLoginPasswordReplacesConflictingCredentials(t *testing.T) {
	clearLoginEnv(t)
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_APP_ID", "300")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/system/user/self/login" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"","data":{"token":{"accessToken":"session-token","accessExpire":"4102444800"}}}`))
	}))
	defer server.Close()
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "default",
		Profiles: map[string]config.Profile{
			"default": {AccessKey: "old-ak", AccessSecret: "old-sk", UserID: 9},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}

	loginOpts = loginOptions{
		method: "password", jsonMode: true, baseURL: server.URL, tenantCode: "tenant-a",
		account: "admin", password: "raw-password",
	}
	cmd, stdout, stderr := newLoginTestCommand()
	if err := runLogin(cmd, nil); err != nil {
		t.Fatalf("runLogin error: %v stderr=%s", err, stderr.String())
	}
	assertJSONLoginResult(t, stdout.String(), "password")
	cfg, err := config.ReadConfig()
	if err != nil {
		t.Fatalf("ReadConfig error: %v", err)
	}
	got := cfg.Profiles["default"]
	if got.Token != "session-token" || got.Account != "admin" || got.Password != "raw-password" {
		t.Fatalf("password profile = %#v", got)
	}
	if got.AccessKey != "" || got.AccessSecret != "" || got.UserID != 0 {
		t.Fatalf("old AK/SK not cleared: %#v", got)
	}
}

// TestRunLoginAccessKeyWithoutUserID 验证直接输入 AK/SK 不要求 userID。
func TestRunLoginAccessKeyWithoutUserID(t *testing.T) {
	clearLoginEnv(t)
	t.Setenv("HOME", t.TempDir())
	t.Setenv("UR_APP_ID", "300")
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "default",
		Profiles: map[string]config.Profile{
			"default": {Token: "old-token", Account: "old-account", Password: "old-password"},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			t.Fatal("missing Authorization header")
		}
		if r.Header.Get("tenant-code") != "tenant-a" {
			t.Fatalf("tenant-code = %q", r.Header.Get("tenant-code"))
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"","data":{"userID":"42"}}`))
	}))
	defer server.Close()

	loginOpts = loginOptions{
		method: "aksk", jsonMode: true, baseURL: server.URL, tenantCode: "tenant-a",
		accessKey: "ak_test", accessSecret: "sk_test",
	}
	cmd, stdout, stderr := newLoginTestCommand()
	if err := runLogin(cmd, nil); err != nil {
		t.Fatalf("runLogin error: %v stderr=%s", err, stderr.String())
	}
	assertJSONLoginResult(t, stdout.String(), "aksk")
	profile, err := config.CurrentProfile()
	if err != nil {
		t.Fatalf("CurrentProfile error: %v", err)
	}
	if profile.AccessKey != "ak_test" || profile.AccessSecret != "sk_test" || profile.UserID != 42 {
		t.Fatalf("AK/SK profile = %#v", profile)
	}
	if profile.Token != "" || profile.Account != "" || profile.Password != "" {
		t.Fatalf("old password credentials not cleared: %#v", profile)
	}
}

// TestResolveLoginSecretPrecedence 验证明文参数、stdin、环境变量的固定优先级。
func TestResolveLoginSecretPrecedence(t *testing.T) {
	t.Setenv("UR_PASSWORD", "env-secret")
	got, err := resolveLoginSecret(strings.NewReader("stdin-secret\n"), "flag-secret", false, "UR_PASSWORD", "password")
	if err != nil || got != "flag-secret" {
		t.Fatalf("flag precedence got=%q err=%v", got, err)
	}
	got, err = resolveLoginSecret(strings.NewReader("stdin-secret\n"), "flag-secret", true, "UR_PASSWORD", "password")
	if err != nil || got != "flag-secret" {
		t.Fatalf("flag must override stdin got=%q err=%v", got, err)
	}
	got, err = resolveLoginSecret(strings.NewReader("stdin-secret\n"), "", true, "UR_PASSWORD", "password")
	if err != nil || got != "stdin-secret" {
		t.Fatalf("stdin precedence got=%q err=%v", got, err)
	}
	got, err = resolveLoginSecret(strings.NewReader("unused"), "", false, "UR_PASSWORD", "password")
	if err != nil || got != "env-secret" {
		t.Fatalf("environment fallback got=%q err=%v", got, err)
	}
}

// TestRunLoginRejectsUnknownMethod 验证未知认证方式返回稳定错误。
func TestRunLoginRejectsUnknownMethod(t *testing.T) {
	clearLoginEnv(t)
	loginOpts = loginOptions{method: "unknown", jsonMode: true}
	cmd, _, _ := newLoginTestCommand()
	if err := runLogin(cmd, nil); err == nil || !strings.Contains(err.Error(), "unsupported login method") {
		t.Fatalf("runLogin error = %v", err)
	}
}

// TestResolveLoginBaseURLUsesExistingProfile 验证升级用户继续使用原私有化平台地址。
func TestResolveLoginBaseURLUsesExistingProfile(t *testing.T) {
	clearLoginEnv(t)
	t.Setenv("HOME", t.TempDir())
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "default",
		Profiles: map[string]config.Profile{
			"default": {BaseURL: "http://private.example/", AppID: "300", TenantCode: "tenant-a", Token: "token"},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}
	got, err := resolveLoginBaseURL("")
	if err != nil || got != "http://private.example" {
		t.Fatalf("resolveLoginBaseURL = %q, %v", got, err)
	}
}

// TestRunSetupUsesPasswordLogin 验证旧 setup 输入顺序保持兼容并立即完成账号密码登录。
func TestRunSetupUsesPasswordLogin(t *testing.T) {
	clearLoginEnv(t)
	t.Setenv("HOME", t.TempDir())
	t.Setenv("UR_APP_ID", "300")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":200,"msg":"","data":{"token":{"accessToken":"setup-token","accessExpire":"4102444800"}}}`))
	}))
	defer server.Close()

	cmd, _, _ := newLoginTestCommand()
	cmd.SetIn(strings.NewReader(server.URL + "\n\ntenant-a\nadmin\nraw-password\n"))
	if err := runSetup(cmd, nil); err != nil {
		t.Fatalf("runSetup error: %v", err)
	}
	profile, err := config.CurrentProfile()
	if err != nil {
		t.Fatalf("CurrentProfile error: %v", err)
	}
	if profile.Token != "setup-token" || profile.Account != "admin" || profile.Password != "raw-password" {
		t.Fatalf("setup profile = %#v", profile)
	}
}

// TestRunCheckJSONReportsEnvironmentAuth 验证 AI 可从 check 输出识别 Sandbox 认证来源和方式。
func TestRunCheckJSONReportsEnvironmentAuth(t *testing.T) {
	clearLoginEnv(t)
	t.Setenv("HOME", t.TempDir())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":200,"msg":"","data":{"userID":"42"}}`))
	}))
	defer server.Close()
	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "300")
	t.Setenv("UR_TENANT_CODE", "tenant-a")
	t.Setenv("UR_ACCESS_KEY", "ak_test")
	t.Setenv("UR_ACCESS_SECRET", "sk_test")
	checkOpts.jsonMode = true
	cmd, stdout, _ := newLoginTestCommand()
	if err := runCheck(cmd, nil); err != nil {
		t.Fatalf("runCheck error: %v", err)
	}
	var result map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(stdout.String())), &result); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	if result["auth_source"] != "environment" || result["auth_method"] != "aksk" {
		t.Fatalf("auth metadata = %#v", result)
	}
	if strings.Contains(stdout.String(), "ak_test") || strings.Contains(stdout.String(), "sk_test") {
		t.Fatalf("secret leaked in check output: %s", stdout.String())
	}
}

// TestRunCheckJSONReportsPartialEnvironmentCredential 验证 env-only 半套凭据返回明确且脱敏的 JSON 错误。
func TestRunCheckJSONReportsPartialEnvironmentCredential(t *testing.T) {
	clearLoginEnv(t)
	t.Setenv("HOME", t.TempDir())
	t.Setenv("UR_BASE_URL", "http://sandbox.invalid")
	t.Setenv("UR_APP_ID", "300")
	t.Setenv("UR_TENANT_CODE", "tenant-a")
	t.Setenv("UR_ACCESS_KEY", "ak_only")
	checkOpts.jsonMode = true
	cmd, stdout, _ := newLoginTestCommand()
	err := runCheck(cmd, nil)
	if err == nil {
		t.Fatal("expected partial credential error")
	}
	var result map[string]any
	if jsonErr := json.Unmarshal([]byte(strings.TrimSpace(stdout.String())), &result); jsonErr != nil {
		t.Fatalf("invalid JSON output %q: %v", stdout.String(), jsonErr)
	}
	if result["auth_status"] != "error" || !strings.Contains(result["auth_status_msg"].(string), "UR_ACCESS_KEY and UR_ACCESS_SECRET") {
		t.Fatalf("unexpected error result: %#v", result)
	}
	if strings.Contains(stdout.String(), "ak_only") {
		t.Fatalf("credential leaked in JSON output: %s", stdout.String())
	}
}

// newLoginTestCommand 创建隔离输入输出的 Cobra 命令。
func newLoginTestCommand() (*cobra.Command, *bytes.Buffer, *bytes.Buffer) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := &cobra.Command{}
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)
	cmd.SetIn(strings.NewReader(""))
	cmd.SetContext(context.Background())
	return cmd, stdout, stderr
}

// assertJSONLoginResult 校验登录 JSON 输出不泄露敏感字段。
func assertJSONLoginResult(t *testing.T, raw, method string) {
	t.Helper()
	var result map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(raw)), &result); err != nil {
		t.Fatalf("invalid JSON output %q: %v", raw, err)
	}
	if result["status"] != "ok" || result["method"] != method {
		t.Fatalf("unexpected result: %#v", result)
	}
	for _, secret := range []string{"raw-password", "session-token", "sk_test"} {
		if strings.Contains(raw, secret) {
			t.Fatalf("secret leaked in JSON output: %q", raw)
		}
	}
}

// clearLoginEnv 清理可能影响命令测试的认证环境变量。
func clearLoginEnv(t *testing.T) {
	t.Helper()
	for _, key := range []string{
		"UR_BASE_URL", "UR_APP_ID", "UR_TENANT_CODE", "UR_TOKEN", "UR_ACCOUNT", "UR_PASSWORD",
		"UR_USER_ID", "UR_ACCESS_KEY", "UR_ACCESS_SECRET", "UR_PROFILE",
	} {
		t.Setenv(key, "")
	}
}
