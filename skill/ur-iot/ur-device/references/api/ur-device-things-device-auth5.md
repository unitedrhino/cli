# ur-device things/device/auth5

设备操作认证 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/auth5/access` | 设备操作认证 | device |
| POST | `/api/v1/things/device/auth5/login` | 设备登录认证 | device |

## 详细说明

### POST `/api/v1/things/device/auth5/access`

**说明**: 设备操作认证

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `action` | string | 否 | 操作 |
| `clientID` | string | 否 | clientID |
| `ip` | string | 否 | 访问的ip地址 |
| `topic` | string | 否 | 主题 |
| `username` | string | 否 | 用户名 |

**请求示例**:
```json
{
  "action": "string",
  "clientID": "string",
  "ip": "string",
  "topic": "string",
  "username": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "result": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth5/access \
  --body '{"action": "string", "clientID": "string", "ip": "string", "topic": "string", "username": "string"}'
```

### POST `/api/v1/things/device/auth5/login`

**说明**: 设备登录认证

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `certificate` | string | 否 | 客户端证书 base64后传过来 |
| `clientID` | string | 是 | clientID |
| `ip` | string | 是 | 访问的ip地址 |
| `password` | string | 否 | 密码 |
| `username` | string | 是 | 用户名 |

**请求示例**:
```json
{
  "certificate": "string",
  "clientID": "string",
  "ip": "string",
  "password": "******",
  "username": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "is_superuser": true,
    "result": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth5/login \
  --body '{"certificate": "string", "clientID": "string", "ip": "string", "password": "******", "username": "string"}'
```
