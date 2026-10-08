# ur-protocol things/protocol/image

下载协议镜像

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/protocol/image/pull` | 下载协议镜像 | admin |

## 详细说明

### POST `/api/v1/things/protocol/image/pull`

**说明**: 下载协议镜像

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
ur api /api/v1/things/protocol/image/pull \
  --body '{"protocolID": "string"}'
```
