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
ur things device info get-list --page 1 --size 10

# 查询设备详情
ur things device info get-one -p xxx -d yyy
```

## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| 日志查询 | 属性/事件/命令/上下线/诊断/异常/SDK 日志详解 | [cli/device-log.md](references/device-log.md) |
| 实时调试 | 属性控制、行为调用、Mock数据、模拟上报详解 | [cli/device-debug.md](references/device-debug.md) |

### 物模型查询

```bash
# 设备级物模型
ur things schema get-list -p xxx -d yyy

# 或产品级物模型
ur things schema get-list -p xxx
```

> 完整命令帮助：`ur things device help`

## 设备实时调试

### 查询设备物模型（调试前必备）

发送任何调试指令前，必须先查询设备物模型确认可用属性/行为/事件：

```bash
# 设备级物模型
ur things schema get-list -p xxx -d yyy

# 或产品级物模型
ur things schema get-list -p xxx
```

返回的 `identifier` 即为调试时使用的 `dataID`。

---

### 1. 属性控制（下行）

向设备发送属性控制指令：

```bash
ur things device control -p xxx -d yyy --data '{"PowerSwitch": 1, "Brightness": 80}'
```

| 参数 | 说明 |
|------|------|
| `productID` + `deviceName` | 目标设备 |
| `--data` | 属性键值对 JSON，key 必须与物模型 `identifier` 完全一致 |

**影子设备控制模式**：
- 默认：指令下发到设备，同时更新影子设备
- 仅影子：`shadowControl` 参数控制行为（平台逻辑）
- 离线时：指令缓存到影子设备，设备上线后自动同步

---

### 2. 行为调用（下行）

调用设备的某个行为（Action）：

```bash
ur things device action send -p xxx -d yyy --data-id OpenValve --input '{"Duration": 30}'
```

| 参数 | 说明 |
|------|------|
| `--data-id` | 行为标识符（来自物模型 Action.identifier） |
| `--input` | 行为输入参数 JSON，key 必须与物模型 `input[].identifier` 一致 |

**获取行为执行结果**：
```bash
ur things device action get -p xxx -d yyy --data-id OpenValve
```

---

### 3. 回复设备行为调用（上行）

当设备主动调用云端行为（`dir: up`）时，云端回复结果：

```bash
ur things device action resp -p xxx -d yyy --data-id ReadMeter --output '{"EP": "1234.56"}'
```

---

### 4. 事件发送（下行通知）

向设备发送下行事件通知：

```bash
# 注：event-send 暂无 CLI 封装，需使用 ur api 直接调用
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
ur things device mock -p xxx -d yyy --data-id Temperature --num 5
```

| 参数 | 说明 |
|------|------|
| `--data-id` | 目标属性/行为/事件标识符 |
| `--num` | 生成数量 |

用途：调试时快速生成合法测试数据，无需手动构造。

---

## 日志类型速查

| 日志类型 | CLI 命令 | 典型用途 |
|---------|---------|----------|
| 属性最新值 | `ur things device log property` | 查看设备当前状态 |
| 属性历史 | `ur things device log property --data-id <id> --arg-func avg` | 分析属性变化趋势 |
| 事件 | `ur things device log event` | 排查告警/故障 |
| 命令 | `ur things device log send` | 审计谁操作了设备 |
| 上下线 | `ur things device log status` | 排查在线/离线问题 |
| 诊断(Hub) | `ur things device log hub` | 查看 MQTT 原始通信 |
| 异常 | `ur things device log abnormal` | 查看超时/错误异常 |
| SDK | `ur things device log sdk` | 查看设备端日志 |

## 典型调试工作流

### 工作流 1：排查设备属性不更新

```bash
# 1. 查看最新属性值
ur things device log property -p xxx -d yyy

# 2. 查看属性历史，确认是否有上报记录
ur things device log property -p xxx -d yyy \
  --data-id Temperature \
  --time-start ... --time-end ... --order 2

# 3. 查看诊断日志，确认 MQTT 消息是否到达平台
ur things device log hub -p xxx -d yyy \
  --actions property --time-start ... --time-end ...

# 4. 查看上下线日志，确认设备是否在线
ur things device log status -p xxx -d yyy \
  --time-start ... --time-end ...
```

### 工作流 2：测试属性控制

```bash
# 1. 查询物模型，确认控制属性
ur things schema get-list -p xxx -d yyy

# 2. 查看当前属性值
ur things device log property -p xxx -d yyy --data-ids PowerSwitch

# 3. 发送控制指令
ur things device control -p xxx -d yyy --data '{"PowerSwitch":1}'

# 4. 查看命令日志确认执行结果
ur things device log send -p xxx -d yyy \
  --actions propertyControlSend --time-start ... --time-end ...
```

### 工作流 3：排查设备离线

```bash
# 1. 查看最新上下线记录
ur things device log status -p xxx -d yyy --order 2 --page 1 --size 5

# 2. 查看异常日志
ur things device log abnormal -p xxx -d yyy --time-start ... --time-end ...

# 3. 查看诊断日志中的连接记录
ur things device log hub -p xxx -d yyy \
  --actions connected,disconnected --time-start ... --time-end ...

# 4. 查看 SDK 日志（如设备支持）
ur things device log sdk -p xxx -d yyy \
  --log-level 2 --time-start ... --time-end ...
```

## 注意事项

1. **时间戳为毫秒**：`--time-start`/`--time-end` 是 Unix 毫秒时间戳
2. **默认查询近30分钟**：前端日志页默认只查最近30分钟，CLI 调用建议显式指定时间范围
3. **物模型必须先查**：调试前务必确认 `dataID` 正确，大小写敏感
4. **离线设备控制**：属性控制指令会缓存到影子设备，设备上线后自动下发
5. **诊断日志最底层**：`hub-log` 是 MQTT 原始通信记录，排查问题时最彻底
6. **命令日志审计**：`send-log` 记录谁（account）、何时、发送了什么指令、结果如何

## CLI 命令缺口

以下 API 暂无对应的 CLI Layer 2 命令，仍需使用 `ur api` 直接调用：

| API | 说明 |
|-----|------|
| `/api/v1/things/device/interact/event-send` | 向设备发送下行事件通知 |
