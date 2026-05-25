# CLI 命令 — 告警场景（alarm scene）

告警场景管理，支持查询、批量创建和删除告警场景。

---

## alarm scene get-list — 查询告警场景列表

### 命令语法

```
ur alarm scene get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --alarm-id |  | 否 | string | 告警规则ID过滤 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm scene get-list --alarm-id 12345
```

### 对应API

- `POST /api/v1/things/rule/alarm/scene/get-list`

---

## alarm scene batch-create — 批量创建告警场景

### 命令语法

```
ur alarm scene batch-create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 告警场景列表JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm scene batch-create --body '{"alarmID":12345,"scenes":[{"sceneID":111,"mode":"trigger"},{"sceneID":112,"mode":"relieve"}]}'
```

### 对应API

- `POST /api/v1/things/rule/alarm/scene/batch-create`

---

## alarm scene delete — 删除告警场景

### 命令语法

```
ur alarm scene delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 告警场景ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur alarm scene delete --id 12345
```

### 对应API

- `POST /api/v1/things/rule/alarm/scene/delete`
