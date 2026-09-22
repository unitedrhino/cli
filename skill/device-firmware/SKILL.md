---
name: device-firmware
description: "Use when bringing a physical device onto 联犀 IoT from zero or maintaining its firmware: create/reuse the product and thing model, provision identity, implement MQTT uplink/downlink, flash safely, diagnose connectivity, publish full OTA, recover credentials, and complete hardware acceptance."
---

# 联犀设备固件零到一接入

本技能面向能够修改并刷新固件的 AI。目标不是只生成代码，而是在明确授权的
环境、产品和设备范围内，把真实设备从未接入推进到可配网、在线、双向通信、
可 OTA、可恢复并有验收证据的状态。

## 开始前

1. 确认目标平台环境、硬件/芯片、固件仓库、串口、Flash 大小和是否允许平台写操作。
2. 先运行 `ur check --json` 复用现有认证；不删除正常 profile，不索取已经存在的凭据。
3. 读取当前分区表和设备身份存储实现。首次写入前备份整片 Flash 与出厂分区。
4. ProductSecret、DeviceSecret、WiFi 密码、绑定 Token、平台 Token 只经环境变量、
   受控文件或标准输入传递，不进入提交、命令历史和普通日志。
5. 已有产品优先复用。创建产品、删除/解绑设备、清 NVS、烧录和发布 OTA 都必须
   限定到用户授权的目标；不把测试设备操作扩大到其他设备。

## 零到一主流程

严格按阶段推进，每一阶段通过完成判据后再进入下一阶段：

1. 平台和设备现状盘点。
2. 创建或核对产品、物模型、设备身份与绑定策略。
3. 实现安全存储、配网、校时、动态注册和 MQTT 连接。
4. 实现属性、事件、行为、基础信息以及请求/回复关联。
5. 需要语音或视觉时，依次通过平台 devicesim、固件生产核心单元测试、
   协议回放 E2E 和真机硬件验收；语音再运行内置真实 Opus 样本的串口循环 E2E。
6. 首刷并用串口、网络和平台日志逐层验收，语音设备再验收音频与表情硬件。
7. 实现并验证全量 OTA、重启确认和 rollback。
8. 验证鉴权恢复、解绑重配、真实断电和稳定运行。

完整执行顺序、输入和完成判据见
[零到一工作流](references/from-zero-workflow.md)。

## 按任务加载参考

- 创建/核对产品、物模型、设备和绑定 Token：读
  [平台初始化](references/platform-bootstrap.md)。
- 动态注册、MQTT 凭据、TLS 与重连：读
  [认证与连接](references/mqtt-auth-registration.md)。
- 属性、事件、行为、基础信息的 Topic 和完整 JSON：读
  [物模型线协议](references/thing-model-wire-protocol.md)。
- BLE/SoftAP 配网、状态机和失败回退：读
  [配网与绑定](references/provisioning.md)。
- 构建、备份、分区级首刷和真实断电：读
  [构建与烧录](references/build-flash.md)。
- 固件上传、平台任务、设备 worker、确认与回滚：读
  [全量 OTA](references/ota.md)，平台命令同时加载 `ur-ota`。
- 联犀语音对话、表情、多轮、打断和当前设备控制：读
  [设备语音 AI](references/voice-ai.md)，平台配置同时加载 `ur-ai`。
- 相机上传、`takePhoto`、图片输入、双击拍照和表情显示生命周期：读
  [设备拍照识图](references/photo-vision.md)，平台配置同时加载 `ur-ai`。
- 故障定位：读 [分层排障](references/troubleshooting.md)。
- 换芯片、RTOS 或板型：读 [移植清单](references/porting-checklist.md)。
- 完整硬件验收与收尾：读 [实机验收](references/acceptance.md)。

## 不变式

- 配网器与主网络管理器不能竞争同一无线栈；成功或失败都必须释放临时资源。
- 获取网络后先校时。动态注册和 MQTT 鉴权都依赖正确的秒级时间戳。
- DeviceSecret 的内存值与持久化值必须同步；只在明确鉴权失败时清密钥，普通断网不清。
- `identifier`、`actionID`、事件 ID 大小写敏感，以真实物模型为准。
- MQTT 回调只解析和入队；耗时硬件动作、下载和写 Flash 在 worker 中执行。
- 联犀 OTA 与其他 OTA 源互斥；新镜像收到匹配的 `reportInfoReply code=200`
  后才标记有效。
- “CLI/API 成功”不等于“设备成功”。必须同时核对设备状态、回复报文和后续上报。
