# CLI 命令 — 设备信息管理

## 命令语法

```
ur things device info <subcommand> [选项]
```

子命令：`get-list`、`get-one`、`create`、`update`、`delete`、`bind`、`unbind`、`count`

---

## get-list — 查询设备列表

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 否 | string | 产品ID，筛选特定产品下的设备 |
| --device-name | -d | 否 | string | 设备名称，精确匹配 |
| --group-id | | 否 | string | 分组ID，筛选特定分组内的设备 |
| --tag | | 否 | string | 标签筛选 |
| --area-id | | 否 | string | 区域ID筛选 |
| --project-id | | 否 | string | 项目ID筛选 |
| --page | | 否 | int | 页码（默认1） |
| --size | | 否 | int | 每页数量（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询所有设备列表
```bash
ur things device info get-list --page 1 --size 10
```

#### 示例2：查询特定产品下的设备
```bash
ur things device info get-list -p p_smartswitch_001 --page 1 --size 20
```

#### 示例3：按区域筛选设备
```bash
ur things device info get-list --area-id area_001 --page 1 --size 50
```

### 对应API

- `POST /api/v1/things/device/info/get-list`

---

## get-one — 查询设备详情

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称（设备唯一ID） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device info get-one -p p_smartswitch_001 -d switch-001
```

### 对应API

- `POST /api/v1/things/device/info/get-one`

---

## create — 创建设备

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称（设备唯一ID） |
| --alias | | 否 | string | 设备别名/显示名称 |
| --tags | | 否 | JSON | 设备标签数组，如 `[{"key":"loc","value":"room1"}]` |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：创建设备
```bash
ur things device info create -p p_smartswitch_001 -d switch-001 --alias "一楼走廊开关"
```

#### 示例2：创建设备并添加标签
```bash
ur things device info create -p p_smartswitch_001 -d switch-002 --alias "二楼开关" --tags '[{"key":"floor","value":"2"}]'
```

### 对应API

- `POST /api/v1/things/device/info/create`

---

## update — 更新设备

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称（设备唯一ID） |
| --alias | | 否 | string | 设备别名/显示名称 |
| --tags | | 否 | JSON | 设备标签数组 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device info update -p p_smartswitch_001 -d switch-001 --alias "新名称"
```

### 对应API

- `POST /api/v1/things/device/info/update`

---

## delete — 删除设备

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称（设备唯一ID） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device info delete -p p_smartswitch_001 -d switch-001
```

### 对应API

- `POST /api/v1/things/device/info/delete`

---

## bind — 绑定设备到网关

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 子设备产品ID |
| --device-name | -d | 是 | string | 子设备名称 |
| --gateway-id | | 否 | string | 网关产品ID |
| --gateway-name | | 否 | string | 网关设备名称 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device info bind -p sub_pid -d sub_001 --gateway-id gateway_pid --gateway-name gateway_001
```

### 对应API

- `POST /api/v1/things/device/info/bind`

---

## unbind — 解绑设备

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称（设备唯一ID） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device info unbind -p p_smartswitch_001 -d switch-001
```

### 对应API

- `POST /api/v1/things/device/info/unbind`

---

## count — 统计设备数量

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 否 | string | 按产品ID筛选统计 |
| --project-id | | 否 | string | 按项目ID筛选统计 |
| --area-id | | 否 | string | 按区域ID筛选统计 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：统计全部设备数量
```bash
ur things device info count
```

#### 示例2：统计某产品下的设备数量
```bash
ur things device info count -p p_smartswitch_001
```

### 对应API

- `POST /api/v1/things/device/info/count`

---

## 注意事项

1. **deviceName 是设备唯一ID**，不是显示名称；显示名称使用 `--alias` 设置
2. 分页从 1 开始，`--page 1` 表示第一页
3. `get-list` 支持多维度筛选，可同时使用 `--product-id`、`--area-id`、`--project-id` 等组合条件
4. 绑定/解绑操作影响网关拓扑关系，子设备通过网关间接接入平台
