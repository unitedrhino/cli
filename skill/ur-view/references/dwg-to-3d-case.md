# DWG 一次图 → 3D 电力组态图（离线复原案例）

适用场景：用户给一张配电一次图 DWG（AutoCAD），要求"把这张 DWG 画成 3D 电力组态图 /
三维组态 / 系统图 3D 化"时，从图纸解析生成可交互的 Three.js 单页 3D 组态。
模板与解析器随技能附带：[assets/dwg-primary-3d/](../assets/dwg-primary-3d/README.md)
（离线 three.js r128 模板 + 纯标准库 DXF 解析器），环境特定取值（租户/项目/图纸名）一律按目标输入替换。

## 三条组态路线怎么选（先选路，再动手）

| 路线 | 输入 | 产出 | 数据形态 | 适用 |
|---|---|---|---|---|
| 2D 一次图叠加（[primary-diagram-case.md](primary-diagram-case.md)） | DWG 渲染 PNG 底图 + 真实设备清单 | GoView 画布（Image + TextCommon） | 真实遥测（WS 实时推送） | 正式监控一次图大屏 |
| 3D 数字孪生（[scene-templates.md](scene-templates.md) + [embed-page-3d-case.md](embed-page-3d-case.md)） | power-station 模板 + 真实设备/物模型 | 场景 ZIP + EmbedPage 接入 | ur-scene SDK 真实绑定 | 正式数字孪生 |
| DWG→3D 图纸复原（本文） | DWG 图纸 | 离线 three.js 单页 HTML | 图纸静态数据 + 模拟 / 外部钩子 | 核图、评审、汇报演示 |

**红线**：正式环境监控/数字孪生**不得**用本路线交付——画面数据是图纸还原 + 模拟；
要接真实数据必须迁到 EmbedPage 路线，或用模板预留钩子 `PRIMARY3D.setCircuitValue()` / `setMode('live')`
自接 IoT 数据源（见文末）。本路线交付说明必须写明哪些内容来自图纸、哪些是模拟值。

**沉淀纪律（防再犯）**：画图任务开始必须先检索并加载已有技能（如本技能），禁止不加载就自行
从 DWG 解析到手写 Three.js 重造轮子；沉淀/更新技能前必须检查目标目录有无同名技能，先备份再写，
禁止直接覆盖。

## 流程（Windows + Git Bash 实战验证，2026-09）

### 1. DWG → DXF（LibreDWG）

```bash
# 只用 win64 包；win32 包缺 DLL，运行报 0xC0000135 (STATUS_DLL_NOT_FOUND)
# GitHub 直连慢：ghfast.top 前缀镜像 + curl -C - 断点续传（11~12 MB 常需续传一次）
curl -sL -C - -o libredwg64.zip \
  "https://ghfast.top/https://github.com/LibreDWG/libredwg/releases/download/0.14/libredwg-0.14-win64.zip"
unzip -t libredwg64.zip && unzip -oq libredwg64.zip -d x64

./x64/dwg2dxf.exe -y -o out.dxf <drawing>.dwg   # -y 覆盖输出
./x64/dwglayers.exe <drawing>.dwg               # 先列图层，定位电气层（开关柜设备/配电--N图层/电气火灾监控系统等）
```

- bash 直跑 win32 exe → exit 127；PowerShell 跑 → -1073741515，两种都是缺 DLL，换 win64。
- `dwg2SVG.exe` 不认 `-o`，输出走 stdout；SVG 线条渲染一般，DXF 文本解析更可靠。
- Linux 环境用发行版 libredwg 包或源码构建，工具同名（`dwg2dxf`/`dwglayers`），参数同理（未在本仓实测，以 `--help` 为准）。

### 2. 解析柜列/回路数据

```bash
python3 assets/dwg-primary-3d/parse_dxf3.py out.dxf all > full.txt
# all = 图层分布 + 全部文字(坐标/字高) + INSERT 块(坐标) + LINE/LWPOLYLINE 统计；纯标准库无依赖
```

从 `full.txt` 按下节规则解读电气内容，构建 `CABINETS` 列表（每柜一条：id/name/width/circuits，
回路含 id/name/breaker/ct/kw/run/fire）与 `SPECIAL` 条目（变压器/进线/总柜/SVG/电容柜）。

### 3. 模板生成 3D 组态（只改数据区）

复制 `assets/dwg-primary-3d/template/`（index.html + lib/，全程离线无 CDN）到输出目录，
**只改** `<script>` 顶部数据区：`CABINETS` / `SPECIAL` / `GAP`（柜中心距，默认 0.72 m）。
柜宽保持真实感：GCS 馈线柜 0.6 m、电容柜 0.8 m、总柜 1.0 m、干变外壳 2.0 m；布局函数自动排柜。
模板已实现（勿重写）：柜面 CanvasTexture 组态面板（回路行 + 实时电流/功率）、三相母排
L1黄/L2绿/L3红/N蓝 + PE 排、能量流光点、10kV 进线电缆、点击柜体相机 tween + 回路详情面板、
模拟数据（0.5 s tick）、消防绿/备用灰/超载 85% 橙告警、`?shot=1` 直达机位与 `?sel=<柜ID>` 自动选中。

### 4. headless 截图验证

```bash
python -m http.server <port>   # 输出目录起服务
# Edge headless（--screenshot 必须绝对路径，相对路径报"拒绝访问"）
"C:/Program Files (x86)/Microsoft/Edge/Application/msedge.exe" \
  --headless=new --disable-gpu --enable-unsafe-swiftshader --hide-scrollbars \
  --window-size=1920,1080 --virtual-time-budget=8000 \
  --screenshot="C:/abs/path/shot.png" "http://localhost:<port>/?shot=1"
```

- 必截两张：`?shot=1`（全景）与 `?shot=1&sel=<柜ID>`（详情面板 + 高亮 + 柜面近景）。
- `--virtual-time-budget` 下 rAF 开场 tween 可能不完成 → `?shot=1` 直达机位是硬要求。
- Linux 同参数用 chromium/chrome headless（未在本仓实测，以实际浏览器为准）。

### 5. 交付说明（诚实标注）

说明哪些柜列/回路参数来自图纸、回路号是否图纸缺失自行编号（示意）、布局柜序真实但间距示意
（图纸是原理图而非平面布置图）、实时数值当前为模拟。

## DXF 解析规则速查（LibreDWG 输出特性）

- `100 AcDbText` 之后的组码 `1` = 文字值；`100 AcDbAttribute` 之后的组码 `2` = 属性 tag，
  两者在同一 ATTRIB 实体里；按"组码行+值行"成对流解析，不要猜顺序。
- TEXT：图层在 `8`，插入点 `10/20`，字高 `40`；INSERT：块名 `2`，位置 `10/20`，缩放 `41/42`，旋转 `50`。
- 电气标注特征：字高 `h=151~301` 的大字（相对图内小注）；柜名标签 `h=301` 集中在同一 y（图顶部一行）。

### 柜列定位与回路归属

- 柜体块（如 `09-GCS馈电柜`）INSERT 的 x = 柜中心；柜名大字标签在 INSERT x + ~136 偏移处，y 一致。
- 柜间距 = 相邻 INSERT x 差；每个 GCS 柜对应**两列回路**（列距 = 柜间距/2），按回路号标签 x 距哪个柜中心最近归属。
- 底部"馈线柜/xxx柜 600x1000x2200"字样是柜体尺寸标注（宽×深×高），柜宽由此取。
- 一个回路的文字五件套（y 相近成组）：断路器 `xA/3P[+分励脱扣+辅助触点]`、CT `BH-0.66 x/5`、
  回路号（如 `1ATJK`/`3APWT`/`B1-1AT1`）、用途名（如"监控中心（主）"）、kW 数值。
- 「备用」回路 → `run:false`；消防类用途（消防风机/消防泵/应急照明/消防电梯/人防）→ `fire:true`。
- 无功补偿柜：容量（如 190kvar）+ 刀开关规格 + 电容/电抗组配置 + 避雷器型号。
- 变压器：型号/容量（如 SCB14-800kVA）、电压比 `10±2×2.5%/0.4kV`、额定 `In:46.2A/1154.7A`、
  `Uk%=6`、接线组别 `Dyn11`、高压进线电缆规格。
- 母线：主母排 `TMY-3x(80x10)+N(80x10)`、PE 排 `TMY-80x6`、母线槽电流等级（如 1600A 密集型）。

## 与联犀 IoT 对接（可选后续）

模板已留外部数据钩子。真实数据路径：`ur things device info get-list` 拿 deviceName/productID →
物模型标识符（Ua/Ia/P/TotalEnergy，取数命令见 [primary-diagram-case.md](primary-diagram-case.md)）→
`PRIMARY3D.setCircuitValue(柜ID, 回路ID, {Ia, P, Energy})` 逐点注入，或改造为 WS 订阅
`project.prop.{projectID}` 批量刷新。**回路名与设备名的映射需人工对一次**——图纸用途名
（如"3层办公用电"）≠ 设备别名；映射表建议存独立映射文件，不要硬编码进模板。

## 已知坑位（勿重新踩）

- THREE.Group 只 push 进 `items` 而未 `scene.add` 会整柜消失，只剩母排。
- headless 虚拟时间下开场相机 tween 可能永不结束，`?shot=1` 直达机位即为此设。
- LibreDWG 的 DXF 组码顺序与 AutoCAD 原生略有差异，必须按行对解析（见 parse_dxf3.py 实现）。
- agent-browser 的 bash sh wrapper 有路径拼接 bug（Windows 实测）；确需使用时直接
  `node <npm-global>/agent-browser/bin/agent-browser.js open <url>`。
