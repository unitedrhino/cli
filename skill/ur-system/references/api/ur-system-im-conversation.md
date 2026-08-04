# ur-system im/conversation

群聊邀请成员 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/im/conversation/add-members` | 群聊邀请成员 | all |
| POST | `/api/v1/system/im/conversation/create` | 创建会话 | all |
| POST | `/api/v1/system/im/conversation/create-ai` | 创建 AI 会话 | all |
| POST | `/api/v1/system/im/conversation/disband` | 群聊解散 | all |
| POST | `/api/v1/system/im/conversation/get` | 获取会话详情 | all |
| POST | `/api/v1/system/im/conversation/get-by-participants` | 按参与者查询单聊会话 | all |
| POST | `/api/v1/system/im/conversation/get-list` | 获取会话列表 | all |
| POST | `/api/v1/system/im/conversation/get-members` | 获取会话成员列表 | all |
| POST | `/api/v1/system/im/conversation/mute-member` | 群聊禁言/取消禁言成员 | all |
| POST | `/api/v1/system/im/conversation/remove-member` | 群聊移除/退出成员 | all |
| POST | `/api/v1/system/im/conversation/set-admin` | 群聊设置/取消管理员 | all |
| POST | `/api/v1/system/im/conversation/transfer-owner` | 群聊转让群主 | all |
| POST | `/api/v1/system/im/conversation/update-setting` | 更新会话设置 | all |

## 详细说明

### POST `/api/v1/system/im/conversation/add-members`

**说明**: 群聊邀请成员

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `participants` | array[ImParticipant] | 是 |  被邀请的参与者列表 |

**请求示例**:
```json
{
  "conversationId": 1,
  "participants": [
    {
      "id": "string",
      "type": "string"
    }
  ]
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
ur api /api/v1/system/im/conversation/add-members \
  --body '{"conversationId": 1, "participants": [{"id": "string", "type": "string"}]}'
```

### POST `/api/v1/system/im/conversation/create`

**说明**: 创建会话

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  会话名称（群聊必填） |
| `participants` | array[ImParticipant] | 是 |  参与者列表 |
| `type` | string | 是 |  会话类型：single / group |

**请求示例**:
```json
{
  "name": "示例名称",
  "participants": [
    {
      "id": "string",
      "type": "string"
    }
  ],
  "type": "string"
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
ur api /api/v1/system/im/conversation/create \
  --body '{"name": "示例名称", "participants": [{"id": "string", "type": "string"}], "type": "string"}'
```

### POST `/api/v1/system/im/conversation/create-ai`

**说明**: 创建 AI 会话

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneId` | integer | 是 |  AI 分身（Clone）ID (格式: int64) |

**请求示例**:
```json
{
  "cloneId": 1
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
ur api /api/v1/system/im/conversation/create-ai \
  --body '{"cloneId": 1}'
```

### POST `/api/v1/system/im/conversation/disband`

**说明**: 群聊解散

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1
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
ur api /api/v1/system/im/conversation/disband \
  --body '{"conversationId": 1}'
```

### POST `/api/v1/system/im/conversation/get`

**说明**: 获取会话详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "info": {
      "avatar": "string",
      "createdTime": 1,
      "id": 1,
      "isMuted": 1,
      "isPinned": 1,
      "lastMsgContent": "string",
      "lastMsgID": 1,
      "lastMsgTime": 1,
      "lastReadMsgID": 1,
      "name": "示例名称",
      "ownerID": 1,
      "status": 1,
      "type": "string",
      "unreadCount": 1,
      "updatedTime": 1
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/conversation/get \
  --body '{"conversationId": 1}'
```

### POST `/api/v1/system/im/conversation/get-by-participants`

**说明**: 按参与者查询单聊会话

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `userID1` | integer | 是 |  用户 1 ID (格式: int64) |
| `userID2` | integer | 是 |  用户 2 ID (格式: int64) |

**请求示例**:
```json
{
  "userID1": 1,
  "userID2": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avatar": "string",
    "createdTime": 1,
    "id": 1,
    "isMuted": 1,
    "isPinned": 1,
    "lastMsgContent": "string",
    "lastMsgID": 1,
    "lastMsgTime": 1,
    "lastReadMsgID": 1,
    "name": "示例名称",
    "ownerID": 1,
    "status": 1,
    "type": "string",
    "unreadCount": 1,
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/conversation/get-by-participants \
  --body '{"userID1": 1, "userID2": 1}'
```

### POST `/api/v1/system/im/conversation/get-list`

**说明**: 获取会话列表

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
        "id": 1,
        "isMuted": 1,
        "isPinned": 1,
        "lastMsgContent": "string",
        "lastMsgID": 1,
        "lastMsgTime": 1,
        "lastReadMsgID": 1,
        "name": "示例名称",
        "ownerID": 1,
        "status": 1,
        "type": "string",
        "unreadCount": 1,
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
ur api /api/v1/system/im/conversation/get-list \
  --body '{"page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/conversation/get-members`

**说明**: 获取会话成员列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1
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
        "groupNickname": "string",
        "isMuted": 1,
        "joinedTime": 1,
        "nickName": "string",
        "participantId": "string",
        "participantType": "string",
        "role": "string",
        "userName": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/conversation/get-members \
  --body '{"conversationId": 1}'
```

### POST `/api/v1/system/im/conversation/mute-member`

**说明**: 群聊禁言/取消禁言成员

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `isMuted` | integer | 是 |  是否禁言：1 禁言 / 0 取消禁言 (格式: int64) |
| `participantId` | string | 是 |  参与者 ID |
| `participantType` | string | 是 |  参与者类型 |

**请求示例**:
```json
{
  "conversationId": 1,
  "isMuted": 1,
  "participantId": "string",
  "participantType": "string"
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
ur api /api/v1/system/im/conversation/mute-member \
  --body '{"conversationId": 1, "isMuted": 1, "participantId": "string", "participantType": "string"}'
```

### POST `/api/v1/system/im/conversation/remove-member`

**说明**: 群聊移除/退出成员

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `participantId` | string | 是 |  参与者 ID |
| `participantType` | string | 是 |  参与者类型 |

**请求示例**:
```json
{
  "conversationId": 1,
  "participantId": "string",
  "participantType": "string"
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
ur api /api/v1/system/im/conversation/remove-member \
  --body '{"conversationId": 1, "participantId": "string", "participantType": "string"}'
```

### POST `/api/v1/system/im/conversation/set-admin`

**说明**: 群聊设置/取消管理员

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `isAdmin` | integer | 是 |  是否管理员：1 设置 / 0 取消 (格式: int64) |
| `participantId` | string | 是 |  参与者 ID |
| `participantType` | string | 是 |  参与者类型 |

**请求示例**:
```json
{
  "conversationId": 1,
  "isAdmin": 1,
  "participantId": "string",
  "participantType": "string"
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
ur api /api/v1/system/im/conversation/set-admin \
  --body '{"conversationId": 1, "isAdmin": 1, "participantId": "string", "participantType": "string"}'
```

### POST `/api/v1/system/im/conversation/transfer-owner`

**说明**: 群聊转让群主

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `newOwnerID` | integer | 是 |  新群主用户 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "newOwnerID": 1
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
ur api /api/v1/system/im/conversation/transfer-owner \
  --body '{"conversationId": 1, "newOwnerID": 1}'
```

### POST `/api/v1/system/im/conversation/update-setting`

**说明**: 更新会话设置

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `isDeleted` | integer | 否 |  是否删除（退出会话列表） (格式: int64) |
| `isMuted` | integer | 否 |  是否免打扰 (格式: int64) |
| `isPinned` | integer | 否 |  是否置顶 (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "isDeleted": 1,
  "isMuted": 1,
  "isPinned": 1
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
ur api /api/v1/system/im/conversation/update-setting \
  --body '{"conversationId": 1, "isDeleted": 1, "isMuted": 1, "isPinned": 1}'
```
