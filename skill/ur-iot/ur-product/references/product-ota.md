# CLI 命令 — OTA 管理（ota）

OTA（Over-The-Air）固件升级管理，包含固件管理、升级任务管理和模块管理。

---

## ota firmware get-list — 查询固件列表

### 命令语法

```
ur things ota firmware get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 否 | string | 产品ID |
| --version |  | 否 | string | 固件版本号 |
| --module |  | 否 | string | 模块名称 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota firmware get-list -p p_smartswitch_001
```

### 对应API

- `POST /api/v1/things/ota/firmware/info/get-list`

---

## ota firmware get-one — 查询固件详情

### 命令语法

```
ur things ota firmware get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 固件ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota firmware get-one --id 12345
```

### 对应API

- `POST /api/v1/things/ota/firmware/info/get-one`

---

## ota firmware create — 创建固件

### 命令语法

```
ur things ota firmware create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 固件信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota firmware create --body '{"productID":"p_smartswitch_001","version":"1.0.1","name":"固件v1.0.1","url":"https://example.com/firmware.bin"}'
```

### 对应API

- `POST /api/v1/things/ota/firmware/info/create`

---

## ota firmware update — 更新固件

### 命令语法

```
ur things ota firmware update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 固件信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota firmware update --body '{"id":12345,"name":"固件v1.0.2"}'
```

### 对应API

- `POST /api/v1/things/ota/firmware/info/update`

---

## ota firmware delete — 删除固件

### 命令语法

```
ur things ota firmware delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 固件ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota firmware delete --id 12345
```

### 对应API

- `POST /api/v1/things/ota/firmware/info/delete`

---

## ota job get-list — 查询 OTA 任务列表

### 命令语法

```
ur things ota job get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 否 | string | 产品ID |
| --firmware-id |  | 否 | string | 固件ID |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota job get-list -p p_smartswitch_001
```

### 对应API

- `POST /api/v1/things/ota/firmware/job/get-list`

---

## ota job get-one — 查询 OTA 任务详情

### 命令语法

```
ur things ota job get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 任务ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota job get-one --id 67890
```

### 对应API

- `POST /api/v1/things/ota/firmware/job/get-one`

---

## ota job create — 创建 OTA 任务

### 命令语法

```
ur things ota job create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 任务信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota job create --body '{"productID":"p_smartswitch_001","firmwareID":12345,"deviceNames":["switch-001","switch-002"]}'
```

### 对应API

- `POST /api/v1/things/ota/firmware/job/create`

---

## ota job update — 更新 OTA 任务

### 命令语法

```
ur things ota job update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 任务信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota job update --body '{"id":67890,"status":2}'
```

### 对应API

- `POST /api/v1/things/ota/firmware/job/update`

---

## ota module get-list — 查询模块列表

### 命令语法

```
ur things ota module get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 否 | string | 产品ID |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota module get-list -p p_smartswitch_001
```

### 对应API

- `POST /api/v1/things/ota/module/info/get-list`

---

## ota module get-one — 查询模块详情

### 命令语法

```
ur things ota module get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 模块ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota module get-one --id 11111
```

### 对应API

- `POST /api/v1/things/ota/module/info/get-one`

---

## ota module create — 创建模块

### 命令语法

```
ur things ota module create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 模块信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota module create --body '{"productID":"p_smartswitch_001","name":"主控模块","code":"main"}'
```

### 对应API

- `POST /api/v1/things/ota/module/info/create`

---

## ota module update — 更新模块

### 命令语法

```
ur things ota module update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 模块信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota module update --body '{"id":11111,"name":"主控模块v2"}'
```

### 对应API

- `POST /api/v1/things/ota/module/info/update`

---

## ota module delete — 删除模块

### 命令语法

```
ur things ota module delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 模块ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things ota module delete --id 11111
```

### 对应API

- `POST /api/v1/things/ota/module/info/delete`

---

## 注意事项

- OTA 进度码：1-100 表示百分比进度，-1=下载失败，-2=校验失败，-3=烧录失败，-4=版本不匹配
- 创建固件时需要提供固件包的下载 URL
- 创建 OTA 任务时需指定目标设备列表
