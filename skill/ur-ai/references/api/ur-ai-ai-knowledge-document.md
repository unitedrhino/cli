# ur-ai ai/knowledge-document

获取切片关联列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/knowledge/document/chunk/relation/get-list` | 获取切片关联列表 | admin |
| POST | `/api/v1/ai/knowledge/document/get-content` | 获取文档全文 | admin |
| POST | `/api/v1/ai/knowledge/document/get-detail` | 获取文档详情 | admin |
| POST | `/api/v1/ai/knowledge/document/get-failed-list` | 获取失败文档列表 | admin |
| POST | `/api/v1/ai/knowledge/document/get-tree` | 获取知识库目录树 | admin |
| POST | `/api/v1/ai/knowledge/document/import` | 导入知识库文档 | admin |
| POST | `/api/v1/ai/knowledge/document/import-zip` | ZIP导入知识库 | admin |

## 详细说明

### POST `/api/v1/ai/knowledge/document/chunk/relation/get-list`

**说明**: 获取切片关联列表

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
ur api /api/v1/ai/knowledge/document/chunk/relation/get-list \
  --body '{"artifactID": 1, "limit": 1, "minScore": 1, "relationTypes": ["string"]}'
```

### POST `/api/v1/ai/knowledge/document/get-content`

**说明**: 获取文档全文

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
ur api /api/v1/ai/knowledge/document/get-content \
  --body '{"documentID": 1}'
```

### POST `/api/v1/ai/knowledge/document/get-detail`

**说明**: 获取文档详情

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
    "articles": [
      {
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
      }
    ],
    "chunks": [
      {
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
      }
    ],
    "compileStatus": "string",
    "concepts": [
      {
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
      }
    ],
    "document": {
      "children": [
        {
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
    "lastCompileTask": {
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
    "lastError": "string",
    "previewText": "string",
    "sourceDownloadURL": "string",
    "summary": {
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
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/get-detail \
  --body '{"documentID": 1}'
```

### POST `/api/v1/ai/knowledge/document/get-failed-list`

**说明**: 获取失败文档列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "knowledgeID": 1,
  "page": {
    "page": 1,
    "pageSize": 1
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "children": [
          {
            "children": [],
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
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/get-failed-list \
  --body '{"knowledgeID": 1, "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/ai/knowledge/document/get-tree`

**说明**: 获取知识库目录树

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
    "knowledgeID": 1,
    "nodes": [
      {
        "children": [
          {
            "children": [],
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
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/get-tree \
  --body '{"knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/document/import`

**说明**: 导入知识库文档

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `items` | array[KnowledgeImportItem] | 是 |  |
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "items": [
    {
      "filePath": "string",
      "name": "string",
      "path": "string",
      "rawContent": "string"
    }
  ],
  "knowledgeID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "importedCount": 1,
    "taskID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/import \
  --body '{"items": [{"filePath": "string", "name": "string", "path": "string", "rawContent": "string"}], "knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/document/import-zip`

**说明**: ZIP导入知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `fileURL` | string | 是 |  |
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "fileURL": "string",
  "knowledgeID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "importedCount": 1,
    "taskID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/import-zip \
  --body '{"fileURL": "string", "knowledgeID": 1}'
```
