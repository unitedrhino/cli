# CLI 命令 — 告警事件（alarm event）

新版告警事件管理（前端告警中心），支持查询、详情、处理、误报标记、统计。

---

## alarm event get-list — 查询告警事件列表

### 命令语法

```
ur alarm event get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --rule-id |  | 否 | string | 告警规则ID过滤（alarmID） |
| --product-id |  | 否 | string | 产品ID过滤 |
| --device-name |  | 否 | string | 设备名称过滤 |
| --status |  | 否 | string | 事件状态过滤，逗号分隔（如 alarmActive,normal,recovered,acked） |
| --level |  | 否 | string | 告警级别过滤，逗号分隔（如 major） |
| --keyword |  | 否 | string | 关键字过滤 |
| --is-recovered |  | 否 | bool | 开关式 flag，仅查已恢复事件 |
| --is-false-alarm |  | 否 | bool | 开关式 flag，仅查误报事件 |
| --start |  | 否 | int | 触发时间范围起（毫秒时间戳） |
| --end |  | 否 | int | 触发时间范围止（毫秒时间戳） |
| --project-id |  | 否 | string | 项目ID（作为 project-id 请求头传递） |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm event get-list --rule-id 1 --size 5 --project-id 365369582443328
ur alarm event get-list --device-name meter-ai-001 --status alarmActive
```

### 对应API

- `POST /api/v1/things/alarm/event/get-list`

---

## alarm event get-one — 查询告警事件详情

### 命令语法

```
ur alarm event get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警事件ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm event get-one --id 5
```

### 对应API

- `POST /api/v1/things/alarm/event/get-one`

---

## alarm event deal — 处理告警事件

### 命令语法

```
ur alarm event deal [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警事件ID |
| --action |  | 是 | string | 处理动作（如 acked，取值以后端枚举为准） |
| --remark |  | 否 | string | 处理备注 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm event deal --id 5 --action acked --remark "已现场确认"
```

### 对应API

- `POST /api/v1/things/alarm/event/deal`

---

## alarm event false-alarm — 标记误报

### 命令语法

```
ur alarm event false-alarm [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警事件ID |
| --reason |  | 否 | string | 误报原因 |
| --reason-type |  | 否 | string | 误报原因类型 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm event false-alarm --id 5 --reason "传感器故障" --reason-type device
```

### 对应API

- `POST /api/v1/things/alarm/event/false-alarm`

---

## alarm event stat — 告警事件统计

### 命令语法

```
ur alarm event stat [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --group-by |  | 否 | string | 统计维度（取值以后端枚举为准） |
| --status |  | 否 | string | 状态过滤，逗号分隔 |
| --level |  | 否 | string | 级别过滤，逗号分隔 |
| --start |  | 否 | int | 触发时间范围起（毫秒时间戳） |
| --end |  | 否 | int | 触发时间范围止（毫秒时间戳） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm event stat --group-by status
```

### 对应API

- `POST /api/v1/things/alarm/event/stat`
