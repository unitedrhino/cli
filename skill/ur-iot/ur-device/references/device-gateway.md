# CLI 命令 — 网关管理

## 命令语法

```
ur things device gateway <subcommand> [选项]
```

子命令：`get-list`、`batch-create`、`batch-delete`

---

## get-list — 查询网关子设备列表

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 网关产品ID |
| --device-name | -d | 是 | string | 网关设备名称 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device gateway get-list -p gateway_pid -d gateway_001
```

### 对应API

- `POST /api/v1/things/device/gateway/get-list`

---

## batch-create — 批量添加子设备到网关

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 网关产品ID |
| --device-name | -d | 是 | string | 网关设备名称 |
| --devices | | 是 | JSON | 子设备列表，JSON数组格式 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device gateway batch-create -p gateway_pid -d gateway_001 \
  --devices '[{"productID":"sub_pid","deviceName":"sub_001"}]'
```

### 对应API

- `POST /api/v1/things/device/gateway/batch-create`

---

## batch-delete — 批量从网关删除子设备

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 网关产品ID |
| --device-name | -d | 是 | string | 网关设备名称 |
| --devices | | 是 | JSON | 子设备列表，JSON数组格式 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device gateway batch-delete -p gateway_pid -d gateway_001 \
  --devices '[{"productID":"sub_pid","deviceName":"sub_001"}]'
```

### 对应API

- `POST /api/v1/things/device/gateway/batch-delete`

---

## 注意事项

1. 网关设备本身也是普通设备，需先通过 `ur things device info create` 创建
2. `--devices` 参数必须是 JSON 数组，每个元素包含 `productID` 和 `deviceName`
3. 批量操作前建议先查询网关当前子设备列表，确认目标设备状态
4. 子设备通过网关间接接入平台，网关离线时子设备也无法通信
