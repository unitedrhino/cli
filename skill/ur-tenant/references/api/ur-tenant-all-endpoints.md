# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/agreement/create` | 添加协议 | admin |
| POST | `/api/v1/system/tenant/agreement/delete` | 删除协议 | admin |
| POST | `/api/v1/system/tenant/agreement/get-list` | 获取协议列表 | admin |
| POST | `/api/v1/system/tenant/agreement/get-one` | 获取协议详情 | public |
| POST | `/api/v1/system/tenant/agreement/update` | 更新协议 | admin |
| POST | `/api/v1/system/tenant/app/create` | 绑定企业应用 | all |
| POST | `/api/v1/system/tenant/app/delete` | 解绑企业应用 | platform |
| POST | `/api/v1/system/tenant/app/get-list` | 获取企业应用列表 | admin |
| POST | `/api/v1/system/tenant/app/get-one` | 获取应用信息 | platform |
| POST | `/api/v1/system/tenant/app/update` | 更新企业应用 | platform |
| POST | `/api/v1/system/tenant/app/menu/get-list` | 获取企业应用菜单列表 | admin |
| POST | `/api/v1/system/tenant/app/menu/update` | 更新企业应用菜单 | admin |
| POST | `/api/v1/system/tenant/config/get-one` | 获取企业配置 | admin |
| POST | `/api/v1/system/tenant/config/update` | 更新企业配置 | admin |
| POST | `/api/v1/system/tenant/core/get-list` | 搜索企业信息 | public |
| POST | `/api/v1/system/tenant/core/get-one` | 获取企业信息 | public |
| POST | `/api/v1/system/tenant/info/create` | 添加企业 | all |
| POST | `/api/v1/system/tenant/info/delete` | 删除企业 | platform |
| POST | `/api/v1/system/tenant/info/get-list` | 获取企业列表 | platform |
| POST | `/api/v1/system/tenant/info/get-one` | 获取企业详情 | admin |
| POST | `/api/v1/system/tenant/info/transfer` | 转让企业（仅企业所有者） | admin |
| POST | `/api/v1/system/tenant/info/update` | 更新企业 | admin |
| POST | `/api/v1/system/tenant/renewal/direct-grant` | 直接授权 | platform |
| POST | `/api/v1/system/tenant/renewal/get-list` | 获取续期列表 | admin |
| POST | `/api/v1/system/tenant/renewal/renew` | 续费 | admin |
| POST | `/api/v1/system/tenant/user/batch-create` | 批量添加用户加入企业 | admin |
| POST | `/api/v1/system/tenant/user/delete` | 删除企业用户 | admin |
| POST | `/api/v1/system/tenant/user/get-list` | 获取企业用户列表 | admin |
| POST | `/api/v1/system/tenant/user/get-one` | 获取企业用户详情,会同时返回所拥有的角色列表 | admin |
| POST | `/api/v1/system/tenant/user/invite` | 邀请用户加入企业 | admin |
| POST | `/api/v1/system/tenant/user/invite-code/gen` | 生成企业用户邀请码 | admin |
| POST | `/api/v1/system/tenant/user/invite-code/get-one` | 获取当前有效的企业用户邀请码 | admin |
| POST | `/api/v1/system/tenant/user/invite-pending/delete` | 删除待处理邀请 | admin |
| POST | `/api/v1/system/tenant/user/invite-pending/get-list` | 获取待处理邀请列表 | admin |
| POST | `/api/v1/system/tenant/user/invite-send` | 发送邀请（支持邮件和手机号） | admin |
| POST | `/api/v1/system/tenant/user/update` | 更新企业用户 | admin |
| POST | `/api/v1/system/tenant/user/role/batch-update` | 更新企业用户的角色列表 | admin |
| POST | `/api/v1/system/tenant/user/role/get-list` | 获取企业用户角色列表 | admin |

## 端点详情

### POST `/api/v1/system/tenant/agreement/create`

**说明**: 添加协议

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  编号 |
| `content` | string | 否 | 协议内容(只有详情会返回) |
| `id` | string | 否 |  编号 |
| `name` | string | 否 |  协议名称 |
| `title` | string | 否 | 协议标题 |

**请求示例**:
```json
{
  "code": "string",
  "content": "string",
  "id": "string",
  "name": "示例名称",
  "title": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/agreement/create \
  --body '{"code": "string", "content": "string", "id": "string", "name": "示例名称", "title": "string"}'
```

### POST `/api/v1/system/tenant/agreement/delete`

**说明**: 删除协议

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

**请求示例**:
```json
{
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
ur api /api/v1/system/tenant/agreement/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/tenant/agreement/get-list`

**说明**: 获取协议列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
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
        "code": "string",
        "content": "string",
        "id": "string",
        "name": "示例名称",
        "title": "string"
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
ur api /api/v1/system/tenant/agreement/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/system/tenant/agreement/get-one`

**说明**: 获取协议详情

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "code": "string",
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "content": "string",
    "id": "string",
    "name": "示例名称",
    "title": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/agreement/get-one \
  --body '{"code": "string", "id": "string"}'
```

### POST `/api/v1/system/tenant/agreement/update`

**说明**: 更新协议

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  编号 |
| `content` | string | 否 | 协议内容(只有详情会返回) |
| `id` | string | 否 |  编号 |
| `name` | string | 否 |  协议名称 |
| `title` | string | 否 | 协议标题 |

**请求示例**:
```json
{
  "code": "string",
  "content": "string",
  "id": "string",
  "name": "示例名称",
  "title": "string"
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
ur api /api/v1/system/tenant/agreement/update \
  --body '{"code": "string", "content": "string", "id": "string", "name": "示例名称", "title": "string"}'
```

### POST `/api/v1/system/tenant/app/create`

**说明**: 绑定企业应用

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  应用ID |
| `licenseCode` | string | 否 |  授权码（付费应用必填，免费/试用应用可留空） |
| `tenantCode` | string | 否 |  指定企业,只有平台可以指定 |

**请求示例**:
```json
{
  "appID": "string",
  "licenseCode": "string",
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
ur api /api/v1/system/tenant/app/create \
  --body '{"appID": "string", "licenseCode": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/app/delete`

**说明**: 解绑企业应用

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  |
| `tenantCode` | string | 否 | 指定企业,只有平台可以指定 |

**请求示例**:
```json
{
  "appID": "string",
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
ur api /api/v1/system/tenant/app/delete \
  --body '{"appID": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/app/get-list`

**说明**: 获取企业应用列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  |
| `appIDs` | array[string] | 否 |  |
| `notAppID` | string | 否 | 排除某个应用ID |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  应用状态 1 正常(默认)，2 禁用，3 过期 (格式: int64) |
| `subType` | string | 否 | 子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `tenantCode` | string | 否 | 指定企业,企业下只能过滤 common或他自己的  common的是通用应用,非common的是企业应用 |
| `type` | string | 否 | 应用类型 web:web页面  native:原生  mini:小程序 |
| `useBy` | string | 否 |  使用对象 client:登录直接授权该应用 admin:管理员使用,需要邀请 |
| `withApp` | boolean | 否 | 同时返回应用信息 (格式: boolean) |

**请求示例**:
```json
{
  "appID": "string",
  "appIDs": [
    "string"
  ],
  "notAppID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": 1,
  "subType": "string",
  "tenantCode": "string",
  "type": "string",
  "useBy": "string",
  "withApp": true
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
ur api /api/v1/system/tenant/app/get-list \
  --body '{"appID": "string", "appIDs": ["string"], "notAppID": "string", "page": {"page": 1, "pageSize": 1}, "status": 1, "subType": "string", "tenantCode": "string", "type": "string", "useBy": "string", "withApp": true}'
```

### POST `/api/v1/system/tenant/app/get-one`

**说明**: 获取应用信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  |
| `tenantCode` | string | 否 | 指定企业,只有平台可以指定 |

**请求示例**:
```json
{
  "appID": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
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
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/app/get-one \
  --body '{"appID": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/app/update`

**说明**: 更新企业应用

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `app` | object | 否 |  |
| `app.agreementsMap` | object | 否 |  按请求协议 code 回填的协议映射 |
| `app.appleConfig` | object | 否 |  |
| `app.appleConfig.appID` | string | 否 |  |
| `app.appleConfig.bundleID` | string | 否 |  |
| `app.appleConfig.keyID` | string | 否 |  |
| `app.appleConfig.privateKey` | string | 否 |  |
| `app.appleConfig.redirectURI` | string | 否 |  |
| `app.appleConfig.teamID` | string | 否 |  |
| `app.avatar` | string | 否 |  头像 更新的时候传filePath即可 |
| `app.createdTime` | string | 否 |  创建时间 |
| `app.desc` | string | 否 |  应用描述 |
| `app.dingConfig` | object | 否 |  |
| `app.dingConfig.appID` | string | 否 |  |
| `app.dingConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `app.dingConfig.appSecret` | string | 否 |  |
| `app.githubConfig` | object | 否 |  |
| `app.githubConfig.appID` | string | 否 |  |
| `app.githubConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `app.githubConfig.appSecret` | string | 否 |  |
| `app.googleConfig` | object | 否 |  |
| `app.googleConfig.appID` | string | 否 |  |
| `app.googleConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `app.googleConfig.appSecret` | string | 否 |  |
| `app.id` | string | 否 |  编号,创建时自动生成(默认),也可以指定 |
| `app.isSysCreated` | integer | 否 |  是否系统创建 1-是 2-否（系统创建的应用不允许删除，绑定时自动添加） (格式: int64) |
| `app.isUseMenu` | integer | 否 |  是否使用菜单 1-是 2-否 (格式: int64) |
| `app.isUseProxy` | integer | 否 |  是否使用代理 1-是 2-否 (格式: int64) |
| `app.loginTypes` | array[string] | 否 | 选项取决于 app里是否配置了 支持的登录类型(不填支持全部登录方式):    "email":邮箱(系统配置里配置了才能开启) "phone":手机号(系统配置里配置了才能开启)  "wxMiniP":微信小程序(app里配置了微信小程序才能开启)  "wxOpen": 微信开放平台登录(app里配置了微信开放才能开启)   "ding":钉钉应用(包含小程序,h5等方式,app里配置了钉钉才能开启)  "pwd":账号密码注册 |
| `app.menuRole` | string | 否 |  菜单角色类型 all-支持两类(菜单需选role) platform-平台管理员 tenant-企业管理员 |
| `app.name` | string | 否 |  应用名称 |
| `app.nativeConfig` | object | 否 |  |
| `app.nativeConfig.filePath` | string | 否 |  文件路径,拿来下载文件 |
| `app.nativeConfig.version` | string | 否 |  应用版本 |
| `app.nativeConfig.versionDesc` | string | 否 | 版本说明 |
| `app.proxy` | string | 否 |  代理配置(JSON字符串) |
| `app.sort` | integer | 否 | 排序,不能小于等于0 (格式: int64) |
| `app.status` | integer | 否 |  应用状态 1:上架, 2:下架(默认) (格式: int64) |
| `app.subType` | string | 否 |  子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `app.tenantCode` | string | 否 |  企业编码,common的是通用应用,非common的是企业应用 |
| `app.tenantName` | string | 否 |  企业名称 |
| `app.trialTime` | integer | 否 | 试用时间(单位为天,为0不限制,如果为0,是免费使用的,如果为-1 则不允许自己创建) (格式: int64) |
| `app.type` | string | 否 |  应用类型 web:web页面  native:原生  mini:小程序 |
| `app.url` | string | 否 |  应用地址 |
| `app.useBy` | string | 否 |  使用对象 client:登录直接授权该应用 admin:管理员使用,需要邀请 |
| `app.wxMiniConfig` | object | 否 |  |
| `app.wxMiniConfig.appID` | string | 否 |  |
| `app.wxMiniConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `app.wxMiniConfig.appSecret` | string | 否 |  |
| `app.wxOpenConfig` | object | 否 |  |
| `app.wxOpenConfig.appID` | string | 否 |  |
| `app.wxOpenConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `app.wxOpenConfig.appSecret` | string | 否 |  |
| `appID` | string | 是 |  应用ID |
| `createdTime` | string | 否 | 只读 |
| `expTime` | integer | 否 | 过期时间(平台才可以修改),为0不限制 (格式: int64) |
| `id` | string | 是 |  编号 |
| `sort` | integer | 否 | 排序,不能小于等于0,默认为1 (格式: int64) |
| `status` | integer | 否 |  应用状态(平台才可以修改) 1 正常(默认)，2 禁用，3 过期 (格式: int64) |
| `tenantCode` | string | 否 | 指定企业,只有平台可以指定 |

**请求示例**:
```json
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
ur api /api/v1/system/tenant/app/update \
  --body '{"app": {"agreementsMap": {}, "appleConfig": {"appID": "string", "bundleID": "string", "keyID": "string", "privateKey": "string", "redirectURI": "string", "teamID": "string"}, "avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "dingConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "githubConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "googleConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "id": "string", "isSysCreated": 1, "isUseMenu": 1, "isUseProxy": 1, "loginTypes": ["string"], "menuRole": "string", "name": "示例名称", "nativeConfig": {"filePath": "string", "version": "string", "versionDesc": "string"}, "proxy": "string", "sort": 1, "status": 1, "subType": "string", "tenantCode": "string", "tenantName": "示例名称", "trialTime": 1, "type": "string", "url": "string", "useBy": "string", "wxMiniConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "wxOpenConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}}, "appID": "string", "createdTime": "string", "expTime": 1, "id": "string", "sort": 1, "status": 1, "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/app/menu/get-list`

**说明**: 获取企业应用菜单列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  |
| `isRetTree` | boolean | 否 | 格式: boolean |
| `isStar` | integer | 否 |  是否收藏 1:是 2:否 (格式: int64) |
| `menuIDs` | array[string] | 否 |  |
| `tenantCode` | string | 否 | 指定企业,只有平台可以指定 |

**请求示例**:
```json
{
  "appID": "string",
  "isRetTree": true,
  "isStar": 1,
  "menuIDs": [
    "string"
  ],
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
ur api /api/v1/system/tenant/app/menu/get-list \
  --body '{"appID": "string", "isRetTree": true, "isStar": 1, "menuIDs": ["string"], "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/app/menu/update`

**说明**: 更新企业应用菜单

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  应用ID,新增必填,不可修改 |
| `children` | array[TenantMenu] | 否 |  |
| `icon` | string | 否 |  图标 |
| `id` | string | 否 |  编号 |
| `isHideInMenu` | integer | 否 |  是否隐藏菜单 1-是 2-否 (格式: int64) |
| `isStar` | integer | 否 | 是否收藏 1:是 2:否 (格式: int64) |
| `menu` | object | 否 |  |
| `menu.appID` | string | 否 |  模块编码 |
| `menu.body` | string | 否 |  前端自定义字段 |
| `menu.children` | array[MenuInfo] | 否 |  子菜单列表 |
| `menu.component` | string | 否 |  页面 |
| `menu.createdTime` | string | 否 |  创建时间 |
| `menu.defaultIcon` | string | 否 |  默认图标 |
| `menu.defaultIsHideInMenu` | integer | 否 |  默认是否隐藏菜单 1-是 2-否 (格式: int64) |
| `menu.defaultName` | string | 否 |  默认 菜单名称 |
| `menu.defaultSort` | integer | 否 |  默认左侧table排序序号 (格式: int64) |
| `menu.id` | string | 否 |  编号 |
| `menu.idPath` | string | 否 |  菜单ID路径 |
| `menu.parentID` | string | 否 |  父菜单ID，一级菜单为1 |
| `menu.path` | string | 否 |  系统的path |
| `menu.redirect` | string | 否 |  路由重定向 |
| `menu.role` | string | 否 |  角色类型 仅App.menuRole=all时有意义 all:两类都可以看 platform:平台管理员 tenant:企业管理员 |
| `menu.tenantCode` | string | 否 |  跟着应用走, common的是通用应用,非common的是企业应用 |
| `menu.type` | integer | 否 |  类型 1：菜单或者页面 2：iframe嵌入 3：外链跳转 (格式: int64) |
| `menuID` | string | 否 |  菜单ID,不可修改 |
| `name` | string | 否 |  菜单名称 |
| `parentMenuID` | string | 否 |  父菜单ID，一级菜单为1 |
| `sort` | integer | 否 |  左侧table排序序号 (格式: int64) |

**请求示例**:
```json
{
  "appID": "string",
  "children": [
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
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/app/menu/update \
  --body '{"appID": "string", "children": [{"appID": "string", "children": [{"appID": "string", "children": [], "icon": "string", "id": "string", "isHideInMenu": 1, "isStar": 1, "menu": {"appID": "...", "body": "...", "children": "...", "component": "...", "createdTime": "...", "defaultIcon": "...", "defaultIsHideInMenu": "...", "defaultName": "...", "defaultSort": "...", "id": "...", "idPath": "...", "parentID": "...", "path": "...", "redirect": "...", "role": "...", "tenantCode": "...", "type": "..."}, "menuID": "string", "name": "示例名称", "parentMenuID": "string", "sort": 1}], "icon": "string", "id": "string", "isHideInMenu": 1, "isStar": 1, "menu": {"appID": "string", "body": "string", "children": [{"appID": "...", "body": "...", "children": "...", "component": "...", "createdTime": "...", "defaultIcon": "...", "defaultIsHideInMenu": "...", "defaultName": "...", "defaultSort": "...", "id": "...", "idPath": "...", "parentID": "...", "path": "...", "redirect": "...", "role": "...", "tenantCode": "...", "type": "..."}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}, "menuID": "string", "name": "示例名称", "parentMenuID": "string", "sort": 1}], "icon": "string", "id": "string", "isHideInMenu": 1, "isStar": 1, "menu": {"appID": "string", "body": "string", "children": [{"appID": "string", "body": "string", "children": [{"appID": "...", "body": "...", "children": "...", "component": "...", "createdTime": "...", "defaultIcon": "...", "defaultIsHideInMenu": "...", "defaultName": "...", "defaultSort": "...", "id": "...", "idPath": "...", "parentID": "...", "path": "...", "redirect": "...", "role": "...", "tenantCode": "...", "type": "..."}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}, "menuID": "string", "name": "示例名称", "parentMenuID": "string", "sort": 1}'
```

### POST `/api/v1/system/tenant/config/get-one`

**说明**: 获取企业配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `tenantCode` | string | 是 |  |

**请求示例**:
```json
{
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string",
    "registerRoleID": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/config/get-one \
  --body '{"tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/config/update`

**说明**: 更新企业配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  |
| `registerRoleID` | string | 否 | 注册分配的角色id |
| `tenantCode` | string | 否 | 企业编码 |

**请求示例**:
```json
{
  "id": "string",
  "registerRoleID": "string",
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
ur api /api/v1/system/tenant/config/update \
  --body '{"id": "string", "registerRoleID": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/core/get-list`

**说明**: 搜索企业信息

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 | 过滤绑定了指定app的企业 |
| `code` | string | 否 |  编号 |
| `name` | string | 否 |  企业名称 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "appID": "string",
  "code": "string",
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
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
  --body '{"appID": "string", "code": "string", "name": "示例名称", "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/system/tenant/core/get-one`

**说明**: 获取企业信息

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "code": "string",
  "id": "string"
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
  --body '{"code": "string", "id": "string"}'
```

### POST `/api/v1/system/tenant/info/create`

**说明**: 添加企业

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  头像,更新时传filePath |
| `code` | string | 否 |  企业编码,创建的时候自动生成 |
| `config` | object | 否 |  |
| `config.id` | string | 否 |  |
| `config.registerRoleID` | string | 否 | 注册分配的角色id |
| `config.tenantCode` | string | 否 | 企业编码 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  应用描述 |
| `id` | string | 否 |  id编号 |
| `licenseCode` | string | 否 |  授权码（创建时可选，写入后应用套餐配额） |
| `name` | string | 否 |  企业名称 |
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
| `resource.quota.deviceMsgPerDay` | integer | 否 |  企业每天设备消息总量上限 (格式: int64) |
| `resource.quota.deviceNum` | integer | 否 |  企业下的设备数量限制，0 为不限制 (格式: int64) |
| `resource.quota.diskSpaceGB` | integer | 否 |  磁盘空间配额（GB） (格式: int64) |
| `resource.quota.userNum` | integer | 否 |  企业下的用户数量限制，0 为不限制 (格式: int64) |
| `resource.usage` | object | 否 |  |
| `resource.usage.aiTokenBalanceUsed` | integer | 否 |  AI Token 余额已用 (格式: int64) |
| `resource.usage.aiTokenMonth` | string | 否 |  AI Token 用量月份 |
| `resource.usage.aiTokenUsed` | integer | 否 |  AI Token 月已用 (格式: int64) |
| `resource.usage.deviceCount` | integer | 否 | 企业下设备统计,只读 (格式: int64) |
| `resource.usage.userCount` | integer | 否 | 企业下用户统计,只读 (格式: int64) |
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
    "id": "string"
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

**说明**: 删除企业

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "code": "string",
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
ur api /api/v1/system/tenant/info/delete \
  --body '{"code": "string", "id": "string"}'
```

### POST `/api/v1/system/tenant/info/get-list`

**说明**: 获取企业列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  企业编号 |
| `isGetAll` | boolean | 否 | 是否返回所有,只有平台管理员有权限 (格式: boolean) |
| `name` | string | 否 |  企业名称 |
| `ownerUserID` | string | 否 |  超级管理员id,只有平台管理员有权限 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 企业状态 (格式: int64) |
| `tenantCodes` | array[string] | 否 |  指定企业编码列表，非空时按编码精确返回 |
| `withOwnerUser` | boolean | 否 | 同时获取管理员核心信息 (格式: boolean) |

**请求示例**:
```json
{
  "code": "string",
  "isGetAll": true,
  "name": "示例名称",
  "ownerUserID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
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
  --body '{"code": "string", "isGetAll": true, "name": "示例名称", "ownerUserID": "string", "page": {"page": 1, "pageSize": 1}, "status": 1, "tenantCodes": ["string"], "withOwnerUser": true}'
```

### POST `/api/v1/system/tenant/info/get-one`

**说明**: 获取企业详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "code": "string",
  "id": "string"
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
  --body '{"code": "string", "id": "string"}'
```

### POST `/api/v1/system/tenant/info/transfer`

**说明**: 转让企业（仅企业所有者）

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `newOwnerUserID` | string | 是 |  新的企业所有者用户ID |

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

**说明**: 更新企业

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  头像,更新时传filePath |
| `code` | string | 否 |  企业编码,创建的时候自动生成 |
| `config` | object | 否 |  |
| `config.id` | string | 否 |  |
| `config.registerRoleID` | string | 否 | 注册分配的角色id |
| `config.tenantCode` | string | 否 | 企业编码 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  应用描述 |
| `id` | string | 否 |  id编号 |
| `licenseCode` | string | 否 |  授权码（创建时可选，写入后应用套餐配额） |
| `name` | string | 否 |  企业名称 |
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
| `resource.quota.deviceMsgPerDay` | integer | 否 |  企业每天设备消息总量上限 (格式: int64) |
| `resource.quota.deviceNum` | integer | 否 |  企业下的设备数量限制，0 为不限制 (格式: int64) |
| `resource.quota.diskSpaceGB` | integer | 否 |  磁盘空间配额（GB） (格式: int64) |
| `resource.quota.userNum` | integer | 否 |  企业下的用户数量限制，0 为不限制 (格式: int64) |
| `resource.usage` | object | 否 |  |
| `resource.usage.aiTokenBalanceUsed` | integer | 否 |  AI Token 余额已用 (格式: int64) |
| `resource.usage.aiTokenMonth` | string | 否 |  AI Token 用量月份 |
| `resource.usage.aiTokenUsed` | integer | 否 |  AI Token 月已用 (格式: int64) |
| `resource.usage.deviceCount` | integer | 否 | 企业下设备统计,只读 (格式: int64) |
| `resource.usage.userCount` | integer | 否 | 企业下用户统计,只读 (格式: int64) |
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
| `tenantCode` | string | 是 |  企业编码 |

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
| `keyword` | string | 否 |  关键字，按企业名、应用名或企业编码模糊匹配 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  到期状态：all 全部 normal 正常 expiring 即将到期 expired 已过期 |
| `targetType` | string | 否 |  续期对象类型：all 全部 tenant 企业 app 应用 |
| `tenantCode` | string | 否 |  指定企业编码，仅平台管理员可传 |

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
| `tenantCode` | string | 否 |  所属企业编码，平台管理员可显式指定 |

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

### POST `/api/v1/system/tenant/user/batch-create`

**说明**: 批量添加用户加入企业

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

**说明**: 删除企业用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `tenantCode` | string | 否 |  企业代码,(平台管理员查看其他企业下的用户时使用,普通企业只能查看自己的用户) |
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

**说明**: 获取企业用户列表

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
| `status` | integer | 否 |  企业状态（1:启用，2:禁用） (格式: int64) |
| `tenantCode` | string | 否 |  企业代码,(平台管理员查看其他企业下的用户时使用,普通企业只能查看自己的用户) |
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

**说明**: 获取企业用户详情,会同时返回所拥有的角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `tenantCode` | string | 否 |  企业代码,(平台管理员查看其他企业下的用户时使用,普通企业只能查看自己的用户) |
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

**说明**: 邀请用户加入企业

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

**说明**: 生成企业用户邀请码

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `roleID` | string | 是 |  分配的角色ID（必填） |
| `tenantCode` | string | 否 |  企业编码（可选，不填取header） |

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

**说明**: 获取当前有效的企业用户邀请码

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
| `tenantCode` | string | 否 |  企业编码（可选，不填取header） |

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
| `tenantCode` | string | 否 |  企业编码（可选，不填取header） |

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
| `tenantCode` | string | 否 |  企业编码（可选，不填取header） |
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

**说明**: 更新企业用户

**权限**: admin

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/user/update \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "deptIDs": ["string"], "deviceCount": 1, "isTenantOwner": 1, "pubTags": {}, "roles": [{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "isFullAuth": 1, "name": "示例名称", "status": 1}], "status": 1, "tags": {}, "tenant": {"appIDs": ["string"], "avatar": "string", "code": "string", "desc": "string", "id": "string", "name": "示例名称", "resource": {"package": {"expTime": "2026-01-01T00:00:00Z", "packageCode": "string", "packageName": "示例名称"}, "quota": {"aiTokenBalance": 1, "aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}, "usage": {"aiTokenBalanceUsed": 1, "aiTokenMonth": "string", "aiTokenUsed": 1, "deviceCount": 1, "userCount": 1}}}, "tenantCode": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "userID": "string"}'
```

### POST `/api/v1/system/tenant/user/role/batch-update`

**说明**: 更新企业用户的角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `roleIDs` | array[string] | 是 |  角色编号列表 |
| `tenantCode` | string | 否 |  企业编码（可选，不填取当前用户企业） |
| `userID` | string | 是 |  用户ID |

**请求示例**:
```json
{
  "roleIDs": [
    "string"
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
ur api /api/v1/system/tenant/user/role/batch-update \
  --body '{"roleIDs": ["string"], "tenantCode": "string", "userID": "string"}'
```

### POST `/api/v1/system/tenant/user/role/get-list`

**说明**: 获取企业用户角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `tenantCode` | string | 否 |  企业编码（可选，不填取当前用户企业） |
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
    "list": [
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
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/tenant/user/role/get-list \
  --body '{"tenantCode": "string", "userID": "string"}'
```
