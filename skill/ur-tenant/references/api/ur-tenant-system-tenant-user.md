# ur-tenant system/tenant/user

批量添加用户加入租户 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/user/batch-create` | 批量添加用户加入租户 | admin |
| POST | `/api/v1/system/tenant/user/delete` | 删除租户用户 | admin |
| POST | `/api/v1/system/tenant/user/get-list` | 获取租户用户列表 | admin |
| POST | `/api/v1/system/tenant/user/get-one` | 获取租户用户详情,会同时返回所拥有的角色列表 | admin |
| POST | `/api/v1/system/tenant/user/invite` | 邀请用户加入租户 | admin |
| POST | `/api/v1/system/tenant/user/invite-code/gen` | 生成租户用户邀请码 | admin |
| POST | `/api/v1/system/tenant/user/invite-code/get-one` | 获取当前有效的租户用户邀请码 | admin |
| POST | `/api/v1/system/tenant/user/invite-pending/delete` | 删除待处理邀请 | admin |
| POST | `/api/v1/system/tenant/user/invite-pending/get-list` | 获取待处理邀请列表 | admin |
| POST | `/api/v1/system/tenant/user/invite-send` | 发送邀请（支持邮件和手机号） | admin |
| POST | `/api/v1/system/tenant/user/update` | 更新租户用户 | admin |

## 详细说明

### POST `/api/v1/system/tenant/user/batch-create`

**说明**: 批量添加用户加入租户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `users` | array[UserTenantCore] | 是 |  用户列表 |

**请求示例**:
```json
{
  "users": [
    {
      "deptIDs": [
        "string"
      ],
      "pubTags": {},
      "roleIDs": [
        "string"
      ],
      "status": 1,
      "tags": {},
      "userID": "string"
    }
  ]
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
ur api /api/v1/system/tenant/user/batch-create \
  --body '{"users": [{"deptIDs": ["string"], "pubTags": {}, "roleIDs": ["string"], "status": 1, "tags": {}, "userID": "string"}]}'
```

### POST `/api/v1/system/tenant/user/delete`

**说明**: 删除租户用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `tenantCode` | string | 否 |  租户代码,(平台管理员查看其他租户下的用户时使用,普通租户只能查看自己的用户) |
| `userID` | string | 是 |  用户ID |

**请求示例**:
```json
{
  "tenantCode": "string",
  "userID": "string"
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
ur api /api/v1/system/tenant/user/delete \
  --body '{"tenantCode": "string", "userID": "string"}'
```

### POST `/api/v1/system/tenant/user/get-list`

**说明**: 获取租户用户列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 |  账号 |
| `deptID` | string | 否 |  部门ID |
| `email` | string | 否 |  邮箱 |
| `hasAccessAreas` | array[integer] | 否 |  拥有访问权限的区域 |
| `nickName` | string | 否 |  昵称 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `phone` | string | 否 |  手机号 |
| `roleCode` | string | 否 |  角色编码 |
| `status` | integer | 否 |  租户状态（1:启用，2:禁用） (格式: int64) |
| `tenantCode` | string | 否 |  租户代码,(平台管理员查看其他租户下的用户时使用,普通租户只能查看自己的用户) |
| `userIDs` | array[string] | 否 |  用户ID列表 |
| `userName` | string | 否 |  用户名 |
| `withRole` | boolean | 否 | 同时返回角色信息 (格式: boolean) |

**请求示例**:
```json
{
  "account": "string",
  "deptID": "string",
  "email": "string",
  "hasAccessAreas": [
    1
  ],
  "nickName": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "phone": "string",
  "roleCode": "string",
  "status": 1,
  "tenantCode": "string",
  "userIDs": [
    "string"
  ],
  "userName": "string",
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
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/user/get-list \
  --body '{"account": "string", "deptID": "string", "email": "string", "hasAccessAreas": [1], "nickName": "string", "page": {"page": 1, "pageSize": 1}, "phone": "string", "roleCode": "string", "status": 1, "tenantCode": "string", "userIDs": ["string"], "userName": "string", "withRole": true}'
```

### POST `/api/v1/system/tenant/user/get-one`

**说明**: 获取租户用户详情,会同时返回所拥有的角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `tenantCode` | string | 否 |  租户代码,(平台管理员查看其他租户下的用户时使用,普通租户只能查看自己的用户) |
| `userID` | string | 是 |  用户ID |

**请求示例**:
```json
{
  "tenantCode": "string",
  "userID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
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
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/user/get-one \
  --body '{"tenantCode": "string", "userID": "string"}'
```

### POST `/api/v1/system/tenant/user/invite`

**说明**: 邀请用户加入租户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accounts` | array[string] | 否 |  账号列表(userIDs没有取这个) |
| `roleCode` | string | 否 |  角色code,如果有重复的,则只授权第一个角色 |
| `roleID` | string | 否 |  角色ID,默认取这个 |
| `userIDs` | array[string] | 否 |  用户ID列表 |

**请求示例**:
```json
{
  "accounts": [
    "string"
  ],
  "roleCode": "string",
  "roleID": "string",
  "userIDs": [
    "string"
  ]
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
ur api /api/v1/system/tenant/user/invite \
  --body '{"accounts": ["string"], "roleCode": "string", "roleID": "string", "userIDs": ["string"]}'
```

### POST `/api/v1/system/tenant/user/invite-code/gen`

**说明**: 生成租户用户邀请码

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `roleID` | string | 是 |  分配的角色ID（必填） |
| `tenantCode` | string | 否 |  租户编码（可选，不填取header） |

**请求示例**:
```json
{
  "roleID": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "roleID": "string",
    "validUntilTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/user/invite-code/gen \
  --body '{"roleID": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/user/invite-code/get-one`

**说明**: 获取当前有效的租户用户邀请码

**权限**: admin

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "roleID": "string",
    "validUntilTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/user/invite-code/get-one \
  --body '{}'
```

### POST `/api/v1/system/tenant/user/invite-pending/delete`

**说明**: 删除待处理邀请

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `address` | string | 是 |  邮箱地址或手机号 |
| `inviteType` | string | 是 |  邀请类型：email/phone |
| `tenantCode` | string | 否 |  租户编码（可选，不填取header） |

**请求示例**:
```json
{
  "address": "string",
  "inviteType": "string",
  "tenantCode": "string"
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
ur api /api/v1/system/tenant/user/invite-pending/delete \
  --body '{"address": "string", "inviteType": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/user/invite-pending/get-list`

**说明**: 获取待处理邀请列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `inviteType` | string | 否 |  邀请类型筛选（可选：email/phone，不填返回全部） |
| `tenantCode` | string | 否 |  租户编码（可选，不填取header） |

**请求示例**:
```json
{
  "inviteType": "string",
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
        "address": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "inviteType": "string",
        "roleID": "string",
        "validUntilTime": "2026-01-01T00:00:00Z"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/user/invite-pending/get-list \
  --body '{"inviteType": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/user/invite-send`

**说明**: 发送邀请（支持邮件和手机号）

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `addresses` | array[string] | 是 |  邀请地址列表（邮箱或手机号，自动识别类型） |
| `roleID` | string | 是 |  分配的角色ID（必填） |
| `tenantCode` | string | 否 |  租户编码（可选，不填取header） |
| `validUntilTime` | string | 否 |  有效期到（毫秒时间戳，不填使用默认值） |

**请求示例**:
```json
{
  "addresses": [
    "string"
  ],
  "roleID": "string",
  "tenantCode": "string",
  "validUntilTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/system/tenant/user/invite-send \
  --body '{"addresses": ["string"], "roleID": "string", "tenantCode": "string", "validUntilTime": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/system/tenant/user/update`

**说明**: 更新租户用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 |  创建时间,只读 |
| `deptIDs` | array[string] | 否 | 部门ID列表,默认展示第一个 |
| `deviceCount` | integer | 否 | 用户所拥有的设备数量统计,只读 (格式: int64) |
| `isTenantOwner` | integer | 否 | 是否是租户管理员,只读 (格式: int64) |
| `pubTags` | object | 否 |  公共的标签,用户自己可以修改 |
| `roles` | array[RoleInfo] | 否 |  角色列表 |
| `status` | integer | 否 | 用户状态(管理员才可修改): 1启用 2禁用 (格式: int64) |
| `tags` | object | 否 |  管理员才可修改的标签 |
| `tenant` | object | 否 |  |
| `tenant.appIDs` | array[string] | 否 |  租户拥有的应用ID列表 |
| `tenant.avatar` | string | 否 |  头像,更新时传filePath |
| `tenant.code` | string | 否 |  租户编码,创建的时候自动生成 |
| `tenant.desc` | string | 否 |  应用描述 |
| `tenant.id` | string | 否 |  id编号 |
| `tenant.name` | string | 否 |  租户名称 |
| `tenant.resource` | object | 否 |  |
| `tenant.resource.package` | object | 否 |  |
| `tenant.resource.quota` | object | 否 |  |
| `tenant.resource.usage` | object | 否 |  |
| `tenantCode` | string | 否 |  租户编码,只读 |
| `user` | object | 否 |  |
| `user.avatar` | string | 否 |  用户头像 |
| `user.createdTime` | string | 否 |  创建时间 |
| `user.email` | string | 否 |  邮箱 |
| `user.lastIP` | string | 否 |  最后登录ip |
| `user.nickName` | string | 否 |  用户的昵称 |
| `user.phone` | string | 否 |  手机号 |
| `user.regIP` | string | 否 |  注册ip |
| `user.userID` | string | 否 |  用户id |
| `user.userName` | string | 否 |  用户名(唯一) |
| `userID` | string | 否 |  用户id,只读 |

**请求示例**:
```json
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
ur api /api/v1/system/tenant/user/update \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "deptIDs": ["string"], "deviceCount": 1, "isTenantOwner": 1, "pubTags": {}, "roles": [{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "isFullAuth": 1, "name": "示例名称", "status": 1}], "status": 1, "tags": {}, "tenant": {"appIDs": ["string"], "avatar": "string", "code": "string", "desc": "string", "id": "string", "name": "示例名称", "resource": {"package": {"expTime": "2026-01-01T00:00:00Z", "packageCode": "string", "packageName": "示例名称"}, "quota": {"aiTokenBalance": 1, "aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}, "usage": {"aiTokenBalanceUsed": 1, "aiTokenMonth": "string", "aiTokenUsed": 1, "deviceCount": 1, "userCount": 1}}}, "tenantCode": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "userID": "string"}'
```
