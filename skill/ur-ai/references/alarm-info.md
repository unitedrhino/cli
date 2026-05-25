# CLI 命令 — 告警规则（alarm info）

告警规则管理，支持查询、创建、更新、删除告警规则。

---

## alarm info get-list — 查询告警规则列表

### 命令语法

```
ur alarm info get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 否 | string | 产品ID过滤 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm info get-list -p p_smartswitch_001
```

### 对应API

- `POST /api/v1/things/rule/alarm/info/get-list`

---

## alarm info get-one — 查询告警规则详情

### 命令语法

```
ur alarm info get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警规则ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm info get-one --id 12345
```

### 对应API

- `POST /api/v1/things/rule/alarm/info/get-one`

---

## alarm info create — 创建告警规则

### 命令语法

```
ur alarm info create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 告警规则JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm info create --body '{"productID":"p_smartswitch_001","name":"温度过高告警","level":1,"trigger":{"type":"device","device":{"productID":"p_smartswitch_001","selectType":"fixed","deviceName":"switch-001","type":"propertyReport","dataID":"Temperature","termType":"gt","values":["50"]}}}'
```

### 对应API

- `POST /api/v1/things/rule/alarm/info/create`

---

## alarm info update — 更新告警规则

### 命令语法

```
ur alarm info update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 告警规则JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm info update --body '{"id":12345,"name":"温度过高告警v2","level":2}'
```

### 对应API

- `POST /api/v1/things/rule/alarm/info/update`

---

## alarm info delete — 删除告警规则

### 命令语法

```
ur alarm info delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警规则ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm info delete --id 12345
```

### 对应API

- `POST /api/v1/things/rule/alarm/info/delete`
