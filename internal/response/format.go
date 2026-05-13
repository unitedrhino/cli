package response

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

var ValidFormats = []string{"json", "raw", "yaml"}

// FormatOptions 控制 API 响应的输出格式。
type FormatOptions struct {
	Format    string // json, raw, yaml
	Transform string // GJSON 路径表达式
}

// FormatOutput 将 API 响应按指定格式序列化输出。
func FormatOutput(data any, opts FormatOptions) ([]byte, error) {
	// 1. 若有 transform，先提取
	if opts.Transform != "" {
		extracted, err := applyTransform(data, opts.Transform)
		if err != nil {
			return nil, err
		}
		data = extracted
	}

	// 2. 按 format 序列化
	switch opts.Format {
	case "", "json":
		return json.MarshalIndent(data, "", "  ")
	case "raw":
		return json.Marshal(data)
	case "yaml":
		return yamlMarshal(data)
	default:
		return nil, fmt.Errorf("unsupported format %q, valid: %s", opts.Format, strings.Join(ValidFormats, ", "))
	}
}

// applyTransform 用 GJSON 语法从 data 中提取值。
func applyTransform(data any, path string) (any, error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal for transform: %w", err)
	}
	result := gjson.GetBytes(raw, path)
	if !result.Exists() {
		return nil, fmt.Errorf("transform path %q not found", path)
	}
	// 将 gjson.Result 转回 Go 值
	var out any
	if err := json.Unmarshal([]byte(result.Raw), &out); err != nil {
		// 非 JSON 值（如字符串、数字）直接返回
		return result.Value(), nil
	}
	return out, nil
}

// yamlMarshal 简单的 YAML 序列化（不引入外部库）。
func yamlMarshal(v any) ([]byte, error) {
	var b []byte
	appendYAML(&b, v, 0)
	return b, nil
}

func appendYAML(b *[]byte, v any, indent int) {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += "  "
	}
	switch val := v.(type) {
	case map[string]any:
		for k, vv := range val {
			if m, ok := vv.(map[string]any); ok && len(m) > 0 {
				*b = append(*b, prefix+k+":\n"...)
				appendYAML(b, m, indent+1)
			} else if a, ok := vv.([]any); ok {
				appendYAMLArray(b, k, a, prefix, indent)
			} else if a, ok := vv.([]map[string]any); ok {
				items := make([]any, len(a))
				for i, item := range a {
					items[i] = item
				}
				appendYAMLArray(b, k, items, prefix, indent)
			} else {
				*b = append(*b, prefix+k+": "+yamlValue(vv)+"\n"...)
			}
		}
	default:
		*b = append(*b, prefix+yamlValue(v)+"\n"...)
	}
}

func appendYAMLArray(b *[]byte, k string, a []any, prefix string, indent int) {
	if len(a) == 0 {
		*b = append(*b, prefix+k+": []\n"...)
		return
	}
	*b = append(*b, prefix+k+":\n"...)
	for _, item := range a {
		if m, isMap := item.(map[string]any); isMap && len(m) > 0 {
			*b = append(*b, prefix+"  -\n"...)
			appendYAML(b, item, indent+2)
		} else {
			*b = append(*b, prefix+"  - "+yamlValue(item)+"\n"...)
		}
	}
}

func yamlValue(v any) string {
	if s, ok := v.(string); ok {
		if strings.ContainsAny(s, ":#{}[]|>") || strings.HasPrefix(s, " ") || strings.HasSuffix(s, " ") {
			return "\"" + s + "\""
		}
		return s
	}
	if b, ok := v.(bool); ok {
		if b {
			return "true"
		}
		return "false"
	}
	if v == nil {
		return "null"
	}
	return fmt.Sprintf("%v", v)
}
