# ur-product things/product/remoteConfig

创建配置 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/product/remote-config/create` | 创建配置 | admin |
| POST | `/api/v1/things/product/remote-config/get-list` | 获取配置列表 | admin |
| POST | `/api/v1/things/product/remote-config/lastest-read` | 获取最新配置 | admin |
| POST | `/api/v1/things/product/remote-config/push-all` | 推送配置 | admin |

## 详细说明

### POST `/api/v1/things/product/remote-config/create`

**说明**: 创建配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 | 配置内容 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "content": "string",
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
ur api /api/v1/things/product/remote-config/create \
  --body '{"content": "string", "productID": "string"}'
```

### POST `/api/v1/things/product/remote-config/get-list`

**说明**: 获取配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "content": "string",
        "createTime": "2026-01-01T00:00:00Z",
        "id": "string"
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
ur api /api/v1/things/product/remote-config/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}, "productID": "string"}'
```

### POST `/api/v1/things/product/remote-config/lastest-read`

**说明**: 获取最新配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "content": "string",
    "createTime": "2026-01-01T00:00:00Z",
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/remote-config/lastest-read \
  --body '{"productID": "string"}'
```

### POST `/api/v1/things/product/remote-config/push-all`

**说明**: 推送配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
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
ur api /api/v1/things/product/remote-config/push-all \
  --body '{"productID": "string"}'
```
