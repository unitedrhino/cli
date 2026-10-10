# ur-device things/device/group

将设备加到多个分组中 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/group/batch-create` | 将设备加到多个分组中 | admin |
| POST | `/api/v1/things/device/group/batch-delete` | 删除设备所在分组 | admin |
| POST | `/api/v1/things/device/group/batch-update` | 更新设备所在分组 | admin |

## 详细说明

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
