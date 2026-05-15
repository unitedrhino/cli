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
> 如需更新，修改 `scripts/lib/swagger.ts` 中的域名定义后重新生成。

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
| 租户管理员 | 设备CRUD、属性控制、批量操作、物模型管理、网关拓扑 | 创建设备、设备属性控制 |
| 普通用户 | 设备分享、收藏、查看权限范围内的设备 | 分享设备给他人、收藏常用设备 |

---

### 租户管理员视角

**权限范围**：设备CRUD、属性控制、批量操作、物模型管理、网关拓扑

#### 典型场景

**创建设备**

手动创建单个设备或批量导入设备

- 涉及 API: `/api/v1/things/device/info/create`, `/api/v1/things/device/info/batch-import`
- 工作流: 选择产品 → 填写设备ID（deviceName） → 设置设备名称（deviceAlias） → 分配区域

请求示例：
```json
{
  "productID": "p_smartswitch_001",
  "deviceName": "switch-001",
  "deviceAlias": "一楼走廊开关"
}
```

**设备属性控制**

开关灯 / 调节温度 / 设置参数（MQTT: $thing/down/property, method: control）

前置条件：
- 需要知道 productID（从产品列表 /api/v1/things/product/info/get-list 获取）
- 需要知道 deviceName（从设备列表 /api/v1/things/device/info/get-list 获取，注意 deviceName≠deviceAlias）
- 设备必须在线（isOnline=1）；离线时命令会缓存到影子设备（期望值）
- data 中属性 key 必须与物模型 identifier 完全一致（大驼峰，先查 /api/v1/things/product/schema/get-list 确认）

- 涉及 API: `/api/v1/things/device/interact/property-control-send`, `/api/v1/things/device/msg/property-latest/get-list`
- 工作流: 查询设备物模型（确认属性标识符） → 读取当前属性值（确认设备在线） → 发送控制命令 → 验证执行结果

请求示例：
```json
{
  "productID": "p_smartswitch_001",
  "deviceName": "switch-001",
  "data": {
    "Power": 1
  }
}
```

**网关拓扑管理**

添加子设备到网关 / 解绑子设备（MQTT: $gateway/up/topo, $gateway/down/topo）

- 涉及 API: `/api/v1/things/device/gateway/batch-create`, `/api/v1/things/device/gateway/batch-delete`
- 工作流: 确认网关设备在线 → 批量添加子设备 → 通知网关更新拓扑

**设备批量操作**

批量导入/更新/绑定设备

- 涉及 API: `/api/v1/things/device/info/batch-import`, `/api/v1/things/device/info/batch-update`, `/api/v1/things/device/info/batch-bind`
- 工作流: 准备批量数据（Excel/JSON） → 调用批量接口 → 检查导入结果


### 普通用户视角

**权限范围**：设备分享、收藏、查看权限范围内的设备

#### 典型场景

**分享设备给他人**

将设备分享给其他用户，设置分享权限

- 涉及 API: `/api/v1/things/user/device/share/create`, `/api/v1/things/user/device/share/get-list`
- 工作流: 选择要分享的设备 → 填写目标用户账号 → 设置分享权限（读/写）

**收藏常用设备**

收藏设备便于快速访问

- 涉及 API: `/api/v1/things/user/device/collect/batch-create`, `/api/v1/things/user/device/collect/get-list`
- 工作流: 选择要收藏的设备 → 收藏成功后在「我的设备」查看


## API 参考

<!-- API_LIST:ur-device -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/auth/access` | 设备操作认证 | device |
| POST | `/api/v1/things/device/auth/login` | 设备登录认证 | device |
| POST | `/api/v1/things/device/auth/register` | 设备自动注册 | device |
| POST | `/api/v1/things/device/auth/root-check` | 鉴定是否是root账号 | device |
| POST | `/api/v1/things/device/auth5/access` | 设备操作认证 | device |
| POST | `/api/v1/things/device/auth5/login` | 设备登录认证 | device |
| POST | `/api/v1/things/device/edge/send/{handle}/{type}` | 设备使用http协议用云端交互,需要在http头中带上mqtt的账号密码(basic auth) | device |
| POST | `/api/v1/things/device/edge/upload-file` | 设备文件直传,需要在http头中带上mqtt的账号密码(basic auth) | device |
| POST | `/api/v1/things/device/gateway/batch-create` | 添加网关子设备 | admin |
| POST | `/api/v1/things/device/gateway/batch-delete` | 解绑子设备 | admin |
| POST | `/api/v1/things/device/gateway/get-list` | 获取子设备列表 | admin |
| POST | `/api/v1/things/device/group/batch-create` | 将设备加到多个分组中 | admin |
| POST | `/api/v1/things/device/group/batch-delete` | 删除设备所在分组 | admin |
| POST | `/api/v1/things/device/group/batch-update` | 更新设备所在分组 | admin |
| POST | `/api/v1/things/device/info/batch-bind` | 批量绑定 | admin |
| POST | `/api/v1/things/device/info/batch-import` | 批量导入设备 | admin |
| POST | `/api/v1/things/device/info/batch-update` | 批量更新设备 | admin |
| POST | `/api/v1/things/device/info/batch-update-import` | 导入批量更新设备 | admin |
| POST | `/api/v1/things/device/info/bind` | 绑定 | admin |
| POST | `/api/v1/things/device/info/bind/token/create` | 创建绑定token | admin |
| POST | `/api/v1/things/device/info/bind/token/get-one` | 绑定token状态查询 | admin |
| POST | `/api/v1/things/device/info/can-bind` | 是否可以绑定设备 | admin |
| POST | `/api/v1/things/device/info/count` | 设备统计详情 | admin |
| POST | `/api/v1/things/device/info/create` | 新增设备 | admin |
| POST | `/api/v1/things/device/info/delete` | 删除设备 | admin |
| POST | `/api/v1/things/device/info/get-list` | 获取设备列表 | admin |
| POST | `/api/v1/things/device/info/get-one` | 获取设备详情 | admin |
| POST | `/api/v1/things/device/info/move` | 转移设备到新设备上 | admin |
| POST | `/api/v1/things/device/info/ota/upgrade` | 设备升级,获取升级包手动升级 | admin |
| POST | `/api/v1/things/device/info/transfer` | 转让设备 | admin |
| POST | `/api/v1/things/device/info/unbind` | 解绑设备 | admin |
| POST | `/api/v1/things/device/info/update` | 更新设备 | admin |
| POST | `/api/v1/things/device/profile/delete` | 删除设备配置 | admin |
| POST | `/api/v1/things/device/profile/get-list` | 获取设备配置列表 | admin |
| POST | `/api/v1/things/device/profile/get-one` | 获取设备配置详情 | admin |
| POST | `/api/v1/things/device/profile/update` | 更新设备配置 | admin |
| POST | `/api/v1/things/device/schema/batch-create` | 批量创建设备物模型 | admin |
| POST | `/api/v1/things/device/schema/batch-delete` | 批量删除设备物模型 | admin |
| POST | `/api/v1/things/device/schema/create` | 创建设备物模型 | admin |
| POST | `/api/v1/things/device/schema/get-list` | 获取设备物模型列表 | admin |
| POST | `/api/v1/things/device/schema/tsl-read` | 获取设备物模型tsl | admin |
| POST | `/api/v1/things/device/schema/update` | 更新设备物模型 | admin |
| POST | `/api/v1/things/device/version/get-list` | 获取设备模块版本列表 | admin |
| POST | `/api/v1/things/device/version/get-one` | 获取设备模块版本详情 | admin |

<!-- END_API_LIST -->

权限: mixed

| 端点 | 说明 | 权限 |
|---|------|------|
| POST /api/v1/things/device/auth/access | 设备操作认证 | device |
| POST /api/v1/things/device/auth/login | 设备登录认证 | device |
| POST /api/v1/things/device/auth/register | 设备自动注册 | device |
| POST /api/v1/things/device/auth/root-check | 鉴定是否是root账号 | device |
| POST /api/v1/things/device/auth5/access | 设备操作认证 | device |
| POST /api/v1/things/device/auth5/login | 设备登录认证 | device |
| POST /api/v1/things/device/edge/send/{handle}/{type} | 设备使用http协议用云端交互,需要在http头中带上mqtt的账号密码(basic auth) | device |
| POST /api/v1/things/device/edge/upload-file | 设备文件直传,需要在http头中带上mqtt的账号密码(basic auth) | device |
| POST /api/v1/things/device/gateway/batch-create | 添加网关子设备 | 管理员 |
| POST /api/v1/things/device/gateway/batch-delete | 解绑子设备 | 管理员 |
| POST /api/v1/things/device/gateway/get-list | 获取子设备列表 | 管理员 |
| POST /api/v1/things/device/group/batch-create | 将设备加到多个分组中 | 管理员 |
| POST /api/v1/things/device/group/batch-delete | 删除设备所在分组 | 管理员 |
| POST /api/v1/things/device/group/batch-update | 更新设备所在分组 | 管理员 |
| POST /api/v1/things/device/info/batch-bind | 批量绑定 | 管理员 |
| POST /api/v1/things/device/info/batch-import | 批量导入设备 | 管理员 |
| POST /api/v1/things/device/info/batch-update | 批量更新设备 | 管理员 |
| POST /api/v1/things/device/info/batch-update-import | 导入批量更新设备 | 管理员 |
| POST /api/v1/things/device/info/bind | 绑定 | 管理员 |
| POST /api/v1/things/device/info/bind/token/create | 创建绑定token | 管理员 |
| POST /api/v1/things/device/info/bind/token/get-one | 绑定token状态查询 | 管理员 |
| POST /api/v1/things/device/info/can-bind | 是否可以绑定设备 | 管理员 |
| POST /api/v1/things/device/info/count | 设备统计详情 | 管理员 |
| POST /api/v1/things/device/info/create | 新增设备 | 管理员 |
| POST /api/v1/things/device/info/delete | 删除设备 | 管理员 |
| POST /api/v1/things/device/info/get-list | 获取设备列表 | 管理员 |
| POST /api/v1/things/device/info/get-one | 获取设备详情 | 管理员 |
| POST /api/v1/things/device/info/move | 转移设备到新设备上 | 管理员 |
| POST /api/v1/things/device/info/ota/upgrade | 设备升级,获取升级包手动升级 | 管理员 |
| POST /api/v1/things/device/info/transfer | 转让设备 | 管理员 |
| POST /api/v1/things/device/info/unbind | 解绑设备 | 管理员 |
| POST /api/v1/things/device/info/update | 更新设备 | 管理员 |
| POST /api/v1/things/device/interact/action-get-one | 获取调用设备行为的结果 | 管理员 |
| POST /api/v1/things/device/interact/action-resp | 回复设备行为调用结果 | 管理员 |
| POST /api/v1/things/device/interact/action-send | 调用设备行为 | 管理员 |
| POST /api/v1/things/device/interact/event-send | 下行事件通知设备 | 管理员 |
| POST /api/v1/things/device/interact/gateway-get-found-send | 请求网关上报拓扑关系 | 管理员 |
| POST /api/v1/things/device/interact/gateway-notify-bind-send | 通知网关绑定子设备 | 管理员 |
| POST /api/v1/things/device/interact/property-control-batch-send | 批量调用设备属性 | 管理员 |
| POST /api/v1/things/device/interact/property-control-get-one | 获取调用设备属性的结果 | 管理员 |
| POST /api/v1/things/device/interact/property-control-send | 调用设备属性 | 管理员 |
| POST /api/v1/things/device/interact/property-get-report-batch-send | 批量请求设备获取设备最新属性 | 管理员 |
| POST /api/v1/things/device/interact/property-get-report-send | 请求设备获取设备最新属性 | 管理员 |
| POST /api/v1/things/device/interact/schema-mock-gen | 生成物模型模拟数据 | 管理员 |
| POST /api/v1/things/device/msg/abnormal-log/get-list | 获取设备异常日志 | 管理员 |
| POST /api/v1/things/device/msg/event-log/get-list | 获取事件历史记录 | 管理员 |
| POST /api/v1/things/device/msg/gateway-can-bind/get-list | 获取网关可以绑定的子设备列表 | 管理员 |
| POST /api/v1/things/device/msg/hub-log/get-list | 获取云端诊断日志 | 管理员 |
| POST /api/v1/things/device/msg/property-agg/by-device/get-list | 弃用 | 管理员 |
| POST /api/v1/things/device/msg/property-agg/get-list | 弃用 | 管理员 |
| POST /api/v1/things/device/msg/property-latest-agg/get-list | 聚合属性最新值 | 管理员 |
| POST /api/v1/things/device/msg/property-latest/get-list | 获取最新属性记录 | 管理员 |
| POST /api/v1/things/device/msg/property-log-agg/by-device/get-list | 聚合属性历史记录,设备维度 | 管理员 |
| POST /api/v1/things/device/msg/property-log-agg/get-list | 聚合属性历史记录 | 管理员 |
| POST /api/v1/things/device/msg/property-log-latest/get-list | 弃用 | 管理员 |
| POST /api/v1/things/device/msg/property-log/batch-get-list | 批量获取单个id属性历史记录 | 管理员 |
| POST /api/v1/things/device/msg/property-log/get-list | 获取单个id属性历史记录 | 管理员 |
| POST /api/v1/things/device/msg/sdk-log/get-list | 获取设备sdk日志 | 管理员 |
| POST /api/v1/things/device/msg/send-log/get-list | 获取设备命令日志 | 管理员 |
| POST /api/v1/things/device/msg/shadow/get-list | 获取设备影子列表 | 管理员 |
| POST /api/v1/things/device/msg/status-log/get-list | 获取设备状态日志 | 管理员 |
| POST /api/v1/things/device/profile/delete | 删除设备配置 | 管理员 |
| POST /api/v1/things/device/profile/get-list | 获取设备配置列表 | 管理员 |
| POST /api/v1/things/device/profile/get-one | 获取设备配置详情 | 管理员 |
| POST /api/v1/things/device/profile/update | 更新设备配置 | 管理员 |
| POST /api/v1/things/device/schema/batch-create | 批量创建设备物模型 | 管理员 |
| POST /api/v1/things/device/schema/batch-delete | 批量删除设备物模型 | 管理员 |
| POST /api/v1/things/device/schema/create | 创建设备物模型 | 管理员 |
| POST /api/v1/things/device/schema/get-list | 获取设备物模型列表 | 管理员 |
| POST /api/v1/things/device/schema/tsl-read | 获取设备物模型tsl | 管理员 |
| POST /api/v1/things/device/schema/update | 更新设备物模型 | 管理员 |
| POST /api/v1/things/device/version/get-list | 获取设备模块版本列表 | 管理员 |
| POST /api/v1/things/device/version/get-one | 获取设备模块版本详情 | 管理员 |
| POST /api/v1/things/user/device/collect/batch-create | 批量收藏设备 | 所有用户 |
| POST /api/v1/things/user/device/collect/batch-delete | 批量取消收藏设备 | 所有用户 |
| POST /api/v1/things/user/device/collect/get-list | 获取收藏设备列表 | 所有用户 |
| POST /api/v1/things/user/device/share/batch-accept | 接受批量分享设备 | 所有用户 |
| POST /api/v1/things/user/device/share/batch-create | 生成批量分享设备二维码 | 所有用户 |
| POST /api/v1/things/user/device/share/batch-delete | 批量取消分享设备 | 所有用户 |
| POST /api/v1/things/user/device/share/batch-get-list | 获取批量分享的设备列表 | 所有用户 |
| POST /api/v1/things/user/device/share/create | 分享设备 | 所有用户 |
| POST /api/v1/things/user/device/share/delete | 取消分享设备 | 所有用户 |
| POST /api/v1/things/user/device/share/get-list | 获取分享设备列表 | 所有用户 |
| POST /api/v1/things/user/device/share/get-one | 获取分享设备详情 | 所有用户 |
| POST /api/v1/things/user/device/share/update | 更新分享设备信息 | 所有用户 |


## 典型业务场景

### 设备列表查询与筛选

**场景描述**：查看所有在线设备 / 查看某产品下设备 / 搜索设备名称

**涉及 API**：
- `/api/v1/things/device/info/get-list`
- `/api/v1/things/device/info/count`
- `/api/v1/things/device/info/get-one`

**工作流**：
1. get-list 获取列表（支持分页、筛选条件）
2. count 获取统计数据（在线/离线数量）
3. get-one 查看单个设备详情

### 设备事件与行为

**场景描述**：设备上报事件 / 调用设备行为（Event: type=info/alert/fault, Action: dir=up/down）

**涉及 API**：
- `/api/v1/things/device/msg/event-log/get-list`
- `/api/v1/things/device/interact/action-send`
- `/api/v1/things/device/interact/action-get-one`

**工作流**：
1. 定义事件参数（params）或行为输入输出（input/output）
2. 设备上报事件或云端调用行为
3. 查询执行结果

### 设备认证与三元组获取

**场景描述**：获取设备 MQTT 认证凭证

**涉及 API**：
- `/api/v1/things/device/auth/login`
- `/api/v1/things/device/auth/register`

**工作流**：
1. 获取三元组（ProductID, DeviceName, DeviceSecret）
2. 设备使用三元组连接 MQTT


## 常用工作流

### 查询设备列表

```bash
ur api /api/v1/things/device/info/get-list \
  --body '{"page":{"page":1,"size":10}}'
```

### 获取设备详情

```bash
ur api /api/v1/things/device/info/get-one \
  --body '{"productID":"xxx","deviceName":"yyy"}'
```

### 控制设备属性

```bash
ur api /api/v1/things/device/interact/property-control-send \
  --body '{"productID":"xxx","deviceName":"yyy","data":{"power":1}}'
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
- 分页从 1 开始：`page.page=1`，不是 0
- 所有接口使用 POST 方法，响应格式: `{code, msg, data}`，code=200 表示成功
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置，请求头需要 `app-id`、`tenant-code` 和认证头
