# ur-protocol things/protocol/service

删除自定义协议服务器 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/protocol/service/delete` | 删除自定义协议服务器 | admin |
| POST | `/api/v1/things/protocol/service/get-list` | 获取自定义协议服务器信息列表 | admin |

## 详细说明

### POST `/api/v1/things/protocol/service/delete`

**说明**: 删除自定义协议服务器

**权限**: admin

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
ur api /api/v1/things/protocol/service/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/protocol/service/get-list`

**说明**: 获取自定义协议服务器信息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
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
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "ip": "string",
        "port": 1,
        "status": 1,
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
ur api /api/v1/things/protocol/service/get-list \
  --body '{"code": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```
