// 文件说明：离线面板工程初始化、严格包校验、可重复打包与本地模拟服务；不加载平台认证。
package paneldev

import (
	"archive/zip"
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/panelpkg"
)

// assets 提供随 CLI 分发的固定 SDK、模拟宿主与示例工程，不依赖 CDN。
//
//go:embed assets/*
var assets embed.FS

// Init 在不存在的目录写出完整可运行示例，绝不覆盖已有开发文件。
func Init(directory string) error { return InitTemplate(directory, "switch") }

// InitTemplate 按固定模板名创建完整工程；未知模板和已有目录均不写入文件。
func InitTemplate(directory, template string) error {
	prefix := "assets/"
	switch template {
	case "switch":
	case "air-conditioner":
		prefix += "air-conditioner/"
	default:
		return fmt.Errorf("未知面板模板：%s（可选 switch、air-conditioner）", template)
	}
	if _, err := os.Lstat(directory); !os.IsNotExist(err) {
		return fmt.Errorf("目标目录已存在或不可访问：%s", directory)
	}
	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}
	for _, name := range []string{"manifest.json", "index.html", "panel-sdk.js", "panel.css", "panel.js", "mock.json", "README.md"} {
		data, err := assets.ReadFile(prefix + name)
		if err != nil {
			return err
		}
		if err = os.WriteFile(filepath.Join(directory, name), data, 0644); err != nil {
			return err
		}
	}
	return nil
}

// Load 读取 HTML、ZIP 或工程目录；目录校验拒绝符号链接并限制展开体积与文件数。
func Load(input string) (*panelpkg.Package, error) {
	info, err := os.Lstat(input)
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return nil, fmt.Errorf("不允许符号链接")
	}
	if !info.IsDir() {
		if info.Size() > panelpkg.MaxCompressed {
			return nil, fmt.Errorf("文件超过 10 MiB")
		}
		data, e := os.ReadFile(input)
		if e != nil {
			return nil, e
		}
		return panelpkg.Decode(filepath.Base(input), data)
	}
	var buffer bytes.Buffer
	writer := zip.NewWriter(&buffer)
	count, total := 0, int64(0)
	err = filepath.WalkDir(input, func(name string, entry os.DirEntry, e error) error {
		if e != nil {
			return e
		}
		relative, e := filepath.Rel(input, name)
		if e != nil {
			return e
		}
		if relative == "." {
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("不允许符号链接：%s", relative)
		}
		if entry.IsDir() {
			if strings.HasPrefix(entry.Name(), ".") || entry.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if relative == "README.md" || relative == "mock.json" || strings.HasPrefix(entry.Name(), ".") {
			return nil
		}
		stat, e := entry.Info()
		if e != nil {
			return e
		}
		count++
		total += stat.Size()
		if count > panelpkg.MaxFiles || total > panelpkg.MaxExpanded {
			return fmt.Errorf("工程文件数或大小超限")
		}
		file, e := writer.Create(filepath.ToSlash(relative))
		if e != nil {
			return e
		}
		source, e := os.Open(name)
		if e != nil {
			return e
		}
		_, e = io.Copy(file, io.LimitReader(source, panelpkg.MaxExpanded+1))
		source.Close()
		return e
	})
	if err != nil {
		return nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}
	return panelpkg.Decode("panel.zip", buffer.Bytes())
}

// Pack 校验后以标准 ZIP 写入新文件，拒绝覆盖；摘要以服务端同源规则计算。
func Pack(input, output string) (*panelpkg.Package, error) {
	p, err := Load(input)
	if err != nil {
		return nil, err
	}
	data, err := p.ZIP()
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, err = file.Write(data); err != nil {
		return nil, err
	}
	return p, nil
}

// Serve 仅监听本机随机端口，重新读取文件支持外部编辑器开发，不代理平台或真实设备 API。
func Serve(ctx context.Context, input string, port int, out io.Writer) error {
	if _, err := Load(input); err != nil {
		return err
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return err
	}
	defer listener.Close()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		switch r.URL.Path {
		case "/":
			data, _ := assets.ReadFile("assets/host.html")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(data)
			return
		case "/__ur_panel/host.js":
			data, _ := assets.ReadFile("assets/host.js")
			w.Header().Set("Content-Type", "application/javascript")
			w.Write(data)
			return
		case "/__ur_panel/manifest":
			p, e := Load(input)
			if e != nil {
				http.Error(w, "工程校验失败，请在终端运行 validate", 422)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p.Manifest)
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/panel/") {
			http.NotFound(w, r)
			return
		}
		p, e := Load(input)
		if e != nil {
			http.Error(w, "工程校验失败", 422)
			return
		}
		name := strings.TrimPrefix(r.URL.Path, "/panel/")
		if name == "" {
			name = "index.html"
		}
		data, ok := p.Files[name]
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Security-Policy", "default-src 'none'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: blob:; font-src 'self' data:; connect-src 'none'; base-uri 'none'; form-action 'none'; frame-src 'none'; sandbox allow-scripts")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cross-Origin-Resource-Policy", "cross-origin")
		w.Header().Set("Referrer-Policy", "no-referrer")
		if strings.HasSuffix(name, ".html") {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		}
		http.ServeContent(w, r, name, time.Time{}, bytes.NewReader(data))
	})
	server := &http.Server{Handler: handler, ReadHeaderTimeout: 5 * time.Second}
	done := make(chan struct{})
	defer close(done)
	go func() {
		select {
		case <-ctx.Done():
			server.Close()
		case <-done:
		}
	}()
	fmt.Fprintf(out, "本地模拟面板：http://%s\n仅模拟设备；不连接平台或蓝牙。使用 Ctrl+C 停止。\n", listener.Addr())
	err = server.Serve(listener)
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}
