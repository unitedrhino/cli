---
name: scene-linkage-skill
description: "Use when generating, validating, or troubleshooting 联犀 IoT scene linkage (场景联动) JSON definitions for triggers (if), conditions (when), and actions (then)."
---

# 场景联动生成与校验 Skill

帮助用户根据业务需求描述生成符合联犀平台规范的场景联动 JSON，并理解 `if`/`when`/`then` 三个核心字段的 JSON 结构。

## 详细参考文档

| 文档 | 内容 |
|------|------|
| [if-schema.md](references/if-schema.md) | If 触发器完整字段：TriggerDevice / TriggerTimer / TriggerWeather |
| [when-conditions.md](references/when-conditions.md) | When 生效条件完整字段：WhenRange / DateRange / TimeRange / Conditions / Term |
| [then-actions.md](references/then-actions.md) | Then 执行动作完整字段：ActionDevice / ActionNotify / ActionAlarm / ActionScene |
| [cases.md](references/cases.md) | 实战案例：日落开关灯、上线通知、GPS 围栏、湿度除湿等完整 JSON |
| [model-schema-mapping.md](references/model-schema-mapping.md) | 场景字段与物模型（dataID / value / schemaAffordance）的对应关系 |
| [scene.md](references/scene.md) | CLI 命令参考、枚举/类型速查表、关键源码文件、前端模型参考 |
| [api/](references/api/) | 场景联动 REST API 端点与请求体字段 |

## 工作流程

1. **理解需求**：用户描述联动场景（如"温度超过30度时开空调并发送通知"）
2. **确定场景类型**：`auto`（自动触发）或 `manual`（手动触发）
3. **构建 If（触发器）**：设备触发 / 定时触发 / 天气触发
4. **构建 When（生效条件）**：有效时间段、无效时间段、附加条件
5. **构建 Then（执行动作）**：设备控制 / 延时 / 通知 / 告警 / 执行其他场景
6. **校验**：使用校验脚本验证 JSON 合法性
7. **输出**：返回校验通过的场景联动 JSON

## 重要说明：CLI 中的字符串 JSON 字段

使用 CLI 命令（如 `ur things scene info create`）时，`if`/`when`/`then` 字段在 `--body` 参数中作为**JSON 序列化后的字符串**传入。实际构造和解析时按以下结构处理：

```json
{
  "type": "auto",
  "name": "温度告警联动",
  "if": "{\"triggers\":[{\"type\":\"device\",...}]}",
  "when": "{\"validRanges\":[],\"invalidRanges\":[],\"conditions\":{...}}",
  "then": "{\"actions\":[{\"type\":\"device\",...}]}"
}
```

**AI 生成时应输出结构化的 JSON 对象，由调用方自行 `JSON.stringify` 后填入 CLI 的 `--body` 参数。**

## Info（场景顶层结构）

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "温度超限告警",
  "desc": "当环境温度超过30度时触发告警并通知管理员",
  "status": 1,
  "if": { "triggers": [...] },
  "when": { "validRanges": [], "invalidRanges": [], "conditions": {...} },
  "then": { "actions": [...] }
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| type | string | 是 | `auto`(自动触发) / `manual`(手动触发) |
| deviceMode | string | 是 | `single`(单设备) / `multi`(多设备) |
| productID | string | 单设备必填 | 产品ID |
| deviceName | string | 单设备必填 | 设备名称 |
| name | string | 是 | 场景名称 |
| desc | string | 否 | 场景描述 |
| status | int | 否 | 1=启用(默认), 2=禁用 |
| if | If | auto必填 | 触发器定义 |
| when | When | 否 | 生效条件（手动触发不生效） |
| then | Then | 是 | 执行动作 |

## If（触发器 Triggers）

```json
{
  "triggers": [
    {
      "type": "device",
      "order": 1,
      "device": {
        "productID": "p123",
        "selectType": "fixed",
        "deviceName": "dev001",
        "type": "propertyReport",
        "mode": "edge",
        "dataID": "Temperature",
        "termType": "gt",
        "values": ["30"]
      }
    }
  ]
}
```

### Trigger 通用字段

| 字段 | 类型 | 说明 |
|------|------|------|
| type | string | `device` / `timer` / `weather` |
| order | int | 排序，从1开始 |
| device | TriggerDevice | 设备触发详情（type=device时必填） |
| timer | TriggerTimer | 定时触发详情（type=timer时必填） |
| weather | TriggerWeather | 天气触发详情（type=weather时必填） |

## When（生效条件）

```json
{
  "validRanges": [
    { "order": 1, "type": "time", "timeRange": { "type": "customRange", "startTime": 28800, "endTime": 64800 } }
  ],
  "invalidRanges": [],
  "conditions": {
    "type": "and",
    "terms": [
      {
        "order": 1,
        "columnType": "property",
        "property": {
          "productID": "p123",
          "deviceName": "dev001",
          "dataID": "Humidity",
          "termType": "lt",
          "values": ["80"]
        }
      }
    ]
  }
}
```

### WhenRange（时间/日期范围）

| 字段 | 类型 | 说明 |
|------|------|------|
| order | int | 排序 |
| type | string | `date`(日期范围) / `time`(时间范围) |
| dateRange | DateRange | type=date时必填 |
| timeRange | TimeRange | type=time时必填 |

**DateRange**：
```json
{ "type": "workday", "startDate": "2026-01-01", "endDate": "2026-12-31", "repeat": "1111100" }
```
- `type`: `allDay`/`workday`/`weekend`/`holiday`/`customRange`/`customWeek`
- `startDate`/`endDate`: 自定义范围时必填（格式 2006-01-02）
- `repeat`: customWeek时必填，7位二进制

**TimeRange**：
```json
{ "type": "customRange", "startTime": 28800, "endTime": 64800 }
```
- `type`: `allDay`/`light`/`night`/`customRange`
- `startTime`/`endTime`: 从0点起的秒数

## Then（执行动作）

```json
{
  "actions": [
    {
      "order": 1,
      "type": "device",
      "device": {
        "productID": "p456",
        "selectType": "fixed",
        "deviceName": "ac001",
        "type": "propertyControl",
        "dataID": "PowerSwitch",
        "value": "1"
      }
    },
    {
      "order": 2,
      "type": "delay",
      "delay": 5
    },
    {
      "order": 3,
      "type": "notify",
      "notify": {
        "type": "message",
        "notifyCode": "ruleScene",
        "userType": "deviceProjectAdmin",
        "params": { "body": "温度超限，已自动开启空调" }
      }
    }
  ]
}
```

### Action 通用字段

| 字段 | 类型 | 说明 |
|------|------|------|
| order | int | 执行顺序，从1开始 |
| type | string | `device`/`delay`/`notify`/`alarm`/`scene` |
| delay | int | type=delay时必填，延时秒数 |
| device | ActionDevice | type=device时必填 |
| notify | ActionNotify | type=notify时必填 |
| alarm | ActionAlarm | type=alarm时必填 |
| scene | ActionScene | type=scene时必填 |

## Compare（比较逻辑）

### 简单比较（单值类型）

```json
{
  "dataID": "Temperature",
  "termType": "gt",
  "values": ["30"]
}
```

### 结构体多条件比较

```json
{
  "dataID": "Position",
  "terms": [
    { "column": "Position.Lat", "termType": "gt", "values": ["30.0"] },
    { "column": "Position.Lng", "termType": "lt", "values": ["120.0"] }
  ]
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| dataID | string | 属性/事件标识符 |
| termType | string | 比较类型 |
| values | string[] | 比较值 |
| terms | Cmp[] | 结构体多条件 |

**Cmp 结构**：
| 字段 | 类型 | 说明 |
|------|------|------|
| column | string | 结构体字段路径，如 `aa.bb` / `aa.1.bb` |
| termType | string | 比较类型 |
| values | string[] | 比较值 |

### 比较类型速查表

| 类型 | 含义 | 所需values长度 |
|------|------|---------------|
| eq | 等于 | 1 |
| not | 不等于 | 1 |
| btw | 在之间（闭区间） | 2 |
| gt | 大于 | 1 |
| gte | 大于等于 | 1 |
| lt | 小于 | 1 |
| lte | 小于等于 | 1 |
| in | 在值列表中 | n |
| all | 全部触发 | 0 |

## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| `ur things scene info` | 场景查询/创建/更新/删除/触发 | [scene.md](references/scene.md) |
| `ur things scene log` | 场景执行日志查询 | [scene.md](references/scene.md) |
| `ur scene validate` | 场景JSON校验 | [scene.md](references/scene.md) |
| `ur scene template` | 场景模板生成 | [scene.md](references/scene.md) |

> 完整命令帮助：`ur things scene help`、`ur scene help`

## 实战案例

### 案例1：温度超限自动开空调+通知

**需求**：当温度传感器上报温度超过30度时，如果当前时间在8:00-18:00之间，自动打开空调，并发送站内信通知项目管理员。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "温度超限自动降温",
  "desc": "工作时间温度超过30度自动开启空调",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "device",
        "order": 1,
        "device": {
          "productID": "tempSensor01",
          "selectType": "fixed",
          "deviceName": "sensor_001",
          "type": "propertyReport",
          "mode": "edge",
          "dataID": "Temperature",
          "termType": "gt",
          "values": ["30"]
        }
      }
    ]
  },
  "when": {
    "validRanges": [
      {
        "order": 1,
        "type": "time",
        "timeRange": {
          "type": "customRange",
          "startTime": 28800,
          "endTime": 64800
        }
      }
    ],
    "invalidRanges": [],
    "conditions": {
      "type": "and",
      "terms": []
    }
  },
  "then": {
    "actions": [
      {
        "order": 1,
        "type": "device",
        "device": {
          "productID": "airConditioner01",
          "selectType": "fixed",
          "deviceName": "ac_001",
          "type": "propertyControl",
          "dataID": "PowerSwitch",
          "value": "1"
        }
      },
      {
        "order": 2,
        "type": "notify",
        "notify": {
          "type": "message",
          "notifyCode": "ruleScene",
          "userType": "deviceProjectAdmin",
          "params": {
            "body": "温度超限，已自动开启空调降温"
          }
        }
      }
    ]
  }
}
```

## 常见错误排查

| 错误 | 原因 | 修复 |
|------|------|------|
| 触发器类型不支持 | type 字段值非法 | 检查 TriggerType 枚举 |
| 设备选择方式不支持 | selectType 值非法 | 检查 SelectType 枚举 |
| dataID不存在 | 标识符不在物模型中 | 核对产品物模型的 identifier |
| 动态条件类型需要填写参数 | values 为空但 termType 不是 all | 补充 values 或改用 all |
| 需要填写地理位置 | 使用日出日落但项目无经纬度 | 在项目设置中配置位置 |
| 时间范围只能在0到24小时 | execAt/execLoop 超出 86400 | 检查秒数是否在 0-86400 范围 |
| 传的值不能为空 | ActionDevice value/values 未填 | 补充 value 或 values |
