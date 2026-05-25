# CLI 命令 — 聚合查询（agg）

查询设备属性的聚合值（平均值、最大值、最小值等）。

---

## agg — 属性聚合查询

### 命令语法

```
ur things agg [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称（不指定则查询产品下所有设备） |
| --data-id | -i | 是 | string | 属性标识符（从物模型获取） |
| --funcs | -f | 是 | string | 聚合函数，逗号分隔 |
| --fill |  | 否 | string | 缺失数据填充模式 |
| --no-first-ts |  | 否 | bool | 不填充最早时间戳 |
| --json | -j | 否 | bool | 输出JSON格式 |
| --help | -h | 否 | bool | 显示帮助信息 |

### 支持的聚合函数

| 函数 | 说明 |
|------|------|
| avg | 平均值 |
| first | 首个值 |
| last | 最新值 |
| count | 计数 |
| twa | 时间加权平均 |
| max | 最大值 |
| min | 最小值 |
| sum | 求和 |

### 使用示例

#### 示例1：查询设备平均CPU使用率
```bash
ur things agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg
```

#### 示例2：查询设备温度的最大值和最小值
```bash
ur things agg -p p_smartswitch_001 -d switch-001 -i Temperature -f max,min
```

#### 示例3：查询产品下所有设备的平均温度
```bash
ur things agg -p p_smartswitch_001 -i Temperature -f avg
```

#### 示例4：JSON输出
```bash
ur things agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg -j
```

### 对应API

- `POST /api/v1/things/device/msg/property-latest-agg/get-list`

### 注意事项

- 属性标识符（data-id）必须从物模型获取
- 使用 `ur things schema get-list -p <productID>` 查询物模型
- 使用 `ur things schema get-list -p <productID> -d <deviceName>` 查询设备物模型
