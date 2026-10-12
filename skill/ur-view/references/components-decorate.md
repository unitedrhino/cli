# Decorates 装饰 / Presets 预置组件

> 本文件是 [components.md](components.md) 的子文档：Decorates 装饰（30）与 Presets 预置（6）类别清单（快照日期 2026-08-03，内容未改动）。

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

## Presets 预置（6）

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| BorderRunningTimer | VBorderRunningTimer | 累计运行时间 | composite | — | 边框容器 + 设备累计运行时间 | 经典数据面板限定 静态/AJAX，走专用数据配置，不走通用 IoT 面板；复合预置组件，redirectComponent → Decorates/BorderCharts/BorderRunningTimer |
| BorderMetricFlipper | VBorderMetricFlipper | 指标翻牌器 | composite | property, deviceStatus, deviceInfo | 边框容器 + 指标翻牌器 | 复合预置组件，redirectComponent → Decorates/BorderCharts/BorderMetricFlipper |
| BorderDeviceStatus | VBorderDeviceStatus | 设备状态 | composite | — | 边框容器 + 设备状态统计 | 经典数据面板限定 静态/AJAX，走专用数据配置，不走通用 IoT 面板；复合预置组件，redirectComponent → Decorates/BorderCharts/BorderDeviceStatus |
| BorderDualChart | VBorderDualChart | 设备对比分析 | composite | property, deviceStatus, deviceInfo | 边框容器 + 双设备对比图表 | 复合预置组件，redirectComponent → Decorates/BorderCharts/BorderDualChart |
| BorderBarChart | VBorderChartBase | 通用图表 | composite | property, deviceStatus, deviceInfo | 边框容器 + 通用柱状图 | 复合预置组件，redirectComponent → Decorates/BorderCharts/BorderBarChart |
| BorderContainer | group | 空白边框容器 | common | — | 空白边框分组容器 | 分组容器（chartKey=group），不参与数据绑定；复合预置组件，redirectComponent → Presets/Container/BorderContainer |
