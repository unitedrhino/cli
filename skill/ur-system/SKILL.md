---
name: ur-system
description: "系统管理：通用接口、应用管理、访问令牌、WebSocket推送、授权码、Hook能力。triggers: 系统管理, 文件上传, WebSocket, 批量接口, 应用管理, 菜单管理, Hook配置, 授权码, License"
metadata:
  hermes:
    tags: [system, websocket, file, license]
---


# ur-system — 系统管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/lib/swagger.ts` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 | 注意 |
|------|------|------|
| **AccessKey** | 令牌唯一标识（16位），用于查询和区分令牌 |  |
| **AccessSecret** | 令牌认证密钥（32位），仅创建时返回一次 | 需妥善保存！ |
| **Bearer JWT** | 用 AccessKey/Secret 自行构造 JWT，Authorization: Bearer 调用 API |  |
| **会话 Token** | 登录获取的 session token，token: header，短期有效 |  |
| **应用** | 系统应用（Web、App 等），tenantCode='common' 为全局应用 |  |
| **菜单** | 应用菜单配置，树形结构 |  |
| **WebSocket 频道** | namespace.type.seg1.seg2... 点分层级格式 |  |
| **频道通配符** | * 匹配单段，> 匹配多段（借鉴 NATS） |  |
| **Hook** | 扩展能力（Webhook） |  |
| **授权码** | 许可证管理 |  |

## 角色权限说明

本域 API 按调用者角色区分权限：

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| 平台管理员 | 创建全局应用、管理应用上下架、配置全局 Hook | 创建全局应用、管理应用上下架 |
| 租户管理员 | 订阅应用到租户、自定义菜单配置、使用通用接口 | 订阅应用到租户、文件上传 |
| 普通用户 | 文件上传、WebSocket 连接、批量接口调用 | 文件上传、批量接口调用 |

---

### 平台管理员视角

**权限范围**：创建全局应用、管理应用上下架、配置全局 Hook

#### 典型场景

**创建全局应用**

创建全局应用（tenantCode='common'），定义菜单树形结构

- 涉及 CLI: `ur app info create`, `ur app menu create`
- 工作流: 创建应用 → 定义菜单树 → 配置登录方式

**管理应用上下架**

控制应用的启用/禁用状态

- 涉及 CLI: `ur app info update`
- 工作流: 选择应用 → 更新状态


### 租户管理员视角

**权限范围**：订阅应用到租户、自定义菜单配置、使用通用接口

#### 典型场景

**订阅应用到租户**

为租户启用应用，自定义菜单配置

- 涉及 CLI: `ur tenant app create`, `ur tenant app menu update`
- 工作流: 选择要启用的应用 → 自定义菜单（名称、图标、排序） → 配置登录方式

**文件上传**

上传文件到 OSS，获取访问路径

- 涉及 CLI: `ur system upload upload-file`
- 工作流: 选择文件 → 上传 → 返回文件路径


### 普通用户视角

**权限范围**：文件上传、WebSocket 连接、批量接口调用

#### 典型场景

**文件上传**

上传文件 / 获取上传地址

- 涉及 CLI: `ur system upload upload-file`, `ur system upload upload-url create`
- 工作流: 获取上传 URL 或直接上传 → 返回文件访问路径

**批量接口调用**

一次请求调用多个 API，减少网络开销

- 涉及 CLI: `ur system batch-agg`
- 工作流: 构造批量请求体 → 一次请求获取多个接口数据


## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| 文件上传 | 上传文件到 OSS | [system-upload.md](references/system-upload.md) |
| 批量聚合 | 批量聚合接口调用 | [system-batch-agg.md](references/system-batch-agg.md) |

另有 `ur user`、`ur dept`、`ur alarm` 等根命令也归属系统管理范畴，详见对应 Skill。

> 完整命令帮助：`ur system help`


## 协议与登录页场景

### 协议资源分层

- `system/agreement`：协议本体（协议库 CRUD + 模板预览）
- `system/app/agreement`：应用与协议的绑定关系

### 登录页展示协议

登录页或未登录公开页需要展示协议时，走：

- `POST /api/v1/system/app/core/get-one`

并传：

```json
{
  "id": "100",
  "withAgreementCodes": ["user-agreement", "privacy-policy"]
}
```

从 `agreementsMap` 中按 code 读取协议内容。
## 典型业务场景

### 访问令牌管理

**场景描述**：创建 API 访问令牌用于第三方集成

**涉及 CLI**：
- `ur user self access-token create`
- `ur user self access-token get-list`

**工作流**：
1. 创建令牌获取 AccessKey/Secret
2. 构造 JWT: {userID, tenantCode, accessKey, exp}
3. 使用 Authorization: Bearer <jwt> 调用 API

### WebSocket 实时推送

**场景描述**：连接 WebSocket 获取实时数据推送

**涉及 CLI**：
- `ur system websocket connect`

**工作流**：
1. 携带 token 连接 WebSocket
2. 订阅频道: namespace.type.seg1.seg2
3. 使用通配符 * 或 > 匹配多个频道
4. 接收实时推送消息

### Hook 扩展配置

**场景描述**：配置 Webhook 处理平台事件

**涉及 CLI**：
- `ur hook capability create`
- `ur hook server create`

**工作流**：
1. 创建 Hook 能力
2. 配置回调地址
3. 处理事件回调


## 常用工作流

### 批量聚合接口

```bash
ur system batch-agg \
  --body '{"apis":[{"path":"/api/v1/system/user/self/get-one","body":{}}]}'
```


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
