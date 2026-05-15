---
name: ur-tenant
description: "租户管理：租户CRUD、用户管理、权限配置、应用绑定。triggers: 租户管理, 创建租户, 租户列表, 租户用户, 租户配置, 邀请用户, 租户应用, SaaS租户"
metadata:
  hermes:
    tags: [tenant, saas, user-management, permission]
---


# ur-tenant — 租户管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/lib/swagger.ts` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 | 注意 |
|------|------|------|
| **租户** | SaaS 租户，独立的业务隔离单元 |  |
| **租户代码** | 租户唯一标识（tenantCode） | 请求头需传 tenant-code |
| **租户用户** | 租户内的用户 |  |
| **租户权限** | 租户级权限配置 |  |
| **租户配置** | 租户个性化配置 |  |

## 角色权限说明

本域 API 按调用者角色区分权限：

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| 平台管理员 | 创建/删除租户、查看所有租户、管理租户配置 | 创建新租户、查看所有租户 |
| 租户管理员 | 管理本租户信息、邀请用户、配置权限 | 查看本租户信息、邀请用户加入租户 |

---

### 平台管理员视角

**权限范围**：创建/删除租户、查看所有租户、管理租户配置

#### 典型场景

**创建新租户**

创建一个新的 SaaS 租户

- 涉及 API: `/api/v1/system/tenant/info/create`
- 工作流: 填写租户名称和代码 → 设置租户管理员 → 绑定应用 → 配置租户参数

**查看所有租户**

查询平台下所有租户列表

- 涉及 API: `/api/v1/system/tenant/info/get-list`, `/api/v1/system/tenant/info/get-one`
- 工作流: 无需传 tenant-code header → 返回所有租户列表

**更新租户信息**

修改租户的基本信息和配置

- 涉及 API: `/api/v1/system/tenant/info/update`, `/api/v1/system/tenant/config/update`
- 工作流: 选择目标租户 → 修改配置项 → 保存更新


### 租户管理员视角

**权限范围**：管理本租户信息、邀请用户、配置权限

#### 典型场景

**查看本租户信息**

获取当前租户的详细信息和配置

- 涉及 API: `/api/v1/system/tenant/info/get-one`, `/api/v1/system/tenant/config/get-one`
- 工作流: 请求需带 tenant-code header → 只能查看本租户信息

**邀请用户加入租户**

邀请新用户或已有用户加入本租户

- 涉及 API: `/api/v1/system/tenant/user/invite`, `/api/v1/system/tenant/user/batch-create`
- 工作流: 填写用户账号 → 分配角色 → 发送邀请

**管理租户用户**

查看/修改/删除租户内的用户

- 涉及 API: `/api/v1/system/tenant/user/get-list`, `/api/v1/system/tenant/user/update`, `/api/v1/system/tenant/user/delete`
- 工作流: 查询用户列表 → 修改用户角色或状态 → 删除用户

**配置租户应用**

为租户启用/禁用应用

- 涉及 API: `/api/v1/system/tenant/app/create`, `/api/v1/system/tenant/app/get-list`
- 工作流: 查看可用应用列表 → 启用应用 → 配置应用菜单


## API 参考

<!-- API_LIST:ur-tenant -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/agreement/create` | 添加协议 | admin |
| POST | `/api/v1/system/tenant/agreement/delete` | 删除协议 | admin |
| POST | `/api/v1/system/tenant/agreement/get-list` | 获取协议列表 | admin |
| POST | `/api/v1/system/tenant/agreement/get-one` | 获取协议详情 | public |
| POST | `/api/v1/system/tenant/agreement/update` | 更新协议 | admin |
| POST | `/api/v1/system/tenant/app/create` | 绑定租户应用 | all |
| POST | `/api/v1/system/tenant/app/delete` | 解绑租户应用 | platform |
| POST | `/api/v1/system/tenant/app/get-list` | 获取租户应用列表 | admin |
| POST | `/api/v1/system/tenant/app/get-one` | 获取应用信息 | platform |
| POST | `/api/v1/system/tenant/app/menu/get-list` | 获取租户应用菜单列表 | admin |
| POST | `/api/v1/system/tenant/app/menu/update` | 更新租户应用菜单 | admin |
| POST | `/api/v1/system/tenant/app/update` | 更新租户应用 | platform |
| POST | `/api/v1/system/tenant/config/get-one` | 获取租户配置 | admin |
| POST | `/api/v1/system/tenant/config/update` | 更新租户配置 | admin |
| POST | `/api/v1/system/tenant/core/get-list` | 搜索租户信息 | public |
| POST | `/api/v1/system/tenant/core/get-one` | 获取租户信息 | public |
| POST | `/api/v1/system/tenant/info/create` | 添加租户 | all |
| POST | `/api/v1/system/tenant/info/delete` | 删除租户 | platform |
| POST | `/api/v1/system/tenant/info/get-list` | 获取租户列表 | platform |
| POST | `/api/v1/system/tenant/info/get-one` | 获取租户详情 | admin |
| POST | `/api/v1/system/tenant/info/transfer` | 转让租户（仅租户所有者） | admin |
| POST | `/api/v1/system/tenant/info/update` | 更新租户 | admin |
| POST | `/api/v1/system/tenant/renewal/direct-grant` | 直接授权 | platform |
| POST | `/api/v1/system/tenant/renewal/get-list` | 获取续期列表 | admin |
| POST | `/api/v1/system/tenant/renewal/renew` | 续费 | admin |
| POST | `/api/v1/system/tenant/user/batch-create` | 批量添加用户加入租户 | admin |
| POST | `/api/v1/system/tenant/user/delete` | 删除租户用户 | admin |
| POST | `/api/v1/system/tenant/user/get-list` | 获取租户用户列表 | admin |
| POST | `/api/v1/system/tenant/user/get-one` | 获取租户用户详情,会同时返回所拥有的角色列表 | admin |
| POST | `/api/v1/system/tenant/user/invite` | 邀请用户加入租户 | admin |
| POST | `/api/v1/system/tenant/user/invite-code/gen` | 生成租户用户邀请码 | admin |
| POST | `/api/v1/system/tenant/user/invite-code/get-one` | 获取当前有效的租户用户邀请码 | admin |
| POST | `/api/v1/system/tenant/user/invite-pending/delete` | 删除待处理邀请 | admin |
| POST | `/api/v1/system/tenant/user/invite-pending/get-list` | 获取待处理邀请列表 | admin |
| POST | `/api/v1/system/tenant/user/invite-send` | 发送邀请（支持邮件和手机号） | admin |
| POST | `/api/v1/system/tenant/user/role/batch-update` | 更新租户用户的角色列表 | admin |
| POST | `/api/v1/system/tenant/user/role/get-list` | 获取租户用户角色列表 | admin |
| POST | `/api/v1/system/tenant/user/update` | 更新租户用户 | admin |

<!-- END_API_LIST -->

权限: mixed

| 端点 | 说明 | 权限 |
|---|------|------|
| POST /api/v1/system/tenant/access/info/batch-create | 批量创建租户操作权限 | 管理员 |
| POST /api/v1/system/tenant/access/info/batch-delete | 批量删除租户操作权限 | 管理员 |
| POST /api/v1/system/tenant/access/info/batch-update | 批量更新租户操作权限 | 管理员 |
| POST /api/v1/system/tenant/access/info/get-list | 获取租户操作权限列表 | 管理员 |
| POST /api/v1/system/tenant/access/info/tree | 获取租户操作权限树 | 管理员 |
| POST /api/v1/system/tenant/agreement/create | 添加协议 | 管理员 |
| POST /api/v1/system/tenant/agreement/delete | 删除协议 | 管理员 |
| POST /api/v1/system/tenant/agreement/get-list | 获取协议列表 | 管理员 |
| POST /api/v1/system/tenant/agreement/get-one | 获取协议详情 | public |
| POST /api/v1/system/tenant/agreement/update | 更新协议 | 管理员 |
| POST /api/v1/system/tenant/app/create | 绑定租户应用 | 平台管理员 |
| POST /api/v1/system/tenant/app/delete | 解绑租户应用 | 平台管理员 |
| POST /api/v1/system/tenant/app/get-list | 获取租户应用列表 | 管理员 |
| POST /api/v1/system/tenant/app/get-one | 获取应用信息 | 平台管理员 |
| POST /api/v1/system/tenant/app/menu/get-list | 获取租户应用菜单列表 | 管理员 |
| POST /api/v1/system/tenant/app/menu/update | 更新租户应用菜单 | 平台管理员 |
| POST /api/v1/system/tenant/app/update | 更新租户应用 | 平台管理员 |
| POST /api/v1/system/tenant/config/get-one | 获取租户配置 | 管理员 |
| POST /api/v1/system/tenant/config/update | 更新租户配置 | 管理员 |
| POST /api/v1/system/tenant/info/create | 添加租户 | 平台管理员 |
| POST /api/v1/system/tenant/info/delete | 删除租户 | 平台管理员 |
| POST /api/v1/system/tenant/info/get-list | 获取租户列表 | 平台管理员 |
| POST /api/v1/system/tenant/info/get-one | 获取租户详情 | 管理员 |
| POST /api/v1/system/tenant/info/transfer | 转让租户（仅租户所有者） | 管理员 |
| POST /api/v1/system/tenant/info/update | 更新租户 | 管理员 |
| POST /api/v1/system/tenant/user/batch-create | 批量添加用户加入租户 | 管理员 |
| POST /api/v1/system/tenant/user/delete | 删除租户用户 | 管理员 |
| POST /api/v1/system/tenant/user/get-list | 获取租户用户列表 | 管理员 |
| POST /api/v1/system/tenant/user/get-one | 获取租户用户详情,会同时返回所拥有的角色列表 | 管理员 |
| POST /api/v1/system/tenant/user/invite | 邀请用户加入租户 | 管理员 |
| POST /api/v1/system/tenant/user/invite-code/gen | 生成租户用户邀请码 | 管理员 |
| POST /api/v1/system/tenant/user/invite-code/get-one | 获取当前有效的租户用户邀请码 | 管理员 |
| POST /api/v1/system/tenant/user/invite-pending/delete | 删除待处理邀请 | 管理员 |
| POST /api/v1/system/tenant/user/invite-pending/get-list | 获取待处理邀请列表 | 管理员 |
| POST /api/v1/system/tenant/user/invite-send | 发送邀请（支持邮件和手机号） | 管理员 |
| POST /api/v1/system/tenant/user/role/batch-update | 更新租户用户的角色列表 | 管理员 |
| POST /api/v1/system/tenant/user/role/get-list | 获取租户用户角色列表 | 管理员 |
| POST /api/v1/system/tenant/user/update | 更新租户用户 | 管理员 |


## 典型业务场景

### 租户创建与配置

**场景描述**：创建租户 / 配置租户信息

**涉及 API**：
- `/api/v1/system/tenant/info/create`
- `/api/v1/system/tenant/config/update`

**工作流**：
1. 创建租户
2. 配置租户参数

### 租户用户管理

**场景描述**：邀请用户 / 分配角色

**涉及 API**：
- `/api/v1/system/tenant/user/invite`
- `/api/v1/system/tenant/user/batch-create`

**工作流**：
1. 邀请用户加入租户
2. 为用户分配角色


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
