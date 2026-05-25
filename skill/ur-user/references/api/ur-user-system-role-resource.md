# ur-user system/role/resource

批量更新角色资源动作权限 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/role/resource/batch-update` | 批量更新角色资源动作权限 | admin |
| POST | `/api/v1/system/role/resource/get-list` | 获取角色资源动作权限列表 | admin |

## 详细说明

### POST `/api/v1/system/role/resource/batch-update`

**说明**: 批量更新角色资源动作权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `list` | array[RoleResourceActionItem] | 是 |  完整权限列表（全量替换） |
| `roleId` | string | 是 |  角色ID |

**请求示例**:
```json
{
  "list": [
    {
      "action": "string",
      "resource": "string"
    }
  ],
  "roleId": "string"
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
ur api /api/v1/system/role/resource/batch-update \
  --body '{"list": [{"action": "string", "resource": "string"}], "roleId": "string"}'
```

### POST `/api/v1/system/role/resource/get-list`

**说明**: 获取角色资源动作权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `resource` | string | 否 |  可选资源筛选 |
| `roleId` | string | 是 |  角色ID |

**请求示例**:
```json
{
  "resource": "string",
  "roleId": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "action": "string",
        "resource": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/role/resource/get-list \
  --body '{"resource": "string", "roleId": "string"}'
```
