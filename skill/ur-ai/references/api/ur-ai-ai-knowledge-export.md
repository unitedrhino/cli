# ur-ai ai/knowledge-export

创建导出任务 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/knowledge/export/create` | 创建导出任务 | admin |
| POST | `/api/v1/ai/knowledge/export/export-zip` | ZIP导出知识库 | admin |
| POST | `/api/v1/ai/knowledge/export/get-status` | 获取导出状态 | admin |

## 详细说明

### POST `/api/v1/ai/knowledge/export/create`

**说明**: 创建导出任务

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
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/export/create \
  --body '{"knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/export/export-zip`

**说明**: ZIP导出知识库

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
    "downloadURL": "string",
    "taskID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/export/export-zip \
  --body '{"knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/export/get-status`

**说明**: 获取导出状态

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `taskID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "taskID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "downloadURL": "string",
    "errorMessage": "string",
    "status": "string",
    "taskID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/export/get-status \
  --body '{"taskID": 1}'
```
