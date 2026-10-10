# ur-rule things/rule/scene/log

获取场景日志列表(弃用)

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/rule/scene/log/get-list` | 获取场景日志列表(弃用) | admin |

## 详细说明

### POST `/api/v1/things/rule/scene/log/get-list`

**说明**: 获取场景日志列表(弃用)

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
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
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "sceneID": "string", "sceneName": "示例名称", "status": 1, "timeRange": {"end": 1, "start": 1}}'
```
