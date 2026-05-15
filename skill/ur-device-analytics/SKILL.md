---
name: ur-device-analytics
description: "设备数据分析：属性历史查询、趋势分析、聚合统计、报表生成。覆盖任何设备的物模型属性数据查询与分析，能源管理是主要示例场景。triggers: 设备数据分析, 属性历史, 能耗统计, 趋势分析, 用能报表, 电力监控, 数据聚合, 设备报表, 历史数据查询"
metadata:
  hermes:
    tags: [device, analytics, data, trend, report, energy]
---

# ur-device-analytics — 设备数据分析

## 核心原则：物模型驱动

**任何设备数据分析的第一步，必须先查询物模型获取可用属性。**

设备属性标识符（`dataID`）不是固定的，不同产品/设备的物模型定义不同：
- 电表可能有：`TotalEnergy`、`P`、`I`、`U`
- 温湿度传感器可能有：`Temperature`、`Humidity`
- 智能灯可能有：`PowerSwitch`、`Brightness`、`Color`

**正确流程**：
```
1. 查询产品/设备物模型 → 获取可用属性标识符列表
2. 根据业务需求选择目标属性（dataID）
3. 构建数据查询参数 → 调用历史数据 API
4. 处理响应数据 → 趋势/报表/可视化
```

## 物模型查询（第一步）

### 产品级物模型

查询某产品下定义的所有属性、事件、行为：

```bash
ur api /api/v1/things/product/schema/get-list \
  --body '{"productID":"xxx"}'
```

响应中的 `identifier` 字段即为数据查询时使用的 `dataID`。

### 设备级物模型

查询具体设备的物模型（含设备级覆盖）：

```bash
ur api /api/v1/things/device/schema/get-list \
  --body '{"productID":"xxx","deviceName":"yyy"}'
```

### 物模型响应关键字段

| 字段 | 说明 |
|------|------|
| `identifier` | 属性标识符 = 数据查询时的 `dataID` |
| `name` | 属性中文名称 |
| `mode` | `r`(只读) / `rw`(可读可写) |
| `define.type` | 数据类型：bool/int/float/string/enum/struct/array/timestamp |
| `define.unit` | 单位（如 `kW·h`、`℃`、`%`） |

## 数据查询 API

<!-- API_LIST:ur-device-analytics -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/msg/abnormal-log/get-list` | 获取设备异常日志 | admin |
| POST | `/api/v1/things/device/msg/event-log/get-list` | 获取事件历史记录 | admin |
| POST | `/api/v1/things/device/msg/gateway-can-bind/get-list` | 获取网关可以绑定的子设备列表 | admin |
| POST | `/api/v1/things/device/msg/hub-log/get-list` | 获取云端诊断日志 | admin |
| POST | `/api/v1/things/device/msg/property-agg/by-device/get-list` | 弃用 | admin |
| POST | `/api/v1/things/device/msg/property-agg/get-list` | 弃用 | admin |
| POST | `/api/v1/things/device/msg/property-latest-agg/get-list` | 聚合属性最新值 | admin |
| POST | `/api/v1/things/device/msg/property-latest/get-list` | 获取最新属性记录 | admin |
| POST | `/api/v1/things/device/msg/property-log-agg/by-device/get-list` | 聚合属性历史记录,设备维度 | admin |
| POST | `/api/v1/things/device/msg/property-log-agg/get-list` | 聚合属性历史记录 | admin |
| POST | `/api/v1/things/device/msg/property-log-latest/get-list` | 弃用 | admin |
| POST | `/api/v1/things/device/msg/property-log/batch-get-list` | 批量获取单个id属性历史记录 | admin |
| POST | `/api/v1/things/device/msg/property-log/get-list` | 获取单个id属性历史记录 | admin |
| POST | `/api/v1/things/device/msg/sdk-log/get-list` | 获取设备sdk日志 | admin |
| POST | `/api/v1/things/device/msg/send-log/get-list` | 获取设备命令日志 | admin |
| POST | `/api/v1/things/device/msg/shadow/get-list` | 获取设备影子列表 | admin |
| POST | `/api/v1/things/device/msg/status-log/get-list` | 获取设备状态日志 | admin |

<!-- END_API_LIST -->

### API 1：单属性历史查询

```bash
ur api /api/v1/things/device/msg/property-log/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataID": "Temperature",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "interval": 1,
    "intervalUnit": "h",
    "argFunc": "avg",
    "fill": "LINEAR"
  }'
```

### API 2：批量属性历史查询

同时查询多个属性或多个时间段：

```bash
ur api /api/v1/things/device/msg/property-log/batch-get-list \
  --body '{
    "reqs": [
      {
        "productID": "xxx", "deviceName": "yyy",
        "dataID": "TotalEnergyChange",
        "timeStart": "1715500800000", "timeEnd": "1715587199999",
        "interval": 1, "intervalUnit": "h", "argFunc": "sum"
      },
      {
        "productID": "xxx", "deviceName": "yyy",
        "dataID": "P",
        "timeStart": "1715500800000", "timeEnd": "1715587199999",
        "interval": 1, "intervalUnit": "h", "argFunc": "avg"
      }
    ]
  }'
```

响应格式：
```json
{
  "code": 200,
  "data": {
    "lists": [
      [{"timestamp": 1715500800000, "value": "12.5", "dataID": "TotalEnergyChange"}],
      [{"timestamp": 1715500800000, "value": "3.2", "dataID": "P"}]
    ]
  }
}
```

### API 3：聚合查询（多指标）

同时查询多个属性，支持不同聚合函数：

```bash
ur api /api/v1/things/device/msg/property-log-agg/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataIDs": ["P", "TotalEnergyChange"],
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "interval": 1,
    "intervalUnit": "h",
    "argFunc": ["avg", "sum"]
  }'
```

### API 4：最新值聚合查询

获取多个属性的当前最新值：

```bash
ur api /api/v1/things/device/msg/property-latest-agg/get-list \
  --body '{
    "productID": "xxx",
    "deviceName": "yyy",
    "dataIDs": ["Temperature", "Humidity", "P"]
  }'
```

## 查询参数速查

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品ID |
| `deviceName` | string | 是 | 设备名称（设备唯一ID） |
| `dataID` | string | 是 | 属性标识符（来自物模型 identifier） |
| `dataIDs` | string[] | agg时 | 多个属性标识符（仅聚合API） |
| `timeStart` | string | 是 | 开始时间戳（毫秒） |
| `timeEnd` | string | 是 | 结束时间戳（毫秒） |
| `interval` | int | 是 | 间隔数值 |
| `intervalUnit` | string | 是 | `h`(时)/`d`(日)/`n`(月)/`y`(年)/`m`(分)/`w`(周) |
| `argFunc` | string/string[] | 是 | 聚合函数：`sum`/`avg`/`max`/`min`/`first`/`last`/`count` |
| `fill` | string | 否 | 缺失值填充：`LINEAR`(线性插值)/`NONE`/`PREV`/`NULL`，默认 `LINEAR` |
| `productCategoryCode` | string | 否 | 品类过滤，如 `dianbiao` |

### 时间粒度对照

| intervalUnit | 含义 | 典型 interval |
|-------------|------|--------------|
| `m` | 分钟 | 1, 5, 15, 30 |
| `h` | 小时 | 1 |
| `d` | 天 | 1 |
| `w` | 周 | 1 |
| `n` | 月 | 1 |
| `y` | 年 | 1 |

### 聚合函数选择

| argFunc | 适用场景 | 典型属性 |
|---------|---------|----------|
| `sum` | 累计消耗量 | TotalEnergyChange、用水量 |
| `avg` | 平均值趋势 | Temperature、Humidity、P |
| `max` | 峰值检测 | P、I、U、Temperature |
| `min` | 谷值检测 | P、I、U、Temperature |
| `first` | 起始读数 | TotalEnergy（集抄起始值） |
| `last` | 结束读数 | TotalEnergy（集抄结束值） |
| `count` | 上报次数统计 | 任意 |

## 能源管理：典型示例场景

能源管理（`org-energy` 应用）是设备数据分析的最完整实践，覆盖 12+ 个分析页面：

### 页面与 API 映射

| 页面 | 路由 | 核心 API | 分析属性 | 聚合方式 |
|------|------|---------|----------|----------|
| 用能概况 | `/energy/analysis/energy-overview` | batch-get-list ×6 | TotalEnergyChange | sum |
| 能耗趋势 | `/energy/analysis/energy-trend` | get-list | TotalEnergyChange | sum |
| 同比环比 | `/energy/analysis/energy-comparison` | batch-get-list | TotalEnergyChange | sum |
| 能耗排名 | `/energy/analysis/energy-ranking` | batch-get-list | TotalEnergyChange | sum |
| 损耗分析 | `/energy/analysis/energy-loss` | batch-get-list | TotalEnergyChange | sum |
| 用能报表 | `/energy/analysis/energy-report` | batch-get-list | TotalEnergyChange | sum |
| 复费率报表 | `/energy/analysis/time-price-report` | batch-get-list | TotalEnergyRate1~4 | sum |
| 集抄明细 | `/energy/analysis/energy-meter` | batch-get-list | TotalEnergy | first/last |
| 数据监控 | `/energy/collection/data-monitoring` | property-log-agg | P + TotalEnergyChange | avg + sum |
| 电力数据 | `/energy/collection/electricity-data` | property-log-agg | P/I/U/TotalEnergy | 多函数 |
| 电力报表 | `/energy/collection/electricity-report` | property-log-agg | P/Q/Ua~c/Ia~c | 快照 |
| 极值报表 | `/energy/collection/extreme-value-report` | property-log-agg | P/Q/Ua~c/Ia~c | max/min |

### 电表类常见属性（示例，非规则）

| 属性标识符 | 含义 | 单位 | 适用分析 |
|-----------|------|------|----------|
| `TotalEnergy` | 总电能（累计读数） | kW·h | 集抄起始/结束值 |
| `TotalEnergyChange` | 电能变化量（消耗） | kW·h | 能耗趋势、排名、报表 |
| `TotalEnergyRate1` | 尖费率电能 | kW·h | 复费率-尖 |
| `TotalEnergyRate2` | 峰费率电能 | kW·h | 复费率-峰 |
| `TotalEnergyRate3` | 平费率电能 | kW·h | 复费率-平 |
| `TotalEnergyRate4` | 谷费率电能 | kW·h | 复费率-谷 |
| `P` | 有功功率 | kW | 功率趋势、极值 |
| `Q` | 无功功率 | kVar | 电力报表 |
| `I` / `Ia` / `Ib` / `Ic` | 电流 | A | 电力数据、极值 |
| `U` / `Ua` / `Ub` / `Uc` | 电压 | V | 电力数据、极值 |

> **重要**：上表属性仅适用于电表类产品（`productCategoryCode: dianbiao`）。分析其他类型设备时，**必须先查询该设备的物模型**获取真实 `identifier`。

### 能源管理树形选择器

能源分析页面使用树形选择器过滤设备范围：

- **区域模式**：`partitionBy: 'areaID'`，按区域层级筛选
- **建筑分组模式**：`partitionBy: 'groupBuildingIDs'`，按建筑分组筛选
- 设备列表过滤：`isOnlyCore: true`（只显示核心设备）

## 典型工作流

### 工作流 1：查询设备某属性的历史趋势

```bash
# 1. 查询物模型，确认可用属性
ur api /api/v1/things/product/schema/get-list \
  --body '{"productID":"p_thermometer_01"}' \
  --fields data.list

# 2. 查询历史数据（以 Temperature 为例）
ur api /api/v1/things/device/msg/property-log/get-list \
  --body '{
    "productID": "p_thermometer_01",
    "deviceName": "room-101",
    "dataID": "Temperature",
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "interval": 1,
    "intervalUnit": "h",
    "argFunc": "avg",
    "fill": "LINEAR"
  }'
```

### 工作流 2：批量对比多个设备的能耗

```bash
# 批量查询多个设备的 TotalEnergyChange
ur api /api/v1/things/device/msg/property-log/batch-get-list \
  --body '{
    "reqs": [
      {"productID":"p_meter_01","deviceName":"floor1-a","dataID":"TotalEnergyChange","timeStart":"...","timeEnd":"...","interval":1,"intervalUnit":"d","argFunc":"sum"},
      {"productID":"p_meter_01","deviceName":"floor1-b","dataID":"TotalEnergyChange","timeStart":"...","timeEnd":"...","interval":1,"intervalUnit":"d","argFunc":"sum"},
      {"productID":"p_meter_01","deviceName":"floor2-a","dataID":"TotalEnergyChange","timeStart":"...","timeEnd":"...","interval":1,"intervalUnit":"d","argFunc":"sum"}
    ]
  }'
```

### 工作流 3：多指标聚合监控面板

```bash
# 同时查询功率和能耗（不同聚合函数）
ur api /api/v1/things/device/msg/property-log-agg/get-list \
  --body '{
    "productID": "p_meter_01",
    "deviceName": "main-meter",
    "dataIDs": ["P", "TotalEnergyChange", "Ia", "Ib", "Ic"],
    "timeStart": "1715500800000",
    "timeEnd": "1715587199999",
    "interval": 1,
    "intervalUnit": "h",
    "argFunc": ["avg", "sum", "max", "max", "max"]
  }'
```

## 数据响应结构

### property-log/get-list 响应

```json
{
  "code": 200,
  "data": {
    "list": [
      {"timestamp": 1715500800000, "value": "25.3", "dataID": "Temperature"},
      {"timestamp": 1715504400000, "value": "25.5", "dataID": "Temperature"}
    ]
  }
}
```

### property-log-agg/get-list 响应

```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "timestamp": 1715500800000,
        "value": {"P": "3.2", "TotalEnergyChange": "12.5", "Ia": "5.1"}
      }
    ]
  }
}
```

### property-latest-agg/get-list 响应

```json
{
  "code": 200,
  "data": {
    "list": [
      {"dataID": "Temperature", "value": "26.1", "timestamp": 1715587200000},
      {"dataID": "Humidity", "value": "60", "timestamp": 1715587200000}
    ]
  }
}
```

## 注意事项

1. **必须先查物模型**：不要假设任何设备有固定属性，始终先调用 schema API 确认
2. **时间戳为毫秒**：`timeStart` / `timeEnd` 是 Unix 毫秒时间戳
3. **value 为字符串**：响应中 `value` 字段是字符串类型，需根据物模型 `define.type` 转换
4. **品类过滤**：能源管理页面统一过滤 `productCategoryCode: dianbiao`，通用分析可不传
5. **批量查询限制**：`batch-get-list` 的 `reqs` 数组长度建议控制在合理范围（如 < 50）
6. **缺失数据处理**：`fill: LINEAR` 适合连续数据（温度、功率），`fill: NONE` 适合离散数据（事件）
7. **累计值 vs 变化量**：`TotalEnergy` 是累计读数（只增不减），`TotalEnergyChange` 是周期消耗量；做趋势分析时用 `TotalEnergyChange`，做集抄读数时用 `TotalEnergy` 的 first/last
