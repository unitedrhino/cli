# ur-device things/device/bind-token

新增设备 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/info/create` | 新增设备 | admin |
| POST | `/api/v1/things/device/info/update` | 更新设备 | admin |

## 详细说明

### POST `/api/v1/things/device/info/create`

**说明**: 新增设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `adcode` | string | 否 | 地区编码 |
| `address` | string | 否 | 所在地址 |
| `area` | object | 否 |  |
| `area.areaID` | string | 否 | 项目区域id（只读） |
| `area.areaIDPath` | string | 否 | 项目区域ids（只读） |
| `area.areaImg` | string | 否 |  |
| `area.areaName` | string | 否 | 项目区域名称（读写） |
| `area.areaNamePath` | string | 否 | （只读） |
| `area.children` | array[AreaInfo] | 否 | 下级项目区域列表（只读） |
| `area.configFile` | string | 否 | 区域配置文件 |
| `area.createdTime` | string | 否 | 创建时间（只读） |
| `area.desc` | string | 否 | 项目区域备注（读写） |
| `area.deviceCount` | integer | 否 | 设备数量统计(只读) (格式: int64) |
| `area.devices` | array[DeviceInfo] | 否 | 区域下设备列表（按需返回） |
| `area.groupCount` | integer | 否 | 分组数量统计(只读) (格式: int64) |
| `area.isLeaf` | integer | 否 | 是否是子节点（只读） (格式: int64) |
| `area.isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `area.lowerLevelCount` | integer | 否 | 下级的数量统计（只读） (格式: int64) |
| `area.parentAreaID` | string | 否 | 上级项目区域id（只读） |
| `area.position` | object | 否 |  |
| `area.position.latitude` | number | 是 | 纬度 (格式: double) |
| `area.position.longitude` | number | 是 | 经度 (格式: double) |
| `area.projectID` | string | 否 | 项目id（只读） |
| `area.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `area.tags` | object | 否 | 自定义标签 |
| `area.tenantCode` | string | 否 | 企业编码（只读） |
| `area.useBy` | string | 否 | 用途: commonSpace default |
| `areaID` | string | 否 | 项目区域id 只读 |
| `categoryID` | string | 否 | 产品品类 |
| `cert` | string | 否 |  设备证书  只读 |
| `cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `createdTime` | string | 否 | 创建时间 只读 |
| `desc` | string | 否 |  |
| `deviceAlias` | string | 否 | 设备别名 读写 |
| `deviceImg` | string | 否 | 设备图片 |
| `deviceName` | string | 是 | 设备名称 读写 |
| `deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `distributor` | object | 否 |  |
| `distributor.id` | string | 否 |  |
| `distributor.idPath` | string | 否 |  |
| `expTime` | string | 否 | 到期时间 |
| `file` | string | 否 | 设备相关文件 |
| `firstBind` | string | 否 |  |
| `firstLogin` | string | 否 | 激活时间 只读 |
| `gateway` | object | 否 |  |
| `gateway.adcode` | string | 否 | 地区编码 |
| `gateway.address` | string | 否 | 所在地址 |
| `gateway.area` | object | 否 |  |
| `gateway.area.areaID` | string | 否 | 项目区域id（只读） |
| `gateway.area.areaIDPath` | string | 否 | 项目区域ids（只读） |
| `gateway.area.areaImg` | string | 否 |  |
| `gateway.area.areaName` | string | 否 | 项目区域名称（读写） |
| `gateway.area.areaNamePath` | string | 否 | （只读） |
| `gateway.area.children` | array[AreaInfo] | 否 | 下级项目区域列表（只读） |
| `gateway.area.configFile` | string | 否 | 区域配置文件 |
| `gateway.area.createdTime` | string | 否 | 创建时间（只读） |
| `gateway.area.desc` | string | 否 | 项目区域备注（读写） |
| `gateway.area.deviceCount` | integer | 否 | 设备数量统计(只读) (格式: int64) |
| `gateway.area.devices` | array[DeviceInfo] | 否 | 区域下设备列表（按需返回） |
| `gateway.area.groupCount` | integer | 否 | 分组数量统计(只读) (格式: int64) |
| `gateway.area.isLeaf` | integer | 否 | 是否是子节点（只读） (格式: int64) |
| `gateway.area.isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `gateway.area.lowerLevelCount` | integer | 否 | 下级的数量统计（只读） (格式: int64) |
| `gateway.area.parentAreaID` | string | 否 | 上级项目区域id（只读） |
| `gateway.area.position` | object | 否 |  |
| `gateway.area.projectID` | string | 否 | 项目id（只读） |
| `gateway.area.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `gateway.area.tags` | object | 否 | 自定义标签 |
| `gateway.area.tenantCode` | string | 否 | 企业编码（只读） |
| `gateway.area.useBy` | string | 否 | 用途: commonSpace default |
| `gateway.areaID` | string | 否 | 项目区域id 只读 |
| `gateway.categoryID` | string | 否 | 产品品类 |
| `gateway.cert` | string | 否 |  设备证书  只读 |
| `gateway.cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `gateway.createdTime` | string | 否 | 创建时间 只读 |
| `gateway.desc` | string | 否 |  |
| `gateway.deviceAlias` | string | 否 | 设备别名 读写 |
| `gateway.deviceImg` | string | 否 | 设备图片 |
| `gateway.deviceName` | string | 是 | 设备名称 读写 |
| `gateway.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `gateway.distributor` | object | 否 |  |
| `gateway.distributor.id` | string | 否 |  |
| `gateway.distributor.idPath` | string | 否 |  |
| `gateway.expTime` | string | 否 | 到期时间 |
| `gateway.file` | string | 否 | 设备相关文件 |
| `gateway.firstBind` | string | 否 |  |
| `gateway.firstLogin` | string | 否 | 激活时间 只读 |
| `gateway.gateway` | object | 否 |  |
| `gateway.gateway.adcode` | string | 否 | 地区编码 |
| `gateway.gateway.address` | string | 否 | 所在地址 |
| `gateway.gateway.area` | object | 否 |  |
| `gateway.gateway.areaID` | string | 否 | 项目区域id 只读 |
| `gateway.gateway.categoryID` | string | 否 | 产品品类 |
| `gateway.gateway.cert` | string | 否 |  设备证书  只读 |
| `gateway.gateway.cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `gateway.gateway.createdTime` | string | 否 | 创建时间 只读 |
| `gateway.gateway.desc` | string | 否 |  |
| `gateway.gateway.deviceAlias` | string | 否 | 设备别名 读写 |
| `gateway.gateway.deviceImg` | string | 否 | 设备图片 |
| `gateway.gateway.deviceName` | string | 是 | 设备名称 读写 |
| `gateway.gateway.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `gateway.gateway.distributor` | object | 否 |  |
| `gateway.gateway.expTime` | string | 否 | 到期时间 |
| `gateway.gateway.file` | string | 否 | 设备相关文件 |
| `gateway.gateway.firstBind` | string | 否 |  |
| `gateway.gateway.firstLogin` | string | 否 | 激活时间 只读 |
| `gateway.gateway.gateway` | object | 否 |  |
| `gateway.gateway.groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `gateway.gateway.groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `gateway.gateway.groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `gateway.gateway.hardInfo` | string | 否 |  模组硬件型号 只读 |
| `gateway.gateway.iccid` | string | 否 | SIM卡卡号 |
| `gateway.gateway.id` | string | 否 |  |
| `gateway.gateway.imei` | string | 否 |  IMEI号信息 只读 |
| `gateway.gateway.isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `gateway.gateway.isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `gateway.gateway.lastBind` | string | 否 | 最后绑定时间 只读 |
| `gateway.gateway.lastIp` | string | 否 | 最后登录的ip地址 |
| `gateway.gateway.lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `gateway.gateway.lastLogin` | string | 否 | 最后上线时间 只读 |
| `gateway.gateway.lastOffline` | string | 否 | 最后下线时间 只读 |
| `gateway.gateway.logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `gateway.gateway.mac` | string | 否 |  MAC号信息 只读 |
| `gateway.gateway.mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `gateway.gateway.needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `gateway.gateway.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `gateway.gateway.owner` | object | 否 |  |
| `gateway.gateway.phone` | string | 否 | 手机号 |
| `gateway.gateway.position` | object | 否 |  |
| `gateway.gateway.productID` | string | 是 | 产品id 只读 |
| `gateway.gateway.productImg` | string | 否 | 产品图片 |
| `gateway.gateway.productName` | string | 否 |  |
| `gateway.gateway.profiles` | object | 否 |  |
| `gateway.gateway.projectID` | string | 否 | 项目id 只读 |
| `gateway.gateway.protocolConf` | object | 否 | 协议配置 |
| `gateway.gateway.ratedPower` | integer | 否 | 格式: int64 |
| `gateway.gateway.remoteAddr` | string | 否 | 远程访问地址 |
| `gateway.gateway.rssi` | integer | 否 | 信号强度 (格式: int64) |
| `gateway.gateway.schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `gateway.gateway.secret` | string | 否 | 设备秘钥 只读 |
| `gateway.gateway.softInfo` | string | 否 |  模组软件版本 只读 |
| `gateway.gateway.sort` | integer | 否 | 排序 (格式: int64) |
| `gateway.gateway.status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `gateway.gateway.subProtocolConf` | object | 否 | 子协议配置 |
| `gateway.gateway.tags` | object | 否 |  设备tag |
| `gateway.gateway.tenantCode` | string | 否 |  |
| `gateway.gateway.userID` | string | 否 |  |
| `gateway.gateway.version` | string | 否 |  固件版本  读写 |
| `gateway.gateway.withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |
| `gateway.groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `gateway.groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `gateway.groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `gateway.hardInfo` | string | 否 |  模组硬件型号 只读 |
| `gateway.iccid` | string | 否 | SIM卡卡号 |
| `gateway.id` | string | 否 |  |
| `gateway.imei` | string | 否 |  IMEI号信息 只读 |
| `gateway.isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `gateway.isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `gateway.lastBind` | string | 否 | 最后绑定时间 只读 |
| `gateway.lastIp` | string | 否 | 最后登录的ip地址 |
| `gateway.lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `gateway.lastLogin` | string | 否 | 最后上线时间 只读 |
| `gateway.lastOffline` | string | 否 | 最后下线时间 只读 |
| `gateway.logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `gateway.mac` | string | 否 |  MAC号信息 只读 |
| `gateway.mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `gateway.needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `gateway.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `gateway.owner` | object | 否 |  |
| `gateway.owner.avatar` | string | 否 |  用户头像 |
| `gateway.owner.createdTime` | string | 否 |  创建时间 |
| `gateway.owner.email` | string | 否 |  邮箱 |
| `gateway.owner.lastIP` | string | 否 |  最后登录ip |
| `gateway.owner.nickName` | string | 否 |  用户的昵称 |
| `gateway.owner.phone` | string | 否 |  手机号 |
| `gateway.owner.regIP` | string | 否 |  注册ip |
| `gateway.owner.userID` | string | 否 |  用户id |
| `gateway.owner.userName` | string | 否 |  用户名(唯一) |
| `gateway.phone` | string | 否 | 手机号 |
| `gateway.position` | object | 否 |  |
| `gateway.position.latitude` | number | 是 | 纬度 (格式: double) |
| `gateway.position.longitude` | number | 是 | 经度 (格式: double) |
| `gateway.productID` | string | 是 | 产品id 只读 |
| `gateway.productImg` | string | 否 | 产品图片 |
| `gateway.productName` | string | 否 |  |
| `gateway.profiles` | object | 否 |  |
| `gateway.projectID` | string | 否 | 项目id 只读 |
| `gateway.protocolConf` | object | 否 | 协议配置 |
| `gateway.ratedPower` | integer | 否 | 格式: int64 |
| `gateway.remoteAddr` | string | 否 | 远程访问地址 |
| `gateway.rssi` | integer | 否 | 信号强度 (格式: int64) |
| `gateway.schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `gateway.secret` | string | 否 | 设备秘钥 只读 |
| `gateway.softInfo` | string | 否 |  模组软件版本 只读 |
| `gateway.sort` | integer | 否 | 排序 (格式: int64) |
| `gateway.status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `gateway.subProtocolConf` | object | 否 | 子协议配置 |
| `gateway.tags` | object | 否 |  设备tag |
| `gateway.tenantCode` | string | 否 |  |
| `gateway.userID` | string | 否 |  |
| `gateway.version` | string | 否 |  固件版本  读写 |
| `gateway.withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |
| `groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `hardInfo` | string | 否 |  模组硬件型号 只读 |
| `iccid` | string | 否 | SIM卡卡号 |
| `id` | string | 否 |  |
| `imei` | string | 否 |  IMEI号信息 只读 |
| `isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `lastBind` | string | 否 | 最后绑定时间 只读 |
| `lastIp` | string | 否 | 最后登录的ip地址 |
| `lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `lastLogin` | string | 否 | 最后上线时间 只读 |
| `lastOffline` | string | 否 | 最后下线时间 只读 |
| `logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `mac` | string | 否 |  MAC号信息 只读 |
| `mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `owner` | object | 否 |  |
| `owner.avatar` | string | 否 |  用户头像 |
| `owner.createdTime` | string | 否 |  创建时间 |
| `owner.email` | string | 否 |  邮箱 |
| `owner.lastIP` | string | 否 |  最后登录ip |
| `owner.nickName` | string | 否 |  用户的昵称 |
| `owner.phone` | string | 否 |  手机号 |
| `owner.regIP` | string | 否 |  注册ip |
| `owner.userID` | string | 否 |  用户id |
| `owner.userName` | string | 否 |  用户名(唯一) |
| `phone` | string | 否 | 手机号 |
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `productID` | string | 是 | 产品id 只读 |
| `productImg` | string | 否 | 产品图片 |
| `productName` | string | 否 |  |
| `profiles` | object | 否 |  |
| `projectID` | string | 否 | 项目id 只读 |
| `protocolConf` | object | 否 | 协议配置 |
| `ratedPower` | integer | 否 | 格式: int64 |
| `remoteAddr` | string | 否 | 远程访问地址 |
| `rssi` | integer | 否 | 信号强度 (格式: int64) |
| `schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `secret` | string | 否 | 设备秘钥 只读 |
| `softInfo` | string | 否 |  模组软件版本 只读 |
| `sort` | integer | 否 | 排序 (格式: int64) |
| `status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `subProtocolConf` | object | 否 | 子协议配置 |
| `tags` | object | 否 |  设备tag |
| `tenantCode` | string | 否 |  |
| `userID` | string | 否 |  |
| `version` | string | 否 |  固件版本  读写 |
| `withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |

**请求示例**:
```json
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
ur api /api/v1/things/device/info/create \
  --body '{"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "...", "name": "...", "parentID": "...", "purpose": "..."}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "...", "name": "...", "parentID": "...", "purpose": "..."}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}'
```

### POST `/api/v1/things/device/info/update`

**说明**: 更新设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `adcode` | string | 否 | 地区编码 |
| `address` | string | 否 | 所在地址 |
| `area` | object | 否 |  |
| `area.areaID` | string | 否 | 项目区域id（只读） |
| `area.areaIDPath` | string | 否 | 项目区域ids（只读） |
| `area.areaImg` | string | 否 |  |
| `area.areaName` | string | 否 | 项目区域名称（读写） |
| `area.areaNamePath` | string | 否 | （只读） |
| `area.children` | array[AreaInfo] | 否 | 下级项目区域列表（只读） |
| `area.configFile` | string | 否 | 区域配置文件 |
| `area.createdTime` | string | 否 | 创建时间（只读） |
| `area.desc` | string | 否 | 项目区域备注（读写） |
| `area.deviceCount` | integer | 否 | 设备数量统计(只读) (格式: int64) |
| `area.devices` | array[DeviceInfo] | 否 | 区域下设备列表（按需返回） |
| `area.groupCount` | integer | 否 | 分组数量统计(只读) (格式: int64) |
| `area.isLeaf` | integer | 否 | 是否是子节点（只读） (格式: int64) |
| `area.isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `area.lowerLevelCount` | integer | 否 | 下级的数量统计（只读） (格式: int64) |
| `area.parentAreaID` | string | 否 | 上级项目区域id（只读） |
| `area.position` | object | 否 |  |
| `area.position.latitude` | number | 是 | 纬度 (格式: double) |
| `area.position.longitude` | number | 是 | 经度 (格式: double) |
| `area.projectID` | string | 否 | 项目id（只读） |
| `area.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `area.tags` | object | 否 | 自定义标签 |
| `area.tenantCode` | string | 否 | 企业编码（只读） |
| `area.useBy` | string | 否 | 用途: commonSpace default |
| `areaID` | string | 否 | 项目区域id 只读 |
| `categoryID` | string | 否 | 产品品类 |
| `cert` | string | 否 |  设备证书  只读 |
| `cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `createdTime` | string | 否 | 创建时间 只读 |
| `desc` | string | 否 |  |
| `deviceAlias` | string | 否 | 设备别名 读写 |
| `deviceImg` | string | 否 | 设备图片 |
| `deviceName` | string | 是 | 设备名称 读写 |
| `deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `distributor` | object | 否 |  |
| `distributor.id` | string | 否 |  |
| `distributor.idPath` | string | 否 |  |
| `expTime` | string | 否 | 到期时间 |
| `file` | string | 否 | 设备相关文件 |
| `firstBind` | string | 否 |  |
| `firstLogin` | string | 否 | 激活时间 只读 |
| `gateway` | object | 否 |  |
| `gateway.adcode` | string | 否 | 地区编码 |
| `gateway.address` | string | 否 | 所在地址 |
| `gateway.area` | object | 否 |  |
| `gateway.area.areaID` | string | 否 | 项目区域id（只读） |
| `gateway.area.areaIDPath` | string | 否 | 项目区域ids（只读） |
| `gateway.area.areaImg` | string | 否 |  |
| `gateway.area.areaName` | string | 否 | 项目区域名称（读写） |
| `gateway.area.areaNamePath` | string | 否 | （只读） |
| `gateway.area.children` | array[AreaInfo] | 否 | 下级项目区域列表（只读） |
| `gateway.area.configFile` | string | 否 | 区域配置文件 |
| `gateway.area.createdTime` | string | 否 | 创建时间（只读） |
| `gateway.area.desc` | string | 否 | 项目区域备注（读写） |
| `gateway.area.deviceCount` | integer | 否 | 设备数量统计(只读) (格式: int64) |
| `gateway.area.devices` | array[DeviceInfo] | 否 | 区域下设备列表（按需返回） |
| `gateway.area.groupCount` | integer | 否 | 分组数量统计(只读) (格式: int64) |
| `gateway.area.isLeaf` | integer | 否 | 是否是子节点（只读） (格式: int64) |
| `gateway.area.isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `gateway.area.lowerLevelCount` | integer | 否 | 下级的数量统计（只读） (格式: int64) |
| `gateway.area.parentAreaID` | string | 否 | 上级项目区域id（只读） |
| `gateway.area.position` | object | 否 |  |
| `gateway.area.projectID` | string | 否 | 项目id（只读） |
| `gateway.area.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `gateway.area.tags` | object | 否 | 自定义标签 |
| `gateway.area.tenantCode` | string | 否 | 企业编码（只读） |
| `gateway.area.useBy` | string | 否 | 用途: commonSpace default |
| `gateway.areaID` | string | 否 | 项目区域id 只读 |
| `gateway.categoryID` | string | 否 | 产品品类 |
| `gateway.cert` | string | 否 |  设备证书  只读 |
| `gateway.cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `gateway.createdTime` | string | 否 | 创建时间 只读 |
| `gateway.desc` | string | 否 |  |
| `gateway.deviceAlias` | string | 否 | 设备别名 读写 |
| `gateway.deviceImg` | string | 否 | 设备图片 |
| `gateway.deviceName` | string | 是 | 设备名称 读写 |
| `gateway.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `gateway.distributor` | object | 否 |  |
| `gateway.distributor.id` | string | 否 |  |
| `gateway.distributor.idPath` | string | 否 |  |
| `gateway.expTime` | string | 否 | 到期时间 |
| `gateway.file` | string | 否 | 设备相关文件 |
| `gateway.firstBind` | string | 否 |  |
| `gateway.firstLogin` | string | 否 | 激活时间 只读 |
| `gateway.gateway` | object | 否 |  |
| `gateway.gateway.adcode` | string | 否 | 地区编码 |
| `gateway.gateway.address` | string | 否 | 所在地址 |
| `gateway.gateway.area` | object | 否 |  |
| `gateway.gateway.areaID` | string | 否 | 项目区域id 只读 |
| `gateway.gateway.categoryID` | string | 否 | 产品品类 |
| `gateway.gateway.cert` | string | 否 |  设备证书  只读 |
| `gateway.gateway.cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `gateway.gateway.createdTime` | string | 否 | 创建时间 只读 |
| `gateway.gateway.desc` | string | 否 |  |
| `gateway.gateway.deviceAlias` | string | 否 | 设备别名 读写 |
| `gateway.gateway.deviceImg` | string | 否 | 设备图片 |
| `gateway.gateway.deviceName` | string | 是 | 设备名称 读写 |
| `gateway.gateway.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `gateway.gateway.distributor` | object | 否 |  |
| `gateway.gateway.expTime` | string | 否 | 到期时间 |
| `gateway.gateway.file` | string | 否 | 设备相关文件 |
| `gateway.gateway.firstBind` | string | 否 |  |
| `gateway.gateway.firstLogin` | string | 否 | 激活时间 只读 |
| `gateway.gateway.gateway` | object | 否 |  |
| `gateway.gateway.groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `gateway.gateway.groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `gateway.gateway.groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `gateway.gateway.hardInfo` | string | 否 |  模组硬件型号 只读 |
| `gateway.gateway.iccid` | string | 否 | SIM卡卡号 |
| `gateway.gateway.id` | string | 否 |  |
| `gateway.gateway.imei` | string | 否 |  IMEI号信息 只读 |
| `gateway.gateway.isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `gateway.gateway.isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `gateway.gateway.lastBind` | string | 否 | 最后绑定时间 只读 |
| `gateway.gateway.lastIp` | string | 否 | 最后登录的ip地址 |
| `gateway.gateway.lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `gateway.gateway.lastLogin` | string | 否 | 最后上线时间 只读 |
| `gateway.gateway.lastOffline` | string | 否 | 最后下线时间 只读 |
| `gateway.gateway.logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `gateway.gateway.mac` | string | 否 |  MAC号信息 只读 |
| `gateway.gateway.mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `gateway.gateway.needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `gateway.gateway.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `gateway.gateway.owner` | object | 否 |  |
| `gateway.gateway.phone` | string | 否 | 手机号 |
| `gateway.gateway.position` | object | 否 |  |
| `gateway.gateway.productID` | string | 是 | 产品id 只读 |
| `gateway.gateway.productImg` | string | 否 | 产品图片 |
| `gateway.gateway.productName` | string | 否 |  |
| `gateway.gateway.profiles` | object | 否 |  |
| `gateway.gateway.projectID` | string | 否 | 项目id 只读 |
| `gateway.gateway.protocolConf` | object | 否 | 协议配置 |
| `gateway.gateway.ratedPower` | integer | 否 | 格式: int64 |
| `gateway.gateway.remoteAddr` | string | 否 | 远程访问地址 |
| `gateway.gateway.rssi` | integer | 否 | 信号强度 (格式: int64) |
| `gateway.gateway.schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `gateway.gateway.secret` | string | 否 | 设备秘钥 只读 |
| `gateway.gateway.softInfo` | string | 否 |  模组软件版本 只读 |
| `gateway.gateway.sort` | integer | 否 | 排序 (格式: int64) |
| `gateway.gateway.status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `gateway.gateway.subProtocolConf` | object | 否 | 子协议配置 |
| `gateway.gateway.tags` | object | 否 |  设备tag |
| `gateway.gateway.tenantCode` | string | 否 |  |
| `gateway.gateway.userID` | string | 否 |  |
| `gateway.gateway.version` | string | 否 |  固件版本  读写 |
| `gateway.gateway.withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |
| `gateway.groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `gateway.groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `gateway.groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `gateway.hardInfo` | string | 否 |  模组硬件型号 只读 |
| `gateway.iccid` | string | 否 | SIM卡卡号 |
| `gateway.id` | string | 否 |  |
| `gateway.imei` | string | 否 |  IMEI号信息 只读 |
| `gateway.isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `gateway.isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `gateway.lastBind` | string | 否 | 最后绑定时间 只读 |
| `gateway.lastIp` | string | 否 | 最后登录的ip地址 |
| `gateway.lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `gateway.lastLogin` | string | 否 | 最后上线时间 只读 |
| `gateway.lastOffline` | string | 否 | 最后下线时间 只读 |
| `gateway.logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `gateway.mac` | string | 否 |  MAC号信息 只读 |
| `gateway.mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `gateway.needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `gateway.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `gateway.owner` | object | 否 |  |
| `gateway.owner.avatar` | string | 否 |  用户头像 |
| `gateway.owner.createdTime` | string | 否 |  创建时间 |
| `gateway.owner.email` | string | 否 |  邮箱 |
| `gateway.owner.lastIP` | string | 否 |  最后登录ip |
| `gateway.owner.nickName` | string | 否 |  用户的昵称 |
| `gateway.owner.phone` | string | 否 |  手机号 |
| `gateway.owner.regIP` | string | 否 |  注册ip |
| `gateway.owner.userID` | string | 否 |  用户id |
| `gateway.owner.userName` | string | 否 |  用户名(唯一) |
| `gateway.phone` | string | 否 | 手机号 |
| `gateway.position` | object | 否 |  |
| `gateway.position.latitude` | number | 是 | 纬度 (格式: double) |
| `gateway.position.longitude` | number | 是 | 经度 (格式: double) |
| `gateway.productID` | string | 是 | 产品id 只读 |
| `gateway.productImg` | string | 否 | 产品图片 |
| `gateway.productName` | string | 否 |  |
| `gateway.profiles` | object | 否 |  |
| `gateway.projectID` | string | 否 | 项目id 只读 |
| `gateway.protocolConf` | object | 否 | 协议配置 |
| `gateway.ratedPower` | integer | 否 | 格式: int64 |
| `gateway.remoteAddr` | string | 否 | 远程访问地址 |
| `gateway.rssi` | integer | 否 | 信号强度 (格式: int64) |
| `gateway.schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `gateway.secret` | string | 否 | 设备秘钥 只读 |
| `gateway.softInfo` | string | 否 |  模组软件版本 只读 |
| `gateway.sort` | integer | 否 | 排序 (格式: int64) |
| `gateway.status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `gateway.subProtocolConf` | object | 否 | 子协议配置 |
| `gateway.tags` | object | 否 |  设备tag |
| `gateway.tenantCode` | string | 否 |  |
| `gateway.userID` | string | 否 |  |
| `gateway.version` | string | 否 |  固件版本  读写 |
| `gateway.withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |
| `groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `hardInfo` | string | 否 |  模组硬件型号 只读 |
| `iccid` | string | 否 | SIM卡卡号 |
| `id` | string | 否 |  |
| `imei` | string | 否 |  IMEI号信息 只读 |
| `isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `lastBind` | string | 否 | 最后绑定时间 只读 |
| `lastIp` | string | 否 | 最后登录的ip地址 |
| `lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `lastLogin` | string | 否 | 最后上线时间 只读 |
| `lastOffline` | string | 否 | 最后下线时间 只读 |
| `logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `mac` | string | 否 |  MAC号信息 只读 |
| `mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `owner` | object | 否 |  |
| `owner.avatar` | string | 否 |  用户头像 |
| `owner.createdTime` | string | 否 |  创建时间 |
| `owner.email` | string | 否 |  邮箱 |
| `owner.lastIP` | string | 否 |  最后登录ip |
| `owner.nickName` | string | 否 |  用户的昵称 |
| `owner.phone` | string | 否 |  手机号 |
| `owner.regIP` | string | 否 |  注册ip |
| `owner.userID` | string | 否 |  用户id |
| `owner.userName` | string | 否 |  用户名(唯一) |
| `phone` | string | 否 | 手机号 |
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `productID` | string | 是 | 产品id 只读 |
| `productImg` | string | 否 | 产品图片 |
| `productName` | string | 否 |  |
| `profiles` | object | 否 |  |
| `projectID` | string | 否 | 项目id 只读 |
| `protocolConf` | object | 否 | 协议配置 |
| `ratedPower` | integer | 否 | 格式: int64 |
| `remoteAddr` | string | 否 | 远程访问地址 |
| `rssi` | integer | 否 | 信号强度 (格式: int64) |
| `schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `secret` | string | 否 | 设备秘钥 只读 |
| `softInfo` | string | 否 |  模组软件版本 只读 |
| `sort` | integer | 否 | 排序 (格式: int64) |
| `status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `subProtocolConf` | object | 否 | 子协议配置 |
| `tags` | object | 否 |  设备tag |
| `tenantCode` | string | 否 |  |
| `userID` | string | 否 |  |
| `version` | string | 否 |  固件版本  读写 |
| `withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |

**请求示例**:
```json
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
ur api /api/v1/things/device/info/update \
  --body '{"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "...", "name": "...", "parentID": "...", "purpose": "..."}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "...", "name": "...", "parentID": "...", "purpose": "..."}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}'
```
