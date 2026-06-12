# ai/chat

> 该 group 共 7 个端点。

- `POST /api/v1/ai/chat/available-tools` [all] 获取当前会话可用的后端工具列表
- `POST /api/v1/ai/chat/check-available` [all] 检测 AI 对话功能是否可用
- `POST /api/v1/ai/chat/completions` [all] 文字对话（指定 agentID）
- `POST /api/v1/ai/chat/tool-result` [all] 前端 Tool 执行结果回传
- `POST /api/v1/ai/embedding/debug` [platform] Embedding配置调试
- `POST /api/v1/ai/llm/debug` [platform] LLM配置调试
- `POST /api/v1/ai/user/completions` [all] 用户 AI 对话（自动查找/创建专属 Clone）
