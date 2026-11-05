# 一次图组件与图库图元

与[一次图绘制流程](primary-diagram-case.md)配合使用。画布采用 GoView 原生组件，每个器件、标注和连接线有稳定独立 ID。

## 组件选择

CLI v0.9.0 起可使用 `ur view component get-one --key ConnectLine --json` 核对注册信息，使用 `ur view builtin-image search --keyword 互感器 --category electric --json` 搜索图元，再 `get-one --id bi-electric-0952 --json` 查看详情。完整参数、版本与错误处理见[搜索说明](component-search.md)。

| 作用 | 组件与注册信息 | 配置要点 |
| --- | --- | --- |
| 保护开关、CT、变压器、隔离开关等设备 | `Image`，`Informations/Mores`，`VImage` / `VCImage` | `dataset` 使用图库返回 URL，保存 `assetId`、`assetSource`、`assetName`，电气符号可用 `tintColor` 与 `symbolRotation` |
| 母线、直角接线、计量二次线、出线箭头、仪表框 | `ConnectLine`，`Decorates/FlowChart`，`VConnectLine` / `VCConnectLine`，`chartFrame: "static"` | 明确 `points` 折点、线宽和颜色，静态一次图设置 `animate: false`、`direction: "static"` |
| 标题、柜号、用途、规格、Ia/Ib/Ic/P 读数 | `TextCommon`，`Informations/Texts`，`VTextCommon` / `VCTextCommon` | `fontSize`、`fontColor`、独立边界和格式，实时值使用 IoT request |

每个组件补齐 `attr`、`styles.animations`、`status`、`events` 和 `request` 默认字段。完整结构可以直接参考[两回路画布示例](../assets/primary-diagram/region-canvas.json)。`attr` 保存位置与尺寸，`chartConfig` 保存组件注册信息，`option` 保存显示配置，`request` 保存数据来源。

仪表框和参数表边界采用闭合 ConnectLine。空 TextCommon 的内容盒没有有效宽高，边框随内容计算，适合承载文字。

## 经看图核对的候选

这些图元仍需与目标图纸核对，尤其注意接法、方向和实际端子。

| 图元 ID | 可见结构与使用方式 |
| --- | --- |
| `bi-symbol-0038` | 竖向引线、断口与触点的保护开关候选，原 PNG 为透明背景。纵向放置角度为 0°，横向为 -90°，型号与状态取实际资料 |
| `bi-electric-0952` | 一次引线穿单圆，附二次引出线。三相 CT 使用三个独立 Image，纵向为 0°，横向为 90° |
| `bi-electric-0930` | 水平断开开关，两侧圆形端点与倾斜触臂，按隔离角色核对。原图为横向长条，横向 0°，纵向 90° |
| `bi-electric-0972` | 上下相交双圆的双绕组候选，接法标注另按原图核实 |
| `bi-electric-0987` | 上圆 Y、下圆三角的耦合符号，使用前核对上下绕组接法 |
| `bi-electric-0980` | 三条渐短横线的接地候选，接地类别与方向按原图确定 |
| `bi-electric-0984` | 电容极板符号，用于明确的电容支路 |
| `bi-electric-0985` | 电抗器候选，按图纸确认接法与方向 |

查询使用 POST `/api/v1/view/builtin-image/get-list` 与 `/api/v1/view/builtin-image/get-one`。图元信息取 `id`、`name`、`description`、`url`、`width`、`height`，鉴权使用物联网应用上下文 `app-id: 200`。

先看原图，再确定角色。图库中的实物电表外形适合设备展示，电气计量关系采用独立仪表功能框。原图缺少精确匹配时记录替代原因，保留角色文字或核对过的基础符号。

## 端口与透明度

Image 采用原图透明通道着色，`fit: "contain"` 保持宽高比。旋转长条符号时使用足够大的方形占位，核对旋转后真实引线端点。深色背景下可用临时浅底预览检查透明边界，再回到金线画面验收。

`bi-electric-0952` 的原 SVG 为 64×64，一次引线位于 x≈19.076、y≈6.099～57.9，二次引线沿 y≈32 向右延伸至 x≈57.882。缩放和旋转后按这些真实端点接线。三相分别排布，并在模型中区分一次端口与二次端口。

连接线的 `attr` 使用折点包围盒，`option.points` 使用相对包围盒的局部坐标。一次线路在开关两端分别结束和起接，二次线从已声明的 CT 端口引到 meter 功能框。

```json
{
  "pathType": "straight",
  "points": [{"x": 0, "y": 0}, {"x": 80, "y": 0}, {"x": 80, "y": 40}],
  "startPoint": {"x": 0, "y": 0},
  "endPoint": {"x": 80, "y": 40},
  "lineWidth": 2,
  "baseColor": "#ffb600",
  "animate": false,
  "direction": "static",
  "arrow": "none"
}
```

出线箭头需要同时设置 `arrow: "end"`、`arrowStyle: "triangle"`、`arrowSize` 和 `endArrowColor`。ConnectLine 的 `lineStyle: "dashed"` 当前没有渲染支持，计量二次线可用较细的静态线表达。

## 读数与设计参数

TextCommon 的 `option.dataset` 初始值为 `--`。显示格式与绑定配置分别填写。

```json
{
  "dataset": "--",
  "fontSize": 19,
  "fontColor": "#ffee00",
  "backgroundColor": "transparent",
  "format": {
    "enabled": true,
    "prefix": "Ia ",
    "suffix": " A",
    "decimals": 1,
    "useThousands": false
  }
}
```

A/B/C 可分别使用黄、绿、红，相色用于读数识别。电气状态颜色应依据明确状态属性和现场规则配置。进线电压与电流各成一列，馈线读数靠近对应器件链。

开关额定电流、CT 变比、母线和电缆规格来自图纸或已有配置，填写模型 `specification`。用户选区资料中尚未提供的规格写待核对，并记录资料来源。完整的实际属性绑定方法见[能源接入指南](primary-diagram-energy.md)。

## 数量与交互

可按一级区域或母线段拆屏，保证全量回路可检索和编辑。先做一个完整区域样板，再按同类结构应用其他区域。CAD 逐路径实体可能达到数千项，需实际检查完整规模下的加载、选中和保存响应。

编辑器中设备 Image、文字 TextCommon、连接 ConnectLine 分别选择和拖动，保存重开核对坐标与绑定。自动化定位器可使用 `.go-shape-box[data-id="<组件ID>"]`，实际内容边界从其 `.edit-content-chart` 获取，外层 shape 包装自身可能为零尺寸。
