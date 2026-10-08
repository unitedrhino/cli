# ur-user system/role/menu

更新角色对应菜单列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/role/menu/batch-update` | 更新角色对应菜单列表 | admin |
| POST | `/api/v1/system/role/menu/get-list` | 获取角色对应菜单列表 | admin |

## 详细说明

### POST `/api/v1/system/role/menu/batch-update`

**说明**: 更新角色对应菜单列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  应用编号 |
| `id` | string | 是 | 角色编号 |
| `menuIDs` | array[string] | 是 | 菜单编号列表 |

**请求示例**:
```json
{
  "appID": "string",
  "id": "string",
  "menuIDs": [
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
ur api /api/v1/system/role/menu/batch-update \
  --body '{"appID": "string", "id": "string", "menuIDs": ["string"]}'
```

### POST `/api/v1/system/role/menu/get-list`

**说明**: 获取角色对应菜单列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  应用编号 |
| `id` | string | 是 | 角色编号 |

**请求示例**:
```json
{
  "appID": "string",
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "menuIDs": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/role/menu/get-list \
  --body '{"appID": "string", "id": "string"}'
```
