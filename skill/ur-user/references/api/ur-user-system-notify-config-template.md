# ur-user system/notify/config/template

删除通知绑定模版 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/notify/config/template/delete` | 删除通知绑定模版 | platform |
| POST | `/api/v1/system/notify/config/template/get-list` | 获取通知绑定模版列表 | platform |
| POST | `/api/v1/system/notify/config/template/update` | 更新通知绑定模版 | platform |

## 详细说明

### POST `/api/v1/system/notify/config/template/delete`

**说明**: 删除通知绑定模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `notifyCode` | string | 是 |  |
| `type` | string | 是 |  |

**请求示例**:
```json
{
  "notifyCode": "string",
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
ur api /api/v1/system/notify/config/template/delete \
  --body '{"notifyCode": "string", "type": "string"}'
```

### POST `/api/v1/system/notify/config/template/get-list`

**说明**: 获取通知绑定模版列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `notifyCode` | string | 是 |  |
| `type` | string | 否 |  |

**请求示例**:
```json
{
  "notifyCode": "string",
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
        "id": "string",
        "notifyCode": "string",
        "templateID": "string",
        "type": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/config/template/get-list \
  --body '{"notifyCode": "string", "type": "string"}'
```

### POST `/api/v1/system/notify/config/template/update`

**说明**: 更新通知绑定模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  |
| `notifyCode` | string | 否 |  |
| `templateID` | string | 否 |  |
| `type` | string | 否 |  |

**请求示例**:
```json
{
  "id": "string",
  "notifyCode": "string",
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
ur api /api/v1/system/notify/config/template/update \
  --body '{"id": "string", "notifyCode": "string", "templateID": "string", "type": "string"}'
```
