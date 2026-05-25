# CLI 命令 — 角色管理

## 命令语法
```
ur tenant user role <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询用户角色列表 |
| `batch-update` | 批量更新用户角色 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page | | 否 | int | 页码，默认 1 |
| --size | | 否 | int | 每页数量，默认 20 |
| --json | -j | 否 | bool | 输出 JSON 格式 |
| --user-id | | 否 | string | 按用户 ID 筛选（get-list 可用） |
| --body | | 条件 | string | 请求体 JSON（batch-update 必填） |

## 使用示例

### 示例 1：查询用户角色列表
```bash
ur tenant user role get-list --page 1 --size 10
```

### 示例 2：按用户 ID 查询角色
```bash
ur tenant user role get-list --user-id 123456
```

### 示例 3：批量更新用户角色
```bash
ur tenant user role batch-update --body '{"userID":123,"roleIDs":[1,2,3]}'
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| get-list | `POST /api/v1/system/tenant/user/role/get-list` |
| batch-update | `POST /api/v1/system/tenant/user/role/batch-update` |
