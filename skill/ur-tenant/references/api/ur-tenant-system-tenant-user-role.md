# ur-tenant system/tenant/user/role

更新企业用户的角色列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/user/role/batch-update` | 更新企业用户的角色列表 | admin |
| POST | `/api/v1/system/tenant/user/role/get-list` | 获取企业用户角色列表 | admin |

## 详细说明

### POST `/api/v1/system/tenant/user/role/batch-update`

**说明**: 更新企业用户的角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `roleIDs` | array[string] | 是 |  角色编号列表 |
| `tenantCode` | string | 否 |  企业编码（可选，不填取当前用户企业） |
| `userID` | string | 是 |  用户ID |

**请求示例**:
```json
{
  "roleIDs": [
    "string"
  ],
  "tenantCode": "string",
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
ur api /api/v1/system/tenant/user/role/batch-update \
  --body '{"roleIDs": ["string"], "tenantCode": "string", "userID": "string"}'
```

### POST `/api/v1/system/tenant/user/role/get-list`

**说明**: 获取企业用户角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `tenantCode` | string | 否 |  企业编码（可选，不填取当前用户企业） |
| `userID` | string | 是 |  用户ID |

**请求示例**:
```json
{
  "tenantCode": "string",
  "userID": "string"
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
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/user/role/get-list \
  --body '{"tenantCode": "string", "userID": "string"}'
```
