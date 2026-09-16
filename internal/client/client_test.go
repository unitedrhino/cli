package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitee.com/unitedrhino/cli/internal/config"
)

func TestDoAPIUsesRuntimeEnvAuthAndHeaders(t *testing.T) {
	var gotBody map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("app-id") != "100" {
			t.Fatalf("app-id = %q", r.Header.Get("app-id"))
		}
		if r.Header.Get("tenant-code") != "platform" {
			t.Fatalf("tenant-code = %q", r.Header.Get("tenant-code"))
		}
		if r.Header.Get("token") != "runtime-token" {
			t.Fatalf("token = %q", r.Header.Get("token"))
		}
		if r.Header.Get("traceparent") != "00-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa-bbbbbbbbbbbbbbbb-01" {
			t.Fatalf("traceparent = %q", r.Header.Get("traceparent"))
		}
		if r.Header.Get("tracestate") != "vendor=test" {
			t.Fatalf("tracestate = %q", r.Header.Get("tracestate"))
		}
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"success","data":{"ok":true}}`))
	}))
	defer server.Close()

	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "100")
	t.Setenv("UR_TENANT_CODE", "platform")
	t.Setenv("UR_TOKEN", "runtime-token")
	t.Setenv("UR_TRACEPARENT", "00-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa-bbbbbbbbbbbbbbbb-01")
	t.Setenv("UR_TRACESTATE", "vendor=test")

	resp, err := DoAPI(context.Background(), APIRequest{
		Path: "/api/v1/system/user/info/get-list",
		Body: map[string]any{"page": map[string]any{"page": 1}},
	})
	if err != nil {
		t.Fatalf("DoAPI error: %v", err)
	}
	if resp.Code != 200 {
		t.Fatalf("resp.Code = %d", resp.Code)
	}
	if gotBody["page"] == nil {
		t.Fatalf("body not sent: %+v", gotBody)
	}
}

func TestDoAPINormalizesQueryStringIntoBody(t *testing.T) {
	var gotPath string
	var gotBody map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if r.URL.RawQuery != "" {
			t.Fatalf("expected normalized request without query string, got %q", r.URL.RawQuery)
		}
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"success","data":{"ok":true}}`))
	}))
	defer server.Close()

	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "100")
	t.Setenv("UR_TENANT_CODE", "platform")
	t.Setenv("UR_TOKEN", "runtime-token")

	_, err := DoAPI(context.Background(), APIRequest{
		Path: "/api/v1/system/role/info/get-one?id=1",
		Body: map[string]any{},
	})
	if err != nil {
		t.Fatalf("DoAPI error: %v", err)
	}
	if gotPath != "/api/v1/system/role/info/get-one" {
		t.Fatalf("path = %q", gotPath)
	}
	if gotBody["id"] != "1" {
		t.Fatalf("body = %+v", gotBody)
	}
}

func TestDoAPINormalizesLegacyInfoGetPathToGetList(t *testing.T) {
	var gotPath string
	var gotBody map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"success","data":{"ok":true}}`))
	}))
	defer server.Close()

	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "100")
	t.Setenv("UR_TENANT_CODE", "platform")
	t.Setenv("UR_TOKEN", "runtime-token")

	_, err := DoAPI(context.Background(), APIRequest{
		Path: "/api/v1/system/role/info/get",
		Body: map[string]any{"codes": []string{"supper"}},
	})
	if err != nil {
		t.Fatalf("DoAPI error: %v", err)
	}
	if gotPath != "/api/v1/system/role/info/get-list" {
		t.Fatalf("path = %q", gotPath)
	}
	if gotBody["codes"] == nil {
		t.Fatalf("body = %+v", gotBody)
	}
}

// TestIsAuthFailureDoesNotTreatBusinessPermissionAsExpired 验证业务权限错误不会触发账号密码刷新。
func TestIsAuthFailureDoesNotTreatBusinessPermissionAsExpired(t *testing.T) {
	if isAuthFailure(APIResponse{Code: 403, Msg: "权限不足"}) {
		t.Fatal("business permission failure must not trigger auth fallback")
	}
	if isAuthFailure(APIResponse{Code: 500, Msg: "invalid request payload"}) {
		t.Fatal("generic invalid message must not trigger auth fallback")
	}
	for _, resp := range []APIResponse{
		{Code: 401, Msg: "unauthorized"},
		{Code: 200, Msg: "token 已过期"},
		{Code: 200, Msg: "登录状态过期"},
		{Code: 1000007, Msg: "尚未登录:认证失败"},
		{Code: 1000002, Msg: "登录过期,请退出重新登录"},
	} {
		if !isAuthFailure(resp) {
			t.Fatalf("expected auth failure: %#v", resp)
		}
	}
}

// TestDoAPIRefreshesExpiredLegacyTokenWithPassword 验证旧 Token 过期后只更新 Session Token 并重试。
func TestDoAPIRefreshesExpiredLegacyTokenWithPassword(t *testing.T) {
	home := t.TempDir()
	for _, key := range []string{
		"UR_BASE_URL", "UR_APP_ID", "UR_TENANT_CODE", "UR_TOKEN", "UR_ACCOUNT", "UR_PASSWORD",
		"UR_USER_ID", "UR_ACCESS_KEY", "UR_ACCESS_SECRET", "UR_PROFILE",
	} {
		t.Setenv(key, "")
	}
	t.Setenv("HOME", home)
	apiCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/system/user/self/login" {
			_, _ = w.Write([]byte(`{"code":200,"msg":"","data":{"token":{"accessToken":"fresh-token","accessExpire":"4102444800"}}}`))
			return
		}
		apiCalls++
		if r.Header.Get("token") == "fresh-token" {
			_, _ = w.Write([]byte(`{"code":200,"msg":"success","data":{"ok":true}}`))
			return
		}
		_, _ = w.Write([]byte(`{"code":401,"msg":"登录过期,请退出重新登录","data":null}`))
	}))
	defer server.Close()
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "legacy",
		Profiles: map[string]config.Profile{
			"legacy": {
				BaseURL: server.URL, AppID: "300", TenantCode: "tenant-a", Role: "admin", Token: "expired-token",
				Account: "admin", Password: "raw-password",
			},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}

	resp, err := DoAPI(context.Background(), APIRequest{Path: "/api/v1/test", Body: map[string]any{}})
	if err != nil {
		t.Fatalf("DoAPI error: %v", err)
	}
	if resp.Code != 200 || apiCalls != 2 {
		t.Fatalf("resp=%#v apiCalls=%d", resp, apiCalls)
	}
	cfg, err := config.ReadConfig()
	if err != nil {
		t.Fatalf("ReadConfig error: %v", err)
	}
	got := cfg.Profiles["legacy"]
	if cfg.CurrentProfile != "legacy" || got.Token != "fresh-token" || got.Account != "admin" || got.Password != "raw-password" || got.Role != "admin" {
		t.Fatalf("legacy profile was not safely refreshed: current=%q profile=%#v", cfg.CurrentProfile, got)
	}
}

// TestDoAPIFallsBackThroughLegacyCredentialCandidates 验证旧 Token、密码均失效时可回退有效 AK/SK。
func TestDoAPIFallsBackThroughLegacyCredentialCandidates(t *testing.T) {
	home := t.TempDir()
	for _, key := range []string{
		"UR_BASE_URL", "UR_APP_ID", "UR_TENANT_CODE", "UR_TOKEN", "UR_ACCOUNT", "UR_PASSWORD",
		"UR_USER_ID", "UR_ACCESS_KEY", "UR_ACCESS_SECRET", "UR_PROFILE",
	} {
		t.Setenv(key, "")
	}
	t.Setenv("HOME", home)
	apiCalls := 0
	loginCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v1/system/user/self/login":
			loginCalls++
			_, _ = w.Write([]byte(`{"code":401,"msg":"账号或密码错误","data":null}`))
		default:
			apiCalls++
			if strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ") {
				_, _ = w.Write([]byte(`{"code":200,"msg":"success","data":{"ok":true}}`))
				return
			}
			_, _ = w.Write([]byte(`{"code":401,"msg":"登录过期,请退出重新登录","data":null}`))
		}
	}))
	defer server.Close()
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "legacy",
		Profiles: map[string]config.Profile{
			"legacy": {
				BaseURL: server.URL, AppID: "300", TenantCode: "tenant-a", Token: "expired-token",
				Account: "admin", Password: "bad-password", AccessKey: "ak_test", AccessSecret: "sk_test",
			},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}

	resp, err := DoAPI(context.Background(), APIRequest{Path: "/api/v1/test", Body: map[string]any{}})
	if err != nil {
		t.Fatalf("DoAPI error: %v", err)
	}
	if resp.Code != 200 || apiCalls != 2 || loginCalls != 1 {
		t.Fatalf("resp=%#v apiCalls=%d loginCalls=%d", resp, apiCalls, loginCalls)
	}
	cfg, err := config.ReadConfig()
	if err != nil {
		t.Fatalf("ReadConfig error: %v", err)
	}
	if got := cfg.Profiles["legacy"]; got.Token != "expired-token" || got.Account != "admin" || got.AccessKey != "ak_test" {
		t.Fatalf("legacy profile was unexpectedly rewritten: %#v", got)
	}
}

// TestDoAPIStopsFallbackOnPasswordNetworkError 验证密码刷新遇到非凭据错误时不会继续尝试 AK/SK。
func TestDoAPIStopsFallbackOnPasswordNetworkError(t *testing.T) {
	home := t.TempDir()
	for _, key := range []string{
		"UR_BASE_URL", "UR_APP_ID", "UR_TENANT_CODE", "UR_TOKEN", "UR_ACCOUNT", "UR_PASSWORD",
		"UR_USER_ID", "UR_ACCESS_KEY", "UR_ACCESS_SECRET", "UR_PROFILE",
	} {
		t.Setenv(key, "")
	}
	t.Setenv("HOME", home)
	apiCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/system/user/self/login" {
			_, _ = w.Write([]byte(`not-json`))
			return
		}
		apiCalls++
		if r.Header.Get("Authorization") != "" {
			t.Fatal("AK/SK must not be tried after password response parse error")
		}
		_, _ = w.Write([]byte(`{"code":401,"msg":"登录过期,请退出重新登录","data":null}`))
	}))
	defer server.Close()
	if err := config.WriteConfig(config.Config{
		CurrentProfile: "legacy",
		Profiles: map[string]config.Profile{
			"legacy": {
				BaseURL: server.URL, AppID: "300", TenantCode: "tenant-a", Token: "expired-token",
				Account: "admin", Password: "raw-password", AccessKey: "ak_test", AccessSecret: "sk_test",
			},
		},
	}); err != nil {
		t.Fatalf("WriteConfig error: %v", err)
	}

	resp, err := DoAPI(context.Background(), APIRequest{Path: "/api/v1/test", Body: map[string]any{}})
	if err != nil {
		t.Fatalf("DoAPI error: %v", err)
	}
	if resp.Code != 401 || apiCalls != 1 {
		t.Fatalf("resp=%#v apiCalls=%d", resp, apiCalls)
	}
}
