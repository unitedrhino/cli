# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/auth/access` | 设备操作认证 | device |
| POST | `/api/v1/things/device/auth/login` | 设备登录认证 | device |
| POST | `/api/v1/things/device/auth/register` | 设备自动注册 | device |
| POST | `/api/v1/things/device/auth/root-check` | 鉴定是否是root账号 | device |
| POST | `/api/v1/things/device/auth5/access` | 设备操作认证 | device |
| POST | `/api/v1/things/device/auth5/login` | 设备登录认证 | device |
| POST | `/api/v1/things/device/edge/send/{handle}/{type}` | 设备使用http协议用云端交互,需要在http头中带上mqtt的账号密码(basic auth) | device |
| POST | `/api/v1/things/device/edge/upload-file` | 设备文件直传,需要在http头中带上mqtt的账号密码(basic auth) | device |
| POST | `/api/v1/things/device/gateway/batch-create` | 添加网关子设备 | admin |
| POST | `/api/v1/things/device/gateway/batch-delete` | 解绑子设备 | admin |
| POST | `/api/v1/things/device/gateway/get-list` | 获取子设备列表 | admin |
| POST | `/api/v1/things/device/group/batch-create` | 将设备加到多个分组中 | admin |
| POST | `/api/v1/things/device/group/batch-delete` | 删除设备所在分组 | admin |
| POST | `/api/v1/things/device/group/batch-update` | 更新设备所在分组 | admin |
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
| POST | `/api/v1/things/device/info/batch-update` | 批量更新设备 | admin |
| POST | `/api/v1/things/device/info/create` | 新增设备 | admin |
| POST | `/api/v1/things/device/info/update` | 更新设备 | admin |
| POST | `/api/v1/things/device/profile/delete` | 删除设备配置 | admin |
| POST | `/api/v1/things/device/profile/get-list` | 获取设备配置列表 | admin |
| POST | `/api/v1/things/device/profile/get-one` | 获取设备配置详情 | admin |
| POST | `/api/v1/things/device/profile/update` | 更新设备配置 | admin |
| POST | `/api/v1/things/device/schema/batch-create` | 批量创建设备物模型 | admin |
| POST | `/api/v1/things/device/schema/batch-delete` | 批量删除设备物模型 | admin |
| POST | `/api/v1/things/device/schema/create` | 创建设备物模型 | admin |
| POST | `/api/v1/things/device/schema/get-list` | 获取设备物模型列表 | admin |
| POST | `/api/v1/things/device/schema/tsl-read` | 获取设备物模型tsl | admin |
| POST | `/api/v1/things/device/schema/update` | 更新设备物模型 | admin |
| POST | `/api/v1/things/device/version/get-list` | 获取设备模块版本列表 | admin |
| POST | `/api/v1/things/device/version/get-one` | 获取设备模块版本详情 | admin |

## 端点详情

### POST `/api/v1/things/device/auth/access`

**说明**: 设备操作认证

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `access` | string | 否 | 操作 |
| `clientID` | string | 否 | clientID |
| `ip` | string | 否 | 访问的ip地址 |
| `topic` | string | 否 | 主题 |
| `username` | string | 否 | 用户名 |

**请求示例**:
```json
{
  "access": "string",
  "clientID": "string",
  "ip": "string",
  "topic": "string",
  "username": "string"
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
ur api /api/v1/things/device/auth/access \
  --body '{"access": "string", "clientID": "string", "ip": "string", "topic": "string", "username": "string"}'
```

### POST `/api/v1/things/device/auth/login`

**说明**: 设备登录认证

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `certificate` | string | 否 | 客户端证书 base64后传过来 |
| `clientID` | string | 是 | clientID |
| `ip` | string | 是 | 访问的ip地址 |
| `password` | string | 否 | 密码 |
| `username` | string | 是 | 用户名 |

**请求示例**:
```json
{
  "certificate": "string",
  "clientID": "string",
  "ip": "string",
  "password": "******",
  "username": "string"
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
ur api /api/v1/things/device/auth/login \
  --body '{"certificate": "string", "clientID": "string", "ip": "string", "password": "******", "username": "string"}'
```

### POST `/api/v1/things/device/auth/register`

**说明**: 设备自动注册

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名称 读写 |
| `nonce` | integer | 是 | 随机数 (格式: int64) |
| `productId` | string | 是 | 产品id 只读 |
| `retEnc` | string | 否 | 返回使用的加密方式: aes128cbc(默认) aes128ecb |
| `signature` | string | 是 | 签名信息 |
| `tenantCode` | string | 否 | 企业号（可选）：产品为common企业或企业号与产品一致时生效 |
| `timestamp` | integer | 是 | 秒级时间戳 (格式: int64) |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "nonce": 1,
  "productId": "string",
  "retEnc": "string",
  "signature": "string",
  "tenantCode": "string",
  "timestamp": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "len": 1,
    "payload": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth/register \
  --body '{"deviceName": "示例名称", "nonce": 1, "productId": "string", "retEnc": "string", "signature": "string", "tenantCode": "string", "timestamp": 1}'
```

### POST `/api/v1/things/device/auth/root-check`

**说明**: 鉴定是否是root账号

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `certificate` | string | 否 | 客户端证书 |
| `clientID` | string | 否 | clientID |
| `ip` | string | 否 | 访问的ip地址 |
| `password` | string | 否 | 密码 |
| `username` | string | 否 | 用户名 |

**请求示例**:
```json
{
  "certificate": "string",
  "clientID": "string",
  "ip": "string",
  "password": "******",
  "username": "string"
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
ur api /api/v1/things/device/auth/root-check \
  --body '{"certificate": "string", "clientID": "string", "ip": "string", "password": "******", "username": "string"}'
```

### POST `/api/v1/things/device/auth5/access`

**说明**: 设备操作认证

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `action` | string | 否 | 操作 |
| `clientID` | string | 否 | clientID |
| `ip` | string | 否 | 访问的ip地址 |
| `topic` | string | 否 | 主题 |
| `username` | string | 否 | 用户名 |

**请求示例**:
```json
{
  "action": "string",
  "clientID": "string",
  "ip": "string",
  "topic": "string",
  "username": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "result": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth5/access \
  --body '{"action": "string", "clientID": "string", "ip": "string", "topic": "string", "username": "string"}'
```

### POST `/api/v1/things/device/auth5/login`

**说明**: 设备登录认证

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `certificate` | string | 否 | 客户端证书 base64后传过来 |
| `clientID` | string | 是 | clientID |
| `ip` | string | 是 | 访问的ip地址 |
| `password` | string | 否 | 密码 |
| `username` | string | 是 | 用户名 |

**请求示例**:
```json
{
  "certificate": "string",
  "clientID": "string",
  "ip": "string",
  "password": "******",
  "username": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "is_superuser": true,
    "result": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth5/login \
  --body '{"certificate": "string", "clientID": "string", "ip": "string", "password": "******", "username": "string"}'
```

### POST `/api/v1/things/device/edge/send/{handle}/{type}`

**说明**: 设备使用http协议用云端交互,需要在http头中带上mqtt的账号密码(basic auth)

**权限**: device

**路径参数**:

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| `handle` | path | string | 是 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `type` | path | string | 是 | 操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |
| `productID` | header | string | 否 | 如果网关类型要操作子设备的topic,需要指定子设备的产品ID和设备ID,如果没有填写则默认是账号里的设备 |
| `deviceName` | header | string | 否 | 如果网关类型要操作子设备的topic,需要指定子设备的产品ID和设备ID,如果没有填写则默认是账号里的设备 |
| `authorization` | header | string | 是 | basic auth  账号密码使用mqtt格式生成的账号密码 |

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/edge/send/{handle}/{type} \
  --body '{}'
```

### POST `/api/v1/things/device/edge/upload-file`

**说明**: 设备文件直传,需要在http头中带上mqtt的账号密码(basic auth)

**权限**: device

**路径参数**:

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| `authorization` | header | string | 是 |  basic auth 账号密码使用mqtt格式生成的账号密码 |

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authorization` | string | 是 | // basic auth 账号密码使用mqtt格式生成的账号密码 |
| `file` | string | 是 | 格式: binary |

**请求示例**:
```json
{
  "authorization": "******",
  "file": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "filePath": "string",
    "fileUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/edge/upload-file \
  --body '{"authorization": "******", "file": "string"}'
```

### POST `/api/v1/things/device/gateway/batch-create`

**说明**: 添加网关子设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `gateWayProductID` | string | 是 | 产品ID |
| `gateWaydeviceName` | string | 是 | 设备名称 |
| `list` | array[DeviceCore] | 否 | 分组tag |

**请求示例**:
```json
{
  "gateWayProductID": "string",
  "gateWaydeviceName": "示例名称",
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
ur api /api/v1/things/device/gateway/batch-create \
  --body '{"gateWayProductID": "string", "gateWaydeviceName": "示例名称", "list": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}]}'
```

### POST `/api/v1/things/device/gateway/batch-delete`

**说明**: 解绑子设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `gateWayProductID` | string | 是 | 产品ID |
| `gateWaydeviceName` | string | 是 | 设备名称 |
| `list` | array[DeviceCore] | 否 | 分组tag |

**请求示例**:
```json
{
  "gateWayProductID": "string",
  "gateWaydeviceName": "示例名称",
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
ur api /api/v1/things/device/gateway/batch-delete \
  --body '{"gateWayProductID": "string", "gateWaydeviceName": "示例名称", "list": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}]}'
```

### POST `/api/v1/things/device/gateway/get-list`

**说明**: 获取子设备列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `gateWayProductID` | string | 是 | 产品ID |
| `gateWaydeviceName` | string | 是 | 设备名称 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "gateWayProductID": "string",
  "gateWaydeviceName": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  }
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
ur api /api/v1/things/device/gateway/get-list \
  --body '{"gateWayProductID": "string", "gateWaydeviceName": "示例名称", "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/things/device/group/batch-create`

**说明**: 将设备加到多个分组中

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名称 |
| `groupIDs` | array[string] | 是 |  |
| `productID` | string | 是 | 产品id |
| `purpose` | string | 否 | 用途 不填默认为default |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "groupIDs": [
    "string"
  ],
  "productID": "string",
  "purpose": "string"
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
ur api /api/v1/things/device/group/batch-create \
  --body '{"deviceName": "示例名称", "groupIDs": ["string"], "productID": "string", "purpose": "string"}'
```

### POST `/api/v1/things/device/group/batch-delete`

**说明**: 删除设备所在分组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名称 |
| `groupIDs` | array[string] | 是 |  |
| `productID` | string | 是 | 产品id |
| `purpose` | string | 否 | 用途 不填默认为default |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "groupIDs": [
    "string"
  ],
  "productID": "string",
  "purpose": "string"
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
ur api /api/v1/things/device/group/batch-delete \
  --body '{"deviceName": "示例名称", "groupIDs": ["string"], "productID": "string", "purpose": "string"}'
```

### POST `/api/v1/things/device/group/batch-update`

**说明**: 更新设备所在分组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名称 |
| `groupIDs` | array[string] | 是 |  |
| `productID` | string | 是 | 产品id |
| `purpose` | string | 否 | 用途 不填默认为default |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "groupIDs": [
    "string"
  ],
  "productID": "string",
  "purpose": "string"
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
ur api /api/v1/things/device/group/batch-update \
  --body '{"deviceName": "示例名称", "groupIDs": ["string"], "productID": "string", "purpose": "string"}'
```

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
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
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
    "page": 1,
    "pageSize": 1
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
  --body '{"areaID": "string", "areaIDPath": "string", "areaIDPaths": ["string"], "areaIDs": ["string"], "deviceAlias": "示例名称", "deviceName": "string", "deviceNameOrAlias": "示例名称", "deviceNames": ["string"], "deviceTypes": [1], "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "expTime": {"cmpType": "string", "value": "string"}, "gateway": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "groupID": "string", "groupIDPath": "string", "groupIDPaths": ["string"], "groupIDs": ["string"], "groupName": "string", "groupPurpose": "string", "hasOwner": 1, "iccid": "string", "isOnline": 1, "isOnlyCore": true, "netType": 1, "notAreaID": "string", "notGroupID": "string", "notVersion": "string", "page": {"page": 1, "pageSize": 1}, "parentGroupID": "string", "position": {"latitude": 1, "longitude": 1}, "productCategoryCode": "string", "productCategoryID": "string", "productCategoryIDs": [1], "productID": "string", "productIDs": ["string"], "property": {}, "propertyIgnoreEmpty": true, "range": 1, "ratedPower": {"cmpType": "string", "value": "string"}, "rssi": {"cmpType": "string", "value": "string"}, "status": 1, "statuses": [1], "tags": {}, "tenantCode": "string", "userID": "string", "versions": ["string"], "withArea": true, "withCollect": 1, "withGateway": true, "withGroups": ["string"], "withOwner": true, "withProfiles": ["string"], "withProperties": ["string"], "withShared": 1}'
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

### POST `/api/v1/things/device/info/batch-update`

**说明**: 批量更新设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 |  |
| `devices` | array[DeviceCore] | 是 |  |
| `ratedPower` | integer | 否 | 格式: int64 |

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
  ],
  "ratedPower": 1
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
ur api /api/v1/things/device/info/batch-update \
  --body '{"areaID": "string", "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "ratedPower": 1}'
```

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

### POST `/api/v1/things/device/profile/delete`

**说明**: 删除设备配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `device` | object | 是 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `device.productName` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
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
ur api /api/v1/things/device/profile/delete \
  --body '{"code": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}}'
```

### POST `/api/v1/things/device/profile/get-list`

**说明**: 获取设备配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `codes` | array[string] | 否 |  |
| `device` | object | 是 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `device.productName` | string | 否 |  |

**请求示例**:
```json
{
  "codes": [
    "string"
  ],
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
  "data": {
    "profiles": [
      {
        "code": "string",
        "device": {
          "deviceName": "示例名称",
          "productID": "string",
          "productName": "string"
        },
        "params": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/profile/get-list \
  --body '{"codes": ["string"], "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}}'
```

### POST `/api/v1/things/device/profile/get-one`

**说明**: 获取设备配置详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `device` | object | 是 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `device.productName` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
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
  "data": {
    "code": "string",
    "device": {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
    },
    "params": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/profile/get-one \
  --body '{"code": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}}'
```

### POST `/api/v1/things/device/profile/update`

**说明**: 更新设备配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `device` | object | 是 |  |
| `device.deviceName` | string | 是 | 设备名称 |
| `device.productID` | string | 是 | 产品ID |
| `device.productName` | string | 否 |  |
| `params` | string | 是 |  |

**请求示例**:
```json
{
  "code": "string",
  "device": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  },
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
ur api /api/v1/things/device/profile/update \
  --body '{"code": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "params": "string"}'
```

### POST `/api/v1/things/device/schema/batch-create`

**说明**: 批量创建设备物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  |
| `identifiers` | array[string] | 否 | 导入的标识符列表 |
| `list` | array[DeviceSchema] | 否 | 如果没有填标识符列表取这里 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "deviceName": "string",
  "identifiers": [
    "string"
  ],
  "list": [
    {
      "affordance": "string",
      "desc": "string",
      "deviceName": "string",
      "extendConfig": "string",
      "funcGroup": 1,
      "identifier": "string",
      "isCanSceneLinkage": 1,
      "name": "示例名称",
      "order": 1,
      "productID": "string",
      "required": 1,
      "tag": 1,
      "type": 1
    }
  ],
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
ur api /api/v1/things/device/schema/batch-create \
  --body '{"deviceName": "string", "identifiers": ["string"], "list": [{"affordance": "string", "desc": "string", "deviceName": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}], "productID": "string"}'
```

### POST `/api/v1/things/device/schema/batch-delete`

**说明**: 批量删除设备物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  |
| `identifiers` | array[string] | 是 | 物模型ID列表 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "deviceName": "string",
  "identifiers": [
    "string"
  ],
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
ur api /api/v1/things/device/schema/batch-delete \
  --body '{"deviceName": "string", "identifiers": ["string"], "productID": "string"}'
```

### POST `/api/v1/things/device/schema/create`

**说明**: 创建设备物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `affordance` | string | 否 |  |
| `desc` | string | 否 | 描述 |
| `deviceName` | string | 否 |  |
| `extendConfig` | string | 否 | 拓展参数 |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifier` | string | 是 | 标识符 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 | 功能名称 |
| `order` | integer | 否 |  排序 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `required` | integer | 否 | 是否必须 1:是 2:否 (格式: int64) |
| `tag` | integer | 是 | 物模型标签 1:自定义 2:可选 3:必选  必选不可删除 (格式: int64) |
| `type` | integer | 是 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

**请求示例**:
```json
{
  "affordance": "string",
  "desc": "string",
  "deviceName": "string",
  "extendConfig": "string",
  "funcGroup": 1,
  "identifier": "string",
  "isCanSceneLinkage": 1,
  "name": "示例名称",
  "order": 1,
  "productID": "string",
  "required": 1,
  "tag": 1,
  "type": 1
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
ur api /api/v1/things/device/schema/create \
  --body '{"affordance": "string", "desc": "string", "deviceName": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}'
```

### POST `/api/v1/things/device/schema/get-list`

**说明**: 获取设备物模型列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `controlMode` | integer | 否 | 控制模式: 1: 可以群控,可以单控  2:只能单控 (格式: int64) |
| `deviceName` | string | 是 |  |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifiers` | array[string] | 否 | 过滤标识符列表 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 是 | 产品id |
| `propertyMode` | string | 否 | 属性类型可用 读写类型: 1:r(只读) 2:rw(可读可写) |
| `tag` | integer | 否 | 过滤条件: 物模型标签 1:自定义 2:可选 3:必选 (格式: int64) |
| `type` | integer | 否 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |
| `types` | array[integer] | 否 |  |
| `withProductSchema` | boolean | 否 | 同时获取产品物模型 (格式: boolean) |

**请求示例**:
```json
{
  "controlMode": 1,
  "deviceName": "string",
  "funcGroup": 1,
  "identifiers": [
    "string"
  ],
  "isCanSceneLinkage": 1,
  "name": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "productID": "string",
  "propertyMode": "string",
  "tag": 1,
  "type": 1,
  "types": [
    1
  ],
  "withProductSchema": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "affordance": "string",
        "desc": "string",
        "deviceName": "string",
        "extendConfig": "string",
        "funcGroup": 1,
        "identifier": "string",
        "isCanSceneLinkage": 1,
        "name": "示例名称",
        "order": 1,
        "productID": "string",
        "required": 1,
        "tag": 1,
        "type": 1
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
ur api /api/v1/things/device/schema/get-list \
  --body '{"controlMode": 1, "deviceName": "string", "funcGroup": 1, "identifiers": ["string"], "isCanSceneLinkage": 1, "name": "string", "page": {"page": 1, "pageSize": 1}, "productID": "string", "propertyMode": "string", "tag": 1, "type": 1, "types": [1], "withProductSchema": true}'
```

### POST `/api/v1/things/device/schema/tsl-read`

**说明**: 获取设备物模型tsl

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "deviceName": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "tsl": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/schema/tsl-read \
  --body '{"deviceName": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/schema/update`

**说明**: 更新设备物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `affordance` | string | 否 |  |
| `desc` | string | 否 | 描述 |
| `deviceName` | string | 否 |  |
| `extendConfig` | string | 否 | 拓展参数 |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifier` | string | 是 | 标识符 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 | 功能名称 |
| `order` | integer | 否 |  排序 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `required` | integer | 否 | 是否必须 1:是 2:否 (格式: int64) |
| `tag` | integer | 是 | 物模型标签 1:自定义 2:可选 3:必选  必选不可删除 (格式: int64) |
| `type` | integer | 是 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

**请求示例**:
```json
{
  "affordance": "string",
  "desc": "string",
  "deviceName": "string",
  "extendConfig": "string",
  "funcGroup": 1,
  "identifier": "string",
  "isCanSceneLinkage": 1,
  "name": "示例名称",
  "order": 1,
  "productID": "string",
  "required": 1,
  "tag": 1,
  "type": 1
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
ur api /api/v1/things/device/schema/update \
  --body '{"affordance": "string", "desc": "string", "deviceName": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}'
```

### POST `/api/v1/things/device/version/get-list`

**说明**: 获取设备模块版本列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `moduleCode` | integer | 否 | 格式: int64 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `productName` | string | 否 |  |

**请求示例**:
```json
{
  "moduleCode": 1,
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "productID": "string",
  "productName": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": "string",
        "moduleCode": 1,
        "productID": "string",
        "productName": "string",
        "version": "string"
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
ur api /api/v1/things/device/version/get-list \
  --body '{"moduleCode": 1, "page": {"page": 1, "pageSize": 1}, "productID": "string", "productName": "string"}'
```

### POST `/api/v1/things/device/version/get-one`

**说明**: 获取设备模块版本详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  |
| `moduleCode` | integer | 否 | 格式: int64 |
| `productID` | string | 否 | 产品id 只读 |
| `productName` | string | 否 |  |

**请求示例**:
```json
{
  "id": "string",
  "moduleCode": 1,
  "productID": "string",
  "productName": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string",
    "moduleCode": 1,
    "productID": "string",
    "productName": "string",
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/version/get-one \
  --body '{"id": "string", "moduleCode": 1, "productID": "string", "productName": "string"}'
```
