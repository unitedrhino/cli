// doc.go 实现 `ur doc` 文档解析命令族：基于 docling 库(github.com/unitedrhino/docling)
// 把文档转换为 outline / markdown / content-list / 无损 Docling JSON。命令只做
// 通用原语——转换与格式发现；单元格、公式、章节等精查交给 jq 等通用工具。
// OCR（--ocr）默认走平台 /api/v1/ai/chat/completions 裸模型代理（与知识库
// OCR 共用租户模型池计费），也可 --ocr-provider openai 直连 OpenAI 兼容网关。
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/client"
	"github.com/spf13/cobra"
	"github.com/unitedrhino/docling"
	"github.com/unitedrhino/docling/cli"
	"github.com/unitedrhino/docling/llmocr"
)

// docParseFlags 收集 parse 子命令参数。
type docParseFlags struct {
	format           string
	section          string
	sheet            string
	layers           string
	out              string
	name             string
	ocr              bool
	ocrProvider      string
	ocrModel         string
	maxPages         int
	pdfMaxFileSizeMB int
	pdfMaxPages      int
}

// docParseOpts 是 parse 参数实例(flag 绑定目标)。
var docParseOpts = &docParseFlags{}

// docParseCmd 是 `ur doc parse`。
var docParseCmd = &cobra.Command{
	Use:   "parse <file|URL|->",
	Short: "解析文档为结构地图/Markdown/内容列表/无损 JSON",
	Long: `解析文档(本地路径、http(s) URL 或 stdin)并输出:
  outline      结构地图:标题树/表格行列/图片/分组(渐进式披露首跳,默认)
  md           Markdown(可 --section 章节过滤)
  content-list 扁平内容列表 JSON(带 SectionPath/PageIdx/BBox,grep 友好)
  json         完整无损 Docling JSON(--out 落盘后用 jq 精查;公式在
               texts[].meta.docling__xlsx_formula,表格单元格经 ref 关联)

示例:
  ur doc parse 报告.pdf --format outline
  ur doc parse 报告.pdf --format md --section 第四章
  ur doc parse 报表.xlsx --format json --out d.json
  jq -r '.texts[]|select(.meta.docling__xlsx_formula)|[.text,.meta.docling__xlsx_formula]|@tsv' d.json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return &CLIError{Message: "需要且只需要一个输入:文件路径 / http(s) URL / -", ExitCode: 2}
		}
		return runDocParse(cmd, args[0])
	},
}

// docFormatsCmd 是 `ur doc formats`。
var docFormatsCmd = &cobra.Command{
	Use:   "formats",
	Short: "列出 doc 解析支持的文件格式",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, e := range []string{"pdf", "docx", "pptx", "xlsx", "csv", "html", "md", "adoc", "txt", "eml", "png", "jpg", "jpeg", "bmp", "webp"} {
			cmd.Println(e)
		}
		return nil
	},
}

// docCmd 是 `ur doc` 父命令。
var docCmd = &cobra.Command{Use: "doc", Short: "文档解析:把 PDF/Office/图片等转为 AI 友好的结构化输出"}

func init() {
	docParseCmd.Flags().StringVar(&docParseOpts.format, "format", "outline", "输出格式: outline|md|content-list|json")
	docParseCmd.Flags().StringVar(&docParseOpts.section, "section", "", "按章节路径前缀过滤(md/content-list)")
	docParseCmd.Flags().StringVar(&docParseOpts.sheet, "sheet", "", "按首级分组名(如工作表)过滤(md/content-list)")
	docParseCmd.Flags().StringVar(&docParseOpts.layers, "layers", "body", "markdown 内容层,逗号分隔: body,furniture,notes,invisible,background")
	docParseCmd.Flags().StringVar(&docParseOpts.out, "out", "", "输出到文件(默认 stdout;json 格式建议落盘后用 jq 查询)")
	docParseCmd.Flags().StringVar(&docParseOpts.name, "name", "", "stdin 模式的文件名(决定解析器)")
	docParseCmd.Flags().BoolVar(&docParseOpts.ocr, "ocr", false, "启用大模型 OCR/结构化视觉(扫描件/乱码页/图片表格自动送模型)")
	docParseCmd.Flags().StringVar(&docParseOpts.ocrProvider, "ocr-provider", "platform", "OCR 通路: platform(平台模型池)|openai(OPENAI_BASE_URL/OPENAI_API_KEY)")
	docParseCmd.Flags().StringVar(&docParseOpts.ocrModel, "ocr-model", "", "OCR 模型名(platform: modelType 默认 large;openai: 默认 gpt-4o)")
	docParseCmd.Flags().IntVar(&docParseOpts.maxPages, "ocr-max-pages", 0, "OCR/视觉最大页数预算(0=不限)")
	docParseCmd.Flags().IntVar(&docParseOpts.pdfMaxFileSizeMB, "pdf-max-file-size-mb", 50, "PDF 最大文件大小(MiB，必须大于 0)")
	docParseCmd.Flags().IntVar(&docParseOpts.pdfMaxPages, "pdf-max-pages", 2000, "PDF 最大页数(必须大于 0)")
	docCmd.AddCommand(docParseCmd)
	docCmd.AddCommand(docFormatsCmd)
	RootCmd.AddCommand(docCmd)
}

// runDocParse 执行读取→解析→格式化→输出。
func runDocParse(cmd *cobra.Command, source string) error {
	if docParseOpts.pdfMaxFileSizeMB <= 0 || docParseOpts.pdfMaxPages <= 0 {
		return &CLIError{Message: "--pdf-max-file-size-mb 和 --pdf-max-pages 必须大于 0", ExitCode: 2}
	}
	pdfMaxBytes := int64(docParseOpts.pdfMaxFileSizeMB) << 20
	maxReadBytes := int64(0)
	if strings.EqualFold(filepath.Ext(docInputName(source, docParseOpts.name)), ".pdf") {
		maxReadBytes = pdfMaxBytes
	}
	name, data, err := readDocSource(source, docParseOpts.name, maxReadBytes)
	if err != nil {
		return &CLIError{Message: err.Error(), ExitCode: 1}
	}
	parseOpts := docling.ParseOptions{PDFLimits: docling.PDFLimits{
		MaxFileBytes: pdfMaxBytes,
		MaxPages:     docParseOpts.pdfMaxPages,
	}}
	if docParseOpts.ocr {
		budget := llmocr.Options{MaxPages: docParseOpts.maxPages}
		switch strings.ToLower(docParseOpts.ocrProvider) {
		case "openai":
			cfg := llmocr.OpenAIConfig{Model: docParseOpts.ocrModel}
			if docParseOpts.ocrModel == "" {
				cfg.Model = os.Getenv("DOCLING_OCR_MODEL")
			}
			ocrClient, ocrErr := llmocr.NewOpenAIClient(cfg)
			if ocrErr != nil {
				return &CLIError{Message: ocrErr.Error(), ExitCode: 1}
			}
			parseOpts.OCRHook = llmocr.NewOCRHook(ocrClient, budget)
			parseOpts.PDFVisualHook = llmocr.NewPDFVisualHook(ocrClient, budget)
		default:
			ocrClient := platformLLMClient{modelType: docParseOpts.ocrModel}
			if ocrClient.modelType == "" {
				ocrClient.modelType = "large"
			}
			parseOpts.OCRHook = llmocr.NewOCRHook(ocrClient, budget)
			parseOpts.PDFVisualHook = llmocr.NewPDFVisualHook(ocrClient, budget)
		}
	}
	doc, err := docling.ParseByExtWithOptions(name, data, parseOpts)
	if err != nil {
		if errors.Is(err, docling.ErrPDFResourceLimit) {
			return &CLIError{Message: "PDF 超出安全限制: " + err.Error(), ExitCode: 1}
		}
		return &CLIError{Message: err.Error(), ExitCode: 1}
	}

	var output []byte
	switch docParseOpts.format {
	case "outline":
		output = []byte(cli.Outline(doc))
	case "md":
		if docParseOpts.section == "" && docParseOpts.sheet == "" {
			output = []byte(doc.ToMarkdownWithOptions(docling.ExportOptions{Layers: cli.ParseLayers(docParseOpts.layers)}))
		} else {
			output = []byte(cli.RenderItemsMarkdown(cli.FilterItems(docling.ToContentList(doc, docling.SourceGolight), docParseOpts.section, docParseOpts.sheet)))
		}
	case "content-list":
		items := docling.ToContentList(doc, docling.SourceGolight)
		if docParseOpts.section != "" || docParseOpts.sheet != "" {
			items = cli.FilterItems(items, docParseOpts.section, docParseOpts.sheet)
		}
		if output, err = docMarshalIndent(items); err != nil {
			return &CLIError{Message: err.Error(), ExitCode: 1}
		}
	case "json":
		if docParseOpts.section != "" || docParseOpts.sheet != "" {
			cmd.PrintErrln("提示: json 格式不受 --section/--sheet 影响,请用 jq 查询")
		}
		if output, err = docMarshalIndent(doc); err != nil {
			return &CLIError{Message: err.Error(), ExitCode: 1}
		}
	default:
		return &CLIError{Message: fmt.Sprintf("未知格式 %q(可选 outline|md|content-list|json)", docParseOpts.format), ExitCode: 2}
	}

	if docParseOpts.out == "" {
		if _, err := cmd.OutOrStdout().Write(output); err != nil {
			return err
		}
		if !bytesHasSuffix(output, '\n') {
			cmd.Println()
		}
		return nil
	}
	return os.WriteFile(docParseOpts.out, output, 0o644)
}

// docInputName 返回读取前可确定的文件名，用于决定是否启用 PDF 输入限流。
func docInputName(source, name string) string {
	if source == "-" {
		return name
	}
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		parsed, err := url.Parse(source)
		if err == nil {
			return path.Base(parsed.Path)
		}
	}
	return source
}

// readDocSource 读取本地文件、http(s) URL 或 stdin。maxBytes 为正时在
// 分配完整缓冲前拒绝超限输入；0 表示沿用非 PDF 的历史限制。
func readDocSource(source, name string, maxBytes int64) (string, []byte, error) {
	if source == "-" {
		if name == "" {
			return "", nil, fmt.Errorf("stdin 模式需要 --name 提供文件名(如 report.xlsx)")
		}
		data, err := readDocInputLimited(os.Stdin, maxBytes)
		return name, data, err
	}
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return downloadDocFile(source, maxBytes)
	}
	info, err := os.Stat(source)
	if err != nil {
		return "", nil, fmt.Errorf("输入不存在且不是 URL: %s", source)
	}
	if maxBytes > 0 && info.Size() > maxBytes {
		return "", nil, fmt.Errorf("PDF 文件超过 %d MiB 上限", maxBytes>>20)
	}
	data, err := os.ReadFile(source)
	return source, data, err
}

// docDownloadTimeout / docDownloadMaxBytes 控制 URL 下载。
const (
	docDownloadTimeout        = 120 * time.Second
	docDownloadMaxBytes int64 = 200 << 20
)

// downloadDocFile 下载 URL 文件(120s 超时、200MB 上限),文件名取路径末段。
func downloadDocFile(source string, maxBytes int64) (string, []byte, error) {
	parsed, err := url.Parse(source)
	if err != nil || parsed.Host == "" {
		return "", nil, fmt.Errorf("非法 URL: %s", source)
	}
	invokable := &http.Client{Timeout: docDownloadTimeout}
	httpResp, err := invokable.Get(source)
	if err != nil {
		return "", nil, fmt.Errorf("下载失败: %w", err)
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("下载失败: http %d", httpResp.StatusCode)
	}
	readLimit := docDownloadMaxBytes
	if maxBytes > 0 && maxBytes < readLimit {
		readLimit = maxBytes
	}
	if httpResp.ContentLength > readLimit {
		return "", nil, fmt.Errorf("文件超过 %d MiB 上限", readLimit>>20)
	}
	data, err := readDocInputLimited(httpResp.Body, readLimit)
	if err != nil {
		return "", nil, err
	}
	name := path.Base(parsed.Path)
	if name == "" || name == "/" || name == "." {
		name = "download"
	}
	return name, data, nil
}

// readDocInputLimited 使用 limit+1 探测超限，避免 URL 或 stdin 无界读入内存。
func readDocInputLimited(reader io.Reader, limit int64) ([]byte, error) {
	if limit <= 0 {
		return io.ReadAll(reader)
	}
	data, err := io.ReadAll(io.LimitReader(reader, limit+1))
	if err != nil {
		return nil, fmt.Errorf("读取文档失败: %w", err)
	}
	if int64(len(data)) > limit {
		return nil, fmt.Errorf("文件超过 %d MiB 上限", limit>>20)
	}
	return data, nil
}

// platformLLMClient 通过平台 /api/v1/ai/chat/completions(agentID=0 裸模型
// 代理)实现 llmocr.LLMClient;认证与 baseURL 复用 ur 既有 client/auth 体系。
type platformLLMClient struct {
	modelType string
}

// Complete 组装平台自有 JSON 格式的对话请求并提取回复文本。
func (p platformLLMClient) Complete(ctx context.Context, req llmocr.VisionRequest) (string, error) {
	contents := make([]map[string]any, 0, len(req.Images)+1)
	if prompt := strings.TrimSpace(req.Prompt); prompt != "" {
		contents = append(contents, map[string]any{"type": "text", "text": prompt})
	}
	for _, img := range req.Images {
		contents = append(contents, map[string]any{"type": "image_url", "imageUrl": img.DataURI})
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/ai/chat/completions",
		Body: map[string]any{
			"agentID":   0,
			"modelType": p.modelType,
			"stream":    false,
			"messages":  []map[string]any{{"role": "user", "contents": contents}},
		},
	})
	if err != nil {
		return "", fmt.Errorf("platform ocr call: %w", err)
	}
	if resp.Code != 200 {
		return "", fmt.Errorf("platform ocr code=%d msg=%s", resp.Code, resp.Msg)
	}
	data, ok := resp.Data.(map[string]any)
	if !ok {
		return "", fmt.Errorf("platform ocr: unexpected data type %T", resp.Data)
	}
	content, _ := data["content"].(string)
	return content, nil
}

// bytesHasSuffix 判断输出是否以指定字节结尾。
func bytesHasSuffix(b []byte, c byte) bool {
	return len(b) > 0 && b[len(b)-1] == c
}

// docMarshalIndent 输出带尾行换行的缩进 JSON。
func docMarshalIndent(v any) ([]byte, error) {
	raw, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(raw, '\n'), nil
}
