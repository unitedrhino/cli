# ur-device things/device/auth

设备操作认证 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/auth/access` | 设备操作认证 | device |
| POST | `/api/v1/things/device/auth/login` | 设备登录认证 | device |
| POST | `/api/v1/things/device/auth/register` | 设备自动注册 | device |
| POST | `/api/v1/things/device/auth/root-check` | 鉴定是否是root账号 | device |

## 详细说明

### POST `/api/v1/things/device/auth/access`

**说明**: 设备操作认证

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `access` | string | 否 | 操作 |
| `clientID` | string | 否 | clientID |
| `ip` | string | 否 | 访问的ip地址 |
| `topic` | string | 否 | 主题 |
| `username` | string | 否 | 用户名 |

**请求示例**:
```json
{
  "access": "string",
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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth/access \
  --body '{"access": "string", "clientID": "string", "ip": "string", "topic": "string", "username": "string"}'
```

### POST `/api/v1/things/device/auth/login`

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth/login \
  --body '{"certificate": "string", "clientID": "string", "ip": "string", "password": "******", "username": "string"}'
```

### POST `/api/v1/things/device/auth/register`

**说明**: 设备自动注册

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 | 设备名称 读写 |
| `nonce` | integer | 是 | 随机数 (格式: int64) |
| `productId` | string | 是 | 产品id 只读 |
| `retEnc` | string | 否 | 返回使用的加密方式: aes128cbc(默认) aes128ecb |
| `signature` | string | 是 | 签名信息 |
| `tenantCode` | string | 否 | 企业号（可选）：产品为common企业或企业号与产品一致时生效 |
| `timestamp` | integer | 是 | 秒级时间戳 (格式: int64) |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "nonce": 1,
  "productId": "string",
  "retEnc": "string",
  "signature": "string",
  "tenantCode": "string",
  "timestamp": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "len": 1,
    "payload": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth/register \
  --body '{"deviceName": "示例名称", "nonce": 1, "productId": "string", "retEnc": "string", "signature": "string", "tenantCode": "string", "timestamp": 1}'
```

### POST `/api/v1/things/device/auth/root-check`

**说明**: 鉴定是否是root账号

**权限**: device

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `certificate` | string | 否 | 客户端证书 |
| `clientID` | string | 否 | clientID |
| `ip` | string | 否 | 访问的ip地址 |
| `password` | string | 否 | 密码 |
| `username` | string | 否 | 用户名 |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/auth/root-check \
  --body '{"certificate": "string", "clientID": "string", "ip": "string", "password": "******", "username": "string"}'
```
