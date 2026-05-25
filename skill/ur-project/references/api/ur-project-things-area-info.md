# ur-project things/area/info

新增项目区域 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/area/info/create` | 新增项目区域 | admin |
| POST | `/api/v1/things/area/info/delete` | 删除项目区域 | admin |
| POST | `/api/v1/things/area/info/get-list` | 获取项目区域列表 | all |
| POST | `/api/v1/things/area/info/get-one` | 获取项目区域详情 | all |
| POST | `/api/v1/things/area/info/update` | 更新项目区域 | admin |

## 详细说明

### POST `/api/v1/things/area/info/create`

**说明**: 新增项目区域

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id（只读） |
| `areaIDPath` | string | 否 | 项目区域ids（只读） |
| `areaImg` | string | 否 |  |
| `areaName` | string | 否 | 项目区域名称（读写） |
| `areaNamePath` | string | 否 | （只读） |
| `children` | array[AreaInfo] | 否 | 下级项目区域列表（只读） |
| `configFile` | string | 否 | 区域配置文件 |
| `createdTime` | string | 否 | 创建时间（只读） |
| `desc` | string | 否 | 项目区域备注（读写） |
| `deviceCount` | integer | 否 | 设备数量统计(只读) (格式: int64) |
| `devices` | array[DeviceInfo] | 否 | 区域下设备列表（按需返回） |
| `groupCount` | integer | 否 | 分组数量统计(只读) (格式: int64) |
| `isLeaf` | integer | 否 | 是否是子节点（只读） (格式: int64) |
| `isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `lowerLevelCount` | integer | 否 | 下级的数量统计（只读） (格式: int64) |
| `parentAreaID` | string | 否 | 上级项目区域id（只读） |
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `projectID` | string | 否 | 项目id（只读） |
| `sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `tags` | object | 否 | 自定义标签 |
| `tenantCode` | string | 否 | 租户编码（只读） |
| `useBy` | string | 否 | 用途: commonSpace default |

**请求示例**:
```json
{
  "areaID": "string",
  "areaIDPath": "string",
  "areaImg": "string",
  "areaName": "示例名称",
  "areaNamePath": "string",
  "children": [
    {
      "areaID": "string",
      "areaIDPath": "string",
      "areaImg": "string",
      "areaName": "示例名称",
      "areaNamePath": "string",
      "children": [
        {
          "areaID": "string",
          "areaIDPath": "string",
          "areaImg": "string",
          "areaName": "示例名称",
          "areaNamePath": "string",
          "children": [],
          "configFile": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [],
          "groupCount": 1,
          "isLeaf": 1,
          "isSysCreated": 1,
          "lowerLevelCount": 1,
          "parentAreaID": "string",
          "position": {
            "latitude": "...",
            "longitude": "..."
          },
          "projectID": "string",
          "sort": 1,
          "tags": {},
          "tenantCode": "string",
          "useBy": "string"
        }
      ],
      "configFile": "string",
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "deviceCount": 1,
      "devices": [
        {
          "adcode": "string",
          "address": "string",
          "area": {
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
          },
          "areaID": "string",
          "categoryID": "string",
          "cert": "string",
          "cloneID": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceAlias": "示例名称",
          "deviceImg": "string",
          "deviceName": "示例名称",
          "deviceType": 1,
          "distributor": {
            "id": "...",
            "idPath": "..."
          },
          "expTime": "2026-01-01T00:00:00Z",
          "file": "string",
          "firstBind": "string",
          "firstLogin": "2026-01-01T00:00:00Z",
          "gateway": {
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
          },
          "groupIDs": [],
          "groupPurpose": "string",
          "groups": [],
          "hardInfo": "string",
          "iccid": "string",
          "id": "string",
          "imei": "string",
          "isEnable": 1,
          "isOnline": 1,
          "lastBind": "2026-01-01T00:00:00Z",
          "lastIp": "string",
          "lastLocalIp": "string",
          "lastLogin": "2026-01-01T00:00:00Z",
          "lastOffline": "2026-01-01T00:00:00Z",
          "logLevel": "string",
          "mac": "string",
          "mobileOperator": 1,
          "needConfirmVersion": "string",
          "netType": 1,
          "owner": {
            "avatar": "...",
            "createdTime": "...",
            "email": "...",
            "lastIP": "...",
            "nickName": "...",
            "phone": "...",
            "regIP": "...",
            "userID": "...",
            "userName": "..."
          },
          "phone": "string",
          "position": {
            "latitude": "...",
            "longitude": "..."
          },
          "productID": "string",
          "productImg": "string",
          "productName": "string",
          "profiles": {},
          "projectID": "string",
          "protocolConf": {},
          "ratedPower": 1,
          "remoteAddr": "string",
          "rssi": 1,
          "schemaAlias": {},
          "secret": "string",
          "softInfo": "string",
          "sort": 1,
          "status": 1,
          "subProtocolConf": {},
          "tags": {},
          "tenantCode": "string",
          "userID": "string",
          "version": "string",
          "withProperties": {}
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
    }
  ],
  "configFile": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "deviceCount": 1,
  "devices": [
    {
      "adcode": "string",
      "address": "string",
      "area": {
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
      "areaID": "string",
      "categoryID": "string",
      "cert": "string",
      "cloneID": "string",
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "deviceAlias": "示例名称",
      "deviceImg": "string",
      "deviceName": "示例名称",
      "deviceType": 1,
      "distributor": {
        "id": "string",
        "idPath": "string"
      },
      "expTime": "2026-01-01T00:00:00Z",
      "file": "string",
      "firstBind": "string",
      "firstLogin": "2026-01-01T00:00:00Z",
      "gateway": {
        "adcode": "string",
        "address": "string",
        "area": {
          "areaID": "string",
          "areaIDPath": "string",
          "areaImg": "string",
          "areaName": "示例名称",
          "areaNamePath": "string",
          "children": [],
          "configFile": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [],
          "groupCount": 1,
          "isLeaf": 1,
          "isSysCreated": 1,
          "lowerLevelCount": 1,
          "parentAreaID": "string",
          "position": {
            "latitude": "...",
            "longitude": "..."
          },
          "projectID": "string",
          "sort": 1,
          "tags": {},
          "tenantCode": "string",
          "useBy": "string"
        },
        "areaID": "string",
        "categoryID": "string",
        "cert": "string",
        "cloneID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceAlias": "示例名称",
        "deviceImg": "string",
        "deviceName": "示例名称",
        "deviceType": 1,
        "distributor": {
          "id": "string",
          "idPath": "string"
        },
        "expTime": "2026-01-01T00:00:00Z",
        "file": "string",
        "firstBind": "string",
        "firstLogin": "2026-01-01T00:00:00Z",
        "gateway": {
          "adcode": "string",
          "address": "string",
          "area": {
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
          },
          "areaID": "string",
          "categoryID": "string",
          "cert": "string",
          "cloneID": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceAlias": "示例名称",
          "deviceImg": "string",
          "deviceName": "示例名称",
          "deviceType": 1,
          "distributor": {
            "id": "...",
            "idPath": "..."
          },
          "expTime": "2026-01-01T00:00:00Z",
          "file": "string",
          "firstBind": "string",
          "firstLogin": "2026-01-01T00:00:00Z",
          "gateway": {
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
          },
          "groupIDs": [],
          "groupPurpose": "string",
          "groups": [],
          "hardInfo": "string",
          "iccid": "string",
          "id": "string",
          "imei": "string",
          "isEnable": 1,
          "isOnline": 1,
          "lastBind": "2026-01-01T00:00:00Z",
          "lastIp": "string",
          "lastLocalIp": "string",
          "lastLogin": "2026-01-01T00:00:00Z",
          "lastOffline": "2026-01-01T00:00:00Z",
          "logLevel": "string",
          "mac": "string",
          "mobileOperator": 1,
          "needConfirmVersion": "string",
          "netType": 1,
          "owner": {
            "avatar": "...",
            "createdTime": "...",
            "email": "...",
            "lastIP": "...",
            "nickName": "...",
            "phone": "...",
            "regIP": "...",
            "userID": "...",
            "userName": "..."
          },
          "phone": "string",
          "position": {
            "latitude": "...",
            "longitude": "..."
          },
          "productID": "string",
          "productImg": "string",
          "productName": "string",
          "profiles": {},
          "projectID": "string",
          "protocolConf": {},
          "ratedPower": 1,
          "remoteAddr": "string",
          "rssi": 1,
          "schemaAlias": {},
          "secret": "string",
          "softInfo": "string",
          "sort": 1,
          "status": 1,
          "subProtocolConf": {},
          "tags": {},
          "tenantCode": "string",
          "userID": "string",
          "version": "string",
          "withProperties": {}
        },
        "groupIDs": [
          "string"
        ],
        "groupPurpose": "string",
        "groups": [
          {
            "id": "...",
            "name": "...",
            "parentID": "...",
            "purpose": "..."
          }
        ],
        "hardInfo": "string",
        "iccid": "string",
        "id": "string",
        "imei": "string",
        "isEnable": 1,
        "isOnline": 1,
        "lastBind": "2026-01-01T00:00:00Z",
        "lastIp": "string",
        "lastLocalIp": "string",
        "lastLogin": "2026-01-01T00:00:00Z",
        "lastOffline": "2026-01-01T00:00:00Z",
        "logLevel": "string",
        "mac": "string",
        "mobileOperator": 1,
        "needConfirmVersion": "string",
        "netType": 1,
        "owner": {
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
        "phone": "string",
        "position": {
          "latitude": 1,
          "longitude": 1
        },
        "productID": "string",
        "productImg": "string",
        "productName": "string",
        "profiles": {},
        "projectID": "string",
        "protocolConf": {},
        "ratedPower": 1,
        "remoteAddr": "string",
        "rssi": 1,
        "schemaAlias": {},
        "secret": "string",
        "softInfo": "string",
        "sort": 1,
        "status": 1,
        "subProtocolConf": {},
        "tags": {},
        "tenantCode": "string",
        "userID": "string",
        "version": "string",
        "withProperties": {}
      },
      "groupIDs": [
        "string"
      ],
      "groupPurpose": "string",
      "groups": [
        {
          "id": "string",
          "name": "string",
          "parentID": "string",
          "purpose": "string"
        }
      ],
      "hardInfo": "string",
      "iccid": "string",
      "id": "string",
      "imei": "string",
      "isEnable": 1,
      "isOnline": 1,
      "lastBind": "2026-01-01T00:00:00Z",
      "lastIp": "string",
      "lastLocalIp": "string",
      "lastLogin": "2026-01-01T00:00:00Z",
      "lastOffline": "2026-01-01T00:00:00Z",
      "logLevel": "string",
      "mac": "string",
      "mobileOperator": 1,
      "needConfirmVersion": "string",
      "netType": 1,
      "owner": {
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
      "phone": "string",
      "position": {
        "latitude": 1,
        "longitude": 1
      },
      "productID": "string",
      "productImg": "string",
      "productName": "string",
      "profiles": {},
      "projectID": "string",
      "protocolConf": {},
      "ratedPower": 1,
      "remoteAddr": "string",
      "rssi": 1,
      "schemaAlias": {},
      "secret": "string",
      "softInfo": "string",
      "sort": 1,
      "status": 1,
      "subProtocolConf": {},
      "tags": {},
      "tenantCode": "string",
      "userID": "string",
      "version": "string",
      "withProperties": {}
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
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "areaID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/area/info/create \
  --body '{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "...", "name": "...", "parentID": "...", "purpose": "..."}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}'
```

### POST `/api/v1/things/area/info/delete`

**说明**: 删除项目区域

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 是 | 项目区域id 只读 |

**请求示例**:
```json
{
  "areaID": "string"
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
ur api /api/v1/things/area/info/delete \
  --body '{"areaID": "string"}'
```

### POST `/api/v1/things/area/info/get-list`

**说明**: 获取项目区域列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaIDs` | array[string] | 否 | 项目区域ids |
| `deviceCount` | object | 否 |  |
| `deviceCount.cmpType` | string | 是 | "=":相等 "!=":不相等 ">":大于">=":大于等于"<":小于"<=":小于等于 "like":模糊查询 |
| `deviceCount.value` | string | 是 |  |
| `groupCount` | object | 否 |  |
| `groupCount.cmpType` | string | 是 | "=":相等 "!=":不相等 ">":大于">=":大于等于"<":小于"<=":小于等于 "like":模糊查询 |
| `groupCount.value` | string | 是 |  |
| `isRetTopLevel` | boolean | 否 | 如果该参数为true则返回除了root节点的有权限的最高层的区域列表 (格式: boolean) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `parentAreaID` | string | 否 |  |
| `projectID` | string | 否 | 项目id |
| `tenantCode` | string | 否 |  |
| `withDevices` | boolean | 否 | 是否同时返回区域下设备列表 (格式: boolean) |

**请求示例**:
```json
{
  "areaIDs": [
    "string"
  ],
  "deviceCount": {
    "cmpType": "string",
    "value": "string"
  },
  "groupCount": {
    "cmpType": "string",
    "value": "string"
  },
  "isRetTopLevel": true,
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
  "parentAreaID": "string",
  "projectID": "string",
  "tenantCode": "string",
  "withDevices": true
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
        "areaIDPath": "string",
        "areaImg": "string",
        "areaName": "示例名称",
        "areaNamePath": "string",
        "children": [
          {
            "areaID": "string",
            "areaIDPath": "string",
            "areaImg": "string",
            "areaName": "示例名称",
            "areaNamePath": "string",
            "children": [],
            "configFile": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "groupCount": 1,
            "isLeaf": 1,
            "isSysCreated": 1,
            "lowerLevelCount": 1,
            "parentAreaID": "string",
            "position": {
              "latitude": "...",
              "longitude": "..."
            },
            "projectID": "string",
            "sort": 1,
            "tags": {},
            "tenantCode": "string",
            "useBy": "string"
          }
        ],
        "configFile": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "devices": [
          {
            "adcode": "string",
            "address": "string",
            "area": {
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
            },
            "areaID": "string",
            "categoryID": "string",
            "cert": "string",
            "cloneID": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceAlias": "示例名称",
            "deviceImg": "string",
            "deviceName": "示例名称",
            "deviceType": 1,
            "distributor": {
              "id": "...",
              "idPath": "..."
            },
            "expTime": "2026-01-01T00:00:00Z",
            "file": "string",
            "firstBind": "string",
            "firstLogin": "2026-01-01T00:00:00Z",
            "gateway": {
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
            },
            "groupIDs": [],
            "groupPurpose": "string",
            "groups": [],
            "hardInfo": "string",
            "iccid": "string",
            "id": "string",
            "imei": "string",
            "isEnable": 1,
            "isOnline": 1,
            "lastBind": "2026-01-01T00:00:00Z",
            "lastIp": "string",
            "lastLocalIp": "string",
            "lastLogin": "2026-01-01T00:00:00Z",
            "lastOffline": "2026-01-01T00:00:00Z",
            "logLevel": "string",
            "mac": "string",
            "mobileOperator": 1,
            "needConfirmVersion": "string",
            "netType": 1,
            "owner": {
              "avatar": "...",
              "createdTime": "...",
              "email": "...",
              "lastIP": "...",
              "nickName": "...",
              "phone": "...",
              "regIP": "...",
              "userID": "...",
              "userName": "..."
            },
            "phone": "string",
            "position": {
              "latitude": "...",
              "longitude": "..."
            },
            "productID": "string",
            "productImg": "string",
            "productName": "string",
            "profiles": {},
            "projectID": "string",
            "protocolConf": {},
            "ratedPower": 1,
            "remoteAddr": "string",
            "rssi": 1,
            "schemaAlias": {},
            "secret": "string",
            "softInfo": "string",
            "sort": 1,
            "status": 1,
            "subProtocolConf": {},
            "tags": {},
            "tenantCode": "string",
            "userID": "string",
            "version": "string",
            "withProperties": {}
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
ur api /api/v1/things/area/info/get-list \
  --body '{"areaIDs": ["string"], "deviceCount": {"cmpType": "string", "value": "string"}, "groupCount": {"cmpType": "string", "value": "string"}, "isRetTopLevel": true, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentAreaID": "string", "projectID": "string", "tenantCode": "string", "withDevices": true}'
```

### POST `/api/v1/things/area/info/get-one`

**说明**: 获取项目区域详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 是 | 项目区域id |
| `projectID` | string | 否 | 项目id 不填选默认项目 |
| `tenantCode` | string | 否 |  |
| `withChildren` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "areaID": "string",
  "projectID": "string",
  "tenantCode": "string",
  "withChildren": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "areaID": "string",
    "areaIDPath": "string",
    "areaImg": "string",
    "areaName": "示例名称",
    "areaNamePath": "string",
    "children": [
      {
        "areaID": "string",
        "areaIDPath": "string",
        "areaImg": "string",
        "areaName": "示例名称",
        "areaNamePath": "string",
        "children": [
          {
            "areaID": "string",
            "areaIDPath": "string",
            "areaImg": "string",
            "areaName": "示例名称",
            "areaNamePath": "string",
            "children": [],
            "configFile": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "groupCount": 1,
            "isLeaf": 1,
            "isSysCreated": 1,
            "lowerLevelCount": 1,
            "parentAreaID": "string",
            "position": {
              "latitude": "...",
              "longitude": "..."
            },
            "projectID": "string",
            "sort": 1,
            "tags": {},
            "tenantCode": "string",
            "useBy": "string"
          }
        ],
        "configFile": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "devices": [
          {
            "adcode": "string",
            "address": "string",
            "area": {
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
            },
            "areaID": "string",
            "categoryID": "string",
            "cert": "string",
            "cloneID": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceAlias": "示例名称",
            "deviceImg": "string",
            "deviceName": "示例名称",
            "deviceType": 1,
            "distributor": {
              "id": "...",
              "idPath": "..."
            },
            "expTime": "2026-01-01T00:00:00Z",
            "file": "string",
            "firstBind": "string",
            "firstLogin": "2026-01-01T00:00:00Z",
            "gateway": {
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
            },
            "groupIDs": [],
            "groupPurpose": "string",
            "groups": [],
            "hardInfo": "string",
            "iccid": "string",
            "id": "string",
            "imei": "string",
            "isEnable": 1,
            "isOnline": 1,
            "lastBind": "2026-01-01T00:00:00Z",
            "lastIp": "string",
            "lastLocalIp": "string",
            "lastLogin": "2026-01-01T00:00:00Z",
            "lastOffline": "2026-01-01T00:00:00Z",
            "logLevel": "string",
            "mac": "string",
            "mobileOperator": 1,
            "needConfirmVersion": "string",
            "netType": 1,
            "owner": {
              "avatar": "...",
              "createdTime": "...",
              "email": "...",
              "lastIP": "...",
              "nickName": "...",
              "phone": "...",
              "regIP": "...",
              "userID": "...",
              "userName": "..."
            },
            "phone": "string",
            "position": {
              "latitude": "...",
              "longitude": "..."
            },
            "productID": "string",
            "productImg": "string",
            "productName": "string",
            "profiles": {},
            "projectID": "string",
            "protocolConf": {},
            "ratedPower": 1,
            "remoteAddr": "string",
            "rssi": 1,
            "schemaAlias": {},
            "secret": "string",
            "softInfo": "string",
            "sort": 1,
            "status": 1,
            "subProtocolConf": {},
            "tags": {},
            "tenantCode": "string",
            "userID": "string",
            "version": "string",
            "withProperties": {}
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
      }
    ],
    "configFile": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "deviceCount": 1,
    "devices": [
      {
        "adcode": "string",
        "address": "string",
        "area": {
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
        "areaID": "string",
        "categoryID": "string",
        "cert": "string",
        "cloneID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceAlias": "示例名称",
        "deviceImg": "string",
        "deviceName": "示例名称",
        "deviceType": 1,
        "distributor": {
          "id": "string",
          "idPath": "string"
        },
        "expTime": "2026-01-01T00:00:00Z",
        "file": "string",
        "firstBind": "string",
        "firstLogin": "2026-01-01T00:00:00Z",
        "gateway": {
          "adcode": "string",
          "address": "string",
          "area": {
            "areaID": "string",
            "areaIDPath": "string",
            "areaImg": "string",
            "areaName": "示例名称",
            "areaNamePath": "string",
            "children": [],
            "configFile": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "groupCount": 1,
            "isLeaf": 1,
            "isSysCreated": 1,
            "lowerLevelCount": 1,
            "parentAreaID": "string",
            "position": {
              "latitude": "...",
              "longitude": "..."
            },
            "projectID": "string",
            "sort": 1,
            "tags": {},
            "tenantCode": "string",
            "useBy": "string"
          },
          "areaID": "string",
          "categoryID": "string",
          "cert": "string",
          "cloneID": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceAlias": "示例名称",
          "deviceImg": "string",
          "deviceName": "示例名称",
          "deviceType": 1,
          "distributor": {
            "id": "string",
            "idPath": "string"
          },
          "expTime": "2026-01-01T00:00:00Z",
          "file": "string",
          "firstBind": "string",
          "firstLogin": "2026-01-01T00:00:00Z",
          "gateway": {
            "adcode": "string",
            "address": "string",
            "area": {
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
            },
            "areaID": "string",
            "categoryID": "string",
            "cert": "string",
            "cloneID": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceAlias": "示例名称",
            "deviceImg": "string",
            "deviceName": "示例名称",
            "deviceType": 1,
            "distributor": {
              "id": "...",
              "idPath": "..."
            },
            "expTime": "2026-01-01T00:00:00Z",
            "file": "string",
            "firstBind": "string",
            "firstLogin": "2026-01-01T00:00:00Z",
            "gateway": {
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
            },
            "groupIDs": [],
            "groupPurpose": "string",
            "groups": [],
            "hardInfo": "string",
            "iccid": "string",
            "id": "string",
            "imei": "string",
            "isEnable": 1,
            "isOnline": 1,
            "lastBind": "2026-01-01T00:00:00Z",
            "lastIp": "string",
            "lastLocalIp": "string",
            "lastLogin": "2026-01-01T00:00:00Z",
            "lastOffline": "2026-01-01T00:00:00Z",
            "logLevel": "string",
            "mac": "string",
            "mobileOperator": 1,
            "needConfirmVersion": "string",
            "netType": 1,
            "owner": {
              "avatar": "...",
              "createdTime": "...",
              "email": "...",
              "lastIP": "...",
              "nickName": "...",
              "phone": "...",
              "regIP": "...",
              "userID": "...",
              "userName": "..."
            },
            "phone": "string",
            "position": {
              "latitude": "...",
              "longitude": "..."
            },
            "productID": "string",
            "productImg": "string",
            "productName": "string",
            "profiles": {},
            "projectID": "string",
            "protocolConf": {},
            "ratedPower": 1,
            "remoteAddr": "string",
            "rssi": 1,
            "schemaAlias": {},
            "secret": "string",
            "softInfo": "string",
            "sort": 1,
            "status": 1,
            "subProtocolConf": {},
            "tags": {},
            "tenantCode": "string",
            "userID": "string",
            "version": "string",
            "withProperties": {}
          },
          "groupIDs": [
            "string"
          ],
          "groupPurpose": "string",
          "groups": [
            {
              "id": "...",
              "name": "...",
              "parentID": "...",
              "purpose": "..."
            }
          ],
          "hardInfo": "string",
          "iccid": "string",
          "id": "string",
          "imei": "string",
          "isEnable": 1,
          "isOnline": 1,
          "lastBind": "2026-01-01T00:00:00Z",
          "lastIp": "string",
          "lastLocalIp": "string",
          "lastLogin": "2026-01-01T00:00:00Z",
          "lastOffline": "2026-01-01T00:00:00Z",
          "logLevel": "string",
          "mac": "string",
          "mobileOperator": 1,
          "needConfirmVersion": "string",
          "netType": 1,
          "owner": {
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
          "phone": "string",
          "position": {
            "latitude": 1,
            "longitude": 1
          },
          "productID": "string",
          "productImg": "string",
          "productName": "string",
          "profiles": {},
          "projectID": "string",
          "protocolConf": {},
          "ratedPower": 1,
          "remoteAddr": "string",
          "rssi": 1,
          "schemaAlias": {},
          "secret": "string",
          "softInfo": "string",
          "sort": 1,
          "status": 1,
          "subProtocolConf": {},
          "tags": {},
          "tenantCode": "string",
          "userID": "string",
          "version": "string",
          "withProperties": {}
        },
        "groupIDs": [
          "string"
        ],
        "groupPurpose": "string",
        "groups": [
          {
            "id": "string",
            "name": "string",
            "parentID": "string",
            "purpose": "string"
          }
        ],
        "hardInfo": "string",
        "iccid": "string",
        "id": "string",
        "imei": "string",
        "isEnable": 1,
        "isOnline": 1,
        "lastBind": "2026-01-01T00:00:00Z",
        "lastIp": "string",
        "lastLocalIp": "string",
        "lastLogin": "2026-01-01T00:00:00Z",
        "lastOffline": "2026-01-01T00:00:00Z",
        "logLevel": "string",
        "mac": "string",
        "mobileOperator": 1,
        "needConfirmVersion": "string",
        "netType": 1,
        "owner": {
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
        "phone": "string",
        "position": {
          "latitude": 1,
          "longitude": 1
        },
        "productID": "string",
        "productImg": "string",
        "productName": "string",
        "profiles": {},
        "projectID": "string",
        "protocolConf": {},
        "ratedPower": 1,
        "remoteAddr": "string",
        "rssi": 1,
        "schemaAlias": {},
        "secret": "string",
        "softInfo": "string",
        "sort": 1,
        "status": 1,
        "subProtocolConf": {},
        "tags": {},
        "tenantCode": "string",
        "userID": "string",
        "version": "string",
        "withProperties": {}
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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/area/info/get-one \
  --body '{"areaID": "string", "projectID": "string", "tenantCode": "string", "withChildren": true}'
```

### POST `/api/v1/things/area/info/update`

**说明**: 更新项目区域

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id（只读） |
| `areaIDPath` | string | 否 | 项目区域ids（只读） |
| `areaImg` | string | 否 |  |
| `areaName` | string | 否 | 项目区域名称（读写） |
| `areaNamePath` | string | 否 | （只读） |
| `children` | array[AreaInfo] | 否 | 下级项目区域列表（只读） |
| `configFile` | string | 否 | 区域配置文件 |
| `createdTime` | string | 否 | 创建时间（只读） |
| `desc` | string | 否 | 项目区域备注（读写） |
| `deviceCount` | integer | 否 | 设备数量统计(只读) (格式: int64) |
| `devices` | array[DeviceInfo] | 否 | 区域下设备列表（按需返回） |
| `groupCount` | integer | 否 | 分组数量统计(只读) (格式: int64) |
| `isLeaf` | integer | 否 | 是否是子节点（只读） (格式: int64) |
| `isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `lowerLevelCount` | integer | 否 | 下级的数量统计（只读） (格式: int64) |
| `parentAreaID` | string | 否 | 上级项目区域id（只读） |
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `projectID` | string | 否 | 项目id（只读） |
| `sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `tags` | object | 否 | 自定义标签 |
| `tenantCode` | string | 否 | 租户编码（只读） |
| `useBy` | string | 否 | 用途: commonSpace default |

**请求示例**:
```json
{
  "areaID": "string",
  "areaIDPath": "string",
  "areaImg": "string",
  "areaName": "示例名称",
  "areaNamePath": "string",
  "children": [
    {
      "areaID": "string",
      "areaIDPath": "string",
      "areaImg": "string",
      "areaName": "示例名称",
      "areaNamePath": "string",
      "children": [
        {
          "areaID": "string",
          "areaIDPath": "string",
          "areaImg": "string",
          "areaName": "示例名称",
          "areaNamePath": "string",
          "children": [],
          "configFile": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [],
          "groupCount": 1,
          "isLeaf": 1,
          "isSysCreated": 1,
          "lowerLevelCount": 1,
          "parentAreaID": "string",
          "position": {
            "latitude": "...",
            "longitude": "..."
          },
          "projectID": "string",
          "sort": 1,
          "tags": {},
          "tenantCode": "string",
          "useBy": "string"
        }
      ],
      "configFile": "string",
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "deviceCount": 1,
      "devices": [
        {
          "adcode": "string",
          "address": "string",
          "area": {
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
          },
          "areaID": "string",
          "categoryID": "string",
          "cert": "string",
          "cloneID": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceAlias": "示例名称",
          "deviceImg": "string",
          "deviceName": "示例名称",
          "deviceType": 1,
          "distributor": {
            "id": "...",
            "idPath": "..."
          },
          "expTime": "2026-01-01T00:00:00Z",
          "file": "string",
          "firstBind": "string",
          "firstLogin": "2026-01-01T00:00:00Z",
          "gateway": {
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
          },
          "groupIDs": [],
          "groupPurpose": "string",
          "groups": [],
          "hardInfo": "string",
          "iccid": "string",
          "id": "string",
          "imei": "string",
          "isEnable": 1,
          "isOnline": 1,
          "lastBind": "2026-01-01T00:00:00Z",
          "lastIp": "string",
          "lastLocalIp": "string",
          "lastLogin": "2026-01-01T00:00:00Z",
          "lastOffline": "2026-01-01T00:00:00Z",
          "logLevel": "string",
          "mac": "string",
          "mobileOperator": 1,
          "needConfirmVersion": "string",
          "netType": 1,
          "owner": {
            "avatar": "...",
            "createdTime": "...",
            "email": "...",
            "lastIP": "...",
            "nickName": "...",
            "phone": "...",
            "regIP": "...",
            "userID": "...",
            "userName": "..."
          },
          "phone": "string",
          "position": {
            "latitude": "...",
            "longitude": "..."
          },
          "productID": "string",
          "productImg": "string",
          "productName": "string",
          "profiles": {},
          "projectID": "string",
          "protocolConf": {},
          "ratedPower": 1,
          "remoteAddr": "string",
          "rssi": 1,
          "schemaAlias": {},
          "secret": "string",
          "softInfo": "string",
          "sort": 1,
          "status": 1,
          "subProtocolConf": {},
          "tags": {},
          "tenantCode": "string",
          "userID": "string",
          "version": "string",
          "withProperties": {}
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
    }
  ],
  "configFile": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "deviceCount": 1,
  "devices": [
    {
      "adcode": "string",
      "address": "string",
      "area": {
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
      "areaID": "string",
      "categoryID": "string",
      "cert": "string",
      "cloneID": "string",
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "deviceAlias": "示例名称",
      "deviceImg": "string",
      "deviceName": "示例名称",
      "deviceType": 1,
      "distributor": {
        "id": "string",
        "idPath": "string"
      },
      "expTime": "2026-01-01T00:00:00Z",
      "file": "string",
      "firstBind": "string",
      "firstLogin": "2026-01-01T00:00:00Z",
      "gateway": {
        "adcode": "string",
        "address": "string",
        "area": {
          "areaID": "string",
          "areaIDPath": "string",
          "areaImg": "string",
          "areaName": "示例名称",
          "areaNamePath": "string",
          "children": [],
          "configFile": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [],
          "groupCount": 1,
          "isLeaf": 1,
          "isSysCreated": 1,
          "lowerLevelCount": 1,
          "parentAreaID": "string",
          "position": {
            "latitude": "...",
            "longitude": "..."
          },
          "projectID": "string",
          "sort": 1,
          "tags": {},
          "tenantCode": "string",
          "useBy": "string"
        },
        "areaID": "string",
        "categoryID": "string",
        "cert": "string",
        "cloneID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceAlias": "示例名称",
        "deviceImg": "string",
        "deviceName": "示例名称",
        "deviceType": 1,
        "distributor": {
          "id": "string",
          "idPath": "string"
        },
        "expTime": "2026-01-01T00:00:00Z",
        "file": "string",
        "firstBind": "string",
        "firstLogin": "2026-01-01T00:00:00Z",
        "gateway": {
          "adcode": "string",
          "address": "string",
          "area": {
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
          },
          "areaID": "string",
          "categoryID": "string",
          "cert": "string",
          "cloneID": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceAlias": "示例名称",
          "deviceImg": "string",
          "deviceName": "示例名称",
          "deviceType": 1,
          "distributor": {
            "id": "...",
            "idPath": "..."
          },
          "expTime": "2026-01-01T00:00:00Z",
          "file": "string",
          "firstBind": "string",
          "firstLogin": "2026-01-01T00:00:00Z",
          "gateway": {
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
          },
          "groupIDs": [],
          "groupPurpose": "string",
          "groups": [],
          "hardInfo": "string",
          "iccid": "string",
          "id": "string",
          "imei": "string",
          "isEnable": 1,
          "isOnline": 1,
          "lastBind": "2026-01-01T00:00:00Z",
          "lastIp": "string",
          "lastLocalIp": "string",
          "lastLogin": "2026-01-01T00:00:00Z",
          "lastOffline": "2026-01-01T00:00:00Z",
          "logLevel": "string",
          "mac": "string",
          "mobileOperator": 1,
          "needConfirmVersion": "string",
          "netType": 1,
          "owner": {
            "avatar": "...",
            "createdTime": "...",
            "email": "...",
            "lastIP": "...",
            "nickName": "...",
            "phone": "...",
            "regIP": "...",
            "userID": "...",
            "userName": "..."
          },
          "phone": "string",
          "position": {
            "latitude": "...",
            "longitude": "..."
          },
          "productID": "string",
          "productImg": "string",
          "productName": "string",
          "profiles": {},
          "projectID": "string",
          "protocolConf": {},
          "ratedPower": 1,
          "remoteAddr": "string",
          "rssi": 1,
          "schemaAlias": {},
          "secret": "string",
          "softInfo": "string",
          "sort": 1,
          "status": 1,
          "subProtocolConf": {},
          "tags": {},
          "tenantCode": "string",
          "userID": "string",
          "version": "string",
          "withProperties": {}
        },
        "groupIDs": [
          "string"
        ],
        "groupPurpose": "string",
        "groups": [
          {
            "id": "...",
            "name": "...",
            "parentID": "...",
            "purpose": "..."
          }
        ],
        "hardInfo": "string",
        "iccid": "string",
        "id": "string",
        "imei": "string",
        "isEnable": 1,
        "isOnline": 1,
        "lastBind": "2026-01-01T00:00:00Z",
        "lastIp": "string",
        "lastLocalIp": "string",
        "lastLogin": "2026-01-01T00:00:00Z",
        "lastOffline": "2026-01-01T00:00:00Z",
        "logLevel": "string",
        "mac": "string",
        "mobileOperator": 1,
        "needConfirmVersion": "string",
        "netType": 1,
        "owner": {
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
        "phone": "string",
        "position": {
          "latitude": 1,
          "longitude": 1
        },
        "productID": "string",
        "productImg": "string",
        "productName": "string",
        "profiles": {},
        "projectID": "string",
        "protocolConf": {},
        "ratedPower": 1,
        "remoteAddr": "string",
        "rssi": 1,
        "schemaAlias": {},
        "secret": "string",
        "softInfo": "string",
        "sort": 1,
        "status": 1,
        "subProtocolConf": {},
        "tags": {},
        "tenantCode": "string",
        "userID": "string",
        "version": "string",
        "withProperties": {}
      },
      "groupIDs": [
        "string"
      ],
      "groupPurpose": "string",
      "groups": [
        {
          "id": "string",
          "name": "string",
          "parentID": "string",
          "purpose": "string"
        }
      ],
      "hardInfo": "string",
      "iccid": "string",
      "id": "string",
      "imei": "string",
      "isEnable": 1,
      "isOnline": 1,
      "lastBind": "2026-01-01T00:00:00Z",
      "lastIp": "string",
      "lastLocalIp": "string",
      "lastLogin": "2026-01-01T00:00:00Z",
      "lastOffline": "2026-01-01T00:00:00Z",
      "logLevel": "string",
      "mac": "string",
      "mobileOperator": 1,
      "needConfirmVersion": "string",
      "netType": 1,
      "owner": {
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
      "phone": "string",
      "position": {
        "latitude": 1,
        "longitude": 1
      },
      "productID": "string",
      "productImg": "string",
      "productName": "string",
      "profiles": {},
      "projectID": "string",
      "protocolConf": {},
      "ratedPower": 1,
      "remoteAddr": "string",
      "rssi": 1,
      "schemaAlias": {},
      "secret": "string",
      "softInfo": "string",
      "sort": 1,
      "status": 1,
      "subProtocolConf": {},
      "tags": {},
      "tenantCode": "string",
      "userID": "string",
      "version": "string",
      "withProperties": {}
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
ur api /api/v1/things/area/info/update \
  --body '{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "...", "name": "...", "parentID": "...", "purpose": "..."}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}'
```
