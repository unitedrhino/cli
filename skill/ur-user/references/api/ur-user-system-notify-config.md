# ur-user system/notify/config

添加通知配置 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/notify/config/create` | 添加通知配置 | platform |
| POST | `/api/v1/system/notify/config/delete` | 删除通知配置 | platform |
| POST | `/api/v1/system/notify/config/get-list` | 获取通知列表 | platform |
| POST | `/api/v1/system/notify/config/get-one` | 获取通知详情 | platform |
| POST | `/api/v1/system/notify/config/send-test` | 测试发送通知 | platform |
| POST | `/api/v1/system/notify/config/tree` | 获取通知列表树 | platform |
| POST | `/api/v1/system/notify/config/update` | 更新通知配置 | platform |

## 详细说明

### POST `/api/v1/system/notify/config/create`

**说明**: 添加通知配置

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  通知类型编码 |
| `desc` | string | 否 |  备注 |
| `group` | string | 否 | 分组 |
| `id` | string | 否 |  id编号 |
| `isEnabled` | integer | 否 | 是否启用 1:启用 2:禁用 (格式: int64) |
| `isRecord` | integer | 否 | 是否记录该消息,是的情况下会将消息存一份到消息中心 (格式: int64) |
| `name` | string | 否 | 通知的命名 |
| `params` | object | 否 | 变量属性 key是变量参数,value是变量描述 |
| `supportTypes` | array[string] | 否 | 支持的通知类型 sms(短信),email(邮箱),dingTalk(钉钉机器人),dingWebhook(钉钉webhook),wxEWebHook(企业微信webhook) |
| `templates` | array[NotifyConfigTemplate] | 否 | 传withTemplates的时候返回,只读 |
| `throttle` | integer | 否 | 防抖间隔秒数,0表示不防抖 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "group": "string",
  "id": "string",
  "isEnabled": 1,
  "isRecord": 1,
  "name": "string",
  "params": {},
  "supportTypes": [
    "string"
  ],
  "templates": [
    {
      "id": "string",
      "notifyCode": "string",
      "templateID": "string",
      "type": "string"
    }
  ],
  "throttle": 1
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
ur api /api/v1/system/notify/config/create \
  --body '{"code": "string", "desc": "string", "group": "string", "id": "string", "isEnabled": 1, "isRecord": 1, "name": "string", "params": {}, "supportTypes": ["string"], "templates": [{"id": "string", "notifyCode": "string", "templateID": "string", "type": "string"}], "throttle": 1}'
```

### POST `/api/v1/system/notify/config/delete`

**说明**: 删除通知配置

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
ur api /api/v1/system/notify/config/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/config/get-list`

**说明**: 获取通知列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  应用编号 |
| `group` | string | 否 | 分组 |
| `isEnabled` | integer | 否 | 是否启用 1:启用 2:禁用 (格式: int64) |
| `name` | string | 否 |  应用名称 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `withTemplates` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "code": "string",
  "group": "string",
  "isEnabled": 1,
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "withTemplates": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": "string",
        "desc": "string",
        "group": "string",
        "id": "string",
        "isEnabled": 1,
        "isRecord": 1,
        "name": "string",
        "params": {},
        "supportTypes": [
          "string"
        ],
        "templates": [
          {
            "id": "string",
            "notifyCode": "string",
            "templateID": "string",
            "type": "string"
          }
        ],
        "throttle": 1
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
ur api /api/v1/system/notify/config/get-list \
  --body '{"code": "string", "group": "string", "isEnabled": 1, "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "withTemplates": true}'
```

### POST `/api/v1/system/notify/config/get-one`

**说明**: 获取通知详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "code": "string",
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "desc": "string",
    "group": "string",
    "id": "string",
    "isEnabled": 1,
    "isRecord": 1,
    "name": "string",
    "params": {},
    "supportTypes": [
      "string"
    ],
    "templates": [
      {
        "id": "string",
        "notifyCode": "string",
        "templateID": "string",
        "type": "string"
      }
    ],
    "throttle": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/config/get-one \
  --body '{"code": "string", "id": "string"}'
```

### POST `/api/v1/system/notify/config/send-test`

**说明**: 测试发送通知

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accounts` | array[string] | 否 | 指定账号列表 |
| `notifyCode` | string | 是 | 通知编码 |
| `params` | object | 否 | 模版变量 |
| `str1` | string | 否 |  |
| `str2` | string | 否 |  |
| `str3` | string | 否 |  |
| `templateID` | string | 否 | 指定模版ID |
| `type` | string | 是 | 渠道类型 sms\|email\|dingTalk\|message等 |

**请求示例**:
```json
{
  "accounts": [
    "string"
  ],
  "notifyCode": "string",
  "params": {},
  "str1": "string",
  "str2": "string",
  "str3": "string",
  "templateID": "string",
  "type": "string"
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
ur api /api/v1/system/notify/config/send-test \
  --body '{"accounts": ["string"], "notifyCode": "string", "params": {}, "str1": "string", "str2": "string", "str3": "string", "templateID": "string", "type": "string"}'
```

### POST `/api/v1/system/notify/config/tree`

**说明**: 获取通知列表树

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  应用编号 |
| `group` | string | 否 | 分组 |
| `isEnabled` | integer | 否 | 是否启用 1:启用 2:禁用 (格式: int64) |
| `name` | string | 否 |  应用名称 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `withTemplates` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "code": "string",
  "group": "string",
  "isEnabled": 1,
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "withTemplates": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "children": [
          {
            "code": "string",
            "desc": "string",
            "group": "string",
            "id": "string",
            "isEnabled": 1,
            "isRecord": 1,
            "name": "string",
            "params": {},
            "supportTypes": [],
            "templates": [],
            "throttle": 1
          }
        ],
        "code": "string",
        "id": "string",
        "name": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/config/tree \
  --body '{"code": "string", "group": "string", "isEnabled": 1, "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "withTemplates": true}'
```

### POST `/api/v1/system/notify/config/update`

**说明**: 更新通知配置

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  通知类型编码 |
| `desc` | string | 否 |  备注 |
| `group` | string | 否 | 分组 |
| `id` | string | 否 |  id编号 |
| `isEnabled` | integer | 否 | 是否启用 1:启用 2:禁用 (格式: int64) |
| `isRecord` | integer | 否 | 是否记录该消息,是的情况下会将消息存一份到消息中心 (格式: int64) |
| `name` | string | 否 | 通知的命名 |
| `params` | object | 否 | 变量属性 key是变量参数,value是变量描述 |
| `supportTypes` | array[string] | 否 | 支持的通知类型 sms(短信),email(邮箱),dingTalk(钉钉机器人),dingWebhook(钉钉webhook),wxEWebHook(企业微信webhook) |
| `templates` | array[NotifyConfigTemplate] | 否 | 传withTemplates的时候返回,只读 |
| `throttle` | integer | 否 | 防抖间隔秒数,0表示不防抖 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "group": "string",
  "id": "string",
  "isEnabled": 1,
  "isRecord": 1,
  "name": "string",
  "params": {},
  "supportTypes": [
    "string"
  ],
  "templates": [
    {
      "id": "string",
      "notifyCode": "string",
      "templateID": "string",
      "type": "string"
    }
  ],
  "throttle": 1
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
ur api /api/v1/system/notify/config/update \
  --body '{"code": "string", "desc": "string", "group": "string", "id": "string", "isEnabled": 1, "isRecord": 1, "name": "string", "params": {}, "supportTypes": ["string"], "templates": [{"id": "string", "notifyCode": "string", "templateID": "string", "type": "string"}], "throttle": 1}'
```
