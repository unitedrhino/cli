# ur-ai things/alarm/notifyTemplate

创建通知内容模板 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/notify-template/create` | 创建通知内容模板 | admin |
| POST | `/api/v1/things/alarm/notify-template/delete` | 删除通知内容模板 | admin |
| POST | `/api/v1/things/alarm/notify-template/get-list` | 获取通知内容模板列表 | admin |
| POST | `/api/v1/things/alarm/notify-template/get-one` | 获取通知内容模板详情 | admin |
| POST | `/api/v1/things/alarm/notify-template/test-send` | 测试发送通知 | admin |
| POST | `/api/v1/things/alarm/notify-template/update` | 更新通知内容模板 | admin |

## 详细说明

### POST `/api/v1/things/alarm/notify-template/create`

**说明**: 创建通知内容模板

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `channel` | string | 是 |  渠道 |
| `desc` | string | 否 |  描述 |
| `headers` | object | 否 |  请求头 |
| `name` | string | 是 |  名称 |
| `recoveryBody` | string | 否 |  恢复内容 |
| `subject` | string | 否 |  主题 |
| `triggerBody` | string | 否 |  触发内容 |

**请求示例**:
```json
{
  "channel": "string",
  "desc": "string",
  "headers": {},
  "name": "示例名称",
  "recoveryBody": "string",
  "subject": "string",
  "triggerBody": "string"
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
ur api /api/v1/things/alarm/notify-template/create \
  --body '{"channel": "string", "desc": "string", "headers": {}, "name": "示例名称", "recoveryBody": "string", "subject": "string", "triggerBody": "string"}'
```

### POST `/api/v1/things/alarm/notify-template/delete`

**说明**: 删除通知内容模板

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  ID |

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
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/notify-template/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/notify-template/get-list`

**说明**: 获取通知内容模板列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `channel` | string | 否 |  渠道 |
| `keyword` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  状态 (格式: int64) |

**请求示例**:
```json
{
  "channel": "string",
  "keyword": "示例名称",
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
  "status": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "channel": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "headers": {},
        "id": "string",
        "name": "示例名称",
        "recoveryBody": "string",
        "refCount": 1,
        "status": 1,
        "subject": "string",
        "triggerBody": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/alarm/notify-template/get-list \
  --body '{"channel": "string", "keyword": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1}'
```

### POST `/api/v1/things/alarm/notify-template/get-one`

**说明**: 获取通知内容模板详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  ID |

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
    "channel": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "headers": {},
    "id": "string",
    "name": "示例名称",
    "recoveryBody": "string",
    "refCount": 1,
    "status": 1,
    "subject": "string",
    "triggerBody": "string",
    "updatedTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/notify-template/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/notify-template/test-send`

**说明**: 测试发送通知

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  模板ID |
| `targets` | object | 否 |  |
| `targets.groupIDs` | array[string] | 否 |  用户组ID列表 |
| `targets.roleIDs` | array[string] | 否 |  角色ID列表 |
| `targets.userIDs` | array[string] | 否 |  用户ID列表 |
| `timing` | string | 是 |  trigger / recovery |

**请求示例**:
```json
{
  "id": "string",
  "targets": {
    "groupIDs": [
      "string"
    ],
    "roleIDs": [
      "string"
    ],
    "userIDs": [
      "string"
    ]
  },
  "timing": "string"
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
ur api /api/v1/things/alarm/notify-template/test-send \
  --body '{"id": "string", "targets": {"groupIDs": ["string"], "roleIDs": ["string"], "userIDs": ["string"]}, "timing": "string"}'
```

### POST `/api/v1/things/alarm/notify-template/update`

**说明**: 更新通知内容模板

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `channel` | string | 否 |  渠道 |
| `desc` | string | 否 |  描述 |
| `headers` | object | 否 |  请求头 |
| `id` | string | 是 |  ID |
| `name` | string | 否 |  名称 |
| `recoveryBody` | string | 否 |  恢复内容 |
| `status` | integer | 否 |  状态 (格式: int64) |
| `subject` | string | 否 |  主题 |
| `triggerBody` | string | 否 |  触发内容 |

**请求示例**:
```json
{
  "channel": "string",
  "desc": "string",
  "headers": {},
  "id": "string",
  "name": "示例名称",
  "recoveryBody": "string",
  "status": 1,
  "subject": "string",
  "triggerBody": "string"
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
ur api /api/v1/things/alarm/notify-template/update \
  --body '{"channel": "string", "desc": "string", "headers": {}, "id": "string", "name": "示例名称", "recoveryBody": "string", "status": 1, "subject": "string", "triggerBody": "string"}'
```
