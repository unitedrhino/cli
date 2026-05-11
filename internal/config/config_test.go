package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeEnvProfileUsesEnvWithoutProfileFile(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
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
	if _, err := os.Stat(filepath.Join(t.TempDir(), ".ur", "config.json")); !os.IsNotExist(err) {
		t.Fatalf("unexpected config file presence: %v", err)
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
