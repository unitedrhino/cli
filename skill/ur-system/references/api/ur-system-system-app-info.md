# ur-system system/app/info

添加应用 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/app/info/create` | 添加应用 | platform |
| POST | `/api/v1/system/app/info/delete` | 删除应用 | platform |
| POST | `/api/v1/system/app/info/get-list` | 获取应用列表 | platform |
| POST | `/api/v1/system/app/info/get-one` | 获取应用详情 | platform |
| POST | `/api/v1/system/app/info/update` | 更新应用 | platform |

## 详细说明

### POST `/api/v1/system/app/info/create`

**说明**: 添加应用

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agreementsMap` | object | 否 |  按请求协议 code 回填的协议映射 |
| `appleConfig` | object | 否 |  |
| `appleConfig.appID` | string | 否 |  |
| `appleConfig.bundleID` | string | 否 |  |
| `appleConfig.keyID` | string | 否 |  |
| `appleConfig.privateKey` | string | 否 |  |
| `appleConfig.redirectURI` | string | 否 |  |
| `appleConfig.teamID` | string | 否 |  |
| `avatar` | string | 否 |  头像 更新的时候传filePath即可 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  应用描述 |
| `dingConfig` | object | 否 |  |
| `dingConfig.appID` | string | 否 |  |
| `dingConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `dingConfig.appSecret` | string | 否 |  |
| `githubConfig` | object | 否 |  |
| `githubConfig.appID` | string | 否 |  |
| `githubConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `githubConfig.appSecret` | string | 否 |  |
| `googleConfig` | object | 否 |  |
| `googleConfig.appID` | string | 否 |  |
| `googleConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `googleConfig.appSecret` | string | 否 |  |
| `id` | string | 否 |  编号,创建时自动生成(默认),也可以指定 |
| `isSysCreated` | integer | 否 |  是否系统创建 1-是 2-否（系统创建的应用不允许删除，绑定时自动添加） (格式: int64) |
| `isUseMenu` | integer | 否 |  是否使用菜单 1-是 2-否 (格式: int64) |
| `isUseProxy` | integer | 否 |  是否使用代理 1-是 2-否 (格式: int64) |
| `loginTypes` | array[string] | 否 | 选项取决于 app里是否配置了 支持的登录类型(不填支持全部登录方式):    "email":邮箱(系统配置里配置了才能开启) "phone":手机号(系统配置里配置了才能开启)  "wxMiniP":微信小程序(app里配置了微信小程序才能开启)  "wxOpen": 微信开放平台登录(app里配置了微信开放才能开启)   "ding":钉钉应用(包含小程序,h5等方式,app里配置了钉钉才能开启)  "pwd":账号密码注册 |
| `menuRole` | string | 否 |  菜单角色类型 all-支持两类(菜单需选role) platform-平台管理员 tenant-租户管理员 |
| `name` | string | 否 |  应用名称 |
| `nativeConfig` | object | 否 |  |
| `nativeConfig.filePath` | string | 否 |  文件路径,拿来下载文件 |
| `nativeConfig.version` | string | 否 |  应用版本 |
| `nativeConfig.versionDesc` | string | 否 | 版本说明 |
| `proxy` | string | 否 |  代理配置(JSON字符串) |
| `sort` | integer | 否 | 排序,不能小于等于0 (格式: int64) |
| `status` | integer | 否 |  应用状态 1:上架, 2:下架(默认) (格式: int64) |
| `subType` | string | 否 |  子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `tenantCode` | string | 否 |  租户编码,common的是通用应用,非common的是租户应用 |
| `tenantName` | string | 否 |  租户名称 |
| `trialTime` | integer | 否 | 试用时间(单位为天,为0不限制,如果为0,是免费使用的,如果为-1 则不允许自己创建) (格式: int64) |
| `type` | string | 否 |  应用类型 web:web页面  native:原生  mini:小程序 |
| `url` | string | 否 |  应用地址 |
| `useBy` | string | 否 |  使用对象 client:登录直接授权该应用 admin:管理员使用,需要邀请 |
| `wxMiniConfig` | object | 否 |  |
| `wxMiniConfig.appID` | string | 否 |  |
| `wxMiniConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `wxMiniConfig.appSecret` | string | 否 |  |
| `wxOpenConfig` | object | 否 |  |
| `wxOpenConfig.appID` | string | 否 |  |
| `wxOpenConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `wxOpenConfig.appSecret` | string | 否 |  |

**请求示例**:
```json
{
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
ur api /api/v1/system/app/info/create \
  --body '{"agreementsMap": {}, "appleConfig": {"appID": "string", "bundleID": "string", "keyID": "string", "privateKey": "string", "redirectURI": "string", "teamID": "string"}, "avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "dingConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "githubConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "googleConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "id": "string", "isSysCreated": 1, "isUseMenu": 1, "isUseProxy": 1, "loginTypes": ["string"], "menuRole": "string", "name": "示例名称", "nativeConfig": {"filePath": "string", "version": "string", "versionDesc": "string"}, "proxy": "string", "sort": 1, "status": 1, "subType": "string", "tenantCode": "string", "tenantName": "示例名称", "trialTime": 1, "type": "string", "url": "string", "useBy": "string", "wxMiniConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "wxOpenConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}}'
```

### POST `/api/v1/system/app/info/delete`

**说明**: 删除应用

**权限**: platform

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
ur api /api/v1/system/app/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/app/info/get-list`

**说明**: 获取应用列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[string] | 否 |  应用ID列表 |
| `isSysCreated` | integer | 否 |  是否系统创建过滤 1-是 2-否 (格式: int64) |
| `name` | string | 否 |  应用名称 |
| `notID` | string | 否 |  排除某个应用ID |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 应用状态 1 上架 2 下架 (格式: int64) |
| `subType` | string | 否 |  子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `tenantIsBind` | integer | 否 |  过滤租户是否已绑定的应用 1-是 2-否,如果要过滤指定租户,通过http头里指定 (格式: int64) |
| `type` | string | 否 |  应用类型 web:web页面  native:原生  mini:小程序 |
| `useBy` | string | 否 |  使用对象 client:登录直接授权该应用 admin:管理员使用,需要邀请 |
| `withAgreementCodes` | array[string] | 否 |  非空时按 code 回填已绑定协议 |

**请求示例**:
```json
{
  "ids": [
    "string"
  ],
  "isSysCreated": 1,
  "name": "示例名称",
  "notID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": 1,
  "subType": "string",
  "tenantIsBind": 1,
  "type": "string",
  "useBy": "string",
  "withAgreementCodes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
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
ur api /api/v1/system/app/info/get-list \
  --body '{"ids": ["string"], "isSysCreated": 1, "name": "示例名称", "notID": "string", "page": {"page": 1, "pageSize": 1}, "status": 1, "subType": "string", "tenantIsBind": 1, "type": "string", "useBy": "string", "withAgreementCodes": ["string"]}'
```

### POST `/api/v1/system/app/info/get-one`

**说明**: 获取应用详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  |
| `withAgreementCodes` | array[string] | 否 |  |

**请求示例**:
```json
{
  "id": "string",
  "withAgreementCodes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/app/info/get-one \
  --body '{"id": "string", "withAgreementCodes": ["string"]}'
```

### POST `/api/v1/system/app/info/update`

**说明**: 更新应用

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agreementsMap` | object | 否 |  按请求协议 code 回填的协议映射 |
| `appleConfig` | object | 否 |  |
| `appleConfig.appID` | string | 否 |  |
| `appleConfig.bundleID` | string | 否 |  |
| `appleConfig.keyID` | string | 否 |  |
| `appleConfig.privateKey` | string | 否 |  |
| `appleConfig.redirectURI` | string | 否 |  |
| `appleConfig.teamID` | string | 否 |  |
| `avatar` | string | 否 |  头像 更新的时候传filePath即可 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  应用描述 |
| `dingConfig` | object | 否 |  |
| `dingConfig.appID` | string | 否 |  |
| `dingConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `dingConfig.appSecret` | string | 否 |  |
| `githubConfig` | object | 否 |  |
| `githubConfig.appID` | string | 否 |  |
| `githubConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `githubConfig.appSecret` | string | 否 |  |
| `googleConfig` | object | 否 |  |
| `googleConfig.appID` | string | 否 |  |
| `googleConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `googleConfig.appSecret` | string | 否 |  |
| `id` | string | 否 |  编号,创建时自动生成(默认),也可以指定 |
| `isSysCreated` | integer | 否 |  是否系统创建 1-是 2-否（系统创建的应用不允许删除，绑定时自动添加） (格式: int64) |
| `isUseMenu` | integer | 否 |  是否使用菜单 1-是 2-否 (格式: int64) |
| `isUseProxy` | integer | 否 |  是否使用代理 1-是 2-否 (格式: int64) |
| `loginTypes` | array[string] | 否 | 选项取决于 app里是否配置了 支持的登录类型(不填支持全部登录方式):    "email":邮箱(系统配置里配置了才能开启) "phone":手机号(系统配置里配置了才能开启)  "wxMiniP":微信小程序(app里配置了微信小程序才能开启)  "wxOpen": 微信开放平台登录(app里配置了微信开放才能开启)   "ding":钉钉应用(包含小程序,h5等方式,app里配置了钉钉才能开启)  "pwd":账号密码注册 |
| `menuRole` | string | 否 |  菜单角色类型 all-支持两类(菜单需选role) platform-平台管理员 tenant-租户管理员 |
| `name` | string | 否 |  应用名称 |
| `nativeConfig` | object | 否 |  |
| `nativeConfig.filePath` | string | 否 |  文件路径,拿来下载文件 |
| `nativeConfig.version` | string | 否 |  应用版本 |
| `nativeConfig.versionDesc` | string | 否 | 版本说明 |
| `proxy` | string | 否 |  代理配置(JSON字符串) |
| `sort` | integer | 否 | 排序,不能小于等于0 (格式: int64) |
| `status` | integer | 否 |  应用状态 1:上架, 2:下架(默认) (格式: int64) |
| `subType` | string | 否 |  子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `tenantCode` | string | 否 |  租户编码,common的是通用应用,非common的是租户应用 |
| `tenantName` | string | 否 |  租户名称 |
| `trialTime` | integer | 否 | 试用时间(单位为天,为0不限制,如果为0,是免费使用的,如果为-1 则不允许自己创建) (格式: int64) |
| `type` | string | 否 |  应用类型 web:web页面  native:原生  mini:小程序 |
| `url` | string | 否 |  应用地址 |
| `useBy` | string | 否 |  使用对象 client:登录直接授权该应用 admin:管理员使用,需要邀请 |
| `wxMiniConfig` | object | 否 |  |
| `wxMiniConfig.appID` | string | 否 |  |
| `wxMiniConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `wxMiniConfig.appSecret` | string | 否 |  |
| `wxOpenConfig` | object | 否 |  |
| `wxOpenConfig.appID` | string | 否 |  |
| `wxOpenConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `wxOpenConfig.appSecret` | string | 否 |  |

**请求示例**:
```json
{
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
ur api /api/v1/system/app/info/update \
  --body '{"agreementsMap": {}, "appleConfig": {"appID": "string", "bundleID": "string", "keyID": "string", "privateKey": "string", "redirectURI": "string", "teamID": "string"}, "avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "dingConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "githubConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "googleConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "id": "string", "isSysCreated": 1, "isUseMenu": 1, "isUseProxy": 1, "loginTypes": ["string"], "menuRole": "string", "name": "示例名称", "nativeConfig": {"filePath": "string", "version": "string", "versionDesc": "string"}, "proxy": "string", "sort": 1, "status": 1, "subType": "string", "tenantCode": "string", "tenantName": "示例名称", "trialTime": 1, "type": "string", "url": "string", "useBy": "string", "wxMiniConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "wxOpenConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}}'
```
