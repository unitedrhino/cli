# Watcher 构建与烧录

## 可复现构建

- 使用项目锁定到摘要的 ESP-IDF 6 容器镜像，不使用漂移的 tag。
- 默认启用 Bluedroid、GATTS、LLSync BLE、OTA rollback 和联犀 OTA。
- 构建脚本必须检查板型、32MB Flash、分区表和关键 Kconfig，失配时直接失败。
- 只上传应用分区的单文件二进制做全量 OTA，并确认它小于目标
  OTA 分区。

ProductSecret 用环境变量或受控标准输入注入构建进程，不写命令历史、
配置示例或产物日志。

## 首刷前备份

确认目标串口后，先在权限 `0700` 的目录中读取：

```bash
python3 -m esptool --port <serial-port> -b 2000000 \
  read-flash 0x0 0x2000000 <backup-dir>/watcher-full-32m.bin
python3 -m esptool --port <serial-port> -b 2000000 \
  read-flash 0x9000 0x32000 <backup-dir>/nvsfactory.bin
chmod 600 <backup-dir>/*.bin
stat -c '%n %s' <backup-dir>/*.bin
sha256sum <backup-dir>/*.bin
```

Watcher 已验证的预期大小是整片 `33554432` 字节、`nvsfactory`
`204800` 字节。偏移仍必须与当前分区表交叉确认。

## 分区级烧录

以 `flash_args` 的文件和偏移为唯一依据。已验证的 Watcher 布局中，
`nvsfactory` 占用 `0x9000..0x3b000`，测试清理只允许擦除其后的
`nvs`；禁止把 merged binary 从 `0x0` 整体写入。

烧录后用 boot banner 和 reset reason 识别启动类型。普通 `app_init`
日志不是掉电证据。

## 真实断电

Watcher 内置电池，只拔 USB 可能仅切换到电池供电。现场验收应：

1. 拔 USB，按当前硬件说明完全关机，避免使用可能触发恢复出厂的长按时长。
2. 确认旧串口实例因设备断开而退出，再重插 USB 并等待设备节点重新枚举。
3. 对比操作前后固件运行时间，并确认 WiFi、MQTT、`reportInfoReply`
   和固件版本都恢复。

若稳定观察期间收到平台明确的 `Reboot` 下行，核对
`actionReply` 和 `RTC_SW_CPU_RST`，并从 MQTT 重连时重新计算
完整观察窗口。
