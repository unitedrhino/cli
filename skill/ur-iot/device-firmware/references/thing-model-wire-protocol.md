# 物模型 MQTT 线协议

所有 Topic 中的 `{productID}`、`{deviceName}` 都取当前设备身份。所有标识符大小写
敏感，并以平台回读后的物模型为准。时间戳是 Unix 毫秒；请求和回复必须复用同一个
`msgToken`。

## 属性

```text
上行 $thing/up/property/{productID}/{deviceName}
下行 $thing/down/property/{productID}/{deviceName}
```

属性上报：

```json
{
  "method": "report",
  "msgToken": "<unique-token>",
  "timestamp": 1687525491000,
  "params": {"Temperature": 23.5, "Power": 1}
}
```

平台回复使用 `method=reportReply`。首次连接和固件重启后还要上报设备基础信息：

```json
{
  "method": "reportInfo",
  "msgToken": "<unique-token>",
  "timestamp": 1687525491000,
  "params": {
    "version": "<firmware-version>",
    "mac": "<mac>",
    "module": "<module-code>",
    "hardInfo": "<hardware-revision>",
    "softInfo": "<build-summary>",
    "localIp": "<local-ip>"
  }
}
```

不适用的 `imei`、`projectID`、`position`、`tags`、`deviceAlias` 可按产品要求省略，
不要伪造。平台回复 `reportInfoReply`；OTA 待确认镜像还必须验证 `code=200` 且
回复关联当前版本的上报后再确认。

属性控制请求示例：

```json
{
  "method": "control",
  "msgToken": "<request-token>",
  "timestamp": 1687525491000,
  "params": {"Power": 1}
}
```

成功回复：

```json
{
  "method": "controlReply",
  "msgToken": "<request-token>",
  "code": 200,
  "msg": "success",
  "data": {"Power": 1}
}
```

设备必须校验类型、范围、枚举和读写权限，执行真实硬件动作后才回复成功，并再上报
最终实际值。兼容读取旧下行中的 `clientToken` 或 `identifier` 时，只在解析层兼容；
新上行仍发送当前 camelCase 合同。

按需实现 `getReport/getReportReply` 和 `getStatus/getStatusReply`。回复的 `data` 必须是
当前设备真实状态，不能回显未经执行的请求值。

## 事件

```text
上行 $thing/up/event/{productID}/{deviceName}
下行 $thing/down/event/{productID}/{deviceName}
```

```json
{
  "method": "eventPost",
  "msgToken": "<unique-token>",
  "version": "1.0",
  "eventID": "<thing-model-event-id>",
  "type": "info",
  "timestamp": 1687525491000,
  "params": {"Code": 0}
}
```

平台回复 `eventReply`。事件 ID、`type` 和参数必须匹配物模型；不要用事件模拟持续状态。

## 行为

```text
下行 $thing/down/action/{productID}/{deviceName}
上行 $thing/up/action/{productID}/{deviceName}
```

请求：

```json
{
  "method": "action",
  "msgToken": "<request-token>",
  "actionID": "<thing-model-action-id>",
  "timestamp": 1687525491000,
  "params": {"Duration": 3}
}
```

回复：

```json
{
  "method": "actionReply",
  "msgToken": "<request-token>",
  "actionID": "<thing-model-action-id>",
  "code": 200,
  "msg": "success",
  "data": {"Result": 1}
}
```

`data` 的字段由行为输出模型决定；`Result` 只是常用示例，不得强加给没有该输出的
产品。兼容旧 `actionId` 只限下行解析。耗时行为进入 worker，MQTT 回调不能阻塞。

## 服务与绑定

```text
上行 $thing/up/service/{productID}/{deviceName}
下行 $thing/down/service/{productID}/{deviceName}
```

用户绑定通过 `method=appBindToken` 上报配网阶段取得的 Token，并等待对应回复。平台
历史协议中存在拼写为 `bingChange` 的绑定变化通知；实现时按目标环境实测字段处理，
不能自行改成看似正确的新拼写。

## 消息处理不变式

- Token 在一次请求/回复链路中保持不变；本机生成的 Token 必须足够唯一。
- 使用 JSON 类型校验，不把字符串数字静默当成数值。
- 对同一 Token 的重复控制或行为做幂等保护，特别是重启、开锁等不可逆动作。
- MQTT 回调只完成长度检查、解析、校验和入队；硬件操作在专用任务中执行。
- 错误回复使用非 200 `code` 和可诊断 `msg`，但不泄露秘密或内部栈信息。
