# ur-system system/common

批量聚合接口请求 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/common/api/batch-agg` | 批量聚合接口请求 | all |
| GET | `/api/v1/system/common/debug` | 调试接口GET | public |
| POST | `/api/v1/system/common/debug` | 调试接口POST | public |
| GET | `/api/v1/system/common/debug-tencent` | 腾讯云调试接口 | public |
| GET | `/api/v1/system/common/download-file` | 下载本地文件 | public |
| POST | `/api/v1/system/common/init-upload-file` | 初始化上传文件 | public |
| POST | `/api/v1/system/common/ntp/get-one` | ntp时间同步 | public |
| POST | `/api/v1/system/common/qr-code/get-one` | 获取小程序二维码 | all |
| POST | `/api/v1/system/common/sys-config/info/get-one` | 读取系统配置信息 | platform |
| POST | `/api/v1/system/common/sys-config/info/update` | 更新系统配置信息 | platform |
| POST | `/api/v1/system/common/third/dept/get-list` | 获取第三方部门列表 | all |
| POST | `/api/v1/system/common/third/dept/get-one` | 获取第三方部门详情 | all |
| POST | `/api/v1/system/common/upload-file` | 文件直传 | all |
| POST | `/api/v1/system/common/upload-url/create` | 获取文件上传地址 | all |
| POST | `/api/v1/system/common/weather/get-one` | 获取天气情况 | all |
| GET | `/api/v1/system/common/websocket/connect` | websocket连接 | all |

## 详细说明

### POST `/api/v1/system/common/api/batch-agg`

**说明**: 批量聚合接口请求

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `reqs` | array[ApiBatchAggReqOne] | 是 |  |

**请求示例**:
```json
{
  "reqs": [
    {
      "body": [
        {}
      ],
      "uri": "string"
    }
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "lists": [
      [
        null
      ]
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/api/batch-agg \
  --body '{"reqs": [{"body": [{}], "uri": "string"}]}'
```

### GET `/api/v1/system/common/debug`

**说明**: 调试接口GET

**权限**: public

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "body": "string",
    "headers": {},
    "requestUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/debug \
  --body '{}'
```

### POST `/api/v1/system/common/debug`

**说明**: 调试接口POST

**权限**: public

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "body": "string",
    "headers": {},
    "requestUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/debug \
  --body '{}'
```

### GET `/api/v1/system/common/debug-tencent`

**说明**: 腾讯云调试接口

**权限**: public

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/debug-tencent \
  --body '{}'
```

### GET `/api/v1/system/common/download-file`

**说明**: 下载本地文件

**权限**: public

**路径参数**:

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| `filePath` | query | string | 是 | 文件的路径 |
| `sign` | query | string | 否 | 签名.如果是私有的需要填此参数 |

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/download-file \
  --body '{}'
```

### POST `/api/v1/system/common/init-upload-file`

**说明**: 初始化上传文件

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `file` | string | 是 | 格式: binary |

**请求示例**:
```json
{
  "file": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "filePath": "string",
    "fileUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/init-upload-file \
  --body '{"file": "string"}'
```

### POST `/api/v1/system/common/ntp/get-one`

**说明**: ntp时间同步

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceSendTime` | string | 是 |  |

**请求示例**:
```json
{
  "deviceSendTime": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "deviceSendTime": "string",
    "serverRecvTime": "string",
    "serverSendTime": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/ntp/get-one \
  --body '{"deviceSendTime": "string"}'
```

### POST `/api/v1/system/common/qr-code/get-one`

**说明**: 获取小程序二维码

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 | 应用ID |
| `envVersion` | string | 是 | 要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版。 |
| `page` | string | 是 | 默认是主页，页面 page，例如 pages/index/index，根路径前不要填加 /，不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面。scancode_time为系统保留参数，不允许配置 |
| `scene` | string | 是 |         	最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式） |

**请求示例**:
```json
{
  "appID": "string",
  "envVersion": "string",
  "page": "string",
  "scene": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "buffer": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/qr-code/get-one \
  --body '{"appID": "string", "envVersion": "string", "page": "string", "scene": "string"}'
```

### POST `/api/v1/system/common/sys-config/info/get-one`

**说明**: 读取系统配置信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "email": {
      "from": "string",
      "host": "string",
      "isEnable": 1,
      "isSsl": 1,
      "nickname": "string",
      "port": 1,
      "secret": "******"
    },
    "log": {
      "loginKeepDays": 1,
      "operKeepDays": 1
    },
    "map": {
      "accessKey": "string",
      "accessSecret": "string",
      "isEnable": 1,
      "type": "gaode"
    },
    "oem": {
      "avatar": "string",
      "footer": {
        "beian": "string",
        "beianLink": "string",
        "companyName": "示例名称",
        "companySiteLink": "string",
        "date": "string",
        "icp": "string",
        "icpLink": "string",
        "isEnable": 1
      },
      "loginFormDescription": "string",
      "loginFormTitle": "string",
      "loginPageDescription": "string",
      "loginPageTitle": "string",
      "sloganImage": "string",
      "sloganImageTemplate": "string",
      "sloganImageType": "template",
      "title": "string"
    },
    "sms": {
      "accessKeyID": "string",
      "accessKeySecret": "string",
      "appID": "string",
      "isEnable": 1,
      "signName": "string",
      "type": "ali"
    },
    "tenant": {
      "defaultLimit": {
        "aiTokenMonthly": 1,
        "deviceMsgPerDay": 1,
        "deviceNum": 1,
        "diskSpaceGB": 1,
        "userNum": 1
      }
    },
    "user": {
      "captchaLen": 1,
      "forceSetPassword": "1",
      "ownerUserID": "string",
      "passLevel": 1,
      "remindSetPassword": "1",
      "userOwnTenantLimit": 1
    },
    "weather": {
      "apiHost": "string",
      "apiKey": "string",
      "isEnable": 1
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/sys-config/info/get-one \
  --body '{}'
```

### POST `/api/v1/system/common/sys-config/info/update`

**说明**: 更新系统配置信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `email` | object | 否 |  |
| `email.from` | string | 否 |  发件人  你自己要发邮件的邮箱 |
| `email.host` | string | 否 |  服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议 |
| `email.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `email.isSsl` | integer | 否 |  是否SSL   是否开启SSL (格式: int64) |
| `email.nickname` | string | 否 |  昵称    发件人昵称 通常为自己的邮箱 |
| `email.port` | integer | 否 |  端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465 (格式: int64) |
| `email.secret` | string | 否 |  密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥 |
| `log` | object | 否 |  |
| `log.loginKeepDays` | integer | 是 |  登录日志保留天数 (格式: int64) |
| `log.operKeepDays` | integer | 是 |  操作日志保留天数 (格式: int64) |
| `map` | object | 否 |  |
| `map.accessKey` | string | 是 |  访问密钥 |
| `map.accessSecret` | string | 是 |  访问密钥 |
| `map.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `map.type` | string | 是 |  地图服务商模式 gaode: 高德 可选: ['gaode'] |
| `oem` | object | 否 |  |
| `oem.avatar` | string | 是 |  Logo图片地址 |
| `oem.footer` | object | 是 |  |
| `oem.footer.beian` | string | 否 |  公安备案号 |
| `oem.footer.beianLink` | string | 否 |  公安备案的跳转链接 |
| `oem.footer.companyName` | string | 是 |  公司名称 |
| `oem.footer.companySiteLink` | string | 否 |  公司网站链接 |
| `oem.footer.date` | string | 否 |  年份或日期 |
| `oem.footer.icp` | string | 否 |  ICP备案号 |
| `oem.footer.icpLink` | string | 否 |  ICP备案的跳转链接 |
| `oem.footer.isEnable` | integer | 是 |  是否启用(1:启用，2:禁用) (格式: int64) |
| `oem.loginFormDescription` | string | 否 |  登录表单描述 |
| `oem.loginFormTitle` | string | 否 |  登录表单标题 |
| `oem.loginPageDescription` | string | 否 |  登录页左侧描述 |
| `oem.loginPageTitle` | string | 否 |  登录页左侧标题 |
| `oem.sloganImage` | string | 否 |  上传的Slogan图片 |
| `oem.sloganImageTemplate` | string | 否 |  Slogan模板ID |
| `oem.sloganImageType` | string | 否 |  Slogan图片类型 可选: ['template', 'upload'] |
| `oem.title` | string | 是 |  系统标题 |
| `sms` | object | 否 |  |
| `sms.accessKeyID` | string | 否 |  AccessKey ID |
| `sms.accessKeySecret` | string | 否 |  AccessKey Secret |
| `sms.appID` | string | 否 |  应用ID 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666,可前往 [短信控制台](https:console.cloud.tencent.com/smsv2/app-manage) 查看 |
| `sms.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `sms.signName` | string | 否 |  签名 |
| `sms.type` | string | 否 |  短信类型, ali: 阿里云 tencent: 腾讯 可选: ['ali', 'tencent'] |
| `tenant` | object | 否 |  |
| `tenant.defaultLimit` | object | 是 |  |
| `tenant.defaultLimit.aiTokenMonthly` | integer | 否 |  默认 AI Token 月配额（单位：token） (格式: int64) |
| `tenant.defaultLimit.deviceMsgPerDay` | integer | 否 |  默认企业每天设备消息总量上限 (格式: int64) |
| `tenant.defaultLimit.deviceNum` | integer | 否 |  企业下的设备数量限制,0为不限制 (格式: int64) |
| `tenant.defaultLimit.diskSpaceGB` | integer | 否 |  默认磁盘空间 GB (格式: int64) |
| `tenant.defaultLimit.userNum` | integer | 否 |  企业下的用户数量限制,0为不限制 (格式: int64) |
| `user` | object | 否 |  |
| `user.captchaLen` | integer | 是 |  验证码长度 (格式: int32) |
| `user.forceSetPassword` | integer | 是 |  是否强制未设置密码的用户登录后设置密码(1:是，2:否) (格式: int64) 可选: ['1', '2'] |
| `user.ownerUserID` | string | 否 |  平台超级管理员ID 非平台超管不返回该参数,不可修改 |
| `user.passLevel` | integer | 是 |  用户密码强度级别 1:长度大于等于8位即可 2:包含数字,大写字母,小写字母,特殊字符起码两种  3:包含数字,大写字母,小写字母,特殊字符起码3种 4:包含数字,大写字母,小写字母,特殊字符   特殊字符: !@#~$%^&*()+\|_ (格式: int64) |
| `user.remindSetPassword` | integer | 是 |  是否提醒未设置密码的用户设置密码(1:是，2:否) (格式: int64) 可选: ['1', '2'] |
| `user.userOwnTenantLimit` | integer | 是 |  用户可直接创建免费企业数量，0表示不可直接创建免费企业，需提供授权码 (格式: int64) |
| `weather` | object | 否 |  |
| `weather.apiHost` | string | 是 |  天气API主机地址 |
| `weather.apiKey` | string | 是 |  天气API密钥 |
| `weather.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |

**请求示例**:
```json
{
  "email": {
    "from": "string",
    "host": "string",
    "isEnable": 1,
    "isSsl": 1,
    "nickname": "string",
    "port": 1,
    "secret": "******"
  },
  "log": {
    "loginKeepDays": 1,
    "operKeepDays": 1
  },
  "map": {
    "accessKey": "string",
    "accessSecret": "string",
    "isEnable": 1,
    "type": "gaode"
  },
  "oem": {
    "avatar": "string",
    "footer": {
      "beian": "string",
      "beianLink": "string",
      "companyName": "示例名称",
      "companySiteLink": "string",
      "date": "string",
      "icp": "string",
      "icpLink": "string",
      "isEnable": 1
    },
    "loginFormDescription": "string",
    "loginFormTitle": "string",
    "loginPageDescription": "string",
    "loginPageTitle": "string",
    "sloganImage": "string",
    "sloganImageTemplate": "string",
    "sloganImageType": "template",
    "title": "string"
  },
  "sms": {
    "accessKeyID": "string",
    "accessKeySecret": "string",
    "appID": "string",
    "isEnable": 1,
    "signName": "string",
    "type": "ali"
  },
  "tenant": {
    "defaultLimit": {
      "aiTokenMonthly": 1,
      "deviceMsgPerDay": 1,
      "deviceNum": 1,
      "diskSpaceGB": 1,
      "userNum": 1
    }
  },
  "user": {
    "captchaLen": 1,
    "forceSetPassword": "1",
    "ownerUserID": "string",
    "passLevel": 1,
    "remindSetPassword": "1",
    "userOwnTenantLimit": 1
  },
  "weather": {
    "apiHost": "string",
    "apiKey": "string",
    "isEnable": 1
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
ur api /api/v1/system/common/sys-config/info/update \
  --body '{"email": {"from": "string", "host": "string", "isEnable": 1, "isSsl": 1, "nickname": "string", "port": 1, "secret": "******"}, "log": {"loginKeepDays": 1, "operKeepDays": 1}, "map": {"accessKey": "string", "accessSecret": "string", "isEnable": 1, "type": "gaode"}, "oem": {"avatar": "string", "footer": {"beian": "string", "beianLink": "string", "companyName": "示例名称", "companySiteLink": "string", "date": "string", "icp": "string", "icpLink": "string", "isEnable": 1}, "loginFormDescription": "string", "loginFormTitle": "string", "loginPageDescription": "string", "loginPageTitle": "string", "sloganImage": "string", "sloganImageTemplate": "string", "sloganImageType": "template", "title": "string"}, "sms": {"accessKeyID": "string", "accessKeySecret": "string", "appID": "string", "isEnable": 1, "signName": "string", "type": "ali"}, "tenant": {"defaultLimit": {"aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}}, "user": {"captchaLen": 1, "forceSetPassword": "1", "ownerUserID": "string", "passLevel": 1, "remindSetPassword": "1", "userOwnTenantLimit": 1}, "weather": {"apiHost": "string", "apiKey": "string", "isEnable": 1}}'
```

### POST `/api/v1/system/common/third/dept/get-list`

**说明**: 获取第三方部门列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 | 父节点 |
| `thirdConfig` | object | 是 |  |
| `thirdConfig.appID` | string | 否 |  |
| `thirdConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `thirdConfig.appSecret` | string | 否 |  |
| `thirdType` | string | 否 | 第三方的类型 ding:钉钉  wxE:企业微信 |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "parentID": "string",
  "thirdConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "thirdType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "children": [
          {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          }
        ],
        "desc": "string",
        "dingTalkID": "string",
        "id": "string",
        "idPath": "string",
        "name": "示例名称",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            }
          ],
          "desc": "string",
          "dingTalkID": "string",
          "id": "string",
          "idPath": "string",
          "name": "示例名称",
          "parent": {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "userCount": 1
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "userCount": 1
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
ur api /api/v1/system/common/third/dept/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}, "parentID": "string", "thirdConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "thirdType": "string"}'
```

### POST `/api/v1/system/common/third/dept/get-one`

**说明**: 获取第三方部门详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  编号 |
| `thirdConfig` | object | 是 |  |
| `thirdConfig.appID` | string | 否 |  |
| `thirdConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `thirdConfig.appSecret` | string | 否 |  |
| `thirdType` | string | 否 | 第三方的类型 ding:钉钉  wxE:企业微信 |
| `withChildren` | boolean | 否 | 是否返回子级 (格式: boolean) |
| `withFather` | boolean | 否 | 是否返回父级 (格式: boolean) |

**请求示例**:
```json
{
  "id": "string",
  "thirdConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "thirdType": "string",
  "withChildren": true,
  "withFather": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "children": [
      {
        "children": [
          {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          }
        ],
        "desc": "string",
        "dingTalkID": "string",
        "id": "string",
        "idPath": "string",
        "name": "示例名称",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            }
          ],
          "desc": "string",
          "dingTalkID": "string",
          "id": "string",
          "idPath": "string",
          "name": "示例名称",
          "parent": {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "userCount": 1
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "userCount": 1
      }
    ],
    "desc": "string",
    "dingTalkID": "string",
    "id": "string",
    "idPath": "string",
    "name": "示例名称",
    "parent": {
      "children": [
        {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            }
          ],
          "desc": "string",
          "dingTalkID": "string",
          "id": "string",
          "idPath": "string",
          "name": "示例名称",
          "parent": {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "userCount": 1
        }
      ],
      "desc": "string",
      "dingTalkID": "string",
      "id": "string",
      "idPath": "string",
      "name": "示例名称",
      "parent": {
        "children": [
          {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          }
        ],
        "desc": "string",
        "dingTalkID": "string",
        "id": "string",
        "idPath": "string",
        "name": "示例名称",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            }
          ],
          "desc": "string",
          "dingTalkID": "string",
          "id": "string",
          "idPath": "string",
          "name": "示例名称",
          "parent": {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "userCount": 1
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "userCount": 1
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "userCount": 1
    },
    "parentID": "string",
    "sort": 1,
    "status": 1,
    "userCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/third/dept/get-one \
  --body '{"id": "string", "thirdConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "thirdType": "string", "withChildren": true, "withFather": true}'
```

### POST `/api/v1/system/common/upload-file`

**说明**: 文件直传

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `file` | string | 是 | 格式: binary |

**请求示例**:
```json
{
  "file": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "filePath": "string",
    "fileUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/upload-file \
  --body '{"file": "string"}'
```

### POST `/api/v1/system/common/upload-url/create`

**说明**: 获取文件上传地址

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `business` | string | 是 | 业务(如产品管理 productManage) |
| `filePath` | string | 是 | 文件路径(带文件名), 当前不支持.svg |
| `isPublic` | boolean | 否 | true时使用公开桶上传,返回永久fileUrl。用于通知富文本中的图片、封面图等需长期可访问的资源 (格式: boolean) |
| `rename` | boolean | 否 | true 文件重命名，false 不重命名(默认) (格式: boolean) |
| `scene` | string | 是 | 场景(业务定义 如产品图片 productImg) |
| `useBy` | string | 否 | 公开桶时必填。上传用途: user=用户级, tenant=企业级, platform=平台级(需supper权限) |

**请求示例**:
```json
{
  "business": "string",
  "filePath": "string",
  "isPublic": true,
  "rename": true,
  "scene": "string",
  "useBy": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "filePath": "string",
    "fileUrl": "string",
    "uploadUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/upload-url/create \
  --body '{"business": "string", "filePath": "string", "isPublic": true, "rename": true, "scene": "string", "useBy": "string"}'
```

### POST `/api/v1/system/common/weather/get-one`

**说明**: 获取天气情况

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `projectID` | string | 否 |  |

**请求示例**:
```json
{
  "position": {
    "latitude": 1,
    "longitude": 1
  },
  "projectID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "air": {
      "aqi": "string",
      "category": "string",
      "co": "string",
      "level": "string",
      "no2": "string",
      "o3": "string",
      "pm10": "string",
      "pm2p5": "string",
      "primary": "string",
      "so2": "string"
    },
    "cloud": "string",
    "dew": "string",
    "feelsLike": "string",
    "humidity": "string",
    "icon": "string",
    "obsTime": "string",
    "precip": "string",
    "pressure": "string",
    "temp": "string",
    "text": "string",
    "vis": "string",
    "wind360": "string",
    "windDir": "string",
    "windScale": "string",
    "windSpeed": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/weather/get-one \
  --body '{"position": {"latitude": 1, "longitude": 1}, "projectID": "string"}'
```

### GET `/api/v1/system/common/websocket/connect`

**说明**: websocket连接

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
ur api /api/v1/system/common/websocket/connect \
  --body '{}'
```
