# ur-system system/agreement

创建协议 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/agreement/create` | 创建协议 | platform |
| POST | `/api/v1/system/agreement/delete` | 删除协议 | platform |
| POST | `/api/v1/system/agreement/get-list` | 获取协议列表 | platform |
| POST | `/api/v1/system/agreement/get-one` | 获取协议详情 | platform |
| POST | `/api/v1/system/agreement/preview-render` | 预览渲染协议 | platform |
| POST | `/api/v1/system/agreement/update` | 更新协议 | platform |

## 详细说明

### POST `/api/v1/system/agreement/create`

**说明**: 创建协议

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `content` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `remark` | string | 否 |  |
| `status` | integer | 否 | 格式: int64 |
| `title` | string | 否 |  |
| `updatedTime` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
  "content": "string",
  "createdTime": "string",
  "id": "string",
  "name": "string",
  "remark": "string",
  "status": 1,
  "title": "string",
  "updatedTime": "string"
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
ur api /api/v1/system/agreement/create \
  --body '{"code": "string", "content": "string", "createdTime": "string", "id": "string", "name": "string", "remark": "string", "status": 1, "title": "string", "updatedTime": "string"}'
```

### POST `/api/v1/system/agreement/delete`

**说明**: 删除协议

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
ur api /api/v1/system/agreement/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/agreement/get-list`

**说明**: 获取协议列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  |
| `code` | string | 否 |  |
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int64 |
| `title` | string | 否 |  |

**请求示例**:
```json
{
  "appID": "string",
  "code": "string",
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
  "title": "string"
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
        "content": "string",
        "createdTime": "string",
        "id": "string",
        "name": "string",
        "remark": "string",
        "status": 1,
        "title": "string",
        "updatedTime": "string"
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
ur api /api/v1/system/agreement/get-list \
  --body '{"appID": "string", "code": "string", "name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "title": "string"}'
```

### POST `/api/v1/system/agreement/get-one`

**说明**: 获取协议详情

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
    "code": "string",
    "content": "string",
    "createdTime": "string",
    "id": "string",
    "name": "string",
    "remark": "string",
    "status": 1,
    "title": "string",
    "updatedTime": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/agreement/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/agreement/preview-render`

**说明**: 预览渲染协议

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  |
| `content` | string | 否 |  |
| `id` | string | 否 |  |
| `tenantCode` | string | 否 |  |

**请求示例**:
```json
{
  "appID": "string",
  "content": "string",
  "id": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "content": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/agreement/preview-render \
  --body '{"appID": "string", "content": "string", "id": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/agreement/update`

**说明**: 更新协议

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `content` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `remark` | string | 否 |  |
| `status` | integer | 否 | 格式: int64 |
| `title` | string | 否 |  |
| `updatedTime` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
  "content": "string",
  "createdTime": "string",
  "id": "string",
  "name": "string",
  "remark": "string",
  "status": 1,
  "title": "string",
  "updatedTime": "string"
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
ur api /api/v1/system/agreement/update \
  --body '{"code": "string", "content": "string", "createdTime": "string", "id": "string", "name": "string", "remark": "string", "status": 1, "title": "string", "updatedTime": "string"}'
```
