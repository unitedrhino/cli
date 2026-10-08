# CLI 命令 — 实时调试详解

## 命令语法

```
ur things device <subcommand> [选项]
```

子命令：`control`、`action`、`mock`、`report`

---

## control — 发送属性控制命令

向设备发送属性控制指令，同时更新影子设备。离线设备上线后会自动同步。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称（设备唯一ID） |
| --data | | 是 | JSON | 属性键值对 JSON，key 必须与物模型 identifier 完全一致 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：控制设备开关
```bash
ur things device control -p p_smartswitch_001 -d switch-001 --data '{"PowerSwitch": 1}'
```

#### 示例2：控制设备亮度
```bash
ur things device control -p p_smartswitch_001 -d switch-001 --data '{"Brightness": 80}'
```

#### 示例3：同时控制多个属性
```bash
ur things device control -p p_smartswitch_001 -d switch-001 \
  --data '{"PowerSwitch": 1, "Brightness": 80, "Color": "#FF0000"}'
```

### 对应API

- `POST /api/v1/things/device/interact/property-control-send`

---

## action — 设备行为调用

子命令：`send`、`get`、`resp`

### action send — 调用设备行为

向设备发送行为调用指令。

#### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --data-id | | 是 | string | 行为标识符（来自物模型 Action.identifier） |
| --input | | 否 | JSON | 行为输入参数 JSON，key 必须与物模型 input[].identifier 一致 |
| --json | -j | 否 | bool | 输出JSON格式 |

#### 使用示例

```bash
ur things device action send -p p_smartswitch_001 -d switch-001 \
  --data-id OpenValve --input '{"Duration": 30}'
```

#### 对应API

- `POST /api/v1/things/device/interact/action-send`

---

### action get — 获取行为执行结果

查询设备行为的执行结果。

#### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --data-id | | 是 | string | 行为标识符 |
| --json | -j | 否 | bool | 输出JSON格式 |

#### 使用示例

```bash
ur things device action get -p p_smartswitch_001 -d switch-001 --data-id OpenValve
```

#### 对应API

- `POST /api/v1/things/device/interact/action-get-one`

---

### action resp — 回复设备行为调用

当设备主动调用云端行为（上行，`dir: up`）时，云端回复结果。

#### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --data-id | | 是 | string | 行为标识符 |
| --output | | 否 | JSON | 行为输出参数 JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

#### 使用示例

```bash
ur things device action resp -p p_smartswitch_001 -d switch-001 \
  --data-id ReadMeter --output '{"EP": "1234.56"}'
```

#### 对应API

- `POST /api/v1/things/device/interact/action-resp`

---

## mock — 生成 Mock 数据

根据设备物模型自动生成符合数据类型约束的 Mock 数据，用于调试时快速生成合法测试数据。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --data-id | | 是 | string | 属性/行为/事件标识符 |
| --num | | 否 | int | 生成数量（默认1） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：生成1条温度Mock数据
```bash
ur things device mock -p p_smartswitch_001 -d switch-001 --data-id Temperature
```

#### 示例2：生成5条Mock数据
```bash
ur things device mock -p p_smartswitch_001 -d switch-001 --data-id Temperature --num 5
```

### 对应API

- `POST /api/v1/things/device/interact/schema-mock-gen`

---

## report — 模拟设备上报

通过 HTTP 模拟设备上报消息。自动获取设备密钥并生成 MQTT 认证凭据，无需手动提供 username/password。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --handle | | 否 | string | 处理类型：`thing`/`ota`/`config`（默认thing） |
| --type | | 否 | string | 消息类型：`property`/`event`/`action`（默认property） |
| --method | | 否 | string | 方法：`report`（默认report） |
| --params | | 是 | JSON | 上报参数 JSON |
| --msg-token | | 否 | string | 消息令牌 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：上报设备属性
```bash
ur things device report -p p_smartswitch_001 -d switch-001 \
  --params '{"Temperature": 25.3, "Humidity": 60}'
```

#### 示例2：上报设备事件
```bash
ur things device report -p p_smartswitch_001 -d switch-001 \
  --type event --params '{"PowerAlarm": {"Voltage": 220}}'
```

#### 示例3：上报设备行为
```bash
ur things device report -p p_smartswitch_001 -d switch-001 \
  --type action --params '{"ReadMeter": {"EP": "1234.56"}}'
```

### 对应API

- `POST /api/v1/things/device/edge/send/{handle}/{type}`

---

## event-send — 向设备发送下行事件通知

目前暂无 Layer 2 CLI 命令，需使用 `ur api` 直接调用。

### 使用示例

```bash
ur api /api/v1/things/device/interact/event-send \
  --body '{
    "productID": "p_smartswitch_001",
    "deviceName": "switch-001",
    "dataID": "ConfigUpdate",
    "params": {"interval": 300}
  }'
```

### 对应API

- `POST /api/v1/things/device/interact/event-send`

---

## 调试前必备：查询物模型

发送任何调试指令前，必须先查询设备物模型确认可用属性/行为/事件：

```bash
# 设备级物模型
ur things schema get-list -p xxx -d yyy

# 或产品级物模型
ur things schema get-list -p xxx
```

返回的 `identifier` 即为调试时使用的 `data-id`。

---

## 典型调试工作流

### 工作流 1：测试属性控制

```bash
# 1. 查询物模型，确认控制属性
ur things schema get-list -p xxx -d yyy

# 2. 查看当前属性值
ur things device log property -p xxx -d yyy --data-ids PowerSwitch

# 3. 发送控制指令
ur things device control -p xxx -d yyy --data '{"PowerSwitch":1}'

# 4. 查看命令日志确认执行结果
ur things device log send -p xxx -d yyy --actions propertyControlSend
```

### 工作流 2：测试行为调用

```bash
# 1. 查询物模型，确认行为定义
ur things schema get-list -p xxx -d yyy

# 2. 调用设备行为
ur things device action send -p xxx -d yyy --data-id OpenValve --input '{"Duration":30}'

# 3. 获取行为执行结果
ur things device action get -p xxx -d yyy --data-id OpenValve
```

### 工作流 3：模拟设备上报

```bash
# 1. 查询物模型确认属性标识符
ur things schema get-list -p xxx -d yyy

# 2. 生成Mock数据验证格式
ur things device mock -p xxx -d yyy --data-id Temperature --num 1

# 3. 模拟设备上报
ur things device report -p xxx -d yyy --params '{"Temperature":25.3}'

# 4. 查看属性日志确认上报成功
ur things device log property -p xxx -d yyy --data-id Temperature
```

---

## 注意事项

1. **物模型必须先查**：调试前务必确认 `data-id` 正确，大小写敏感
2. **data 字段 key 必须与物模型 identifier 完全一致**（大驼峰命名）
3. **离线设备控制**：属性控制指令会缓存到影子设备，设备上线后自动下发
4. **report 自动认证**：无需手动提供 username/password，自动获取设备密钥生成 MQTT 凭据
5. **影子设备控制模式**：默认指令下发到设备同时更新影子设备；离线时指令缓存到影子设备
