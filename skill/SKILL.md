---
name: ur-api
description: "Use when calling 联犀 SaaS 平台 API: device management, user management, product management, tenant management, AI management, project management, system management, or IoT device operations. triggers: API调用, 设备列表, IoT设备控制, 设备注册, 访问令牌, 物模型管理, OTA升级, 场景联动, 权限配置, swagger schema, 项目管理, 区域管理, 系统管理"
---
# ur-api — 联犀 SaaS 平台 API 工具

## 概述

当前主实现位于 `backend/cli/ur`，通过 Go CLI `ur` 提供能力；runtime / AI 调用只走 `UR_*` 环境变量，不再依赖 profile / `~/.ur/config.json`。

支持两种认证方式调用平台 API：

1. **Device Auth**（推荐）：`ur login` → 浏览器授权 → 获取 AccessKey/Secret → JWT → `Authorization: Bearer` header。
2. **AccessKey/JWT**：直接配置 AccessKey + AccessSecret → HS256 JWT → `Authorization: Bearer` header。

所有接口均为 POST 方法。请求格式 `{code, msg, data}`。

## 角色权限区分

**全栈权限总索引**（契约 / 运行时 / 前端 / 消歧）：仓库根目录 `docs/中台/功能说明/权限体系总览.md`。

本工具按调用者角色区分 API 权限，各域 SKILL.md 包含「角色权限说明」章节：

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

### 当前用户身份：`user/self/get-one` 与 .api 单一事实来源

前端或自动化工具需要区分 **平台管理员 / 组织（租户）管理员 / 普通用户** 时，应以 **HTTP 接口返回结构** 为准，不要只依赖本地 profile 配置，运行时以 `UR_*` 环境变量为准。

- **端点**：`POST /api/v1/system/user/self/get-one`（请求体可含 `withTenant`，定义见 `backend/core/service/apisvr/http/system/user/self/types.api` 中 `UserSelfGetOneReq`）。
- **响应类型**：`UserInfo`，其中与权限上下文相关的主要嵌套为 **`tenants[]` → `UserTenant`**。
- **组织侧管理员线索**：在同一 `tenantCode` 下的 `UserTenant` 上查看 **`isTenantOwner`**（是否组织管理员）及 **`roles[]`**。
- **角色编码**：`UserTenant.roles` 的元素类型为 **`RoleInfo`**，其中 **`code`**（如初始化中的 `admin` / `client` / `supper` 等，以环境实际数据为准）用于区分角色类型。
- **平台侧上下文**：结合当前请求头 **`tenant-code`**（例如平台类应用固定为 `platform`，见下文「SaaS 应用权限配置」）与上述 **`tenantCode` + `roles[].code`** 综合判断；**响应体没有单独的 `isPlatformAdmin` 布尔字段**，不要在文档中虚构该字段。

**类型定义位置（单一事实来源，优先于本 SKILL 的文字描述）**：

- `UserInfo` / `UserTenant`：`backend/core/service/apisvr/http/system/user/info.api`
- `RoleInfo`：`backend/core/service/apisvr/http/system/role.api`

**前端能力枚举（完整表）**：推导用能力 key、与 `authType` 的对应、Vben 衔接说明见仓库根目录 `apps/web/docs/capability-keys.md`。

> **说明**：`ur-user/SKILL.md` 由 `ur generate-skills` 自动生成，字段级约定统一写在本主 `SKILL.md`，避免与子域生成文件重复且被覆盖。

## AI 快速决策

### 第一步：确认当前用户角色

调用任何 API 前，**先运行 check 确认角色**，错误的角色会导致 403：

```bash
ur check
```

输出示例：
```
[OK] 当前 profile: prod-admin
[OK] 当前角色: 租户管理员（可调用 admin/all 权限接口）
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

### 第三步：如果任务是 IoT AI 工具调用设计，再继续细分

以下内容不单独作为顶层 Skill，而是由 `ur-api` 第一层继续路由，避免多个顶层 Skill 语义重叠：

| 用户意图 / 关键词 | 加载子域 | 说明 |
|----------------|---------|------|
| MCP 工具、frontend tools、设备工具语义、物模型调用顺序、实时属性查询 | `ur-iot-device` | 关注设备工具本体与工具语义对齐 |
| 所有设备可见、设备上下文、session 绑定、工具作用目标、前端不负责绑定设备 | `ur-iot-context` | 关注 Web 前端全设备可见场景下的上下文约束 |
| Web、Win AI、MQTT、UDP、客户端差异、共存迁移 | `ur-iot-client` | 关注多客户端执行链路差异 |

版本说明见：`references/iot-tools-versions.md`

> **角色不足时不要尝试调用**：如果 check 显示为 `user` 但任务需要 `admin`，直接告知用户切换角色，不要尝试然后 403。

### 数量 / 列表类问题

- 如果用户问“有几个”“多少个”“当前有哪些”这类业务数据问题，优先直接调用真实的 `get-list` 接口。
- 数量统一优先从返回的 `data.total` 读取；不要自己数当前页 `list.length`，除非接口根本没有 `total`。
- 不要先调用 `schema` 再决定是否查询数据；`schema` 只用于查看已知接口的字段结构，不用于查询业务结果。
- 角色数量 / 角色列表常见真实路径：`/api/v1/system/role/info/get-list`
- 应用列表常见真实路径：
  - 平台应用：`/api/v1/system/app/info/get-list`
  - 当前用户可见应用：`/api/v1/system/user/self/app/get-list`

### API 路径选择与 404 回退策略（强制）

**绝对禁止臆造不存在的 API 路径。** 以下行为严格禁止：

1. **不要拼接或猜测路径**：严禁根据自然语言描述自行拼接 `/api/v1/...` 路径。例如用户提到“设备接入配置”，**严禁**直接调用 `/api/v1/things/device/access/config/get-list` 或 `/api/v1/things/device/access/config`，这些路径不存在。
2. **不要假设近义词路径**：设备管理域的真实接口以 `/api/v1/things/device/info/...` 和 `/api/v1/things/device/interact/...` 等为准，不要为“接入”“配置”“绑定”等概念编造新路径。
3. **404 后必须停止猜测**：如果一次 API 调用返回 `404 page not found`，**必须立即停止继续尝试其他相似路径**。正确的回退顺序是：
   - 第一步：调用 `ur schema <已知前缀>` 查看该前缀下真实存在的接口列表（例如 `ur schema /api/v1/things/device/`）。
   - 第二步：查阅本 SKILL 的「高频端点速查」或各子域 SKILL.md 中的端点列表。
   - 第三步：如果仍不确定，调用 `knowledge_search` 检索知识库中关于该业务的真实接口文档。
   - **严禁在 404 后继续凭直觉尝试另一个相似路径。**
4. **优先使用已知真实路径**：只有在本 SKILL 中明确列出的路径、或通过 `schema` 确认存在的路径、或各子域 SKILL.md 中记录的路径，才是可调用的真实路径。

**多域任务**：先加载主域，按需引用次域。例如"设备告警+通知" = ur-ai（告警规则）+ ur-user（通知配置）。

**不确定时**：用 schema 命令确认路径归属：
```bash
ur schema /api/v1/things/device/
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

## 快速开始

### 方式 1：Device Auth（推荐，默认）

通过浏览器完成授权，无需手动输入账号密码：

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

已有 AccessKey/Secret 时直接配置：

```bash
# 交互式配置（手动输入 baseURL、appID、tenantCode、account、password）
ur setup

# 或手动编辑 ~/.ur/config.json
```

需要先在平台 UI 创建访问令牌：用户设置 → 访问令牌 → 创建。

### 方式 3：环境变量

```bash
# Device Auth / JWT 模式
export UR_BASE_URL=http://localhost:7777
export UR_APP_ID=77
export UR_ACCESS_KEY=xxx
export UR_ACCESS_SECRET=xxx
export UR_USER_ID=12345

# 或临时使用已有 token
export UR_TOKEN=xxx

# 指定应用上下文（等效于 --app）
export UR_APP=iot
```

## 认证原理

### Device Auth 模式（ur login）

```
ur login → 生成 setup code → 浏览器授权 → 创建访问令牌 → 点击「完成 CLI 绑定」
    → CLI 轮询 /setup-check → 获取 AK/SK → 保存到 ~/.ur/config.json
```

- 无需输入账号密码
- 授权通过后自动保存 AccessKey + AccessSecret + TenantCode
- 后续调用 API 时使用 AK/SK 生成 JWT

### AccessKey/JWT 模式

```
AccessKey + AccessSecret → HS256 JWT → Authorization: Bearer header
```

持有 AccessKey + AccessSecret 的客户端可**自行构造 JWT**，通过 `Authorization: Bearer <jwt>` Header 直接调用任意业务 API，无需先登录获取会话 Token。

#### 认证方式对比

| 特性 | Device Auth（ur login） | AccessKey/JWT（ur setup） |
|------|------------------------|--------------------------|
| 获取方式 | 浏览器授权（setup code） | 手动输入 AK/SK 或账号密码 |
| 有效期 | 由访问令牌设置决定 | 由客户端控制（建议 1~24 小时） |
| 适用场景 | 人机交互、首次配置 | 服务间调用、自动化脚本 |
| 撤销方式 | 删除访问令牌 | 删除访问令牌 |

#### JWT Payload 字段

| 字段名 | 类型 | JSON 名称 | 说明 |
|--------|------|-----------|------|
| `UserID` | int64 | `"userID"` | **必须为字符串格式** |
| `TenantCode` | string | `"tenantCode"` | 租户代码，可为空 |
| `AccessKey` | string | `"accessKey"` | 令牌标识 |
| `exp` | int64 | `"exp"` | Unix 秒时间戳（可选） |

#### 关键 HTTP Header

| Header | 值 | 说明 |
|---|--------|------|
| `Authorization` | Bearer \<jwt\> | JWT 模式（device auth / AK） |
| `app-id` | 应用 ID（如 `77`） | 必须 |
| `tenant-code` | 租户代码（如 `default`） | 必须 |

**关键陷阱**: JWT 中 `userID` 必须为字符串格式（`"userID": "12345"` 而非数字 `12345`），否则 401 错误。

## CLI 用法

### 全局选项

```bash
# 查看版本
ur --version
ur -v

# 切换应用上下文（支持 iot, platform-manage, org-manage, org-energy, console）
ur --app iot api /api/v1/things/device/info/get-list
UR_APP=iot ur api /api/v1/things/device/info/get-list
```

### 认证与配置

```bash
# Device Auth 授权（推荐首次使用）
ur login

# AI 模式：分步授权（获取 URL 和 setupCode，不阻塞）
ur login --no-wait --json
ur login --setup-code ABC123 --json

# 交互式配置 baseURL/appID/tenantCode（高级/手动配置）
ur setup

# 验证配置 + 连通性
ur check

# 管理多环境配置
ur config --list
ur config --use prod

# 查看当前 token
ur token --decode
ur token --raw
```

### API 调用

```bash
# 基本调用
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'

# 输出格式控制
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}' --format yaml
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}' --format raw

# 字段筛选（只保留指定字段）
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}' --fields code,data.total,data.list

# 摘要模式（列表只保留前 5 条）
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}' --summarize

# GJSON 路径提取
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}' --transform data.list.0.deviceName

# 保存输出到文件
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}' --output result.json

# 调试模式（打印完整 HTTP 请求/响应，敏感头已脱敏）
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}' --debug

# 自定义请求头
ur api /api/v1/things/device/info/get-list -H "X-Custom-Header: value" --body '{}'

# 从文件读取请求 body（适合大型 JSON）
ur api /api/v1/things/protocol/script/update --body-file /tmp/payload.json

# 临时覆盖连接配置做一次性调用（通过环境变量）
UR_BASE_URL=http://120.25.49.238:7777 UR_APP_ID=77 UR_TENANT_CODE=platform ur check
UR_BASE_URL=http://1.13.180.134:7777 UR_APP_ID=200 UR_TENANT_CODE=platform \
  ur api /api/v1/system/user/self/get-one
```

### 物模型命令

```bash
# 生成物模型模板
ur model template property --json
ur model template event --yaml --output event.yaml
ur model template action --json
ur model template full --yaml --output model.yaml

# 校验物模型 JSON
ur model validate /tmp/model.json

# 从物模型生成协议脚本框架
ur model generate-script /tmp/model.json --mode property --output script.go
```

### 场景联动命令

```bash
# 生成场景联动模板
ur scene template auto     # 自动触发场景
ur scene template manual   # 手动触发场景

# 校验场景联动 JSON
ur scene validate /tmp/scene.json
```

### 协议脚本命令

```bash
# 生成协议脚本模板
ur script template up-before    # 上行前处理
ur script template up-after     # 上行后处理
ur script template down-before  # 下行前处理
ur script template down-after   # 下行后处理

# 校验协议脚本
ur script validate /tmp/script.go
```

### Schema 与补全

```bash
# 查看 API schema
ur schema
ur schema --json
ur schema --auth-type admin
ur schema /api/v1/things/device/info/create

# Shell 补全（安装后新开终端生效）
ur completion bash >> ~/.bashrc
ur completion zsh >> ~/.zshrc
ur completion fish > ~/.config/fish/completions/ur.fish
```

### 输出选项说明

| 选项 | 说明 | 示例 |
|------|------|------|
| `--format json` | 美化 JSON（默认） | `ur api ... --format json` |
| `--format raw` | 单行 JSON | `ur api ... --format raw` |
| `--format yaml` | YAML 格式 | `ur api ... --format yaml` |
| `--transform PATH` | GJSON 路径提取 | `--transform data.list.0.name` |
| `--fields SELECTORS` | 字段筛选（逗号分隔） | `--fields code,data.total` |
| `--summarize` | 摘要模式（列表截断前 5 条） | `ur api ... --summarize` |
| `--output FILE` | 保存到文件 | `--output result.json` |
| `--debug` | 打印 HTTP 详情（敏感头脱敏） | `ur api ... --debug` |
| `--header, -H` | 自定义请求头 | `-H "X-Request-ID: abc"` |

> **注意**：`--fields`、`--summarize`、`--transform` 三者互斥，只能同时用其中一个。

## API 通用约定

- 全部使用 POST 方法（非 RESTful）
- 请求头: Content-Type: application/json, token 或 Authorization, app-id, tenant-code
- 响应格式: `{code, msg, data}`，其中 code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
- 枚举/状态字段必须加 `json:",status,string"` tag

Swagger 中 x-auth-type 标注了权限层级：
- `all`（所有登录用户）、 `admin`（租户管理员+平台管理员）、 `platform`（仅平台管理员）

## 常见错误排查

> 📘 完整排查指南见 `references/troubleshooting.md`

### 登录失败 - 租户校验

登录时请求头 `tenant-code` 需与用户所属租户匹配：

| 场景 | 行为 |
|------|------|
| 用户属于租户，租户已绑定 App | ✅ 登录成功 |
| 用户属于租户，租户**未**绑定 App | ❌ "该企业未绑定该应用，无法登录" |
| 用户不属于租户，App 是 `client` 类型 | 自动注册到 App 所属租户 |
| 不带租户登录，用户有绑定该 App 的租户 | ✅ 登录成功 |
| 不带租户登录，用户无绑定该 App 的租户（admin 应用） | ❌ "您没有权限访问该应用" |

**解决方案**：先不带 `tenant-code` 登录，查看 `userInfo.tenants` 确认用户所属租户。

### 401 认证失败

| 原因 | 排查方法 | 解决方案 |
|------|----------|----------|
| JWT 中 `userID` 格式错误 | 检查是否为字符串格式 | 改为 `"userID": "12345"` |
| `accessSecret` 填错 | 检查大小写和完整 32 位 | 复制创建令牌时返回的完整 secret |
| JWT 已过期 | 检查 `exp` 字段 | 重新生成 JWT |
| 访问令牌过期 | Device Auth / JWT 模式 | 重新运行 `ur login` 或重新生成 JWT |

**Token 自动刷新**：CLI 在调用 `ur api` 时，若遇到认证失败（code=401 或响应消息包含认证相关关键词），会**自动使用保存的账号密码重新登录获取新 token**，保存到配置文件后**自动重试原请求一次**。此行为无需手动干预，但若自动刷新失败（如账号密码已变更），会返回认证错误。

### 403 权限不足

| 错误信息 | 原因 | 解决方案 |
|----------|------|----------|
| 权限不足 | 需要管理员权限 | 使用管理员账号或令牌 |
| 租户权限不足 | 需要 platform 权限 | 使用平台管理员账号 |
| 数据权限不足 | 无该项目/区域权限 | 申请数据权限或联系管理员 |
| 该企业未绑定该应用 | 租户未开通该 App | 联系平台管理员绑定应用 |

### 业务错误 (code != 200)

CLI 调用 `ur api` 时，若业务返回 code 不为 200，会在 stderr 输出友好的错误提示，同时仍将完整响应输出到 stdout（或 `--output` 文件），方便查看具体错误信息：

```
[错误] 业务返回 code=500: 数据不存在
```

| 错误码 | 常见原因 | 排查方法 |
|--------|----------|----------|
| 参数错误 | 必填字段缺失、格式错误 | 检查请求 body 格式 |
| 数据不存在 | ID 错误或已删除 | 先查询确认 ID 正确 |
| 数据已存在 | 唯一键冲突 | 检查名称/编码是否重复 |
| 状态错误 | 操作不允许当前状态 | 检查业务状态流转规则 |

### 连接问题

| 错误 | 原因 | 解决方案 |
|------|------|----------|
| ECONNREFUSED | 服务未启动或端口错误 | 检查 `baseURL` 配置 |
| ETIMEDOUT | 网络超时 | 检查网络连通性或增加超时 |
| SSL 错误 | HTTPS 证书问题 | 检查证书或使用 HTTP |

## 各域 API 概览

详细 API 端点和业务场景见各子域 SKILL.md：

- **ur-device/SKILL.md** — 设备管理（85 端点）：设备 CRUD、属性控制、网关子设备、设备认证、物模型
- **ur-product/SKILL.md** — 产品管理（86 端点）：产品 CRUD、品类、物模型、协议脚本、OTA 固件
- **ur-project/SKILL.md** — 项目管理（41 端点）：项目、区域、分组、数据权限、用户权限申请
- **ur-user/SKILL.md** — 用户管理（153 端点）：登录注册、角色、部门、字典、定时任务、通知
- **ur-system/SKILL.md** — 系统管理（53 端点）：通用接口、应用、菜单、Hook、授权码
- **ur-tenant/SKILL.md** — 租户管理（38 端点）：租户 CRUD、用户管理、权限配置、应用绑定
- **ur-ai/SKILL.md** — AI 管理（36 端点）：Agent、Skill、会话、告警、场景联动

### 常用工作流示例

查询设备列表：
```bash
ur api /api/v1/things/device/info/get-list \
  --body '{"page":{"page":1,"size":10}}'
```

控制设备属性：
```bash
ur api /api/v1/things/device/interact/property-control-send \
  --body '{"productID":"xxx","deviceName":"yyy","data":{"power":1}}'
```

创建访问令牌：
```bash
ur api /api/v1/system/user/self/access-token/create \
  --body '{"name":"my-token","expTime":86400}'
```

## 高频端点速查（Top 15）

最常用的 15 个端点，直接复制使用：

| 操作 | 端点 | 最小请求体 |
|------|------|-----------|
| 设备列表 | POST /api/v1/things/device/info/get-list | `{"page":{"page":1,"size":10}}` |
| 当前用户信息 | POST /api/v1/system/user/self/get-one | `{}` |
| 设备列表 | POST /api/v1/things/device/info/get-list | `{"page":{"page":1,"size":10}}` |
| 控制设备属性 | POST /api/v1/things/device/interact/property-control-send | `{"productID":"x","deviceName":"x","data":{"Key":值}}` |
| 设备最新属性 | POST /api/v1/things/device/msg/property-latest/get-list | `{"productID":"x","deviceName":"x"}` |
| 产品列表 | POST /api/v1/things/product/info/get-list | `{"page":{"page":1,"size":10}}` |
| 产品物模型 | POST /api/v1/things/product/schema/get-list | `{"productID":"x"}` |
| 项目列表 | POST /api/v1/things/project/crud/get-list | `{"page":{"page":1,"size":10}}` |
| 创建访问令牌 | POST /api/v1/system/user/self/access-token/create | `{"name":"my-token","expTime":86400}` |
| 租户列表（平台管理员）| POST /api/v1/system/tenant/info/get-list | `{"page":{"page":1,"size":10}}` |
| 场景列表 | POST /api/v1/things/scene/info/get-list | `{"page":{"page":1,"size":10}}` |
| 手动触发场景 | POST /api/v1/things/scene/info/manually-trigger | `{"id":"<sceneID>"}` |
| 告警规则列表 | POST /api/v1/things/alarm/info/get-list | `{"page":{"page":1,"size":10}}` |
| 文件上传 | POST /api/v1/system/common/upload-file | multipart/form-data |
| 批量接口调用 | POST /api/v1/system/common/api/batch-agg | `{"apis":[{"path":"/api/v1/system/user/self/get-one","body":{}}]}` |

---

## 跨域工作流

### 工作流 1：全新 IoT 设备从零上线

**角色**：租户管理员 | **涉及域**：ur-product → ur-device → ur-ai（可选）

```bash
# 1. 创建产品（ur-product）
ur api /api/v1/things/product/info/create \
  --body '{"name":"智能开关","protocolCode":"urMqtt","deviceType":1}'
# → 获得 productID

# 2. 定义物模型属性（ur-product）
ur api /api/v1/things/schema/create \
  --body '{"productID":"<step1>","identifier":"Power","name":"开关","accessMode":"rw","dataType":{"type":"bool"}}'

# 3. 创建设备（ur-device）
ur api /api/v1/things/device/info/create \
  --body '{"productID":"<step1>","deviceName":"switch-001","deviceAlias":"一楼走廊开关"}'
# → 获得三元组（含 deviceSecret）

# 4. 验证属性（ur-device）
ur api /api/v1/things/device/msg/property-latest/get-list \
  --body '{"productID":"<step1>","deviceName":"switch-001"}'

# 5. 可选：创建场景联动（ur-ai）
ur api /api/v1/things/scene/info/create \
  --body '{"name":"自动关灯","triggers":[{"type":"timer","cron":"0 22 * * *"}]}'
```

**关键依赖**：步骤 3 依赖步骤 1 的 productID；设备控制依赖设备在线（isOnline=1）。

---

### 工作流 2：新租户完整初始化

**角色**：平台管理员 | **涉及域**：ur-tenant → ur-system → ur-tenant

```bash
# 1. 创建租户（ur-tenant，需平台管理员）
ur api /api/v1/system/tenant/info/create \
  --body '{"name":"Acme Corp","tenantCode":"acme"}'

# 2. 为租户绑定应用（ur-system，需平台管理员）
ur api /api/v1/system/tenant/app/create \
  --body '{"tenantCode":"acme","appID":77}'

# 3. 邀请租户管理员（ur-tenant）
ur api /api/v1/system/tenant/user/invite \
  --body '{"tenantCode":"acme","account":"admin@acme.com"}'
```

---

### 工作流 3：设备超阈值告警 + 通知

**角色**：租户管理员 | **涉及域**：ur-ai → ur-user

```bash
# 1. 创建告警规则（ur-ai）
ur api /api/v1/things/alarm/info/create \
  --body '{"name":"温度过高","productID":"<productID>","alarmLevel":1}'

# 2. 创建通知配置（ur-user）
ur api /api/v1/system/notify/config/create \
  --body '{"name":"设备告警通知","throttle":300}'

# 3. 创建通知模板（ur-user）
ur api /api/v1/system/notify/template/create \
  --body '{"configID":"<step2>","notifyType":"inner","template":"设备${deviceName}告警"}'
```

---

## 新增接口维护规范

> 📘 详细规范见 `docs/中台/功能说明/API角色权限/设计方案.md`

### 新增接口检查清单

1. **标注权限**：在 .api 文件中添加 `// x-auth-type: admin|all|platform`
2. **更新 Skill**：
   - 新增域 → 在 `backend/cli/ur/internal/config/app.go` 的 `Features()` 中添加功能模块定义
   - 新增接口 → 确认后端 `.api` 文件中已标注 `// x-auth-type: admin|all|platform`
3. **重新生成**：
   ```bash
   cd backend/cli/ur
   go run ./cmd/ur-iot generate-skills
   # 或批量生成所有应用
   for app in ur-platform-manage ur-iot ur-org-manage ur-org-energy ur-console; do
     go run ./cmd/$app generate-skills
   done
   ```

### 角色权限标注指南

| 标注值 | 适用场景 |
|--------|---------|
| `platform` | 跨租户操作、租户 CRUD、全局配置 |
| `admin` | 本租户内 CRUD、用户管理、角色分配 |
| `all` | 个人信息、设备分享、查看权限范围内数据 |
| `mixed` | 不同角色有不同权限（需在 `roleViews` 中详细说明） |

## 前端应用权限映射

不同类型的前端应用可调用的 API authType 不同：

| 应用 | 类型 | TENANT_CODE | 可调用 API authType |
|------|------|-------------|---------------------|
| platform-manage | 平台 | platform | platform, admin, all |
| platform-iot | 平台 | platform | platform, admin, all |
| org-manage | 组织 | 用户租户 | admin, all |
| org-iot | 组织 | 用户租户 | admin, all |
| org-energy | 组织 | 用户租户 | admin, all |
| console | 公共 | 用户租户 | all |

### 平台应用（platform-manage, platform-iot）

- **环境变量**: `VITE_GLOB_TENANT_CODE=platform`
- **请求头**: 固定发送 `tenant-code: platform`
- **可访问 API**: 所有 authType（platform/admin/all）
- **典型场景**: 租户管理、应用配置、平台级数据查看

### 组织应用（org-manage, org-iot, org-energy）

- **环境变量**: 无 `VITE_GLOB_TENANT_CODE`
- **请求头**: 租户代码由登录用户决定
- **可访问 API**: admin、all
- **典型场景**: 设备管理、用户管理、租户级配置

### 公共应用（console）

- **环境变量**: `VITE_GLOB_APPID=77`
- **请求头**: 租户代码由登录用户决定
- **可访问 API**: all
- **典型场景**: 应用入口、租户切换、个人信息管理

---

## 参考文档

| 文档 | 说明 |
|------|------|
| `references/troubleshooting.md` | 常见问题排查指南（登录失败、权限错误、设备控制异常） |
| `references/error-codes.md` | 错误码速查表 |
| `references/quick-reference.md` | 快速参考（路径前缀映射、枚举值、CLI schema） |
| `references/api-conventions.md` | API 通用约定 |
| `references/iot-tools-versions.md` | IoT Tools 版本说明 |

详细 API 端点见各子域 SKILL.md：ur-device、ur-product、ur-project、ur-user、ur-system、ur-tenant、ur-ai
