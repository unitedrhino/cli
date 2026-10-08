# CLI 命令 — 区域管理（area）

区域管理，支持区域信息查询/创建/更新/删除和区域配置管理。

---

## area info get-list — 查询区域列表

### 命令语法

```
ur things area info get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --name |  | 否 | string | 区域名称过滤 |
| --parent-id |  | 否 | string | 父区域ID过滤 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things area info get-list
```

### 对应API

- `POST /api/v1/things/area/info/get-list`

---

## area info get-one — 查询区域详情

### 命令语法

```
ur things area info get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 区域ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things area info get-one --id 12345
```

### 对应API

- `POST /api/v1/things/area/info/get-one`

---

## area info create — 创建区域

### 命令语法

```
ur things area info create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 区域信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things area info create --body '{"name":"一楼大厅","parentID":0,"desc":"主楼一层公共区域"}'
```

### 对应API

- `POST /api/v1/things/area/info/create`

---

## area info update — 更新区域

### 命令语法

```
ur things area info update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 区域信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things area info update --body '{"id":12345,"name":"一楼大厅A区"}'
```

### 对应API

- `POST /api/v1/things/area/info/update`

---

## area info delete — 删除区域

### 命令语法

```
ur things area info delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 区域ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things area info delete --id 12345
```

### 对应API

- `POST /api/v1/things/area/info/delete`

---

## area profile get-list — 查询区域配置列表

### 命令语法

```
ur things area profile get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things area profile get-list
```

### 对应API

- `POST /api/v1/things/area/profile/get-list`

---

## area profile get-one — 查询区域配置详情

### 命令语法

```
ur things area profile get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 区域ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things area profile get-one --id 12345
```

### 对应API

- `POST /api/v1/things/area/profile/get-one`

---

## area profile update — 更新区域配置

### 命令语法

```
ur things area profile update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 区域配置JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things area profile update --body '{"id":12345,"config":{"temperatureThreshold":26}}'
```

### 对应API

- `POST /api/v1/things/area/profile/update`
