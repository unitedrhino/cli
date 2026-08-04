# ur-ai ai/memory/debug

调试-归档记录列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/memory/debug/archives/get-list` | 调试-归档记录列表 | platform |
| POST | `/api/v1/ai/memory/debug/dream-runs/get-list` | 调试-Dream运行记录列表 | platform |
| POST | `/api/v1/ai/memory/debug/overview` | 记忆调试总览 | platform |
| POST | `/api/v1/ai/memory/debug/pipeline-state/get-list` | 调试-流水线状态列表 | platform |
| POST | `/api/v1/ai/memory/debug/traces/get-list` | 调试-召回轨迹列表 | platform |

## 详细说明

### POST `/api/v1/ai/memory/debug/archives/get-list`

**说明**: 调试-归档记录列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `archiveReason` | string | 否 |  归档原因过滤 |
| `cloneID` | string | 否 |  分身ID（不传则查全部） |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "archiveReason": "string",
  "cloneID": "string",
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
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "ownerID": "string",
        "ownerType": "string",
        "recordID": "string",
        "tenantCode": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/memory/debug/archives/get-list \
  --body '{"archiveReason": "string", "cloneID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/ai/memory/debug/dream-runs/get-list`

**说明**: 调试-Dream运行记录列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | string | 否 |  分身ID（不传则查全部） |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态过滤 running/success/skipped/failed |

**请求示例**:
```json
{
  "cloneID": "string",
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
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "inputWindowEnd": "string",
        "inputWindowStart": "string",
        "ownerID": "string",
        "ownerType": "string",
        "status": "string",
        "summary": "string",
        "tenantCode": "string",
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
ur api /api/v1/ai/memory/debug/dream-runs/get-list \
  --body '{"cloneID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": "string"}'
```

### POST `/api/v1/ai/memory/debug/overview`

**说明**: 记忆调试总览

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "archiveCount": "string",
    "dreamRunCount": "string",
    "l1Count": "string",
    "l2Count": "string",
    "lastErrorTime": "2026-01-01T00:00:00Z",
    "lastWriteTime": "2026-01-01T00:00:00Z",
    "pipelineStates": [
      {
        "id": "string",
        "ownerID": "string",
        "ownerType": "string",
        "stateKey": "string",
        "stateValue": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "profileCount": "string",
    "recordCount": "string",
    "traceCount": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/memory/debug/overview \
  --body '{}'
```

### POST `/api/v1/ai/memory/debug/pipeline-state/get-list`

**说明**: 调试-流水线状态列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | string | 否 |  分身ID（不传则查全部） |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "cloneID": "string",
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
        "id": "string",
        "ownerID": "string",
        "ownerType": "string",
        "stateKey": "string",
        "stateValue": "string",
        "tenantCode": "string",
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
ur api /api/v1/ai/memory/debug/pipeline-state/get-list \
  --body '{"cloneID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/ai/memory/debug/traces/get-list`

**说明**: 调试-召回轨迹列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | string | 否 |  分身ID（不传则查全部） |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `query` | string | 否 |  查询内容模糊搜索 |

**请求示例**:
```json
{
  "cloneID": "string",
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
        "ownerID": "string",
        "path": "string",
        "query": "string",
        "resultSummary": "string",
        "tenantCode": "string",
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
ur api /api/v1/ai/memory/debug/traces/get-list \
  --body '{"cloneID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "query": "string"}'
```
