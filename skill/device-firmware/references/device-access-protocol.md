# Watcher 联犀接入协议

主题和字段以当前平台后端契约为权威来源，不从历史 SDK 示例推断。

## 动态注册与鉴权恢复

- DeviceSecret 为空时通过 `$ext/up/register/{pid}/{dev}` 动态注册。
- 成功后同步更新运行时内存和 NVS；空密钥写入也必须同步清理两处。
- 只在 MQTT 明确返回 `-116` 或 `-117` 连续三次时清
  DeviceSecret，保留 WiFi 并重新注册。普通 DNS、TCP、TLS 或无线断网
  不得清密钥。
- 重试使用 `30/60/120/240/300` 秒封顶退避，成功后清零。

## 物模型契约

| 方向 | 主题 | method |
|---|---|---|
| 上行 | `$thing/up/property/{pid}/{dev}` | `report`、`controlReply`、`reportInfo` |
| 下行 | `$thing/down/property/{pid}/{dev}` | `control`、`reportInfoReply` |
| 上行 | `$thing/up/action/{pid}/{dev}` | `actionReply` |
| 下行 | `$thing/down/action/{pid}/{dev}` | `action` |

- 关联字段统一使用 `msgToken`。
- `controlReply` 使用 `code:200`、`msg`、`data`。
- `actionReply` 保留 `actionID`，使用 `code:200`、`msg`、`data.Result`。
- 下行可兼容读取历史 `clientToken`、`identifier`、`actionId`，上行不再发送旧 snake_case 结构。

## LLSync BLE

- Service `FFF0`，App 写 `FFE1`，设备通知 `FFE3`，协议版本 `0x02`。
- BlueZ 验收按设备查询、STA 模式、WiFi、连接、绑定 Token 的顺序发送。
- 拿到 IP 后先 SNTP 校时；空 DeviceSecret 必须走动态注册，不得误判为已注册。
- 外层等待时间覆盖 LLSync 内层超时；Token 完成前继续广播和接受重连。
- 失败或超时先清本轮 WiFi 凭据再回退热点配网；成功后释放蓝牙栈并重启。
- 关闭 LLSync BLE OTA，升级只走经验收的 MQTT 全量 OTA。

## 全量 OTA

- 下行：`$ota/down/upgrade/{pid}/{dev}`，`method:upgrade + msgToken + data`。
- 上行：`$ota/up/progress/{pid}/{dev}`，`method:progress + msgToken + params.step/desc`。
- 只接受单文件 `isDiff=1`、`signMethod=MD5`、目标版本不同且不超过 OTA 分区的包。
- worker 校验 HTTP 200、Content-Length、实际大小、MD5、目标版本、ESP 镜像和备用分区。
- 失败码：`-1` 下载/初始化、`-2` 完整性、`-3` 写入/切换、`-4` 版本或包类型。
- 成功上报 100 后重启；匹配的 `reportInfoReply code=200` 才确认新镜像。
- 启用联犀 OTA 时跳过其他公共 OTA 检查，不让两个升级源同时决定目标版本。
