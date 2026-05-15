---
name: ur-device-debug
description: "设备调试与日志查询：属性/事件/命令/上下线/异常/诊断/SDK日志查询，属性控制、行为调用、事件发送、Mock数据生成、WebSocket实时调试。triggers: 设备调试, 设备日志, 属性控制, 行为调用, 事件发送, 上下线日志, 诊断日志, SDK日志, 异常日志, 命令日志, 实时调试"
metadata:
  hermes:
    tags: [device, debug, log, trace, monitor, diagnose]
---

# ur-device-debug — 设备调试与日志

## 概述

设备调试 Skill 覆盖两大能力：
1. **日志查询** — 查询设备历史消息（属性、事件、命令、上下线、异常、诊断、SDK）
2. **实时调试** — 向设备发送控制指令、调用行为、模拟事件，通过 WebSocket 实时观察设备响应

## 前置条件

调试设备需要知道 `productID` 和 `deviceName`（设备唯一ID）。如不确定：

```bash
# 查询设备列表
ur api /api/v1/things/device/info/get-list \
  --body '{"page":{"page":1,"size":10}}'

# 查询设备详情
ur api /api/v1/things/device/info/get-one \
  --body '{"productID":"xxx","deviceName":"yyy"}'
```

## 设备日志查询

<!-- API_LIST:ur-device-debug -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/interact/action-get-one` | 获取调用设备行为的结果 | admin |
| POST | `/api/v1/things/device/interact/action-resp` | 回复设备行为调用结果 | admin |
| POST | `/api/v1/things/device/interact/action-send` | 调用设备行为 | admin |
| POST | `/api/v1/things/device/interact/event-send` | 下行事件通知设备 | admin |
| POST | `/api/v1/things/device/interact/gateway-get-found-send` | 请求网关上报拓扑关系 | admin |
| POST | `/api/v1/things/device/interact/gateway-notify-bind-send` | 通知网关绑定子设备 | admin |
| POST | `/api/v1/things/device/interact/property-control-batch-send` | 批量调用设备属性 | admin |
| POST | `/api/v1/things/device/interact/property-control-get-one` | 获取调用设备属性的结果 | admin |
| POST | `/api/v1/things/device/interact/property-control-send` | 调用设备属性 | admin |
| POST | `/api/v1/things/device/interact/property-get-report-batch-send` | 批量请求设备获取设备最新属性 | admin |
| POST | `/api/v1/things/device/interact/property-get-report-send` | 请求设备获取设备最新属性 | admin |
| POST | `/api/v1/things/device/interact/schema-mock-gen` | 生成物模型模拟数据 | admin |

<!-- END_API_LIST -->

### 通用查询参数

所有日志查询共享以下基础参数：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品ID |
| `deviceName` | string | 是 | 设备名称（设备唯一ID） |
| `timeStart` | string | 否 | 开始时间戳（毫秒） |
| `timeEnd` | string | 否 | 结束时间戳（毫秒） |
| `page` | object | 否 | `{ page: 1, size: 20 }` 分页 |

> 日志查询默认时间范围为**近30分钟**（前端行为）。CLI 调用时建议显式指定 `timeStart`/`timeEnd`。

---

### 1. 属性日志

#### 最新属性值

查询设备当前各属性的最新值：

```bash
ur api /api/v1/things/device/msg/property-latest/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataIDs": ["Temperature", "Humidity"],
    "ignoreEmpty": true
  }'
```

| 专属参数 | 说明 |
|---------|------|
| `dataIDs` | 属性标识符数组，不传则返回全部 |
| `ignoreEmpty` | 是否跳过空值 |

响应字段：
```json
{
  "list": [
    {"dataID": "Temperature", "dataName": "温度", "value": "25.3", "timestamp": 1715500800000}
  ]
}
```

#### 属性历史记录

查询单个属性的历史变化：

```bash
ur api /api/v1/things/device/msg/property-log/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataID": "Temperature",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "argFunc": "avg",
    "interval": 1,
    "intervalUnit": "h",
    "order": 2
  }'
```

| 专属参数 | 说明 |
|---------|------|
| `dataID` | 单个属性标识符 |
| `argFunc` | 聚合函数：`avg`/`first`/`last`/`count`/`twa` |
| `interval` + `intervalUnit` | 时间窗口聚合 |
| `order` | 排序：1=升序(旧→新)，2=降序(新→旧) |

---

### 2. 事件日志

```bash
ur api /api/v1/things/device/msg/event-log/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "types": ["info", "alert"],
    "dataID": "PowerAlarm"
  }'
```

| 专属参数 | 说明 |
|---------|------|
| `types` | 事件类型过滤：`info`(信息)/`alert`(告警)/`fault`(故障) |
| `dataID` | 特定事件标识符 |
| `deviceNames` | 可批量查询多个设备（传数组） |

响应字段：
```json
{
  "list": [
    {"timestamp": 1715500800000, "type": "alert", "dataID": "PowerAlarm", "params": {"Voltage": "220"}}
  ]
}
```

---

### 3. 命令日志（发送记录）

查询向设备发送指令的历史记录：

```bash
ur api /api/v1/things/device/msg/send-log/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "actions": ["propertyControlSend"],
    "resultCode": 200,
    "withUser": true
  }'
```

| 专属参数 | 说明 |
|---------|------|
| `actions` | 指令类型：`propertyControlSend`/`propertyGetReportSend`/`actionSend` |
| `resultCode` | 结果码过滤：200=成功 |
| `dataID` / `dataIDs` | 属性/行为标识符过滤 |
| `withUser` | 是否包含操作人信息 |

响应字段：
```json
{
  "list": [
    {
      "timestamp": 1715500800000,
      "account": "admin",
      "action": "propertyControlSend",
      "dataID": "PowerSwitch",
      "content": "{\"PowerSwitch\":1}",
      "resultCode": 200
    }
  ]
}
```

---

### 4. 上下线日志

```bash
ur api /api/v1/things/device/msg/status-log/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "status": 1
  }'
```

| 专属参数 | 说明 |
|---------|------|
| `status` | 状态过滤：1=上线，2=离线 |

响应字段：
```json
{
  "list": [
    {"timestamp": 1715500800000, "status": 1, "productID": "xxx", "deviceName": "yyy"}
  ]
}
```

---

### 5. 诊断日志（Hub日志）

查询设备与平台之间的 MQTT 通信原始日志：

```bash
ur api /api/v1/things/device/msg/hub-log/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "actions": ["property", "event"],
    "topics": ["$thing/up/property"],
    "content": "report"
  }'
```

| 专属参数 | 说明 |
|---------|------|
| `actions` | MQTT 动作：`connected`/`disconnected`/`property`/`event`/`action`/`thing` |
| `topics` | MQTT Topic 过滤 |
| `content` | 内容模糊匹配 |
| `requestID` | 按请求 token 精确过滤 |

响应字段：
```json
{
  "list": [
    {
      "timestamp": 1715500800000,
      "action": "property",
      "topic": "$thing/up/property/xxx/yyy",
      "content": "{...}",
      "resultCode": 200,
      "respPayload": "{...}",
      "traceID": "abc123",
      "requestID": "req456"
    }
  ]
}
```

---

### 6. 异常日志

```bash
ur api /api/v1/things/device/msg/abnormal-log/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "action": 1,
    "type": "timeout"
  }'
```

| 专属参数 | 说明 |
|---------|------|
| `action` | 1=触发，2=恢复 |
| `type` | 异常类型 |

---

### 7. SDK日志

```bash
ur api /api/v1/things/device/msg/sdk-log/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "logLevel": 4
  }'
```

| 专属参数 | 说明 |
|---------|------|
| `logLevel` | 1=关闭, 2=错误, 3=警告, 4=信息, 5=调试 |

---

## 设备实时调试

### 查询设备物模型（调试前必备）

发送任何调试指令前，必须先查询设备物模型确认可用属性/行为/事件：

```bash
# 设备级物模型
ur api /api/v1/things/device/schema/get-list \
  --body '{"productID":"xxx","deviceName":"yyy"}'

# 或产品级物模型
ur api /api/v1/things/product/schema/get-list \
  --body '{"productID":"xxx"}'
```

返回的 `identifier` 即为调试时使用的 `dataID`。

---

### 1. 属性控制（下行）

向设备发送属性控制指令：

```bash
ur api /api/v1/things/device/interact/property-control-send \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "data": {"PowerSwitch": 1, "Brightness": 80}
  }'
```

| 参数 | 说明 |
|------|------|
| `productID` + `deviceName` | 目标设备 |
| `data` | 属性键值对，key 必须与物模型 `identifier` 完全一致 |

**影子设备控制模式**：
- 默认：指令下发到设备，同时更新影子设备
- 仅影子：`shadowControl` 参数控制行为（平台逻辑）
- 离线时：指令缓存到影子设备，设备上线后自动同步

---

### 2. 行为调用（下行）

调用设备的某个行为（Action）：

```bash
ur api /api/v1/things/device/interact/action-send \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataID": "OpenValve",
    "input": {"Duration": 30}
  }'
```

| 参数 | 说明 |
|------|------|
| `dataID` | 行为标识符（来自物模型 Action.identifier） |
| `input` | 行为输入参数，key 必须与物模型 `input[].identifier` 一致 |

**获取行为执行结果**：
```bash
ur api /api/v1/things/device/interact/action-get-one \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataID": "OpenValve"
  }'
```

---

### 3. 回复设备行为调用（上行）

当设备主动调用云端行为（`dir: up`）时，云端回复结果：

```bash
ur api /api/v1/things/device/interact/action-resp \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataID": "ReadMeter",
    "output": {"EP": "1234.56"}
  }'
```

---

### 4. 事件发送（下行通知）

向设备发送下行事件通知：

```bash
ur api /api/v1/things/device/interact/event-send \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataID": "ConfigUpdate",
    "params": {"interval": 300}
  }'
```

---

### 5. 生成 Mock 数据

根据设备物模型自动生成符合数据类型约束的 Mock 数据：

```bash
ur api /api/v1/things/device/interact/schema-mock-gen \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataID": "Temperature",
    "num": 5
  }'
```

| 参数 | 说明 |
|------|------|
| `dataID` | 目标属性/行为/事件标识符 |
| `num` | 生成数量 |

用途：调试时快速生成合法测试数据，无需手动构造。

---

## 日志类型速查

| 日志类型 | API 端点 | 典型用途 |
|---------|---------|----------|
| 属性最新值 | `property-latest/get-list` | 查看设备当前状态 |
| 属性历史 | `property-log/get-list` | 分析属性变化趋势 |
| 事件 | `event-log/get-list` | 排查告警/故障 |
| 命令 | `send-log/get-list` | 审计谁操作了设备 |
| 上下线 | `status-log/get-list` | 排查在线/离线问题 |
| 诊断(Hub) | `hub-log/get-list` | 查看 MQTT 原始通信 |
| 异常 | `abnormal-log/get-list` | 查看超时/错误异常 |
| SDK | `sdk-log/get-list` | 查看设备端日志 |

## 典型调试工作流

### 工作流 1：排查设备属性不更新

```bash
# 1. 查看最新属性值
ur api /api/v1/things/device/msg/property-latest/get-list \
  --body '{"productID":"xxx","deviceName":"yyy"}'

# 2. 查看属性历史，确认是否有上报记录
ur api /api/v1/things/device/msg/property-log/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","dataID":"Temperature","timeStart":"...","timeEnd":"...","order":2}'

# 3. 查看诊断日志，确认 MQTT 消息是否到达平台
ur api /api/v1/things/device/msg/hub-log/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","actions":["property"],"timeStart":"...","timeEnd":"..."}'

# 4. 查看上下线日志，确认设备是否在线
ur api /api/v1/things/device/msg/status-log/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","timeStart":"...","timeEnd":"..."}'
```

### 工作流 2：测试属性控制

```bash
# 1. 查询物模型，确认控制属性
ur api /api/v1/things/device/schema/get-list \
  --body '{"productID":"xxx","deviceName":"yyy"}'

# 2. 查看当前属性值
ur api /api/v1/things/device/msg/property-latest/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","dataIDs":["PowerSwitch"]}'

# 3. 发送控制指令
ur api /api/v1/things/device/interact/property-control-send \
  --body '{"productID":"xxx","deviceName":"yyy","data":{"PowerSwitch":1}}'

# 4. 查看命令日志确认执行结果
ur api /api/v1/things/device/msg/send-log/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","actions":["propertyControlSend"],"timeStart":"...","timeEnd":"..."}'
```

### 工作流 3：排查设备离线

```bash
# 1. 查看最新上下线记录
ur api /api/v1/things/device/msg/status-log/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","order":2,"page":{"page":1,"size":5}}'

# 2. 查看异常日志
ur api /api/v1/things/device/msg/abnormal-log/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","timeStart":"...","timeEnd":"..."}'

# 3. 查看诊断日志中的连接记录
ur api /api/v1/things/device/msg/hub-log/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","actions":["connected","disconnected"],"timeStart":"...","timeEnd":"..."}'

# 4. 查看 SDK 日志（如设备支持）
ur api /api/v1/things/device/msg/sdk-log/get-list \
  --body '{"productID":"xxx","deviceName":"yyy","logLevel":2,"timeStart":"...","timeEnd":"..."}'
```

## 注意事项

1. **时间戳为毫秒**：`timeStart`/`timeEnd` 是 Unix 毫秒时间戳
2. **默认查询近30分钟**：前端日志页默认只查最近30分钟，CLI 调用建议显式指定时间范围
3. **物模型必须先查**：调试前务必确认 `dataID` 正确，大小写敏感
4. **离线设备控制**：属性控制指令会缓存到影子设备，设备上线后自动下发
5. **诊断日志最底层**：`hub-log` 是 MQTT 原始通信记录，排查问题时最彻底
6. **命令日志审计**：`send-log` 记录谁（account）、何时、发送了什么指令、结果如何
