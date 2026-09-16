# 认证方式

## AI 决策顺序

先运行：

```bash
ur check --json
```

- `auth_status=ok`：复用现有 Sandbox 环境或历史 profile，直接调用 API。
- 缺少认证：再选择 `ur login --method device|password|aksk`；默认方式为 `device`。
- 不要因 CLI 升级删除 `~/.ur/config.json`、重新执行 `setup` 或索要 `userID`。

## Device Flow（默认）

```
ur login → 生成 setup code → 浏览器授权 → 创建访问令牌 → 点击「完成第三方客户端绑定」
    → CLI 轮询 /setup-check → 获取 AK/SK → 保存到 ~/.ur/config.json
```

- 无需输入账号密码
- 授权通过后自动保存 AccessKey + AccessSecret + TenantCode
- 后续调用 API 时使用 AK/SK 生成 JWT

```bash
# AI 环境使用分步授权
ur login --method device --no-wait --json           # 获取授权 URL
ur login --method device --setup-code <CODE> --json # 用户确认后完成授权
```

## 账号密码

密码必须是原始密码，CLI 会做且只做一次 SHA-256。不要预先摘要。

```bash
UR_PASSWORD='<原始密码>' ur login --method password \
  --account '<账号>' --tenant-code '<企业编码>' --json

printf '%s' "$UR_PASSWORD" | ur login --method password \
  --account '<账号>' --tenant-code '<企业编码>' --password-stdin --json
```

`--password` 明文参数仅为兼容；环境变量或 `--password-stdin` 更适合 AI/Sandbox。直接调用登录 HTTP 接口时才由调用方发送 SHA-256 摘要与 `pwdType: 1`；`pwdType: 2` 的 MD5 仅为旧版兼容。

## AccessKey/JWT（程序化访问）

```bash
UR_ACCESS_SECRET='<AccessSecret>' ur login --method aksk \
  --access-key '<AccessKey>' --tenant-code '<企业编码>' --json

printf '%s' "$UR_ACCESS_SECRET" | ur login --method aksk \
  --access-key '<AccessKey>' --tenant-code '<企业编码>' \
  --access-secret-stdin --json
```

`userID` 不是 AK/SK 登录的必填输入；CLI 可使用兼容值签名，服务端按 AccessKey 关联真实用户。需要先在平台 UI 创建访问令牌：用户设置 → 访问令牌 → 创建。

```
AccessKey + AccessSecret → HS256 JWT → Authorization: Bearer header
```

## Sandbox 环境变量（env-only）

```bash
export UR_BASE_URL='<平台地址>'
export UR_APP_ID='<应用ID>'
export UR_TENANT_CODE='<企业编码>'

# 以下认证组任选一组
export UR_TOKEN='<Session Token>'
# export UR_ACCESS_KEY='<AccessKey>' UR_ACCESS_SECRET='<AccessSecret>'
# export UR_ACCOUNT='<账号>' UR_PASSWORD='<原始密码>'
```

设置 `UR_BASE_URL` 后，CLI 不读取磁盘 profile 补齐认证，也不持久化 Sandbox 凭据。环境认证优先级为 Token → AK/SK →账号密码；AK/SK、账号密码都必须成对注入。`UR_USER_ID` 可选。

## 历史 profile 兼容

- 旧 profile 只有 `baseURL/appID/tenantCode/account/password` 时，`check/api/token` 会自动登录并只补充 Session Token。
- 遗留 Token 过期后，CLI 使用保存的账号密码刷新；账号密码被后端明确拒绝时，才回退到同一历史 profile 中完整的 AK/SK。
- 网络错误、响应解析错误和业务错误不会被当作凭据失效，也不会盲目切换认证方式。
- 新的显式 `login` 会替换冲突认证字段；普通升级和 Token 刷新不会清空旧账号密码。
- `ur setup` 仅保留为人类终端的账号密码兼容向导。

## 认证方式对比

| 方式 | 获取方式 | 适用场景 |
|------|----------|----------|
| Device Flow | 浏览器授权（setup code） | 无现成凭据时的默认方式 |
| Password | 原始账号密码换取 Session Token | 已有业务账号的自动化环境 |
| AK/SK | AccessKey + AccessSecret 自签 JWT | 服务间调用、长期自动化 |
| Sandbox env-only | 运行时直接注入上述任一凭据组 | 隔离执行环境，不落盘 |

## JWT Payload 字段

| 字段名 | JSON 名称 | 说明 |
|--------|-----------|------|
| `UserID` | `"userID"` | 兼容字段；缺失时 CLI 可使用字符串 `"0"` |
| `TenantCode` | `"tenantCode"` | 当前认证上下文的企业代码 |
| `AccessKey` | `"accessKey"` | 令牌标识 |
| `exp` | `"exp"` | Unix 秒时间戳（可选） |

## 关键 HTTP Header

| Header | 值 | 说明 |
|--------|-----|------|
| `Authorization` | Bearer \<jwt\> | JWT 模式 |
| `app-id` | 应用 ID（如 `77`） | 必须 |
| `tenant-code` | 企业代码 | 必须 |

> 不要把 AK/SK 自签 JWT 填入 `UR_TOKEN`。`UR_TOKEN` 仅表示服务端颁发的 Session Token。
