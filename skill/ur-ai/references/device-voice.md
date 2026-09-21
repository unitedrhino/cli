# 设备语音会话、表情与 MCP 控制

本指南描述联犀平台侧设备语音助手的配置与验证。设备固件实现、刷写、UDP 加密和
真机验收同时读取 `device-firmware/references/voice-ai.md`。

## 配置对象

运行时配置由 Product、AgentGroup、Agent、CloneGroup、Clone 和 MCP 服务共同组成：

1. Product 的默认 Agent 决定设备首次绑定的助手。
2. Clone 绑定具体设备，CloneGroup 与 AgentGroup 必须用于同一设备用途。
3. Agent 声明 `text`、`voice` 能力，并显式选择 LLM、ASR、TTS。
4. 设备控制关闭全量 MCP，只绑定审核过的 IoT MCP 服务。
5. session 注入 ProductID、DeviceName 和物模型，模型不能自行选择其他设备。

配置必须走现有 API/CLI并回读，不直接修改数据库。任何环境操作前先读对应环境说明；
密钥、访问令牌和音频内容不进入提交或普通日志。

## 会话合同

- MQTT 上行 `$thing/up/ai/{ProductID}/{DeviceName}`，下行 `$thing/down/ai/{ProductID}/{DeviceName}`。
- `sessionCreate`、`audioStart`、`sessionClose` 使用 `msgToken` 关联请求和回复。
- `sessionCreated code=200` 返回 sessionId、transport 与 UDP AES 参数。
- `respSttDelta/Done` 在 `respCreated` 前按当前 session/音频轮次处理，不要求 `respId`；
  创建响应后以 `respCreated.data.respId` 关联文本、音频、表情和 tool 状态。
- `respTextDelta` 可在队列拥塞时丢弃，终态事件不得丢弃。
- `respAudioStart` 不能证明已经产生音频；首个真实 UDP 帧才表示开始播报。
- `respAudioDone` 后设备应排空本地播放缓冲再恢复监听。
- 多轮复用同一 session；打断先 `respCancel`，再启动新一轮。
- 设备 MQTT 使用 clean session 时，自动重连后先恢复 AI 和物模型订阅，再允许创建新 session。

完整时序和加密封包以 `backend/things/tools/devicesim` 为准，不从旧客户端或示例猜协议。

## 表情合同

`respEmotion` 必须携带当前 `respId`。合法 emoji/emotion 映射以
`backend/things/tools/devicesim/emotion_test.go` 的 21 项白名单为准。未知值由设备回退
`neutral`，不能让任意模型文本直接选择设备资源名。

## MCP 最小授权

面向当前设备的语音助手只开放：

- `get_device_properties`
- `device_property_control`
- `device_action_send`

工具参数中的 ProductID、DeviceName 取自 session 上下文。工具返回成功后仍需核对：
物模型下行、设备 `controlReply/actionReply`、真实状态变化和后续属性上报。仅有模型文本
或 toolCallResult 不代表设备执行成功。

## 基线验证

先运行 devicesim，不通过时禁止用刷固件来试错：

```bash
DEVICESIM_TEST_PATTERN='Test(MultiTurnVoiceChat|VoiceInterruptDuringTTSThenContinue|VoiceTurnWithoutAudioStopStillGetsSTT|TextToTTSAudioEvents|EmojiEmotionText|DeviceControlSuccess)$' \
bash shell/devicesim-oneclick-test.sh
```

重点用例：

| 用例 | 验证目标 |
|---|---|
| `TestMultiTurnVoiceChat` | 同 session 多轮与上下文 |
| `TestVoiceInterruptDuringTTSThenContinue` | 播报中打断并继续 |
| `TestVoiceTurnWithoutAudioStopStillGetsSTT` | 服务端异常恢复能力 |
| `TestTextToTTSAudioEvents` | TTS 真实音频事件 |
| `TestEmojiEmotionText` | 表情白名单和 respId |
| `TestDeviceControlSuccess` | MCP 到物模型回复闭环 |

保存脱敏的方法序列与阶段耗时。延迟门限建议为 AudioStop 到首个音频帧小于 8 秒、
STTDone 到 TextDone 小于 15 秒。

## 分层排障

| 层级 | 判定方法 | 处理方向 |
|---|---|---|
| 平台 | devicesim 也失败 | Agent/模型/MCP 配置、ASR/LLM/TTS、UDP 服务 |
| 协议 | devicesim 通过，真机无 STT | MQTT token、UDP 预热、AES nonce、序号与 Opus 参数 |
| 播放 | 有文本无声音 | TTS 是否真实产帧、UDP 下行、Opus 解码、扬声器 |
| 控制 | 有回复无设备变化 | AgentGroup、MCP 绑定、物模型 identifier、reply 合同 |
| 生命周期 | 多轮或重连串话 | session/respId 关联、旧消息、断线后旧 UDP 未关闭 |
| 资源 | OTA 后平台离线且 MQTT 分配失败 | 检查音频初始化后的内部 RAM；大报文队列只存指针，payload 放 PSRAM并完整释放 |

平台排障以 devicesim 为第一层 oracle；真机只负责验证固件时序和硬件链路。
