# CLI 命令 — 用户管理

## 命令语法
```
ur user info <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询用户列表 |
| `get-one` | 查询用户详情 |
| `create` | 创建用户 |
| `update` | 更新用户 |
| `delete` | 删除用户 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page | | 否 | int | 页码，默认 1 |
| --size | | 否 | int | 每页数量，默认 20 |
| --json | -j | 否 | bool | 输出 JSON 格式 |
| --account | | 否 | string | 按账号筛选（get-list 可用） |
| --status | | 否 | string | 按状态筛选（get-list 可用） |
| --id | | 条件 | string | 用户 ID（get-one/delete 必填） |
| --body | | 条件 | string | 请求体 JSON（create/update 必填） |

## 使用示例

### 示例 1：查询用户列表
```bash
ur user info get-list --page 1 --size 10
```

### 示例 2：按账号筛选用户
```bash
ur user info get-list --account admin --page 1 --size 10
```

### 示例 3：查询用户详情
```bash
ur user info get-one --id 123456
```

### 示例 4：创建用户
```bash
ur user info create --body '{"account":"newuser","userName":"新用户","password":"123456"}'
```

### 示例 5：更新用户
```bash
ur user info update --body '{"id":123456,"userName":"更新后的名称"}'
```

### 示例 6：删除用户
```bash
ur user info delete --id 123456
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| get-list | `POST /api/v1/system/user/info/get-list` |
| get-one | `POST /api/v1/system/user/info/get-one` |
| create | `POST /api/v1/system/user/info/create` |
| update | `POST /api/v1/system/user/info/update` |
| delete | `POST /api/v1/system/user/info/delete` |
