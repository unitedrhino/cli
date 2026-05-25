# ur-user system/user/self/accessToken

创建访问令牌 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/self/access-token/create` | 创建访问令牌 | all |
| POST | `/api/v1/system/user/self/access-token/delete` | 删除访问令牌 | all |
| POST | `/api/v1/system/user/self/access-token/get-list` | 获取访问令牌列表 | all |
| POST | `/api/v1/system/user/self/access-token/get-one` | 获取访问令牌详情 | all |
| POST | `/api/v1/system/user/self/access-token/update` | 更新访问令牌 | all |

## 详细说明

### POST `/api/v1/system/user/self/access-token/create`

**说明**: 创建访问令牌

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  令牌描述 |
| `expTime` | string | 否 |  到期时间 |
| `tenantCodes` | array[string] | 否 |  可访问的租户列表 |

**请求示例**:
```json
{
  "desc": "string",
  "expTime": "2026-01-01T00:00:00Z",
  "tenantCodes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "accessKey": "string",
    "accessSecret": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "expTime": "2026-01-01T00:00:00Z",
    "id": "string",
    "tenantCodes": [
      "string"
    ],
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/access-token/create \
  --body '{"desc": "string", "expTime": "2026-01-01T00:00:00Z", "tenantCodes": ["string"]}'
```

### POST `/api/v1/system/user/self/access-token/delete`

**说明**: 删除访问令牌

**权限**: all

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
ur api /api/v1/system/user/self/access-token/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/user/self/access-token/get-list`

**说明**: 获取访问令牌列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessKey` | string | 否 |  按访问密钥标识过滤 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "accessKey": "string",
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
        "accessKey": "string",
        "accessSecret": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "expTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "tenantCodes": [
          "string"
        ],
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
ur api /api/v1/system/user/self/access-token/get-list \
  --body '{"accessKey": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/user/self/access-token/get-one`

**说明**: 获取访问令牌详情

**权限**: all

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
    "accessKey": "string",
    "accessSecret": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "expTime": "2026-01-01T00:00:00Z",
    "id": "string",
    "tenantCodes": [
      "string"
    ],
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/access-token/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/user/self/access-token/update`

**说明**: 更新访问令牌

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  令牌描述 |
| `expTime` | string | 否 |  到期时间 |
| `id` | string | 是 |  令牌ID |
| `tenantCodes` | array[string] | 否 |  可访问的租户列表 |

**请求示例**:
```json
{
  "desc": "string",
  "expTime": "2026-01-01T00:00:00Z",
  "id": "string",
  "tenantCodes": [
    "string"
  ]
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
ur api /api/v1/system/user/self/access-token/update \
  --body '{"desc": "string", "expTime": "2026-01-01T00:00:00Z", "id": "string", "tenantCodes": ["string"]}'
```
