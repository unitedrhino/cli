package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/auth"
	"gitee.com/unitedrhino/cli/internal/config"
)

type APIRequest struct {
	Path    string
	Body    map[string]any
	Headers map[string]string
}

type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func DoAPI(ctx context.Context, req APIRequest) (APIResponse, error) {
	baseURL, err := config.GetBaseURL()
	if err != nil {
		return APIResponse{}, err
	}
	appID, err := config.GetAppID()
	if err != nil {
		return APIResponse{}, err
	}
	tenantCode, err := config.GetTenantCode()
	if err != nil {
		return APIResponse{}, err
	}
	if req.Body == nil {
		req.Body = map[string]any{}
	}
	path, body, err := normalizeRequestPathAndBody(req.Path, req.Body)
	if err != nil {
		return APIResponse{}, err
	}
	req.Path = path
	req.Body = body
	rawBody, err := json.Marshal(req.Body)
	if err != nil {
		return APIResponse{}, fmt.Errorf("marshal body: %w", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+req.Path, bytes.NewReader(rawBody))
	if err != nil {
		return APIResponse{}, fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("app-id", appID)
	httpReq.Header.Set("tenant-code", tenantCode)
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}
	if traceparent := strings.TrimSpace(os.Getenv("UR_TRACEPARENT")); traceparent != "" {
		httpReq.Header.Set("traceparent", traceparent)
	}
	if tracestate := strings.TrimSpace(os.Getenv("UR_TRACESTATE")); tracestate != "" {
		httpReq.Header.Set("tracestate", tracestate)
	}
	authHeaders, err := auth.ResolveAuthHeaders(ctx)
	if err != nil {
		return APIResponse{}, err
	}
	for key, value := range authHeaders {
		httpReq.Header.Set(key, value)
	}
	resp, err := (&http.Client{Timeout: 30 * time.Second}).Do(httpReq)
	if err != nil {
		return APIResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()
	rawResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, fmt.Errorf("read response: %w", err)
	}
	if reqTraceparent := httpReq.Header.Get("traceparent"); reqTraceparent != "" || resp.Header.Get("traceparent") != "" {
		log.Printf("ur-api trace path=%s reqTraceparent=%s respTraceparent=%s status=%d", req.Path, reqTraceparent, resp.Header.Get("traceparent"), resp.StatusCode)
	}
	var out APIResponse
	if err := json.Unmarshal(rawResp, &out); err != nil {
		return APIResponse{}, fmt.Errorf("decode response status=%d body=%s: %w", resp.StatusCode, strings.TrimSpace(string(rawResp)), err)
	}
	return out, nil
}

func normalizeRequestPathAndBody(path string, body map[string]any) (string, map[string]any, error) {
	parsed, err := url.Parse(path)
	if err != nil {
		return "", nil, fmt.Errorf("parse path: %w", err)
	}
	if parsed.RawQuery == "" {
		return normalizeLegacyInfoGetPath(path, body), body, nil
	}
	normalized := make(map[string]any, len(body)+len(parsed.Query()))
	for k, v := range body {
		normalized[k] = v
	}
	for key, values := range parsed.Query() {
		if _, exists := normalized[key]; exists || len(values) == 0 {
			continue
		}
		normalized[key] = values[0]
	}
	parsed.RawQuery = ""
	return normalizeLegacyInfoGetPath(parsed.String(), normalized), normalized, nil
}

func normalizeLegacyInfoGetPath(path string, body map[string]any) string {
	if !strings.HasSuffix(path, "/info/get") {
		return path
	}
	if bodyContainsSingularID(body) {
		return strings.TrimSuffix(path, "/get") + "/get-one"
	}
	return strings.TrimSuffix(path, "/get") + "/get-list"
}

func bodyContainsSingularID(body map[string]any) bool {
	for key := range body {
		trimmed := strings.TrimSpace(key)
		if trimmed == "" {
			continue
		}
		lower := strings.ToLower(trimmed)
		if lower == "id" || strings.HasSuffix(lower, "id") {
			if lower == "ids" || strings.HasSuffix(lower, "ids") {
				continue
			}
			return true
		}
	}
	return false
}
