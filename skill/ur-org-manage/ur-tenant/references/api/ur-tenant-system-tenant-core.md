# ur-tenant system/tenant/core

搜索租户信息 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/core/get-list` | 搜索租户信息 | public |
| POST | `/api/v1/system/tenant/core/get-one` | 获取租户信息 | public |

## 详细说明

### POST `/api/v1/system/tenant/core/get-list`

**说明**: 搜索租户信息

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 | 过滤绑定了指定app的租户 |
| `code` | string | 否 |  编号 |
| `name` | string | 否 |  租户名称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "appID": "string",
  "code": "string",
  "name": "示例名称",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "appIDs": [
          "string"
        ],
        "avatar": "string",
        "code": "string",
        "desc": "string",
        "id": "string",
        "name": "示例名称",
        "resource": {
          "package": {
            "expTime": "2026-01-01T00:00:00Z",
            "packageCode": "string",
            "packageName": "示例名称"
          },
          "quota": {
            "aiTokenBalance": 1,
            "aiTokenMonthly": 1,
            "deviceMsgPerDay": 1,
            "deviceNum": 1,
            "diskSpaceGB": 1,
            "userNum": 1
          },
          "usage": {
            "aiTokenBalanceUsed": 1,
            "aiTokenMonth": "string",
            "aiTokenUsed": 1,
            "deviceCount": 1,
            "userCount": 1
          }
        }
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
ur api /api/v1/system/tenant/core/get-list \
  --body '{"appID": "string", "code": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/tenant/core/get-one`

**说明**: 获取租户信息

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | integer | 否 |  id (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "appIDs": [
      "string"
    ],
    "avatar": "string",
    "code": "string",
    "desc": "string",
    "id": "string",
    "name": "示例名称",
    "resource": {
      "package": {
        "expTime": "2026-01-01T00:00:00Z",
        "packageCode": "string",
        "packageName": "示例名称"
      },
      "quota": {
        "aiTokenBalance": 1,
        "aiTokenMonthly": 1,
        "deviceMsgPerDay": 1,
        "deviceNum": 1,
        "diskSpaceGB": 1,
        "userNum": 1
      },
      "usage": {
        "aiTokenBalanceUsed": 1,
        "aiTokenMonth": "string",
        "aiTokenUsed": 1,
        "deviceCount": 1,
        "userCount": 1
      }
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/core/get-one \
  --body '{"code": "string", "id": 1}'
```
