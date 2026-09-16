---
name: ai-tool
description: "Use when developing or validating AI Tool artifacts or rule-engine AI tasks, including executor.js, skill.md, manifest.json, scheduled HTTP calls, platform API workflows, device-status forwarding, cross-device control, and Bun sandbox execution."
---

# ai-tool — AI 工具开发

## 核心概念

| 概念 | 说明 |
|------|------|
| executor.js | Bun 单次执行逻辑；调平台用 `Bun.spawn` 执行 `ur`，调外部 HTTP 用 `fetch` |
| skill.md | 用户可见的工具说明、参数、输出、副作用和失败语义 |
| manifest.json | 工具元信息、Bun 启动命令、输入与输出定义 |
| 平台 API | 使用内置 `ur` CLI，项目接口传 `--project-id`，认证只从执行环境继承 |
| 外部 HTTP | 使用 Bun `fetch`，不向外部主机转发 `UR_*` 密钥或平台认证头 |

## 规则引擎 AI 任务

规则引擎 AI 任务是通用自动化工具，不是设备数据模拟器。常见链路包括：

- 调用用户指定的 HTTP 接口；
- 用 `ur` CLI 查询当前项目的设备状态，再把最小必要结果 POST 到用户指定的 HTTP 接口；
- 根据查询结果控制另一台设备；
- 巡检、汇总、转发或模拟数据。

平台的 timedsvr 负责单次、间隔、每日和每周调度。`executor.js` 每次只执行一轮业务，不使用 `setInterval`、常驻循环、系统 cron 或自建调度器。用户说“每 5 分钟”时，工具只生成单次逻辑，频率留给发布后的调度表单。

开发时遵守以下合同：

1. 工具通用输入仅保留字符串 `projectID` 和可选 `prompt`；不用 JavaScript `Number` 承载项目 ID。
2. 调平台前先读 `ur-api` 中对应业务子域，再用 `ur --help` 与子命令 `--help` 核对真实命令；不猜 API 路径。
3. 组合任务按业务语义串行：上游查询成功且结果合法后才执行下游转发或控制，任一步失败都以非零退出。
4. 外部 HTTP 的 URL、method、body 和所需业务认证不明确时先询问，不猜测目标。外部请求设置有界超时，不发送 `UR_ACCESS_KEY`、`UR_ACCESS_SECRET`、`UR_TOKEN` 或平台认证头。
5. `Bun.spawn` 必须以 `const exitCode = await proc.exited` 获取最终退出码，不读运行中可能为 `null` 的 `proc.exitCode`。失败时抛错或 `process.exit(1)`，不在 stdout 打印 `code=500` 后正常退出。
6. stdout 最后一行只输出成功 JSON，过程日志写 stderr。用户只要生成或保存代码时，不自动试运行。

设备模拟或控制是上述通用能力的一个分支。只有进入该分支时才读 `ur-device/references/device-control.md`，按用户目的区分云端改值、模拟设备上报、实体控制和只生成样例。不应把该分支的接口、参数或澄清问题套用到普通 HTTP、巡检或数据转发任务。

> 下方组件标签与 `document.md` 章节仅适用于旧版可视化工具，不是规则引擎 AI 任务的产物合同。

## 快速开始

### 1. 获取三件套

```bash
ur ai-tool artifact get --id 5 --output-dir ./workspace
# 输出：workspace/executor.js, workspace/document.md, workspace/manifest.json
```

### 2. AI 编辑

```bash
ur ai-tool edit --id 5 --instruction "添加一个 OEE 趋势图，x轴时间 y轴OEE值"
# LLM 自动修改三件套并保存到 S3
```

### 3. 校验

```bash
ur ai-tool validate --id 5
# 检查变量一致性、组件白名单、JS 安全、JSON 格式
```

### 4. 渲染预览

```bash
ur ai-tool render --id 5 --output preview.json
# 将 document.md 解析为 components + variables JSON
```

### 5. 沙箱运行

```bash
ur ai-tool run --id 5 --inputs '{"timeRange":"24h","line":"Line 01"}'
# 返回 runID，轮询 run-status 获取结果
```

## 开发工作流

### 新建工具

```bash
# 1. 创建工具（草稿）
# 注：create 暂无 CLI 封装，需使用 ur api 直接调用
ur api /api/v1/ai/tool/create \
  --body '{"agentGroupID":"1","code":"oee-analysis","name":"OEE分析"}'

# 2. AI 生成初始三件套
ur ai-tool edit --id <new_id> --instruction "创建一个 OEE 分析工具：读取设备生产数据，计算 OEE，展示趋势图和停机事件表"

# 3. 校验
ur ai-tool validate --id <new_id>

# 4. 沙箱测试
ur ai-tool run --id <new_id> --inputs '{"timeRange":"24h"}'

# 5. 发布
# 注：publish 暂无 CLI 封装，需使用 ur api 直接调用
ur api /api/v1/ai/tool/publish --body '{"id":"<id>"}'
```

### 修改已有工具

```bash
# 1. AI 编辑
ur ai-tool edit --id 5 --instruction "把停机事件表按产线分组"

# 2. 校验 + 测试
ur ai-tool validate --id 5
ur ai-tool run --id 5 --inputs '{"timeRange":"24h"}'
```

### 本地直接修改

```bash
# 1. 拉取到本地
ur ai-tool artifact get --id 5 --output-dir ./workspace

# 2. 手动编辑文件
vim ./workspace/executor.js
vim ./workspace/document.md
vim ./workspace/manifest.json

# 3. 校验
ur ai-tool validate --id 5

# 4. 保存回平台
ur ai-tool artifact save --id 5 --dir ./workspace

# 5. 沙箱测试
ur ai-tool run --id 5 --inputs '{}'
```

## document.md 组件标签

### 语法

```markdown
# 页面标题

## 关键指标
<metric title="平均OEE" value="{{oee}}" unit="%" />

## 趋势图
<chart type="line" data="{{oeeTrend}}" x="time" y="value" title="OEE趋势" />

## 设备状态流转
<mermaid-diagram chart="{{flowDef}}" />

## 停机事件
<table-cpt data="{{downtimeEvents}}" />

## 执行步骤
<steps data="{{steps}}" />

## 状态
<status value="{{status}}" />

## 告警
<alert type="warning" message="{{alertMsg}}" />

## 原始数据
<json-view data="{{rawData}}" />
```

### 白名单组件

| 标签名 | 组件 | Props | 依赖 |
|--------|------|-------|------|
| `<chart>` | Chart | type, data, x, y, title, height | ECharts |
| `<metric>` | Metric | title, value, unit, trend, status | 纯 HTML/CSS |
| `<table-cpt>` | Table | data, columns, maxHeight, pageSize | 纯 HTML/CSS |
| `<steps>` | Steps | data[]: {id, title, status, summary} | 纯 HTML/CSS |
| `<status>` | Status | value | 纯 HTML/CSS |
| `<alert>` | Alert | type, message | 纯 HTML/CSS |
| `<mermaid-diagram>` | Mermaid | chart, theme | Mermaid.js |
| `<json-view>` | JsonView | data | 纯 HTML/CSS |

### 变量绑定

- `{{varName}}` — 从 state 读取顶层字段
- `{{obj.key}}` — 从 state 读取嵌套字段（暂未实现）
- 变量在 executor.js 中通过 `runtime.set("varName", value)` 写入

## CLI 命令缺口

以下操作暂无对应的 CLI 命令，仍需使用 `ur api` 直接调用：

| API | 说明 |
|-----|------|
| `/api/v1/ai/tool/create` | 创建工具 |
| `/api/v1/ai/tool/publish` | 发布工具 |

## 注意事项

- 联犀平台业务接口使用 POST；外部 HTTP 按用户指定的 method，不强制改为 POST
- 运行时通过 `UR_*` 环境变量注入认证
- executor.js 禁止使用 eval、Function、require
- 组件标签使用 kebab-case，属性使用双引号
- 校验命令本地执行，不消耗平台资源
- AI 编辑通过 LLM 接口实现
