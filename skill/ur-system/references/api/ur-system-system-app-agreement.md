# ur-system system/app/agreement

批量更新应用协议绑定 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/app/agreement/bind-batch-update` | 批量更新应用协议绑定 | platform |
| POST | `/api/v1/system/app/agreement/get-bind-list` | 获取应用协议绑定列表 | platform |

## 详细说明

### POST `/api/v1/system/app/agreement/bind-batch-update`

**说明**: 批量更新应用协议绑定

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agreementID` | string | 否 |  |
| `appID` | string | 否 |  |
| `list` | array[AppAgreementBind] | 是 |  |

**请求示例**:
```json
{
  "agreementID": "string",
  "appID": "string",
  "list": [
    {
      "agreement": {
        "code": "string",
        "content": "string",
        "createdTime": "string",
        "id": "string",
        "name": "string",
        "remark": "string",
        "status": 1,
        "title": "string",
        "updatedTime": "string"
      },
      "agreementID": "string",
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
      "id": "string",
      "sort": 1,
      "status": 1,
      "type": "string"
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
ur api /api/v1/system/app/agreement/bind-batch-update \
  --body '{"agreementID": "string", "appID": "string", "list": [{"agreement": {"code": "string", "content": "string", "createdTime": "string", "id": "string", "name": "string", "remark": "string", "status": 1, "title": "string", "updatedTime": "string"}, "agreementID": "string", "app": {"agreementsMap": {}, "appleConfig": {"appID": "string", "bundleID": "string", "keyID": "string", "privateKey": "string", "redirectURI": "string", "teamID": "string"}, "avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "dingConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "githubConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "googleConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "id": "string", "isSysCreated": 1, "isUseMenu": 1, "isUseProxy": 1, "loginTypes": ["string"], "menuRole": "string", "name": "示例名称", "nativeConfig": {"filePath": "string", "version": "string", "versionDesc": "string"}, "proxy": "string", "sort": 1, "status": 1, "subType": "string", "tenantCode": "string", "tenantName": "示例名称", "trialTime": 1, "type": "string", "url": "string", "useBy": "string", "wxMiniConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "wxOpenConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}}, "appID": "string", "id": "string", "sort": 1, "status": 1, "type": "string"}]}'
```

### POST `/api/v1/system/app/agreement/get-bind-list`

**说明**: 获取应用协议绑定列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agreementID` | string | 否 |  |
| `appID` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int64 |
| `type` | string | 否 |  |
| `withAgreement` | boolean | 否 | 格式: boolean |
| `withApp` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "agreementID": "string",
  "appID": "string",
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
  "status": 1,
  "type": "string",
  "withAgreement": true,
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
        "agreement": {
          "code": "string",
          "content": "string",
          "createdTime": "string",
          "id": "string",
          "name": "string",
          "remark": "string",
          "status": 1,
          "title": "string",
          "updatedTime": "string"
        },
        "agreementID": "string",
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
        "id": "string",
        "sort": 1,
        "status": 1,
        "type": "string"
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
ur api /api/v1/system/app/agreement/get-bind-list \
  --body '{"agreementID": "string", "appID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "type": "string", "withAgreement": true, "withApp": true}'
```
