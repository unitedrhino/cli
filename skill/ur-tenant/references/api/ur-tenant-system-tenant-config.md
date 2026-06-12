# ur-tenant system/tenant/config

获取企业配置 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/config/get-one` | 获取企业配置 | admin |
| POST | `/api/v1/system/tenant/config/update` | 更新企业配置 | admin |

## 详细说明

### POST `/api/v1/system/tenant/config/get-one`

**说明**: 获取企业配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `tenantCode` | string | 是 |  |

**请求示例**:
```json
{
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string",
    "registerRoleID": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/config/get-one \
  --body '{"tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/config/update`

**说明**: 更新企业配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  |
| `registerRoleID` | string | 否 | 注册分配的角色id |
| `tenantCode` | string | 否 | 企业编码 |

**请求示例**:
```json
{
  "id": "string",
  "registerRoleID": "string",
  "tenantCode": "string"
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
ur api /api/v1/system/tenant/config/update \
  --body '{"id": "string", "registerRoleID": "string", "tenantCode": "string"}'
```
