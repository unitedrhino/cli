# CLI 命令 — 项目分组（project group）

项目设备分组管理，支持分组信息管理和设备批量添加。

---

## project group info get-list — 查询项目分组列表

### 命令语法

```
ur things project group info get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --name |  | 否 | string | 分组名称过滤 |
| --group-id |  | 否 | string | 分组ID过滤 |
| --parent-id |  | 否 | string | 父分组ID过滤 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project group info get-list
```

### 对应API

- `POST /api/v1/things/group/info/get-list`

---

## project group info create — 创建项目分组

### 命令语法

```
ur things project group info create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --name |  | 是 | string | 分组名称 |
| --desc |  | 否 | string | 分组描述 |
| --parent-id |  | 否 | string | 父分组ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project group info create --name "一楼设备" --desc "一楼所有智能设备"
```

### 对应API

- `POST /api/v1/things/group/info/create`

---

## project group info delete — 删除项目分组

### 命令语法

```
ur things project group info delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --group-id |  | 是 | string | 分组ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project group info delete --group-id 12345
```

### 对应API

- `POST /api/v1/things/group/info/delete`

---

## project group device batch-create — 批量添加设备到分组

### 命令语法

```
ur things project group device batch-create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --group-id |  | 是 | string | 分组ID |
| --devices |  | 是 | JSON | 设备列表JSON数组 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project group device batch-create --group-id 12345 --devices '[{"productID":"p_smartswitch_001","deviceName":"switch-001"},{"productID":"p_smartswitch_001","deviceName":"switch-002"}]'
```

### 对应API

- `POST /api/v1/things/group/device/batch-create`
