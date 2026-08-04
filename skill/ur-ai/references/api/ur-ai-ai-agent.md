# ur-ai ai/agent

创建助手

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/agent/info/create` | 创建助手 | admin |

## 详细说明

### POST `/api/v1/ai/agent/info/create`

**说明**: 创建助手

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `asrConfigID` | string | 否 |  ASR配置ID |
| `asrSpeed` | string | 否 |  ASR语速（旧版，兼容保留） |
| `avatar` | string | 否 |  助手头像URL |
| `capabilities` | array[string] | 否 |  能力列表 |
| `code` | string | 否 |  助手编码（租户内唯一） |
| `customPrompt` | string | 否 |  自定义提示词（旧版，兼容保留） |
| `description` | string | 否 |  助手描述 |
| `groupID` | integer | 否 |  所属助手组ID (格式: int64) |
| `id` | integer | 否 |  助手ID (格式: int64) |
| `isAllKb` | integer | 否 |  是否使用全部知识库 1-是 2-否 (格式: int64) |
| `isAllMcp` | integer | 否 |  是否使用全部MCP 1-是 2-否 (格式: int64) |
| `isAllSkill` | integer | 否 |  是否使用全部Skills 1-是 2-否 (格式: int64) |
| `isPublic` | integer | 否 |  是否公开 1-是 2-否 (格式: int64) |
| `knowledgeBaseIds` | string | 否 |  知识库ID列表（JSON字符串） |
| `llmConfigID` | string | 否 |  LLM配置ID |
| `llmModelType` | string | 否 |  LLM模型级别 |
| `llmParams` | string | 否 |  LLM参数（JSON字符串） |
| `mcpServiceIds` | string | 否 |  MCP服务ID列表（JSON字符串） |
| `name` | string | 是 |  助手名称 |
| `skillIds` | string | 否 |  Skill ID列表（JSON字符串） |
| `sort` | integer | 否 |  排序权重 (格式: int64) |
| `status` | string | 否 |  助手状态 1-启用 2-禁用 |
| `systemPrompt` | string | 否 |  系统提示词 |
| `tenantCode` | string | 否 |  租户编码 |
| `ttsConfigID` | string | 否 |  TTS配置ID |
| `userID` | integer | 否 |  用户ID (格式: int64) |
| `visionConfigID` | string | 否 |  视觉配置ID |
| `voiceParams` | string | 否 |  语音参数（JSON字符串） |
| `welcomeMessage` | string | 否 |  欢迎语 |

**请求示例**:
```json
{
  "asrConfigID": "string",
  "asrSpeed": "string",
  "avatar": "string",
  "capabilities": [
    "string"
  ],
  "code": "string",
  "customPrompt": "string",
  "description": "string",
  "groupID": 1,
  "id": 1,
  "isAllKb": 1,
  "isAllMcp": 1,
  "isAllSkill": 1,
  "isPublic": 1,
  "knowledgeBaseIds": "string",
  "llmConfigID": "string",
  "llmModelType": "string",
  "llmParams": "string",
  "mcpServiceIds": "string",
  "name": "示例名称",
  "skillIds": "string",
  "sort": 1,
  "status": "string",
  "systemPrompt": "string",
  "tenantCode": "string",
  "ttsConfigID": "string",
  "userID": 1,
  "visionConfigID": "string",
  "voiceParams": "string",
  "welcomeMessage": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/agent/info/create \
  --body '{"asrConfigID": "string", "asrSpeed": "string", "avatar": "string", "capabilities": ["string"], "code": "string", "customPrompt": "string", "description": "string", "groupID": 1, "id": 1, "isAllKb": 1, "isAllMcp": 1, "isAllSkill": 1, "isPublic": 1, "knowledgeBaseIds": "string", "llmConfigID": "string", "llmModelType": "string", "llmParams": "string", "mcpServiceIds": "string", "name": "示例名称", "skillIds": "string", "sort": 1, "status": "string", "systemPrompt": "string", "tenantCode": "string", "ttsConfigID": "string", "userID": 1, "visionConfigID": "string", "voiceParams": "string", "welcomeMessage": "string"}'
```
