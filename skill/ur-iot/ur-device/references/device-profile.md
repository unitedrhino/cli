# CLI 命令 — 设备配置

## 命令语法

```
ur things device profile <subcommand> [选项]
```

子命令：`get-list`、`get-one`、`update`、`delete`

---

## get-list — 查询设备配置列表

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 否 | string | 产品ID筛选 |
| --device-name | -d | 否 | string | 设备名称筛选 |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device profile get-list -p product_id --page 1 --size 20
```

### 对应API

- `POST /api/v1/things/device/profile/get-list`

---

## get-one — 查询设备配置详情

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device profile get-one -p product_id -d device_name
```

### 对应API

- `POST /api/v1/things/device/profile/get-one`

---

## update — 更新设备配置

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --data | | 是 | JSON | 配置数据，JSON对象格式 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device profile update -p product_id -d device_name --data '{"key":"value"}'
```

### 对应API

- `POST /api/v1/things/device/profile/update`

---

## delete — 删除设备配置

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device profile delete -p product_id -d device_name
```

### 对应API

- `POST /api/v1/things/device/profile/delete`

---

## 注意事项

1. 设备配置用于存储设备的个性化参数，如阈值、告警规则等
2. `--data` 参数必须是 JSON 对象，键值对根据具体业务需求定义
3. 删除配置不会影响设备本身，仅清除配置数据
