# ur-device things/device/info

批量绑定 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/info/batch-bind` | 批量绑定 | admin |
| POST | `/api/v1/things/device/info/batch-import` | 批量导入设备 | admin |
| POST | `/api/v1/things/device/info/batch-update-import` | 导入批量更新设备 | admin |
| POST | `/api/v1/things/device/info/bind` | 绑定 | admin |
| POST | `/api/v1/things/device/info/bind/token/create` | 创建绑定token | admin |
| POST | `/api/v1/things/device/info/bind/token/get-one` | 绑定token状态查询 | admin |
| POST | `/api/v1/things/device/info/can-bind` | 是否可以绑定设备 | admin |
| POST | `/api/v1/things/device/info/count` | 设备统计详情 | admin |
| POST | `/api/v1/things/device/info/delete` | 删除设备 | admin |
| POST | `/api/v1/things/device/info/get-list` | 获取设备列表 | admin |
| POST | `/api/v1/things/device/info/get-one` | 获取设备详情 | admin |
| POST | `/api/v1/things/device/info/move` | 转移设备到新设备上 | admin |
| POST | `/api/v1/things/device/info/ota/upgrade` | 设备升级,获取升级包手动升级 | admin |
| POST | `/api/v1/things/device/info/transfer` | 转让设备 | admin |
| POST | `/api/v1/things/device/info/unbind` | 解绑设备 | admin |

## 详细说明

### POST `/api/v1/things/device/info/batch-bind`

**说明**: 批量绑定

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 |  |
| `devices` | array[DeviceCore] | 是 |  |

**请求示例**:
```json
{
  "areaID": "string",
  "devices": [
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
  "data": {
    "errs": [
      {
        "code": 1,
        "deviceName": "示例名称",
        "msg": "string",
        "productID": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/batch-bind \
  --body '{"areaID": "string", "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}]}'
```

### POST `/api/v1/things/device/info/batch-import`

**说明**: 批量导入设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `file` | string | 是 | 格式: binary |

**请求示例**:
```json
{
  "file": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "errdata": [
      {
        "address": "string",
        "deviceAlias": "示例名称",
        "deviceName": "示例名称",
        "gatewayDeviceName": "string",
        "gatewayProductID": "string",
        "logLevel": "string",
        "position": "string",
        "productID": "string",
        "row": 1,
        "secret": "string",
        "tags": "string",
        "tips": "string"
      }
    ],
    "headers": {
      "address": "string",
      "deviceAlias": "示例名称",
      "deviceName": "示例名称",
      "gatewayDeviceName": "string",
      "gatewayProductID": "string",
      "logLevel": "string",
      "position": "string",
      "productID": "string",
      "row": 1,
      "secret": "string",
      "tags": "string",
      "tips": "string"
    },
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/batch-import \
  --body '{"file": "string"}'
```

### POST `/api/v1/things/device/info/batch-update-import`

**说明**: 导入批量更新设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `file` | string | 是 | 格式: binary |

**请求示例**:
```json
{
  "file": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "errCount": 1,
    "errDetail": [
      {
        "code": 1,
        "device": {
          "deviceName": "示例名称",
          "productID": "string",
          "productName": "string"
        },
        "msg": "string"
      }
    ],
    "ignoreCount": 1,
    "succCount": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/batch-update-import \
  --body '{"file": "string"}'
```

### POST `/api/v1/things/device/info/bind`

**说明**: 绑定

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 |  |
| `device` | object | 是 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `device.productName` | string | 否 |  |
| `expand` | object | 否 | 拓展字段 |
| `token` | string | 否 | 配网token,中绑定需要 |

**请求示例**:
```json
{
  "areaID": "string",
  "device": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  },
  "expand": {},
  "token": "string"
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
ur api /api/v1/things/device/info/bind \
  --body '{"areaID": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "expand": {}, "token": "string"}'
```

### POST `/api/v1/things/device/info/bind/token/create`

**说明**: 创建绑定token

**权限**: admin

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "status": 1,
    "token": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/bind/token/create \
  --body '{}'
```

### POST `/api/v1/things/device/info/bind/token/get-one`

**说明**: 绑定token状态查询

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `token` | string | 是 | 配网token |

**请求示例**:
```json
{
  "token": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "status": 1,
    "token": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/bind/token/get-one \
  --body '{"token": "string"}'
```

### POST `/api/v1/things/device/info/can-bind`

**说明**: 是否可以绑定设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `device` | object | 是 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `device.productName` | string | 否 |  |

**请求示例**:
```json
{
  "device": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
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
ur api /api/v1/things/device/info/can-bind \
  --body '{"device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}}'
```

### POST `/api/v1/things/device/info/count`

**说明**: 设备统计详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `countTypes` | array[string] | 是 | 总数:total,在线设备数:status,设备类型: deviceType |
| `rangeIDs` | array[string] | 是 | 统计的id列表, areaID或者groupID |
| `rangeType` | string | 是 | 0:全部 1:area 2:group |

**请求示例**:
```json
{
  "countTypes": [
    "string"
  ],
  "rangeIDs": [
    "string"
  ],
  "rangeType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "count": {},
        "rangeID": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/count \
  --body '{"countTypes": ["string"], "rangeIDs": ["string"], "rangeType": "string"}'
```

### POST `/api/v1/things/device/info/delete`

**说明**: 删除设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名称 读写 |
| `productID` | string | 是 | 产品id 只读 |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "productID": "string"
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
ur api /api/v1/things/device/info/delete \
  --body '{"deviceName": "示例名称", "productID": "string"}'
```

### POST `/api/v1/things/device/info/get-list`

**说明**: 获取设备列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域ids |
| `areaIDPath` | string | 否 | 区域路径过滤 |
| `areaIDPaths` | array[string] | 否 | 区域路径过滤 |
| `areaIDs` | array[string] | 否 | 项目区域ids |
| `deviceAlias` | string | 否 | 过滤条件:模糊查询 设备别名 |
| `deviceName` | string | 否 | 过滤条件:模糊查询 设备名 |
| `deviceNameOrAlias` | string | 否 | 过滤条件:模糊查询 设备名或别名 |
| `deviceNames` | array[string] | 否 | 过滤条件:精准查询 设备名 |
| `deviceTypes` | array[integer] | 否 | 设备类型:1:设备,2:网关,3:子设备设备类型:1:设备,2:网关,3:子设备 |
| `devices` | array[DeviceCore] | 否 |  |
| `expTime` | object | 否 |  |
| `expTime.cmpType` | string | 是 | "=":相等 "!=":不相等 ">":大于">=":大于等于"<":小于"<=":小于等于 "like":模糊查询 |
| `expTime.value` | string | 是 |  |
| `gateway` | object | 否 |  |
| `gateway.deviceName` | string | 是 | 设备名称 |
| `gateway.productID` | string | 是 | 产品ID |
| `gateway.productName` | string | 否 |  |
| `groupID` | string | 否 |  |
| `groupIDPath` | string | 否 |  |
| `groupIDPaths` | array[string] | 否 |  |
| `groupIDs` | array[string] | 否 |  |
| `groupName` | string | 否 |  |
| `groupPurpose` | string | 否 | 设备分组用途 不填默认为default |
| `hasOwner` | integer | 否 | 是否被人拥有,1为是 2为否 (格式: int64) |
| `iccid` | string | 否 | SIM卡卡号 |
| `isOnline` | integer | 否 |  在线状态过滤  1离线 2在线 (格式: int64) |
| `isOnlyCore` | boolean | 否 | 只返回核心信息 (格式: boolean) |
| `netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网 (格式: int64) |
| `notAreaID` | string | 否 |  |
| `notGroupID` | string | 否 |  |
| `notVersion` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `parentGroupID` | string | 否 |  |
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `productCategoryCode` | string | 否 | 品类编码 |
| `productCategoryID` | string | 否 |  |
| `productCategoryIDs` | array[integer] | 否 |  |
| `productID` | string | 否 | 产品id 为空时获取所有产品 |
| `productIDs` | array[string] | 否 | 产品id 为空时获取所有产品 |
| `property` | object | 否 | 设备最新属性过滤,key为属性的id,如果是结构体则key为 aaa.bbb 数组为aaa.1 |
| `propertyIgnoreEmpty` | boolean | 否 |  设备属性是否忽略空值,空值是否返回 (格式: boolean) |
| `range` | integer | 否 | 过滤条件:距离坐标点固定范围内的设备 单位：米 (格式: int64) |
| `ratedPower` | object | 否 |  |
| `ratedPower.cmpType` | string | 是 | "=":相等 "!=":不相等 ">":大于">=":大于等于"<":小于"<=":小于等于 "like":模糊查询 |
| `ratedPower.value` | string | 是 |  |
| `rssi` | object | 否 |  |
| `rssi.cmpType` | string | 是 | "=":相等 "!=":不相等 ">":大于">=":大于等于"<":小于"<=":小于等于 "like":模糊查询 |
| `rssi.value` | string | 是 |  |
| `status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `statuses` | array[integer] | 否 |  |
| `tags` | object | 否 |  key tag过滤查询,非模糊查询 为tag的名,value为tag对应的值 |
| `tenantCode` | string | 否 |  |
| `userID` | string | 否 |  |
| `versions` | array[string] | 否 |  |
| `withArea` | boolean | 否 | 同时返回区域信息 (格式: boolean) |
| `withCollect` | integer | 否 |  过滤收藏的设备(这里只获取收藏的设备) 1: 同时获取收藏的设备 2:只获取收藏的设备 (格式: int64) |
| `withGateway` | boolean | 否 | 同时返回子设备绑定的网关信息 (格式: boolean) |
| `withGroups` | array[string] | 否 |  |
| `withOwner` | boolean | 否 | 同时获取拥有人的信息 (格式: boolean) |
| `withProfiles` | array[string] | 否 |  |
| `withProperties` | array[string] | 否 | 如果不为nil,如果为空,获取设备所有最新属性 如果传了属性列表,则会返回属性列表,如果没有匹配的则不会返回 |
| `withShared` | integer | 否 |  过滤分享的设备1: 同时获取分享的设备 2:只获取分享的设备 (格式: int64) |

**请求示例**:
```json
{
  "areaID": "string",
  "areaIDPath": "string",
  "areaIDPaths": [
    "string"
  ],
  "areaIDs": [
    "string"
  ],
  "deviceAlias": "示例名称",
  "deviceName": "string",
  "deviceNameOrAlias": "示例名称",
  "deviceNames": [
    "string"
  ],
  "deviceTypes": [
    1
  ],
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
    }
  ],
  "expTime": {
    "cmpType": "string",
    "value": "string"
  },
  "gateway": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  },
  "groupID": "string",
  "groupIDPath": "string",
  "groupIDPaths": [
    "string"
  ],
  "groupIDs": [
    "string"
  ],
  "groupName": "string",
  "groupPurpose": "string",
  "hasOwner": 1,
  "iccid": "string",
  "isOnline": 1,
  "isOnlyCore": true,
  "netType": 1,
  "notAreaID": "string",
  "notGroupID": "string",
  "notVersion": "string",
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
  "parentGroupID": "string",
  "position": {
    "latitude": 1,
    "longitude": 1
  },
  "productCategoryCode": "string",
  "productCategoryID": "string",
  "productCategoryIDs": [
    1
  ],
  "productID": "string",
  "productIDs": [
    "string"
  ],
  "property": {},
  "propertyIgnoreEmpty": true,
  "range": 1,
  "ratedPower": {
    "cmpType": "string",
    "value": "string"
  },
  "rssi": {
    "cmpType": "string",
    "value": "string"
  },
  "status": 1,
  "statuses": [
    1
  ],
  "tags": {},
  "tenantCode": "string",
  "userID": "string",
  "versions": [
    "string"
  ],
  "withArea": true,
  "withCollect": 1,
  "withGateway": true,
  "withGroups": [
    "string"
  ],
  "withOwner": true,
  "withProfiles": [
    "string"
  ],
  "withProperties": [
    "string"
  ],
  "withShared": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
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
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/get-list \
  --body '{"areaID": "string", "areaIDPath": "string", "areaIDPaths": ["string"], "areaIDs": ["string"], "deviceAlias": "示例名称", "deviceName": "string", "deviceNameOrAlias": "示例名称", "deviceNames": ["string"], "deviceTypes": [1], "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "expTime": {"cmpType": "string", "value": "string"}, "gateway": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "groupID": "string", "groupIDPath": "string", "groupIDPaths": ["string"], "groupIDs": ["string"], "groupName": "string", "groupPurpose": "string", "hasOwner": 1, "iccid": "string", "isOnline": 1, "isOnlyCore": true, "netType": 1, "notAreaID": "string", "notGroupID": "string", "notVersion": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentGroupID": "string", "position": {"latitude": 1, "longitude": 1}, "productCategoryCode": "string", "productCategoryID": "string", "productCategoryIDs": [1], "productID": "string", "productIDs": ["string"], "property": {}, "propertyIgnoreEmpty": true, "range": 1, "ratedPower": {"cmpType": "string", "value": "string"}, "rssi": {"cmpType": "string", "value": "string"}, "status": 1, "statuses": [1], "tags": {}, "tenantCode": "string", "userID": "string", "versions": ["string"], "withArea": true, "withCollect": 1, "withGateway": true, "withGroups": ["string"], "withOwner": true, "withProfiles": ["string"], "withProperties": ["string"], "withShared": 1}'
```

### POST `/api/v1/things/device/info/get-one`

**说明**: 获取设备详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名称 读写 |
| `isOnlyCore` | boolean | 否 | 只返回核心信息 (格式: boolean) |
| `productID` | string | 否 | 产品id 为空时获取所有产品 |
| `propertyIgnoreEmpty` | boolean | 否 |  设备属性是否忽略空值,空值是否返回 (格式: boolean) |
| `withArea` | boolean | 否 | 同时返回区域信息 (格式: boolean) |
| `withGateway` | boolean | 否 | 同时返回子设备绑定的网关信息 (格式: boolean) |
| `withOwner` | boolean | 否 | 同时获取拥有人的信息 (格式: boolean) |
| `withProfiles` | array[string] | 否 |  |
| `withProperties` | array[string] | 否 | 如果不为nil,如果为空,获取设备所有最新属性 如果传了属性列表,则会返回属性列表 |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "isOnlyCore": true,
  "productID": "string",
  "propertyIgnoreEmpty": true,
  "withArea": true,
  "withGateway": true,
  "withOwner": true,
  "withProfiles": [
    "string"
  ],
  "withProperties": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
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
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/get-one \
  --body '{"deviceName": "示例名称", "isOnlyCore": true, "productID": "string", "propertyIgnoreEmpty": true, "withArea": true, "withGateway": true, "withOwner": true, "withProfiles": ["string"], "withProperties": ["string"]}'
```

### POST `/api/v1/things/device/info/move`

**说明**: 转移设备到新设备上

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `new` | object | 是 |  |
| `new.deviceName` | string | 是 | 设备名称 |
| `new.productID` | string | 是 | 产品ID |
| `new.productName` | string | 否 |  |
| `old` | object | 是 |  |
| `old.deviceName` | string | 是 | 设备名称 |
| `old.productID` | string | 是 | 产品ID |
| `old.productName` | string | 否 |  |
| `withTarget` | array[string] | 否 | 设备配置-profile,schemaAlias:profile |

**请求示例**:
```json
{
  "new": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  },
  "old": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  },
  "withTarget": [
    "string"
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
ur api /api/v1/things/device/info/move \
  --body '{"new": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "old": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "withTarget": ["string"]}'
```

### POST `/api/v1/things/device/info/ota/upgrade`

**说明**: 设备升级,获取升级包手动升级

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  设备名称 |
| `productID` | string | 是 |  产品ID |
| `version` | string | 是 |  当前版本,如果和设备的当前版本不一致,会更新设备的版本,如果和固件的版本一致,这个固件的升级进度会变成成功 |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "productID": "string",
  "version": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "firmware": {
      "desc": "string",
      "extra": "string",
      "files": [
        {
          "fileMd5": "string",
          "fileUrl": "string",
          "name": "string",
          "signature": "string",
          "size": 1
        }
      ],
      "isDiff": 1,
      "name": "示例名称",
      "signMethod": "string",
      "version": "string"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/info/ota/upgrade \
  --body '{"deviceName": "示例名称", "productID": "string", "version": "string"}'
```

### POST `/api/v1/things/device/info/transfer`

**说明**: 转让设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 转移到项目时指定区域ID |
| `device` | object | 否 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `device.productName` | string | 否 |  |
| `devices` | array[DeviceCore] | 否 |  |
| `isCleanData` | integer | 是 | 是否清除数据:1是 2否 (格式: int64) |
| `projectID` | string | 否 |  |
| `srcProjectID` | string | 否 | 指定原来的项目ID,不写使用头 |
| `transferTo` | integer | 是 | 转让给: 1: 某个人 2: 自己的某个项目 (格式: int64) |
| `userID` | string | 否 |  |

**请求示例**:
```json
{
  "areaID": "string",
  "device": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  },
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
    }
  ],
  "isCleanData": 1,
  "projectID": "string",
  "srcProjectID": "string",
  "transferTo": 1,
  "userID": "string"
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
ur api /api/v1/things/device/info/transfer \
  --body '{"areaID": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "isCleanData": 1, "projectID": "string", "srcProjectID": "string", "transferTo": 1, "userID": "string"}'
```

### POST `/api/v1/things/device/info/unbind`

**说明**: 解绑设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名称 读写 |
| `nonce` | integer | 否 | 随机数 (格式: int64) |
| `productID` | string | 是 | 产品id 只读 |
| `secretType` | string | 否 | 采用秘钥的类型 |
| `signType` | string | 否 | hmacsha256 hmacsha1 hmacmd5(默认) |
| `signature` | string | 否 | 签名 |
| `timestamp` | integer | 否 | 秒时间戳 (格式: int64) |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "nonce": 1,
  "productID": "string",
  "secretType": "string",
  "signType": "string",
  "signature": "string",
  "timestamp": 1
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
ur api /api/v1/things/device/info/unbind \
  --body '{"deviceName": "示例名称", "nonce": 1, "productID": "string", "secretType": "string", "signType": "string", "signature": "string", "timestamp": 1}'
```
