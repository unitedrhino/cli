# CLI 命令 — 物模型管理

## 命令语法

```
ur things schema <subcommand> [选项]
```

子命令：`get-list`、`create`、`update`、`delete`、`tsl-read`、`tsl-import`、`browse`

---

## get-list — 查询物模型列表

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称（查询设备级物模型时使用） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：查询产品级物模型
```bash
ur things schema get-list -p p_smartswitch_001
```

#### 示例2：查询设备级物模型
```bash
ur things schema get-list -p p_smartswitch_001 -d switch-001
```

### 对应API

- 产品级：`POST /api/v1/things/product/schema/get-list`
- 设备级：`POST /api/v1/things/device/schema/get-list`

---

## create — 创建物模型

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称（创建设备级物模型时使用） |
| --schema | | 是 | JSON | 物模型定义 JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema create -p p_smartswitch_001 \
  --schema '{"identifier":"Temperature","name":"温度","mode":"r","define":{"type":"float","unit":"℃"}}'
```

### 对应API

- `POST /api/v1/things/product/schema/create`

---

## update — 更新物模型

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称 |
| --schema | | 是 | JSON | 物模型定义 JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema update -p p_smartswitch_001 \
  --schema '{"identifier":"Temperature","name":"温度","mode":"rw","define":{"type":"float","unit":"℃"}}'
```

### 对应API

- `POST /api/v1/things/product/schema/update`

---

## delete — 删除物模型

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --data-id | | 是 | string | 物模型数据标识符 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema delete -p p_smartswitch_001 --data-id Temperature
```

### 对应API

- `POST /api/v1/things/product/schema/delete`

---

## tsl-read — 读取 TSL

读取产品的 TSL（Thing Specification Language）定义。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称（读取设备级TSL时使用） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：读取产品TSL
```bash
ur things schema tsl-read -p p_smartswitch_001
```

#### 示例2：读取设备TSL
```bash
ur things schema tsl-read -p p_smartswitch_001 -d switch-001
```

### 对应API

- 产品级：`POST /api/v1/things/product/schema/tsl-read`
- 设备级：`POST /api/v1/things/device/schema/tsl-read`

---

## tsl-import — 导入 TSL

导入 TSL（Thing Specification Language）定义到产品。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --tsl | | 是 | JSON | TSL JSON 数据 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things schema tsl-import -p p_smartswitch_001 --tsl '{"properties":[{"identifier":"Temperature"}]}'
```

### 对应API

- `POST /api/v1/things/product/schema/tsl-import`

---

## browse — 浏览 API 端点

浏览 Swagger API 端点列表，支持按路径和认证类型过滤。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| path | | 否 | string | API 路径过滤，如 `/api/v1/things/device/info/...` |
| --json | -j | 否 | bool | 输出JSON格式 |
| --auth-type | | 否 | string | 认证类型过滤 |

### 使用示例

#### 示例1：浏览所有 API
```bash
ur things schema browse
```

#### 示例2：按路径过滤
```bash
ur things schema browse /api/v1/things/device/info/...
```

#### 示例3：JSON 格式输出
```bash
ur things schema browse /api/v1/things/device/info/... -j
```

---

## 注意事项

1. **identifier 大小写敏感**：物模型标识符使用大驼峰命名（如 `CurrentTemperature`）
2. **产品级 vs 设备级**：不指定 `--device-name` 时操作产品级物模型，指定时操作设备级物模型
3. **TSL 导入会覆盖**：导入 TSL 会覆盖产品原有物模型定义，请谨慎操作
4. **删除前确认**：删除物模型会影响依赖该物模型的设备数据解析，请确认无影响后再执行
5. `browse` 默认行为：不带子命令时默认执行 browse，兼容旧版行为
