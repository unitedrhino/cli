package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
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
	Debug   bool
}

type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// DoAPI 发送 API 请求，若遇到认证失败会自动刷新 token 并重试一次。
func DoAPI(ctx context.Context, req APIRequest) (APIResponse, error) {
	resp, err := doAPIOnce(ctx, req)
	if err != nil {
		return APIResponse{}, err
	}
	if isAuthFailure(resp) {
		if _, refreshErr := auth.RefreshToken(ctx); refreshErr == nil {
			if req.Debug {
				log.Println("[debug] token refreshed, retrying request...")
			}
			return doAPIOnce(ctx, req)
		}
	}
	return resp, nil
}

func doAPIOnce(ctx context.Context, req APIRequest) (APIResponse, error) {
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
	if req.Debug {
		logDebugRequest(httpReq, rawBody)
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
	if req.Debug {
		logDebugResponse(resp, rawResp)
	}
	// traceparent 日志仅在 debug 模式下输出，避免泄露到 stderr 被 sandbox runtime 捕获后暴露给 LLM
	if req.Debug {
		if reqTraceparent := httpReq.Header.Get("traceparent"); reqTraceparent != "" || resp.Header.Get("traceparent") != "" {
			log.Printf("ur-api trace path=%s reqTraceparent=%s respTraceparent=%s status=%d", req.Path, reqTraceparent, resp.Header.Get("traceparent"), resp.StatusCode)
		}
	}
	var out APIResponse
	if err := json.Unmarshal(rawResp, &out); err != nil {
		return APIResponse{}, fmt.Errorf("decode response status=%d body=%s: %w", resp.StatusCode, strings.TrimSpace(string(rawResp)), err)
	}
	return out, nil
}

// UploadFileMultipart 以 multipart/form-data 上传文件，适配 /api/v1/system/common/upload-file 等接口。
// 与 DoAPI 一致注入 app-id/tenant-code/认证头，并支持 401 自动刷新重试。
func UploadFileMultipart(ctx context.Context, path, fieldName, fileName string, fileData []byte, form map[string]string) (APIResponse, error) {
	resp, err := uploadFileMultipartOnce(ctx, path, fieldName, fileName, fileData, form)
	if err != nil {
		return APIResponse{}, err
	}
	if isAuthFailure(resp) {
		if _, refreshErr := auth.RefreshToken(ctx); refreshErr == nil {
			return uploadFileMultipartOnce(ctx, path, fieldName, fileName, fileData, form)
		}
	}
	return resp, nil
}

// uploadFileMultipartOnce 执行一次 multipart 上传（不含 401 重试）
func uploadFileMultipartOnce(ctx context.Context, path, fieldName, fileName string, fileData []byte, form map[string]string) (APIResponse, error) {
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

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	part, err := mw.CreateFormFile(fieldName, fileName)
	if err != nil {
		return APIResponse{}, fmt.Errorf("create multipart file field: %w", err)
	}
	if _, err := part.Write(fileData); err != nil {
		return APIResponse{}, fmt.Errorf("write multipart file: %w", err)
	}
	for key, value := range form {
		if err := mw.WriteField(key, value); err != nil {
			return APIResponse{}, fmt.Errorf("write multipart field %s: %w", key, err)
		}
	}
	if err := mw.Close(); err != nil {
		return APIResponse{}, fmt.Errorf("close multipart writer: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+path, &buf)
	if err != nil {
		return APIResponse{}, fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", mw.FormDataContentType())
	httpReq.Header.Set("app-id", appID)
	httpReq.Header.Set("tenant-code", tenantCode)
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

	resp, err := (&http.Client{Timeout: 60 * time.Second}).Do(httpReq)
	if err != nil {
		return APIResponse{}, fmt.Errorf("do upload request: %w", err)
	}
	defer resp.Body.Close()
	rawResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, fmt.Errorf("read upload response: %w", err)
	}
	var out APIResponse
	if err := json.Unmarshal(rawResp, &out); err != nil {
		return APIResponse{}, fmt.Errorf("decode upload response status=%d body=%s: %w", resp.StatusCode, strings.TrimSpace(string(rawResp)), err)
	}
	return out, nil
}

// isAuthFailure 判断响应是否为认证/授权失败。
func isAuthFailure(resp APIResponse) bool {
	if resp.Code == 401 {
		return true
	}
	msg := strings.ToLower(resp.Msg)
	authKeywords := []string{"token", "认证", "登录", "unauthorized", "未授权", "权限", "expire", "过期", "invalid"}
	for _, kw := range authKeywords {
		if strings.Contains(msg, kw) {
			return true
		}
	}
	return false
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

// 脱敏敏感 header
var sensitiveHeaders = []string{"token", "authorization", "cookie", "set-cookie", "x-api-key", "api-key", "password"}

func isSensitiveHeader(key string) bool {
	lower := strings.ToLower(key)
	for _, h := range sensitiveHeaders {
		if lower == h || strings.Contains(lower, h) {
			return true
		}
	}
	return false
}

func redactHeaders(headers http.Header) http.Header {
	redacted := headers.Clone()
	for key := range redacted {
		if isSensitiveHeader(key) {
			redacted.Set(key, "<REDACTED>")
		}
	}
	return redacted
}

func logDebugRequest(req *http.Request, body []byte) {
	log.Println("[debug] ---------- HTTP Request ----------")
	log.Printf("[debug] %s %s", req.Method, req.URL.String())
	for key, values := range redactHeaders(req.Header) {
		for _, v := range values {
			log.Printf("[debug] %s: %s", key, v)
		}
	}
	if len(body) > 0 {
		log.Printf("[debug] Body: %s", string(body))
	}
	log.Println("[debug] ----------------------------------")
}

func logDebugResponse(resp *http.Response, body []byte) {
	log.Println("[debug] ---------- HTTP Response ---------")
	log.Printf("[debug] Status: %d %s", resp.StatusCode, resp.Status)
	for key, values := range resp.Header {
		for _, v := range values {
			log.Printf("[debug] %s: %s", key, v)
		}
	}
	if len(body) > 0 {
		const maxBodyLen = 4096
		if len(body) > maxBodyLen {
			log.Printf("[debug] Body: %s... (%d bytes truncated)", string(body[:maxBodyLen]), len(body))
		} else {
			log.Printf("[debug] Body: %s", string(body))
		}
	}
	log.Println("[debug] ----------------------------------")
}
