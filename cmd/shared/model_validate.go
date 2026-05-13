package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func runModelValidate(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: model validate <file>")
		return 2
	}

	filePath := args[0]
	var content []byte
	var err error

	if filePath == "-" {
		content, err = io.ReadAll(os.Stdin)
	} else {
		content, err = os.ReadFile(filePath)
	}
	if err != nil {
		fmt.Fprintf(stderr, "读取失败: %v\n", err)
		return 1
	}

	var data map[string]any
	if err := json.Unmarshal(content, &data); err != nil {
		fmt.Fprintf(stderr, "JSON 解析错误: %v\n", err)
		return 1
	}

	v := newModelValidator()
	v.validateModel(data, "")

	if len(v.errors) > 0 {
		fmt.Fprintf(stdout, "❌ 校验失败，共 %d 个错误:\n", len(v.errors))
		for _, e := range v.errors {
			fmt.Fprintf(stdout, "  %s\n", e)
		}
		return 1
	}

	if len(v.warnings) > 0 {
		fmt.Fprintf(stdout, "✅ 结构校验通过\n")
		fmt.Fprintf(stdout, "⚠️  共 %d 个警告:\n", len(v.warnings))
		for _, w := range v.warnings {
			fmt.Fprintf(stdout, "  %s\n", w)
		}
	} else {
		fmt.Fprintln(stdout, "✅ 校验通过")
	}
	return 0
}

type modelValidator struct {
	errors   []string
	warnings []string
}

func newModelValidator() *modelValidator {
	return &modelValidator{}
}

func (v *modelValidator) error(path, msg string) {
	v.errors = append(v.errors, fmt.Sprintf("[%s] %s", path, msg))
}

func (v *modelValidator) warn(path, msg string) {
	v.warnings = append(v.warnings, fmt.Sprintf("[%s] %s", path, msg))
}

var validDataTypes = map[string]bool{
	"bool":      true,
	"int":       true,
	"float":     true,
	"string":    true,
	"struct":    true,
	"enum":      true,
	"array":     true,
	"timestamp": true,
	"matrix":    true,
}

var validModes = map[string]bool{
	"r":  true,
	"rw": true,
}

var validEventTypes = map[string]bool{
	"info":   true,
	"alert":  true,
	"fault":  true,
}

var validDirs = map[string]bool{
	"up":   true,
	"down": true,
}

func (v *modelValidator) validateModel(data map[string]any, path string) {
	if data == nil {
		v.error(path, "物模型 JSON 应为 object 类型")
		return
	}

	// 校验 properties
	if props, ok := data["properties"].([]any); ok {
		for i, p := range props {
			if prop, ok := p.(map[string]any); ok {
				v.validateProperty(prop, fmt.Sprintf("%s.properties[%d]", path, i))
			} else {
				v.error(fmt.Sprintf("%s.properties[%d]", path, i), "应为 object 类型")
			}
		}
	}

	// 校验 events
	if events, ok := data["events"].([]any); ok {
		for i, e := range events {
			if evt, ok := e.(map[string]any); ok {
				v.validateEvent(evt, fmt.Sprintf("%s.events[%d]", path, i))
			} else {
				v.error(fmt.Sprintf("%s.events[%d]", path, i), "应为 object 类型")
			}
		}
	}

	// 校验 actions
	if actions, ok := data["actions"].([]any); ok {
		for i, a := range actions {
			if act, ok := a.(map[string]any); ok {
				v.validateAction(act, fmt.Sprintf("%s.actions[%d]", path, i))
			} else {
				v.error(fmt.Sprintf("%s.actions[%d]", path, i), "应为 object 类型")
			}
		}
	}
}

func (v *modelValidator) validateProperty(data map[string]any, path string) {
	v.validateString(path+".identifier", data["identifier"], true)
	v.validateString(path+".name", data["name"], true)

	if define, ok := data["define"].(map[string]any); ok {
		v.validateDefine(define, path+".define")
	} else {
		v.error(path+".define", "必填字段缺失")
	}

	if mode, ok := data["mode"].(string); ok {
		if !validModes[mode] {
			v.error(path+".mode", fmt.Sprintf("非法值 '%s'，允许: r, rw", mode))
		}
	}
}

func (v *modelValidator) validateEvent(data map[string]any, path string) {
	v.validateString(path+".identifier", data["identifier"], true)
	v.validateString(path+".name", data["name"], true)

	if evtType, ok := data["type"].(string); ok {
		if !validEventTypes[evtType] {
			v.error(path+".type", fmt.Sprintf("非法值 '%s'，允许: info, alert, fault", evtType))
		}
	}

	if dir, ok := data["dir"].(string); ok {
		if !validDirs[dir] {
			v.error(path+".dir", fmt.Sprintf("非法值 '%s'，允许: up, down", dir))
		}
	}

	if params, ok := data["params"].([]any); ok {
		for i, p := range params {
			if param, ok := p.(map[string]any); ok {
				v.validateParam(param, fmt.Sprintf("%s.params[%d]", path, i))
			} else {
				v.error(fmt.Sprintf("%s.params[%d]", path, i), "应为 object 类型")
			}
		}
	}
}

func (v *modelValidator) validateAction(data map[string]any, path string) {
	v.validateString(path+".identifier", data["identifier"], true)
	v.validateString(path+".name", data["name"], true)

	if dir, ok := data["dir"].(string); ok {
		if !validDirs[dir] {
			v.error(path+".dir", fmt.Sprintf("非法值 '%s'，允许: up, down", dir))
		}
	}

	if input, ok := data["input"].([]any); ok {
		for i, p := range input {
			if param, ok := p.(map[string]any); ok {
				v.validateParam(param, fmt.Sprintf("%s.input[%d]", path, i))
			} else {
				v.error(fmt.Sprintf("%s.input[%d]", path, i), "应为 object 类型")
			}
		}
	}

	if output, ok := data["output"].([]any); ok {
		for i, p := range output {
			if param, ok := p.(map[string]any); ok {
				v.validateParam(param, fmt.Sprintf("%s.output[%d]", path, i))
			} else {
				v.error(fmt.Sprintf("%s.output[%d]", path, i), "应为 object 类型")
			}
		}
	}
}

func (v *modelValidator) validateParam(data map[string]any, path string) {
	v.validateString(path+".identifier", data["identifier"], true)
	v.validateString(path+".name", data["name"], false)

	if define, ok := data["define"].(map[string]any); ok {
		v.validateDefine(define, path+".define")
	}
}

func (v *modelValidator) validateDefine(data map[string]any, path string) {
	dataType, ok := data["type"].(string)
	if !ok || dataType == "" {
		v.error(path+".type", "必填字段缺失")
		return
	}
	if !validDataTypes[dataType] {
		v.error(path+".type", fmt.Sprintf("非法值 '%s'，允许: bool, int, float, string, struct, enum, array, timestamp, matrix", dataType))
	}

	if dataType == "enum" {
		if mapping, ok := data["mapping"].(map[string]any); !ok || len(mapping) == 0 {
			v.warn(path+".mapping", "enum 类型建议定义 mapping")
		}
	}
}

func (v *modelValidator) validateString(path string, value any, required bool) {
	if value == nil {
		if required {
			v.error(path, "必填字段缺失")
		}
		return
	}
	if _, ok := value.(string); !ok {
		v.error(path, fmt.Sprintf("应为 string 类型，实际为 %T", value))
	}
}
