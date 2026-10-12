# Charts 图表组件（23）

> 本文件是 [components.md](components.md) 的子文档：Charts 图表类 23 个组件完整表（快照日期 2026-08-03，内容未改动）。

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
