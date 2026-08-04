# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/user/area/apply/create` | 申请用户区域权限 | all |
| POST | `/api/v1/things/user/device/collect/batch-create` | 批量收藏设备 | all |
| POST | `/api/v1/things/user/device/collect/batch-delete` | 批量取消收藏设备 | all |
| POST | `/api/v1/things/user/device/collect/get-list` | 获取收藏设备列表 | all |
| POST | `/api/v1/things/user/device/share/batch-accept` | 接受批量分享设备 | all |
| POST | `/api/v1/things/user/device/share/batch-create` | 生成批量分享设备二维码 | all |
| POST | `/api/v1/things/user/device/share/batch-delete` | 批量取消分享设备 | all |
| POST | `/api/v1/things/user/device/share/batch-delete-token` | 删除批量分享 Token | all |
| POST | `/api/v1/things/user/device/share/batch-get-list` | 获取批量分享的设备列表 | all |
| POST | `/api/v1/things/user/device/share/batch-get-token-list` | 获取批量分享 Token 列表 | all |
| POST | `/api/v1/things/user/device/share/create` | 分享设备 | all |
| POST | `/api/v1/things/user/device/share/delete` | 取消分享设备 | all |
| POST | `/api/v1/things/user/device/share/get-list` | 获取分享设备列表 | all |
| POST | `/api/v1/things/user/device/share/get-one` | 获取分享设备详情 | all |
| POST | `/api/v1/things/user/device/share/share-code-accept` | 通过分享码接受设备 | all |
| POST | `/api/v1/things/user/device/share/share-code-delete` | 删除分享码 | all |
| POST | `/api/v1/things/user/device/share/share-code-gen` | 生成设备分享码 | all |
| POST | `/api/v1/things/user/device/share/share-code-get-info` | 通过分享码查询分享信息 | public |
| POST | `/api/v1/things/user/device/share/share-code-get-list` | 获取分享码列表 | all |
| POST | `/api/v1/things/user/device/share/update` | 更新分享设备信息 | all |

## 端点详情

### POST `/api/v1/things/user/area/apply/create`

**说明**: 申请用户区域权限

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | integer | 是 |  区域ID (格式: int64) |
| `authType` | integer | 是 |  权限类型 1:读权限 4:管理权限 (格式: int64) |

**请求示例**:
```json
{
  "areaID": 1,
  "authType": 1
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
ur api /api/v1/things/user/area/apply/create \
  --body '{"areaID": 1, "authType": 1}'
```

### POST `/api/v1/things/user/device/collect/batch-create`

**说明**: 批量收藏设备

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `devices` | array[DeviceCore] | 是 |  |

**请求示例**:
```json
{
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string"
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
ur api /api/v1/things/user/device/collect/batch-create \
  --body '{"devices": [{"deviceName": "示例名称", "productID": "string"}]}'
```

### POST `/api/v1/things/user/device/collect/batch-delete`

**说明**: 批量取消收藏设备

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `devices` | array[DeviceCore] | 是 |  |

**请求示例**:
```json
{
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string"
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
ur api /api/v1/things/user/device/collect/batch-delete \
  --body '{"devices": [{"deviceName": "示例名称", "productID": "string"}]}'
```

### POST `/api/v1/things/user/device/collect/get-list`

**说明**: 获取收藏设备列表

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
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
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/collect/get-list \
  --body '{}'
```

### POST `/api/v1/things/user/device/share/batch-accept`

**说明**: 接受批量分享设备

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `devices` | array[DeviceCore] | 否 | 批量设备信息 |
| `shareToken` | string | 否 | 分享后的用于生成二维码的字符串 |

**请求示例**:
```json
{
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string"
    }
  ],
  "shareToken": "string"
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
ur api /api/v1/things/user/device/share/batch-accept \
  --body '{"devices": [{"deviceName": "示例名称", "productID": "string"}], "shareToken": "string"}'
```

### POST `/api/v1/things/user/device/share/batch-create`

**说明**: 生成批量分享设备二维码

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessPerm` | object | 否 | 系统功能权限 2:读写权限 3读权限 |
| `authType` | integer | 否 | 授权类型:1:全部授权 2:部分授权 (格式: int64) |
| `createdTime` | string | 否 |  |
| `devices` | array[DeviceCore] | 否 | 批量设备信息 |
| `expTime` | integer | 否 | 到期时间 (格式: int64) |
| `schemaPerm` | object | 否 | 普通功能权限 2:读写权限 3读权限 |
| `useBy` | string | 否 | 用途 |

**请求示例**:
```json
{
  "accessPerm": {},
  "authType": 1,
  "createdTime": "string",
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string"
    }
  ],
  "expTime": 1,
  "schemaPerm": {},
  "useBy": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "shareToken": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/batch-create \
  --body '{"accessPerm": {}, "authType": 1, "createdTime": "string", "devices": [{"deviceName": "示例名称", "productID": "string"}], "expTime": 1, "schemaPerm": {}, "useBy": "string"}'
```

### POST `/api/v1/things/user/device/share/batch-delete`

**说明**: 批量取消分享设备

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[integer] | 否 |  |
| `projectID` | string | 否 |  |

**请求示例**:
```json
{
  "ids": [
    1
  ],
  "projectID": "string"
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
ur api /api/v1/things/user/device/share/batch-delete \
  --body '{"ids": [1], "projectID": "string"}'
```

### POST `/api/v1/things/user/device/share/batch-delete-token`

**说明**: 删除批量分享 Token

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `shareToken` | string | 是 | 要删除的分享 Token |

**请求示例**:
```json
{
  "shareToken": "string"
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
ur api /api/v1/things/user/device/share/batch-delete-token \
  --body '{"shareToken": "string"}'
```

### POST `/api/v1/things/user/device/share/batch-get-list`

**说明**: 获取批量分享的设备列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `shareToken` | string | 否 |  |

**请求示例**:
```json
{
  "shareToken": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "accessPerm": {},
    "authType": 1,
    "createdTime": "string",
    "devices": [
      {
        "deviceAlias": "示例名称",
        "deviceName": "示例名称",
        "productID": "string",
        "productImg": "string",
        "productName": "string",
        "useBy": "string"
      }
    ],
    "expTime": 1,
    "schemaPerm": {},
    "useBy": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/batch-get-list \
  --body '{"shareToken": "string"}'
```

### POST `/api/v1/things/user/device/share/batch-get-token-list`

**说明**: 获取批量分享 Token 列表

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "authType": 1,
        "createdTime": "2026-01-01T00:00:00Z",
        "deviceCount": "string",
        "expTime": 1,
        "shareToken": "string",
        "useBy": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/batch-get-token-list \
  --body '{}'
```

### POST `/api/v1/things/user/device/share/create`

**说明**: 分享设备

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessPerm` | object | 否 | 系统功能权限 2:读写权限 3读权限 |
| `authType` | integer | 否 | 授权类型:1:全部授权 2:部分授权 (格式: int64) |
| `createdTime` | string | 否 |  |
| `device` | object | 否 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `expTime` | string | 否 | 到期时间 |
| `id` | string | 否 |  |
| `projectID` | string | 否 |  |
| `schemaPerm` | object | 否 | 普通功能权限 2:读写权限 3读权限 |
| `sharedUserAccount` | string | 否 |  |
| `sharedUserID` | string | 否 | 分享的对象 |
| `useBy` | string | 否 | 用途 |
| `user` | object | 否 |  |
| `user.avatar` | string | 否 |  用户头像 |
| `user.createdTime` | string | 否 |  创建时间 |
| `user.email` | string | 否 |  邮箱 |
| `user.lastIP` | string | 否 |  最后登录ip |
| `user.nickName` | string | 否 |  用户的昵称 |
| `user.phone` | string | 否 |  手机号 |
| `user.regIP` | string | 否 |  注册ip |
| `user.userID` | string | 否 |  用户id |
| `user.userName` | string | 否 |  用户名(唯一) |

**请求示例**:
```json
{
  "accessPerm": {},
  "authType": 1,
  "createdTime": "string",
  "device": {
    "deviceName": "示例名称",
    "productID": "string"
  },
  "expTime": "2026-01-01T00:00:00Z",
  "id": "string",
  "projectID": "string",
  "schemaPerm": {},
  "sharedUserAccount": "string",
  "sharedUserID": "string",
  "useBy": "string",
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
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/create \
  --body '{"accessPerm": {}, "authType": 1, "createdTime": "string", "device": {"deviceName": "示例名称", "productID": "string"}, "expTime": "2026-01-01T00:00:00Z", "id": "string", "projectID": "string", "schemaPerm": {}, "sharedUserAccount": "string", "sharedUserID": "string", "useBy": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}}'
```

### POST `/api/v1/things/user/device/share/delete`

**说明**: 取消分享设备

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `device` | object | 否 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `id` | string | 否 |  |
| `useBy` | string | 否 | 用途 |

**请求示例**:
```json
{
  "device": {
    "deviceName": "示例名称",
    "productID": "string"
  },
  "id": "string",
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
ur api /api/v1/things/user/device/share/delete \
  --body '{"device": {"deviceName": "示例名称", "productID": "string"}, "id": "string", "useBy": "string"}'
```

### POST `/api/v1/things/user/device/share/get-list`

**说明**: 获取分享设备列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `device` | object | 否 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `withUser` | boolean | 否 | 同时返回用户信息 (格式: boolean) |

**请求示例**:
```json
{
  "device": {
    "deviceName": "示例名称",
    "productID": "string"
  },
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
  "withUser": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "accessPerm": {},
        "authType": 1,
        "createdTime": "string",
        "device": {
          "deviceName": "示例名称",
          "productID": "string"
        },
        "expTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "projectID": "string",
        "schemaPerm": {},
        "sharedUserAccount": "string",
        "sharedUserID": "string",
        "useBy": "string",
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
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/get-list \
  --body '{"device": {"deviceName": "示例名称", "productID": "string"}, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "withUser": true}'
```

### POST `/api/v1/things/user/device/share/get-one`

**说明**: 获取分享设备详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `device` | object | 否 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `id` | string | 否 |  |
| `useBy` | string | 否 | 用途 |

**请求示例**:
```json
{
  "device": {
    "deviceName": "示例名称",
    "productID": "string"
  },
  "id": "string",
  "useBy": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "accessPerm": {},
    "authType": 1,
    "createdTime": "string",
    "device": {
      "deviceName": "示例名称",
      "productID": "string"
    },
    "expTime": "2026-01-01T00:00:00Z",
    "id": "string",
    "projectID": "string",
    "schemaPerm": {},
    "sharedUserAccount": "string",
    "sharedUserID": "string",
    "useBy": "string",
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
ur api /api/v1/things/user/device/share/get-one \
  --body '{"device": {"deviceName": "示例名称", "productID": "string"}, "id": "string", "useBy": "string"}'
```

### POST `/api/v1/things/user/device/share/share-code-accept`

**说明**: 通过分享码接受设备

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 | 分享码 |

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
    "devices": [
      {
        "deviceAlias": "示例名称",
        "deviceName": "示例名称",
        "productID": "string",
        "productImg": "string",
        "productName": "string",
        "useBy": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/share-code-accept \
  --body '{"code": "string"}'
```

### POST `/api/v1/things/user/device/share/share-code-delete`

**说明**: 删除分享码

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 | 要删除的分享码 |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/share-code-delete \
  --body '{"code": "string"}'
```

### POST `/api/v1/things/user/device/share/share-code-gen`

**说明**: 生成设备分享码

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessPerm` | object | 否 | 系统功能权限 |
| `authType` | integer | 否 | 授权类型:1全部授权 2部分授权 (格式: int64) |
| `devices` | array[DeviceCore] | 是 | 要分享的设备列表（至少1个） |
| `expTime` | string | 否 | 分享过期时间（0表示不限制） |
| `schemaPerm` | object | 否 | 物模型权限 |
| `useBy` | string | 否 | 用途 |
| `validHours` | integer | 否 | 分享码有效时长（小时），默认24，最大72（3天） (格式: int64) |

**请求示例**:
```json
{
  "accessPerm": {},
  "authType": 1,
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string"
    }
  ],
  "expTime": "2026-01-01T00:00:00Z",
  "schemaPerm": {},
  "useBy": "string",
  "validHours": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "deviceCount": "string",
    "validUntilTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/share-code-gen \
  --body '{"accessPerm": {}, "authType": 1, "devices": [{"deviceName": "示例名称", "productID": "string"}], "expTime": "2026-01-01T00:00:00Z", "schemaPerm": {}, "useBy": "string", "validHours": 1}'
```

### POST `/api/v1/things/user/device/share/share-code-get-info`

**说明**: 通过分享码查询分享信息

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 | 分享码 |

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
    "accessPerm": {},
    "authType": 1,
    "deviceCount": "string",
    "devices": [
      {
        "deviceAlias": "示例名称",
        "deviceName": "示例名称",
        "productID": "string",
        "productImg": "string",
        "productName": "string",
        "useBy": "string"
      }
    ],
    "expTime": 1,
    "schemaPerm": {},
    "validUntilTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/share-code-get-info \
  --body '{"code": "string"}'
```

### POST `/api/v1/things/user/device/share/share-code-get-list`

**说明**: 获取分享码列表

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "accessPerm": {},
        "authType": 1,
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "deviceCount": "string",
        "devices": [
          {
            "deviceAlias": "示例名称",
            "deviceName": "示例名称",
            "productID": "string",
            "productImg": "string",
            "productName": "string",
            "useBy": "string"
          }
        ],
        "expTime": 1,
        "schemaPerm": {},
        "useBy": "string",
        "validUntilTime": "2026-01-01T00:00:00Z"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/user/device/share/share-code-get-list \
  --body '{}'
```

### POST `/api/v1/things/user/device/share/update`

**说明**: 更新分享设备信息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessPerm` | object | 否 | 系统功能权限 2:读写权限 3读权限 |
| `authType` | integer | 否 | 授权类型:1:全部授权 2:部分授权 (格式: int64) |
| `createdTime` | string | 否 |  |
| `device` | object | 否 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `expTime` | string | 否 | 到期时间 |
| `id` | string | 否 |  |
| `projectID` | string | 否 |  |
| `schemaPerm` | object | 否 | 普通功能权限 2:读写权限 3读权限 |
| `sharedUserAccount` | string | 否 |  |
| `sharedUserID` | string | 否 | 分享的对象 |
| `useBy` | string | 否 | 用途 |
| `user` | object | 否 |  |
| `user.avatar` | string | 否 |  用户头像 |
| `user.createdTime` | string | 否 |  创建时间 |
| `user.email` | string | 否 |  邮箱 |
| `user.lastIP` | string | 否 |  最后登录ip |
| `user.nickName` | string | 否 |  用户的昵称 |
| `user.phone` | string | 否 |  手机号 |
| `user.regIP` | string | 否 |  注册ip |
| `user.userID` | string | 否 |  用户id |
| `user.userName` | string | 否 |  用户名(唯一) |

**请求示例**:
```json
{
  "accessPerm": {},
  "authType": 1,
  "createdTime": "string",
  "device": {
    "deviceName": "示例名称",
    "productID": "string"
  },
  "expTime": "2026-01-01T00:00:00Z",
  "id": "string",
  "projectID": "string",
  "schemaPerm": {},
  "sharedUserAccount": "string",
  "sharedUserID": "string",
  "useBy": "string",
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
ur api /api/v1/things/user/device/share/update \
  --body '{"accessPerm": {}, "authType": 1, "createdTime": "string", "device": {"deviceName": "示例名称", "productID": "string"}, "expTime": "2026-01-01T00:00:00Z", "id": "string", "projectID": "string", "schemaPerm": {}, "sharedUserAccount": "string", "sharedUserID": "string", "useBy": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}}'
```
