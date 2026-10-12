# API 通用约定

- 全部使用 POST 方法
- 请求头: `Content-Type: application/json`, `Authorization: Bearer <jwt>`, `app-id`, `tenant-code`
- 响应格式: `{code, msg, data}`，其中 `code=200` 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
- 枚举/状态字段必须加 `json:",status,string"` tag

Swagger 中 `x-auth-type` 标注权限层级：
- `all` — 所有登录用户
- `admin` — 企业管理员 + 平台管理员
- `platform` — 仅平台管理员


## 数量 / 列表类问题

- 如果用户问"有几个""多少个""当前有哪些"这类业务数据问题，优先直接调用真实的 `get-list` 接口。
- 数量统一优先从返回的 `data.total` 读取；不要自己数当前页 `list.length`，除非接口根本没有 `total`。
- 不要先调用 `schema` 再决定是否查询数据；`schema` 只用于查看已知接口的字段结构，不用于查询业务结果。
