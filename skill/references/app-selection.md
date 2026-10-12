# 子域与端点清单

本文件是 `SKILL.md`（ur-api 主入口）的子文档。主文档保留「按任务选择子域」核心表与业务分组导航，本文件收录 IoT AI 工具专题与 API 域划分。

## 第三步：IoT AI 工具调用设计

### 第三步：IoT AI 工具调用设计

下列三个客户端专题只随 CLI 完整技能包和 SaaS 整理副本提供，位于 `ur-iot/` 内；公开 skills 源仓不包含这些专题。

| 用户意图 / 关键词 | 加载子域 | 说明 |
|----------------|---------|------|
| MCP 工具、frontend tools、设备工具语义、物模型调用顺序、实时属性查询 | `ur-iot-device` | 关注设备工具本体与工具语义对齐 |
| 所有设备可见、设备上下文、session 绑定、工具作用目标、前端不负责绑定设备 | `ur-iot-context` | 关注 Web 前端全设备可见场景下的上下文约束 |
| Web、Win AI、MQTT、UDP、客户端差异、共存迁移 | `ur-iot-client` | 关注多客户端执行链路差异 |

版本说明见：`references/iot-tools-versions.md`

## API 域划分

共 7 个子域，覆盖 492 个端点（100% 覆盖）：

| 域 | 端点数 | 说明 |
|---|--------|------|
| ur-device | 85 | 设备管理：设备CRUD、状态、属性控制、认证、物模型 |
| ur-product | 86 | 产品管理：产品CRUD、物模型、协议脚本、OTA |
| ur-project | 41 | 项目管理：项目、区域、分组、数据权限、用户权限 |
| ur-user | 153 | 用户管理：登录、角色、部门、字典、任务、通知、日志 |
| ur-system | 53 | 系统管理：通用接口、应用管理、授权码、Hook能力 |
| ur-tenant | 38 | 企业管理：企业CRUD、用户管理、权限配置 |
| ur-ai | 36 | AI管理：Agent、告警、规则引擎、场景联动 |

详细 API 端点见各子域 SKILL.md：ur-device、ur-product、ur-project、ur-user、ur-system、ur-tenant、ur-ai
