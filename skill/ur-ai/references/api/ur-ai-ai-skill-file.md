# ur-ai ai/skill/file

创建技能文件 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/skill/file/create` | 创建技能文件 | admin |
| POST | `/api/v1/ai/skill/file/delete` | 删除文件 | admin |
| POST | `/api/v1/ai/skill/file/get-one` | 获取文件内容 | admin |
| POST | `/api/v1/ai/skill/file/update` | 更新文件内容 | admin |

## 详细说明

### POST `/api/v1/ai/skill/file/create`

**说明**: 创建技能文件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 |  文件内容 |
| `isDir` | boolean | 是 |  是否为目录 (格式: boolean) |
| `path` | string | 是 |  文件路径 |
| `skillID` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "content": "string",
  "isDir": true,
  "path": "string",
  "skillID": 1
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
ur api /api/v1/ai/skill/file/create \
  --body '{"content": "string", "isDir": true, "path": "string", "skillID": 1}'
```

### POST `/api/v1/ai/skill/file/delete`

**说明**: 删除文件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `path` | string | 是 |  文件路径 |
| `skillID` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "path": "string",
  "skillID": 1
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
ur api /api/v1/ai/skill/file/delete \
  --body '{"path": "string", "skillID": 1}'
```

### POST `/api/v1/ai/skill/file/get-one`

**说明**: 获取文件内容

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `path` | string | 是 |  文件路径 |
| `skillID` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "path": "string",
  "skillID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "content": "string",
    "path": "string",
    "size": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/file/get-one \
  --body '{"path": "string", "skillID": 1}'
```

### POST `/api/v1/ai/skill/file/update`

**说明**: 更新文件内容

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 |  文件内容 |
| `path` | string | 是 |  文件路径 |
| `skillID` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "content": "string",
  "path": "string",
  "skillID": 1
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
ur api /api/v1/ai/skill/file/update \
  --body '{"content": "string", "path": "string", "skillID": 1}'
```
