# ur-protocol things/protocol/info

新增自定义协议 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/protocol/info/create` | 新增自定义协议 | admin |
| POST | `/api/v1/things/protocol/info/delete` | 删除自定义协议 | admin |
| POST | `/api/v1/things/protocol/info/get-list` | 获取自定义协议信息列表 | admin |
| POST | `/api/v1/things/protocol/info/get-one` | 获取自定义协议详情 | admin |
| POST | `/api/v1/things/protocol/info/update` | 更新自定义协议 | admin |

## 详细说明

### POST `/api/v1/things/protocol/info/create`

**说明**: 新增自定义协议

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |   iThings,iThings-thingsboard,wumei,aliyun,huaweiyun,tuya |
| `configFields` | array[ProtocolConfigField] | 否 | 配置字段列表,没有可以不传 |
| `containerPort` | integer | 否 |  容器端口 (格式: int64) |
| `deployStatus` | string | 否 |  部署状态: notInstalled/downloaded/running/stopped |
| `desc` | string | 否 |  |
| `deviceFields` | array[ProtocolConfigField] | 否 | 设备级的配置字段列表,没有可以不传 |
| `endpoints` | array[string] | 否 |  |
| `envVars` | object | 否 |  环境变量 |
| `etcdKey` | string | 否 |  |
| `id` | string | 否 |  |
| `image` | string | 否 |  镜像地址 |
| `isBuiltin` | integer | 否 |  是否内置 1:是 2:否 (格式: int64) |
| `isEnableSyncDevice` | integer | 是 | 是否支持主动同步设备信息 (格式: int64) |
| `isEnableSyncProduct` | integer | 是 | 是否支持主动同步产品信息 (格式: int64) |
| `name` | string | 否 |  |
| `productFields` | array[ProtocolConfigField] | 否 | 产品级的配置字段列表,没有可以不传 |
| `transProtocol` | string | 否 |  传输协议: mqtt,tcp,udp |
| `type` | string | 否 | 协议类型: 普通设备(默认):normal  音视频:media |
| `version` | string | 否 |  镜像版本 |

**请求示例**:
```json
{
  "code": "string",
  "configFields": [
    {
      "defaultValue": "string",
      "group": "string",
      "id": "string",
      "isRequired": true,
      "key": "string",
      "label": "string",
      "mapping": {},
      "max": "string",
      "min": "string",
      "placeholder": "string",
      "sort": 1,
      "type": "string"
    }
  ],
  "containerPort": 1,
  "deployStatus": "string",
  "desc": "string",
  "deviceFields": [
    {
      "defaultValue": "string",
      "group": "string",
      "id": "string",
      "isRequired": true,
      "key": "string",
      "label": "string",
      "mapping": {},
      "max": "string",
      "min": "string",
      "placeholder": "string",
      "sort": 1,
      "type": "string"
    }
  ],
  "endpoints": [
    "string"
  ],
  "envVars": {},
  "etcdKey": "string",
  "id": "string",
  "image": "string",
  "isBuiltin": 1,
  "isEnableSyncDevice": 1,
  "isEnableSyncProduct": 1,
  "name": "string",
  "productFields": [
    {
      "defaultValue": "string",
      "group": "string",
      "id": "string",
      "isRequired": true,
      "key": "string",
      "label": "string",
      "mapping": {},
      "max": "string",
      "min": "string",
      "placeholder": "string",
      "sort": 1,
      "type": "string"
    }
  ],
  "transProtocol": "string",
  "type": "string",
  "version": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/info/create \
  --body '{"code": "string", "configFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "containerPort": 1, "deployStatus": "string", "desc": "string", "deviceFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "endpoints": ["string"], "envVars": {}, "etcdKey": "string", "id": "string", "image": "string", "isBuiltin": 1, "isEnableSyncDevice": 1, "isEnableSyncProduct": 1, "name": "string", "productFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "transProtocol": "string", "type": "string", "version": "string"}'
```

### POST `/api/v1/things/protocol/info/delete`

**说明**: 删除自定义协议

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

**请求示例**:
```json
{
  "id": 1
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
ur api /api/v1/things/protocol/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/protocol/info/get-list`

**说明**: 获取自定义协议信息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `codes` | array[string] | 否 |  |
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `transProtocol` | string | 否 |  传输协议: mqtt,tcp,udp |
| `type` | string | 否 | 协议类型: 普通设备(默认):normal  音视频:media |

**请求示例**:
```json
{
  "code": "string",
  "codes": [
    "string"
  ],
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
  "transProtocol": "string",
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": "string",
        "configFields": [
          {
            "defaultValue": "string",
            "group": "string",
            "id": "string",
            "isRequired": true,
            "key": "string",
            "label": "string",
            "mapping": {},
            "max": "string",
            "min": "string",
            "placeholder": "string",
            "sort": 1,
            "type": "string"
          }
        ],
        "containerPort": 1,
        "deployStatus": "string",
        "desc": "string",
        "deviceFields": [
          {
            "defaultValue": "string",
            "group": "string",
            "id": "string",
            "isRequired": true,
            "key": "string",
            "label": "string",
            "mapping": {},
            "max": "string",
            "min": "string",
            "placeholder": "string",
            "sort": 1,
            "type": "string"
          }
        ],
        "endpoints": [
          "string"
        ],
        "envVars": {},
        "etcdKey": "string",
        "id": "string",
        "image": "string",
        "isBuiltin": 1,
        "isEnableSyncDevice": 1,
        "isEnableSyncProduct": 1,
        "name": "string",
        "productFields": [
          {
            "defaultValue": "string",
            "group": "string",
            "id": "string",
            "isRequired": true,
            "key": "string",
            "label": "string",
            "mapping": {},
            "max": "string",
            "min": "string",
            "placeholder": "string",
            "sort": 1,
            "type": "string"
          }
        ],
        "transProtocol": "string",
        "type": "string",
        "version": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/info/get-list \
  --body '{"code": "string", "codes": ["string"], "name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "transProtocol": "string", "type": "string"}'
```

### POST `/api/v1/things/protocol/info/get-one`

**说明**: 获取自定义协议详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | integer | 否 |  id (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "configFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "containerPort": 1,
    "deployStatus": "string",
    "desc": "string",
    "deviceFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "endpoints": [
      "string"
    ],
    "envVars": {},
    "etcdKey": "string",
    "id": "string",
    "image": "string",
    "isBuiltin": 1,
    "isEnableSyncDevice": 1,
    "isEnableSyncProduct": 1,
    "name": "string",
    "productFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "transProtocol": "string",
    "type": "string",
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/info/get-one \
  --body '{"code": "string", "id": 1}'
```

### POST `/api/v1/things/protocol/info/update`

**说明**: 更新自定义协议

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |   iThings,iThings-thingsboard,wumei,aliyun,huaweiyun,tuya |
| `configFields` | array[ProtocolConfigField] | 否 | 配置字段列表,没有可以不传 |
| `containerPort` | integer | 否 |  容器端口 (格式: int64) |
| `deployStatus` | string | 否 |  部署状态: notInstalled/downloaded/running/stopped |
| `desc` | string | 否 |  |
| `deviceFields` | array[ProtocolConfigField] | 否 | 设备级的配置字段列表,没有可以不传 |
| `endpoints` | array[string] | 否 |  |
| `envVars` | object | 否 |  环境变量 |
| `etcdKey` | string | 否 |  |
| `id` | string | 否 |  |
| `image` | string | 否 |  镜像地址 |
| `isBuiltin` | integer | 否 |  是否内置 1:是 2:否 (格式: int64) |
| `isEnableSyncDevice` | integer | 是 | 是否支持主动同步设备信息 (格式: int64) |
| `isEnableSyncProduct` | integer | 是 | 是否支持主动同步产品信息 (格式: int64) |
| `name` | string | 否 |  |
| `productFields` | array[ProtocolConfigField] | 否 | 产品级的配置字段列表,没有可以不传 |
| `transProtocol` | string | 否 |  传输协议: mqtt,tcp,udp |
| `type` | string | 否 | 协议类型: 普通设备(默认):normal  音视频:media |
| `version` | string | 否 |  镜像版本 |

**请求示例**:
```json
{
  "code": "string",
  "configFields": [
    {
      "defaultValue": "string",
      "group": "string",
      "id": "string",
      "isRequired": true,
      "key": "string",
      "label": "string",
      "mapping": {},
      "max": "string",
      "min": "string",
      "placeholder": "string",
      "sort": 1,
      "type": "string"
    }
  ],
  "containerPort": 1,
  "deployStatus": "string",
  "desc": "string",
  "deviceFields": [
    {
      "defaultValue": "string",
      "group": "string",
      "id": "string",
      "isRequired": true,
      "key": "string",
      "label": "string",
      "mapping": {},
      "max": "string",
      "min": "string",
      "placeholder": "string",
      "sort": 1,
      "type": "string"
    }
  ],
  "endpoints": [
    "string"
  ],
  "envVars": {},
  "etcdKey": "string",
  "id": "string",
  "image": "string",
  "isBuiltin": 1,
  "isEnableSyncDevice": 1,
  "isEnableSyncProduct": 1,
  "name": "string",
  "productFields": [
    {
      "defaultValue": "string",
      "group": "string",
      "id": "string",
      "isRequired": true,
      "key": "string",
      "label": "string",
      "mapping": {},
      "max": "string",
      "min": "string",
      "placeholder": "string",
      "sort": 1,
      "type": "string"
    }
  ],
  "transProtocol": "string",
  "type": "string",
  "version": "string"
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
ur api /api/v1/things/protocol/info/update \
  --body '{"code": "string", "configFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "containerPort": 1, "deployStatus": "string", "desc": "string", "deviceFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "endpoints": ["string"], "envVars": {}, "etcdKey": "string", "id": "string", "image": "string", "isBuiltin": 1, "isEnableSyncDevice": 1, "isEnableSyncProduct": 1, "name": "string", "productFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "transProtocol": "string", "type": "string", "version": "string"}'
```
