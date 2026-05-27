# ur-ai ai/tool

创建工具 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/tool/create` | 创建工具 | admin |
| POST | `/api/v1/ai/tool/delete` | 删除工具 | admin |
| POST | `/api/v1/ai/tool/disable` | 停用工具 | admin |
| POST | `/api/v1/ai/tool/get-list` | 获取工具列表 | admin |
| POST | `/api/v1/ai/tool/get-one` | 获取工具详情 | admin |
| POST | `/api/v1/ai/tool/publish` | 发布工具 | admin |
| POST | `/api/v1/ai/tool/save-artifact` | 保存三件套 | admin |
| POST | `/api/v1/ai/tool/update` | 更新工具 | admin |

## 详细说明

### POST `/api/v1/ai/tool/create`

**说明**: 创建工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentGroupID` | string | 是 |  关联AgentGroup ID |
| `code` | string | 是 |  工具编码 |
| `description` | string | 否 |  工具描述 |
| `name` | string | 是 |  工具名称 |

**请求示例**:
```json
{
  "agentGroupID": "string",
  "code": "string",
  "description": "string",
  "name": "示例名称"
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
ur api /api/v1/ai/tool/create \
  --body '{"agentGroupID": "string", "code": "string", "description": "string", "name": "示例名称"}'
```

### POST `/api/v1/ai/tool/delete`

**说明**: 删除工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
ur api /api/v1/ai/tool/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/tool/disable`

**说明**: 停用工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
ur api /api/v1/ai/tool/disable \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/tool/get-list`

**说明**: 获取工具列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentGroupID` | string | 否 |  AgentGroup过滤 |
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态过滤 |

**请求示例**:
```json
{
  "agentGroupID": "string",
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "agentGroupID": "string",
        "code": "string",
        "createdTime": 1,
        "description": "string",
        "groupName": "示例名称",
        "id": "string",
        "name": "示例名称",
        "status": "string",
        "tenantCode": "string",
        "version": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/get-list \
  --body '{"agentGroupID": "string", "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "status": "string"}'
```

### POST `/api/v1/ai/tool/get-one`

**说明**: 获取工具详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  工具ID |

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
    "artifact": {
      "documentMd": "string",
      "executorJs": "string",
      "manifestJson": "string"
    },
    "info": {
      "agentGroupID": "string",
      "code": "string",
      "createdTime": 1,
      "description": "string",
      "groupName": "示例名称",
      "id": "string",
      "name": "示例名称",
      "status": "string",
      "tenantCode": "string",
      "version": "string"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/ai/tool/publish`

**说明**: 发布工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
ur api /api/v1/ai/tool/publish \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/tool/save-artifact`

**说明**: 保存三件套

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `documentMd` | string | 否 |  document.md 内容 |
| `executorJs` | string | 是 |  executor.js 内容 |
| `id` | string | 是 |  工具ID |
| `manifestJson` | string | 否 |  manifest.json 内容 |

**请求示例**:
```json
{
  "documentMd": "string",
  "executorJs": "string",
  "id": "string",
  "manifestJson": "string"
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
ur api /api/v1/ai/tool/save-artifact \
  --body '{"documentMd": "string", "executorJs": "string", "id": "string", "manifestJson": "string"}'
```

### POST `/api/v1/ai/tool/update`

**说明**: 更新工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `description` | string | 否 |  工具描述 |
| `id` | string | 是 |  工具ID |
| `name` | string | 否 |  工具名称 |

**请求示例**:
```json
{
  "description": "string",
  "id": "string",
  "name": "示例名称"
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
ur api /api/v1/ai/tool/update \
  --body '{"description": "string", "id": "string", "name": "示例名称"}'
```
