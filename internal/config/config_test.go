package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeEnvProfileUsesEnvWithoutProfileFile(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_BASE_URL", "http://127.0.0.1:7777")
	t.Setenv("UR_APP_ID", "100")
	t.Setenv("UR_TENANT_CODE", "platform")
	t.Setenv("UR_TOKEN", "runtime-token")

	profile, ok := RuntimeProfileFromEnv()
	if !ok {
		t.Fatal("expected runtime env profile")
	}
	if profile.BaseURL != "http://127.0.0.1:7777" {
		t.Fatalf("baseURL = %q", profile.BaseURL)
	}
	if profile.Token != "runtime-token" {
		t.Fatalf("token = %q", profile.Token)
	}
	if _, err := os.Stat(filepath.Join(home, ".ur", "config.json")); !os.IsNotExist(err) {
		t.Fatalf("unexpected config file presence: %v", err)
	}
}

// TestResolveAuthContextKeepsLegacySingleCredentialProfiles 验证旧版纯 Token 与纯 AK/SK profile 行为不变。
func TestResolveAuthContextKeepsLegacySingleCredentialProfiles(t *testing.T) {
	for _, tc := range []struct {
		name    string
		profile Profile
		method  AuthMethod
	}{
		{name: "token", profile: Profile{Token: "session-token"}, method: AuthMethodToken},
		{name: "aksk", profile: Profile{AccessKey: "ak_test", AccessSecret: "sk_test"}, method: AuthMethodAccessKey},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("HOME", t.TempDir())
			for _, key := range []string{"UR_BASE_URL", "UR_TOKEN", "UR_ACCOUNT", "UR_PASSWORD", "UR_ACCESS_KEY", "UR_ACCESS_SECRET", "UR_USER_ID"} {
				t.Setenv(key, "")
			}
			tc.profile.BaseURL = "http://profile"
			tc.profile.AppID = "300"
			tc.profile.TenantCode = "tenant-a"
			if err := WriteConfig(Config{CurrentProfile: "legacy", Profiles: map[string]Profile{"legacy": tc.profile}}); err != nil {
				t.Fatalf("WriteConfig error: %v", err)
			}
			ctx, err := ResolveAuthContext()
			if err != nil {
				t.Fatalf("ResolveAuthContext error: %v", err)
			}
			if ctx.Source != AuthSourceProfile || ctx.Method != tc.method {
				t.Fatalf("source/method = %q/%q", ctx.Source, ctx.Method)
			}
		})
	}
}

func TestTenantCodePrefersEnvOverProfile(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	cfg := Config{
		CurrentProfile: "default",
		Profiles: map[string]Profile{
			"default": {
				BaseURL:    "http://127.0.0.1:7777",
				AppID:      "77",
				TenantCode: "from-profile",
			},
		},
	}
	if err := WriteConfig(cfg); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}
	t.Setenv("UR_TENANT_CODE", "from-env")
	got, err := GetTenantCode()
	if err != nil {
		t.Fatalf("GetTenantCode error: %v", err)
	}
	if got != "from-env" {
		t.Fatalf("tenantCode = %q, want from-env", got)
	}
}

// TestResolveAuthContextAllowsAccessKeyWithoutUserID 验证 Sandbox 仅注入 AK/SK 时无需 userID。
func TestResolveAuthContextAllowsAccessKeyWithoutUserID(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("UR_BASE_URL", "http://127.0.0.1:7777")
	t.Setenv("UR_APP_ID", "100")
	t.Setenv("UR_TENANT_CODE", "platform")
	t.Setenv("UR_ACCESS_KEY", "ak_test")
	t.Setenv("UR_ACCESS_SECRET", "sk_test")

	ctx, err := ResolveAuthContext()
	if err != nil {
		t.Fatalf("ResolveAuthContext error: %v", err)
	}
	if ctx.AccessKey != "ak_test" || ctx.AccessSecret != "sk_test" {
		t.Fatalf("unexpected AK/SK context: %#v", ctx)
	}
	if ctx.Source != AuthSourceEnvironment || ctx.Method != AuthMethodAccessKey {
		t.Fatalf("source/method = %q/%q", ctx.Source, ctx.Method)
	}
}

// TestResolveAuthContextRejectsPartialEnvironmentCredential 验证半套 Sandbox 凭据不会回退磁盘配置。
func TestResolveAuthContextRejectsPartialEnvironmentCredential(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	if err := WriteConfig(Config{
		CurrentProfile: "default",
		Profiles: map[string]Profile{
			"default": {BaseURL: "http://profile", AppID: "77", TenantCode: "profile", Token: "disk-token"},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}
	t.Setenv("UR_BASE_URL", "http://sandbox")
	t.Setenv("UR_APP_ID", "100")
	t.Setenv("UR_TENANT_CODE", "sandbox")
	t.Setenv("UR_ACCESS_KEY", "ak_only")

	_, err := ResolveAuthContext()
	if err == nil || err.Error() != "UR_ACCESS_KEY and UR_ACCESS_SECRET must be set together" {
		t.Fatalf("ResolveAuthContext error = %v", err)
	}
}

// TestResolveAuthContextPrefersLegacyPasswordBeforeAccessKey 验证旧版混合 profile 优先账号密码。
func TestResolveAuthContextPrefersLegacyPasswordBeforeAccessKey(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	if err := WriteConfig(Config{
		CurrentProfile: "legacy",
		Profiles: map[string]Profile{
			"legacy": {
				BaseURL: "http://127.0.0.1:7777", AppID: "300", TenantCode: "tenant-a",
				Account: "admin", Password: "raw-password", AccessKey: "old-ak", AccessSecret: "old-sk",
			},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}

	ctx, err := ResolveAuthContext()
	if err != nil {
		t.Fatalf("ResolveAuthContext error: %v", err)
	}
	if ctx.Method != AuthMethodPassword || ctx.Account != "admin" || ctx.Password != "raw-password" {
		t.Fatalf("unexpected legacy context: %#v", ctx)
	}
	if ctx.AccessKey != "old-ak" || ctx.AccessSecret != "old-sk" {
		t.Fatalf("legacy fallback credentials were lost: %#v", ctx)
	}
}

// TestReplaceProfileAuthAndUpdateSessionToken 验证显式切换认证会清除冲突字段，而刷新只更新 Token。
func TestReplaceProfileAuthAndUpdateSessionToken(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	if err := WriteConfig(Config{
		CurrentProfile: "default",
		Profiles: map[string]Profile{
			"default": {
				BaseURL: "http://old", AppID: "77", TenantCode: "old", Role: "admin",
				Token: "old-token", Account: "old-account", Password: "old-password",
				UserID: 9, AccessKey: "old-ak", AccessSecret: "old-sk",
			},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}
	if err := os.Chmod(ConfigPath(), 0o640); err != nil {
		t.Fatalf("Chmod error: %v", err)
	}

	if err := ReplaceProfileAuth(Profile{
		BaseURL: "http://new", AppID: "300", TenantCode: "tenant-new", Role: "admin",
		Token: "session-token", Account: "new-account", Password: "new-password",
	}, AuthMethodPassword); err != nil {
		t.Fatalf("ReplaceProfileAuth error: %v", err)
	}
	if err := UpdateSessionToken("refreshed-token"); err != nil {
		t.Fatalf("UpdateSessionToken error: %v", err)
	}

	raw, err := os.ReadFile(filepath.Join(home, ".ur", "config.json"))
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	var cfg Config
	if err := json.Unmarshal(raw, &cfg); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	got := cfg.Profiles["default"]
	if got.Token != "refreshed-token" || got.Account != "new-account" || got.Password != "new-password" {
		t.Fatalf("password credentials not preserved: %#v", got)
	}
	if got.AccessKey != "" || got.AccessSecret != "" || got.UserID != 0 {
		t.Fatalf("conflicting AK/SK credentials not cleared: %#v", got)
	}
	info, err := os.Stat(ConfigPath())
	if err != nil {
		t.Fatalf("Stat error: %v", err)
	}
	if info.Mode().Perm() != 0o640 {
		t.Fatalf("config permissions changed to %o", info.Mode().Perm())
	}
}

// TestUpdateSessionTokenDoesNotPersistRuntimeEnvironment 验证 Sandbox 登录不会创建本地配置。
func TestUpdateSessionTokenDoesNotPersistRuntimeEnvironment(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_BASE_URL", "http://sandbox")
	t.Setenv("UR_APP_ID", "100")
	t.Setenv("UR_TENANT_CODE", "platform")
	t.Setenv("UR_ACCOUNT", "admin")
	t.Setenv("UR_PASSWORD", "raw-password")

	if err := UpdateSessionToken("runtime-token"); err != nil {
		t.Fatalf("UpdateSessionToken error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(home, ".ur", "config.json")); !os.IsNotExist(err) {
		t.Fatalf("runtime credentials were persisted: %v", err)
	}
}

// TestEnvironmentTokenOverridesStoredCredentialsWithoutBaseURL 验证旧用法可只覆盖认证变量并复用 profile 上下文。
func TestEnvironmentTokenOverridesStoredCredentialsWithoutBaseURL(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	if err := WriteConfig(Config{
		CurrentProfile: "default",
		Profiles: map[string]Profile{
			"default": {BaseURL: "http://profile", AppID: "77", TenantCode: "tenant-a", Account: "admin", Password: "raw-password"},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}
	t.Setenv("UR_TOKEN", "environment-token")

	ctx, err := ResolveAuthContext()
	if err != nil {
		t.Fatalf("ResolveAuthContext error: %v", err)
	}
	if ctx.Source != AuthSourceEnvironment || ctx.Method != AuthMethodToken || ctx.Token != "environment-token" {
		t.Fatalf("unexpected auth context: %#v", ctx)
	}
}

// TestRuntimeEnvironmentRequiresAppAndTenantContext 验证 Sandbox 不会静默回退错误的默认上下文。
func TestRuntimeEnvironmentRequiresAppAndTenantContext(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("UR_BASE_URL", "http://sandbox")
	t.Setenv("UR_TOKEN", "runtime-token")
	t.Setenv("UR_APP_ID", "")
	t.Setenv("UR_TENANT_CODE", "")

	if _, err := GetAppID(); err == nil || err.Error() != "UR_APP_ID is required in runtime environment" {
		t.Fatalf("GetAppID error = %v", err)
	}
	if _, err := GetTenantCode(); err == nil || err.Error() != "UR_TENANT_CODE is required in runtime environment" {
		t.Fatalf("GetTenantCode error = %v", err)
	}
}
