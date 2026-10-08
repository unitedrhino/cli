# ur-user system/dept/role

批量添加部门角色 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/dept/role/batch-create` | 批量添加部门角色 | admin |
| POST | `/api/v1/system/dept/role/batch-delete` | 批量移除部门角色 | admin |
| POST | `/api/v1/system/dept/role/get-list` | 获取部门角色列表 | admin |

## 详细说明

### POST `/api/v1/system/dept/role/batch-create`

**说明**: 批量添加部门角色

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 是 |  部门ID |
| `roleIDs` | array[string] | 是 |  角色ID列表 |

**请求示例**:
```json
{
  "deptID": "string",
  "roleIDs": [
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
ur api /api/v1/system/dept/role/batch-create \
  --body '{"deptID": "string", "roleIDs": ["string"]}'
```

### POST `/api/v1/system/dept/role/batch-delete`

**说明**: 批量移除部门角色

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 是 |  部门ID |
| `roleIDs` | array[string] | 是 |  角色ID列表 |

**请求示例**:
```json
{
  "deptID": "string",
  "roleIDs": [
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
ur api /api/v1/system/dept/role/batch-delete \
  --body '{"deptID": "string", "roleIDs": ["string"]}'
```

### POST `/api/v1/system/dept/role/get-list`

**说明**: 获取部门角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 是 |  部门ID |

**请求示例**:
```json
{
  "deptID": "string"
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
        "desc": "string",
        "id": "string",
        "isFullAuth": 1,
        "name": "示例名称",
        "status": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dept/role/get-list \
  --body '{"deptID": "string"}'
```
