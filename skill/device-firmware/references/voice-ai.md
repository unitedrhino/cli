# 联犀设备语音 AI 从零接入

本指南用于把已经具备联犀 MQTT、物模型与 OTA 能力的设备接入语音对话、表情和
当前设备控制。协议行为以 `backend/things/tools/devicesim` 为唯一基准：先证明平台
可用，再验证固件协议，最后验证麦克风、扬声器和显示硬件。

## 完成标准

只有同时满足以下条件才算完成：

1. devicesim 的 ASR、LLM、TTS、UDP、表情、多轮、打断和 MCP 控制回归通过。
2. 真机可通过唤醒词和按键建会话，连续完成至少三轮对话并正确播放尾音。
3. 表情严格采用平台 21 项白名单，未知值显示 `neutral`。
4. 查询、属性控制和行为调用均有 MCP 调用、物模型下行、设备回复、状态变化和后续上报五层证据。
5. MQTT 断开后废弃旧 session，恢复网络后无需重新配网或注册；普通断网不清 DeviceSecret。
6. OTA 前关闭语音会话，升级期间拒绝新会话；升级、断电和 15 分钟稳定性验收通过。

## 1. 平台对象与安全边界

先确认环境、ProductID、DeviceName、Agent、AgentGroup、Clone 与 MCP 服务 ID。读取平台
环境说明后用 `ur check --json` 确认身份，所有变更通过现有 API 或 CLI 完成，不直接写库。

产品的默认 Agent、设备 Clone 与 CloneGroup 共同决定最终运行配置。建议设备语音助手：

- AgentGroup 的用途为 `device`，与设备 CloneGroup 对齐。
- Agent capabilities 至少包含 `text`、`voice`。
- 显式选择 LLM、ASR、TTS 配置，不依赖未核验的历史默认值。
- 关闭全量 MCP，只绑定设备所需 MCP 服务。
- 仅开放 `get_device_properties`、`device_property_control`、`device_action_send`。
- 目标 ProductID、DeviceName 由 session 注入，模型输出不能覆盖目标设备。

变更后必须逐项回读 Agent、产品默认 Agent、Clone、CloneGroup 和 MCP 绑定。API 返回
`code=200` 只表示配置写入成功，不能替代 devicesim 与真机验证。

## 2. 先运行 devicesim 平台基线

在仓库根目录执行完整回归。凭据通过已有 profile 或环境变量注入，不写入命令、文档或日志：

```bash
DEVICESIM_TEST_PATTERN='Test(MultiTurnVoiceChat|VoiceInterruptDuringTTSThenContinue|VoiceTurnWithoutAudioStopStillGetsSTT|TextToTTSAudioEvents|EmojiEmotionText|DeviceControlSuccess)$' \
DEVICESIM_AUDIO_SAMPLE_RATE=16000 \
bash shell/devicesim-oneclick-test.sh
```

独立能力测试用于缩小问题范围：

```bash
DEVICESIM_TEST_PATTERN='Test(VoiceASRNumberSample|VoiceASRQuerySample|TextToTTSAudioEvents)$' \
bash shell/devicesim-oneclick-test.sh
```

稳定性与延迟入口：

```bash
DEVICESIM_TEST_PATTERN='TestConcurrentStabilityRandomMix$' \
DEVICESIM_TEST_TIMEOUT=5400s \
DEVICESIM_STABILITY_DEVICES=5 \
DEVICESIM_STABILITY_ROUNDS=100 \
DEVICESIM_STABILITY_INTERRUPT_WINDOW=5 \
bash shell/devicesim-oneclick-test.sh
```

保留脱敏的方法序列、通过用例与阶段耗时。devicesim 不通过时不要刷真机；先修平台配置、
ASR/LLM/TTS、MCP 或 UDP 服务端问题。

## 3. MQTT 会话协议

设备复用物联网主 MQTT 客户端，不建立第二连接：

- 上行：`$thing/up/ai/{ProductID}/{DeviceName}`
- 下行：`$thing/down/ai/{ProductID}/{DeviceName}`
- 会话请求通过 `msgToken` 关联。
- `respSttDelta/Done` 发生在 `respCreated` 之前，按当前 session 与当前音频轮次接收；
  `respCreated` 之后的文本、音频、表情和 tool 状态再通过 `sessionId + respId` 关联。

固定时序：

1. 发送 `sessionClose` 清理遗留会话，匹配 token 等待 `sessionClosed`；超时允许继续。
2. 发送 `sessionCreate`，只接受同 token、`code=200` 的 `sessionCreated`。
3. 校验 `sessionId`、`transport=udp`、服务地址、端口、16 字节 AES key 和 nonce。
4. 在新 UDP 通道发送两帧与真实编码参数一致的 Opus 静音，预热服务端路由。
5. 发送 `audioStart`；同 token 的 `audioStarted code=200` 到达前只做有界缓存，不上传真实语音。
6. 本地 VAD 检测一句结束后发送一次 `audioStop`。
7. 接收 `respSttDone`、`respCreated`、文本、音频、表情和工具状态；不得假定到达顺序。
8. `respAudioDone` 后等本地解码播放缓冲清空，再进入下一轮；不重建 session。
9. 退出时发送 `sessionClose`，正常路径等待 `sessionClosed`。

MQTT 回调只复制 payload 并入队。`respTextDelta` 在拥塞时允许丢弃；
`respSttDone`、`respCreated`、`respEmotion`、`respTextDone`、`respAudioStart`、
`respAudioDone` 必须进入无损队列或溢出队列。重复终态、旧 token、旧 session 和错误
`respId` 必须安全忽略；不得因为 STT 没有 `respId` 而丢弃识别结果。

## 4. UDP 与音频

UDP 数据报固定为 16 字节头加 AES-CTR 密文。帧头、nonce 字段覆盖、序号、大小端和
AES 实现必须直接对照 `devicesim/udp.go`；公共小智协议与联犀协议复用同一封包实现，
避免两份算法漂移。

音频参数以设备编码器的真实输出和 `sessionCreate.audioParams` 一致为准。Watcher 的 codec
以 24 kHz 采集，但现有 AFE/Opus 上行链路会重采样并编码为 16 kHz、单声道、60 ms，
因此联犀会话声明 16 kHz。不能只按 codec 采样率误报 24 kHz。

下行处理要求：

- `respAudioStart` 只表示待播放；收到首个真实 UDP Opus 帧后才进入播报状态。
- `respAudioDone` 不立即截断扬声器，等待本地队列播放完成。
- 播报中检测到用户插话，先发 `respCancel`，再开始新一轮 `audioStart`。
- MQTT 断开立即关闭 UDP、丢弃旧 session；SDK 使用 `clean_session=1` 时，重连成功必须
  先恢复 AI/property/action/OTA 与解绑订阅，再向语音层宣布已连接，由下次唤醒创建新 session。
- INFO 日志只记录方法、短 session/resp ID 与耗时，不记录密钥、nonce、音频或对话全文。

## 5. 表情和界面

`respEmotion` 必须与当前 `respId` 匹配。允许的 emoji/emotion 映射以
`backend/things/tools/devicesim/emotion_test.go` 的 21 项白名单为准；未知、空值或过期
消息统一回退 `neutral`。STT 显示用户字幕，文本 delta 增量累积助手字幕并对终态去重。
tool 状态只能作界面提示，不能代替真实物模型控制。

## 6. 当前设备控制验收

先查物模型确认标识符，再分别验证查询、属性和行为。示例命令仅用于人工对照：

```bash
ur things schema get-list -p <product-id> -d <device-name>
ur things device control -p <product-id> -d <device-name> --data '{"Volume":35,"Brightness":40}'
ur things device action send -p <product-id> -d <device-name> --data-id SendMessage --input '{"Message":"语音验收"}'
```

每次语音控制都要采集五层证据：

1. Agent 选择了允许的 MCP tool，参数中的目标设备来自 session 上下文。
2. 平台发出正确 property/action 下行。
3. 设备返回同 token 的 `controlReply` 或 `actionReply`，且 `code=200`。
4. 屏幕、音量、亮度或重启等真实状态发生变化。
5. 后续属性上报与设备状态一致；Reboot 必须先回复，再只重启一次。

模型口头说“已完成”不属于控制成功证据。

## 7. 构建、OTA 与真机验收

构建前执行资源门禁，并校验 IDF 版本、板型、Flash、分区、rollback、BLE、联犀 MQTT/OTA、
本地唤醒词和语音 AI 配置。固件不得包含平台凭据、WiFi 密码或公共小智服务 URL。

首次刷写遵循 [构建与烧录](build-flash.md)，先备份整片 Flash 和 `nvsfactory`，使用分区级
`flash_args`，不覆盖出厂分区。已有可 OTA 版本优先按 [全量 OTA](ota.md) 升级，并通过
`reportInfoReply code=200` 确认新镜像。

真机至少完成：

- 唤醒词和按键各建会话一次。
- 三轮连续聊天、多轮记忆、播报完成后继续、播报中打断。
- `happy`、`sad`、`angry`、`thinking`、`neutral` 显示；协议测试覆盖全部 21 类。
- 属性查询、属性设置、消息显示和重启行为。
- AudioStop 到首个音频帧小于 8 秒，STTDone 到 TextDone 小于 15 秒。
- 断网恢复、一次真实断电、15 分钟在线稳定观察。

## 8. 分层诊断

| 现象 | 先查 | 常见根因 |
|---|---|---|
| devicesim 失败 | 平台配置与服务日志 | Agent/模型/MCP 绑定错误，ASR/LLM/TTS 或 UDP 服务故障 |
| devicesim 通过、真机无 STT | MQTT 方法序列和 UDP 向量 | 未等 `audioStarted`、预热失败、nonce/序号错误、Opus 参数不一致 |
| 有文本无声音 | TTS 事件与 UDP 收包 | TTS 无帧、UDP 下行、Opus 解码或扬声器故障 |
| 有声音但尾音被截 | 播放队列 | 收到 `respAudioDone` 后过早切换状态 |
| 有回复无控制 | MCP 与物模型日志 | AgentGroup/MCP 绑定、identifier、reply token 或目标上下文错误 |
| 多轮串话 | session/resp 短 ID | 旧 session 未清理、未校验 respId、重复终态未去重 |
| 打断无效 | VAD 与方法序列 | 未发 `respCancel`，或未重新等待 `audioStarted` |
| 断网后异常 | MQTT/UDP 生命周期 | 继续使用旧 UDP/session、clean session 重连后漏订阅，或普通断网误清 DeviceSecret |

排障顺序不可颠倒：devicesim 平台基线 → 报文/加密向量 → Watcher 固件 → 麦克风、
扬声器和显示硬件。
