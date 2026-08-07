---
name: ur-view
description: "大屏可视化管理：大屏（GoView）项目 CRUD、画布 JSON 本地编辑与推送、发布/取消发布、素材库管理、IoT 数据绑定、页面截图调优、一次图/配电 CAD 底图+实时数据叠加案例（多模态核对）。triggers: 大屏, 数据可视化, GoView, 画布, 组件, 实时数据, 看板, 编辑大屏, 发布大屏, 大屏截图, 可视化大屏, bigscreen, view, 一次图, 配电, CAD, dwg, 底图"
metadata:
  hermes:
    tags: [view, bigscreen, goview, visualization, iot]
---

# ur-view — 大屏可视化（GoView）

> 手写维护文件。文末 API_LIST 标记区域由 `scripts/generate-api-lists.py` 自动重写，其余章节手写维护；正文中请勿再出现同名标记文本，否则会被生成器误吞。

大屏域 API 全部为 **POST** `/api/v1/view/*`，鉴权 **app-id=200**（iot 应用上下文）。CLI 提供两层能力：

- `ur view screen`：大屏项目 CRUD、发布管理、画布迁移（pull/push）、本地校验（validate）、组件摘要（describe）、页面截图（screenshot）
- `ur view asset`：素材库（图片/视频/音频）查询、上传、删除
- 兜底：`ur api /api/v1/view/... --body '{...}'` 直接调用任意端点（端点清单见文末 API_LIST 区）

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

---

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

### content 顶层结构

| 字段 | 说明 |
|------|------|
| `editCanvasConfig` | 画布配置：`width` / `height` / `background` / `backgroundImage` / `chartThemeColor` / `previewScaleType` / `projectName` |
| `componentList` | 组件元素数组（见下节结构速览） |
| `requestGlobalConfig` | 全局请求配置（缺失时 validate 仅给 warning，组件用默认配置） |

### 组件元素字段

| 字段 | 说明 |
|------|------|
| `id` | 组件实例 ID，画布内全局唯一 |
| `title` | 组件标题（显示名） |
| `chartKey` / `conKey` | 渲染/配置面板注册键，约定 `chartKey="V"+key`、`conKey="VC"+key` |
| `chartConfig` | 组件元数据：`key` / `chartKey` / `conKey` / `title` / `category` / `package` / `chartFrame` |
| `attr` | 布局：`x` / `y` / `w` / `h` / `zIndex` |
| `option` | 组件专有配置（样式 + 静态 `dataset`） |
| `styles` | 通用样式 |
| `request` | 数据绑定：`requestDataType` + 各类型对应配置（IoT 见下节） |
| `status` | `{lock, hide}` 锁定与隐藏 |
| `events` | `{baseEvent, advancedEvents, interactEvents}` 事件三件套 |
| `isGroup` / `groupList` | 分组组件标记与子组件列表（分组无独立 key 体系） |

---

## 图片素材工作流

```bash
# 1. 上传（一条命令串两步：system/common/upload-file 拿 URL → view/asset/upload 登记）
ur view asset upload -f ./logo.png --name "企业LOGO" --tags "logo,brand"

# 2. 拿到返回的 url 后，在画布 JSON 中引用：
#    - Image 组件（Informations/Mores/Image）：option.dataset = "<url>"
#    - 画布背景图：editCanvasConfig.backgroundImage = "<url>"

# 3. 管理素材
ur view asset get-list --type image --page 1 --size 20
ur view asset get-list --name logo
ur view asset delete --id <assetID>
```

上传细节：
- 第一步走 `system/common/upload-file` 的 **multipart** 表单，与前端封面上传同口径：`isPublic=true` + `business=view` + `scene=goView/asset` + `useBy=user`，返回**公开桶永久 URL**（`public-*.cos.*`），可直接在画布 JSON 中长期引用
- `--tags` 传逗号分隔字符串（如 `"logo,brand"`），后端按字符串存储
- 上传时按扩展名自动推断类型：image（jpg/jpeg/png/gif/svg/webp/bmp/ico）、video（mp4/webm/mov）、audio（mp3/wav/ogg），其余归 `other`

---

## IoT 数据配置手册

组件数据绑定由 `request.requestDataType` 决定：

| requestDataType | 类型 | 说明 |
|-----------------|------|------|
| 0 | 静态 | 用 `option.dataset` 内置数据，不发请求 |
| 1 | AJAX | 走 `request.requestUrl` 等 HTTP 配置 |
| 2 | Pond | 公共数据池 |
| 5 | IoT 设备数据 | 走 `request.requestIoTDeviceConfig`，接设备属性/状态/信息 |

`requestDataType=5` 时 `requestIoTDeviceConfig` 必填，核心字段：`queryType`（`property` / `deviceStatus` / `deviceInfo`）、`productID`、`deviceNames`、`dataIDs`、`mode`（`latest` / `log`）。**注意**：`deviceStatus` 按产品维度统计在线状态，**不需要** `deviceNames`/`dataIDs`（真实画布即只填 `productID`，可留空数组）；`property`/`deviceInfo` 必须填非空 `deviceNames` 与 `dataIDs`（validate 会按此规则校验）。完整字段表见 [references/components.md](references/components.md) 附录 C。

### queryType 支持矩阵摘要（以附录 D 真实矩阵为准）

| 组件分组 | property | deviceStatus | deviceInfo |
|----------|----------|--------------|------------|
| 折线组（LineCommon 等） | ✓（仅 log） | ✗ | ✗ |
| 饼图组（PieCircle / PieCommon） | ✓（仅 latest） | ✓ | ✗ |
| 柱状组（BarCommon / BarCrossrange / BarLine） | ✓ | ✓ | ✗ |
| 文字组（TextCommon / FlipperNumber） | ✓ | ✓ | ✗ |
| 表格-信息组（TablesBasic / TableList） | ✓ | ✓ | ✓ |
| 表格-滚动组（TableScrollBoard） | ✓ | ✓ | ✗ |
| 雷达组（Radar） | ✓ | ✓ | ✗ |
| 进度组（Process / WaterPolo） | ✓（仅 latest） | ✓ | ✗ |
| 胶囊组（CapsuleChart） | ✓（仅 latest） | ✓ | ✓ |
| 其他可绑定数据组件 | ✓ | ✓ | ✓ |
| 静态/装饰组件 | ✗ | ✗ | ✗ |

单值自渲染组件（Dial / PieCircle / Process / WaterPolo）拿到 IoT 单值后自行渲染，不走 ECharts dataset 覆盖。

### 常见配方

**① 单值实时展示 → Dial / PieCircle / Process / WaterPolo**

`queryType=property` + `mode=latest`，单设备单属性取最新值，组件自渲染：

```json
"request": {
  "requestDataType": 5,
  "requestIoTDeviceConfig": {
    "queryType": "property",
    "mode": "latest",
    "productID": "p_temp_001",
    "deviceNames": ["sensor-001"],
    "dataIDs": ["CurrentTemperature"]
  }
}
```

**② 多设备同属性对比 → BarCommon（或 CapsuleChart）**

`queryType=property`，`deviceNames` 填多台设备、`dataIDs` 填同一属性：

```json
"requestIoTDeviceConfig": {
  "queryType": "property",
  "productID": "p_temp_001",
  "deviceNames": ["sensor-001", "sensor-002", "sensor-003"],
  "dataIDs": ["CurrentTemperature"]
}
```

**③ 趋势曲线 → LineCommon（折线组仅支持 log 模式）**

`queryType=property` + `mode=log`，配 `logTimeRange` / `logInterval`：

```json
"requestIoTDeviceConfig": {
  "queryType": "property",
  "mode": "log",
  "logTimeRange": "24h",
  "logInterval": "1h",
  "productID": "p_temp_001",
  "deviceNames": ["sensor-001"],
  "dataIDs": ["CurrentTemperature"]
}
```

**④ 设备在线状态占比 → PieCircle / PieCommon（饼图组）**

`queryType=deviceStatus`，统计所选设备在线/离线占比：

```json
"requestIoTDeviceConfig": {
  "queryType": "deviceStatus",
  "productID": "p_temp_001",
  "deviceNames": ["sensor-001", "sensor-002"],
  "dataIDs": ["CurrentTemperature"]
}
```

**⑤ 设备信息卡片/列表 → TableList / TablesBasic（表格-信息组）**

`queryType=deviceInfo` 展示设备信息字段。**注意**：`TextCommon` 的 `getQueryTypeHint` 提示语声称支持设备信息，但 `isQueryTypeSupported` 实际判定文字组**禁用 deviceInfo**——不要用 TextCommon 绑设备信息，以运行时矩阵为准。

**⑥ 大数字单值 → Number / CountDown（装饰组，queryType 全放行）**

`queryType=property` + `mode=latest`，适合做大字号的单值看板。

---

## 组件清单

GoView 共 **77 个静态注册组件**，分 8 类：Charts 图表（23）、Informations 信息（13）、Tables 表格（3）、Decorates 装饰（30）、Icons 图标（1）、Photos 图片（0）、Presets 预置（6）、Interact 交互（1）。Icons 面板另有 127 个动态图标条目（`uim:` / `line-md:` / `wi:` 前缀，统一重定向到 Icon 组件）；Photos 无静态组件，素材库/本地上传/共享图片运行时动态生成，重定向到 Image 组件。

完整清单（key/chartKey/chartFrame/IoT 支持矩阵/适用场景）见 **[references/components.md](references/components.md)**。

## 组件结构速览

可直接抄改的 `componentList` 元素骨架。示例一：静态柱状图（BarCommon，requestDataType=0）：

```json
{
  "id": "comp-bar-001",
  "title": "产量对比",
  "chartKey": "VBarCommon",
  "conKey": "VCBarCommon",
  "chartConfig": {
    "key": "BarCommon",
    "chartKey": "VBarCommon",
    "conKey": "VCBarCommon",
    "title": "柱状图",
    "category": "Charts",
    "categoryName": "图表",
    "package": "Charts",
    "chartFrame": "echarts",
    "image": "bar.png"
  },
  "attr": { "x": 60, "y": 120, "w": 500, "h": 300, "zIndex": 1 },
  "option": {
    "dataset": {
      "dimensions": ["product", "data1"],
      "source": [
        { "product": "一车间", "data1": 120 },
        { "product": "二车间", "data1": 200 }
      ]
    }
  },
  "styles": {},
  "request": { "requestDataType": 0 },
  "status": { "lock": false, "hide": false },
  "events": { "baseEvent": {}, "advancedEvents": {}, "interactEvents": [] }
}
```

示例二：IoT 绑定表盘（Dial，requestDataType=5，单值自渲染）：

```json
{
  "id": "comp-dial-001",
  "title": "实时温度",
  "chartKey": "VDial",
  "conKey": "VCDial",
  "chartConfig": {
    "key": "Dial",
    "chartKey": "VDial",
    "conKey": "VCDial",
    "title": "表盘",
    "category": "Charts",
    "categoryName": "图表",
    "package": "Charts",
    "chartFrame": "common",
    "image": "dial.png"
  },
  "attr": { "x": 640, "y": 120, "w": 300, "h": 300, "zIndex": 2 },
  "option": { "dataset": 66 },
  "styles": {},
  "request": {
    "requestDataType": 5,
    "requestIoTDeviceConfig": {
      "queryType": "property",
      "mode": "latest",
      "productID": "p_temp_001",
      "deviceNames": ["sensor-001"],
      "dataIDs": ["CurrentTemperature"]
    }
  },
  "status": { "lock": false, "hide": false },
  "events": { "baseEvent": {}, "advancedEvents": {}, "interactEvents": [] }
}
```

分组组件：`"isGroup": true` + `"groupList": [...]`，无独立 chartKey/conKey，validate 会递归校验子组件。

---

## 一次图 + 实时数据绑定案例

案例专项内容独立成文，主文档仅留索引（避免占用上下文）：
见 [references/primary-diagram-case.md](references/primary-diagram-case.md) ——
一次图背景 + 回路实时 U/I/P/电能叠加、多模态核对流程、全量节点布局、
列表封面回写与完整 CLI 命令序列（含生产案例实录，环境取值占位符化）。

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
- **组件数据快照**：`references/components.md` 由 `scripts/gen-view-components.py` 生成（快照日期 2026-08-03），前端组件变更后重新运行 `python3 scripts/gen-view-components.py` 刷新。
