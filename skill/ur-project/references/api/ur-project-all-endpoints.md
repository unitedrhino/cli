# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/area/info/create` | 新增项目区域 | admin |
| POST | `/api/v1/things/area/info/delete` | 删除项目区域 | admin |
| POST | `/api/v1/things/area/info/get-list` | 获取项目区域列表 | all |
| POST | `/api/v1/things/area/info/get-one` | 获取项目区域详情 | all |
| POST | `/api/v1/things/area/info/update` | 更新项目区域 | admin |
| POST | `/api/v1/things/area/profile/get-list` | 获取区域配置列表 | admin |
| POST | `/api/v1/things/area/profile/get-one` | 获取区域配置详情 | admin |
| POST | `/api/v1/things/area/profile/update` | 更新区域配置 | admin |
| POST | `/api/v1/things/data/area/batch-delete` | 删除授权区域权限 | admin |
| POST | `/api/v1/things/data/area/batch-update` | 更新授权区域权限 | admin |
| POST | `/api/v1/things/data/area/get-list` | 获取区域权限列表 | admin |
| POST | `/api/v1/things/data/area/user/apply/deal` | 授权区域权限 | admin |
| POST | `/api/v1/things/data/area/user/apply/get-list` | 获取区域权限列表 | admin |
| POST | `/api/v1/things/data/project/batch-create` | 批量创建授权项目权限 | admin |
| POST | `/api/v1/things/data/project/batch-delete` | 批量删除授权项目权限 | admin |
| POST | `/api/v1/things/data/project/create` | 创建授权项目权限 | admin |
| POST | `/api/v1/things/data/project/delete` | 删除授权项目权限 | admin |
| POST | `/api/v1/things/data/project/get-list` | 获取项目权限列表 | admin |
| POST | `/api/v1/things/group/device/batch-create` | 添加分组设备 | admin |
| POST | `/api/v1/things/group/device/batch-delete` | 删除分组设备 | admin |
| POST | `/api/v1/things/group/device/batch-update` | 更新分组设备 | admin |
| POST | `/api/v1/things/group/info/create` | 创建分组 | admin |
| POST | `/api/v1/things/group/info/delete` | 删除分组 | admin |
| POST | `/api/v1/things/group/info/get-list` | 获取分组列表 | admin |
| POST | `/api/v1/things/group/info/get-one` | 获取分组详情信息 | admin |
| POST | `/api/v1/things/group/info/update` | 更新分组信息 | admin |
| POST | `/api/v1/things/project/crud/create` | 新增项目crud | admin |
| POST | `/api/v1/things/project/crud/update` | 更新项目crud | admin |
| POST | `/api/v1/things/project/crud/delete` | 删除项目crud | admin |
| POST | `/api/v1/things/project/crud/get-list` | 获取项目crud列表 | admin |
| POST | `/api/v1/things/project/crud/get-one` | 获取项目crud详情 | admin |
| POST | `/api/v1/things/project/info/create` | 新增项目 | admin |
| POST | `/api/v1/things/project/info/delete` | 删除项目 | admin |
| POST | `/api/v1/things/project/info/get-list` | 获取项目列表 | all |
| POST | `/api/v1/things/project/info/get-one` | 获取项目详情 | all |
| POST | `/api/v1/things/project/info/update` | 更新项目 | admin |
| POST | `/api/v1/things/project/profile/get-list` | 获取项目配置列表 | admin |
| POST | `/api/v1/things/project/profile/get-one` | 获取项目配置详情 | admin |
| POST | `/api/v1/things/project/profile/update` | 更新项目配置 | admin |

## 端点详情

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

### POST `/api/v1/things/area/profile/get-list`

**说明**: 获取区域配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 是 | 项目区域id（只读） |
| `codes` | array[string] | 是 |  |

**请求示例**:
```json
{
  "areaID": "string",
  "codes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "profiles": [
      {
        "areaID": "string",
        "code": "string",
        "params": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/area/profile/get-list \
  --body '{"areaID": "string", "codes": ["string"]}'
```

### POST `/api/v1/things/area/profile/get-one`

**说明**: 获取区域配置详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 是 | 项目区域id（只读） |
| `code` | string | 是 |  |

**请求示例**:
```json
{
  "areaID": "string",
  "code": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "areaID": "string",
    "code": "string",
    "params": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/area/profile/get-one \
  --body '{"areaID": "string", "code": "string"}'
```

### POST `/api/v1/things/area/profile/update`

**说明**: 更新区域配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 是 | 项目区域id（只读） |
| `code` | string | 是 |  |
| `params` | string | 是 |  |

**请求示例**:
```json
{
  "areaID": "string",
  "code": "string",
  "params": "string"
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
ur api /api/v1/things/area/profile/update \
  --body '{"areaID": "string", "code": "string", "params": "string"}'
```

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
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
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
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
  --body '{"areaID": "string", "authTypes": [1], "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "withAreaInfo": true, "withUserInfo": true}'
```

### POST `/api/v1/things/data/project/batch-create`

**说明**: 批量创建授权项目权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authType` | integer | 是 |  1:管理权限,可以修改别人的权限,及读写权限 2:读写权限,可以读写该权限 3:读权限,只能读,不能修改 (格式: int64) |
| `projectID` | string | 是 | 权限数据ID |
| `targetIDs` | array[string] | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "authType": 1,
  "projectID": "string",
  "targetIDs": [
    "string"
  ],
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
ur api /api/v1/things/data/project/batch-create \
  --body '{"authType": 1, "projectID": "string", "targetIDs": ["string"], "targetType": "string"}'
```

### POST `/api/v1/things/data/project/batch-delete`

**说明**: 批量删除授权项目权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `projectID` | string | 否 | 权限数据ID |
| `targetIDs` | array[string] | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "projectID": "string",
  "targetIDs": [
    "string"
  ],
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
ur api /api/v1/things/data/project/batch-delete \
  --body '{"projectID": "string", "targetIDs": ["string"], "targetType": "string"}'
```

### POST `/api/v1/things/data/project/create`

**说明**: 创建授权项目权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authType` | integer | 是 |  1:管理权限,可以修改别人的权限,及读写权限 2:读写权限,可以读写该权限 3:读权限,只能读,不能修改 (格式: int64) |
| `projectID` | string | 是 | 权限数据ID |
| `targetID` | string | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
  "authType": 1,
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
    "authType": 1,
    "projectID": "string",
    "targetID": "string",
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
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/data/project/create \
  --body '{"authType": 1, "projectID": "string", "targetID": "string", "targetType": "string"}'
```

### POST `/api/v1/things/data/project/delete`

**说明**: 删除授权项目权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `projectID` | string | 否 | 权限数据ID |
| `targetID` | string | 是 | 用户ID |
| `targetType` | string | 是 |  |

**请求示例**:
```json
{
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
ur api /api/v1/things/data/project/delete \
  --body '{"projectID": "string", "targetID": "string", "targetType": "string"}'
```

### POST `/api/v1/things/data/project/get-list`

**说明**: 获取项目权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `projectID` | string | 否 | 权限数据ID |
| `targetID` | string | 否 | 用户ID |
| `targetType` | string | 是 |  |

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
        "authType": 1,
        "projectID": "string",
        "targetID": "string",
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
ur api /api/v1/things/data/project/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "projectID": "string", "targetID": "string", "targetType": "string"}'
```

### POST `/api/v1/things/group/device/batch-create`

**说明**: 添加分组设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupID` | string | 是 | 分组ID |
| `list` | array[DeviceCore] | 否 | 分组tag |

**请求示例**:
```json
{
  "groupID": "string",
  "list": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
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
ur api /api/v1/things/group/device/batch-create \
  --body '{"groupID": "string", "list": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}]}'
```

### POST `/api/v1/things/group/device/batch-delete`

**说明**: 删除分组设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupID` | string | 是 | 分组ID |
| `list` | array[DeviceCore] | 否 | 分组tag |

**请求示例**:
```json
{
  "groupID": "string",
  "list": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
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
ur api /api/v1/things/group/device/batch-delete \
  --body '{"groupID": "string", "list": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}]}'
```

### POST `/api/v1/things/group/device/batch-update`

**说明**: 更新分组设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupID` | string | 是 | 分组ID |
| `list` | array[DeviceCore] | 否 | 分组tag |

**请求示例**:
```json
{
  "groupID": "string",
  "list": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
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
ur api /api/v1/things/group/device/batch-update \
  --body '{"groupID": "string", "list": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}]}'
```

### POST `/api/v1/things/group/info/create`

**说明**: 创建分组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 区域ID |
| `children` | array[GroupInfo] | 否 |  |
| `createdTime` | string | 否 | 创建时间 |
| `desc` | string | 否 | 分组描述 |
| `deviceCount` | integer | 否 | 设备数量统计 (格式: int64) |
| `devices` | array[DeviceCore] | 否 |  |
| `files` | object | 否 | key 是用途, value是文件的路径, |
| `id` | string | 否 | 分组ID |
| `idPath` | string | 否 | 分组ID路径 |
| `isLeaf` | integer | 否 | 格式: int64 |
| `name` | string | 否 | 分组名称 |
| `parentID` | string | 否 | 父组ID |
| `productID` | string | 否 | 产品ID |
| `productName` | string | 否 | 产品ID |
| `projectID` | string | 否 | 项目ID |
| `purpose` | string | 否 | 用途 不填默认为default |
| `tags` | object | 否 | 分组tag |

**请求示例**:
```json
{
  "areaID": "string",
  "children": [
    {
      "areaID": "string",
      "children": [
        {
          "areaID": "string",
          "children": [],
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [],
          "files": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "name": "示例名称",
          "parentID": "string",
          "productID": "string",
          "productName": "string",
          "projectID": "string",
          "purpose": "string",
          "tags": {}
        }
      ],
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "deviceCount": 1,
      "devices": [
        {
          "deviceName": "示例名称",
          "productID": "string",
          "productName": "string"
        }
      ],
      "files": {},
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "name": "示例名称",
      "parentID": "string",
      "productID": "string",
      "productName": "string",
      "projectID": "string",
      "purpose": "string",
      "tags": {}
    }
  ],
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "deviceCount": 1,
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
    }
  ],
  "files": {},
  "id": "string",
  "idPath": "string",
  "isLeaf": 1,
  "name": "示例名称",
  "parentID": "string",
  "productID": "string",
  "productName": "string",
  "projectID": "string",
  "purpose": "string",
  "tags": {}
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/group/info/create \
  --body '{"areaID": "string", "children": [{"areaID": "string", "children": [{"areaID": "string", "children": [], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}'
```

### POST `/api/v1/things/group/info/delete`

**说明**: 删除分组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
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
ur api /api/v1/things/group/info/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/group/info/get-list`

**说明**: 获取分组列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 区域ID |
| `name` | string | 否 | 分组名称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 | 父组ID |
| `productID` | string | 否 | 产品ID |
| `purpose` | string | 否 | 用途 不填默认为default |
| `tags` | object | 否 | 分组tag |

**请求示例**:
```json
{
  "areaID": "string",
  "name": "示例名称",
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
  "parentID": "string",
  "productID": "string",
  "purpose": "string",
  "tags": {}
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
        "children": [
          {
            "areaID": "string",
            "children": [],
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "files": {},
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "name": "示例名称",
            "parentID": "string",
            "productID": "string",
            "productName": "string",
            "projectID": "string",
            "purpose": "string",
            "tags": {}
          }
        ],
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "devices": [
          {
            "deviceName": "示例名称",
            "productID": "string",
            "productName": "string"
          }
        ],
        "files": {},
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "name": "示例名称",
        "parentID": "string",
        "productID": "string",
        "productName": "string",
        "projectID": "string",
        "purpose": "string",
        "tags": {}
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
ur api /api/v1/things/group/info/get-list \
  --body '{"areaID": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentID": "string", "productID": "string", "purpose": "string", "tags": {}}'
```

### POST `/api/v1/things/group/info/get-one`

**说明**: 获取分组详情信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  id |
| `purpose` | string | 否 | 用途 不填默认为default |
| `withChildren` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "id": "string",
  "purpose": "string",
  "withChildren": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "areaID": "string",
    "children": [
      {
        "areaID": "string",
        "children": [
          {
            "areaID": "string",
            "children": [],
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "files": {},
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "name": "示例名称",
            "parentID": "string",
            "productID": "string",
            "productName": "string",
            "projectID": "string",
            "purpose": "string",
            "tags": {}
          }
        ],
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "devices": [
          {
            "deviceName": "示例名称",
            "productID": "string",
            "productName": "string"
          }
        ],
        "files": {},
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "name": "示例名称",
        "parentID": "string",
        "productID": "string",
        "productName": "string",
        "projectID": "string",
        "purpose": "string",
        "tags": {}
      }
    ],
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "deviceCount": 1,
    "devices": [
      {
        "deviceName": "示例名称",
        "productID": "string",
        "productName": "string"
      }
    ],
    "files": {},
    "id": "string",
    "idPath": "string",
    "isLeaf": 1,
    "name": "示例名称",
    "parentID": "string",
    "productID": "string",
    "productName": "string",
    "projectID": "string",
    "purpose": "string",
    "tags": {}
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/group/info/get-one \
  --body '{"id": "string", "purpose": "string", "withChildren": true}'
```

### POST `/api/v1/things/group/info/update`

**说明**: 更新分组信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 区域ID |
| `children` | array[GroupInfo] | 否 |  |
| `createdTime` | string | 否 | 创建时间 |
| `desc` | string | 否 | 分组描述 |
| `deviceCount` | integer | 否 | 设备数量统计 (格式: int64) |
| `devices` | array[DeviceCore] | 否 |  |
| `files` | object | 否 | key 是用途, value是文件的路径, |
| `id` | string | 否 | 分组ID |
| `idPath` | string | 否 | 分组ID路径 |
| `isLeaf` | integer | 否 | 格式: int64 |
| `name` | string | 否 | 分组名称 |
| `parentID` | string | 否 | 父组ID |
| `productID` | string | 否 | 产品ID |
| `productName` | string | 否 | 产品ID |
| `projectID` | string | 否 | 项目ID |
| `purpose` | string | 否 | 用途 不填默认为default |
| `tags` | object | 否 | 分组tag |

**请求示例**:
```json
{
  "areaID": "string",
  "children": [
    {
      "areaID": "string",
      "children": [
        {
          "areaID": "string",
          "children": [],
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [],
          "files": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "name": "示例名称",
          "parentID": "string",
          "productID": "string",
          "productName": "string",
          "projectID": "string",
          "purpose": "string",
          "tags": {}
        }
      ],
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "deviceCount": 1,
      "devices": [
        {
          "deviceName": "示例名称",
          "productID": "string",
          "productName": "string"
        }
      ],
      "files": {},
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "name": "示例名称",
      "parentID": "string",
      "productID": "string",
      "productName": "string",
      "projectID": "string",
      "purpose": "string",
      "tags": {}
    }
  ],
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "deviceCount": 1,
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
    }
  ],
  "files": {},
  "id": "string",
  "idPath": "string",
  "isLeaf": 1,
  "name": "示例名称",
  "parentID": "string",
  "productID": "string",
  "productName": "string",
  "projectID": "string",
  "purpose": "string",
  "tags": {}
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
ur api /api/v1/things/group/info/update \
  --body '{"areaID": "string", "children": [{"areaID": "string", "children": [{"areaID": "string", "children": [], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}'
```

### POST `/api/v1/things/project/crud/create`

**说明**: 新增项目crud

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 只读 |
| `id` | string | 否 |  |
| `params` | string | 否 | 普通业务参数,如果是文件类型,则需要以 Img,File结尾,后端会进行文件的处理,上传方式可以参考设备图片 |
| `purpose` | string | 否 | 用途 新增必填 |
| `sort` | integer | 否 | 排序 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": "string",
  "id": "string",
  "params": "string",
  "purpose": "string",
  "sort": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/crud/create \
  --body '{"createdTime": "string", "id": "string", "params": "string", "purpose": "string", "sort": 1}'
```

### POST `/api/v1/things/project/crud/update`

**说明**: 更新项目crud

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 只读 |
| `id` | string | 否 |  |
| `params` | string | 否 | 普通业务参数,如果是文件类型,则需要以 Img,File结尾,后端会进行文件的处理,上传方式可以参考设备图片 |
| `purpose` | string | 否 | 用途 新增必填 |
| `sort` | integer | 否 | 排序 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": "string",
  "id": "string",
  "params": "string",
  "purpose": "string",
  "sort": 1
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
ur api /api/v1/things/project/crud/update \
  --body '{"createdTime": "string", "id": "string", "params": "string", "purpose": "string", "sort": 1}'
```

### POST `/api/v1/things/project/crud/delete`

**说明**: 删除项目crud

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
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
ur api /api/v1/things/project/crud/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/project/crud/get-list`

**说明**: 获取项目crud列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `params` | object | 否 | params过滤查询,key 是params的key, value是对象,说明: {"cmpType":"比较类型(如=、>、like)","value":"比较值(需与类型匹配)","castTo":"数据类型(int/float/string,选填，填则转换)"}，其中 like 仅支持 string 类型，value 需以字符串格式传入 示例: {"aaa":{"cmpType":">","castTo":"int","value":"100"} |
| `purpose` | string | 是 | 用途必填 |

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
  "params": {},
  "purpose": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "string",
        "id": "string",
        "params": "string",
        "purpose": "string",
        "sort": 1
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
ur api /api/v1/things/project/crud/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "params": {}, "purpose": "string"}'
```

### POST `/api/v1/things/project/crud/get-one`

**说明**: 获取项目crud详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdTime": "string",
    "id": "string",
    "params": "string",
    "purpose": "string",
    "sort": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/crud/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/project/info/create`

**说明**: 新增项目

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `address` | string | 否 |  可选：项目详细地址 |
| `adminUserID` | string | 否 |  可选：管理员用户id，默认当前登录用户 |
| `area` | number | 否 |  可选：建筑面积（单位平米） (格式: float) |
| `attachments` | array[Attachment] | 否 |  可选：附件列表 |
| `desc` | string | 否 |  可选：项目备注 |
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `ppsm` | integer | 否 |  可选：每平方米功耗 (格式: int64) |
| `projectImg` | string | 否 |  可选：项目图片 |
| `projectName` | string | 是 |  必填：项目名称 |
| `sort` | integer | 否 |  可选：排序标记，默认为1 (格式: int64) |
| `status` | integer | 否 |  可选：项目状态 1:正常 2:禁用 3:过期 (格式: int64) |
| `tags` | object | 否 |  可选：自定义标签 |
| `type` | string | 否 |  可选：项目类型 |

**请求示例**:
```json
{
  "address": "string",
  "adminUserID": "string",
  "area": 1,
  "attachments": [
    {
      "fileName": "string",
      "filePath": "string",
      "fileUrl": "string",
      "id": "string",
      "useBy": "string"
    }
  ],
  "desc": "string",
  "position": {
    "latitude": 1,
    "longitude": 1
  },
  "ppsm": 1,
  "projectImg": "string",
  "projectName": "示例名称",
  "sort": 1,
  "status": 1,
  "tags": {},
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "projectID": "string",
    "withAdminUser": true,
    "withTopAreas": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/info/create \
  --body '{"address": "string", "adminUserID": "string", "area": 1, "attachments": [{"fileName": "string", "filePath": "string", "fileUrl": "string", "id": "string", "useBy": "string"}], "desc": "string", "position": {"latitude": 1, "longitude": 1}, "ppsm": 1, "projectImg": "string", "projectName": "示例名称", "sort": 1, "status": 1, "tags": {}, "type": "string"}'
```

### POST `/api/v1/things/project/info/delete`

**说明**: 删除项目

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `projectID` | string | 是 | 项目id 只读 |
| `withAdminUser` | boolean | 否 | 同时获取管理员核心信息 (格式: boolean) |
| `withTopAreas` | boolean | 否 | 同时返回顶层的区域列表 (格式: boolean) |

**请求示例**:
```json
{
  "projectID": "string",
  "withAdminUser": true,
  "withTopAreas": true
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
ur api /api/v1/things/project/info/delete \
  --body '{"projectID": "string", "withAdminUser": true, "withTopAreas": true}'
```

### POST `/api/v1/things/project/info/get-list`

**说明**: 获取项目列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmStatus` | integer | 否 | 报警状态(只读) （1正常 2提醒 3一般 4严重 5紧急 6超紧急） (格式: int64) |
| `isGetAll` | boolean | 否 | 是否返回所有,只有平台管理员有权限 (格式: boolean) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `projectIDs` | array[string] | 否 | 过滤项目id列表 |
| `projectName` | string | 否 | 过滤项目名称 |
| `status` | integer | 否 | 项目状态  1 正常，2-禁用，3-过期 (格式: int64) |
| `tenantCode` | string | 否 | 指定租户(只有default租户下的管理员才有效) |
| `type` | string | 否 | 项目类型 |
| `withAdminUser` | boolean | 否 | 同时获取管理员核心信息 (格式: boolean) |
| `withTopAreas` | boolean | 否 | 同时返回顶层的区域列表 (格式: boolean) |

**请求示例**:
```json
{
  "alarmStatus": 1,
  "isGetAll": true,
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
  "projectIDs": [
    "string"
  ],
  "projectName": "示例名称",
  "status": 1,
  "tenantCode": "string",
  "type": "string",
  "withAdminUser": true,
  "withTopAreas": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "address": "string",
        "adminUser": {
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
        "adminUserID": "string",
        "alarmStatus": 1,
        "area": 1,
        "areaCount": 1,
        "areas": [
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
        "attachments": [
          {
            "fileName": "string",
            "filePath": "string",
            "fileUrl": "string",
            "id": "string",
            "useBy": "string"
          }
        ],
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "deviceOnlineCount": 1,
        "isSysCreated": 1,
        "position": {
          "latitude": 1,
          "longitude": 1
        },
        "ppsm": 1,
        "projectID": "string",
        "projectImg": "string",
        "projectName": "示例名称",
        "sort": 1,
        "status": 1,
        "tags": {},
        "tenantCode": "string",
        "type": "string",
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
ur api /api/v1/things/project/info/get-list \
  --body '{"alarmStatus": 1, "isGetAll": true, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "projectIDs": ["string"], "projectName": "示例名称", "status": 1, "tenantCode": "string", "type": "string", "withAdminUser": true, "withTopAreas": true}'
```

### POST `/api/v1/things/project/info/get-one`

**说明**: 获取项目详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `projectID` | string | 是 | 项目id 只读 |
| `withAdminUser` | boolean | 否 | 同时获取管理员核心信息 (格式: boolean) |
| `withTopAreas` | boolean | 否 | 同时返回顶层的区域列表 (格式: boolean) |

**请求示例**:
```json
{
  "projectID": "string",
  "withAdminUser": true,
  "withTopAreas": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "address": "string",
    "adminUser": {
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
    "adminUserID": "string",
    "alarmStatus": 1,
    "area": 1,
    "areaCount": 1,
    "areas": [
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
    "attachments": [
      {
        "fileName": "string",
        "filePath": "string",
        "fileUrl": "string",
        "id": "string",
        "useBy": "string"
      }
    ],
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "deviceCount": 1,
    "deviceOnlineCount": 1,
    "isSysCreated": 1,
    "position": {
      "latitude": 1,
      "longitude": 1
    },
    "ppsm": 1,
    "projectID": "string",
    "projectImg": "string",
    "projectName": "示例名称",
    "sort": 1,
    "status": 1,
    "tags": {},
    "tenantCode": "string",
    "type": "string",
    "userCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/info/get-one \
  --body '{"projectID": "string", "withAdminUser": true, "withTopAreas": true}'
```

### POST `/api/v1/things/project/info/update`

**说明**: 更新项目

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `address` | string | 否 | 项目详细地址（读写） |
| `adminUser` | object | 否 |  |
| `adminUser.avatar` | string | 否 |  用户头像 |
| `adminUser.createdTime` | string | 否 |  创建时间 |
| `adminUser.email` | string | 否 |  邮箱 |
| `adminUser.lastIP` | string | 否 |  最后登录ip |
| `adminUser.nickName` | string | 否 |  用户的昵称 |
| `adminUser.phone` | string | 否 |  手机号 |
| `adminUser.regIP` | string | 否 |  注册ip |
| `adminUser.userID` | string | 否 |  用户id |
| `adminUser.userName` | string | 否 |  用户名(唯一) |
| `adminUserID` | string | 否 |  管理员用户id（读写） |
| `alarmStatus` | integer | 否 | 报警状态(只读) (格式: int64) |
| `area` | number | 否 | 建筑面积(单位平米) (格式: float) |
| `areaCount` | integer | 否 | 项目下的区域统计(只读) (格式: int64) |
| `areas` | array[AreaInfo] | 否 |  |
| `attachments` | array[Attachment] | 否 | 附件 |
| `createdTime` | string | 否 | 创建时间（只读） |
| `desc` | string | 否 | 项目备注（读写） |
| `deviceCount` | integer | 否 | 项目下的设备统计(只读) (格式: int64) |
| `deviceOnlineCount` | integer | 否 | 项目下的设备在线统计(只读) (格式: int64) |
| `isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `ppsm` | integer | 否 | 每平方米功耗 (格式: int64) |
| `projectID` | string | 否 | 项目id（只读） |
| `projectImg` | string | 否 |  |
| `projectName` | string | 否 | 项目名称（读写） |
| `sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `status` | integer | 否 | 项目状态  1 正常，2-禁用，3-过期 (格式: int64) |
| `tags` | object | 否 | 自定义标签 |
| `tenantCode` | string | 否 | 租户编码（只读） |
| `type` | string | 否 | 项目类型 |
| `userCount` | integer | 否 | 项目下的用户统计(只读) (格式: int64) |

**请求示例**:
```json
{
  "address": "string",
  "adminUser": {
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
  "adminUserID": "string",
  "alarmStatus": 1,
  "area": 1,
  "areaCount": 1,
  "areas": [
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
  "attachments": [
    {
      "fileName": "string",
      "filePath": "string",
      "fileUrl": "string",
      "id": "string",
      "useBy": "string"
    }
  ],
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "deviceCount": 1,
  "deviceOnlineCount": 1,
  "isSysCreated": 1,
  "position": {
    "latitude": 1,
    "longitude": 1
  },
  "ppsm": 1,
  "projectID": "string",
  "projectImg": "string",
  "projectName": "示例名称",
  "sort": 1,
  "status": 1,
  "tags": {},
  "tenantCode": "string",
  "type": "string",
  "userCount": 1
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
ur api /api/v1/things/project/info/update \
  --body '{"address": "string", "adminUser": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "adminUserID": "string", "alarmStatus": 1, "area": 1, "areaCount": 1, "areas": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "attachments": [{"fileName": "string", "filePath": "string", "fileUrl": "string", "id": "string", "useBy": "string"}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "deviceOnlineCount": 1, "isSysCreated": 1, "position": {"latitude": 1, "longitude": 1}, "ppsm": 1, "projectID": "string", "projectImg": "string", "projectName": "示例名称", "sort": 1, "status": 1, "tags": {}, "tenantCode": "string", "type": "string", "userCount": 1}'
```

### POST `/api/v1/things/project/profile/get-list`

**说明**: 获取项目配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `codes` | array[string] | 是 |  |

**请求示例**:
```json
{
  "codes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "profiles": [
      {
        "code": "string",
        "params": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/profile/get-list \
  --body '{"codes": ["string"]}'
```

### POST `/api/v1/things/project/profile/get-one`

**说明**: 获取项目配置详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "params": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/profile/get-one \
  --body '{"code": "string"}'
```

### POST `/api/v1/things/project/profile/update`

**说明**: 更新项目配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `params` | string | 是 |  |

**请求示例**:
```json
{
  "code": "string",
  "params": "string"
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
ur api /api/v1/things/project/profile/update \
  --body '{"code": "string", "params": "string"}'
```
