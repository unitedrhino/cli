# API 通用约定

- 全部使用 POST 方法
- 请求头: `Content-Type: application/json`, `Authorization: Bearer <jwt>`, `app-id`, `tenant-code`
- 响应格式: `{code, msg, data}`，其中 `code=200` 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
- 枚举/状态字段必须加 `json:",status,string"` tag

Swagger 中 `x-auth-type` 标注权限层级：
- `all` — 所有登录用户
- `admin` — 租户管理员 + 平台管理员
- `platform` — 仅平台管理员
