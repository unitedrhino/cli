# ur-protocol things/protocol/sync

设备同步(如果该协议不支持会返回不支持) 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/protocol/sync/device` | 设备同步(如果该协议不支持会返回不支持) | admin |
| POST | `/api/v1/things/protocol/sync/product` | 产品同步(如果该协议不支持会返回不支持) | admin |

## 详细说明

### POST `/api/v1/things/protocol/sync/device`

**说明**: 设备同步(如果该协议不支持会返回不支持)

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |   协议code |
| `productID` | string | 是 |  |

**请求示例**:
```json
{
  "code": "string",
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "deviceNames": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/sync/device \
  --body '{"code": "string", "productID": "string"}'
```

### POST `/api/v1/things/protocol/sync/product`

**说明**: 产品同步(如果该协议不支持会返回不支持)

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |

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
ur api /api/v1/things/protocol/sync/product \
  --body '{"code": "string"}'
```
