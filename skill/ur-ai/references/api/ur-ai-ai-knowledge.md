# ur-ai ai/knowledge

创建知识库 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/knowledge/info/create` | 创建知识库 | admin |
| POST | `/api/v1/ai/knowledge/info/delete` | 删除知识库 | admin |
| POST | `/api/v1/ai/knowledge/info/get-list` | 获取知识库列表 | admin |
| POST | `/api/v1/ai/knowledge/info/get-one` | 获取知识库详情 | admin |
| POST | `/api/v1/ai/knowledge/info/search` | 搜索知识库切片 | admin |
| POST | `/api/v1/ai/knowledge/info/update` | 更新知识库 | admin |

## 详细说明

### POST `/api/v1/ai/knowledge/info/create`

**说明**: 创建知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `compiledDocumentCount` | integer | 否 | 格式: int64 |
| `createdTime` | integer | 否 | 格式: int64 |
| `desc` | string | 否 |  |
| `documentCount` | integer | 否 | 格式: int64 |
| `failedDocumentCount` | integer | 否 | 格式: int64 |
| `id` | integer | 否 | 格式: int64 |
| `lastCompileTime` | integer | 否 | 格式: int64 |
| `name` | string | 是 |  |
| `scope` | string | 否 | 归属层级：platform=平台公共库(common)，tenant=租户级库 |
| `status` | string | 否 |  |
| `tenantCode` | string | 否 |  |
| `updatedTime` | integer | 否 | 格式: int64 |

**请求示例**:
```json
{
  "compiledDocumentCount": 1,
  "createdTime": 1,
  "desc": "string",
  "documentCount": 1,
  "failedDocumentCount": 1,
  "id": 1,
  "lastCompileTime": 1,
  "name": "string",
  "scope": "string",
  "status": "string",
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
ur api /api/v1/ai/knowledge/info/create \
  --body '{"compiledDocumentCount": 1, "createdTime": 1, "desc": "string", "documentCount": 1, "failedDocumentCount": 1, "id": 1, "lastCompileTime": 1, "name": "string", "scope": "string", "status": "string", "tenantCode": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/knowledge/info/delete`

**说明**: 删除知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 | 格式: int64 |

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
ur api /api/v1/ai/knowledge/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/knowledge/info/get-list`

**说明**: 获取知识库列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `scope` | string | 否 | 可选过滤：platform=仅平台公共库，tenant=仅本租户库，空=全部可见范围 |
| `status` | string | 否 |  |
| `tenantCode` | string | 否 |  |

**请求示例**:
```json
{
  "name": "string",
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
  "scope": "string",
  "status": "string",
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
        "compiledDocumentCount": 1,
        "createdTime": 1,
        "desc": "string",
        "documentCount": 1,
        "failedDocumentCount": 1,
        "id": 1,
        "lastCompileTime": 1,
        "name": "string",
        "scope": "string",
        "status": "string",
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
ur api /api/v1/ai/knowledge/info/get-list \
  --body '{"name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "scope": "string", "status": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/knowledge/info/get-one`

**说明**: 获取知识库详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 | 格式: int64 |

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
    "compiledDocumentCount": 1,
    "createdTime": 1,
    "desc": "string",
    "documentCount": 1,
    "failedDocumentCount": 1,
    "id": 1,
    "lastCompileTime": 1,
    "name": "string",
    "scope": "string",
    "status": "string",
    "tenantCode": "string",
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/knowledge/info/search`

**说明**: 搜索知识库切片

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 | 格式: int64 |
| `knowledgeBaseIDs` | array[integer] | 否 |  |
| `minScore` | number | 否 | 格式: double |
| `query` | string | 是 |  |
| `topK` | integer | 否 | 格式: int64 |

**请求示例**:
```json
{
  "agentID": 1,
  "knowledgeBaseIDs": [
    1
  ],
  "minScore": 1,
  "query": "string",
  "topK": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "chunks": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "confidenceBand": "string",
    "documents": [
      {
        "chunkCount": 1,
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "score": 1,
        "snippet": "string"
      }
    ],
    "expandedChunks": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "list": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "originalQuery": "string",
    "rewrittenQuery": "string",
    "searchReady": true,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/info/search \
  --body '{"agentID": 1, "knowledgeBaseIDs": [1], "minScore": 1, "query": "string", "topK": 1}'
```

### POST `/api/v1/ai/knowledge/info/update`

**说明**: 更新知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `compiledDocumentCount` | integer | 否 | 格式: int64 |
| `createdTime` | integer | 否 | 格式: int64 |
| `desc` | string | 否 |  |
| `documentCount` | integer | 否 | 格式: int64 |
| `failedDocumentCount` | integer | 否 | 格式: int64 |
| `id` | integer | 否 | 格式: int64 |
| `lastCompileTime` | integer | 否 | 格式: int64 |
| `name` | string | 是 |  |
| `scope` | string | 否 | 归属层级：platform=平台公共库(common)，tenant=租户级库 |
| `status` | string | 否 |  |
| `tenantCode` | string | 否 |  |
| `updatedTime` | integer | 否 | 格式: int64 |

**请求示例**:
```json
{
  "compiledDocumentCount": 1,
  "createdTime": 1,
  "desc": "string",
  "documentCount": 1,
  "failedDocumentCount": 1,
  "id": 1,
  "lastCompileTime": 1,
  "name": "string",
  "scope": "string",
  "status": "string",
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
ur api /api/v1/ai/knowledge/info/update \
  --body '{"compiledDocumentCount": 1, "createdTime": 1, "desc": "string", "documentCount": 1, "failedDocumentCount": 1, "id": 1, "lastCompileTime": 1, "name": "string", "scope": "string", "status": "string", "tenantCode": "string", "updatedTime": 1}'
```
