# ur-project things/project/profile

获取项目配置列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/project/profile/get-list` | 获取项目配置列表 | admin |
| POST | `/api/v1/things/project/profile/get-one` | 获取项目配置详情 | admin |
| POST | `/api/v1/things/project/profile/update` | 更新项目配置 | admin |

## 详细说明

### POST `/api/v1/things/project/profile/get-list`

**说明**: 获取项目配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `codes` | array[string] | 是 |  |

**请求示例**:
```json
{
  "codes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "profiles": [
      {
        "code": "string",
        "params": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/profile/get-list \
  --body '{"codes": ["string"]}'
```

### POST `/api/v1/things/project/profile/get-one`

**说明**: 获取项目配置详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "params": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/profile/get-one \
  --body '{"code": "string"}'
```

### POST `/api/v1/things/project/profile/update`

**说明**: 更新项目配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `params` | string | 是 |  |

**请求示例**:
```json
{
  "code": "string",
  "params": "string"
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
ur api /api/v1/things/project/profile/update \
  --body '{"code": "string", "params": "string"}'
```
