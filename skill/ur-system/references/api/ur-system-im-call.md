# ur-system im/call

查询群通话进行状态 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| GET | `/api/v1/system/im/call/active` | 查询群通话进行状态 | all |
| POST | `/api/v1/system/im/call/signal` | 通话信令统一入口 | all |

## 详细说明

### GET `/api/v1/system/im/call/active`

**说明**: 查询群通话进行状态

**权限**: all

**路径参数**:

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| `conversationId` | query | integer | 是 |  群聊会话 ID |

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "active": true,
    "callId": "string",
    "callType": "string",
    "participantIds": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/call/active \
  --body '{}'
```

### POST `/api/v1/system/im/call/signal`

**说明**: 通话信令统一入口

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `callId` | string | 否 |  通话 ID（invite 类信令为空，由服务端生成） |
| `callType` | string | 否 |  通话媒体类型：voice / video |
| `calleeId` | string | 否 |  1:1 被叫用户 ID |
| `conversationId` | integer | 否 |  关联的 IM 会话 ID (格式: int64) |
| `members` | array[string] | 否 |  群通话邀请成员用户 ID 列表 |
| `payload` | string | 否 |  扩展 JSON（media_state 的 mute/video 状态等） |
| `to` | string | 否 |  media_state 定向目标用户 ID（空=1:1 对端/群通话全员） |
| `type` | string | 是 |  信令类型：call_invite/call_accept/call_reject/call_cancel/call_end/call_heartbeat/group_invite/group_join/group_leave/media_state |

**请求示例**:
```json
{
  "callId": "string",
  "callType": "string",
  "calleeId": "string",
  "conversationId": 1,
  "members": [
    "string"
  ],
  "payload": "string",
  "to": "string",
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "callId": "string",
    "event": "string",
    "notifyUserIDs": [
      1
    ],
    "payload": "string",
    "state": "string",
    "zlmPushUrl": "string",
    "zlmServer": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/call/signal \
  --body '{"callId": "string", "callType": "string", "calleeId": "string", "conversationId": 1, "members": ["string"], "payload": "string", "to": "string", "type": "string"}'
```
