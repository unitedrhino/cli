# ur-ai ai/knowledge-compile

获取编译状态 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/knowledge/compile/get-status` | 获取编译状态 | admin |
| POST | `/api/v1/ai/knowledge/compile/retry` | 触发知识库重编译 | admin |

## 详细说明

### POST `/api/v1/ai/knowledge/compile/get-status`

**说明**: 获取编译状态

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "knowledgeID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "compiledDocumentCount": 1,
    "currentTask": {
      "createdTime": 1,
      "errorMessage": "string",
      "finishedTime": 1,
      "id": 1,
      "knowledgeID": 1,
      "startedTime": 1,
      "status": "string",
      "triggerType": "string",
      "updatedTime": 1
    },
    "documentCount": 1,
    "failedDocumentCount": 1,
    "knowledgeID": 1,
    "relationMessage": "string",
    "relationStatus": "string",
    "searchFailedDocumentCount": 1,
    "searchMessage": "string",
    "searchReadyDocumentCount": 1,
    "searchStatus": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/compile/get-status \
  --body '{"knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/compile/retry`

**说明**: 触发知识库重编译

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |
| `mode` | string | 否 |  |

**请求示例**:
```json
{
  "knowledgeID": 1,
  "mode": "string"
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
ur api /api/v1/ai/knowledge/compile/retry \
  --body '{"knowledgeID": 1, "mode": "string"}'
```
