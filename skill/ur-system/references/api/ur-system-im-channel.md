# ur-system im/channel

绑定渠道端点 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/im/channel/bind` | 绑定渠道端点 | all |
| POST | `/api/v1/system/im/channel/get-list` | 查询渠道绑定列表 | all |
| POST | `/api/v1/system/im/channel/set-mute` | 设置端点免打扰 | all |
| POST | `/api/v1/system/im/channel/unbind` | 解绑渠道端点 | all |

## 详细说明

### POST `/api/v1/system/im/channel/bind`

**说明**: 绑定渠道端点

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `channelId` | string | 是 |  渠道标识（device 为 {productID}:{deviceName}） |
| `channelName` | string | 否 |  渠道展示名（设备名称） |
| `channelType` | string | 是 |  渠道类型：device / ... |

**请求示例**:
```json
{
  "channelId": "string",
  "channelName": "示例名称",
  "channelType": "string"
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
ur api /api/v1/system/im/channel/bind \
  --body '{"channelId": "string", "channelName": "示例名称", "channelType": "string"}'
```

### POST `/api/v1/system/im/channel/get-list`

**说明**: 查询渠道绑定列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "channelId": "string",
        "channelName": "string",
        "channelType": "string",
        "createdTime": 1,
        "id": 1,
        "isMuted": 1,
        "userId": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/channel/get-list \
  --body '{}'
```

### POST `/api/v1/system/im/channel/set-mute`

**说明**: 设置端点免打扰

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  绑定记录 ID (格式: int64) |
| `isMuted` | integer | 是 |  0-下发 1-免打扰 (格式: int64) |

**请求示例**:
```json
{
  "id": 1,
  "isMuted": 1
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
ur api /api/v1/system/im/channel/set-mute \
  --body '{"id": 1, "isMuted": 1}'
```

### POST `/api/v1/system/im/channel/unbind`

**说明**: 解绑渠道端点

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  绑定记录 ID (格式: int64) |

**请求示例**:
```json
{
  "id": 1
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
ur api /api/v1/system/im/channel/unbind \
  --body '{"id": 1}'
```
