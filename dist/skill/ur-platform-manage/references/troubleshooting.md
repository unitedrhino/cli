---
name: ur-api-troubleshooting
description: "ur-api 常见问题排查指南：登录失败、权限错误、设备控制异常、API 调用问题。triggers: 登录失败, 401错误, 403错误, 权限不足, 设备控制失败, API错误, token过期"
---

# ur-api 常见问题排查

## 一、登录与认证问题

### 1.1 登录失败 - 租户校验错误

**错误示例**：
```json
{"code":100006,"msg":"参数错误","detail":"tenantCode not eq uc:xxx t:yyy"}
```

**原因**：请求头 `tenant-code` 与用户所属租户不匹配。

**解决方案**：

| 场景 | 正确做法 |
|------|---------|
| 用户属于多个租户 | 选择正确的 `tenant-code` 或不传 |
| 用户不属于该租户 | 切换到用户所属的租户 |
| 不确定用户属于哪些租户 | 先不带 `tenant-code` 登录，查看 `userInfo.tenants` |

**登录租户校验规则**：

| 场景 | 行为 |
|------|------|
| 用户属于租户，租户已绑定 App | ✅ 登录成功 |
| 用户属于租户，租户**未**绑定 App | ❌ "该企业未绑定该应用，无法登录" |
| 用户不属于租户，App 是 `client` 类型 | 自动注册到 App 所属租户 |
| 用户不属于租户，App 是 `admin` 类型 | ❌ 报错 |
| 不带租户登录，用户有绑定该 App 的租户 | ✅ 登录成功 |
| 不带租户登录，用户无绑定该 App 的租户（admin 应用） | ❌ "您没有权限访问该应用" |

### 1.2 密码错误

**错误示例**：
```json
{"code":100001,"msg":"账号或密码错误"}
```

**排查步骤**：
1. 确认账号密码正确（默认 `administrator` / `iThings666`）
2. 检查 `pwdType` 参数：`1`=明文密码，`2`=MD5 加密后

### 1.3 Token 过期

**错误示例**：
```json
{"code":1000007,"msg":"登录失效"}
```

**解决方案**：
```bash
# 重新登录
ur login

# 或使用 --profile 自动刷新
ur api /api/v1/... --profile prod-admin
```

---

## 二、权限问题

### 2.1 403 权限不足

**错误示例**：
```json
{"code":100003,"msg":"权限不足"}
```

**排查步骤**：

1. **确认当前角色**：
   ```bash
   ur check
   ```

2. **检查 API 权限要求**：
   ```bash
   ur schema /api/v1/xxx --json | jq '.authType'
   ```

3. **角色与权限对照**：

| 角色 | 可调用 authType |
|------|----------------|
| 平台管理员 | platform, admin, all |
| 租户管理员 | admin, all |
| 普通用户 | all |

### 2.2 数据权限不足

**错误示例**：
```json
{"code":100003,"msg":"数据权限不足"}
```

**原因**：用户没有该项目/区域的访问权限。

**解决方案**：
1. 检查用户区域权限：`/api/v1/system/user/data/area/get-list`
2. 检查用户项目权限：`/api/v1/system/user/data/project/get-list`
3. 申请数据权限或联系管理员授权

### 2.3 租户未绑定应用

**错误示例**：
```json
{"code":100006,"msg":"该企业未绑定该应用，无法登录"}
```

**原因**：用户所属租户未开通当前 `app-id` 对应的应用。

**解决方案**：
1. 平台管理员绑定应用：`/api/v1/system/tenant/app/create`
2. 或切换到已绑定该应用的租户

---

## 三、设备控制问题

### 3.1 设备离线

**现象**：控制命令发送成功但设备无响应。

**排查步骤**：
```bash
# 查询设备在线状态
ur api /api/v1/things/device/info/get-one \
  --body '{"productID":"xxx","deviceName":"yyy"}'
# 检查 isOnline 字段：1=在线，2=离线
```

**说明**：离线设备的控制命令会缓存到**影子设备**（期望值），设备上线后自动同步。

### 3.2 属性标识符不匹配

**错误示例**：
```json
{"code":100001,"msg":"属性不存在"}
```

**原因**：`data` 中的 key 与物模型 `identifier` 不一致。

**解决方案**：
```bash
# 查询物模型
ur api /api/v1/things/product/schema/get-list \
  --body '{"productID":"xxx"}'
# 确认属性标识符（通常是大驼峰，如 CurrentTemperature）
```

### 3.3 deviceName vs deviceAlias 混淆

**常见错误**：用显示名称（`deviceAlias`）作为 `deviceName` 参数。

| 字段 | 含义 | 用途 |
|------|------|------|
| `deviceName` | 设备ID（唯一标识） | API 参数 |
| `deviceAlias` | 设备显示名称 | 列表展示 |

**正确做法**：
```bash
# 先查询获取 deviceName
ur api /api/v1/things/device/info/get-list \
  --body '{"page":{"page":1,"size":10},"deviceAlias":"一楼开关"}'

# 使用返回的 deviceName 进行控制
```

---

## 四、JWT 构造问题

### 4.1 userID 格式错误

**错误示例**：
```json
{"code":401,"msg":"认证失败"}
```

**原因**：JWT payload 中 `userID` 是数字而非字符串。

**正确格式**：
```json
{
  "userID": "12345",    // ✅ 字符串
  "tenantCode": "platform",
  "accessKey": "xxx",
  "exp": 1712345678
}
```

**错误格式**：
```json
{
  "userID": 12345,      // ❌ 数字
  ...
}
```

### 4.2 JWT 过期

**解决方案**：重新生成 JWT，设置合理的 `exp`（建议 1-24 小时）。

```bash
# 使用 ur-api 工具自动构造 JWT
ur api /api/v1/... --access-key xxx --access-secret xxx --user-id 12345
```

---

## 五、网络连接问题

### 5.1 连接被拒绝

**错误示例**：
```
Error: connect ECONNREFUSED 127.0.0.1:7777
```

**排查步骤**：
1. 确认服务已启动：`curl http://localhost:7777/api/v1/system/common/debug`
2. 检查 `baseURL` 配置：`ur config --list`

### 5.2 超时

**错误示例**：
```
Error: timeout of 30000ms exceeded
```

**解决方案**：
1. 检查网络连通性
2. 对于大量数据查询，分批请求减少单次响应时间

---

## 六、常见业务错误码

| 错误码 | 含义 | 常见原因 |
|--------|------|----------|
| `100000` | 系统错误 | 内部异常，联系管理员 |
| `100001` | 参数错误 | 必填字段缺失、格式错误 |
| `100002` | 数据不存在 | ID 错误或已删除 |
| `100003` | 权限不足 | 角色权限不够 |
| `100004` | 数据已存在 | 唯一键冲突 |
| `100005` | 未授权 | Token 无效 |
| `100006` | 业务校验失败 | 业务规则校验不通过 |
| `1000007` | 登录失效 | Token 过期 |

---

## 七、调试技巧

### 7.1 开启详细输出

```bash
ur api /api/v1/xxx --verbose
```

### 7.2 解码 Token

```bash
ur token --decode
```

### 7.3 查看完整 Schema

```bash
ur schema /api/v1/xxx --json
```

### 7.4 从文件读取请求体

```bash
# 适合大型 JSON
ur api /api/v1/xxx --body-file /tmp/payload.json
```