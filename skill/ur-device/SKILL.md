---
name: ur-device
description: "设备管理：设备的增删改查、状态查询、属性控制、批量操作、设备认证、物模型。triggers: 设备管理, 设备列表, 设备控制, 设备在线状态, 属性上报, 网关, 三元组, 设备认证, 设备影子, OTA升级, 设备分享, 设备收藏, 物模型查询"
metadata:
  hermes:
    tags: [device, iot, control, ota, thing-model]
---


# ur-device — 设备管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/generate-api-lists.py` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 | 注意 |
|------|------|------|
| **deviceName** | 设备ID（唯一标识） | 不是设备名称！ |
| **deviceAlias** | 设备名称/别名 | 用于显示 |
| **三元组** | ProductID + DeviceName + DeviceSecret，设备身份凭证 |  |
| **物模型** | Property属性 + Event事件 + Action行为，标识符大驼峰命名 |  |
| **影子设备** | 设备期望状态的云端缓存，离线时可设置期望值 |  |
| **网关** | 管理子设备的代理设备，子设备通过网关间接接入 |  |
| **MQTT Topic** | $thing/up/property上报, $thing/down/property控制, $thing/up/event事件 |  |

## 角色权限说明

本域 API 按调用者角色区分权限：

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| 企业管理员 | 设备CRUD、属性控制、批量操作、物模型管理、网关拓扑 | 创建设备、设备属性控制 |
| 普通用户 | 设备分享、收藏、查看权限范围内的设备 | 分享设备给他人、收藏常用设备 |

---

### 企业管理员视角

**权限范围**：设备CRUD、属性控制、批量操作、物模型管理、网关拓扑

#### 典型场景

**创建设备**

手动创建单个设备或批量导入设备

- 涉及 CLI: `ur things device info create`（批量导入暂无 CLI，见下方缺口）
- 工作流: 选择产品 → 填写设备ID（deviceName） → 设置设备名称（deviceAlias） → 分配区域

请求示例：
```bash
ur things device info create -p p_smartswitch_001 -d switch-001 --alias "一楼走廊开关"
```

**设备属性控制**

开关灯 / 调节温度 / 设置参数（MQTT: $thing/down/property, method: control）

前置条件：
- 需要知道 productID（从产品列表 `ur things product info get-list` 获取）
- 需要知道 deviceName（从设备列表 `ur things device info get-list` 获取，注意 deviceName≠deviceAlias）
- 设备必须在线（isOnline=1）；离线时命令会缓存到影子设备（期望值）
- data 中属性 key 必须与物模型 identifier 完全一致（大驼峰，先查 `ur things schema get-list -p <productID>` 确认）

- 涉及 CLI: `ur things device control`, `ur things device log property`
- 工作流: 查询设备物模型（确认属性标识符） → 读取当前属性值（确认设备在线） → 发送控制命令 → 验证执行结果

请求示例：
```bash
ur things device control -p p_smartswitch_001 -d switch-001 --data '{"Power": 1}'
```

**网关拓扑管理**

添加子设备到网关 / 解绑子设备（MQTT: $gateway/up/topo, $gateway/down/topo）

- 涉及 CLI: `ur device gateway batch-create`, `ur device gateway batch-delete`
- 工作流: 确认网关设备在线 → 批量添加子设备 → 通知网关更新拓扑

**设备批量操作**

批量导入/更新/绑定设备

- 涉及 CLI: 批量导入/更新/绑定暂无 CLI 命令，见下方缺口
- 工作流: 准备批量数据（Excel/JSON） → 调用批量接口 → 检查导入结果


### 普通用户视角

**权限范围**：设备分享、收藏、查看权限范围内的设备

#### 典型场景

**分享设备给他人**

将设备分享给其他用户，设置分享权限

- 涉及 CLI: 设备分享暂无 CLI 命令，见下方缺口
- 工作流: 选择要分享的设备 → 填写目标用户账号 → 设置分享权限（读/写）

**收藏常用设备**

收藏设备便于快速访问

- 涉及 CLI: 设备收藏暂无 CLI 命令，见下方缺口
- 工作流: 选择要收藏的设备 → 收藏成功后在「我的设备」查看


## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| 设备信息管理 | 查询/创建/更新/删除设备、绑定/解绑网关、统计数量 | [cli/device-info.md](references/device-info.md) |
| 设备日志查询 | 属性/事件/命令/上下线/诊断/异常/SDK 日志 | [cli/device-log.md](references/device-log.md) |
| 网关管理 | 查询子设备列表、批量添加/删除子设备 | [cli/device-gateway.md](references/device-gateway.md) |
| 设备分组 | 分组增删改查、批量添加/移除设备 | [cli/device-group.md](references/device-group.md) |
| 设备配置 | 查询/更新/删除设备配置 | [cli/device-profile.md](references/device-profile.md) |
| 属性控制与行为 | 属性控制、行为调用、Mock数据、模拟上报 | [cli/device-control.md](references/device-control.md) |
| 物模型管理 | 查询/创建/更新/删除物模型、TSL导入/读取、API浏览 | [cli/schema.md](references/schema.md) |

> 完整命令帮助：`ur things device help`

## 典型业务场景

### 设备列表查询与筛选

**场景描述**：查看所有在线设备 / 查看某产品下设备 / 搜索设备名称

**涉及 CLI 命令**：
- `ur things device info get-list`
- `ur things device info count`
- `ur things device info get-one`

**工作流**：
1. `get-list` 获取列表（支持分页、筛选条件）
2. `count` 获取统计数据（在线/离线数量）
3. `get-one` 查看单个设备详情

### 设备事件与行为

**场景描述**：设备上报事件 / 调用设备行为（Event: type=info/alert/fault, Action: dir=up/down）

**涉及 CLI 命令**：
- `ur things device log event`
- `ur things device action send`
- `ur things device action get`

**工作流**：
1. 定义事件参数（params）或行为输入输出（input/output）
2. 设备上报事件或云端调用行为
3. 查询执行结果

### 设备认证与三元组获取

**场景描述**：获取设备 MQTT 认证凭证

**涉及 CLI**: 设备认证登录/注册暂无 CLI 命令（为设备端 API），见下方缺口

**工作流**：
1. 获取三元组（ProductID, DeviceName, DeviceSecret）
2. 设备使用三元组连接 MQTT


## 常用工作流

### 查询设备列表

```bash
ur things device info get-list --page 1 --size 10
```

### 获取设备详情

```bash
ur things device info get-one -p xxx -d yyy
```

### 控制设备属性

```bash
ur things device control -p xxx -d yyy --data '{"power":1}'
```


## 设备数据分析

设备属性历史数据的查询与分析见子域 Skill：
- **`ur-device-analytics/SKILL.md`** — 属性历史查询、趋势分析、聚合统计、报表生成

## 设备调试

设备日志查询和实时调试见子域 Skill：
- **`ur-device-debug/SKILL.md`** — 设备日志查询、属性控制、行为调用、事件发送、Mock数据生成

覆盖内容：
- 7 种日志查询：属性、事件、命令、上下线、异常、诊断(Hub)、SDK
- 4 种调试操作：属性控制、行为调用、事件发送、Mock数据生成
- 典型排障工作流：属性不更新、控制测试、离线排查

## 注意事项

- **deviceName 是设备唯一 ID（非显示名称）**，控制/查询时必须同时提供 productID + deviceName
- 属性标识符使用**大驼峰命名**（如 `CurrentTemperature`），`data` 字段 key 必须与物模型 identifier 完全一致
- 控制属性前先查询 isOnline=1 确认在线；离线设备命令会缓存到**影子设备**（期望值）
- protocolCode 默认值为 `urMqtt`（非旧值 `iThings`），2026-03-18 已修正
- 分页从 1 开始：`--page 1`，不是 0
- 所有接口使用 POST 方法，响应格式: `{code, msg, data}`，code=200 表示成功
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置，请求头需要 `app-id`、`tenant-code` 和认证头

## CLI 命令缺口

以下 API 暂无对应的 CLI Layer 2 命令，仍需使用 `ur api` 直接调用：

| API | 说明 | 状态 |
|-----|------|------|
| `/api/v1/things/device/info/batch-import` | 批量导入设备 | 缺口 |
| `/api/v1/things/device/info/batch-update` | 批量更新设备 | 缺口 |
| `/api/v1/things/device/info/batch-bind` | 批量绑定设备 | 缺口 |
| `/api/v1/things/device/gateway/batch-create` | 批量添加网关子设备 | **已可用** `ur device gateway batch-create` |
| `/api/v1/things/device/gateway/batch-delete` | 批量删除网关子设备 | **已可用** `ur device gateway batch-delete` |
| `/api/v1/things/device/gateway/get-list` | 查询网关子设备列表 | **已可用** `ur device gateway get-list` |
| `/api/v1/things/group/info/create` | 创建设备分组 | **已可用** `ur device group create` |
| `/api/v1/things/group/info/delete` | 删除设备分组 | **已可用** `ur device group delete` |
| `/api/v1/things/group/info/get-list` | 查询设备分组列表 | **已可用** `ur device group get-list` |
| `/api/v1/things/group/device/batch-create` | 批量添加设备到分组 | **已可用** `ur device group batch-create-device` |
| `/api/v1/things/group/device/batch-delete` | 批量从分组删除设备 | **已可用** `ur device group batch-delete-device` |
| `/api/v1/things/device/profile/get-list` | 查询设备配置列表 | **已可用** `ur device profile get-list` |
| `/api/v1/things/device/profile/get-one` | 查询设备配置详情 | **已可用** `ur device profile get-one` |
| `/api/v1/things/device/profile/update` | 更新设备配置 | **已可用** `ur device profile update` |
| `/api/v1/things/device/profile/delete` | 删除设备配置 | **已可用** `ur device profile delete` |
| `/api/v1/things/user/device/share/create` | 分享设备给他人 | 缺口 |
| `/api/v1/things/user/device/share/get-list` | 查询设备分享列表 | 缺口 |
| `/api/v1/things/user/device/collect/batch-create` | 批量收藏设备 | 缺口 |
| `/api/v1/things/user/device/collect/get-list` | 查询收藏设备列表 | 缺口 |
| `/api/v1/things/device/auth/login` | 设备认证登录 | 缺口 |
| `/api/v1/things/device/auth/register` | 设备认证注册 | 缺口 |
