# ur-user system/notify/notification

复制手动通知 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/notify/notification/copy` | 复制手动通知 | platform |
| POST | `/api/v1/system/notify/notification/create` | 创建手动通知 | platform |
| POST | `/api/v1/system/notify/notification/delete` | 删除手动通知 | platform |
| POST | `/api/v1/system/notify/notification/estimate-users` | 预估送达用户数 | platform |
| POST | `/api/v1/system/notify/notification/index` | 获取手动通知列表 | platform |
| POST | `/api/v1/system/notify/notification/read` | 获取手动通知详情 | platform |
| POST | `/api/v1/system/notify/notification/revoke` | 撤回手动通知 | platform |
| POST | `/api/v1/system/notify/notification/send` | 发送手动通知 | platform |
| POST | `/api/v1/system/notify/notification/update` | 更新手动通知 | platform |

## 详细说明

### POST `/api/v1/system/notify/notification/copy`

**说明**: 复制手动通知

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
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/notification/copy \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/create`

**说明**: 创建手动通知

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionText` | string | 否 | 操作按钮文案 |
| `actionURL` | string | 否 | 操作跳转URL |
| `actualUsers` | integer | 否 | 实际送达数,只读 (格式: int64) |
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `audienceType` | integer | 否 | 受众类型 1:指定用户 2:全局 (格式: int64) |
| `channels` | array[string] | 否 | 发送渠道列表 |
| `content` | string | 是 | 内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `createdTime` | integer | 否 | 创建时间,只读 (格式: int64) |
| `creatorID` | string | 否 | 创建者ID,只读 |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `estimatedUsers` | integer | 否 | 预计送达数,只读 (格式: int64) |
| `expireTime` | integer | 否 | 有效期(毫秒时间戳) (格式: int64) |
| `group` | string | 否 | 分组 |
| `id` | string | 否 | 通知ID |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `priority` | integer | 否 | 优先级 1:紧急 2:重要 3:普通 (格式: int64) |
| `scheduleTime` | integer | 否 | 计划发送时间(毫秒时间戳) (格式: int64) |
| `sentTime` | integer | 否 | 实际发送时间(毫秒时间戳),只读 (格式: int64) |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `status` | integer | 否 | 状态 1:草稿 2:已计划 3:发送中 4:已发送 5:已撤回 (格式: int64) |
| `targetUserIDs` | array[string] | 否 | 目标用户ID列表 |
| `tenantCode` | string | 否 | 租户编码,空=平台级 |
| `title` | string | 是 | 标题 |
| `videoUrl` | string | 否 |  news专属：视频URL |

**请求示例**:
```json
{
  "actionText": "string",
  "actionURL": "string",
  "actualUsers": 1,
  "allowDismiss": "string",
  "audienceType": 1,
  "channels": [
    "string"
  ],
  "content": "string",
  "coverImage": "string",
  "createdTime": 1,
  "creatorID": "string",
  "emailFormat": "string",
  "estimatedUsers": 1,
  "expireTime": 1,
  "group": "string",
  "id": "string",
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "priority": 1,
  "scheduleTime": 1,
  "sentTime": 1,
  "severity": "string",
  "showAsModal": "string",
  "status": 1,
  "targetUserIDs": [
    "string"
  ],
  "tenantCode": "string",
  "title": "string",
  "videoUrl": "string"
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
ur api /api/v1/system/notify/notification/create \
  --body '{"actionText": "string", "actionURL": "string", "actualUsers": 1, "allowDismiss": "string", "audienceType": 1, "channels": ["string"], "content": "string", "coverImage": "string", "createdTime": 1, "creatorID": "string", "emailFormat": "string", "estimatedUsers": 1, "expireTime": 1, "group": "string", "id": "string", "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "priority": 1, "scheduleTime": 1, "sentTime": 1, "severity": "string", "showAsModal": "string", "status": 1, "targetUserIDs": ["string"], "tenantCode": "string", "title": "string", "videoUrl": "string"}'
```

### POST `/api/v1/system/notify/notification/delete`

**说明**: 删除手动通知

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
ur api /api/v1/system/notify/notification/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/estimate-users`

**说明**: 预估送达用户数

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
    "estimatedCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/notification/estimate-users \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/index`

**说明**: 获取手动通知列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `group` | string | 否 | 分组筛选 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `priority` | integer | 否 | 优先级筛选 (格式: int64) |
| `status` | integer | 否 | 状态筛选 (格式: int64) |
| `tenantCode` | string | 否 | 租户编码筛选 |
| `title` | string | 否 | 标题模糊搜索 |

**请求示例**:
```json
{
  "group": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "priority": 1,
  "status": 1,
  "tenantCode": "string",
  "title": "string"
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
        "actualUsers": 1,
        "allowDismiss": "string",
        "audienceType": 1,
        "channels": [
          "string"
        ],
        "content": "string",
        "coverImage": "string",
        "createdTime": 1,
        "creatorID": "string",
        "emailFormat": "string",
        "estimatedUsers": 1,
        "expireTime": 1,
        "group": "string",
        "id": "string",
        "modalShowEndTime": "2026-01-01T00:00:00Z",
        "modalShowStartTime": "2026-01-01T00:00:00Z",
        "priority": 1,
        "scheduleTime": 1,
        "sentTime": 1,
        "severity": "string",
        "showAsModal": "string",
        "status": 1,
        "targetUserIDs": [
          "string"
        ],
        "tenantCode": "string",
        "title": "string",
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
ur api /api/v1/system/notify/notification/index \
  --body '{"group": "string", "page": {"page": 1, "pageSize": 1}, "priority": 1, "status": 1, "tenantCode": "string", "title": "string"}'
```

### POST `/api/v1/system/notify/notification/read`

**说明**: 获取手动通知详情

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
    "actionText": "string",
    "actionURL": "string",
    "actualUsers": 1,
    "allowDismiss": "string",
    "audienceType": 1,
    "channels": [
      "string"
    ],
    "content": "string",
    "coverImage": "string",
    "createdTime": 1,
    "creatorID": "string",
    "emailFormat": "string",
    "estimatedUsers": 1,
    "expireTime": 1,
    "group": "string",
    "id": "string",
    "modalShowEndTime": "2026-01-01T00:00:00Z",
    "modalShowStartTime": "2026-01-01T00:00:00Z",
    "priority": 1,
    "scheduleTime": 1,
    "sentTime": 1,
    "severity": "string",
    "showAsModal": "string",
    "status": 1,
    "targetUserIDs": [
      "string"
    ],
    "tenantCode": "string",
    "title": "string",
    "videoUrl": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/notification/read \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/revoke`

**说明**: 撤回手动通知

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
ur api /api/v1/system/notify/notification/revoke \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/send`

**说明**: 发送手动通知

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 通知ID |
| `scheduleTime` | integer | 否 | 计划发送时间,不填立即发送 (格式: int64) |

**请求示例**:
```json
{
  "id": "string",
  "scheduleTime": 1
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
ur api /api/v1/system/notify/notification/send \
  --body '{"id": "string", "scheduleTime": 1}'
```

### POST `/api/v1/system/notify/notification/update`

**说明**: 更新手动通知

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionText` | string | 否 | 操作按钮文案 |
| `actionURL` | string | 否 | 操作跳转URL |
| `actualUsers` | integer | 否 | 实际送达数,只读 (格式: int64) |
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `audienceType` | integer | 否 | 受众类型 1:指定用户 2:全局 (格式: int64) |
| `channels` | array[string] | 否 | 发送渠道列表 |
| `content` | string | 是 | 内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `createdTime` | integer | 否 | 创建时间,只读 (格式: int64) |
| `creatorID` | string | 否 | 创建者ID,只读 |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `estimatedUsers` | integer | 否 | 预计送达数,只读 (格式: int64) |
| `expireTime` | integer | 否 | 有效期(毫秒时间戳) (格式: int64) |
| `group` | string | 否 | 分组 |
| `id` | string | 否 | 通知ID |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `priority` | integer | 否 | 优先级 1:紧急 2:重要 3:普通 (格式: int64) |
| `scheduleTime` | integer | 否 | 计划发送时间(毫秒时间戳) (格式: int64) |
| `sentTime` | integer | 否 | 实际发送时间(毫秒时间戳),只读 (格式: int64) |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `status` | integer | 否 | 状态 1:草稿 2:已计划 3:发送中 4:已发送 5:已撤回 (格式: int64) |
| `targetUserIDs` | array[string] | 否 | 目标用户ID列表 |
| `tenantCode` | string | 否 | 租户编码,空=平台级 |
| `title` | string | 是 | 标题 |
| `videoUrl` | string | 否 |  news专属：视频URL |

**请求示例**:
```json
{
  "actionText": "string",
  "actionURL": "string",
  "actualUsers": 1,
  "allowDismiss": "string",
  "audienceType": 1,
  "channels": [
    "string"
  ],
  "content": "string",
  "coverImage": "string",
  "createdTime": 1,
  "creatorID": "string",
  "emailFormat": "string",
  "estimatedUsers": 1,
  "expireTime": 1,
  "group": "string",
  "id": "string",
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "priority": 1,
  "scheduleTime": 1,
  "sentTime": 1,
  "severity": "string",
  "showAsModal": "string",
  "status": 1,
  "targetUserIDs": [
    "string"
  ],
  "tenantCode": "string",
  "title": "string",
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
ur api /api/v1/system/notify/notification/update \
  --body '{"actionText": "string", "actionURL": "string", "actualUsers": 1, "allowDismiss": "string", "audienceType": 1, "channels": ["string"], "content": "string", "coverImage": "string", "createdTime": 1, "creatorID": "string", "emailFormat": "string", "estimatedUsers": 1, "expireTime": 1, "group": "string", "id": "string", "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "priority": 1, "scheduleTime": 1, "sentTime": 1, "severity": "string", "showAsModal": "string", "status": 1, "targetUserIDs": ["string"], "tenantCode": "string", "title": "string", "videoUrl": "string"}'
```
