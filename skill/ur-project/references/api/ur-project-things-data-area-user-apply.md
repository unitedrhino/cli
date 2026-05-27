# ur-project things/data/area/user/apply

授权区域权限 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/data/area/user/apply/deal` | 授权区域权限 | admin |
| POST | `/api/v1/things/data/area/user/apply/get-list` | 获取区域权限列表 | admin |

## 详细说明

### POST `/api/v1/things/data/area/user/apply/deal`

**说明**: 授权区域权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[string] | 是 | 权限数据IDs |
| `isApprove` | boolean | 是 | 是否同意 (格式: boolean) |

**请求示例**:
```json
{
  "ids": [
    "string"
  ],
  "isApprove": true
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
ur api /api/v1/things/data/area/user/apply/deal \
  --body '{"ids": ["string"], "isApprove": true}'
```

### POST `/api/v1/things/data/area/user/apply/get-list`

**说明**: 获取区域权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目id |
| `authTypes` | array[integer] | 否 | 权限类型 1:读权限,只能读,不能写 4:管理权限,可以修改别人的权限 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `withAreaInfo` | boolean | 否 | 把区域信息附带上 (格式: boolean) |
| `withUserInfo` | boolean | 否 | 把用户信息带上 (格式: boolean) |

**请求示例**:
```json
{
  "areaID": "string",
  "authTypes": [
    1
  ],
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "withAreaInfo": true,
  "withUserInfo": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "areaInfo": {
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
        "createdTime": "string",
        "id": "string",
        "userID": "string",
        "userInfo": {
          "avatar": "string",
          "email": "string",
          "nickName": "string",
          "phone": "string",
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
ur api /api/v1/things/data/area/user/apply/get-list \
  --body '{"areaID": "string", "authTypes": [1], "page": {"page": 1, "pageSize": 1}, "withAreaInfo": true, "withUserInfo": true}'
```
