# screenshot 截图细节

> 本文件是 [SKILL.md](../SKILL.md) 的子文档：screenshot 的 URL 模板、`--url-template` 覆盖、登录态与排障细节（自 SKILL.md 拆分整理，内容未改动）。

## 截图细节（发布态 / 编辑态）

- **screenshot 依赖本机 agent-browser**：依次执行 `open → wait（毫秒）→ screenshot`；`--wait` 单位是秒。前端地址取 `--front-base` 或环境变量 `UR_FRONT_BASE_URL`。实测：发布页默认模板 `{base}/app/iot/#/view/{id}`（`/view/` 兼容路由始终取已发布快照；`/iot/big-bigscreen/published/{id}` 在 iot 应用内会落入应用外壳不渲染画布，勿用），编辑态默认 `{base}/app/iot/#/iot/big-bigscreen/editor/{id}`；部署形态不同可用 `--url-template` 覆盖（支持 `{base}` / `{id}` 占位符）。发布页需登录态，建议先用 agent-browser 登录 iot 应用（租户管理员 + 选择企业）保持会话。
