# dwg-primary-3d — DWG 一次图转 3D 组态模板与解析器

配合 [references/dwg-to-3d-case.md](../../references/dwg-to-3d-case.md) 使用（流程、解析规则、
路线选择、坑位都在案例文档，本文只说资产本身）。

## 目录结构

| 路径 | 用途 |
|---|---|
| `parse_dxf3.py` | DXF 文字/INSERT/图层提取器（纯标准库，按"组码行+值行"成对解析） |
| `template/index.html` | 已验证可跑的 3D 一次图组态模板（复制后只改顶部数据区 `CABINETS`/`SPECIAL`/`GAP`） |
| `template/lib/three.min.js` | 本地 three.js r128（离线，无 CDN） |
| `template/lib/OrbitControls.js` | 轨道控制器（离线） |

## 边界与口径

- 模板 `index.html` 内置的柜列/回路数据是**源图纸示例**（1# 变 800kVA 低压 GCS 系统），换图纸时整体替换，不得当作现场资料引用。
- 本资产产出的 3D 组态为**图纸复原 + 模拟数据**，仅用于核图、评审、汇报演示；正式监控/数字孪生走 EmbedPage 路线（见案例文档三路线表），不得用本资产直接交付。
- 外部数据钩子：`PRIMARY3D.setCircuitValue(柜ID, 回路ID, {Ia, P, Energy})` + `setMode('live')`；URL 参数 `?shot=1`（headless 截图机位）、`?sel=<柜ID>`（自动选中）。

## 来源与验证记录

模板与解析器来自 2026-09 实战（源图 1.dwg：1# 变压器 SCB14-800kVA 低压系统，GCS 柜
T1D3~T1D9），经 LibreDWG 解析 → 模板生成 → Edge headless 截图端到端验证后沉淀；
沉淀时已剥离环境特定信息（租户/项目/地址），保留通用电气参数作示例。
