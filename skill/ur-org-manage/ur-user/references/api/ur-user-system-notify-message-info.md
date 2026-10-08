# ur-user system/notify/message/info

删除消息 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/notify/message/info/delete` | 删除消息 | platform |
| POST | `/api/v1/system/notify/message/info/get-list` | 获取消息列表 | platform |
| POST | `/api/v1/system/notify/message/info/send` | 发送消息 | platform |
| POST | `/api/v1/system/notify/message/info/update` | 更新消息 | platform |

## 详细说明

### POST `/api/v1/system/notify/message/info/delete`

**说明**: 删除消息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/message/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/message/info/get-list`

**说明**: 获取消息列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `group` | string | 否 | 分组 |
| `isGlobal` | integer | 否 | 全局消息 (格式: int64) |
| `notificationID` | string | 否 | 关联手动通知ID |
| `notifyCode` | string | 否 | 通知编码 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "group": "string",
  "isGlobal": 1,
  "notificationID": "string",
  "notifyCode": "string",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
        "actionText": "string",
        "actionURL": "string",
        "allowDismiss": "string",
        "body": "string",
        "coverImage": "string",
        "createdTime": "string",
        "emailFormat": "string",
        "expireTime": 1,
        "group": "string",
        "id": "string",
        "isGlobal": 1,
        "modalShowEndTime": "2026-01-01T00:00:00Z",
        "modalShowStartTime": "2026-01-01T00:00:00Z",
        "notificationID": "string",
        "notifyCode": "string",
        "notifyName": "示例名称",
        "notifyTime": 1,
        "priority": 1,
        "severity": "string",
        "showAsModal": "string",
        "str1": "string",
        "str2": "string",
        "str3": "string",
        "subject": "string",
        "type": "string",
        "videoUrl": "string"
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
ur api /api/v1/system/notify/message/info/get-list \
  --body '{"group": "string", "isGlobal": 1, "notificationID": "string", "notifyCode": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/notify/message/info/send`

**说明**: 发送消息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `body` | string | 是 |  |
| `isGlobal` | integer | 是 | 全局消息 (格式: int64) |
| `notifyCode` | string | 是 | 通知的code |
| `notifyTime` | integer | 否 | 通知时间,不填立刻发送 (格式: int64) |
| `str1` | string | 否 |  |
| `str2` | string | 否 |  |
| `str3` | string | 否 |  |
| `subject` | string | 是 | 通知主题 |
| `userIDs` | array[string] | 否 | 指定用户ID |

**请求示例**:
```json
{
  "body": "string",
  "isGlobal": 1,
  "notifyCode": "string",
  "notifyTime": 1,
  "str1": "string",
  "str2": "string",
  "str3": "string",
  "subject": "string",
  "userIDs": [
    "string"
  ]
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
ur api /api/v1/system/notify/message/info/send \
  --body '{"body": "string", "isGlobal": 1, "notifyCode": "string", "notifyTime": 1, "str1": "string", "str2": "string", "str3": "string", "subject": "string", "userIDs": ["string"]}'
```

### POST `/api/v1/system/notify/message/info/update`

**说明**: 更新消息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionText` | string | 否 | 操作按钮文案 |
| `actionURL` | string | 否 | 操作跳转URL |
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `body` | string | 是 | 内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `createdTime` | string | 是 |  |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `expireTime` | integer | 否 | 有效期(毫秒时间戳) (格式: int64) |
| `group` | string | 是 | 分组 |
| `id` | string | 是 |  |
| `isGlobal` | integer | 是 | 格式: int64 |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `notificationID` | string | 否 | 关联手动通知ID |
| `notifyCode` | string | 是 |  |
| `notifyName` | string | 否 | 通知的名称,只读 |
| `notifyTime` | integer | 是 | 通知时间 (格式: int64) |
| `priority` | integer | 否 | 优先级 1:紧急 2:重要 3:普通 (格式: int64) |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `str1` | string | 是 |  |
| `str2` | string | 是 |  |
| `str3` | string | 是 |  |
| `subject` | string | 是 | 主题 |
| `type` | string | 否 | 渠道类型 sms\|email\|dingTalk\|message等 |
| `videoUrl` | string | 否 |  news专属：视频URL |

**请求示例**:
```json
{
  "actionText": "string",
  "actionURL": "string",
  "allowDismiss": "string",
  "body": "string",
  "coverImage": "string",
  "createdTime": "string",
  "emailFormat": "string",
  "expireTime": 1,
  "group": "string",
  "id": "string",
  "isGlobal": 1,
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "notificationID": "string",
  "notifyCode": "string",
  "notifyName": "示例名称",
  "notifyTime": 1,
  "priority": 1,
  "severity": "string",
  "showAsModal": "string",
  "str1": "string",
  "str2": "string",
  "str3": "string",
  "subject": "string",
  "type": "string",
  "videoUrl": "string"
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
ur api /api/v1/system/notify/message/info/update \
  --body '{"actionText": "string", "actionURL": "string", "allowDismiss": "string", "body": "string", "coverImage": "string", "createdTime": "string", "emailFormat": "string", "expireTime": 1, "group": "string", "id": "string", "isGlobal": 1, "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "notificationID": "string", "notifyCode": "string", "notifyName": "示例名称", "notifyTime": 1, "priority": 1, "severity": "string", "showAsModal": "string", "str1": "string", "str2": "string", "str3": "string", "subject": "string", "type": "string", "videoUrl": "string"}'
```
