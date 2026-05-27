# ur-ota things/ota/firmware/info

添加升级包 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/ota/firmware/info/create` | 添加升级包 | admin |
| POST | `/api/v1/things/ota/firmware/info/delete` | 删除升级包 | admin |
| POST | `/api/v1/things/ota/firmware/info/get-list` | 升级包列表 | admin |
| POST | `/api/v1/things/ota/firmware/info/get-one` | 查询升级包 | admin |
| POST | `/api/v1/things/ota/firmware/info/update` | 修改升级包 | admin |

## 详细说明

### POST `/api/v1/things/ota/firmware/info/create`

**说明**: 添加升级包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `extra` | string | 否 |  |
| `filePaths` | array[string] | 是 |  |
| `isDiff` | integer | 是 | 格式: int64 |
| `isNeedToVerify` | integer | 是 | 格式: int64 |
| `moduleCode` | string | 否 |  |
| `name` | string | 是 |  |
| `productID` | string | 是 |  |
| `signMethod` | string | 是 |  |
| `srcVersion` | string | 是 |  |
| `version` | string | 是 |  |

**请求示例**:
```json
{
  "desc": "string",
  "extra": "string",
  "filePaths": [
    "string"
  ],
  "isDiff": 1,
  "isNeedToVerify": 1,
  "moduleCode": "string",
  "name": "string",
  "productID": "string",
  "signMethod": "string",
  "srcVersion": "string",
  "version": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/ota/firmware/info/create \
  --body '{"desc": "string", "extra": "string", "filePaths": ["string"], "isDiff": 1, "isNeedToVerify": 1, "moduleCode": "string", "name": "string", "productID": "string", "signMethod": "string", "srcVersion": "string", "version": "string"}'
```

### POST `/api/v1/things/ota/firmware/info/delete`

**说明**: 删除升级包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

**请求示例**:
```json
{
  "id": 1
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
ur api /api/v1/things/ota/firmware/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/ota/firmware/info/get-list`

**说明**: 升级包列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 |  |

**请求示例**:
```json
{
  "name": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "extra": "string",
        "fileList": [
          {
            "fileMd5": "string",
            "filePath": "string",
            "name": "string",
            "signature": "string"
          }
        ],
        "id": "string",
        "isDiff": 1,
        "isNeedToVerify": 1,
        "moduleCode": "string",
        "name": "示例名称",
        "productID": "string",
        "productName": "示例名称",
        "signMethod": "string",
        "status": 1,
        "version": "string"
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
ur api /api/v1/things/ota/firmware/info/get-list \
  --body '{"name": "string", "page": {"page": 1, "pageSize": 1}, "productID": "string"}'
```

### POST `/api/v1/things/ota/firmware/info/get-one`

**说明**: 查询升级包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

**请求示例**:
```json
{
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "extra": "string",
    "fileList": [
      {
        "fileMd5": "string",
        "filePath": "string",
        "name": "string",
        "signature": "string"
      }
    ],
    "id": "string",
    "isDiff": 1,
    "isNeedToVerify": 1,
    "moduleCode": "string",
    "name": "示例名称",
    "productID": "string",
    "productName": "示例名称",
    "signMethod": "string",
    "status": 1,
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/ota/firmware/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/things/ota/firmware/info/update`

**说明**: 修改升级包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `extra` | string | 否 |  |
| `id` | string | 是 |  |
| `name` | string | 否 |  |

**请求示例**:
```json
{
  "desc": "string",
  "extra": "string",
  "id": "string",
  "name": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/ota/firmware/info/update \
  --body '{"desc": "string", "extra": "string", "id": "string", "name": "string"}'
```
