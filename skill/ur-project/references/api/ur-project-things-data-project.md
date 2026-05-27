# ur-project things/data/project

批量创建授权项目权限 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/data/project/batch-create` | 批量创建授权项目权限 | admin |
| POST | `/api/v1/things/data/project/batch-delete` | 批量删除授权项目权限 | admin |
| POST | `/api/v1/things/data/project/create` | 创建授权项目权限 | admin |
| POST | `/api/v1/things/data/project/delete` | 删除授权项目权限 | admin |
| POST | `/api/v1/things/data/project/get-list` | 获取项目权限列表 | admin |

## 详细说明

### POST `/api/v1/things/data/project/batch-create`

**说明**: 批量创建授权项目权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authType` | integer | 是 |  1:管理权限,可以修改别人的权限,及读写权限 2:读写权限,可以读写该权限 3:读权限,只能读,不能修改 (格式: int64) |
| `projectID` | string | 是 | 权限数据ID |
| `targetIDs` | array[string] | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "authType": 1,
  "projectID": "string",
  "targetIDs": [
    "string"
  ],
  "targetType": "string"
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
ur api /api/v1/things/data/project/batch-create \
  --body '{"authType": 1, "projectID": "string", "targetIDs": ["string"], "targetType": "string"}'
```

### POST `/api/v1/things/data/project/batch-delete`

**说明**: 批量删除授权项目权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `projectID` | string | 否 | 权限数据ID |
| `targetIDs` | array[string] | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "projectID": "string",
  "targetIDs": [
    "string"
  ],
  "targetType": "string"
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
ur api /api/v1/things/data/project/batch-delete \
  --body '{"projectID": "string", "targetIDs": ["string"], "targetType": "string"}'
```

### POST `/api/v1/things/data/project/create`

**说明**: 创建授权项目权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authType` | integer | 是 |  1:管理权限,可以修改别人的权限,及读写权限 2:读写权限,可以读写该权限 3:读权限,只能读,不能修改 (格式: int64) |
| `projectID` | string | 是 | 权限数据ID |
| `targetID` | string | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "authType": 1,
  "projectID": "string",
  "targetID": "string",
  "targetType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "authType": 1,
    "projectID": "string",
    "targetID": "string",
    "updatedTime": 1,
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
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/data/project/create \
  --body '{"authType": 1, "projectID": "string", "targetID": "string", "targetType": "string"}'
```

### POST `/api/v1/things/data/project/delete`

**说明**: 删除授权项目权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `projectID` | string | 否 | 权限数据ID |
| `targetID` | string | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "projectID": "string",
  "targetID": "string",
  "targetType": "string"
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
ur api /api/v1/things/data/project/delete \
  --body '{"projectID": "string", "targetID": "string", "targetType": "string"}'
```

### POST `/api/v1/things/data/project/get-list`

**说明**: 获取项目权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `projectID` | string | 否 | 权限数据ID |
| `targetID` | string | 否 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "projectID": "string",
  "targetID": "string",
  "targetType": "string"
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
        "projectID": "string",
        "targetID": "string",
        "updatedTime": 1,
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
        }
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
ur api /api/v1/things/data/project/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}, "projectID": "string", "targetID": "string", "targetType": "string"}'
```
