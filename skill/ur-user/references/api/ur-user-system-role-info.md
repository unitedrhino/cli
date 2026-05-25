# ur-user system/role/info

添加角色 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/role/info/create` | 添加角色 | admin |
| POST | `/api/v1/system/role/info/delete` | 删除角色 | admin |
| POST | `/api/v1/system/role/info/get-list` | 获取角色列表 | admin |
| POST | `/api/v1/system/role/info/update` | 更新角色 | admin |

## 详细说明

### POST `/api/v1/system/role/info/create`

**说明**: 添加角色

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 | 角色编码 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  备注 |
| `id` | string | 否 |  编号 |
| `isFullAuth` | integer | 否 |  是否完全授权（1:是，2:否） (格式: int64) |
| `name` | string | 否 |  角色名称 |
| `status` | integer | 否 |  角色状态 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "isFullAuth": 1,
  "name": "示例名称",
  "status": 1
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
ur api /api/v1/system/role/info/create \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "isFullAuth": 1, "name": "示例名称", "status": 1}'
```

### POST `/api/v1/system/role/info/delete`

**说明**: 删除角色

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
ur api /api/v1/system/role/info/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/role/info/get-list`

**说明**: 获取角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `codes` | array[string] | 否 | 角色编码过滤 |
| `ids` | array[string] | 否 |  |
| `name` | string | 否 | 按名称查找角色 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 按状态查找角色 (格式: int64) |

**请求示例**:
```json
{
  "codes": [
    "string"
  ],
  "ids": [
    "string"
  ],
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
  "status": 1
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
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "id": "string",
        "isFullAuth": 1,
        "name": "示例名称",
        "status": 1
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
ur api /api/v1/system/role/info/get-list \
  --body '{"codes": ["string"], "ids": ["string"], "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1}'
```

### POST `/api/v1/system/role/info/update`

**说明**: 更新角色

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 | 角色编码 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  备注 |
| `id` | string | 否 |  编号 |
| `isFullAuth` | integer | 否 |  是否完全授权（1:是，2:否） (格式: int64) |
| `name` | string | 否 |  角色名称 |
| `status` | integer | 否 |  角色状态 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "isFullAuth": 1,
  "name": "示例名称",
  "status": 1
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
ur api /api/v1/system/role/info/update \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "isFullAuth": 1, "name": "示例名称", "status": 1}'
```
