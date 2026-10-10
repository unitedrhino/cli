# CLI 命令 — 聚合查询详解

## 命令语法

```
ur things agg [选项]
```

查询设备属性的聚合值（平均值、最大值、最小值等）。

---

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称（不指定则查询产品下所有设备） |
| --data-id | -i | 是 | string | 属性标识符（来自物模型 identifier） |
| --funcs | -f | 是 | string | 聚合函数，逗号分隔 |
| --fill | | 否 | string | 缺失值填充模式 |
| --no-first-ts | | 否 | bool | 不填充最早时间戳 |
| --json | -j | 否 | bool | 输出JSON格式 |
| --help | -h | 否 | bool | 显示帮助信息 |

---

## 聚合函数说明

| 函数 | 说明 | 适用场景 | 典型属性 |
|------|------|---------|----------|
| `avg` | 平均值 | 平均值趋势 | Temperature、Humidity、P |
| `max` | 最大值 | 峰值检测 | P、I、U、Temperature |
| `min` | 最小值 | 谷值检测 | P、I、U、Temperature |
| `sum` | 求和 | 累计消耗量 | TotalEnergyChange、用水量 |
| `first` | 首个值 | 起始读数 | TotalEnergy（集抄起始值） |
| `last` | 最后值 | 结束读数 | TotalEnergy（集抄结束值） |
| `count` | 计数 | 上报次数统计 | 任意 |
| `twa` | 时间加权平均 | 时间加权趋势 | 功率、温度 |

---

## 使用示例

### 示例1：查询单设备单属性的平均值

```bash
ur things agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg
```

### 示例2：查询单设备多聚合函数

```bash
ur things agg -p p_smartswitch_001 -d switch-001 -i Temperature -f max,min
```

### 示例3：查询产品下所有设备的平均值（不指定设备）

```bash
ur things agg -p p_smartswitch_001 -i Temperature -f avg
```

### 示例4：JSON格式输出

```bash
ur things agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg -j
```

### 示例5：指定填充模式

```bash
ur things agg -p p_smartswitch_001 -d switch-001 -i Temperature -f avg --fill LINEAR
```

---

## 对应API

- `POST /api/v1/things/device/msg/property-latest-agg/get-list`

---

## 请求体结构

```json
{
  "productID": "p_smartswitch_001",
  "deviceName": "switch-001",
  "aggs": [
    {
      "dataID": "Temperature",
      "argFuncs": ["avg", "max"],
      "fill": "LINEAR",
      "noFirstTs": false
    }
  ]
}
```

---

## 响应结构

```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "dataID": "Temperature",
        "value": "26.5",
        "timestamp": 1715587200000
      }
    ]
  }
}
```

---

## 数据分析工作流

### 工作流1：查询设备某属性的聚合值

```bash
# 1. 查询物模型，确认可用属性
ur things schema get-list -p p_thermometer_01

# 2. 查询聚合值
ur things agg -p p_thermometer_01 -d room-101 -i Temperature -f avg
```

### 工作流2：多指标聚合监控

```bash
# 同时查询功率和能耗（不同聚合函数）
ur things agg -p p_meter_01 -d main-meter -i P -f avg
ur things agg -p p_meter_01 -d main-meter -i TotalEnergyChange -f sum
```

---

## 注意事项

1. **必须先查物模型**：属性标识符（`data-id`）必须从物模型获取，不同产品/设备的物模型定义不同
2. **device-name 可选**：不指定 `--device-name` 时查询产品下所有设备的聚合值
3. **funcs 支持多函数**：使用逗号分隔多个聚合函数，如 `-f max,min,avg`
4. **fill 模式**：`LINEAR` 适合连续数据（温度、功率），`NONE` 适合离散数据
5. **value 为字符串**：响应中 `value` 字段是字符串类型，需根据物模型 `define.type` 转换
6. 使用 `ur api /api/v1/things/product/schema/get-list` 查询产品物模型
7. 使用 `ur api /api/v1/things/device/schema/get-list` 查询设备物模型
