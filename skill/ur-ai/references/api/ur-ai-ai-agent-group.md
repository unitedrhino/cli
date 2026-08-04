# ur-ai ai/agent-group

创建助手组 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/agent/group/create` | 创建助手组 | admin |
| POST | `/api/v1/ai/agent/group/delete` | 删除助手组 | admin |
| POST | `/api/v1/ai/agent/group/get-list` | 获取助手组列表 | admin |
| POST | `/api/v1/ai/agent/group/get-one` | 获取助手组详情 | admin |
| POST | `/api/v1/ai/agent/group/update` | 更新助手组 | admin |

## 详细说明

### POST `/api/v1/ai/agent/group/create`

**说明**: 创建助手组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `asrConfigID` | string | 否 |  共享默认ASR配置ID |
| `code` | string | 是 |  助手组编码（租户内唯一） |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `defaultAgentId` | integer | 否 |  默认决策AgentID（0=未设置） (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  助手组ID (格式: int64) |
| `knowledgeBaseIds` | string | 否 |  共享知识库ID列表（JSON字符串，供组内Agent复用） |
| `llmConfigID` | string | 否 |  共享默认LLM配置ID |
| `mcpServiceIds` | string | 否 |  共享MCP服务ID列表（JSON字符串，供组内Agent复用） |
| `name` | string | 是 |  助手组名称 |
| `purpose` | string | 否 |  用途分类 default/customer_service/... |
| `skillIds` | string | 否 |  平台分发Skill ID列表（JSON字符串，供组内Agent复用，不表示Group或Clone拥有learned skills） |
| `status` | string | 否 |  状态 1-启用 2-禁用 |
| `systemPrompt` | string | 否 |  共享系统提示词模板（供组内Agent决策时继承） |
| `tenantCode` | string | 否 |  租户编码，"common"=平台级 |
| `ttsConfigID` | string | 否 |  共享默认TTS配置ID |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "asrConfigID": "string",
  "code": "string",
  "createdTime": 1,
  "defaultAgentId": 1,
  "desc": "string",
  "id": 1,
  "knowledgeBaseIds": "string",
  "llmConfigID": "string",
  "mcpServiceIds": "string",
  "name": "示例名称",
  "purpose": "string",
  "skillIds": "string",
  "status": "string",
  "systemPrompt": "string",
  "tenantCode": "string",
  "ttsConfigID": "string",
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
ur api /api/v1/ai/agent/group/create \
  --body '{"asrConfigID": "string", "code": "string", "createdTime": 1, "defaultAgentId": 1, "desc": "string", "id": 1, "knowledgeBaseIds": "string", "llmConfigID": "string", "mcpServiceIds": "string", "name": "示例名称", "purpose": "string", "skillIds": "string", "status": "string", "systemPrompt": "string", "tenantCode": "string", "ttsConfigID": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/agent/group/delete`

**说明**: 删除助手组

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
ur api /api/v1/ai/agent/group/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/agent/group/get-list`

**说明**: 获取助手组列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `purpose` | string | 否 |  用途分类过滤 |
| `tenantCode` | string | 否 |  租户编码过滤 |

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
  "purpose": "string",
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
        "asrConfigID": "string",
        "code": "string",
        "createdTime": 1,
        "defaultAgentId": 1,
        "desc": "string",
        "id": 1,
        "knowledgeBaseIds": "string",
        "llmConfigID": "string",
        "mcpServiceIds": "string",
        "name": "示例名称",
        "purpose": "string",
        "skillIds": "string",
        "status": "string",
        "systemPrompt": "string",
        "tenantCode": "string",
        "ttsConfigID": "string",
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
ur api /api/v1/ai/agent/group/get-list \
  --body '{"name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "purpose": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/agent/group/get-one`

**说明**: 获取助手组详情

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
    "asrConfigID": "string",
    "code": "string",
    "createdTime": 1,
    "defaultAgentId": 1,
    "desc": "string",
    "id": 1,
    "knowledgeBaseIds": "string",
    "llmConfigID": "string",
    "mcpServiceIds": "string",
    "name": "示例名称",
    "purpose": "string",
    "skillIds": "string",
    "status": "string",
    "systemPrompt": "string",
    "tenantCode": "string",
    "ttsConfigID": "string",
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/agent/group/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/agent/group/update`

**说明**: 更新助手组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `asrConfigID` | string | 否 |  共享默认ASR配置ID |
| `code` | string | 是 |  助手组编码（租户内唯一） |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `defaultAgentId` | integer | 否 |  默认决策AgentID（0=未设置） (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  助手组ID (格式: int64) |
| `knowledgeBaseIds` | string | 否 |  共享知识库ID列表（JSON字符串，供组内Agent复用） |
| `llmConfigID` | string | 否 |  共享默认LLM配置ID |
| `mcpServiceIds` | string | 否 |  共享MCP服务ID列表（JSON字符串，供组内Agent复用） |
| `name` | string | 是 |  助手组名称 |
| `purpose` | string | 否 |  用途分类 default/customer_service/... |
| `skillIds` | string | 否 |  平台分发Skill ID列表（JSON字符串，供组内Agent复用，不表示Group或Clone拥有learned skills） |
| `status` | string | 否 |  状态 1-启用 2-禁用 |
| `systemPrompt` | string | 否 |  共享系统提示词模板（供组内Agent决策时继承） |
| `tenantCode` | string | 否 |  租户编码，"common"=平台级 |
| `ttsConfigID` | string | 否 |  共享默认TTS配置ID |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "asrConfigID": "string",
  "code": "string",
  "createdTime": 1,
  "defaultAgentId": 1,
  "desc": "string",
  "id": 1,
  "knowledgeBaseIds": "string",
  "llmConfigID": "string",
  "mcpServiceIds": "string",
  "name": "示例名称",
  "purpose": "string",
  "skillIds": "string",
  "status": "string",
  "systemPrompt": "string",
  "tenantCode": "string",
  "ttsConfigID": "string",
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
ur api /api/v1/ai/agent/group/update \
  --body '{"asrConfigID": "string", "code": "string", "createdTime": 1, "defaultAgentId": 1, "desc": "string", "id": 1, "knowledgeBaseIds": "string", "llmConfigID": "string", "mcpServiceIds": "string", "name": "示例名称", "purpose": "string", "skillIds": "string", "status": "string", "systemPrompt": "string", "tenantCode": "string", "ttsConfigID": "string", "updatedTime": 1}'
```
