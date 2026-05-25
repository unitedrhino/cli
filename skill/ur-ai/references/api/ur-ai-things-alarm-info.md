# ur-ai things/alarm/info

新增告警 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/info/create` | 新增告警 | admin |
| POST | `/api/v1/things/alarm/info/delete` | 删除告警 | admin |
| POST | `/api/v1/things/alarm/info/get-list` | 获取告警信息列表 | admin |
| POST | `/api/v1/things/alarm/info/get-one` | 获取告警信息 | admin |
| POST | `/api/v1/things/alarm/info/update` | 更新告警 | admin |

## 详细说明

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
