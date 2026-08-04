# ur-ai things/alarm/notifyRecord

获取通知记录列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/notify-record/get-list` | 获取通知记录列表 | admin |
| POST | `/api/v1/things/alarm/notify-record/resend` | 重新发送通知 | admin |

## 详细说明

### POST `/api/v1/things/alarm/notify-record/get-list`

**说明**: 获取通知记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 否 |  告警规则ID |
| `alarmName` | string | 否 |  规则名称模糊查询 |
| `channel` | array[string] | 否 |  渠道列表 |
| `createdTimeRange` | object | 否 |  |
| `createdTimeRange.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `createdTimeRange.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |
| `eventID` | string | 否 |  事件ID |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | array[string] | 否 |  状态列表 |
| `targetName` | string | 否 |  告警对象模糊查询 |
| `timing` | string | 否 |  trigger / recovery |

**请求示例**:
```json
{
  "alarmID": "string",
  "alarmName": "示例名称",
  "channel": [
    "string"
  ],
  "createdTimeRange": {
    "end": 1,
    "start": 1
  },
  "eventID": "string",
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
  "status": [
    "string"
  ],
  "targetName": "string",
  "timing": "string"
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
ur api /api/v1/things/alarm/notify-record/get-list \
  --body '{"alarmID": "string", "alarmName": "示例名称", "channel": ["string"], "createdTimeRange": {"end": 1, "start": 1}, "eventID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": ["string"], "targetName": "string", "timing": "string"}'
```

### POST `/api/v1/things/alarm/notify-record/resend`

**说明**: 重新发送通知

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  通知记录ID |

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
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/notify-record/resend \
  --body '{"id": "string"}'
```
