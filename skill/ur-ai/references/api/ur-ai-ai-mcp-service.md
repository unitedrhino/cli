# ur-ai ai/mcp/service

创建MCP服务 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/mcp/service/create` | 创建MCP服务 | admin |
| POST | `/api/v1/ai/mcp/service/delete` | 删除MCP服务 | admin |
| POST | `/api/v1/ai/mcp/service/get-list` | 获取MCP服务列表 | admin |
| POST | `/api/v1/ai/mcp/service/get-one` | 获取MCP服务详情 | admin |
| POST | `/api/v1/ai/mcp/service/refresh-tools` | 刷新MCP服务工具列表 | admin |
| POST | `/api/v1/ai/mcp/service/update` | 更新MCP服务 | admin |

## 详细说明

### POST `/api/v1/ai/mcp/service/create`

**说明**: 创建MCP服务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authConfig` | string | 否 |  认证配置（JSON字符串） |
| `authType` | string | 否 |  认证类型 |
| `description` | string | 否 |  服务描述 |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `endpoint` | string | 是 |  服务端点URL |
| `id` | integer | 否 |  MCP服务ID (格式: int64) |
| `name` | string | 是 |  服务名称 |
| `tenantCode` | string | 否 |  租户编码 |
| `tools` | string | 否 |  工具列表（JSON字符串） |

**请求示例**:
```json
{
  "authConfig": "string",
  "authType": "string",
  "description": "string",
  "enabled": true,
  "endpoint": "string",
  "id": 1,
  "name": "示例名称",
  "tenantCode": "string",
  "tools": "string"
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
ur api /api/v1/ai/mcp/service/create \
  --body '{"authConfig": "string", "authType": "string", "description": "string", "enabled": true, "endpoint": "string", "id": 1, "name": "示例名称", "tenantCode": "string", "tools": "string"}'
```

### POST `/api/v1/ai/mcp/service/delete`

**说明**: 删除MCP服务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  MCP服务ID (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "id": 1,
  "tenantCode": "string"
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
ur api /api/v1/ai/mcp/service/delete \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/mcp/service/get-list`

**说明**: 获取MCP服务列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `enabled` | boolean | 否 |  启用状态过滤 (格式: boolean) |
| `nameLike` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "enabled": true,
  "nameLike": "示例名称",
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
        "authConfig": "string",
        "authType": "string",
        "description": "string",
        "enabled": true,
        "endpoint": "string",
        "id": 1,
        "name": "示例名称",
        "tenantCode": "string",
        "tools": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/service/get-list \
  --body '{"enabled": true, "nameLike": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "tenantCode": "string"}'
```

### POST `/api/v1/ai/mcp/service/get-one`

**说明**: 获取MCP服务详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  MCP服务ID (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "id": 1,
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "authConfig": "string",
    "authType": "string",
    "description": "string",
    "enabled": true,
    "endpoint": "string",
    "id": 1,
    "name": "示例名称",
    "tenantCode": "string",
    "tools": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/service/get-one \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/mcp/service/refresh-tools`

**说明**: 刷新MCP服务工具列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  MCP服务ID (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "id": 1,
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "tools": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/service/refresh-tools \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/mcp/service/update`

**说明**: 更新MCP服务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authConfig` | string | 否 |  认证配置（JSON字符串） |
| `authType` | string | 否 |  认证类型 |
| `description` | string | 否 |  服务描述 |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `endpoint` | string | 是 |  服务端点URL |
| `id` | integer | 否 |  MCP服务ID (格式: int64) |
| `name` | string | 是 |  服务名称 |
| `tenantCode` | string | 否 |  租户编码 |
| `tools` | string | 否 |  工具列表（JSON字符串） |

**请求示例**:
```json
{
  "authConfig": "string",
  "authType": "string",
  "description": "string",
  "enabled": true,
  "endpoint": "string",
  "id": 1,
  "name": "示例名称",
  "tenantCode": "string",
  "tools": "string"
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
ur api /api/v1/ai/mcp/service/update \
  --body '{"authConfig": "string", "authType": "string", "description": "string", "enabled": true, "endpoint": "string", "id": 1, "name": "示例名称", "tenantCode": "string", "tools": "string"}'
```
