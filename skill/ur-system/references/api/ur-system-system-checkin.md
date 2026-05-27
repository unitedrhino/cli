# ur-system system/checkIn

用户签到 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/check-in/do` | 用户签到 | admin |
| POST | `/api/v1/system/check-in/get-list` | 签到记录列表 | admin |
| POST | `/api/v1/system/check-in/point-balance/get` | 获取当前用户积分余额 | admin |
| POST | `/api/v1/system/check-in/point-log/adjust` | 管理员调整积分 | admin |
| POST | `/api/v1/system/check-in/point-log/get-list` | 积分流水列表 | admin |

## 详细说明

### POST `/api/v1/system/check-in/do`

**说明**: 用户签到

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "alreadyDone": true,
    "pointDelta": "string",
    "totalBalance": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/check-in/do \
  --body '{}'
```

### POST `/api/v1/system/check-in/get-list`

**说明**: 签到记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `endTime` | string | 否 |  结束时间(Unix毫秒) |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `startTime` | string | 否 |  开始时间(Unix毫秒) |
| `userID` | string | 否 |  管理员可指定用户ID |

**请求示例**:
```json
{
  "endTime": "2026-01-01T00:00:00Z",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "startTime": "2026-01-01T00:00:00Z",
  "userID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "checkDate": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "isContinuous": true,
        "pointDelta": "string",
        "userID": "string"
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
ur api /api/v1/system/check-in/get-list \
  --body '{"endTime": "2026-01-01T00:00:00Z", "page": {"page": 1, "pageSize": 1}, "startTime": "2026-01-01T00:00:00Z", "userID": "string"}'
```

### POST `/api/v1/system/check-in/point-balance/get`

**说明**: 获取当前用户积分余额

**权限**: admin

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "balance": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/check-in/point-balance/get \
  --body '{}'
```

### POST `/api/v1/system/check-in/point-log/adjust`

**说明**: 管理员调整积分

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `delta` | string | 是 |  积分变动量(正增负减,不为0) |
| `remark` | string | 是 |  调整备注(必填) |
| `targetUserID` | string | 是 |  目标用户ID |

**请求示例**:
```json
{
  "delta": "string",
  "remark": "string",
  "targetUserID": "string"
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
ur api /api/v1/system/check-in/point-log/adjust \
  --body '{"delta": "string", "remark": "string", "targetUserID": "string"}'
```

### POST `/api/v1/system/check-in/point-log/get-list`

**说明**: 积分流水列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `endTime` | string | 否 |  结束时间(Unix毫秒) |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `source` | string | 否 |  来源过滤: checkIn/admin |
| `startTime` | string | 否 |  开始时间(Unix毫秒) |
| `userID` | string | 否 |  管理员可指定用户ID |

**请求示例**:
```json
{
  "endTime": "2026-01-01T00:00:00Z",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "source": "string",
  "startTime": "2026-01-01T00:00:00Z",
  "userID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "balance": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "delta": "string",
        "id": "string",
        "remark": "string",
        "source": "string",
        "userID": "string"
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
ur api /api/v1/system/check-in/point-log/get-list \
  --body '{"endTime": "2026-01-01T00:00:00Z", "page": {"page": 1, "pageSize": 1}, "source": "string", "startTime": "2026-01-01T00:00:00Z", "userID": "string"}'
```
