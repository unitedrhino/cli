# ur-protocol things/protocol/container

安装协议容器 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/protocol/container/install` | 安装协议容器 | admin |
| POST | `/api/v1/things/protocol/container/status` | 查询协议容器状态 | admin |
| POST | `/api/v1/things/protocol/container/stop` | 停止协议容器 | admin |
| POST | `/api/v1/things/protocol/container/uninstall` | 卸载协议容器 | admin |
| POST | `/api/v1/things/protocol/container/update` | 更新协议容器 | admin |

## 详细说明

### POST `/api/v1/things/protocol/container/install`

**说明**: 安装协议容器

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `protocolID` | string | 是 |  |
| `version` | string | 否 |  |

**请求示例**:
```json
{
  "protocolID": "string",
  "version": "string"
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
ur api /api/v1/things/protocol/container/install \
  --body '{"protocolID": "string", "version": "string"}'
```

### POST `/api/v1/things/protocol/container/status`

**说明**: 查询协议容器状态

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `protocolID` | string | 是 |  |

**请求示例**:
```json
{
  "protocolID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "containerID": "string",
    "image": "string",
    "status": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/protocol/container/status \
  --body '{"protocolID": "string"}'
```

### POST `/api/v1/things/protocol/container/stop`

**说明**: 停止协议容器

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `protocolID` | string | 是 |  |

**请求示例**:
```json
{
  "protocolID": "string"
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
ur api /api/v1/things/protocol/container/stop \
  --body '{"protocolID": "string"}'
```

### POST `/api/v1/things/protocol/container/uninstall`

**说明**: 卸载协议容器

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `protocolID` | string | 是 |  |
| `removeImage` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "protocolID": "string",
  "removeImage": true
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
ur api /api/v1/things/protocol/container/uninstall \
  --body '{"protocolID": "string", "removeImage": true}'
```

### POST `/api/v1/things/protocol/container/update`

**说明**: 更新协议容器

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `protocolID` | string | 是 |  |
| `version` | string | 否 |  |

**请求示例**:
```json
{
  "protocolID": "string",
  "version": "string"
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
ur api /api/v1/things/protocol/container/update \
  --body '{"protocolID": "string", "version": "string"}'
```
