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
> 如需更新，修改 `scripts/lib/swagger.ts` 中的域名定义后重新生成。

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
| 租户管理员 | 用户CRUD、角色分配、部门管理、字典管理 | 创建用户、角色与权限管理 |
| 普通用户 | 个人信息、登录、修改密码、访问令牌 | 用户登录、获取个人信息 |

---

### 租户管理员视角

**权限范围**：用户CRUD、角色分配、部门管理、字典管理

#### 典型场景

**创建用户**

在租户内创建新用户

- 涉及 API: `/api/v1/system/user/info/create`
- 工作流: 填写用户基本信息 → 分配角色 → 设置部门

**角色与权限管理**

创建角色 / 分配权限 / 分配用户

- 涉及 API: `/api/v1/system/role/info/create`, `/api/v1/system/role/access/update`, `/api/v1/system/role/user/batch-create`
- 工作流: 创建角色 → 配置角色权限（菜单/按钮） → 为用户分配角色

**部门管理**

管理组织架构 / 分配用户到部门

- 涉及 API: `/api/v1/system/dept/info/create`, `/api/v1/system/dept/user/batch-create`
- 工作流: 创建部门树形结构 → 分配用户到部门 → 设置部门负责人

**字典管理**

配置系统枚举值 / 数据字典

- 涉及 API: `/api/v1/system/dict/type/create`, `/api/v1/system/dict/data/create`
- 工作流: 创建字典类型 → 添加字典数据项


### 普通用户视角

**权限范围**：个人信息、登录、修改密码、访问令牌

#### 典型场景

**用户登录**

使用账号密码登录获取 session token

- 涉及 API: `/api/v1/system/user/self/login`
- 工作流: 填写账号密码 → 选择登录类型（pwd/sms） → 获取 token

**获取个人信息**

查看当前登录用户的详细信息

- 涉及 API: `/api/v1/system/user/self/get-one`
- 工作流: 无需参数，从 token 解析用户ID → 返回用户详细信息

**修改密码**

修改当前用户的登录密码

- 涉及 API: `/api/v1/system/user/self/change-pwd`
- 工作流: 验证旧密码 → 设置新密码 → 重新登录

**访问令牌管理**

创建/查看 API 访问令牌（AccessKey）

- 涉及 API: `/api/v1/system/user/self/access-token/create`, `/api/v1/system/user/self/access-token/get-list`
- 工作流: 创建令牌获取 AccessKey/Secret → 使用令牌调用 API


## API 参考

<!-- API_LIST:ur-user -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/dept/info/create` | 添加部门详情 | admin |
| POST | `/api/v1/system/dept/info/delete` | 删除部门 | admin |
| POST | `/api/v1/system/dept/info/get-list` | 获取部门列表 | admin |
| POST | `/api/v1/system/dept/info/get-one` | 获取部门详情 | admin |
| POST | `/api/v1/system/dept/info/update` | 更新部门 | admin |
| POST | `/api/v1/system/dept/sync-job/create` | 添加同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/delete` | 删除同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/execute` | 执行同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/get-list` | 获取同步任务列表 | admin |
| POST | `/api/v1/system/dept/sync-job/get-one` | 获取同步任务详情 | admin |
| POST | `/api/v1/system/dept/sync-job/update` | 更新同步任务 | admin |
| POST | `/api/v1/system/dept/user/batch-create` | 批量授权部门用户 | admin |
| POST | `/api/v1/system/dept/user/batch-delete` | 批量取消授权部门用户 | admin |
| POST | `/api/v1/system/dept/user/get-list` | 获取部门授权列表 | admin |
| POST | `/api/v1/system/dict/detail/create` | 添加字典详情 | platform |
| POST | `/api/v1/system/dict/detail/delete` | 删除字典详情 | platform |
| POST | `/api/v1/system/dict/detail/get-list` | 获取字典详情列表 | platform |
| POST | `/api/v1/system/dict/detail/get-one` | 获取字典详情单个 | platform |
| POST | `/api/v1/system/dict/detail/update` | 更新字典详情 | platform |
| POST | `/api/v1/system/dict/info/batch-export` | 批量导出字典信息 | platform |
| POST | `/api/v1/system/dict/info/batch-import` | 批量导入字典信息 | platform |
| POST | `/api/v1/system/dict/info/create` | 添加字典信息 | platform |
| POST | `/api/v1/system/dict/info/delete` | 删除字典信息 | platform |
| POST | `/api/v1/system/dict/info/get-list` | 获取字典信息列表 | all |
| POST | `/api/v1/system/dict/info/get-one` | 获取字典信息详情 | platform |
| POST | `/api/v1/system/dict/info/update` | 更新字典信息 | platform |
| POST | `/api/v1/system/log/login/get-list` | 获取登录日志列表 | admin |
| POST | `/api/v1/system/log/oper/get-list` | 获取操作日志列表 | admin |
| POST | `/api/v1/system/notify/config/create` | 添加通知配置 | platform |
| POST | `/api/v1/system/notify/config/delete` | 删除通知配置 | platform |
| POST | `/api/v1/system/notify/config/get-list` | 获取通知列表 | platform |
| POST | `/api/v1/system/notify/config/get-one` | 获取通知详情 | platform |
| POST | `/api/v1/system/notify/config/send-test` | 测试发送通知 | platform |
| POST | `/api/v1/system/notify/config/template/delete` | 删除通知绑定模版 | platform |
| POST | `/api/v1/system/notify/config/template/get-list` | 获取通知绑定模版列表 | platform |
| POST | `/api/v1/system/notify/config/template/update` | 更新通知绑定模版 | platform |
| POST | `/api/v1/system/notify/config/tree` | 获取通知列表树 | platform |
| POST | `/api/v1/system/notify/config/update` | 更新通知配置 | platform |
| POST | `/api/v1/system/notify/message/info/delete` | 删除消息 | platform |
| POST | `/api/v1/system/notify/message/info/get-list` | 获取消息列表 | platform |
| POST | `/api/v1/system/notify/message/info/send` | 发送消息 | platform |
| POST | `/api/v1/system/notify/message/info/update` | 更新消息 | platform |
| POST | `/api/v1/system/notify/news/create` | 创建资讯 | platform |
| POST | `/api/v1/system/notify/news/delete` | 删除资讯 | platform |
| POST | `/api/v1/system/notify/news/info` | 获取资讯详情 | platform |
| POST | `/api/v1/system/notify/news/list` | 获取资讯列表 | platform |
| POST | `/api/v1/system/notify/news/update` | 更新资讯 | platform |
| POST | `/api/v1/system/notify/notification/copy` | 复制手动通知 | platform |
| POST | `/api/v1/system/notify/notification/create` | 创建手动通知 | platform |
| POST | `/api/v1/system/notify/notification/delete` | 删除手动通知 | platform |
| POST | `/api/v1/system/notify/notification/estimate-users` | 预估送达用户数 | platform |
| POST | `/api/v1/system/notify/notification/index` | 获取手动通知列表 | platform |
| POST | `/api/v1/system/notify/notification/read` | 获取手动通知详情 | platform |
| POST | `/api/v1/system/notify/notification/revoke` | 撤回手动通知 | platform |
| POST | `/api/v1/system/notify/notification/send` | 发送手动通知 | platform |
| POST | `/api/v1/system/notify/notification/update` | 更新手动通知 | platform |
| POST | `/api/v1/system/notify/template/create` | 添加通知模版 | platform |
| POST | `/api/v1/system/notify/template/delete` | 删除通知模版 | platform |
| POST | `/api/v1/system/notify/template/get-list` | 获取通知模版列表 | platform |
| POST | `/api/v1/system/notify/template/get-one` | 获取通知模版详情 | platform |
| POST | `/api/v1/system/notify/template/update` | 更新通知模版 | platform |
| POST | `/api/v1/system/role/app/batch-update` | 更新APP权限 | admin |
| POST | `/api/v1/system/role/app/get-list` | 获取APP权限列表 | admin |
| POST | `/api/v1/system/role/info/create` | 添加角色 | admin |
| POST | `/api/v1/system/role/info/delete` | 删除角色 | admin |
| POST | `/api/v1/system/role/info/get-list` | 获取角色列表 | admin |
| POST | `/api/v1/system/role/info/update` | 更新角色 | admin |
| POST | `/api/v1/system/role/menu/batch-update` | 更新角色对应菜单列表 | admin |
| POST | `/api/v1/system/role/menu/get-list` | 获取角色对应菜单列表 | admin |
| POST | `/api/v1/system/role/resource/batch-update` | 批量更新角色资源动作权限 | admin |
| POST | `/api/v1/system/role/resource/get-list` | 获取角色资源动作权限列表 | admin |
| POST | `/api/v1/system/user/data/area/get-list` | 获取区域权限列表 | admin |
| POST | `/api/v1/system/user/data/project/get-list` | 获取项目权限列表 | admin |
| POST | `/api/v1/system/user/dept/batch-create` | 新增用户的部门列表 | admin |
| POST | `/api/v1/system/user/dept/batch-delete` | 删除用户的部门列表 | admin |
| POST | `/api/v1/system/user/info/create` | 创建用户信息 | admin |
| POST | `/api/v1/system/user/info/delete` | 刪除用户 | admin |
| POST | `/api/v1/system/user/info/get-list` | 查询用户信息列表 | admin |
| POST | `/api/v1/system/user/info/get-one` | 获取用户信息 | admin |
| POST | `/api/v1/system/user/info/update` | 更新用户基本数据 | admin |
| POST | `/api/v1/system/user/self/access-token/create` | 创建访问令牌 | all |
| POST | `/api/v1/system/user/self/access-token/delete` | 删除访问令牌 | all |
| POST | `/api/v1/system/user/self/access-token/get-list` | 获取访问令牌列表 | all |
| POST | `/api/v1/system/user/self/access-token/get-one` | 获取访问令牌详情 | all |
| POST | `/api/v1/system/user/self/access-token/update` | 更新访问令牌 | all |
| POST | `/api/v1/system/user/self/app/get-list` | 获取用户应用列表 | all |
| POST | `/api/v1/system/user/self/app/get-one` | 获取用户应用详情 | all |
| POST | `/api/v1/system/user/self/bind-account` | 绑定账号 | all |
| POST | `/api/v1/system/user/self/cancel` | 注销用户 | all |
| POST | `/api/v1/system/user/self/captcha` | 获取验证码 | public |
| POST | `/api/v1/system/user/self/change-pwd` | 更新用户密码 | all |
| POST | `/api/v1/system/user/self/forget-pwd` | 忘记密码 | public |
| POST | `/api/v1/system/user/self/get-one` | 获取用户信息 | all |
| POST | `/api/v1/system/user/self/login` | 用户登录 | public |
| POST | `/api/v1/system/user/self/logout` | 用户登出 | all |
| POST | `/api/v1/system/user/self/menu/get-list` | 获取用户菜单列表 | all |
| POST | `/api/v1/system/user/self/message/get-list` | 用户消息列表 | all |
| POST | `/api/v1/system/user/self/message/get-pending` | 用户待处理消息 | all |
| POST | `/api/v1/system/user/self/message/handle` | 用户消息标记已处理 | all |
| POST | `/api/v1/system/user/self/message/mark-all-read` | 用户消息全部已读 | all |
| POST | `/api/v1/system/user/self/message/multi-delete` | 用户消息批量删除 | all |
| POST | `/api/v1/system/user/self/message/multi-is-read` | 用户消息批量已读 | all |
| POST | `/api/v1/system/user/self/message/statistics` | 用户消息统计 | all |
| POST | `/api/v1/system/user/self/notify-preference/read` | 用户通知偏好读取 | all |
| POST | `/api/v1/system/user/self/notify-preference/update` | 用户通知偏好更新 | all |
| POST | `/api/v1/system/user/self/openclaw/setup-check` | 查询 CLI 绑定状态 | public |
| POST | `/api/v1/system/user/self/openclaw/setup-complete` | 完成 CLI 绑定 | all |
| POST | `/api/v1/system/user/self/profile/get-list` | 获取用户配置列表 | all |
| POST | `/api/v1/system/user/self/profile/get-one` | 获取用户配置详情 | all |
| POST | `/api/v1/system/user/self/profile/update` | 更新用户配置 | all |
| POST | `/api/v1/system/user/self/register` | 普通用户注册 | public |
| POST | `/api/v1/system/user/self/resource/action/get-list` | 获取用户资源动作权限列表 | all |
| POST | `/api/v1/system/user/self/tenant/delete` | 退出当前租户 | all |
| POST | `/api/v1/system/user/self/tenant/get-list` | 获取用户所处的租户列表 | all |
| POST | `/api/v1/system/user/self/tenant/get-one` | 获取当前用户在当前租户的详情 | all |
| POST | `/api/v1/system/user/self/tenant/join` | 用户加入租户（通过邀请码、邮件或手机邀请） | all |
| POST | `/api/v1/system/user/self/tenant/update` | 更新当前用户在当前租户的信息 | all |
| POST | `/api/v1/system/user/self/third-auth/start` | 第三方登录授权起跳 | public |
| POST | `/api/v1/system/user/self/third-login` | 第三方登录回调换平台登录态 | public |
| POST | `/api/v1/system/user/self/third-register` | 第三方补全注册 | public |
| POST | `/api/v1/system/user/self/update` | 更新用户基本数据 | all |
| POST | `/api/v1/system/user/self/user/search` | 精准搜索用户 | all |
| POST | `/api/v1/system/user/tenant/get-list` | 用户所处的租户列表 | all |

<!-- END_API_LIST -->

权限: mixed

| 端点 | 说明 | 权限 |
|---|------|------|
| POST /api/v1/system/access/api/create | 添加接口 | 平台管理员 |
| POST /api/v1/system/access/api/delete | 删除接口 | 平台管理员 |
| POST /api/v1/system/access/api/get-list | 获取接口列表 | 平台管理员 |
| POST /api/v1/system/access/api/update | 更新接口 | 平台管理员 |
| POST /api/v1/system/access/info/batch-import | 批量导入授权 | 管理员 |
| POST /api/v1/system/access/info/create | 添加访问权限信息 | 平台管理员 |
| POST /api/v1/system/access/info/delete | 删除访问权限信息 | 平台管理员 |
| POST /api/v1/system/access/info/get-list | 获取访问权限信息列表 | 平台管理员 |
| POST /api/v1/system/access/info/tree | 获取访问权限信息树结构 | 平台管理员 |
| POST /api/v1/system/access/info/update | 更新访问权限信息 | 平台管理员 |
| POST /api/v1/system/check-in/do | 用户签到 | 管理员 |
| POST /api/v1/system/check-in/get-list | 签到记录列表 | 管理员 |
| POST /api/v1/system/check-in/point-balance/get | 获取当前用户积分余额 | 管理员 |
| POST /api/v1/system/check-in/point-log/adjust | 管理员调整积分 | 管理员 |
| POST /api/v1/system/check-in/point-log/get-list | 积分流水列表 | 管理员 |
| POST /api/v1/system/dept/info/create | 添加部门详情 | 管理员 |
| POST /api/v1/system/dept/info/delete | 删除部门 | 管理员 |
| POST /api/v1/system/dept/info/get-list | 获取部门列表 | 管理员 |
| POST /api/v1/system/dept/info/get-one | 获取部门详情 | 管理员 |
| POST /api/v1/system/dept/info/update | 更新部门 | 管理员 |
| POST /api/v1/system/dept/sync-job/create | 添加同步任务 | 管理员 |
| POST /api/v1/system/dept/sync-job/delete | 删除同步任务 | 管理员 |
| POST /api/v1/system/dept/sync-job/execute | 执行同步任务 | 管理员 |
| POST /api/v1/system/dept/sync-job/get-list | 获取同步任务列表 | 管理员 |
| POST /api/v1/system/dept/sync-job/get-one | 获取同步任务详情 | 管理员 |
| POST /api/v1/system/dept/sync-job/update | 更新同步任务 | 管理员 |
| POST /api/v1/system/dept/user/batch-create | 批量授权部门用户 | 管理员 |
| POST /api/v1/system/dept/user/batch-delete | 批量取消授权部门用户 | 管理员 |
| POST /api/v1/system/dept/user/get-list | 获取部门授权列表 | 管理员 |
| POST /api/v1/system/dict/detail/create | 添加字典详情 | 平台管理员 |
| POST /api/v1/system/dict/detail/delete | 删除字典详情 | 平台管理员 |
| POST /api/v1/system/dict/detail/get-list | 获取字典详情列表 | 管理员 |
| POST /api/v1/system/dict/detail/get-one | 获取字典详情单个 | 管理员 |
| POST /api/v1/system/dict/detail/update | 更新字典详情 | 平台管理员 |
| POST /api/v1/system/dict/info/batch-export | 批量导出字典信息 | 平台管理员 |
| POST /api/v1/system/dict/info/batch-import | 批量导入字典信息 | 平台管理员 |
| POST /api/v1/system/dict/info/create | 添加字典信息 | 平台管理员 |
| POST /api/v1/system/dict/info/delete | 删除字典信息 | 平台管理员 |
| POST /api/v1/system/dict/info/get-list | 获取字典信息列表 | 管理员 |
| POST /api/v1/system/dict/info/get-one | 获取字典信息详情 | 管理员 |
| POST /api/v1/system/dict/info/update | 更新字典信息 | 平台管理员 |
| POST /api/v1/system/job/task/cancel | 取消执行任务 | 平台管理员 |
| POST /api/v1/system/job/task/group/create | 新建任务分组 | 平台管理员 |
| POST /api/v1/system/job/task/group/delete | 删除任务分组 | 平台管理员 |
| POST /api/v1/system/job/task/group/get-list | 获取任务分组列表 | 平台管理员 |
| POST /api/v1/system/job/task/group/get-one | 获取任务分组详情 | 平台管理员 |
| POST /api/v1/system/job/task/group/update | 更新任务分组 | 平台管理员 |
| POST /api/v1/system/job/task/info/create | 创建任务 | 平台管理员 |
| POST /api/v1/system/job/task/info/delete | 删除任务 | 平台管理员 |
| POST /api/v1/system/job/task/info/get-list | 获取任务列表 | 平台管理员 |
| POST /api/v1/system/job/task/info/get-one | 获取任务详情 | 平台管理员 |
| POST /api/v1/system/job/task/info/start | 启动任务 | 平台管理员 |
| POST /api/v1/system/job/task/info/stop | 停止任务 | 平台管理员 |
| POST /api/v1/system/job/task/info/update | 更新任务 | 平台管理员 |
| POST /api/v1/system/job/task/send | 发送延时请求 | 平台管理员 |
| POST /api/v1/system/log/login/get-list | 获取登录日志列表 | 管理员 |
| POST /api/v1/system/log/oper/get-list | 获取操作日志列表 | 管理员 |
| POST /api/v1/system/notify/config/create | 添加通知配置 | 平台管理员 |
| POST /api/v1/system/notify/config/delete | 删除通知配置 | 平台管理员 |
| POST /api/v1/system/notify/config/get-list | 获取通知列表 | 管理员 |
| POST /api/v1/system/notify/config/get-one | 获取通知详情 | 管理员 |
| POST /api/v1/system/notify/config/send-test | 测试发送通知 | 管理员 |
| POST /api/v1/system/notify/config/template/delete | 删除通知绑定模版 | 管理员 |
| POST /api/v1/system/notify/config/template/get-list | 获取通知绑定模版列表 | 管理员 |
| POST /api/v1/system/notify/config/template/update | 更新通知绑定模版 | 管理员 |
| POST /api/v1/system/notify/config/tree | 获取通知列表树 | 管理员 |
| POST /api/v1/system/notify/config/update | 更新通知配置 | 平台管理员 |
| POST /api/v1/system/notify/message/info/delete | 删除消息 | 管理员 |
| POST /api/v1/system/notify/message/info/get-list | 获取消息列表 | 管理员 |
| POST /api/v1/system/notify/message/info/send | 发送消息 | 管理员 |
| POST /api/v1/system/notify/message/info/update | 更新消息 | 管理员 |
| POST /api/v1/system/notify/news/create | 创建资讯 | 管理员 |
| POST /api/v1/system/notify/news/delete | 删除资讯 | 管理员 |
| POST /api/v1/system/notify/news/info | 获取资讯详情 | 管理员 |
| POST /api/v1/system/notify/news/list | 获取资讯列表 | 管理员 |
| POST /api/v1/system/notify/news/update | 更新资讯 | 管理员 |
| POST /api/v1/system/notify/notification/copy | 复制手动通知 | 管理员 |
| POST /api/v1/system/notify/notification/create | 创建手动通知 | 管理员 |
| POST /api/v1/system/notify/notification/delete | 删除手动通知 | 管理员 |
| POST /api/v1/system/notify/notification/estimate-users | 预估送达用户数 | 管理员 |
| POST /api/v1/system/notify/notification/index | 获取手动通知列表 | 管理员 |
| POST /api/v1/system/notify/notification/read | 获取手动通知详情 | 管理员 |
| POST /api/v1/system/notify/notification/revoke | 撤回手动通知 | 管理员 |
| POST /api/v1/system/notify/notification/send | 发送手动通知 | 管理员 |
| POST /api/v1/system/notify/notification/update | 更新手动通知 | 管理员 |
| POST /api/v1/system/notify/template/create | 添加通知模版 | 管理员 |
| POST /api/v1/system/notify/template/delete | 删除通知模版 | 管理员 |
| POST /api/v1/system/notify/template/get-list | 获取通知模版列表 | 管理员 |
| POST /api/v1/system/notify/template/get-one | 获取通知模版详情 | 管理员 |
| POST /api/v1/system/notify/template/update | 更新通知模版 | 管理员 |
| POST /api/v1/system/ops/feedback/create | 添加帮助与反馈 | 所有用户 |
| POST /api/v1/system/ops/feedback/get-list | 获取帮助与反馈 | 所有用户 |
| POST /api/v1/system/ops/feedback/update | 更新帮助与反馈 | 所有用户 |
| POST /api/v1/system/ops/work-order/create | 添加工单 | 所有用户 |
| POST /api/v1/system/ops/work-order/get-list | 获取工单列表 | 所有用户 |
| POST /api/v1/system/ops/work-order/update | 更新工单 | 所有用户 |
| POST /api/v1/system/role/access/batch-update | 更新操作权限 | 管理员 |
| POST /api/v1/system/role/access/get-list | 获取操作权限列表 | 管理员 |
| POST /api/v1/system/role/app/batch-update | 更新APP权限 | 管理员 |
| POST /api/v1/system/role/app/get-list | 获取APP权限列表 | 管理员 |
| POST /api/v1/system/role/info/create | 添加角色 | 管理员 |
| POST /api/v1/system/role/info/delete | 删除角色 | 管理员 |
| POST /api/v1/system/role/info/get-list | 获取角色列表 | 管理员 |
| POST /api/v1/system/role/info/update | 更新角色 | 管理员 |
| POST /api/v1/system/role/menu/batch-update | 更新角色对应菜单列表 | 管理员 |
| POST /api/v1/system/role/menu/get-list | 获取角色对应菜单列表 | 管理员 |
| POST /api/v1/system/user/data/area/get-list | 获取区域权限列表 | 管理员 |
| POST /api/v1/system/user/data/project/get-list | 获取项目权限列表 | 管理员 |
| POST /api/v1/system/user/dept/batch-create | 新增用户的部门列表 | 管理员 |
| POST /api/v1/system/user/dept/batch-delete | 删除用户的部门列表 | 管理员 |
| POST /api/v1/system/user/info/create | 创建用户信息 | 管理员 |
| POST /api/v1/system/user/info/delete | 刪除用户 | 管理员 |
| POST /api/v1/system/user/info/get-list | 查询用户信息列表 | 管理员 |
| POST /api/v1/system/user/info/get-one | 获取用户信息 | 管理员 |
| POST /api/v1/system/user/info/update | 更新用户基本数据 | 管理员 |
| POST /api/v1/system/user/self/access-token/create | 创建访问令牌 | 所有用户 |
| POST /api/v1/system/user/self/access-token/delete | 删除访问令牌 | 所有用户 |
| POST /api/v1/system/user/self/access-token/get-list | 获取访问令牌列表 | 所有用户 |
| POST /api/v1/system/user/self/access-token/get-one | 获取访问令牌详情 | 所有用户 |
| POST /api/v1/system/user/self/access-token/update | 更新访问令牌 | 所有用户 |
| POST /api/v1/system/user/self/access/tree | 获取用户授权树 | 所有用户 |
| POST /api/v1/system/user/self/app/get-list | 获取用户应用列表 | 所有用户 |
| POST /api/v1/system/user/self/app/get-one | 获取用户应用详情 | 所有用户 |
| POST /api/v1/system/user/self/bind-account | 绑定账号 | 所有用户 |
| POST /api/v1/system/user/self/cancel | 注销用户 | 所有用户 |
| POST /api/v1/system/user/self/captcha | 获取验证码 | public |
| POST /api/v1/system/user/self/change-pwd | 更新用户密码 | 所有用户 |
| POST /api/v1/system/user/self/forget-pwd | 忘记密码 | public |
| POST /api/v1/system/user/self/get-one | 获取用户信息 | 所有用户 |
| POST /api/v1/system/user/self/login | 用户登录 | public |
| POST /api/v1/system/user/self/logout | 用户登出 | 所有用户 |
| POST /api/v1/system/user/self/menu/get-list | 获取用户菜单列表 | 所有用户 |
| POST /api/v1/system/user/self/message/get-list | 用户消息列表 | 所有用户 |
| POST /api/v1/system/user/self/message/get-pending | 用户待处理消息 | 所有用户 |
| POST /api/v1/system/user/self/message/handle | 用户消息标记已处理 | 所有用户 |
| POST /api/v1/system/user/self/message/mark-all-read | 用户消息全部已读 | 所有用户 |
| POST /api/v1/system/user/self/message/multi-delete | 用户消息批量删除 | 所有用户 |
| POST /api/v1/system/user/self/message/multi-is-read | 用户消息批量已读 | 所有用户 |
| POST /api/v1/system/user/self/message/statistics | 用户消息统计 | 所有用户 |
| POST /api/v1/system/user/self/notify-preference/read | 用户通知偏好读取 | 所有用户 |
| POST /api/v1/system/user/self/notify-preference/update | 用户通知偏好更新 | 所有用户 |
| POST /api/v1/system/user/self/profile/get-list | 获取用户配置列表 | 所有用户 |
| POST /api/v1/system/user/self/profile/get-one | 获取用户配置详情 | 所有用户 |
| POST /api/v1/system/user/self/profile/update | 更新用户配置 | 所有用户 |
| POST /api/v1/system/user/self/register | 普通用户注册 | public |
| POST /api/v1/system/user/self/tenant/delete | 退出当前租户 | 所有用户 |
| POST /api/v1/system/user/self/tenant/get-list | 获取用户所处的租户列表 | 所有用户 |
| POST /api/v1/system/user/self/tenant/get-one | 获取当前用户在当前租户的详情 | 所有用户 |
| POST /api/v1/system/user/self/tenant/join | 用户加入租户（通过邀请码、邮件或手机邀请） | 所有用户 |
| POST /api/v1/system/user/self/tenant/update | 更新当前用户在当前租户的信息 | 所有用户 |
| POST /api/v1/system/user/self/update | 更新用户基本数据 | 所有用户 |
| POST /api/v1/system/user/self/user/search | 精准搜索用户 | 所有用户 |
| POST /api/v1/system/user/tenant/get-list | 用户所处的租户列表 | 所有用户 |


## 典型业务场景

### 用户登录与个人信息

**场景描述**：登录 / 获取个人信息 / 修改密码

**涉及 API**：
- `/api/v1/system/user/self/login`
- `/api/v1/system/user/self/get-one`
- `/api/v1/system/user/self/change-pwd`

**工作流**：
1. 用户名密码登录
2. 获取个人信息
3. 修改密码

### 访问令牌管理

**场景描述**：创建 API 访问令牌 / 查看令牌列表（AccessKey/Secret 用于 JWT 构造）

**涉及 API**：
- `/api/v1/system/user/self/access-token/create`
- `/api/v1/system/user/self/access-token/get-list`

**工作流**：
1. 创建令牌（获取 AccessKey/Secret）
2. 构造 JWT: {userID, tenantCode, accessKey, exp}
3. 使用 Authorization: Bearer <jwt> 调用 API

### 通知配置与发送

**场景描述**：配置设备告警通知 / 发送批量通知（防抖机制防止告警风暴）

**涉及 API**：
- `/api/v1/system/notify/config/create`
- `/api/v1/system/notify/template/create`
- `/api/v1/system/notify/message/send`

**工作流**：
1. 创建事件配置（code、throttle 防抖间隔）
2. 创建通知模板（按渠道：短信/邮件/钉钉）
3. 触发事件发送通知

### 角色与权限管理

**场景描述**：创建角色 / 分配权限 / 分配用户

**涉及 API**：
- `/api/v1/system/role/info/create`
- `/api/v1/system/role/access/update`
- `/api/v1/system/role/user/batch-create`

**工作流**：
1. 创建角色
2. 配置角色权限
3. 为用户分配角色


## 常用工作流

### 获取当前用户信息

```bash
ur api /api/v1/system/user/self/get-one \
  --body '{}'
```


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
