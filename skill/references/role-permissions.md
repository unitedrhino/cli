# 角色权限与用户身份判定

本文件是 `SKILL.md`（ur-api 主入口）的子文档。主文档保留「先 check」入口与导航，本文件收录角色总览、按角色区分的域、可操作范围与接口侧身份判定细节。

**全栈权限总索引**（契约 / 运行时 / 前端 / 消歧）：仓库根目录 `docs/中台/功能说明/权限体系总览.md`。

## 角色总览

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| **平台管理员** | 跨企业操作、企业 CRUD、查看所有企业 | 创建企业、管理企业配置 |
| **企业管理员** | 本企业内 CRUD、用户管理、角色分配 | 创建设备、管理用户、配置项目 |
| **普通用户** | 个人信息、设备分享、查看权限范围内数据 | 登录、修改密码、分享设备 |

## 按角色区分的域

| 域 | 平台管理员 | 企业管理员 | 普通用户 |
|---|-----------|-----------|---------|
| ur-tenant | 创建/删除企业、查看所有企业 | 管理本企业、邀请用户 | 无权限 |
| ur-device | 无权限 | 设备 CRUD、属性控制 | 设备分享、收藏 |
| ur-user | 无权限 | 用户 CRUD、角色管理 | 个人信息、修改密码 |
| ur-product | 无权限 | 产品 CRUD、物模型管理 | 无权限 |
| ur-project | 无权限 | 项目 CRUD、区域管理 | 查看权限内项目 |
| ur-ai | 无权限 | Agent 配置、告警管理 | 无权限 |
| ur-view | 无权限 | 大屏 CRUD、画布编辑、发布管理、素材管理 | 无权限 |

## 角色与可操作范围（AI 快速决策第一步）

| 角色 | 可操作范围 | 配置方式 |
|------|----------|---------|
| **平台管理员** | 所有域（platform/admin/all 接口） | 先复用环境/profile；必要时用 `ur login` 切换认证 |
| **企业管理员** | 本企业 CRUD（admin/all 接口） | 先复用环境/profile；必要时用 `ur login` 切换认证 |
| **普通用户** | 个人信息、设备分享（仅 all 接口） | 先复用环境/profile；必要时用 `ur login` 切换认证 |

> **切换角色**：重新运行 `ur login`（用不同账号授权），或用 `--app` 切换应用上下文：
> ```bash
> ur --app platform-manage api /api/v1/...
> ```

> **角色不足时不要尝试调用**：如果 check 显示为 `user` 但任务需要 `admin`，直接告知用户切换角色，不要尝试然后 403。

## 当前用户身份（接口侧判定）

前端或自动化工具需要区分角色时，应以 **HTTP 接口返回结构** 为准：

- **端点**：`POST /api/v1/system/user/self/get-one`
- **组织侧管理员线索**：`tenants[]` → `UserTenant` 上的 `isTenantOwner` 及 `roles[]`
- **角色编码**：`UserTenant.roles` 的元素类型 `RoleInfo` 中的 `code`（如 `admin` / `client` / `supper`）
- **平台侧上下文**：结合请求头 `tenant-code` 与 `tenantCode` + `roles[].code` 综合判断

**类型定义单一事实来源**：
- `UserInfo` / `UserTenant`：`backend/core/service/apisvr/http/system/user/info.api`
- `RoleInfo`：`backend/core/service/apisvr/http/system/role.api`

> **说明**：`ur-org-manage/ur-user/SKILL.md` 由 `ur generate-skills` 自动生成，字段级约定统一写在本主 `SKILL.md`，避免与子域生成文件重复且被覆盖。
