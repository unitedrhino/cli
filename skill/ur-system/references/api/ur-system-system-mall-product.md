# ur-system system/mall/product

创建商品 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/mall/product/create` | 创建商品 | platform |
| POST | `/api/v1/system/mall/product/delete` | 删除商品 | platform |
| POST | `/api/v1/system/mall/product/get-list` | 获取商品列表 | admin |
| POST | `/api/v1/system/mall/product/get-one` | 获取商品详情 | admin |
| POST | `/api/v1/system/mall/product/update` | 更新商品 | platform |

## 详细说明

### POST `/api/v1/system/mall/product/create`

**说明**: 创建商品

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `grantSchema` | string | 否 |  |
| `productCode` | string | 是 |  |
| `productName` | string | 是 |  |
| `productType` | string | 是 |  |
| `skus` | array[MallProductSku] | 否 |  |
| `status` | integer | 否 | 格式: int32 |

**请求示例**:
```json
{
  "desc": "string",
  "grantSchema": "string",
  "productCode": "string",
  "productName": "string",
  "productType": "string",
  "skus": [
    {
      "bindMode": "string",
      "configValues": {},
      "durationDays": 1,
      "licenseMode": "string",
      "skuCode": "string",
      "skuName": "string",
      "status": 1
    }
  ],
  "status": 1
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
ur api /api/v1/system/mall/product/create \
  --body '{"desc": "string", "grantSchema": "string", "productCode": "string", "productName": "string", "productType": "string", "skus": [{"bindMode": "string", "configValues": {}, "durationDays": 1, "licenseMode": "string", "skuCode": "string", "skuName": "string", "status": 1}], "status": 1}'
```

### POST `/api/v1/system/mall/product/delete`

**说明**: 删除商品

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/product/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/mall/product/get-list`

**说明**: 获取商品列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productCode` | string | 否 |  |
| `productName` | string | 否 |  |
| `productType` | string | 否 |  |
| `status` | integer | 否 | 格式: int32 |

**请求示例**:
```json
{
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
  },
  "productCode": "string",
  "productName": "string",
  "productType": "string",
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
        "createdTime": "string",
        "desc": "string",
        "grantSchema": "string",
        "id": "string",
        "productCode": "string",
        "productName": "string",
        "productType": "string",
        "skus": [
          {
            "bindMode": "string",
            "configValues": {},
            "durationDays": 1,
            "licenseMode": "string",
            "skuCode": "string",
            "skuName": "string",
            "status": 1
          }
        ],
        "status": 1
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/product/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productCode": "string", "productName": "string", "productType": "string", "status": 1}'
```

### POST `/api/v1/system/mall/product/get-one`

**说明**: 获取商品详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

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
    "createdTime": "string",
    "desc": "string",
    "grantSchema": "string",
    "id": "string",
    "productCode": "string",
    "productName": "string",
    "productType": "string",
    "skus": [
      {
        "bindMode": "string",
        "configValues": {},
        "durationDays": 1,
        "licenseMode": "string",
        "skuCode": "string",
        "skuName": "string",
        "status": 1
      }
    ],
    "status": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/product/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/mall/product/update`

**说明**: 更新商品

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `grantSchema` | string | 否 |  |
| `id` | string | 是 |  |
| `productName` | string | 否 |  |
| `productType` | string | 否 |  |
| `skus` | array[MallProductSku] | 否 |  |
| `status` | integer | 否 | 格式: int32 |

**请求示例**:
```json
{
  "desc": "string",
  "grantSchema": "string",
  "id": "string",
  "productName": "string",
  "productType": "string",
  "skus": [
    {
      "bindMode": "string",
      "configValues": {},
      "durationDays": 1,
      "licenseMode": "string",
      "skuCode": "string",
      "skuName": "string",
      "status": 1
    }
  ],
  "status": 1
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
ur api /api/v1/system/mall/product/update \
  --body '{"desc": "string", "grantSchema": "string", "id": "string", "productName": "string", "productType": "string", "skus": [{"bindMode": "string", "configValues": {}, "durationDays": 1, "licenseMode": "string", "skuCode": "string", "skuName": "string", "status": 1}], "status": 1}'
```
