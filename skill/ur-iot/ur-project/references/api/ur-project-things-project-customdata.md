# ur-project things/project/customData

新增项目crud 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/project/crud/create` | 新增项目crud | admin |
| POST | `/api/v1/things/project/crud/update` | 更新项目crud | admin |

## 详细说明

### POST `/api/v1/things/project/crud/create`

**说明**: 新增项目crud

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 只读 |
| `id` | string | 否 |  |
| `params` | string | 否 | 普通业务参数,如果是文件类型,则需要以 Img,File结尾,后端会进行文件的处理,上传方式可以参考设备图片 |
| `purpose` | string | 否 | 用途 新增必填 |
| `sort` | integer | 否 | 排序 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": "string",
  "id": "string",
  "params": "string",
  "purpose": "string",
  "sort": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/crud/create \
  --body '{"createdTime": "string", "id": "string", "params": "string", "purpose": "string", "sort": 1}'
```

### POST `/api/v1/things/project/crud/update`

**说明**: 更新项目crud

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 只读 |
| `id` | string | 否 |  |
| `params` | string | 否 | 普通业务参数,如果是文件类型,则需要以 Img,File结尾,后端会进行文件的处理,上传方式可以参考设备图片 |
| `purpose` | string | 否 | 用途 新增必填 |
| `sort` | integer | 否 | 排序 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": "string",
  "id": "string",
  "params": "string",
  "purpose": "string",
  "sort": 1
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
ur api /api/v1/things/project/crud/update \
  --body '{"createdTime": "string", "id": "string", "params": "string", "purpose": "string", "sort": 1}'
```
