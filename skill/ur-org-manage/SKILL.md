---
name: ur-api
description: "ur-api — 联犀 SaaS 平台统一 API 工具（涵盖所有应用）"
metadata:
  hermes:
    tags: [api, cli, saas, iot, platform, org, energy, console]
---

# ur-api — 联犀 SaaS 平台

> **配置检查**：如果尚未配置联犀连接，请先运行 `ur-api login --no-wait`，按指引在浏览器中完成授权。`setup` 命令是终端交互式的，在 AI 聊天环境中无法使用。

## 应用选择

根据当前操作的前端应用选择对应的 CLI：

| CLI | 前端应用 | AppID | TenantCode | 可调用权限 |
|-----|---------|-------|------------|-----------|
| `ur-platform-manage` | 平台管理 | 100 | platform | platform, admin, all, public |
| `ur-iot` | 物联网 | 200 | 用户输入 | platform, admin, all, public |
| `ur-org-manage` | 组织管理 | 300 | 用户输入 | admin, all, public |
| `ur-org-energy` | 能源管理 | 1000 | 用户输入 | admin, all, public |
| `ur-console` | 控制台 | 77 | 用户输入 | all, public |

## 常见查询速查

以下是最常用的查询场景及对应 API。**优先尝试这些接口，无需查阅子 skill**。

| 查询场景 | 推荐 CLI | API 命令示例 | 权限 |
|---------|---------|-------------|------|
| 查询【我的】应用列表 | `ur-console` | `api /api/v1/system/user/self/app/get-list` | all（任何登录用户） |
| 查询【我的】菜单列表 | `ur-console` | `api /api/v1/system/user/self/menu/get-list` | all（任何登录用户） |
| 查询当前用户信息 | `ur-console` | `api /api/v1/system/user/self/get-one` | all（任何登录用户） |
| 查询设备列表 | `ur-iot` | `api /api/v1/things/device/info/get-list` | admin/tenant |
| 查询产品列表 | `ur-iot` | `api /api/v1/things/product/info/get-list` | admin/tenant |
| 查询项目列表 | `ur-iot` | `api /api/v1/things/project/info/get-list` | admin/tenant |
| 发起流程/办审批 | `ur-org-manage` | 专用命令 `flow process launch` / `flow task consent`（详见其 references/flow-approval.md） | all（任何登录用户） |
| 流程定义/表单管理 | `ur-org-manage` | 专用命令 `flow def create` / `flow def deploy`（仅租户管理员可用） | admin（租户管理员） |

> **空结果说明**：如果 `user/self/app/get-list` 返回空列表（`list: []`），表示**当前用户没有任何应用权限**。
> 此时应直接告知用户『您当前没有分配任何应用』，**不要**再去 `system/app/info/get-list`（platform 权限）查找。
> `system/app/info/get-list` 是平台管理员专属接口，普通用户调用会返回权限不足。

## API 端点

共 745 个可调用端点（涵盖所有应用）。

> **注意**：本 skill 为统一索引。上方「常见查询速查」已列出最常用 API，**优先尝试速查表中的接口**。
> 如需查看更多端点，按以下三级渐进式查阅（文档越小越优先）：
> 1. 查阅领域索引（只含 group 名称和数量，约 2KB）：
>    - `skill_view(name="ur-api", filePath="references/system-index.md")` — system 领域
>    - `skill_view(name="ur-api", filePath="references/things-index.md")` — things 领域
>    - `skill_view(name="ur-api", filePath="references/ai-index.md")` — AI 领域
> 2. 从索引中找到目标 group，查阅 group 详细文件（约 0.5-3KB）：
>    - `skill_view(name="ur-api", filePath="references/groups/{group文件名}.md")`
> 3. 如需查看全部端点，调用 `skill_view(name="ur-api", filePath="swagger-index.md")`。

### 分类索引

| 分类 | 端点数量 |
|------|---------|
| things/product/config | 1 |
| system/notify/message/info | 4 |
| system/tenant/app | 5 |
| system/tenant/core | 2 |
| things/alarm/capability | 1 |
| things/product/category | 11 |
| system/dept/user | 3 |
| things/alarm/info | 7 |
| things/product/remoteConfig | 4 |
| system/ops/workOrder | 5 |
| system/user/dept | 2 |
| things/schema/common | 8 |
| ai/clone | 5 |
| system/dept/role | 3 |
| system/route | 3 |
| system/tenant/user | 11 |
| things/alarm/conditionTemplate | 5 |
| things/protocol/image | 1 |
| system/app/group | 4 |
| system/resource/api | 4 |
| things/alarm/notifyTemplate | 6 |
| things/ota/firmware/job | 4 |
| things/product/schema | 7 |
| things/scene/info | 6 |
| ai/mcp/resolve | 1 |
| system/app/agreement | 2 |
| system/checkIn | 5 |
| system/flow | 39 |
| system/user/tenant | 1 |
| things/device/auth | 4 |
| ai/memory/debug | 5 |
| system/tenant/info | 6 |
| things/alarm/ai | 1 |
| things/area/profile | 3 |
| things/protocol/sync | 2 |
| system/mall/package | 5 |
| system/notify/config/template | 3 |
| system/role/app | 2 |
| system/role/resource | 2 |
| system/tenant/renewal | 3 |
| things/ai/mcp | 4 |
| things/group/info | 5 |
| things/scene/log | 1 |
| ai/clone-task | 10 |
| AI中台 | 15 |
| system/dept/syncJob | 6 |
| things/device/bind-token | 2 |
| things/group/device | 3 |
| things/project/customData | 2 |
| ai/agent-group | 5 |
| things/protocol/container | 5 |
| ai/knowledge-document | 11 |
| ai/tool | 10 |
| things/alarm/notifyRecord | 2 |
| things/data/area | 3 |
| things/device/ota | 1 |
| system/dict/info | 7 |
| things/ota/firmware/device | 4 |
| things/project/crud | 3 |
| system/log | 2 |
| system/notify/config | 7 |
| system/tenant/app/menu | 2 |
| system/user/info | 5 |
| things/protocol/config | 5 |
| things/protocol/info | 5 |
| things/protocol/service | 2 |
| things/user/device/collect | 3 |
| ai/agent | 1 |
| ai/clone/memory | 11 |
| ai/tool/run | 3 |
| system/app/core | 1 |
| system/mcp | 1 |
| system/tenant/user/role | 2 |
| things/area/info | 5 |
| things/data/area/user/apply | 2 |
| system/notify/notification | 9 |
| things/device/interact/gateway | 2 |
| things/device/profile | 4 |
| things/project/info | 5 |
| things/user/area/apply | 1 |
| ai/speaker/embedding | 7 |
| system/common | 18 |
| system/ops/feedback | 3 |
| system/user/self | 32 |
| things/protocol/script | 8 |
| ai/mcp/tools | 1 |
| things/device/schema | 6 |
| system/mall/product | 5 |
| system/role/menu | 2 |
| things/project/profile | 3 |
| system/agreement | 6 |
| system/dept/info | 5 |
| system/mall/license | 5 |
| system/app/info | 6 |
| system/user/self/tenant | 5 |
| things/device/auth5 | 2 |
| things/device/gateway | 3 |
| things/device/pair | 2 |
| ai/knowledge-tool | 3 |
| ai/session | 8 |
| ai/udp | 1 |
| system/app/menu | 4 |
| system/notify/template | 5 |
| system/role/info | 4 |
| system/tenant/config | 2 |
| system/user/self/accessToken | 5 |
|  | 5 |
| ai/knowledge-export | 3 |
| system/user/data | 2 |
| things/config/info | 2 |
| things/data/user | 1 |
| things/product/custom | 2 |
| things/product/info | 8 |
| things/user/device/share | 17 |
| system/init | 1 |
| things/data/project | 5 |
| things/device/interact | 11 |
| things/rule/scene/info | 6 |
| ai/chat | 8 |
| ai/config | 7 |
| ai/control | 3 |
| ai/knowledge | 6 |
| things/device/edge | 2 |
| things/ota/module/info | 5 |
| system/tenant/agreement | 5 |
| things/alarm/event | 5 |
| things/hook | 1 |
| things/ota/firmware/info | 5 |
| ai/clone-group | 5 |
| ai/skill | 8 |
| system/user/self/thirdparty | 3 |
| things/device/group | 3 |
| things/device/info | 15 |
| things/protocol/script/device | 5 |
| ai/skill/file | 4 |
| system/app/role | 6 |
| system/dict/detail | 5 |
| system/notify/news | 5 |
| things/rule/scene/log | 1 |
| system/hook/server | 5 |
| system/job/task | 14 |
| ai/mcp/service | 6 |
| system/config/core | 1 |
| things/device/msg | 17 |
| things/device/version | 2 |
| ai/knowledge-compile | 2 |
| system/hook/capability | 4 |

## 使用示例

```bash
# 配置（以 iot 为例，其他应用同理）
ur-iot setup

# 验证连通性
ur-iot check

# 调用 API
ur-iot api /api/v1/system/user/self/get-one
```
