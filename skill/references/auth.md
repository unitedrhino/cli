# 认证方式

## Device Auth（推荐）

```
ur login → 生成 setup code → 浏览器授权 → 创建访问令牌 → 点击「完成第三方客户端绑定」
    → CLI 轮询 /setup-check → 获取 AK/SK → 保存到 ~/.ur/config.json
```

- 无需输入账号密码
- 授权通过后自动保存 AccessKey + AccessSecret + TenantCode
- 后续调用 API 时使用 AK/SK 生成 JWT

```bash
# AI 环境使用分步授权
ur login --no-wait --json          # 获取授权 URL
ur login --setup-code <CODE> --json # 用户确认后完成授权
```

## AccessKey/JWT（程序化访问）

已有 AccessKey/Secret 时直接配置：

```bash
ur setup  # 交互式配置
```

或手动编辑 `~/.ur/config.json`。需要先在平台 UI 创建访问令牌：用户设置 → 访问令牌 → 创建。

```
AccessKey + AccessSecret → HS256 JWT → Authorization: Bearer header
```

## 环境变量

```bash
export UR_BASE_URL=http://localhost:7777
export UR_APP_ID=77
export UR_ACCESS_KEY=xxx
export UR_ACCESS_SECRET=xxx
export UR_USER_ID=12345
export UR_TOKEN=xxx        # 临时使用已有 token
export UR_APP=iot          # 指定应用上下文
```

## 认证方式对比

| 特性 | Device Auth | AccessKey/JWT |
|------|-------------|---------------|
| 获取方式 | 浏览器授权（setup code） | 手动输入 AK/SK 或账号密码 |
| 有效期 | 由访问令牌设置决定 | 由客户端控制（建议 1~24 小时） |
| 适用场景 | 人机交互、首次配置 | 服务间调用、自动化脚本 |
| 撤销方式 | 删除访问令牌 | 删除访问令牌 |

## JWT Payload 字段

| 字段名 | JSON 名称 | 说明 |
|--------|-----------|------|
| `UserID` | `"userID"` | **必须为字符串格式** |
| `TenantCode` | `"tenantCode"` | 企业代码，可为空 |
| `AccessKey` | `"accessKey"` | 令牌标识 |
| `exp` | `"exp"` | Unix 秒时间戳（可选） |

## 关键 HTTP Header

| Header | 值 | 说明 |
|--------|-----|------|
| `Authorization` | Bearer \<jwt\> | JWT 模式 |
| `app-id` | 应用 ID（如 `77`） | 必须 |
| `tenant-code` | 企业代码 | 必须 |

> **关键陷阱**: JWT 中 `userID` 必须为字符串格式（`"userID": "12345"` 而非数字 `12345`），否则 401 错误。
