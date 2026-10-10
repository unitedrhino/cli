// 文件说明：通用 API 的 SSE 请求入口，复用普通 API 的鉴权和请求上下文。
package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// OpenAPIStream 启动一个不设总超时的 SSE 请求，调用方负责关闭响应体。
func OpenAPIStream(ctx context.Context, req APIRequest) (*http.Response, error) {
	httpReq, rawBody, err := buildAPIRequest(ctx, req, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "text/event-stream")
	if req.Debug {
		logDebugRequest(httpReq, rawBody)
	}
	resp, err := (&http.Client{}).Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("open SSE request: %w", err)
	}
	if resp.StatusCode != http.StatusOK || !strings.Contains(resp.Header.Get("Content-Type"), "text/event-stream") {
		defer resp.Body.Close()
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 8192))
		return nil, fmt.Errorf("预期 SSE 响应，实际 HTTP %d Content-Type=%s: %s", resp.StatusCode, resp.Header.Get("Content-Type"), strings.TrimSpace(string(body)))
	}
	return resp, nil
}
