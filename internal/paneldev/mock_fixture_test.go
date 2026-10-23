// 文件说明：通过真实本地HTTP服务验证面板库夹具读取、缺省及格式错误，不访问平台。
package paneldev

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestOfflineMockFixture 验证.data/mock.json初值按实际工程返回，编辑错误可见，不静默替代为默认值。
func TestOfflineMockFixture(t *testing.T) {
	project := filepath.Join(t.TempDir(), "panel")
	if err := Init(project); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reader, writer := io.Pipe()
	defer reader.Close()
	done := make(chan error, 1)
	go func() { done <- Serve(ctx, project, 0, writer); writer.Close() }()
	line, err := bufio.NewReader(reader).ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	url := strings.TrimSpace(strings.TrimPrefix(line, "本地模拟面板：")) + "/__ur_panel/mock"
	client := http.Client{Timeout: time.Second}
	get := func(status int) map[string]any {
		t.Helper()
		r, e := client.Get(url)
		if e != nil {
			t.Fatal(e)
		}
		defer r.Body.Close()
		if r.StatusCode != status {
			t.Fatalf("状态=%d，期望%d", r.StatusCode, status)
		}
		if status != 200 {
			return nil
		}
		var data map[string]any
		if e = json.NewDecoder(r.Body).Decode(&data); e != nil {
			t.Fatal(e)
		}
		return data
	}
	if data := get(200); len(data) != 0 {
		t.Fatal(data)
	}
	fixture := filepath.Join(project, ".data", "mock.json")
	if err = os.WriteFile(fixture, []byte(`{"power":true,"level":395,"unknown":null}`), 0644); err != nil {
		t.Fatal(err)
	}
	data := get(200)
	if data["power"] != true || data["level"] != float64(395) {
		t.Fatal(data)
	}
	if _, exists := data["unknown"]; !exists {
		t.Fatal("未知初值丢失")
	}
	for _, invalid := range []string{`null`, `[]`, `{"power":`} {
		if err = os.WriteFile(fixture, []byte(invalid), 0644); err != nil {
			t.Fatal(err)
		}
		get(422)
	}
	cancel()
	reader.Close()
	select {
	case err = <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("服务未退出")
	}
}
