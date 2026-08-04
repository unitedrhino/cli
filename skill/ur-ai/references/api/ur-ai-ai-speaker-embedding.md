# ur-ai ai/speaker/embedding

注册声纹 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/speaker/embedding/create` | 注册声纹 | admin |
| POST | `/api/v1/ai/speaker/embedding/delete` | 删除声纹 | admin |
| POST | `/api/v1/ai/speaker/embedding/extract` | 提取声纹特征向量 | admin |
| POST | `/api/v1/ai/speaker/embedding/get-list` | 获取声纹列表 | admin |
| POST | `/api/v1/ai/speaker/embedding/get-one` | 获取声纹详情 | admin |
| POST | `/api/v1/ai/speaker/embedding/import` | 批量导入声纹 | admin |
| POST | `/api/v1/ai/speaker/embedding/update` | 更新声纹 | admin |

## 详细说明

### POST `/api/v1/ai/speaker/embedding/create`

**说明**: 注册声纹

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `audioPath` | string | 否 |  原始音频存储路径 |
| `cloneGroupID` | integer | 是 |  所属Clone组ID (格式: int64) |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `embedding` | string | 否 |  声纹特征向量（JSON数组，仅查询时返回） |
| `id` | integer | 否 |  声纹ID (格式: int64) |
| `personaPrompt` | string | 否 |  个性化提示词 |
| `speakerName` | string | 是 |  说话人名称 |
| `status` | string | 否 |  状态 1-启用 2-禁用 |
| `tenantCode` | string | 否 |  租户编码 |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "audioPath": "string",
  "cloneGroupID": 1,
  "createdTime": 1,
  "embedding": "string",
  "id": 1,
  "personaPrompt": "string",
  "speakerName": "示例名称",
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
ur api /api/v1/ai/speaker/embedding/create \
  --body '{"audioPath": "string", "cloneGroupID": 1, "createdTime": 1, "embedding": "string", "id": 1, "personaPrompt": "string", "speakerName": "示例名称", "status": "string", "tenantCode": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/speaker/embedding/delete`

**说明**: 删除声纹

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  声纹ID (格式: int64) |

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
ur api /api/v1/ai/speaker/embedding/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/speaker/embedding/extract`

**说明**: 提取声纹特征向量

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `audioBase64` | string | 是 |  音频数据（Base64编码的PCM或WAV） |
| `sampleRate` | integer | 否 |  采样率（默认16000） (格式: int64) |

**请求示例**:
```json
{
  "audioBase64": "string",
  "sampleRate": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "audioQuality": {
      "checks": [
        {
          "message": "string",
          "name": "示例名称",
          "passed": true,
          "threshold": 1,
          "value": 1
        }
      ],
      "level": "string",
      "message": "string",
      "passed": true,
      "score": 1
    },
    "embedding": [
      1
    ],
    "embeddingQuality": {
      "checks": [
        {
          "message": "string",
          "name": "示例名称",
          "passed": true,
          "threshold": 1,
          "value": 1
        }
      ],
      "level": "string",
      "message": "string",
      "passed": true,
      "score": 1
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/speaker/embedding/extract \
  --body '{"audioBase64": "string", "sampleRate": 1}'
```

### POST `/api/v1/ai/speaker/embedding/get-list`

**说明**: 获取声纹列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneGroupID` | integer | 是 |  按Clone组过滤 (格式: int64) |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `speakerName` | string | 否 |  说话人名称模糊搜索 |
| `status` | integer | 否 |  状态过滤 (格式: int64) |

**请求示例**:
```json
{
  "cloneGroupID": 1,
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
  "speakerName": "示例名称",
  "status": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "audioPath": "string",
        "cloneGroupID": 1,
        "createdTime": 1,
        "embedding": "string",
        "id": 1,
        "personaPrompt": "string",
        "speakerName": "示例名称",
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
ur api /api/v1/ai/speaker/embedding/get-list \
  --body '{"cloneGroupID": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "speakerName": "示例名称", "status": 1}'
```

### POST `/api/v1/ai/speaker/embedding/get-one`

**说明**: 获取声纹详情

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
    "audioPath": "string",
    "cloneGroupID": 1,
    "createdTime": 1,
    "embedding": "string",
    "id": 1,
    "personaPrompt": "string",
    "speakerName": "示例名称",
    "status": "string",
    "tenantCode": "string",
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/speaker/embedding/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/speaker/embedding/import`

**说明**: 批量导入声纹

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneGroupID` | integer | 是 |  Clone组ID (格式: int64) |
| `speakers` | array[SpeakerImportItem] | 是 |  声纹列表 |

**请求示例**:
```json
{
  "cloneGroupID": 1,
  "speakers": [
    {
      "audioBase64": "string",
      "embedding": [
        1
      ],
      "personaPrompt": "string",
      "speakerName": "示例名称"
    }
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "failed": 1,
    "imported": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/speaker/embedding/import \
  --body '{"cloneGroupID": 1, "speakers": [{"audioBase64": "string", "embedding": [1], "personaPrompt": "string", "speakerName": "示例名称"}]}'
```

### POST `/api/v1/ai/speaker/embedding/update`

**说明**: 更新声纹

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `audioPath` | string | 否 |  原始音频存储路径 |
| `cloneGroupID` | integer | 是 |  所属Clone组ID (格式: int64) |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `embedding` | string | 否 |  声纹特征向量（JSON数组，仅查询时返回） |
| `id` | integer | 否 |  声纹ID (格式: int64) |
| `personaPrompt` | string | 否 |  个性化提示词 |
| `speakerName` | string | 是 |  说话人名称 |
| `status` | string | 否 |  状态 1-启用 2-禁用 |
| `tenantCode` | string | 否 |  租户编码 |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "audioPath": "string",
  "cloneGroupID": 1,
  "createdTime": 1,
  "embedding": "string",
  "id": 1,
  "personaPrompt": "string",
  "speakerName": "示例名称",
  "status": "string",
  "tenantCode": "string",
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
ur api /api/v1/ai/speaker/embedding/update \
  --body '{"audioPath": "string", "cloneGroupID": 1, "createdTime": 1, "embedding": "string", "id": 1, "personaPrompt": "string", "speakerName": "示例名称", "status": "string", "tenantCode": "string", "updatedTime": 1}'
```
