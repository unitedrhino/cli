# CLI 命令 — 应用与菜单

## 命令语法

### 应用管理
```
ur user self app <subcommand> [选项]
```

### 菜单管理
```
ur user self menu <subcommand> [选项]
```

## 子命令

### 应用管理

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询应用列表 |
| `get-one` | 查询应用详情 |

### 菜单管理

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询菜单列表 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page | | 否 | int | 页码，默认 1 |
| --size | | 否 | int | 每页数量，默认 20 |
| --json | -j | 否 | bool | 输出 JSON 格式 |
| --body | | 条件 | string | 请求体 JSON（app get-one 必填） |
| --app-id | | 否 | string | 应用 ID（menu get-list 可用） |
| --is-common | | 否 | string | 是否通用菜单（menu get-list 可用） |

## 使用示例

### 示例 1：查询应用列表
```bash
ur user self app get-list --page 1 --size 10
```

### 示例 2：查询应用详情
```bash
ur user self app get-one --body '{"id":"100"}'
```

### 示例 3：查询菜单列表
```bash
ur user self menu get-list --app-id 100
```

## 对应 API

| 功能 | 子命令 | API 端点 | 权限 |
|------|--------|----------|------|
| 应用 | get-list | `POST /api/v1/system/user/self/app/get-list` | all（所有登录用户） |
| 应用 | get-one | `POST /api/v1/system/user/self/app/get-one` | all（所有登录用户） |
| 菜单 | get-list | `POST /api/v1/system/user/self/menu/get-list` | all（所有登录用户） |

## 接口区分说明

| 场景 | 使用接口 | 权限 |
|------|---------|------|
| 查询【当前用户】可访问的应用列表 | `POST /api/v1/system/user/self/app/get-list` | all |
| 查询【平台】所有应用（管理用） | `POST /api/v1/system/app/info/get-list` | platform（仅平台管理员） |

> 当用户询问【我的应用】【我有哪些应用】等涉及当前用户可用应用的场景时，**必须使用** `/api/v1/system/user/self/app/get-list`，不可使用 `/api/v1/system/app/info/get-list`（后者仅平台管理员可调用）。
