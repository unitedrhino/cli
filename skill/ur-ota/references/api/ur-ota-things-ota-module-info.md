# ur-ota things/ota/module/info

创建模块 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/ota/module/info/create` | 创建模块 | admin |
| POST | `/api/v1/things/ota/module/info/delete` | 删除模块 | admin |
| POST | `/api/v1/things/ota/module/info/get-list` | 获取模块列表 | admin |
| POST | `/api/v1/things/ota/module/info/get-one` | 查询模块详情 | admin |
| POST | `/api/v1/things/ota/module/info/update` | 更新模块 | admin |

## 详细说明

### POST `/api/v1/things/ota/module/info/create`

**说明**: 创建模块

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `productID` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "string",
  "desc": "string",
  "id": "string",
  "name": "string",
  "productID": "string"
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
ur api /api/v1/things/ota/module/info/create \
  --body '{"code": "string", "createdTime": "string", "desc": "string", "id": "string", "name": "string", "productID": "string"}'
```

### POST `/api/v1/things/ota/module/info/delete`

**说明**: 删除模块

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
ur api /api/v1/things/ota/module/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/ota/module/info/get-list`

**说明**: 获取模块列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 |  |

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
        "code": "string",
        "createdTime": "string",
        "desc": "string",
        "id": "string",
        "name": "string",
        "productID": "string"
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
ur api /api/v1/things/ota/module/info/get-list \
  --body '{"name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string"}'
```

### POST `/api/v1/things/ota/module/info/get-one`

**说明**: 查询模块详情

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
    "createdTime": "string",
    "desc": "string",
    "id": "string",
    "name": "string",
    "productID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/ota/module/info/get-one \
  --body '{"code": "string", "id": 1}'
```

### POST `/api/v1/things/ota/module/info/update`

**说明**: 更新模块

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `productID` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "string",
  "desc": "string",
  "id": "string",
  "name": "string",
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
ur api /api/v1/things/ota/module/info/update \
  --body '{"code": "string", "createdTime": "string", "desc": "string", "id": "string", "name": "string", "productID": "string"}'
```
