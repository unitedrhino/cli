package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func runModelTemplate(args []string, stdout, stderr io.Writer) int {
	templateType := "full"
	format := "json"
	outputPath := ""

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "property", "event", "action", "full":
			templateType = args[i]
		case "--json":
			format = "json"
		case "--yaml":
			format = "yaml"
		case "--output":
			if i+1 < len(args) {
				outputPath = args[i+1]
				i++
			}
		}
	}

	var content []byte
	var err error

	switch templateType {
	case "property":
		content, err = modelPropertyTemplate(format)
	case "event":
		content, err = modelEventTemplate(format)
	case "action":
		content, err = modelActionTemplate(format)
	case "full":
		content, err = modelFullTemplate(format)
	}

	if err != nil {
		fmt.Fprintf(stderr, "生成模板失败: %v\n", err)
		return 1
	}

	if outputPath != "" {
		if err := os.WriteFile(outputPath, content, 0644); err != nil {
			fmt.Fprintf(stderr, "写入文件失败: %v\n", err)
			return 1
		}
		fmt.Fprintf(stdout, "模板已保存: %s\n", outputPath)
	} else {
		fmt.Fprintln(stdout, string(content))
	}
	return 0
}

func modelPropertyTemplate(format string) ([]byte, error) {
	tmpl := map[string]any{
		"identifier": "temperature",
		"name":       "温度",
		"desc":       "设备当前温度",
		"mode":       "r",
		"define": map[string]any{
			"type": "float",
			"min":  "-40",
			"max":  "80",
			"step": "0.1",
			"unit": "°C",
		},
		"isUseShadow":       false,
		"isCanSceneLinkage": 1,
		"funcGroup":         1,
		"userPerm":          1,
		"required":          true,
		"order":             1,
	}
	if format == "yaml" {
		return yamlMarshal(tmpl)
	}
	return json.MarshalIndent(tmpl, "", "  ")
}

func modelEventTemplate(format string) ([]byte, error) {
	tmpl := map[string]any{
		"identifier": "alarm",
		"name":       "告警事件",
		"desc":       "设备异常告警",
		"type":       "alert",
		"dir":        "up",
		"params": []map[string]any{
			{
				"identifier": "code",
				"name":       "告警码",
				"define": map[string]any{
					"type": "string",
				},
			},
			{
				"identifier": "message",
				"name":       "告警信息",
				"define": map[string]any{
					"type": "string",
				},
			},
		},
		"required": true,
		"order":    1,
	}
	if format == "yaml" {
		return yamlMarshal(tmpl)
	}
	return json.MarshalIndent(tmpl, "", "  ")
}

func modelActionTemplate(format string) ([]byte, error) {
	tmpl := map[string]any{
		"identifier": "switch",
		"name":       "开关控制",
		"desc":       "控制设备开关状态",
		"dir":        "down",
		"input": []map[string]any{
			{
				"identifier": "state",
				"name":       "开关状态",
				"define": map[string]any{
					"type": "bool",
				},
				"required": true,
			},
		},
		"output": []map[string]any{
			{
				"identifier": "result",
				"name":       "执行结果",
				"define": map[string]any{
					"type": "bool",
				},
				"required": true,
			},
		},
		"required": true,
		"order":    1,
	}
	if format == "yaml" {
		return yamlMarshal(tmpl)
	}
	return json.MarshalIndent(tmpl, "", "  ")
}

func modelFullTemplate(format string) ([]byte, error) {
	tmpl := map[string]any{
		"version": "1.0",
		"properties": []map[string]any{
			{
				"identifier": "temperature",
				"name":       "温度",
				"desc":       "设备当前温度",
				"mode":       "r",
				"define": map[string]any{
					"type": "float",
					"min":  "-40",
					"max":  "80",
					"step": "0.1",
					"unit": "°C",
				},
				"isUseShadow":       false,
				"isCanSceneLinkage": 1,
				"funcGroup":         1,
				"userPerm":          1,
				"required":          true,
				"order":             1,
			},
			{
				"identifier": "humidity",
				"name":       "湿度",
				"desc":       "环境湿度",
				"mode":       "r",
				"define": map[string]any{
					"type": "float",
					"min":  "0",
					"max":  "100",
					"step": "0.1",
					"unit": "%",
				},
				"isUseShadow":       false,
				"isCanSceneLinkage": 1,
				"funcGroup":         1,
				"userPerm":          1,
				"required":          true,
				"order":             2,
			},
		},
		"events": []map[string]any{
			{
				"identifier": "alarm",
				"name":       "告警事件",
				"desc":       "设备异常告警",
				"type":       "alert",
				"dir":        "up",
				"params": []map[string]any{
					{
						"identifier": "code",
						"name":       "告警码",
						"define": map[string]any{
							"type": "string",
						},
					},
				},
				"required": true,
				"order":    1,
			},
		},
		"actions": []map[string]any{
			{
				"identifier": "switch",
				"name":       "开关控制",
				"desc":       "控制设备开关状态",
				"dir":        "down",
				"input": []map[string]any{
					{
						"identifier": "state",
						"name":       "开关状态",
						"define": map[string]any{
							"type": "bool",
						},
						"required": true,
					},
				},
				"output": []map[string]any{
					{
						"identifier": "result",
						"name":       "执行结果",
						"define": map[string]any{
							"type": "bool",
						},
						"required": true,
					},
				},
				"required": true,
				"order":    1,
			},
		},
	}
	if format == "yaml" {
		return yamlMarshal(tmpl)
	}
	return json.MarshalIndent(tmpl, "", "  ")
}

// yamlMarshal 简单的 YAML 序列化（用缩进实现，不引入外部库）
func yamlMarshal(v map[string]any) ([]byte, error) {
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
	return fmt.Sprintf("%v", v)
}
