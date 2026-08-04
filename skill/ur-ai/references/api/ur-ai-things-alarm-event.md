# ur-ai things/alarm/event

处理告警事件 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/event/deal` | 处理告警事件 | admin |
| POST | `/api/v1/things/alarm/event/false-alarm` | 标记误报 | admin |
| POST | `/api/v1/things/alarm/event/get-list` | 获取告警事件列表 | admin |
| POST | `/api/v1/things/alarm/event/get-one` | 获取告警事件详情 | admin |
| POST | `/api/v1/things/alarm/event/stat` | 告警事件统计 | admin |

## 详细说明

### POST `/api/v1/things/alarm/event/deal`

**说明**: 处理告警事件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `action` | string | 是 |  ack / shield / unshield |
| `eventID` | string | 是 |  事件ID |
| `remark` | string | 否 |  备注 |

**请求示例**:
```json
{
  "action": "string",
  "eventID": "string",
  "remark": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/event/deal \
  --body '{"action": "string", "eventID": "string", "remark": "string"}'
```

### POST `/api/v1/things/alarm/event/false-alarm`

**说明**: 标记误报

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `eventID` | string | 是 |  事件ID |
| `reason` | string | 否 |  原因 |
| `reasonType` | string | 是 |  configurationError / dataError / drill / other |

**请求示例**:
```json
{
  "eventID": "string",
  "reason": "string",
  "reasonType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/event/false-alarm \
  --body '{"eventID": "string", "reason": "string", "reasonType": "string"}'
```

### POST `/api/v1/things/alarm/event/get-list`

**说明**: 获取告警事件列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 否 |  告警规则ID |
| `deviceName` | string | 否 |  设备名 |
| `isFalseAlarm` | boolean | 否 |  是否误报 (格式: boolean) |
| `isRecovered` | boolean | 否 |  是否已恢复 (格式: boolean) |
| `keyword` | string | 否 |  名称模糊搜索 |
| `level` | array[string] | 否 |  级别列表 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 |  产品ID |
| `status` | array[string] | 否 |  状态列表 |
| `triggerTimeRange` | object | 否 |  |
| `triggerTimeRange.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `triggerTimeRange.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |

**请求示例**:
```json
{
  "alarmID": "string",
  "deviceName": "string",
  "isFalseAlarm": true,
  "isRecovered": true,
  "keyword": "示例名称",
  "level": [
    "string"
  ],
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
  "productID": "string",
  "status": [
    "string"
  ],
  "triggerTimeRange": {
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
        "alarmID": "string",
        "alarmName": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "deviceName": "string",
        "falseAlarm": true,
        "falseAlarmReason": "string",
        "falseAlarmReasonType": "string",
        "firstTriggerTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "lastTriggerTime": "2026-01-01T00:00:00Z",
        "level": "string",
        "productID": "string",
        "recoveryTime": "2026-01-01T00:00:00Z",
        "status": "string",
        "target": {
          "areas": [
            1
          ],
          "categoryID": "string",
          "devices": [
            {
              "alias": "...",
              "name": "...",
              "productID": "..."
            }
          ],
          "groups": [
            1
          ],
          "productID": "string",
          "selectType": "string"
        },
        "targetName": "string",
        "triggerContent": "string",
        "triggerCount": 1,
        "triggerValue": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/alarm/event/get-list \
  --body '{"alarmID": "string", "deviceName": "string", "isFalseAlarm": true, "isRecovered": true, "keyword": "示例名称", "level": ["string"], "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string", "status": ["string"], "triggerTimeRange": {"end": 1, "start": 1}}'
```

### POST `/api/v1/things/alarm/event/get-one`

**说明**: 获取告警事件详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  ID |

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
    "dealRecords": [
      {
        "action": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "eventID": "string",
        "fromStatus": "string",
        "id": "string",
        "operatorID": "string",
        "operatorName": "string",
        "remark": "string",
        "toStatus": "string"
      }
    ],
    "event": {
      "alarmID": "string",
      "alarmName": "string",
      "createdTime": "2026-01-01T00:00:00Z",
      "deviceName": "string",
      "falseAlarm": true,
      "falseAlarmReason": "string",
      "falseAlarmReasonType": "string",
      "firstTriggerTime": "2026-01-01T00:00:00Z",
      "id": "string",
      "lastTriggerTime": "2026-01-01T00:00:00Z",
      "level": "string",
      "productID": "string",
      "recoveryTime": "2026-01-01T00:00:00Z",
      "status": "string",
      "target": {
        "areas": [
          1
        ],
        "categoryID": "string",
        "devices": [
          {
            "alias": "示例名称",
            "name": "string",
            "productID": "string"
          }
        ],
        "groups": [
          1
        ],
        "productID": "string",
        "selectType": "string"
      },
      "targetName": "string",
      "triggerContent": "string",
      "triggerCount": 1,
      "triggerValue": "string",
      "updatedTime": "2026-01-01T00:00:00Z"
    },
    "notifyRecords": [
      {
        "alarmID": "string",
        "alarmName": "示例名称",
        "channel": "string",
        "content": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "eventID": "string",
        "externalResp": "string",
        "failReason": "string",
        "id": "string",
        "ladderOrder": 1,
        "retryTimes": 1,
        "sentTime": "2026-01-01T00:00:00Z",
        "status": "string",
        "targetName": "示例名称",
        "targets": {
          "groupIDs": [
            "string"
          ],
          "roleIDs": [
            "string"
          ],
          "userIDs": [
            "string"
          ]
        },
        "templateID": "string",
        "timing": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/event/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/event/stat`

**说明**: 告警事件统计

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupBy` | string | 是 |  level / status / alarm / date / product |
| `level` | array[string] | 否 |  级别列表 |
| `status` | array[string] | 否 |  状态列表 |
| `triggerTimeRange` | object | 否 |  |
| `triggerTimeRange.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `triggerTimeRange.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |

**请求示例**:
```json
{
  "groupBy": "string",
  "level": [
    "string"
  ],
  "status": [
    "string"
  ],
  "triggerTimeRange": {
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
        "key": "string",
        "name": "string",
        "recovered": 1,
        "total": 1,
        "unRecovered": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/event/stat \
  --body '{"groupBy": "string", "level": ["string"], "status": ["string"], "triggerTimeRange": {"end": 1, "start": 1}}'
```
