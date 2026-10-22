// 文件说明：验证不可信面板包的真实归一化、ZIP 导出回导、路径与资源边界。
package panelpkg

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"os"
	"strings"
	"testing"
)

// fixtureHTML 提供没有第二份清单的正常运行入口。
func fixtureHTML(t *testing.T) []byte {
	t.Helper()
	return []byte(`<!doctype html><html><head><style>body{color:red}</style></head><body>环境面板</body></html>`)
}

// fixtureMetadata 提供独立根元数据，字段与真实包合同一致。
func fixtureMetadata(t *testing.T) []byte {
	t.Helper()
	m := Manifest{FormatVersion: 1, Name: "环境面板", Entry: "index.html", SDKVersion: SDKVersion, Platforms: []string{"web", "h5"}, Requirements: []Requirement{{Kind: "property", Identifier: "temperature", Required: true, Define: map[string]any{"type": "float", "unit": "℃"}, Access: "read"}}}
	b, e := json.Marshal(m)
	if e != nil {
		t.Fatal(e)
	}
	return b
}

// fixturePackage 按唯一正式导入路径校验真实 ZIP，禁止测试绕过必需清单。
func fixturePackage(t *testing.T) (*Package, error) {
	t.Helper()
	return Decode("panel.zip", fixtureZIP(t, map[string][]byte{"meta.json": fixtureMetadata(t), "index.html": fixtureHTML(t)}, zip.Deflate, false))
}

// fixtureZIP 创建真实不同压缩方式和条目属性的压缩包。
func fixtureZIP(t *testing.T, files map[string][]byte, method uint16, symlink bool) []byte {
	t.Helper()
	var out bytes.Buffer
	w := zip.NewWriter(&out)
	for name, b := range files {
		h := &zip.FileHeader{Name: name, Method: method}
		if symlink {
			h.SetMode(os.ModeSymlink | 0777)
		}
		f, e := w.CreateHeader(h)
		if e != nil {
			t.Fatal(e)
		}
		if _, e = f.Write(b); e != nil {
			t.Fatal(e)
		}
	}
	if e := w.Close(); e != nil {
		t.Fatal(e)
	}
	return out.Bytes()
}

// TestPackageRoundTrip 证明根 meta.json 的 ZIP 导出和回导摘要一致。
func TestPackageRoundTrip(t *testing.T) {
	p, e := fixturePackage(t)
	if e != nil {
		t.Fatal(e)
	}
	if strings.Contains(string(p.Files["index.html"]), "ur-panel-manifest") {
		t.Fatal("运行入口仍含重复清单")
	}
	z, e := p.ZIP()
	if e != nil {
		t.Fatal(e)
	}
	for _, data := range [][]byte{z, fixtureZIP(t, p.Files, zip.Deflate, false)} {
		q, e := Decode("round.zip", data)
		if e != nil {
			t.Fatal(e)
		}
		if p.Hash != q.Hash {
			t.Fatalf("回导内容不等价 %s != %s", p.Hash, q.Hash)
		}
	}
}

// TestPackageRejectsUnsafeInputs 验证越界、缺失资源、非法清单与外部依赖均无法形成资产。
func TestPackageRejectsUnsafeInputs(t *testing.T) {
	base, e := fixturePackage(t)
	if e != nil {
		t.Fatal(e)
	}
	cases := map[string][]byte{"traversal": fixtureZIP(t, map[string][]byte{"../index.html": fixtureHTML(t)}, zip.Store, false), "symlink": fixtureZIP(t, base.Files, zip.Store, true), "missing-entry": fixtureZIP(t, map[string][]byte{"meta.json": base.Files["meta.json"]}, zip.Store, false), "broken": []byte("not zip")}
	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			if _, e := Decode("bad.zip", data); e == nil {
				t.Fatal("非法包被接受")
			}
		})
	}
	for _, html := range []string{strings.Replace(string(fixtureHTML(t)), "</body>", `<script src="https://cdn.example/a.js"></script></body>`, 1), strings.Replace(string(fixtureHTML(t)), "</body>", `<img src="assets/missing.png"></body>`, 1), strings.Replace(string(fixtureHTML(t)), "</body>", `<iframe src="anything"></iframe></body>`, 1)} {
		if _, e := Decode("bad.zip", fixtureZIP(t, map[string][]byte{"meta.json": fixtureMetadata(t), "index.html": []byte(html)}, zip.Deflate, false)); e == nil {
			t.Fatal("非法 HTML 被接受")
		}
	}
}

// TestPackageLocalResources 验证多资源包必须包含被引用的脚本/样式，资源内容参与摘要。
func TestPackageLocalResources(t *testing.T) {
	p, e := fixturePackage(t)
	if e != nil {
		t.Fatal(e)
	}
	files := p.Files
	files["index.html"] = bytes.Replace(files["index.html"], []byte("</head>"), []byte(`<link rel="stylesheet" href="assets/panel.css"><script src="assets/panel.js"></script></head>`), 1)
	files["assets/panel.css"] = []byte(`body{background:url(icon.svg)}`)
	files["assets/icon.svg"] = []byte(`<svg xmlns="http://www.w3.org/2000/svg"></svg>`)
	files["assets/panel.js"] = []byte(`window.panelLoaded=true`)
	q, e := Decode("resources.zip", fixtureZIP(t, files, zip.Deflate, false))
	if e != nil {
		t.Fatal(e)
	}
	if len(q.Inventory) != 5 || q.Hash == p.Hash {
		t.Fatal("资源未进入清单或摘要")
	}
	delete(files, "assets/icon.svg")
	if _, e = Decode("missing.zip", fixtureZIP(t, files, zip.Store, false)); e == nil {
		t.Fatal("缺失 CSS 依赖仍被接受")
	}
}

// TestCompressedResourceExport 验证大但可压缩的资源不会在导出时膨胀到无法回导。
func TestCompressedResourceExport(t *testing.T) {
	p, e := fixturePackage(t)
	if e != nil {
		t.Fatal(e)
	}
	p.Files["large.js"] = bytes.Repeat([]byte("/* local resource */\n"), 600000)
	large, e := Decode("input.zip", fixtureZIP(t, p.Files, zip.Deflate, false))
	if e != nil {
		t.Fatal(e)
	}
	exported, e := large.ZIP()
	if e != nil {
		t.Fatal(e)
	}
	imported, e := Decode("roundtrip.zip", exported)
	if e != nil || imported.Hash != large.Hash {
		t.Fatalf("导出不能回导：%v", e)
	}
}

// TestPackageCover 验证截图完整性、导出回导及内容摘要，拒绝伪装图片与超大尺寸。
func TestPackageCover(t *testing.T) {
	var out bytes.Buffer
	if e := png.Encode(&out, image.NewRGBA(image.Rect(0, 0, 64, 40))); e != nil {
		t.Fatal(e)
	}
	p, e := Decode("cover.zip", fixtureZIP(t, map[string][]byte{"meta.json": fixtureMetadata(t), "index.html": fixtureHTML(t), CoverFile: out.Bytes()}, zip.Deflate, false))
	if e != nil {
		t.Fatal(e)
	}
	exported, e := p.ZIP()
	if e != nil {
		t.Fatal(e)
	}
	again, e := Decode("again.zip", exported)
	if e != nil || !bytes.Equal(again.Files[CoverFile], out.Bytes()) {
		t.Fatalf("封面未完整保留: %v", e)
	}
	without, e := fixturePackage(t)
	if e != nil {
		t.Fatal(e)
	}
	if without.Hash == p.Hash {
		t.Fatal("封面未参与摘要")
	}
	for name, data := range map[string][]byte{"SVG伪装": []byte(`<svg onload="alert(1)"></svg>`), "PNG截断": out.Bytes()[:30]} {
		t.Run(name, func(t *testing.T) {
			if _, e := Decode("bad.zip", fixtureZIP(t, map[string][]byte{"meta.json": fixtureMetadata(t), "index.html": fixtureHTML(t), CoverFile: data}, zip.Store, false)); e == nil {
				t.Fatal("损坏封面未拒绝")
			}
		})
	}
	out.Reset()
	if e := png.Encode(&out, image.NewRGBA(image.Rect(0, 0, 1921, 1))); e != nil {
		t.Fatal(e)
	}
	if _, e := Decode("large.zip", fixtureZIP(t, map[string][]byte{"meta.json": fixtureMetadata(t), "index.html": fixtureHTML(t), CoverFile: out.Bytes()}, zip.Store, false)); e == nil {
		t.Fatal("超大封面未拒绝")
	}
}

// TestPackageMetadataBundle 验证新清单、资料目录、终端和完整导出不会丢失；拒绝两份清单及路径逃逸。
func TestPackageMetadataBundle(t *testing.T) {
	base, err := fixturePackage(t)
	if err != nil {
		t.Fatal(err)
	}
	m := base.Manifest
	m.Readme = "README.md"
	m.Skills = []string{"skills/panel-development.md"}
	m.ThingModel = ".data/thing-model.json"
	m.Cover = CoverFile
	metadata, _ := json.Marshal(m)
	files := map[string][]byte{"meta.json": metadata, "index.html": base.Files["index.html"], "README.md": []byte("# 开发说明"), "skills/panel-development.md": []byte("# 面板约束"), ".data/thing-model.json": []byte(`{"version":"1.0","properties":[]}`)}
	p, err := Decode("new.zip", fixtureZIP(t, files, zip.Deflate, false))
	if err != nil {
		t.Fatal(err)
	}
	if p.Manifest.Platforms[0] != "web" || p.Manifest.Readme != "README.md" || p.Files["manifest.json"] != nil {
		t.Fatal("新元数据未生效")
	}
	z, err := p.ZIP()
	if err != nil {
		t.Fatal(err)
	}
	q, err := Decode("round.zip", z)
	if err != nil || q.Hash != p.Hash {
		t.Fatalf("元数据包回导不稳定：%v", err)
	}
	files["manifest.json"] = metadata
	if _, err = Decode("duplicate.zip", fixtureZIP(t, files, zip.Deflate, false)); err == nil {
		t.Fatal("接受双份清单")
	}
	delete(files, "manifest.json")
	m.Readme = "../README.md"
	files["meta.json"], _ = json.Marshal(m)
	if _, err = Decode("escape.zip", fixtureZIP(t, files, zip.Deflate, false)); err == nil {
		t.Fatal("接受资料路径逃逸")
	}
}

// TestPanelZIPMetaRequired 只接受根目录有meta.json的ZIP，不能用旧清单或内嵌清单代替。
func TestPanelZIPMetaRequired(t *testing.T) {
	metadata, _ := json.Marshal(Manifest{FormatVersion: 1, Name: "严格ZIP", Entry: "index.html", SDKVersion: SDKVersion, Platforms: []string{"web"}, Requirements: []Requirement{}})
	for name, data := range map[string][]byte{
		"single.html":  fixtureHTML(t),
		"embedded.zip": fixtureZIP(t, map[string][]byte{"meta.json": metadata, "index.html": []byte(`<html><head><script type="application/json" id="ur-panel-manifest">` + string(metadata) + `</script></head></html>`)}, zip.Deflate, false),
		"legacy.zip":   fixtureZIP(t, map[string][]byte{"index.html": []byte("<html><body>旧包</body></html>"), "manifest.json": metadata}, zip.Deflate, false),
		"nested.zip":   fixtureZIP(t, map[string][]byte{"index.html": []byte("<html></html>"), "panel/meta.json": metadata}, zip.Deflate, false),
	} {
		if _, e := Decode(name, data); e == nil {
			t.Fatalf("缺少根meta.json的输入被接受：%s", name)
		}
	}
	p, e := Decode("strict.zip", fixtureZIP(t, map[string][]byte{"index.html": []byte("<html><body>面板</body></html>"), "meta.json": metadata}, zip.Deflate, false))
	if e != nil {
		t.Fatal(e)
	}
	exported, e := p.ZIP()
	if e != nil {
		t.Fatal(e)
	}
	round, e := Decode("export.zip", exported)
	if e != nil || round.Files["meta.json"] == nil || round.Hash != p.Hash {
		t.Fatal("导出没有保留meta与摘要", e)
	}
}
