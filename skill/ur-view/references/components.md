# GoView 大屏组件清单

> 快照日期：2026-08-03
> 源码路径：`/home/ubuntu/saas/apps/web/packages/bigscreen/src/goview/packages/components`（仓库内相对路径 `apps/web/packages/bigscreen/src/goview/packages/components`）
> 本文件由 `scripts/gen-view-components.py` 自动生成，请勿手工编辑；前端组件变更后重新运行脚本即可刷新。
> 表中 package/category/chartFrame 为画布存档 JSON 实际存储的枚举值字符串（如 `Charts`/`Bars`/`echarts`）。

## 总览

| 分类 | 组件数 | 说明 |
| --- | --- | --- |
| Charts 图表 | 23 | — |
| Informations 信息 | 13 | — |
| Tables 表格 | 3 | — |
| Decorates 装饰 | 30 | — |
| Icons 图标 | 1 | 另有 127 个图标库动态条目（uim:/line-md:/wi: 前缀），统一重定向到 Icon 组件 |
| Photos 图片 | 0 | 无静态组件，资源库/本地上传/共享图片运行时动态生成，重定向到 Informations/Mores/Image |
| Presets 预置 | 6 | 业务复合组件，通过 redirectComponent 指向底层实现 |
| Interact 交互 | 1 | — |
| **合计** | **77** | 静态注册组件总数 |

> 本清单已按类别拆分到子文档（内容同为 2026-08-03 快照，总览与附录保留在本文件）：
> - [components-chart.md](components-chart.md) — Charts 图表（23）完整表
> - [components-info.md](components-info.md) — Informations 信息（13）/ Tables 表格（3）/ Icons 图标（1）/ Photos 图片（0）/ Interact 交互（1）
> - [components-decorate.md](components-decorate.md) — Decorates 装饰（30）/ Presets 预置（6）
> - [components-iot-config.md](components-iot-config.md) — RequestIoTDeviceConfigType 字段表（附录 C）与 IoT queryType 支持矩阵（附录 D）
> 注：本文件由 `scripts/gen-view-components.py` 生成，重新生成后需同步维护上述子文档。

## 附录 A：组件目录结构

每个常规组件目录固定包含 5 个文件：

| 文件 | 职责 |
| --- | --- |
| `index.ts` | ConfigType 元数据（key/chartKey/conKey/title/category/categoryName/package/chartFrame/image） |
| `config.ts` | 默认配置类（attr/option/events 等，option 内含默认 dataset） |
| `index.vue` | 画布渲染组件（按 chartKey 注册） |
| `config.vue` | 右侧配置面板组件（按 conKey 注册） |
| `data.json` | 静态示例数据（静态数据源模式使用） |

例外：Icons/Photos 条目动态生成，无独立目录；Presets 通过 `redirectComponent` 复用底层实现组件的 5 文件。

## 附录 B：前端新增组件步骤（简要）

1. 在对应包/子分类下新建组件目录，补齐 `index.ts` / `config.ts` / `index.vue` / `config.vue` / `data.json` 五个文件。
2. 在子分类 `index.ts`（如 `Charts/Bars/index.ts`）中 import 并加入默认导出数组，完成注册。
3. 新增子分类时同步维护包目录 `index.type.ts` 的 `ChatCategoryEnum` / `ChatCategoryEnumName` 枚举。
4. 需要支持 IoT 数据源时，按组件形态更新 `ChartDataIoTDevice/queryTypeSupport.ts` 中的组件集合（折线/饼/柱/文字/表格/雷达/进度等分组），单值自渲染组件同步更新 `useChartDataFetch.hook.ts` 的 `IOT_SINGLE_VALUE_SELF_RENDER_CHARTS`。
5. 重新运行 `scripts/gen-view-components.py` 刷新 components.json 与本清单。
