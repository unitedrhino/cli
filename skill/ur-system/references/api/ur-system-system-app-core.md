# ur-system system/app/core

无需登录获取应用信息

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/app/core/get-one` | 无需登录获取应用信息 | public |

## 详细说明

### POST `/api/v1/system/app/core/get-one`

**说明**: 无需登录获取应用信息

**权限**: public

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
ur api /api/v1/system/app/core/get-one \
  --body '{"id": "string", "withAgreementCodes": ["string"]}'
```
