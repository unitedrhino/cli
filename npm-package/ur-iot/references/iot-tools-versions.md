# IoT Tools 版本说明

## v1.0-mcp-compatible

### 目标

- 保持现有 MCP 路径不变
- Web 前端允许开始接入 frontend tools
- 设备侧继续保留 MCP

### 规则

- Web 前端可新增 frontend tools，但不能破坏既有 MCP 工具行为
- Win AI、MQTT、UDP 视为 MCP 优先客户端
- 工具名和结果语义尽量与 MCP 保持一致

## v1.1-all-device-visible

### 目标

- 在 Web 前端支持所有可访问设备可见、可搜索
- 明确前端不负责切换当前设备，也不负责绑定设备上下文

### 规则

- AI 在未拿到明确目标设备时，应先查询设备列表、让用户补充设备信息，或依赖后端上下文能力
- 设备列表态、会话态、执行态必须可解释且一致
- 与 MCP 共存期间，Web 不要求承担 `_sessionID` 替代者职责
- 目标设备解析需同时支持两类意图：单设备控制、区域 + 类型批量控制
