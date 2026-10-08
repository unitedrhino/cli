# CLI 命令 — 项目配置（project profile）

项目配置管理，支持查询和更新项目配置。

---

## project profile get-list — 查询项目配置列表

### 命令语法

```
ur things project profile get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project profile get-list
```

### 对应API

- `POST /api/v1/things/project/profile/get-list`

---

## project profile get-one — 查询项目配置详情

### 命令语法

```
ur things project profile get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 项目ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project profile get-one --id 12345
```

### 对应API

- `POST /api/v1/things/project/profile/get-one`

---

## project profile update — 更新项目配置

### 命令语法

```
ur things project profile update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 项目配置JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things project profile update --body '{"id":12345,"config":{"theme":"dark"}}'
```

### 对应API

- `POST /api/v1/things/project/profile/update`
