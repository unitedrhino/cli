---
name: ur-api
description: "Use when calling 联犀 SaaS 平台 API, bringing physical IoT firmware onto the platform, or parsing documents: product and device management, provisioning, MQTT, uplink/downlink, debugging, OTA, authentication recovery, project management, system management, bigscreen visualization, and document parsing. triggers: API调用, CLI登录, CLI升级, _notice, Sandbox认证, 账号密码登录, AKSK, 设备列表, IoT设备控制, 设备固件, 固件接入, BLE配网, MQTT, OTA升级, 鉴权恢复, 设备注册, 访问令牌, 物模型管理, 场景联动, 权限配置, swagger schema, 项目管理, 区域管理, 系统管理, 大屏, 数据可视化, GoView, 文档解析, 解析PDF, 解析Excel, 读附件, ur doc"
---

# ur-api — 联犀 SaaS 平台 API 工具

CLI 主实现位于独立仓库 `unitedrhino/cli`，通过 Go CLI `ur` 提供能力。Sandbox 可只用 `UR_*` 环境变量；本地已配置的旧 profile 继续自动兼容，无需迁移或重新登录。

所有接口均为 POST 方法。请求格式 `{code, msg, data}`。

---

## 认证前置流程（AI 必读）

1. 先运行 `ur check --json`。如果环境变量或历史 profile 已可用，直接调用 API；不要删除配置、重新执行 `setup` 或要求用户再次提供凭据。
2. 只有在没有可用认证时，才运行 `ur login --method device|password|aksk`；未指定 `--method` 时默认 `device`。
3. Sandbox 设置 `UR_BASE_URL` 后采用 env-only 模式：只使用同一环境中的完整凭据组，不读取磁盘 profile 补齐缺失值，也不把环境凭据写入磁盘。

支持的运行时凭据组按 `UR_TOKEN` → `UR_ACCESS_KEY` + `UR_ACCESS_SECRET` → `UR_ACCOUNT` + `UR_PASSWORD` 的顺序选择。AK/SK 不要求 `UR_USER_ID`；成对变量缺一项时应修正 Sandbox 注入，不要从 profile 混合补值。

账号密码入口接收**原始密码**，CLI 会做且只做一次 SHA-256；禁止先摘要再传入，否则会二次摘要。直接调用 `/api/v1/system/user/self/login` 时，调用方才需要发送 SHA-256 摘要并使用 `pwdType: 1`；MD5/`pwdType: 2` 仅为历史兼容。

敏感值优先通过 Sandbox 环境变量或 stdin 注入。`--password`、`--access-secret` 明文参数仅用于兼容，可能进入 shell 历史或进程列表；不得把密码、AccessSecret 或完整 Token 写入 skill、日志、截图和版本库。

业务命令的 JSON 结果可能包含 `_notice.update` 或 `_notice.skills`。先完成用户当前请求，再按 `command` 简短提示或执行用户已授权的升级；统一使用 `ur upgrade` 更新 CLI 与客户端 Skills。详细规则见 `references/cli-usage.md` 的“CLI 更新提示与处理”。

## 文档解析(ur doc)

用户上传或引用文档(PDF/Word/PPT/Excel/图片/邮件等)需要读取内容时,使用 `ur doc parse`:先 `--format outline` 看结构地图,再按需 `--format md` 通读/`--section` 取章节,或 `--format json` 落盘后用 jq 精查(excel 公式溯源、单元格坐标)。扫描件/图片加 `--ocr`(默认平台模型池计费)。完整用法见 `ur-doc/SKILL.md`。

## 角色权限区分

**全栈权限总索引**（契约 / 运行时 / 前端 / 消歧）：仓库根目录 `docs/中台/功能说明/权限体系总览.md`。

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| **平台管理员** | 跨企业操作、企业 CRUD、查看所有企业 | 创建企业、管理企业配置 |
| **企业管理员** | 本企业内 CRUD、用户管理、角色分配 | 创建设备、管理用户、配置项目 |
| **普通用户** | 个人信息、设备分享、查看权限范围内数据 | 登录、修改密码、分享设备 |

### 按角色区分的域

| 域 | 平台管理员 | 企业管理员 | 普通用户 |
|---|-----------|-----------|---------|
| ur-tenant | 创建/删除企业、查看所有企业 | 管理本企业、邀请用户 | 无权限 |
| ur-device | 无权限 | 设备 CRUD、属性控制 | 设备分享、收藏 |
| ur-user | 无权限 | 用户 CRUD、角色管理 | 个人信息、修改密码 |
| ur-product | 无权限 | 产品 CRUD、物模型管理 | 无权限 |
| ur-ota | 无权限 | 固件包、模块与升级任务管理 | 无权限 |
| ur-project | 无权限 | 项目 CRUD、区域管理 | 查看权限内项目 |
| ur-ai | 无权限 | Agent 配置、告警管理 | 无权限 |
| ur-view | 无权限 | 大屏 CRUD、画布编辑、发布管理、素材管理 | 无权限 |

### 当前用户身份

调用任何联犀平台 API 前，**先运行 check 确认角色**，错误的角色会导致 403。纯外部 HTTP 任务不依赖平台角色，不为此运行 `ur check`：

```bash
ur check --json
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

规则引擎 AI 任务、定时 HTTP、设备状态查询后转发或跨设备控制，先读 [AI 工具开发](ai-tool/SKILL.md)，再按其中的具体平台操作加载对应子域。设备数据模拟只是通用任务的一个可选分支，不要把模拟澄清问题或接口合同套用到普通 HTTP、巡检和转发任务。

### 第一步：确认当前用户角色

```bash
ur check
```

| 角色 | 可操作范围 | 配置方式 |
|------|----------|---------|
| **平台管理员** | 所有域（platform/admin/all 接口） | 先复用环境/profile；必要时用 `ur login` 切换认证 |
| **企业管理员** | 本企业 CRUD（admin/all 接口） | 先复用环境/profile；必要时用 `ur login` 切换认证 |
| **普通用户** | 个人信息、设备分享（仅 all 接口） | 先复用环境/profile；必要时用 `ur login` 切换认证 |

> **切换角色**：重新运行 `ur login`（用不同账号授权），或用 `--app` 切换应用上下文：
> ```bash
> ur --app platform-manage api /api/v1/...
> ```

### 第二步：根据任务选择子域

| 用户意图 / 关键词 | 加载子域 | 最低所需角色 |
|----------------|---------|------------|
| 设备列表、设备控制、属性上报、MQTT、三元组、网关、OTA升级设备 | `ur-device` | admin（控制/CRUD），user（分享/收藏） |
| 设备固件从零接入、编译烧录、配网、MQTT 契约、OTA 回滚、鉴权自愈 | `device-firmware` | 固件操作无平台角色；平台初始化通常需 admin |
| OTA 升级包、模块、任务创建和进度核验 | `ur-ota` | admin |
| 产品定义、物模型、协议脚本、OTA固件包、品类 | `ur-product` | admin |
| 项目、区域、分组、数据权限申请 | `ur-project` | admin（管理），user（申请权限） |
| 登录、用户信息、角色、部门、字典、通知、访问令牌、修改密码 | `ur-user` | admin（CRUD），user（自身信息） |
| 文件上传、WebSocket、批量接口、应用管理、Hook | `ur-system` | user（上传/WS），admin（应用），platform（全局应用） |
| 创建企业、企业列表、企业配置、邀请用户加入企业 | `ur-tenant` | **platform**（创建/查看全部），admin（管理本企业） |
| AI Agent、场景联动、告警规则、数字分身、AI会话 | `ur-ai` | admin |
| 大屏、数据可视化、GoView、画布编辑、发布大屏、看板、素材库 | `ur-view` | admin |

### 第三步：IoT AI 工具调用设计

设备模拟或控制先读 [属性控制与模拟数据](ur-device/references/device-control.md)：区分云端改值、模拟上报、实体控制与只生成样例；泛称“模拟数据”时先澄清目的，不默认下发。

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
   - 第一步：查阅本 SKILL 的「高频端点速查」或各子域 SKILL.md 中的端点列表。
   - 第二步：如果仍不确定，调用 `knowledge_search` 检索知识库中关于该业务的真实接口文档。
4. **优先使用已知真实路径**：只有在本 SKILL 中明确列出的路径、或各子域 SKILL.md 中记录的路径，才是可调用的真实路径。

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
| ur-tenant | 38 | 企业管理：企业CRUD、用户管理、权限配置 |
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

大屏可视化子域：
- `ur-view/SKILL.md` — 大屏（GoView）项目 CRUD、画布 JSON 本地编辑闭环（pull/validate/push/publish/screenshot）、素材库管理、IoT 数据绑定配方

设备端固件子域：
- `device-firmware/SKILL.md` — 设备从产品/物模型初始化到固件接入、首刷、调试、全量 OTA、鉴权恢复和实机验收

OTA 管理子域：
- `ur-ota/SKILL.md` — 升级包上传登记、模块和升级任务管理；设备端协议与回滚由 `device-firmware` 负责

---

## 快速开始

```bash
# 总是先检查现有 Sandbox 环境或历史 profile
ur check --json
```

检查通过后直接调用 API。仅在返回缺少认证时选择以下一种登录方式。

### 方式 1：Device Flow（默认）

```bash
# 1. 运行 login，CLI 生成 setup code 并输出授权 URL
ur login --method device

# 2. 在浏览器中打开 URL，登录控制台后创建/选择访问令牌
# 3. 点击「完成 CLI 绑定」
# 4. CLI 自动轮询获取 AK/SK，保存到 ~/.ur/config.json

# 验证连通性
ur check --json

# 调用 API
ur api /api/v1/system/user/self/get-one
```

### 方式 2：账号密码

```bash
# 推荐由环境或 stdin 提供秘密；password 必须是原始密码
UR_PASSWORD='<原始密码>' ur login --method password \
  --account '<账号>' --tenant-code '<企业编码>' --json
# 或：printf '%s' "$UR_PASSWORD" | ur login --method password \
#       --account '<账号>' --tenant-code '<企业编码>' --password-stdin --json
```

### 方式 3：AccessKey/JWT

```bash
UR_ACCESS_SECRET='<AccessSecret>' ur login --method aksk \
  --access-key '<AccessKey>' --tenant-code '<企业编码>' --json
```

AK/SK 登录不要求 `userID`。需要先在平台 UI 创建访问令牌：用户设置 → 访问令牌 → 创建。

### 方式 4：Sandbox 环境变量（无需 login）

```bash
export UR_BASE_URL='<平台地址>'
export UR_APP_ID='<应用ID>'
export UR_TENANT_CODE='<企业编码>'

# 以下三组任选一组，不要混用半套凭据：
export UR_TOKEN='<Session Token>'
# export UR_ACCESS_KEY='<AccessKey>' UR_ACCESS_SECRET='<AccessSecret>'
# export UR_ACCOUNT='<账号>' UR_PASSWORD='<原始密码>'

ur check --json
```

旧 `~/.ur/config.json` 中仅有账号/密码，或同时遗留 Token、AK/SK 的 profile 仍会自动兼容；不要为了升级主动重写或删除它。`ur setup` 仅保留为人类终端的兼容向导，不是 AI 首选入口。

---

## 参考文档

| 文档 | 说明 |
|------|------|
| `references/auth.md` | 认证方式详解（Device Auth / JWT / 环境变量） |
| `references/cli-usage.md` | CLI 命令用法（更新通知、多客户端 Skills 分发、API 调用、物模型、场景联动、协议脚本、schema） |
| `references/api-conventions.md` | API 通用约定（请求格式、响应格式、分页、权限标注） |
| `references/troubleshooting.md` | 常见问题排查（登录失败、401/403、连接问题、问题反馈） |
| `references/quick-reference.md` | 高频端点速查（Top 15 最常用端点） |
| `references/workflows.md` | 跨域工作流示例（设备上线、企业初始化、告警通知） |
| `references/maintenance.md` | 新增接口维护规范（检查清单、权限标注指南） |
| `references/frontend-permissions.md` | 前端应用权限映射（平台应用 / 组织应用 / 公共应用） |

详细 API 端点见各子域 SKILL.md：ur-device、ur-product、ur-project、ur-user、ur-system、ur-tenant、ur-ai
