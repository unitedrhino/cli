# ur-ai ai/mcp/resolve

解析MCP服务端点

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/mcp/resolve/resolve` | 解析MCP服务端点 | all |

## 详细说明

### POST `/api/v1/ai/mcp/resolve/resolve`

**说明**: 解析MCP服务端点

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 是 |  服务名称 |
| `tenantCode` | string | 否 |  企业编码 |

**请求示例**:
```json
{
  "name": "示例名称",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "authConfig": "string",
    "authType": "string",
    "enabled": true,
    "endpoint": "string",
    "tools": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/resolve/resolve \
  --body '{"name": "示例名称", "tenantCode": "string"}'
```
