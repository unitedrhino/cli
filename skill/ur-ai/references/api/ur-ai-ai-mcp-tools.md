# ur-ai ai/mcp/tools

获取会话MCP工具列表

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/mcp/tools/get-tools` | 获取会话MCP工具列表 | all |

## 详细说明

### POST `/api/v1/ai/mcp/tools/get-tools`

**说明**: 获取会话MCP工具列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `sessionID` | string | 是 |  会话ID |
| `tenantCode` | string | 否 |  企业编码 |

**请求示例**:
```json
{
  "sessionID": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "timestamp": 1,
    "tools": {},
    "totalCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/tools/get-tools \
  --body '{"sessionID": "string", "tenantCode": "string"}'
```
