---
name: ur-tenant
description: "企业管理：企业CRUD、用户管理、权限配置、应用绑定。triggers: 企业管理, 创建企业, 企业列表, 企业用户, 企业配置, 邀请用户, 企业应用, SaaS企业"
metadata:
  hermes:
    tags: [tenant, saas, user-management, permission]
---


# ur-tenant — 企业管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/generate-api-lists.py` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 | 注意 |
|------|------|------|
| **企业** | SaaS 企业，独立的业务隔离单元 |  |
| **企业代码** | 企业唯一标识（tenantCode） | 请求头需传 tenant-code |
| **企业用户** | 企业内的用户 |  |
| **企业权限** | 企业级权限配置 |  |
| **企业配置** | 企业个性化配置 |  |

## 角色权限说明

本域 API 按调用者角色区分权限：

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| 平台管理员 | 创建/删除企业、查看所有企业、管理企业配置 | 创建新企业、查看所有企业 |
| 企业管理员 | 管理本企业信息、邀请用户、配置权限 | 查看本企业信息、邀请用户加入企业 |

## 平台专属接口说明

本 Skill 涵盖的接口中，以下接口仅限 **平台管理员**（`platform` 权限）调用：

| 子域 | 路径前缀 | 说明 |
|------|---------|------|
| 企业列表 | `/api/v1/system/tenant/info/get-list` | 查询平台下所有企业 |
| 删除企业 | `/api/v1/system/tenant/info/delete` | 删除指定企业 |
| 企业续费授权 | `/api/v1/system/tenant/renewal/*` | 企业续费与直接授权 |
| 企业应用管理 | `/api/v1/system/tenant/app/*` | 平台视角管理企业应用绑定 |

> **权限说明**：
> - `system/tenant/info/get-list` 是**平台视角**的企业列表，仅平台管理员可调用
> - 企业管理员查看**本企业信息**请使用 `ur system/tenant/info get-one`（权限 `admin`）
> - 企业管理员管理**本企业用户**请使用 `ur system/tenant/user get-list`（权限 `admin`）

---

### 平台管理员视角

**权限范围**：创建/删除企业、查看所有企业、管理企业配置

#### 典型场景

**创建新企业**

创建一个新的 SaaS 企业

- 涉及 CLI: `ur system/tenant/info create`
- 工作流: 填写企业名称和代码 → 设置企业管理员 → 绑定应用 → 配置企业参数

**查看所有企业**

查询平台下所有企业列表

- 涉及 CLI: `ur system/tenant/info get-list`
- 工作流: 无需传 tenant-code header → 返回所有企业列表

**更新企业信息**

修改企业的基本信息和配置

- 涉及 CLI: `ur system/tenant/info update`, `ur system/tenant/config update`
- 工作流: 选择目标企业 → 修改配置项 → 保存更新


### 企业管理员视角

**权限范围**：管理本企业信息、邀请用户、配置权限

#### 典型场景

**查看本企业信息**

获取当前企业的详细信息和配置

- 涉及 CLI: `ur system/tenant/info get-one`, `ur system/tenant/config get-one`
- 工作流: 请求需带 tenant-code header → 只能查看本企业信息

**邀请用户加入企业**

邀请新用户或已有用户加入本企业

- 涉及 CLI: `ur tenant user invite`, `ur tenant user batch-create`
- 工作流: 填写用户账号 → 分配角色 → 发送邀请

**管理企业用户**

查看/修改/删除企业内的用户

- 涉及 CLI: `ur tenant user get-list`, `ur tenant user update`, `ur tenant user delete`
- 工作流: 查询用户列表 → 修改用户角色或状态 → 删除用户

**配置企业应用**

为企业启用/禁用应用

- 涉及 CLI: `ur system/tenant/app create`, `ur system/tenant/app get-list`
- 工作流: 查看可用应用列表 → 启用应用 → 配置应用菜单


## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| 企业用户管理 | 查询/创建/更新/删除企业用户 | [tenant-user.md](references/tenant-user.md) |
| 邀请管理 | 邀请用户/发送邀请/邀请码/待处理邀请 | [tenant-invite.md](references/tenant-invite.md) |
| 角色管理 | 查询/批量更新用户角色 | [tenant-role.md](references/tenant-role.md) |

> 完整命令帮助：`ur tenant help`


## 典型业务场景

### 企业创建与配置

**场景描述**：创建企业 / 配置企业信息

**涉及 CLI**：
- `ur system/tenant/info create`
- `ur system/tenant/config update`

**工作流**：
1. 创建企业
2. 配置企业参数

### 企业用户管理

**场景描述**：邀请用户 / 分配角色

**涉及 CLI**：
- `ur tenant user invite`
- `ur tenant user batch-create`

**工作流**：
1. 邀请用户加入企业
2. 为用户分配角色


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
