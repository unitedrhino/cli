# CLI 命令 — 物模型管理（schema）

物模型（Thing Model/Schema）管理，支持查询、创建、更新、删除物模型，以及 TSL 导入/读取和 API 端点浏览。

---

## schema get-list — 查询物模型列表

### 命令语法

```
ur things schema get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称（指定则查询设备级物模型） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询产品物模型
```bash
ur things schema get-list -p p_smartswitch_001
```

#### 示例2：查询设备物模型
```bash
ur things schema get-list -p p_smartswitch_001 -d switch-001
```

#### 示例3：JSON输出
```bash
ur things schema get-list -p p_smartswitch_001 -j
```

### 对应API

- 产品级：`POST /api/v1/things/product/schema/get-list`
- 设备级：`POST /api/v1/things/device/schema/get-list`

---

## schema create — 创建物模型

### 命令语法

```
ur things schema create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称（指定则创建设备级物模型） |
| --schema |  | 是 | JSON | 物模型定义JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema create -p p_smartswitch_001 --schema '{"identifier":"PowerSwitch","name":"电源开关","mode":"rw","define":{"type":"bool"}}'
```

### 对应API

- `POST /api/v1/things/product/schema/create`

---

## schema update — 更新物模型

### 命令语法

```
ur things schema update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称 |
| --schema |  | 是 | JSON | 物模型定义JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema update -p p_smartswitch_001 --schema '{"identifier":"PowerSwitch","name":"电源开关","mode":"rw","define":{"type":"bool"}}'
```

### 对应API

- `POST /api/v1/things/product/schema/update`

---

## schema delete — 删除物模型

### 命令语法

```
ur things schema delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --data-id |  | 是 | string | 物模型数据标识符 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema delete -p p_smartswitch_001 --data-id PowerSwitch
```

### 对应API

- `POST /api/v1/things/product/schema/delete`

---

## schema tsl-read — 读取 TSL

### 命令语法

```
ur things schema tsl-read [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema tsl-read -p p_smartswitch_001
```

### 对应API

- 产品级：`POST /api/v1/things/product/schema/tsl-read`
- 设备级：`POST /api/v1/things/device/schema/tsl-read`

---

## schema tsl-import — 导入 TSL

### 命令语法

```
ur things schema tsl-import [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --tsl |  | 是 | JSON | TSL JSON数据 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema tsl-import -p p_smartswitch_001 --tsl '{"properties":[{"identifier":"Temperature","name":"温度","mode":"r","define":{"type":"float","min":"-40","max":"80","step":"0.1","unit":"°C"}}]}'
```

### 对应API

- `POST /api/v1/things/product/schema/tsl-import`

---

## schema browse — 浏览 API 端点

### 命令语法

```
ur things schema browse [路径] [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| path |  | 否 | string | API路径过滤（如 /api/v1/things/device） |
| --auth-type |  | 否 | string | 认证类型过滤 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：浏览所有端点
```bash
ur things schema browse
```

#### 示例2：按路径过滤
```bash
ur things schema browse /api/v1/things/device
```

#### 示例3：JSON输出
```bash
ur things schema browse /api/v1/things/device -j
```

### 说明

- 从本地 swagger 文档加载 API 端点列表
- 默认按当前应用上下文过滤认证类型
- 显式指定 `--auth-type` 可覆盖默认过滤
