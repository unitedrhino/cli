# ur-system system/hook/capability

新增Hook能力 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/hook/capability/create` | 新增Hook能力 | platform |
| POST | `/api/v1/system/hook/capability/delete` | 删除Hook能力 | platform |
| POST | `/api/v1/system/hook/capability/get-list` | 获取Hook能力列表 | platform |
| POST | `/api/v1/system/hook/capability/update` | 更新Hook能力 | platform |

## 详细说明

### POST `/api/v1/system/hook/capability/create`

**说明**: 新增Hook能力

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  能力编码 |
| `desc` | string | 否 |  描述 |
| `id` | string | 否 |  能力ID |
| `kind` | string | 否 |  能力类型 sync=同步扩展点 async=异步事件通知（说明性元数据，默认 sync） |
| `serverID` | string | 否 |  关联服务ID |
| `subCode` | string | 否 |  子编码 |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "id": "string",
  "kind": "string",
  "serverID": "string",
  "subCode": "string"
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
ur api /api/v1/system/hook/capability/create \
  --body '{"code": "string", "desc": "string", "id": "string", "kind": "string", "serverID": "string", "subCode": "string"}'
```

### POST `/api/v1/system/hook/capability/delete`

**说明**: 删除Hook能力

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
ur api /api/v1/system/hook/capability/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/hook/capability/get-list`

**说明**: 获取Hook能力列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  能力编码筛选 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `serverID` | integer | 否 |  服务ID筛选 (格式: int64) |
| `subCode` | string | 否 |  子编码筛选 |

**请求示例**:
```json
{
  "code": "string",
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
  "serverID": 1,
  "subCode": "string"
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
        "desc": "string",
        "id": "string",
        "kind": "string",
        "serverID": "string",
        "subCode": "string"
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
ur api /api/v1/system/hook/capability/get-list \
  --body '{"code": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "serverID": 1, "subCode": "string"}'
```

### POST `/api/v1/system/hook/capability/update`

**说明**: 更新Hook能力

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  能力编码 |
| `desc` | string | 否 |  描述 |
| `id` | string | 否 |  能力ID |
| `kind` | string | 否 |  能力类型 sync=同步扩展点 async=异步事件通知（说明性元数据，默认 sync） |
| `serverID` | string | 否 |  关联服务ID |
| `subCode` | string | 否 |  子编码 |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "id": "string",
  "kind": "string",
  "serverID": "string",
  "subCode": "string"
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
ur api /api/v1/system/hook/capability/update \
  --body '{"code": "string", "desc": "string", "id": "string", "kind": "string", "serverID": "string", "subCode": "string"}'
```
