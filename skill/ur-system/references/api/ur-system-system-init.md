# ur-system system/init

初始化系统

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/common/system/init` | 初始化系统 | public |

## 详细说明

### POST `/api/v1/system/common/system/init`

**说明**: 初始化系统

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `config` | object | 是 |  |
| `config.email` | object | 否 |  |
| `config.email.from` | string | 否 |  发件人  你自己要发邮件的邮箱 |
| `config.email.host` | string | 否 |  服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议 |
| `config.email.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `config.email.isSsl` | integer | 否 |  是否SSL   是否开启SSL (格式: int64) |
| `config.email.nickname` | string | 否 |  昵称    发件人昵称 通常为自己的邮箱 |
| `config.email.port` | integer | 否 |  端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465 (格式: int64) |
| `config.email.secret` | string | 否 |  密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥 |
| `config.log` | object | 否 |  |
| `config.log.loginKeepDays` | integer | 是 |  登录日志保留天数 (格式: int64) |
| `config.log.operKeepDays` | integer | 是 |  操作日志保留天数 (格式: int64) |
| `config.map` | object | 否 |  |
| `config.map.accessKey` | string | 是 |  访问密钥 |
| `config.map.accessSecret` | string | 是 |  访问密钥 |
| `config.map.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `config.map.type` | string | 是 |  地图服务商模式 gaode: 高德 可选: ['gaode'] |
| `config.oem` | object | 否 |  |
| `config.oem.avatar` | string | 是 |  Logo图片地址 |
| `config.oem.footer` | object | 是 |  |
| `config.oem.loginFormDescription` | string | 否 |  登录表单描述 |
| `config.oem.loginFormTitle` | string | 否 |  登录表单标题 |
| `config.oem.loginPageDescription` | string | 否 |  登录页左侧描述 |
| `config.oem.loginPageTitle` | string | 否 |  登录页左侧标题 |
| `config.oem.sloganImage` | string | 否 |  上传的Slogan图片 |
| `config.oem.sloganImageTemplate` | string | 否 |  Slogan模板ID |
| `config.oem.sloganImageType` | string | 否 |  Slogan图片类型 可选: ['template', 'upload'] |
| `config.oem.title` | string | 是 |  系统标题 |
| `config.sms` | object | 否 |  |
| `config.sms.accessKeyID` | string | 否 |  AccessKey ID |
| `config.sms.accessKeySecret` | string | 否 |  AccessKey Secret |
| `config.sms.appID` | string | 否 |  应用ID 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666,可前往 [短信控制台](https:console.cloud.tencent.com/smsv2/app-manage) 查看 |
| `config.sms.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `config.sms.signName` | string | 否 |  签名 |
| `config.sms.type` | string | 否 |  短信类型, ali: 阿里云 tencent: 腾讯 可选: ['ali', 'tencent'] |
| `config.tenant` | object | 否 |  |
| `config.tenant.defaultLimit` | object | 是 |  |
| `config.user` | object | 否 |  |
| `config.user.captchaLen` | integer | 是 |  验证码长度 (格式: int32) |
| `config.user.forceSetPassword` | integer | 是 |  是否强制未设置密码的用户登录后设置密码(1:是，2:否) (格式: int64) 可选: ['1', '2'] |
| `config.user.ownerUserID` | string | 否 |  平台超级管理员ID 非平台超管不返回该参数,不可修改 |
| `config.user.passLevel` | integer | 是 |  用户密码强度级别 1:长度大于等于8位即可 2:包含数字,大写字母,小写字母,特殊字符起码两种  3:包含数字,大写字母,小写字母,特殊字符起码3种 4:包含数字,大写字母,小写字母,特殊字符   特殊字符: !@#~$%^&*()+\|_ (格式: int64) |
| `config.user.remindSetPassword` | integer | 是 |  是否提醒未设置密码的用户设置密码(1:是，2:否) (格式: int64) 可选: ['1', '2'] |
| `config.user.userOwnTenantLimit` | integer | 是 |  用户可直接创建免费租户数量，0表示不可直接创建免费租户，需提供授权码 (格式: int64) |
| `config.weather` | object | 否 |  |
| `config.weather.apiHost` | string | 是 |  天气API主机地址 |
| `config.weather.apiKey` | string | 是 |  天气API密钥 |
| `config.weather.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `ownerUser` | object | 是 |  |
| `ownerUser.avatar` | string | 否 |  可选头像，创建的时候传filePath |
| `ownerUser.email` | string | 否 |  可选邮箱 |
| `ownerUser.nickName` | string | 否 |  可选昵称 |
| `ownerUser.password` | string | 是 |  密码原文 |
| `ownerUser.phone` | string | 否 |  可选手机号 |
| `ownerUser.userName` | string | 是 |  用户账号 |

**请求示例**:
```json
{
  "config": {
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
  "ownerUser": {
    "avatar": "string",
    "email": "string",
    "nickName": "string",
    "password": "******",
    "phone": "string",
    "userName": "string"
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/system/init \
  --body '{"config": {"email": {"from": "string", "host": "string", "isEnable": 1, "isSsl": 1, "nickname": "string", "port": 1, "secret": "******"}, "log": {"loginKeepDays": 1, "operKeepDays": 1}, "map": {"accessKey": "string", "accessSecret": "string", "isEnable": 1, "type": "gaode"}, "oem": {"avatar": "string", "footer": {"beian": "string", "beianLink": "string", "companyName": "示例名称", "companySiteLink": "string", "date": "string", "icp": "string", "icpLink": "string", "isEnable": 1}, "loginFormDescription": "string", "loginFormTitle": "string", "loginPageDescription": "string", "loginPageTitle": "string", "sloganImage": "string", "sloganImageTemplate": "string", "sloganImageType": "template", "title": "string"}, "sms": {"accessKeyID": "string", "accessKeySecret": "string", "appID": "string", "isEnable": 1, "signName": "string", "type": "ali"}, "tenant": {"defaultLimit": {"aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}}, "user": {"captchaLen": 1, "forceSetPassword": "1", "ownerUserID": "string", "passLevel": 1, "remindSetPassword": "1", "userOwnTenantLimit": 1}, "weather": {"apiHost": "string", "apiKey": "string", "isEnable": 1}}, "ownerUser": {"avatar": "string", "email": "string", "nickName": "string", "password": "******", "phone": "string", "userName": "string"}}'
```
