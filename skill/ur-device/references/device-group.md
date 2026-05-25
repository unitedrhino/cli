# CLI 命令 — 设备分组

## 命令语法

```
ur things device group <subcommand> [选项]
```

子命令：`get-list`、`create`、`delete`、`batch-create-device`、`batch-delete-device`

---

## get-list — 查询设备分组列表

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --name | | 否 | string | 按分组名称筛选 |
| --group-id | | 否 | string | 按分组ID筛选 |
| --parent-id | | 否 | string | 按父分组ID筛选 |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device group get-list --page 1 --size 20
```

### 对应API

- `POST /api/v1/things/group/info/get-list`

---

## create — 创建设备分组

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --name | | 是 | string | 分组名称 |
| --desc | | 否 | string | 分组描述 |
| --parent-id | | 否 | string | 父分组ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device group create --name "My Group" --desc "Description"
```

### 对应API

- `POST /api/v1/things/group/info/create`

---

## delete — 删除设备分组

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --group-id | | 是 | string | 分组ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device group delete --group-id 123
```

### 对应API

- `POST /api/v1/things/group/info/delete`

---

## batch-create-device — 批量添加设备到分组

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --group-id | | 是 | string | 分组ID |
| --devices | | 是 | JSON | 设备列表，JSON数组格式 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device group batch-create-device --group-id 123 \
  --devices '[{"productID":"pid","deviceName":"dev1"}]'
```

### 对应API

- `POST /api/v1/things/group/device/batch-create`

---

## batch-delete-device — 批量从分组删除设备

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --group-id | | 是 | string | 分组ID |
| --devices | | 是 | JSON | 设备列表，JSON数组格式 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device group batch-delete-device --group-id 123 \
  --devices '[{"productID":"pid","deviceName":"dev1"}]'
```

### 对应API

- `POST /api/v1/things/group/device/batch-delete`

---

## 注意事项

1. 分组支持层级结构，通过 `--parent-id` 创建子分组
2. `--devices` 参数必须是 JSON 数组，每个元素包含 `productID` 和 `deviceName`
3. 删除分组不会删除分组内的设备，仅解除分组关系
4. 设备可以同时属于多个分组
