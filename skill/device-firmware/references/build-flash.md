# 固件构建、备份与首刷

先从目标仓库的板级配置、分区表和构建脚本取得事实。本文件不给所有设备套用同一
Flash 大小、偏移或串口。

## 可复现构建

1. 固定 SDK、工具链、依赖锁和容器镜像摘要，记录 Git SHA 与固件版本。
2. 构建脚本主动校验板型、Flash 大小、分区表、TLS、rollback、网络和配网组件。
3. 生成 bootloader、分区表、应用和其他实际需要的分区文件，以及权威 `flash_args`。
4. 确认应用镜像小于目标 OTA 分区，并校验镜像声明芯片与目标硬件一致。
5. 输出应用单文件的字节数、SHA-256、MD5、版本和 Git SHA，供首刷与 OTA 对照。

ProductSecret 只通过环境变量、受控标准输入或 CI secret 注入，不写入命令历史、默认
配置、构建日志或提交。构建后扫描固件和差异中是否出现测试 WiFi、Token 和秘密。

## 首刷前备份

从实际 Flash 容量和分区表计算读取范围。创建权限 `0700` 的受控目录，至少备份：

- 整片 Flash；
- 出厂身份/校准分区；
- 若需要恢复现场状态，再单独备份业务 NVS。

ESP 系列示例：

```bash
python3 -m esptool --port <serial-port> -b <safe-baud> \
  read-flash 0x0 <flash-size> <secure-backup-dir>/full-flash.bin
python3 -m esptool --port <serial-port> -b <safe-baud> \
  read-flash <factory-offset> <factory-size> <secure-backup-dir>/factory.bin
chmod 600 <secure-backup-dir>/*.bin
stat -c '%n %s' <secure-backup-dir>/*.bin
sha256sum <secure-backup-dir>/*.bin
```

读取后必须校验文件大小等于请求长度，并记录 SHA-256。备份可能含秘密，不能提交到
仓库或普通附件。

## 分区级烧录

以本次构建生成的 `flash_args`/manifest 为唯一偏移来源。逐项核对 bootloader、分区表、
应用、OTA 元数据、NVS 和出厂分区；优先只写需要更新的分区。禁止用从 `0x0` 开始的
merged binary 覆盖中间保留分区，除非目标硬件的正式量产流程明确如此且已授权。

测试需要清凭据时只擦除已确认的业务 NVS 范围。不得擦除出厂校准、MAC、设备证书、
序列号、`nvsfactory` 或 bootloader。擦除和写入后回读关键分区或用芯片工具校验。

## 启动与串口证据

烧录后记录：

- ROM/bootloader banner、reset reason、启动分区和镜像版本；
- 分区表摘要、Flash 探测容量和芯片型号；
- NVS/身份读取、配网、IP、校时、注册、MQTT 和订阅阶段；
- 首次 `reportInfo` Token 及平台对应回复。

普通应用初始化日志不是掉电或重启证据。软件行为触发的重启、watchdog、brownout 和
真实电源循环必须分别判定。

## Watcher 已验证配置（仅适用于匹配硬件）

Watcher ESP32-S3/32MB 配置使用项目锁定的 ESP-IDF 6 镜像摘要，构建门禁核对
32MB 分区、Bluedroid/GATTS、LLSync、联犀 OTA 和 rollback。其已验证出厂分区为：

```text
offset 0x9000, size 0x32000, expected bytes 204800
full flash expected bytes 33554432
```

这些数值必须再次与当前分区表交叉确认，不能复制到其他板型。Watcher 测试清理只允许
擦除确认位于出厂分区之后的业务 `nvs`。

## 真实断电

先确认设备是否有电池、超级电容或第二供电。只拔 USB 可能只是切换电源。按硬件说明
完全断电，确认串口节点消失，再恢复供电并等待重新枚举。随后验证固件版本、启动分区、
WiFi、DeviceSecret、MQTT 和 `reportInfoReply` 均恢复。

避免使用可能触发恢复出厂的长按时长。真实断电需要现场操作时，AI 在完成其余自动化
步骤后明确给出单次操作窗口并等待证据，不能把软件重启冒充断电验收。
