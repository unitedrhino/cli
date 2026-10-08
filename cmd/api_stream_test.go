// 文件说明：验证通用 API 命令读取 SSE 时的鉴权、流式输出和错误处理。
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/pflag"
)

// TestAPIStream 验证 SSE 事件被逐行保留为机器可读 JSON。
func TestAPIStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/system/client-debug/stream" || r.Header.Get("token") != "test-token" || r.Header.Get("app-id") != "200" {
			t.Errorf("请求鉴权或路径错误: path=%s app-id=%s", r.URL.Path, r.Header.Get("app-id"))
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["userID"] != "42" {
			t.Errorf("请求体错误: %v %v", body, err)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = fmt.Fprint(w, "event: waiting\ndata: {\"sessionID\":\"abc\"}\n\nevent: log\ndata: {\"kind\":\"log\",\ndata: \"log\":{\"seq\":1}}\n\n")
	}))
	defer server.Close()
	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "200")
	t.Setenv("UR_TENANT_CODE", "test")
	t.Setenv("UR_TOKEN", "test-token")
	apiCmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if values, ok := flag.Value.(pflag.SliceValue); ok {
			_ = values.Replace(nil)
		} else {
			_ = flag.Value.Set(flag.DefValue)
		}
		flag.Changed = false
	})
	previous := RootCmd.OutOrStdout()
	defer RootCmd.SetOut(previous)
	var output bytes.Buffer
	RootCmd.SetOut(&output)
	RootCmd.SetArgs([]string{"api", "/api/v1/system/client-debug/stream", "--stream", "--body", `{"userID":"42"}`})
	if err := RootCmd.Execute(); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("应收到两个 SSE 事件，实际: %q", output.String())
	}
	for i, want := range []string{"waiting", "log"} {
		var event struct {
			Name string          `json:"event"`
			Data json.RawMessage `json:"data"`
		}
		if err := json.Unmarshal([]byte(lines[i]), &event); err != nil || event.Name != want || !json.Valid(event.Data) {
			t.Fatalf("事件 %d 无效: %s %v", i, lines[i], err)
		}
	}
}
