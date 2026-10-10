# ur-system system/hook/server

新增Hook服务 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/hook/server/create` | 新增Hook服务 | platform |
| POST | `/api/v1/system/hook/server/delete` | 删除Hook服务 | platform |
| POST | `/api/v1/system/hook/server/get-list` | 获取Hook服务列表 | platform |
| POST | `/api/v1/system/hook/server/get-one` | 获取Hook服务详情 | platform |
| POST | `/api/v1/system/hook/server/update` | 更新Hook服务 | platform |

## 详细说明

### POST `/api/v1/system/hook/server/create`

**说明**: 新增Hook服务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authConfig` | object | 否 |  鉴权配置，hmac 时为 {"secret":"xxx"} |
| `authType` | string | 否 |  鉴权类型 none/hmac/custom |
| `capabilities` | array[HookCapabilityInfo] | 否 |  关联的能力列表 |
| `desc` | string | 否 |  描述 |
| `endpoint` | string | 否 |  服务端点URL |
| `failPolicy` | string | 否 |  失败策略 fail/ignore |
| `headers` | object | 否 |  自定义请求头（authType=custom 时使用） |
| `id` | string | 否 |  服务ID |
| `maxRetry` | integer | 否 |  最大重试次数 (格式: int64) |
| `name` | string | 否 |  服务名称（租户内唯一） |
| `status` | integer | 否 |  状态 1:启用 2:禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码：common=全平台可用，其余=仅该租户（创建时默认 common） |
| `timeoutSec` | integer | 否 |  超时秒数 (格式: int64) |

**请求示例**:
```json
{
  "authConfig": {},
  "authType": "string",
  "capabilities": [
    {
      "code": "string",
      "desc": "string",
      "id": "string",
      "kind": "string",
      "serverID": "string",
      "subCode": "string"
    }
  ],
  "desc": "string",
  "endpoint": "string",
  "failPolicy": "string",
  "headers": {},
  "id": "string",
  "maxRetry": 1,
  "name": "示例名称",
  "status": 1,
  "tenantCode": "string",
  "timeoutSec": 1
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
ur api /api/v1/system/hook/server/create \
  --body '{"authConfig": {}, "authType": "string", "capabilities": [{"code": "string", "desc": "string", "id": "string", "kind": "string", "serverID": "string", "subCode": "string"}], "desc": "string", "endpoint": "string", "failPolicy": "string", "headers": {}, "id": "string", "maxRetry": 1, "name": "示例名称", "status": 1, "tenantCode": "string", "timeoutSec": 1}'
```

### POST `/api/v1/system/hook/server/delete`

**说明**: 删除Hook服务

**权限**: platform

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
ur api /api/v1/system/hook/server/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/hook/server/get-list`

**说明**: 获取Hook服务列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  服务名称筛选 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  状态筛选 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码筛选 |

**请求示例**:
```json
{
  "name": "示例名称",
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
        "authConfig": {},
        "authType": "string",
        "capabilities": [
          {
            "code": "string",
            "desc": "string",
            "id": "string",
            "kind": "string",
            "serverID": "string",
            "subCode": "string"
          }
        ],
        "desc": "string",
        "endpoint": "string",
        "failPolicy": "string",
        "headers": {},
        "id": "string",
        "maxRetry": 1,
        "name": "示例名称",
        "status": 1,
        "tenantCode": "string",
        "timeoutSec": 1
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
ur api /api/v1/system/hook/server/get-list \
  --body '{"name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "tenantCode": "string"}'
```

### POST `/api/v1/system/hook/server/get-one`

**说明**: 获取Hook服务详情

**权限**: platform

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
  "data": {
    "authConfig": {},
    "authType": "string",
    "capabilities": [
      {
        "code": "string",
        "desc": "string",
        "id": "string",
        "kind": "string",
        "serverID": "string",
        "subCode": "string"
      }
    ],
    "desc": "string",
    "endpoint": "string",
    "failPolicy": "string",
    "headers": {},
    "id": "string",
    "maxRetry": 1,
    "name": "示例名称",
    "status": 1,
    "tenantCode": "string",
    "timeoutSec": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/hook/server/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/hook/server/update`

**说明**: 更新Hook服务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authConfig` | object | 否 |  鉴权配置，hmac 时为 {"secret":"xxx"} |
| `authType` | string | 否 |  鉴权类型 none/hmac/custom |
| `capabilities` | array[HookCapabilityInfo] | 否 |  关联的能力列表 |
| `desc` | string | 否 |  描述 |
| `endpoint` | string | 否 |  服务端点URL |
| `failPolicy` | string | 否 |  失败策略 fail/ignore |
| `headers` | object | 否 |  自定义请求头（authType=custom 时使用） |
| `id` | string | 否 |  服务ID |
| `maxRetry` | integer | 否 |  最大重试次数 (格式: int64) |
| `name` | string | 否 |  服务名称（租户内唯一） |
| `status` | integer | 否 |  状态 1:启用 2:禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码：common=全平台可用，其余=仅该租户（创建时默认 common） |
| `timeoutSec` | integer | 否 |  超时秒数 (格式: int64) |

**请求示例**:
```json
{
  "authConfig": {},
  "authType": "string",
  "capabilities": [
    {
      "code": "string",
      "desc": "string",
      "id": "string",
      "kind": "string",
      "serverID": "string",
      "subCode": "string"
    }
  ],
  "desc": "string",
  "endpoint": "string",
  "failPolicy": "string",
  "headers": {},
  "id": "string",
  "maxRetry": 1,
  "name": "示例名称",
  "status": 1,
  "tenantCode": "string",
  "timeoutSec": 1
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
ur api /api/v1/system/hook/server/update \
  --body '{"authConfig": {}, "authType": "string", "capabilities": [{"code": "string", "desc": "string", "id": "string", "kind": "string", "serverID": "string", "subCode": "string"}], "desc": "string", "endpoint": "string", "failPolicy": "string", "headers": {}, "id": "string", "maxRetry": 1, "name": "示例名称", "status": 1, "tenantCode": "string", "timeoutSec": 1}'
```
