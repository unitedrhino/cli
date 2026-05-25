# CLI 命令 — 项目管理（project info）

项目信息管理，支持查询、创建、更新、删除项目。

---

## project info get-list — 查询项目列表

### 命令语法

```
ur things project info get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --name |  | 否 | string | 项目名称过滤 |
| --area-id |  | 否 | string | 区域ID过滤 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project info get-list
```

### 对应API

- `POST /api/v1/things/project/info/get-list`

---

## project info get-one — 查询项目详情

### 命令语法

```
ur things project info get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 项目ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project info get-one --id 12345
```

### 对应API

- `POST /api/v1/things/project/info/get-one`

---

## project info create — 创建项目

### 命令语法

```
ur things project info create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 项目信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project info create --body '{"name":"智能楼宇项目","desc":"某写字楼智能化改造"}'
```

### 对应API

- `POST /api/v1/things/project/info/create`

---

## project info update — 更新项目

### 命令语法

```
ur things project info update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 项目信息JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project info update --body '{"id":12345,"name":"智能楼宇项目v2"}'
```

### 对应API

- `POST /api/v1/things/project/info/update`

---

## project info delete — 删除项目

### 命令语法

```
ur things project info delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 项目ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project info delete --id 12345
```

### 对应API

- `POST /api/v1/things/project/info/delete`
