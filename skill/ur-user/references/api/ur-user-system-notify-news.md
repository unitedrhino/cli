# ur-user system/notify/news

创建资讯 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/notify/news/create` | 创建资讯 | platform |
| POST | `/api/v1/system/notify/news/delete` | 删除资讯 | platform |
| POST | `/api/v1/system/notify/news/info` | 获取资讯详情 | platform |
| POST | `/api/v1/system/notify/news/list` | 获取资讯列表 | platform |
| POST | `/api/v1/system/notify/news/update` | 更新资讯 | platform |

## 详细说明

### POST `/api/v1/system/notify/news/create`

**说明**: 创建资讯

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionURL` | string | 否 |  跳转链接 |
| `body` | string | 是 |  正文内容（HTML格式） |
| `coverImage` | string | 否 |  封面图URL |
| `createdTime` | integer | 否 |  创建时间，只读 (格式: int64) |
| `expireTime` | string | 否 |  过期时间（毫秒时间戳） |
| `id` | string | 否 |  资讯ID |
| `notifyTime` | string | 否 |  发布时间（毫秒时间戳） |
| `status` | string | 否 |  状态 1:草稿 2:已发布 3:已下架 |
| `summary` | string | 否 |  摘要 |
| `title` | string | 是 |  标题 |

**请求示例**:
```json
{
  "actionURL": "string",
  "body": "string",
  "coverImage": "string",
  "createdTime": 1,
  "expireTime": "2026-01-01T00:00:00Z",
  "id": "string",
  "notifyTime": "2026-01-01T00:00:00Z",
  "status": "string",
  "summary": "string",
  "title": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/news/create \
  --body '{"actionURL": "string", "body": "string", "coverImage": "string", "createdTime": 1, "expireTime": "2026-01-01T00:00:00Z", "id": "string", "notifyTime": "2026-01-01T00:00:00Z", "status": "string", "summary": "string", "title": "string"}'
```

### POST `/api/v1/system/notify/news/delete`

**说明**: 删除资讯

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
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
ur api /api/v1/system/notify/news/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/notify/news/info`

**说明**: 获取资讯详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "actionURL": "string",
    "body": "string",
    "coverImage": "string",
    "createdTime": 1,
    "expireTime": "2026-01-01T00:00:00Z",
    "id": "string",
    "notifyTime": "2026-01-01T00:00:00Z",
    "status": "string",
    "summary": "string",
    "title": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/news/info \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/notify/news/list`

**说明**: 获取资讯列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `notifyTimeEnd` | integer | 否 |  发布时间结束 (格式: int64) |
| `notifyTimeStart` | integer | 否 |  发布时间开始 (格式: int64) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态筛选 |
| `title` | string | 否 |  标题模糊搜索 |

**请求示例**:
```json
{
  "notifyTimeEnd": 1,
  "notifyTimeStart": 1,
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
  },
  "status": "string",
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
        "actionURL": "string",
        "body": "string",
        "coverImage": "string",
        "createdTime": 1,
        "expireTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "notifyTime": "2026-01-01T00:00:00Z",
        "status": "string",
        "summary": "string",
        "title": "string"
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
ur api /api/v1/system/notify/news/list \
  --body '{"notifyTimeEnd": 1, "notifyTimeStart": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": "string", "title": "string"}'
```

### POST `/api/v1/system/notify/news/update`

**说明**: 更新资讯

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionURL` | string | 否 |  跳转链接 |
| `body` | string | 是 |  正文内容（HTML格式） |
| `coverImage` | string | 否 |  封面图URL |
| `createdTime` | integer | 否 |  创建时间，只读 (格式: int64) |
| `expireTime` | string | 否 |  过期时间（毫秒时间戳） |
| `id` | string | 否 |  资讯ID |
| `notifyTime` | string | 否 |  发布时间（毫秒时间戳） |
| `status` | string | 否 |  状态 1:草稿 2:已发布 3:已下架 |
| `summary` | string | 否 |  摘要 |
| `title` | string | 是 |  标题 |

**请求示例**:
```json
{
  "actionURL": "string",
  "body": "string",
  "coverImage": "string",
  "createdTime": 1,
  "expireTime": "2026-01-01T00:00:00Z",
  "id": "string",
  "notifyTime": "2026-01-01T00:00:00Z",
  "status": "string",
  "summary": "string",
  "title": "string"
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
ur api /api/v1/system/notify/news/update \
  --body '{"actionURL": "string", "body": "string", "coverImage": "string", "createdTime": 1, "expireTime": "2026-01-01T00:00:00Z", "id": "string", "notifyTime": "2026-01-01T00:00:00Z", "status": "string", "summary": "string", "title": "string"}'
```
