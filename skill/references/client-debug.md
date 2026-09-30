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

客户端用户**本次会话只授权一次**；后续白名单动作不再弹窗，页面持续展示控制标识并可立即取消。停止、断线、退出登录或过期后授权失效，新会话重新授权。每次提交均要核对 `/message` 响应的 `code=200` 和 `data.accepted=true`；这只表示平台接受请求。必须从 `event=result` 行的 `data.commandID`、`data.status` 和 `data.data` 判断执行结果，再决定下一条动作。允许动作只有 `app.snapshot`、`device-list.refresh`、`app.navigate`，导航页面白名单仅 `/pages/home/index` 和 `/packageUser/pages/settings/index`。

## 4. 结束与异常

停止运行 `--stream` 的命令，或让客户端用户在设置页点“停止调试”。确认后续 `/message` 请求被拒收；不保留长期日志会话。若出现鉴权或会话过期错误，重新检查登录态、目标实例和当前会话 ID，不要复用旧 ID。截图和任意坐标点击尚未接入。

接口契约只有 `POST /api/v1/system/client-debug/stream` 和 `POST /api/v1/system/client-debug/message`。前者返回 SSE；后者交换 `ack/logs/control.request/control.grant/control.deny/command/result/stop`。平台通过既有 `user.notify.<userID>` WebSocket 频道通知目标实例，没有新增客户端控制接口。
