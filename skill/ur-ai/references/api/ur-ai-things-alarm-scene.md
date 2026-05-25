# ur-ai things/alarm/scene

更新告警和场景的关联 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/scene/batch-create` | 更新告警和场景的关联 | admin |
| POST | `/api/v1/things/alarm/scene/delete` | 删除告警和场景的关联 | admin |
| POST | `/api/v1/things/alarm/scene/get-list` | 获取告警和场景的关联列表 | admin |

## 详细说明

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
