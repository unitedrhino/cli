package response

import (
	"fmt"
	"strings"
)

// FilterFields 根据 selectors 裁剪 data，保留指定字段。
// selector 语法：
//   - "a.b.c"       保留 a.b.c 完整子树
//   - "a.b.{x,y}"   只保留 a.b 下的 x 和 y 字段
func FilterFields(data map[string]any, selectors []string) (map[string]any, error) {
	if len(selectors) == 0 {
		return data, nil
	}

	parsed := make([]selector, 0, len(selectors))
	for _, s := range selectors {
		sel, err := parseSelector(strings.TrimSpace(s))
		if err != nil {
			return nil, fmt.Errorf("parse selector %q: %w", s, err)
		}
		parsed = append(parsed, sel)
	}

	out := make(map[string]any)
	for _, sel := range parsed {
		applySelector(out, data, sel)
	}
	return out, nil
}

type selector struct {
	parts  []string   // 路径段，如 ["data", "list"]
	fields []string   // 若最后一段是 {a,b}，则 fields=[a,b]；否则 nil
}

func parseSelector(s string) (selector, error) {
	if s == "" {
		return selector{}, fmt.Errorf("empty selector")
	}

	// 检查最后一段是否是 {a,b,c}
	if idx := strings.LastIndex(s, ".{"); idx != -1 && strings.HasSuffix(s, "}") {
		prefix := s[:idx]
		inner := s[idx+2 : len(s)-1] // skip ".{" and "}"
		fields := splitFields(inner)
		if prefix == "" {
			return selector{parts: []string{}, fields: fields}, nil
		}
		return selector{parts: strings.Split(prefix, "."), fields: fields}, nil
	}

	return selector{parts: strings.Split(s, ".")}, nil
}

func splitFields(s string) []string {
	var out []string
	for _, f := range strings.Split(s, ",") {
		f = strings.TrimSpace(f)
		if f != "" {
			out = append(out, f)
		}
	}
	return out
}

func applySelector(dst, src map[string]any, sel selector) {
	if len(sel.parts) == 0 {
		// 根级选择，如 "{a,b}"
		if sel.fields != nil {
			for _, f := range sel.fields {
				if v, ok := src[f]; ok {
					dst[f] = v
				}
			}
		}
		return
	}

	walkAndSet(dst, src, sel.parts, 0, sel.fields)
}

func walkAndSet(dst, src map[string]any, parts []string, idx int, fields []string) {
	key := parts[idx]
	v, ok := src[key]
	if !ok {
		return
	}

	if idx == len(parts)-1 {
		// 最后一段
		if fields != nil {
			// 只保留指定子字段
			sub, isMap := v.(map[string]any)
			if !isMap {
				dst[key] = v
				return
			}
			filtered := make(map[string]any)
			for _, f := range fields {
				if fv, ok := sub[f]; ok {
					filtered[f] = fv
				}
			}
			dst[key] = filtered
		} else {
			dst[key] = deepCopy(v)
		}
		return
	}

	sub, isMap := v.(map[string]any)
	if !isMap {
		// 中间路径不是 map，无法继续深入，直接保留
		dst[key] = deepCopy(v)
		return
	}

	nextDst, ok := dst[key].(map[string]any)
	if !ok {
		nextDst = make(map[string]any)
		dst[key] = nextDst
	}
	walkAndSet(nextDst, sub, parts, idx+1, fields)
}

func deepCopy(v any) any {
	switch x := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, vv := range x {
			out[k] = deepCopy(vv)
		}
		return out
	case []any:
		out := make([]any, len(x))
		for i, vv := range x {
			out[i] = deepCopy(vv)
		}
		return out
	default:
		return x
	}
}
