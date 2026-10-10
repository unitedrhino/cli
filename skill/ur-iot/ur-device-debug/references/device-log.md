# CLI 命令 — 设备日志查询详解

## 命令语法

```
ur things device log <subcommand> [选项]
```

子命令：`property`、`event`、`send`、`status`、`hub`、`abnormal`、`sdk`

---

## 通用参数

所有日志查询子命令共享以下通用参数：

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称（设备唯一ID） |
| --time-start | | 否 | string | 开始时间戳（毫秒） |
| --time-end | | 否 | string | 结束时间戳（毫秒） |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

---

## property — 查询属性日志

查询设备属性最新值或历史记录。不指定 `--data-id` 时查询最新值，指定 `--data-id` 时查询历史记录。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --data-id | | 否 | string | 属性标识符（查询历史时必填） |
| --data-ids | | 否 | string | 多个属性标识符，逗号分隔（查询最新值时使用） |
| --ignore-empty | | 否 | bool | 跳过空值 |
| --arg-func | | 否 | string | 聚合函数：`avg`/`first`/`last`/`count`/`twa` |
| --interval | | 否 | int | 时间窗口间隔 |
| --interval-unit | | 否 | string | 时间窗口单位：`s`/`m`/`h`/`d` |
| --order | | 否 | int | 排序：1=升序，2=降序 |
| --time-start | | 否 | string | 开始时间戳（毫秒） |
| --time-end | | 否 | string | 结束时间戳（毫秒） |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询设备最新属性值
```bash
ur things device log property -p p_smartswitch_001 -d switch-001
```

#### 示例2：查询温度历史平均值（按小时聚合）
```bash
ur things device log property -p p_smartswitch_001 -d switch-001 \
  --data-id Temperature --arg-func avg --interval 1 --interval-unit h
```

#### 示例3：查询多个属性最新值（忽略空值）
```bash
ur things device log property -p p_smartswitch_001 -d switch-001 \
  --data-ids Temperature,Humidity --ignore-empty
```

#### 示例4：查询特定时间范围的属性历史
```bash
ur things device log property -p p_smartswitch_001 -d switch-001 \
  --data-id Temperature --time-start 1715500800000 --time-end 1715587199999 \
  --arg-func avg --interval 1 --interval-unit h --order 2
```

### 对应API

- 最新值：`POST /api/v1/things/device/msg/property-latest/get-list`
- 历史记录：`POST /api/v1/things/device/msg/property-log/get-list`

---

## event — 查询事件日志

查询设备上报的事件记录，支持按事件类型和标识符筛选。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --types | | 否 | string | 事件类型，逗号分隔：`info`/`alert`/`fault` |
| --data-id | | 否 | string | 特定事件标识符 |
| --time-start | | 否 | string | 开始时间戳（毫秒） |
| --time-end | | 否 | string | 结束时间戳（毫秒） |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询所有事件日志
```bash
ur things device log event -p p_smartswitch_001 -d switch-001
```

#### 示例2：查询告警和故障事件
```bash
ur things device log event -p p_smartswitch_001 -d switch-001 --types alert,fault
```

#### 示例3：查询特定事件
```bash
ur things device log event -p p_smartswitch_001 -d switch-001 --data-id PowerAlarm
```

### 对应API

- `POST /api/v1/things/device/msg/event-log/get-list`

---

## send — 查询命令发送日志

查询向设备发送的命令记录，用于审计谁操作了设备。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --actions | | 否 | string | 命令类型，逗号分隔：`propertyControlSend`/`propertyGetReportSend`/`actionSend` |
| --result-code | | 否 | int | 结果码筛选：200=成功 |
| --data-id | | 否 | string | 属性/行为标识符筛选 |
| --with-user | | 否 | bool | 包含操作人信息 |
| --time-start | | 否 | string | 开始时间戳（毫秒） |
| --time-end | | 否 | string | 结束时间戳（毫秒） |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询所有命令日志
```bash
ur things device log send -p p_smartswitch_001 -d switch-001
```

#### 示例2：查询属性控制命令（含操作人信息）
```bash
ur things device log send -p p_smartswitch_001 -d switch-001 \
  --actions propertyControlSend --with-user
```

#### 示例3：查询成功的命令
```bash
ur things device log send -p p_smartswitch_001 -d switch-001 --result-code 200
```

### 对应API

- `POST /api/v1/things/device/msg/send-log/get-list`

---

## status — 查询上下线日志

查询设备的在线/离线状态变更记录。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --status | | 否 | int | 状态筛选：1=上线，2=下线 |
| --time-start | | 否 | string | 开始时间戳（毫秒） |
| --time-end | | 否 | string | 结束时间戳（毫秒） |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询所有上下线记录
```bash
ur things device log status -p p_smartswitch_001 -d switch-001
```

#### 示例2：查询上线记录
```bash
ur things device log status -p p_smartswitch_001 -d switch-001 --status 1
```

### 对应API

- `POST /api/v1/things/device/msg/status-log/get-list`

---

## hub — 查询诊断日志（MQTT通信）

查询设备与平台之间的 MQTT 原始通信记录，是排查问题时最底层的日志。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --actions | | 否 | string | MQTT动作，逗号分隔：`connected`/`disconnected`/`property`/`event`/`action`/`thing` |
| --topics | | 否 | string | MQTT主题，逗号分隔 |
| --content | | 否 | string | 内容模糊匹配 |
| --request-id | | 否 | string | 请求令牌精确匹配 |
| --time-start | | 否 | string | 开始时间戳（毫秒） |
| --time-end | | 否 | string | 结束时间戳（毫秒） |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询所有诊断日志
```bash
ur things device log hub -p p_smartswitch_001 -d switch-001
```

#### 示例2：查询属性相关的MQTT消息
```bash
ur things device log hub -p p_smartswitch_001 -d switch-001 \
  --actions property --content report
```

#### 示例3：按请求ID精确查询
```bash
ur things device log hub -p p_smartswitch_001 -d switch-001 --request-id req_xxx
```

### 对应API

- `POST /api/v1/things/device/msg/hub-log/get-list`

---

## abnormal — 查询异常日志

查询设备异常记录，如超时、错误等。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --action | | 否 | int | 动作筛选：1=触发，2=恢复 |
| --type | | 否 | string | 异常类型 |
| --time-start | | 否 | string | 开始时间戳（毫秒） |
| --time-end | | 否 | string | 结束时间戳（毫秒） |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询所有异常日志
```bash
ur things device log abnormal -p p_smartswitch_001 -d switch-001
```

#### 示例2：查询触发的异常
```bash
ur things device log abnormal -p p_smartswitch_001 -d switch-001 --action 1
```

### 对应API

- `POST /api/v1/things/device/msg/abnormal-log/get-list`

---

## sdk — 查询SDK日志

查询设备端 SDK 上报的日志记录。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --log-level | | 否 | int | 日志级别：1=关闭，2=错误，3=警告，4=信息，5=调试 |
| --time-start | | 否 | string | 开始时间戳（毫秒） |
| --time-end | | 否 | string | 结束时间戳（毫秒） |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询所有SDK日志
```bash
ur things device log sdk -p p_smartswitch_001 -d switch-001
```

#### 示例2：查询错误级别日志
```bash
ur things device log sdk -p p_smartswitch_001 -d switch-001 --log-level 2
```

### 对应API

- `POST /api/v1/things/device/msg/sdk-log/get-list`

---

## 日志类型速查表

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

---

## 注意事项

1. **时间戳为毫秒**：`--time-start`/`--time-end` 是 Unix 毫秒时间戳
2. **默认查询近30分钟**：前端日志页默认只查最近30分钟，CLI 调用建议显式指定时间范围
3. **property 子命令两种模式**：不指定 `--data-id` 时查询最新值，指定 `--data-id` 时查询历史记录
4. **hub-log 最底层**：诊断日志是 MQTT 原始通信记录，排查问题时最彻底
5. **send-log 审计**：命令日志记录谁（account）、何时、发送了什么指令、结果如何
6. **分页从1开始**：`--page 1` 表示第一页
