# ur-ai AI中台

删除助手 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/agent/info/delete` | 删除助手 | admin |
| POST | `/api/v1/ai/agent/info/get-list` | 获取助手列表 | admin |
| POST | `/api/v1/ai/agent/info/get-one` | 获取助手详情 | admin |
| POST | `/api/v1/ai/agent/info/update` | 更新助手 | admin |
| POST | `/api/v1/ai/clone/snapshot/create` | 创建快照 | admin |
| POST | `/api/v1/ai/clone/snapshot/delete` | 删除快照 | admin |
| POST | `/api/v1/ai/clone/snapshot/get-list` | 获取快照列表 | admin |
| POST | `/api/v1/ai/clone/snapshot/get-one` | 获取快照详情 | admin |
| POST | `/api/v1/ai/clone/snapshot/restore` | 恢复快照 | admin |
| POST | `/api/v1/ai/skill/approve` | 审批技能 | admin |
| POST | `/api/v1/ai/skill/file/get-list` | 列出技能包内文件 | admin |
| POST | `/api/v1/ai/skill/review` | AI审阅 | admin |
| POST | `/api/v1/ai/skill/submit` | 提交审核 | admin |
| POST | `/api/v1/ai/skill/test` | 测试技能 | admin |
| POST | `/api/v1/ai/skill/validate` | 验证技能 | admin |

## 详细说明

### POST `/api/v1/ai/agent/info/delete`

**说明**: 删除助手

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  助手ID (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "id": 1,
  "tenantCode": "string"
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
ur api /api/v1/ai/agent/info/delete \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/agent/info/get-list`

**说明**: 获取助手列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupID` | integer | 否 |  按助手组ID过滤 (格式: int64) |
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态过滤 |
| `tenantCode` | string | 否 |  租户编码过滤 |

**请求示例**:
```json
{
  "groupID": 1,
  "name": "示例名称",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
  },
  "status": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
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
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/agent/info/get-list \
  --body '{"groupID": 1, "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/agent/info/get-one`

**说明**: 获取助手详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  助手ID (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "id": 1,
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
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
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/agent/info/get-one \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/agent/info/update`

**说明**: 更新助手

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/agent/info/update \
  --body '{"asrConfigID": "string", "asrSpeed": "string", "avatar": "string", "capabilities": ["string"], "code": "string", "customPrompt": "string", "description": "string", "groupID": 1, "id": 1, "isAllKb": 1, "isAllMcp": 1, "isAllSkill": 1, "isPublic": 1, "knowledgeBaseIds": "string", "llmConfigID": "string", "llmModelType": "string", "llmParams": "string", "mcpServiceIds": "string", "name": "示例名称", "skillIds": "string", "sort": 1, "status": "string", "systemPrompt": "string", "tenantCode": "string", "ttsConfigID": "string", "userID": 1, "visionConfigID": "string", "voiceParams": "string", "welcomeMessage": "string"}'
```

### POST `/api/v1/ai/clone/snapshot/create`

**说明**: 创建快照

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `description` | string | 否 |  快照描述 |
| `isAuto` | integer | 否 |  是否自动创建 1-是 2-否 (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "description": "string",
  "isAuto": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "checksum": "string",
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/snapshot/create \
  --body '{"cloneID": 1, "description": "string", "isAuto": 1}'
```

### POST `/api/v1/ai/clone/snapshot/delete`

**说明**: 删除快照

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `snapshotID` | integer | 是 |  快照ID (格式: int64) |

**请求示例**:
```json
{
  "snapshotID": 1
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
ur api /api/v1/ai/clone/snapshot/delete \
  --body '{"snapshotID": 1}'
```

### POST `/api/v1/ai/clone/snapshot/get-list`

**说明**: 获取快照列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "checksum": "string",
        "cloneID": 1,
        "createdTime": 1,
        "description": "string",
        "fileCount": 1,
        "id": 1,
        "isAuto": 1,
        "sizeBytes": 1,
        "status": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/snapshot/get-list \
  --body '{"cloneID": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/ai/clone/snapshot/get-one`

**说明**: 获取快照详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `snapshotID` | integer | 是 |  快照ID (格式: int64) |

**请求示例**:
```json
{
  "snapshotID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "checksum": "string",
    "cloneID": 1,
    "createdTime": 1,
    "description": "string",
    "fileCount": 1,
    "id": 1,
    "isAuto": 1,
    "sizeBytes": 1,
    "status": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/snapshot/get-one \
  --body '{"snapshotID": 1}'
```

### POST `/api/v1/ai/clone/snapshot/restore`

**说明**: 恢复快照

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `snapshotID` | integer | 是 |  快照ID (格式: int64) |

**请求示例**:
```json
{
  "snapshotID": 1
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
ur api /api/v1/ai/clone/snapshot/restore \
  --body '{"snapshotID": 1}'
```

### POST `/api/v1/ai/skill/approve`

**说明**: 审批技能

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `approved` | boolean | 是 |  是否通过 (格式: boolean) |
| `comment` | string | 否 |  审批意见 |
| `id` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "approved": true,
  "comment": "string",
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "createdTime": 1,
    "desc": "string",
    "id": 1,
    "manifest": "string",
    "name": "示例名称",
    "ossPath": "string",
    "parentID": 1,
    "publishStatus": "string",
    "reviewResult": "string",
    "scope": "string",
    "sort": 1,
    "status": 1,
    "tenantCode": "string",
    "testResult": "string",
    "updatedTime": 1,
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/approve \
  --body '{"approved": true, "comment": "string", "id": 1}'
```

### POST `/api/v1/ai/skill/file/get-list`

**说明**: 列出技能包内文件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `path` | string | 否 |  目录路径（默认根目录） |
| `skillID` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "path": "string",
  "skillID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "content": "string",
        "isDir": true,
        "path": "string",
        "size": 1,
        "updatedTime": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/file/get-list \
  --body '{"path": "string", "skillID": 1}'
```

### POST `/api/v1/ai/skill/review`

**说明**: AI审阅

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `force` | boolean | 否 |  是否强制通过 (格式: boolean) |
| `id` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "force": true,
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "createdTime": 1,
    "desc": "string",
    "id": 1,
    "manifest": "string",
    "name": "示例名称",
    "ossPath": "string",
    "parentID": 1,
    "publishStatus": "string",
    "reviewResult": "string",
    "scope": "string",
    "sort": 1,
    "status": 1,
    "tenantCode": "string",
    "testResult": "string",
    "updatedTime": 1,
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/review \
  --body '{"force": true, "id": 1}'
```

### POST `/api/v1/ai/skill/submit`

**说明**: 提交审核

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "createdTime": 1,
    "desc": "string",
    "id": 1,
    "manifest": "string",
    "name": "示例名称",
    "ossPath": "string",
    "parentID": 1,
    "publishStatus": "string",
    "reviewResult": "string",
    "scope": "string",
    "sort": 1,
    "status": 1,
    "tenantCode": "string",
    "testResult": "string",
    "updatedTime": 1,
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/submit \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/skill/test`

**说明**: 测试技能

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  技能ID (格式: int64) |
| `testInput` | string | 否 |  测试输入 |

**请求示例**:
```json
{
  "id": 1,
  "testInput": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "durationMs": 1,
    "errors": [
      "string"
    ],
    "output": "string",
    "success": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/test \
  --body '{"id": 1, "testInput": "string"}'
```

### POST `/api/v1/ai/skill/validate`

**说明**: 验证技能

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "errors": [
      "string"
    ],
    "success": true,
    "warnings": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/validate \
  --body '{"id": 1}'
```
