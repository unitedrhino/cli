# CLI 命令 — 告警通知记录（alarm notify-record）

新版告警通知记录管理（前端告警中心），支持查询与手动重发。

---

## alarm notify-record get-list — 查询通知记录列表

### 命令语法

```
ur alarm notify-record get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --event-id |  | 否 | string | 告警事件ID过滤 |
| --alarm-id |  | 否 | string | 告警规则ID过滤 |
| --alarm-name |  | 否 | string | 告警规则名称过滤 |
| --channel |  | 否 | string | 通知渠道过滤，逗号分隔（如 callback,webhook） |
| --status |  | 否 | string | 发送状态过滤，逗号分隔（如 success,fail） |
| --target-name |  | 否 | string | 通知目标名称过滤 |
| --timing |  | 否 | string | 通知时机过滤（如 trigger/recovery） |
| --start |  | 否 | int | 创建时间范围起（毫秒时间戳） |
| --end |  | 否 | int | 创建时间范围止（毫秒时间戳） |
| --project-id |  | 否 | string | 项目ID（作为 project-id 请求头传递） |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm notify-record get-list --event-id 5 --size 10 --project-id 365369582443328
ur alarm notify-record get-list --status fail
```

### 对应API

- `POST /api/v1/things/alarm/notify-record/get-list`

---

## alarm notify-record resend — 重发通知

### 命令语法

```
ur alarm notify-record resend [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 通知记录ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm notify-record resend --id 3
```

### 对应API

- `POST /api/v1/things/alarm/notify-record/resend`
