# Informations / Tables / Icons / Photos / Interact 组件

> 本文件是 [components.md](components.md) 的子文档：Informations 信息（13）、Tables 表格（3）、Icons 图标（1）、Photos 图片（0）、Interact 交互（1）类别清单（快照日期 2026-08-03，内容未改动）。

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

## Icons 图标（1）

Icons 分类只有一个真实组件 `Icon`（chartFrame=static）。图标面板中的条目由 `Icons/MaterialLine`、`Icons/Common`、`Icons/Weather` 三个 index.ts 基于 IconConfig 动态生成（共 26 + 68 + 33 = 127 个），条目的 `image`/`dataset` 为图标名（`line-md:`/`uim:`/`wi:` 前缀），`redirectComponent` 统一指向 `Icons/Default/Icon`。

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| Icon | VIcon | 图标 | static | — | 矢量图标（图标库统一入口） | 无数据绑定 |

## Photos 图片（0）

Photos 分类没有静态组件：资源库（Library）条目运行时从资源库 API 加载，本地（Local）条目来自用户上传并持久化在 localStorage，共享（Share）为远程图片列表；三类条目均基于 `Informations/Mores/Image` 的 ImageConfig 生成，`chartFrame=static`，`redirectComponent` 指向 `Informations/Mores/Image`。

## Interact 交互（1）

| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |
| --- | --- | --- | --- | --- | --- | --- |
| DeviceInteractButton | VDeviceInteractButton | 设备交互按钮 | static | — | 设备控制交互按钮 | 无数据绑定 |
