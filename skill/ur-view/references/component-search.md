# 组件与图元搜索

CLI v0.9.0 起提供组件目录和内置图元的专用查询命令。命令、本文、一次图案例及随包 Skills 同步发布。先执行 `ur --version` 核对安装版本，旧版执行 `ur upgrade --version v0.9.2`，升级会同步随包 Skills 与已识别的客户端副本。安装后仍使用旧技能的 AI 客户端需重开会话。

v0.9.2 统一修复 CLI 二进制替换、随包 Skills 同步和单独 Skills 更新的跨文件系统处理。旧版首次跨盘升级报错时，可重新运行官方安装脚本，Linux/macOS 使用以下命令，脚本同时安装 CLI 与 Skills。

```bash
curl -fsSL https://doc.unitedrhino.com/cli/install.sh | UR_TAG=v0.9.2 bash
```

## 离线组件目录

组件决定画布的渲染与数据能力。以下命令读取随包快照，无需登录。

```bash
ur view component search --keyword 连线 --json
ur view component search --keyword TextCommon --json
ur view component get-one --key Image --json
ur view component get-one --key ConnectLine --json
ur view component get-one --key TextCommon --json
```

`search` 的 `--keyword` 可省略，返回全部组件。关键词匹配名称、key、chartKey、conKey、包和分类，英文忽略大小写。`get-one --key` 使用精确注册 key。

JSON 成功结果为 `{code:200,msg,data}`。搜索的 `data` 包含 `list`、`total`、`snapshotDate`、`source`。每项包含注册 key、标题、分类、渲染框架、`iotQueryTypes` 和 `singleValueSelfRender`，预置组件另有 `redirectComponent`。详情的 `data` 是单个组件。空搜索成功返回空数组，详情不存在返回失败。

`ConnectLine` 保留供存量画布加载和校验，当前常用面板使用 `FlowChartLine`。新画布应按组件实际配置选择，存量一次图中的 `ConnectLine` 可继续查询和编辑。目录为快照，目标平台需已部署对应组件。

## 在线内置图元

图元是供 Image 引用的设备图形。先执行 `ur check --json` 复用认证，查询强制使用物联网应用上下文 `app-id=200`，企业范围沿用当前配置。

```bash
ur view builtin-image search --keyword 圆形 --category electric --page 1 --size 10 --json
ur view builtin-image search --keyword 保护开关 --status confirmed --size 10 --json
ur view builtin-image get-one --id bi-electric-0952 --json
```

| 参数 | 含义 |
| --- | --- |
| `--keyword` | 服务端匹配名称、别名、标签和说明 |
| `--category` | 一级分类 ID，如 electric、symbol、industrial |
| `--sub-category` | 二级分类名称 |
| `--tags` | 逗号分隔标签，命中任一 |
| `--format` | svg、png、jpg、gif |
| `--status` | confirmed 或 unconfirmed |
| `--page` / `--size` | 页码从 1 开始，默认每页 48 条，上限 100 条 |
| `--id` | 详情查询的稳定图元 ID |
| `--json` | 返回 code/msg/data 结构化结果 |

搜索复用 POST `/api/v1/view/builtin-image/get-list`，详情复用 POST `/api/v1/view/builtin-image/get-one`。返回保留服务端字段，包括 `id`、`name`、`description`、`url`、`thumbUrl`、`width`、`height`、分类与确认状态。列表的 `data.list` 为本页结果，`data.total` 为命中总数。排序与匹配规则由平台决定。

未带 `--json` 时输出格式化的 data。参数无效、认证失败、业务失败均返回非零退出码，业务失败的 JSON 保留服务端 code/msg。认证或网络层失败可能只输出 stderr，脚本需同时检查退出码。

## 从搜索到能源配电图

1. 搜索 Image、连线与 TextCommon，核对目标平台组件和 IoT 支持矩阵。
2. 搜索图元、查询详情并查看原图，核对器件角色、方向和端点。confirmed 仍需与目标图纸核对。
3. 从已有画布或组件配置复制完整默认结构，设备 Image 保存 `assetId`、`assetSource: "builtin"`、`assetName`，`dataset` 引用返回 URL。相对 URL 按目标站点解析。
4. 按区域或图纸组织独立器件、文字、连线。电表逐台独立绑定，缺少数据显示 `--`，推测连接与设计参数标记待核对。
5. 使用 `ur view screen validate -f <canvas.json>`、`describe` 核查，普通生成保存草稿，获得发布授权后才 `push --publish`。
6. 按[能源接入指南](primary-diagram-energy.md)设置区域 profile，进入能耗中心“电力集抄 → 配电图”，回读绑定并实测目标表实时更新。

完整绘制与两回路示例见[一次图案例](primary-diagram-case.md)，器件与端口选择见[组件与图元](primary-diagram-components.md)。
