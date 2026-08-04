# ur-user system/dept/user

批量授权部门用户 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/dept/user/batch-create` | 批量授权部门用户 | admin |
| POST | `/api/v1/system/dept/user/batch-delete` | 批量取消授权部门用户 | admin |
| POST | `/api/v1/system/dept/user/get-list` | 获取部门授权列表 | admin |

## 详细说明

### POST `/api/v1/system/dept/user/batch-create`

**说明**: 批量授权部门用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authType` | integer | 是 | 授权类型: 1 管理权限,可以修改该部门及修改部门人员的授权 2 普通权限 (格式: int64) |
| `deptID` | string | 否 |  |
| `isAuthChildren` | integer | 是 | 同时授权子节点 (格式: int64) |
| `userIDs` | array[string] | 是 |  |

**请求示例**:
```json
{
  "authType": 1,
  "deptID": "string",
  "isAuthChildren": 1,
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
ur api /api/v1/system/dept/user/batch-create \
  --body '{"authType": 1, "deptID": "string", "isAuthChildren": 1, "userIDs": ["string"]}'
```

### POST `/api/v1/system/dept/user/batch-delete`

**说明**: 批量取消授权部门用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 是 |  |
| `ids` | array[string] | 是 |  |

**请求示例**:
```json
{
  "deptID": "string",
  "ids": [
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
ur api /api/v1/system/dept/user/batch-delete \
  --body '{"deptID": "string", "ids": ["string"]}'
```

### POST `/api/v1/system/dept/user/get-list`

**说明**: 获取部门授权列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "deptID": "string",
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
        "authType": 1,
        "deptID": "string",
        "deptIdPath": "string",
        "id": "string",
        "isAuthChildren": 1,
        "user": {
          "avatar": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "email": "string",
          "lastIP": "string",
          "nickName": "string",
          "phone": "string",
          "regIP": "string",
          "userID": "string",
          "userName": "string"
        },
        "userID": "string"
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
ur api /api/v1/system/dept/user/get-list \
  --body '{"deptID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```
