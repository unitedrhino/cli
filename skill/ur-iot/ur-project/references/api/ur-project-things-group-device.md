# ur-project things/group/device

添加分组设备 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/group/device/batch-create` | 添加分组设备 | admin |
| POST | `/api/v1/things/group/device/batch-delete` | 删除分组设备 | admin |
| POST | `/api/v1/things/group/device/batch-update` | 更新分组设备 | admin |

## 详细说明

### POST `/api/v1/things/group/device/batch-create`

**说明**: 添加分组设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupID` | string | 是 | 分组ID |
| `list` | array[DeviceCore] | 否 | 分组tag |

**请求示例**:
```json
{
  "groupID": "string",
  "list": [
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
ur api /api/v1/things/group/device/batch-create \
  --body '{"groupID": "string", "list": [{"deviceName": "示例名称", "productID": "string"}]}'
```

### POST `/api/v1/things/group/device/batch-delete`

**说明**: 删除分组设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupID` | string | 是 | 分组ID |
| `list` | array[DeviceCore] | 否 | 分组tag |

**请求示例**:
```json
{
  "groupID": "string",
  "list": [
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
ur api /api/v1/things/group/device/batch-delete \
  --body '{"groupID": "string", "list": [{"deviceName": "示例名称", "productID": "string"}]}'
```

### POST `/api/v1/things/group/device/batch-update`

**说明**: 更新分组设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupID` | string | 是 | 分组ID |
| `list` | array[DeviceCore] | 否 | 分组tag |

**请求示例**:
```json
{
  "groupID": "string",
  "list": [
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
ur api /api/v1/things/group/device/batch-update \
  --body '{"groupID": "string", "list": [{"deviceName": "示例名称", "productID": "string"}]}'
```
