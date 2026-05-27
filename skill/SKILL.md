---
name: ur-api
description: "Use when calling 联犀 SaaS 平台 API: device management, user management, product management, tenant management, AI management, project management, system management, or IoT device operations. triggers: API调用, 设备列表, IoT设备控制, 设备注册, 访问令牌, 物模型管理, OTA升级, 场景联动, 权限配置, swagger schema, 项目管理, 区域管理, 系统管理, 问题反馈, 提交反馈, 反馈类型, 使用问题, 业务受损, 业务不可用"
---

# ur-api — 联犀 SaaS 平台 API 工具

当前主实现位于 `backend/cli/ur`，通过 Go CLI `ur` 提供能力；runtime / AI 调用只走 `UR_*` 环境变量，不再依赖 profile / `~/.ur/config.json`。

所有接口均为 POST 方法。请求格式 `{code, msg, data}`。

---

## 角色权限区分

**全栈权限总索引**（契约 / 运行时 / 前端 / 消歧）：仓库根目录 `docs/中台/功能说明/权限体系总览.md`。

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| **平台管理员** | 跨租户操作、租户 CRUD、查看所有租户 | 创建租户、管理租户配置 |
| **租户管理员** | 本租户内 CRUD、用户管理、角色分配 | 创建设备、管理用户、配置项目 |
| **普通用户** | 个人信息、设备分享、查看权限范围内数据 | 登录、修改密码、分享设备 |

### 按角色区分的域

| 域 | 平台管理员 | 租户管理员 | 普通用户 |
|---|-----------|-----------|---------|
| ur-tenant | 创建/删除租户、查看所有租户 | 管理本租户、邀请用户 | 无权限 |
| ur-device | 无权限 | 设备 CRUD、属性控制 | 设备分享、收藏 |
| ur-user | 无权限 | 用户 CRUD、角色管理 | 个人信息、修改密码 |
| ur-product | 无权限 | 产品 CRUD、物模型管理 | 无权限 |
| ur-project | 无权限 | 项目 CRUD、区域管理 | 查看权限内项目 |
| ur-ai | 无权限 | Agent 配置、告警管理 | 无权限 |

### 当前用户身份

调用任何 API 前，**先运行 check 确认角色**，错误的角色会导致 403：

```bash
ur check
```

前端或自动化工具需要区分角色时，应以 **HTTP 接口返回结构** 为准：

- **端点**：`POST /api/v1/system/user/self/get-one`
- **组织侧管理员线索**：`tenants[]` → `UserTenant` 上的 `isTenantOwner` 及 `roles[]`
- **角色编码**：`UserTenant.roles` 的元素类型 `RoleInfo` 中的 `code`（如 `admin` / `client` / `supper`）
- **平台侧上下文**：结合请求头 `tenant-code` 与 `tenantCode` + `roles[].code` 综合判断

**类型定义单一事实来源**：
- `UserInfo` / `UserTenant`：`backend/core/service/apisvr/http/system/user/info.api`
- `RoleInfo`：`backend/core/service/apisvr/http/system/role.api`

> **说明**：`ur-user/SKILL.md` 由 `ur generate-skills` 自动生成，字段级约定统一写在本主 `SKILL.md`，避免与子域生成文件重复且被覆盖。

---

## AI 快速决策

### 第一步：确认当前用户角色

```bash
ur check
```

| 角色 | 可操作范围 | 配置方式 |
|------|----------|---------|
| **平台管理员** | 所有域（platform/admin/all 接口） | `ur login` 用平台管理员账号授权 |
| **租户管理员** | 本租户 CRUD（admin/all 接口） | `ur login` 用租户管理员账号授权 |
| **普通用户** | 个人信息、设备分享（仅 all 接口） | `ur login` 用普通用户账号授权 |

> **切换角色**：重新运行 `ur login`（用不同账号授权），或用 `--app` 切换应用上下文：
> ```bash
> ur --app platform-manage api /api/v1/...
> ```

### 第二步：根据任务选择子域

| 用户意图 / 关键词 | 加载子域 | 最低所需角色 |
|----------------|---------|------------|
| 设备列表、设备控制、属性上报、MQTT、三元组、网关、OTA升级设备 | `ur-device` | admin（控制/CRUD），user（分享/收藏） |
| 产品定义、物模型、协议脚本、OTA固件包、品类 | `ur-product` | admin |
| 项目、区域、分组、数据权限申请 | `ur-project` | admin（管理），user（申请权限） |
| 登录、用户信息、角色、部门、字典、通知、访问令牌、修改密码 | `ur-user` | admin（CRUD），user（自身信息） |
| 文件上传、WebSocket、批量接口、应用管理、Hook | `ur-system` | user（上传/WS），admin（应用），platform（全局应用） |
| 创建租户、租户列表、租户配置、邀请用户加入租户 | `ur-tenant` | **platform**（创建/查看全部），admin（管理本租户） |
| AI Agent、场景联动、告警规则、数字分身、AI会话 | `ur-ai` | admin |

### 第三步：IoT AI 工具调用设计

| 用户意图 / 关键词 | 加载子域 | 说明 |
|----------------|---------|------|
| MCP 工具、frontend tools、设备工具语义、物模型调用顺序、实时属性查询 | `ur-iot-device` | 关注设备工具本体与工具语义对齐 |
| 所有设备可见、设备上下文、session 绑定、工具作用目标、前端不负责绑定设备 | `ur-iot-context` | 关注 Web 前端全设备可见场景下的上下文约束 |
| Web、Win AI、MQTT、UDP、客户端差异、共存迁移 | `ur-iot-client` | 关注多客户端执行链路差异 |

版本说明见：`references/iot-tools-versions.md`

> **角色不足时不要尝试调用**：如果 check 显示为 `user` 但任务需要 `admin`，直接告知用户切换角色，不要尝试然后 403。

### 数量 / 列表类问题

- 如果用户问"有几个""多少个""当前有哪些"这类业务数据问题，优先直接调用真实的 `get-list` 接口。
- 数量统一优先从返回的 `data.total` 读取；不要自己数当前页 `list.length`，除非接口根本没有 `total`。
- 不要先调用 `schema` 再决定是否查询数据；`schema` 只用于查看已知接口的字段结构，不用于查询业务结果。

### API 路径选择与 404 回退策略（强制）

**绝对禁止臆造不存在的 API 路径。** 以下行为严格禁止：

1. **不要拼接或猜测路径**：严禁根据自然语言描述自行拼接 `/api/v1/...` 路径。
2. **不要假设近义词路径**：设备管理域的真实接口以 `/api/v1/things/device/info/...` 和 `/api/v1/things/device/interact/...` 等为准。
3. **404 后必须停止猜测**：如果一次 API 调用返回 `404 page not found`，**必须立即停止继续尝试其他相似路径**。正确的回退顺序是：
   - 第一步：调用 `ur schema <已知前缀>` 查看该前缀下真实存在的接口列表。
   - 第二步：查阅本 SKILL 的「高频端点速查」或各子域 SKILL.md 中的端点列表。
   - 第三步：如果仍不确定，调用 `knowledge_search` 检索知识库中关于该业务的真实接口文档。
4. **优先使用已知真实路径**：只有在本 SKILL 中明确列出的路径、或通过 `schema` 确认存在的路径、或各子域 SKILL.md 中记录的路径，才是可调用的真实路径。

**不确定时**：用 schema 命令确认路径归属：
```bash
ur schema /api/v1/things/device/info/get-list
```

---

## API 域划分

共 7 个子域，覆盖 492 个端点（100% 覆盖）：

| 域 | 端点数 | 说明 |
|---|--------|------|
| ur-device | 85 | 设备管理：设备CRUD、状态、属性控制、认证、物模型 |
| ur-product | 86 | 产品管理：产品CRUD、物模型、协议脚本、OTA |
| ur-project | 41 | 项目管理：项目、区域、分组、数据权限、用户权限 |
| ur-user | 153 | 用户管理：登录、角色、部门、字典、任务、通知、日志 |
| ur-system | 53 | 系统管理：通用接口、应用管理、授权码、Hook能力 |
| ur-tenant | 38 | 租户管理：租户CRUD、用户管理、权限配置 |
| ur-ai | 36 | AI管理：Agent、告警、规则引擎、场景联动 |

详细端点列表见各子域 SKILL.md：
- `ur-device/SKILL.md`
- `ur-product/SKILL.md`
- `ur-project/SKILL.md`
- `ur-user/SKILL.md`
- `ur-system/SKILL.md`
- `ur-tenant/SKILL.md`
- `ur-ai/SKILL.md`

IoT AI 工具迁移相关子域：
- `ur-iot-device/SKILL.md`
- `ur-iot-context/SKILL.md`
- `ur-iot-client/SKILL.md`

设备数据分析子域：
- `ur-device-analytics/SKILL.md` — 属性历史查询、趋势分析、聚合统计、报表生成（物模型驱动）

设备调试子域：
- `ur-device-debug/SKILL.md` — 设备日志查询（属性/事件/命令/上下线/异常/诊断/SDK）、实时调试（属性控制/行为调用/事件发送/Mock数据）

---

## 快速开始

### 方式 1：Device Auth（推荐，默认）

```bash
# 1. 运行 login，CLI 生成 setup code 并输出授权 URL
ur login

# 2. 在浏览器中打开 URL，登录控制台后创建/选择访问令牌
# 3. 点击「完成 CLI 绑定」
# 4. CLI 自动轮询获取 AK/SK，保存到 ~/.ur/config.json

# 验证连通性
ur check

# 调用 API
ur api /api/v1/system/user/self/get-one
```

### 方式 2：AccessKey/JWT（程序化访问）

```bash
# 交互式配置
ur setup

# 或手动编辑 ~/.ur/config.json
```

需要先在平台 UI 创建访问令牌：用户设置 → 访问令牌 → 创建。

### 方式 3：环境变量

```bash
export UR_BASE_URL=http://localhost:7777
export UR_APP_ID=77
export UR_ACCESS_KEY=xxx
export UR_ACCESS_SECRET=xxx
export UR_USER_ID=12345
export UR_TOKEN=xxx
export UR_APP=iot
```

---

## 参考文档

| 文档 | 说明 |
|------|------|
| `references/auth.md` | 认证方式详解（Device Auth / JWT / 环境变量） |
| `references/cli-usage.md` | CLI 命令用法（全局选项、API调用、物模型、场景联动、协议脚本、schema） |
| `references/api-conventions.md` | API 通用约定（请求格式、响应格式、分页、权限标注） |
| `references/troubleshooting.md` | 常见问题排查（登录失败、401/403、连接问题、问题反馈） |
| `references/quick-reference.md` | 高频端点速查（Top 15 最常用端点） |
| `references/workflows.md` | 跨域工作流示例（设备上线、租户初始化、告警通知） |
| `references/maintenance.md` | 新增接口维护规范（检查清单、权限标注指南） |
| `references/frontend-permissions.md` | 前端应用权限映射（平台应用 / 组织应用 / 公共应用） |

详细 API 端点见各子域 SKILL.md：ur-device、ur-product、ur-project、ur-user、ur-system、ur-tenant、ur-ai
