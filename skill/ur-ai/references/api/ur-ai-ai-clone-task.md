# ur-ai ai/clone-task

创建clone任务 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/clone-task/create` | 创建clone任务 | admin |
| POST | `/api/v1/ai/clone-task/delete` | 删除clone任务 | admin |
| POST | `/api/v1/ai/clone-task/get-list` | 获取clone任务列表 | admin |
| POST | `/api/v1/ai/clone-task/get-one` | 获取clone任务详情 | admin |
| POST | `/api/v1/ai/clone-task/get-runs` | 获取clone任务执行日志 | admin |
| POST | `/api/v1/ai/clone-task/get-subscribe-topics` | 获取订阅主题列表 | admin |
| POST | `/api/v1/ai/clone-task/run` | 手动触发clone任务 | admin |
| POST | `/api/v1/ai/clone-task/start` | 启用clone任务 | admin |
| POST | `/api/v1/ai/clone-task/stop` | 禁用clone任务 | admin |
| POST | `/api/v1/ai/clone-task/update` | 更新clone任务 | admin |

## 详细说明

### POST `/api/v1/ai/clone-task/create`

**说明**: 创建clone任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneGroupID` | string | 是 |  CloneGroup ID |
| `cloneID` | string | 是 |  绑定的Clone ID |
| `desc` | string | 否 |  任务描述 |
| `name` | string | 是 |  任务名称 |
| `prompt` | string | 否 |  AI运行的提示词(ai_run时必填) |
| `repeatTimes` | integer | 否 |  最大执行次数，0=无限 (格式: int64) |
| `scheduleDailyTimeSec` | integer | 否 |  daily:当天秒数(0-86399) (格式: int64) |
| `scheduleInterval` | integer | 否 |  interval:间隔分钟数 (格式: int64) |
| `scheduleTime` | string | 否 |  once:ISO时间 "2026-06-01T08:00:00" |
| `scheduleType` | string | 是 |  调度类型 none/once/interval/daily/weekly |
| `scheduleWeeklyDay` | integer | 否 |  weekly:0=周日1=周一...6=周六 (格式: int64) |
| `scheduleWeeklyTimeSec` | integer | 否 |  weekly:当天秒数(0-86399) (格式: int64) |
| `subscribePayloadFilter` | string | 否 |  订阅消息过滤条件JSON |
| `subscribeTopic` | string | 否 |  订阅消息触发时选择的主题编码 |
| `taskType` | string | 是 |  任务类型 ai_run/tool_run |
| `timeoutSeconds` | integer | 否 |  超时时间，默认300 (格式: int64) |
| `toolID` | integer | 否 |  AiTool ID(tool_run时必填) (格式: int64) |
| `toolInputs` | string | 否 |  工具输入JSON(tool_run时选填) |
| `triggerType` | string | 否 |  触发类型 schedule/subscribe_message，默认schedule |

**请求示例**:
```json
{
  "cloneGroupID": "string",
  "cloneID": "string",
  "desc": "string",
  "name": "示例名称",
  "prompt": "string",
  "repeatTimes": 1,
  "scheduleDailyTimeSec": 1,
  "scheduleInterval": 1,
  "scheduleTime": "2026-01-01T00:00:00Z",
  "scheduleType": "string",
  "scheduleWeeklyDay": 1,
  "scheduleWeeklyTimeSec": 1,
  "subscribePayloadFilter": "string",
  "subscribeTopic": "string",
  "taskType": "string",
  "timeoutSeconds": 1,
  "toolID": 1,
  "toolInputs": "string",
  "triggerType": "string"
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
ur api /api/v1/ai/clone-task/create \
  --body '{"cloneGroupID": "string", "cloneID": "string", "desc": "string", "name": "示例名称", "prompt": "string", "repeatTimes": 1, "scheduleDailyTimeSec": 1, "scheduleInterval": 1, "scheduleTime": "2026-01-01T00:00:00Z", "scheduleType": "string", "scheduleWeeklyDay": 1, "scheduleWeeklyTimeSec": 1, "subscribePayloadFilter": "string", "subscribeTopic": "string", "taskType": "string", "timeoutSeconds": 1, "toolID": 1, "toolInputs": "string", "triggerType": "string"}'
```

### POST `/api/v1/ai/clone-task/delete`

**说明**: 删除clone任务

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
ur api /api/v1/ai/clone-task/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone-task/get-list`

**说明**: 获取clone任务列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 |  按Agent过滤 (格式: int64) |
| `cloneGroupID` | integer | 否 |  按CloneGroup过滤 (格式: int64) |
| `cloneID` | integer | 否 |  按Clone过滤 (格式: int64) |
| `enabled` | integer | 否 |  按启用状态过滤 (格式: int64) |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `scheduleType` | string | 否 |  按调度类型过滤 |
| `taskType` | string | 否 |  按任务类型过滤 |
| `triggerType` | string | 否 |  按触发类型过滤 |

**请求示例**:
```json
{
  "agentID": 1,
  "cloneGroupID": 1,
  "cloneID": 1,
  "enabled": 1,
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
  "scheduleType": "string",
  "taskType": "string",
  "triggerType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "agentID": 1,
        "cloneGroupID": "string",
        "cloneID": "string",
        "createdTime": 1,
        "desc": "string",
        "enabled": "string",
        "id": 1,
        "lastError": "string",
        "lastRunAt": 1,
        "lastRunID": "string",
        "lastStatus": "string",
        "name": "示例名称",
        "nextRunAt": 1,
        "prompt": "string",
        "repeatCompleted": 1,
        "repeatTimes": 1,
        "scheduleDailyTimeSec": 1,
        "scheduleInterval": 1,
        "scheduleTime": "2026-01-01T00:00:00Z",
        "scheduleType": "string",
        "scheduleWeeklyDay": 1,
        "scheduleWeeklyTimeSec": 1,
        "sessionID": "string",
        "state": "string",
        "subscribePayloadFilter": "string",
        "subscribeTopic": "string",
        "taskType": "string",
        "tenantCode": "string",
        "timeoutSeconds": 1,
        "toolCodeSnapshot": "string",
        "toolID": 1,
        "toolInputsTemplate": "string",
        "toolManifestSnapshot": "string",
        "toolVersionSnapshot": 1,
        "triggerType": "string",
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
ur api /api/v1/ai/clone-task/get-list \
  --body '{"agentID": 1, "cloneGroupID": 1, "cloneID": 1, "enabled": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "scheduleType": "string", "taskType": "string", "triggerType": "string"}'
```

### POST `/api/v1/ai/clone-task/get-one`

**说明**: 获取clone任务详情

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
    "agentID": 1,
    "cloneGroupID": "string",
    "cloneID": "string",
    "createdTime": 1,
    "desc": "string",
    "enabled": "string",
    "id": 1,
    "lastError": "string",
    "lastRunAt": 1,
    "lastRunID": "string",
    "lastStatus": "string",
    "name": "示例名称",
    "nextRunAt": 1,
    "prompt": "string",
    "repeatCompleted": 1,
    "repeatTimes": 1,
    "scheduleDailyTimeSec": 1,
    "scheduleInterval": 1,
    "scheduleTime": "2026-01-01T00:00:00Z",
    "scheduleType": "string",
    "scheduleWeeklyDay": 1,
    "scheduleWeeklyTimeSec": 1,
    "sessionID": "string",
    "state": "string",
    "subscribePayloadFilter": "string",
    "subscribeTopic": "string",
    "taskType": "string",
    "tenantCode": "string",
    "timeoutSeconds": 1,
    "toolCodeSnapshot": "string",
    "toolID": 1,
    "toolInputsTemplate": "string",
    "toolManifestSnapshot": "string",
    "toolVersionSnapshot": 1,
    "triggerType": "string",
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone-task/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone-task/get-runs`

**说明**: 获取clone任务执行日志

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `taskID` | string | 是 |  任务ID |

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
  "taskID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": 1,
        "durationMs": 1,
        "endTime": 1,
        "errorMsg": "string",
        "id": 1,
        "logs": "string",
        "output": "string",
        "runID": "string",
        "sessionID": "string",
        "startTime": 1,
        "status": "string",
        "taskID": 1,
        "toolID": 1,
        "triggerType": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone-task/get-runs \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "taskID": "string"}'
```

### POST `/api/v1/ai/clone-task/get-subscribe-topics`

**说明**: 获取订阅主题列表

**权限**: admin

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": "string",
        "description": "示例名称"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone-task/get-subscribe-topics \
  --body '{}'
```

### POST `/api/v1/ai/clone-task/run`

**说明**: 手动触发clone任务

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
ur api /api/v1/ai/clone-task/run \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone-task/start`

**说明**: 启用clone任务

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
ur api /api/v1/ai/clone-task/start \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone-task/stop`

**说明**: 禁用clone任务

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
ur api /api/v1/ai/clone-task/stop \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone-task/update`

**说明**: 更新clone任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  任务描述 |
| `id` | string | 是 |  任务ID |
| `name` | string | 否 |  任务名称 |
| `prompt` | string | 否 |  AI运行的提示词 |
| `repeatTimes` | integer | 否 |  最大执行次数 (格式: int64) |
| `scheduleDailyTimeSec` | integer | 否 |  daily:当天秒数 (格式: int64) |
| `scheduleInterval` | integer | 否 |  interval:间隔分钟数 (格式: int64) |
| `scheduleTime` | string | 否 |  once:ISO时间 |
| `scheduleType` | string | 否 |  调度类型 |
| `scheduleWeeklyDay` | integer | 否 |  weekly:周几 (格式: int64) |
| `scheduleWeeklyTimeSec` | integer | 否 |  weekly:当天秒数 (格式: int64) |
| `subscribePayloadFilter` | string | 否 |  订阅消息过滤条件JSON |
| `subscribeTopic` | string | 否 |  订阅消息触发时选择的主题编码 |
| `timeoutSeconds` | integer | 否 |  超时时间 (格式: int64) |
| `toolInputs` | string | 否 |  工具输入JSON |
| `triggerType` | string | 否 |  触发类型 schedule/subscribe_message |

**请求示例**:
```json
{
  "desc": "string",
  "id": "string",
  "name": "示例名称",
  "prompt": "string",
  "repeatTimes": 1,
  "scheduleDailyTimeSec": 1,
  "scheduleInterval": 1,
  "scheduleTime": "2026-01-01T00:00:00Z",
  "scheduleType": "string",
  "scheduleWeeklyDay": 1,
  "scheduleWeeklyTimeSec": 1,
  "subscribePayloadFilter": "string",
  "subscribeTopic": "string",
  "timeoutSeconds": 1,
  "toolInputs": "string",
  "triggerType": "string"
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
ur api /api/v1/ai/clone-task/update \
  --body '{"desc": "string", "id": "string", "name": "示例名称", "prompt": "string", "repeatTimes": 1, "scheduleDailyTimeSec": 1, "scheduleInterval": 1, "scheduleTime": "2026-01-01T00:00:00Z", "scheduleType": "string", "scheduleWeeklyDay": 1, "scheduleWeeklyTimeSec": 1, "subscribePayloadFilter": "string", "subscribeTopic": "string", "timeoutSeconds": 1, "toolInputs": "string", "triggerType": "string"}'
```
