# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/info/create` | 新增告警 | admin |
| POST | `/api/v1/things/alarm/info/delete` | 删除告警 | admin |
| POST | `/api/v1/things/alarm/info/get-list` | 获取告警信息列表 | admin |
| POST | `/api/v1/things/alarm/info/get-one` | 获取告警信息 | admin |
| POST | `/api/v1/things/alarm/info/update` | 更新告警 | admin |
| POST | `/api/v1/things/alarm/record/deal` | 处理告警 | admin |
| POST | `/api/v1/things/alarm/record/get-list` | 获取告警记录列表 | admin |
| POST | `/api/v1/things/alarm/scene/batch-create` | 更新告警和场景的关联 | admin |
| POST | `/api/v1/things/alarm/scene/delete` | 删除告警和场景的关联 | admin |
| POST | `/api/v1/things/alarm/scene/get-list` | 获取告警和场景的关联列表 | admin |

## 端点详情

### POST `/api/v1/things/alarm/info/create`

**说明**: 新增告警

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accounts` | array[string] | 否 | 账号 |
| `code` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `level` | integer | 否 | 告警配置级别（1提醒 2一般 3严重 4紧急 5超紧急） (格式: int64) |
| `name` | string | 否 |  |
| `notifies` | array[AlarmNotify] | 否 | 通知 |
| `sceneIDs` | array[integer] | 否 | 绑定的场景列表 |
| `scenes` | array[SceneCore] | 否 | 绑定的场景列表,只读 |
| `status` | integer | 否 | 状态: 1启用 2禁用 (格式: int64) |
| `userIDs` | array[string] | 否 | 指定用户ID |

**请求示例**:
```json
{
  "accounts": [
    "string"
  ],
  "code": "string",
  "createdTime": "string",
  "desc": "string",
  "id": "string",
  "level": 1,
  "name": "string",
  "notifies": [
    {
      "templateID": "string",
      "templateName": "示例名称",
      "type": "string"
    }
  ],
  "sceneIDs": [
    1
  ],
  "scenes": [
    {
      "desc": "string",
      "id": "string",
      "name": "string",
      "status": 1
    }
  ],
  "status": 1,
  "userIDs": [
    "string"
  ]
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
ur api /api/v1/things/alarm/info/create \
  --body '{"accounts": ["string"], "code": "string", "createdTime": "string", "desc": "string", "id": "string", "level": 1, "name": "string", "notifies": [{"templateID": "string", "templateName": "示例名称", "type": "string"}], "sceneIDs": [1], "scenes": [{"desc": "string", "id": "string", "name": "string", "status": 1}], "status": 1, "userIDs": ["string"]}'
```

### POST `/api/v1/things/alarm/info/delete`

**说明**: 删除告警

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
ur api /api/v1/things/alarm/info/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/info/get-list`

**说明**: 获取告警信息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `name` | string | 否 | 告警名模糊查询 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "name": "string",
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
        "accounts": [
          "string"
        ],
        "code": "string",
        "createdTime": "string",
        "desc": "string",
        "id": "string",
        "level": 1,
        "name": "string",
        "notifies": [
          {
            "templateID": "string",
            "templateName": "示例名称",
            "type": "string"
          }
        ],
        "sceneIDs": [
          1
        ],
        "scenes": [
          {
            "desc": "string",
            "id": "string",
            "name": "string",
            "status": 1
          }
        ],
        "status": 1,
        "userIDs": [
          "string"
        ]
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
ur api /api/v1/things/alarm/info/get-list \
  --body '{"code": "string", "name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/things/alarm/info/get-one`

**说明**: 获取告警信息

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
  "data": {
    "accounts": [
      "string"
    ],
    "code": "string",
    "createdTime": "string",
    "desc": "string",
    "id": "string",
    "level": 1,
    "name": "string",
    "notifies": [
      {
        "templateID": "string",
        "templateName": "示例名称",
        "type": "string"
      }
    ],
    "sceneIDs": [
      1
    ],
    "scenes": [
      {
        "desc": "string",
        "id": "string",
        "name": "string",
        "status": 1
      }
    ],
    "status": 1,
    "userIDs": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/info/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/info/update`

**说明**: 更新告警

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accounts` | array[string] | 否 | 账号 |
| `code` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `level` | integer | 否 | 告警配置级别（1提醒 2一般 3严重 4紧急 5超紧急） (格式: int64) |
| `name` | string | 否 |  |
| `notifies` | array[AlarmNotify] | 否 | 通知 |
| `sceneIDs` | array[integer] | 否 | 绑定的场景列表 |
| `scenes` | array[SceneCore] | 否 | 绑定的场景列表,只读 |
| `status` | integer | 否 | 状态: 1启用 2禁用 (格式: int64) |
| `userIDs` | array[string] | 否 | 指定用户ID |

**请求示例**:
```json
{
  "accounts": [
    "string"
  ],
  "code": "string",
  "createdTime": "string",
  "desc": "string",
  "id": "string",
  "level": 1,
  "name": "string",
  "notifies": [
    {
      "templateID": "string",
      "templateName": "示例名称",
      "type": "string"
    }
  ],
  "sceneIDs": [
    1
  ],
  "scenes": [
    {
      "desc": "string",
      "id": "string",
      "name": "string",
      "status": 1
    }
  ],
  "status": 1,
  "userIDs": [
    "string"
  ]
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
ur api /api/v1/things/alarm/info/update \
  --body '{"accounts": ["string"], "code": "string", "createdTime": "string", "desc": "string", "id": "string", "level": 1, "name": "string", "notifies": [{"templateID": "string", "templateName": "示例名称", "type": "string"}], "sceneIDs": [1], "scenes": [{"desc": "string", "id": "string", "name": "string", "status": 1}], "status": 1, "userIDs": ["string"]}'
```

### POST `/api/v1/things/alarm/record/deal`

**说明**: 处理告警

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `handle` | integer | 是 | 告警处理结果 1:无效(告警无效) 2:派单(走工作流) (格式: int64) |
| `id` | string | 是 | 告警记录ID |

**请求示例**:
```json
{
  "handle": 1,
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
ur api /api/v1/things/alarm/record/deal \
  --body '{"handle": 1, "id": "string"}'
```

### POST `/api/v1/things/alarm/record/get-list`

**说明**: 获取告警记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmCode` | string | 否 |  |
| `alarmID` | string | 否 | 告警记录ID |
| `alarmName` | string | 否 |  |
| `areaID` | string | 否 | 区域ID过滤 |
| `areaIDPath` | string | 否 |  |
| `dealStatus` | array[integer] | 否 | 告警记录状态（1告警中 2忽略 3正在处理 4已处理） |
| `deviceAlias` | string | 否 | 触发设备名称 |
| `deviceName` | string | 否 | 触发设备ID |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `timeRange` | object | 否 |  |
| `timeRange.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `timeRange.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |

**请求示例**:
```json
{
  "alarmCode": "string",
  "alarmID": "string",
  "alarmName": "string",
  "areaID": "string",
  "areaIDPath": "string",
  "dealStatus": [
    1
  ],
  "deviceAlias": "示例名称",
  "deviceName": "string",
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
  "timeRange": {
    "end": 1,
    "start": 1
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
        "alarmCount": 1,
        "alarmID": "string",
        "alarmName": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "dealStatus": 1,
        "desc": "string",
        "deviceAlias": "示例名称",
        "deviceName": "string",
        "id": "string",
        "lastAlarm": 1,
        "level": 1,
        "productID": "string",
        "sceneID": "string",
        "sceneName": "示例名称",
        "triggerType": "string"
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
ur api /api/v1/things/alarm/record/get-list \
  --body '{"alarmCode": "string", "alarmID": "string", "alarmName": "string", "areaID": "string", "areaIDPath": "string", "dealStatus": [1], "deviceAlias": "示例名称", "deviceName": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "timeRange": {"end": 1, "start": 1}}'
```

### POST `/api/v1/things/alarm/scene/batch-create`

**说明**: 更新告警和场景的关联

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 是 | 告警配置ID |
| `sceneIDs` | array[integer] | 是 | 场景id |

**请求示例**:
```json
{
  "alarmID": "string",
  "sceneIDs": [
    1
  ]
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
ur api /api/v1/things/alarm/scene/batch-create \
  --body '{"alarmID": "string", "sceneIDs": [1]}'
```

### POST `/api/v1/things/alarm/scene/delete`

**说明**: 删除告警和场景的关联

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 是 | 告警配置ID |
| `sceneID` | string | 是 | 场景id |

**请求示例**:
```json
{
  "alarmID": "string",
  "sceneID": "string"
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
ur api /api/v1/things/alarm/scene/delete \
  --body '{"alarmID": "string", "sceneID": "string"}'
```

### POST `/api/v1/things/alarm/scene/get-list`

**说明**: 获取告警和场景的关联列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 是 | 告警配置ID |

**请求示例**:
```json
{
  "alarmID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "alarmID": "string",
    "scenes": [
      {
        "areaID": "string",
        "avatar": "string",
        "body": "string",
        "createdTime": "string",
        "desc": "string",
        "deviceAlias": "string",
        "deviceMode": "string",
        "deviceName": "string",
        "flowPath": [
          {
            "info": "string",
            "subType": "string",
            "type": "string"
          }
        ],
        "id": "string",
        "if": "string",
        "isCommon": 1,
        "lastRunTime": 1,
        "logo": "string",
        "name": "string",
        "productID": "string",
        "reason": "string",
        "status": 1,
        "tag": "string",
        "then": "string",
        "type": "string",
        "updatedTime": 1,
        "when": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/scene/get-list \
  --body '{"alarmID": "string"}'
```
