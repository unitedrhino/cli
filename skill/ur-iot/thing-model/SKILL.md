---
name: thing-model-skill
description: "Use when generating, validating, or troubleshooting 联犀 IoT thing-model JSON definitions for properties, events, actions, and schema constraints."
---

# 物模型生成与校验 Skill

帮助用户根据设备需求描述生成符合联犀平台规范的物模型 JSON，并通过 Go 代码校验。

## 工作流程

1. **理解需求**：用户描述设备类型和功能（如"三相电表"、"水表"、"智能灯"）
2. **生成物模型 JSON**：根据需求生成符合 `schema.Model` 结构的 JSON
3. **校验**：使用校验工具验证 JSON 合法性
4. **输出**：返回校验通过的物模型 JSON

## Model 顶层结构

```json
{
  "version": "1.0",
  "properties": [...],
  "events": [...],
  "actions": [...],
  "profile": { "productID": "" }
}
```

至少包含 properties/events/actions 中的一个。所有 identifier 全局不可重复。

## Property（属性）

```json
{
  "identifier": "Temperature",
  "tag": 1,
  "name": "环境温度",
  "desc": "当前环境温度",
  "required": false,
  "extendConfig": "",
  "isShareAuthPerm": 0,
  "order": 0,
  "mode": "rw",
  "define": { "type": "float", "min": "-20", "max": "55", "step": "0.1", "unit": "℃" },
  "isUseShadow": false,
  "isCanSceneLinkage": 0,
  "funcGroup": 1,
  "userPerm": 3,
  "controlMode": 1,
  "recordMode": 1,
  "isPassword": 2
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| identifier | string | 标识符，大驼峰，1-128字符，字母数字下划线连字符，不能数字开头 |
| tag | int | 1:自定义 2:可选 3:必选(不可删除) |
| name | string | 功能名称，≤128字符 |
| desc | string | 描述，≤80字符 |
| mode | string | "r"(只读) 或 "rw"(可读可写，默认) |
| define | Define | 数据类型定义（见下方） |
| funcGroup | int | 1:普通功能 2:系统功能 |
| userPerm | int | 1:只读 3:可读可写 |
| controlMode | int | 1:群控+单控 2:仅单控 |
| recordMode | int | 1:全部记录 2:差异记录 3:不记录 |
| isPassword | int | 1:是密码 2:非密码 |

## Event（事件）

```json
{
  "identifier": "PowerAlarm",
  "tag": 1,
  "name": "电量告警",
  "desc": "电量低于阈值时告警",
  "type": "alert",
  "params": [
    { "identifier": "Voltage", "name": "电压", "define": { "type": "float", "min": "0", "max": "300", "step": "0.1", "unit": "V" } }
  ],
  "dir": "up"
}
```

| 字段 | 说明 |
|------|------|
| type | "info"(信息) / "alert"(告警) / "fault"(故障) |
| params | 事件参数，最多20个，identifier 不可重复 |
| dir | "up"(设备→云) / "down"(云→设备) |

## Action（行为）

```json
{
  "identifier": "OpenValve",
  "tag": 1,
  "name": "开阀",
  "desc": "远程开阀控制",
  "dir": "down",
  "input": [
    { "identifier": "Duration", "name": "持续时间", "define": { "type": "int", "min": "0", "max": "3600", "unit": "s" } }
  ],
  "output": [
    { "identifier": "Result", "name": "执行结果", "define": { "type": "bool", "mapping": {"0": "失败", "1": "成功"} } }
  ]
}
```

| 字段 | 说明 |
|------|------|
| dir | "down"(云端调用设备，默认) / "up"(设备调用云端) |
| input | 输入参数，最多20个 |
| output | 输出参数，最多20个，无返回参数时设备可不回复 |

## 数据类型定义（Define）

### bool 布尔型
```json
{ "type": "bool", "mapping": {"0": "关", "1": "开"} }
```
- mapping 必须恰好 2 个键："0" 和 "1"，值≤50字符
- 不填时自动补充 `{"0":"关","1":"开"}`

### int 整数型
```json
{ "type": "int", "min": "-20", "max": "55", "step": "1", "unit": "℃" }
```
- min 默认 -9999999999999，max 默认 9999999999999，step 默认 1
- unit ≤12字符，min/max/step 均为 string 类型

### float 浮点型
```json
{ "type": "float", "min": "0", "max": "16", "step": "0.001", "unit": "A" }
```
- 规则同 int，step 默认 0.001

### string 字符串
```json
{ "type": "string", "max": "2048" }
```
- max 为最大字节长度，默认 2048

### enum 枚举型
```json
{ "type": "enum", "mapping": {"0": "低档", "1": "中档", "2": "高档"} }
```
- 键必须为整数字符串，值≤50字符，mapping 不能为空

### timestamp 时间型
```json
{ "type": "timestamp" }
```
- 毫秒级 Unix 时间戳，无额外参数

### struct 结构体
```json
{
  "type": "struct",
  "specs": [
    { "identifier": "Red", "name": "红色", "dataType": { "type": "int", "min": "0", "max": "255" } },
    { "identifier": "Green", "name": "绿色", "dataType": { "type": "int", "min": "0", "max": "255" } }
  ]
}
```
- specs 最多 20 个字段，identifier 不可重复

### array 数组
```json
{ "type": "array", "max": "30", "arrayInfo": { "type": "int", "min": "0", "max": "1" } }
```
- max 为数组最大长度（≤1024），必填
- arrayInfo 为元素类型定义，必填

## 命名规范

- 标识符使用**大驼峰命名法**（PascalCase），不带标点
- 必须以字母开头，只包含字母、数字、下划线、连字符
- 同一功能多例化可加下划线编号：`Switch_1`, `Switch_2`
- **属性**用名词或断言：`Temperature`, `IsClosed`, `PostiveFlux`
- **行为**用动词开头：`GetColor`, `OpenValve`, `ToggleSwitch`
- **事件**描述事件含义：`PowerAlarm`, `Error`, `Warning`
- 英文标识符≤50字符，中文名称≤20字符

## 校验规则速查

| 规则 | 约束 |
|------|------|
| identifier 长度 | 1-128字符 |
| identifier 字符集 | A-Z, a-z, 0-9, _, - |
| identifier 首字符 | 不能是数字 |
| identifier 唯一性 | 全局（properties+events+actions）不可重复 |
| name 长度 | ≤128字符（中文算1个） |
| desc 长度 | ≤80字符 |
| struct specs 数量 | ≤20个 |
| params 数量 | ≤20个 |
| bool mapping | 恰好2个键："0"和"1" |
| enum mapping 键 | 必须为整数字符串 |
| enum mapping 值长度 | ≤50字符 |
| int/float unit 长度 | ≤12字符 |
| array max | ≤1024，必须>0 |
| array arrayInfo | 必填 |

## 校验方法

### 方式 1：CLI 校验（推荐）

```bash
ur model validate model.json
```

校验内容：
- identifier 命名规范和唯一性
- 必填字段完整性
- 枚举值合法性（mode、type、dir 等）
- Define 结构合法性

### 方式 2：Go 代码校验

```bash
cd backend/things/share/domain/schema
go test -v -run TestValidateThingModel ./...
```

校验测试文件：`backend/things/share/domain/schema/validate_thingmodel_test.go`

```go
func TestValidateThingModel(t *testing.T) {
    jsonStr := `这里粘贴生成的物模型JSON`
    model, err := ValidateWithFmt([]byte(jsonStr))
    if err != nil {
        t.Fatalf("校验失败: %v", err)
    }
    // 格式化输出校验后的结果（会自动补充默认值）
    out, _ := json.MarshalIndent(model, "", "  ")
    t.Logf("校验通过:\n%s", string(out))
}
```

校验通过后，`ValidateWithFmt` 会自动补充默认值（如 bool 的 mapping、int 的 min/max/step 等）。

### 生成模板

```bash
ur model template property --json   # 属性模板
ur model template event --yaml      # 事件模板
ur model template action --json     # 行为模板
ur model template full --yaml       # 完整物模型模板
```

## MQTT 协议使用参考

物模型定义完成后，设备通过以下 MQTT Topic 与平台交互：

| 功能 | Topic | method |
|------|-------|--------|
| 属性上报 | `$thing/up/property/{ProductID}/{DeviceName}` | report |
| 属性控制 | `$thing/down/property/{ProductID}/{DeviceName}` | control |
| 事件上报 | `$thing/up/event/{ProductID}/{DeviceName}` | eventPost |
| 行为调用 | `$thing/down/action/{ProductID}/{DeviceName}` | action |
| 批量上报 | `$thing/up/property/{ProductID}/{DeviceName}` | packReport |

属性上报示例：
```json
{ "method": "report", "msgToken": "123", "params": { "Temperature": 25.5, "Humidity": 60 } }
```

## 关键源码文件

| 文件 | 用途 |
|------|------|
| `backend/things/share/domain/schema/model.go` | Model/Property/Event/Action/Define 结构体 |
| `backend/things/share/domain/schema/define.go` | DataType 常量、枚举定义 |
| `backend/things/share/domain/schema/validate.go` | 校验逻辑和约束常量 |

## 参考文档

- 物模型概念: `doc/docs/use/00.介绍/06.基本概念/02.物模型/`
- 命名规范: `doc/docs/use/00.介绍/06.基本概念/02.物模型/02.英文标识符命名规范/README.md`
- 数据类型详解: `doc/docs/use/00.介绍/06.基本概念/02.物模型/03.数据类型详解/README.md`
- 物模型协议: `doc/docs/use/10.设备接入/01.联犀协议/14.物模型协议/README.md`

## 子文档索引

本 Skill 的详细内容已拆分为子文档，按任务加载：

- [实战案例](references/examples.md) —— CJT188 水表、DLT645 三相电表从 Excel 定义到物模型 JSON 的完整转换示例。
- [批量导入](references/batch-import.md) —— 通用物模型（CommonSchema）JSON 格式、Excel 转换脚本、打包上传与批量导入 API、结果解读与常见错误。
- [Affordance 说明](references/affordance.md) —— 三种 Affordance 结构、与场景联动字段的对应关系、CommonSchema 中 affordance 字段的写法与源码位置。
