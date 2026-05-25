# ur-device things/device/schema

批量创建设备物模型 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/schema/batch-create` | 批量创建设备物模型 | admin |
| POST | `/api/v1/things/device/schema/batch-delete` | 批量删除设备物模型 | admin |
| POST | `/api/v1/things/device/schema/create` | 创建设备物模型 | admin |
| POST | `/api/v1/things/device/schema/get-list` | 获取设备物模型列表 | admin |
| POST | `/api/v1/things/device/schema/tsl-read` | 获取设备物模型tsl | admin |
| POST | `/api/v1/things/device/schema/update` | 更新设备物模型 | admin |

## 详细说明

### POST `/api/v1/things/device/schema/batch-create`

**说明**: 批量创建设备物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  |
| `identifiers` | array[string] | 否 | 导入的标识符列表 |
| `list` | array[DeviceSchema] | 否 | 如果没有填标识符列表取这里 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "deviceName": "string",
  "identifiers": [
    "string"
  ],
  "list": [
    {
      "affordance": "string",
      "desc": "string",
      "deviceName": "string",
      "extendConfig": "string",
      "funcGroup": 1,
      "identifier": "string",
      "isCanSceneLinkage": 1,
      "name": "示例名称",
      "order": 1,
      "productID": "string",
      "required": 1,
      "tag": 1,
      "type": 1
    }
  ],
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
ur api /api/v1/things/device/schema/batch-create \
  --body '{"deviceName": "string", "identifiers": ["string"], "list": [{"affordance": "string", "desc": "string", "deviceName": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}], "productID": "string"}'
```

### POST `/api/v1/things/device/schema/batch-delete`

**说明**: 批量删除设备物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  |
| `identifiers` | array[string] | 是 | 物模型ID列表 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "deviceName": "string",
  "identifiers": [
    "string"
  ],
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
ur api /api/v1/things/device/schema/batch-delete \
  --body '{"deviceName": "string", "identifiers": ["string"], "productID": "string"}'
```

### POST `/api/v1/things/device/schema/create`

**说明**: 创建设备物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `affordance` | string | 否 |  |
| `desc` | string | 否 | 描述 |
| `deviceName` | string | 否 |  |
| `extendConfig` | string | 否 | 拓展参数 |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifier` | string | 是 | 标识符 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 | 功能名称 |
| `order` | integer | 否 |  排序 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `required` | integer | 否 | 是否必须 1:是 2:否 (格式: int64) |
| `tag` | integer | 是 | 物模型标签 1:自定义 2:可选 3:必选  必选不可删除 (格式: int64) |
| `type` | integer | 是 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

**请求示例**:
```json
{
  "affordance": "string",
  "desc": "string",
  "deviceName": "string",
  "extendConfig": "string",
  "funcGroup": 1,
  "identifier": "string",
  "isCanSceneLinkage": 1,
  "name": "示例名称",
  "order": 1,
  "productID": "string",
  "required": 1,
  "tag": 1,
  "type": 1
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
ur api /api/v1/things/device/schema/create \
  --body '{"affordance": "string", "desc": "string", "deviceName": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}'
```

### POST `/api/v1/things/device/schema/get-list`

**说明**: 获取设备物模型列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `controlMode` | integer | 否 | 控制模式: 1: 可以群控,可以单控  2:只能单控 (格式: int64) |
| `deviceName` | string | 是 |  |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifiers` | array[string] | 否 | 过滤标识符列表 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 是 | 产品id |
| `propertyMode` | string | 否 | 属性类型可用 读写类型: 1:r(只读) 2:rw(可读可写) |
| `tag` | integer | 否 | 过滤条件: 物模型标签 1:自定义 2:可选 3:必选 (格式: int64) |
| `type` | integer | 否 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |
| `types` | array[integer] | 否 |  |
| `withProductSchema` | boolean | 否 | 同时获取产品物模型 (格式: boolean) |

**请求示例**:
```json
{
  "controlMode": 1,
  "deviceName": "string",
  "funcGroup": 1,
  "identifiers": [
    "string"
  ],
  "isCanSceneLinkage": 1,
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
  "productID": "string",
  "propertyMode": "string",
  "tag": 1,
  "type": 1,
  "types": [
    1
  ],
  "withProductSchema": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "affordance": "string",
        "desc": "string",
        "deviceName": "string",
        "extendConfig": "string",
        "funcGroup": 1,
        "identifier": "string",
        "isCanSceneLinkage": 1,
        "name": "示例名称",
        "order": 1,
        "productID": "string",
        "required": 1,
        "tag": 1,
        "type": 1
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
ur api /api/v1/things/device/schema/get-list \
  --body '{"controlMode": 1, "deviceName": "string", "funcGroup": 1, "identifiers": ["string"], "isCanSceneLinkage": 1, "name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string", "propertyMode": "string", "tag": 1, "type": 1, "types": [1], "withProductSchema": true}'
```

### POST `/api/v1/things/device/schema/tsl-read`

**说明**: 获取设备物模型tsl

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "deviceName": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "tsl": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/schema/tsl-read \
  --body '{"deviceName": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/schema/update`

**说明**: 更新设备物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `affordance` | string | 否 |  |
| `desc` | string | 否 | 描述 |
| `deviceName` | string | 否 |  |
| `extendConfig` | string | 否 | 拓展参数 |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifier` | string | 是 | 标识符 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 | 功能名称 |
| `order` | integer | 否 |  排序 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `required` | integer | 否 | 是否必须 1:是 2:否 (格式: int64) |
| `tag` | integer | 是 | 物模型标签 1:自定义 2:可选 3:必选  必选不可删除 (格式: int64) |
| `type` | integer | 是 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

**请求示例**:
```json
{
  "affordance": "string",
  "desc": "string",
  "deviceName": "string",
  "extendConfig": "string",
  "funcGroup": 1,
  "identifier": "string",
  "isCanSceneLinkage": 1,
  "name": "示例名称",
  "order": 1,
  "productID": "string",
  "required": 1,
  "tag": 1,
  "type": 1
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
ur api /api/v1/things/device/schema/update \
  --body '{"affordance": "string", "desc": "string", "deviceName": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}'
```
