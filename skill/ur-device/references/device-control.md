# CLI 命令 — 属性控制与行为

## 先按用户意图选路径（规则引擎 AI 任务必读）

用户不需要提供接口名或参数。先区分“仅修改云端”“模拟设备上报”和“控制实体设备”，不能把三者当成同一操作：

| 用户目的 | 写入方式 | 验证依据 |
|---|---|---|
| 云端模拟、演示值、只改平台属性且不下发 | `ur api /api/v1/things/device/interact/property-control-send`，显式 `shadowControl=4` | 最新属性 + `propertyControlSend` 操作日志；不承诺标准上报的历史/告警链路 |
| 模拟设备上报，验证历史、规则、告警 | `ur api /api/v1/things/device/simulate/report`，走标准属性上报链路 | 最新属性 + 用户要求的历史/规则/告警结果；可能触发自动化，限已授权测试设备 |
| 真正开灯、设置实体设备属性 | 属性控制接口，根据明确要求选择实时下发或设备影子模式 | 控制响应与设备回执/实际状态；不得用云端改值假装控制成功 |
| 只生成样例、不写平台 | 生成本地值或使用 `mock`，不调用写入接口 | 展示生成结果，不宣称设备值已改变 |

只说“模拟数据”时先询问要仅改云端还是测试上报链路；不能默认上报或向实体下发。用户只要方案/代码时，不得自动试运行。明确云端模拟不要求设备在线；`shadowControl=0/2` 可能在上线后下发，不能替代 4。

### 目标、项目与物模型

1. 从当前任务上下文取得项目 ID，始终保持字符串。项目接口使用 `ur api ... --project-id "$UR_PROJECT_ID"`；先核实 `ur api --help` 支持该参数。旧 CLI 不支持时提示升级，不静默省略项目上下文。
2. 查询当前项目的设备，核对 `productID`、`deviceName`（唯一标识，不是显示别名）及项目归属；查不到或同名多个时询问，不选列表第一项或其他项目设备。未明确授权不得自动新建设备。
3. 查询设备合并后的物模型，按返回的 identifier 原样使用，不能假设必须大驼峰；检查类型、范围、枚举、步长与当前用户权限。物模型 mode 与用户授权不是一回事：实体控制须确认设备支持写入；云端模拟不能仅因传感器属性 mode=r 就改走其他接口或擅自修改物模型，仍由云端接口校验权限与数据。用户范围超出物模型时提示冲突，不能私自扩大。
4. 平台注入任务创建者的 AK/SK、应用和项目环境；继承环境即可。不在代码/对话中写密钥，不构造 UR_TOKEN，不借用其他用户身份。

### 云端模拟的完整请求合同

以下示例仅在查询确认目标设备、`temperature` 标识和数值合法后替换占位符执行：

```bash
ur api /api/v1/things/device/interact/property-control-send \
  --project-id "$UR_PROJECT_ID" \
  --body '{"productID":"<productID>","deviceName":"<deviceName>","data":"{\"temperature\":25.5}","shadowControl":4}'
```

这里 HTTP 请求的 `data` 是 **JSON 字符串**，不是对象。在 JavaScript 中先 `JSON.stringify({ [identifier]: value })` 作为 body.data，再序列化整个 body 传给 `--body`。CLI 的 `--data` 参数是另一层包装，不能据此推断 HTTP data 类型。

控制成功必须同时满足：CLI 退出码 0、响应外层 `code===200`、存在内层 `data.code===200`；内层缺失或失败均报错。再查询最新属性和操作日志核对值与时间，不能只看外层 200。

模式 4 保证此控制调用不直接向设备下发，但仍可能发布平台属性通知/Hook。用户要求完全不影响实体时，应先确认测试设备未绑定会下发的规则/告警联动；不要承诺“不会触发任何平台事件”。

### 模拟上报的不同合同

`/api/v1/things/device/simulate/report` 的 body.data 则是属性值为字符串的对象，例如 `{"projectID":"<projectID>","productID":"<productID>","deviceName":"<deviceName>","data":{"temperature":"25.5"}}`。它返回外层成功码，不能套用控制接口的内层 code 要求。先核实环境已提供该接口；不要为模拟测试索取设备密钥、改用需要设备认证的 edge 上报接口或混用两种 data 类型。用户要求不影响实体时，还必须检查相关规则/告警联动是否会间接下发；无法确认安全时先询问，不先上报再解释副作用。

### 生成 executor.js 时

- 脚本每次只执行一次业务操作，频率/次数交给任务调度，不写无限循环或内部定时器。
- 对于物模型允许的 20～30、步长 0.5，离散随机公式为 `20 + Math.floor(Math.random() * 21) * 0.5`；检查上下界，其他范围按真实物模型计算。
- 用 `Bun.spawn` 的参数数组调用 ur，继承环境，不拼接 shell 命令。并行读取 stdout/stderr，并以 `const exitCode = await proc.exited` 取得最终退出码；禁止读取可能仍为 `null` 的 `proc.exitCode`。
- 项目输入若提供，必须是非空、无首尾空白且与执行上下文一致的字符串，禁止 Number/String 强转掩盖错误；创建者凭证缺失时明确失败。
- 失败必须抛错或 `console.error(...)` 后 `process.exit(1)`；禁止 catch 后在 stdout 输出 `{code:500}` 再正常退出，因为进程 0 会造成假成功。成功才在 stdout 最后一行输出结果 JSON。生成并实际保存 executor.js、manifest.json、skill.md 后才告知完成，不把聊天中的代码块当作已保存产物。

上面的规则适用于云端模拟与任务开发；以下 control/report 命令是其他交互用途，不能替代已经确认的业务路径。

## 命令语法

```
ur things device <subcommand> [选项]
```

子命令：`control`、`action`、`mock`、`report`、`upload`

---

## control — 发送属性控制命令

向设备发送属性控制指令；是否写影子由所选控制模式决定。云端模拟请使用上方明确设置模式 4 的通用接口调用。

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
2. **data 字段 key 必须与查询到的物模型 identifier 完全一致**，不可假设大小写风格
3. **离线设备控制**：是否缓存影子取决于控制模式；仅云端改值的模式 4 不会在上线后补发实体控制
4. **传统 report 命令自动认证**：会获取设备密钥生成 MQTT 凭据；此行为不适用于上方规则引擎任务的模拟上报路径
5. **upload 输出 curl 命令**：实际文件上传需使用输出的 curl 命令执行
