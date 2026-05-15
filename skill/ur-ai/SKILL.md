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
> 如需更新，修改 `scripts/lib/swagger.ts` 中的域名定义后重新生成。

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

## API 参考

<!-- API_LIST:ur-ai -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
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

<!-- END_API_LIST -->

权限: 管理员

| 端点 | 说明 | 权限 |
|---|------|------|
| POST /api/v1/things/ai/mcp/message | MCP消息发送 | 所有用户 |
| POST /api/v1/things/ai/mcp/sse | SSE连接（POST） | 所有用户 |
| POST /api/v1/things/alarm/info/create | 新增告警 | 管理员 |
| POST /api/v1/things/alarm/info/delete | 删除告警 | 管理员 |
| POST /api/v1/things/alarm/info/get-list | 获取告警信息列表 | 管理员 |
| POST /api/v1/things/alarm/info/get-one | 获取告警信息 | 管理员 |
| POST /api/v1/things/alarm/info/update | 更新告警 | 管理员 |
| POST /api/v1/things/alarm/record/deal | 处理告警 | 管理员 |
| POST /api/v1/things/alarm/record/get-list | 获取告警记录列表 | 管理员 |
| POST /api/v1/things/alarm/scene/batch-create | 更新告警和场景的关联 | 管理员 |
| POST /api/v1/things/alarm/scene/delete | 删除告警和场景的关联 | 管理员 |
| POST /api/v1/things/alarm/scene/get-list | 获取告警和场景的关联列表 | 管理员 |
| POST /api/v1/things/rule/alarm/info/create | 新增告警(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/info/delete | 删除告警(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/info/get-list | 获取告警信息列表(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/info/get-one | 获取告警信息(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/info/update | 更新告警(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/record/deal | 处理告警(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/record/get-list | 获取告警记录列表(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/scene/batch-create | 更新告警和场景的关联(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/scene/delete | 删除告警和场景的关联(弃用) | 管理员 |
| POST /api/v1/things/rule/alarm/scene/get-list | 获取告警和场景的关联列表(弃用) | 管理员 |
| POST /api/v1/things/rule/scene/info/create | 新增场景(弃用) | 管理员 |
| POST /api/v1/things/rule/scene/info/delete | 删除场景(弃用) | 管理员 |
| POST /api/v1/things/rule/scene/info/get-list | 获取场景信息列表(弃用) | 管理员 |
| POST /api/v1/things/rule/scene/info/get-one | 获取场景信息详情(弃用) | 管理员 |
| POST /api/v1/things/rule/scene/info/manually-trigger | 手动触发场景(弃用) | 管理员 |
| POST /api/v1/things/rule/scene/info/update | 更新场景(弃用) | 管理员 |
| POST /api/v1/things/rule/scene/log/get-list | 获取场景日志列表(弃用) | 管理员 |
| POST /api/v1/things/scene/info/create | 新增场景 | 管理员 |
| POST /api/v1/things/scene/info/delete | 删除场景 | 管理员 |
| POST /api/v1/things/scene/info/get-list | 获取场景信息列表 | 管理员 |
| POST /api/v1/things/scene/info/get-one | 获取场景信息详情 | 管理员 |
| POST /api/v1/things/scene/info/manually-trigger | 手动触发场景 | 管理员 |
| POST /api/v1/things/scene/info/update | 更新场景 | 管理员 |
| POST /api/v1/things/scene/log/get-list | 获取场景日志列表 | 管理员 |


## 典型业务场景

### AI Agent 与设备绑定

**场景描述**：为设备启用AI能力 / 配置智能助手（产品defaultAgentID → 设备绑定 → 自动创建Clone）

**涉及 API**：
- `/api/v1/things/ai/agent/get-one`
- `/api/v1/things/ai/clone/create`
- `/api/v1/things/ai/clone/get-list`

**工作流**：
1. 产品设置defaultAgentID
2. 设备绑定Agent
3. 系统自动创建Clone（一设备一Clone）
4. 设备通过MQTT或API与AI交互

### AI MQTT 交互协议

**场景描述**：设备与AI实时对话 / 语音交互（MQTT: $thing/up/ai, $thing/down/ai）

**涉及 API**：
- `/api/v1/things/ai/mcp/message`
- `/api/v1/things/ai/mcp/sse`

**工作流**：
1. sessionCreate：创建会话，返回sessionID + UDP参数
2. inputSend：发送文本/音频输入
3. respTextDone/respAudio：接收响应
4. UDP通道：加密音频流传输

### AgentGroup 配置继承

**场景描述**：创建 AgentGroup 定义共享配置，Agent 继承并覆盖

**涉及 API**：
- `/api/v1/things/ai/agent-group/create`
- `/api/v1/things/ai/agent/create`

**工作流**：
1. 创建 AgentGroup（llm_config_id, mcp_service_ids, skill_ids）
2. 创建 Agent 并关联 group_id
3. Agent 配置覆盖 Group 配置
4. 最终配置 = Group ∪ Agent（Agent 优先）

### 告警管理

**场景描述**：创建告警规则 / 处理告警 / 通知推送

**涉及 API**：
- `/api/v1/things/alarm/info/create`
- `/api/v1/things/alarm/record/deal`

**工作流**：
1. 定义告警条件（属性阈值、设备状态）
2. 配置告警动作（通知、场景联动）
3. 触发告警后处理

### 场景联动

**场景描述**：创建自动化场景 / 触发场景（触发器+条件+动作）

**涉及 API**：
- `/api/v1/things/scene/info/create`
- `/api/v1/things/scene/info/manually-trigger`

**工作流**：
1. 定义触发器（定时、属性变化、事件）
2. 设置条件判断
3. 配置执行动作（设备控制、通知）


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
