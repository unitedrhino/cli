# 画布 JSON 结构与组件元素骨架

> 本文件是 [SKILL.md](../SKILL.md) 的子文档：画布 content 顶层结构、组件元素字段与 componentList 元素骨架示例（自 SKILL.md 拆分整理，内容未改动）。

## content 顶层结构

| 字段 | 说明 |
|------|------|
| `editCanvasConfig` | 画布配置：`width` / `height` / `background` / `backgroundImage` / `chartThemeColor` / `previewScaleType` / `projectName` |
| `componentList` | 组件元素数组（见下节结构速览） |
| `requestGlobalConfig` | 全局请求配置（缺失时 validate 仅给 warning，组件用默认配置） |

## 组件元素字段

| 字段 | 说明 |
|------|------|
| `id` | 组件实例 ID，画布内全局唯一 |
| `title` | 组件标题（显示名） |
| `chartKey` / `conKey` | 渲染/配置面板注册键，约定 `chartKey="V"+key`、`conKey="VC"+key` |
| `chartConfig` | 组件元数据：`key` / `chartKey` / `conKey` / `title` / `category` / `package` / `chartFrame` |
| `attr` | 布局：`x` / `y` / `w` / `h` / `zIndex` |
| `option` | 组件专有配置（样式 + 静态 `dataset`） |
| `styles` | 通用样式 |
| `request` | 数据绑定：`requestDataType` + 各类型对应配置（IoT 见 [iot-data-binding.md](iot-data-binding.md)） |
| `status` | `{lock, hide}` 锁定与隐藏 |
| `events` | `{baseEvent, advancedEvents, interactEvents}` 事件三件套 |
| `isGroup` / `groupList` | 分组组件标记与子组件列表（分组无独立 key 体系） |

## 组件结构速览

可直接抄改的 `componentList` 元素骨架。示例一：静态柱状图（BarCommon，requestDataType=0）：

```json
{
  "id": "comp-bar-001",
  "title": "产量对比",
  "chartKey": "VBarCommon",
  "conKey": "VCBarCommon",
  "chartConfig": {
    "key": "BarCommon",
    "chartKey": "VBarCommon",
    "conKey": "VCBarCommon",
    "title": "柱状图",
    "category": "Charts",
    "categoryName": "图表",
    "package": "Charts",
    "chartFrame": "echarts",
    "image": "bar.png"
  },
  "attr": { "x": 60, "y": 120, "w": 500, "h": 300, "zIndex": 1 },
  "option": {
    "dataset": {
      "dimensions": ["product", "data1"],
      "source": [
        { "product": "一车间", "data1": 120 },
        { "product": "二车间", "data1": 200 }
      ]
    }
  },
  "styles": {},
  "request": { "requestDataType": 0 },
  "status": { "lock": false, "hide": false },
  "events": { "baseEvent": {}, "advancedEvents": {}, "interactEvents": [] }
}
```

示例二：IoT 绑定表盘（Dial，requestDataType=5，单值自渲染）：

```json
{
  "id": "comp-dial-001",
  "title": "实时温度",
  "chartKey": "VDial",
  "conKey": "VCDial",
  "chartConfig": {
    "key": "Dial",
    "chartKey": "VDial",
    "conKey": "VCDial",
    "title": "表盘",
    "category": "Charts",
    "categoryName": "图表",
    "package": "Charts",
    "chartFrame": "common",
    "image": "dial.png"
  },
  "attr": { "x": 640, "y": 120, "w": 300, "h": 300, "zIndex": 2 },
  "option": { "dataset": 66 },
  "styles": {},
  "request": {
    "requestDataType": 5,
    "requestIoTDeviceConfig": {
      "queryType": "property",
      "mode": "latest",
      "productID": "p_temp_001",
      "deviceNames": ["sensor-001"],
      "dataIDs": ["CurrentTemperature"]
    }
  },
  "status": { "lock": false, "hide": false },
  "events": { "baseEvent": {}, "advancedEvents": {}, "interactEvents": [] }
}
```

分组组件：`"isGroup": true` + `"groupList": [...]`，无独立 chartKey/conKey，validate 会递归校验子组件。
