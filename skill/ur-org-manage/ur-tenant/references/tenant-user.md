# CLI 命令 — 企业用户管理

## 命令语法
```
ur tenant user <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询企业用户列表 |
| `get-one` | 查询企业用户详情 |
| `batch-create` | 批量创建企业用户 |
| `update` | 更新企业用户 |
| `delete` | 删除企业用户 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page | | 否 | int | 页码，默认 1 |
| --size | | 否 | int | 每页数量，默认 20 |
| --json | -j | 否 | bool | 输出 JSON 格式 |
| --account | | 否 | string | 按账号筛选（get-list 可用） |
| --status | | 否 | string | 按状态筛选（get-list 可用） |
| --id | | 条件 | string | 用户 ID（get-one/delete 必填） |
| --body | | 条件 | string | 请求体 JSON（batch-create/update 必填） |

## 使用示例

### 示例 1：查询企业用户列表
```bash
ur tenant user get-list --page 1 --size 10
```

### 示例 2：按账号筛选企业用户
```bash
ur tenant user get-list --account admin --page 1 --size 10
```

### 示例 3：查询企业用户详情
```bash
ur tenant user get-one --id 123456
```

### 示例 4：批量创建企业用户
```bash
ur tenant user batch-create --body '{"users":[{"account":"user1","roleID":1}]}'
```

### 示例 5：更新企业用户
```bash
ur tenant user update --body '{"id":123456,"status":1}'
```

### 示例 6：删除企业用户
```bash
ur tenant user delete --id 123456
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| get-list | `POST /api/v1/system/tenant/user/get-list` |
| get-one | `POST /api/v1/system/tenant/user/get-one` |
| batch-create | `POST /api/v1/system/tenant/user/batch-create` |
| update | `POST /api/v1/system/tenant/user/update` |
| delete | `POST /api/v1/system/tenant/user/delete` |
