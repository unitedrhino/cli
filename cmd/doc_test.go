// doc_test.go 覆盖 ur doc 命令：parse 的 json/outline 输出、excel 公式场景
// 与平台 OCR 客户端（env-only 认证 + 平台自有响应信封解析）。
package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/unitedrhino/docling/llmocr"
	"github.com/xuri/excelize/v2"
)

// resetDocParseOpts 保存/恢复包级参数实例，避免用例间串扰。
func resetDocParseOpts(t *testing.T) {
	saved := *docParseOpts
	t.Cleanup(func() { *docParseOpts = saved })
	*docParseOpts = docParseFlags{
		format: "outline", layers: "body", ocrProvider: "platform",
		pdfMaxFileSizeMB: 50, pdfMaxPages: 2000,
	}
}

// TestDocParseRejectsInvalidPDFLimits 验证安全限制不能通过非正参数关闭。
func TestDocParseRejectsInvalidPDFLimits(t *testing.T) {
	resetDocParseOpts(t)
	docParseOpts.pdfMaxPages = 0
	err := runDocParse(&cobra.Command{}, filepath.Join(t.TempDir(), "missing.pdf"))
	if err == nil || !strings.Contains(err.Error(), "必须大于 0") {
		t.Fatalf("invalid PDF limit err=%v", err)
	}
}

// TestReadDocSourceRejectsOversizedLocalPDF 验证本地 PDF 在 os.ReadFile 前
// 通过文件元数据拒绝超限输入。
func TestReadDocSourceRejectsOversizedLocalPDF(t *testing.T) {
	file := filepath.Join(t.TempDir(), "oversized.pdf")
	if err := os.WriteFile(file, []byte("12345"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, err := readDocSource(file, "", 4); err == nil || !strings.Contains(err.Error(), "超过") {
		t.Fatalf("oversized local PDF err=%v", err)
	}
}

// TestReadDocInputLimited 验证 URL 与 stdin 共用的 limit+1 读取器不会返回
// 部分文档或在超限后继续增长缓冲。
func TestReadDocInputLimited(t *testing.T) {
	if _, err := readDocInputLimited(strings.NewReader("12345"), 4); err == nil || !strings.Contains(err.Error(), "超过") {
		t.Fatalf("limited reader err=%v", err)
	}
	data, err := readDocInputLimited(strings.NewReader("1234"), 4)
	if err != nil || string(data) != "1234" {
		t.Fatalf("boundary read data=%q err=%v", data, err)
	}
}

// TestDownloadDocFileRejectsOversizedStream 验证无 Content-Length 的分块响应
// 仍由 LimitReader 在内存缓冲超过边界前终止。
func TestDownloadDocFileRejectsOversizedStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.(http.Flusher).Flush()
		_, _ = io.WriteString(w, "12345")
	}))
	defer server.Close()
	if _, _, err := downloadDocFile(server.URL+"/report.pdf", 4); err == nil || !strings.Contains(err.Error(), "超过") {
		t.Fatalf("oversized URL PDF err=%v", err)
	}
}

// TestDocParseRejectsPDFPageLimit 验证页数参数会透传到 Go Docling，并以
// 可操作的 CLI 资源超限错误退出。
func TestDocParseRejectsPDFPageLimit(t *testing.T) {
	resetDocParseOpts(t)
	docParseOpts.pdfMaxPages = 1
	file := filepath.Join(t.TempDir(), "two-pages.pdf")
	if err := os.WriteFile(file, buildCLITestPDF(2), 0o644); err != nil {
		t.Fatal(err)
	}
	err := runDocParse(&cobra.Command{}, file)
	if err == nil || !strings.Contains(err.Error(), "PDF 超出安全限制") || !strings.Contains(err.Error(), "pages 2") {
		t.Fatalf("page limit err=%v", err)
	}
}

// buildCLITestPDF 构造指定页数的最小 PDF，供 CLI 页数限制测试使用。
func buildCLITestPDF(pageCount int) []byte {
	objects := []string{"<< /Type /Catalog /Pages 2 0 R >>"}
	kids := make([]string, 0, pageCount)
	for page := 0; page < pageCount; page++ {
		kids = append(kids, fmt.Sprintf("%d 0 R", page+3))
	}
	objects = append(objects, fmt.Sprintf("<< /Type /Pages /Kids [%s] /Count %d >>", strings.Join(kids, " "), pageCount))
	for range pageCount {
		objects = append(objects, "<< /Type /Page /Parent 2 0 R /MediaBox [0 0 200 200] >>")
	}
	var output bytes.Buffer
	output.WriteString("%PDF-1.4\n")
	offsets := make([]int, len(objects))
	for index, object := range objects {
		offsets[index] = output.Len()
		fmt.Fprintf(&output, "%d 0 obj\n%s\nendobj\n", index+1, object)
	}
	xref := output.Len()
	fmt.Fprintf(&output, "xref\n0 %d\n0000000000 65535 f \n", len(objects)+1)
	for _, offset := range offsets {
		fmt.Fprintf(&output, "%010d 00000 n \n", offset)
	}
	fmt.Fprintf(&output, "trailer\n<< /Size %d /Root 1 0 R >>\nstartxref\n%d\n%%%%EOF\n", len(objects)+1, xref)
	return output.Bytes()
}

// mustTestFormulaXLSX 构造含公式的 xlsx：A1=合计 B1=SUM(A2:A3) A2=1 A3=2。
func mustTestFormulaXLSX(t *testing.T) string {
	t.Helper()
	file := excelize.NewFile()
	t.Cleanup(func() { _ = file.Close() })
	if err := file.SetCellValue("Sheet1", "A1", "合计"); err != nil {
		t.Fatal(err)
	}
	if err := file.SetCellFormula("Sheet1", "B1", "SUM(A2:A3)"); err != nil {
		t.Fatal(err)
	}
	for cell, value := range map[string]int{"A2": 1, "A3": 2} {
		if err := file.SetCellValue("Sheet1", cell, value); err != nil {
			t.Fatal(err)
		}
	}
	path := filepath.Join(t.TempDir(), "report.xlsx")
	var buf bytes.Buffer
	if err := file.Write(&buf); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

// TestDocFormatsCommand 验证 formats 子命令。
func TestDocFormatsCommand(t *testing.T) {
	var out bytes.Buffer
	cmd := &cobra.Command{}
	cmd.SetOut(&out)
	if err := docFormatsCmd.RunE(cmd, nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "pdf") || !strings.Contains(out.String(), "xlsx") {
		t.Fatalf("formats=%s", out.String())
	}
}

// TestDocParseOutline 验证 outline 含 sheet 名与表格信息。
func TestDocParseOutline(t *testing.T) {
	resetDocParseOpts(t)
	var out bytes.Buffer
	cmd := &cobra.Command{}
	cmd.SetOut(&out)
	if err := runDocParse(cmd, mustTestFormulaXLSX(t)); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Sheet1", "表格:"} {
		if !strings.Contains(out.String(), want) {
			t.Fatalf("outline missing %q:\n%s", want, out.String())
		}
	}
}

// TestDocParseJSONFormulaScenario 校准 excel 公式 jq 查询路径：公式在
// texts[].meta.docling__xlsx_formula，表格单元格 ref.$ref 反查。
func TestDocParseJSONFormulaScenario(t *testing.T) {
	resetDocParseOpts(t)
	docParseOpts.format = "json"
	docParseOpts.out = filepath.Join(t.TempDir(), "doc.json")
	cmd := &cobra.Command{}
	cmd.SetOut(&bytes.Buffer{})
	cmd.SetErr(&bytes.Buffer{})
	if err := runDocParse(cmd, mustTestFormulaXLSX(t)); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(docParseOpts.out)
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Texts []struct {
			Meta map[string]any `json:"meta"`
		} `json:"texts"`
		Tables []struct {
			Data *struct {
				TableCells []struct {
					Ref *struct {
						Ref string `json:"$ref"`
					} `json:"ref"`
				} `json:"table_cells"`
			} `json:"data"`
		} `json:"tables"`
	}
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	formulaIdx := -1
	for i, item := range doc.Texts {
		if f, ok := item.Meta["docling__xlsx_formula"].(string); ok && f == "SUM(A2:A3)" {
			formulaIdx = i
		}
	}
	if formulaIdx < 0 {
		t.Fatal("formula meta not found")
	}
	linked := false
	for _, table := range doc.Tables {
		if table.Data == nil {
			continue
		}
		for _, cell := range table.Data.TableCells {
			if cell.Ref != nil && cell.Ref.Ref == "#/texts/"+itoa(formulaIdx) {
				linked = true
			}
		}
	}
	if !linked {
		t.Fatalf("no cell ref points to #/texts/%d", formulaIdx)
	}
}

// TestPlatformLLMClient 验证平台 OCR 客户端：env-only 认证、请求组包、
// 平台响应信封解析与业务码处理。
func TestPlatformLLMClient(t *testing.T) {
	var gotPath, gotToken, gotBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotToken = r.Header.Get("token")
		body, _ := io.ReadAll(r.Body)
		gotBody = string(body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"msg":"ok","data":{"content":"页面识别文本","role":"assistant"}}`))
	}))
	defer server.Close()

	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "200")
	t.Setenv("UR_TENANT_CODE", "test-tenant")
	t.Setenv("UR_TOKEN", "test-token-123")

	got, err := (platformLLMClient{modelType: "large"}).Complete(context.Background(), llmocr.VisionRequest{
		Prompt: "识别这一页",
		Images: []llmocr.ImageInput{{MIMEType: "application/pdf", DataURI: "data:application/pdf;base64,QUJD"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != "页面识别文本" {
		t.Fatalf("got=%q", got)
	}
	if gotPath != "/api/v1/ai/chat/completions" {
		t.Fatalf("path=%s", gotPath)
	}
	if gotToken != "test-token-123" {
		t.Fatalf("token header=%q", gotToken)
	}
	for _, want := range []string{`"agentID":0`, `"modelType":"large"`, `"imageUrl":"data:application/pdf;base64,QUJD"`} {
		if !strings.Contains(gotBody, want) {
			t.Fatalf("body missing %s:\n%s", want, gotBody)
		}
	}

	// 业务码非 200 → 报错
	bad := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":500,"msg":"模型未配置","data":null}`))
	}))
	defer bad.Close()
	t.Setenv("UR_BASE_URL", bad.URL)
	t.Setenv("UR_APP_ID", "200")
	t.Setenv("UR_TENANT_CODE", "test-tenant")
	if _, err := (platformLLMClient{modelType: "large"}).Complete(context.Background(), llmocr.VisionRequest{Prompt: "p"}); err == nil || !strings.Contains(err.Error(), "模型未配置") {
		t.Fatalf("err=%v", err)
	}
}

// itoa 极简 int 转字符串(测试用)。
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	digits := ""
	for n > 0 {
		digits = string(rune('0'+n%10)) + digits
		n /= 10
	}
	return digits
}
