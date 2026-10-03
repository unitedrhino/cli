# 客户端实时调试：AI 操作流程

用于已接入调试桥的客户端报错、白屏、设备状态异常。首批客户端为 `mobile-uniapp` 与 `iot-client-app`；其他 Web、Windows、鸿蒙客户端需实际接入后才可使用。平台只转发短期日志，不保存日志正文。会话最长 30 分钟。

## 1. 确认目标并建立实时流

若目标环境已包含实例连接日志，可先按时间窗口、目标账号及应用过滤 `client-debug.instance.connected`，读取 `userID/clientInstanceID/appCode/connID`。新版客户端通过 WebSocket 的 `device-id` 传递安装实例；旧端或未重连可能没有日志。多个候选时核实身份，不能直接选择最后一条；历史连接记录不等于当前在线，仍须等待 SSE 的 `ready`。没有日志再使用设置页 ID，不额外新增实例查询 API。

先检查 `ur api --help` 是否实际支持 `--stream`；版本号或升级成功提示不代表未发布主线功能已可用。授权与动作确认均绑定原会话对象；停止、替换、过期、撤销后的迟到确认不得执行动作。

先运行 `ur check --json`，复用已有认证。请客户端用户从设置页复制**调试实例 ID**，并确认其 `userID`；不要把 IoT 设备 ID 当实例 ID。用通用 API 命令持续读取 SSE：

```bash
ur api /api/v1/system/client-debug/stream --stream --body '{"userID":"<userID>","clientInstanceID":"<实例ID>","filter":{"levels":["error","warn"],"names":["console.error","vue.error","app.error","promise.unhandled"]}}'
```

输出每行是 `{"event":"...","data":{...}}`。先从 `event=waiting` 行读取 `data.sessionID`，等 `event=ready` 后让用户复现。**保持这个终端运行**；关闭流即撤销会话。如果没有 `ready`，核对用户、实例 ID、登录态、WebSocket 连通性和客户端版本。

## 2. 读日志并按需缩小范围

从 `event=log` 行的 `data.log` 读取 `name`、`level`、`message`、`seq`、`data.value.stack`。直接 `console.error(Error)` 的堆栈通常在日志的 `data.value.args[].stack`。普通名称可精确匹配或用 `prefix.*`、`*`；`ble.*.raw` 类原始包必须精确选中完整名称，例如 `ble.llsync.raw`，并包含 `info` 级别。原始包可能含配网凭据，仅在用户知情的短期会话中使用。

收到 `dropped`，或相邻日志 `seq` 有 `gap`，表示流不完整；缩小名称/级别并重建会话复测，不能据此断言“没有报错”。如果现有日志不足以定位，在统一 logger 增加有稳定名称且脱敏的日志，重新构建客户端后复测。

## 3. 会话一次授权，逐条核对结果

在另一个终端使用同一个 `sessionID`，先请求操控授权：

```bash
ur api /api/v1/system/client-debug/message --body '{"sessionID":"<sessionID>","kind":"control.request"}'
```

先确认 `/message` 响应的 `code=200` 和 `data.accepted=true`，再等 SSE 的 `event=control` 行出现 `data.status=granted`。用户拒绝时停止操控请求，继续通过日志排查。随后每次只提交一条动作，`commandID` 在本会话内唯一：

```bash
ur api /api/v1/system/client-debug/message --body '{"sessionID":"<sessionID>","kind":"command","commandID":"diag-001","action":"app.snapshot"}'
ur api /api/v1/system/client-debug/message --body '{"sessionID":"<sessionID>","kind":"command","commandID":"diag-002","action":"device-list.refresh"}'
ur api /api/v1/system/client-debug/message --body '{"sessionID":"<sessionID>","kind":"command","commandID":"diag-003","action":"app.navigate","args":{"path":"/packageUser/pages/settings/index"}}'
```

客户端用户**本次会话只授权一次**；后续白名单动作不再弹窗，页面持续展示控制标识并可立即取消。停止、断线、退出登录或过期后授权失效，新会话重新授权。每次提交均要核对 `/message` 响应的 `code=200` 和 `data.accepted=true`；这只表示平台接受请求。必须从 `event=result` 行的 `data.commandID`、`data.status` 和 `data.data` 判断执行结果，再决定下一条动作。基础诊断动作包括 `app.snapshot`、`device-list.refresh`、`app.navigate`，导航页面白名单仅 `/pages/home/index` 和 `/packageUser/pages/settings/index`。

## 4. 读取控件并自行点击

新版客户端和后端配套增加 `ui.inspect`、`ui.tap`、`ui.input`，仍走上述两个接口，不执行脚本或任意坐标点击。旧客户端不支持时须升级，不用直接业务 API 冒充点击。

授权后发送 `ui.inspect`，实际成功结果包含 `page/ticket/controls/state`。每个控件仅有 `id/label/role/enabled`，不包含输入框值。当前开放主页设备卡片、设备详情的设置入口、设置内 Wi-Fi 管理、准备页、Wi-Fi 表单以及进度/结果页；未开放的页面会明确拒绝，不凭空猜控件。

```bash
ur api /api/v1/system/client-debug/message --body '{"sessionID":"<sessionID>","kind":"command","commandID":"ui-001","action":"ui.inspect"}'
ur api /api/v1/system/client-debug/message --body '{"sessionID":"<sessionID>","kind":"command","commandID":"ui-002","action":"ui.tap","args":{"ticket":"<本次快照ticket>","targetID":"<快照控件id>"}}'
```

输入使用 `ui.input`，参数为 `ticket/targetID/value`；密码通过受限 stdin 请求体发送，不放命令参数、shell 历史、任务日志或截图。成功结果只表示处理函数执行；之后再次读取控件和非敏感页面状态，核验实际跳转、阶段及成功结果。

每次点击或输入后重新 inspect，票据单次使用；页面隐藏、切换、控件重排、禁用、停止或撤权后旧票据无效。在途操作拒绝并发；撤销不会撤回已发出的业务请求。系统蓝牙/Wi-Fi/定位权限弹窗不能由此自动允许，恢复出厂、删除、解绑不开放。远程点击调用与手动点击相同的页面处理函数，业务权限和配网状态机保持不变。

若 ui.inspect 失败，用 app.snapshot 中 controls.registered/routeMatches 区分页面未注册与路由不匹配；这两个布尔值不读取表单或控件工厂。原生导航无回调时五秒返回失败，超时不是导航成功，也不能据此断言蓝牙失败；回读当前页面再定位。

最小复测：执行两端真实注册器和调试桥单测，验证隐藏/跳页/列表重排、票据重放、禁用控件、撤权、超长输入及密码不回读；配套后端验证会话鉴权、参数边界、转发/回执关联和 HTTP/操作日志正文屏蔽。最后在实际手机客户端自行走设备卡片→设备设置→Wi-Fi 管理→准备→填表→结果，记录真实 `result`，不能把模拟测试或 accepted 当作手机 E2E 通过。

## 5. 结束与异常

停止运行 `--stream` 的命令，或让客户端用户在设置页点“停止调试”。确认后续 `/message` 请求被拒收；不保留长期日志会话。若出现鉴权或会话过期错误，重新检查登录态、目标实例和当前会话 ID，不要复用旧 ID。截图和任意坐标点击尚未接入。

接口契约只有 `POST /api/v1/system/client-debug/stream` 和 `POST /api/v1/system/client-debug/message`。前者返回 SSE；后者交换 `ack/logs/control.request/control.grant/control.deny/command/result/stop`。平台通过既有 `user.notify.<userID>` WebSocket 频道通知目标实例，没有新增客户端控制接口。
