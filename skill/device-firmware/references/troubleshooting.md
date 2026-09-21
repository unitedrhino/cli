# 联犀设备分层排障

按层定位，上一层没有证据前不要修改下一层代码。每次只改变一个变量，并用
ProductID、DeviceName、`msgToken`、时间窗口和固件版本关联证据。

## 1. 启动与持久化

检查 boot banner、reset reason、运行分区、固件版本、分区表、NVS 初始化和设备名。

- 反复复位：先区分 watchdog、brownout、软件重启和 rollback。
- 身份每次丢失：核对 NVS namespace/key、提交结果和读取长度；不要先重刷整片。
- 版本没变：核对启动分区和镜像版本字符串，不以构建文件名判断。

## 2. 配网与 IP

检查 BLE/SoftAP 状态、SSID 长度、WiFi reason code、DHCP 地址和网关。

- BLE 可见但不能写：核对 FFF0/FFE1/FFE3、通知订阅和 SDK 版本帧。
- 写入成功但不上线：确认配网器是否释放无线栈，主 WifiManager 是否随后启动。
- 反复回配网：区分候选 WiFi 失败与已保存凭据读取失败。

## 3. DNS、TLS 与时间

分别记录 DNS 解析、TCP 建连、TLS 验证和 SNTP 状态。TLS 证书“尚未生效/已过期”常由
设备时间错误导致。不能通过关闭证书验证掩盖时钟、CA 或主机名问题。

## 4. 动态注册

检查：产品是否启用对应 `autoRegister`、DeviceName 是否已占用、Unix 秒时间戳、签名
原文顺序、ProductSecret 是否被错误 Base64 解码、下行是否先订阅、响应路径是
`data.payload` 还是旧 `payload`、AES key/IV 和 `psk` 持久化。

平台返回“设备已注册”时先查设备生命周期状态；不要不断清 NVS 重试。若平台没有正式
恢复待注册态接口，停止并报告，不直接改数据库。

## 5. 正常 MQTT 鉴权

逐字记录脱敏后的 clientID 结构、username 字段数、expiry 与设备时钟差、算法后缀，
并确认 DeviceSecret 是 Base64 解码后作为 HMAC key。

- `-116/-117`：达到连续阈值才清 DeviceSecret 并重新注册。
- 连接超时/断网：保留密钥，只做网络重连。
- 连接后收不到下行：检查订阅 ACK、Topic 的产品和设备段及 ACL。

## 6. Topic 与物模型

先看原始 Hub 日志，再看属性/事件/命令解析日志：

```bash
ur things device log hub -p <productID> -d <deviceName>
ur things device log property -p <productID> -d <deviceName>
ur things device log event -p <productID> -d <deviceName>
ur things device log send -p <productID> -d <deviceName>
ur things device log status -p <productID> -d <deviceName>
ur things device log abnormal -p <productID> -d <deviceName>
ur things device log sdk -p <productID> -d <deviceName>
```

CLI 参数不确定时先执行对应 `--help`。日志时间范围使用平台要求的毫秒值。常见问题：
identifier 大小写错误、把秒当毫秒、`clientToken`/`msg_token` 混入新上行、数值越界、
动作输出与物模型不一致、回复使用新 Token。

## 7. 控制和行为

用同一 Token 核对：平台下行 → 设备解析 → worker 执行 → reply → 后续 report。

- reply 成功但硬件没变化：执行顺序错误，不能先回成功。
- 硬件变化但平台失败：检查 reply Topic、method、Token、`actionID` 和 `data` 类型。
- 重启两次：重复消息没有幂等去重，或 MQTT 回调与 worker 都执行了动作。

## 8. OTA

按任务下行、HTTP、长度/MD5、分区写入、启动分区、重启、`reportInfoReply`、平台任务
依次定位。平台任务卡住时同时查询任务详情和设备进度日志；只看到 HTTP 200 不能证明
镜像写入成功。新镜像反复回旧版通常是确认条件未满足、自检失败或版本字段不一致。

## 9. 证据最小集

提交问题时保留：固件 Git SHA 和版本、分区表摘要、reset reason、脱敏设备身份、失败
阶段、SDK 错误码、相关 Token、平台原始 Hub 消息、设备日志时间窗口和最后一次成功
状态。秘密、WiFi 密码、完整签名 URL 和动态注册明文不得进入证据。
