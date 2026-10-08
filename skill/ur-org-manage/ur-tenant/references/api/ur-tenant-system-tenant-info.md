# ur-tenant system/tenant/info

添加租户 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/info/create` | 添加租户 | all |
| POST | `/api/v1/system/tenant/info/delete` | 删除租户 | platform |
| POST | `/api/v1/system/tenant/info/get-list` | 获取租户列表 | platform |
| POST | `/api/v1/system/tenant/info/get-one` | 获取租户详情 | admin |
| POST | `/api/v1/system/tenant/info/transfer` | 转让租户（仅租户所有者） | admin |
| POST | `/api/v1/system/tenant/info/update` | 更新租户 | admin |

## 详细说明

### POST `/api/v1/system/tenant/info/create`

**说明**: 添加租户

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  头像,更新时传filePath |
| `code` | string | 否 |  租户编码,创建的时候自动生成 |
| `config` | object | 否 |  |
| `config.id` | string | 否 |  |
| `config.registerRoleID` | string | 否 | 注册分配的角色id |
| `config.tenantCode` | string | 否 | 租户编码 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  应用描述 |
| `id` | string | 否 |  id编号 |
| `licenseCode` | string | 否 |  授权码（创建时可选，写入后应用套餐配额） |
| `name` | string | 否 |  租户名称 |
| `ownerUser` | object | 否 |  |
| `ownerUser.avatar` | string | 否 |  用户头像 |
| `ownerUser.createdTime` | string | 否 |  创建时间 |
| `ownerUser.email` | string | 否 |  邮箱 |
| `ownerUser.lastIP` | string | 否 |  最后登录ip |
| `ownerUser.nickName` | string | 否 |  用户的昵称 |
| `ownerUser.phone` | string | 否 |  手机号 |
| `ownerUser.regIP` | string | 否 |  注册ip |
| `ownerUser.userID` | string | 否 |  用户id |
| `ownerUser.userName` | string | 否 |  用户名(唯一) |
| `ownerUserID` | string | 否 |  超级管理员id |
| `resource` | object | 否 |  |
| `resource.package` | object | 否 |  |
| `resource.package.expTime` | string | 否 |  企业套餐到期时间（毫秒时间戳） |
| `resource.package.packageCode` | string | 否 |  当前套餐编码 |
| `resource.package.packageName` | string | 否 |  当前套餐名称 |
| `resource.quota` | object | 否 |  |
| `resource.quota.aiTokenBalance` | integer | 否 |  AI Token 充值余额 (格式: int64) |
| `resource.quota.aiTokenMonthly` | integer | 否 |  AI Token 月配额 (格式: int64) |
| `resource.quota.deviceMsgPerDay` | integer | 否 |  租户每天设备消息总量上限 (格式: int64) |
| `resource.quota.deviceNum` | integer | 否 |  租户下的设备数量限制，0 为不限制 (格式: int64) |
| `resource.quota.diskSpaceGB` | integer | 否 |  磁盘空间配额（GB） (格式: int64) |
| `resource.quota.userNum` | integer | 否 |  租户下的用户数量限制，0 为不限制 (格式: int64) |
| `resource.usage` | object | 否 |  |
| `resource.usage.aiTokenBalanceUsed` | integer | 否 |  AI Token 余额已用 (格式: int64) |
| `resource.usage.aiTokenMonth` | string | 否 |  AI Token 用量月份 |
| `resource.usage.aiTokenUsed` | integer | 否 |  AI Token 月已用 (格式: int64) |
| `resource.usage.deviceCount` | integer | 否 | 租户下设备统计,只读 (格式: int64) |
| `resource.usage.userCount` | integer | 否 | 租户下用户统计,只读 (格式: int64) |
| `status` | integer | 否 |  租戶状态: 1启用 2禁用 (格式: int64) |

**请求示例**:
```json
{
  "avatar": "string",
  "code": "string",
  "config": {
    "id": "string",
    "registerRoleID": "string",
    "tenantCode": "string"
  },
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "licenseCode": "string",
  "name": "示例名称",
  "ownerUser": {
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
  "ownerUserID": "string",
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
  },
  "status": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/info/create \
  --body '{"avatar": "string", "code": "string", "config": {"id": "string", "registerRoleID": "string", "tenantCode": "string"}, "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "licenseCode": "string", "name": "示例名称", "ownerUser": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "ownerUserID": "string", "resource": {"package": {"expTime": "2026-01-01T00:00:00Z", "packageCode": "string", "packageName": "示例名称"}, "quota": {"aiTokenBalance": 1, "aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}, "usage": {"aiTokenBalanceUsed": 1, "aiTokenMonth": "string", "aiTokenUsed": 1, "deviceCount": 1, "userCount": 1}}, "status": 1}'
```

### POST `/api/v1/system/tenant/info/delete`

**说明**: 删除租户

**权限**: platform

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/info/delete \
  --body '{"code": "string", "id": 1}'
```

### POST `/api/v1/system/tenant/info/get-list`

**说明**: 获取租户列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  租户编号 |
| `isGetAll` | boolean | 否 | 是否返回所有,只有平台管理员有权限 (格式: boolean) |
| `name` | string | 否 |  租户名称 |
| `ownerUserID` | string | 否 |  超级管理员id,只有平台管理员有权限 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 租户状态 (格式: int64) |
| `tenantCodes` | array[string] | 否 |  指定租户编码列表，非空时按编码精确返回 |
| `withOwnerUser` | boolean | 否 | 同时获取管理员核心信息 (格式: boolean) |

**请求示例**:
```json
{
  "code": "string",
  "isGetAll": true,
  "name": "示例名称",
  "ownerUserID": "string",
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
  "status": 1,
  "tenantCodes": [
    "string"
  ],
  "withOwnerUser": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "avatar": "string",
        "code": "string",
        "config": {
          "id": "string",
          "registerRoleID": "string",
          "tenantCode": "string"
        },
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "id": "string",
        "licenseCode": "string",
        "name": "示例名称",
        "ownerUser": {
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
        "ownerUserID": "string",
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
        },
        "status": 1
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
ur api /api/v1/system/tenant/info/get-list \
  --body '{"code": "string", "isGetAll": true, "name": "示例名称", "ownerUserID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "tenantCodes": ["string"], "withOwnerUser": true}'
```

### POST `/api/v1/system/tenant/info/get-one`

**说明**: 获取租户详情

**权限**: admin

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
    "avatar": "string",
    "code": "string",
    "config": {
      "id": "string",
      "registerRoleID": "string",
      "tenantCode": "string"
    },
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "licenseCode": "string",
    "name": "示例名称",
    "ownerUser": {
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
    "ownerUserID": "string",
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
    },
    "status": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/info/get-one \
  --body '{"code": "string", "id": 1}'
```

### POST `/api/v1/system/tenant/info/transfer`

**说明**: 转让租户（仅租户所有者）

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `newOwnerUserID` | string | 是 |  新的租户所有者用户ID |

**请求示例**:
```json
{
  "newOwnerUserID": "string"
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
ur api /api/v1/system/tenant/info/transfer \
  --body '{"newOwnerUserID": "string"}'
```

### POST `/api/v1/system/tenant/info/update`

**说明**: 更新租户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  头像,更新时传filePath |
| `code` | string | 否 |  租户编码,创建的时候自动生成 |
| `config` | object | 否 |  |
| `config.id` | string | 否 |  |
| `config.registerRoleID` | string | 否 | 注册分配的角色id |
| `config.tenantCode` | string | 否 | 租户编码 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  应用描述 |
| `id` | string | 否 |  id编号 |
| `licenseCode` | string | 否 |  授权码（创建时可选，写入后应用套餐配额） |
| `name` | string | 否 |  租户名称 |
| `ownerUser` | object | 否 |  |
| `ownerUser.avatar` | string | 否 |  用户头像 |
| `ownerUser.createdTime` | string | 否 |  创建时间 |
| `ownerUser.email` | string | 否 |  邮箱 |
| `ownerUser.lastIP` | string | 否 |  最后登录ip |
| `ownerUser.nickName` | string | 否 |  用户的昵称 |
| `ownerUser.phone` | string | 否 |  手机号 |
| `ownerUser.regIP` | string | 否 |  注册ip |
| `ownerUser.userID` | string | 否 |  用户id |
| `ownerUser.userName` | string | 否 |  用户名(唯一) |
| `ownerUserID` | string | 否 |  超级管理员id |
| `resource` | object | 否 |  |
| `resource.package` | object | 否 |  |
| `resource.package.expTime` | string | 否 |  企业套餐到期时间（毫秒时间戳） |
| `resource.package.packageCode` | string | 否 |  当前套餐编码 |
| `resource.package.packageName` | string | 否 |  当前套餐名称 |
| `resource.quota` | object | 否 |  |
| `resource.quota.aiTokenBalance` | integer | 否 |  AI Token 充值余额 (格式: int64) |
| `resource.quota.aiTokenMonthly` | integer | 否 |  AI Token 月配额 (格式: int64) |
| `resource.quota.deviceMsgPerDay` | integer | 否 |  租户每天设备消息总量上限 (格式: int64) |
| `resource.quota.deviceNum` | integer | 否 |  租户下的设备数量限制，0 为不限制 (格式: int64) |
| `resource.quota.diskSpaceGB` | integer | 否 |  磁盘空间配额（GB） (格式: int64) |
| `resource.quota.userNum` | integer | 否 |  租户下的用户数量限制，0 为不限制 (格式: int64) |
| `resource.usage` | object | 否 |  |
| `resource.usage.aiTokenBalanceUsed` | integer | 否 |  AI Token 余额已用 (格式: int64) |
| `resource.usage.aiTokenMonth` | string | 否 |  AI Token 用量月份 |
| `resource.usage.aiTokenUsed` | integer | 否 |  AI Token 月已用 (格式: int64) |
| `resource.usage.deviceCount` | integer | 否 | 租户下设备统计,只读 (格式: int64) |
| `resource.usage.userCount` | integer | 否 | 租户下用户统计,只读 (格式: int64) |
| `status` | integer | 否 |  租戶状态: 1启用 2禁用 (格式: int64) |

**请求示例**:
```json
{
  "avatar": "string",
  "code": "string",
  "config": {
    "id": "string",
    "registerRoleID": "string",
    "tenantCode": "string"
  },
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "licenseCode": "string",
  "name": "示例名称",
  "ownerUser": {
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
  "ownerUserID": "string",
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
  },
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
ur api /api/v1/system/tenant/info/update \
  --body '{"avatar": "string", "code": "string", "config": {"id": "string", "registerRoleID": "string", "tenantCode": "string"}, "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "licenseCode": "string", "name": "示例名称", "ownerUser": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "ownerUserID": "string", "resource": {"package": {"expTime": "2026-01-01T00:00:00Z", "packageCode": "string", "packageName": "示例名称"}, "quota": {"aiTokenBalance": 1, "aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}, "usage": {"aiTokenBalanceUsed": 1, "aiTokenMonth": "string", "aiTokenUsed": 1, "deviceCount": 1, "userCount": 1}}, "status": 1}'
```
