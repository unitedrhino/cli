# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/scene/info/create` | 新增场景 | admin |
| POST | `/api/v1/things/scene/info/delete` | 删除场景 | admin |
| POST | `/api/v1/things/scene/info/get-list` | 获取场景信息列表 | admin |
| POST | `/api/v1/things/scene/info/get-one` | 获取场景信息详情 | admin |
| POST | `/api/v1/things/scene/info/manually-trigger` | 手动触发场景 | admin |
| POST | `/api/v1/things/scene/info/update` | 更新场景 | admin |
| POST | `/api/v1/things/scene/log/get-list` | 获取场景日志列表 | admin |

## 端点详情

### POST `/api/v1/things/scene/info/create`

**说明**: 新增场景

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
ur api /api/v1/things/scene/info/create \
  --body '{"areaID": "string", "avatar": "string", "body": "string", "createdTime": "string", "desc": "string", "deviceAlias": "string", "deviceMode": "string", "deviceName": "string", "flowPath": [{"info": "string", "subType": "string", "type": "string"}], "id": "string", "if": "string", "isCommon": 1, "lastRunTime": 1, "logo": "string", "name": "string", "productID": "string", "reason": "string", "status": 1, "tag": "string", "then": "string", "type": "string", "updatedTime": 1, "when": "string"}'
```

### POST `/api/v1/things/scene/info/delete`

**说明**: 删除场景

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
ur api /api/v1/things/scene/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/scene/info/get-list`

**说明**: 获取场景信息列表

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
ur api /api/v1/things/scene/info/get-list \
  --body '{"alarmID": "string", "areaID": "string", "deviceFilterMode": 1, "deviceMode": "string", "deviceName": "string", "hasActionType": "string", "isCommon": 1, "isOnlyCore": 1, "name": "string", "page": {"page": 1, "pageSize": 1}, "productID": "string", "sceneIDs": [1], "status": 1, "tag": "string", "type": "string"}'
```

### POST `/api/v1/things/scene/info/get-one`

**说明**: 获取场景信息详情

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
ur api /api/v1/things/scene/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/things/scene/info/manually-trigger`

**说明**: 手动触发场景

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
ur api /api/v1/things/scene/info/manually-trigger \
  --body '{"id": 1}'
```

### POST `/api/v1/things/scene/info/update`

**说明**: 更新场景

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
ur api /api/v1/things/scene/info/update \
  --body '{"areaID": "string", "avatar": "string", "body": "string", "createdTime": "string", "desc": "string", "deviceAlias": "string", "deviceMode": "string", "deviceName": "string", "flowPath": [{"info": "string", "subType": "string", "type": "string"}], "id": "string", "if": "string", "isCommon": 1, "lastRunTime": 1, "logo": "string", "name": "string", "productID": "string", "reason": "string", "status": 1, "tag": "string", "then": "string", "type": "string", "updatedTime": 1, "when": "string"}'
```

### POST `/api/v1/things/scene/log/get-list`

**说明**: 获取场景日志列表

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
ur api /api/v1/things/scene/log/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}, "sceneID": "string", "sceneName": "示例名称", "status": 1, "timeRange": {"end": 1, "start": 1}}'
```
