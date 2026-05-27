# ur-iot-user things/user/device/share

接受批量分享设备 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/user/device/share/batch-accept` | 接受批量分享设备 | all |
| POST | `/api/v1/things/user/device/share/batch-create` | 生成批量分享设备二维码 | all |
| POST | `/api/v1/things/user/device/share/batch-delete` | 批量取消分享设备 | all |
| POST | `/api/v1/things/user/device/share/batch-get-list` | 获取批量分享的设备列表 | all |
| POST | `/api/v1/things/user/device/share/create` | 分享设备 | all |
| POST | `/api/v1/things/user/device/share/delete` | 取消分享设备 | all |
| POST | `/api/v1/things/user/device/share/get-list` | 获取分享设备列表 | all |
| POST | `/api/v1/things/user/device/share/get-one` | 获取分享设备详情 | all |
| POST | `/api/v1/things/user/device/share/update` | 更新分享设备信息 | all |

## 详细说明

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
      "productID": "string",
      "productName": "string"
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
  --body '{"devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "shareToken": "string"}'
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
      "productID": "string",
      "productName": "string"
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
  --body '{"accessPerm": {}, "authType": 1, "createdTime": "string", "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "expTime": 1, "schemaPerm": {}, "useBy": "string"}'
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
| `device.productName` | string | 否 |  |
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
    "productID": "string",
    "productName": "string"
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
  --body '{"accessPerm": {}, "authType": 1, "createdTime": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "expTime": "2026-01-01T00:00:00Z", "id": "string", "projectID": "string", "schemaPerm": {}, "sharedUserAccount": "string", "sharedUserID": "string", "useBy": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}}'
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
| `device.productName` | string | 否 |  |
| `id` | string | 否 |  |
| `useBy` | string | 否 | 用途 |

**请求示例**:
```json
{
  "device": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
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
  --body '{"device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "id": "string", "useBy": "string"}'
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
| `device.productName` | string | 否 |  |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `withUser` | boolean | 否 | 同时返回用户信息 (格式: boolean) |

**请求示例**:
```json
{
  "device": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  },
  "page": {
    "page": 1,
    "pageSize": 1
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
          "productID": "string",
          "productName": "string"
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
  --body '{"device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "page": {"page": 1, "pageSize": 1}, "withUser": true}'
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
| `device.productName` | string | 否 |  |
| `id` | string | 否 |  |
| `useBy` | string | 否 | 用途 |

**请求示例**:
```json
{
  "device": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
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
      "productID": "string",
      "productName": "string"
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
  --body '{"device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "id": "string", "useBy": "string"}'
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
| `device.productName` | string | 否 |  |
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
    "productID": "string",
    "productName": "string"
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
  --body '{"accessPerm": {}, "authType": 1, "createdTime": "string", "device": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "expTime": "2026-01-01T00:00:00Z", "id": "string", "projectID": "string", "schemaPerm": {}, "sharedUserAccount": "string", "sharedUserID": "string", "useBy": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}}'
```
