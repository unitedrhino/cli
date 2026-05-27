# ur-tenant system/tenant/renewal

直接授权 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/renewal/direct-grant` | 直接授权 | platform |
| POST | `/api/v1/system/tenant/renewal/get-list` | 获取续期列表 | admin |
| POST | `/api/v1/system/tenant/renewal/renew` | 续费 | admin |

## 详细说明

### POST `/api/v1/system/tenant/renewal/direct-grant`

**说明**: 直接授权

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `confirmReplace` | boolean | 否 |  确认替换不同套餐 (格式: boolean) |
| `count` | integer | 否 |  数量倍数，默认1 (格式: int64) |
| `skuCode` | string | 否 |  SKU编码，sourceType=product时必填 |
| `sourceID` | string | 是 |  来源ID |
| `sourceType` | string | 是 |  来源类型：product 商品 package 套餐 |
| `tenantCode` | string | 是 |  租户编码 |

**请求示例**:
```json
{
  "confirmReplace": true,
  "count": 1,
  "skuCode": "string",
  "sourceID": "string",
  "sourceType": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "confirmMessage": "string",
    "granted": true,
    "needsConfirm": true,
    "newExpireTime": "2026-01-01T00:00:00Z",
    "newQuota": {},
    "oldExpireTime": "2026-01-01T00:00:00Z",
    "packageCode": "string",
    "packageName": "示例名称",
    "targetID": "string",
    "targetType": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/renewal/direct-grant \
  --body '{"confirmReplace": true, "count": 1, "skuCode": "string", "sourceID": "string", "sourceType": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/renewal/get-list`

**说明**: 获取续期列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `expireWithinDays` | integer | 否 |  到期天数筛选，传 30 表示仅返回 30 天内到期数据 (格式: int64) |
| `keyword` | string | 否 |  关键字，按企业名、应用名或租户编码模糊匹配 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  到期状态：all 全部 normal 正常 expiring 即将到期 expired 已过期 |
| `targetType` | string | 否 |  续期对象类型：all 全部 tenant 企业 app 应用 |
| `tenantCode` | string | 否 |  指定租户编码，仅平台管理员可传 |

**请求示例**:
```json
{
  "expireWithinDays": 1,
  "keyword": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": "string",
  "targetType": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "appCode": "string",
        "appID": "string",
        "categories": [
          "string"
        ],
        "expireTime": "2026-01-01T00:00:00Z",
        "grantItems": [
          {
            "grantKey": "string",
            "grantType": "string",
            "grantValue": "string"
          }
        ],
        "packageCode": "string",
        "packageName": "示例名称",
        "quota": {},
        "remainDays": 1,
        "renewable": true,
        "status": "string",
        "targetID": "string",
        "targetName": "示例名称",
        "targetType": "string",
        "tenantCode": "string",
        "trialing": true,
        "usage": {}
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
ur api /api/v1/system/tenant/renewal/get-list \
  --body '{"expireWithinDays": 1, "keyword": "string", "page": {"page": 1, "pageSize": 1}, "status": "string", "targetType": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/renewal/renew`

**说明**: 续费

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `confirmReplace` | boolean | 否 |  确认替换不同套餐（plan类型不同商品时需要） (格式: boolean) |
| `licenseCode` | string | 是 |  授权码 |
| `targetID` | string | 是 |  续期对象标识，企业为 tenantCode，应用为 tenantAppID |
| `targetType` | string | 是 |  续期对象类型：tenant 企业 app 应用 |
| `tenantCode` | string | 否 |  所属租户编码，平台管理员可显式指定 |

**请求示例**:
```json
{
  "confirmReplace": true,
  "licenseCode": "string",
  "targetID": "string",
  "targetType": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "confirmMessage": "string",
    "needsConfirm": true,
    "newExpireTime": "2026-01-01T00:00:00Z",
    "newQuota": {},
    "oldExpireTime": "2026-01-01T00:00:00Z",
    "packageCode": "string",
    "packageName": "示例名称",
    "renewed": true,
    "targetID": "string",
    "targetType": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/renewal/renew \
  --body '{"confirmReplace": true, "licenseCode": "string", "targetID": "string", "targetType": "string", "tenantCode": "string"}'
```
