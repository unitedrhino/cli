# IoT 数据配置与常见配方

> 本文件是 [SKILL.md](../SKILL.md) 的子文档：requestDataType 数据绑定说明、queryType 支持矩阵摘要与常见绑定配方（自 SKILL.md 拆分整理，内容未改动；完整字段表见 [components-iot-config.md](components-iot-config.md) 附录 C）。

## IoT 数据配置手册

组件数据绑定由 `request.requestDataType` 决定：

| requestDataType | 类型 | 说明 |
|-----------------|------|------|
| 0 | 静态 | 用 `option.dataset` 内置数据，不发请求 |
| 1 | AJAX | 走 `request.requestUrl` 等 HTTP 配置 |
| 2 | Pond | 公共数据池 |
| 5 | IoT 设备数据 | 走 `request.requestIoTDeviceConfig`，接设备属性/状态/信息 |

`requestDataType=5` 时 `requestIoTDeviceConfig` 必填，核心字段：`queryType`（`property` / `deviceStatus` / `deviceInfo`）、`productID`、`deviceNames`、`dataIDs`、`mode`（`latest` / `log`）。**注意**：`deviceStatus` 按产品维度统计在线状态，**不需要** `deviceNames`/`dataIDs`（真实画布即只填 `productID`，可留空数组）；`property`/`deviceInfo` 必须填非空 `deviceNames` 与 `dataIDs`（validate 会按此规则校验）。完整字段表见 [components-iot-config.md](components-iot-config.md) 附录 C。

### queryType 支持矩阵摘要（以 [components-iot-config.md](components-iot-config.md) 附录 D 真实矩阵为准）

| 组件分组 | property | deviceStatus | deviceInfo |
|----------|----------|--------------|------------|
| 折线组（LineCommon 等） | ✓（仅 log） | ✗ | ✗ |
| 饼图组（PieCircle / PieCommon） | ✓（仅 latest） | ✓ | ✗ |
| 柱状组（BarCommon / BarCrossrange / BarLine） | ✓ | ✓ | ✗ |
| 文字组（TextCommon / FlipperNumber） | ✓ | ✓ | ✗ |
| 表格-信息组（TablesBasic / TableList） | ✓ | ✓ | ✓ |
| 表格-滚动组（TableScrollBoard） | ✓ | ✓ | ✗ |
| 雷达组（Radar） | ✓ | ✓ | ✗ |
| 进度组（Process / WaterPolo） | ✓（仅 latest） | ✓ | ✗ |
| 胶囊组（CapsuleChart） | ✓（仅 latest） | ✓ | ✓ |
| 其他可绑定数据组件 | ✓ | ✓ | ✓ |
| 静态/装饰组件 | ✗ | ✗ | ✗ |

单值自渲染组件（Dial / PieCircle / Process / WaterPolo）拿到 IoT 单值后自行渲染，不走 ECharts dataset 覆盖。

### 常见配方

**① 单值实时展示 → Dial / PieCircle / Process / WaterPolo**

`queryType=property` + `mode=latest`，单设备单属性取最新值，组件自渲染：

```json
"request": {
  "requestDataType": 5,
  "requestIoTDeviceConfig": {
    "queryType": "property",
    "mode": "latest",
    "productID": "p_temp_001",
    "deviceNames": ["sensor-001"],
    "dataIDs": ["CurrentTemperature"]
  }
}
```

**② 多设备同属性对比 → BarCommon（或 CapsuleChart）**

`queryType=property`，`deviceNames` 填多台设备、`dataIDs` 填同一属性：

```json
"requestIoTDeviceConfig": {
  "queryType": "property",
  "productID": "p_temp_001",
  "deviceNames": ["sensor-001", "sensor-002", "sensor-003"],
  "dataIDs": ["CurrentTemperature"]
}
```

**③ 趋势曲线 → LineCommon（折线组仅支持 log 模式）**

`queryType=property` + `mode=log`，配 `logTimeRange` / `logInterval`：

```json
"requestIoTDeviceConfig": {
  "queryType": "property",
  "mode": "log",
  "logTimeRange": "24h",
  "logInterval": "1h",
  "productID": "p_temp_001",
  "deviceNames": ["sensor-001"],
  "dataIDs": ["CurrentTemperature"]
}
```

**④ 设备在线状态占比 → PieCircle / PieCommon（饼图组）**

`queryType=deviceStatus`，统计所选设备在线/离线占比：

```json
"requestIoTDeviceConfig": {
  "queryType": "deviceStatus",
  "productID": "p_temp_001",
  "deviceNames": ["sensor-001", "sensor-002"],
  "dataIDs": ["CurrentTemperature"]
}
```

**⑤ 设备信息卡片/列表 → TableList / TablesBasic（表格-信息组）**

`queryType=deviceInfo` 展示设备信息字段。**注意**：`TextCommon` 的 `getQueryTypeHint` 提示语声称支持设备信息，但 `isQueryTypeSupported` 实际判定文字组**禁用 deviceInfo**——不要用 TextCommon 绑设备信息，以运行时矩阵为准。

**⑥ 大数字单值 → Number / CountDown（装饰组，queryType 全放行）**

`queryType=property` + `mode=latest`，适合做大字号的单值看板。
