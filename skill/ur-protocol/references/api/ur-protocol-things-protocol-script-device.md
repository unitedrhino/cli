# ur-protocol things/protocol/script/device

新增协议脚本设备 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/protocol/script/device/create` | 新增协议脚本设备 | admin |
| POST | `/api/v1/things/protocol/script/device/delete` | 删除协议脚本设备 | admin |
| POST | `/api/v1/things/protocol/script/device/get-list` | 获取协议脚本设备列表 | admin |
| POST | `/api/v1/things/protocol/script/device/get-one` | 获取协议脚本设备详情 | admin |
| POST | `/api/v1/things/protocol/script/device/update` | 更新协议脚本设备 | admin |

## 详细说明

### POST `/api/v1/things/protocol/script/device/create`

**说明**: 新增协议脚本设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间 |
| `device` | object | 否 |  |
| `device.adcode` | string | 否 | 地区编码 |
| `device.address` | string | 否 | 所在地址 |
| `device.area` | object | 否 |  |
| `device.area.areaID` | string | 否 | 项目区域id（只读） |
| `device.area.areaIDPath` | string | 否 | 项目区域ids（只读） |
| `device.area.areaImg` | string | 否 |  |
| `device.area.areaName` | string | 否 | 项目区域名称（读写） |
| `device.area.areaNamePath` | string | 否 | （只读） |
| `device.area.children` | array[AreaInfo] | 否 | 下级项目区域列表（只读） |
| `device.area.configFile` | string | 否 | 区域配置文件 |
| `device.area.createdTime` | string | 否 | 创建时间（只读） |
| `device.area.desc` | string | 否 | 项目区域备注（读写） |
| `device.area.deviceCount` | integer | 否 | 设备数量统计(只读) (格式: int64) |
| `device.area.devices` | array[DeviceInfo] | 否 | 区域下设备列表（按需返回） |
| `device.area.groupCount` | integer | 否 | 分组数量统计(只读) (格式: int64) |
| `device.area.isLeaf` | integer | 否 | 是否是子节点（只读） (格式: int64) |
| `device.area.isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `device.area.lowerLevelCount` | integer | 否 | 下级的数量统计（只读） (格式: int64) |
| `device.area.parentAreaID` | string | 否 | 上级项目区域id（只读） |
| `device.area.position` | object | 否 |  |
| `device.area.projectID` | string | 否 | 项目id（只读） |
| `device.area.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `device.area.tags` | object | 否 | 自定义标签 |
| `device.area.tenantCode` | string | 否 | 企业编码（只读） |
| `device.area.useBy` | string | 否 | 用途: commonSpace default |
| `device.areaID` | string | 否 | 项目区域id 只读 |
| `device.categoryID` | string | 否 | 产品品类 |
| `device.cert` | string | 否 |  设备证书  只读 |
| `device.cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `device.createdTime` | string | 否 | 创建时间 只读 |
| `device.desc` | string | 否 |  |
| `device.deviceAlias` | string | 否 | 设备别名 读写 |
| `device.deviceImg` | string | 否 | 设备图片 |
| `device.deviceName` | string | 是 | 设备名称 读写 |
| `device.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `device.distributor` | object | 否 |  |
| `device.distributor.id` | string | 否 |  |
| `device.distributor.idPath` | string | 否 |  |
| `device.expTime` | string | 否 | 到期时间 |
| `device.file` | string | 否 | 设备相关文件 |
| `device.firstBind` | string | 否 |  |
| `device.firstLogin` | string | 否 | 激活时间 只读 |
| `device.gateway` | object | 否 |  |
| `device.gateway.adcode` | string | 否 | 地区编码 |
| `device.gateway.address` | string | 否 | 所在地址 |
| `device.gateway.area` | object | 否 |  |
| `device.gateway.areaID` | string | 否 | 项目区域id 只读 |
| `device.gateway.categoryID` | string | 否 | 产品品类 |
| `device.gateway.cert` | string | 否 |  设备证书  只读 |
| `device.gateway.cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `device.gateway.createdTime` | string | 否 | 创建时间 只读 |
| `device.gateway.desc` | string | 否 |  |
| `device.gateway.deviceAlias` | string | 否 | 设备别名 读写 |
| `device.gateway.deviceImg` | string | 否 | 设备图片 |
| `device.gateway.deviceName` | string | 是 | 设备名称 读写 |
| `device.gateway.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `device.gateway.distributor` | object | 否 |  |
| `device.gateway.expTime` | string | 否 | 到期时间 |
| `device.gateway.file` | string | 否 | 设备相关文件 |
| `device.gateway.firstBind` | string | 否 |  |
| `device.gateway.firstLogin` | string | 否 | 激活时间 只读 |
| `device.gateway.gateway` | object | 否 |  |
| `device.gateway.groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `device.gateway.groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `device.gateway.groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `device.gateway.hardInfo` | string | 否 |  模组硬件型号 只读 |
| `device.gateway.iccid` | string | 否 | SIM卡卡号 |
| `device.gateway.id` | string | 否 |  |
| `device.gateway.imei` | string | 否 |  IMEI号信息 只读 |
| `device.gateway.isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `device.gateway.isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `device.gateway.lastBind` | string | 否 | 最后绑定时间 只读 |
| `device.gateway.lastIp` | string | 否 | 最后登录的ip地址 |
| `device.gateway.lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `device.gateway.lastLogin` | string | 否 | 最后上线时间 只读 |
| `device.gateway.lastOffline` | string | 否 | 最后下线时间 只读 |
| `device.gateway.logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `device.gateway.mac` | string | 否 |  MAC号信息 只读 |
| `device.gateway.mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `device.gateway.needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `device.gateway.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `device.gateway.owner` | object | 否 |  |
| `device.gateway.phone` | string | 否 | 手机号 |
| `device.gateway.position` | object | 否 |  |
| `device.gateway.productID` | string | 是 | 产品id 只读 |
| `device.gateway.productImg` | string | 否 | 产品图片 |
| `device.gateway.productName` | string | 否 |  |
| `device.gateway.profiles` | object | 否 |  |
| `device.gateway.projectID` | string | 否 | 项目id 只读 |
| `device.gateway.protocolConf` | object | 否 | 协议配置 |
| `device.gateway.ratedPower` | integer | 否 | 格式: int64 |
| `device.gateway.remoteAddr` | string | 否 | 远程访问地址 |
| `device.gateway.rssi` | integer | 否 | 信号强度 (格式: int64) |
| `device.gateway.schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `device.gateway.secret` | string | 否 | 设备秘钥 只读 |
| `device.gateway.softInfo` | string | 否 |  模组软件版本 只读 |
| `device.gateway.sort` | integer | 否 | 排序 (格式: int64) |
| `device.gateway.status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `device.gateway.subProtocolConf` | object | 否 | 子协议配置 |
| `device.gateway.tags` | object | 否 |  设备tag |
| `device.gateway.tenantCode` | string | 否 |  |
| `device.gateway.userID` | string | 否 |  |
| `device.gateway.version` | string | 否 |  固件版本  读写 |
| `device.gateway.withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |
| `device.groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `device.groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `device.groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `device.hardInfo` | string | 否 |  模组硬件型号 只读 |
| `device.iccid` | string | 否 | SIM卡卡号 |
| `device.id` | string | 否 |  |
| `device.imei` | string | 否 |  IMEI号信息 只读 |
| `device.isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `device.isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `device.lastBind` | string | 否 | 最后绑定时间 只读 |
| `device.lastIp` | string | 否 | 最后登录的ip地址 |
| `device.lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `device.lastLogin` | string | 否 | 最后上线时间 只读 |
| `device.lastOffline` | string | 否 | 最后下线时间 只读 |
| `device.logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `device.mac` | string | 否 |  MAC号信息 只读 |
| `device.mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `device.needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `device.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `device.owner` | object | 否 |  |
| `device.owner.avatar` | string | 否 |  用户头像 |
| `device.owner.createdTime` | string | 否 |  创建时间 |
| `device.owner.email` | string | 否 |  邮箱 |
| `device.owner.lastIP` | string | 否 |  最后登录ip |
| `device.owner.nickName` | string | 否 |  用户的昵称 |
| `device.owner.phone` | string | 否 |  手机号 |
| `device.owner.regIP` | string | 否 |  注册ip |
| `device.owner.userID` | string | 否 |  用户id |
| `device.owner.userName` | string | 否 |  用户名(唯一) |
| `device.phone` | string | 否 | 手机号 |
| `device.position` | object | 否 |  |
| `device.position.latitude` | number | 是 | 纬度 (格式: double) |
| `device.position.longitude` | number | 是 | 经度 (格式: double) |
| `device.productID` | string | 是 | 产品id 只读 |
| `device.productImg` | string | 否 | 产品图片 |
| `device.productName` | string | 否 |  |
| `device.profiles` | object | 否 |  |
| `device.projectID` | string | 否 | 项目id 只读 |
| `device.protocolConf` | object | 否 | 协议配置 |
| `device.ratedPower` | integer | 否 | 格式: int64 |
| `device.remoteAddr` | string | 否 | 远程访问地址 |
| `device.rssi` | integer | 否 | 信号强度 (格式: int64) |
| `device.schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `device.secret` | string | 否 | 设备秘钥 只读 |
| `device.softInfo` | string | 否 |  模组软件版本 只读 |
| `device.sort` | integer | 否 | 排序 (格式: int64) |
| `device.status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `device.subProtocolConf` | object | 否 | 子协议配置 |
| `device.tags` | object | 否 |  设备tag |
| `device.tenantCode` | string | 否 |  |
| `device.userID` | string | 否 |  |
| `device.version` | string | 否 |  固件版本  读写 |
| `device.withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |
| `deviceName` | string | 否 |  |
| `id` | string | 否 |  |
| `priority` | integer | 否 | 执行优先级 (格式: int64) |
| `product` | object | 否 |  |
| `product.authMode` | integer | 否 | 认证方式:1:账密认证,2:秘钥认证 (格式: int64) |
| `product.autoRegister` | integer | 否 | 动态注册:1:关闭,2:打开,3:打开并自动创建设备,4:在前面的基础上绑定没有也自动创建 (格式: int64) |
| `product.bindLevel` | integer | 否 | 绑定级别: 1:强绑定(默认,只有用户解绑之后才能绑定) 2:中绑定(可以通过token强制解绑设备) 3:弱绑定(app可以内部解绑被绑定的设备) (格式: int64) |
| `product.category` | object | 否 |  |
| `product.category.avatar` | string | 否 |  品类图片 |
| `product.category.children` | array[ProductCategory] | 否 |  |
| `product.category.code` | string | 否 | 品类编码,全局唯一,创建后不可修改,创建时必填 |
| `product.category.desc` | string | 否 | 描述 |
| `product.category.deviceCount` | integer | 否 | 设备数量 (格式: int64) |
| `product.category.id` | string | 否 |  |
| `product.category.idPath` | array[integer] | 否 | 只读 |
| `product.category.isLeaf` | integer | 否 | 格式: int64 |
| `product.category.name` | string | 否 | 品类名称 |
| `product.category.order` | integer | 否 | 排序,必须大于等于1,为0代表不修改 (格式: int64) |
| `product.category.parentID` | string | 否 |  |
| `product.categoryID` | string | 否 | 产品品类 |
| `product.compatMode` | integer | 否 | 协议兼容模式 1:旧productID(大写D) 2:新productId(小写d,默认) (格式: int64) |
| `product.config` | object | 否 |  |
| `product.config.devInit` | object | 否 |  |
| `product.config.productID` | string | 是 |  |
| `product.createdTime` | string | 否 | 创建时间 只读 |
| `product.customUi` | object | 否 | 自定义ui,key是端的类型(web-client  mini-client) value是以下类型的对象{version:123(版本号,只读),isUpdateUi:bool(是否更新ui),path:string(前端路径,如果需要修改,需要将isUpdateUi置为true并在这个参数中传入压缩包的filePath)} |
| `product.defaultAgentId` | integer | 否 | 默认绑定的 Agent ID，0 表示未绑定 (格式: int64) |
| `product.desc` | string | 否 | 描述 |
| `product.deviceSchemaMode` | integer | 否 |  设备物模型模式:1:手动创建,2:设备自动创建 3: 设备自动创建及上报无定义自动创建 4: 设备自动创建及上报无定义自动创建(数字类型只使用浮点) (格式: int64) |
| `product.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `product.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `product.onlineHandle` | integer | 否 | 在线处理:1: 自动 2: 永远在线 (格式: int64) |
| `product.productID` | string | 否 | 产品id 只读 |
| `product.productImg` | string | 否 | 产品图片 |
| `product.productName` | string | 否 | 产品名称 |
| `product.protocol` | object | 否 |  |
| `product.protocol.code` | string | 否 |   iThings,iThings-thingsboard,wumei,aliyun,huaweiyun,tuya |
| `product.protocol.configFields` | array[ProtocolConfigField] | 否 | 配置字段列表,没有可以不传 |
| `product.protocol.containerPort` | integer | 否 |  容器端口 (格式: int64) |
| `product.protocol.deployStatus` | string | 否 |  部署状态: notInstalled/downloaded/running/stopped |
| `product.protocol.desc` | string | 否 |  |
| `product.protocol.deviceFields` | array[ProtocolConfigField] | 否 | 设备级的配置字段列表,没有可以不传 |
| `product.protocol.endpoints` | array[string] | 否 |  |
| `product.protocol.envVars` | object | 否 |  环境变量 |
| `product.protocol.etcdKey` | string | 否 |  |
| `product.protocol.id` | string | 否 |  |
| `product.protocol.image` | string | 否 |  镜像地址 |
| `product.protocol.isBuiltin` | integer | 否 |  是否内置 1:是 2:否 (格式: int64) |
| `product.protocol.isEnableSyncDevice` | integer | 是 | 是否支持主动同步设备信息 (格式: int64) |
| `product.protocol.isEnableSyncProduct` | integer | 是 | 是否支持主动同步产品信息 (格式: int64) |
| `product.protocol.name` | string | 否 |  |
| `product.protocol.productFields` | array[ProtocolConfigField] | 否 | 产品级的配置字段列表,没有可以不传 |
| `product.protocol.transProtocol` | string | 否 |  传输协议: mqtt,tcp,udp |
| `product.protocol.type` | string | 否 | 协议类型: 普通设备(默认):normal  音视频:media |
| `product.protocol.version` | string | 否 |  镜像版本 |
| `product.protocolCode` | string | 否 | 协议code,默认urMqtt,设备下发只会发送给主协议  urMqtt,urHttp,wumei,aliyun,huaweiyun,tuya |
| `product.protocolConf` | object | 否 | 协议配置 |
| `product.sceneMode` | string | 否 | 场景模式 读写类型: r(只读) rw(可读可写) none(不参与场景) |
| `product.secret` | string | 否 | 动态注册产品秘钥 只读 |
| `product.status` | integer | 否 | 产品状态 1:启用 2:禁用 3:开发中 (格式: int64) |
| `product.subProtocolCode` | string | 否 | 子协议,主协议和子协议传输类型必须不相同, 设备控制下发只会发送给主协议, 当设备是音视频设备但是控制协议需要单独走的时候就可以把主协议定义为普通协议,子协议定义为音视频协议,这样就能实现音视频走音视频协议,控制走子协议 |
| `product.subProtocolConf` | object | 否 | 子协议协议配置 |
| `product.tags` | object | 否 |  产品tag |
| `product.tenantCode` | string | 否 | 企业号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `product.trialTime` | string | 否 | 试用时间(单位为天,为0不限制) |
| `productID` | string | 否 |  |
| `script` | object | 否 |  |
| `script.createdTime` | string | 否 | 创建时间 |
| `script.desc` | string | 否 |  |
| `script.id` | string | 否 |  |
| `script.name` | string | 否 |  |
| `script.script` | string | 否 | 脚本内容 |
| `script.status` | integer | 否 | 状态:是否启用 (格式: int64) |
| `script.tenantCode` | string | 否 | 企业号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `script.triggerDir` | integer | 否 | 1:up 2:down (格式: int64) |
| `script.triggerHandle` | string | 否 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `script.triggerTimer` | integer | 否 | 收到前处理1:before 2:after (格式: int64) |
| `script.triggerType` | string | 否 |  操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |
| `scriptID` | string | 否 |  |
| `status` | integer | 否 | 状态:是否启用 (格式: int64) |
| `tenantCode` | string | 否 | 企业号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `triggerSrc` | integer | 否 | product:1 device:2 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "device": {
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
  "deviceName": "string",
  "id": "string",
  "priority": 1,
  "product": {
    "authMode": 1,
    "autoRegister": 1,
    "bindLevel": 1,
    "category": {
      "avatar": "string",
      "children": [
        {
          "avatar": "string",
          "children": [],
          "code": "string",
          "desc": "string",
          "deviceCount": 1,
          "id": "string",
          "idPath": [],
          "isLeaf": 1,
          "name": "示例名称",
          "order": 1,
          "parentID": "string"
        }
      ],
      "code": "string",
      "desc": "string",
      "deviceCount": 1,
      "id": "string",
      "idPath": [
        1
      ],
      "isLeaf": 1,
      "name": "示例名称",
      "order": 1,
      "parentID": "string"
    },
    "categoryID": "string",
    "compatMode": 1,
    "config": {
      "devInit": {
        "areaID": "string",
        "deptID": "string",
        "projectID": "string",
        "tenantCode": "string"
      },
      "productID": "string"
    },
    "createdTime": "2026-01-01T00:00:00Z",
    "customUi": {},
    "defaultAgentId": 1,
    "desc": "string",
    "deviceSchemaMode": 1,
    "deviceType": 1,
    "netType": 1,
    "onlineHandle": 1,
    "productID": "string",
    "productImg": "string",
    "productName": "示例名称",
    "protocol": {
      "code": "string",
      "configFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "containerPort": 1,
      "deployStatus": "string",
      "desc": "string",
      "deviceFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "endpoints": [
        "string"
      ],
      "envVars": {},
      "etcdKey": "string",
      "id": "string",
      "image": "string",
      "isBuiltin": 1,
      "isEnableSyncDevice": 1,
      "isEnableSyncProduct": 1,
      "name": "string",
      "productFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "transProtocol": "string",
      "type": "string",
      "version": "string"
    },
    "protocolCode": "string",
    "protocolConf": {},
    "sceneMode": "string",
    "secret": "string",
    "status": 1,
    "subProtocolCode": "string",
    "subProtocolConf": {},
    "tags": {},
    "tenantCode": "string",
    "trialTime": "2026-01-01T00:00:00Z"
  },
  "productID": "string",
  "script": {
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "name": "string",
    "script": "string",
    "status": 1,
    "tenantCode": "string",
    "triggerDir": 1,
    "triggerHandle": "string",
    "triggerTimer": 1,
    "triggerType": "string"
  },
  "scriptID": "string",
  "status": 1,
  "tenantCode": "string",
  "triggerSrc": 1
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
ur api /api/v1/things/protocol/script/device/create \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "device": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "...", "name": "...", "parentID": "...", "purpose": "..."}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "deviceName": "string", "id": "string", "priority": 1, "product": {"authMode": 1, "autoRegister": 1, "bindLevel": 1, "category": {"avatar": "string", "children": [{"avatar": "string", "children": [], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}, "categoryID": "string", "compatMode": 1, "config": {"devInit": {"areaID": "string", "deptID": "string", "projectID": "string", "tenantCode": "string"}, "productID": "string"}, "createdTime": "2026-01-01T00:00:00Z", "customUi": {}, "defaultAgentId": 1, "desc": "string", "deviceSchemaMode": 1, "deviceType": 1, "netType": 1, "onlineHandle": 1, "productID": "string", "productImg": "string", "productName": "示例名称", "protocol": {"code": "string", "configFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "containerPort": 1, "deployStatus": "string", "desc": "string", "deviceFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "endpoints": ["string"], "envVars": {}, "etcdKey": "string", "id": "string", "image": "string", "isBuiltin": 1, "isEnableSyncDevice": 1, "isEnableSyncProduct": 1, "name": "string", "productFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "transProtocol": "string", "type": "string", "version": "string"}, "protocolCode": "string", "protocolConf": {}, "sceneMode": "string", "secret": "string", "status": 1, "subProtocolCode": "string", "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "trialTime": "2026-01-01T00:00:00Z"}, "productID": "string", "script": {"createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "name": "string", "script": "string", "status": 1, "tenantCode": "string", "triggerDir": 1, "triggerHandle": "string", "triggerTimer": 1, "triggerType": "string"}, "scriptID": "string", "status": 1, "tenantCode": "string", "triggerSrc": 1}'
```

### POST `/api/v1/things/protocol/script/device/delete`

**说明**: 删除协议脚本设备

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
ur api /api/v1/things/protocol/script/device/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/protocol/script/device/get-list`

**说明**: 获取协议脚本设备列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 否 |  |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 |  |
| `scriptID` | string | 否 |  |
| `status` | integer | 否 | 状态:是否启用 (格式: int64) |
| `triggerSrc` | integer | 否 | product:1 device:2 (格式: int64) |
| `withDevice` | boolean | 否 | 格式: boolean |
| `withScript` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "deviceName": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "productID": "string",
  "scriptID": "string",
  "status": 1,
  "triggerSrc": 1,
  "withDevice": true,
  "withScript": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "device": {
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
        "deviceName": "string",
        "id": "string",
        "priority": 1,
        "product": {
          "authMode": 1,
          "autoRegister": 1,
          "bindLevel": 1,
          "category": {
            "avatar": "string",
            "children": [],
            "code": "string",
            "desc": "string",
            "deviceCount": 1,
            "id": "string",
            "idPath": [],
            "isLeaf": 1,
            "name": "示例名称",
            "order": 1,
            "parentID": "string"
          },
          "categoryID": "string",
          "compatMode": 1,
          "config": {
            "devInit": {
              "areaID": "...",
              "deptID": "...",
              "projectID": "...",
              "tenantCode": "..."
            },
            "productID": "string"
          },
          "createdTime": "2026-01-01T00:00:00Z",
          "customUi": {},
          "defaultAgentId": 1,
          "desc": "string",
          "deviceSchemaMode": 1,
          "deviceType": 1,
          "netType": 1,
          "onlineHandle": 1,
          "productID": "string",
          "productImg": "string",
          "productName": "示例名称",
          "protocol": {
            "code": "string",
            "configFields": [],
            "containerPort": 1,
            "deployStatus": "string",
            "desc": "string",
            "deviceFields": [],
            "endpoints": [],
            "envVars": {},
            "etcdKey": "string",
            "id": "string",
            "image": "string",
            "isBuiltin": 1,
            "isEnableSyncDevice": 1,
            "isEnableSyncProduct": 1,
            "name": "string",
            "productFields": [],
            "transProtocol": "string",
            "type": "string",
            "version": "string"
          },
          "protocolCode": "string",
          "protocolConf": {},
          "sceneMode": "string",
          "secret": "string",
          "status": 1,
          "subProtocolCode": "string",
          "subProtocolConf": {},
          "tags": {},
          "tenantCode": "string",
          "trialTime": "2026-01-01T00:00:00Z"
        },
        "productID": "string",
        "script": {
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "id": "string",
          "name": "string",
          "script": "string",
          "status": 1,
          "tenantCode": "string",
          "triggerDir": 1,
          "triggerHandle": "string",
          "triggerTimer": 1,
          "triggerType": "string"
        },
        "scriptID": "string",
        "status": 1,
        "tenantCode": "string",
        "triggerSrc": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/script/device/get-list \
  --body '{"deviceName": "string", "page": {"page": 1, "pageSize": 1}, "productID": "string", "scriptID": "string", "status": 1, "triggerSrc": 1, "withDevice": true, "withScript": true}'
```

### POST `/api/v1/things/protocol/script/device/get-one`

**说明**: 获取协议脚本设备详情

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
    "createdTime": "2026-01-01T00:00:00Z",
    "device": {
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
    "deviceName": "string",
    "id": "string",
    "priority": 1,
    "product": {
      "authMode": 1,
      "autoRegister": 1,
      "bindLevel": 1,
      "category": {
        "avatar": "string",
        "children": [
          {
            "avatar": "string",
            "children": [],
            "code": "string",
            "desc": "string",
            "deviceCount": 1,
            "id": "string",
            "idPath": [],
            "isLeaf": 1,
            "name": "示例名称",
            "order": 1,
            "parentID": "string"
          }
        ],
        "code": "string",
        "desc": "string",
        "deviceCount": 1,
        "id": "string",
        "idPath": [
          1
        ],
        "isLeaf": 1,
        "name": "示例名称",
        "order": 1,
        "parentID": "string"
      },
      "categoryID": "string",
      "compatMode": 1,
      "config": {
        "devInit": {
          "areaID": "string",
          "deptID": "string",
          "projectID": "string",
          "tenantCode": "string"
        },
        "productID": "string"
      },
      "createdTime": "2026-01-01T00:00:00Z",
      "customUi": {},
      "defaultAgentId": 1,
      "desc": "string",
      "deviceSchemaMode": 1,
      "deviceType": 1,
      "netType": 1,
      "onlineHandle": 1,
      "productID": "string",
      "productImg": "string",
      "productName": "示例名称",
      "protocol": {
        "code": "string",
        "configFields": [
          {
            "defaultValue": "string",
            "group": "string",
            "id": "string",
            "isRequired": true,
            "key": "string",
            "label": "string",
            "mapping": {},
            "max": "string",
            "min": "string",
            "placeholder": "string",
            "sort": 1,
            "type": "string"
          }
        ],
        "containerPort": 1,
        "deployStatus": "string",
        "desc": "string",
        "deviceFields": [
          {
            "defaultValue": "string",
            "group": "string",
            "id": "string",
            "isRequired": true,
            "key": "string",
            "label": "string",
            "mapping": {},
            "max": "string",
            "min": "string",
            "placeholder": "string",
            "sort": 1,
            "type": "string"
          }
        ],
        "endpoints": [
          "string"
        ],
        "envVars": {},
        "etcdKey": "string",
        "id": "string",
        "image": "string",
        "isBuiltin": 1,
        "isEnableSyncDevice": 1,
        "isEnableSyncProduct": 1,
        "name": "string",
        "productFields": [
          {
            "defaultValue": "string",
            "group": "string",
            "id": "string",
            "isRequired": true,
            "key": "string",
            "label": "string",
            "mapping": {},
            "max": "string",
            "min": "string",
            "placeholder": "string",
            "sort": 1,
            "type": "string"
          }
        ],
        "transProtocol": "string",
        "type": "string",
        "version": "string"
      },
      "protocolCode": "string",
      "protocolConf": {},
      "sceneMode": "string",
      "secret": "string",
      "status": 1,
      "subProtocolCode": "string",
      "subProtocolConf": {},
      "tags": {},
      "tenantCode": "string",
      "trialTime": "2026-01-01T00:00:00Z"
    },
    "productID": "string",
    "script": {
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "id": "string",
      "name": "string",
      "script": "string",
      "status": 1,
      "tenantCode": "string",
      "triggerDir": 1,
      "triggerHandle": "string",
      "triggerTimer": 1,
      "triggerType": "string"
    },
    "scriptID": "string",
    "status": 1,
    "tenantCode": "string",
    "triggerSrc": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/script/device/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/protocol/script/device/update`

**说明**: 更新协议脚本设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间 |
| `device` | object | 否 |  |
| `device.adcode` | string | 否 | 地区编码 |
| `device.address` | string | 否 | 所在地址 |
| `device.area` | object | 否 |  |
| `device.area.areaID` | string | 否 | 项目区域id（只读） |
| `device.area.areaIDPath` | string | 否 | 项目区域ids（只读） |
| `device.area.areaImg` | string | 否 |  |
| `device.area.areaName` | string | 否 | 项目区域名称（读写） |
| `device.area.areaNamePath` | string | 否 | （只读） |
| `device.area.children` | array[AreaInfo] | 否 | 下级项目区域列表（只读） |
| `device.area.configFile` | string | 否 | 区域配置文件 |
| `device.area.createdTime` | string | 否 | 创建时间（只读） |
| `device.area.desc` | string | 否 | 项目区域备注（读写） |
| `device.area.deviceCount` | integer | 否 | 设备数量统计(只读) (格式: int64) |
| `device.area.devices` | array[DeviceInfo] | 否 | 区域下设备列表（按需返回） |
| `device.area.groupCount` | integer | 否 | 分组数量统计(只读) (格式: int64) |
| `device.area.isLeaf` | integer | 否 | 是否是子节点（只读） (格式: int64) |
| `device.area.isSysCreated` | integer | 否 | 是否是系统创建的,系统创建的只有管理员可以删除 (格式: int64) |
| `device.area.lowerLevelCount` | integer | 否 | 下级的数量统计（只读） (格式: int64) |
| `device.area.parentAreaID` | string | 否 | 上级项目区域id（只读） |
| `device.area.position` | object | 否 |  |
| `device.area.projectID` | string | 否 | 项目id（只读） |
| `device.area.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `device.area.tags` | object | 否 | 自定义标签 |
| `device.area.tenantCode` | string | 否 | 企业编码（只读） |
| `device.area.useBy` | string | 否 | 用途: commonSpace default |
| `device.areaID` | string | 否 | 项目区域id 只读 |
| `device.categoryID` | string | 否 | 产品品类 |
| `device.cert` | string | 否 |  设备证书  只读 |
| `device.cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `device.createdTime` | string | 否 | 创建时间 只读 |
| `device.desc` | string | 否 |  |
| `device.deviceAlias` | string | 否 | 设备别名 读写 |
| `device.deviceImg` | string | 否 | 设备图片 |
| `device.deviceName` | string | 是 | 设备名称 读写 |
| `device.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `device.distributor` | object | 否 |  |
| `device.distributor.id` | string | 否 |  |
| `device.distributor.idPath` | string | 否 |  |
| `device.expTime` | string | 否 | 到期时间 |
| `device.file` | string | 否 | 设备相关文件 |
| `device.firstBind` | string | 否 |  |
| `device.firstLogin` | string | 否 | 激活时间 只读 |
| `device.gateway` | object | 否 |  |
| `device.gateway.adcode` | string | 否 | 地区编码 |
| `device.gateway.address` | string | 否 | 所在地址 |
| `device.gateway.area` | object | 否 |  |
| `device.gateway.areaID` | string | 否 | 项目区域id 只读 |
| `device.gateway.categoryID` | string | 否 | 产品品类 |
| `device.gateway.cert` | string | 否 |  设备证书  只读 |
| `device.gateway.cloneID` | string | 否 | 绑定的 Clone ID，0 表示未绑定，只读 |
| `device.gateway.createdTime` | string | 否 | 创建时间 只读 |
| `device.gateway.desc` | string | 否 |  |
| `device.gateway.deviceAlias` | string | 否 | 设备别名 读写 |
| `device.gateway.deviceImg` | string | 否 | 设备图片 |
| `device.gateway.deviceName` | string | 是 | 设备名称 读写 |
| `device.gateway.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `device.gateway.distributor` | object | 否 |  |
| `device.gateway.expTime` | string | 否 | 到期时间 |
| `device.gateway.file` | string | 否 | 设备相关文件 |
| `device.gateway.firstBind` | string | 否 |  |
| `device.gateway.firstLogin` | string | 否 | 激活时间 只读 |
| `device.gateway.gateway` | object | 否 |  |
| `device.gateway.groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `device.gateway.groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `device.gateway.groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `device.gateway.hardInfo` | string | 否 |  模组硬件型号 只读 |
| `device.gateway.iccid` | string | 否 | SIM卡卡号 |
| `device.gateway.id` | string | 否 |  |
| `device.gateway.imei` | string | 否 |  IMEI号信息 只读 |
| `device.gateway.isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `device.gateway.isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `device.gateway.lastBind` | string | 否 | 最后绑定时间 只读 |
| `device.gateway.lastIp` | string | 否 | 最后登录的ip地址 |
| `device.gateway.lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `device.gateway.lastLogin` | string | 否 | 最后上线时间 只读 |
| `device.gateway.lastOffline` | string | 否 | 最后下线时间 只读 |
| `device.gateway.logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `device.gateway.mac` | string | 否 |  MAC号信息 只读 |
| `device.gateway.mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `device.gateway.needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `device.gateway.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `device.gateway.owner` | object | 否 |  |
| `device.gateway.phone` | string | 否 | 手机号 |
| `device.gateway.position` | object | 否 |  |
| `device.gateway.productID` | string | 是 | 产品id 只读 |
| `device.gateway.productImg` | string | 否 | 产品图片 |
| `device.gateway.productName` | string | 否 |  |
| `device.gateway.profiles` | object | 否 |  |
| `device.gateway.projectID` | string | 否 | 项目id 只读 |
| `device.gateway.protocolConf` | object | 否 | 协议配置 |
| `device.gateway.ratedPower` | integer | 否 | 格式: int64 |
| `device.gateway.remoteAddr` | string | 否 | 远程访问地址 |
| `device.gateway.rssi` | integer | 否 | 信号强度 (格式: int64) |
| `device.gateway.schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `device.gateway.secret` | string | 否 | 设备秘钥 只读 |
| `device.gateway.softInfo` | string | 否 |  模组软件版本 只读 |
| `device.gateway.sort` | integer | 否 | 排序 (格式: int64) |
| `device.gateway.status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `device.gateway.subProtocolConf` | object | 否 | 子协议配置 |
| `device.gateway.tags` | object | 否 |  设备tag |
| `device.gateway.tenantCode` | string | 否 |  |
| `device.gateway.userID` | string | 否 |  |
| `device.gateway.version` | string | 否 |  固件版本  读写 |
| `device.gateway.withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |
| `device.groupIDs` | array[string] | 否 | 如果要更新分组,必须填写groupPurpose |
| `device.groupPurpose` | string | 否 | 更新的时候使用,将设备该用途下的分组进行更新 |
| `device.groups` | array[GroupCore] | 否 | 所在分组信息,只读 |
| `device.hardInfo` | string | 否 |  模组硬件型号 只读 |
| `device.iccid` | string | 否 | SIM卡卡号 |
| `device.id` | string | 否 |  |
| `device.imei` | string | 否 |  IMEI号信息 只读 |
| `device.isEnable` | integer | 否 | 是否启用 (格式: int64) |
| `device.isOnline` | integer | 否 |  在线状态  1离线 2在线 只读 (格式: int64) |
| `device.lastBind` | string | 否 | 最后绑定时间 只读 |
| `device.lastIp` | string | 否 | 最后登录的ip地址 |
| `device.lastLocalIp` | string | 否 | 最后登录局域网的ip地址 |
| `device.lastLogin` | string | 否 | 最后上线时间 只读 |
| `device.lastOffline` | string | 否 | 最后下线时间 只读 |
| `device.logLevel` | string | 否 |  日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写 |
| `device.mac` | string | 否 |  MAC号信息 只读 |
| `device.mobileOperator` | integer | 否 | 移动运营商:1)无 2)移动 3)联通 4)电信 5)广电 (格式: int64) |
| `device.needConfirmVersion` | string | 否 | 待确认升级的版本 |
| `device.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `device.owner` | object | 否 |  |
| `device.owner.avatar` | string | 否 |  用户头像 |
| `device.owner.createdTime` | string | 否 |  创建时间 |
| `device.owner.email` | string | 否 |  邮箱 |
| `device.owner.lastIP` | string | 否 |  最后登录ip |
| `device.owner.nickName` | string | 否 |  用户的昵称 |
| `device.owner.phone` | string | 否 |  手机号 |
| `device.owner.regIP` | string | 否 |  注册ip |
| `device.owner.userID` | string | 否 |  用户id |
| `device.owner.userName` | string | 否 |  用户名(唯一) |
| `device.phone` | string | 否 | 手机号 |
| `device.position` | object | 否 |  |
| `device.position.latitude` | number | 是 | 纬度 (格式: double) |
| `device.position.longitude` | number | 是 | 经度 (格式: double) |
| `device.productID` | string | 是 | 产品id 只读 |
| `device.productImg` | string | 否 | 产品图片 |
| `device.productName` | string | 否 |  |
| `device.profiles` | object | 否 |  |
| `device.projectID` | string | 否 | 项目id 只读 |
| `device.protocolConf` | object | 否 | 协议配置 |
| `device.ratedPower` | integer | 否 | 格式: int64 |
| `device.remoteAddr` | string | 否 | 远程访问地址 |
| `device.rssi` | integer | 否 | 信号强度 (格式: int64) |
| `device.schemaAlias` | object | 否 | 设备物模型别名,如果是结构体类型则key为xxx.xxx |
| `device.secret` | string | 否 | 设备秘钥 只读 |
| `device.softInfo` | string | 否 |  模组软件版本 只读 |
| `device.sort` | integer | 否 | 排序 (格式: int64) |
| `device.status` | integer | 否 | 设备状态 1-未激活，2-在线，3-离线 4-异常(频繁上下线,告警中) (格式: int64) |
| `device.subProtocolConf` | object | 否 | 子协议配置 |
| `device.tags` | object | 否 |  设备tag |
| `device.tenantCode` | string | 否 |  |
| `device.userID` | string | 否 |  |
| `device.version` | string | 否 |  固件版本  读写 |
| `device.withProperties` | object | 否 | 获取的属性列表,如果不传withProperty,则不会返回 |
| `deviceName` | string | 否 |  |
| `id` | string | 否 |  |
| `priority` | integer | 否 | 执行优先级 (格式: int64) |
| `product` | object | 否 |  |
| `product.authMode` | integer | 否 | 认证方式:1:账密认证,2:秘钥认证 (格式: int64) |
| `product.autoRegister` | integer | 否 | 动态注册:1:关闭,2:打开,3:打开并自动创建设备,4:在前面的基础上绑定没有也自动创建 (格式: int64) |
| `product.bindLevel` | integer | 否 | 绑定级别: 1:强绑定(默认,只有用户解绑之后才能绑定) 2:中绑定(可以通过token强制解绑设备) 3:弱绑定(app可以内部解绑被绑定的设备) (格式: int64) |
| `product.category` | object | 否 |  |
| `product.category.avatar` | string | 否 |  品类图片 |
| `product.category.children` | array[ProductCategory] | 否 |  |
| `product.category.code` | string | 否 | 品类编码,全局唯一,创建后不可修改,创建时必填 |
| `product.category.desc` | string | 否 | 描述 |
| `product.category.deviceCount` | integer | 否 | 设备数量 (格式: int64) |
| `product.category.id` | string | 否 |  |
| `product.category.idPath` | array[integer] | 否 | 只读 |
| `product.category.isLeaf` | integer | 否 | 格式: int64 |
| `product.category.name` | string | 否 | 品类名称 |
| `product.category.order` | integer | 否 | 排序,必须大于等于1,为0代表不修改 (格式: int64) |
| `product.category.parentID` | string | 否 |  |
| `product.categoryID` | string | 否 | 产品品类 |
| `product.compatMode` | integer | 否 | 协议兼容模式 1:旧productID(大写D) 2:新productId(小写d,默认) (格式: int64) |
| `product.config` | object | 否 |  |
| `product.config.devInit` | object | 否 |  |
| `product.config.productID` | string | 是 |  |
| `product.createdTime` | string | 否 | 创建时间 只读 |
| `product.customUi` | object | 否 | 自定义ui,key是端的类型(web-client  mini-client) value是以下类型的对象{version:123(版本号,只读),isUpdateUi:bool(是否更新ui),path:string(前端路径,如果需要修改,需要将isUpdateUi置为true并在这个参数中传入压缩包的filePath)} |
| `product.defaultAgentId` | integer | 否 | 默认绑定的 Agent ID，0 表示未绑定 (格式: int64) |
| `product.desc` | string | 否 | 描述 |
| `product.deviceSchemaMode` | integer | 否 |  设备物模型模式:1:手动创建,2:设备自动创建 3: 设备自动创建及上报无定义自动创建 4: 设备自动创建及上报无定义自动创建(数字类型只使用浮点) (格式: int64) |
| `product.deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `product.netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `product.onlineHandle` | integer | 否 | 在线处理:1: 自动 2: 永远在线 (格式: int64) |
| `product.productID` | string | 否 | 产品id 只读 |
| `product.productImg` | string | 否 | 产品图片 |
| `product.productName` | string | 否 | 产品名称 |
| `product.protocol` | object | 否 |  |
| `product.protocol.code` | string | 否 |   iThings,iThings-thingsboard,wumei,aliyun,huaweiyun,tuya |
| `product.protocol.configFields` | array[ProtocolConfigField] | 否 | 配置字段列表,没有可以不传 |
| `product.protocol.containerPort` | integer | 否 |  容器端口 (格式: int64) |
| `product.protocol.deployStatus` | string | 否 |  部署状态: notInstalled/downloaded/running/stopped |
| `product.protocol.desc` | string | 否 |  |
| `product.protocol.deviceFields` | array[ProtocolConfigField] | 否 | 设备级的配置字段列表,没有可以不传 |
| `product.protocol.endpoints` | array[string] | 否 |  |
| `product.protocol.envVars` | object | 否 |  环境变量 |
| `product.protocol.etcdKey` | string | 否 |  |
| `product.protocol.id` | string | 否 |  |
| `product.protocol.image` | string | 否 |  镜像地址 |
| `product.protocol.isBuiltin` | integer | 否 |  是否内置 1:是 2:否 (格式: int64) |
| `product.protocol.isEnableSyncDevice` | integer | 是 | 是否支持主动同步设备信息 (格式: int64) |
| `product.protocol.isEnableSyncProduct` | integer | 是 | 是否支持主动同步产品信息 (格式: int64) |
| `product.protocol.name` | string | 否 |  |
| `product.protocol.productFields` | array[ProtocolConfigField] | 否 | 产品级的配置字段列表,没有可以不传 |
| `product.protocol.transProtocol` | string | 否 |  传输协议: mqtt,tcp,udp |
| `product.protocol.type` | string | 否 | 协议类型: 普通设备(默认):normal  音视频:media |
| `product.protocol.version` | string | 否 |  镜像版本 |
| `product.protocolCode` | string | 否 | 协议code,默认urMqtt,设备下发只会发送给主协议  urMqtt,urHttp,wumei,aliyun,huaweiyun,tuya |
| `product.protocolConf` | object | 否 | 协议配置 |
| `product.sceneMode` | string | 否 | 场景模式 读写类型: r(只读) rw(可读可写) none(不参与场景) |
| `product.secret` | string | 否 | 动态注册产品秘钥 只读 |
| `product.status` | integer | 否 | 产品状态 1:启用 2:禁用 3:开发中 (格式: int64) |
| `product.subProtocolCode` | string | 否 | 子协议,主协议和子协议传输类型必须不相同, 设备控制下发只会发送给主协议, 当设备是音视频设备但是控制协议需要单独走的时候就可以把主协议定义为普通协议,子协议定义为音视频协议,这样就能实现音视频走音视频协议,控制走子协议 |
| `product.subProtocolConf` | object | 否 | 子协议协议配置 |
| `product.tags` | object | 否 |  产品tag |
| `product.tenantCode` | string | 否 | 企业号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `product.trialTime` | string | 否 | 试用时间(单位为天,为0不限制) |
| `productID` | string | 否 |  |
| `script` | object | 否 |  |
| `script.createdTime` | string | 否 | 创建时间 |
| `script.desc` | string | 否 |  |
| `script.id` | string | 否 |  |
| `script.name` | string | 否 |  |
| `script.script` | string | 否 | 脚本内容 |
| `script.status` | integer | 否 | 状态:是否启用 (格式: int64) |
| `script.tenantCode` | string | 否 | 企业号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `script.triggerDir` | integer | 否 | 1:up 2:down (格式: int64) |
| `script.triggerHandle` | string | 否 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `script.triggerTimer` | integer | 否 | 收到前处理1:before 2:after (格式: int64) |
| `script.triggerType` | string | 否 |  操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |
| `scriptID` | string | 否 |  |
| `status` | integer | 否 | 状态:是否启用 (格式: int64) |
| `tenantCode` | string | 否 | 企业号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `triggerSrc` | integer | 否 | product:1 device:2 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "device": {
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
  "deviceName": "string",
  "id": "string",
  "priority": 1,
  "product": {
    "authMode": 1,
    "autoRegister": 1,
    "bindLevel": 1,
    "category": {
      "avatar": "string",
      "children": [
        {
          "avatar": "string",
          "children": [],
          "code": "string",
          "desc": "string",
          "deviceCount": 1,
          "id": "string",
          "idPath": [],
          "isLeaf": 1,
          "name": "示例名称",
          "order": 1,
          "parentID": "string"
        }
      ],
      "code": "string",
      "desc": "string",
      "deviceCount": 1,
      "id": "string",
      "idPath": [
        1
      ],
      "isLeaf": 1,
      "name": "示例名称",
      "order": 1,
      "parentID": "string"
    },
    "categoryID": "string",
    "compatMode": 1,
    "config": {
      "devInit": {
        "areaID": "string",
        "deptID": "string",
        "projectID": "string",
        "tenantCode": "string"
      },
      "productID": "string"
    },
    "createdTime": "2026-01-01T00:00:00Z",
    "customUi": {},
    "defaultAgentId": 1,
    "desc": "string",
    "deviceSchemaMode": 1,
    "deviceType": 1,
    "netType": 1,
    "onlineHandle": 1,
    "productID": "string",
    "productImg": "string",
    "productName": "示例名称",
    "protocol": {
      "code": "string",
      "configFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "containerPort": 1,
      "deployStatus": "string",
      "desc": "string",
      "deviceFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "endpoints": [
        "string"
      ],
      "envVars": {},
      "etcdKey": "string",
      "id": "string",
      "image": "string",
      "isBuiltin": 1,
      "isEnableSyncDevice": 1,
      "isEnableSyncProduct": 1,
      "name": "string",
      "productFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "transProtocol": "string",
      "type": "string",
      "version": "string"
    },
    "protocolCode": "string",
    "protocolConf": {},
    "sceneMode": "string",
    "secret": "string",
    "status": 1,
    "subProtocolCode": "string",
    "subProtocolConf": {},
    "tags": {},
    "tenantCode": "string",
    "trialTime": "2026-01-01T00:00:00Z"
  },
  "productID": "string",
  "script": {
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "name": "string",
    "script": "string",
    "status": 1,
    "tenantCode": "string",
    "triggerDir": 1,
    "triggerHandle": "string",
    "triggerTimer": 1,
    "triggerType": "string"
  },
  "scriptID": "string",
  "status": 1,
  "tenantCode": "string",
  "triggerSrc": 1
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
ur api /api/v1/things/protocol/script/device/update \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "device": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [{"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": 1, "longitude": 1}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "string", "areaIDPath": "string", "areaImg": "string", "areaName": "示例名称", "areaNamePath": "string", "children": [], "configFile": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "groupCount": 1, "isLeaf": 1, "isSysCreated": 1, "lowerLevelCount": 1, "parentAreaID": "string", "position": {"latitude": "...", "longitude": "..."}, "projectID": "string", "sort": 1, "tags": {}, "tenantCode": "string", "useBy": "string"}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "string", "idPath": "string"}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "string", "address": "string", "area": {"areaID": "...", "areaIDPath": "...", "areaImg": "...", "areaName": "...", "areaNamePath": "...", "children": "...", "configFile": "...", "createdTime": "...", "desc": "...", "deviceCount": "...", "devices": "...", "groupCount": "...", "isLeaf": "...", "isSysCreated": "...", "lowerLevelCount": "...", "parentAreaID": "...", "position": "...", "projectID": "...", "sort": "...", "tags": "...", "tenantCode": "...", "useBy": "..."}, "areaID": "string", "categoryID": "string", "cert": "string", "cloneID": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceAlias": "示例名称", "deviceImg": "string", "deviceName": "示例名称", "deviceType": 1, "distributor": {"id": "...", "idPath": "..."}, "expTime": "2026-01-01T00:00:00Z", "file": "string", "firstBind": "string", "firstLogin": "2026-01-01T00:00:00Z", "gateway": {"adcode": "...", "address": "...", "area": "...", "areaID": "...", "categoryID": "...", "cert": "...", "cloneID": "...", "createdTime": "...", "desc": "...", "deviceAlias": "...", "deviceImg": "...", "deviceName": "...", "deviceType": "...", "distributor": "...", "expTime": "...", "file": "...", "firstBind": "...", "firstLogin": "...", "gateway": "...", "groupIDs": "...", "groupPurpose": "...", "groups": "...", "hardInfo": "...", "iccid": "...", "id": "...", "imei": "...", "isEnable": "...", "isOnline": "...", "lastBind": "...", "lastIp": "...", "lastLocalIp": "...", "lastLogin": "...", "lastOffline": "...", "logLevel": "...", "mac": "...", "mobileOperator": "...", "needConfirmVersion": "...", "netType": "...", "owner": "...", "phone": "...", "position": "...", "productID": "...", "productImg": "...", "productName": "...", "profiles": "...", "projectID": "...", "protocolConf": "...", "ratedPower": "...", "remoteAddr": "...", "rssi": "...", "schemaAlias": "...", "secret": "...", "softInfo": "...", "sort": "...", "status": "...", "subProtocolConf": "...", "tags": "...", "tenantCode": "...", "userID": "...", "version": "...", "withProperties": "..."}, "groupIDs": [], "groupPurpose": "string", "groups": [], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "...", "createdTime": "...", "email": "...", "lastIP": "...", "nickName": "...", "phone": "...", "regIP": "...", "userID": "...", "userName": "..."}, "phone": "string", "position": {"latitude": "...", "longitude": "..."}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "...", "name": "...", "parentID": "...", "purpose": "..."}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "groupIDs": ["string"], "groupPurpose": "string", "groups": [{"id": "string", "name": "string", "parentID": "string", "purpose": "string"}], "hardInfo": "string", "iccid": "string", "id": "string", "imei": "string", "isEnable": 1, "isOnline": 1, "lastBind": "2026-01-01T00:00:00Z", "lastIp": "string", "lastLocalIp": "string", "lastLogin": "2026-01-01T00:00:00Z", "lastOffline": "2026-01-01T00:00:00Z", "logLevel": "string", "mac": "string", "mobileOperator": 1, "needConfirmVersion": "string", "netType": 1, "owner": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "phone": "string", "position": {"latitude": 1, "longitude": 1}, "productID": "string", "productImg": "string", "productName": "string", "profiles": {}, "projectID": "string", "protocolConf": {}, "ratedPower": 1, "remoteAddr": "string", "rssi": 1, "schemaAlias": {}, "secret": "string", "softInfo": "string", "sort": 1, "status": 1, "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "userID": "string", "version": "string", "withProperties": {}}, "deviceName": "string", "id": "string", "priority": 1, "product": {"authMode": 1, "autoRegister": 1, "bindLevel": 1, "category": {"avatar": "string", "children": [{"avatar": "string", "children": [], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}, "categoryID": "string", "compatMode": 1, "config": {"devInit": {"areaID": "string", "deptID": "string", "projectID": "string", "tenantCode": "string"}, "productID": "string"}, "createdTime": "2026-01-01T00:00:00Z", "customUi": {}, "defaultAgentId": 1, "desc": "string", "deviceSchemaMode": 1, "deviceType": 1, "netType": 1, "onlineHandle": 1, "productID": "string", "productImg": "string", "productName": "示例名称", "protocol": {"code": "string", "configFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "containerPort": 1, "deployStatus": "string", "desc": "string", "deviceFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "endpoints": ["string"], "envVars": {}, "etcdKey": "string", "id": "string", "image": "string", "isBuiltin": 1, "isEnableSyncDevice": 1, "isEnableSyncProduct": 1, "name": "string", "productFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "transProtocol": "string", "type": "string", "version": "string"}, "protocolCode": "string", "protocolConf": {}, "sceneMode": "string", "secret": "string", "status": 1, "subProtocolCode": "string", "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "trialTime": "2026-01-01T00:00:00Z"}, "productID": "string", "script": {"createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "name": "string", "script": "string", "status": 1, "tenantCode": "string", "triggerDir": 1, "triggerHandle": "string", "triggerTimer": 1, "triggerType": "string"}, "scriptID": "string", "status": 1, "tenantCode": "string", "triggerSrc": 1}'
```
