---
name: ur-api
description: "Use when calling 联犀 SaaS 平台 API: device management, user management, product management, tenant management, AI management, project management, system management, bigscreen visualization, client debugging, or IoT device operations. triggers: API调用, CLI登录, CLI升级, _notice, Sandbox认证, 账号密码登录, AKSK, 设备列表, IoT设备控制, 设备注册, 访问令牌, 物模型管理, OTA升级, 场景联动, 权限配置, swagger schema, 项目管理, 区域管理, 系统管理, 客户端调试, 控制台报错, 远程诊断, 问题反馈, 提交反馈, 反馈类型, 使用问题, 业务受损, 业务不可用, 大屏, 数据可视化, GoView, 画布, 看板, 编辑大屏, 发布大屏, 文档解析, 解析PDF, 解析Excel, 读附件, ur doc"
---

# ur-api — 联犀 SaaS 平台 API 工具

CLI 主实现位于独立仓库 `unitedrhino/cli`，通过 Go CLI `ur` 提供能力。Sandbox 可只用 `UR_*` 环境变量；本地已配置的旧 profile 继续自动兼容，无需迁移或重新登录。

所有接口均为 POST 方法。请求格式 `{code, msg, data}`。

---

## 认证前置流程（AI 必读）

1. 先运行 `ur check --json`。如果环境变量或历史 profile 已可用，直接调用 API；不要删除配置、重新执行 `setup` 或要求用户再次提供凭据。
2. 只有在没有可用认证时，才运行 `ur login --method device|password|aksk`；未指定 `--method` 时默认 `device`。
3. Sandbox 设置 `UR_BASE_URL` 后采用 env-only 模式：只使用同一环境中的完整凭据组，不读取磁盘 profile 补齐缺失值，也不把环境凭据写入磁盘。

支持的运行时凭据组按 `UR_TOKEN` → `UR_ACCESS_KEY` + `UR_ACCESS_SECRET` → `UR_ACCOUNT` + `UR_PASSWORD` 的顺序选择。AK/SK 不要求 `UR_USER_ID`；成对变量缺一项时应修正 Sandbox 注入，不要从 profile 混合补值。

账号密码入口接收**原始密码**，CLI 会做且只做一次 SHA-256；禁止先摘要再传入，否则会二次摘要。直接调用 `/api/v1/system/user/self/login` 时，调用方才需要发送 SHA-256 摘要并使用 `pwdType: 1`；MD5/`pwdType: 2` 仅为历史兼容。

敏感值优先通过 Sandbox 环境变量或 stdin 注入。`--password`、`--access-secret` 明文参数仅用于兼容，可能进入 shell 历史或进程列表；不得把密码、AccessSecret 或完整 Token 写入 skill、日志、截图和版本库。

业务命令的 JSON 结果可能包含 `_notice.update` 或 `_notice.skills`：先完成用户当前请求，再按 `command` 简短提示或执行用户已授权的升级（统一用 `ur upgrade`）；详细规则见 `references/cli-usage.md` 的“CLI 更新提示与处理”。

## 生成文件交付（AI 产文件给用户下载）

用户要求生成文件（Excel/报表/图片等）并下载时，直接在沙箱内完成，不要把步骤留给用户：
1. 用沙箱内工具生成文件（如 Python openpyxl 写 xlsx；文件名建议用英文）
2. 执行 `ur system upload-file --file <文件路径>` 上传到临时桶（凭据已注入环境变量，无需登录）
3. 回复中写 markdown 链接 `[文件名](/oss/temporary/<filePath>)` 交付——`<filePath>` 取命令输出 `filePath:` 行的对象路径；页面会渲染为可下载的文件卡片
4. 只允许 `/oss/` 开头的相对路径（临时桶前缀 `/oss/temporary/`、公开桶 `/oss/public/`）；不要使用 `http(s)://` 完整地址（如 COS 签名 URL，有时效且不渲染卡片），不要拼 `http://allinone:7777` 等沙箱内地址。需要把链接发到对话外直接打开时，改用 `--public --business <业务> --use-by user` 上传公开桶

## 快速开始

先运行 `ur check --json`；检查通过后直接调用 API，仅在返回缺少认证时选择以下一种登录方式：

- 方式 1（默认）Device Flow：`ur login --method device`；方式 2 账号密码：`ur login --method password`（密码必须是原始密码）；方式 3 AccessKey/JWT：`ur login --method aksk`（不要求 `userID`）
- 方式 4 Sandbox 环境变量（无需 login）：`UR_BASE_URL` + `UR_APP_ID` + `UR_TENANT_CODE` 及任一完整凭据组

四种方式的完整命令示例与历史 profile 兼容说明见 `references/auth.md`。

---

## 生成文件交付（AI 产文件给用户下载）

用户要求生成文件（Excel/报表/图片等）并下载时，直接在沙箱内完成，不要把步骤留给用户：
1. 用沙箱内工具生成文件（如 Python openpyxl 写 xlsx；文件名建议用英文）
2. 执行 `ur system upload-file --file <文件路径>` 上传到临时桶（凭据已注入环境变量，无需登录）
3. 从命令输出的 markdown 行**原样复制** `[文件名](/oss/temporary/....xlsx)` 写进回复——页面会渲染为可下载的文件卡片；fileUri 以命令输出为准，不要自行拼接或修改路径
4. fileUri 的 `/oss/` 相对路径即为正确形态（页面经登录态代理下载）；不要拼 `http://allinone:7777` 等沙箱内地址。需要把链接发到对话外直接打开时，改用 `--public --business <业务> --use-by user` 上传公开桶（返回永久 URL）

## 文档解析(ur doc)

用户上传或引用文档(PDF/Word/PPT/Excel/图片/邮件等)需要读取内容时,使用 `ur doc parse`:先 `--format outline` 看结构地图,再按需 `--format md` 通读/`--section` 取章节,或 `--format json` 落盘后用 jq 精查;扫描件/图片加 `--ocr`。完整用法见 `ur-doc/SKILL.md`。

## 客户端实时调试

排查已接入调试桥的客户端控制台报错、白屏或状态异常时，先读 [客户端调试 AI 流程](references/client-debug.md)；使用通用 `ur api --stream` 保持 SSE 会话，逐条核对执行结果，结束时关闭流。

---

## AI 快速决策

### 第一步：确认当前用户角色

调用任何 API 前，**先运行 check 确认角色**，错误的角色会导致 403。角色不足时直接告知用户切换，不要尝试后 403。角色表、接口范围（platform/admin/all）、切换方式（`ur login` 或用 `--app` 切换应用上下文）与接口侧身份判定（`/api/v1/system/user/self/get-one`、`UserTenant`/`RoleInfo` 字段与类型定义来源）见 `references/role-permissions.md`。

### 第二步：根据任务选择子域

| 用户意图 / 关键词 | 加载子域 | 最低所需角色 |
|----------------|---------|------------|
| 设备列表、设备控制、属性上报、MQTT、三元组、网关、OTA升级设备 | `ur-device` | admin（控制/CRUD），user（分享/收藏） |
| 产品定义、物模型、协议脚本、OTA固件包、品类 | `ur-product` | admin |
| 项目、区域、分组、数据权限申请 | `ur-project` | admin（管理），user（申请权限） |
| 登录、用户信息、角色、部门、字典、通知、访问令牌、修改密码 | `ur-user` | admin（CRUD），user（自身信息） |
| 文件上传、WebSocket、批量接口、应用管理、Hook | `ur-system` | user（上传/WS），admin（应用），platform（全局应用） |
| 客户端控制台报错、实时日志、AI 远程诊断 | [客户端调试 AI 流程](references/client-debug.md) | 目标用户本人或平台管理员 |
| 创建企业、企业列表、企业配置、邀请用户加入企业 | `ur-tenant` | **platform**（创建/查看全部），admin（管理本企业） |
| AI Agent、场景联动、告警规则、数字分身、AI会话 | `ur-ai` | admin |
| 大屏、数据可视化、GoView、画布编辑、发布大屏、看板、素材库 | `ur-view` | admin |

设备模拟、测试数据或属性控制先读 [设备控制与模拟](ur-iot/ur-device/references/device-control.md)，区分纯生成、云端改值、模拟上报和实体控制；不能用仅修改云端属性证明实体控制成功。

三个客户端专题（ur-iot-device / ur-iot-context / ur-iot-client）的适用场景表见 `references/app-selection.md`。

### 数量 / 列表类问题

问“有几个”“多少个”“当前有哪些”时直接调用真实的 `get-list` 接口，数量优先从返回的 `data.total` 读取；不要先调用 `schema` 再决定是否查询数据。完整规则见 `references/api-conventions.md`。

### API 路径选择与 404 回退策略（强制）

**绝对禁止臆造不存在的 API 路径**：不要拼接或猜测 `/api/v1/...` 路径，不要假设近义词路径；调用一次返回 `404 page not found` 后**必须立即停止猜测**，改为查阅本 SKILL 的「高频端点速查」或各子域 SKILL.md 的端点列表，仍不确定时调用 `knowledge_search` 检索知识库。完整回退策略见 `references/troubleshooting.md`。

---

## API 域划分

共 7 个子域，覆盖 492 个端点（100% 覆盖）；域划分表见 `references/app-selection.md`。

## 业务分组导航

按任务选择业务组，再按组内导航读取子模块；各层入口统一使用 `SKILL.md`。

| 业务组 | 内容 | 入口 |
|--------|------|------|
| 物联网 | 设备、产品、项目、固件、OTA、物模型、协议、联动、数据与调试 | [ur-iot](ur-iot/SKILL.md) |
| 组织管理 | 企业、用户、权限、系统与审批 | [ur-org-manage](ur-org-manage/SKILL.md) |
| AI | Agent、告警、设备智能与 AI 工具 | [ur-ai](ur-ai/SKILL.md) |
| 大屏 | GoView、数据绑定与场景模板 | [ur-view](ur-view/SKILL.md) |
| 文档解析 | PDF、Office、图片与 CAD 解析 | [ur-doc](ur-doc/SKILL.md) |

---

## 参考文档

| 文档 | 说明 |
|------|------|
| `references/auth.md` | 认证方式详解（Device Auth / JWT / 环境变量） |
| `references/cli-usage.md` | CLI 命令用法（更新通知、多客户端 Skills 分发、API 调用、物模型、场景联动、协议脚本、schema） |
| `references/client-debug.md` | 客户端实时调试的 SSE、日志过滤、操控确认与排障流程 |
| `references/api-conventions.md` | API 通用约定（请求格式、响应格式、分页、权限标注） |
| `references/troubleshooting.md` | 常见问题排查（登录失败、401/403、连接问题、问题反馈） |
| `references/quick-reference.md` | 高频端点速查（Top 15 最常用端点） |
| `references/workflows.md` | 跨域工作流示例（设备上线、企业初始化、告警通知） |
| `references/maintenance.md` | 新增接口维护规范（检查清单、权限标注指南） |
| `references/frontend-permissions.md` | 前端应用权限映射（平台应用 / 组织应用 / 公共应用） |
| `references/role-permissions.md` | 角色权限细节（角色总览、按角色区分的域、可操作范围、接口侧身份判定） |
| `references/app-selection.md` | 子域与端点清单（API 域划分、IoT AI 工具专题与应用场景表） |
| `references/upload-file.md` | 文件上传（`ur system upload-file` 参数、示例与临时桶/公开桶说明） |
