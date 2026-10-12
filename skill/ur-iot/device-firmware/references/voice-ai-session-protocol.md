# 语音会话协议与音频链路

> 本文件是 [voice-ai.md](voice-ai.md) 的子文档，收录原指南第 4、5 节：MQTT 会话时序、打断与 VAD、队列与内存约束、停顿后无结果恢复，以及 UDP 封包与音频要求。

## 4. MQTT 会话协议

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

播报中单击打断时，`respCancel` 后服务端不保证继续发送旧轮 `respAudioDone`。
设备必须主动清除本地解码尾音、旧轮超时及表情保持状态，并切回监听，发送新一轮
`audioStart` 后等待匹配的 `audioStarted`。回归需要覆盖“取消后没有终态”，
不能只在测试里补发 audioDone 让流程通过。Watcher 的
`scripts.tests.test_ur_button_interrupt` 直接编译实际按键处理函数并替换硬件边界，
验证播报中打断、监听中再次按键结束以及空闲按键启动；真机再核对取消后的新一轮
audioStarted、STT 和可听见的回复。

平台模拟必须分别覆盖直接 audioStart 打断和先 respCancel 再 audioStart：
`TestVoiceInterruptDuringTTSThenContinue` 与 `TestVoiceCancelDuringTTSThenContinue`。
后者才覆盖实体按键取消路径；两者都要求新轮 STT、文本和实际语音，不能仅检查停止播报。
如果服务端已识别但设备无 STT，核对是否在新回复准备期间退出 voice loop：
audioStop 不能只用“LLM/TTS 是否运行”判断结束，还需保护 hold、排队及准备中的回复。
后台回归使用 `TestAudioStopPreservesPendingRecognition`、
`TestAudioStopPreservesDequeuedReply`、`TestEmptyASRDoesNotCancelPendingReply`。
取消旧轮出现 context canceled 属正常现象；只有关联新轮方法序列与上下文生命周期后，
才能判定是否误取消。排查前 fetch 最新主线并核验实际部署二进制，不能用本地 HEAD 代替。

`audioStarted` 成功不代表收音正常。若随后数百毫秒内就出现 `audioStop`、没有 STT，
先核对 VAD 边沿与本轮起始时间，不要直接归因于网络。AFE 的短静音事件不是完整句尾：
Watcher 使用 300ms 启动保护窗、至少 180ms 有效语音和连续 700ms 静音判定，
在协议 worker 中检查句尾；VAD 回调和 audioStarted 回执都不能立即结束收音。
从唤醒检测或播报切入新轮时，清理 AFE 残留须由 fetch 任务按代际执行，不能跨任务
直接 reset。可移植 `UrAiVadEndpoint` 测试应覆盖启动短尾音、句中 100ms 停顿、
恢复说话取消结束判定、新轮清零及短按键噪声；这些模拟不替代实际麦克风验收。

MQTT 回调只复制 payload 并入队。`respTextDelta` 在拥塞时允许丢弃；
`respSttDone`、`respCreated`、`respEmotion`、`respTextDone`、`respAudioStart`、
`respAudioDone` 必须进入无损队列或溢出队列。重复终态、旧 token、旧 session 和错误
`respId` 必须安全忽略；不得因为 STT 没有 `respId` 而丢弃识别结果。

### 停顿后无结果与等待态

`audioStop`后上行可能已经关闭，即使应用仍为listening也不能据此判断正在收音。
先关联本轮token，确认有无STT、respCreated和新audioStarted，区分等待回复与监听。
Watcher当前无结果恢复保留60秒预算：仅无STT、无响应创建的已停止音频轮次，
主任务再次核对连接、session、token、页面互斥及晚到结果后，发出respCancel→audioStart；
新audioStarted确认前不得上传。正常模型/图片等待、audioStart超时不走此恢复。
这不是即时免唤醒；待回复期间说话仍需独立验收，不宣称该路径已支持说话即触发。
`scripts.tests.test_ur_no_result_recovery`直接编译实际超时函数、恢复函数和应用事件消费，
覆盖60秒边界、重复、旧token、晚到STT/respId、断网和菜单互斥；连续十次通过后，
仍需真机复现空识别轮次并验证恢复后的实际麦克风STT与完整回复。

Watcher 的内部 RAM 同时承载音频任务栈和 MQTT SDK。跨任务 AI 上行队列只保存小型描述符，
完整 JSON payload 应按需分配到 PSRAM，并在发布成功、入队失败和断线清队列时逐项释放。
禁止用 `队列深度 × 最大报文长度` 的固定元素预占内部 RAM；这种实现可能通过编译和协议
单测，却在真机音频初始化后令 MQTTClient 因连续内存不足而创建失败。

新增相机等功能后，还必须在真机创建 UDP 通道后检查内部 DMA 总余量和最大连续块，
不能只检查启动时总空闲堆。若 sessionCreated 已成功，随后出现
`insufficient internal DMA memory for voice session`，应排查任务栈和新增长驻队列；
网络错误文案不能作为断网证据。只在任务上下文使用的行为/回执队列可通过
`xQueueCreateWithCaps` 放入 PSRAM；不写 Flash 的相机 worker 栈可使用
`xTaskCreateWithCaps`，销毁必须配对 `vTaskDeleteWithCaps`。保留 SPI DMA 保护阈值，
修复后重复真机按键启动、语音回复和拍照；主机测试无法证明实际堆布局满足要求。
还需检查实际播报期间的余量：若出现 `esp-aes: Failed to allocate memory` 与
`AES-CTR operation failed: -132`，即使 MQTT/STT/文本成功也可能因加解密丢帧而断音。
纯 Opus 编解码任务的较大栈可迁入 PSRAM，须保留任务退出的配对释放与 OTA 停止流程；
不要降低内存门禁或把这类丢帧当成网络抖动。

## 5. UDP 与音频

UDP 数据报固定为 16 字节头加 AES-CTR 密文。帧头、nonce 字段覆盖、序号、大小端和
AES 实现必须直接对照 `devicesim/udp.go`；公共小智协议与联犀协议复用同一封包实现，
避免两份算法漂移。

联犀语音会话、devicesim、UDP channel 和固定测试样本统一声明 16 kHz、单声道、
60 ms，不能只修改其中一处。Watcher 音频硬件以 24 kHz 采集，AFE 转为 16 kHz 后编码
Opus 上行；这两个采样率属于不同层，不能因为硬件是 24 kHz 就把会话伪装为 24 kHz。

Opus 解码器可从同一 16 kHz 码流原生输出 8/12/16/24/48 kHz PCM。Watcher 收到下行
16 kHz Opus 时应直接让解码器输出扬声器所需的 24 kHz PCM，不再串联外部 16→24 kHz
重采样器。曾验证 ESP 音效库重采样器对一帧返回的容量小于实际写入量，并在返回调用方
前越界破坏堆；只在调用后比较 `actual_output` 与容量无法阻止该类越界。

下行处理要求：

- `respAudioStart` 只表示待播放；收到首个真实 UDP Opus 帧后才进入播报状态。
- `respAudioDone` 不立即截断扬声器，等待本地队列播放完成。
- 播报中检测到用户插话，先发 `respCancel`，再开始新一轮 `audioStart`。
- MQTT 断开立即关闭 UDP、丢弃旧 session；SDK 使用 `clean_session=1` 时，重连成功必须
  先恢复 AI/property/action/OTA 与解绑订阅，再向语音层宣布已连接，由下次唤醒创建新 session。
- INFO 日志只记录方法、短 session/resp ID 与耗时，不记录密钥、nonce、音频或对话全文。

