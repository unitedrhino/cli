# CLI 命令 — 属性控制与行为

子文档导航：

- [device-control-commands.md](device-control-commands.md) — 命令语法与参数、control/action/mock/report/upload 示例、对应 API、注意事项
- [device-control-executor.md](device-control-executor.md) — 生成 executor.js 时（规则引擎 AI 任务开发规范）

## 先按用户意图选路径（规则引擎 AI 任务必读）

用户不需要提供接口名或参数。先按意图区分“纯生成”“仅修改云端”“模拟设备上报”和“控制实体设备”，不能把四者当成同一操作：

| 用户目的 | 写入方式 | 验证依据 |
|---|---|---|
| 生成、Mock、给一份样例数据，不写平台 | 直接执行 `ur things device mock ... -j` | 只返回符合物模型的数据内容，不修改属性、不产生上报、不控制设备 |
| 云端模拟、演示值、只改平台属性且不下发 | `ur things device control ... --shadow-control 4 --project-id <项目ID>` | 最新属性 + `propertyControlSend` 操作日志；不承诺标准上报的历史/告警链路 |
| 模拟设备上报，验证历史、规则、告警 | `ur things device simulate-report ... --project-id <项目ID>`，走标准属性上报链路 | 最新属性 + 用户要求的历史/规则/告警结果；可能触发自动化，限已授权测试设备 |
| 真正开灯、设置实体设备属性 | 属性控制接口，根据明确要求选择实时下发或设备影子模式 | 控制响应与设备回执/实际状态；不得用云端改值假装控制成功 |

“生成”“Mock”“样例数据”明确表示纯内容生成时没有平台副作用，直接使用 `mock`，不再询问“云端修改还是设备上报”。只有缺少 `productID`，或用户明确要求设备级合并物模型却缺少 `deviceName` 时才补问。用户明确说“修改云端属性”“模拟设备上报”或“控制设备”时，才进入对应写操作路径并执行其安全确认。用户只要方案/代码时，不得自动试运行。明确云端模拟不要求设备在线；`shadowControl=0/2` 可能在上线后下发，不能替代 4。

### 目标、项目与物模型

1. 从当前任务上下文取得项目 ID，始终保持字符串。`mock` 默认读取 `UR_PROJECT_ID`，环境没有时才显式传 `--project-id`；其他设备领域命令按帮助传入项目上下文。旧 CLI 不支持本文命令或参数时明确提示升级，禁止默认退回 `ur api`。
2. 纯 Mock 已知产品或设备标识时直接执行，不预先查询设备列表或完整物模型；`schema-mock-gen` 自身会读取并校验产品物模型或设备合并物模型。只有用户要求解释字段、限定生成范围，或排查生成失败时才查询物模型。
3. 写操作需要定位目标时，使用 `ur things device info get-list --project-id "$UR_PROJECT_ID" -j` 核对 `productID`、`deviceName`（唯一标识，不是显示别名）及项目归属；查不到或同名多个时询问，不选列表第一项或其他项目设备。查询物模型使用 `ur things schema get-list -p '<productID>' --project-id "$UR_PROJECT_ID" -j`，不要写成对应的 `ur api` 路径。按返回的 identifier 原样使用，不能假设大小写风格；检查类型、范围、枚举、步长与当前用户权限。
4. 平台注入任务创建者的 AK/SK、应用和项目环境；继承环境即可。不在代码/对话中写密钥，不构造 UR_TOKEN，不借用其他用户身份。

### 云端模拟的完整请求合同

以下示例仅在查询确认目标设备、`temperature` 标识和数值合法后替换占位符执行：

```bash
ur things device control \
  --product-id '<productID>' \
  --device-name '<deviceName>' \
  --data '{"temperature":25.5}' \
  --shadow-control 4 \
  --project-id "$UR_PROJECT_ID" \
  --json
```

命令的 `--data` 接受 JSON 对象，并负责按后端合同编码成字符串字段；生成代码不要自行拼两层 JSON，也不要改用通用 API 绕过命令校验。

控制成功必须同时满足：CLI 退出码 0、响应外层 `code===200`、存在内层 `data.code===200`；内层缺失或失败均报错。再查询最新属性和操作日志核对值与时间，不能只看外层 200。

模式 4 保证此控制调用不直接向设备下发，但仍可能发布平台属性通知/Hook。用户要求完全不影响实体时，应先确认测试设备未绑定会下发的规则/告警联动；不要承诺“不会触发任何平台事件”。

### 模拟上报的不同合同

使用 `ur things device simulate-report --product-id '<productID>' --device-name '<deviceName>' --data '{"temperature":"25.5"}' --project-id "$UR_PROJECT_ID" --json`。这里 `--data` 的属性值必须是字符串，命令调用 `/api/v1/things/device/simulate/report` 并按外层成功码判断结果，不能套用控制接口的内层 code 要求。先核实 CLI 与环境已提供该命令；不要为模拟测试索取设备密钥、改用需要设备认证的传统 `report` 命令或混用两种 data 类型。用户要求不影响实体时，还必须检查相关规则/告警联动是否会间接下发；无法确认安全时先询问，不先上报再解释副作用。

### 生成 executor.js 时

规则引擎 AI 任务开发的脚本生成规则、调用约定与完成标准见 [device-control-executor.md](device-control-executor.md)。

上面的规则适用于云端模拟与任务开发；以下 control/report 命令是其他交互用途，不能替代已经确认的业务路径。

## 命令语法

```
ur things device <subcommand> [选项]
```

子命令：`control`、`action`、`mock`、`simulate-report`、`report`、`upload`

各子命令的参数、示例与对应 API 见 [device-control-commands.md](device-control-commands.md)。
