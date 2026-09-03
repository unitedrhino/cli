// setup_test.go — 绑定码登记（InitSetup）单元测试，用 httptest 模拟后端，不依赖真实环境
package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestInitSetup_Success 校验登记成功与请求体格式
func TestInitSetup_Success(t *testing.T) {
	var gotPath, gotCode string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		gotCode = body["setupCode"]
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"code":200,"msg":"","data":{"success":true}}`))
	}))
	defer srv.Close()

	if err := InitSetup(context.Background(), srv.URL, "ABCD2345"); err != nil {
		t.Fatalf("InitSetup: %v", err)
	}
	if gotPath != "/api/v1/system/user/self/thirdparty/setup-init" {
		t.Errorf("path = %q", gotPath)
	}
	if gotCode != "ABCD2345" {
		t.Errorf("body setupCode = %q", gotCode)
	}
}

// TestInitSetup_NotFound 旧版后端返回 404 时应返回可识别的 ErrSetupInitUnsupported
func TestInitSetup_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer srv.Close()

	err := InitSetup(context.Background(), srv.URL, "ABCD2345")
	if !errors.Is(err, ErrSetupInitUnsupported) {
		t.Fatalf("want ErrSetupInitUnsupported, got %v", err)
	}
}

// TestInitSetup_APIError 后端返回非 200 code 时应报错（携带后端 msg）
func TestInitSetup_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"code":500,"msg":"内部错误","data":null}`))
	}))
	defer srv.Close()

	err := InitSetup(context.Background(), srv.URL, "ABCD2345")
	if err == nil || errors.Is(err, ErrSetupInitUnsupported) {
		t.Fatalf("api error should not be ErrSetupInitUnsupported, got %v", err)
	}
}
