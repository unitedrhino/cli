// view_validate_test.go — 大屏画布内容校验（validateScreenContent）单元测试
package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	viewdata "gitee.com/unitedrhino/cli/cmd/view/data"
)

// loadValidScreenContent 读取 testdata 中的合法画布样例作为各反例的变异基准
func loadValidScreenContent(t *testing.T) map[string]any {
	t.Helper()
	raw, err := os.ReadFile("testdata/view_screen_valid.json")
	if err != nil {
		t.Fatalf("read valid fixture: %v", err)
	}
	var content map[string]any
	if err := json.Unmarshal(raw, &content); err != nil {
		t.Fatalf("parse valid fixture: %v", err)
	}
	return content
}

// TestPushRejectsMissingImagePackage 验证缺少图片模块路径时，推送在发起请求前被阻断。
func TestPushRejectsMissingImagePackage(t *testing.T) {
	content := loadValidScreenContent(t)
	comp := firstComponent(t, content)
	comp["chartKey"], comp["conKey"] = "VImage", "VCImage"
	comp["chartConfig"] = map[string]any{"key": "Image", "chartKey": "VImage", "conKey": "VCImage", "category": "Mores"}
	raw, err := json.Marshal(content)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "canvas.json")
	if err := os.WriteFile(path, raw, 0600); err != nil {
		t.Fatal(err)
	}
	var out, stderr bytes.Buffer
	if code := runViewScreenPush(context.Background(), []string{"-f", path, "--id", "123"}, &out, &stderr); code != 1 || !strings.Contains(stderr.String(), "chartConfig.package") {
		t.Fatalf("应拒绝缺失 package 的画布，code=%d stderr=%s", code, stderr.String())
	}
}

// TestImageRegistrationInGroups 验证图片模块路径在分组内同样会被校验，修正后可以通过。
func TestImageRegistrationInGroups(t *testing.T) {
	content := loadValidScreenContent(t)
	image := firstComponent(t, content)
	image["chartKey"], image["conKey"] = "VImage", "VCImage"
	image["chartConfig"] = map[string]any{"key": "Image", "chartKey": "VImage", "conKey": "VCImage", "category": "Mores"}
	image["option"] = map[string]any{"dataset": "https://example.com/image.png"}
	content["componentList"] = []any{map[string]any{"id": "group", "isGroup": true, "groupList": []any{image}}}
	if !hasErrorField(errorFields(validateScreenContent(content)), "chartConfig.package") {
		t.Fatal("分组内缺少图片 package 时应报错")
	}
	image["chartConfig"].(map[string]any)["package"] = "Informations"
	if errs := errorFields(validateScreenContent(content)); len(errs) != 0 {
		t.Fatalf("图片配置修正后应通过：%v", errs)
	}
}

// mutateScreenContent 深拷贝基准 content 并应用变异函数
func mutateScreenContent(t *testing.T, mutate func(content map[string]any)) map[string]any {
	t.Helper()
	base := loadValidScreenContent(t)
	raw, err := json.Marshal(base)
	if err != nil {
		t.Fatalf("marshal base: %v", err)
	}
	var content map[string]any
	if err := json.Unmarshal(raw, &content); err != nil {
		t.Fatalf("deep copy base: %v", err)
	}
	mutate(content)
	return content
}

// firstComponent 取 componentList 的第一个组件（静态 ECharts 组件）
func firstComponent(t *testing.T, content map[string]any) map[string]any {
	t.Helper()
	list, ok := content["componentList"].([]any)
	if !ok || len(list) == 0 {
		t.Fatalf("fixture has no componentList")
	}
	comp, ok := list[0].(map[string]any)
	if !ok {
		t.Fatalf("first component is not object")
	}
	return comp
}

// iotComponent 取 componentList 的第二个组件（IoT 绑定组件）
func iotComponent(t *testing.T, content map[string]any) map[string]any {
	t.Helper()
	list, ok := content["componentList"].([]any)
	if !ok || len(list) < 2 {
		t.Fatalf("fixture has no IoT component")
	}
	comp, ok := list[1].(map[string]any)
	if !ok {
		t.Fatalf("IoT component is not object")
	}
	return comp
}

// errorFields 收集校验结果中全部 error 级问题的字段路径
func errorFields(issues []viewIssue) []string {
	var fields []string
	for _, issue := range issues {
		if issue.Level == "error" {
			fields = append(fields, issue.Field)
		}
	}
	return fields
}

// hasErrorField 判断 error 字段列表中是否存在包含指定子串的条目
func hasErrorField(fields []string, substr string) bool {
	for _, f := range fields {
		if strings.Contains(f, substr) {
			return true
		}
	}
	return false
}

// TestValidateScreenContent_Valid 正例：完整合法 content 不应产生任何 error
func TestValidateScreenContent_Valid(t *testing.T) {
	content := loadValidScreenContent(t)
	issues := validateScreenContent(content)
	if errs := errorFields(issues); len(errs) > 0 {
		t.Fatalf("valid fixture should have no error, got: %v (all issues: %+v)", errs, issues)
	}
}

// TestValidateScreenContent_Errors 反例：逐项校验非法场景产生对应 error
func TestValidateScreenContent_Errors(t *testing.T) {
	cases := []struct {
		name      string
		mutate    func(content map[string]any)
		wantField string
	}{
		{
			name: "图片组件缺少 package",
			mutate: func(content map[string]any) {
				comp := firstComponent(t, content)
				comp["chartKey"], comp["conKey"] = "VImage", "VCImage"
				comp["chartConfig"] = map[string]any{"key": "Image", "chartKey": "VImage", "conKey": "VCImage", "category": "Mores"}
			},
			wantField: "chartConfig.package",
		},
		{
			name: "组件分类错误",
			mutate: func(content map[string]any) {
				firstComponent(t, content)["chartConfig"].(map[string]any)["category"] = "Charts"
			},
			wantField: "chartConfig.category",
		},
		{
			name: "仅顶层有注册键",
			mutate: func(content map[string]any) {
				delete(firstComponent(t, content)["chartConfig"].(map[string]any), "chartKey")
			},
			wantField: "chartKey",
		},
		{
			name: "分组缺少子组件列表",
			mutate: func(content map[string]any) {
				content["componentList"] = []any{map[string]any{"id": "group", "isGroup": true}}
			},
			wantField: "groupList",
		},
		{
			name: "无效重定向路径",
			mutate: func(content map[string]any) {
				firstComponent(t, content)["chartConfig"].(map[string]any)["redirectComponent"] = "wrong/Image"
			},
			wantField: "chartConfig.redirectComponent",
		},
		{
			name: "非法组件 key",
			mutate: func(content map[string]any) {
				comp := firstComponent(t, content)
				comp["chartConfig"].(map[string]any)["key"] = "NotExistChart"
			},
			wantField: "chartConfig.key",
		},
		{
			name: "chartKey 前缀不一致",
			mutate: func(content map[string]any) {
				firstComponent(t, content)["chartKey"] = "XBarCommon"
			},
			wantField: "chartKey",
		},
		{
			name: "conKey 前缀不一致",
			mutate: func(content map[string]any) {
				firstComponent(t, content)["conKey"] = "VBarCommon"
			},
			wantField: "conKey",
		},
		{
			name: "组件 id 重复",
			mutate: func(content map[string]any) {
				iotComponent(t, content)["id"] = "comp-bar-001"
			},
			wantField: "id",
		},
		{
			name: "attr.w 为 0",
			mutate: func(content map[string]any) {
				firstComponent(t, content)["attr"].(map[string]any)["w"] = float64(0)
			},
			wantField: "attr.w",
		},
		{
			name: "attr.h 为负数",
			mutate: func(content map[string]any) {
				firstComponent(t, content)["attr"].(map[string]any)["h"] = float64(-100)
			},
			wantField: "attr.h",
		},
		{
			name: "requestDataType=5 缺 requestIoTDeviceConfig",
			mutate: func(content map[string]any) {
				delete(iotComponent(t, content)["request"].(map[string]any), "requestIoTDeviceConfig")
			},
			wantField: "request.requestIoTDeviceConfig",
		},
		{
			name: "queryType 不在组件支持矩阵（LineCommon 配 latest）",
			mutate: func(content map[string]any) {
				iot := iotComponent(t, content)["request"].(map[string]any)["requestIoTDeviceConfig"].(map[string]any)
				iot["queryType"] = "latest"
			},
			wantField: "queryType",
		},
		{
			name: "productID 为空",
			mutate: func(content map[string]any) {
				iot := iotComponent(t, content)["request"].(map[string]any)["requestIoTDeviceConfig"].(map[string]any)
				iot["productID"] = ""
			},
			wantField: "productID",
		},
		{
			name: "deviceNames 为空数组",
			mutate: func(content map[string]any) {
				iot := iotComponent(t, content)["request"].(map[string]any)["requestIoTDeviceConfig"].(map[string]any)
				iot["deviceNames"] = []any{}
			},
			wantField: "deviceNames",
		},
		{
			name: "dataIDs 为空数组",
			mutate: func(content map[string]any) {
				iot := iotComponent(t, content)["request"].(map[string]any)["requestIoTDeviceConfig"].(map[string]any)
				iot["dataIDs"] = []any{}
			},
			wantField: "dataIDs",
		},
		{
			name: "画布宽度非正数",
			mutate: func(content map[string]any) {
				content["editCanvasConfig"].(map[string]any)["width"] = float64(0)
			},
			wantField: "editCanvasConfig.width",
		},
		{
			name: "styles 缺失（animations 渲染崩溃）",
			mutate: func(content map[string]any) {
				delete(firstComponent(t, content), "styles")
			},
			wantField: "styles",
		},
		{
			name: "styles.animations 不是数组",
			mutate: func(content map[string]any) {
				firstComponent(t, content)["styles"].(map[string]any)["animations"] = "not-array"
			},
			wantField: "styles.animations",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			content := mutateScreenContent(t, tc.mutate)
			issues := validateScreenContent(content)
			fields := errorFields(issues)
			if !hasErrorField(fields, tc.wantField) {
				t.Fatalf("expect error on field %q, got error fields: %v (all issues: %+v)", tc.wantField, fields, issues)
			}
		})
	}
}

// TestValidateScreenContent_InvalidFixture 校验 testdata 中的非法样例确实产生 error
func TestValidateScreenContent_InvalidFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/view_screen_invalid.json")
	if err != nil {
		t.Fatalf("read invalid fixture: %v", err)
	}
	var content map[string]any
	if err := json.Unmarshal(raw, &content); err != nil {
		t.Fatalf("parse invalid fixture: %v", err)
	}
	issues := validateScreenContent(content)
	fields := errorFields(issues)
	if len(fields) == 0 {
		t.Fatalf("invalid fixture should produce errors")
	}
	if !hasErrorField(fields, "queryType") {
		t.Fatalf("invalid fixture should fail on queryType, got: %v", fields)
	}
}

// TestValidateScreenContent_WarningsNotErrors 校验 status/events 缺失只产生 warning
func TestValidateScreenContent_WarningsNotErrors(t *testing.T) {
	content := mutateScreenContent(t, func(content map[string]any) {
		comp := firstComponent(t, content)
		delete(comp, "status")
		delete(comp, "events")
	})
	issues := validateScreenContent(content)
	if errs := errorFields(issues); len(errs) > 0 {
		t.Fatalf("missing status/events should only warn, got errors: %v", errs)
	}
	if len(issues) == 0 {
		t.Fatalf("expect warnings for missing status/events")
	}
}

// TestValidateScreenContent_EchartsOptionWarns 校验 ECharts 组件 option 过简仅告警不阻断
func TestValidateScreenContent_EchartsOptionWarns(t *testing.T) {
	content := mutateScreenContent(t, func(content map[string]any) {
		firstComponent(t, content)["option"] = map[string]any{"dataset": map[string]any{}}
	})
	issues := validateScreenContent(content)
	if errs := errorFields(issues); len(errs) > 0 {
		t.Fatalf("ECharts option 过简应只告警，got errors: %v", errs)
	}
	warned := false
	for _, issue := range issues {
		if issue.Level == "warning" && strings.Contains(issue.Field, "option") {
			warned = true
			break
		}
	}
	if !warned {
		t.Fatalf("expect warning on ECharts option, got issues: %+v", issues)
	}
}

// TestValidateScreenContent_RealCanvasShape 校验真实画布形态（chartKey/conKey 在
// chartConfig 内、deviceStatus 查询不带 deviceNames/dataIDs）不产生误报
func TestValidateScreenContent_RealCanvasShape(t *testing.T) {
	content := mutateScreenContent(t, func(content map[string]any) {
		// 模拟真实画布：顶层无 chartKey/conKey（仅存在于 chartConfig）
		for _, comp := range content["componentList"].([]any) {
			c := comp.(map[string]any)
			delete(c, "chartKey")
			delete(c, "conKey")
		}
		// 第二个组件换成支持 deviceStatus 的 PieCircle，模拟设备在线状态统计
		comp := iotComponent(t, content)
		cc := comp["chartConfig"].(map[string]any)
		cc["key"] = "PieCircle"
		cc["chartKey"] = "VPieCircle"
		cc["conKey"] = "VCPieCircle"
		cc["category"] = "Pies"
		comp["key"] = "PieCircle"
		comp["request"] = map[string]any{
			"requestDataType": 5,
			"requestIoTDeviceConfig": map[string]any{
				"queryType":   "deviceStatus",
				"mode":        "latest",
				"productID":   "14",
				"deviceNames": []any{},
				"dataIDs":     []any{},
			},
		}
	})
	issues := validateScreenContent(content)
	if errs := errorFields(issues); len(errs) > 0 {
		t.Fatalf("real canvas shape should pass, got errors: %v (all issues: %+v)", errs, issues)
	}
}

// TestValidateScreenContent_UnsupportedIoTComponent 校验不支持 IoT 的组件绑定 IoT 报 error
func TestValidateScreenContent_UnsupportedIoTComponent(t *testing.T) {
	// 动态选取一个 IoTQueryTypes 为空（不支持 IoT）的组件，
	// 使本用例对占位版/完整版 components.json 数据均健壮
	snapshot, err := viewdata.Load()
	if err != nil {
		t.Fatalf("load components snapshot: %v", err)
	}
	unsupportedKey := ""
	for _, c := range snapshot.Components {
		if len(c.IoTQueryTypes) == 0 {
			unsupportedKey = c.Key
			break
		}
	}
	if unsupportedKey == "" {
		t.Skip("components.json 中所有组件均支持 IoT，无可测对象")
	}

	content := mutateScreenContent(t, func(content map[string]any) {
		comp := firstComponent(t, content)
		comp["chartConfig"].(map[string]any)["key"] = unsupportedKey
		comp["key"] = unsupportedKey
		comp["chartKey"] = "V" + unsupportedKey
		comp["conKey"] = "VC" + unsupportedKey
		comp["request"] = map[string]any{
			"requestDataType": float64(5),
			"requestIoTDeviceConfig": map[string]any{
				"queryType":   "property",
				"productID":   "p1",
				"deviceNames": []any{"d1"},
				"dataIDs":     []any{"Temperature"},
			},
		}
	})
	issues := validateScreenContent(content)
	if fields := errorFields(issues); !hasErrorField(fields, "queryType") {
		t.Fatalf("expect error for IoT binding on unsupported component %s, got: %v", unsupportedKey, fields)
	}
}
