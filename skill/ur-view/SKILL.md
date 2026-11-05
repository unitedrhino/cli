---
name: ur-view
description: "大屏与 GoView 画布管理、图元组态、设备实时数据绑定、按区域或图纸绘制电力一次图，以及能源配电图和 3D 场景接入时使用。"
metadata:
  hermes:
    tags: [view, bigscreen, goview, visualization, iot]
---

# ur-view — 大屏可视化（GoView）

> 手写维护文件。文末 API_LIST 标记区域由 `scripts/generate-api-lists.py` 自动重写，其余章节手写维护；正文中请勿再出现同名标记文本，否则会被生成器误吞。

大屏域 API 全部为 **POST** `/api/v1/view/*`，鉴权 **app-id=200**（iot 应用上下文）。CLI 提供两层能力：

- `ur view screen`：大屏项目 CRUD、发布管理、画布迁移（pull/push）、本地校验（validate）、组件摘要（describe）、页面截图（screenshot）
- `ur view asset`：素材库（图片/视频/音频）查询、上传、删除
- `ur view component search/get-one`：离线查询组件目录、注册 key 与 IoT 支持矩阵（CLI v0.9.0 起）
- `ur view builtin-image search/get-one`：在线搜索图元及查询来源 ID、说明、URL 与尺寸（CLI v0.9.0 起）
- 兜底：`ur api /api/v1/view/... --body '{...}'` 直接调用任意端点（端点清单见文末 API_LIST 区）

## 电力一次图绘制、组件、案例与能源接入

搜索组件或图元、核对命令版本和参数时，读取[组件与图元搜索](references/component-search.md)。CLI 命令、Skills、案例和发布说明同步交付，升级使用 `ur upgrade --version v0.9.0`。

按区域生成、按图纸复原、区域加图纸匹配时，读取[一次图绘制与案例](references/primary-diagram-case.md)。其中包含完整 T2 图纸复原和按实际区域拆图两个案例，以及可导入的两回路画布示例。

选取保护开关、三相 CT、母线、仪表和实时读数时，读取[组件与图库图元](references/primary-diagram-components.md)。设备使用 Image，接线使用 ConnectLine，文字和测量值使用 TextCommon。

接入能耗中心“电力集抄 → 配电图”时，读取[能源接入与数据绑定](references/primary-diagram-energy.md)。文档包含设备属性绑定、实时订阅、保存发布、区域 profile 请求与回读验收。

## 核心概念

| 概念 | 说明 | 注意 |
|------|------|------|
| **project（大屏项目）** | 一个大屏即一个 project，`projectID` 为唯一标识 | CRUD 走 `/api/v1/view/project/*` |
| **content（画布 JSON）** | 大屏的完整画布描述，字符串化 JSON 存在 project detail 里 | 顶层结构：`editCanvasConfig` / `componentList` / `requestGlobalConfig` |
| **发布快照** | 发布 = `project/update {projectID, status:1}`，取消发布 status:2；无独立发布接口 | 发布页读已发布快照（`detail/get-one forView=true`），编辑态改动不影响发布页，**改完必须重新 publish** |
| **组件元素** | `componentList` 数组中的一项，含定位（attr）、渲染（chartConfig/option）、数据（request）、事件（events） | 结构速览见下文 |
| **asset（素材）** | 资源库条目（image/video/audio/other），URL 可在画布 JSON 中引用 | 上传是两步链路，CLI 已串成一条命令 |
| **WS 实时频道** | 大屏已绑定项目后，属性实时推送频道 `project.prop.{projectID}`、设备上下线频道 `project.conn.{projectID}` | 发布页实时数据依赖这两个频道 |

## 角色权限

| 角色 | 权限范围 |
|------|---------|
| 企业管理员 | 大屏 CRUD、画布编辑、发布管理、素材管理（app-id=200，iot 应用上下文） |
| 平台管理员 / 普通用户 | 无大屏管理入口 |

企业边界与菜单判定，以及能源模板验收合同见[大屏验收指南](references/bigscreen-acceptance.md)。

## 本地编辑工作流（AI 调优闭环，核心）

`detail/update` 保存接口**纯透传不做校验**，因此 CLI 在 `validate` / `push` 前置结构校验。推荐闭环：

```bash
# 1. 拉取远端编辑态画布到本地（同时落盘 <file>.meta.json 记录 projectID 等元信息）
ur view screen pull --id <screen-id> -o ./bigscreen.json

# 2. 本地编辑 bigscreen.json（改布局、改样式、绑 IoT 数据）

# 3. 本地校验（error 级问题退出码非 0；warning 不阻断）
ur view screen validate -f ./bigscreen.json

# 4. 查看组件状态摘要（title/key/坐标/层级/显隐/数据绑定，便于确认改对了组件）
ur view screen describe -f ./bigscreen.json
ur view screen describe -f ./bigscreen.json --json   # 结构化输出

# 5. 推送回远端并发布（push 默认先跑 validate，有 error 会拒绝；--force 跳过校验）
ur view screen push -f ./bigscreen.json --publish
# projectID 解析顺序：--id 参数 > <file>.meta.json > 文件名 bigscreen-<id>.json

# 6. 截图看效果（发布态页面；依赖本机 agent-browser）
ur view screen screenshot --id <screen-id> -o ./out.png \
  --front-base <front-base> --wait 8

# 7. 看图 → 不满意回到第 2 步迭代
```

辅助命令：

```bash
# 随时查看远端画布组件状态（不落盘）
ur view screen describe --id <screen-id>

# 截编辑态页面（编辑器界面，用于核对画布而非发布效果）
ur view screen screenshot --id <screen-id> -o ./edit.png \
  --front-base <front-base> --edit

# 前端地址也可走环境变量，省去每次 --front-base
export UR_FRONT_BASE_URL=<front-base>
```

### 命令速查

| 子命令 | 用途 | 常用 flags |
|--------|------|-----------|
| `ur view screen get-list` | 大屏项目列表 | `--body '{"page":{"page":1,"size":10}}'` `--json` |
| `ur view screen get-one` | 项目详情（元信息） | `--id <projectID>` |
| `ur view screen create` | 创建大屏（须绑定项目） | `--body '{"name":"xx","type":"screen","projectID":"<物联网项目ID>"}'` |
| `ur view screen update` | 更新项目元信息 | `--body '{"id":"...","name":"xx","desc":"..."}'` |
| `ur view screen delete` | 删除大屏 | `--id <projectID>` |
| `ur view screen publish` | 发布（status=1） | `--id <projectID>` |
| `ur view screen unpublish` | 取消发布（status=2） | `--id <projectID>` |
| `ur view screen pull` | 拉画布到本地 | `--id <projectID>` `-o <file>`（默认 `./bigscreen-<id>.json`） |
| `ur view screen push` | 推画布回远端 | `-f <file>` `[--id]` `[--publish]` `[--force]` |
| `ur view screen validate` | 校验画布 | `-f <file>` 或 `--id <projectID>` |
| `ur view screen describe` | 组件摘要 | `-f <file>` 或 `--id <projectID>` `[--json]` |
| `ur view screen screenshot` | 页面截图 | `--id <projectID>` `-o <png>` `[--front-base]` `[--wait 秒]` `[--edit]` `[--url-template]` |

## 画布、素材与实时数据

画布字段及完整组件骨架见[画布结构](references/canvas-structure.md)，素材上传和引用见[素材工作流](references/asset-workflow.md)，IoT 配置、支持矩阵与绑定配方见[数据绑定指南](references/iot-data-binding.md)。

## 组件清单

CLI 当前组件快照包含 **81 项**，涵盖面板注册组件及存量画布仍使用的 ConnectLine。按 8 类组织，具体数量与注册信息以 `ur view component search --json` 和生成清单为准。Icons 的动态图标和 Photos 的运行时素材统一重定向到 Icon、Image。

完整清单（key/chartKey/chartFrame/IoT 支持矩阵/适用场景）见 **[references/components.md](references/components.md)**。

## 3D 场景页接入案例（EmbedPage 内嵌页面，数字孪生）

见 [references/embed-page-3d-case.md](references/embed-page-3d-case.md) ——
自研/AI 生成 three.js 场景页 zip 托管上传（upload-zip）、EmbedPage 画布接入与锚点批量绑定、
封面补录（goView/projectIndexImage）、agent-browser 验证与缓存/代理排障。
完整方法论与可复用案例包在 saas 主仓 `docs/大屏/功能说明/内嵌页面组件/`。

## DWG 一次图 → 3D 组态图案例（离线复原）

用户要求"把这张 DWG 画成 3D 电力组态图/三维组态"时走本路线：
见 [references/dwg-to-3d-case.md](references/dwg-to-3d-case.md) ——
三条组态路线选择地图（2D 一次图叠加 / 3D 数字孪生 / DWG 图纸复原）、LibreDWG→DXF→
柜列/回路解析规则、模板数据区改法与 headless 截图验证、诚实标注原则、联犀 IoT 对接钩子；
模板与解析器随技能附带（[assets/dwg-primary-3d/](assets/dwg-primary-3d/README.md)）。
正式监控/数字孪生不得用本路线交付（数据为图纸还原+模拟）。

## API 端点索引

<!-- API_LIST:ur-view -->

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| `可视化管理` | 删除资源 (37 个端点) | [ur-view-可视化管理.md](references/api/ur-view-可视化管理.md) |

[查看全部端点](references/api/ur-view-all-endpoints.md)


<!-- END_API_LIST -->

## 注意事项

- **发布快照机制**：编辑态改动不影响发布页；`push --publish` 或单独 `publish` 后发布页才更新。截图验证效果时确认截的是发布态（默认）且已重新发布。
- **validate / push 校验规则摘要**（push 默认前置校验，error 阻断，`--force` 跳过）：
  - 顶层：`editCanvasConfig` 存在且 `width`/`height` 为正数；`componentList` 为数组；`requestGlobalConfig` 缺失仅 warning。
  - 组件级：`id` 非空且全局唯一；`chartConfig.key` 非空且在组件 key 表内；`chartKey == "V"+key`、`conKey == "VC"+key`；`attr.w`/`attr.h` 为正数，`x`/`y`/`zIndex` 为数值；`status` / `events` 三件套缺失仅 warning。
  - IoT 级：`requestDataType=5` 时 `requestIoTDeviceConfig` 必填，`queryType` 必须在该组件支持矩阵内，`productID` 非空、`deviceNames` / `dataIDs` 为非空数组。
- **screenshot 依赖本机 agent-browser**：依次执行 `open → wait（毫秒）→ screenshot`；`--wait` 单位是秒。前端地址取 `--front-base` 或环境变量 `UR_FRONT_BASE_URL`。实测：发布页默认模板 `{base}/app/iot/#/view/{id}`（`/view/` 兼容路由始终取已发布快照；`/iot/big-bigscreen/published/{id}` 在 iot 应用内会落入应用外壳不渲染画布，勿用），编辑态默认 `{base}/app/iot/#/iot/big-bigscreen/editor/{id}`；部署形态不同可用 `--url-template` 覆盖（支持 `{base}` / `{id}` 占位符）。发布页需登录态，建议先用 agent-browser 登录 iot 应用（租户管理员 + 选择企业）保持会话。
- **鉴权口径**：view 域全部走 iot 应用上下文（app-id=200），运行时通过 `UR_*` 环境变量注入认证，先 `ur check` 确认身份。
- **用户可见文案遵循根级文案规范**（禁用平台侧组织隔离术语，见仓库根 `AGENTS.md`）；代码层字段可保留原命名。
- **组件数据快照**：`references/components.md` 由 `scripts/gen-view-components.py` 生成（快照日期 2026-11-05），前端组件变更后重新运行 `python3 scripts/gen-view-components.py` 刷新。


## 建筑与配电站通用源码模板

制作建筑能耗、分层功率分布或配电站机房时，使用 [通用场景源码指南](references/scene-templates.md)。完整源码、本地依赖、真实数据适配、打包和测试随技能提供；默认无绑定、无模拟回退，不依赖业务案例仓库。
