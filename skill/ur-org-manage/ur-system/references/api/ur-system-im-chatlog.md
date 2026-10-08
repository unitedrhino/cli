# ur-system im/chatlog

获取消息列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/im/chatlog/get-list` | 获取消息列表 | all |
| POST | `/api/v1/system/im/chatlog/mark-read` | 标记已读 | all |
| POST | `/api/v1/system/im/chatlog/search` | 搜索聊天记录 | all |
| POST | `/api/v1/system/im/message/recall` | 撤回消息 | all |
| POST | `/api/v1/system/im/message/send` | 发送消息 | all |

## 详细说明

### POST `/api/v1/system/im/chatlog/get-list`

**说明**: 获取消息列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
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
        "clientMsgId": "string",
        "content": "string",
        "conversationId": 1,
        "extra": "string",
        "id": 1,
        "isAtAll": 1,
        "mentionUserIDs": [
          1
        ],
        "msgType": "string",
        "sendTime": 1,
        "senderId": "string",
        "senderType": "string",
        "seq": 1,
        "status": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/chatlog/get-list \
  --body '{"conversationId": 1, "page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/chatlog/mark-read`

**说明**: 标记已读

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `readSeq` | integer | 否 |  已读到的序号 (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "readSeq": 1
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
ur api /api/v1/system/im/chatlog/mark-read \
  --body '{"conversationId": 1, "readSeq": 1}'
```

### POST `/api/v1/system/im/chatlog/search`

**说明**: 搜索聊天记录

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 否 |  会话 ID（可选，限定单个会话） (格式: int64) |
| `keyword` | string | 是 |  搜索关键词 |
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "keyword": "string",
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
        "clientMsgId": "string",
        "content": "string",
        "conversationId": 1,
        "extra": "string",
        "id": 1,
        "isAtAll": 1,
        "mentionUserIDs": [
          1
        ],
        "msgType": "string",
        "sendTime": 1,
        "senderId": "string",
        "senderType": "string",
        "seq": 1,
        "status": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/chatlog/search \
  --body '{"conversationId": 1, "keyword": "string", "page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/message/recall`

**说明**: 撤回消息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `msgId` | integer | 是 |  消息 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "msgId": 1
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
ur api /api/v1/system/im/message/recall \
  --body '{"conversationId": 1, "msgId": 1}'
```

### POST `/api/v1/system/im/message/send`

**说明**: 发送消息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `clientMsgId` | string | 否 |  客户端消息 ID |
| `content` | string | 是 |  消息内容 |
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `extra` | string | 否 |  扩展字段 |
| `isAtAll` | integer | 否 |  是否 @ 全体成员 (格式: int64) |
| `mentionUserIDs` | array[integer] | 否 |  @ 的用户列表 |
| `msgType` | string | 是 |  消息类型 |

**请求示例**:
```json
{
  "clientMsgId": "string",
  "content": "string",
  "conversationId": 1,
  "extra": "string",
  "isAtAll": 1,
  "mentionUserIDs": [
    1
  ],
  "msgType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "sendTime": 1,
    "seq": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/message/send \
  --body '{"clientMsgId": "string", "content": "string", "conversationId": 1, "extra": "string", "isAtAll": 1, "mentionUserIDs": [1], "msgType": "string"}'
```
