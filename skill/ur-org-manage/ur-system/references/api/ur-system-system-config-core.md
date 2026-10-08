# ur-system system/config/core

读取系统配置信息(无需登录)

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/common/sys-config/core/get-one` | 读取系统配置信息(无需登录) | public |

## 详细说明

### POST `/api/v1/system/common/sys-config/core/get-one`

**说明**: 读取系统配置信息(无需登录)

**权限**: public

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
ur api /api/v1/system/common/sys-config/core/get-one \
  --body '{}'
```
