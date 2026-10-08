# ur-device things/device/ota

批量更新设备

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/info/batch-update` | 批量更新设备 | admin |

## 详细说明

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
      "productID": "string"
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
  --body '{"areaID": "string", "devices": [{"deviceName": "示例名称", "productID": "string"}], "ratedPower": 1}'
```
