// 文件说明：在无认证空目录验证模板、离线打包、回导和路径拒绝。
package paneldev

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestLocalDevelopment 验证离线开发完整链路不依赖账号或平台连接。
func TestLocalDevelopment(t *testing.T) {
	root := t.TempDir()
	project := filepath.Join(root, "panel")
	if e := Init(project); e != nil {
		t.Fatal(e)
	}
	if e := Init(project); e == nil {
		t.Fatal("初始化覆盖已有目录")
	}
	p, e := Load(project)
	if e != nil {
		t.Fatal(e)
	}
	if len(p.Files) != 3 {
		t.Fatal(p.Inventory)
	}
	out := filepath.Join(root, "panel.zip")
	q, e := Pack(project, out)
	if e != nil || q.Hash != p.Hash {
		t.Fatalf("打包失败：%v", e)
	}
	r, e := Load(out)
	if e != nil || r.Hash != p.Hash {
		t.Fatalf("回导摘要不一致：%v", e)
	}
	if _, e = Pack(project, out); e == nil {
		t.Fatal("覆盖导出文件")
	}
	if e = os.Symlink(filepath.Join(project, "index.html"), filepath.Join(project, "hidden.js")); e != nil {
		t.Fatal(e)
	}
	if _, e = Load(project); e == nil {
		t.Fatal("接受符号链接")
	}
}

// TestOfflineServer 验证仅本机绑定、资产 CSP 及无通用平台代理。
func TestOfflineServer(t *testing.T) {
	project := filepath.Join(t.TempDir(), "panel")
	if e := Init(project); e != nil {
		t.Fatal(e)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reader, writer := io.Pipe()
	defer reader.Close()
	done := make(chan error, 1)
	go func() { done <- Serve(ctx, project, 0, writer); writer.Close() }()
	line, e := bufio.NewReader(reader).ReadString('\n')
	if e != nil {
		t.Fatal(e)
	}
	url := strings.TrimSpace(strings.TrimPrefix(line, "本地模拟面板："))
	if !strings.HasPrefix(url, "http://127.0.0.1:") {
		t.Fatal(url)
	}
	client := http.Client{Timeout: time.Second}
	for _, path := range []string{"/", "/panel/index.html", "/panel/panel-sdk.js", "/__ur_panel/manifest"} {
		resp, e := client.Get(url + path)
		if e != nil {
			t.Fatal(e)
		}
		if resp.StatusCode != 200 {
			t.Fatal(path, resp.StatusCode)
		}
		if strings.HasPrefix(path, "/panel/") && !strings.Contains(resp.Header.Get("Content-Security-Policy"), "connect-src 'none'") {
			t.Fatal("缺少 CSP")
		}
		resp.Body.Close()
	}
	resp, e := client.Get(url + "/api/v1/things/device/info/get-one")
	if e != nil {
		t.Fatal(e)
	}
	resp.Body.Close()
	if resp.StatusCode != 404 {
		t.Fatal("模拟服务开放平台接口")
	}
	cancel()
	reader.Close()
	select {
	case e = <-done:
		if e != nil {
			t.Fatal(e)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("服务未退出")
	}
}
