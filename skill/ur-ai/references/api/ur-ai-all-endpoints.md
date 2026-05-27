# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/agent/group/create` | 创建助手组 | admin |
| POST | `/api/v1/ai/agent/group/delete` | 删除助手组 | admin |
| POST | `/api/v1/ai/agent/group/get-list` | 获取助手组列表 | admin |
| POST | `/api/v1/ai/agent/group/get-one` | 获取助手组详情 | admin |
| POST | `/api/v1/ai/agent/group/update` | 更新助手组 | admin |
| POST | `/api/v1/ai/agent/info/create` | 创建助手 | admin |
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
| POST | `/api/v1/ai/app/chat-send` | 发送对话消息 | - |
| POST | `/api/v1/ai/app/sse/chat-run` | sse模式ai聊天运行 | - |
| POST | `/api/v1/ai/app/sse/workflow-run` | sse模式执行 workflow | - |
| POST | `/api/v1/ai/app/workflow-run` | 执行 workflow | - |
| POST | `/api/v1/ai/mcp/run` | 添加规则链 | - |
| POST | `/api/v1/ai/chat/available-tools` | 获取当前会话可用的后端工具列表 | all |
| POST | `/api/v1/ai/chat/check-available` | 检测 AI 对话功能是否可用 | all |
| POST | `/api/v1/ai/chat/completions` | 文字对话（指定 agentID） | all |
| POST | `/api/v1/ai/chat/tool-result` | 前端 Tool 执行结果回传 | all |
| POST | `/api/v1/ai/embedding/debug` | Embedding配置调试 | platform |
| POST | `/api/v1/ai/llm/debug` | LLM配置调试 | platform |
| POST | `/api/v1/ai/user/completions` | 用户 AI 对话（自动查找/创建专属 Clone） | all |
| POST | `/api/v1/ai/clone/create` | 创建分身 | admin |
| POST | `/api/v1/ai/clone/delete` | 删除分身 | admin |
| POST | `/api/v1/ai/clone/get-list` | 获取分身列表 | admin |
| POST | `/api/v1/ai/clone/get-one` | 获取分身详情 | admin |
| POST | `/api/v1/ai/clone/update` | 更新分身 | admin |
| POST | `/api/v1/ai/clone/memory/archives/get-list` | 获取归档记录列表 | admin |
| POST | `/api/v1/ai/clone/memory/compact` | 压缩记忆 | admin |
| POST | `/api/v1/ai/clone/memory/create` | 添加记忆 | admin |
| POST | `/api/v1/ai/clone/memory/delete` | 删除记忆 | admin |
| POST | `/api/v1/ai/clone/memory/dream-runs/get-list` | 获取 Dream 运行记录列表 | admin |
| POST | `/api/v1/ai/clone/memory/dream/run` | 手动触发 Dream 整理 | admin |
| POST | `/api/v1/ai/clone/memory/entities/get-list` | 获取记忆实体列表 | admin |
| POST | `/api/v1/ai/clone/memory/profile/get-one` | 获取记忆画像 | admin |
| POST | `/api/v1/ai/clone/memory/records/get-list` | 获取记忆记录列表 | admin |
| POST | `/api/v1/ai/clone/memory/search` | 搜索记忆 | admin |
| POST | `/api/v1/ai/clone/memory/traces/get-list` | 获取召回轨迹列表 | admin |
| POST | `/api/v1/ai/config/info/create` | 创建配置 | admin |
| POST | `/api/v1/ai/config/info/delete` | 删除配置 | admin |
| POST | `/api/v1/ai/config/info/get-default` | 获取默认配置 | admin |
| POST | `/api/v1/ai/config/info/get-list` | 获取配置列表 | admin |
| POST | `/api/v1/ai/config/info/get-one` | 获取配置详情 | admin |
| POST | `/api/v1/ai/config/info/test` | 测试配置 | admin |
| POST | `/api/v1/ai/config/info/update` | 更新配置 | admin |
| POST | `/api/v1/ai/control/abort` | 中断当前对话 | all |
| POST | `/api/v1/ai/control/goodbye` | 结束语音会话 | all |
| POST | `/api/v1/ai/control/listen` | 开始语音监听 | all |
| POST | `/api/v1/ai/knowledge/compile/get-status` | 获取编译状态 | admin |
| POST | `/api/v1/ai/knowledge/compile/retry` | 触发知识库重编译 | admin |
| POST | `/api/v1/ai/knowledge/document/chunk/relation/get-list` | 获取切片关联列表 | admin |
| POST | `/api/v1/ai/knowledge/document/get-content` | 获取文档全文 | admin |
| POST | `/api/v1/ai/knowledge/document/get-detail` | 获取文档详情 | admin |
| POST | `/api/v1/ai/knowledge/document/get-failed-list` | 获取失败文档列表 | admin |
| POST | `/api/v1/ai/knowledge/document/get-tree` | 获取知识库目录树 | admin |
| POST | `/api/v1/ai/knowledge/document/import` | 导入知识库文档 | admin |
| POST | `/api/v1/ai/knowledge/document/import-zip` | ZIP导入知识库 | admin |
| POST | `/api/v1/ai/knowledge/export/create` | 创建导出任务 | admin |
| POST | `/api/v1/ai/knowledge/export/export-zip` | ZIP导出知识库 | admin |
| POST | `/api/v1/ai/knowledge/export/get-status` | 获取导出状态 | admin |
| POST | `/api/v1/ai/knowledge/info/create` | 创建知识库 | admin |
| POST | `/api/v1/ai/knowledge/info/delete` | 删除知识库 | admin |
| POST | `/api/v1/ai/knowledge/info/get-list` | 获取知识库列表 | admin |
| POST | `/api/v1/ai/knowledge/info/get-one` | 获取知识库详情 | admin |
| POST | `/api/v1/ai/knowledge/info/search` | 搜索知识库切片 | admin |
| POST | `/api/v1/ai/knowledge/info/update` | 更新知识库 | admin |
| POST | `/api/v1/ai/knowledge/tool/get-chunk-relations` | CLI/调试-获取切片关联 | admin |
| POST | `/api/v1/ai/knowledge/tool/get-document-content` | CLI/调试-获取文档全文 | admin |
| POST | `/api/v1/ai/knowledge/tool/search` | CLI/调试-搜索知识库 | admin |
| POST | `/api/v1/ai/mcp/resolve/resolve` | 解析MCP服务端点 | all |
| POST | `/api/v1/ai/mcp/service/create` | 创建MCP服务 | admin |
| POST | `/api/v1/ai/mcp/service/delete` | 删除MCP服务 | admin |
| POST | `/api/v1/ai/mcp/service/get-list` | 获取MCP服务列表 | admin |
| POST | `/api/v1/ai/mcp/service/get-one` | 获取MCP服务详情 | admin |
| POST | `/api/v1/ai/mcp/service/refresh-tools` | 刷新MCP服务工具列表 | admin |
| POST | `/api/v1/ai/mcp/service/update` | 更新MCP服务 | admin |
| POST | `/api/v1/ai/mcp/tools/get-tools` | 获取会话MCP工具列表 | all |
| POST | `/api/v1/ai/memory/debug/archives/get-list` | 调试-归档记录列表 | platform |
| POST | `/api/v1/ai/memory/debug/dream-runs/get-list` | 调试-Dream运行记录列表 | platform |
| POST | `/api/v1/ai/memory/debug/overview` | 记忆调试总览 | platform |
| POST | `/api/v1/ai/memory/debug/pipeline-state/get-list` | 调试-流水线状态列表 | platform |
| POST | `/api/v1/ai/memory/debug/traces/get-list` | 调试-召回轨迹列表 | platform |
| POST | `/api/v1/ai/session/archive` | 归档会话 | admin |
| POST | `/api/v1/ai/session/create` | 创建会话 | admin |
| POST | `/api/v1/ai/session/delete` | 删除会话 | admin |
| POST | `/api/v1/ai/session/end` | 结束会话 | admin |
| POST | `/api/v1/ai/session/get-list` | 获取会话列表 | admin |
| POST | `/api/v1/ai/session/get-one` | 获取会话详情 | admin |
| POST | `/api/v1/ai/session/inject` | 注入消息到会话 | admin |
| POST | `/api/v1/ai/session/messages` | 获取会话消息列表 | admin |
| POST | `/api/v1/ai/skill/create` | 创建技能 | admin |
| POST | `/api/v1/ai/skill/delete` | 删除技能 | admin |
| POST | `/api/v1/ai/skill/download` | 下载技能zip包 | admin |
| POST | `/api/v1/ai/skill/get-list` | 获取技能列表 | admin |
| POST | `/api/v1/ai/skill/get-one` | 获取技能详情 | admin |
| POST | `/api/v1/ai/skill/publish` | 发布新版本 | admin |
| POST | `/api/v1/ai/skill/update` | 更新技能 | admin |
| POST | `/api/v1/ai/skill/upload` | 上传技能zip包 | admin |
| POST | `/api/v1/ai/skill/file/create` | 创建技能文件 | admin |
| POST | `/api/v1/ai/skill/file/delete` | 删除文件 | admin |
| POST | `/api/v1/ai/skill/file/get-one` | 获取文件内容 | admin |
| POST | `/api/v1/ai/skill/file/update` | 更新文件内容 | admin |
| POST | `/api/v1/ai/tool/create` | 创建工具 | admin |
| POST | `/api/v1/ai/tool/delete` | 删除工具 | admin |
| POST | `/api/v1/ai/tool/disable` | 停用工具 | admin |
| POST | `/api/v1/ai/tool/get-list` | 获取工具列表 | admin |
| POST | `/api/v1/ai/tool/get-one` | 获取工具详情 | admin |
| POST | `/api/v1/ai/tool/publish` | 发布工具 | admin |
| POST | `/api/v1/ai/tool/save-artifact` | 保存三件套 | admin |
| POST | `/api/v1/ai/tool/update` | 更新工具 | admin |
| POST | `/api/v1/ai/tool/run` | 手动运行工具 | admin |
| POST | `/api/v1/ai/tool/run-history` | 执行历史列表 | admin |
| POST | `/api/v1/ai/tool/run-status` | 查询执行状态 | admin |
| POST | `/api/v1/ai/udp/token/create` | 创建UDP会话令牌 | all |
| POST | `/api/v1/things/ai/mcp/message` | MCP消息发送 | all |
| POST | `/api/v1/things/ai/mcp/run` | Stateless MCP HTTP | all |
| GET | `/api/v1/things/ai/mcp/sse` | SSE连接 | all |
| POST | `/api/v1/things/ai/mcp/sse` | SSE连接（POST） | all |
| POST | `/api/v1/things/alarm/info/create` | 新增告警 | admin |
| POST | `/api/v1/things/alarm/info/delete` | 删除告警 | admin |
| POST | `/api/v1/things/alarm/info/get-list` | 获取告警信息列表 | admin |
| POST | `/api/v1/things/alarm/info/get-one` | 获取告警信息 | admin |
| POST | `/api/v1/things/alarm/info/update` | 更新告警 | admin |
| POST | `/api/v1/things/alarm/record/deal` | 处理告警 | admin |
| POST | `/api/v1/things/alarm/record/get-list` | 获取告警记录列表 | admin |
| POST | `/api/v1/things/alarm/scene/batch-create` | 更新告警和场景的关联 | admin |
| POST | `/api/v1/things/alarm/scene/delete` | 删除告警和场景的关联 | admin |
| POST | `/api/v1/things/alarm/scene/get-list` | 获取告警和场景的关联列表 | admin |

## 端点详情

### POST `/api/v1/ai/agent/group/create`

**说明**: 创建助手组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `asrConfigID` | string | 否 |  共享默认ASR配置ID |
| `code` | string | 是 |  助手组编码（租户内唯一） |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `defaultAgentId` | integer | 否 |  默认决策AgentID（0=未设置） (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  助手组ID (格式: int64) |
| `knowledgeBaseIds` | string | 否 |  共享知识库ID列表（JSON字符串，供组内Agent复用） |
| `llmConfigID` | string | 否 |  共享默认LLM配置ID |
| `mcpServiceIds` | string | 否 |  共享MCP服务ID列表（JSON字符串，供组内Agent复用） |
| `name` | string | 是 |  助手组名称 |
| `purpose` | string | 否 |  用途分类 default/customer_service/... |
| `skillIds` | string | 否 |  平台分发Skill ID列表（JSON字符串，供组内Agent复用，不表示Group或Clone拥有learned skills） |
| `status` | string | 否 |  状态 1-启用 2-禁用 |
| `systemPrompt` | string | 否 |  共享系统提示词模板（供组内Agent决策时继承） |
| `tenantCode` | string | 否 |  租户编码，"common"=平台级 |
| `ttsConfigID` | string | 否 |  共享默认TTS配置ID |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "asrConfigID": "string",
  "code": "string",
  "createdTime": 1,
  "defaultAgentId": 1,
  "desc": "string",
  "id": 1,
  "knowledgeBaseIds": "string",
  "llmConfigID": "string",
  "mcpServiceIds": "string",
  "name": "示例名称",
  "purpose": "string",
  "skillIds": "string",
  "status": "string",
  "systemPrompt": "string",
  "tenantCode": "string",
  "ttsConfigID": "string",
  "updatedTime": 1
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
ur api /api/v1/ai/agent/group/create \
  --body '{"asrConfigID": "string", "code": "string", "createdTime": 1, "defaultAgentId": 1, "desc": "string", "id": 1, "knowledgeBaseIds": "string", "llmConfigID": "string", "mcpServiceIds": "string", "name": "示例名称", "purpose": "string", "skillIds": "string", "status": "string", "systemPrompt": "string", "tenantCode": "string", "ttsConfigID": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/agent/group/delete`

**说明**: 删除助手组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/agent/group/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/agent/group/get-list`

**说明**: 获取助手组列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `purpose` | string | 否 |  用途分类过滤 |
| `tenantCode` | string | 否 |  租户编码过滤 |

**请求示例**:
```json
{
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "purpose": "string",
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
        "code": "string",
        "createdTime": 1,
        "defaultAgentId": 1,
        "desc": "string",
        "id": 1,
        "knowledgeBaseIds": "string",
        "llmConfigID": "string",
        "mcpServiceIds": "string",
        "name": "示例名称",
        "purpose": "string",
        "skillIds": "string",
        "status": "string",
        "systemPrompt": "string",
        "tenantCode": "string",
        "ttsConfigID": "string",
        "updatedTime": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/agent/group/get-list \
  --body '{"name": "示例名称", "page": {"page": 1, "pageSize": 1}, "purpose": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/agent/group/get-one`

**说明**: 获取助手组详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
    "asrConfigID": "string",
    "code": "string",
    "createdTime": 1,
    "defaultAgentId": 1,
    "desc": "string",
    "id": 1,
    "knowledgeBaseIds": "string",
    "llmConfigID": "string",
    "mcpServiceIds": "string",
    "name": "示例名称",
    "purpose": "string",
    "skillIds": "string",
    "status": "string",
    "systemPrompt": "string",
    "tenantCode": "string",
    "ttsConfigID": "string",
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/agent/group/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/agent/group/update`

**说明**: 更新助手组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `asrConfigID` | string | 否 |  共享默认ASR配置ID |
| `code` | string | 是 |  助手组编码（租户内唯一） |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `defaultAgentId` | integer | 否 |  默认决策AgentID（0=未设置） (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  助手组ID (格式: int64) |
| `knowledgeBaseIds` | string | 否 |  共享知识库ID列表（JSON字符串，供组内Agent复用） |
| `llmConfigID` | string | 否 |  共享默认LLM配置ID |
| `mcpServiceIds` | string | 否 |  共享MCP服务ID列表（JSON字符串，供组内Agent复用） |
| `name` | string | 是 |  助手组名称 |
| `purpose` | string | 否 |  用途分类 default/customer_service/... |
| `skillIds` | string | 否 |  平台分发Skill ID列表（JSON字符串，供组内Agent复用，不表示Group或Clone拥有learned skills） |
| `status` | string | 否 |  状态 1-启用 2-禁用 |
| `systemPrompt` | string | 否 |  共享系统提示词模板（供组内Agent决策时继承） |
| `tenantCode` | string | 否 |  租户编码，"common"=平台级 |
| `ttsConfigID` | string | 否 |  共享默认TTS配置ID |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "asrConfigID": "string",
  "code": "string",
  "createdTime": 1,
  "defaultAgentId": 1,
  "desc": "string",
  "id": 1,
  "knowledgeBaseIds": "string",
  "llmConfigID": "string",
  "mcpServiceIds": "string",
  "name": "示例名称",
  "purpose": "string",
  "skillIds": "string",
  "status": "string",
  "systemPrompt": "string",
  "tenantCode": "string",
  "ttsConfigID": "string",
  "updatedTime": 1
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
ur api /api/v1/ai/agent/group/update \
  --body '{"asrConfigID": "string", "code": "string", "createdTime": 1, "defaultAgentId": 1, "desc": "string", "id": 1, "knowledgeBaseIds": "string", "llmConfigID": "string", "mcpServiceIds": "string", "name": "示例名称", "purpose": "string", "skillIds": "string", "status": "string", "systemPrompt": "string", "tenantCode": "string", "ttsConfigID": "string", "updatedTime": 1}'
```

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
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态过滤 |
| `tenantCode` | string | 否 |  租户编码过滤 |

**请求示例**:
```json
{
  "groupID": 1,
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
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
  --body '{"groupID": 1, "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "status": "string", "tenantCode": "string"}'
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
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "page": {
    "page": 1,
    "pageSize": 1
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
  --body '{"cloneID": 1, "page": {"page": 1, "pageSize": 1}}'
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

### POST `/api/v1/ai/app/chat-send`

**说明**: 发送对话消息

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appCode` | string | 是 | 执行的应用 |
| `conversationID` | string | 否 | （选填）会话 ID，需要基于之前的聊天记录继续对话，必须传之前消息的 conversation_id。 |
| `inputs` | object | 否 |  |
| `query` | string | 是 | 用户输入/提问内容 |
| `user` | string | 是 | 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。 |

**请求示例**:
```json
{
  "appCode": "string",
  "conversationID": "string",
  "inputs": {},
  "query": "string",
  "user": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "answer": "string",
    "conversationID": "string",
    "createdAt": 1,
    "event": "string",
    "messageID": "string",
    "mode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/app/chat-send \
  --body '{"appCode": "string", "conversationID": "string", "inputs": {}, "query": "string", "user": "string"}'
```

### POST `/api/v1/ai/app/sse/chat-run`

**说明**: sse模式ai聊天运行

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appCode` | string | 是 | 执行的应用 |
| `conversationID` | string | 否 | （选填）会话 ID，需要基于之前的聊天记录继续对话，必须传之前消息的 conversation_id。 |
| `inputs` | object | 否 |  |
| `query` | string | 是 | 用户输入/提问内容 |
| `user` | string | 是 | 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。 |

**请求示例**:
```json
{
  "appCode": "string",
  "conversationID": "string",
  "inputs": {},
  "query": "string",
  "user": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "answer": "string",
    "conversationID": "string",
    "createdAt": 1,
    "event": "string",
    "messageID": "string",
    "mode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/app/sse/chat-run \
  --body '{"appCode": "string", "conversationID": "string", "inputs": {}, "query": "string", "user": "string"}'
```

### POST `/api/v1/ai/app/sse/workflow-run`

**说明**: sse模式执行 workflow

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appCode` | string | 是 | 执行的应用 |
| `inputs` | object | 否 |  |
| `user` | string | 是 | 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。 |

**请求示例**:
```json
{
  "appCode": "string",
  "inputs": {},
  "user": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdAt": 1,
    "elapsedTime": 1,
    "error": "string",
    "event": "string",
    "finishedAt": 1,
    "inputs": {},
    "outputs": {},
    "status": "string",
    "taskID": "string",
    "workflowID": "string",
    "workflowRunID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/app/sse/workflow-run \
  --body '{"appCode": "string", "inputs": {}, "user": "string"}'
```

### POST `/api/v1/ai/app/workflow-run`

**说明**: 执行 workflow

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appCode` | string | 是 | 执行的应用 |
| `inputs` | object | 否 |  |
| `user` | string | 是 | 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。 |

**请求示例**:
```json
{
  "appCode": "string",
  "inputs": {},
  "user": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdAt": 1,
    "elapsedTime": 1,
    "error": "string",
    "event": "string",
    "finishedAt": 1,
    "inputs": {},
    "outputs": {},
    "status": "string",
    "taskID": "string",
    "workflowID": "string",
    "workflowRunID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/app/workflow-run \
  --body '{"appCode": "string", "inputs": {}, "user": "string"}'
```

### POST `/api/v1/ai/mcp/run`

**说明**: 添加规则链

**权限**: -

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

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
ur api /api/v1/ai/mcp/run \
  --body '{}'
```

### POST `/api/v1/ai/chat/available-tools`

**说明**: 获取当前会话可用的后端工具列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 是 |  助手ID (格式: int64) |
| `audioConfig` | object | 否 |  |
| `audioConfig.channels` | integer | 否 |  声道数，默认 1 (格式: int32) |
| `audioConfig.format` | string | 否 |  音频格式：opus/mp3/wav，默认 opus |
| `audioConfig.sampleRate` | integer | 否 |  采样率，默认 24000 (格式: int32) |
| `cloneID` | integer | 否 |  分身ID（0=不使用分身记忆） (格式: int64) |
| `currentAppID` | string | 否 |  当前页面所属应用ID |
| `currentAppName` | string | 否 |  当前页面所属应用名称 |
| `currentRoute` | string | 否 |  当前前端路由 |
| `messages` | array[ChatMessage] | 是 |  对话消息列表 |
| `outputModalities` | array[string] | 否 |  输出模态：text / audio，默认 ["text"] |
| `sessionID` | string | 否 |  会话ID（为空则自动生成） |
| `stream` | boolean | 否 |  是否流式返回（SSE） (格式: boolean) |
| `toolID` | integer | 否 |  工具ID（用于工具编辑场景，注入文件操作工具） (格式: int64) |
| `tools` | array[ToolDef] | 否 |  前端注册的工具定义列表 |

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
  --body '{"agentID": 1, "audioConfig": {"channels": 1, "format": "string", "sampleRate": 1}, "cloneID": 1, "currentAppID": "string", "currentAppName": "示例名称", "currentRoute": "string", "messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "outputModalities": ["string"], "sessionID": "string", "stream": true, "toolID": 1, "tools": [{"available": "string", "description": "string", "formSchema": {"cancelLabel": "string", "description": "string", "fields": [{"default": "...", "key": "...", "label": "...", "max": "...", "min": "...", "options": "...", "placeholder": "...", "required": "...", "type": "..."}], "submitLabel": "string", "title": "string"}, "interactionType": "string", "name": "示例名称", "parameters": {}}]}'
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
| `agentID` | integer | 是 |  助手ID (格式: int64) |
| `audioConfig` | object | 否 |  |
| `audioConfig.channels` | integer | 否 |  声道数，默认 1 (格式: int32) |
| `audioConfig.format` | string | 否 |  音频格式：opus/mp3/wav，默认 opus |
| `audioConfig.sampleRate` | integer | 否 |  采样率，默认 24000 (格式: int32) |
| `cloneID` | integer | 否 |  分身ID（0=不使用分身记忆） (格式: int64) |
| `currentAppID` | string | 否 |  当前页面所属应用ID |
| `currentAppName` | string | 否 |  当前页面所属应用名称 |
| `currentRoute` | string | 否 |  当前前端路由 |
| `messages` | array[ChatMessage] | 是 |  对话消息列表 |
| `outputModalities` | array[string] | 否 |  输出模态：text / audio，默认 ["text"] |
| `sessionID` | string | 否 |  会话ID（为空则自动生成） |
| `stream` | boolean | 否 |  是否流式返回（SSE） (格式: boolean) |
| `toolID` | integer | 否 |  工具ID（用于工具编辑场景，注入文件操作工具） (格式: int64) |
| `tools` | array[ToolDef] | 否 |  前端注册的工具定义列表 |

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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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
  --body '{"agentID": 1, "audioConfig": {"channels": 1, "format": "string", "sampleRate": 1}, "cloneID": 1, "currentAppID": "string", "currentAppName": "示例名称", "currentRoute": "string", "messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "outputModalities": ["string"], "sessionID": "string", "stream": true, "toolID": 1, "tools": [{"available": "string", "description": "string", "formSchema": {"cancelLabel": "string", "description": "string", "fields": [{"default": "...", "key": "...", "label": "...", "max": "...", "min": "...", "options": "...", "placeholder": "...", "required": "...", "type": "..."}], "submitLabel": "string", "title": "string"}, "interactionType": "string", "name": "示例名称", "parameters": {}}]}'
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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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

### POST `/api/v1/ai/clone/create`

**说明**: 创建分身

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 是 |  关联Agent ID (格式: int64) |
| `code` | string | 否 |  分身编码（全局唯一，可为空） |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  分身ID (格式: int64) |
| `metadata` | string | 否 |  扩展元数据（JSON字符串） |
| `name` | string | 是 |  分身名称 |
| `personaPrompt` | string | 否 |  人设补充提示词 |
| `status` | integer | 否 |  状态 1-启用 2-禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "code": "string",
  "createdTime": 1,
  "desc": "string",
  "id": 1,
  "metadata": "string",
  "name": "示例名称",
  "personaPrompt": "string",
  "status": 1,
  "tenantCode": "string",
  "updatedTime": 1
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
ur api /api/v1/ai/clone/create \
  --body '{"agentID": 1, "code": "string", "createdTime": 1, "desc": "string", "id": 1, "metadata": "string", "name": "示例名称", "personaPrompt": "string", "status": 1, "tenantCode": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/clone/delete`

**说明**: 删除分身

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone/get-list`

**说明**: 获取分身列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 |  按Agent ID过滤 (格式: int64) |
| `code` | string | 否 |  按code精确查询 |
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  状态过滤 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码过滤 |

**请求示例**:
```json
{
  "agentID": 1,
  "code": "string",
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": 1,
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
        "agentID": 1,
        "code": "string",
        "createdTime": 1,
        "desc": "string",
        "id": 1,
        "metadata": "string",
        "name": "示例名称",
        "personaPrompt": "string",
        "status": 1,
        "tenantCode": "string",
        "updatedTime": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/get-list \
  --body '{"agentID": 1, "code": "string", "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "status": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/clone/get-one`

**说明**: 获取分身详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
    "agentID": 1,
    "code": "string",
    "createdTime": 1,
    "desc": "string",
    "id": 1,
    "metadata": "string",
    "name": "示例名称",
    "personaPrompt": "string",
    "status": 1,
    "tenantCode": "string",
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone/update`

**说明**: 更新分身

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 是 |  关联Agent ID (格式: int64) |
| `code` | string | 否 |  分身编码（全局唯一，可为空） |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  分身ID (格式: int64) |
| `metadata` | string | 否 |  扩展元数据（JSON字符串） |
| `name` | string | 是 |  分身名称 |
| `personaPrompt` | string | 否 |  人设补充提示词 |
| `status` | integer | 否 |  状态 1-启用 2-禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "code": "string",
  "createdTime": 1,
  "desc": "string",
  "id": 1,
  "metadata": "string",
  "name": "示例名称",
  "personaPrompt": "string",
  "status": 1,
  "tenantCode": "string",
  "updatedTime": 1
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
ur api /api/v1/ai/clone/update \
  --body '{"agentID": 1, "code": "string", "createdTime": 1, "desc": "string", "id": 1, "metadata": "string", "name": "示例名称", "personaPrompt": "string", "status": 1, "tenantCode": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/clone/memory/archives/get-list`

**说明**: 获取归档记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `archiveReason` | string | 否 |  归档原因过滤 |
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "archiveReason": "string",
  "cloneID": 1,
  "page": {
    "page": 1,
    "pageSize": 1
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
        "archiveReason": "string",
        "cloneID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "recordID": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/archives/get-list \
  --body '{"archiveReason": "string", "cloneID": 1, "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/ai/clone/memory/compact`

**说明**: 压缩记忆

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/compact \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/clone/memory/create`

**说明**: 添加记忆

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `content` | string | 是 |  记忆内容 |
| `importance` | number | 否 |  重要性评分 (格式: double) |
| `keywords` | string | 否 |  关键词（JSON字符串） |
| `memoryType` | string | 是 |  记忆类型 |

**请求示例**:
```json
{
  "cloneID": 1,
  "content": "string",
  "importance": 1,
  "keywords": "string",
  "memoryType": "string"
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
ur api /api/v1/ai/clone/memory/create \
  --body '{"cloneID": 1, "content": "string", "importance": 1, "keywords": "string", "memoryType": "string"}'
```

### POST `/api/v1/ai/clone/memory/delete`

**说明**: 删除记忆

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `memoryID` | integer | 是 |  记忆ID (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "memoryID": 1
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
ur api /api/v1/ai/clone/memory/delete \
  --body '{"cloneID": 1, "memoryID": 1}'
```

### POST `/api/v1/ai/clone/memory/dream-runs/get-list`

**说明**: 获取 Dream 运行记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态过滤 running/success/skipped/failed |

**请求示例**:
```json
{
  "cloneID": 1,
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "cloneID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "inputWindowEnd": "string",
        "inputWindowStart": "string",
        "status": "string",
        "summary": "string",
        "triggerType": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/dream-runs/get-list \
  --body '{"cloneID": 1, "page": {"page": 1, "pageSize": 1}, "status": "string"}'
```

### POST `/api/v1/ai/clone/memory/dream/run`

**说明**: 手动触发 Dream 整理

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `triggerType` | string | 否 |  触发类型 manual/repair |

**请求示例**:
```json
{
  "cloneID": 1,
  "triggerType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "archivedCount": "string",
    "newRecordID": "string",
    "runID": "string",
    "status": "string",
    "summary": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/dream/run \
  --body '{"cloneID": 1, "triggerType": "string"}'
```

### POST `/api/v1/ai/clone/memory/entities/get-list`

**说明**: 获取记忆实体列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `entityKey` | string | 否 |  实体键过滤 |
| `entityType` | string | 否 |  实体类型过滤 |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "entityKey": "string",
  "entityType": "string",
  "page": {
    "page": 1,
    "pageSize": 1
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
        "confidence": 1,
        "entityKey": "string",
        "entityName": "string",
        "entityType": "string",
        "recordCount": "string",
        "role": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/entities/get-list \
  --body '{"cloneID": 1, "entityKey": "string", "entityType": "string", "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/ai/clone/memory/profile/get-one`

**说明**: 获取记忆画像

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "ownerID": "string",
    "ownerType": "string",
    "profileText": "string",
    "updatedTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/profile/get-one \
  --body '{"cloneID": 1}'
```

### POST `/api/v1/ai/clone/memory/records/get-list`

**说明**: 获取记忆记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `endTime` | string | 否 |  结束时间戳（秒） |
| `memoryKind` | string | 否 |  记忆类别 |
| `memoryLevel` | string | 否 |  记忆层级 l1/l2/profile |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `query` | string | 否 |  搜索关键词（摘要/内容模糊匹配） |
| `startTime` | string | 否 |  开始时间戳（秒） |
| `status` | string | 否 |  状态 1-活跃 2-已衰减 3-已归档 |

**请求示例**:
```json
{
  "cloneID": 1,
  "endTime": "2026-01-01T00:00:00Z",
  "memoryKind": "string",
  "memoryLevel": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "query": "string",
  "startTime": "2026-01-01T00:00:00Z",
  "status": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "cloneID": 1,
        "content": "string",
        "id": "string",
        "importance": 1,
        "keywords": "string",
        "memoryKind": "string",
        "memoryLevel": "string",
        "sessionID": "string",
        "status": "string",
        "summary": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/records/get-list \
  --body '{"cloneID": 1, "endTime": "2026-01-01T00:00:00Z", "memoryKind": "string", "memoryLevel": "string", "page": {"page": 1, "pageSize": 1}, "query": "string", "startTime": "2026-01-01T00:00:00Z", "status": "string"}'
```

### POST `/api/v1/ai/clone/memory/search`

**说明**: 搜索记忆

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `memoryType` | string | 否 |  记忆类型过滤 |
| `query` | string | 是 |  搜索关键词 |
| `topK` | integer | 否 |  返回条数（默认10） (格式: int64) |

**请求示例**:
```json
{
  "cloneID": 1,
  "memoryType": "string",
  "query": "string",
  "topK": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "accessCount": 1,
        "cloneID": 1,
        "content": "string",
        "createdTime": 1,
        "id": 1,
        "importance": 1,
        "keywords": "string",
        "memoryType": "string",
        "sessionID": "string",
        "status": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/search \
  --body '{"cloneID": 1, "memoryType": "string", "query": "string", "topK": 1}'
```

### POST `/api/v1/ai/clone/memory/traces/get-list`

**说明**: 获取召回轨迹列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | integer | 是 |  分身ID (格式: int64) |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `query` | string | 否 |  查询内容模糊搜索 |

**请求示例**:
```json
{
  "cloneID": 1,
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "query": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "enoughAt": "string",
        "path": "string",
        "query": "string",
        "resultSummary": "string",
        "traceID": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/clone/memory/traces/get-list \
  --body '{"cloneID": 1, "page": {"page": 1, "pageSize": 1}, "query": "string"}'
```

### POST `/api/v1/ai/config/info/create`

**说明**: 创建配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 否 |  配置类型：llm/tts/asr/vad |
| `configID` | string | 否 |  配置唯一标识 |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `id` | integer | 否 |  配置ID (格式: int64) |
| `inputModalities` | string | 否 |  输入模态（JSON字符串） |
| `isDefault` | boolean | 否 |  是否默认 (格式: boolean) |
| `jsonData` | string | 否 |  配置参数（JSON字符串） |
| `modelType` | string | 否 |  模型级别：small/medium/large/xlarge |
| `name` | string | 否 |  配置名称 |
| `outputModalities` | string | 否 |  输出模态（JSON字符串） |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "configID": "string",
  "enabled": true,
  "id": 1,
  "inputModalities": "string",
  "isDefault": true,
  "jsonData": "string",
  "modelType": "string",
  "name": "示例名称",
  "outputModalities": "string",
  "tenantCode": "string"
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
ur api /api/v1/ai/config/info/create \
  --body '{"cfgType": "string", "configID": "string", "enabled": true, "id": 1, "inputModalities": "string", "isDefault": true, "jsonData": "string", "modelType": "string", "name": "示例名称", "outputModalities": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/delete`

**说明**: 删除配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  配置ID (格式: int64) |
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
ur api /api/v1/ai/config/info/delete \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/get-default`

**说明**: 获取默认配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 是 |  配置类型 |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "cfgType": "string",
    "configID": "string",
    "enabled": true,
    "id": 1,
    "inputModalities": "string",
    "isDefault": true,
    "jsonData": "string",
    "modelType": "string",
    "name": "示例名称",
    "outputModalities": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/config/info/get-default \
  --body '{"cfgType": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/get-list`

**说明**: 获取配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 否 |  配置类型过滤 |
| `onlyEnabled` | boolean | 否 |  仅返回启用的 (格式: boolean) |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "onlyEnabled": true,
  "page": {
    "page": 1,
    "pageSize": 1
  },
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
        "cfgType": "string",
        "configID": "string",
        "enabled": true,
        "id": 1,
        "inputModalities": "string",
        "isDefault": true,
        "jsonData": "string",
        "modelType": "string",
        "name": "示例名称",
        "outputModalities": "string",
        "tenantCode": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/config/info/get-list \
  --body '{"cfgType": "string", "onlyEnabled": true, "page": {"page": 1, "pageSize": 1}, "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/get-one`

**说明**: 获取配置详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  配置ID (格式: int64) |
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
    "cfgType": "string",
    "configID": "string",
    "enabled": true,
    "id": 1,
    "inputModalities": "string",
    "isDefault": true,
    "jsonData": "string",
    "modelType": "string",
    "name": "示例名称",
    "outputModalities": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/config/info/get-one \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/test`

**说明**: 测试配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 是 |  配置类型，目前先支持 llm |
| `configID` | string | 否 |  配置唯一标识，仅用于调试展示 |
| `inputModalities` | string | 否 |  输入模态（JSON字符串） |
| `jsonData` | string | 是 |  配置参数（JSON字符串） |
| `messages` | array[ChatMessage] | 是 |  调试输入消息 |
| `modelType` | string | 否 |  模型级别 |
| `name` | string | 否 |  配置名称，仅用于调试展示 |
| `outputModalities` | string | 否 |  输出模态（JSON字符串） |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "configID": "string",
  "inputModalities": "string",
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
  "modelType": "string",
  "name": "示例名称",
  "outputModalities": "string",
  "tenantCode": "string"
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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
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
    "role": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/config/info/test \
  --body '{"cfgType": "string", "configID": "string", "inputModalities": "string", "jsonData": "string", "messages": [{"contents": [{"audioData": "string", "audioFormat": "string", "fileMime": "string", "fileName": "string", "fileUrl": "string", "imageUrl": "string", "text": "string", "type": "string", "videoUrl": "string"}], "role": "string", "toolCallID": "string", "toolCalls": [{"arguments": "string", "id": "string", "name": "示例名称"}]}], "modelType": "string", "name": "示例名称", "outputModalities": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/config/info/update`

**说明**: 更新配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 否 |  配置类型：llm/tts/asr/vad |
| `configID` | string | 否 |  配置唯一标识 |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `id` | integer | 否 |  配置ID (格式: int64) |
| `inputModalities` | string | 否 |  输入模态（JSON字符串） |
| `isDefault` | boolean | 否 |  是否默认 (格式: boolean) |
| `jsonData` | string | 否 |  配置参数（JSON字符串） |
| `modelType` | string | 否 |  模型级别：small/medium/large/xlarge |
| `name` | string | 否 |  配置名称 |
| `outputModalities` | string | 否 |  输出模态（JSON字符串） |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "configID": "string",
  "enabled": true,
  "id": 1,
  "inputModalities": "string",
  "isDefault": true,
  "jsonData": "string",
  "modelType": "string",
  "name": "示例名称",
  "outputModalities": "string",
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
ur api /api/v1/ai/config/info/update \
  --body '{"cfgType": "string", "configID": "string", "enabled": true, "id": 1, "inputModalities": "string", "isDefault": true, "jsonData": "string", "modelType": "string", "name": "示例名称", "outputModalities": "string", "tenantCode": "string"}'
```

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

### POST `/api/v1/ai/knowledge/compile/get-status`

**说明**: 获取编译状态

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "knowledgeID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "compiledDocumentCount": 1,
    "currentTask": {
      "createdTime": 1,
      "errorMessage": "string",
      "finishedTime": 1,
      "id": 1,
      "knowledgeID": 1,
      "startedTime": 1,
      "status": "string",
      "triggerType": "string",
      "updatedTime": 1
    },
    "documentCount": 1,
    "failedDocumentCount": 1,
    "knowledgeID": 1,
    "relationMessage": "string",
    "relationStatus": "string",
    "searchFailedDocumentCount": 1,
    "searchMessage": "string",
    "searchReadyDocumentCount": 1,
    "searchStatus": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/compile/get-status \
  --body '{"knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/compile/retry`

**说明**: 触发知识库重编译

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |
| `mode` | string | 否 |  |

**请求示例**:
```json
{
  "knowledgeID": 1,
  "mode": "string"
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
ur api /api/v1/ai/knowledge/compile/retry \
  --body '{"knowledgeID": 1, "mode": "string"}'
```

### POST `/api/v1/ai/knowledge/document/chunk/relation/get-list`

**说明**: 获取切片关联列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `artifactID` | integer | 是 | 格式: int64 |
| `limit` | integer | 否 | 格式: int64 |
| `minScore` | number | 否 | 格式: double |
| `relationTypes` | array[string] | 否 |  |

**请求示例**:
```json
{
  "artifactID": 1,
  "limit": 1,
  "minScore": 1,
  "relationTypes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "artifact": {
      "artifactType": "string",
      "charEnd": 1,
      "charStart": 1,
      "chunkIndex": 1,
      "content": "string",
      "createdTime": 1,
      "documentID": 1,
      "headingPath": "string",
      "id": 1,
      "knowledgeID": 1,
      "relatedCount": 1,
      "sourceDocumentIDs": [
        1
      ],
      "title": "string",
      "tokenEstimate": 1,
      "updatedTime": 1
    },
    "relations": [
      {
        "cosineScore": 1,
        "evidence": "string",
        "fromArtifactID": 1,
        "fromDocumentID": 1,
        "id": 1,
        "lexicalOverlap": 1,
        "llmConfidence": 1,
        "relationType": "string",
        "score": 1,
        "targetArtifact": {
          "artifactType": "string",
          "charEnd": 1,
          "charStart": 1,
          "chunkIndex": 1,
          "content": "string",
          "createdTime": 1,
          "documentID": 1,
          "headingPath": "string",
          "id": 1,
          "knowledgeID": 1,
          "relatedCount": 1,
          "sourceDocumentIDs": [
            1
          ],
          "title": "string",
          "tokenEstimate": 1,
          "updatedTime": 1
        },
        "targetDocument": {
          "children": [
            {
              "children": "...",
              "compileStatus": "...",
              "id": "...",
              "knowledgeID": "...",
              "lastError": "...",
              "name": "...",
              "nodeType": "...",
              "parentID": "...",
              "parseStatus": "...",
              "path": "...",
              "searchMessage": "...",
              "searchStatus": "...",
              "uri": "..."
            }
          ],
          "compileStatus": "string",
          "id": 1,
          "knowledgeID": 1,
          "lastError": "string",
          "name": "string",
          "nodeType": "string",
          "parentID": 1,
          "parseStatus": "string",
          "path": "string",
          "searchMessage": "string",
          "searchStatus": "string",
          "uri": "string"
        },
        "toArtifactID": 1,
        "toDocumentID": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/chunk/relation/get-list \
  --body '{"artifactID": 1, "limit": 1, "minScore": 1, "relationTypes": ["string"]}'
```

### POST `/api/v1/ai/knowledge/document/get-content`

**说明**: 获取文档全文

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `documentID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "documentID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "documentID": 1,
    "parsedFullText": "string",
    "sourceTextReady": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/get-content \
  --body '{"documentID": 1}'
```

### POST `/api/v1/ai/knowledge/document/get-detail`

**说明**: 获取文档详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `documentID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "documentID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "articles": [
      {
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "createdTime": 1,
        "documentID": 1,
        "headingPath": "string",
        "id": 1,
        "knowledgeID": 1,
        "relatedCount": 1,
        "sourceDocumentIDs": [
          1
        ],
        "title": "string",
        "tokenEstimate": 1,
        "updatedTime": 1
      }
    ],
    "chunks": [
      {
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "createdTime": 1,
        "documentID": 1,
        "headingPath": "string",
        "id": 1,
        "knowledgeID": 1,
        "relatedCount": 1,
        "sourceDocumentIDs": [
          1
        ],
        "title": "string",
        "tokenEstimate": 1,
        "updatedTime": 1
      }
    ],
    "compileStatus": "string",
    "concepts": [
      {
        "artifactType": "string",
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "createdTime": 1,
        "documentID": 1,
        "headingPath": "string",
        "id": 1,
        "knowledgeID": 1,
        "relatedCount": 1,
        "sourceDocumentIDs": [
          1
        ],
        "title": "string",
        "tokenEstimate": 1,
        "updatedTime": 1
      }
    ],
    "document": {
      "children": [
        {
          "children": [
            {
              "children": "...",
              "compileStatus": "...",
              "id": "...",
              "knowledgeID": "...",
              "lastError": "...",
              "name": "...",
              "nodeType": "...",
              "parentID": "...",
              "parseStatus": "...",
              "path": "...",
              "searchMessage": "...",
              "searchStatus": "...",
              "uri": "..."
            }
          ],
          "compileStatus": "string",
          "id": 1,
          "knowledgeID": 1,
          "lastError": "string",
          "name": "string",
          "nodeType": "string",
          "parentID": 1,
          "parseStatus": "string",
          "path": "string",
          "searchMessage": "string",
          "searchStatus": "string",
          "uri": "string"
        }
      ],
      "compileStatus": "string",
      "id": 1,
      "knowledgeID": 1,
      "lastError": "string",
      "name": "string",
      "nodeType": "string",
      "parentID": 1,
      "parseStatus": "string",
      "path": "string",
      "searchMessage": "string",
      "searchStatus": "string",
      "uri": "string"
    },
    "lastCompileTask": {
      "createdTime": 1,
      "errorMessage": "string",
      "finishedTime": 1,
      "id": 1,
      "knowledgeID": 1,
      "startedTime": 1,
      "status": "string",
      "triggerType": "string",
      "updatedTime": 1
    },
    "lastError": "string",
    "previewText": "string",
    "sourceDownloadURL": "string",
    "summary": {
      "artifactType": "string",
      "charEnd": 1,
      "charStart": 1,
      "chunkIndex": 1,
      "content": "string",
      "createdTime": 1,
      "documentID": 1,
      "headingPath": "string",
      "id": 1,
      "knowledgeID": 1,
      "relatedCount": 1,
      "sourceDocumentIDs": [
        1
      ],
      "title": "string",
      "tokenEstimate": 1,
      "updatedTime": 1
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/get-detail \
  --body '{"documentID": 1}'
```

### POST `/api/v1/ai/knowledge/document/get-failed-list`

**说明**: 获取失败文档列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "knowledgeID": 1,
  "page": {
    "page": 1,
    "pageSize": 1
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
        "children": [
          {
            "children": [],
            "compileStatus": "string",
            "id": 1,
            "knowledgeID": 1,
            "lastError": "string",
            "name": "string",
            "nodeType": "string",
            "parentID": 1,
            "parseStatus": "string",
            "path": "string",
            "searchMessage": "string",
            "searchStatus": "string",
            "uri": "string"
          }
        ],
        "compileStatus": "string",
        "id": 1,
        "knowledgeID": 1,
        "lastError": "string",
        "name": "string",
        "nodeType": "string",
        "parentID": 1,
        "parseStatus": "string",
        "path": "string",
        "searchMessage": "string",
        "searchStatus": "string",
        "uri": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/get-failed-list \
  --body '{"knowledgeID": 1, "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/ai/knowledge/document/get-tree`

**说明**: 获取知识库目录树

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "knowledgeID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "knowledgeID": 1,
    "nodes": [
      {
        "children": [
          {
            "children": [],
            "compileStatus": "string",
            "id": 1,
            "knowledgeID": 1,
            "lastError": "string",
            "name": "string",
            "nodeType": "string",
            "parentID": 1,
            "parseStatus": "string",
            "path": "string",
            "searchMessage": "string",
            "searchStatus": "string",
            "uri": "string"
          }
        ],
        "compileStatus": "string",
        "id": 1,
        "knowledgeID": 1,
        "lastError": "string",
        "name": "string",
        "nodeType": "string",
        "parentID": 1,
        "parseStatus": "string",
        "path": "string",
        "searchMessage": "string",
        "searchStatus": "string",
        "uri": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/get-tree \
  --body '{"knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/document/import`

**说明**: 导入知识库文档

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `items` | array[KnowledgeImportItem] | 是 |  |
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "items": [
    {
      "filePath": "string",
      "name": "string",
      "path": "string",
      "rawContent": "string"
    }
  ],
  "knowledgeID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "importedCount": 1,
    "taskID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/import \
  --body '{"items": [{"filePath": "string", "name": "string", "path": "string", "rawContent": "string"}], "knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/document/import-zip`

**说明**: ZIP导入知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `fileURL` | string | 是 |  |
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "fileURL": "string",
  "knowledgeID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "importedCount": 1,
    "taskID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/document/import-zip \
  --body '{"fileURL": "string", "knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/export/create`

**说明**: 创建导出任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "knowledgeID": 1
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
ur api /api/v1/ai/knowledge/export/create \
  --body '{"knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/export/export-zip`

**说明**: ZIP导出知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `knowledgeID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "knowledgeID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "downloadURL": "string",
    "taskID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/export/export-zip \
  --body '{"knowledgeID": 1}'
```

### POST `/api/v1/ai/knowledge/export/get-status`

**说明**: 获取导出状态

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `taskID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "taskID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "downloadURL": "string",
    "errorMessage": "string",
    "status": "string",
    "taskID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/export/get-status \
  --body '{"taskID": 1}'
```

### POST `/api/v1/ai/knowledge/info/create`

**说明**: 创建知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `compiledDocumentCount` | integer | 否 | 格式: int64 |
| `createdTime` | integer | 否 | 格式: int64 |
| `desc` | string | 否 |  |
| `documentCount` | integer | 否 | 格式: int64 |
| `failedDocumentCount` | integer | 否 | 格式: int64 |
| `id` | integer | 否 | 格式: int64 |
| `lastCompileTime` | integer | 否 | 格式: int64 |
| `name` | string | 是 |  |
| `status` | string | 否 |  |
| `tenantCode` | string | 否 |  |
| `updatedTime` | integer | 否 | 格式: int64 |

**请求示例**:
```json
{
  "compiledDocumentCount": 1,
  "createdTime": 1,
  "desc": "string",
  "documentCount": 1,
  "failedDocumentCount": 1,
  "id": 1,
  "lastCompileTime": 1,
  "name": "string",
  "status": "string",
  "tenantCode": "string",
  "updatedTime": 1
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
ur api /api/v1/ai/knowledge/info/create \
  --body '{"compiledDocumentCount": 1, "createdTime": 1, "desc": "string", "documentCount": 1, "failedDocumentCount": 1, "id": 1, "lastCompileTime": 1, "name": "string", "status": "string", "tenantCode": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/knowledge/info/delete`

**说明**: 删除知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 | 格式: int64 |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/knowledge/info/get-list`

**说明**: 获取知识库列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  |
| `tenantCode` | string | 否 |  |

**请求示例**:
```json
{
  "name": "string",
  "page": {
    "page": 1,
    "pageSize": 1
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
        "compiledDocumentCount": 1,
        "createdTime": 1,
        "desc": "string",
        "documentCount": 1,
        "failedDocumentCount": 1,
        "id": 1,
        "lastCompileTime": 1,
        "name": "string",
        "status": "string",
        "tenantCode": "string",
        "updatedTime": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/info/get-list \
  --body '{"name": "string", "page": {"page": 1, "pageSize": 1}, "status": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/knowledge/info/get-one`

**说明**: 获取知识库详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 | 格式: int64 |

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
    "compiledDocumentCount": 1,
    "createdTime": 1,
    "desc": "string",
    "documentCount": 1,
    "failedDocumentCount": 1,
    "id": 1,
    "lastCompileTime": 1,
    "name": "string",
    "status": "string",
    "tenantCode": "string",
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/knowledge/info/search`

**说明**: 搜索知识库切片

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 | 格式: int64 |
| `knowledgeBaseIDs` | array[integer] | 否 |  |
| `minScore` | number | 否 | 格式: double |
| `query` | string | 是 |  |
| `topK` | integer | 否 | 格式: int64 |

**请求示例**:
```json
{
  "agentID": 1,
  "knowledgeBaseIDs": [
    1
  ],
  "minScore": 1,
  "query": "string",
  "topK": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "chunks": [
      {
        "artifactID": 1,
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "confidenceBand": "string",
    "documents": [
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
    "expandedChunks": [
      {
        "artifactID": 1,
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "list": [
      {
        "artifactID": 1,
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "originalQuery": "string",
    "rewrittenQuery": "string",
    "searchReady": true,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/info/search \
  --body '{"agentID": 1, "knowledgeBaseIDs": [1], "minScore": 1, "query": "string", "topK": 1}'
```

### POST `/api/v1/ai/knowledge/info/update`

**说明**: 更新知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `compiledDocumentCount` | integer | 否 | 格式: int64 |
| `createdTime` | integer | 否 | 格式: int64 |
| `desc` | string | 否 |  |
| `documentCount` | integer | 否 | 格式: int64 |
| `failedDocumentCount` | integer | 否 | 格式: int64 |
| `id` | integer | 否 | 格式: int64 |
| `lastCompileTime` | integer | 否 | 格式: int64 |
| `name` | string | 是 |  |
| `status` | string | 否 |  |
| `tenantCode` | string | 否 |  |
| `updatedTime` | integer | 否 | 格式: int64 |

**请求示例**:
```json
{
  "compiledDocumentCount": 1,
  "createdTime": 1,
  "desc": "string",
  "documentCount": 1,
  "failedDocumentCount": 1,
  "id": 1,
  "lastCompileTime": 1,
  "name": "string",
  "status": "string",
  "tenantCode": "string",
  "updatedTime": 1
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
ur api /api/v1/ai/knowledge/info/update \
  --body '{"compiledDocumentCount": 1, "createdTime": 1, "desc": "string", "documentCount": 1, "failedDocumentCount": 1, "id": 1, "lastCompileTime": 1, "name": "string", "status": "string", "tenantCode": "string", "updatedTime": 1}'
```

### POST `/api/v1/ai/knowledge/tool/get-chunk-relations`

**说明**: CLI/调试-获取切片关联

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `artifactID` | integer | 是 | 格式: int64 |
| `limit` | integer | 否 | 格式: int64 |
| `minScore` | number | 否 | 格式: double |
| `relationTypes` | array[string] | 否 |  |

**请求示例**:
```json
{
  "artifactID": 1,
  "limit": 1,
  "minScore": 1,
  "relationTypes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "artifact": {
      "artifactType": "string",
      "charEnd": 1,
      "charStart": 1,
      "chunkIndex": 1,
      "content": "string",
      "createdTime": 1,
      "documentID": 1,
      "headingPath": "string",
      "id": 1,
      "knowledgeID": 1,
      "relatedCount": 1,
      "sourceDocumentIDs": [
        1
      ],
      "title": "string",
      "tokenEstimate": 1,
      "updatedTime": 1
    },
    "relations": [
      {
        "cosineScore": 1,
        "evidence": "string",
        "fromArtifactID": 1,
        "fromDocumentID": 1,
        "id": 1,
        "lexicalOverlap": 1,
        "llmConfidence": 1,
        "relationType": "string",
        "score": 1,
        "targetArtifact": {
          "artifactType": "string",
          "charEnd": 1,
          "charStart": 1,
          "chunkIndex": 1,
          "content": "string",
          "createdTime": 1,
          "documentID": 1,
          "headingPath": "string",
          "id": 1,
          "knowledgeID": 1,
          "relatedCount": 1,
          "sourceDocumentIDs": [
            1
          ],
          "title": "string",
          "tokenEstimate": 1,
          "updatedTime": 1
        },
        "targetDocument": {
          "children": [
            {
              "children": "...",
              "compileStatus": "...",
              "id": "...",
              "knowledgeID": "...",
              "lastError": "...",
              "name": "...",
              "nodeType": "...",
              "parentID": "...",
              "parseStatus": "...",
              "path": "...",
              "searchMessage": "...",
              "searchStatus": "...",
              "uri": "..."
            }
          ],
          "compileStatus": "string",
          "id": 1,
          "knowledgeID": 1,
          "lastError": "string",
          "name": "string",
          "nodeType": "string",
          "parentID": 1,
          "parseStatus": "string",
          "path": "string",
          "searchMessage": "string",
          "searchStatus": "string",
          "uri": "string"
        },
        "toArtifactID": 1,
        "toDocumentID": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/tool/get-chunk-relations \
  --body '{"artifactID": 1, "limit": 1, "minScore": 1, "relationTypes": ["string"]}'
```

### POST `/api/v1/ai/knowledge/tool/get-document-content`

**说明**: CLI/调试-获取文档全文

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `documentID` | integer | 是 | 格式: int64 |

**请求示例**:
```json
{
  "documentID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "documentID": 1,
    "parsedFullText": "string",
    "sourceTextReady": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/tool/get-document-content \
  --body '{"documentID": 1}'
```

### POST `/api/v1/ai/knowledge/tool/search`

**说明**: CLI/调试-搜索知识库

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 | 格式: int64 |
| `knowledgeBaseIDs` | array[integer] | 否 |  |
| `minScore` | number | 否 | 格式: double |
| `query` | string | 是 |  |
| `topK` | integer | 否 | 格式: int64 |

**请求示例**:
```json
{
  "agentID": 1,
  "knowledgeBaseIDs": [
    1
  ],
  "minScore": 1,
  "query": "string",
  "topK": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "chunks": [
      {
        "artifactID": 1,
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "confidenceBand": "string",
    "documents": [
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
    "expandedChunks": [
      {
        "artifactID": 1,
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "list": [
      {
        "artifactID": 1,
        "charEnd": 1,
        "charStart": 1,
        "chunkIndex": 1,
        "content": "string",
        "documentID": 1,
        "documentName": "string",
        "knowledgeID": 1,
        "path": "string",
        "relationType": "string",
        "score": 1,
        "snippet": "string",
        "sourceArtifactID": 1,
        "sourceType": "string"
      }
    ],
    "originalQuery": "string",
    "rewrittenQuery": "string",
    "searchReady": true,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/knowledge/tool/search \
  --body '{"agentID": 1, "knowledgeBaseIDs": [1], "minScore": 1, "query": "string", "topK": 1}'
```

### POST `/api/v1/ai/mcp/resolve/resolve`

**说明**: 解析MCP服务端点

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 是 |  服务名称 |
| `tenantCode` | string | 否 |  租户编码 |

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

### POST `/api/v1/ai/mcp/service/create`

**说明**: 创建MCP服务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authConfig` | string | 否 |  认证配置（JSON字符串） |
| `authType` | string | 否 |  认证类型 |
| `description` | string | 否 |  服务描述 |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `endpoint` | string | 是 |  服务端点URL |
| `id` | integer | 否 |  MCP服务ID (格式: int64) |
| `name` | string | 是 |  服务名称 |
| `tenantCode` | string | 否 |  租户编码 |
| `tools` | string | 否 |  工具列表（JSON字符串） |

**请求示例**:
```json
{
  "authConfig": "string",
  "authType": "string",
  "description": "string",
  "enabled": true,
  "endpoint": "string",
  "id": 1,
  "name": "示例名称",
  "tenantCode": "string",
  "tools": "string"
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
ur api /api/v1/ai/mcp/service/create \
  --body '{"authConfig": "string", "authType": "string", "description": "string", "enabled": true, "endpoint": "string", "id": 1, "name": "示例名称", "tenantCode": "string", "tools": "string"}'
```

### POST `/api/v1/ai/mcp/service/delete`

**说明**: 删除MCP服务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  MCP服务ID (格式: int64) |
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
ur api /api/v1/ai/mcp/service/delete \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/mcp/service/get-list`

**说明**: 获取MCP服务列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `enabled` | boolean | 否 |  启用状态过滤 (格式: boolean) |
| `nameLike` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "enabled": true,
  "nameLike": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
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
        "authConfig": "string",
        "authType": "string",
        "description": "string",
        "enabled": true,
        "endpoint": "string",
        "id": 1,
        "name": "示例名称",
        "tenantCode": "string",
        "tools": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/service/get-list \
  --body '{"enabled": true, "nameLike": "示例名称", "page": {"page": 1, "pageSize": 1}, "tenantCode": "string"}'
```

### POST `/api/v1/ai/mcp/service/get-one`

**说明**: 获取MCP服务详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  MCP服务ID (格式: int64) |
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
    "authConfig": "string",
    "authType": "string",
    "description": "string",
    "enabled": true,
    "endpoint": "string",
    "id": 1,
    "name": "示例名称",
    "tenantCode": "string",
    "tools": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/service/get-one \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/mcp/service/refresh-tools`

**说明**: 刷新MCP服务工具列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  MCP服务ID (格式: int64) |
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
    "tools": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/mcp/service/refresh-tools \
  --body '{"id": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/mcp/service/update`

**说明**: 更新MCP服务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authConfig` | string | 否 |  认证配置（JSON字符串） |
| `authType` | string | 否 |  认证类型 |
| `description` | string | 否 |  服务描述 |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `endpoint` | string | 是 |  服务端点URL |
| `id` | integer | 否 |  MCP服务ID (格式: int64) |
| `name` | string | 是 |  服务名称 |
| `tenantCode` | string | 否 |  租户编码 |
| `tools` | string | 否 |  工具列表（JSON字符串） |

**请求示例**:
```json
{
  "authConfig": "string",
  "authType": "string",
  "description": "string",
  "enabled": true,
  "endpoint": "string",
  "id": 1,
  "name": "示例名称",
  "tenantCode": "string",
  "tools": "string"
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
ur api /api/v1/ai/mcp/service/update \
  --body '{"authConfig": "string", "authType": "string", "description": "string", "enabled": true, "endpoint": "string", "id": 1, "name": "示例名称", "tenantCode": "string", "tools": "string"}'
```

### POST `/api/v1/ai/mcp/tools/get-tools`

**说明**: 获取会话MCP工具列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `sessionID` | string | 是 |  会话ID |
| `tenantCode` | string | 否 |  租户编码 |

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

### POST `/api/v1/ai/memory/debug/archives/get-list`

**说明**: 调试-归档记录列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `archiveReason` | string | 否 |  归档原因过滤 |
| `cloneID` | string | 否 |  分身ID（不传则查全部） |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "archiveReason": "string",
  "cloneID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
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
        "archiveReason": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "ownerID": "string",
        "ownerType": "string",
        "recordID": "string",
        "tenantCode": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/memory/debug/archives/get-list \
  --body '{"archiveReason": "string", "cloneID": "string", "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/ai/memory/debug/dream-runs/get-list`

**说明**: 调试-Dream运行记录列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | string | 否 |  分身ID（不传则查全部） |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态过滤 running/success/skipped/failed |

**请求示例**:
```json
{
  "cloneID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "inputWindowEnd": "string",
        "inputWindowStart": "string",
        "ownerID": "string",
        "ownerType": "string",
        "status": "string",
        "summary": "string",
        "tenantCode": "string",
        "triggerType": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/memory/debug/dream-runs/get-list \
  --body '{"cloneID": "string", "page": {"page": 1, "pageSize": 1}, "status": "string"}'
```

### POST `/api/v1/ai/memory/debug/overview`

**说明**: 记忆调试总览

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "archiveCount": "string",
    "dreamRunCount": "string",
    "l1Count": "string",
    "l2Count": "string",
    "lastErrorTime": "2026-01-01T00:00:00Z",
    "lastWriteTime": "2026-01-01T00:00:00Z",
    "pipelineStates": [
      {
        "id": "string",
        "ownerID": "string",
        "ownerType": "string",
        "stateKey": "string",
        "stateValue": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "profileCount": "string",
    "recordCount": "string",
    "traceCount": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/memory/debug/overview \
  --body '{}'
```

### POST `/api/v1/ai/memory/debug/pipeline-state/get-list`

**说明**: 调试-流水线状态列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | string | 否 |  分身ID（不传则查全部） |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "cloneID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
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
        "id": "string",
        "ownerID": "string",
        "ownerType": "string",
        "stateKey": "string",
        "stateValue": "string",
        "tenantCode": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/memory/debug/pipeline-state/get-list \
  --body '{"cloneID": "string", "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/ai/memory/debug/traces/get-list`

**说明**: 调试-召回轨迹列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneID` | string | 否 |  分身ID（不传则查全部） |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `query` | string | 否 |  查询内容模糊搜索 |

**请求示例**:
```json
{
  "cloneID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "query": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "enoughAt": "string",
        "ownerID": "string",
        "path": "string",
        "query": "string",
        "resultSummary": "string",
        "tenantCode": "string",
        "traceID": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/memory/debug/traces/get-list \
  --body '{"cloneID": "string", "page": {"page": 1, "pageSize": 1}, "query": "string"}'
```

### POST `/api/v1/ai/session/archive`

**说明**: 归档会话

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/archive \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/session/create`

**说明**: 创建会话

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 是 |  关联Agent ID (格式: int64) |
| `channel` | string | 否 |  渠道 text/voice |
| `cloneID` | integer | 否 |  关联Clone ID（0=无分身） (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |
| `userID` | integer | 否 |  用户ID (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "channel": "string",
  "cloneID": 1,
  "tenantCode": "string",
  "userID": 1
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
ur api /api/v1/ai/session/create \
  --body '{"agentID": 1, "channel": "string", "cloneID": 1, "tenantCode": "string", "userID": 1}'
```

### POST `/api/v1/ai/session/delete`

**说明**: 删除会话

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `sessionID` | string | 是 |  会话ID |
| `tenantCode` | string | 否 |  租户编码 |

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
    "success": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/delete \
  --body '{"sessionID": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/session/end`

**说明**: 结束会话

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/end \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/session/get-list`

**说明**: 获取会话列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentID` | integer | 否 |  按Agent ID过滤 (格式: int64) |
| `cloneID` | integer | 否 |  按Clone ID过滤 (格式: int64) |
| `endTime` | integer | 否 |  结束时间戳 (格式: int64) |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `startTime` | integer | 否 |  开始时间戳 (格式: int64) |
| `status` | integer | 否 |  状态过滤 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码过滤 |
| `userID` | integer | 否 |  按用户ID过滤 (格式: int64) |

**请求示例**:
```json
{
  "agentID": 1,
  "cloneID": 1,
  "endTime": 1,
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "startTime": 1,
  "status": 1,
  "tenantCode": "string",
  "userID": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "agentID": 1,
        "channel": "string",
        "cloneID": 1,
        "createdTime": 1,
        "id": 1,
        "lastActiveAt": 1,
        "sessionID": "string",
        "status": 1,
        "tenantCode": "string",
        "transcriptOssKey": "string",
        "turnCount": 1,
        "userID": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/get-list \
  --body '{"agentID": 1, "cloneID": 1, "endTime": 1, "page": {"page": 1, "pageSize": 1}, "startTime": 1, "status": 1, "tenantCode": "string", "userID": 1}'
```

### POST `/api/v1/ai/session/get-one`

**说明**: 获取会话详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  会话ID (格式: int64) |
| `sessionID` | string | 否 |  会话唯一标识（与ID二选一） |

**请求示例**:
```json
{
  "id": 1,
  "sessionID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "agentID": 1,
    "channel": "string",
    "cloneID": 1,
    "createdTime": 1,
    "id": 1,
    "lastActiveAt": 1,
    "sessionID": "string",
    "status": 1,
    "tenantCode": "string",
    "transcriptOssKey": "string",
    "turnCount": 1,
    "userID": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/get-one \
  --body '{"id": 1, "sessionID": "string"}'
```

### POST `/api/v1/ai/session/inject`

**说明**: 注入消息到会话

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `message` | string | 是 |  注入消息内容 |
| `messageType` | string | 否 |  消息类型 |
| `sessionID` | string | 是 |  会话ID |

**请求示例**:
```json
{
  "message": "string",
  "messageType": "string",
  "sessionID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "message": "string",
    "success": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/inject \
  --body '{"message": "string", "messageType": "string", "sessionID": "string"}'
```

### POST `/api/v1/ai/session/messages`

**说明**: 获取会话消息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `sessionID` | string | 是 |  会话ID |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "sessionID": "string",
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
        "aiResponse": "string",
        "keyTopics": "string",
        "summary": "string",
        "turnIndex": 1,
        "userInput": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/session/messages \
  --body '{"page": {"page": 1, "pageSize": 1}, "sessionID": "string", "tenantCode": "string"}'
```

### POST `/api/v1/ai/skill/create`

**说明**: 创建技能

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  技能编码 |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  技能ID (格式: int64) |
| `manifest` | string | 否 |  skill.yaml内容（JSON字符串） |
| `name` | string | 是 |  技能名称 |
| `ossPath` | string | 否 |  OSS存储路径 |
| `parentID` | integer | 否 |  父技能ID（0=顶级） (格式: int64) |
| `publishStatus` | string | 否 |  发布状态：draft/uploaded/validated/review_done/test_passed/approved/published/deprecated |
| `reviewResult` | string | 否 |  AI审阅结果（JSON） |
| `scope` | string | 否 |  范围：tenant=租户级 platform=平台级 |
| `sort` | integer | 否 |  排序 (格式: int64) |
| `status` | integer | 否 |  状态 1-启用 2-禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码，"common"=平台级 |
| `testResult` | string | 否 |  测试结果（JSON） |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |
| `version` | string | 否 |  版本号 |

**请求示例**:
```json
{
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
ur api /api/v1/ai/skill/create \
  --body '{"code": "string", "createdTime": 1, "desc": "string", "id": 1, "manifest": "string", "name": "示例名称", "ossPath": "string", "parentID": 1, "publishStatus": "string", "reviewResult": "string", "scope": "string", "sort": 1, "status": 1, "tenantCode": "string", "testResult": "string", "updatedTime": 1, "version": "string"}'
```

### POST `/api/v1/ai/skill/delete`

**说明**: 删除技能

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/skill/download`

**说明**: 下载技能zip包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
    "ossPath": "string",
    "zipData": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/download \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/skill/get-list`

**说明**: 获取技能列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | integer | 否 |  父技能ID过滤 (格式: int64) |
| `status` | integer | 否 |  状态过滤 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码过滤 |

**请求示例**:
```json
{
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "parentID": 1,
  "status": 1,
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
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/get-list \
  --body '{"name": "示例名称", "page": {"page": 1, "pageSize": 1}, "parentID": 1, "status": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/skill/get-one`

**说明**: 获取技能详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
ur api /api/v1/ai/skill/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/skill/publish`

**说明**: 发布新版本

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `skillID` | integer | 是 |  技能ID (格式: int64) |
| `version` | string | 是 |  新版本号 |

**请求示例**:
```json
{
  "skillID": 1,
  "version": "string"
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
ur api /api/v1/ai/skill/publish \
  --body '{"skillID": 1, "version": "string"}'
```

### POST `/api/v1/ai/skill/update`

**说明**: 更新技能

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  技能编码 |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  技能ID (格式: int64) |
| `manifest` | string | 否 |  skill.yaml内容（JSON字符串） |
| `name` | string | 是 |  技能名称 |
| `ossPath` | string | 否 |  OSS存储路径 |
| `parentID` | integer | 否 |  父技能ID（0=顶级） (格式: int64) |
| `publishStatus` | string | 否 |  发布状态：draft/uploaded/validated/review_done/test_passed/approved/published/deprecated |
| `reviewResult` | string | 否 |  AI审阅结果（JSON） |
| `scope` | string | 否 |  范围：tenant=租户级 platform=平台级 |
| `sort` | integer | 否 |  排序 (格式: int64) |
| `status` | integer | 否 |  状态 1-启用 2-禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码，"common"=平台级 |
| `testResult` | string | 否 |  测试结果（JSON） |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |
| `version` | string | 否 |  版本号 |

**请求示例**:
```json
{
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
ur api /api/v1/ai/skill/update \
  --body '{"code": "string", "createdTime": 1, "desc": "string", "id": 1, "manifest": "string", "name": "示例名称", "ossPath": "string", "parentID": 1, "publishStatus": "string", "reviewResult": "string", "scope": "string", "sort": 1, "status": 1, "tenantCode": "string", "testResult": "string", "updatedTime": 1, "version": "string"}'
```

### POST `/api/v1/ai/skill/upload`

**说明**: 上传技能zip包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  技能编码 |
| `name` | string | 是 |  技能名称 |
| `parentID` | integer | 否 |  父技能ID (格式: int64) |
| `tenantCode` | string | 否 |  租户编码 |
| `zipData` | string | 是 |  Base64编码的zip包数据 |

**请求示例**:
```json
{
  "code": "string",
  "name": "示例名称",
  "parentID": 1,
  "tenantCode": "string",
  "zipData": "string"
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
ur api /api/v1/ai/skill/upload \
  --body '{"code": "string", "name": "示例名称", "parentID": 1, "tenantCode": "string", "zipData": "string"}'
```

### POST `/api/v1/ai/skill/file/create`

**说明**: 创建技能文件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 |  文件内容 |
| `isDir` | boolean | 是 |  是否为目录 (格式: boolean) |
| `path` | string | 是 |  文件路径 |
| `skillID` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "content": "string",
  "isDir": true,
  "path": "string",
  "skillID": 1
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
ur api /api/v1/ai/skill/file/create \
  --body '{"content": "string", "isDir": true, "path": "string", "skillID": 1}'
```

### POST `/api/v1/ai/skill/file/delete`

**说明**: 删除文件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `path` | string | 是 |  文件路径 |
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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/file/delete \
  --body '{"path": "string", "skillID": 1}'
```

### POST `/api/v1/ai/skill/file/get-one`

**说明**: 获取文件内容

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `path` | string | 是 |  文件路径 |
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
    "content": "string",
    "path": "string",
    "size": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/file/get-one \
  --body '{"path": "string", "skillID": 1}'
```

### POST `/api/v1/ai/skill/file/update`

**说明**: 更新文件内容

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 |  文件内容 |
| `path` | string | 是 |  文件路径 |
| `skillID` | integer | 是 |  技能ID (格式: int64) |

**请求示例**:
```json
{
  "content": "string",
  "path": "string",
  "skillID": 1
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
ur api /api/v1/ai/skill/file/update \
  --body '{"content": "string", "path": "string", "skillID": 1}'
```

### POST `/api/v1/ai/tool/create`

**说明**: 创建工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentGroupID` | string | 是 |  关联AgentGroup ID |
| `code` | string | 是 |  工具编码 |
| `description` | string | 否 |  工具描述 |
| `name` | string | 是 |  工具名称 |

**请求示例**:
```json
{
  "agentGroupID": "string",
  "code": "string",
  "description": "string",
  "name": "示例名称"
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
ur api /api/v1/ai/tool/create \
  --body '{"agentGroupID": "string", "code": "string", "description": "string", "name": "示例名称"}'
```

### POST `/api/v1/ai/tool/delete`

**说明**: 删除工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/tool/disable`

**说明**: 停用工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/disable \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/tool/get-list`

**说明**: 获取工具列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agentGroupID` | string | 否 |  AgentGroup过滤 |
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态过滤 |

**请求示例**:
```json
{
  "agentGroupID": "string",
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "status": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "agentGroupID": "string",
        "code": "string",
        "createdTime": 1,
        "description": "string",
        "groupName": "示例名称",
        "id": "string",
        "name": "示例名称",
        "status": "string",
        "tenantCode": "string",
        "version": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/get-list \
  --body '{"agentGroupID": "string", "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "status": "string"}'
```

### POST `/api/v1/ai/tool/get-one`

**说明**: 获取工具详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  工具ID |

**请求示例**:
```json
{
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "artifact": {
      "documentMd": "string",
      "executorJs": "string",
      "manifestJson": "string"
    },
    "info": {
      "agentGroupID": "string",
      "code": "string",
      "createdTime": 1,
      "description": "string",
      "groupName": "示例名称",
      "id": "string",
      "name": "示例名称",
      "status": "string",
      "tenantCode": "string",
      "version": "string"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/ai/tool/publish`

**说明**: 发布工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/publish \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/tool/save-artifact`

**说明**: 保存三件套

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `documentMd` | string | 否 |  document.md 内容 |
| `executorJs` | string | 是 |  executor.js 内容 |
| `id` | string | 是 |  工具ID |
| `manifestJson` | string | 否 |  manifest.json 内容 |

**请求示例**:
```json
{
  "documentMd": "string",
  "executorJs": "string",
  "id": "string",
  "manifestJson": "string"
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
ur api /api/v1/ai/tool/save-artifact \
  --body '{"documentMd": "string", "executorJs": "string", "id": "string", "manifestJson": "string"}'
```

### POST `/api/v1/ai/tool/update`

**说明**: 更新工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `description` | string | 否 |  工具描述 |
| `id` | string | 是 |  工具ID |
| `name` | string | 否 |  工具名称 |

**请求示例**:
```json
{
  "description": "string",
  "id": "string",
  "name": "示例名称"
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
ur api /api/v1/ai/tool/update \
  --body '{"description": "string", "id": "string", "name": "示例名称"}'
```

### POST `/api/v1/ai/tool/run`

**说明**: 手动运行工具

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `inputs` | string | 是 |  输入参数（JSON字符串） |
| `toolID` | string | 是 |  工具ID |

**请求示例**:
```json
{
  "inputs": "string",
  "toolID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "runID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/run \
  --body '{"inputs": "string", "toolID": "string"}'
```

### POST `/api/v1/ai/tool/run-history`

**说明**: 执行历史列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `toolID` | string | 是 |  工具ID |

**请求示例**:
```json
{
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "toolID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "durationMs": "string",
        "errorMsg": "string",
        "logs": "string",
        "runID": "string",
        "state": "string",
        "status": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/run-history \
  --body '{"page": {"page": 1, "pageSize": 1}, "toolID": "string"}'
```

### POST `/api/v1/ai/tool/run-status`

**说明**: 查询执行状态

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `runID` | string | 是 |  运行标识UUID |

**请求示例**:
```json
{
  "runID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "durationMs": "string",
    "errorMsg": "string",
    "logs": "string",
    "runID": "string",
    "state": "string",
    "status": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/tool/run-status \
  --body '{"runID": "string"}'
```

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

### POST `/api/v1/things/alarm/info/create`

**说明**: 新增告警

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accounts` | array[string] | 否 | 账号 |
| `code` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `level` | integer | 否 | 告警配置级别（1提醒 2一般 3严重 4紧急 5超紧急） (格式: int64) |
| `name` | string | 否 |  |
| `notifies` | array[AlarmNotify] | 否 | 通知 |
| `sceneIDs` | array[integer] | 否 | 绑定的场景列表 |
| `scenes` | array[SceneCore] | 否 | 绑定的场景列表,只读 |
| `status` | integer | 否 | 状态: 1启用 2禁用 (格式: int64) |
| `userIDs` | array[string] | 否 | 指定用户ID |

**请求示例**:
```json
{
  "accounts": [
    "string"
  ],
  "code": "string",
  "createdTime": "string",
  "desc": "string",
  "id": "string",
  "level": 1,
  "name": "string",
  "notifies": [
    {
      "templateID": "string",
      "templateName": "示例名称",
      "type": "string"
    }
  ],
  "sceneIDs": [
    1
  ],
  "scenes": [
    {
      "desc": "string",
      "id": "string",
      "name": "string",
      "status": 1
    }
  ],
  "status": 1,
  "userIDs": [
    "string"
  ]
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
ur api /api/v1/things/alarm/info/create \
  --body '{"accounts": ["string"], "code": "string", "createdTime": "string", "desc": "string", "id": "string", "level": 1, "name": "string", "notifies": [{"templateID": "string", "templateName": "示例名称", "type": "string"}], "sceneIDs": [1], "scenes": [{"desc": "string", "id": "string", "name": "string", "status": 1}], "status": 1, "userIDs": ["string"]}'
```

### POST `/api/v1/things/alarm/info/delete`

**说明**: 删除告警

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/alarm/info/get-list`

**说明**: 获取告警信息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `name` | string | 否 | 告警名模糊查询 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "name": "string",
  "page": {
    "page": 1,
    "pageSize": 1
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
        "accounts": [
          "string"
        ],
        "code": "string",
        "createdTime": "string",
        "desc": "string",
        "id": "string",
        "level": 1,
        "name": "string",
        "notifies": [
          {
            "templateID": "string",
            "templateName": "示例名称",
            "type": "string"
          }
        ],
        "sceneIDs": [
          1
        ],
        "scenes": [
          {
            "desc": "string",
            "id": "string",
            "name": "string",
            "status": 1
          }
        ],
        "status": 1,
        "userIDs": [
          "string"
        ]
      }
    ],
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/info/get-list \
  --body '{"code": "string", "name": "string", "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/things/alarm/info/get-one`

**说明**: 获取告警信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

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
    "accounts": [
      "string"
    ],
    "code": "string",
    "createdTime": "string",
    "desc": "string",
    "id": "string",
    "level": 1,
    "name": "string",
    "notifies": [
      {
        "templateID": "string",
        "templateName": "示例名称",
        "type": "string"
      }
    ],
    "sceneIDs": [
      1
    ],
    "scenes": [
      {
        "desc": "string",
        "id": "string",
        "name": "string",
        "status": 1
      }
    ],
    "status": 1,
    "userIDs": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/things/alarm/info/update`

**说明**: 更新告警

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accounts` | array[string] | 否 | 账号 |
| `code` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `level` | integer | 否 | 告警配置级别（1提醒 2一般 3严重 4紧急 5超紧急） (格式: int64) |
| `name` | string | 否 |  |
| `notifies` | array[AlarmNotify] | 否 | 通知 |
| `sceneIDs` | array[integer] | 否 | 绑定的场景列表 |
| `scenes` | array[SceneCore] | 否 | 绑定的场景列表,只读 |
| `status` | integer | 否 | 状态: 1启用 2禁用 (格式: int64) |
| `userIDs` | array[string] | 否 | 指定用户ID |

**请求示例**:
```json
{
  "accounts": [
    "string"
  ],
  "code": "string",
  "createdTime": "string",
  "desc": "string",
  "id": "string",
  "level": 1,
  "name": "string",
  "notifies": [
    {
      "templateID": "string",
      "templateName": "示例名称",
      "type": "string"
    }
  ],
  "sceneIDs": [
    1
  ],
  "scenes": [
    {
      "desc": "string",
      "id": "string",
      "name": "string",
      "status": 1
    }
  ],
  "status": 1,
  "userIDs": [
    "string"
  ]
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
ur api /api/v1/things/alarm/info/update \
  --body '{"accounts": ["string"], "code": "string", "createdTime": "string", "desc": "string", "id": "string", "level": 1, "name": "string", "notifies": [{"templateID": "string", "templateName": "示例名称", "type": "string"}], "sceneIDs": [1], "scenes": [{"desc": "string", "id": "string", "name": "string", "status": 1}], "status": 1, "userIDs": ["string"]}'
```

### POST `/api/v1/things/alarm/record/deal`

**说明**: 处理告警

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `handle` | integer | 是 | 告警处理结果 1:无效(告警无效) 2:派单(走工作流) (格式: int64) |
| `id` | string | 是 | 告警记录ID |

**请求示例**:
```json
{
  "handle": 1,
  "id": "string"
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
ur api /api/v1/things/alarm/record/deal \
  --body '{"handle": 1, "id": "string"}'
```

### POST `/api/v1/things/alarm/record/get-list`

**说明**: 获取告警记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmCode` | string | 否 |  |
| `alarmID` | string | 否 | 告警记录ID |
| `alarmName` | string | 否 |  |
| `areaID` | string | 否 | 区域ID过滤 |
| `areaIDPath` | string | 否 |  |
| `dealStatus` | array[integer] | 否 | 告警记录状态（1告警中 2忽略 3正在处理 4已处理） |
| `deviceAlias` | string | 否 | 触发设备名称 |
| `deviceName` | string | 否 | 触发设备ID |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `timeRange` | object | 否 |  |
| `timeRange.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `timeRange.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |

**请求示例**:
```json
{
  "alarmCode": "string",
  "alarmID": "string",
  "alarmName": "string",
  "areaID": "string",
  "areaIDPath": "string",
  "dealStatus": [
    1
  ],
  "deviceAlias": "示例名称",
  "deviceName": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "timeRange": {
    "end": 1,
    "start": 1
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
        "alarmCount": 1,
        "alarmID": "string",
        "alarmName": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "dealStatus": 1,
        "desc": "string",
        "deviceAlias": "示例名称",
        "deviceName": "string",
        "id": "string",
        "lastAlarm": 1,
        "level": 1,
        "productID": "string",
        "sceneID": "string",
        "sceneName": "示例名称",
        "triggerType": "string"
      }
    ],
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/record/get-list \
  --body '{"alarmCode": "string", "alarmID": "string", "alarmName": "string", "areaID": "string", "areaIDPath": "string", "dealStatus": [1], "deviceAlias": "示例名称", "deviceName": "string", "page": {"page": 1, "pageSize": 1}, "timeRange": {"end": 1, "start": 1}}'
```

### POST `/api/v1/things/alarm/scene/batch-create`

**说明**: 更新告警和场景的关联

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 是 | 告警配置ID |
| `sceneIDs` | array[integer] | 是 | 场景id |

**请求示例**:
```json
{
  "alarmID": "string",
  "sceneIDs": [
    1
  ]
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
ur api /api/v1/things/alarm/scene/batch-create \
  --body '{"alarmID": "string", "sceneIDs": [1]}'
```

### POST `/api/v1/things/alarm/scene/delete`

**说明**: 删除告警和场景的关联

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 是 | 告警配置ID |
| `sceneID` | string | 是 | 场景id |

**请求示例**:
```json
{
  "alarmID": "string",
  "sceneID": "string"
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
ur api /api/v1/things/alarm/scene/delete \
  --body '{"alarmID": "string", "sceneID": "string"}'
```

### POST `/api/v1/things/alarm/scene/get-list`

**说明**: 获取告警和场景的关联列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmID` | string | 是 | 告警配置ID |

**请求示例**:
```json
{
  "alarmID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "alarmID": "string",
    "scenes": [
      {
        "areaID": "string",
        "avatar": "string",
        "body": "string",
        "createdTime": "string",
        "desc": "string",
        "deviceAlias": "string",
        "deviceMode": "string",
        "deviceName": "string",
        "flowPath": [
          {
            "info": "string",
            "subType": "string",
            "type": "string"
          }
        ],
        "id": "string",
        "if": "string",
        "isCommon": 1,
        "lastRunTime": 1,
        "logo": "string",
        "name": "string",
        "productID": "string",
        "reason": "string",
        "status": 1,
        "tag": "string",
        "then": "string",
        "type": "string",
        "updatedTime": 1,
        "when": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/scene/get-list \
  --body '{"alarmID": "string"}'
```
