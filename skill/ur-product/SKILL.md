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
> 如需更新，修改 `scripts/generate-api-lists.py` 中的域名定义后重新生成。

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

## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| 物模型管理 | 查询/创建/更新/删除物模型，TSL导入/读取，API浏览 | [schema.md](references/product-schema.md) |
| OTA固件管理 | 固件查询/创建/更新/删除 | [ota.md](references/product-ota.md) |
| OTA任务管理 | OTA升级任务查询/创建/更新 | [ota.md](references/product-ota.md) |
| OTA模块管理 | 模块查询/创建/更新/删除 | [ota.md](references/product-ota.md) |
| 物模型模板 | 模板生成/结构验证/脚本生成 | [model.md](references/product-model.md) |
| 协议脚本 | 脚本验证/模板生成 | [script.md](references/product-script.md) |
| 聚合查询 | 设备属性聚合值查询 | [agg.md](references/product-agg.md) |

> 完整命令帮助：`ur things help`

## 典型业务场景

### 创建产品并定义物模型

**场景描述**：创建新产品 / 定义物模型 / 配置协议

**前置条件**：
- 品类 ID 可选（通过 `ur things product category get-list` 查询，不传则为无品类产品）

**涉及 CLI**：
- `ur things product info create`
- `ur things product info get-list`
- `ur things schema get-list`

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

**涉及 CLI**：
- `ur things schema get-list`
- `ur things schema create`
- `ur things schema update`

**工作流**：
1. 查询现有物模型
2. 添加属性/服务/事件
3. 配置数据类型（Define）
4. 更新或删除物模型元素

### 物模型 MQTT 协议

**场景描述**：设备与平台物模型交互（Topic: $thing/up/property, $thing/down/property）

**涉及 CLI**：
- `ur things device info get-list`
- `ur things device info create`

**工作流**：
1. 属性上报: $thing/up/property, method: report
2. 属性控制: $thing/down/property, method: control
3. 事件上报: $thing/up/event, method: eventPost
4. 行为调用: $thing/down/action, method: action

### 协议脚本开发

**场景描述**：自定义Go脚本处理设备消息（yaegi解释器，支持json/gjson/utils/deviceMsg/dm包）

**涉及 CLI**：
- `ur things protocol script create`
- `ur things protocol script update`
- `ur things protocol script get-list`

**工作流**：
1. 选择触发模式: UpBefore/UpAfter/DownBefore/DownAfter
2. 编写Handle函数处理PublishMsg
3. 本地测试验证（script_test.go）
4. 部署并监控日志

### OTA 固件升级

**场景描述**：上传固件 / 创建升级任务 / 监控升级进度（step: 1-100百分比, -1~-4错误码）

**涉及 CLI**：
- `ur things ota info create`
- `ur things ota info get-list`
- `ur things device info get-list`

**工作流**：
1. 上传固件包（差分/全量）
2. 创建升级任务（选择目标设备）
3. 监控升级进度（step: 1→100 或 -1~-4）

### 通用物模型批量导入

**场景描述**：批量导入通用物模型定义（Excel → JSON → ZIP → OSS → CLI导入）

**涉及 CLI**：
- `ur things schema common batch-import`
- `ur things schema common get-list`

**工作流**：
1. 准备Excel模板（identifier/name/define）
2. 转换为CommonSchemaInfo JSON格式
3. 打包ZIP并上传OSS
4. 调用 batch-import CLI 导入


## 常用工作流

### 查询产品列表

```bash
ur things product info get-list
```

### 查询物模型列表

```bash
ur things schema get-list
```


## 注意事项

- 协议脚本使用 **yaegi Go 解释器**，支持 json/gjson/utils/deviceMsg/dm 包，**不支持 CGO**
- OTA 进度码：1-100 表示百分比，-1=下载失败，-2=校验失败，-3=烧录失败，-4=版本不匹配
- 物模型属性标识符使用**大驼峰命名**（如 `CurrentTemperature`），设备上报/控制时 key 必须完全一致
- 创建产品时 `protocolCode` 应填 `urMqtt`
- 所有接口使用 POST 方法，响应格式: `{code, msg, data}`，code=200 表示成功
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置，请求头需要 `app-id`、`tenant-code` 和认证头
