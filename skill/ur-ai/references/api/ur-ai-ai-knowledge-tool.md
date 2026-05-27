# ur-ai ai/knowledge-tool

CLI/调试-获取切片关联 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/knowledge/tool/get-chunk-relations` | CLI/调试-获取切片关联 | admin |
| POST | `/api/v1/ai/knowledge/tool/get-document-content` | CLI/调试-获取文档全文 | admin |
| POST | `/api/v1/ai/knowledge/tool/search` | CLI/调试-搜索知识库 | admin |

## 详细说明

### POST `/api/v1/ai/knowledge/tool/get-chunk-relations`

**说明**: CLI/调试-获取切片关联

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `artifactID` | integer | 是 | 格式: int64 |
| `limit` | integer | 否 | 格式: int64 |
| `minScore` | number | 否 | 格式: double |
| `relationTypes` | array[string] | 否 |  |

**请求示例**:
```json
{
  "artifactID": 1,
  "limit": 1,
  "minScore": 1,
  "relationTypes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "artifact": {
      "artifactType": "string",
      "charEnd": 1,
      "charStart": 1,
      "chunkIndex": 1,
      "content": "string",
      "createdTime": 1,
      "documentID": 1,
      "headingPath": "string",
      "id": 1,
      "knowledgeID": 1,
      "relatedCount": 1,
      "sourceDocumentIDs": [
        1
      ],
      "title": "string",
      "tokenEstimate": 1,
      "updatedTime": 1
    },
    "relations": [
      {
        "cosineScore": 1,
        "evidence": "string",
        "fromArtifactID": 1,
        "fromDocumentID": 1,
        "id": 1,
        "lexicalOverlap": 1,
        "llmConfidence": 1,
        "relationType": "string",
        "score": 1,
        "targetArtifact": {
          "artifactType": "string",
          "charEnd": 1,
          "charStart": 1,
          "chunkIndex": 1,
          "content": "string",
          "createdTime": 1,
          "documentID": 1,
          "headingPath": "string",
          "id": 1,
          "knowledgeID": 1,
          "relatedCount": 1,
          "sourceDocumentIDs": [
            1
          ],
          "title": "string",
          "tokenEstimate": 1,
          "updatedTime": 1
        },
        "targetDocument": {
          "children": [
            {
              "children": "...",
              "compileStatus": "...",
              "id": "...",
              "knowledgeID": "...",
              "lastError": "...",
              "name": "...",
              "nodeType": "...",
              "parentID": "...",
              "parseStatus": "...",
              "path": "...",
              "searchMessage": "...",
              "searchStatus": "...",
              "uri": "..."
            }
          ],
          "compileStatus": "string",
          "id": 1,
          "knowledgeID": 1,
          "lastError": "string",
          "name": "string",
          "nodeType": "string",
          "parentID": 1,
          "parseStatus": "string",
          "path": "string",
          "searchMessage": "string",
          "searchStatus": "string",
          "uri": "string"
        },
        "toArtifactID": 1,
        "toDocumentID": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/tool/get-chunk-relations \
  --body '{"artifactID": 1, "limit": 1, "minScore": 1, "relationTypes": ["string"]}'
```

### POST `/api/v1/ai/knowledge/tool/get-document-content`

**说明**: CLI/调试-获取文档全文

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `documentID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "documentID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "documentID": 1,
    "parsedFullText": "string",
    "sourceTextReady": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/tool/get-document-content \
  --body '{"documentID": 1}'
```

### POST `/api/v1/ai/knowledge/tool/search`

**说明**: CLI/调试-搜索知识库

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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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
ur api /api/v1/ai/knowledge/tool/search \
  --body '{"agentID": 1, "knowledgeBaseIDs": [1], "minScore": 1, "query": "string", "topK": 1}'
```
