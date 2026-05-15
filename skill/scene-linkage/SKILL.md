---
name: scene-linkage-skill
description: "Use when generating, validating, or troubleshooting 联犀 IoT scene linkage (场景联动) JSON definitions for triggers (if), conditions (when), and actions (then)."
---

# 场景联动生成与校验 Skill

帮助用户根据业务需求描述生成符合联犀平台规范的场景联动 JSON，并理解 `if`/`when`/`then` 三个核心字段的 JSON 结构。

## 工作流程

1. **理解需求**：用户描述联动场景（如"温度超过30度时开空调并发送通知"）
2. **确定场景类型**：`auto`（自动触发）或 `manual`（手动触发）
3. **构建 If（触发器）**：设备触发 / 定时触发 / 天气触发
4. **构建 When（生效条件）**：有效时间段、无效时间段、附加条件
5. **构建 Then（执行动作）**：设备控制 / 延时 / 通知 / 告警 / 执行其他场景
6. **校验**：使用校验脚本验证 JSON 合法性
7. **输出**：返回校验通过的场景联动 JSON

## 重要说明：API 中的字符串 JSON 字段

后端 API 接口（如 `scene-info.ts`）中 `if`/`when`/`then` 字段类型为 `string`，但这些字符串内部是**JSON 序列化后的对象**。实际构造和解析时按以下结构处理：

```json
{
  "type": "auto",
  "name": "温度告警联动",
  "if": "{\"triggers\":[{\"type\":\"device\",...}]}",
  "when": "{\"validRanges\":[],\"invalidRanges\":[],\"conditions\":{...}}",
  "then": "{\"actions\":[{\"type\":\"device\",...}]}"
}
```

**AI 生成时应输出结构化的 JSON 对象，由调用方自行 `JSON.stringify` 后填入 API 字段。**

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

### TriggerDevice（设备触发）

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| productID | string | 是 | 产品ID |
| selectType | string | 是 | `all`/`fixed`/`area`/`areaWithChildren`/`group` |
| deviceName | string | fixed时必填 | 设备名 |
| areaID | string | area时必填 | 区域ID |
| groupID | string | group时必填 | 分组ID |
| mode | string | 否 | `edge`(边缘触发,默认) / `level`(水平触发) |
| type | string | 是 | `connected`/`disConnected`/`propertyReport`/`eventReport` |
| stateKeep | StateKeep | 否 | 状态保持配置 |
| dataID | string | property/event时必填 | 属性/事件标识符 |
| termType | string | 否 | 比较类型：`eq`/`not`/`btw`/`gt`/`gte`/`lt`/`lte`/`in`/`all` |
| values | string[] | 否 | 比较值列表 |
| terms | Cmp[] | 否 | 多条件比较（结构体类型用） |

**StateKeep 状态保持**：
```json
{ "type": "duration", "value": 60 }
```
- `type`: `duration`(持续秒数) / `repeat`(重复次数)
- `value`: 持续时间(秒) 或 重复次数

### TriggerTimer（定时触发）

```json
{
  "execType": "at",
  "execAt": 28800,
  "repeatType": "week",
  "execRepeat": "1111100"
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| execType | string | `at`(定点) / `sunRises`(日出) / `sunSet`(日落) / `loop`(循环) |
| execAdd | int | 日出日落前后偏移秒数（-10800 ~ 10800） |
| execAt | int | 执行时间点，从0点起的秒数（如 8:00 = 28800） |
| execLoopStart | int | 循环开始时间（秒） |
| execLoopEnd | int | 循环结束时间（秒） |
| execLoop | int | 循环间隔（秒） |
| repeatType | string | `once`/`week`/`mount`/`allDay`/`customRange` |
| execRepeat | string | 二进制重复规则：周(7位) / 月(31位) |
| execRepeatStart | string | 自定义范围开始日期（2006-01-02） |
| execRepeatEnd | string | 自定义范围结束日期（2006-01-02） |

**execRepeat 二进制规则**：
- 周重复：`"1111100"` = 周一到周五（从高位到低位：周日至周六）
- 月重复：`"1111111111111111111111111111111"` = 每天

### TriggerWeather（天气触发）

```json
{
  "type": "temp",
  "termType": "gt",
  "values": ["35"]
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| type | string | `temp`(温度) / `humidity`(湿度) |
| termType | string | 比较类型，同设备触发 |
| values | string[] | 比较值 |

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

### Conditions（附加条件）

| 字段 | 类型 | 说明 |
|------|------|------|
| type | string | `and`(全部满足) / `or`(任一满足) |
| terms | Term[] | 条件项列表 |

**Term**：
```json
{
  "order": 1,
  "columnType": "property",
  "property": { ... },
  "weather": { ... },
  "time": { ... }
}
```

- `columnType`: `property` / `event` / `time` / `weather`
- `property`: TermProperty（含 productID、deviceName、Compare）
- `weather`: TermWeather（同 TriggerWeather）
- `time`: TermTime（系统时间或日出日落偏移比较）

**TermTime**：
```json
{ "type": "sys", "termType": "btw", "values": ["28800", "64800"] }
```
- `type`: `sys`(系统时间) / `sunRises`(日出后) / `sunSet`(日落后)
- `termType`/`values`: 同比较逻辑

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

### ActionDevice（设备动作）

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| productID | string | 是 | 产品ID |
| selectType | string | 是 | `all`/`fixed`/`area`/`areaWithChildren`/`group` |
| deviceName | string | fixed时必填 | 设备名 |
| areaID | string | area时必填 | 区域ID |
| groupID | string | group时必填 | 分组ID |
| type | string | 是 | `propertyControl`(属性控制) / `action`(调用行为) |
| dataID | string | 否 | 属性/行为标识符 |
| value | string | 单值时必填 | 发送的值（结构体类型传 JSON 对象字符串） |
| values | DeviceValue[] | 多值时必填 | 多属性控制列表 |

**DeviceValue**：
```json
{ "dataID": "Brightness", "value": "80" }
```

**结构体类型 value 示例**：
```json
{
  "dataID": "Color",
  "value": "{\"Red\":255,\"Green\":0,\"Blue\":0}"
}
```

### ActionNotify（通知动作）

```json
{
  "type": "message",
  "notifyCode": "ruleScene",
  "userType": "deviceProjectAdmin",
  "accounts": ["admin"],
  "userIDs": ["123"],
  "params": { "body": "告警内容" }
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| type | string | `sms`/`email`/`dingTalk`/`wxMini`/`message`/`phoneCall`/`wxEWebhook` |
| notifyCode | string | `ruleScene`(场景联动通知) / `ruleDeviceAlarm`(设备告警通知) |
| userType | string | `account`/`userID`/`deviceOwner`/`deviceProjectAdmin`/`deviceAreaAdmin`/`deviceProjectAll`/`deviceAreaAll` |
| accounts | string[] | userType=account时必填 |
| userIDs | string[] | userType=userID时必填 |
| params | map[string]string | 模板参数，如 `{"body":"告警内容"}` |

### ActionAlarm（告警动作）

```json
{ "mode": "trigger" }
```
- `mode`: `trigger`(触发告警) / `relieve`(解除告警)

### ActionScene（执行其他场景）

```json
{
  "sceneID": "123",
  "status": 1
}
```
- `sceneID`: 目标场景ID
- `status`: 自动化场景时修改的状态（1=启用, 2=禁用）

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

## 枚举/类型速查表

| 类型 | 可选值 |
|------|--------|
| SceneType | `manual`, `auto` |
| DeviceMode | `single`, `multi` |
| Status | `1`(启用), `2`(禁用) |
| TriggerType | `device`, `timer`, `weather`, `manual`, `scene` |
| SelectType | `all`, `fixed`, `area`, `areaWithChildren`, `group` |
| TriggerDeviceMode | `edge`, `level` |
| TriggerDeviceType | `connected`, `disConnected`, `propertyReport`, `eventReport` |
| ExecType | `at`, `sunRises`, `sunSet`, `loop` |
| RepeatType | `once`, `week`, `mount`, `allDay`, `customRange` |
| WeatherType | `temp`, `humidity` |
| WhenRangeType | `date`, `time` |
| DateRangeType | `allDay`, `workday`, `weekend`, `holiday`, `customRange`, `customWeek` |
| TimeRangeType | `allDay`, `light`, `night`, `customRange` |
| TermCondType | `and`, `or` |
| TermColumnType | `property`, `event`, `time`, `weather` |
| ActionType | `device`, `delay`, `notify`, `alarm`, `scene` |
| ActionDeviceType | `propertyControl`, `action` |
| NotifyType | `sms`, `email`, `dingTalk`, `wxMini`, `message`, `phoneCall`, `wxEWebhook` |
| NotifyUserType | `account`, `userID`, `deviceOwner`, `deviceProjectAdmin`, `deviceAreaAdmin`, `deviceProjectAll`, `deviceAreaAll` |
| AlarmMode | `trigger`, `relieve` |
| CmpType | `eq`, `not`, `btw`, `gt`, `gte`, `lt`, `lte`, `in`, `all` |
| StateKeepType | `duration`, `repeat` |

<!-- API_LIST:scene-linkage -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/scene/info/create` | 新增场景 | admin |
| POST | `/api/v1/things/scene/info/delete` | 删除场景 | admin |
| POST | `/api/v1/things/scene/info/get-list` | 获取场景信息列表 | admin |
| POST | `/api/v1/things/scene/info/get-one` | 获取场景信息详情 | admin |
| POST | `/api/v1/things/scene/info/manually-trigger` | 手动触发场景 | admin |
| POST | `/api/v1/things/scene/info/update` | 更新场景 | admin |
| POST | `/api/v1/things/scene/log/get-list` | 获取场景日志列表 | admin |

<!-- END_API_LIST -->

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

### 案例2：定时开关灯（日出日落模式）

**需求**：每天日出后30分钟自动关闭路灯，日落后30分钟自动开启路灯。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "路灯定时开关",
  "desc": "根据日出日落时间自动控制路灯",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "timer",
        "order": 1,
        "timer": {
          "execType": "sunRises",
          "execAdd": 1800,
          "repeatType": "allDay"
        }
      },
      {
        "type": "timer",
        "order": 2,
        "timer": {
          "execType": "sunSet",
          "execAdd": 1800,
          "repeatType": "allDay"
        }
      }
    ]
  },
  "when": {
    "validRanges": [],
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
          "productID": "streetLight01",
          "selectType": "all",
          "type": "propertyControl",
          "dataID": "PowerSwitch",
          "value": "0"
        }
      }
    ]
  }
}
```

**注意**：日出日落模式需要项目配置地理位置（经纬度），否则校验失败。

### 案例3：设备上线通知+延时执行

**需求**：当任意设备上线时，延时10秒后发送钉钉通知给指定账号。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "设备上线通知",
  "desc": "设备上线后延时通知",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "device",
        "order": 1,
        "device": {
          "productID": "smartDevice01",
          "selectType": "all",
          "type": "connected"
        }
      }
    ]
  },
  "when": {
    "validRanges": [],
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
        "type": "delay",
        "delay": 10
      },
      {
        "order": 2,
        "type": "notify",
        "notify": {
          "type": "dingTalk",
          "notifyCode": "ruleScene",
          "userType": "account",
          "accounts": ["admin"],
          "params": {
            "body": "有设备上线"
          }
        }
      }
    ]
  }
}
```

### 案例4：结构体属性条件触发

**需求**：当设备的GPS位置纬度大于30且经度小于120时触发告警。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "GPS围栏告警",
  "desc": "设备超出指定地理围栏时触发",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "device",
        "order": 1,
        "device": {
          "productID": "gpsDevice01",
          "selectType": "fixed",
          "deviceName": "gps_001",
          "type": "propertyReport",
          "dataID": "Location",
          "terms": [
            { "column": "Location.Lat", "termType": "gt", "values": ["30.0"] },
            { "column": "Location.Lng", "termType": "lt", "values": ["120.0"] }
          ]
        }
      }
    ]
  },
  "when": {
    "validRanges": [],
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
        "type": "alarm",
        "alarm": {
          "mode": "trigger"
        }
      }
    ]
  }
}
```

### 案例5：When 条件组合（工作时间 + 湿度条件）

**需求**：在工作日的 9:00-18:00，当湿度超过80%时开启除湿机。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "高湿度自动除湿",
  "desc": "工作日高湿度自动开启除湿机",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "device",
        "order": 1,
        "device": {
          "productID": "humiditySensor01",
          "selectType": "fixed",
          "deviceName": "sensor_001",
          "type": "propertyReport",
          "dataID": "Humidity",
          "termType": "gt",
          "values": ["80"]
        }
      }
    ]
  },
  "when": {
    "validRanges": [
      {
        "order": 1,
        "type": "date",
        "dateRange": {
          "type": "workday"
        }
      },
      {
        "order": 2,
        "type": "time",
        "timeRange": {
          "type": "customRange",
          "startTime": 32400,
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
          "productID": "dehumidifier01",
          "selectType": "fixed",
          "deviceName": "dehu_001",
          "type": "propertyControl",
          "dataID": "PowerSwitch",
          "value": "1"
        }
      }
    ]
  }
}
```

## CLI 校验工具

场景联动 JSON 可使用 CLI 命令进行基础结构验证：

```bash
ur scene validate scene.json
```

校验内容：
- 必填字段是否缺失
- 枚举值是否合法
- Compare 的 termType 与 values 数量是否匹配
- Action 的类型与对应详情字段是否一致

也支持生成模板：

```bash
ur scene template auto     # 生成自动触发场景模板
ur scene template manual   # 生成手动触发场景模板
```

## 与物模型的关联

场景联动中的 `dataID` 对应物模型的 `identifier`。生成场景联动 JSON 前，需要了解目标产品的物模型定义：

| 场景联动字段 | 物模型对应 |
|-------------|-----------|
| dataID | Property.identifier / Event.identifier / Action.identifier |
| termType + values | 需匹配 Property.Define.type 的数据类型 |
| schemaAffordance | 只读，后端自动填充为物模型 affordance JSON |

**affordance 说明**：`schemaAffordance` 是后端自动填充的只读字段，包含物模型的完整定义（类型、范围、枚举映射等），用于前端展示和校验。其结构对应 `schema.PropertyAffordance` / `schema.EventAffordance` / `schema.ActionAffordance`。

生成场景联动时，AI 不需要填写 `schemaAffordance`，但需要确保：
1. `dataID` 存在于目标产品的物模型中
2. `value` 的数据类型与物模型定义一致
3. 结构体类型的 `value` 是合法的 JSON 对象字符串

## 关键源码文件

| 文件 | 用途 |
|------|------|
| `backend/things/service/udsvr/internal/domain/scene/scene.go` | Info 结构体、场景类型枚举 |
| `backend/things/service/udsvr/internal/domain/scene/if.go` | Trigger 结构体、触发器类型 |
| `backend/things/service/udsvr/internal/domain/scene/triggerDevice.go` | 设备触发详情 |
| `backend/things/service/udsvr/internal/domain/scene/triggerTimer.go` | 定时触发详情 |
| `backend/things/service/udsvr/internal/domain/scene/triggerWeather.go` | 天气触发详情 |
| `backend/things/service/udsvr/internal/domain/scene/when.go` | When 结构体、时间/日期范围 |
| `backend/things/service/udsvr/internal/domain/scene/whenCondition.go` | Conditions 和 Term 结构体 |
| `backend/things/service/udsvr/internal/domain/scene/whenConditionTermProperty.go` | 属性条件 Term |
| `backend/things/service/udsvr/internal/domain/scene/whenConditionTermTime.go` | 时间条件 Term |
| `backend/things/service/udsvr/internal/domain/scene/whenConditionTermWeather.go` | 天气条件 Term |
| `backend/things/service/udsvr/internal/domain/scene/then.go` | Then 结构体 |
| `backend/things/service/udsvr/internal/domain/scene/thenAction.go` | Action 通用结构 |
| `backend/things/service/udsvr/internal/domain/scene/thenActionDevice.go` | 设备动作详情 |
| `backend/things/service/udsvr/internal/domain/scene/thenActionNotify.go` | 通知动作详情 |
| `backend/things/service/udsvr/internal/domain/scene/thenActionAlarm.go` | 告警动作详情 |
| `backend/things/service/udsvr/internal/domain/scene/thenActionScene.go` | 场景动作详情 |
| `backend/things/service/udsvr/internal/domain/scene/compare.go` | Compare 比较逻辑 |
| `backend/things/service/udsvr/internal/domain/scene/time.go` | 时间范围、重复类型枚举 |
| `backend/things/service/udsvr/internal/domain/scene/stateKeep.go` | 状态保持配置 |
| `backend/things/share/domain/schema/affordance.go` | Affordance 结构定义 |
| `backend/things/share/domain/schema/model.go` | 物模型 Property/Event/Action 定义 |

## 前端模型参考

| 文件 | 用途 |
|------|------|
| `apps/web/packages/api/src/things-api/models/scene-info.ts` | API 接口 SceneInfo 定义 |

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
