# ur-system system/mcp

Stateless MCP HTTP

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/mcp/run` | Stateless MCP HTTP | all |

## 详细说明

### POST `/api/v1/system/mcp/run`

**说明**: Stateless MCP HTTP

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mcp/run \
  --body '{}'
```
