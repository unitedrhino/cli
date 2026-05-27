# ur-ai 未分组

发送对话消息 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/app/chat-send` | 发送对话消息 | - |
| POST | `/api/v1/ai/app/sse/chat-run` | sse模式ai聊天运行 | - |
| POST | `/api/v1/ai/app/sse/workflow-run` | sse模式执行 workflow | - |
| POST | `/api/v1/ai/app/workflow-run` | 执行 workflow | - |
| POST | `/api/v1/ai/mcp/run` | 添加规则链 | - |

## 详细说明

### POST `/api/v1/ai/app/chat-send`

**说明**: 发送对话消息

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appCode` | string | 是 | 执行的应用 |
| `conversationID` | string | 否 | （选填）会话 ID，需要基于之前的聊天记录继续对话，必须传之前消息的 conversation_id。 |
| `inputs` | object | 否 |  |
| `query` | string | 是 | 用户输入/提问内容 |
| `user` | string | 是 | 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。 |

**请求示例**:
```json
{
  "appCode": "string",
  "conversationID": "string",
  "inputs": {},
  "query": "string",
  "user": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "answer": "string",
    "conversationID": "string",
    "createdAt": 1,
    "event": "string",
    "messageID": "string",
    "mode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/app/chat-send \
  --body '{"appCode": "string", "conversationID": "string", "inputs": {}, "query": "string", "user": "string"}'
```

### POST `/api/v1/ai/app/sse/chat-run`

**说明**: sse模式ai聊天运行

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appCode` | string | 是 | 执行的应用 |
| `conversationID` | string | 否 | （选填）会话 ID，需要基于之前的聊天记录继续对话，必须传之前消息的 conversation_id。 |
| `inputs` | object | 否 |  |
| `query` | string | 是 | 用户输入/提问内容 |
| `user` | string | 是 | 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。 |

**请求示例**:
```json
{
  "appCode": "string",
  "conversationID": "string",
  "inputs": {},
  "query": "string",
  "user": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "answer": "string",
    "conversationID": "string",
    "createdAt": 1,
    "event": "string",
    "messageID": "string",
    "mode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/app/sse/chat-run \
  --body '{"appCode": "string", "conversationID": "string", "inputs": {}, "query": "string", "user": "string"}'
```

### POST `/api/v1/ai/app/sse/workflow-run`

**说明**: sse模式执行 workflow

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appCode` | string | 是 | 执行的应用 |
| `inputs` | object | 否 |  |
| `user` | string | 是 | 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。 |

**请求示例**:
```json
{
  "appCode": "string",
  "inputs": {},
  "user": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdAt": 1,
    "elapsedTime": 1,
    "error": "string",
    "event": "string",
    "finishedAt": 1,
    "inputs": {},
    "outputs": {},
    "status": "string",
    "taskID": "string",
    "workflowID": "string",
    "workflowRunID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/app/sse/workflow-run \
  --body '{"appCode": "string", "inputs": {}, "user": "string"}'
```

### POST `/api/v1/ai/app/workflow-run`

**说明**: 执行 workflow

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appCode` | string | 是 | 执行的应用 |
| `inputs` | object | 否 |  |
| `user` | string | 是 | 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。 |

**请求示例**:
```json
{
  "appCode": "string",
  "inputs": {},
  "user": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdAt": 1,
    "elapsedTime": 1,
    "error": "string",
    "event": "string",
    "finishedAt": 1,
    "inputs": {},
    "outputs": {},
    "status": "string",
    "taskID": "string",
    "workflowID": "string",
    "workflowRunID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/app/workflow-run \
  --body '{"appCode": "string", "inputs": {}, "user": "string"}'
```

### POST `/api/v1/ai/mcp/run`

**说明**: 添加规则链

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/run \
  --body '{}'
```
