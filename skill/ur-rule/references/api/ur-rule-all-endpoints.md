# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/rule/alarm/info/create` | 新增告警(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/info/delete` | 删除告警(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/info/get-list` | 获取告警信息列表(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/info/get-one` | 获取告警信息(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/info/update` | 更新告警(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/record/deal` | 处理告警(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/record/get-list` | 获取告警记录列表(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/scene/batch-create` | 更新告警和场景的关联(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/scene/delete` | 删除告警和场景的关联(弃用) | admin |
| POST | `/api/v1/things/rule/alarm/scene/get-list` | 获取告警和场景的关联列表(弃用) | admin |
| POST | `/api/v1/things/rule/scene/info/create` | 新增场景(弃用) | admin |
| POST | `/api/v1/things/rule/scene/info/delete` | 删除场景(弃用) | admin |
| POST | `/api/v1/things/rule/scene/info/get-list` | 获取场景信息列表(弃用) | admin |
| POST | `/api/v1/things/rule/scene/info/get-one` | 获取场景信息详情(弃用) | admin |
| POST | `/api/v1/things/rule/scene/info/manually-trigger` | 手动触发场景(弃用) | admin |
| POST | `/api/v1/things/rule/scene/info/update` | 更新场景(弃用) | admin |
| POST | `/api/v1/things/rule/scene/log/get-list` | 获取场景日志列表(弃用) | admin |

## 端点详情

### POST `/api/v1/things/rule/alarm/info/create`

**说明**: 新增告警(弃用)

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
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/rule/alarm/info/create \
  --body '{"accounts": ["string"], "code": "string", "createdTime": "string", "desc": "string", "id": "string", "level": 1, "name": "string", "notifies": [{"templateID": "string", "templateName": "示例名称", "type": "string"}], "sceneIDs": [1], "scenes": [{"desc": "string", "id": "string", "name": "string", "status": 1}], "status": 1, "userIDs": ["string"]}'
```

### POST `/api/v1/things/rule/alarm/info/delete`

**说明**: 删除告警(弃用)

**权限**: admin

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
ur api /api/v1/things/rule/alarm/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/rule/alarm/info/get-list`

**说明**: 获取告警信息列表(弃用)

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `name` | string | 否 | 告警名模糊查询 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "name": "string",
  "page": {
    "page": 1,
    "pageSize": 1
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
ur api /api/v1/things/rule/alarm/info/get-list \
  --body '{"code": "string", "name": "string", "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/things/rule/alarm/info/get-one`

**说明**: 获取告警信息(弃用)

**权限**: admin

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
ur api /api/v1/things/rule/alarm/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/things/rule/alarm/info/update`

**说明**: 更新告警(弃用)

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
ur api /api/v1/things/rule/alarm/info/update \
  --body '{"accounts": ["string"], "code": "string", "createdTime": "string", "desc": "string", "id": "string", "level": 1, "name": "string", "notifies": [{"templateID": "string", "templateName": "示例名称", "type": "string"}], "sceneIDs": [1], "scenes": [{"desc": "string", "id": "string", "name": "string", "status": 1}], "status": 1, "userIDs": ["string"]}'
```

### POST `/api/v1/things/rule/alarm/record/deal`

**说明**: 处理告警(弃用)

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
ur api /api/v1/things/rule/alarm/record/deal \
  --body '{"handle": 1, "id": "string"}'
```

### POST `/api/v1/things/rule/alarm/record/get-list`

**说明**: 获取告警记录列表(弃用)

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
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
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
    "page": 1,
    "pageSize": 1
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
ur api /api/v1/things/rule/alarm/record/get-list \
  --body '{"alarmCode": "string", "alarmID": "string", "alarmName": "string", "areaID": "string", "areaIDPath": "string", "dealStatus": [1], "deviceAlias": "示例名称", "deviceName": "string", "page": {"page": 1, "pageSize": 1}, "timeRange": {"end": 1, "start": 1}}'
```

### POST `/api/v1/things/rule/alarm/scene/batch-create`

**说明**: 更新告警和场景的关联(弃用)

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
ur api /api/v1/things/rule/alarm/scene/batch-create \
  --body '{"alarmID": "string", "sceneIDs": [1]}'
```

### POST `/api/v1/things/rule/alarm/scene/delete`

**说明**: 删除告警和场景的关联(弃用)

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
ur api /api/v1/things/rule/alarm/scene/delete \
  --body '{"alarmID": "string", "sceneID": "string"}'
```

### POST `/api/v1/things/rule/alarm/scene/get-list`

**说明**: 获取告警和场景的关联列表(弃用)

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
ur api /api/v1/things/rule/alarm/scene/get-list \
  --body '{"alarmID": "string"}'
```

### POST `/api/v1/things/rule/scene/info/create`

**说明**: 新增场景(弃用)

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id 只读 |
| `avatar` | string | 否 |  头像 |
| `body` | string | 否 | 自定义字段 |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `deviceAlias` | string | 否 | 只读 |
| `deviceMode` | string | 否 | 设备模式 single:单设备 multi:多设备 |
| `deviceName` | string | 否 |  |
| `flowPath` | array[SceneFlowInfo] | 否 | 执行路径 |
| `id` | string | 否 |  |
| `if` | string | 否 |  |
| `isCommon` | integer | 否 | 是否是常用的 (格式: int64) |
| `lastRunTime` | integer | 否 | 格式: int64 |
| `logo` | string | 否 |  |
| `name` | string | 否 |  |
| `productID` | string | 否 |  |
| `reason` | string | 否 | 异常情况的描述说明 |
| `status` | integer | 否 | 状态: 状态（1启用 2禁用 3异常） (格式: int64) |
| `tag` | string | 否 | 标签 admin: 管理员 normal: 普通 |
| `then` | string | 否 |  |
| `type` | string | 否 |  |
| `updatedTime` | integer | 否 | 格式: int64 |
| `when` | string | 否 |  |

**请求示例**:
```json
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
ur api /api/v1/things/rule/scene/info/create \
  --body '{"areaID": "string", "avatar": "string", "body": "string", "createdTime": "string", "desc": "string", "deviceAlias": "string", "deviceMode": "string", "deviceName": "string", "flowPath": [{"info": "string", "subType": "string", "type": "string"}], "id": "string", "if": "string", "isCommon": 1, "lastRunTime": 1, "logo": "string", "name": "string", "productID": "string", "reason": "string", "status": 1, "tag": "string", "then": "string", "type": "string", "updatedTime": 1, "when": "string"}'
```

### POST `/api/v1/things/rule/scene/info/delete`

**说明**: 删除场景(弃用)

**权限**: admin

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
ur api /api/v1/things/rule/scene/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/rule/scene/info/get-list`

**说明**: 获取场景信息列表(弃用)

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 否 | 告警id |
| `areaID` | string | 否 |  |
| `deviceFilterMode` | integer | 否 | 设备过滤模式: 1,过滤触发或执行(默认) 2,过滤触发 3,过滤执行 (格式: int64) |
| `deviceMode` | string | 否 | 设备模式 |
| `deviceName` | string | 否 | 设备名 |
| `hasActionType` | string | 否 | 过滤有某个执行类型 |
| `isCommon` | integer | 否 | 是否是常用的 (格式: int64) |
| `isOnlyCore` | integer | 否 | 不返回if,when和then (格式: int64) |
| `name` | string | 否 | 场景名模糊查询 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 | 产品id |
| `sceneIDs` | array[integer] | 否 | 根据场景ID来过滤 |
| `status` | integer | 否 | 状态（1启用 2禁用 3异常） (格式: int64) |
| `tag` | string | 否 | 标签 admin: 管理员 normal: 普通 |
| `type` | string | 否 | 触发类型 auto manual:手动触发 |

**请求示例**:
```json
{
  "alarmID": "string",
  "areaID": "string",
  "deviceFilterMode": 1,
  "deviceMode": "string",
  "deviceName": "string",
  "hasActionType": "string",
  "isCommon": 1,
  "isOnlyCore": 1,
  "name": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "productID": "string",
  "sceneIDs": [
    1
  ],
  "status": 1,
  "tag": "string",
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
ur api /api/v1/things/rule/scene/info/get-list \
  --body '{"alarmID": "string", "areaID": "string", "deviceFilterMode": 1, "deviceMode": "string", "deviceName": "string", "hasActionType": "string", "isCommon": 1, "isOnlyCore": 1, "name": "string", "page": {"page": 1, "pageSize": 1}, "productID": "string", "sceneIDs": [1], "status": 1, "tag": "string", "type": "string"}'
```

### POST `/api/v1/things/rule/scene/info/get-one`

**说明**: 获取场景信息详情(弃用)

**权限**: admin

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
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/rule/scene/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/things/rule/scene/info/manually-trigger`

**说明**: 手动触发场景(弃用)

**权限**: admin

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
ur api /api/v1/things/rule/scene/info/manually-trigger \
  --body '{"id": 1}'
```

### POST `/api/v1/things/rule/scene/info/update`

**说明**: 更新场景(弃用)

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id 只读 |
| `avatar` | string | 否 |  头像 |
| `body` | string | 否 | 自定义字段 |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `deviceAlias` | string | 否 | 只读 |
| `deviceMode` | string | 否 | 设备模式 single:单设备 multi:多设备 |
| `deviceName` | string | 否 |  |
| `flowPath` | array[SceneFlowInfo] | 否 | 执行路径 |
| `id` | string | 否 |  |
| `if` | string | 否 |  |
| `isCommon` | integer | 否 | 是否是常用的 (格式: int64) |
| `lastRunTime` | integer | 否 | 格式: int64 |
| `logo` | string | 否 |  |
| `name` | string | 否 |  |
| `productID` | string | 否 |  |
| `reason` | string | 否 | 异常情况的描述说明 |
| `status` | integer | 否 | 状态: 状态（1启用 2禁用 3异常） (格式: int64) |
| `tag` | string | 否 | 标签 admin: 管理员 normal: 普通 |
| `then` | string | 否 |  |
| `type` | string | 否 |  |
| `updatedTime` | integer | 否 | 格式: int64 |
| `when` | string | 否 |  |

**请求示例**:
```json
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
ur api /api/v1/things/rule/scene/info/update \
  --body '{"areaID": "string", "avatar": "string", "body": "string", "createdTime": "string", "desc": "string", "deviceAlias": "string", "deviceMode": "string", "deviceName": "string", "flowPath": [{"info": "string", "subType": "string", "type": "string"}], "id": "string", "if": "string", "isCommon": 1, "lastRunTime": 1, "logo": "string", "name": "string", "productID": "string", "reason": "string", "status": 1, "tag": "string", "then": "string", "type": "string", "updatedTime": 1, "when": "string"}'
```

### POST `/api/v1/things/rule/scene/log/get-list`

**说明**: 获取场景日志列表(弃用)

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `sceneID` | string | 否 | 场景id |
| `sceneName` | string | 否 | 场景名称 |
| `status` | integer | 否 |  状态（1成功 2失败） (格式: int64) |
| `timeRange` | object | 否 |  |
| `timeRange.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `timeRange.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "sceneID": "string",
  "sceneName": "示例名称",
  "status": 1,
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
        "actions": [
          {
            "alarm": {
              "mode": "..."
            },
            "code": 1,
            "delay": 1,
            "device": {
              "deviceAlias": "...",
              "deviceName": "...",
              "productID": "...",
              "productName": "...",
              "type": "...",
              "values": "..."
            },
            "msg": "string",
            "msgToken": "string",
            "scene": {
              "sceneID": "...",
              "sceneName": "..."
            },
            "status": 1,
            "type": "string"
          }
        ],
        "areaID": "string",
        "createdTime": 1,
        "sceneID": "string",
        "sceneName": "示例名称",
        "status": 1,
        "trigger": {
          "device": {
            "dataID": "string",
            "dataName": "string",
            "deviceAlias": "string",
            "deviceName": "string",
            "productID": "string",
            "type": "string",
            "value": "string"
          },
          "time": 1,
          "type": "string"
        },
        "type": "string"
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
ur api /api/v1/things/rule/scene/log/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}, "sceneID": "string", "sceneName": "示例名称", "status": 1, "timeRange": {"end": 1, "start": 1}}'
```
