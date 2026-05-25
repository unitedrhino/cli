# ur-user system/user/self

获取用户应用列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/self/app/get-list` | 获取用户应用列表 | all |
| POST | `/api/v1/system/user/self/app/get-one` | 获取用户应用详情 | all |
| POST | `/api/v1/system/user/self/bind-account` | 绑定账号 | all |
| POST | `/api/v1/system/user/self/cancel` | 注销用户 | all |
| POST | `/api/v1/system/user/self/captcha` | 获取验证码 | public |
| POST | `/api/v1/system/user/self/change-pwd` | 更新用户密码 | all |
| POST | `/api/v1/system/user/self/forget-pwd` | 忘记密码 | public |
| POST | `/api/v1/system/user/self/get-one` | 获取用户信息 | all |
| POST | `/api/v1/system/user/self/login` | 用户登录 | public |
| POST | `/api/v1/system/user/self/logout` | 用户登出 | all |
| POST | `/api/v1/system/user/self/menu/get-list` | 获取用户菜单列表 | all |
| POST | `/api/v1/system/user/self/message/get-list` | 用户消息列表 | all |
| POST | `/api/v1/system/user/self/message/get-pending` | 用户待处理消息 | all |
| POST | `/api/v1/system/user/self/message/handle` | 用户消息标记已处理 | all |
| POST | `/api/v1/system/user/self/message/mark-all-read` | 用户消息全部已读 | all |
| POST | `/api/v1/system/user/self/message/multi-delete` | 用户消息批量删除 | all |
| POST | `/api/v1/system/user/self/message/multi-is-read` | 用户消息批量已读 | all |
| POST | `/api/v1/system/user/self/message/statistics` | 用户消息统计 | all |
| POST | `/api/v1/system/user/self/notify-preference/read` | 用户通知偏好读取 | all |
| POST | `/api/v1/system/user/self/notify-preference/update` | 用户通知偏好更新 | all |
| POST | `/api/v1/system/user/self/profile/get-list` | 获取用户配置列表 | all |
| POST | `/api/v1/system/user/self/profile/get-one` | 获取用户配置详情 | all |
| POST | `/api/v1/system/user/self/profile/update` | 更新用户配置 | all |
| POST | `/api/v1/system/user/self/register` | 普通用户注册 | public |
| POST | `/api/v1/system/user/self/resource/action/get-list` | 获取用户资源动作权限列表 | all |
| POST | `/api/v1/system/user/self/third-auth/start` | 第三方登录授权起跳 | public |
| POST | `/api/v1/system/user/self/third-login` | 第三方登录回调换平台登录态 | public |
| POST | `/api/v1/system/user/self/third-register` | 第三方补全注册 | public |
| POST | `/api/v1/system/user/self/update` | 更新用户基本数据 | all |
| POST | `/api/v1/system/user/self/user/search` | 精准搜索用户 | all |

## 详细说明

### POST `/api/v1/system/user/self/app/get-list`

**说明**: 获取用户应用列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `notAppID` | string | 否 |  排除指定应用ID（如排除控制台大厅 77） |

**请求示例**:
```json
{
  "notAppID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "app": {
          "agreementsMap": {},
          "appleConfig": {
            "appID": "string",
            "bundleID": "string",
            "keyID": "string",
            "privateKey": "string",
            "redirectURI": "string",
            "teamID": "string"
          },
          "avatar": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "dingConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          },
          "githubConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          },
          "googleConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          },
          "id": "string",
          "isSysCreated": 1,
          "isUseMenu": 1,
          "isUseProxy": 1,
          "loginTypes": [
            "string"
          ],
          "menuRole": "string",
          "name": "示例名称",
          "nativeConfig": {
            "filePath": "string",
            "version": "string",
            "versionDesc": "string"
          },
          "proxy": "string",
          "sort": 1,
          "status": 1,
          "subType": "string",
          "tenantCode": "string",
          "tenantName": "示例名称",
          "trialTime": 1,
          "type": "string",
          "url": "string",
          "useBy": "string",
          "wxMiniConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          },
          "wxOpenConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          }
        },
        "appID": "string",
        "createdTime": "string",
        "expTime": 1,
        "id": "string",
        "sort": 1,
        "status": 1,
        "tenantCode": "string"
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
ur api /api/v1/system/user/self/app/get-list \
  --body '{"notAppID": "string"}'
```

### POST `/api/v1/system/user/self/app/get-one`

**说明**: 获取用户应用详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  |
| `subType` | string | 是 | wx:微信小程序  dingTalk:钉钉小程序 |
| `type` | string | 是 |   app:应用  mini:小程序 |

**请求示例**:
```json
{
  "appID": "string",
  "subType": "string",
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "appID": "string",
    "name": "示例名称",
    "tenantCodes": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/app/get-one \
  --body '{"appID": "string", "subType": "string", "type": "string"}'
```

### POST `/api/v1/system/user/self/bind-account`

**说明**: 绑定账号

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 手机号注册时填写手机号 email填写邮箱 |
| `code` | string | 否 | 验证码    微信登录填code 账号密码登录时填写密码 |
| `codeID` | string | 否 | 验证码编号 微信登录填state |
| `type` | string | 是 | 验证方式:phone手机号注册 wxOpen 微信开放平台登录 wxIn 微信内 wxMiniP 微信小程序 可选: ['phone', 'email', 'wxOpen', 'wxIn', 'wxMiniP', 'dingApp'] |

**请求示例**:
```json
{
  "account": "string",
  "code": "******",
  "codeID": "string",
  "type": "phone"
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
ur api /api/v1/system/user/self/bind-account \
  --body '{"account": "string", "code": "******", "codeID": "string", "type": "phone"}'
```

### POST `/api/v1/system/user/self/cancel`

**说明**: 注销用户

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/cancel \
  --body '{}'
```

### POST `/api/v1/system/user/self/captcha`

**说明**: 获取验证码

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 短信验证时填写手机号,邮箱验证时填写邮箱 |
| `code` | string | 否 | 验证码    微信邮箱验证登录填code |
| `codeID` | string | 否 | 验证码编号 微信邮箱验证登录填state |
| `type` | string | 是 | 验证方式:短信验证,图片验证码 可选: ['phone', 'image', 'email'] |
| `use` | string | 是 | 用途 可选: ['login', 'register', 'changePwd', 'forgetPwd', 'bindAccount'] |

**请求示例**:
```json
{
  "account": "string",
  "code": "string",
  "codeID": "string",
  "type": "phone",
  "use": "login"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "codeID": "string",
    "expire": 1,
    "url": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/captcha \
  --body '{"account": "string", "code": "string", "codeID": "string", "type": "phone", "use": "login"}'
```

### POST `/api/v1/system/user/self/change-pwd`

**说明**: 更新用户密码

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 | 验证码    微信登录填code 账号密码登录时填写密码 |
| `codeID` | string | 否 | 验证码编号 微信登录填state |
| `oldPassword` | string | 否 | pwd方式需要填写 如果原来没有设置过密码可以不填 输入原密码修改密码(sha256格式并需要输入图像验证码) |
| `password` | string | 否 | 密码 |
| `type` | string | 是 | 验证方式: phone手机号 email邮箱  pwd:原密码 可选: ['phone', 'email', 'pwd'] |

**请求示例**:
```json
{
  "code": "******",
  "codeID": "string",
  "oldPassword": "******",
  "password": "******",
  "type": "phone"
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
ur api /api/v1/system/user/self/change-pwd \
  --body '{"code": "******", "codeID": "string", "oldPassword": "******", "password": "******", "type": "phone"}'
```

### POST `/api/v1/system/user/self/forget-pwd`

**说明**: 忘记密码

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 手机号注册时填写手机号 账号密码注册时填写userName |
| `code` | string | 否 | 验证码    微信登录填code 账号密码登录时填写密码 |
| `codeID` | string | 否 | 验证码编号 微信登录填state |
| `password` | string | 是 | 密码 |
| `type` | string | 是 | 验证方式: phone手机号 email邮箱 可选: ['phone', 'email'] |

**请求示例**:
```json
{
  "account": "******",
  "code": "******",
  "codeID": "string",
  "password": "******",
  "type": "phone"
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
ur api /api/v1/system/user/self/forget-pwd \
  --body '{"account": "******", "code": "******", "codeID": "string", "password": "******", "type": "phone"}'
```

### POST `/api/v1/system/user/self/get-one`

**说明**: 获取用户信息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `withTenant` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
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
ur api /api/v1/system/user/self/get-one \
  --body '{"withTenant": true}'
```

### POST `/api/v1/system/user/self/login`

**说明**: 用户登录

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 登录账号(支持用户名,手机号,邮箱) 账号密码登录时需要填写 |
| `code` | string | 否 | 验证码    微信邮箱验证登录填code |
| `codeID` | string | 否 | 验证码编号 微信邮箱验证登录填state |
| `loginType` | string | 是 | 验证类型 phone 手机号 wxOpen 微信开放平台 wxIn 微信内 wxMiniP 微信小程序 pwd 账号密码 email 邮箱 可选: ['phone', 'wxOpen', 'wxIn', 'wxMiniP', 'dingApp', 'pwd', 'email'] |
| `password` | string | 否 | 密码，建议md5转换 密码登录时需要填写 |
| `pwdType` | integer | 否 | 账号密码登录时需要填写.0,无密码 1，明文 2，md5加密 (格式: int32) |

**请求示例**:
```json
{
  "account": "******",
  "code": "string",
  "codeID": "string",
  "loginType": "phone",
  "password": "******",
  "pwdType": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "info": {
      "avatar": "string",
      "email": "string",
      "nickName": "string",
      "phone": "string",
      "userID": "string",
      "userName": "string"
    },
    "token": {
      "accessExpire": "2026-01-01T00:00:00Z",
      "accessToken": "string",
      "refreshAfter": "2026-01-01T00:00:00Z"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/login \
  --body '{"account": "******", "code": "string", "codeID": "string", "loginType": "phone", "password": "******", "pwdType": 1}'
```

### POST `/api/v1/system/user/self/logout`

**说明**: 用户登出

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/logout \
  --body '{}'
```

### POST `/api/v1/system/user/self/menu/get-list`

**说明**: 获取用户菜单列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  应用编号 |
| `isCommon` | integer | 否 |  是否常用菜单 1-是 2-否 (格式: int64) |

**请求示例**:
```json
{
  "appID": "string",
  "isCommon": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "appID": "string",
        "children": [
          {
            "appID": "string",
            "children": [],
            "icon": "string",
            "id": "string",
            "isHideInMenu": 1,
            "isStar": 1,
            "menu": {
              "appID": "...",
              "body": "...",
              "children": "...",
              "component": "...",
              "createdTime": "...",
              "defaultIcon": "...",
              "defaultIsHideInMenu": "...",
              "defaultName": "...",
              "defaultSort": "...",
              "id": "...",
              "idPath": "...",
              "parentID": "...",
              "path": "...",
              "redirect": "...",
              "role": "...",
              "tenantCode": "...",
              "type": "..."
            },
            "menuID": "string",
            "name": "示例名称",
            "parentMenuID": "string",
            "sort": 1
          }
        ],
        "icon": "string",
        "id": "string",
        "isHideInMenu": 1,
        "isStar": 1,
        "menu": {
          "appID": "string",
          "body": "string",
          "children": [
            {
              "appID": "...",
              "body": "...",
              "children": "...",
              "component": "...",
              "createdTime": "...",
              "defaultIcon": "...",
              "defaultIsHideInMenu": "...",
              "defaultName": "...",
              "defaultSort": "...",
              "id": "...",
              "idPath": "...",
              "parentID": "...",
              "path": "...",
              "redirect": "...",
              "role": "...",
              "tenantCode": "...",
              "type": "..."
            }
          ],
          "component": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "defaultIcon": "string",
          "defaultIsHideInMenu": 1,
          "defaultName": "示例名称",
          "defaultSort": 1,
          "id": "string",
          "idPath": "string",
          "parentID": "string",
          "path": "string",
          "redirect": "string",
          "role": "string",
          "tenantCode": "string",
          "type": 1
        },
        "menuID": "string",
        "name": "示例名称",
        "parentMenuID": "string",
        "sort": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/menu/get-list \
  --body '{"appID": "string", "isCommon": 1}'
```

### POST `/api/v1/system/user/self/message/get-list`

**说明**: 用户消息列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间范围 |
| `group` | string | 否 | 分组 |
| `isHandled` | integer | 否 | 是否已处理 1:已处理 2:未处理 (格式: int64) |
| `isRead` | integer | 否 | 是否已读 1:已读 2:未读 (格式: int64) |
| `notifyCode` | string | 否 | 通知编码 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `str1` | string | 否 |  |
| `str2` | string | 否 |  |
| `str3` | string | 否 |  |
| `type` | string | 否 | 渠道类型 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "group": "string",
  "isHandled": 1,
  "isRead": 1,
  "notifyCode": "string",
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
  "str1": "string",
  "str2": "string",
  "str3": "string",
  "type": "string"
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
        "group": "string",
        "id": "string",
        "isHandled": 1,
        "isRead": 1,
        "message": {
          "actionText": "string",
          "actionURL": "string",
          "allowDismiss": "string",
          "body": "string",
          "coverImage": "string",
          "createdTime": "string",
          "emailFormat": "string",
          "expireTime": 1,
          "group": "string",
          "id": "string",
          "isGlobal": 1,
          "modalShowEndTime": "2026-01-01T00:00:00Z",
          "modalShowStartTime": "2026-01-01T00:00:00Z",
          "notificationID": "string",
          "notifyCode": "string",
          "notifyName": "示例名称",
          "notifyTime": 1,
          "priority": 1,
          "severity": "string",
          "showAsModal": "string",
          "str1": "string",
          "str2": "string",
          "str3": "string",
          "subject": "string",
          "type": "string",
          "videoUrl": "string"
        },
        "messageID": "string",
        "type": "string",
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
ur api /api/v1/system/user/self/message/get-list \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "group": "string", "isHandled": 1, "isRead": 1, "notifyCode": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "str1": "string", "str2": "string", "str3": "string", "type": "string"}'
```

### POST `/api/v1/system/user/self/message/get-pending`

**说明**: 用户待处理消息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

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
        "createdTime": "2026-01-01T00:00:00Z",
        "group": "string",
        "id": "string",
        "isHandled": 1,
        "isRead": 1,
        "message": {
          "actionText": "string",
          "actionURL": "string",
          "allowDismiss": "string",
          "body": "string",
          "coverImage": "string",
          "createdTime": "string",
          "emailFormat": "string",
          "expireTime": 1,
          "group": "string",
          "id": "string",
          "isGlobal": 1,
          "modalShowEndTime": "2026-01-01T00:00:00Z",
          "modalShowStartTime": "2026-01-01T00:00:00Z",
          "notificationID": "string",
          "notifyCode": "string",
          "notifyName": "示例名称",
          "notifyTime": 1,
          "priority": 1,
          "severity": "string",
          "showAsModal": "string",
          "str1": "string",
          "str2": "string",
          "str3": "string",
          "subject": "string",
          "type": "string",
          "videoUrl": "string"
        },
        "messageID": "string",
        "type": "string",
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
ur api /api/v1/system/user/self/message/get-pending \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/user/self/message/handle`

**说明**: 用户消息标记已处理

**权限**: all

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
ur api /api/v1/system/user/self/message/handle \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/user/self/message/mark-all-read`

**说明**: 用户消息全部已读

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/message/mark-all-read \
  --body '{}'
```

### POST `/api/v1/system/user/self/message/multi-delete`

**说明**: 用户消息批量删除

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[string] | 是 |  |

**请求示例**:
```json
{
  "ids": [
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
ur api /api/v1/system/user/self/message/multi-delete \
  --body '{"ids": ["string"]}'
```

### POST `/api/v1/system/user/self/message/multi-is-read`

**说明**: 用户消息批量已读

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[string] | 是 |  |

**请求示例**:
```json
{
  "ids": [
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
ur api /api/v1/system/user/self/message/multi-is-read \
  --body '{"ids": ["string"]}'
```

### POST `/api/v1/system/user/self/message/statistics`

**说明**: 用户消息统计

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "totalPending": 1,
    "totalUnread": 1,
    "unreadCounts": {}
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/message/statistics \
  --body '{}'
```

### POST `/api/v1/system/user/self/notify-preference/read`

**说明**: 用户通知偏好读取

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "quietChannels": [
      "string"
    ],
    "quietEndTime": "2026-01-01T00:00:00Z",
    "quietStartTime": "2026-01-01T00:00:00Z",
    "subscriptions": [
      {
        "group": "string",
        "isEnabled": 1,
        "type": "string"
      }
    ],
    "tenantCode": "string",
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/notify-preference/read \
  --body '{}'
```

### POST `/api/v1/system/user/self/notify-preference/update`

**说明**: 用户通知偏好更新

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `quietChannels` | array[string] | 否 | 免打扰渠道列表 |
| `quietEndTime` | string | 否 | 免打扰结束时间 HH:MM |
| `quietStartTime` | string | 否 | 免打扰开始时间 HH:MM |
| `subscriptions` | array[UserNotifySubscription] | 否 | 订阅控制列表 |
| `tenantCode` | string | 否 | 租户编码,只读 |
| `userID` | string | 否 | 用户ID,只读 |

**请求示例**:
```json
{
  "quietChannels": [
    "string"
  ],
  "quietEndTime": "2026-01-01T00:00:00Z",
  "quietStartTime": "2026-01-01T00:00:00Z",
  "subscriptions": [
    {
      "group": "string",
      "isEnabled": 1,
      "type": "string"
    }
  ],
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
ur api /api/v1/system/user/self/notify-preference/update \
  --body '{"quietChannels": ["string"], "quietEndTime": "2026-01-01T00:00:00Z", "quietStartTime": "2026-01-01T00:00:00Z", "subscriptions": [{"group": "string", "isEnabled": 1, "type": "string"}], "tenantCode": "string", "userID": "string"}'
```

### POST `/api/v1/system/user/self/profile/get-list`

**说明**: 获取用户配置列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `codes` | array[string] | 是 |  |

**请求示例**:
```json
{
  "codes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "profiles": [
      {
        "code": "string",
        "params": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/profile/get-list \
  --body '{"codes": ["string"]}'
```

### POST `/api/v1/system/user/self/profile/get-one`

**说明**: 获取用户配置详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `withProjects` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "code": "string",
  "withProjects": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "params": "string",
    "projects": [
      {
        "address": "string",
        "adminUser": {
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
        "adminUserID": "string",
        "alarmStatus": 1,
        "area": 1,
        "areaCount": 1,
        "areas": [
          {
            "areaID": "string",
            "areaIDPath": "string",
            "areaImg": "string",
            "areaName": "示例名称",
            "areaNamePath": "string",
            "children": [],
            "configFile": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "groupCount": 1,
            "isLeaf": 1,
            "isSysCreated": 1,
            "lowerLevelCount": 1,
            "parentAreaID": "string",
            "position": {
              "latitude": "...",
              "longitude": "..."
            },
            "projectID": "string",
            "sort": 1,
            "tags": {},
            "tenantCode": "string",
            "useBy": "string"
          }
        ],
        "attachments": [
          {
            "fileName": "string",
            "filePath": "string",
            "fileUrl": "string",
            "id": "string",
            "useBy": "string"
          }
        ],
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "deviceOnlineCount": 1,
        "isSysCreated": 1,
        "position": {
          "latitude": 1,
          "longitude": 1
        },
        "ppsm": 1,
        "projectID": "string",
        "projectImg": "string",
        "projectName": "示例名称",
        "sort": 1,
        "status": 1,
        "tags": {},
        "tenantCode": "string",
        "type": "string",
        "userCount": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/profile/get-one \
  --body '{"code": "string", "withProjects": true}'
```

### POST `/api/v1/system/user/self/profile/update`

**说明**: 更新用户配置

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `params` | string | 是 |  |

**请求示例**:
```json
{
  "code": "string",
  "params": "string"
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
ur api /api/v1/system/user/self/profile/update \
  --body '{"code": "string", "params": "string"}'
```

### POST `/api/v1/system/user/self/register`

**说明**: 普通用户注册

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 手机号注册时填写手机号 账号密码注册时填写userName |
| `code` | string | 是 | 验证码    微信登录填code 账号密码登录时填写密码 |
| `codeID` | string | 否 | 验证码编号 微信登录填state |
| `expand` | object | 否 | 拓展, 微信登录方式 phoneCode:获取手机号code  手机号和邮箱注册: wxOpenCode:如果需要同时绑定微信则在这里填写开放平台的用户code |
| `isWithLogin` | boolean | 否 | 注册成功同时返回token,自动登录 (格式: boolean) |
| `nickName` | string | 否 |  用户的昵称 |
| `password` | string | 否 | 密码 |
| `regType` | string | 是 | 可选: ['phone', 'email', 'wxOpen', 'wxIn', 'wxMiniP', 'pwd', 'dingApp'] |

**请求示例**:
```json
{
  "account": "******",
  "code": "******",
  "codeID": "string",
  "expand": {},
  "isWithLogin": true,
  "nickName": "string",
  "password": "******",
  "regType": "phone"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "info": {
      "avatar": "string",
      "email": "string",
      "nickName": "string",
      "phone": "string",
      "userID": "string",
      "userName": "string"
    },
    "token": {
      "accessExpire": "2026-01-01T00:00:00Z",
      "accessToken": "string",
      "refreshAfter": "2026-01-01T00:00:00Z"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/register \
  --body '{"account": "******", "code": "******", "codeID": "string", "expand": {}, "isWithLogin": true, "nickName": "string", "password": "******", "regType": "phone"}'
```

### POST `/api/v1/system/user/self/resource/action/get-list`

**说明**: 获取用户资源动作权限列表

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "action": "string",
        "resource": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/resource/action/get-list \
  --body '{}'
```

### POST `/api/v1/system/user/self/third-auth/start`

**说明**: 第三方登录授权起跳

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `platform` | string | 否 |  |
| `provider` | string | 是 | 可选: ['google', 'github', 'apple'] |

**请求示例**:
```json
{
  "platform": "string",
  "provider": "google"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "authUrl": "string",
    "state": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/third-auth/start \
  --body '{"platform": "string", "provider": "google"}'
```

### POST `/api/v1/system/user/self/third-login`

**说明**: 第三方登录回调换平台登录态

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `platform` | string | 否 |  |
| `provider` | string | 是 | 可选: ['google', 'github', 'apple'] |
| `state` | string | 是 |  |

**请求示例**:
```json
{
  "code": "string",
  "platform": "string",
  "provider": "google",
  "state": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "login": {
      "info": {
        "avatar": "string",
        "email": "string",
        "nickName": "string",
        "phone": "string",
        "userID": "string",
        "userName": "string"
      },
      "token": {
        "accessExpire": "2026-01-01T00:00:00Z",
        "accessToken": "string",
        "refreshAfter": "2026-01-01T00:00:00Z"
      }
    },
    "message": "string",
    "pendingToken": "string",
    "provider": "string",
    "status": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/third-login \
  --body '{"code": "string", "platform": "string", "provider": "google", "state": "string"}'
```

### POST `/api/v1/system/user/self/third-register`

**说明**: 第三方补全注册

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 是 |  |
| `code` | string | 是 |  |
| `codeID` | string | 否 |  |
| `nickName` | string | 否 |  |
| `password` | string | 否 |  |
| `pendingToken` | string | 是 |  |
| `regType` | string | 是 | 可选: ['phone', 'email'] |

**请求示例**:
```json
{
  "account": "string",
  "code": "string",
  "codeID": "string",
  "nickName": "string",
  "password": "string",
  "pendingToken": "string",
  "regType": "phone"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "info": {
      "avatar": "string",
      "email": "string",
      "nickName": "string",
      "phone": "string",
      "userID": "string",
      "userName": "string"
    },
    "token": {
      "accessExpire": "2026-01-01T00:00:00Z",
      "accessToken": "string",
      "refreshAfter": "2026-01-01T00:00:00Z"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/third-register \
  --body '{"account": "string", "code": "string", "codeID": "string", "nickName": "string", "password": "string", "pendingToken": "string", "regType": "phone"}'
```

### POST `/api/v1/system/user/self/update`

**说明**: 更新用户基本数据

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  头像（传图片的filePath） |
| `nickName` | string | 否 |  用户昵称 |
| `sex` | integer | 否 |  性别（1:男，2:女，0:未知） (格式: int64) |

**请求示例**:
```json
{
  "avatar": "string",
  "nickName": "string",
  "sex": 1
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
ur api /api/v1/system/user/self/update \
  --body '{"avatar": "string", "nickName": "string", "sex": 1}'
```

### POST `/api/v1/system/user/self/user/search`

**说明**: 精准搜索用户

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 是 |  |

**请求示例**:
```json
{
  "account": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avatar": "string",
    "nickName": "string",
    "status": 1,
    "userID": "string",
    "userName": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/user/search \
  --body '{"account": "string"}'
```
