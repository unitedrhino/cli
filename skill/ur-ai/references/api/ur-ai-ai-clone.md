# ur-ai ai/clone

创建分身 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/clone/create` | 创建分身 | admin |
| POST | `/api/v1/ai/clone/delete` | 删除分身 | admin |
| POST | `/api/v1/ai/clone/get-list` | 获取分身列表 | admin |
| POST | `/api/v1/ai/clone/get-one` | 获取分身详情 | admin |
| POST | `/api/v1/ai/clone/update` | 更新分身 | admin |

## 详细说明

### POST `/api/v1/ai/clone/create`

**说明**: 创建分身

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 是 |  关联Agent ID (格式: int64) |
| `cloneGroupID` | integer | 否 |  所属CloneGroup ID (格式: int64) |
| `code` | string | 否 |  分身编码（全局唯一，可为空） |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  分身ID (格式: int64) |
| `metadata` | string | 否 |  扩展元数据（JSON字符串） |
| `name` | string | 是 |  分身名称 |
| `personaPrompt` | string | 否 |  人设补充提示词 |
| `status` | integer | 否 |  状态 1-启用 2-禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "cloneGroupID": 1,
  "code": "string",
  "createdTime": 1,
  "desc": "string",
  "id": 1,
  "metadata": "string",
  "name": "示例名称",
  "personaPrompt": "string",
  "status": 1,
  "tenantCode": "string",
  "updatedTime": 1
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
ur api /api/v1/ai/clone/create \
  --body '{"agentID": 1, "cloneGroupID": 1, "code": "string", "createdTime": 1, "desc": "string", "id": 1, "metadata": "string", "name": "示例名称", "personaPrompt": "string", "status": 1, "tenantCode": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/clone/delete`

**说明**: 删除分身

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
ur api /api/v1/ai/clone/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone/get-list`

**说明**: 获取分身列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 |  按Agent ID过滤 (格式: int64) |
| `cloneGroupID` | integer | 否 |  按CloneGroup ID过滤 (格式: int64) |
| `code` | string | 否 |  按code精确查询 |
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  状态过滤 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码过滤 |

**请求示例**:
```json
{
  "agentID": 1,
  "cloneGroupID": 1,
  "code": "string",
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
        "agentID": 1,
        "cloneGroupID": 1,
        "code": "string",
        "createdTime": 1,
        "desc": "string",
        "id": 1,
        "metadata": "string",
        "name": "示例名称",
        "personaPrompt": "string",
        "status": 1,
        "tenantCode": "string",
        "updatedTime": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/get-list \
  --body '{"agentID": 1, "cloneGroupID": 1, "code": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/clone/get-one`

**说明**: 获取分身详情

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
  "data": {
    "agentID": 1,
    "cloneGroupID": 1,
    "code": "string",
    "createdTime": 1,
    "desc": "string",
    "id": 1,
    "metadata": "string",
    "name": "示例名称",
    "personaPrompt": "string",
    "status": 1,
    "tenantCode": "string",
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone/update`

**说明**: 更新分身

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 是 |  关联Agent ID (格式: int64) |
| `cloneGroupID` | integer | 否 |  所属CloneGroup ID (格式: int64) |
| `code` | string | 否 |  分身编码（全局唯一，可为空） |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  分身ID (格式: int64) |
| `metadata` | string | 否 |  扩展元数据（JSON字符串） |
| `name` | string | 是 |  分身名称 |
| `personaPrompt` | string | 否 |  人设补充提示词 |
| `status` | integer | 否 |  状态 1-启用 2-禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "cloneGroupID": 1,
  "code": "string",
  "createdTime": 1,
  "desc": "string",
  "id": 1,
  "metadata": "string",
  "name": "示例名称",
  "personaPrompt": "string",
  "status": 1,
  "tenantCode": "string",
  "updatedTime": 1
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
ur api /api/v1/ai/clone/update \
  --body '{"agentID": 1, "cloneGroupID": 1, "code": "string", "createdTime": 1, "desc": "string", "id": 1, "metadata": "string", "name": "示例名称", "personaPrompt": "string", "status": 1, "tenantCode": "string", "updatedTime": 1}'
```
