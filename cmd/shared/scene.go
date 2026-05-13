package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// runScene 场景联动相关命令
//   scene validate <file>     校验场景联动 JSON
//   scene template auto       生成自动触发模板
//   scene template manual     生成手动触发模板
func runScene(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: scene validate <file>       # 校验场景联动 JSON")
		fmt.Fprintln(stderr, "       scene template auto         # 生成自动触发模板")
		fmt.Fprintln(stderr, "       scene template manual       # 生成手动触发模板")
		return 2
	}

	subCmd := args[0]
	subArgs := args[1:]

	switch subCmd {
	case "validate":
		return runSceneValidate(subArgs, stdout, stderr)
	case "template":
		return runSceneTemplate(subArgs, stdout, stderr)
	default:
		fmt.Fprintf(stderr, "未知的 scene 子命令: %s\n", subCmd)
		return 2
	}
}

func runSceneValidate(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: scene validate <file>")
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

	v := newSceneValidator()
	v.validateInfo(data, "")

	if len(v.errors) > 0 {
		fmt.Fprintf(stdout, "❌ 校验失败，共 %d 个错误:\n", len(v.errors))
		for _, e := range v.errors {
			fmt.Fprintf(stdout, "  %s\n", e)
		}
	} else {
		fmt.Fprintln(stdout, "✅ 校验通过")
	}
	if len(v.warnings) > 0 {
		fmt.Fprintf(stdout, "⚠️  共 %d 个警告:\n", len(v.warnings))
		for _, w := range v.warnings {
			fmt.Fprintf(stdout, "  %s\n", w)
		}
	}

	if len(v.errors) > 0 {
		return 1
	}
	return 0
}

func runSceneTemplate(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: scene template [auto|manual]")
		return 2
	}

	templateType := args[0]
	var tmpl map[string]any
	switch templateType {
	case "auto":
		tmpl = sceneAutoTemplate()
	case "manual":
		tmpl = sceneManualTemplate()
	default:
		fmt.Fprintf(stderr, "错误: 模板类型应为 auto 或 manual，实际: %s\n", templateType)
		return 2
	}

	out, _ := json.MarshalIndent(tmpl, "", "  ")
	fmt.Fprintln(stdout, string(out))
	return 0
}

// sceneAutoTemplate 自动触发模板
func sceneAutoTemplate() map[string]any {
	return map[string]any{
		"type":       "auto",
		"deviceMode": "multi",
		"name":       "",
		"desc":       "",
		"status":     1,
		"if": map[string]any{
			"triggers": []map[string]any{
				{
					"type":  "device",
					"order": 1,
					"device": map[string]any{
						"productID":  "",
						"selectType": "fixed",
						"deviceName": "",
						"type":       "propertyReport",
						"dataID":     "",
						"termType":   "gt",
						"values":     []string{""},
					},
				},
			},
		},
		"when": map[string]any{
			"validRanges":   []any{},
			"invalidRanges": []any{},
			"conditions": map[string]any{
				"type":  "and",
				"terms": []any{},
			},
		},
		"then": map[string]any{
			"actions": []map[string]any{
				{
					"order": 1,
					"type":  "device",
					"device": map[string]any{
						"productID":  "",
						"selectType": "fixed",
						"deviceName": "",
						"type":       "propertyControl",
						"dataID":     "",
						"value":      "",
					},
				},
			},
		},
	}
}

// sceneManualTemplate 手动触发模板
func sceneManualTemplate() map[string]any {
	return map[string]any{
		"type":       "manual",
		"deviceMode": "multi",
		"name":       "",
		"desc":       "",
		"status":     1,
		"when": map[string]any{
			"validRanges":   []any{},
			"invalidRanges": []any{},
			"conditions": map[string]any{
				"type":  "and",
				"terms": []any{},
			},
		},
		"then": map[string]any{
			"actions": []map[string]any{
				{
					"order": 1,
					"type":  "device",
					"device": map[string]any{
						"productID":  "",
						"selectType": "fixed",
						"deviceName": "",
						"type":       "propertyControl",
						"dataID":     "",
						"value":      "",
					},
				},
			},
		},
	}
}

// ---------- 校验器 ----------

type sceneValidator struct {
	errors   []string
	warnings []string
}

func newSceneValidator() *sceneValidator {
	return &sceneValidator{}
}

func (v *sceneValidator) error(path, msg string) {
	v.errors = append(v.errors, fmt.Sprintf("[%s] %s", path, msg))
}

func (v *sceneValidator) warn(path, msg string) {
	v.warnings = append(v.warnings, fmt.Sprintf("[%s] %s", path, msg))
}

// 枚举集合
var (
	sceneTypes          = setOf("auto", "manual")
	deviceModes         = setOf("single", "multi")
	triggerTypes        = setOf("device", "timer", "weather")
	selectTypes         = setOf("all", "fixed", "area", "areaWithChildren", "group")
	triggerDeviceModes  = setOf("edge", "level")
	triggerDeviceTypes  = setOf("connected", "disConnected", "propertyReport", "eventReport")
	execTypes           = setOf("at", "sunRises", "sunSet", "loop")
	repeatTypes         = setOf("once", "week", "mount", "allDay", "customRange")
	weatherTypes        = setOf("temp", "humidity")
	whenRangeTypes      = setOf("date", "time")
	dateRangeTypes      = setOf("allDay", "workday", "weekend", "holiday", "customRange", "customWeek")
	timeRangeTypes      = setOf("allDay", "light", "night", "customRange")
	termCondTypes       = setOf("and", "or")
	termColumnTypes     = setOf("property", "event", "time", "weather")
	actionTypes         = setOf("device", "delay", "notify", "alarm", "scene")
	actionDeviceTypes   = setOf("propertyControl", "action")
	notifyTypes         = setOf("sms", "email", "dingTalk", "wxMini", "message", "phoneCall", "wxEWebhook")
	notifyCodes         = setOf("ruleScene", "ruleDeviceAlarm")
	notifyUserTypes     = setOf("account", "userID", "deviceOwner", "deviceProjectAdmin", "deviceAreaAdmin", "deviceProjectAll", "deviceAreaAll")
	alarmModes          = setOf("trigger", "relieve")
	cmpTypes            = setOf("eq", "not", "btw", "gt", "gte", "lt", "lte", "in", "all")
	stateKeepTypes      = setOf("duration", "repeat")
	cmpValuesCount      = map[string]int{"eq": 1, "not": 1, "btw": 2, "gt": 1, "gte": 1, "lt": 1, "lte": 1, "in": -1, "all": 0}
)

func setOf(vals ...string) map[string]struct{} {
	s := make(map[string]struct{})
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (v *sceneValidator) validateInfo(data map[string]any, path string) {
	if data == nil {
		v.error(path, "场景联动 JSON 应为 object 类型")
		return
	}

	v.validateStringEnum(path+".type", data["type"], sceneTypes, true)
	v.validateStringEnum(path+".deviceMode", data["deviceMode"], deviceModes, true)
	v.validateString(path+".name", data["name"], true)

	sceneType, _ := data["type"].(string)
	deviceMode, _ := data["deviceMode"].(string)

	if deviceMode == "single" {
		v.validateString(path+".productID", data["productID"], true)
		v.validateString(path+".deviceName", data["deviceName"], true)
	}

	if status, ok := data["status"]; ok && status != nil {
		if s, ok := status.(float64); !ok || (s != 1 && s != 2) {
			v.error(path+".status", fmt.Sprintf("非法值 %v，允许: 1, 2", status))
		}
	}

	if sceneType == "auto" {
		if ifData, ok := data["if"].(map[string]any); ok {
			v.validateIf(ifData, path+".if")
		} else {
			v.error(path+".if", "自动触发类型必填")
		}
	}

	if whenData, ok := data["when"].(map[string]any); ok {
		v.validateWhen(whenData, path+".when")
	}

	if thenData, ok := data["then"].(map[string]any); ok {
		v.validateThen(thenData, path+".then")
	} else {
		v.error(path+".then", "执行动作必填")
	}
}

func (v *sceneValidator) validateIf(data map[string]any, path string) {
	triggers, ok := data["triggers"].([]any)
	if !ok {
		v.error(path+".triggers", "应为 array 类型")
		return
	}
	for i, t := range triggers {
		if trigger, ok := t.(map[string]any); ok {
			v.validateTrigger(trigger, fmt.Sprintf("%s.triggers[%d]", path, i))
		} else {
			v.error(fmt.Sprintf("%s.triggers[%d]", path, i), "应为 object 类型")
		}
	}
}

func (v *sceneValidator) validateTrigger(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], triggerTypes, true)
	triggerType, _ := data["type"].(string)

	switch triggerType {
	case "device":
		if device, ok := data["device"].(map[string]any); ok {
			v.validateTriggerDevice(device, path+".device")
		} else {
			v.error(path+".device", "设备触发类型必填")
		}
	case "timer":
		if timer, ok := data["timer"].(map[string]any); ok {
			v.validateTriggerTimer(timer, path+".timer")
		} else {
			v.error(path+".timer", "定时触发类型必填")
		}
	case "weather":
		if weather, ok := data["weather"].(map[string]any); ok {
			v.validateTriggerWeather(weather, path+".weather")
		} else {
			v.error(path+".weather", "天气触发类型必填")
		}
	}
}

func (v *sceneValidator) validateTriggerDevice(data map[string]any, path string) {
	v.validateString(path+".productID", data["productID"], true)
	v.validateStringEnum(path+".selectType", data["selectType"], selectTypes, true)
	v.validateStringEnum(path+".mode", data["mode"], triggerDeviceModes, false)
	v.validateStringEnum(path+".type", data["type"], triggerDeviceTypes, true)

	selectType, _ := data["selectType"].(string)
	if selectType == "fixed" {
		v.validateString(path+".deviceName", data["deviceName"], true)
	}

	deviceType, _ := data["type"].(string)
	if deviceType == "propertyReport" || deviceType == "eventReport" {
		v.validateString(path+".dataID", data["dataID"], true)
	}

	if sk, ok := data["stateKeep"].(map[string]any); ok {
		v.validateStateKeep(sk, path+".stateKeep")
	}

	if _, hasTermType := data["termType"]; hasTermType {
		v.validateCompare(data, path)
	}
	if terms, ok := data["terms"].([]any); ok {
		v.validateCmps(terms, path+".terms")
	}
}

func (v *sceneValidator) validateTriggerTimer(data map[string]any, path string) {
	v.validateStringEnum(path+".execType", data["execType"], execTypes, true)
	v.validateStringEnum(path+".repeatType", data["repeatType"], repeatTypes, false)

	execType, _ := data["execType"].(string)
	if execType == "at" {
		if execAt, ok := data["execAt"].(float64); ok {
			if execAt < 0 || execAt > 86400 {
				v.error(path+".execAt", fmt.Sprintf("应在 0-86400 之间，实际: %v", execAt))
			}
		}
	} else if execType == "sunRises" || execType == "sunSet" {
		execAdd := 0.0
		if v, ok := data["execAdd"].(float64); ok {
			execAdd = v
		}
		if execAdd < -10800 || execAdd > 10800 {
			v.error(path+".execAdd", fmt.Sprintf("应在 -10800~10800 之间，实际: %v", execAdd))
		}
	}
}

func (v *sceneValidator) validateTriggerWeather(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], weatherTypes, true)
	v.validateCompare(data, path)
}

func (v *sceneValidator) validateStateKeep(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], stateKeepTypes, true)
	v.validateInt(path+".value", data["value"], true)
}

func (v *sceneValidator) validateWhen(data map[string]any, path string) {
	if ranges, ok := data["validRanges"].([]any); ok {
		v.validateWhenRanges(ranges, path+".validRanges")
	}
	if ranges, ok := data["invalidRanges"].([]any); ok {
		v.validateWhenRanges(ranges, path+".invalidRanges")
	}
	if conds, ok := data["conditions"].(map[string]any); ok {
		v.validateConditions(conds, path+".conditions")
	}
}

func (v *sceneValidator) validateWhenRanges(data []any, path string) {
	for i, r := range data {
		if rang, ok := r.(map[string]any); ok {
			v.validateWhenRange(rang, fmt.Sprintf("%s[%d]", path, i))
		} else {
			v.error(fmt.Sprintf("%s[%d]", path, i), "应为 object 类型")
		}
	}
}

func (v *sceneValidator) validateWhenRange(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], whenRangeTypes, true)
	rangeType, _ := data["type"].(string)
	if rangeType == "date" {
		if dr, ok := data["dateRange"].(map[string]any); ok {
			v.validateDateRange(dr, path+".dateRange")
		}
	} else if rangeType == "time" {
		if tr, ok := data["timeRange"].(map[string]any); ok {
			v.validateTimeRange(tr, path+".timeRange")
		}
	}
}

func (v *sceneValidator) validateDateRange(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], dateRangeTypes, true)
	dateType, _ := data["type"].(string)
	if dateType == "customRange" {
		v.validateString(path+".startDate", data["startDate"], true)
		v.validateString(path+".endDate", data["endDate"], true)
	}
}

func (v *sceneValidator) validateTimeRange(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], timeRangeTypes, true)
	timeType, _ := data["type"].(string)
	if timeType == "customRange" {
		if start, ok := data["startTime"].(float64); ok {
			if start < 0 || start > 86400 {
				v.error(path+".startTime", fmt.Sprintf("应在 0-86400 之间，实际: %v", start))
			}
		}
		if end, ok := data["endTime"].(float64); ok {
			if end < 0 || end > 86400 {
				v.error(path+".endTime", fmt.Sprintf("应在 0-86400 之间，实际: %v", end))
			}
		}
	}
}

func (v *sceneValidator) validateConditions(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], termCondTypes, false)
	if terms, ok := data["terms"].([]any); ok {
		for i, t := range terms {
			if term, ok := t.(map[string]any); ok {
				v.validateTerm(term, fmt.Sprintf("%s.terms[%d]", path, i))
			} else {
				v.error(fmt.Sprintf("%s.terms[%d]", path, i), "应为 object 类型")
			}
		}
	}
}

func (v *sceneValidator) validateTerm(data map[string]any, path string) {
	v.validateStringEnum(path+".columnType", data["columnType"], termColumnTypes, true)
	columnType, _ := data["columnType"].(string)
	switch columnType {
	case "property":
		if p, ok := data["property"].(map[string]any); ok {
			v.validateTermProperty(p, path+".property")
		}
	case "weather":
		if w, ok := data["weather"].(map[string]any); ok {
			v.validateTriggerWeather(w, path+".weather")
		}
	case "time":
		if t, ok := data["time"].(map[string]any); ok {
			v.validateTermTime(t, path+".time")
		}
	}
}

func (v *sceneValidator) validateTermProperty(data map[string]any, path string) {
	v.validateString(path+".productID", data["productID"], true)
	v.validateString(path+".deviceName", data["deviceName"], true)
	if _, hasTermType := data["termType"]; hasTermType {
		v.validateCompare(data, path)
	}
	if terms, ok := data["terms"].([]any); ok {
		v.validateCmps(terms, path+".terms")
	}
}

func (v *sceneValidator) validateTermTime(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], setOf("sys", "sunRises", "sunSet"), true)
	v.validateCompare(data, path)
}

func (v *sceneValidator) validateThen(data map[string]any, path string) {
	actions, ok := data["actions"].([]any)
	if !ok {
		v.error(path+".actions", "应为 array 类型")
		return
	}
	for i, a := range actions {
		if action, ok := a.(map[string]any); ok {
			v.validateAction(action, fmt.Sprintf("%s.actions[%d]", path, i))
		} else {
			v.error(fmt.Sprintf("%s.actions[%d]", path, i), "应为 object 类型")
		}
	}
}

func (v *sceneValidator) validateAction(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], actionTypes, true)
	actionType, _ := data["type"].(string)

	switch actionType {
	case "delay":
		v.validateInt(path+".delay", data["delay"], true)
	case "device":
		if device, ok := data["device"].(map[string]any); ok {
			v.validateActionDevice(device, path+".device")
		} else {
			v.error(path+".device", "设备动作类型必填")
		}
	case "notify":
		if notify, ok := data["notify"].(map[string]any); ok {
			v.validateActionNotify(notify, path+".notify")
		} else {
			v.error(path+".notify", "通知动作类型必填")
		}
	case "alarm":
		if alarm, ok := data["alarm"].(map[string]any); ok {
			v.validateActionAlarm(alarm, path+".alarm")
		} else {
			v.error(path+".alarm", "告警动作类型必填")
		}
	case "scene":
		if scene, ok := data["scene"].(map[string]any); ok {
			v.validateActionScene(scene, path+".scene")
		} else {
			v.error(path+".scene", "场景动作类型必填")
		}
	}
}

func (v *sceneValidator) validateActionDevice(data map[string]any, path string) {
	v.validateString(path+".productID", data["productID"], true)
	v.validateStringEnum(path+".selectType", data["selectType"], selectTypes, true)
	v.validateStringEnum(path+".type", data["type"], actionDeviceTypes, true)
	selectType, _ := data["selectType"].(string)
	if selectType == "fixed" {
		v.validateString(path+".deviceName", data["deviceName"], true)
	}

	hasDataID := data["dataID"] != nil && data["dataID"] != ""
	hasValues := false
	if values, ok := data["values"].([]any); ok && len(values) > 0 {
		hasValues = true
		for i, val := range values {
			if vv, ok := val.(map[string]any); ok {
				v.validateDeviceValue(vv, fmt.Sprintf("%s.values[%d]", path, i))
			} else {
				v.error(fmt.Sprintf("%s.values[%d]", path, i), "应为 object 类型")
			}
		}
	}
	if !hasDataID && !hasValues {
		v.error(path, "设备动作需填写 dataID 或 values")
	}
}

func (v *sceneValidator) validateDeviceValue(data map[string]any, path string) {
	v.validateString(path+".dataID", data["dataID"], true)
	v.validateString(path+".value", data["value"], true)
}

func (v *sceneValidator) validateActionNotify(data map[string]any, path string) {
	v.validateStringEnum(path+".type", data["type"], notifyTypes, true)
	v.validateStringEnum(path+".notifyCode", data["notifyCode"], notifyCodes, true)
	v.validateStringEnum(path+".userType", data["userType"], notifyUserTypes, true)
	userType, _ := data["userType"].(string)
	if userType == "account" {
		v.validateArray(path+".accounts", data["accounts"], true)
	} else if userType == "userID" {
		v.validateArray(path+".userIDs", data["userIDs"], true)
	}
}

func (v *sceneValidator) validateActionAlarm(data map[string]any, path string) {
	v.validateStringEnum(path+".mode", data["mode"], alarmModes, true)
}

func (v *sceneValidator) validateActionScene(data map[string]any, path string) {
	v.validateInt(path+".sceneID", data["sceneID"], true)
}

func (v *sceneValidator) validateCompare(data map[string]any, path string) {
	termType, ok := data["termType"].(string)
	if !ok || termType == "" {
		return
	}
	v.validateStringEnum(path+".termType", termType, cmpTypes, false)
	expected, known := cmpValuesCount[termType]
	if !known {
		return
	}
	values := data["values"]
	if values == nil {
		values = []any{}
	}
	var vals []any
	if vv, ok := values.([]any); ok {
		vals = vv
	}
	if expected == 0 {
		if len(vals) > 0 {
			v.warn(path+".values", fmt.Sprintf("termType='%s' 时 values 应为空或省略", termType))
		}
	} else if expected == -1 {
		if len(vals) == 0 {
			v.error(path+".values", fmt.Sprintf("termType='%s' 需要至少1个值", termType))
		}
	} else if len(vals) != expected {
		v.error(path+".values", fmt.Sprintf("termType='%s' 需要 %d 个值，实际: %d", termType, expected, len(vals)))
	}
}

func (v *sceneValidator) validateCmps(data []any, path string) {
	for i, c := range data {
		if cmp, ok := c.(map[string]any); ok {
			v.validateString(fmt.Sprintf("%s[%d].column", path, i), cmp["column"], true)
			v.validateCompare(cmp, fmt.Sprintf("%s[%d]", path, i))
		} else {
			v.error(fmt.Sprintf("%s[%d]", path, i), "应为 object 类型")
		}
	}
}

// ---------- 基础校验辅助 ----------

func (v *sceneValidator) validateString(path string, value any, required bool) {
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

func (v *sceneValidator) validateInt(path string, value any, required bool) {
	if value == nil {
		if required {
			v.error(path, "必填字段缺失")
		}
		return
	}
	if _, ok := value.(float64); !ok {
		v.error(path, fmt.Sprintf("应为 number 类型，实际为 %T", value))
	}
}

func (v *sceneValidator) validateArray(path string, value any, required bool) {
	if value == nil {
		if required {
			v.error(path, "必填字段缺失")
		}
		return
	}
	if _, ok := value.([]any); !ok {
		v.error(path, fmt.Sprintf("应为 array 类型，实际为 %T", value))
	}
}

func (v *sceneValidator) validateStringEnum(path string, value any, allowed map[string]struct{}, required bool) {
	if value == nil {
		if required {
			v.error(path, "必填字段缺失")
		}
		return
	}
	s, ok := value.(string)
	if !ok {
		v.error(path, fmt.Sprintf("应为 string 类型，实际为 %T", value))
		return
	}
	if _, ok := allowed[s]; !ok {
		var list []string
		for k := range allowed {
			list = append(list, k)
		}
		v.error(path, fmt.Sprintf("非法值 '%s'，允许: %v", s, list))
	}
}

func printSceneHelp(stdout io.Writer) {
	fmt.Fprintln(stdout, "场景联动命令:")
	fmt.Fprintln(stdout, "  scene validate <file>     校验场景联动 JSON 文件")
	fmt.Fprintln(stdout, "  scene template auto       生成自动触发场景模板")
	fmt.Fprintln(stdout, "  scene template manual     生成手动触发场景模板")
}
