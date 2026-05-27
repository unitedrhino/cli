# ur-user system/notify/template

添加通知模版 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/notify/template/create` | 添加通知模版 | platform |
| POST | `/api/v1/system/notify/template/delete` | 删除通知模版 | platform |
| POST | `/api/v1/system/notify/template/get-list` | 获取通知模版列表 | platform |
| POST | `/api/v1/system/notify/template/get-one` | 获取通知模版详情 | platform |
| POST | `/api/v1/system/notify/template/update` | 更新通知模版 | platform |

## 详细说明

### POST `/api/v1/system/notify/template/create`

**说明**: 添加通知模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `body` | string | 否 | 默认模版内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `desc` | string | 否 |  备注 |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `id` | string | 否 |  id编号 |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `name` | string | 否 | 通知的命名 |
| `notifyCode` | string | 否 | 对应的配置Code |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `subject` | string | 否 | 默认消息主题(邮箱需要填写) |
| `templateCode` | string | 否 |  短信通知模版编码(短信必填) |
| `type` | string | 否 | 对应的配置类型 sms:短信, email:邮箱, dingTalk:钉钉机器人, dingWebhook:钉钉webhook, wxMini:微信小程序推送, wxEWebHook:企业微信webhook, message:站内信通知, phoneCall:电话通知, banner:横幅通知, news:资讯通知, wxOfficialAccount:微信公众号 可选: ['sms', 'email', 'dingTalk', 'dingWebhook', 'wxMini', 'wxEWebHook', 'message', 'phoneCall', 'banner', 'news', 'wxOfficialAccount'] |
| `videoUrl` | string | 否 |  news专属：视频URL |
| `webhook` | string | 否 | 钉钉webhook及企微webhook |

**请求示例**:
```json
{
  "allowDismiss": "string",
  "body": "string",
  "coverImage": "string",
  "desc": "string",
  "emailFormat": "string",
  "id": "string",
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "name": "string",
  "notifyCode": "string",
  "severity": "string",
  "showAsModal": "string",
  "subject": "string",
  "templateCode": "string",
  "type": "sms",
  "videoUrl": "string",
  "webhook": "string"
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
ur api /api/v1/system/notify/template/create \
  --body '{"allowDismiss": "string", "body": "string", "coverImage": "string", "desc": "string", "emailFormat": "string", "id": "string", "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "name": "string", "notifyCode": "string", "severity": "string", "showAsModal": "string", "subject": "string", "templateCode": "string", "type": "sms", "videoUrl": "string", "webhook": "string"}'
```

### POST `/api/v1/system/notify/template/delete`

**说明**: 删除通知模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
ur api /api/v1/system/notify/template/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/template/get-list`

**说明**: 获取通知模版列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `notifyCode` | string | 否 |  应用编号 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `type` | string | 否 | 对应的配置类型 sms email |

**请求示例**:
```json
{
  "name": "string",
  "notifyCode": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "allowDismiss": "string",
        "body": "string",
        "coverImage": "string",
        "desc": "string",
        "emailFormat": "string",
        "id": "string",
        "modalShowEndTime": "2026-01-01T00:00:00Z",
        "modalShowStartTime": "2026-01-01T00:00:00Z",
        "name": "string",
        "notifyCode": "string",
        "severity": "string",
        "showAsModal": "string",
        "subject": "string",
        "templateCode": "string",
        "type": "sms",
        "videoUrl": "string",
        "webhook": "string"
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
ur api /api/v1/system/notify/template/get-list \
  --body '{"name": "string", "notifyCode": "string", "page": {"page": 1, "pageSize": 1}, "type": "string"}'
```

### POST `/api/v1/system/notify/template/get-one`

**说明**: 获取通知模版详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "data": {
    "allowDismiss": "string",
    "body": "string",
    "coverImage": "string",
    "desc": "string",
    "emailFormat": "string",
    "id": "string",
    "modalShowEndTime": "2026-01-01T00:00:00Z",
    "modalShowStartTime": "2026-01-01T00:00:00Z",
    "name": "string",
    "notifyCode": "string",
    "severity": "string",
    "showAsModal": "string",
    "subject": "string",
    "templateCode": "string",
    "type": "sms",
    "videoUrl": "string",
    "webhook": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/template/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/template/update`

**说明**: 更新通知模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `body` | string | 否 | 默认模版内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `desc` | string | 否 |  备注 |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `id` | string | 否 |  id编号 |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `name` | string | 否 | 通知的命名 |
| `notifyCode` | string | 否 | 对应的配置Code |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `subject` | string | 否 | 默认消息主题(邮箱需要填写) |
| `templateCode` | string | 否 |  短信通知模版编码(短信必填) |
| `type` | string | 否 | 对应的配置类型 sms:短信, email:邮箱, dingTalk:钉钉机器人, dingWebhook:钉钉webhook, wxMini:微信小程序推送, wxEWebHook:企业微信webhook, message:站内信通知, phoneCall:电话通知, banner:横幅通知, news:资讯通知, wxOfficialAccount:微信公众号 可选: ['sms', 'email', 'dingTalk', 'dingWebhook', 'wxMini', 'wxEWebHook', 'message', 'phoneCall', 'banner', 'news', 'wxOfficialAccount'] |
| `videoUrl` | string | 否 |  news专属：视频URL |
| `webhook` | string | 否 | 钉钉webhook及企微webhook |

**请求示例**:
```json
{
  "allowDismiss": "string",
  "body": "string",
  "coverImage": "string",
  "desc": "string",
  "emailFormat": "string",
  "id": "string",
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "name": "string",
  "notifyCode": "string",
  "severity": "string",
  "showAsModal": "string",
  "subject": "string",
  "templateCode": "string",
  "type": "sms",
  "videoUrl": "string",
  "webhook": "string"
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
ur api /api/v1/system/notify/template/update \
  --body '{"allowDismiss": "string", "body": "string", "coverImage": "string", "desc": "string", "emailFormat": "string", "id": "string", "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "name": "string", "notifyCode": "string", "severity": "string", "showAsModal": "string", "subject": "string", "templateCode": "string", "type": "sms", "videoUrl": "string", "webhook": "string"}'
```
