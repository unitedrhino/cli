# ur-project things/data/area

删除授权区域权限 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/data/area/batch-delete` | 删除授权区域权限 | admin |
| POST | `/api/v1/things/data/area/batch-update` | 更新授权区域权限 | admin |
| POST | `/api/v1/things/data/area/get-list` | 获取区域权限列表 | admin |

## 详细说明

### POST `/api/v1/things/data/area/batch-delete`

**说明**: 删除授权区域权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaIDs` | array[string] | 是 |  |
| `projectID` | string | 否 | 项目id |
| `targetID` | string | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "areaIDs": [
    "string"
  ],
  "projectID": "string",
  "targetID": "string",
  "targetType": "string"
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
ur api /api/v1/things/data/area/batch-delete \
  --body '{"areaIDs": ["string"], "projectID": "string", "targetID": "string", "targetType": "string"}'
```

### POST `/api/v1/things/data/area/batch-update`

**说明**: 更新授权区域权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areas` | array[DataArea] | 是 | 权限数据IDs |
| `projectID` | string | 否 | 项目id |
| `targetID` | string | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "areas": [
    {
      "areaID": "string",
      "authType": 1,
      "isAuthChildren": 1
    }
  ],
  "projectID": "string",
  "targetID": "string",
  "targetType": "string"
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
ur api /api/v1/things/data/area/batch-update \
  --body '{"areas": [{"areaID": "string", "authType": 1, "isAuthChildren": 1}], "projectID": "string", "targetID": "string", "targetType": "string"}'
```

### POST `/api/v1/things/data/area/get-list`

**说明**: 获取区域权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `projectID` | string | 否 | 项目id |
| `targetID` | string | 否 | 用户ID |
| `targetType` | string | 否 |  |

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
  },
  "projectID": "string",
  "targetID": "string",
  "targetType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "": {
          "areaID": "string",
          "areaIDPath": "string",
          "areaImg": "string",
          "areaName": "示例名称",
          "areaNamePath": "string",
          "children": [
            {
              "areaID": "...",
              "areaIDPath": "...",
              "areaImg": "...",
              "areaName": "...",
              "areaNamePath": "...",
              "children": "...",
              "configFile": "...",
              "createdTime": "...",
              "desc": "...",
              "deviceCount": "...",
              "devices": "...",
              "groupCount": "...",
              "isLeaf": "...",
              "isSysCreated": "...",
              "lowerLevelCount": "...",
              "parentAreaID": "...",
              "position": "...",
              "projectID": "...",
              "sort": "...",
              "tags": "...",
              "tenantCode": "...",
              "useBy": "..."
            }
          ],
          "configFile": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [
            {
              "adcode": "...",
              "address": "...",
              "area": "...",
              "areaID": "...",
              "categoryID": "...",
              "cert": "...",
              "cloneID": "...",
              "createdTime": "...",
              "desc": "...",
              "deviceAlias": "...",
              "deviceImg": "...",
              "deviceName": "...",
              "deviceType": "...",
              "distributor": "...",
              "expTime": "...",
              "file": "...",
              "firstBind": "...",
              "firstLogin": "...",
              "gateway": "...",
              "groupIDs": "...",
              "groupPurpose": "...",
              "groups": "...",
              "hardInfo": "...",
              "iccid": "...",
              "id": "...",
              "imei": "...",
              "isEnable": "...",
              "isOnline": "...",
              "lastBind": "...",
              "lastIp": "...",
              "lastLocalIp": "...",
              "lastLogin": "...",
              "lastOffline": "...",
              "logLevel": "...",
              "mac": "...",
              "mobileOperator": "...",
              "needConfirmVersion": "...",
              "netType": "...",
              "owner": "...",
              "phone": "...",
              "position": "...",
              "productID": "...",
              "productImg": "...",
              "productName": "...",
              "profiles": "...",
              "projectID": "...",
              "protocolConf": "...",
              "ratedPower": "...",
              "remoteAddr": "...",
              "rssi": "...",
              "schemaAlias": "...",
              "secret": "...",
              "softInfo": "...",
              "sort": "...",
              "status": "...",
              "subProtocolConf": "...",
              "tags": "...",
              "tenantCode": "...",
              "userID": "...",
              "version": "...",
              "withProperties": "..."
            }
          ],
          "groupCount": 1,
          "isLeaf": 1,
          "isSysCreated": 1,
          "lowerLevelCount": 1,
          "parentAreaID": "string",
          "position": {
            "latitude": 1,
            "longitude": 1
          },
          "projectID": "string",
          "sort": 1,
          "tags": {},
          "tenantCode": "string",
          "useBy": "string"
        },
        "authType": 1,
        "isAuthChildren": 1,
        "targetID": "string",
        "targetType": "string",
        "updatedTime": 1,
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
ur api /api/v1/things/data/area/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "projectID": "string", "targetID": "string", "targetType": "string"}'
```
