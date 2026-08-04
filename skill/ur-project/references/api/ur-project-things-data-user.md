# ur-project things/data/user

获取有区域权限的用户列表

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/data/user/get-list` | 获取有区域权限的用户列表 | admin |

## 详细说明

### POST `/api/v1/things/data/user/get-list`

**说明**: 获取有区域权限的用户列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaIDs` | array[string] | 是 |  区域ID列表（必填） |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `projectID` | string | 否 |  项目ID（可选，不传则用上下文） |

**请求示例**:
```json
{
  "areaIDs": [
    "string"
  ],
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
  "projectID": "string"
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
        "email": "string",
        "nickName": "string",
        "phone": "string",
        "userID": "string",
        "userName": "string"
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
ur api /api/v1/things/data/user/get-list \
  --body '{"areaIDs": ["string"], "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "projectID": "string"}'
```
