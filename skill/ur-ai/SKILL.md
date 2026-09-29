---
name: ur-ai
description: "AI Agent管理、告警规则、规则引擎、场景联动。triggers: AI管理, AI Agent, 场景联动, 告警规则, 告警记录, 数字分身, AI会话, 规则引擎, AgentGroup, Clone"
metadata:
  hermes:
    tags: [ai, agent, rule, alarm, automation]
---


# ur-ai — AI 管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/generate-api-lists.py` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 |
|------|------|
| **AgentGroup** | 平台级分组，配置继承，purpose用途分类（device/user/platform） |
| **Agent** | 智能体配置，覆盖Group配置，声明capabilities（text/voice/image） |
| **Clone** | 数字分身，独立记忆，绑定设备或会话，persona_prompt人设补充 |
| **Session** | 会话，短期记忆（滑动窗口），通道类型（text/voice/multimodal） |
| **配置继承** | 有效配置 = Group配置 ∪ Agent配置（Agent优先） |
| **双层记忆** | 层一：Redis向量数据库；层二：PostgreSQL摘要持久化 |
| **MCP安全模型** | platform scope可访问内网，tenant scope禁止访问内网 |
| **UDP音频通道** | sessionCreated返回udpHost/udpPort/aesKey/nonce，音频AES-CTR加密传输 |
| **Skill** | Agent能力定义，MCP工具、API调用能力 |
| **告警** | 设备告警规则，条件触发→动作执行 |
| **场景** | 自动化场景，触发器+条件+动作 |
| **Agent-设备绑定** | 产品defaultAgentID→设备绑定Agent→自动创建Clone |

## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| AI 对话 | 裸 LLM 调用（agentID=0，无上下文注入） | 本文"裸 LLM 调用"章节 |
| 告警规则 | 告警规则查询/创建/更新/删除/启停/手动评估 | [alarm-rule.md](references/alarm-rule.md) |
| 告警事件 | 告警事件查询/详情/处理/误报/统计 | [alarm-event.md](references/alarm-event.md) |
| 通知记录 | 告警通知记录查询/重发 | [alarm-notify-record.md](references/alarm-notify-record.md) |
| 场景联动 | 场景查询/创建/更新/删除/触发 | [scene.md](references/scene.md) |
| 场景日志 | 场景执行日志查询 | [scene.md](references/scene.md) |
| 场景模板 | 场景JSON校验/模板生成 | [scene.md](references/scene.md) |

> 完整命令帮助：`ur alarm help`、`ur scene help`

## 典型业务场景

### AI Agent 与设备绑定

**场景描述**：为设备启用AI能力 / 配置智能助手（产品defaultAgentID → 设备绑定 → 自动创建Clone）

**涉及 CLI**：
- `ur ai agent get-one`
- `ur ai clone create`
- `ur ai clone get-list`

**工作流**：
1. 产品设置defaultAgentID
2. 设备绑定Agent
3. 系统自动创建Clone（一设备一Clone）
4. 设备通过MQTT或API与AI交互

### 裸 LLM 调用（agentID=0）

**场景描述**：urcli / 外部代码调用 AI，不需要 Agent 上下文注入（无提示词、无 Skill、无知识库、无记忆、无会话记录）

**涉及 API**：`POST /api/v1/ai/chat/completions`（agentID=0 或不传）

**行为特征**：
- 不注入 Agent/Group 系统提示词
- 不加载 Skill 索引
- 不检索知识库
- 不加载 Clone 人格/记忆
- 不加载会话历史
- 不注入页面上下文
- 不记录会话摘要、不触发记忆提取
- **参与企业 token 计费**（与现有接口走同一套计费链路）

**请求参数**：
```json
{
  "agentID": 0,
  "messages": [{"role": "user", "contents": [{"type": "text", "text": "你好"}]}],
  "modelType": "small",
  "stream": false
}
```

| 参数 | 必填 | 说明 |
|------|------|------|
| agentID | 可选 | 0 或不传 = 裸 LLM 调用 |
| messages | 必填 | 对话消息列表 |
| modelType | 可选 | LLM 模型类型，默认 "small" |
| stream | 可选 | 是否流式返回（SSE），默认 false |

**CLI 用法**：
```bash
ur ai chat -m "帮我写一个 Go 的 hello world"
ur ai chat -m "解释 IoT 协议" --model-type large
ur ai chat -m "写一首诗" --stream
```

### AI MQTT 交互协议

**场景描述**：设备与AI实时对话 / 语音交互（MQTT: $thing/up/ai, $thing/down/ai）

**涉及 CLI**：
- `ur ai mcp message`
- `ur ai mcp sse`

**工作流**：
1. sessionCreate：创建会话，返回sessionID + UDP参数
2. inputSend：发送文本/音频输入
3. respTextDone/respAudio：接收响应
4. UDP通道：加密音频流传输

### AgentGroup 配置继承

**场景描述**：创建 AgentGroup 定义共享配置，Agent 继承并覆盖

**涉及 CLI**：
- `ur ai agent-group create`
- `ur ai agent create`

**工作流**：
1. 创建 AgentGroup（llm_config_id, mcp_service_ids, skill_ids）
2. 创建 Agent 并关联 group_id
3. Agent 配置覆盖 Group 配置
4. 最终配置 = Group ∪ Agent（Agent 优先）

### 告警管理

**场景描述**：创建告警规则 / 处理告警 / 通知推送

**涉及 CLI**：
- `ur alarm rule get-list` / `ur alarm rule get-one` / `ur alarm rule create` / `ur alarm rule update` / `ur alarm rule status-update` / `ur alarm rule evaluate-trigger`
- `ur alarm event get-list` / `ur alarm event get-one` / `ur alarm event deal` / `ur alarm event false-alarm` / `ur alarm event stat`
- `ur alarm notify-record get-list` / `ur alarm notify-record resend`

**工作流**：
1. 定义告警条件（属性阈值、设备状态）
2. 配置告警动作（通知、场景联动）
3. 触发告警后处理

### 场景联动

**场景描述**：创建自动化场景 / 触发场景（触发器+条件+动作）

**涉及 CLI**：
- `ur scene info create`
- `ur scene info get-list`
- `ur scene info get-one`
- `ur scene info update`
- `ur scene info delete`
- `ur scene info trigger`

**工作流**：
1. 定义触发器（定时、属性变化、事件）
2. 设置条件判断
3. 配置执行动作（设备控制、通知）


## 注意事项

- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- CLI 命令自动处理请求头（`app-id`、`tenant-code` 和认证头）
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
