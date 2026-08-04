# ur-tenant system/tenant/agreement

添加协议 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/agreement/create` | 添加协议 | admin |
| POST | `/api/v1/system/tenant/agreement/delete` | 删除协议 | admin |
| POST | `/api/v1/system/tenant/agreement/get-list` | 获取协议列表 | admin |
| POST | `/api/v1/system/tenant/agreement/get-one` | 获取协议详情 | public |
| POST | `/api/v1/system/tenant/agreement/update` | 更新协议 | admin |

## 详细说明

### POST `/api/v1/system/tenant/agreement/create`

**说明**: 添加协议

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  编号 |
| `content` | string | 否 | 协议内容(只有详情会返回) |
| `id` | string | 否 |  编号 |
| `name` | string | 否 |  协议名称 |
| `title` | string | 否 | 协议标题 |

**请求示例**:
```json
{
  "code": "string",
  "content": "string",
  "id": "string",
  "name": "示例名称",
  "title": "string"
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
ur api /api/v1/system/tenant/agreement/create \
  --body '{"code": "string", "content": "string", "id": "string", "name": "示例名称", "title": "string"}'
```

### POST `/api/v1/system/tenant/agreement/delete`

**说明**: 删除协议

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
ur api /api/v1/system/tenant/agreement/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/tenant/agreement/get-list`

**说明**: 获取协议列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
        "code": "string",
        "content": "string",
        "id": "string",
        "name": "示例名称",
        "title": "string"
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
ur api /api/v1/system/tenant/agreement/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/tenant/agreement/get-one`

**说明**: 获取协议详情

**权限**: public

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
    "content": "string",
    "id": "string",
    "name": "示例名称",
    "title": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/agreement/get-one \
  --body '{"code": "string", "id": 1}'
```

### POST `/api/v1/system/tenant/agreement/update`

**说明**: 更新协议

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  编号 |
| `content` | string | 否 | 协议内容(只有详情会返回) |
| `id` | string | 否 |  编号 |
| `name` | string | 否 |  协议名称 |
| `title` | string | 否 | 协议标题 |

**请求示例**:
```json
{
  "code": "string",
  "content": "string",
  "id": "string",
  "name": "示例名称",
  "title": "string"
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
ur api /api/v1/system/tenant/agreement/update \
  --body '{"code": "string", "content": "string", "id": "string", "name": "示例名称", "title": "string"}'
```
