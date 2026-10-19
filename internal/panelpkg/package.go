// 文件说明：校验并归一化不可信 HTML 面板包，生成不依赖 ZIP 元数据的内容摘要；不执行包内脚本。
package panelpkg

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"path"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

// SDKVersion 是当前正式面板支持的 SDK 合同版本。
const SDKVersion = "1.0.0"

// MaxCompressed/MaxExpanded/MaxFiles 限定面板包的压缩大小、展开大小与文件数，防止解包资源耗尽。
const (
	MaxCompressed = 10 << 20
	MaxExpanded   = 30 << 20
	MaxFiles      = 100
)

// Requirement 声明面板需要的单设备物模型能力；Define/Input 保留平台的参数定义结构。
type Requirement struct {
	Kind       string         `json:"kind"`             // 属性、下行行为或事件命名空间。
	Identifier string         `json:"identifier"`       // 真实物模型标识符。
	Required   bool           `json:"required"`         // 缺少时是否必须拒绝绑定。
	Define     map[string]any `json:"define,omitempty"` // 属性类型、单位及枚举要求。
	Access     string         `json:"access,omitempty"` // 属性 read/write 访问需求。
	Dir        string         `json:"dir,omitempty"`    // 行为方向，只接受 down。
	Input      map[string]any `json:"input,omitempty"`  // 行为输入参数合同。
}

// UnmarshalJSON 强制能力声明显式选择 required，避免遗漏字段被静默当成可选能力。
func (r *Requirement) UnmarshalJSON(data []byte) error {
	type rawRequirement Requirement
	var fields map[string]json.RawMessage
	if e := json.Unmarshal(data, &fields); e != nil {
		return e
	}
	required, exists := fields["required"]
	if !exists || string(required) == "null" {
		return fmt.Errorf("能力声明必须显式填写布尔 required")
	}
	var value rawRequirement
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if e := decoder.Decode(&value); e != nil {
		return e
	}
	*r = Requirement(value)
	return nil
}

// Manifest 是包的不可变元数据与能力合同，不描述 UI 组件和布局。
type Manifest struct {
	FormatVersion int           `json:"formatVersion"`     // 包格式版本，固定为 1。
	Name          string        `json:"name"`              // 可展示的面板名称。
	Desc          string        `json:"desc,omitempty"`    // 作者说明。
	Version       string        `json:"version,omitempty"` // 作者内容版本，不是可变发布指针。
	Entry         string        `json:"entry"`             // 根入口 index.html。
	SDKVersion    string        `json:"sdkVersion"`        // 运行所需 SDK 版本。
	Platforms     []string      `json:"platforms"`         // 作者声明的终端能力范围。
	Requirements  []Requirement `json:"requirements"`      // 所需物模型功能，可为空数组。
}

// File 描述归一化运行文件的字节数和摘要，供客户端下载校验。
type File struct {
	Path   string `json:"path"`
	Bytes  int64  `json:"bytes"`
	SHA256 string `json:"sha256"`
}

// Package 保存校验后的运行文件；只有这个对象才能写入面板资产区域。
type Package struct {
	Manifest   Manifest
	Files      map[string][]byte
	Hash       string
	Inventory  []File
	TotalBytes int64
}

// ValidateManifest 检查包字段、能力声明及可支持的 SDK，返回明确的输入错误。
func ValidateManifest(m Manifest) error {
	if m.FormatVersion != 1 || m.Entry != "index.html" || m.SDKVersion != SDKVersion {
		return fmt.Errorf("需要 formatVersion=1、entry=index.html、sdkVersion=%s", SDKVersion)
	}
	if strings.TrimSpace(m.Name) == "" || utf8.RuneCountInString(m.Name) > 80 || utf8.RuneCountInString(m.Desc) > 2000 || len(m.Version) > 64 {
		return fmt.Errorf("面板名称、说明或内容版本无效")
	}
	if len(m.Platforms) == 0 || m.Requirements == nil || len(m.Requirements) > MaxFiles {
		return fmt.Errorf("需要 platforms 和 requirements 数组")
	}
	seen := map[string]bool{}
	for _, p := range m.Platforms {
		if seen[p] || !contains([]string{"web", "h5", "appAndroid", "appIos", "mini"}, p) {
			return fmt.Errorf("终端声明无效或重复：%s", p)
		}
		seen[p] = true
	}
	seen = map[string]bool{}
	for _, r := range m.Requirements {
		key := r.Kind + ":" + r.Identifier
		if seen[key] || !regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{0,99}$`).MatchString(r.Identifier) || contains([]string{"constructor", "prototype", "__proto__"}, r.Identifier) {
			return fmt.Errorf("功能标识符无效或重复：%s", key)
		}
		seen[key] = true
		switch r.Kind {
		case "property":
			if !contains([]string{"read", "write"}, r.Access) || !contains([]string{"bool", "int", "float", "string", "enum", "timestamp", "struct", "array"}, stringValue(r.Define["type"])) {
				return fmt.Errorf("属性 %s 缺少有效类型或访问声明", r.Identifier)
			}
		case "action":
			if r.Dir != "down" {
				return fmt.Errorf("行为 %s 只允许 down", r.Identifier)
			}
		case "event":
			if r.Required {
				return fmt.Errorf("当前 SDK 不支持必需的物模型事件：%s", r.Identifier)
			}
		default:
			return fmt.Errorf("未知功能命名空间：%s", r.Kind)
		}
	}
	return nil
}

// Decode 将 HTML 或 ZIP 输入解析成不可变、可重复导出的运行包；文件名只用于判断格式，不用于归属。
func Decode(name string, data []byte) (*Package, error) {
	if len(data) == 0 || len(data) > MaxCompressed {
		return nil, fmt.Errorf("面板文件为空或超过 10 MiB")
	}
	files := map[string][]byte{}
	if strings.EqualFold(path.Ext(name), ".html") || strings.EqualFold(path.Ext(name), ".htm") {
		files["index.html"] = data
	} else if strings.EqualFold(path.Ext(name), ".zip") {
		z, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
		if err != nil {
			return nil, fmt.Errorf("ZIP 无效：%w", err)
		}
		if len(z.File) > MaxFiles {
			return nil, fmt.Errorf("ZIP 文件数超过 %d", MaxFiles)
		}
		var total int64
		for _, f := range z.File {
			if f.FileInfo().IsDir() {
				if !validPath(strings.TrimSuffix(f.Name, "/")) {
					return nil, fmt.Errorf("ZIP 目录路径无效")
				}
				continue
			}
			if !validPath(f.Name) || !f.Mode().IsRegular() || f.Flags&1 != 0 || !allowedExtension(f.Name) {
				return nil, fmt.Errorf("ZIP 文件路径或类型无效：%s", f.Name)
			}
			if _, ok := files[f.Name]; ok {
				return nil, fmt.Errorf("ZIP 存在重复文件：%s", f.Name)
			}
			if f.UncompressedSize64 > MaxExpanded || total+int64(f.UncompressedSize64) > MaxExpanded {
				return nil, fmt.Errorf("ZIP 展开超过 30 MiB")
			}
			reader, e := f.Open()
			if e != nil {
				return nil, e
			}
			content, e := io.ReadAll(io.LimitReader(reader, MaxExpanded-total+1))
			reader.Close()
			if e != nil {
				return nil, fmt.Errorf("ZIP 文件校验失败：%w", e)
			}
			total += int64(len(content))
			if total > MaxExpanded {
				return nil, fmt.Errorf("ZIP 实际展开超限")
			}
			files[f.Name] = content
		}
	} else {
		return nil, fmt.Errorf("只支持 HTML 或 ZIP")
	}
	return normalize(files)
}

// normalize 提取内嵌清单、统一入口序列化、检查静态依赖并计算规范化内容指纹。
func normalize(files map[string][]byte) (*Package, error) {
	entry, ok := files["index.html"]
	if !ok {
		return nil, fmt.Errorf("缺少根入口 index.html")
	}
	root, err := html.Parse(bytes.NewReader(entry))
	if err != nil {
		return nil, err
	}
	manifestBytes := files["manifest.json"]
	embedded := 0
	var walk func(*html.Node) error
	walk = func(n *html.Node) error {
		for child := n.FirstChild; child != nil; {
			next := child.NextSibling
			if child.Type == html.ElementNode && child.Data == "script" && attr(child, "id") == "ur-panel-manifest" {
				embedded++
				if embedded > 1 || attr(child, "type") != "application/json" {
					return fmt.Errorf("重复或无效的内嵌清单")
				}
				text := ""
				for c := child.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						text += c.Data
					}
				}
				if len(manifestBytes) > 0 {
					a, e := canonicalJSON(manifestBytes)
					if e != nil {
						return e
					}
					b, e := canonicalJSON([]byte(text))
					if e != nil {
						return e
					}
					if !bytes.Equal(a, b) {
						return fmt.Errorf("内嵌清单与根清单不一致")
					}
				} else {
					manifestBytes = []byte(text)
				}
				n.RemoveChild(child)
			} else {
				if e := walk(child); e != nil {
					return e
				}
			}
			child = next
		}
		return nil
	}
	if err = walk(root); err != nil {
		return nil, err
	}
	var m Manifest
	decoder := json.NewDecoder(bytes.NewReader(manifestBytes))
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(&m); err != nil {
		return nil, fmt.Errorf("清单无效：%w", err)
	}
	if decoder.Decode(new(any)) != io.EOF {
		return nil, fmt.Errorf("清单包含多余 JSON")
	}
	if err = ValidateManifest(m); err != nil {
		return nil, err
	}
	var rendered bytes.Buffer
	if err = html.Render(&rendered, root); err != nil {
		return nil, err
	}
	files["index.html"] = rendered.Bytes()
	metadata, _ := json.Marshal(m)
	files["manifest.json"], err = canonicalJSON(metadata)
	if err != nil {
		return nil, err
	}
	for name, data := range files {
		if strings.EqualFold(path.Ext(name), ".html") {
			doc, e := html.Parse(bytes.NewReader(data))
			if e != nil {
				return nil, e
			}
			if e = checkHTMLResources(doc, name, files); e != nil {
				return nil, e
			}
			// 文件方式运行没有 HTTP 响应头，必须把相同限制固定进包内 HTML。
			installCSP(doc)
			var secured bytes.Buffer
			if e = html.Render(&secured, doc); e != nil {
				return nil, e
			}
			files[name] = secured.Bytes()
		}
		if strings.EqualFold(path.Ext(name), ".css") {
			if e := checkCSSResources(string(data), name, files); e != nil {
				return nil, e
			}
		}
	}
	p := &Package{Manifest: m, Files: files, Inventory: []File{}}
	keys := make([]string, 0, len(files))
	for name := range files {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	var fingerprint strings.Builder
	for _, name := range keys {
		b := files[name]
		sum := sha256.Sum256(b)
		digest := hex.EncodeToString(sum[:])
		p.Inventory = append(p.Inventory, File{name, int64(len(b)), digest})
		p.TotalBytes += int64(len(b))
		fmt.Fprintf(&fingerprint, "%d:%s:%d:%s\n", len([]byte(name)), name, len(b), digest)
	}
	if len(p.Inventory) > MaxFiles || p.TotalBytes > MaxExpanded {
		return nil, fmt.Errorf("规范化包展开超限")
	}
	sum := sha256.Sum256([]byte(fingerprint.String()))
	p.Hash = hex.EncodeToString(sum[:])
	return p, nil
}

// ZIP 导出标准 DEFLATE 包，避免可导入的压缩资源在导出时膨胀超出回导限制。
func (p *Package) ZIP() ([]byte, error) {
	var out bytes.Buffer
	w := zip.NewWriter(&out)
	for _, f := range p.Inventory {
		header := &zip.FileHeader{Name: f.Path, Method: zip.Deflate}
		entry, e := w.CreateHeader(header)
		if e != nil {
			return nil, e
		}
		if _, e = entry.Write(p.Files[f.Path]); e != nil {
			return nil, e
		}
	}
	if e := w.Close(); e != nil {
		return nil, e
	}
	if out.Len() > MaxCompressed {
		return nil, fmt.Errorf("标准导出包超过 10 MiB")
	}
	return out.Bytes(), nil
}

// canonicalJSON 对对象键排序，保留数组顺序与数值原文，拒绝多段 JSON。
func canonicalJSON(data []byte) ([]byte, error) {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	var value any
	if e := d.Decode(&value); e != nil {
		return nil, e
	}
	if d.Decode(new(any)) != io.EOF {
		return nil, fmt.Errorf("多余 JSON")
	}
	return json.Marshal(value)
}

// contains 判断声明值是否属于固定白名单。
func contains(items []string, v string) bool {
	for _, x := range items {
		if x == v {
			return true
		}
	}
	return false
}

// stringValue 读取参数合同中的字符串字段，其他类型返回空串以触发校验失败。
func stringValue(v any) string { s, _ := v.(string); return s }

// validPath 只接受包内无歧义相对路径，拒绝编码穿越、控制符和操作系统差异路径。
func validPath(p string) bool {
	return p != "" && len(strings.Split(p, "/")) <= 8 && p == path.Clean(p) && !strings.HasPrefix(p, "/") && !strings.ContainsAny(p, "\\:%?#\x00\r\n") && p != ".." && !strings.HasPrefix(p, "../")
}

// allowedExtension 将可托管资源限制为静态网页文件。
func allowedExtension(p string) bool {
	return contains([]string{".html", ".js", ".mjs", ".css", ".json", ".svg", ".png", ".jpg", ".jpeg", ".gif", ".webp", ".ico", ".woff", ".woff2", ".ttf"}, strings.ToLower(path.Ext(p)))
}

// attr 获取 HTML 属性，缺省返回空串。
func attr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

// checkReference 校验资源为包内相对路径或内联 data，禁止外部依赖和路径逃逸。
func checkReference(ref, from string, files map[string][]byte) error {
	ref = strings.TrimSpace(ref)
	if ref == "" || strings.HasPrefix(ref, "#") || strings.HasPrefix(ref, "data:") {
		return nil
	}
	clean := strings.Split(strings.Split(ref, "#")[0], "?")[0]
	if !validPath(clean) {
		return fmt.Errorf("%s 含无效或外部资源：%s", from, ref)
	}
	target := path.Join(path.Dir(from), clean)
	if !validPath(target) {
		return fmt.Errorf("资源路径逃逸：%s", ref)
	}
	if _, ok := files[target]; !ok {
		return fmt.Errorf("%s 缺少资源 %s", from, target)
	}
	return nil
}

// checkHTMLResources 校验 HTML 声明的静态资源及内联 CSS；网络权限另由宿主 CSP 限定。
func checkHTMLResources(n *html.Node, from string, files map[string][]byte) error {
	if n.Type == html.ElementNode {
		if contains([]string{"base", "iframe", "object", "embed"}, n.Data) {
			return fmt.Errorf("面板禁止嵌入 base/iframe/object/embed")
		}
		for _, a := range n.Attr {
			if a.Key == "src" || a.Key == "poster" || (a.Key == "href" && n.Data == "link") {
				if e := checkReference(a.Val, from, files); e != nil {
					return e
				}
			}
			if a.Key == "srcset" {
				return fmt.Errorf("请将 srcset 资源改为明确包内 src")
			}
			if a.Key == "style" {
				if e := checkCSSResources(a.Val, from, files); e != nil {
					return e
				}
			}
		}
		if n.Data == "style" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					if e := checkCSSResources(c.Data, from, files); e != nil {
						return e
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if e := checkHTMLResources(c, from, files); e != nil {
			return e
		}
	}
	return nil
}

// checkCSSResources 拒绝外部 import，并检查 url() 中的包内依赖。
func checkCSSResources(css, from string, files map[string][]byte) error {
	if regexp.MustCompile(`(?i)@import\b`).MatchString(css) {
		return fmt.Errorf("请将 CSS import 合并为包内样式文件")
	}
	for _, m := range regexp.MustCompile(`(?i)url\(\s*["']?([^\s"')]+)["']?\s*\)`).FindAllStringSubmatch(css, -1) {
		if e := checkReference(m[1], from, files); e != nil {
			return e
		}
	}
	return nil
}

// installCSP 将固定资源限制放在文档头首位；重复归一化先移除自己的旧标记，保持摘要稳定。
func installCSP(root *html.Node) {
	var head *html.Node
	var walk func(*html.Node)
	walk = func(node *html.Node) {
		for child := node.FirstChild; child != nil; {
			next := child.NextSibling
			if child.Type == html.ElementNode && child.Data == "meta" && attr(child, "id") == "ur-panel-csp" {
				node.RemoveChild(child)
			} else {
				if child.Type == html.ElementNode && child.Data == "head" {
					head = child
				}
				walk(child)
			}
			child = next
		}
	}
	walk(root)
	if head == nil {
		return
	}
	meta := &html.Node{Type: html.ElementNode, Data: "meta", Attr: []html.Attribute{{Key: "id", Val: "ur-panel-csp"}, {Key: "http-equiv", Val: "Content-Security-Policy"}, {Key: "content", Val: "default-src 'none'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: blob:; font-src 'self' data:; connect-src 'none'; base-uri 'none'; form-action 'none'; frame-src 'none'"}}}
	head.InsertBefore(meta, head.FirstChild)
}
