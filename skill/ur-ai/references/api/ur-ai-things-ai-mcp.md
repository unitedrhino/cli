# ur-ai things/ai/mcp

MCP消息发送 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/ai/mcp/message` | MCP消息发送 | all |
| POST | `/api/v1/things/ai/mcp/run` | Stateless MCP HTTP | all |
| GET | `/api/v1/things/ai/mcp/sse` | SSE连接 | all |
| POST | `/api/v1/things/ai/mcp/sse` | SSE连接（POST） | all |

## 详细说明

### POST `/api/v1/things/ai/mcp/message`

**说明**: MCP消息发送

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
ur api /api/v1/things/ai/mcp/message \
  --body '{}'
```

### POST `/api/v1/things/ai/mcp/run`

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
ur api /api/v1/things/ai/mcp/run \
  --body '{}'
```

### GET `/api/v1/things/ai/mcp/sse`

**说明**: SSE连接

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
ur api /api/v1/things/ai/mcp/sse \
  --body '{}'
```

### POST `/api/v1/things/ai/mcp/sse`

**说明**: SSE连接（POST）

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
ur api /api/v1/things/ai/mcp/sse \
  --body '{}'
```
