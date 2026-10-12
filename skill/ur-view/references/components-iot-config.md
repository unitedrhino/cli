# IoT 数据绑定字段表与 queryType 支持矩阵

> 本文件是 [components.md](components.md) 的子文档：附录 C（RequestIoTDeviceConfigType 字段表）、附录 D（IoT queryType 支持矩阵）及补充说明（快照日期 2026-08-03，内容未改动）。

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
