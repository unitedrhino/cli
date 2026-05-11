## API 通用约定

### 请求格式

 所有 API 使用 **POST** 方法，禁止使用 GET/PUT/DELETE。

请求头:
 Content-Type: application/json
 Authorization: Bearer <jwt>
 app-id: <appID>
 tenant-code: <tenantCode>

请求体: JSON 格式。

### 响应格式

 `{code, msg, data}`

 code=200 表示成功。

### 分页格式

 请求: `{page: {page: 1, size: 10}}`
  响应: `{page: {page: 1, size: 10}, total: 0, list: [...]}`

  page 从 1 开始， size >= 1.

### 权限层级

 所有 API 端点在 swagger JSON 中标注 `x-auth-type`:

 - `all` — 所有登录用户可访问
 - `admin` — 租户管理员及以上
 - `platform` — 仅平台管理员

### 字段命名注意

 `deviceName` = 设备ID（不是显示名称）， `deviceAlias` = 设备显示名称/别名

 `productID` = 产品ID（不是产品名称)
 `protocolCode` = 协议代码（默认 `urMqtt`）
 `userID` 在 JWT 中必须为字符串格式 `"userID": "12345"` 而数字
### 常见错误码

 | code | 含义 | 处理 |
 |------|------|---------|
 | 200 | 成功 | - |
 | 400 | 参数错误 | 检查请求参数 |
 | 401 | 未认证失败 | JWT 无效或过期， 检查 accessKey/Secret |
 | 403 | 权限不足 | 需要更高权限用户 | | 500 | 服务器错误 | 联系管理员 |
