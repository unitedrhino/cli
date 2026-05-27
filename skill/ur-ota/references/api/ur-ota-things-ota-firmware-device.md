# ur-ota things/ota/firmware/device

取消指定任务下的升级 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/ota/firmware/device/cancel` | 取消指定任务下的升级 | admin |
| POST | `/api/v1/things/ota/firmware/device/confirm` | 确认升级设备 | admin |
| POST | `/api/v1/things/ota/firmware/device/get-list` | 查询升级设备列表 | admin |
| POST | `/api/v1/things/ota/firmware/device/retry` | 重试设备升级 | admin |

## 详细说明

### POST `/api/v1/things/ota/firmware/device/cancel`

**说明**: 取消指定任务下的升级

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceNames` | array[string] | 是 |  |
| `firmwareID` | string | 是 |  |
| `jobID` | string | 是 |  |

**请求示例**:
```json
{
  "deviceNames": [
    "string"
  ],
  "firmwareID": "string",
  "jobID": "string"
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
ur api /api/v1/things/ota/firmware/device/cancel \
  --body '{"deviceNames": ["string"], "firmwareID": "string", "jobID": "string"}'
```

### POST `/api/v1/things/ota/firmware/device/confirm`

**说明**: 确认升级设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  |
| `productID` | string | 是 |  |

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
ur api /api/v1/things/ota/firmware/device/confirm \
  --body '{"deviceName": "string", "productID": "string"}'
```

### POST `/api/v1/things/ota/firmware/device/get-list`

**说明**: 查询升级设备列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 否 |  设备名称 |
| `firmwareID` | string | 否 |  任务状态 |
| `isOnline` | integer | 否 | 设备是否在线 (格式: int64) |
| `jobID` | string | 否 |  作业ID |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `srcVersion` | string | 否 |  |
| `status` | integer | 否 | 设备升级作业状态。1：待确认。 2：待推送。 3：已推送。  4：升级中。 5:升级成功 6: 升级失败. 7:已取消 (格式: int64) |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "firmwareID": "string",
  "isOnline": 1,
  "jobID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "srcVersion": "string",
  "status": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "string",
        "destVersion": "string",
        "detail": "string",
        "deviceName": "string",
        "firmwareID": "string",
        "id": "string",
        "isOnline": 1,
        "jobID": "string",
        "productID": "string",
        "productName": "string",
        "srcVersion": "string",
        "status": 1,
        "step": 1,
        "updatedTime": 1
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
ur api /api/v1/things/ota/firmware/device/get-list \
  --body '{"deviceName": "示例名称", "firmwareID": "string", "isOnline": 1, "jobID": "string", "page": {"page": 1, "pageSize": 1}, "srcVersion": "string", "status": 1}'
```

### POST `/api/v1/things/ota/firmware/device/retry`

**说明**: 重试设备升级

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceNames` | array[string] | 是 |  |
| `firmwareID` | string | 是 |  |
| `jobID` | string | 是 |  |

**请求示例**:
```json
{
  "deviceNames": [
    "string"
  ],
  "firmwareID": "string",
  "jobID": "string"
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
ur api /api/v1/things/ota/firmware/device/retry \
  --body '{"deviceNames": ["string"], "firmwareID": "string", "jobID": "string"}'
```
