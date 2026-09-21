---
name: device-firmware
description: "Use when developing or validating 联犀 IoT device firmware for Watcher/ESP32, including ESP-IDF builds, protected flashing, LLSync BLE provisioning, MQTT thing-model messages, full OTA, authentication recovery, unbind/reprovision, and hardware acceptance."
---

# 联犀设备固件

面向 SenseCAP Watcher（ESP32-S3、32MB Flash）及同类联犀设备固件。
本技能处理设备端构建、烧录、配网和生命周期验收；产品、设备、
OTA 任务等平台操作继续使用 `ur-product`、`ur-device` 和 `ur-ota`。

## 安全边界

1. 首次烧录前备份整片 Flash 和出厂分区，核对大小与 SHA-256。
2. 只按构建产物的 `flash_args` 写分区文件；不用 merged binary 覆盖
   NVS 空洞，不擦除出厂凭据。
3. ProductSecret 只在构建进程注入。日志、对话和版本库不得包含
   ProductSecret、DeviceSecret、WiFi 密码、绑定 Token 或完整注册响应。
4. 远程环境、产品 ID、分区偏移和串口名都要从当前项目确认；
   不把测试值复制到生产。

## 关键不变式

- BLE 与主应用 WifiManager 不并行初始化。临时 STA 获取 IP 后先
  SNTP 校时，再以空 DeviceSecret 走动态注册。
- 配网失败或超时时清理本轮 WiFi 凭据并释放蓝牙栈；成功绑定
  后重启，不留两套无线网络管理器竞争。
- MQTT 只在明确鉴权错误连续达到阈值时清 DeviceSecret；普通断网
  不清密钥或 WiFi。
- 联犀 OTA 与公共 OTA 源互斥。新镜像只在匹配的
  `reportInfoReply code=200` 后标记有效，否则保留 rollback 保护。

## 工作流

1. 构建、备份、分区级烧录和真实断电：读
   [构建与烧录](references/build-flash.md)。
2. 动态注册、LLSync BLE、物模型回执、鉴权恢复和 OTA 主题：读
   [接入协议](references/device-access-protocol.md)。
3. 平台准备、验收顺序、解绑重配和敏感临时文件收尾：读
   [实机验收](references/acceptance.md)。
