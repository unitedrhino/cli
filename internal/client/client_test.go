package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
