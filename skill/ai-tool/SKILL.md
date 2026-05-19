---
name: ai-tool
description: "AI Tool 三件套开发：executor.js + document.md + manifest.json。AI工具编辑、校验、沙箱运行、组件标签渲染。triggers: ai-tool, AI工具, 沙箱开发, executor.js, document.md, 组件标签, artifact, QuickJS, {{var}}"
metadata:
  hermes:
    tags: [ai, tool, sandbox, quickjs, artifact, render, validate]
---

# ai-tool — AI 工具开发

## 核心概念

| 概念 | 说明 |
|------|------|
| executor.js | QuickJS 执行逻辑，runtime.set/patch 更新状态，tier0.query 查询数据 |
| document.md | 用户可见文档，Markdown + 组件标签，{{var}} 变量绑定 |
| manifest.json | 元信息：title, runtime, inputs[], output, permissions |
| runtime API | runtime.set(key,value), runtime.patch(values), runtime.log(level,msg) |
| tier0 API | tier0.query(topic,options) 通过 os.exec("ur") 调平台 API |

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
# 1. 通过 API 创建工具（草稿）
ur api /api/v1/ai/tool/create \
  --body '{"agentGroupID":"1","code":"oee-analysis","name":"OEE分析"}'

# 2. AI 生成初始三件套
ur ai-tool edit --id <new_id> --instruction "创建一个 OEE 分析工具：读取设备生产数据，计算 OEE，展示趋势图和停机事件表"

# 3. 校验
ur ai-tool validate --id <new_id>

# 4. 沙箱测试
ur ai-tool run --id <new_id> --inputs '{"timeRange":"24h"}'

# 5. 发布
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

## API 索引

<!-- API_LIST:ai-tool -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/tool/get-list` | 获取工具列表 | admin |
| POST | `/api/v1/ai/tool/get-one` | 获取工具详情（含三件套）| admin |
| POST | `/api/v1/ai/tool/create` | 创建工具 | admin |
| POST | `/api/v1/ai/tool/update` | 更新工具 | admin |
| POST | `/api/v1/ai/tool/delete` | 删除工具 | admin |
| POST | `/api/v1/ai/tool/publish` | 发布工具 | admin |
| POST | `/api/v1/ai/tool/disable` | 停用工具 | admin |
| POST | `/api/v1/ai/tool/save-artifact` | 保存三件套 | admin |
| POST | `/api/v1/ai/tool/run` | 运行工具 | admin |
| POST | `/api/v1/ai/tool/run-status` | 查询运行状态 | admin |
| POST | `/api/v1/ai/tool/run-history` | 运行历史 | admin |

<!-- END_API_LIST -->

## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证
- executor.js 禁止使用 eval、Function、require
- 组件标签使用 kebab-case，属性使用双引号
- 校验命令本地执行，不消耗平台资源
- AI 编辑通过 `/api/v1/ai/chat/completions` 实现
