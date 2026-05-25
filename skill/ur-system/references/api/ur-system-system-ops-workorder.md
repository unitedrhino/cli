# ur-system system/ops/workOrder

添加工单 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/ops/work-order/create` | 添加工单 | all |
| POST | `/api/v1/system/ops/work-order/get-list` | 获取工单列表 | all |
| POST | `/api/v1/system/ops/work-order/update` | 更新工单 | all |

## 详细说明

### POST `/api/v1/system/ops/work-order/create`

**说明**: 添加工单

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id |
| `createdTime` | string | 否 | 创建时间 |
| `finishedTime` | integer | 否 | 处理完成时间 (格式: int64) |
| `handleTime` | integer | 否 | 处理时间 (格式: int64) |
| `id` | string | 否 |  |
| `issueDesc` | string | 否 | 问题描述 |
| `number` | string | 否 | 编号 |
| `params` | object | 否 | 参数 json格式 |
| `raiseUserID` | string | 否 | 问题提出的用户 |
| `status` | integer | 否 |  状态（1 待处理 2:处理中 3:已完成） (格式: int64) |
| `type` | string | 否 | 工单类型: deviceMaintenance:设备维修工单 可选: ['[deviceMaintenance]'] |

**请求示例**:
```json
{
  "areaID": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "finishedTime": 1,
  "handleTime": 1,
  "id": "string",
  "issueDesc": "string",
  "number": "string",
  "params": {},
  "raiseUserID": "string",
  "status": 1,
  "type": "[deviceMaintenance]"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/ops/work-order/create \
  --body '{"areaID": "string", "createdTime": "2026-01-01T00:00:00Z", "finishedTime": 1, "handleTime": 1, "id": "string", "issueDesc": "string", "number": "string", "params": {}, "raiseUserID": "string", "status": 1, "type": "[deviceMaintenance]"}'
```

### POST `/api/v1/system/ops/work-order/get-list`

**说明**: 获取工单列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 区域ID过滤 |
| `number` | string | 否 | 编号 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int64 |
| `type` | string | 否 | 工单类型: deviceMaintenance:设备维修工单 |

**请求示例**:
```json
{
  "areaID": "string",
  "number": "string",
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
  "status": 1,
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "finishedTime": 1,
        "handleTime": 1,
        "id": "string",
        "issueDesc": "string",
        "number": "string",
        "params": {},
        "raiseUserID": "string",
        "status": 1,
        "type": "[deviceMaintenance]"
      }
    ],
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/ops/work-order/get-list \
  --body '{"areaID": "string", "number": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "type": "string"}'
```

### POST `/api/v1/system/ops/work-order/update`

**说明**: 更新工单

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id |
| `createdTime` | string | 否 | 创建时间 |
| `finishedTime` | integer | 否 | 处理完成时间 (格式: int64) |
| `handleTime` | integer | 否 | 处理时间 (格式: int64) |
| `id` | string | 否 |  |
| `issueDesc` | string | 否 | 问题描述 |
| `number` | string | 否 | 编号 |
| `params` | object | 否 | 参数 json格式 |
| `raiseUserID` | string | 否 | 问题提出的用户 |
| `status` | integer | 否 |  状态（1 待处理 2:处理中 3:已完成） (格式: int64) |
| `type` | string | 否 | 工单类型: deviceMaintenance:设备维修工单 可选: ['[deviceMaintenance]'] |

**请求示例**:
```json
{
  "areaID": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "finishedTime": 1,
  "handleTime": 1,
  "id": "string",
  "issueDesc": "string",
  "number": "string",
  "params": {},
  "raiseUserID": "string",
  "status": 1,
  "type": "[deviceMaintenance]"
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
ur api /api/v1/system/ops/work-order/update \
  --body '{"areaID": "string", "createdTime": "2026-01-01T00:00:00Z", "finishedTime": 1, "handleTime": 1, "id": "string", "issueDesc": "string", "number": "string", "params": {}, "raiseUserID": "string", "status": 1, "type": "[deviceMaintenance]"}'
```
