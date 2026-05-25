# CLI 命令 — 告警记录（alarm record）

告警记录管理，支持查询告警记录和处理告警。

---

## alarm record get-list — 查询告警记录列表

### 命令语法

```
ur alarm record get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --alarm-id |  | 否 | string | 告警规则ID过滤 |
| --product-id | -p | 否 | string | 产品ID过滤 |
| --level |  | 否 | int | 告警等级过滤 |
| --deal-state |  | 否 | int | 处理状态过滤 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm record get-list --alarm-id 12345
```

### 对应API

- `POST /api/v1/things/rule/alarm/record/get-list`

---

## alarm record deal — 处理告警记录

### 命令语法

```
ur alarm record deal [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 处理信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm record deal --body '{"id":67890,"dealResult":"已处理","dealDesc":"检查设备运行正常"}'
```

### 对应API

- `POST /api/v1/things/rule/alarm/record/deal`
