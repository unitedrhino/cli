# ur-project things/project/info

新增项目 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/project/info/create` | 新增项目 | admin |
| POST | `/api/v1/things/project/info/delete` | 删除项目 | admin |
| POST | `/api/v1/things/project/info/get-list` | 获取项目列表 | all |
| POST | `/api/v1/things/project/info/get-one` | 获取项目详情 | all |
| POST | `/api/v1/things/project/info/update` | 更新项目 | admin |

## 详细说明

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
