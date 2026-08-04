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

## Charts 图表（23）

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| BarCommon | VBarCommon | 柱状图 | echarts | property, deviceStatus | 分类数值对比柱状图 | — |
| BarCrossrange | VBarCrossrange | 横向柱状图 | echarts | property, deviceStatus | 横向条形排名对比 | — |
| CapsuleChart | VCapsuleChart | 胶囊柱图 | common | property, deviceStatus, deviceInfo | 多设备同属性最新值胶囊柱对比 | — |
| BarLine | VBarLine | 柱状图 & 折线图 | echarts | property, deviceStatus | 柱线混合双轴对比 | — |
| PieCommon | VPieCommon | 饼图 | echarts | property, deviceStatus | 组成占比饼图 | — |
| PieCircle | VPieCircle | 饼图-环形 | echarts | property, deviceStatus | 环形占比 / 单值百分比 | 单值自渲染，IoT 数据不走 ECharts dataset 覆盖 |
| Process | VProcess | NaiveUI-进度 | （未声明） | property, deviceStatus | 单值进度条 | 未声明 chartFrame，按 dataset 判定数据绑定；单值自渲染，IoT 数据不走 ECharts dataset 覆盖 |
| WaterPolo | VWaterPolo | 水球图 | common | property, deviceStatus | 单值水位球 | 单值自渲染，IoT 数据不走 ECharts dataset 覆盖 |
| LineCommon | VLineCommon | 折线图 | echarts | property | 属性历史时序趋势 | — |
| LineLinearSingle | VLineLinearSingle | 单折线渐变图 | echarts | property | 单系列线性趋势 | — |
| LineGradientSingle | VLineGradientSingle | 单折线渐变面积图 | echarts | property | 单系列渐变趋势 | — |
| LineGradients | VLineGradients | 双折线渐变面积图 | echarts | property | 渐变面积趋势图 | — |
| ScatterCommon | VScatterCommon | 散点图 | echarts | property, deviceStatus, deviceInfo | 双数值散点分布 | — |
| ScatterLogarithmicRegression | VScatterLogarithmicRegression | 对数回归散点图 | echarts | property, deviceStatus, deviceInfo | 对数回归散点 | — |
| Radar | VRadar | 雷达图 | common | property, deviceStatus | 多维指标雷达对比 | — |
| Funnel | VFunnel | 漏斗图 | echarts | property, deviceStatus, deviceInfo | 流程转化漏斗 | — |
| Heatmap | VHeatmap | 热力图 | common | property, deviceStatus, deviceInfo | 密度 / 矩阵热力图 | — |
| TreeMap | VTreeMap | 树形分布 | common | property, deviceStatus, deviceInfo | 层级占比矩形树图 | — |
| Graph | VGraph | 关系图 | common | property, deviceStatus, deviceInfo | 关系拓扑图 | — |
| Sankey | VSankey | 桑基图 | common | property, deviceStatus, deviceInfo | 流量去向桑基图 | — |
| Dial | VDial | 表盘 | common | property, deviceStatus, deviceInfo | 单值仪表盘 | 单值自渲染，IoT 数据不走 ECharts dataset 覆盖 |
| MapBase | VMapBase | 地图(可选省份) | common | property, deviceStatus, deviceInfo | 地理区域分布图 | — |
| MapAmap | VMapAmap | 高德地图 | common | property, deviceStatus, deviceInfo | 高德地图点位展示 | — |

## Informations 信息（13）

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| TextCommon | VTextCommon | 文字 | common | property, deviceStatus | 单值文本 / 单项信息展示 | — |
| TextBarrage | VTextBarrage | 弹幕文字 | common | property, deviceStatus, deviceInfo | 滚动弹幕信息 | — |
| TextGradient | VTextGradient | 渐变文字 | naiveUI | property, deviceStatus, deviceInfo | 渐变标题文字 | — |
| InputsDate | VInputsDate | 时间选择器 | static | — | 日期筛选控件 | 无数据绑定 |
| InputsSelect | VInputsSelect | 下拉选择器 | static | — | 下拉筛选控件 | 无数据绑定 |
| InputsTab | VInputsTab | 标签选择器 | static | — | 页签切换控件 | 无数据绑定 |
| InputsPagination | VInputsPagination | 分页 | static | — | 分页控件 | 无数据绑定 |
| InputsInput | VInputsInput | 输入框 | static | — | 输入框控件 | 无数据绑定 |
| Image | VImage | 图片 | common | property, deviceStatus, deviceInfo | 图片展示 | — |
| ImageCarousel | VImageCarousel | 轮播图 | naiveUI | property, deviceStatus, deviceInfo | 图片轮播 | — |
| Iframe | VIframe | 远程网页 | common | property, deviceStatus, deviceInfo | 内嵌第三方网页 | — |
| Video | VVideo | 视频 | common | property, deviceStatus, deviceInfo | 视频播放 | — |
| WordCloud | VWordCloud | 词云 | common | property, deviceStatus, deviceInfo | 关键词词云 | — |

## Tables 表格（3）

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| TableList | VTableList | 滚动排名列表 | common | property, deviceStatus, deviceInfo | 设备信息滚动排行列表 | — |
| TableScrollBoard | VTableScrollBoard | 轮播列表 | common | property, deviceStatus | 滚动数据表格 | — |
| TablesBasic | VTablesBasic | 基础分页表格 | common | property, deviceStatus, deviceInfo | 基础表格 / 多设备字段对比 | — |

## Decorates 装饰（30）

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| Border01 | VBorder01 | 边框-01 | static | — | 纯装饰组件 | 无数据绑定 |
| Border02 | VBorder02 | 边框-02 | static | — | 纯装饰组件 | 无数据绑定 |
| Border03 | VBorder03 | 边框-03 | static | — | 纯装饰组件 | 无数据绑定 |
| Border04 | VBorder04 | 边框-04 | static | — | 纯装饰组件 | 无数据绑定 |
| Border05 | VBorder05 | 边框-05 | static | — | 纯装饰组件 | 无数据绑定 |
| Border06 | VBorder06 | 边框-06 | static | — | 纯装饰组件 | 无数据绑定 |
| Border07 | VBorder07 | 边框-07 | static | — | 纯装饰组件 | 无数据绑定 |
| Border08 | VBorder08 | 边框-08 | static | — | 纯装饰组件 | 无数据绑定 |
| Border09 | VBorder09 | 边框-09 | static | — | 纯装饰组件 | 无数据绑定 |
| Border10 | VBorder10 | 边框-10 | static | — | 纯装饰组件 | 无数据绑定 |
| Border11 | VBorder11 | 边框-11 | static | — | 纯装饰组件 | 无数据绑定 |
| Border12 | VBorder12 | 边框-12 | static | — | 纯装饰组件 | 无数据绑定 |
| Border13 | VBorder13 | 边框-13 | static | — | 纯装饰组件 | 无数据绑定 |
| Decorates01 | VDecorates01 | 装饰-01 | static | — | 纯装饰组件 | 无数据绑定 |
| Decorates02 | VDecorates02 | 装饰-02 | static | — | 纯装饰组件 | 无数据绑定 |
| Decorates03 | VDecorates03 | 装饰-03 | static | — | 纯装饰组件 | 无数据绑定 |
| Decorates04 | VDecorates04 | 装饰-04 | static | — | 纯装饰组件 | 无数据绑定 |
| Decorates05 | VDecorates05 | 装饰-05 | static | — | 纯装饰组件 | 无数据绑定 |
| Decorates06 | VDecorates06 | 装饰-06 | static | — | 纯装饰组件 | 无数据绑定 |
| FlowChartLine | VFlowChartLine | 流程线 | static | — | 纯装饰组件 | 无数据绑定 |
| ThreeEarth01 | VThreeEarth01 | 三维地球 | common | property, deviceStatus, deviceInfo | 三维地球展示 | — |
| Number | VNumber | 数字计数 | common | property, deviceStatus, deviceInfo | 大数字单值展示 | — |
| TimeCommon | VTimeCommon | 通用时间 | static | — | 当前时间展示 | 无数据绑定 |
| Clock | VClock | 时钟 | static | — | 时钟展示 | 无数据绑定 |
| FullScreen | VFullScreen | 全屏按钮 | static | — | 全屏切换按钮 | 无数据绑定 |
| CountDown | VCountDown | 倒计时 | common | property, deviceStatus, deviceInfo | 倒计时 | — |
| FlipperNumber | VFlipperNumber | 数字翻牌-需动态触发 | common | property, deviceStatus | 翻牌器单值展示 | — |
| PipelineH | VPipelineH | 管道-横向 | （未声明） | — | 横向管道装饰 | 未声明 chartFrame，按 dataset 判定数据绑定 |
| PipelineV | VPipelineV | 管道-纵向 | （未声明） | — | 纵向管道装饰 | 未声明 chartFrame，按 dataset 判定数据绑定 |
| CirclePoint | VCirclePoint | 圆点光环 | static | — | 圆点装饰 | 无数据绑定 |

## Icons 图标（1）

Icons 分类只有一个真实组件 `Icon`（chartFrame=static）。图标面板中的条目由 `Icons/MaterialLine`、`Icons/Common`、`Icons/Weather` 三个 index.ts 基于 IconConfig 动态生成（共 26 + 68 + 33 = 127 个），条目的 `image`/`dataset` 为图标名（`line-md:`/`uim:`/`wi:` 前缀），`redirectComponent` 统一指向 `Icons/Default/Icon`。

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| Icon | VIcon | 图标 | static | — | 矢量图标（图标库统一入口） | 无数据绑定 |

## Photos 图片（0）

Photos 分类没有静态组件：资源库（Library）条目运行时从资源库 API 加载，本地（Local）条目来自用户上传并持久化在 localStorage，共享（Share）为远程图片列表；三类条目均基于 `Informations/Mores/Image` 的 ImageConfig 生成，`chartFrame=static`，`redirectComponent` 指向 `Informations/Mores/Image`。

## Presets 预置（6）

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| BorderRunningTimer | VBorderRunningTimer | 累计运行时间 | composite | — | 边框容器 + 设备累计运行时间 | 经典数据面板限定 静态/AJAX，走专用数据配置，不走通用 IoT 面板；复合预置组件，redirectComponent → Decorates/BorderCharts/BorderRunningTimer |
| BorderMetricFlipper | VBorderMetricFlipper | 指标翻牌器 | composite | property, deviceStatus, deviceInfo | 边框容器 + 指标翻牌器 | 复合预置组件，redirectComponent → Decorates/BorderCharts/BorderMetricFlipper |
| BorderDeviceStatus | VBorderDeviceStatus | 设备状态 | composite | — | 边框容器 + 设备状态统计 | 经典数据面板限定 静态/AJAX，走专用数据配置，不走通用 IoT 面板；复合预置组件，redirectComponent → Decorates/BorderCharts/BorderDeviceStatus |
| BorderDualChart | VBorderDualChart | 设备对比分析 | composite | property, deviceStatus, deviceInfo | 边框容器 + 双设备对比图表 | 复合预置组件，redirectComponent → Decorates/BorderCharts/BorderDualChart |
| BorderBarChart | VBorderChartBase | 通用图表 | composite | property, deviceStatus, deviceInfo | 边框容器 + 通用柱状图 | 复合预置组件，redirectComponent → Decorates/BorderCharts/BorderBarChart |
| BorderContainer | group | 空白边框容器 | common | — | 空白边框分组容器 | 分组容器（chartKey=group），不参与数据绑定；复合预置组件，redirectComponent → Presets/Container/BorderContainer |

## Interact 交互（1）

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| DeviceInteractButton | VDeviceInteractButton | 设备交互按钮 | static | — | 设备控制交互按钮 | 无数据绑定 |

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

## 附录 C：RequestIoTDeviceConfigType 字段表

来源：`apps/web/packages/bigscreen/src/goview/store/modules/chartEditStore/chartEditStore.type.ts`。

| 字段名 | 类型 | 含义 |
| --- | --- | --- |
| `aggFunc` | `string` | 聚合函数（log 模式单选） |
| `aggFuncs` | `string[]` | 多聚合（agg 模式，后续迭代） |
| `areaIDs` | `string[]` | — |
| `chartMode` | `'aggregate' \| 'ranking' \| 'share' \| 'singleStat' \| 'traditional' \| 'trend'` | — |
| `aggregateViewMode` | `'ranking' \| 'share'` | — |
| `customTimeRange` | `[number, number]` | [startTime, endTime] 毫秒时间戳 |
| `dataID` | `string` | 旧字段（向后兼容） |
| `dataIDNames` | `Record<string, string>` | dataID → 属性中文名 映射（由组件从物模型 schema 填充） |
| `dataIDs` | `string[]` | 属性选择（统一为数组） |
| `dataUnits` | `Record<string, string>` | 属性单位映射（ECharts 轴标题自动填充） |
| `dataView` | `'agg' \| 'device' \| 'prop'` | — |
| `deviceName` | `string` | 旧字段（向后兼容） |
| `deviceNames` | `string[]` | 设备选择（统一为数组） |
| `groupBy` | `'area' \| 'device' \| 'none'` | — |
| `intervalUnit` | `'d' \| 'h' \| 'n' \| 'y'` | — |
| `latestTimeCustomFormat` | `string` | — |
| `latestTimeFormat` | `string` | prop+latest 时间格式化 |
| `limit` | `number` | — |
| `logAggEnabled` | `boolean` | 旧字段（向后兼容，可选） |
| `logAggFunc` | `string` | — |
| `logAggInterval` | `string` | — |
| `logAggregationMode` | `'day' \| 'intraday' \| 'month' \| 'week'` | — |
| `logInterval` | `string` | log 模式聚合间隔（如 '10m', '1h'） |
| `logTimeRange` | `'1h' \| '6h' \| '7d' \| '24h' \| '30d' \| 'custom'` | — |
| `mapping` | `object` | ── BI 映射配置（Phase 3.2 字段映射器） ── |
| `mapping.isUserOverride` | `boolean` | — |
| `mapping.mappingMode` | `'nameValue' \| 'xy'` | Phase 3.2 新字段 |
| `mapping.mode` | `'nameValue' \| 'xy'` | Phase 3.1 旧字段（向后兼容） |
| `mapping.nameField` | `string` | nameValue 模式（饼图/漏斗） |
| `mapping.valueField` | `string` | — |
| `mapping.xDimension` | `'device' \| 'property' \| 'time'` | — |
| `mapping.xField` | `string` | XY 模式（折线/柱状/散点） |
| `mapping.yFields` | `string[]` | — |
| `mapping.yMetrics` | `string[]` | — |
| `mode` | `'latest' \| 'log'` | — |
| `outputView` | `'distribution' \| 'fields'` | 数据视图（系统自动推导，不再由用户手动选择） ；'prop': 单设备·多属性；'device': 多设备·同属性；'agg': 同属性·多聚合 |
| `productCategoryCode` | `string` | — |
| `productCategoryID` | `number \| string` | — |
| `productID` | `string` | 共有 |
| `queryType` | `'deviceInfo' \| 'deviceStatus' \| 'property'` | — |
| `scope` | `'area' \| 'device' \| 'tenant'` | 模板级聚合配置（TASK-045 新增） |
| `selectedFields` | `string[]` | — |
| `sortBy` | `'valueAsc' \| 'valueDesc'` | — |
| `timeRange` | `'custom' \| 'month' \| 'today' \| 'week' \| 'year' \| 'yesterday'` | — |
| `templateAutoField` | `string` | — |

## 附录 D：IoT queryType 支持矩阵（按组件分组）

| 组件分组 | 成员 chartKey | property | deviceStatus | deviceInfo | 备注 |
| --- | --- | --- | --- | --- | --- |
| 折线组 | LineCommon / LineGradients / LineGradientSingle / LineLinearSingle | ✓ | ✗ | ✗ | 仅属性历史时序（log 模式） |
| 饼图组 | PieCircle / PieCommon | ✓ | ✓ | ✗ | 属性仅 latest 快照，适合组成占比 |
| 柱状组 | BarCommon / BarCrossrange / BarLine | ✓ | ✓ | ✗ | 数值分类对比 |
| 文字组 | TextCommon / FlipperNumber | ✓ | ✓ | ✗ | 单值文本 / 翻牌器 |
| 表格-信息组 | TablesBasic / TableList | ✓ | ✓ | ✓ | 支持设备信息字段与状态统计 |
| 表格-滚动组 | TableScrollBoard | ✓ | ✓ | ✗ | 滚动表格 |
| 雷达组 | Radar | ✓ | ✓ | ✗ | 多维数值对比 |
| 进度组 | Process / WaterPolo | ✓ | ✓ | ✗ | 单值比例，属性仅 latest |
| 胶囊组 | CapsuleChart | ✓ | ✓ | ✓ | queryType 默认全放行；属性仅 latest 模式（多设备同属性对比） |
| 其他可绑定数据组件 | 其余 chartFrame 非 static 且有 dataset 的组件 | ✓ | ✓ | ✓ | 源码 isQueryTypeSupported 默认全部放行 |
| 静态/装饰组件 | chartFrame=static 或无 dataset | ✗ | ✗ | ✗ | iotQueryTypes 为空数组 |

补充说明：
- `property` 查询类型在源码中对所有可绑定数据的组件恒为支持；折线组仅支持 log（历史时序）模式，饼图/胶囊/进度组仅支持 latest（最新值）模式。
- `singleValueSelfRender=true` 的组件（VPieCircle / VProcess / VWaterPolo / VDial）拿到 IoT 单值后自行渲染，不走 ECharts dataset 覆盖。
- BorderDeviceStatus / BorderRunningTimer 在经典数据面板中被限定为 静态/AJAX，不开放通用 IoT 面板。
- 注意：源码 `getQueryTypeHint` 中 TextCommon 的提示语称“支持设备信息字段”，但 `isQueryTypeSupported` 实际判定文字组禁用 deviceInfo，两处不一致；本清单以后者（运行时实际判定）为准。
