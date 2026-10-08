# ur-project things/area/profile

获取区域配置列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/area/profile/get-list` | 获取区域配置列表 | admin |
| POST | `/api/v1/things/area/profile/get-one` | 获取区域配置详情 | admin |
| POST | `/api/v1/things/area/profile/update` | 更新区域配置 | admin |

## 详细说明

### POST `/api/v1/things/area/profile/get-list`

**说明**: 获取区域配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 是 | 项目区域id（只读） |
| `codes` | array[string] | 是 |  |

**请求示例**:
```json
{
  "areaID": "string",
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
        "areaID": "string",
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
ur api /api/v1/things/area/profile/get-list \
  --body '{"areaID": "string", "codes": ["string"]}'
```

### POST `/api/v1/things/area/profile/get-one`

**说明**: 获取区域配置详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 是 | 项目区域id（只读） |
| `code` | string | 是 |  |

**请求示例**:
```json
{
  "areaID": "string",
  "code": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "areaID": "string",
    "code": "string",
    "params": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/area/profile/get-one \
  --body '{"areaID": "string", "code": "string"}'
```

### POST `/api/v1/things/area/profile/update`

**说明**: 更新区域配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 是 | 项目区域id（只读） |
| `code` | string | 是 |  |
| `params` | string | 是 |  |

**请求示例**:
```json
{
  "areaID": "string",
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
ur api /api/v1/things/area/profile/update \
  --body '{"areaID": "string", "code": "string", "params": "string"}'
```
