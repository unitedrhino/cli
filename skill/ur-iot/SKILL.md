---
name: ur-iot
description: "ur-iot — 联犀 SaaS 平台物联网总览。覆盖设备管理、产品管理、项目管理、场景联动、OTA升级、数据流转、告警管理。triggers: 物联网, IoT, 设备管理, 产品管理, 项目管理, 场景联动, OTA升级, 物模型, 协议网关"
metadata:
  hermes:
    tags: [iot, device, product, project, ota, protocol, thing-model]
---

# ur-iot — 物联网

> **本文件是物联网总览入口**，具体操作请查阅对应的子域 Skill。
>
> 如需使用 CLI 命令，优先使用 `ur things`、`ur user`、`ur device` 等 Layer 2 命令。
> 如某功能暂无 CLI 命令，可使用 `ur api /api/v1/...` 直接调用，详见各子域的 `references/api/` 文档。

## 应用信息

- **AppID**: 200
- **TenantCode**: platform
- **可调用权限**: platform, admin, all, public

## 功能架构与子域索引

| 功能域 | 说明 | 对应 Skill | CLI 命令根 |
|--------|------|-----------|-----------|
| **设备管理** | 设备CRUD、状态查询、属性控制、批量操作、网关拓扑、设备分享 | `ur-device` | `ur things device` |
| **设备调试** | 日志查询（属性/事件/命令/上下线/诊断/SDK）、实时调试 | `ur-device-debug` | `ur things device log` |
| **设备分析** | 属性历史查询、趋势分析、聚合统计 | `ur-device-analytics` | `ur things device log property` |
| **产品管理** | 产品CRUD、物模型、品类管理、OTA升级包 | `ur-product` | `ur things product` |
| **项目管理** | 项目/区域/分组CRUD、设备分配、数据权限 | `ur-project` | `ur things project` |
| **场景联动** | 场景规则编排（if/when/then）、手动触发、执行日志 | `scene-linkage` | `ur things scene` |
| **告警管理** | 告警配置、告警记录处理、告警与场景关联 | `ur-ai` | `ur alarm` |
| **AI 能力** | AI Agent、MCP 协议、设备智能绑定 | `ur-ai` | `ur ai` |
| **用户管理** | 登录注册、角色权限、部门、字典、消息 | `ur-user` | `ur user` |
| **企业管理** | 企业CRUD、用户邀请、应用绑定、续费 | `ur-tenant` | `ur tenant` |
| **系统管理** | 应用/菜单、文件上传、Hook、批量聚合 | `ur-system` | `ur system` |
| **物模型** | 物模型规范、TSL导入导出、Mock数据生成 | `thing-model` | `ur things schema` |
| **协议脚本** | 协议网关、协议容器、协议脚本管理 | `protocol-script` | - |
| **AI 工具** | AI 工具创建、发布、编排 | `ai-tool` | - |

## 常用 CLI 命令速查

### 设备相关

```bash
# 查询设备列表
ur things device info get-list --page 1 --size 10

# 获取设备详情
ur things device info get-one -p <productID> -d <deviceName>

# 控制设备属性
ur things device control -p <productID> -d <deviceName> --data '{"PowerSwitch":1}'

# 查询设备日志
ur things device log property -p <productID> -d <deviceName>
ur things device log event -p <productID> -d <deviceName>
ur things device log status -p <productID> -d <deviceName>

# 查询物模型
ur things schema get-list -p <productID>
```

### 产品相关

```bash
# 查询产品列表
ur things product info get-list --page 1 --size 10

# 查询产品物模型
ur things schema get-list -p <productID>

# OTA 升级包管理
ur things ota info get-list --page 1 --size 10
```

### 项目相关

```bash
# 查询项目列表
ur things project info get-list --page 1 --size 10

# 查询区域列表
ur things area info get-list --page 1 --size 10

# 查询分组列表
ur things group info get-list --page 1 --size 10
```

### 用户与企业

```bash
# 登录
ur user self login --userName <账号> --password <密码>

# 获取当前用户信息
ur user self get-one

# 查询企业用户列表
ur tenant user get-list --page 1 --size 10

# 邀请用户加入企业
ur tenant user invite --userID <用户ID>
```

### 系统通用

```bash
# 文件上传
ur system upload upload-file --file <文件路径>

# 批量聚合接口调用
ur system batch-agg --body '{"apis":[{"path":"/api/v1/system/user/self/get-one","body":{}}]}'
```

### 场景与告警

```bash
# 查询场景列表
ur things scene info get-list --page 1 --size 10

# 手动触发场景
ur things scene info manually-trigger --id <sceneID>

# 查询告警规则
ur alarm rule get-list --page 1 --size 10

# 查询告警事件
ur alarm event get-list --size 10

# 处理告警事件
ur alarm event deal --id <eventID> --action acked
```

## 典型工作流

### 工作流 1：设备接入全流程

1. **创建产品** → `ur things product info create`
2. **配置物模型** → `ur things schema get-list` / `ur things schema create`
3. **创建设备** → `ur things device info create`
4. **获取三元组** → `ur things device info get-one`（查看 cert 字段）
5. **设备连接 MQTT** → 设备端使用三元组连接
6. **查看设备状态** → `ur things device info get-one`（查看 isOnline）
7. **控制设备** → `ur things device control`

### 工作流 2：场景联动配置

1. **查询场景** → `ur things scene info get-list`
2. **创建场景** → `ur things scene info create --body '{...}'`
3. **验证场景 JSON** → `ur scene validate --body '{...}'`
4. **手动触发测试** → `ur things scene info manually-trigger --id <id>`
5. **查看执行日志** → `ur things scene log get-list`

### 工作流 3：OTA 固件升级

1. **创建 OTA 模块** → `ur things ota module create`（如需要）
2. **上传升级包** → `ur system upload upload-file`
3. **创建升级包** → `ur things ota info create`
4. **创建升级任务** → `ur api /api/v1/things/ota/firmware/job/create`
5. **查看升级进度** → `ur api /api/v1/things/ota/firmware/device/get-list`

## CLI 命令缺口

以下高频功能暂无 Layer 2 CLI 命令，需使用 `ur api` 直接调用：

| 功能 | API 路径 | 对应 Skill |
|------|---------|-----------|
| 批量导入设备 | `/api/v1/things/device/info/batch-import` | ur-device |
| 设备分享 | `/api/v1/things/user/device/share/create` | ur-device |
| 设备收藏 | `/api/v1/things/user/device/collect/batch-create` | ur-device |
| 设备认证（三元组） | `/api/v1/things/device/auth/login` | ur-device |
| OTA 升级任务 | `/api/v1/things/ota/firmware/job/create` | ur-product |
| TSL 导入 | `/api/v1/things/product/schema/tsl-import` | thing-model |
| 协议容器管理 | `/api/v1/things/protocol/container/install` | protocol-script |
| 协议脚本调试 | `/api/v1/things/protocol/script/debug` | protocol-script |
| AI 工具创建 | `/api/v1/ai/tool/create` | ai-tool |

> 完整 API 端点速查见各子域 `references/api/{domain}-all-endpoints.md`
