# ur-system system/job/task

取消执行任务 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/job/task/cancel` | 取消执行任务 | platform |
| POST | `/api/v1/system/job/task/group/create` | 新建任务分组 | platform |
| POST | `/api/v1/system/job/task/group/delete` | 删除任务分组 | platform |
| POST | `/api/v1/system/job/task/group/get-list` | 获取任务分组列表 | platform |
| POST | `/api/v1/system/job/task/group/get-one` | 获取任务分组详情 | platform |
| POST | `/api/v1/system/job/task/group/update` | 更新任务分组 | platform |
| POST | `/api/v1/system/job/task/info/create` | 创建任务 | platform |
| POST | `/api/v1/system/job/task/info/delete` | 删除任务 | platform |
| POST | `/api/v1/system/job/task/info/get-list` | 获取任务列表 | platform |
| POST | `/api/v1/system/job/task/info/get-one` | 获取任务详情 | platform |
| POST | `/api/v1/system/job/task/info/start` | 启动任务 | platform |
| POST | `/api/v1/system/job/task/info/stop` | 停止任务 | platform |
| POST | `/api/v1/system/job/task/info/update` | 更新任务 | platform |
| POST | `/api/v1/system/job/task/send` | 发送延时请求 | platform |

## 详细说明

### POST `/api/v1/system/job/task/cancel`

**说明**: 取消执行任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `taskID` | string | 是 | 运行的taskID |

**请求示例**:
```json
{
  "taskID": "string"
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
ur api /api/v1/system/job/task/cancel \
  --body '{"taskID": "string"}'
```

### POST `/api/v1/system/job/task/group/create`

**说明**: 新建任务分组

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  任务组编码 |
| `createdTime` | string | 否 |  创建时间（Unix 毫秒） |
| `desc` | string | 否 |  描述 |
| `id` | string | 否 |  任务组ID |
| `name` | string | 否 |  组名 |
| `status` | integer | 否 |  状态 1:启用 2:禁用 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "name": "string",
  "status": 1
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
ur api /api/v1/system/job/task/group/create \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "name": "string", "status": 1}'
```

### POST `/api/v1/system/job/task/group/delete`

**说明**: 删除任务分组

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string"
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
ur api /api/v1/system/job/task/group/delete \
  --body '{"code": "string"}'
```

### POST `/api/v1/system/job/task/group/get-list`

**说明**: 获取任务分组列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  状态过滤 1:启用 2:禁用 (格式: int64) |

**请求示例**:
```json
{
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
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
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "id": "string",
        "name": "string",
        "status": 1
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
ur api /api/v1/system/job/task/group/get-list \
  --body '{"name": "示例名称", "page": {"page": 1, "pageSize": 1}, "status": 1}'
```

### POST `/api/v1/system/job/task/group/get-one`

**说明**: 获取任务分组详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "name": "string",
    "status": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/job/task/group/get-one \
  --body '{"code": "string"}'
```

### POST `/api/v1/system/job/task/group/update`

**说明**: 更新任务分组

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  任务组编码 |
| `createdTime` | string | 否 |  创建时间（Unix 毫秒） |
| `desc` | string | 否 |  描述 |
| `id` | string | 否 |  任务组ID |
| `name` | string | 否 |  组名 |
| `status` | integer | 否 |  状态 1:启用 2:禁用 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "name": "string",
  "status": 1
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
ur api /api/v1/system/job/task/group/update \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "name": "string", "status": 1}'
```

### POST `/api/v1/system/job/task/info/create`

**说明**: 创建任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  任务编码 |
| `createdTime` | string | 否 |  创建时间（Unix 毫秒） |
| `cronExpr` | string | 否 |  Cron 表达式（type=1） |
| `desc` | string | 否 |  描述 |
| `executeAt` | string | 否 |  执行时间 Unix 毫秒（type=2） |
| `groupCode` | string | 是 |  关联分组编码 |
| `id` | string | 否 |  任务ID |
| `lastRunTime` | string | 否 |  上次执行时间 Unix 毫秒 |
| `name` | string | 否 |  任务名称 |
| `params` | string | 否 |  任务参数（JSON 字符串） |
| `priority` | integer | 否 |  优先级 1-10 (格式: int64) |
| `retriedCount` | integer | 否 |  已重试次数 (格式: int64) |
| `retryPolicy` | object | 否 |  |
| `retryPolicy.backoff` | string | 否 |  退避策略：fixed / exponential |
| `retryPolicy.maxRetry` | integer | 否 |  最大重试次数，默认 3 (格式: int64) |
| `retryPolicy.retryDelay` | integer | 否 |  重试间隔基数（秒），默认 10 (格式: int64) |
| `status` | integer | 否 |  状态 1:WaitRun 2:Running 3:WaitStop 4:Stopped 5:WaitDelete (格式: int64) |
| `type` | integer | 否 |  类型 1:定时(Cron) 2:延时(Delay) (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "cronExpr": "string",
  "desc": "string",
  "executeAt": "2026-01-01T00:00:00Z",
  "groupCode": "string",
  "id": "string",
  "lastRunTime": "2026-01-01T00:00:00Z",
  "name": "示例名称",
  "params": "string",
  "priority": 1,
  "retriedCount": 1,
  "retryPolicy": {
    "backoff": "string",
    "maxRetry": 1,
    "retryDelay": 1
  },
  "status": 1,
  "type": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "cronExpr": "string",
    "desc": "string",
    "executeAt": "2026-01-01T00:00:00Z",
    "groupCode": "string",
    "id": "string",
    "lastRunTime": "2026-01-01T00:00:00Z",
    "name": "示例名称",
    "params": "string",
    "priority": 1,
    "retriedCount": 1,
    "retryPolicy": {
      "backoff": "string",
      "maxRetry": 1,
      "retryDelay": 1
    },
    "status": 1,
    "type": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/job/task/info/create \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "cronExpr": "string", "desc": "string", "executeAt": "2026-01-01T00:00:00Z", "groupCode": "string", "id": "string", "lastRunTime": "2026-01-01T00:00:00Z", "name": "示例名称", "params": "string", "priority": 1, "retriedCount": 1, "retryPolicy": {"backoff": "string", "maxRetry": 1, "retryDelay": 1}, "status": 1, "type": 1}'
```

### POST `/api/v1/system/job/task/info/delete`

**说明**: 删除任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
ur api /api/v1/system/job/task/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/info/get-list`

**说明**: 获取任务列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupCode` | string | 是 |  关联分组编码 |
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  按状态过滤 (格式: int64) |
| `type` | integer | 否 |  按类型过滤 1:定时 2:延时 (格式: int64) |

**请求示例**:
```json
{
  "groupCode": "string",
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": 1,
  "type": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "cronExpr": "string",
        "desc": "string",
        "executeAt": "2026-01-01T00:00:00Z",
        "groupCode": "string",
        "id": "string",
        "lastRunTime": "2026-01-01T00:00:00Z",
        "name": "示例名称",
        "params": "string",
        "priority": 1,
        "retriedCount": 1,
        "retryPolicy": {
          "backoff": "string",
          "maxRetry": 1,
          "retryDelay": 1
        },
        "status": 1,
        "type": 1
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
ur api /api/v1/system/job/task/info/get-list \
  --body '{"groupCode": "string", "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "status": 1, "type": 1}'
```

### POST `/api/v1/system/job/task/info/get-one`

**说明**: 获取任务详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
    "code": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "cronExpr": "string",
    "desc": "string",
    "executeAt": "2026-01-01T00:00:00Z",
    "groupCode": "string",
    "id": "string",
    "lastRunTime": "2026-01-01T00:00:00Z",
    "name": "示例名称",
    "params": "string",
    "priority": 1,
    "retriedCount": 1,
    "retryPolicy": {
      "backoff": "string",
      "maxRetry": 1,
      "retryDelay": 1
    },
    "status": 1,
    "type": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/job/task/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/info/start`

**说明**: 启动任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
ur api /api/v1/system/job/task/info/start \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/info/stop`

**说明**: 停止任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
ur api /api/v1/system/job/task/info/stop \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/info/update`

**说明**: 更新任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  任务编码 |
| `createdTime` | string | 否 |  创建时间（Unix 毫秒） |
| `cronExpr` | string | 否 |  Cron 表达式（type=1） |
| `desc` | string | 否 |  描述 |
| `executeAt` | string | 否 |  执行时间 Unix 毫秒（type=2） |
| `groupCode` | string | 是 |  关联分组编码 |
| `id` | string | 否 |  任务ID |
| `lastRunTime` | string | 否 |  上次执行时间 Unix 毫秒 |
| `name` | string | 否 |  任务名称 |
| `params` | string | 否 |  任务参数（JSON 字符串） |
| `priority` | integer | 否 |  优先级 1-10 (格式: int64) |
| `retriedCount` | integer | 否 |  已重试次数 (格式: int64) |
| `retryPolicy` | object | 否 |  |
| `retryPolicy.backoff` | string | 否 |  退避策略：fixed / exponential |
| `retryPolicy.maxRetry` | integer | 否 |  最大重试次数，默认 3 (格式: int64) |
| `retryPolicy.retryDelay` | integer | 否 |  重试间隔基数（秒），默认 10 (格式: int64) |
| `status` | integer | 否 |  状态 1:WaitRun 2:Running 3:WaitStop 4:Stopped 5:WaitDelete (格式: int64) |
| `type` | integer | 否 |  类型 1:定时(Cron) 2:延时(Delay) (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "cronExpr": "string",
  "desc": "string",
  "executeAt": "2026-01-01T00:00:00Z",
  "groupCode": "string",
  "id": "string",
  "lastRunTime": "2026-01-01T00:00:00Z",
  "name": "示例名称",
  "params": "string",
  "priority": 1,
  "retriedCount": 1,
  "retryPolicy": {
    "backoff": "string",
    "maxRetry": 1,
    "retryDelay": 1
  },
  "status": 1,
  "type": 1
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
ur api /api/v1/system/job/task/info/update \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "cronExpr": "string", "desc": "string", "executeAt": "2026-01-01T00:00:00Z", "groupCode": "string", "id": "string", "lastRunTime": "2026-01-01T00:00:00Z", "name": "示例名称", "params": "string", "priority": 1, "retriedCount": 1, "retryPolicy": {"backoff": "string", "maxRetry": 1, "retryDelay": 1}, "status": 1, "type": 1}'
```

### POST `/api/v1/system/job/task/send`

**说明**: 发送延时请求

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 | 任务code |
| `groupCode` | string | 是 | 组需要提前创建好 |
| `option` | object | 否 |  |
| `option.deadline` | integer | 否 | 截止时间  秒时间戳 (格式: int64) |
| `option.priority` | integer | 否 | 优先级: 6:critical 最高优先级  3: default 普通优先级 1:low 低优先级以下两个参数优先使用ProcessIn (格式: int64) |
| `option.processAt` | integer | 否 |  固定时间发 秒时间戳 (格式: int64) |
| `option.processIn` | integer | 否 | 多久之后发 秒数 (格式: int64) |
| `option.taskID` | string | 否 | 指定运行的taskID |
| `option.timeout` | integer | 否 | 超时时间 优先使用 秒数 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "groupCode": "string",
  "option": {
    "deadline": 1,
    "priority": 1,
    "processAt": 1,
    "processIn": 1,
    "taskID": "string",
    "timeout": 1
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "taskID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/job/task/send \
  --body '{"code": "string", "groupCode": "string", "option": {"deadline": 1, "priority": 1, "processAt": 1, "processIn": 1, "taskID": "string", "timeout": 1}}'
```
