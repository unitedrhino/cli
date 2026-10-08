# ur-user system/role/app

更新APP权限 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/role/app/batch-update` | 更新APP权限 | admin |
| POST | `/api/v1/system/role/app/get-list` | 获取APP权限列表 | admin |

## 详细说明

### POST `/api/v1/system/role/app/batch-update`

**说明**: 更新APP权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appIDs` | array[string] | 是 | App列表数据 |
| `id` | string | 是 | 角色ID |

**请求示例**:
```json
{
  "appIDs": [
    "string"
  ],
  "id": "string"
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
ur api /api/v1/system/role/app/batch-update \
  --body '{"appIDs": ["string"], "id": "string"}'
```

### POST `/api/v1/system/role/app/get-list`

**说明**: 获取APP权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 角色ID |

**请求示例**:
```json
{
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "appIDs": [
      "string"
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/role/app/get-list \
  --body '{"id": "string"}'
```
