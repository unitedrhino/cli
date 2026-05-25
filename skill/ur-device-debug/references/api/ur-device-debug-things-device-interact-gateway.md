# ur-device-debug things/device/interact/gateway

请求网关上报拓扑关系 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/interact/gateway-get-found-send` | 请求网关上报拓扑关系 | admin |
| POST | `/api/v1/things/device/interact/gateway-notify-bind-send` | 通知网关绑定子设备 | admin |

## 详细说明

### POST `/api/v1/things/device/interact/gateway-get-found-send`

**说明**: 请求网关上报拓扑关系

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名(必填) |
| `productID` | string | 是 | 产品id (必填) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/interact/gateway-get-found-send \
  --body '{"deviceName": "string", "productID": "string"}'
```

### POST `/api/v1/things/device/interact/gateway-notify-bind-send`

**说明**: 通知网关绑定子设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `gateway` | object | 是 |  |
| `gateway.deviceName` | string | 是 | 设备名称 |
| `gateway.productID` | string | 是 | 产品ID |
| `gateway.productName` | string | 否 |  |
| `subDevices` | array[DeviceCore] | 是 | 如果是不同的产品,则传这个字段,上面两个参数填了优先使用 |

**请求示例**:
```json
{
  "gateway": {
    "deviceName": "示例名称",
    "productID": "string",
    "productName": "string"
  },
  "subDevices": [
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
ur api /api/v1/things/device/interact/gateway-notify-bind-send \
  --body '{"gateway": {"deviceName": "示例名称", "productID": "string", "productName": "string"}, "subDevices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}]}'
```
