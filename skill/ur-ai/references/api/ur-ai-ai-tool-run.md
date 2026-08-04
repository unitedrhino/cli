# ur-ai ai/tool/run

手动运行工具 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/tool/run` | 手动运行工具 | admin |
| POST | `/api/v1/ai/tool/run-history` | 执行历史列表 | admin |
| POST | `/api/v1/ai/tool/run-status` | 查询执行状态 | admin |

## 详细说明

### POST `/api/v1/ai/tool/run`

**说明**: 手动运行工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `inputs` | string | 是 |  输入参数（JSON字符串） |
| `toolID` | string | 是 |  工具ID |

**请求示例**:
```json
{
  "inputs": "string",
  "toolID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "runID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/run \
  --body '{"inputs": "string", "toolID": "string"}'
```

### POST `/api/v1/ai/tool/run-history`

**说明**: 执行历史列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `toolID` | string | 是 |  工具ID |

**请求示例**:
```json
{
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
  "toolID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "durationMs": "string",
        "errorMsg": "string",
        "logs": "string",
        "output": "string",
        "runID": "string",
        "state": "string",
        "status": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/run-history \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "toolID": "string"}'
```

### POST `/api/v1/ai/tool/run-status`

**说明**: 查询执行状态

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `runID` | string | 是 |  运行标识UUID |

**请求示例**:
```json
{
  "runID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "durationMs": "string",
    "errorMsg": "string",
    "logs": "string",
    "output": "string",
    "runID": "string",
    "state": "string",
    "status": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/run-status \
  --body '{"runID": "string"}'
```
