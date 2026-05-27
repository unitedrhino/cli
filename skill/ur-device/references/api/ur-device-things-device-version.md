# ur-device things/device/version

获取设备模块版本列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/version/get-list` | 获取设备模块版本列表 | admin |
| POST | `/api/v1/things/device/version/get-one` | 获取设备模块版本详情 | admin |

## 详细说明

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
