# ur-ai ai/config

创建配置 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/config/info/create` | 创建配置 | admin |
| POST | `/api/v1/ai/config/info/delete` | 删除配置 | admin |
| POST | `/api/v1/ai/config/info/get-default` | 获取默认配置 | admin |
| POST | `/api/v1/ai/config/info/get-list` | 获取配置列表 | admin |
| POST | `/api/v1/ai/config/info/get-one` | 获取配置详情 | admin |
| POST | `/api/v1/ai/config/info/test` | 测试配置 | admin |
| POST | `/api/v1/ai/config/info/update` | 更新配置 | admin |

## 详细说明

### POST `/api/v1/ai/config/info/create`

**说明**: 创建配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 否 |  配置类型：llm/tts/asr/vad |
| `configID` | string | 否 |  配置唯一标识 |
| `contextLength` | integer | 否 |  上下文长度 (格式: int64) |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `id` | integer | 否 |  配置ID (格式: int64) |
| `inputModalities` | string | 否 |  输入模态（JSON字符串） |
| `isDefault` | boolean | 否 |  是否默认 (格式: boolean) |
| `jsonData` | string | 否 |  配置参数（JSON字符串） |
| `modelType` | string | 否 |  模型级别：small/medium/large/xlarge |
| `name` | string | 否 |  配置名称 |
| `outputModalities` | string | 否 |  输出模态（JSON字符串） |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "configID": "string",
  "contextLength": 1,
  "enabled": true,
  "id": 1,
  "inputModalities": "string",
  "isDefault": true,
  "jsonData": "string",
  "modelType": "string",
  "name": "示例名称",
  "outputModalities": "string",
  "tenantCode": "string"
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
ur api /api/v1/ai/config/info/create \
  --body '{"cfgType": "string", "configID": "string", "contextLength": 1, "enabled": true, "id": 1, "inputModalities": "string", "isDefault": true, "jsonData": "string", "modelType": "string", "name": "示例名称", "outputModalities": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/delete`

**说明**: 删除配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  配置ID (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "id": 1,
  "tenantCode": "string"
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
ur api /api/v1/ai/config/info/delete \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/get-default`

**说明**: 获取默认配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 是 |  配置类型 |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "cfgType": "string",
    "configID": "string",
    "contextLength": 1,
    "enabled": true,
    "id": 1,
    "inputModalities": "string",
    "isDefault": true,
    "jsonData": "string",
    "modelType": "string",
    "name": "示例名称",
    "outputModalities": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/config/info/get-default \
  --body '{"cfgType": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/get-list`

**说明**: 获取配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 否 |  配置类型过滤 |
| `onlyEnabled` | boolean | 否 |  仅返回启用的 (格式: boolean) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "onlyEnabled": true,
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
        "cfgType": "string",
        "configID": "string",
        "contextLength": 1,
        "enabled": true,
        "id": 1,
        "inputModalities": "string",
        "isDefault": true,
        "jsonData": "string",
        "modelType": "string",
        "name": "示例名称",
        "outputModalities": "string",
        "tenantCode": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/config/info/get-list \
  --body '{"cfgType": "string", "onlyEnabled": true, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/get-one`

**说明**: 获取配置详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  配置ID (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "id": 1,
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "cfgType": "string",
    "configID": "string",
    "contextLength": 1,
    "enabled": true,
    "id": 1,
    "inputModalities": "string",
    "isDefault": true,
    "jsonData": "string",
    "modelType": "string",
    "name": "示例名称",
    "outputModalities": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/config/info/get-one \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/test`

**说明**: 测试配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 是 |  配置类型，目前先支持 llm |
| `configID` | string | 否 |  配置唯一标识，仅用于调试展示 |
| `inputModalities` | string | 否 |  输入模态（JSON字符串） |
| `jsonData` | string | 是 |  配置参数（JSON字符串） |
| `messages` | array[ChatMessage] | 是 |  调试输入消息 |
| `modelType` | string | 否 |  模型级别 |
| `name` | string | 否 |  配置名称，仅用于调试展示 |
| `outputModalities` | string | 否 |  输出模态（JSON字符串） |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "configID": "string",
  "inputModalities": "string",
  "jsonData": "string",
  "messages": [
    {
      "contents": [
        {
          "audioData": "string",
          "audioFormat": "string",
          "fileMime": "string",
          "fileName": "string",
          "fileUrl": "string",
          "imageUrl": "string",
          "text": "string",
          "type": "string",
          "videoUrl": "string"
        }
      ],
      "role": "string",
      "toolCallID": "string",
      "toolCalls": [
        {
          "arguments": "string",
          "id": "string",
          "name": "示例名称"
        }
      ]
    }
  ],
  "modelType": "string",
  "name": "示例名称",
  "outputModalities": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "answerPolicy": "string",
    "audioData": "string",
    "audioFormat": "string",
    "citations": [
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
    "content": "string",
    "contents": [
      {
        "audioData": "string",
        "audioFormat": "string",
        "fileMime": "string",
        "fileName": "string",
        "fileUrl": "string",
        "imageUrl": "string",
        "text": "string",
        "type": "string",
        "videoUrl": "string"
      }
    ],
    "expandedCitations": [
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
    "matchedDocuments": [
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
    "rewrittenQuery": "string",
    "role": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/config/info/test \
  --body '{"cfgType": "string", "configID": "string", "inputModalities": "string", "jsonData": "string", "messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "modelType": "string", "name": "示例名称", "outputModalities": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/update`

**说明**: 更新配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 否 |  配置类型：llm/tts/asr/vad |
| `configID` | string | 否 |  配置唯一标识 |
| `contextLength` | integer | 否 |  上下文长度 (格式: int64) |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `id` | integer | 否 |  配置ID (格式: int64) |
| `inputModalities` | string | 否 |  输入模态（JSON字符串） |
| `isDefault` | boolean | 否 |  是否默认 (格式: boolean) |
| `jsonData` | string | 否 |  配置参数（JSON字符串） |
| `modelType` | string | 否 |  模型级别：small/medium/large/xlarge |
| `name` | string | 否 |  配置名称 |
| `outputModalities` | string | 否 |  输出模态（JSON字符串） |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "configID": "string",
  "contextLength": 1,
  "enabled": true,
  "id": 1,
  "inputModalities": "string",
  "isDefault": true,
  "jsonData": "string",
  "modelType": "string",
  "name": "示例名称",
  "outputModalities": "string",
  "tenantCode": "string"
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
ur api /api/v1/ai/config/info/update \
  --body '{"cfgType": "string", "configID": "string", "contextLength": 1, "enabled": true, "id": 1, "inputModalities": "string", "isDefault": true, "jsonData": "string", "modelType": "string", "name": "示例名称", "outputModalities": "string", "tenantCode": "string"}'
```
