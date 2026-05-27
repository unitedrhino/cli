# ur-tenant system/tenant/app

绑定租户应用 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/app/create` | 绑定租户应用 | all |
| POST | `/api/v1/system/tenant/app/delete` | 解绑租户应用 | platform |
| POST | `/api/v1/system/tenant/app/get-list` | 获取租户应用列表 | admin |
| POST | `/api/v1/system/tenant/app/get-one` | 获取应用信息 | platform |
| POST | `/api/v1/system/tenant/app/update` | 更新租户应用 | platform |

## 详细说明

### POST `/api/v1/system/tenant/app/create`

**说明**: 绑定租户应用

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  应用ID |
| `licenseCode` | string | 否 |  授权码（付费应用必填，免费/试用应用可留空） |
| `tenantCode` | string | 否 |  指定租户,只有平台可以指定 |

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

**说明**: 解绑租户应用

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  |
| `tenantCode` | string | 否 | 指定租户,只有平台可以指定 |

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

**说明**: 获取租户应用列表

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
| `tenantCode` | string | 否 | 指定租户,租户下只能过滤 common或他自己的  common的是通用应用,非common的是租户应用 |
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
| `tenantCode` | string | 否 | 指定租户,只有平台可以指定 |

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

**说明**: 更新租户应用

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
| `app.menuRole` | string | 否 |  菜单角色类型 all-支持两类(菜单需选role) platform-平台管理员 tenant-租户管理员 |
| `app.name` | string | 否 |  应用名称 |
| `app.nativeConfig` | object | 否 |  |
| `app.nativeConfig.filePath` | string | 否 |  文件路径,拿来下载文件 |
| `app.nativeConfig.version` | string | 否 |  应用版本 |
| `app.nativeConfig.versionDesc` | string | 否 | 版本说明 |
| `app.proxy` | string | 否 |  代理配置(JSON字符串) |
| `app.sort` | integer | 否 | 排序,不能小于等于0 (格式: int64) |
| `app.status` | integer | 否 |  应用状态 1:上架, 2:下架(默认) (格式: int64) |
| `app.subType` | string | 否 |  子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `app.tenantCode` | string | 否 |  租户编码,common的是通用应用,非common的是租户应用 |
| `app.tenantName` | string | 否 |  租户名称 |
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
| `tenantCode` | string | 否 | 指定租户,只有平台可以指定 |

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
