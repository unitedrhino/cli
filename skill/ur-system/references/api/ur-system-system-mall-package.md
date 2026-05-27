# ur-system system/mall/package

创建套餐 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/mall/package/create` | 创建套餐 | platform |
| POST | `/api/v1/system/mall/package/delete` | 删除套餐 | platform |
| POST | `/api/v1/system/mall/package/get-list` | 获取套餐列表 | admin |
| POST | `/api/v1/system/mall/package/get-one` | 获取套餐详情 | admin |
| POST | `/api/v1/system/mall/package/update` | 更新套餐 | platform |

## 详细说明

### POST `/api/v1/system/mall/package/create`

**说明**: 创建套餐

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `category` | string | 否 |  |
| `desc` | string | 否 |  |
| `items` | array[MallPackageItem] | 否 |  |
| `packageCode` | string | 是 |  |
| `packageName` | string | 是 |  |
| `status` | integer | 否 | 格式: int32 |
| `verifyMethod` | string | 否 |  |

**请求示例**:
```json
{
  "category": "string",
  "desc": "string",
  "items": [
    {
      "productID": "string",
      "skuCode": "string",
      "sort": 1,
      "status": 1
    }
  ],
  "packageCode": "string",
  "packageName": "string",
  "status": 1,
  "verifyMethod": "string"
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
ur api /api/v1/system/mall/package/create \
  --body '{"category": "string", "desc": "string", "items": [{"productID": "string", "skuCode": "string", "sort": 1, "status": 1}], "packageCode": "string", "packageName": "string", "status": 1, "verifyMethod": "string"}'
```

### POST `/api/v1/system/mall/package/delete`

**说明**: 删除套餐

**权限**: platform

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
ur api /api/v1/system/mall/package/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/mall/package/get-list`

**说明**: 获取套餐列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `category` | string | 否 |  |
| `packageCode` | string | 否 |  |
| `packageName` | string | 否 |  |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int32 |

**请求示例**:
```json
{
  "category": "string",
  "packageCode": "string",
  "packageName": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "category": "string",
        "createdTime": "string",
        "desc": "string",
        "id": "string",
        "items": [
          {
            "productID": "string",
            "skuCode": "string",
            "sort": 1,
            "status": 1
          }
        ],
        "packageCode": "string",
        "packageName": "string",
        "status": 1,
        "verifyMethod": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/package/get-list \
  --body '{"category": "string", "packageCode": "string", "packageName": "string", "page": {"page": 1, "pageSize": 1}, "status": 1}'
```

### POST `/api/v1/system/mall/package/get-one`

**说明**: 获取套餐详情

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
    "category": "string",
    "createdTime": "string",
    "desc": "string",
    "id": "string",
    "items": [
      {
        "productID": "string",
        "skuCode": "string",
        "sort": 1,
        "status": 1
      }
    ],
    "packageCode": "string",
    "packageName": "string",
    "status": 1,
    "verifyMethod": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/package/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/mall/package/update`

**说明**: 更新套餐

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `category` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 是 |  |
| `items` | array[MallPackageItem] | 否 |  |
| `packageName` | string | 否 |  |
| `status` | integer | 否 | 格式: int32 |
| `verifyMethod` | string | 否 |  |

**请求示例**:
```json
{
  "category": "string",
  "desc": "string",
  "id": "string",
  "items": [
    {
      "productID": "string",
      "skuCode": "string",
      "sort": 1,
      "status": 1
    }
  ],
  "packageName": "string",
  "status": 1,
  "verifyMethod": "string"
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
ur api /api/v1/system/mall/package/update \
  --body '{"category": "string", "desc": "string", "id": "string", "items": [{"productID": "string", "skuCode": "string", "sort": 1, "status": 1}], "packageName": "string", "status": 1, "verifyMethod": "string"}'
```
