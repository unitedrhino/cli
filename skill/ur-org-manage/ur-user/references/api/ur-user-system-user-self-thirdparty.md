# ur-user system/user/self/thirdparty

查询第三方客户端绑定状态 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/self/thirdparty/setup-check` | 查询第三方客户端绑定状态 | public |
| POST | `/api/v1/system/user/self/thirdparty/setup-complete` | 完成第三方客户端绑定 | all |

## 详细说明

### POST `/api/v1/system/user/self/thirdparty/setup-check`

**说明**: 查询第三方客户端绑定状态

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `setupCode` | string | 是 |  CLI 生成的绑定码 |

**请求示例**:
```json
{
  "setupCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "accessKey": "string",
    "accessSecret": "string",
    "status": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/thirdparty/setup-check \
  --body '{"setupCode": "string"}'
```

### POST `/api/v1/system/user/self/thirdparty/setup-complete`

**说明**: 完成第三方客户端绑定

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessKey` | string | 是 |  访问密钥标识（前端传入，用于写入 Redis） |
| `accessSecret` | string | 是 |  访问密钥（前端传入，用于写入 Redis） |
| `accessTokenID` | string | 是 |  访问令牌 ID |
| `setupCode` | string | 是 |  CLI 生成的绑定码 |

**请求示例**:
```json
{
  "accessKey": "string",
  "accessSecret": "string",
  "accessTokenID": "string",
  "setupCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "success": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/thirdparty/setup-complete \
  --body '{"accessKey": "string", "accessSecret": "string", "accessTokenID": "string", "setupCode": "string"}'
```
