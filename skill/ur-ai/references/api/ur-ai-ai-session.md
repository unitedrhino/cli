# ur-ai ai/session

归档会话 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/session/archive` | 归档会话 | admin |
| POST | `/api/v1/ai/session/create` | 创建会话 | admin |
| POST | `/api/v1/ai/session/delete` | 删除会话 | admin |
| POST | `/api/v1/ai/session/end` | 结束会话 | admin |
| POST | `/api/v1/ai/session/get-list` | 获取会话列表 | admin |
| POST | `/api/v1/ai/session/get-one` | 获取会话详情 | admin |
| POST | `/api/v1/ai/session/inject` | 注入消息到会话 | admin |
| POST | `/api/v1/ai/session/messages` | 获取会话消息列表 | admin |

## 详细说明

### POST `/api/v1/ai/session/archive`

**说明**: 归档会话

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
ur api /api/v1/ai/session/archive \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/session/create`

**说明**: 创建会话

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 是 |  关联Agent ID (格式: int64) |
| `channel` | string | 否 |  渠道 text/voice |
| `cloneID` | integer | 否 |  关联Clone ID（0=无分身） (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |
| `userID` | integer | 否 |  用户ID (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "channel": "string",
  "cloneID": 1,
  "tenantCode": "string",
  "userID": 1
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
ur api /api/v1/ai/session/create \
  --body '{"agentID": 1, "channel": "string", "cloneID": 1, "tenantCode": "string", "userID": 1}'
```

### POST `/api/v1/ai/session/delete`

**说明**: 删除会话

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `sessionID` | string | 是 |  会话ID |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "sessionID": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "success": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/delete \
  --body '{"sessionID": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/session/end`

**说明**: 结束会话

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
ur api /api/v1/ai/session/end \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/session/get-list`

**说明**: 获取会话列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 |  按Agent ID过滤 (格式: int64) |
| `cloneID` | integer | 否 |  按Clone ID过滤 (格式: int64) |
| `endTime` | integer | 否 |  结束时间戳 (格式: int64) |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `startTime` | integer | 否 |  开始时间戳 (格式: int64) |
| `status` | integer | 否 |  状态过滤 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码过滤 |
| `userID` | integer | 否 |  按用户ID过滤 (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "cloneID": 1,
  "endTime": 1,
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "startTime": 1,
  "status": 1,
  "tenantCode": "string",
  "userID": 1
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
        "channel": "string",
        "cloneID": 1,
        "createdTime": 1,
        "id": 1,
        "lastActiveAt": 1,
        "sessionID": "string",
        "status": 1,
        "tenantCode": "string",
        "transcriptOssKey": "string",
        "turnCount": 1,
        "userID": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/get-list \
  --body '{"agentID": 1, "cloneID": 1, "endTime": 1, "page": {"page": 1, "pageSize": 1}, "startTime": 1, "status": 1, "tenantCode": "string", "userID": 1}'
```

### POST `/api/v1/ai/session/get-one`

**说明**: 获取会话详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  会话ID (格式: int64) |
| `sessionID` | string | 否 |  会话唯一标识（与ID二选一） |

**请求示例**:
```json
{
  "id": 1,
  "sessionID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "agentID": 1,
    "channel": "string",
    "cloneID": 1,
    "createdTime": 1,
    "id": 1,
    "lastActiveAt": 1,
    "sessionID": "string",
    "status": 1,
    "tenantCode": "string",
    "transcriptOssKey": "string",
    "turnCount": 1,
    "userID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/get-one \
  --body '{"id": 1, "sessionID": "string"}'
```

### POST `/api/v1/ai/session/inject`

**说明**: 注入消息到会话

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `message` | string | 是 |  注入消息内容 |
| `messageType` | string | 否 |  消息类型 |
| `sessionID` | string | 是 |  会话ID |

**请求示例**:
```json
{
  "message": "string",
  "messageType": "string",
  "sessionID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "message": "string",
    "success": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/inject \
  --body '{"message": "string", "messageType": "string", "sessionID": "string"}'
```

### POST `/api/v1/ai/session/messages`

**说明**: 获取会话消息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `sessionID` | string | 是 |  会话ID |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "sessionID": "string",
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
        "aiResponse": "string",
        "keyTopics": "string",
        "summary": "string",
        "turnIndex": 1,
        "userInput": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/messages \
  --body '{"page": {"page": 1, "pageSize": 1}, "sessionID": "string", "tenantCode": "string"}'
```
