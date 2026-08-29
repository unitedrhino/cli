# CLI 命令 — 告警规则（alarm rule）

新版告警规则管理（前端告警中心），支持查询、创建、更新、删除、启停、手动评估。

---

## alarm rule get-list — 查询告警规则列表

### 命令语法

```
ur alarm rule get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --name |  | 否 | string | 规则名称关键字过滤（keyword） |
| --product-id |  | 否 | string | 产品ID过滤 |
| --status |  | 否 | int | 状态过滤（1=启用 0=停用） |
| --level |  | 否 | string | 告警级别过滤，逗号分隔多个 |
| --start |  | 否 | int | 创建时间范围起（毫秒时间戳） |
| --end |  | 否 | int | 创建时间范围止（毫秒时间戳） |
| --project-id |  | 否 | string | 项目ID（作为 project-id 请求头传递，规则按项目隔离时建议带上） |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm rule get-list --product-id 66 --size 10 --project-id 365369582443328
```

### 对应API

- `POST /api/v1/things/alarm/info/get-list`

---

## alarm rule get-one — 查询告警规则详情

### 命令语法

```
ur alarm rule get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警规则ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm rule get-one --id 1
```

### 对应API

- `POST /api/v1/things/alarm/info/get-one`

---

## alarm rule create — 创建告警规则

### 命令语法

```
ur alarm rule create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 告警规则JSON（name/target/if 等字段以 swagger AlarmInfoCreateReq 为准） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm rule create --body '{"name":"温度过高告警","target":{"selectType":"all","productID":"66"},"if":{"triggers":[{"id":"1","type":"and","level":"major","criteria":[{"id":"1","type":"plain","order":1,"plain":{"leaves":[{"order":0,"dataID":"Temperature","sourceType":1,"compare":{"termType":"gt","values":["50"]},"period":60,"aggregateMethod":"last","continuePeriod":1}]}}]}]}}'
```

### 对应API

- `POST /api/v1/things/alarm/info/create`

---

## alarm rule update — 更新告警规则

### 命令语法

```
ur alarm rule update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 告警规则JSON（含 id） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm rule update --body '{"id":"1","name":"温度过高告警v2","desc":"阈值调整"}'
```

### 对应API

- `POST /api/v1/things/alarm/info/update`

---

## alarm rule delete — 删除告警规则

### 命令语法

```
ur alarm rule delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警规则ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm rule delete --id 1
```

### 对应API

- `POST /api/v1/things/alarm/info/delete`

---

## alarm rule status-update — 启用/停用告警规则

### 命令语法

```
ur alarm rule status-update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警规则ID |
| --status |  | 是 | int | 目标状态（1=启用 0=停用） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm rule status-update --id 1 --status 0
```

### 对应API

- `POST /api/v1/things/alarm/info/status-update`

---

## alarm rule evaluate-trigger — 手动评估告警规则

### 命令语法

```
ur alarm rule evaluate-trigger [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | int（可重复或逗号分隔） | 要评估的告警规则ID（alarmIDs 数组） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm rule evaluate-trigger --id 1
ur alarm rule evaluate-trigger --id 1,2
```

### 对应API

- `POST /api/v1/things/alarm/info/evaluate-trigger`
