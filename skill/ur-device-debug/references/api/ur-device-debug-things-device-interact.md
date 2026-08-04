# ur-device-debug things/device/interact

获取调用设备行为的结果 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/interact/action-get-one` | 获取调用设备行为的结果 | admin |
| POST | `/api/v1/things/device/interact/action-resp` | 回复设备行为调用结果 | admin |
| POST | `/api/v1/things/device/interact/action-send` | 调用设备行为 | admin |
| POST | `/api/v1/things/device/interact/event-send` | 下行事件通知设备 | admin |
| POST | `/api/v1/things/device/interact/message-send` | 下行透传消息给设备 | admin |
| POST | `/api/v1/things/device/interact/property-control-batch-send` | 批量调用设备属性 | admin |
| POST | `/api/v1/things/device/interact/property-control-get-one` | 获取调用设备属性的结果 | admin |
| POST | `/api/v1/things/device/interact/property-control-send` | 调用设备属性 | admin |
| POST | `/api/v1/things/device/interact/property-get-report-batch-send` | 批量请求设备获取设备最新属性 | admin |
| POST | `/api/v1/things/device/interact/property-get-report-send` | 请求设备获取设备最新属性 | admin |
| POST | `/api/v1/things/device/interact/schema-mock-gen` | 生成物模型模拟数据 | admin |

## 详细说明

### POST `/api/v1/things/device/interact/action-get-one`

**说明**: 获取调用设备行为的结果

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名 |
| `msgToken` | string | 是 | 调用id |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "deviceName": "string",
  "msgToken": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": 1,
    "msg": "string",
    "msgToken": "string",
    "outputParams": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/action-get-one \
  --body '{"deviceName": "string", "msgToken": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/interact/action-resp`

**说明**: 回复设备行为调用结果

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | integer | 是 | 设备返回状态码 (格式: int64) |
| `deviceName` | string | 是 | 设备名 |
| `msg` | string | 是 | 返回状态 |
| `msgToken` | string | 是 | 调用id |
| `outputParams` | string | 否 | 输出参数 注意：此字段可能返回 null，表示取不到有效值。 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "code": 1,
  "deviceName": "string",
  "msg": "string",
  "msgToken": "string",
  "outputParams": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/action-resp \
  --body '{"code": 1, "deviceName": "string", "msg": "string", "msgToken": "string", "outputParams": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/interact/action-send`

**说明**: 调用设备行为

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionID` | string | 是 | 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义 |
| `deviceName` | string | 是 | 设备名 |
| `inputParams` | string | 否 | 输入参数 |
| `isAsync` | boolean | 否 | 是否异步操作 异步情况通过获取接口来获取 (格式: boolean) |
| `option` | object | 否 |  |
| `option.requestTimeout` | integer | 否 | 请求超时,超时后会进行重试 (格式: int64) |
| `option.retryInterval` | integer | 否 | 重试间隔 (格式: int64) |
| `option.timeoutToFail` | integer | 否 | 超时失败时间 (格式: int64) |
| `outputParams` | string | 否 | 输出参数(模拟的时候才需要填写) ,如果填写该参数,那么不会下发给设备,直接记录操作记录 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "actionID": "string",
  "deviceName": "string",
  "inputParams": "string",
  "isAsync": true,
  "option": {
    "requestTimeout": 1,
    "retryInterval": 1,
    "timeoutToFail": 1
  },
  "outputParams": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": 1,
    "msg": "string",
    "msgToken": "string",
    "outputParams": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/action-send \
  --body '{"actionID": "string", "deviceName": "string", "inputParams": "string", "isAsync": true, "option": {"requestTimeout": 1, "retryInterval": 1, "timeoutToFail": 1}, "outputParams": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/interact/event-send`

**说明**: 下行事件通知设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名 |
| `eventID` | string | 是 | 产品数据模板中事件功能的标识符，由开发者自行根据设备的应用场景定义 |
| `params` | string | 是 | 通知内容,json字符串 |
| `productID` | string | 是 | 产品id 获取产品id下的所有设备信息 |
| `type` | string | 否 | 事件类型:info(默认)：信息,alert：告警,fault：故障 |

**请求示例**:
```json
{
  "deviceName": "string",
  "eventID": "string",
  "params": "string",
  "productID": "string",
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/event-send \
  --body '{"deviceName": "string", "eventID": "string", "params": "string", "productID": "string", "type": "string"}'
```

### POST `/api/v1/things/device/interact/message-send`

**说明**: 下行透传消息给设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 目标设备名称 |
| `handle` | string | 是 | 下行通道，仅支持thing、ota、config、log、shadow、gateway、ext |
| `payload` | string | 是 | 原始文本载荷 |
| `productID` | string | 是 | 目标产品ID |
| `type` | string | 是 | 消息类型，必须是单个Topic路径段 |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "handle": "string",
  "payload": "string",
  "productID": "string",
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/message-send \
  --body '{"deviceName": "示例名称", "handle": "string", "payload": "string", "productID": "string", "type": "string"}'
```

### POST `/api/v1/things/device/interact/property-control-batch-send`

**说明**: 批量调用设备属性

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id,传了优先从项目区域中获取设备列表 |
| `areaIDPath` | string | 否 | 根据区域路径控制设备 |
| `data` | string | 是 | 属性数据, JSON格式字符串, 注意字段需要在物模型属性里定义 |
| `deviceNames` | array[string] | 否 | 设备名列表 |
| `devices` | array[DeviceCore] | 否 | 如果是不同的产品,则传这个字段,上面两个参数填了优先使用 |
| `groupID` | string | 否 | 分组ID,传了会从分组下获取设备 |
| `isAsync` | boolean | 否 | 是否异步操作 异步情况通过获取接口来获取 (格式: boolean) |
| `productID` | string | 否 | 产品id |
| `shadowControl` | integer | 否 | 设备影子控制 0:自动,当设备不在线的时候设置设备影子,设备在线时直接下发给设备 1:只实时下发,不在线报错 2:如果有设备影子只修改影子,没有的也不下发 (格式: int64) |
| `syncTimeout` | integer | 否 | 同步超时时间(单位:秒,不填默认3秒) (格式: int64) |

**请求示例**:
```json
{
  "areaID": "string",
  "areaIDPath": "string",
  "data": "string",
  "deviceNames": [
    "string"
  ],
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string"
    }
  ],
  "groupID": "string",
  "isAsync": true,
  "productID": "string",
  "shadowControl": 1,
  "syncTimeout": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": 1,
        "deviceName": "string",
        "msg": "string",
        "msgToken": "string",
        "productID": "string",
        "sysCode": 1,
        "sysMsg": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/property-control-batch-send \
  --body '{"areaID": "string", "areaIDPath": "string", "data": "string", "deviceNames": ["string"], "devices": [{"deviceName": "示例名称", "productID": "string"}], "groupID": "string", "isAsync": true, "productID": "string", "shadowControl": 1, "syncTimeout": 1}'
```

### POST `/api/v1/things/device/interact/property-control-get-one`

**说明**: 获取调用设备属性的结果

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名 |
| `msgToken` | string | 是 | 调用id |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "deviceName": "string",
  "msgToken": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": 1,
    "msg": "string",
    "msgToken": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/property-control-get-one \
  --body '{"deviceName": "string", "msgToken": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/interact/property-control-send`

**说明**: 调用设备属性

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `data` | string | 是 | 属性数据, JSON格式字符串, 注意字段需要在物模型属性里定义 |
| `deviceName` | string | 是 | 设备名 |
| `isAsync` | boolean | 否 | 是否异步操作 异步情况通过获取接口来获取 (格式: boolean) |
| `productID` | string | 是 | 产品id |
| `shadowControl` | integer | 否 | 设备影子控制 0:自动,当设备不在线的时候设置设备影子,设备在线时直接下发给设备 1:只实时下发,不在线报错 2:如果有设备影子只修改影子,没有的也不下发 3:只修改云端的值 4:只修改云端的值并记录操作日志 (格式: int64) |
| `syncTimeout` | integer | 否 | 同步超时时间(单位:秒,不填默认3秒) (格式: int64) |
| `withProfile` | object | 否 |  |
| `withProfile.code` | string | 是 |  |
| `withProfile.params` | string | 是 |  |

**请求示例**:
```json
{
  "data": "string",
  "deviceName": "string",
  "isAsync": true,
  "productID": "string",
  "shadowControl": 1,
  "syncTimeout": 1,
  "withProfile": {
    "code": "string",
    "params": "string"
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": 1,
    "msg": "string",
    "msgToken": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/property-control-send \
  --body '{"data": "string", "deviceName": "string", "isAsync": true, "productID": "string", "shadowControl": 1, "syncTimeout": 1, "withProfile": {"code": "string", "params": "string"}}'
```

### POST `/api/v1/things/device/interact/property-get-report-batch-send`

**说明**: 批量请求设备获取设备最新属性

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 控制一个区域的设备 |
| `areaIDPath` | string | 否 | 根据区域路径控制设备 |
| `dataIDs` | array[string] | 否 | 获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条 |
| `deviceNames` | array[string] | 否 | 设备名 |
| `devices` | array[DeviceCore] | 否 | 如果是不同的产品,则传这个字段,上面两个参数填了优先使用 |
| `groupID` | string | 否 | 控制一个组的设备 |
| `productID` | string | 否 | 产品id |

**请求示例**:
```json
{
  "areaID": "string",
  "areaIDPath": "string",
  "dataIDs": [
    "string"
  ],
  "deviceNames": [
    "string"
  ],
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string"
    }
  ],
  "groupID": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": 1,
        "deviceName": "string",
        "msg": "string",
        "msgToken": "string",
        "params": "string",
        "productID": "string",
        "sysCode": 1,
        "sysMsg": "string",
        "timestamp": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/property-get-report-batch-send \
  --body '{"areaID": "string", "areaIDPath": "string", "dataIDs": ["string"], "deviceNames": ["string"], "devices": [{"deviceName": "示例名称", "productID": "string"}], "groupID": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/interact/property-get-report-send`

**说明**: 请求设备获取设备最新属性

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dataIDs` | array[string] | 否 | 获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条 |
| `deviceName` | string | 是 | 设备名(必填) |
| `productID` | string | 是 | 产品id (必填) |

**请求示例**:
```json
{
  "dataIDs": [
    "string"
  ],
  "deviceName": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": 1,
    "msg": "string",
    "msgToken": "string",
    "params": "string",
    "timestamp": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/property-get-report-send \
  --body '{"dataIDs": ["string"], "deviceName": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/interact/schema-mock-gen`

**说明**: 生成物模型模拟数据

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dataIDs` | array[string] | 否 | 属性不填为返回全部, 事件和行为必须填一个 |
| `deviceName` | string | 否 | 设备名 |
| `productID` | string | 是 | 产品id |
| `type` | integer | 是 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

**请求示例**:
```json
{
  "dataIDs": [
    "string"
  ],
  "deviceName": "string",
  "productID": "string",
  "type": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "params": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/schema-mock-gen \
  --body '{"dataIDs": ["string"], "deviceName": "string", "productID": "string", "type": 1}'
```
