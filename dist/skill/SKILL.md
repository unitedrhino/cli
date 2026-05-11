---
name: ur-api
description: "联犀 SaaS 平台 API 命令行工具 — 按前端应用拆分为 5 个独立 CLI，每个绑定固定的 app-id 和 tenant-code"
---

# ur-api — 联犀 SaaS 平台 API 工具（应用导向）

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
| 管理租户、用户、应用、授权 | `ur-platform-manage` | 平台管理员操作 |
| 管理设备、产品、项目、OTA | `ur-iot` | IoT 设备管理 |
| 管理组织用户、角色、AI 智能体 | `ur-org-manage` | 租户管理员操作 |
| 能耗分析、电力集抄、预付费 | `ur-org-energy` | 能源管理 |
| 个人信息、访问令牌、续期 | `ur-console` | 普通用户操作 |

## 各应用详细文档

- [ur-platform-manage](ur-platform-manage/SKILL.md) — 平台管理
- [ur-iot](ur-iot/SKILL.md) — 物联网
- [ur-org-manage](ur-org-manage/SKILL.md) — 组织管理
- [ur-org-energy](ur-org-energy/SKILL.md) — 能源管理
- [ur-console](ur-console/SKILL.md) — 控制台

## 通用用法

```bash
# 配置（首次使用）
ur-iot setup

# 验证连通性
ur-iot check

# 调用 API
ur-iot api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'

# 查看 API schema
ur-iot schema

# 查看 token
ur-iot token --decode
```

## 认证方式

支持两种认证方式：

1. **密码登录**（推荐）：`ur-iot login` → session token → `token:` header
2. **AccessKey/JWT**：AccessKey + AccessSecret → HS256 JWT → `Authorization: Bearer` header

所有接口均为 POST 方法。请求格式 `{code, msg, data}`。
