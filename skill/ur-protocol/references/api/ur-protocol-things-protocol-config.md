# ur-protocol things/protocol/config

创建协议配置 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/protocol/config/create` | 创建协议配置 | admin |
| POST | `/api/v1/things/protocol/config/delete` | 删除协议配置 | admin |
| POST | `/api/v1/things/protocol/config/get-list` | 获取协议配置列表 | admin |
| POST | `/api/v1/things/protocol/config/get-one` | 获取协议配置详情 | admin |
| POST | `/api/v1/things/protocol/config/update` | 更新协议配置 | admin |

## 详细说明

### POST `/api/v1/things/protocol/config/create`

**说明**: 创建协议配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `config` | object | 否 |  配置键值对 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  描述（可清空） |
| `id` | string | 否 |  |
| `protocolCode` | string | 否 |  关联协议Code |
| `protocolID` | string | 否 |  关联协议ID |
| `updatedTime` | string | 否 |  更新时间 |

**请求示例**:
```json
{
  "config": {},
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "protocolCode": "string",
  "protocolID": "string",
  "updatedTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/protocol/config/create \
  --body '{"config": {}, "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "protocolCode": "string", "protocolID": "string", "updatedTime": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/protocol/config/delete`

**说明**: 删除协议配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

**请求示例**:
```json
{
  "id": 1
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
ur api /api/v1/things/protocol/config/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/protocol/config/get-list`

**说明**: 获取协议配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `protocolCode` | string | 否 |  按协议Code过滤 |
| `protocolID` | string | 否 |  按协议ID过滤 |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "protocolCode": "string",
  "protocolID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "config": {},
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "id": "string",
        "protocolCode": "string",
        "protocolID": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/config/get-list \
  --body '{"page": {"page": 1, "pageSize": 1}, "protocolCode": "string", "protocolID": "string"}'
```

### POST `/api/v1/things/protocol/config/get-one`

**说明**: 获取协议配置详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

**请求示例**:
```json
{
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "config": {},
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "protocolCode": "string",
    "protocolID": "string",
    "updatedTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/config/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/things/protocol/config/update`

**说明**: 更新协议配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `config` | object | 否 |  配置键值对 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  描述（可清空） |
| `id` | string | 否 |  |
| `protocolCode` | string | 否 |  关联协议Code |
| `protocolID` | string | 否 |  关联协议ID |
| `updatedTime` | string | 否 |  更新时间 |

**请求示例**:
```json
{
  "config": {},
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "protocolCode": "string",
  "protocolID": "string",
  "updatedTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/protocol/config/update \
  --body '{"config": {}, "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "protocolCode": "string", "protocolID": "string", "updatedTime": "2026-01-01T00:00:00Z"}'
```
