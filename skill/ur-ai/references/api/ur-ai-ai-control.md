# ur-ai ai/control

中断当前对话 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/control/abort` | 中断当前对话 | all |
| POST | `/api/v1/ai/control/goodbye` | 结束语音会话 | all |
| POST | `/api/v1/ai/control/listen` | 开始语音监听 | all |

## 详细说明

### POST `/api/v1/ai/control/abort`

**说明**: 中断当前对话

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `sessionID` | string | 是 |  会话ID |

**请求示例**:
```json
{
  "sessionID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/control/abort \
  --body '{"sessionID": "string"}'
```

### POST `/api/v1/ai/control/goodbye`

**说明**: 结束语音会话

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `sessionID` | string | 是 |  会话ID |

**请求示例**:
```json
{
  "sessionID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/control/goodbye \
  --body '{"sessionID": "string"}'
```

### POST `/api/v1/ai/control/listen`

**说明**: 开始语音监听

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `channels` | integer | 否 |  声道数 (格式: int32) |
| `frameDuration` | integer | 否 |  帧时长（ms） (格式: int32) |
| `params` | object | 否 |  扩展参数 |
| `sampleRate` | integer | 否 |  采样率 (格式: int32) |
| `sessionID` | string | 是 |  会话ID |

**请求示例**:
```json
{
  "channels": 1,
  "frameDuration": 1,
  "params": {},
  "sampleRate": 1,
  "sessionID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/control/listen \
  --body '{"channels": 1, "frameDuration": 1, "params": {}, "sampleRate": 1, "sessionID": "string"}'
```
