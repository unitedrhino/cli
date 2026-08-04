# ur-ai ai/clone/memory

获取归档记录列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/clone/memory/archives/get-list` | 获取归档记录列表 | admin |
| POST | `/api/v1/ai/clone/memory/compact` | 压缩记忆 | admin |
| POST | `/api/v1/ai/clone/memory/create` | 添加记忆 | admin |
| POST | `/api/v1/ai/clone/memory/delete` | 删除记忆 | admin |
| POST | `/api/v1/ai/clone/memory/dream-runs/get-list` | 获取 Dream 运行记录列表 | admin |
| POST | `/api/v1/ai/clone/memory/dream/run` | 手动触发 Dream 整理 | admin |
| POST | `/api/v1/ai/clone/memory/entities/get-list` | 获取记忆实体列表 | admin |
| POST | `/api/v1/ai/clone/memory/profile/get-one` | 获取记忆画像 | admin |
| POST | `/api/v1/ai/clone/memory/records/get-list` | 获取记忆记录列表 | admin |
| POST | `/api/v1/ai/clone/memory/search` | 搜索记忆 | admin |
| POST | `/api/v1/ai/clone/memory/traces/get-list` | 获取召回轨迹列表 | admin |

## 详细说明

### POST `/api/v1/ai/clone/memory/archives/get-list`

**说明**: 获取归档记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `archiveReason` | string | 否 |  归档原因过滤 |
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "archiveReason": "string",
  "cloneID": 1,
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
        "archiveReason": "string",
        "cloneID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "recordID": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/archives/get-list \
  --body '{"archiveReason": "string", "cloneID": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/ai/clone/memory/compact`

**说明**: 压缩记忆

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
ur api /api/v1/ai/clone/memory/compact \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone/memory/create`

**说明**: 添加记忆

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `content` | string | 是 |  记忆内容 |
| `importance` | number | 否 |  重要性评分 (格式: double) |
| `keywords` | string | 否 |  关键词（JSON字符串） |
| `memoryType` | string | 是 |  记忆类型 |

**请求示例**:
```json
{
  "cloneID": 1,
  "content": "string",
  "importance": 1,
  "keywords": "string",
  "memoryType": "string"
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
ur api /api/v1/ai/clone/memory/create \
  --body '{"cloneID": 1, "content": "string", "importance": 1, "keywords": "string", "memoryType": "string"}'
```

### POST `/api/v1/ai/clone/memory/delete`

**说明**: 删除记忆

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `memoryID` | integer | 是 |  记忆ID (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "memoryID": 1
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
ur api /api/v1/ai/clone/memory/delete \
  --body '{"cloneID": 1, "memoryID": 1}'
```

### POST `/api/v1/ai/clone/memory/dream-runs/get-list`

**说明**: 获取 Dream 运行记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态过滤 running/success/skipped/failed |

**请求示例**:
```json
{
  "cloneID": 1,
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
  "status": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "cloneID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "inputWindowEnd": "string",
        "inputWindowStart": "string",
        "status": "string",
        "summary": "string",
        "triggerType": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/dream-runs/get-list \
  --body '{"cloneID": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": "string"}'
```

### POST `/api/v1/ai/clone/memory/dream/run`

**说明**: 手动触发 Dream 整理

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `triggerType` | string | 否 |  触发类型 manual/repair |

**请求示例**:
```json
{
  "cloneID": 1,
  "triggerType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "archivedCount": "string",
    "newRecordID": "string",
    "runID": "string",
    "status": "string",
    "summary": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/dream/run \
  --body '{"cloneID": 1, "triggerType": "string"}'
```

### POST `/api/v1/ai/clone/memory/entities/get-list`

**说明**: 获取记忆实体列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `entityKey` | string | 否 |  实体键过滤 |
| `entityType` | string | 否 |  实体类型过滤 |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "entityKey": "string",
  "entityType": "string",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
        "confidence": 1,
        "entityKey": "string",
        "entityName": "string",
        "entityType": "string",
        "recordCount": "string",
        "role": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/entities/get-list \
  --body '{"cloneID": 1, "entityKey": "string", "entityType": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/ai/clone/memory/profile/get-one`

**说明**: 获取记忆画像

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "ownerID": "string",
    "ownerType": "string",
    "profileText": "string",
    "updatedTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/profile/get-one \
  --body '{"cloneID": 1}'
```

### POST `/api/v1/ai/clone/memory/records/get-list`

**说明**: 获取记忆记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `endTime` | string | 否 |  结束时间戳（秒） |
| `memoryKind` | string | 否 |  记忆类别 |
| `memoryLevel` | string | 否 |  记忆层级 l1/l2/profile |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `query` | string | 否 |  搜索关键词（摘要/内容模糊匹配） |
| `startTime` | string | 否 |  开始时间戳（秒） |
| `status` | string | 否 |  状态 1-活跃 2-已衰减 3-已归档 |

**请求示例**:
```json
{
  "cloneID": 1,
  "endTime": "2026-01-01T00:00:00Z",
  "memoryKind": "string",
  "memoryLevel": "string",
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
  "query": "string",
  "startTime": "2026-01-01T00:00:00Z",
  "status": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "cloneID": 1,
        "content": "string",
        "id": "string",
        "importance": 1,
        "keywords": "string",
        "memoryKind": "string",
        "memoryLevel": "string",
        "sessionID": "string",
        "status": "string",
        "summary": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/records/get-list \
  --body '{"cloneID": 1, "endTime": "2026-01-01T00:00:00Z", "memoryKind": "string", "memoryLevel": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "query": "string", "startTime": "2026-01-01T00:00:00Z", "status": "string"}'
```

### POST `/api/v1/ai/clone/memory/search`

**说明**: 搜索记忆

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `memoryType` | string | 否 |  记忆类型过滤 |
| `query` | string | 是 |  搜索关键词 |
| `topK` | integer | 否 |  返回条数（默认10） (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "memoryType": "string",
  "query": "string",
  "topK": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "accessCount": 1,
        "cloneID": 1,
        "content": "string",
        "createdTime": 1,
        "id": 1,
        "importance": 1,
        "keywords": "string",
        "memoryType": "string",
        "sessionID": "string",
        "status": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/search \
  --body '{"cloneID": 1, "memoryType": "string", "query": "string", "topK": 1}'
```

### POST `/api/v1/ai/clone/memory/traces/get-list`

**说明**: 获取召回轨迹列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `query` | string | 否 |  查询内容模糊搜索 |

**请求示例**:
```json
{
  "cloneID": 1,
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
  "query": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "enoughAt": "string",
        "path": "string",
        "query": "string",
        "resultSummary": "string",
        "traceID": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/traces/get-list \
  --body '{"cloneID": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "query": "string"}'
```
