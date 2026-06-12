# ur-user system/user/tenant

用户所处的企业列表

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/tenant/get-list` | 用户所处的企业列表 | all |

## 详细说明

### POST `/api/v1/system/user/tenant/get-list`

**说明**: 用户所处的企业列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  用户状态（1:启用，2:禁用） (格式: int64) |
| `userID` | string | 是 |  |
| `withRole` | boolean | 否 |  同时返回角色信息 (格式: boolean) |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": 1,
  "userID": "string",
  "withRole": true
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
        "deptIDs": [
          "string"
        ],
        "deviceCount": 1,
        "isTenantOwner": 1,
        "pubTags": {},
        "roles": [
          {
            "code": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "id": "string",
            "isFullAuth": 1,
            "name": "示例名称",
            "status": 1
          }
        ],
        "status": 1,
        "tags": {},
        "tenant": {
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
              "expTime": "...",
              "packageCode": "...",
              "packageName": "..."
            },
            "quota": {
              "aiTokenBalance": "...",
              "aiTokenMonthly": "...",
              "deviceMsgPerDay": "...",
              "deviceNum": "...",
              "diskSpaceGB": "...",
              "userNum": "..."
            },
            "usage": {
              "aiTokenBalanceUsed": "...",
              "aiTokenMonth": "...",
              "aiTokenUsed": "...",
              "deviceCount": "...",
              "userCount": "..."
            }
          }
        },
        "tenantCode": "string",
        "user": {
          "avatar": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "email": "string",
          "lastIP": "string",
          "nickName": "string",
          "phone": "string",
          "regIP": "string",
          "userID": "string",
          "userName": "string"
        },
        "userID": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/tenant/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}, "status": 1, "userID": "string", "withRole": true}'
```
