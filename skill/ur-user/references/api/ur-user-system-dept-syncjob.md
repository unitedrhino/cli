# ur-user system/dept/syncJob

添加同步任务 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/dept/sync-job/create` | 添加同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/delete` | 删除同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/execute` | 执行同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/get-list` | 获取同步任务列表 | admin |
| POST | `/api/v1/system/dept/sync-job/get-one` | 获取同步任务详情 | admin |
| POST | `/api/v1/system/dept/sync-job/update` | 更新同步任务 | admin |

## 详细说明

### POST `/api/v1/system/dept/sync-job/create`

**说明**: 添加同步任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间 |
| `direction` | integer | 否 |  同步的方向,1上游同步到联犀(默认),2联犀同步到下游 (格式: int64) |
| `fieldMap` | object | 否 | 用户字段映射,左边是联犀的字段,右边是第三方的,不填写就是全量映射 |
| `id` | string | 否 |  编号 |
| `isAddSync` | integer | 否 | 新增人员自动同步,默认为1 (格式: int64) |
| `syncDeptIDs` | array[string] | 否 | 同步的第三方部门id列表,不填为同步全部 |
| `syncMode` | integer | 否 | 同步模式: 1:手动(默认) 2: 定时同步(半小时) 3: 实时同步 (格式: int64) |
| `thirdConfig` | object | 否 |  |
| `thirdConfig.appID` | string | 否 |  |
| `thirdConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `thirdConfig.appSecret` | string | 否 |  |
| `thirdType` | string | 否 | 同步的类型 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "direction": 1,
  "fieldMap": {},
  "id": "string",
  "isAddSync": 1,
  "syncDeptIDs": [
    "string"
  ],
  "syncMode": 1,
  "thirdConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "thirdType": "string"
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
ur api /api/v1/system/dept/sync-job/create \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "direction": 1, "fieldMap": {}, "id": "string", "isAddSync": 1, "syncDeptIDs": ["string"], "syncMode": 1, "thirdConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "thirdType": "string"}'
```

### POST `/api/v1/system/dept/sync-job/delete`

**说明**: 删除同步任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
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
ur api /api/v1/system/dept/sync-job/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/dept/sync-job/execute`

**说明**: 执行同步任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `jobID` | string | 否 |  编号 |

**请求示例**:
```json
{
  "jobID": "string"
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
ur api /api/v1/system/dept/sync-job/execute \
  --body '{"jobID": "string"}'
```

### POST `/api/v1/system/dept/sync-job/get-list`

**说明**: 获取同步任务列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `direction` | integer | 否 |  同步的方向,1上游同步到联犀(默认),2联犀同步到下游 (格式: int64) |
| `name` | string | 否 |  名称 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 | 父节点 |
| `status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |

**请求示例**:
```json
{
  "direction": 1,
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "parentID": "string",
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
        "createdTime": "2026-01-01T00:00:00Z",
        "direction": 1,
        "fieldMap": {},
        "id": "string",
        "isAddSync": 1,
        "syncDeptIDs": [
          "string"
        ],
        "syncMode": 1,
        "thirdConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "thirdType": "string"
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
ur api /api/v1/system/dept/sync-job/get-list \
  --body '{"direction": 1, "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "parentID": "string", "status": 1}'
```

### POST `/api/v1/system/dept/sync-job/get-one`

**说明**: 获取同步任务详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  编号 |
| `withChildren` | boolean | 否 | 是否返回子级 (格式: boolean) |
| `withFather` | boolean | 否 | 是否返回父级 (格式: boolean) |

**请求示例**:
```json
{
  "id": "string",
  "withChildren": true,
  "withFather": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdTime": "2026-01-01T00:00:00Z",
    "direction": 1,
    "fieldMap": {},
    "id": "string",
    "isAddSync": 1,
    "syncDeptIDs": [
      "string"
    ],
    "syncMode": 1,
    "thirdConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "thirdType": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dept/sync-job/get-one \
  --body '{"id": "string", "withChildren": true, "withFather": true}'
```

### POST `/api/v1/system/dept/sync-job/update`

**说明**: 更新同步任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间 |
| `direction` | integer | 否 |  同步的方向,1上游同步到联犀(默认),2联犀同步到下游 (格式: int64) |
| `fieldMap` | object | 否 | 用户字段映射,左边是联犀的字段,右边是第三方的,不填写就是全量映射 |
| `id` | string | 否 |  编号 |
| `isAddSync` | integer | 否 | 新增人员自动同步,默认为1 (格式: int64) |
| `syncDeptIDs` | array[string] | 否 | 同步的第三方部门id列表,不填为同步全部 |
| `syncMode` | integer | 否 | 同步模式: 1:手动(默认) 2: 定时同步(半小时) 3: 实时同步 (格式: int64) |
| `thirdConfig` | object | 否 |  |
| `thirdConfig.appID` | string | 否 |  |
| `thirdConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `thirdConfig.appSecret` | string | 否 |  |
| `thirdType` | string | 否 | 同步的类型 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "direction": 1,
  "fieldMap": {},
  "id": "string",
  "isAddSync": 1,
  "syncDeptIDs": [
    "string"
  ],
  "syncMode": 1,
  "thirdConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "thirdType": "string"
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
ur api /api/v1/system/dept/sync-job/update \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "direction": 1, "fieldMap": {}, "id": "string", "isAddSync": 1, "syncDeptIDs": ["string"], "syncMode": 1, "thirdConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "thirdType": "string"}'
```
