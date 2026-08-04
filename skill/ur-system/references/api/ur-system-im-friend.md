# ur-system im/friend

发起好友申请 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/im/friend/apply` | 发起好友申请 | all |
| POST | `/api/v1/system/im/friend/delete` | 删除好友 | all |
| POST | `/api/v1/system/im/friend/get-list` | 获取好友列表 | all |
| POST | `/api/v1/system/im/friend/handle` | 处理好友申请 | all |
| POST | `/api/v1/system/im/friend/request/get-list` | 获取好友申请列表 | all |

## 详细说明

### POST `/api/v1/system/im/friend/apply`

**说明**: 发起好友申请

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `recvUserID` | integer | 是 |  被申请人用户 ID (格式: int64) |
| `reqMsg` | string | 否 |  申请附言 |

**请求示例**:
```json
{
  "recvUserID": 1,
  "reqMsg": "string"
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
ur api /api/v1/system/im/friend/apply \
  --body '{"recvUserID": 1, "reqMsg": "string"}'
```

### POST `/api/v1/system/im/friend/delete`

**说明**: 删除好友

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `friendUserID` | integer | 是 |  好友用户 ID (格式: int64) |

**请求示例**:
```json
{
  "friendUserID": 1
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
ur api /api/v1/system/im/friend/delete \
  --body '{"friendUserID": 1}'
```

### POST `/api/v1/system/im/friend/get-list`

**说明**: 获取好友列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "avatar": "string",
        "createdTime": 1,
        "friendUserID": 1,
        "id": 1,
        "nickName": "string",
        "remark": "string",
        "userID": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/friend/get-list \
  --body '{"page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/friend/handle`

**说明**: 处理好友申请

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `handleMsg` | string | 否 |  处理附言 |
| `handleResult` | integer | 是 |  处理结果：1 同意 / 2 拒绝 (格式: int64) |
| `requestID` | integer | 是 |  申请记录 ID (格式: int64) |

**请求示例**:
```json
{
  "handleMsg": "string",
  "handleResult": 1,
  "requestID": 1
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
ur api /api/v1/system/im/friend/handle \
  --body '{"handleMsg": "string", "handleResult": 1, "requestID": 1}'
```

### POST `/api/v1/system/im/friend/request/get-list`

**说明**: 获取好友申请列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `direction` | integer | 是 |  方向：0 我收到的 / 1 我发起的 (格式: int64) |
| `handleResult` | integer | 是 |  处理结果过滤：-1 全部 (格式: int64) |
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "direction": 1,
  "handleResult": 1,
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": 1,
        "handleMsg": "string",
        "handleResult": 1,
        "id": 1,
        "recvUserAvatar": "string",
        "recvUserID": 1,
        "recvUserNickName": "string",
        "reqMsg": "string",
        "reqUserAvatar": "string",
        "reqUserID": 1,
        "reqUserNickName": "string",
        "updatedTime": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/friend/request/get-list \
  --body '{"direction": 1, "handleResult": 1, "page": 1, "pageSize": 1}'
```
