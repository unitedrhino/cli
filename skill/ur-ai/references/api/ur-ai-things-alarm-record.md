# ur-ai things/alarm/record

处理告警 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/record/deal` | 处理告警 | admin |
| POST | `/api/v1/things/alarm/record/get-list` | 获取告警记录列表 | admin |

## 详细说明

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
