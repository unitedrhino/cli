# ur-ai ai/chat

获取当前会话可用的后端工具列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/chat/available-tools` | 获取当前会话可用的后端工具列表 | all |
| POST | `/api/v1/ai/chat/check-available` | 检测 AI 对话功能是否可用 | all |
| POST | `/api/v1/ai/chat/completions` | 文字对话（指定 agentID） | all |
| POST | `/api/v1/ai/chat/tool-result` | 前端 Tool 执行结果回传 | all |
| POST | `/api/v1/ai/embedding/debug` | Embedding配置调试 | platform |
| POST | `/api/v1/ai/llm/debug` | LLM配置调试 | platform |
| POST | `/api/v1/ai/tool-dev/completions` | 工具编辑 AI 对话（自动绑定 tool_assistant） | all |
| POST | `/api/v1/ai/user/completions` | 用户 AI 对话（自动查找/创建专属 Clone） | all |

## 详细说明

### POST `/api/v1/ai/chat/available-tools`

**说明**: 获取当前会话可用的后端工具列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 |  助手ID（0=裸LLM调用，不注入任何上下文） (格式: int64) |
| `audioConfig` | object | 否 |  |
| `audioConfig.channels` | integer | 否 |  声道数，默认 1 (格式: int32) |
| `audioConfig.format` | string | 否 |  音频格式：opus/mp3/wav，默认 opus |
| `audioConfig.sampleRate` | integer | 否 |  采样率，默认 24000 (格式: int32) |
| `cloneID` | integer | 否 |  分身ID（agentID=0 时忽略） (格式: int64) |
| `currentAppID` | string | 否 |  当前页面所属应用ID（agentID=0 时忽略） |
| `currentAppName` | string | 否 |  当前页面所属应用名称（agentID=0 时忽略） |
| `currentRoute` | string | 否 |  当前前端路由（agentID=0 时忽略） |
| `messages` | array[ChatMessage] | 是 |  对话消息列表 |
| `modelType` | string | 否 |  LLM 模型类型（agentID=0 时生效，默认 "small"） |
| `outputModalities` | array[string] | 否 |  输出模态：text / audio，默认 ["text"] |
| `sessionID` | string | 否 |  会话ID（agentID=0 时忽略） |
| `stream` | boolean | 否 |  是否流式返回（SSE） (格式: boolean) |
| `toolID` | integer | 否 |  工具ID（agentID=0 时忽略） (格式: int64) |
| `tools` | array[ToolDef] | 否 |  前端注册的工具定义列表（agentID=0 时忽略） |

**请求示例**:
```json
{
  "agentID": 1,
  "audioConfig": {
    "channels": 1,
    "format": "string",
    "sampleRate": 1
  },
  "cloneID": 1,
  "currentAppID": "string",
  "currentAppName": "string",
  "currentRoute": "string",
  "messages": [
    {
      "contents": [
        {
          "audioData": "string",
          "audioFormat": "string",
          "fileMime": "string",
          "fileName": "string",
          "fileUrl": "string",
          "imageUrl": "string",
          "text": "string",
          "type": "string",
          "videoUrl": "string"
        }
      ],
      "role": "string",
      "toolCallID": "string",
      "toolCalls": [
        {
          "arguments": "string",
          "id": "string",
          "name": "示例名称"
        }
      ]
    }
  ],
  "modelType": "string",
  "outputModalities": [
    "string"
  ],
  "sessionID": "string",
  "stream": true,
  "toolID": 1,
  "tools": [
    {
      "available": "string",
      "description": "string",
      "formSchema": {
        "cancelLabel": "string",
        "description": "string",
        "fields": [
          {
            "default": "...",
            "key": "...",
            "label": "...",
            "max": "...",
            "min": "...",
            "options": "...",
            "placeholder": "...",
            "required": "...",
            "type": "..."
          }
        ],
        "submitLabel": "string",
        "title": "string"
      },
      "interactionType": "string",
      "name": "示例名称",
      "parameters": {}
    }
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "tools": [
      {
        "available": "string",
        "description": "string",
        "formSchema": {
          "cancelLabel": "string",
          "description": "string",
          "fields": [
            {
              "default": "...",
              "key": "...",
              "label": "...",
              "max": "...",
              "min": "...",
              "options": "...",
              "placeholder": "...",
              "required": "...",
              "type": "..."
            }
          ],
          "submitLabel": "string",
          "title": "string"
        },
        "interactionType": "string",
        "name": "示例名称",
        "parameters": {}
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/chat/available-tools \
  --body '{"agentID": 1, "audioConfig": {"channels": 1, "format": "string", "sampleRate": 1}, "cloneID": 1, "currentAppID": "string", "currentAppName": "string", "currentRoute": "string", "messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "modelType": "string", "outputModalities": ["string"], "sessionID": "string", "stream": true, "toolID": 1, "tools": [{"available": "string", "description": "string", "formSchema": {"cancelLabel": "string", "description": "string", "fields": [{"default": "...", "key": "...", "label": "...", "max": "...", "min": "...", "options": "...", "placeholder": "...", "required": "...", "type": "..."}], "submitLabel": "string", "title": "string"}, "interactionType": "string", "name": "示例名称", "parameters": {}}]}'
```

### POST `/api/v1/ai/chat/check-available`

**说明**: 检测 AI 对话功能是否可用

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "available": true,
    "reason": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/chat/check-available \
  --body '{}'
```

### POST `/api/v1/ai/chat/completions`

**说明**: 文字对话（指定 agentID）

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 |  助手ID（0=裸LLM调用，不注入任何上下文） (格式: int64) |
| `audioConfig` | object | 否 |  |
| `audioConfig.channels` | integer | 否 |  声道数，默认 1 (格式: int32) |
| `audioConfig.format` | string | 否 |  音频格式：opus/mp3/wav，默认 opus |
| `audioConfig.sampleRate` | integer | 否 |  采样率，默认 24000 (格式: int32) |
| `cloneID` | integer | 否 |  分身ID（agentID=0 时忽略） (格式: int64) |
| `currentAppID` | string | 否 |  当前页面所属应用ID（agentID=0 时忽略） |
| `currentAppName` | string | 否 |  当前页面所属应用名称（agentID=0 时忽略） |
| `currentRoute` | string | 否 |  当前前端路由（agentID=0 时忽略） |
| `messages` | array[ChatMessage] | 是 |  对话消息列表 |
| `modelType` | string | 否 |  LLM 模型类型（agentID=0 时生效，默认 "small"） |
| `outputModalities` | array[string] | 否 |  输出模态：text / audio，默认 ["text"] |
| `sessionID` | string | 否 |  会话ID（agentID=0 时忽略） |
| `stream` | boolean | 否 |  是否流式返回（SSE） (格式: boolean) |
| `toolID` | integer | 否 |  工具ID（agentID=0 时忽略） (格式: int64) |
| `tools` | array[ToolDef] | 否 |  前端注册的工具定义列表（agentID=0 时忽略） |

**请求示例**:
```json
{
  "agentID": 1,
  "audioConfig": {
    "channels": 1,
    "format": "string",
    "sampleRate": 1
  },
  "cloneID": 1,
  "currentAppID": "string",
  "currentAppName": "string",
  "currentRoute": "string",
  "messages": [
    {
      "contents": [
        {
          "audioData": "string",
          "audioFormat": "string",
          "fileMime": "string",
          "fileName": "string",
          "fileUrl": "string",
          "imageUrl": "string",
          "text": "string",
          "type": "string",
          "videoUrl": "string"
        }
      ],
      "role": "string",
      "toolCallID": "string",
      "toolCalls": [
        {
          "arguments": "string",
          "id": "string",
          "name": "示例名称"
        }
      ]
    }
  ],
  "modelType": "string",
  "outputModalities": [
    "string"
  ],
  "sessionID": "string",
  "stream": true,
  "toolID": 1,
  "tools": [
    {
      "available": "string",
      "description": "string",
      "formSchema": {
        "cancelLabel": "string",
        "description": "string",
        "fields": [
          {
            "default": "...",
            "key": "...",
            "label": "...",
            "max": "...",
            "min": "...",
            "options": "...",
            "placeholder": "...",
            "required": "...",
            "type": "..."
          }
        ],
        "submitLabel": "string",
        "title": "string"
      },
      "interactionType": "string",
      "name": "示例名称",
      "parameters": {}
    }
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "answerPolicy": "string",
    "audioData": "string",
    "audioFormat": "string",
    "citations": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "confidenceBand": "string",
    "content": "string",
    "contents": [
      {
        "audioData": "string",
        "audioFormat": "string",
        "fileMime": "string",
        "fileName": "string",
        "fileUrl": "string",
        "imageUrl": "string",
        "text": "string",
        "type": "string",
        "videoUrl": "string"
      }
    ],
    "expandedCitations": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "matchedDocuments": [
      {
        "chunkCount": 1,
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "score": 1,
        "snippet": "string"
      }
    ],
    "rewrittenQuery": "string",
    "role": "string",
    "sessionID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/chat/completions \
  --body '{"agentID": 1, "audioConfig": {"channels": 1, "format": "string", "sampleRate": 1}, "cloneID": 1, "currentAppID": "string", "currentAppName": "string", "currentRoute": "string", "messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "modelType": "string", "outputModalities": ["string"], "sessionID": "string", "stream": true, "toolID": 1, "tools": [{"available": "string", "description": "string", "formSchema": {"cancelLabel": "string", "description": "string", "fields": [{"default": "...", "key": "...", "label": "...", "max": "...", "min": "...", "options": "...", "placeholder": "...", "required": "...", "type": "..."}], "submitLabel": "string", "title": "string"}, "interactionType": "string", "name": "示例名称", "parameters": {}}]}'
```

### POST `/api/v1/ai/chat/tool-result`

**说明**: 前端 Tool 执行结果回传

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `error` | string | 否 |  执行失败时的错误信息 |
| `result` | string | 是 |  工具执行结果，JSON 字符串格式 |
| `sessionID` | string | 是 |  当前会话ID |
| `toolCallID` | string | 是 |  工具调用ID，对应 SSE tool_call_execute 事件中的 toolCallID |

**请求示例**:
```json
{
  "error": "string",
  "result": "string",
  "sessionID": "string",
  "toolCallID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "sessionID": "string",
    "status": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/chat/tool-result \
  --body '{"error": "string", "result": "string", "sessionID": "string", "toolCallID": "string"}'
```

### POST `/api/v1/ai/embedding/debug`

**说明**: Embedding配置调试

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `jsonData` | string | 是 |  模型配置 JSON 字符串 |
| `text` | string | 是 |  待编码文本 |

**请求示例**:
```json
{
  "jsonData": "string",
  "text": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "dimension": 1,
    "model": "示例名称",
    "preview": [
      1
    ],
    "usage": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/embedding/debug \
  --body '{"jsonData": "string", "text": "string"}'
```

### POST `/api/v1/ai/llm/debug`

**说明**: LLM配置调试

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 是 |  配置类型，当前仅支持 llm |
| `jsonData` | string | 是 |  模型配置 JSON 字符串 |
| `messages` | array[ChatMessage] | 是 |  调试消息（复用现有 ChatMessage） |
| `outputModalities` | array[string] | 否 |  输出模态 |

**请求示例**:
```json
{
  "cfgType": "string",
  "jsonData": "string",
  "messages": [
    {
      "contents": [
        {
          "audioData": "string",
          "audioFormat": "string",
          "fileMime": "string",
          "fileName": "string",
          "fileUrl": "string",
          "imageUrl": "string",
          "text": "string",
          "type": "string",
          "videoUrl": "string"
        }
      ],
      "role": "string",
      "toolCallID": "string",
      "toolCalls": [
        {
          "arguments": "string",
          "id": "string",
          "name": "示例名称"
        }
      ]
    }
  ],
  "outputModalities": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "audioData": "string",
    "audioFormat": "string",
    "content": "string",
    "raw": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/llm/debug \
  --body '{"cfgType": "string", "jsonData": "string", "messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "outputModalities": ["string"]}'
```

### POST `/api/v1/ai/tool-dev/completions`

**说明**: 工具编辑 AI 对话（自动绑定 tool_assistant）

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `messages` | array[ChatMessage] | 是 |  对话消息列表 |
| `sessionID` | string | 否 |  会话ID（为空自动生成） |
| `stream` | boolean | 否 |  是否流式返回（SSE） (格式: boolean) |
| `toolID` | integer | 否 |  工具ID（首次需传，后续可从 session 恢复） (格式: int64) |

**请求示例**:
```json
{
  "messages": [
    {
      "contents": [
        {
          "audioData": "string",
          "audioFormat": "string",
          "fileMime": "string",
          "fileName": "string",
          "fileUrl": "string",
          "imageUrl": "string",
          "text": "string",
          "type": "string",
          "videoUrl": "string"
        }
      ],
      "role": "string",
      "toolCallID": "string",
      "toolCalls": [
        {
          "arguments": "string",
          "id": "string",
          "name": "示例名称"
        }
      ]
    }
  ],
  "sessionID": "string",
  "stream": true,
  "toolID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "answerPolicy": "string",
    "audioData": "string",
    "audioFormat": "string",
    "citations": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "confidenceBand": "string",
    "content": "string",
    "contents": [
      {
        "audioData": "string",
        "audioFormat": "string",
        "fileMime": "string",
        "fileName": "string",
        "fileUrl": "string",
        "imageUrl": "string",
        "text": "string",
        "type": "string",
        "videoUrl": "string"
      }
    ],
    "expandedCitations": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "matchedDocuments": [
      {
        "chunkCount": 1,
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "score": 1,
        "snippet": "string"
      }
    ],
    "rewrittenQuery": "string",
    "role": "string",
    "sessionID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool-dev/completions \
  --body '{"messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "sessionID": "string", "stream": true, "toolID": 1}'
```

### POST `/api/v1/ai/user/completions`

**说明**: 用户 AI 对话（自动查找/创建专属 Clone）

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `assistantCode` | string | 否 |  助手编码（如 tool_assistant），指定时优先使用该 Agent |
| `currentAppID` | string | 否 |  当前页面所属应用ID |
| `currentAppName` | string | 否 |  当前页面所属应用名称 |
| `currentRoute` | string | 否 |  当前前端路由 |
| `messages` | array[ChatMessage] | 是 |  对话消息列表 |
| `outputModalities` | array[string] | 否 |  输出模态：text / audio，默认 ["text"] |
| `sessionID` | string | 否 |  会话ID（为空自动生成） |
| `stream` | boolean | 否 |  是否流式返回（SSE） (格式: boolean) |
| `toolID` | integer | 否 |  工具ID（用于工具编辑场景，注入文件操作工具） (格式: int64) |
| `tools` | array[ToolDef] | 否 |  前端注册的工具定义列表 |

**请求示例**:
```json
{
  "assistantCode": "string",
  "currentAppID": "string",
  "currentAppName": "示例名称",
  "currentRoute": "string",
  "messages": [
    {
      "contents": [
        {
          "audioData": "string",
          "audioFormat": "string",
          "fileMime": "string",
          "fileName": "string",
          "fileUrl": "string",
          "imageUrl": "string",
          "text": "string",
          "type": "string",
          "videoUrl": "string"
        }
      ],
      "role": "string",
      "toolCallID": "string",
      "toolCalls": [
        {
          "arguments": "string",
          "id": "string",
          "name": "示例名称"
        }
      ]
    }
  ],
  "outputModalities": [
    "string"
  ],
  "sessionID": "string",
  "stream": true,
  "toolID": 1,
  "tools": [
    {
      "available": "string",
      "description": "string",
      "formSchema": {
        "cancelLabel": "string",
        "description": "string",
        "fields": [
          {
            "default": "...",
            "key": "...",
            "label": "...",
            "max": "...",
            "min": "...",
            "options": "...",
            "placeholder": "...",
            "required": "...",
            "type": "..."
          }
        ],
        "submitLabel": "string",
        "title": "string"
      },
      "interactionType": "string",
      "name": "示例名称",
      "parameters": {}
    }
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "answerPolicy": "string",
    "audioData": "string",
    "audioFormat": "string",
    "citations": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "confidenceBand": "string",
    "content": "string",
    "contents": [
      {
        "audioData": "string",
        "audioFormat": "string",
        "fileMime": "string",
        "fileName": "string",
        "fileUrl": "string",
        "imageUrl": "string",
        "text": "string",
        "type": "string",
        "videoUrl": "string"
      }
    ],
    "expandedCitations": [
      {
        "artifactID": 1,
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "mediaURI": "string",
        "mimeType": "string",
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "matchedDocuments": [
      {
        "chunkCount": 1,
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "score": 1,
        "snippet": "string"
      }
    ],
    "rewrittenQuery": "string",
    "role": "string",
    "sessionID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/user/completions \
  --body '{"assistantCode": "string", "currentAppID": "string", "currentAppName": "示例名称", "currentRoute": "string", "messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "outputModalities": ["string"], "sessionID": "string", "stream": true, "toolID": 1, "tools": [{"available": "string", "description": "string", "formSchema": {"cancelLabel": "string", "description": "string", "fields": [{"default": "...", "key": "...", "label": "...", "max": "...", "min": "...", "options": "...", "placeholder": "...", "required": "...", "type": "..."}], "submitLabel": "string", "title": "string"}, "interactionType": "string", "name": "示例名称", "parameters": {}}]}'
```
