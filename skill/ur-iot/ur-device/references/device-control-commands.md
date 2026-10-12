# CLI 命令 — 属性控制与行为：命令参考

本文件是 device-control.md 的子文档，收录 `ur things device` 各子命令的参数说明、使用示例与对应 API，以及注意事项。

## control — 发送属性控制命令

向设备发送属性控制指令；是否写影子由所选控制模式决定。云端模拟使用 `--shadow-control 4`。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 是 | string | 设备名称 |
| --data | | 是 | JSON | 属性键值对 JSON |
| --shadow-control | | 否 | int | 控制模式；云端模拟固定使用 4 |
| --project-id | | 否 | string | 项目 ID，默认读取 `UR_PROJECT_ID` |
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

根据产品物模型或设备合并物模型生成符合数据类型约束的内容。此命令只生成数据，不写平台、不产生上报、不控制设备。默认生成全部属性、1 份并使用 JSON 输出；返回的 `data.items` 可直接传给后续代码。

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --product-id | -p | 是 | string | 产品ID |
| --device-name | -d | 否 | string | 设备名称；省略时使用产品物模型，传入时使用设备合并物模型 |
| --type | -t | 否 | string | `property/event/action` 或 `1/2/3`，默认 `property` |
| --data-id | | 否 | string | 可重复；省略表示生成全部相关内容 |
| --num | -n | 否 | int | 生成数量，1～100（默认1） |
| --project-id | | 否 | string | 项目 ID，默认读取 `UR_PROJECT_ID` |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

#### 示例1：产品级全部属性
```bash
ur things device mock -p 2D -j
```

#### 示例2：设备级全部属性
```bash
ur things device mock -p 2D -d yanshi-dev01 -j
```

#### 示例3：指定属性
```bash
ur things device mock -p 2D -d yanshi-dev01 --data-id temperature -j
```

#### 示例4：生成多份
```bash
ur things device mock -p 2D -d yanshi-dev01 --data-id temperature --num 5 -j
```

#### 示例5：全部上行事件
```bash
ur things device mock -p 2D --type event -j
```

#### 示例6：全部下行行为
```bash
ur things device mock -p 2D --type action -j
```

命令会把后端 `data.params` 解析为数组，例如 `data.items[0]` 是第一份完整样例；生成多份时任一请求失败则整体失败，不输出残缺结果。

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

1. **纯 Mock 无需先查物模型**：已知产品或设备后直接调用；解释字段、限定范围或排查失败时再查询
2. **data 字段 key 必须与查询到的物模型 identifier 完全一致**，不可假设大小写风格
3. **离线设备控制**：是否缓存影子取决于控制模式；仅云端改值的模式 4 不会在上线后补发实体控制
4. **传统 report 命令自动认证**：会获取设备密钥生成 MQTT 凭据；此行为不适用于上方规则引擎任务的模拟上报路径
5. **upload 输出 curl 命令**：实际文件上传需使用输出的 curl 命令执行
