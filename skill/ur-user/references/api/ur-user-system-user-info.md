# ur-user system/user/info

创建用户信息 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/info/create` | 创建用户信息 | admin |
| POST | `/api/v1/system/user/info/delete` | 刪除用户 | admin |
| POST | `/api/v1/system/user/info/get-list` | 查询用户信息列表 | admin |
| POST | `/api/v1/system/user/info/get-one` | 获取用户信息 | admin |
| POST | `/api/v1/system/user/info/update` | 更新用户基本数据 | admin |

## 详细说明

### POST `/api/v1/system/user/info/create`

**说明**: 创建用户信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `info` | object | 是 |  |
| `info.avatar` | string | 否 |  头像 |
| `info.email` | string | 否 |  邮箱 |
| `info.nickName` | string | 否 |  昵称 |
| `info.phone` | string | 否 |  手机号 |
| `info.userID` | string | 否 |  用户id |
| `info.userName` | string | 否 |  用户名 |
| `roleIDs` | array[string] | 否 | 平台的角色编号列表 |

**请求示例**:
```json
{
  "info": {
    "avatar": "string",
    "email": "string",
    "nickName": "string",
    "phone": "string",
    "userID": "string",
    "userName": "string"
  },
  "roleIDs": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/info/create \
  --body '{"info": {"avatar": "string", "email": "string", "nickName": "string", "phone": "string", "userID": "string", "userName": "string"}, "roleIDs": ["string"]}'
```

### POST `/api/v1/system/user/info/delete`

**说明**: 刪除用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `userID` | string | 否 |  用户id |

**请求示例**:
```json
{
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
ur api /api/v1/system/user/info/delete \
  --body '{"userID": "string"}'
```

### POST `/api/v1/system/user/info/get-list`

**说明**: 查询用户信息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 |  |
| `email` | string | 否 |  邮箱 |
| `nickName` | string | 否 |  用户的昵称 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `phone` | string | 否 |  手机号 |
| `userIDs` | array[string] | 否 |  |
| `userName` | string | 否 | 用户名(唯一) |

**请求示例**:
```json
{
  "account": "string",
  "email": "string",
  "nickName": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "phone": "string",
  "userIDs": [
    "string"
  ],
  "userName": "string"
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
        "email": "string",
        "nickName": "string",
        "phone": "string",
        "userID": "string",
        "userName": "string"
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
ur api /api/v1/system/user/info/get-list \
  --body '{"account": "string", "email": "string", "nickName": "string", "page": {"page": 1, "pageSize": 1}, "phone": "string", "userIDs": ["string"], "userName": "string"}'
```

### POST `/api/v1/system/user/info/get-one`

**说明**: 获取用户信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `userID` | string | 否 |  用户id |
| `withTenant` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "userID": "string",
  "withTenant": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avatar": "string",
    "email": "string",
    "nickName": "string",
    "phone": "string",
    "userID": "string",
    "userName": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/info/get-one \
  --body '{"userID": "string", "withTenant": true}'
```

### POST `/api/v1/system/user/info/update`

**说明**: 更新用户基本数据

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  头像（传图片的filePath） |
| `email` | string | 否 |  邮箱,平台管理员可修改 |
| `nickName` | string | 否 |  用户昵称 |
| `phone` | string | 否 |  手机号,平台管理员可修改 |
| `sex` | integer | 否 |  性别（1:男，2:女，0:未知） (格式: int64) |
| `userID` | string | 是 |  用户ID（必填） |
| `userName` | string | 否 |  登录用户名,,平台管理员可修改 |

**请求示例**:
```json
{
  "avatar": "string",
  "email": "string",
  "nickName": "string",
  "phone": "string",
  "sex": 1,
  "userID": "string",
  "userName": "string"
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
ur api /api/v1/system/user/info/update \
  --body '{"avatar": "string", "email": "string", "nickName": "string", "phone": "string", "sex": 1, "userID": "string", "userName": "string"}'
```
