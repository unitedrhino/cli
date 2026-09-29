# 联犀全量 OTA：设备端闭环

平台资源的上传、固件记录和任务创建同时加载 `ur-ota`。本文件聚焦设备端协议、
安全写入、确认和回滚。

## 下行与进度

```text
下行 $ota/down/upgrade/{productID}/{deviceName}
上行 $ota/up/progress/{productID}/{deviceName}
```

升级请求：

```json
{
  "method": "upgrade",
  "msgToken": "<request-token>",
  "timestamp": 1687525491000,
  "data": {
    "size": 1048576,
    "signature": "<signature-if-present>",
    "version": "2.0.0",
    "isDiff": 1,
    "name": "firmware.bin",
    "fileUrl": "<temporary-download-url>",
    "signMethod": "MD5",
    "fileMd5": "<32-hex-md5>",
    "extData": ""
  }
}
```

当前平台语义中 `isDiff=1` 是全量包，`2` 是差分包。只实现全量时必须拒绝其他值，
不要按字段英文直觉反转。进度回复复用下行 Token：

```json
{
  "method": "progress",
  "msgToken": "<request-token>",
  "params": {"step": 35, "desc": "downloading", "module": "<module-code>"}
}
```

`step` 正常为 1–100。推荐设备失败码：`-1` 下载/初始化失败、`-2` 完整性校验失败、
`-3` 写入或切换分区失败、`-4` 版本或包类型不匹配。`desc` 保持简短，不包含 URL
签名、密钥或内部堆栈。

## 接收门禁

解析后、开始下载前同时校验：

- `method=upgrade`、Token、URL、版本、大小、算法和摘要字段齐全；
- 单文件全量包，`isDiff=1` 且 `signMethod=MD5`；
- 目标版本不同于当前版本，也不同于正在执行的任务；
- 声明大小大于 0 且不超过目标 OTA 分区；
- 当前无其他 OTA worker，升级源没有被公共 OTA 或 BLE OTA 占用；
- URL scheme 和 TLS 策略符合产品安全要求。

拒绝时报告对应失败码，不改启动分区。

## Worker 流程

MQTT 回调只校验并复制不可变任务参数，然后启动唯一 OTA worker。worker：

1. 选择非当前启动分区，核对分区大小和写权限。
2. 发起 HTTP 请求，验证状态码、重定向策略和声明长度。
3. 流式写入备用分区，同时累计实际长度和 MD5；不把整包读入 RAM。
4. 保持 MQTT 主任务 Yield；通过互斥状态快照上报进度，避免 worker 直接并发操作客户端。
5. 下载结束校验声明大小、实际大小、MD5、镜像头/芯片类型和固件版本。
6. 完成平台/芯片 SDK 的镜像验证，再设置下次启动分区。
7. 持久化“待确认版本 + Token”，上报 100，等待消息实际发送后重启。

任何失败都终止 OTA handle、保持原启动分区并清理临时状态。写入失败后不要继续调用
切换分区 API。

## 新镜像确认与回滚

启用芯片 SDK 的 rollback 机制。新镜像第一次启动时：

1. 读取待确认状态，确认运行版本与目标版本一致。
2. 完成基本硬件、自检、WiFi、MQTT 和下行订阅。
3. 发送包含当前版本的 `reportInfo`。
4. 收到与该消息关联的 `reportInfoReply code=200` 后，才调用 SDK 标记镜像有效。
5. 清理待确认状态并继续正常运行。

在确认前 watchdog、异常重启或自检失败应触发 bootloader rollback。不能因“成功启动到
`app_main`”提前确认，也不能只凭 OTA 任务已创建就判定成功。

## OTA 验收矩阵

至少验证：

| 场景 | 期望 |
|---|---|
| 正常旧版到新版 | 1–100、重启、新版本 reportInfo、平台任务成功 |
| 重复版本 | `-4`，不写 Flash |
| 差分包或未知算法 | `-4`，不写 Flash |
| 大小超过分区 | 拒绝，启动分区不变 |
| MD5 错误/下载截断 | `-2`，仍运行旧版 |
| 写入或切换失败 | `-3`，仍运行旧版 |
| 新镜像确认前异常重启 | 自动 rollback 到旧版 |

最终核对设备串口版本、平台设备 `reportInfo` 版本和 OTA 任务结果三者一致。
