# ur-user system/user/dept

新增用户的部门列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/dept/batch-create` | 新增用户的部门列表 | admin |
| POST | `/api/v1/system/user/dept/batch-delete` | 删除用户的部门列表 | admin |

## 详细说明

### POST `/api/v1/system/user/dept/batch-create`

**说明**: 新增用户的部门列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptIDs` | array[string] | 是 | 角色编号列表 |
| `userID` | string | 是 |  用户id |

**请求示例**:
```json
{
  "deptIDs": [
    "string"
  ],
  "userID": "string"
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
ur api /api/v1/system/user/dept/batch-create \
  --body '{"deptIDs": ["string"], "userID": "string"}'
```

### POST `/api/v1/system/user/dept/batch-delete`

**说明**: 删除用户的部门列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptIDs` | array[string] | 是 | 角色编号列表 |
| `userID` | string | 是 |  用户id |

**请求示例**:
```json
{
  "deptIDs": [
    "string"
  ],
  "userID": "string"
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
ur api /api/v1/system/user/dept/batch-delete \
  --body '{"deptIDs": ["string"], "userID": "string"}'
```
