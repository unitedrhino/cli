# ur-user system/user/self/tenant

退出当前企业 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/self/tenant/delete` | 退出当前企业 | all |
| POST | `/api/v1/system/user/self/tenant/get-list` | 获取用户所处的企业列表 | all |
| POST | `/api/v1/system/user/self/tenant/get-one` | 获取当前用户在当前企业的详情 | all |
| POST | `/api/v1/system/user/self/tenant/join` | 用户加入企业（通过邀请码、邮件或手机邀请） | all |
| POST | `/api/v1/system/user/self/tenant/update` | 更新当前用户在当前企业的信息 | all |

## 详细说明

### POST `/api/v1/system/user/self/tenant/delete`

**说明**: 退出当前企业

**权限**: all

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
ur api /api/v1/system/user/self/tenant/delete \
  --body '{}'
```

### POST `/api/v1/system/user/self/tenant/get-list`

**说明**: 获取用户所处的企业列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `withRole` | boolean | 否 |  同时返回角色信息 (格式: boolean) |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
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
ur api /api/v1/system/user/self/tenant/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}, "withRole": true}'
```

### POST `/api/v1/system/user/self/tenant/get-one`

**说明**: 获取当前用户在当前企业的详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `withRole` | boolean | 否 | 同时返回角色信息 (格式: boolean) |
| `withUser` | boolean | 否 | 同时返回用户信息 (格式: boolean) |

**请求示例**:
```json
{
  "withRole": true,
  "withUser": true
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
ur api /api/v1/system/user/self/tenant/get-one \
  --body '{"withRole": true, "withUser": true}'
```

### POST `/api/v1/system/user/self/tenant/join`

**说明**: 用户加入企业（通过邀请码、邮件或手机邀请）

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  邀请码（使用邀请码加入时填写） |
| `method` | string | 否 |  加入方法：code（邀请码）、email（邮件邀请）、phone（手机邀请） |
| `tenantCode` | string | 否 |  企业编码（可选，某些场景可能需要明确指定） |

**请求示例**:
```json
{
  "code": "string",
  "method": "string",
  "tenantCode": "string"
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
ur api /api/v1/system/user/self/tenant/join \
  --body '{"code": "string", "method": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/user/self/tenant/update`

**说明**: 更新当前用户在当前企业的信息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 |  创建时间,只读 |
| `deptIDs` | array[string] | 否 | 部门ID列表,默认展示第一个 |
| `deviceCount` | integer | 否 | 用户所拥有的设备数量统计,只读 (格式: int64) |
| `isTenantOwner` | integer | 否 | 是否是企业管理员,只读 (格式: int64) |
| `pubTags` | object | 否 |  公共的标签,用户自己可以修改 |
| `roles` | array[RoleInfo] | 否 |  角色列表 |
| `status` | integer | 否 | 用户状态(管理员才可修改): 1启用 2禁用 (格式: int64) |
| `tags` | object | 否 |  管理员才可修改的标签 |
| `tenant` | object | 否 |  |
| `tenant.appIDs` | array[string] | 否 |  企业拥有的应用ID列表 |
| `tenant.avatar` | string | 否 |  头像,更新时传filePath |
| `tenant.code` | string | 否 |  企业编码,创建的时候自动生成 |
| `tenant.desc` | string | 否 |  应用描述 |
| `tenant.id` | string | 否 |  id编号 |
| `tenant.name` | string | 否 |  企业名称 |
| `tenant.resource` | object | 否 |  |
| `tenant.resource.package` | object | 否 |  |
| `tenant.resource.quota` | object | 否 |  |
| `tenant.resource.usage` | object | 否 |  |
| `tenantCode` | string | 否 |  企业编码,只读 |
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
  "data": {
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/tenant/update \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "deptIDs": ["string"], "deviceCount": 1, "isTenantOwner": 1, "pubTags": {}, "roles": [{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "isFullAuth": 1, "name": "示例名称", "status": 1}], "status": 1, "tags": {}, "tenant": {"appIDs": ["string"], "avatar": "string", "code": "string", "desc": "string", "id": "string", "name": "示例名称", "resource": {"package": {"expTime": "2026-01-01T00:00:00Z", "packageCode": "string", "packageName": "示例名称"}, "quota": {"aiTokenBalance": 1, "aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}, "usage": {"aiTokenBalanceUsed": 1, "aiTokenMonth": "string", "aiTokenUsed": 1, "deviceCount": 1, "userCount": 1}}}, "tenantCode": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "userID": "string"}'
```
