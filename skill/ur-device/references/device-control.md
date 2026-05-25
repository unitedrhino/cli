# CLI 命令 — 属性控制与行为

## 命令语法

```
ur things device <subcommand> [选项]
```

子命令：`control`、`action`、`mock`、`report`、`upload`

---

## control — 发送属性控制命令

向设备发送属性控制指令，同时更新影子设备。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --data | | 是 | JSON | 属性键值对 JSON |
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
  --data '{"PowerSwitch": 1, "Brightness": 80}'
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
| --data-id | | 是 | string | 行为标识符 |
| --input | | 否 | JSON | 行为输入参数 JSON |
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

当设备主动调用云端行为（上行）时，云端回复结果。

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

根据设备物模型自动生成符合数据类型约束的 Mock 数据。

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

通过 HTTP 模拟设备上报消息，自动获取设备密钥并生成 MQTT 认证凭据。

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

## upload — 设备文件上传

生成设备文件上传的 curl 命令，自动获取设备密钥。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --file | -f | 是 | string | 要上传的文件路径 |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur things device upload -p p_smartswitch_001 -d switch-001 -f /path/to/file.txt
```

### 对应API

- `POST /api/v1/things/device/edge/upload-file`

---

## 注意事项

1. **物模型必须先查**：控制属性或调用行为前，务必确认 `data-id` 正确，大小写敏感
2. **data 字段 key 必须与物模型 identifier 完全一致**（大驼峰命名）
3. **离线设备控制**：属性控制指令会缓存到影子设备，设备上线后自动下发
4. **report 自动认证**：无需手动提供 username/password，自动获取设备密钥生成 MQTT 凭据
5. **upload 输出 curl 命令**：实际文件上传需使用输出的 curl 命令执行
