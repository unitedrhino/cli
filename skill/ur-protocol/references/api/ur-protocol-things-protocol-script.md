# ur-protocol things/protocol/script

批量导出协议脚本 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/protocol/script/batch-export` | 批量导出协议脚本 | admin |
| POST | `/api/v1/things/protocol/script/batch-import` | 批量导入协议脚本 | admin |
| POST | `/api/v1/things/protocol/script/create` | 新增协议脚本 | admin |
| POST | `/api/v1/things/protocol/script/debug` | 协议脚本调试 | admin |
| POST | `/api/v1/things/protocol/script/delete` | 删除协议脚本 | admin |
| POST | `/api/v1/things/protocol/script/get-list` | 获取协议脚本列表 | admin |
| POST | `/api/v1/things/protocol/script/get-one` | 获取协议脚本详情 | admin |
| POST | `/api/v1/things/protocol/script/update` | 更新协议脚本 | admin |

## 详细说明

### POST `/api/v1/things/protocol/script/batch-export`

**说明**: 批量导出协议脚本

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[string] | 否 |  |

**请求示例**:
```json
{
  "ids": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "scriptsFile": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/script/batch-export \
  --body '{"ids": ["string"]}'
```

### POST `/api/v1/things/protocol/script/batch-import`

**说明**: 批量导入协议脚本

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `scriptsFile` | string | 是 | 协议脚本数据文件的 OSS 地址（ZIP 压缩包） |

**请求示例**:
```json
{
  "scriptsFile": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "errCount": 1,
    "ignoreCount": 1,
    "succCount": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/script/batch-import \
  --body '{"scriptsFile": "string"}'
```

### POST `/api/v1/things/protocol/script/create`

**说明**: 新增协议脚本

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间 |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `script` | string | 否 | 脚本内容 |
| `status` | integer | 否 | 状态:是否启用 (格式: int64) |
| `tenantCode` | string | 否 | 租户号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `triggerDir` | integer | 否 | 1:up 2:down (格式: int64) |
| `triggerHandle` | string | 否 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `triggerTimer` | integer | 否 | 收到前处理1:before 2:after (格式: int64) |
| `triggerType` | string | 否 |  操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "name": "string",
  "script": "string",
  "status": 1,
  "tenantCode": "string",
  "triggerDir": 1,
  "triggerHandle": "string",
  "triggerTimer": 1,
  "triggerType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/script/create \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "name": "string", "script": "string", "status": 1, "tenantCode": "string", "triggerDir": 1, "triggerHandle": "string", "triggerTimer": 1, "triggerType": "string"}'
```

### POST `/api/v1/things/protocol/script/debug`

**说明**: 协议脚本调试

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `req` | object | 是 |  |
| `req.deviceName` | string | 否 |  |
| `req.explain` | string | 否 | 内部使用的拓展字段 |
| `req.handle` | string | 否 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `req.payload` | string | 否 |  |
| `req.productID` | string | 否 |  |
| `req.protocolCode` | string | 否 | 如果有该字段则回复的时候也会带上该字段 |
| `req.timestamp` | string | 否 | 毫秒时间戳 |
| `req.type` | string | 否 | 操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |
| `resp` | object | 否 |  |
| `resp.deviceName` | string | 否 |  |
| `resp.explain` | string | 否 | 内部使用的拓展字段 |
| `resp.handle` | string | 否 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `resp.payload` | string | 否 |  |
| `resp.productID` | string | 否 |  |
| `resp.protocolCode` | string | 否 | 如果有该字段则回复的时候也会带上该字段 |
| `resp.timestamp` | string | 否 | 毫秒时间戳 |
| `resp.type` | string | 否 | 操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |
| `script` | string | 是 | 脚本内容 |
| `triggerDir` | integer | 是 | 1:up 2:down (格式: int64) |
| `triggerTimer` | integer | 是 | 收到前处理1:before 2:after (格式: int64) |

**请求示例**:
```json
{
  "req": {
    "deviceName": "string",
    "explain": "string",
    "handle": "string",
    "payload": "string",
    "productID": "string",
    "protocolCode": "string",
    "timestamp": "2026-01-01T00:00:00Z",
    "type": "string"
  },
  "resp": {
    "deviceName": "string",
    "explain": "string",
    "handle": "string",
    "payload": "string",
    "productID": "string",
    "protocolCode": "string",
    "timestamp": "2026-01-01T00:00:00Z",
    "type": "string"
  },
  "script": "string",
  "triggerDir": 1,
  "triggerTimer": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "logs": [
      "string"
    ],
    "out": {
      "deviceName": "string",
      "explain": "string",
      "handle": "string",
      "payload": "string",
      "productID": "string",
      "protocolCode": "string",
      "timestamp": "2026-01-01T00:00:00Z",
      "type": "string"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/script/debug \
  --body '{"req": {"deviceName": "string", "explain": "string", "handle": "string", "payload": "string", "productID": "string", "protocolCode": "string", "timestamp": "2026-01-01T00:00:00Z", "type": "string"}, "resp": {"deviceName": "string", "explain": "string", "handle": "string", "payload": "string", "productID": "string", "protocolCode": "string", "timestamp": "2026-01-01T00:00:00Z", "type": "string"}, "script": "string", "triggerDir": 1, "triggerTimer": 1}'
```

### POST `/api/v1/things/protocol/script/delete`

**说明**: 删除协议脚本

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
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
ur api /api/v1/things/protocol/script/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/protocol/script/get-list`

**说明**: 获取协议脚本列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 状态:是否启用 (格式: int64) |
| `triggerDir` | integer | 否 | 1:up 2:down (格式: int64) |
| `triggerHandle` | string | 否 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `triggerTimer` | integer | 否 | 收到前处理1:before 2:after (格式: int64) |
| `triggerType` | string | 否 |  操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |

**请求示例**:
```json
{
  "name": "string",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
  },
  "status": 1,
  "triggerDir": 1,
  "triggerHandle": "string",
  "triggerTimer": 1,
  "triggerType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "id": "string",
        "name": "string",
        "script": "string",
        "status": 1,
        "tenantCode": "string",
        "triggerDir": 1,
        "triggerHandle": "string",
        "triggerTimer": 1,
        "triggerType": "string"
      }
    ],
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/script/get-list \
  --body '{"name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "triggerDir": 1, "triggerHandle": "string", "triggerTimer": 1, "triggerType": "string"}'
```

### POST `/api/v1/things/protocol/script/get-one`

**说明**: 获取协议脚本详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "name": "string",
    "script": "string",
    "status": 1,
    "tenantCode": "string",
    "triggerDir": 1,
    "triggerHandle": "string",
    "triggerTimer": 1,
    "triggerType": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/script/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/protocol/script/update`

**说明**: 更新协议脚本

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间 |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `script` | string | 否 | 脚本内容 |
| `status` | integer | 否 | 状态:是否启用 (格式: int64) |
| `tenantCode` | string | 否 | 租户号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `triggerDir` | integer | 否 | 1:up 2:down (格式: int64) |
| `triggerHandle` | string | 否 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `triggerTimer` | integer | 否 | 收到前处理1:before 2:after (格式: int64) |
| `triggerType` | string | 否 |  操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "name": "string",
  "script": "string",
  "status": 1,
  "tenantCode": "string",
  "triggerDir": 1,
  "triggerHandle": "string",
  "triggerTimer": 1,
  "triggerType": "string"
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
ur api /api/v1/things/protocol/script/update \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "name": "string", "script": "string", "status": 1, "tenantCode": "string", "triggerDir": 1, "triggerHandle": "string", "triggerTimer": 1, "triggerType": "string"}'
```
