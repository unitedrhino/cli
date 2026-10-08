---
name: ur-user
description: "用户管理：登录注册、个人信息、角色权限、部门管理、字典管理、定时任务、通知管理、日志查询。triggers: 用户管理, 用户登录, 角色管理, 部门管理, 字典管理, 通知配置, 访问令牌, 修改密码, 定时任务, 登录注册"
metadata:
  hermes:
    tags: [user, auth, role, department, system]
---


# ur-user — 用户管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/generate-api-lists.py` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 |
|------|------|
| **用户** | 平台用户 |
| **角色** | 权限集合，用于用户权限分组 |
| **部门** | 组织架构树形结构 |
| **字典** | 系统枚举值配置 |
| **通知事件** | code、name、group、supportTypes、throttle（防抖间隔） |
| **通知模板** | 按渠道独立配置：短信/邮件/钉钉/站内信等 9 种 |
| **防抖机制** | 同一事件 throttle 秒内只发一次，防止告警风暴 |
| **定时任务** | 定时执行的批处理任务 |

## 角色权限说明

本域 API 按调用者角色区分权限：

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| 企业管理员 | 用户CRUD、角色分配、部门管理、字典管理 | 创建用户、角色与权限管理 |
| 普通用户 | 个人信息、登录、修改密码、访问令牌 | 用户登录、获取个人信息 |

## 平台专属接口说明

本 Skill 涵盖的接口中，以下子域仅限 **平台管理员**（`platform` 权限）调用：

| 子域 | 路径前缀 | 说明 |
|------|---------|------|
| 字典管理 | `/api/v1/system/dict/*` | 字典类型与字典项的 CRUD、批量导入导出 |
| 通知配置 | `/api/v1/system/notify/config/*` | 通知事件配置、通知模板绑定 |
| 通知模板 | `/api/v1/system/notify/template/*` | 通知模板 CRUD |
| 手动通知 | `/api/v1/system/notify/notification/*` | 手动通知/公告的创建、发送、撤回 |
| 资讯管理 | `/api/v1/system/notify/news/*` | 资讯/新闻的 CRUD |
| 消息管理 | `/api/v1/system/notify/message/info/*` | 平台级消息发送与管理 |

> **权限说明**：
> - 上述 `system/notify/*` 和 `system/dict/*` 接口为**平台级管理接口**，仅平台管理员可调用
> - 普通用户/企业管理员**查询字典**请使用公开接口或前端内置字典缓存
> - 普通用户**接收通知/消息**无需调用上述管理接口，由系统根据事件自动推送

---

### 企业管理员视角

**权限范围**：用户CRUD、角色分配、部门管理、字典管理

#### 典型场景

**创建用户**

在企业内创建新用户

- 涉及 CLI: `ur user info create`
- 工作流: 填写用户基本信息 → 分配角色 → 设置部门

**角色与权限管理**

创建角色 / 分配权限 / 分配用户

- 涉及 CLI: `ur role info create`, `ur role access batch-update`, `ur role user batch-create`
- 工作流: 创建角色 → 配置角色权限（菜单/按钮） → 为用户分配角色

**部门管理**

管理组织架构 / 分配用户到部门

- 涉及 CLI: `ur dept info create`, `ur dept user batch-create`
- 工作流: 创建部门树形结构 → 分配用户到部门 → 设置部门负责人

**字典管理**

配置系统枚举值 / 数据字典

- 涉及 CLI: `ur dict info create`, `ur dict detail create`
- 工作流: 创建字典类型 → 添加字典数据项


### 普通用户视角

**权限范围**：个人信息、登录、修改密码、访问令牌

#### 典型场景

**用户登录**

使用账号密码登录获取 session token

- 涉及 CLI: `ur user self login`
- 工作流: 填写账号密码 → 选择登录类型（pwd/sms） → 获取 token

**获取个人信息**

查看当前登录用户的详细信息

- 涉及 CLI: `ur user self get-one`
- 工作流: 无需参数，从 token 解析用户ID → 返回用户详细信息

**查询我的应用列表**

获取当前登录用户可访问的应用列表（区别于平台管理员的应用管理接口）

- 涉及 CLI: `ur user self app get-list`
- 对应 API: `POST /api/v1/system/user/self/app/get-list`
- 工作流: 无需参数，根据用户所在企业的应用开通情况自动过滤
- ⚠️ 注意: 查询【我的应用】请使用此接口，不要使用 `/api/v1/system/app/info/get-list`（后者仅平台管理员可用）

**修改密码**

修改当前用户的登录密码

- 涉及 CLI: `ur user self change-pwd`
- 工作流: 验证旧密码 → 设置新密码 → 重新登录

**访问令牌管理**

创建/查看 API 访问令牌（AccessKey）

- 涉及 CLI: `ur user self access-token create`, `ur user self access-token get-list`
- 工作流: 创建令牌获取 AccessKey/Secret → 使用令牌调用 API


## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| 用户管理 | 查询/创建/更新/删除用户 | [user-info.md](references/user-info.md) |
| 个人中心 | 登录/注册/信息/密码 | [user-self.md](references/user-self.md) |
| 访问令牌 | 查询/创建/更新/删除访问令牌 | [user-access-token.md](references/user-access-token.md) |
| 企业管理 | 查询/更新/删除/加入企业 | [user-tenant.md](references/user-tenant.md) |
| 应用与菜单 | **查询当前用户可用应用列表**（`ur user self app get-list`）/详情、查询菜单列表 | [user-app-menu.md](references/user-app-menu.md) |
| 消息管理 | 查询/处理/标记消息 | [user-message.md](references/user-message.md) |
| 通知偏好 | 读取/更新通知偏好 | [user-notify.md](references/user-notify.md) |
| 第三方客户端绑定 | 检查/完成第三方客户端绑定 | [user-thirdparty.md](references/user-thirdparty.md) |
| 部门用户管理 | 批量创建/删除部门用户 | [user-dept.md](references/user-dept.md) |
| 数据权限 | 查询项目/区域权限列表 | [user-data.md](references/user-data.md) |

> 完整命令帮助：`ur user help`



## 典型业务场景

### 用户登录与个人信息

**场景描述**：登录 / 获取个人信息 / 修改密码

**涉及 CLI**：
- `ur user self login`
- `ur user self get-one`
- `ur user self change-pwd`

**工作流**：
1. 用户名密码登录
2. 获取个人信息
3. 修改密码

### 访问令牌管理

**场景描述**：创建 API 访问令牌 / 查看令牌列表（AccessKey/Secret 用于 JWT 构造）

**涉及 CLI**：
- `ur user self access-token create`
- `ur user self access-token get-list`

**工作流**：
1. 创建令牌（获取 AccessKey/Secret）
2. 构造 JWT: {userID, tenantCode, accessKey, exp}
3. 使用 Authorization: Bearer <jwt> 调用 API

### 通知配置与发送

**场景描述**：配置设备告警通知 / 发送批量通知（防抖机制防止告警风暴）

**涉及 CLI**：
- `ur notify config create`
- `ur notify template create`
- `ur notify message send`

**工作流**：
1. 创建事件配置（code、throttle 防抖间隔）
2. 创建通知模板（按渠道：短信/邮件/钉钉）
3. 触发事件发送通知

### 角色与权限管理

**场景描述**：创建角色 / 分配权限 / 分配用户

**涉及 CLI**：
- `ur role info create`
- `ur role access batch-update`
- `ur role user batch-create`

**工作流**：
1. 创建角色
2. 配置角色权限
3. 为用户分配角色


## 常用工作流

### 获取当前用户信息

```bash
ur user self get-one
```


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
