# ur-system system/mall/license

批量创建授权码 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/mall/license/batch-create` | 批量创建授权码 | platform |
| POST | `/api/v1/system/mall/license/create` | 创建授权码 | platform |
| POST | `/api/v1/system/mall/license/get-list` | 授权码列表 | admin |
| POST | `/api/v1/system/mall/license/get-one` | 授权码详情 | admin |
| POST | `/api/v1/system/mall/license/revoke` | 撤销授权码 | platform |

## 详细说明

### POST `/api/v1/system/mall/license/batch-create`

**说明**: 批量创建授权码

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `count` | integer | 否 | 格式: int32 |
| `desc` | string | 否 |  |
| `issuer` | string | 否 |  |
| `skuCode` | string | 否 |  |
| `sourceID` | string | 是 |  |
| `sourceType` | string | 是 |  |

**请求示例**:
```json
{
  "count": 1,
  "desc": "string",
  "issuer": "string",
  "skuCode": "string",
  "sourceID": "string",
  "sourceType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "count": 1,
    "licenseCodes": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/license/batch-create \
  --body '{"count": 1, "desc": "string", "issuer": "string", "skuCode": "string", "sourceID": "string", "sourceType": "string"}'
```

### POST `/api/v1/system/mall/license/create`

**说明**: 创建授权码

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `issuer` | string | 否 |  |
| `skuCode` | string | 否 |  |
| `sourceID` | string | 是 |  |
| `sourceType` | string | 是 |  |

**请求示例**:
```json
{
  "desc": "string",
  "issuer": "string",
  "skuCode": "string",
  "sourceID": "string",
  "sourceType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "licenseCode": "string",
    "withSource": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/license/create \
  --body '{"desc": "string", "issuer": "string", "skuCode": "string", "sourceID": "string", "sourceType": "string"}'
```

### POST `/api/v1/system/mall/license/get-list`

**说明**: 授权码列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `boundTargetID` | string | 否 |  |
| `licenseCode` | string | 否 |  |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `skuCode` | string | 否 |  |
| `sourceID` | string | 否 |  |
| `sourceType` | string | 否 |  |
| `status` | string | 否 |  |
| `withSource` | boolean | 否 |  是否返回来源对象详情（商品或套餐） (格式: boolean) |

**请求示例**:
```json
{
  "boundTargetID": "string",
  "licenseCode": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "skuCode": "string",
  "sourceID": "string",
  "sourceType": "string",
  "status": "string",
  "withSource": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "boundMeta": {},
        "boundTargetID": "string",
        "boundTime": "string",
        "createdTime": "string",
        "desc": "string",
        "grantResult": "string",
        "grantSnapshot": "string",
        "id": "string",
        "issuer": "string",
        "licenseCode": "string",
        "skuCode": "string",
        "sourceID": "string",
        "sourcePackage": {
          "category": "string",
          "createdTime": "string",
          "desc": "string",
          "id": "string",
          "items": [
            {
              "productID": "...",
              "skuCode": "...",
              "sort": "...",
              "status": "..."
            }
          ],
          "packageCode": "string",
          "packageName": "string",
          "status": 1,
          "verifyMethod": "string"
        },
        "sourceProduct": {
          "createdTime": "string",
          "desc": "string",
          "grantSchema": "string",
          "id": "string",
          "productCode": "string",
          "productName": "string",
          "productType": "string",
          "skus": [
            {
              "bindMode": "...",
              "configValues": "...",
              "durationDays": "...",
              "licenseMode": "...",
              "skuCode": "...",
              "skuName": "...",
              "status": "..."
            }
          ],
          "status": 1
        },
        "sourceType": "string",
        "status": "string",
        "validFrom": "string",
        "validUntil": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/license/get-list \
  --body '{"boundTargetID": "string", "licenseCode": "string", "page": {"page": 1, "pageSize": 1}, "skuCode": "string", "sourceID": "string", "sourceType": "string", "status": "string", "withSource": true}'
```

### POST `/api/v1/system/mall/license/get-one`

**说明**: 授权码详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `licenseCode` | string | 是 |  授权码 |
| `withSource` | boolean | 否 |  是否返回来源对象详情（商品或套餐） (格式: boolean) |

**请求示例**:
```json
{
  "licenseCode": "string",
  "withSource": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "boundMeta": {},
    "boundTargetID": "string",
    "boundTime": "string",
    "createdTime": "string",
    "desc": "string",
    "grantResult": "string",
    "grantSnapshot": "string",
    "id": "string",
    "issuer": "string",
    "licenseCode": "string",
    "skuCode": "string",
    "sourceID": "string",
    "sourcePackage": {
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
    "sourceProduct": {
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
    "sourceType": "string",
    "status": "string",
    "validFrom": "string",
    "validUntil": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/license/get-one \
  --body '{"licenseCode": "string", "withSource": true}'
```

### POST `/api/v1/system/mall/license/revoke`

**说明**: 撤销授权码

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `licenseCode` | string | 是 |  授权码 |
| `withSource` | boolean | 否 |  是否返回来源对象详情（商品或套餐） (格式: boolean) |

**请求示例**:
```json
{
  "licenseCode": "string",
  "withSource": true
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
ur api /api/v1/system/mall/license/revoke \
  --body '{"licenseCode": "string", "withSource": true}'
```
