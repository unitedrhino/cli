---
name: ur-product
description: "产品管理 + 物模型：产品的增删改查、物模型管理、协议脚本、OTA固件管理。triggers: 产品定义, 产品管理, 物模型, 协议脚本, OTA固件, 品类, 产品列表, 属性定义, 事件定义, 行为定义, Thing Model"
metadata:
  hermes:
    tags: [product, thing-model, protocol, ota, firmware]
---


# ur-product — 产品管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/lib/swagger.ts` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 |
|------|------|
| **产品** | 设备的模板，定义物模型、协议脚本 |
| **品类** | 产品分类，预设物模型模板（如智能灯、传感器） |
| **物模型** | Property（属性）+ Event（事件）+ Action（行为），标识符大驼峰命名 |
| **数据类型** | bool/int/float/string/enum/timestamp/struct/array，Define结构定义 |
| **协议脚本** | yaegi Go解释器，支持UpBefore/UpAfter/DownBefore/DownAfter四种触发模式 |
| **OTA** | 固件升级管理，支持差分/全量升级，进度码1-100/-1~-4 |
| **三元组** | ProductID + DeviceName + DeviceSecret，设备身份凭证 |

## API 参考

<!-- API_LIST:ur-product -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/product/category/batch-export` | 批量导出产品品类 | admin |
| POST | `/api/v1/things/product/category/batch-import` | 批量导入产品品类 | admin |
| POST | `/api/v1/things/product/category/create` | 新增产品品类 | admin |
| POST | `/api/v1/things/product/category/delete` | 删除产品品类 | admin |
| POST | `/api/v1/things/product/category/get-list` | 获取产品品类列表 | admin |
| POST | `/api/v1/things/product/category/get-one` | 获取产品品类详情 | admin |
| POST | `/api/v1/things/product/category/schema/batch-create` | 批量新增产品品类物模型 | admin |
| POST | `/api/v1/things/product/category/schema/batch-delete` | 批量删除产品品类物模型 | admin |
| POST | `/api/v1/things/product/category/schema/batch-update` | 批量更新产品品类物模型 | admin |
| POST | `/api/v1/things/product/category/schema/get-list` | 获取产品品类物模型列表 | admin |
| POST | `/api/v1/things/product/category/update` | 更新产品品类 | admin |
| POST | `/api/v1/things/product/config/update` | 更新配置 | admin |
| POST | `/api/v1/things/product/custom/get-one` | 获取产品自定义信息详情 | admin |
| POST | `/api/v1/things/product/custom/update` | 更新产品自定义信息 | admin |
| POST | `/api/v1/things/product/info/batch-export` | 批量导出产品 | admin |
| POST | `/api/v1/things/product/info/batch-import` | 批量导入产品 | admin |
| POST | `/api/v1/things/product/info/create` | 新增产品 | admin |
| POST | `/api/v1/things/product/info/delete` | 删除产品 | admin |
| POST | `/api/v1/things/product/info/get-list` | 获取产品信息列表 | admin |
| POST | `/api/v1/things/product/info/get-one` | 获取产品详情 | admin |
| POST | `/api/v1/things/product/info/init` | 初始化产品 | admin |
| POST | `/api/v1/things/product/info/update` | 更新产品 | admin |
| POST | `/api/v1/things/product/remote-config/create` | 创建配置 | admin |
| POST | `/api/v1/things/product/remote-config/get-list` | 获取配置列表 | admin |
| POST | `/api/v1/things/product/remote-config/lastest-read` | 获取最新配置 | admin |
| POST | `/api/v1/things/product/remote-config/push-all` | 推送配置 | admin |
| POST | `/api/v1/things/product/schema/batch-create` | 批量创建产品物模型 | admin |
| POST | `/api/v1/things/product/schema/create` | 创建产品物模型 | admin |
| POST | `/api/v1/things/product/schema/delete` | 删除产品物模型 | admin |
| POST | `/api/v1/things/product/schema/get-list` | 获取产品物模型 | admin |
| POST | `/api/v1/things/product/schema/tsl-import` | 导入产品物模型tsl | admin |
| POST | `/api/v1/things/product/schema/tsl-read` | 获取产品物模型tsl | admin |
| POST | `/api/v1/things/product/schema/update` | 更新产品物模型 | admin |

<!-- END_API_LIST -->

权限: 管理员

| 端点 | 说明 | 权限 |
|---|------|------|
| POST /api/v1/things/ota/firmware/device/cancel | 取消指定任务下的升级 | 管理员 |
| POST /api/v1/things/ota/firmware/device/confirm | 确认升级设备 | 管理员 |
| POST /api/v1/things/ota/firmware/device/get-list | 查询升级设备列表 | 管理员 |
| POST /api/v1/things/ota/firmware/device/retry | 重试设备升级 | 管理员 |
| POST /api/v1/things/ota/firmware/info/create | 添加升级包 | 管理员 |
| POST /api/v1/things/ota/firmware/info/delete | 删除升级包 | 管理员 |
| POST /api/v1/things/ota/firmware/info/get-list | 升级包列表 | 管理员 |
| POST /api/v1/things/ota/firmware/info/get-one | 查询升级包 | 管理员 |
| POST /api/v1/things/ota/firmware/info/update | 修改升级包 | 管理员 |
| POST /api/v1/things/ota/firmware/job/create | 创建升级任务 | 管理员 |
| POST /api/v1/things/ota/firmware/job/get-list | 获取升级包下的升级任务批次列表 | 管理员 |
| POST /api/v1/things/ota/firmware/job/get-one | 查询指定升级批次的详情 | 管理员 |
| POST /api/v1/things/ota/firmware/job/update | 更新升级批次 | 管理员 |
| POST /api/v1/things/ota/module/info/create | 创建模块 | 管理员 |
| POST /api/v1/things/ota/module/info/delete | 删除模块 | 管理员 |
| POST /api/v1/things/ota/module/info/get-list | 获取模块列表 | 管理员 |
| POST /api/v1/things/ota/module/info/get-one | 查询模块详情 | 管理员 |
| POST /api/v1/things/ota/module/info/update | 更新模块 | 管理员 |
| POST /api/v1/things/product/category/batch-export | 批量导出产品品类 | 管理员 |
| POST /api/v1/things/product/category/batch-import | 批量导入产品品类 | 管理员 |
| POST /api/v1/things/product/category/create | 新增产品品类 | 管理员 |
| POST /api/v1/things/product/category/delete | 删除产品品类 | 管理员 |
| POST /api/v1/things/product/category/get-list | 获取产品品类列表 | 管理员 |
| POST /api/v1/things/product/category/get-one | 获取产品品类详情 | 管理员 |
| POST /api/v1/things/product/category/schema/batch-create | 批量新增产品品类物模型 | 管理员 |
| POST /api/v1/things/product/category/schema/batch-delete | 批量删除产品品类物模型 | 管理员 |
| POST /api/v1/things/product/category/schema/batch-update | 批量更新产品品类物模型 | 管理员 |
| POST /api/v1/things/product/category/schema/get-list | 获取产品品类物模型列表 | 管理员 |
| POST /api/v1/things/product/category/update | 更新产品品类 | 管理员 |
| POST /api/v1/things/product/config/update | 更新配置 | 管理员 |
| POST /api/v1/things/product/custom/get-one | 获取产品自定义信息详情 | 管理员 |
| POST /api/v1/things/product/custom/update | 更新产品自定义信息 | 管理员 |
| POST /api/v1/things/product/info/batch-export | 批量导出产品 | 管理员 |
| POST /api/v1/things/product/info/batch-import | 批量导入产品 | 管理员 |
| POST /api/v1/things/product/info/create | 新增产品 | 管理员 |
| POST /api/v1/things/product/info/delete | 删除产品 | 管理员 |
| POST /api/v1/things/product/info/get-list | 获取产品信息列表 | 管理员 |
| POST /api/v1/things/product/info/get-one | 获取产品详情 | 管理员 |
| POST /api/v1/things/product/info/init | 初始化产品 | 管理员 |
| POST /api/v1/things/product/info/update | 更新产品 | 管理员 |
| POST /api/v1/things/product/remote-config/create | 创建配置 | 管理员 |
| POST /api/v1/things/product/remote-config/get-list | 获取配置列表 | 管理员 |
| POST /api/v1/things/product/remote-config/lastest-read | 获取最新配置 | 管理员 |
| POST /api/v1/things/product/remote-config/push-all | 推送配置 | 管理员 |
| POST /api/v1/things/product/schema/batch-create | 批量创建产品物模型 | 管理员 |
| POST /api/v1/things/product/schema/create | 创建产品物模型 | 管理员 |
| POST /api/v1/things/product/schema/delete | 删除产品物模型 | 管理员 |
| POST /api/v1/things/product/schema/get-list | 获取产品物模型 | 管理员 |
| POST /api/v1/things/product/schema/tsl-import | 导入产品物模型tsl | 管理员 |
| POST /api/v1/things/product/schema/tsl-read | 获取产品物模型tsl | 管理员 |
| POST /api/v1/things/product/schema/update | 更新产品物模型 | 管理员 |
| POST /api/v1/things/protocol/config/create | 创建协议配置 | 管理员 |
| POST /api/v1/things/protocol/config/delete | 删除协议配置 | 管理员 |
| POST /api/v1/things/protocol/config/get-list | 获取协议配置列表 | 管理员 |
| POST /api/v1/things/protocol/config/get-one | 获取协议配置详情 | 管理员 |
| POST /api/v1/things/protocol/config/update | 更新协议配置 | 管理员 |
| POST /api/v1/things/protocol/info/create | 新增自定义协议 | 管理员 |
| POST /api/v1/things/protocol/info/delete | 删除自定义协议 | 管理员 |
| POST /api/v1/things/protocol/info/get-list | 获取自定义协议信息列表 | 管理员 |
| POST /api/v1/things/protocol/info/get-one | 获取自定义协议详情 | 管理员 |
| POST /api/v1/things/protocol/info/update | 更新自定义协议 | 管理员 |
| POST /api/v1/things/protocol/script/batch-export | 批量导出协议脚本 | 管理员 |
| POST /api/v1/things/protocol/script/batch-import | 批量导入协议脚本 | 管理员 |
| POST /api/v1/things/protocol/script/create | 新增协议脚本 | 管理员 |
| POST /api/v1/things/protocol/script/debug | 协议脚本调试 | 管理员 |
| POST /api/v1/things/protocol/script/delete | 删除协议脚本 | 管理员 |
| POST /api/v1/things/protocol/script/device/create | 新增协议脚本设备 | 管理员 |
| POST /api/v1/things/protocol/script/device/delete | 删除协议脚本设备 | 管理员 |
| POST /api/v1/things/protocol/script/device/get-list | 获取协议脚本设备列表 | 管理员 |
| POST /api/v1/things/protocol/script/device/get-one | 获取协议脚本设备详情 | 管理员 |
| POST /api/v1/things/protocol/script/device/update | 更新协议脚本设备 | 管理员 |
| POST /api/v1/things/protocol/script/get-list | 获取协议脚本列表 | 管理员 |
| POST /api/v1/things/protocol/script/get-one | 获取协议脚本详情 | 管理员 |
| POST /api/v1/things/protocol/script/update | 更新协议脚本 | 管理员 |
| POST /api/v1/things/protocol/service/delete | 删除自定义协议服务器 | 管理员 |
| POST /api/v1/things/protocol/service/get-list | 获取自定义协议服务器信息列表 | 管理员 |
| POST /api/v1/things/protocol/sync/device | 设备同步(如果该协议不支持会返回不支持) | 管理员 |
| POST /api/v1/things/protocol/sync/product | 产品同步(如果该协议不支持会返回不支持) | 管理员 |
| POST /api/v1/things/schema/common/batch-export | 批量导出通用物模型 | 管理员 |
| POST /api/v1/things/schema/common/batch-import | 批量导入通用物模型 | 管理员 |
| POST /api/v1/things/schema/common/create | 新增通用物模型 | 管理员 |
| POST /api/v1/things/schema/common/delete | 删除通用物模型 | 管理员 |
| POST /api/v1/things/schema/common/get-list | 获取通用物模型列表 | 管理员 |
| POST /api/v1/things/schema/common/init | 初始化通用物模型 | 管理员 |
| POST /api/v1/things/schema/common/mock-gen | 从参数生成物模型定义 | 管理员 |
| POST /api/v1/things/schema/common/update | 更新通用物模型 | 管理员 |


## 典型业务场景

### 创建产品并定义物模型

**场景描述**：创建新产品 / 定义物模型 / 配置协议

**前置条件**：
- 品类 ID 可选（从 /api/v1/things/product/category/get-list 查询，不传则为无品类产品）

**涉及 API**：
- `/api/v1/things/product/info/create`
- `/api/v1/things/schema/create`
- `/api/v1/things/protocol/script/create`

**工作流**：
1. 创建产品（protocolCode 填 urMqtt，deviceType: 1=直连/2=网关/3=子设备）
2. 定义物模型（Property属性/Event事件/Action行为）
3. 配置协议脚本（可选）

**请求示例**：
```json
{
  "name": "智能开关",
  "protocolCode": "urMqtt",
  "deviceType": 1
}
```

### 物模型管理

**场景描述**：查看/修改产品物模型（属性mode: r/rw，事件type: info/alert/fault，行为dir: up/down）

**涉及 API**：
- `/api/v1/things/schema/get-list`
- `/api/v1/things/schema/create`
- `/api/v1/things/schema/update`
- `/api/v1/things/schema/delete`

**工作流**：
1. 查询现有物模型
2. 添加属性/服务/事件
3. 配置数据类型（Define）
4. 更新或删除物模型元素

### 物模型 MQTT 协议

**场景描述**：设备与平台物模型交互（Topic: $thing/up/property, $thing/down/property）

**涉及 API**：
- `/api/v1/things/device/interact/property-control-send`
- `/api/v1/things/device/msg/property-latest/get-list`

**工作流**：
1. 属性上报: $thing/up/property, method: report
2. 属性控制: $thing/down/property, method: control
3. 事件上报: $thing/up/event, method: eventPost
4. 行为调用: $thing/down/action, method: action

### 协议脚本开发

**场景描述**：自定义Go脚本处理设备消息（yaegi解释器，支持json/gjson/utils/deviceMsg/dm包）

**涉及 API**：
- `/api/v1/things/protocol/script/create`
- `/api/v1/things/protocol/script/update`
- `/api/v1/things/protocol/script/debug`

**工作流**：
1. 选择触发模式: UpBefore/UpAfter/DownBefore/DownAfter
2. 编写Handle函数处理PublishMsg
3. 本地测试验证（script_test.go）
4. 部署并监控日志

### OTA 固件升级

**场景描述**：上传固件 / 创建升级任务 / 监控升级进度（step: 1-100百分比, -1~-4错误码）

**涉及 API**：
- `/api/v1/things/ota/firmware/info/create`
- `/api/v1/things/ota/firmware/job/create`
- `/api/v1/things/ota/firmware/device/get-list`

**工作流**：
1. 上传固件包（差分/全量）
2. 创建升级任务（选择目标设备）
3. 监控升级进度（step: 1→100 或 -1~-4）

### 通用物模型批量导入

**场景描述**：批量导入通用物模型定义（Excel → JSON → ZIP → OSS → API导入）

**涉及 API**：
- `/api/v1/system/common/upload-file`
- `/api/v1/things/schema/common/batch-import`

**工作流**：
1. 准备Excel模板（identifier/name/define）
2. 转换为CommonSchemaInfo JSON格式
3. 打包ZIP并上传OSS
4. 调用batch-import API导入


## 常用工作流

### 查询产品列表

```bash
ur api /api/v1/things/product/info/get-list \
  --body '{"page":{"page":1,"size":10}}'
```

### 查询物模型列表

```bash
ur api /api/v1/things/schema/get-list \
  --body '{"productID":"xxx"}'
```


## 注意事项

- 协议脚本使用 **yaegi Go 解释器**，支持 json/gjson/utils/deviceMsg/dm 包，**不支持 CGO**
- OTA 进度码：1-100 表示百分比，-1=下载失败，-2=校验失败，-3=烧录失败，-4=版本不匹配
- 物模型属性标识符使用**大驼峰命名**（如 `CurrentTemperature`），设备上报/控制时 key 必须完全一致
- 创建产品时 `protocolCode` 应填 `urMqtt`
- 所有接口使用 POST 方法，响应格式: `{code, msg, data}`，code=200 表示成功
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置，请求头需要 `app-id`、`tenant-code` 和认证头
