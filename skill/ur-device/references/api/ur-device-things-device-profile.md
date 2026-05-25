# ur-device things/device/profile

删除设备配置 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/profile/delete` | 删除设备配置 | admin |
| POST | `/api/v1/things/device/profile/get-list` | 获取设备配置列表 | admin |
| POST | `/api/v1/things/device/profile/get-one` | 获取设备配置详情 | admin |
| POST | `/api/v1/things/device/profile/update` | 更新设备配置 | admin |

## 详细说明

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
