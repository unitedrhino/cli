// doc_glb_test.go 验证 GLB 经 CLI 输入、格式化、章节筛选和错误退出的完整路径。
package cmd

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/unitedrhino/docling"
)

// buildDocTestGLB 将模型 JSON 封装为 GLB 2.0，用于命令路径验证而非几何兼容性验收。
func buildDocTestGLB(metadata string) []byte {
	chunk := []byte(metadata)
	for len(chunk)%4 != 0 {
		chunk = append(chunk, ' ')
	}
	data := make([]byte, 20+len(chunk))
	copy(data, "glTF")
	binary.LittleEndian.PutUint32(data[4:8], 2)
	binary.LittleEndian.PutUint32(data[8:12], uint32(len(data)))
	binary.LittleEndian.PutUint32(data[12:16], uint32(len(chunk)))
	binary.LittleEndian.PutUint32(data[16:20], 0x4E4F534A)
	copy(data[20:], chunk)
	return data
}

// docTestGLBModel 包含父子节点和超出浮点安全整数范围的设备编号，验证属性精度与来源。
const docTestGLBModel = `{"asset":{"version":"2.0"},"scene":0,"scenes":[{"name":"机房","nodes":[0]}],"nodes":[{"name":"设备组","children":[1]},{"name":"循环泵","extras":{"deviceNo":9007199254740993,"power":"5.5kW"}}]}`

// mustDocTestGLB 写入模型夹具，返回可用于 CLI 的本地路径。
func mustDocTestGLB(t *testing.T, metadata string) string {
	t.Helper()
	file := filepath.Join(t.TempDir(), "模型.GLB")
	if err := os.WriteFile(file, buildDocTestGLB(metadata), 0o644); err != nil {
		t.Fatal(err)
	}
	return file
}

// TestDocParseGLBOutputs 验证四种输出均可读取模型，结构化输出保留节点来源与精确属性。
func TestDocParseGLBOutputs(t *testing.T) {
	for _, format := range []string{"outline", "md", "content-list", "json"} {
		t.Run(format, func(t *testing.T) {
			resetDocParseOpts(t)
			docParseOpts.format = format
			var out bytes.Buffer
			cmd := &cobra.Command{}
			cmd.SetOut(&out)
			if err := docParseCmd.RunE(cmd, []string{mustDocTestGLB(t, docTestGLBModel)}); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(out.String(), "循环泵") {
				t.Fatalf("%s 没有模型节点: %s", format, out.String())
			}
			if format == "outline" {
				return
			}
			for _, want := range []string{"#/nodes/1", "9007199254740993", "5.5kW"} {
				if !strings.Contains(out.String(), want) {
					t.Fatalf("%s 丢失属性或来源 %q: %s", format, want, out.String())
				}
			}
			if format == "json" {
				var doc docling.DoclingDocument
				if err := json.Unmarshal(out.Bytes(), &doc); err != nil {
					t.Fatal(err)
				}
				found := false
				for _, item := range doc.Texts {
					var source struct {
						Pointer string `json:"pointer"` // Pointer 定位原 GLB JSON 对象。
					}
					if raw, ok := item.Meta["glb:source"]; ok {
						if err := json.Unmarshal(raw, &source); err != nil {
							t.Fatal(err)
						}
						found = found || source.Pointer == "#/nodes/1"
					}
				}
				if !found {
					t.Fatal("Docling JSON 没有节点来源元数据")
				}
			}
		})
	}
}

// TestDocParseGLBSection 验证筛选泵节点时属性仍属于原节点章节，不混入其他场景正文。
func TestDocParseGLBSection(t *testing.T) {
	for _, format := range []string{"md", "content-list"} {
		t.Run(format, func(t *testing.T) {
			resetDocParseOpts(t)
			docParseOpts.format = format
			docParseOpts.section = "循环泵"
			var out bytes.Buffer
			cmd := &cobra.Command{}
			cmd.SetOut(&out)
			if err := runDocParse(cmd, mustDocTestGLB(t, docTestGLBModel)); err != nil {
				t.Fatal(err)
			}
			for _, want := range []string{"循环泵", "9007199254740993", "#/nodes/1"} {
				if !strings.Contains(out.String(), want) {
					t.Fatalf("筛选结果缺少 %q: %s", want, out.String())
				}
			}
			if strings.Contains(out.String(), "机房") {
				t.Fatalf("筛选结果混入无关场景: %s", out.String())
			}
			if format == "content-list" {
				var items []docling.Item
				if err := json.Unmarshal(out.Bytes(), &items); err != nil {
					t.Fatal(err)
				}
				if len(items) == 0 {
					t.Fatal("筛选后列表为空")
				}
				for _, item := range items {
					if !strings.Contains(strings.Join(item.SectionPath, "/"), "循环泵") {
						t.Fatalf("来源章节丢失: %+v", item)
					}
				}
			}
		})
	}
}

// TestDocParseGLBURL 验证 URL 路径决定 GLB 解析器，输出文件仍保留模型属性。
func TestDocParseGLBURL(t *testing.T) {
	resetDocParseOpts(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(buildDocTestGLB(docTestGLBModel))
	}))
	defer server.Close()
	docParseOpts.format = "md"
	docParseOpts.out = filepath.Join(t.TempDir(), "model.md")
	if err := runDocParse(&cobra.Command{}, server.URL+"/model.glb?download=1"); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(docParseOpts.out)
	if err != nil || !bytes.Contains(raw, []byte("9007199254740993")) {
		t.Fatalf("URL 模型输出不正确: %s, %v", raw, err)
	}
}

// TestDocParseGLBRejectsInvalid 验证坏容器与循环节点不会生成成功正文或覆盖已有输出。
func TestDocParseGLBRejectsInvalid(t *testing.T) {
	for name, data := range map[string][]byte{
		"容器损坏": []byte("invalid glb"),
		"节点循环": buildDocTestGLB(`{"asset":{"version":"2.0"},"nodes":[{"children":[0]}]}`),
	} {
		t.Run(name, func(t *testing.T) {
			resetDocParseOpts(t)
			file := filepath.Join(t.TempDir(), "bad.glb")
			if err := os.WriteFile(file, data, 0o644); err != nil {
				t.Fatal(err)
			}
			docParseOpts.out = filepath.Join(t.TempDir(), "existing.md")
			if err := os.WriteFile(docParseOpts.out, []byte("已有内容"), 0o644); err != nil {
				t.Fatal(err)
			}
			err := runDocParse(&cobra.Command{}, file)
			cliErr, ok := err.(*CLIError)
			if !ok || cliErr.ExitCode != 1 {
				t.Fatalf("坏模型没有以错误退出: %v", err)
			}
			out, err := os.ReadFile(docParseOpts.out)
			if err != nil || string(out) != "已有内容" {
				t.Fatalf("失败解析覆盖已有输出: %s, %v", out, err)
			}
		})
	}
}
