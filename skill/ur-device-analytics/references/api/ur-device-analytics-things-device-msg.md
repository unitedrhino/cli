# ur-device-analytics things/device/msg

获取设备异常日志 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/msg/abnormal-log/get-list` | 获取设备异常日志 | admin |
| POST | `/api/v1/things/device/msg/event-log/get-list` | 获取事件历史记录 | admin |
| POST | `/api/v1/things/device/msg/gateway-can-bind/get-list` | 获取网关可以绑定的子设备列表 | admin |
| POST | `/api/v1/things/device/msg/hub-log/get-list` | 获取云端诊断日志 | admin |
| POST | `/api/v1/things/device/msg/property-agg/by-device/get-list` | 弃用 | admin |
| POST | `/api/v1/things/device/msg/property-agg/get-list` | 弃用 | admin |
| POST | `/api/v1/things/device/msg/property-latest-agg/get-list` | 聚合属性最新值 | admin |
| POST | `/api/v1/things/device/msg/property-latest/get-list` | 获取最新属性记录 | admin |
| POST | `/api/v1/things/device/msg/property-log-agg/by-device/get-list` | 聚合属性历史记录,设备维度 | admin |
| POST | `/api/v1/things/device/msg/property-log-agg/get-list` | 聚合属性历史记录 | admin |
| POST | `/api/v1/things/device/msg/property-log-latest/get-list` | 弃用 | admin |
| POST | `/api/v1/things/device/msg/property-log/batch-get-list` | 批量获取单个id属性历史记录 | admin |
| POST | `/api/v1/things/device/msg/property-log/get-list` | 获取单个id属性历史记录 | admin |
| POST | `/api/v1/things/device/msg/sdk-log/get-list` | 获取设备sdk日志 | admin |
| POST | `/api/v1/things/device/msg/send-log/get-list` | 获取设备命令日志 | admin |
| POST | `/api/v1/things/device/msg/shadow/get-list` | 获取设备影子列表 | admin |
| POST | `/api/v1/things/device/msg/status-log/get-list` | 获取设备状态日志 | admin |

## 详细说明

### POST `/api/v1/things/device/msg/abnormal-log/get-list`

**说明**: 获取设备异常日志

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `action` | integer | 否 | 触发1 还是解除2 (格式: int64) |
| `areaID` | string | 否 | 只有管理员有权限 |
| `areaIDPath` | string | 否 | 只有管理员有权限 |
| `areaIDs` | array[string] | 否 | 只有管理员有权限 |
| `belongGroup` | object | 否 | key是group的purpose, value是里面包含的分组id 只有partitionBy 传该参数的时候才会返回 |
| `deviceName` | string | 否 | 设备名 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productCategoryCode` | string | 否 | 品类编码 |
| `productCategoryID` | string | 否 | 只有管理员有权限 |
| `productID` | string | 否 | 产品id |
| `projectID` | string | 否 | 只有管理员有权限 |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |
| `type` | string | 否 | 异常类型 |
| `user` | object | 否 |  |
| `user.avatar` | string | 否 |  用户头像 |
| `user.createdTime` | string | 否 |  创建时间 |
| `user.email` | string | 否 |  邮箱 |
| `user.lastIP` | string | 否 |  最后登录ip |
| `user.nickName` | string | 否 |  用户的昵称 |
| `user.phone` | string | 否 |  手机号 |
| `user.regIP` | string | 否 |  注册ip |
| `user.userID` | string | 否 |  用户id |
| `user.userName` | string | 否 |  用户名(唯一) |

**请求示例**:
```json
{
  "action": 1,
  "areaID": "string",
  "areaIDPath": "string",
  "areaIDs": [
    "string"
  ],
  "belongGroup": {},
  "deviceName": "string",
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
  "productCategoryCode": "string",
  "productCategoryID": "string",
  "productID": "string",
  "projectID": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z",
  "type": "string",
  "user": {
    "avatar": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "email": "string",
    "lastIP": "string",
    "nickName": "string",
    "phone": "string",
    "regIP": "string",
    "userID": "string",
    "userName": "string"
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "action": 1,
        "deviceName": "string",
        "productID": "string",
        "reason": "string",
        "timestamp": "2026-01-01T00:00:00Z",
        "traceID": "string",
        "type": "string"
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
ur api /api/v1/things/device/msg/abnormal-log/get-list \
  --body '{"action": 1, "areaID": "string", "areaIDPath": "string", "areaIDs": ["string"], "belongGroup": {}, "deviceName": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productCategoryCode": "string", "productCategoryID": "string", "productID": "string", "projectID": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z", "type": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}}'
```

### POST `/api/v1/things/device/msg/event-log/get-list`

**说明**: 获取事件历史记录

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dataID` | string | 否 | 获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条,如果没有获取到的不会返回值 |
| `deviceNames` | array[string] | 否 | 设备名(不填获取产品下所有设备) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |
| `types` | array[string] | 否 | 类型 事件类型: 信息:info  告警alert  故障:fault |

**请求示例**:
```json
{
  "dataID": "string",
  "deviceNames": [
    "string"
  ],
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
  "productID": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z",
  "types": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "dataID": "string",
        "params": "string",
        "timestamp": "2026-01-01T00:00:00Z",
        "type": "string"
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
ur api /api/v1/things/device/msg/event-log/get-list \
  --body '{"dataID": "string", "deviceNames": ["string"], "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z", "types": ["string"]}'
```

### POST `/api/v1/things/device/msg/gateway-can-bind/get-list`

**说明**: 获取网关可以绑定的子设备列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `gateway` | object | 是 |  |
| `gateway.deviceName` | string | 是 | 设备名称 |
| `gateway.productID` | string | 是 | 产品ID |
| `gateway.productName` | string | 否 |  |

**请求示例**:
```json
{
  "gateway": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "subDevices": [
      {
        "deviceName": "示例名称",
        "productID": "string",
        "productName": "string"
      }
    ],
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/msg/gateway-can-bind/get-list \
  --body '{"gateway": {"deviceName": "示例名称", "productID": "string", "productName": "string"}}'
```

### POST `/api/v1/things/device/msg/hub-log/get-list`

**说明**: 获取云端诊断日志

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actions` | array[string] | 否 | 过滤操作类型 connected:上线 disconnected:下线  property:属性 event:事件 action:操作 thing:物模型提交的操作为匹配的日志 |
| `content` | string | 否 | 过滤内容 |
| `deviceName` | string | 否 | 设备名 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |
| `requestID` | string | 否 | 过滤请求ID |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |
| `topics` | array[string] | 否 | 过滤主题 |

**请求示例**:
```json
{
  "actions": [
    "string"
  ],
  "content": "string",
  "deviceName": "string",
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
  "productID": "string",
  "requestID": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z",
  "topics": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "action": "string",
        "content": "string",
        "requestID": "string",
        "respPayload": "string",
        "resultCode": "string",
        "timestamp": "2026-01-01T00:00:00Z",
        "topic": "string",
        "traceID": "string"
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
ur api /api/v1/things/device/msg/hub-log/get-list \
  --body '{"actions": ["string"], "content": "string", "deviceName": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string", "requestID": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z", "topics": ["string"]}'
```

### POST `/api/v1/things/device/msg/property-agg/by-device/get-list`

**说明**: 弃用

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `aggs` | array[DeviceMsgPropertyAggByDevice] | 是 | 聚合对象 |
| `interval` | integer | 否 | 间隔 如果这个值不为零值 则时间的开始和结束必须有效及聚合函数不应该为空 (格式: int64) |
| `intervalUnit` | string | 否 | 间隔单位 a (毫秒,默认), d (天), h (小时), m (分钟), n (月), s (秒), u (微秒), w (周), y (年)  则时间的开始和结束必须有效及聚合函数不应该为空 |
| `partitionBy` | string | 否 | 切分数据,可以填写deviceName |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |

**请求示例**:
```json
{
  "aggs": [
    {
      "argFuncs": [
        "string"
      ],
      "dataID": "string",
      "device": {
        "deviceName": "示例名称",
        "productID": "string",
        "productName": "string"
      },
      "fill": "string",
      "noFirstTs": true
    }
  ],
  "interval": 1,
  "intervalUnit": "2026-01-01T00:00:00Z",
  "partitionBy": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "areaIDPath": "string",
        "areaIDs": [
          "string"
        ],
        "belongGroup": {},
        "deviceName": "示例名称",
        "productID": "string",
        "projectID": "string",
        "tenantCode": "string",
        "values": [
          {
            "dataID": "string",
            "dataName": "示例名称",
            "timeWindow": "2026-01-01T00:00:00Z",
            "values": {}
          }
        ]
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/msg/property-agg/by-device/get-list \
  --body '{"aggs": [{"argFuncs": ["string"], "dataID": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "fill": "string", "noFirstTs": true}], "interval": 1, "intervalUnit": "2026-01-01T00:00:00Z", "partitionBy": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/device/msg/property-agg/get-list`

**说明**: 弃用

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `aggs` | array[DeviceMsgPropertyAgg] | 是 | 聚合对象 |
| `areaID` | string | 否 | 只有管理员有权限 |
| `areaIDPath` | string | 否 | 只有管理员有权限 |
| `areaIDs` | array[string] | 否 | 只有管理员有权限 |
| `belongGroup` | object | 否 | key是group的purpose, value是里面包含的分组id 只有partitionBy 传该参数的时候才会返回 |
| `deviceName` | string | 否 | 设备名(不填获取产品下所有设备) |
| `deviceNames` | array[string] | 否 | 设备名(不填获取产品下所有设备,只有管理员有权限) |
| `interval` | integer | 是 | 间隔 如果这个值不为零值 则时间的开始和结束必须有效及聚合函数不应该为空 (格式: int64) |
| `intervalUnit` | string | 否 | 间隔单位 a (毫秒,默认), d (天), h (小时), m (分钟), n (月), s (秒), u (微秒), w (周), y (年)  则时间的开始和结束必须有效及聚合函数不应该为空 |
| `partitionBy` | string | 否 | 切分数据,可以填写deviceName |
| `productCategoryCode` | string | 否 | 品类编码 |
| `productCategoryID` | string | 否 | 产品品类id,通用物模型有效 |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |
| `projectID` | string | 否 | 只有管理员有权限 |
| `tenantCode` | string | 否 | 只有管理员有权限 |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |

**请求示例**:
```json
{
  "aggs": [
    {
      "argFuncs": [
        "string"
      ],
      "dataID": "string",
      "fill": "string",
      "noFirstTs": true
    }
  ],
  "areaID": "string",
  "areaIDPath": "string",
  "areaIDs": [
    "string"
  ],
  "belongGroup": {},
  "deviceName": "string",
  "deviceNames": [
    "string"
  ],
  "interval": 1,
  "intervalUnit": "2026-01-01T00:00:00Z",
  "partitionBy": "string",
  "productCategoryCode": "string",
  "productCategoryID": "string",
  "productID": "string",
  "projectID": "string",
  "tenantCode": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "areaIDPath": "string",
        "areaIDs": [
          "string"
        ],
        "belongGroup": {},
        "deviceName": "示例名称",
        "productID": "string",
        "projectID": "string",
        "tenantCode": "string",
        "values": [
          {
            "dataID": "string",
            "dataName": "示例名称",
            "timeWindow": "2026-01-01T00:00:00Z",
            "values": {}
          }
        ]
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/msg/property-agg/get-list \
  --body '{"aggs": [{"argFuncs": ["string"], "dataID": "string", "fill": "string", "noFirstTs": true}], "areaID": "string", "areaIDPath": "string", "areaIDs": ["string"], "belongGroup": {}, "deviceName": "string", "deviceNames": ["string"], "interval": 1, "intervalUnit": "2026-01-01T00:00:00Z", "partitionBy": "string", "productCategoryCode": "string", "productCategoryID": "string", "productID": "string", "projectID": "string", "tenantCode": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/device/msg/property-latest-agg/get-list`

**说明**: 聚合属性最新值

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `aggs` | array[DeviceMsgPropertyAgg] | 是 | 聚合对象 |
| `areaID` | string | 否 | 只有管理员有权限 |
| `areaIDPath` | string | 否 | 只有管理员有权限 |
| `areaIDs` | array[string] | 否 | 只有管理员有权限 |
| `belongGroup` | object | 否 | key是group的purpose, value是里面包含的分组id 只有partitionBy 传该参数的时候才会返回 |
| `deviceName` | string | 否 | 设备名(不填获取产品下所有设备) |
| `deviceNames` | array[string] | 否 | 设备名(不填获取产品下所有设备,只有管理员有权限) |
| `partitionBy` | string | 否 | 切分数据,可以填写deviceName |
| `productCategoryCode` | string | 否 | 品类编码 |
| `productCategoryID` | string | 否 | 产品品类id,通用物模型有效 |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |
| `projectID` | string | 否 | 只有管理员有权限 |
| `tenantCode` | string | 否 | 只有管理员有权限 |

**请求示例**:
```json
{
  "aggs": [
    {
      "argFuncs": [
        "string"
      ],
      "dataID": "string",
      "fill": "string",
      "noFirstTs": true
    }
  ],
  "areaID": "string",
  "areaIDPath": "string",
  "areaIDs": [
    "string"
  ],
  "belongGroup": {},
  "deviceName": "string",
  "deviceNames": [
    "string"
  ],
  "partitionBy": "string",
  "productCategoryCode": "string",
  "productCategoryID": "string",
  "productID": "string",
  "projectID": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "areaIDPath": "string",
        "areaIDs": [
          "string"
        ],
        "belongGroup": {},
        "deviceName": "示例名称",
        "productID": "string",
        "projectID": "string",
        "tenantCode": "string",
        "values": [
          {
            "dataID": "string",
            "dataName": "示例名称",
            "values": {}
          }
        ]
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/msg/property-latest-agg/get-list \
  --body '{"aggs": [{"argFuncs": ["string"], "dataID": "string", "fill": "string", "noFirstTs": true}], "areaID": "string", "areaIDPath": "string", "areaIDs": ["string"], "belongGroup": {}, "deviceName": "string", "deviceNames": ["string"], "partitionBy": "string", "productCategoryCode": "string", "productCategoryID": "string", "productID": "string", "projectID": "string", "tenantCode": "string"}'
```

### POST `/api/v1/things/device/msg/property-latest/get-list`

**说明**: 获取最新属性记录

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dataIDs` | array[string] | 否 | 获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条,如果没有获取到的不会返回值 |
| `deviceName` | string | 否 | 设备名 |
| `ignoreEmpty` | boolean | 否 |  设备属性是否忽略空值,空值是否返回 (格式: boolean) |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |

**请求示例**:
```json
{
  "dataIDs": [
    "string"
  ],
  "deviceName": "string",
  "ignoreEmpty": true,
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
        "areaID": "string",
        "areaIDPath": "string",
        "areaIDs": [
          "string"
        ],
        "belongGroup": {},
        "dataID": "string",
        "dataName": "示例名称",
        "deviceName": "示例名称",
        "projectID": "string",
        "tenantCode": "string",
        "timestamp": "2026-01-01T00:00:00Z",
        "value": "string"
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
ur api /api/v1/things/device/msg/property-latest/get-list \
  --body '{"dataIDs": ["string"], "deviceName": "string", "ignoreEmpty": true, "productID": "string"}'
```

### POST `/api/v1/things/device/msg/property-log-agg/by-device/get-list`

**说明**: 聚合属性历史记录,设备维度

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `aggs` | array[DeviceMsgPropertyAggByDevice] | 是 | 聚合对象 |
| `interval` | integer | 否 | 间隔 如果这个值不为零值 则时间的开始和结束必须有效及聚合函数不应该为空 (格式: int64) |
| `intervalUnit` | string | 否 | 间隔单位 a (毫秒,默认), d (天), h (小时), m (分钟), n (月), s (秒), u (微秒), w (周), y (年)  则时间的开始和结束必须有效及聚合函数不应该为空 |
| `partitionBy` | string | 否 | 切分数据,可以填写deviceName |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |

**请求示例**:
```json
{
  "aggs": [
    {
      "argFuncs": [
        "string"
      ],
      "dataID": "string",
      "device": {
        "deviceName": "示例名称",
        "productID": "string",
        "productName": "string"
      },
      "fill": "string",
      "noFirstTs": true
    }
  ],
  "interval": 1,
  "intervalUnit": "2026-01-01T00:00:00Z",
  "partitionBy": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "areaIDPath": "string",
        "areaIDs": [
          "string"
        ],
        "belongGroup": {},
        "deviceName": "示例名称",
        "productID": "string",
        "projectID": "string",
        "tenantCode": "string",
        "values": [
          {
            "dataID": "string",
            "dataName": "示例名称",
            "timeWindow": "2026-01-01T00:00:00Z",
            "values": {}
          }
        ]
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/msg/property-log-agg/by-device/get-list \
  --body '{"aggs": [{"argFuncs": ["string"], "dataID": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "fill": "string", "noFirstTs": true}], "interval": 1, "intervalUnit": "2026-01-01T00:00:00Z", "partitionBy": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/device/msg/property-log-agg/get-list`

**说明**: 聚合属性历史记录

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `aggs` | array[DeviceMsgPropertyAgg] | 是 | 聚合对象 |
| `areaID` | string | 否 | 只有管理员有权限 |
| `areaIDPath` | string | 否 | 只有管理员有权限 |
| `areaIDs` | array[string] | 否 | 只有管理员有权限 |
| `belongGroup` | object | 否 | key是group的purpose, value是里面包含的分组id 只有partitionBy 传该参数的时候才会返回 |
| `deviceName` | string | 否 | 设备名(不填获取产品下所有设备) |
| `deviceNames` | array[string] | 否 | 设备名(不填获取产品下所有设备,只有管理员有权限) |
| `interval` | integer | 是 | 间隔 如果这个值不为零值 则时间的开始和结束必须有效及聚合函数不应该为空 (格式: int64) |
| `intervalUnit` | string | 否 | 间隔单位 a (毫秒,默认), d (天), h (小时), m (分钟), n (月), s (秒), u (微秒), w (周), y (年)  则时间的开始和结束必须有效及聚合函数不应该为空 |
| `partitionBy` | string | 否 | 切分数据,可以填写deviceName |
| `productCategoryCode` | string | 否 | 品类编码 |
| `productCategoryID` | string | 否 | 产品品类id,通用物模型有效 |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |
| `projectID` | string | 否 | 只有管理员有权限 |
| `tenantCode` | string | 否 | 只有管理员有权限 |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |

**请求示例**:
```json
{
  "aggs": [
    {
      "argFuncs": [
        "string"
      ],
      "dataID": "string",
      "fill": "string",
      "noFirstTs": true
    }
  ],
  "areaID": "string",
  "areaIDPath": "string",
  "areaIDs": [
    "string"
  ],
  "belongGroup": {},
  "deviceName": "string",
  "deviceNames": [
    "string"
  ],
  "interval": 1,
  "intervalUnit": "2026-01-01T00:00:00Z",
  "partitionBy": "string",
  "productCategoryCode": "string",
  "productCategoryID": "string",
  "productID": "string",
  "projectID": "string",
  "tenantCode": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "areaIDPath": "string",
        "areaIDs": [
          "string"
        ],
        "belongGroup": {},
        "deviceName": "示例名称",
        "productID": "string",
        "projectID": "string",
        "tenantCode": "string",
        "values": [
          {
            "dataID": "string",
            "dataName": "示例名称",
            "timeWindow": "2026-01-01T00:00:00Z",
            "values": {}
          }
        ]
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/msg/property-log-agg/get-list \
  --body '{"aggs": [{"argFuncs": ["string"], "dataID": "string", "fill": "string", "noFirstTs": true}], "areaID": "string", "areaIDPath": "string", "areaIDs": ["string"], "belongGroup": {}, "deviceName": "string", "deviceNames": ["string"], "interval": 1, "intervalUnit": "2026-01-01T00:00:00Z", "partitionBy": "string", "productCategoryCode": "string", "productCategoryID": "string", "productID": "string", "projectID": "string", "tenantCode": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/device/msg/property-log-latest/get-list`

**说明**: 弃用

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dataIDs` | array[string] | 否 | 获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条,如果没有获取到的不会返回值 |
| `deviceName` | string | 否 | 设备名 |
| `ignoreEmpty` | boolean | 否 |  设备属性是否忽略空值,空值是否返回 (格式: boolean) |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |

**请求示例**:
```json
{
  "dataIDs": [
    "string"
  ],
  "deviceName": "string",
  "ignoreEmpty": true,
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
        "areaID": "string",
        "areaIDPath": "string",
        "areaIDs": [
          "string"
        ],
        "belongGroup": {},
        "dataID": "string",
        "dataName": "示例名称",
        "deviceName": "示例名称",
        "projectID": "string",
        "tenantCode": "string",
        "timestamp": "2026-01-01T00:00:00Z",
        "value": "string"
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
ur api /api/v1/things/device/msg/property-log-latest/get-list \
  --body '{"dataIDs": ["string"], "deviceName": "string", "ignoreEmpty": true, "productID": "string"}'
```

### POST `/api/v1/things/device/msg/property-log/batch-get-list`

**说明**: 批量获取单个id属性历史记录

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `reqs` | array[DeviceMsgPropertyLogGetListReq] | 是 |  |

**请求示例**:
```json
{
  "reqs": [
    {
      "areaID": "string",
      "areaIDPath": "string",
      "areaIDs": [
        "string"
      ],
      "argFunc": "2026-01-01T00:00:00Z",
      "belongGroup": {},
      "dataID": "string",
      "deviceName": "string",
      "deviceNames": [
        "string"
      ],
      "fill": "string",
      "interval": 1,
      "intervalUnit": "2026-01-01T00:00:00Z",
      "noFirstTs": true,
      "order": 1,
      "page": {
        "orders": [
          {
            "field": "...",
            "sort": "..."
          }
        ],
        "page": 1,
        "size": 1
      },
      "partitionBy": "string",
      "productCategoryCode": "string",
      "productCategoryID": "string",
      "productID": "string",
      "projectID": "string",
      "tenantCode": "string",
      "timeEnd": "2026-01-01T00:00:00Z",
      "timeStart": "2026-01-01T00:00:00Z"
    }
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "lists": [
      [
        null
      ]
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/msg/property-log/batch-get-list \
  --body '{"reqs": [{"areaID": "string", "areaIDPath": "string", "areaIDs": ["string"], "argFunc": "2026-01-01T00:00:00Z", "belongGroup": {}, "dataID": "string", "deviceName": "string", "deviceNames": ["string"], "fill": "string", "interval": 1, "intervalUnit": "2026-01-01T00:00:00Z", "noFirstTs": true, "order": 1, "page": {"orders": [{"field": "...", "sort": "..."}], "page": 1, "size": 1}, "partitionBy": "string", "productCategoryCode": "string", "productCategoryID": "string", "productID": "string", "projectID": "string", "tenantCode": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}]}'
```

### POST `/api/v1/things/device/msg/property-log/get-list`

**说明**: 获取单个id属性历史记录

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 只有管理员有权限 |
| `areaIDPath` | string | 否 | 只有管理员有权限 |
| `areaIDs` | array[string] | 否 | 只有管理员有权限 |
| `argFunc` | string | 否 | 聚合函数 avg:平均值 first:第一个参数 last:最后一个参数 count:总数 twa: 时间加权平均函数 参考: https:docs.taosdata.com/reference/taos-sql/function/#apercentile |
| `belongGroup` | object | 否 | key是group的purpose, value是里面包含的分组id 只有partitionBy 传该参数的时候才会返回 |
| `dataID` | string | 否 | 获取的具体标识符的数据 |
| `deviceName` | string | 否 | 设备名(不填获取产品下所有设备) |
| `deviceNames` | array[string] | 否 | 设备名(不填获取产品下所有设备,只有管理员有权限) |
| `fill` | string | 否 | 指定窗口区间数据缺失的情况下的填充模式 参考: https:docs.taosdata.com/reference/taos-sql/distinguished/#fill-%E5%AD%90%E5%8F%A5 |
| `interval` | integer | 否 | 间隔 如果这个值不为零值 则时间的开始和结束必须有效及聚合函数不应该为空 (格式: int64) |
| `intervalUnit` | string | 否 | 间隔单位 a (毫秒,默认), d (天), h (小时), m (分钟), n (月), s (秒), u (微秒), w (周), y (年)  则时间的开始和结束必须有效及聚合函数不应该为空 |
| `noFirstTs` | boolean | 否 | 时间戳填充不填充最早的值,聚合模式使用 (格式: boolean) |
| `order` | integer | 否 | 时间排序 1:aes(默认,从久到近排序) 2:desc(时间从近到久排序) (格式: int64) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `partitionBy` | string | 否 | 切分数据,可以填写deviceName |
| `productCategoryCode` | string | 否 | 品类编码 |
| `productCategoryID` | string | 否 | 产品品类id,通用物模型有效 |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |
| `projectID` | string | 否 | 只有管理员有权限 |
| `tenantCode` | string | 否 | 只有管理员有权限 |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |

**请求示例**:
```json
{
  "areaID": "string",
  "areaIDPath": "string",
  "areaIDs": [
    "string"
  ],
  "argFunc": "2026-01-01T00:00:00Z",
  "belongGroup": {},
  "dataID": "string",
  "deviceName": "string",
  "deviceNames": [
    "string"
  ],
  "fill": "string",
  "interval": 1,
  "intervalUnit": "2026-01-01T00:00:00Z",
  "noFirstTs": true,
  "order": 1,
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
  "partitionBy": "string",
  "productCategoryCode": "string",
  "productCategoryID": "string",
  "productID": "string",
  "projectID": "string",
  "tenantCode": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "areaIDPath": "string",
        "areaIDs": [
          "string"
        ],
        "belongGroup": {},
        "dataID": "string",
        "dataName": "示例名称",
        "deviceName": "示例名称",
        "projectID": "string",
        "tenantCode": "string",
        "timestamp": "2026-01-01T00:00:00Z",
        "value": "string"
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
ur api /api/v1/things/device/msg/property-log/get-list \
  --body '{"areaID": "string", "areaIDPath": "string", "areaIDs": ["string"], "argFunc": "2026-01-01T00:00:00Z", "belongGroup": {}, "dataID": "string", "deviceName": "string", "deviceNames": ["string"], "fill": "string", "interval": 1, "intervalUnit": "2026-01-01T00:00:00Z", "noFirstTs": true, "order": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "partitionBy": "string", "productCategoryCode": "string", "productCategoryID": "string", "productID": "string", "projectID": "string", "tenantCode": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/device/msg/sdk-log/get-list`

**说明**: 获取设备sdk日志

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 否 | 设备名 |
| `logLevel` | integer | 否 | 等级 (格式: int32) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |

**请求示例**:
```json
{
  "deviceName": "string",
  "logLevel": 1,
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
  "productID": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "content": "string",
        "loglevel": 1,
        "timestamp": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/device/msg/sdk-log/get-list \
  --body '{"deviceName": "string", "logLevel": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/device/msg/send-log/get-list`

**说明**: 获取设备命令日志

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actions` | array[string] | 否 |  |
| `areaID` | string | 否 | 只有管理员有权限 |
| `areaIDPath` | string | 否 | 只有管理员有权限 |
| `areaIDs` | array[string] | 否 | 只有管理员有权限 |
| `dataID` | string | 否 | 模糊查询 |
| `dataIDs` | array[string] | 否 | 精确查询 |
| `deviceName` | string | 否 | 设备名 |
| `groupIDPaths` | array[string] | 否 | 只有管理员有权限 |
| `groupIDs` | array[string] | 否 | 只有管理员有权限 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productCategoryCode` | string | 否 | 品类编码 |
| `productCategoryID` | string | 否 | 只有管理员有权限 |
| `productID` | string | 否 | 产品id |
| `projectID` | string | 否 | 只有管理员有权限 |
| `resultCode` | integer | 否 | 请求结果状态,200为成功 (格式: int64) |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |
| `userID` | string | 否 |  |
| `withUser` | boolean | 否 | 同时返回用户信息 (格式: boolean) |

**请求示例**:
```json
{
  "actions": [
    "string"
  ],
  "areaID": "string",
  "areaIDPath": "string",
  "areaIDs": [
    "string"
  ],
  "dataID": "string",
  "dataIDs": [
    "string"
  ],
  "deviceName": "string",
  "groupIDPaths": [
    "string"
  ],
  "groupIDs": [
    "string"
  ],
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
  "productCategoryCode": "string",
  "productCategoryID": "string",
  "productID": "string",
  "projectID": "string",
  "resultCode": 1,
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z",
  "userID": "string",
  "withUser": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "account": "string",
        "action": "string",
        "content": "string",
        "dataID": "string",
        "deviceName": "string",
        "productID": "string",
        "resultCode": 1,
        "timestamp": "2026-01-01T00:00:00Z",
        "traceID": "string",
        "user": {
          "avatar": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "email": "string",
          "lastIP": "string",
          "nickName": "string",
          "phone": "string",
          "regIP": "string",
          "userID": "string",
          "userName": "string"
        },
        "userID": "string"
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
ur api /api/v1/things/device/msg/send-log/get-list \
  --body '{"actions": ["string"], "areaID": "string", "areaIDPath": "string", "areaIDs": ["string"], "dataID": "string", "dataIDs": ["string"], "deviceName": "string", "groupIDPaths": ["string"], "groupIDs": ["string"], "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productCategoryCode": "string", "productCategoryID": "string", "productID": "string", "projectID": "string", "resultCode": 1, "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z", "userID": "string", "withUser": true}'
```

### POST `/api/v1/things/device/msg/shadow/get-list`

**说明**: 获取设备影子列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dataIDs` | array[string] | 否 | 获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条,如果没有获取到的不会返回值 |
| `deviceName` | string | 否 | 设备名 |
| `ignoreEmpty` | boolean | 否 |  设备属性是否忽略空值,空值是否返回 (格式: boolean) |
| `productID` | string | 否 | 产品id 获取产品id下的所有设备信息 |

**请求示例**:
```json
{
  "dataIDs": [
    "string"
  ],
  "deviceName": "string",
  "ignoreEmpty": true,
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
        "dataID": "string",
        "updatedDeviceTime": 1,
        "value": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/msg/shadow/get-list \
  --body '{"dataIDs": ["string"], "deviceName": "string", "ignoreEmpty": true, "productID": "string"}'
```

### POST `/api/v1/things/device/msg/status-log/get-list`

**说明**: 获取设备状态日志

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 否 | 设备名 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 | 产品id |
| `status` | integer | 否 | 在线状态 (格式: int64) |
| `timeEnd` | string | 否 | 时间的结束(毫秒时间戳) |
| `timeStart` | string | 否 | 获取时间的开始(毫秒时间戳) |

**请求示例**:
```json
{
  "deviceName": "string",
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
  "productID": "string",
  "status": 1,
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "deviceName": "string",
        "productID": "string",
        "status": 1,
        "timestamp": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/device/msg/status-log/get-list \
  --body '{"deviceName": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string", "status": 1, "timeEnd": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```
