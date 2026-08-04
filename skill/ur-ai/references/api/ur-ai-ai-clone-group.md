# ur-ai ai/clone-group

创建Clone组 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/clone-group/create` | 创建Clone组 | admin |
| POST | `/api/v1/ai/clone-group/delete` | 删除Clone组 | admin |
| POST | `/api/v1/ai/clone-group/get-list` | 获取Clone组列表 | admin |
| POST | `/api/v1/ai/clone-group/get-one` | 获取Clone组详情 | admin |
| POST | `/api/v1/ai/clone-group/update` | 更新Clone组 | admin |

## 详细说明

### POST `/api/v1/ai/clone-group/create`

**说明**: 创建Clone组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentGroupID` | integer | 是 |  关联AgentGroup ID (格式: int64) |
| `agentGroupName` | string | 否 |  关联AgentGroup名称 |
| `authType` | string | 否 |  权限类型 |
| `code` | string | 是 |  Clone组编码 |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `description` | string | 否 |  描述 |
| `id` | integer | 否 |  Clone组ID (格式: int64) |
| `isDefault` | string | 否 |  是否默认组 1-是 2-否 |
| `name` | string | 是 |  Clone组名称 |
| `status` | string | 否 |  状态 1-启用 2-禁用 |
| `tenantCode` | string | 否 |  租户编码 |
| `thirdID` | string | 否 |  第三方业务ID（如 projectID） |
| `ttsConfigID` | string | 否 |  组级TTS配置ID |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |
| `voiceParams` | string | 否 |  组级语音参数（JSON字符串） |

**请求示例**:
```json
{
  "agentGroupID": 1,
  "agentGroupName": "示例名称",
  "authType": "string",
  "code": "string",
  "createdTime": 1,
  "description": "string",
  "id": 1,
  "isDefault": "string",
  "name": "示例名称",
  "status": "string",
  "tenantCode": "string",
  "thirdID": "string",
  "ttsConfigID": "string",
  "updatedTime": 1,
  "voiceParams": "string"
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
ur api /api/v1/ai/clone-group/create \
  --body '{"agentGroupID": 1, "agentGroupName": "示例名称", "authType": "string", "code": "string", "createdTime": 1, "description": "string", "id": 1, "isDefault": "string", "name": "示例名称", "status": "string", "tenantCode": "string", "thirdID": "string", "ttsConfigID": "string", "updatedTime": 1, "voiceParams": "string"}'
```

### POST `/api/v1/ai/clone-group/delete`

**说明**: 删除Clone组

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
ur api /api/v1/ai/clone-group/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone-group/get-list`

**说明**: 获取Clone组列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentGroupID` | integer | 否 |  按AgentGroup ID过滤 (格式: int64) |
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  状态过滤 (格式: int64) |
| `thirdID` | string | 否 |  按第三方业务ID过滤 |

**请求示例**:
```json
{
  "agentGroupID": 1,
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
  "thirdID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "agentGroupID": 1,
        "agentGroupName": "示例名称",
        "authType": "string",
        "code": "string",
        "createdTime": 1,
        "description": "string",
        "id": 1,
        "isDefault": "string",
        "name": "示例名称",
        "status": "string",
        "tenantCode": "string",
        "thirdID": "string",
        "ttsConfigID": "string",
        "updatedTime": 1,
        "voiceParams": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone-group/get-list \
  --body '{"agentGroupID": 1, "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "thirdID": "string"}'
```

### POST `/api/v1/ai/clone-group/get-one`

**说明**: 获取Clone组详情

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
    "agentGroupID": 1,
    "agentGroupName": "示例名称",
    "authType": "string",
    "code": "string",
    "createdTime": 1,
    "description": "string",
    "id": 1,
    "isDefault": "string",
    "name": "示例名称",
    "status": "string",
    "tenantCode": "string",
    "thirdID": "string",
    "ttsConfigID": "string",
    "updatedTime": 1,
    "voiceParams": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone-group/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone-group/update`

**说明**: 更新Clone组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentGroupID` | integer | 是 |  关联AgentGroup ID (格式: int64) |
| `agentGroupName` | string | 否 |  关联AgentGroup名称 |
| `authType` | string | 否 |  权限类型 |
| `code` | string | 是 |  Clone组编码 |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `description` | string | 否 |  描述 |
| `id` | integer | 否 |  Clone组ID (格式: int64) |
| `isDefault` | string | 否 |  是否默认组 1-是 2-否 |
| `name` | string | 是 |  Clone组名称 |
| `status` | string | 否 |  状态 1-启用 2-禁用 |
| `tenantCode` | string | 否 |  租户编码 |
| `thirdID` | string | 否 |  第三方业务ID（如 projectID） |
| `ttsConfigID` | string | 否 |  组级TTS配置ID |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |
| `voiceParams` | string | 否 |  组级语音参数（JSON字符串） |

**请求示例**:
```json
{
  "agentGroupID": 1,
  "agentGroupName": "示例名称",
  "authType": "string",
  "code": "string",
  "createdTime": 1,
  "description": "string",
  "id": 1,
  "isDefault": "string",
  "name": "示例名称",
  "status": "string",
  "tenantCode": "string",
  "thirdID": "string",
  "ttsConfigID": "string",
  "updatedTime": 1,
  "voiceParams": "string"
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
ur api /api/v1/ai/clone-group/update \
  --body '{"agentGroupID": 1, "agentGroupName": "示例名称", "authType": "string", "code": "string", "createdTime": 1, "description": "string", "id": 1, "isDefault": "string", "name": "示例名称", "status": "string", "tenantCode": "string", "thirdID": "string", "ttsConfigID": "string", "updatedTime": 1, "voiceParams": "string"}'
```
