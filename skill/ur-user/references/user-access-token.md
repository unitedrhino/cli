# CLI 命令 — 访问令牌

## 命令语法
```
ur user self access-token <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询访问令牌列表 |
| `get-one` | 查询访问令牌详情 |
| `create` | 创建访问令牌 |
| `update` | 更新访问令牌 |
| `delete` | 删除访问令牌 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page | | 否 | int | 页码，默认 1 |
| --size | | 否 | int | 每页数量，默认 20 |
| --json | -j | 否 | bool | 输出 JSON 格式 |
| --id | | 条件 | string | 令牌 ID（get-one/delete 必填） |
| --body | | 条件 | string | 请求体 JSON（create/update 必填） |

## 使用示例

### 示例 1：查询访问令牌列表
```bash
ur user self access-token get-list --page 1 --size 10
```

### 示例 2：查询访问令牌详情
```bash
ur user self access-token get-one --id 123456
```

### 示例 3：创建访问令牌
```bash
ur user self access-token create --body '{"name":"api-token","expiresAt":"2026-12-31"}'
```

### 示例 4：更新访问令牌
```bash
ur user self access-token update --body '{"id":123456,"name":"updated-token"}'
```

### 示例 5：删除访问令牌
```bash
ur user self access-token delete --id 123456
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| get-list | `POST /api/v1/system/user/self/access-token/get-list` |
| get-one | `POST /api/v1/system/user/self/access-token/get-one` |
| create | `POST /api/v1/system/user/self/access-token/create` |
| update | `POST /api/v1/system/user/self/access-token/update` |
| delete | `POST /api/v1/system/user/self/access-token/delete` |
