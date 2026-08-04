# ur-ai things/alarm/capability

获取告警能力矩阵

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/capability/get-one` | 获取告警能力矩阵 | admin |

## 详细说明

### POST `/api/v1/things/alarm/capability/get-one`

**说明**: 获取告警能力矩阵

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "criteria": [
      {
        "allowedChildren": [
          "string"
        ],
        "type": "string"
      }
    ],
    "notify": {
      "callbackAuthTypes": [
        "string"
      ],
      "callbackMethods": [
        "string"
      ],
      "channels": [
        "string"
      ],
      "timings": [
        "string"
      ]
    },
    "sources": [
      {
        "aggregates": [
          {
            "compares": {},
            "dataType": "string",
            "values": []
          }
        ],
        "compareTypes": [
          "string"
        ],
        "dataTypes": [
          "string"
        ],
        "sourceType": "string",
        "supportsFilters": true
      }
    ],
    "targets": [
      "string"
    ],
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/capability/get-one \
  --body '{}'
```
