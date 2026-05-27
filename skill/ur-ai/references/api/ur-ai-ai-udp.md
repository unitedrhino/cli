# ur-ai ai/udp

创建UDP会话令牌

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/udp/token/create` | 创建UDP会话令牌 | all |

## 详细说明

### POST `/api/v1/ai/udp/token/create`

**说明**: 创建UDP会话令牌

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 是 |  助手ID (格式: int64) |
| `audioParamsJson` | string | 否 |  音频参数（JSON） |
| `mcpToolsJson` | string | 否 |  MCP工具清单（JSON） |
| `sessionTtlSec` | integer | 否 |  会话有效期（秒） (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "audioParamsJson": "string",
  "mcpToolsJson": "string",
  "sessionTtlSec": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "aesKey": "string",
    "expireAt": 1,
    "nonce": "string",
    "replyFormat": "string",
    "replyTopic": "string",
    "sessionID": "string",
    "token": "string",
    "udpHost": "string",
    "udpPort": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/udp/token/create \
  --body '{"agentID": 1, "audioParamsJson": "string", "mcpToolsJson": "string", "sessionTtlSec": 1}'
```
