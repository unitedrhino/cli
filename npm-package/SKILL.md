---
name: ur-api
description: "联犀 SaaS 平台 API 命令行工具 — 按前端应用拆分为 5 个独立 CLI，每个绑定固定的 app-id 和 tenant-code"
metadata:
  hermes:
    tags: [api, cli, saas, iot]
---

# ur-api — 联犀 SaaS 平台 API 工具（应用导向）

## ⚠️ 配置指引（首次使用必读）

**禁止在当前对话中向用户索取 AccessKey、AccessSecret、密码等敏感信息。**

### AI 环境中唯一可用的配置方式

`ur-xxx setup` 是终端交互式命令（需要逐行输入账号密码），**在当前对话环境中无法使用**。

**唯一可行的配置方式**是设备授权流（Device Flow），AI 分两步完成：

**第 1 步 — 获取授权 URL**：
```bash
ur-iot login --no-wait --json
```
输出 JSON 示例：
```json
{
  "verification_url": "https://console.unitedrhino.com/user/access-tokens?setup=ABC123",
  "setup_code": "ABC123",
  "expires_in": 600,
  "hint": "在浏览器中打开 verification_url 完成授权，然后执行: login --setup-code ABC123"
}
```
AI 解析 JSON，向用户展示 `verification_url` 和操作步骤。

**第 2 步 — 完成授权**（用户确认在浏览器中点击「完成 CLI 绑定」后）：
```bash
ur-iot login --setup-code ABC123 --json
```
输出 JSON 示例：
```json
{
  "event": "authorization_complete",
  "tenant_code": "t1",
  "access_key": "ak_xxxx",
  "access_secret": "sk_xxxx"
}
```

**全程无需在当前对话中输入任何密钥。**

### 人类终端模式（非 AI 环境）

如果用户在本地终端直接操作：
```bash
# 阻塞模式：生成 URL → 等待用户浏览器授权 → 自动保存配置
ur-iot login

# 或指定地址跳过交互选择
ur login --base-url https://saas.unitedrhino.com
```

### 认证方式说明

CLI 支持两种认证机制：
1. **Session Token**（`ur-iot login` 自动获取）：`token:` header
2. **AccessKey/JWT**（login 完成后自动保存）：AccessKey + AccessSecret → HS256 JWT → `Authorization: Bearer` header

所有接口均为 POST 方法。请求格式 `{code, msg, data}`。

---

## 应用选择

根据当前操作的前端应用选择对应的 CLI：

| CLI | 前端应用 | AppID | TenantCode | 可调用权限 |
|-----|---------|-------|------------|-----------|
| `ur-platform-manage` | 平台管理 | 100 | platform | platform, admin, all |
| `ur-iot` | 物联网 | 200 | platform | platform, admin, all |
| `ur-org-manage` | 组织管理 | 300 | 用户输入 | admin, all |
| `ur-org-energy` | 能源管理 | 1000 | 用户输入 | admin, all |
| `ur-console` | 控制台 | 77 | platform | all |

## 快速决策

| 用户意图 | 使用 CLI | 说明 |
|---------|---------|------|
| 管理租户、用户、应用、授权 | `ur-platform-manage` | 平台管理员操作，跨租户 |
| 管理设备、产品、项目、OTA、协议 | `ur-iot` | 物联网（**平台/租户共用**，tenant-code=platform） |
| 管理组织用户、角色、AI 智能体 | `ur-org-manage` | 组织管理（租户级，tenant-code 需输入） |
| 能耗分析、电力集抄、预付费 | `ur-org-energy` | 能源管理（租户级，tenant-code 需输入） |
| 个人信息、访问令牌、续期 | `ur-console` | 控制台（平台级，tenant-code=platform） |

### 物联网 vs 能源管理 — 能力重叠与分工

`ur-org-energy` **复用了物联网的核心设备能力**，但聚焦能源场景：

| 能力 | `ur-iot` | `ur-org-energy` | 说明 |
|------|---------|----------------|------|
| 设备管理 | ✅ 产品 + 设备 CRUD、物模型 | ✅ 设备控制台、分组、区域 | 能源管理不管理产品定义，只使用已创建的设备 |
| 场景联动 | ✅ 项目级场景编辑 | ✅ 自动化 + 一键场景 | 两者都支持，数据互通 |
| 区域管理 | ✅ 区域 CRUD | ✅ 区域 CRUD | 共用同一套区域数据 |
| OTA 升级 | ✅ 固件包、批量升级 | ❌ | 能源管理不涉及固件升级 |
| 数据流转 | ✅ 协议网关、协议脚本 | ❌ | 能源管理不涉及协议开发 |
| 能耗分析 | ❌ | ✅ 用能概况、同比环比、趋势、排名、损耗 | 能源特有 |
| 电力集抄 | ❌ | ✅ 实时监控、历史数据、极值报表 | 能源特有 |
| 预付费管理 | ❌ | ✅ 房间、租客、充值、消费 | 能源特有 |

### ur-iot 按角色可见菜单

`ur-iot` **不是平台管理员专属**，租户管理员也能使用大部分功能。

**平台管理员（platform）可见全部菜单**：
- 信息面板、设备地图
- 设备管理（产品管理、设备管理、**通用物模型**、**产品品类**）
- 项目管理（项目列表、项目详情、场景编辑）
- 区域管理
- OTA 升级（升级包列表、模块列表、批量升级）
- 数据流转（**协议网关**、协议脚本）

**租户管理员（admin）可见菜单**（不含 `authority: ['platform']` 标记的）：
- 信息面板、设备地图
- 设备管理（产品管理、设备管理）
- 项目管理（项目列表、项目详情、场景编辑）
- 区域管理
- OTA 升级（升级包列表、模块列表、批量升级）
- 数据流转（协议脚本）

> **平台专属菜单**（租户管理员不可见）：通用物模型、产品品类、协议网关
> 
> 前端路由单一事实来源：`apps/web/apps/iot/src/router/routes/modules/iot.ts`

**决策规则**：
1. 如果是**能源业务**（能耗分析、电力集抄、预付费）→ 用 `ur-org-energy`
2. 如果是**设备/产品/项目管理/OTA** → 平台管理员和租户管理员都用 `ur-iot`
3. 如果是**协议开发**（协议网关、通用物模型、产品品类）→ 必须用 `ur-iot`（且需平台管理员权限）

## 各应用详细文档

- [ur-platform-manage](ur-platform-manage/SKILL.md) — 平台管理
- [ur-iot](ur-iot/SKILL.md) — 物联网
- [ur-org-manage](ur-org-manage/SKILL.md) — 组织管理
- [ur-org-energy](ur-org-energy/SKILL.md) — 能源管理
- [ur-console](ur-console/SKILL.md) — 控制台

## 通用用法

```bash
# 验证连通性
ur-iot check

# 调用 API
ur-iot api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'

# 查看 API schema
ur-iot schema

# 查看 token
ur-iot token --decode
```
