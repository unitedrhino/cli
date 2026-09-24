# 联犀设备语音 AI 从零接入

本指南用于把已经具备联犀 MQTT、物模型与 OTA 能力的设备接入语音对话、表情和
当前设备控制。协议行为以 `backend/things/tools/devicesim` 为唯一基准：先证明平台
可用，再验证固件协议，最后验证麦克风、扬声器和显示硬件。

设备还需要拍照识图、图片输入或表情可见期时，同时读取
[设备拍照识图](photo-vision.md)。

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
DEVICESIM_TEST_PATTERN='Test(MultiTurnVoiceChat|VoiceInterruptDuringTTSThenContinue|VoiceCancelDuringTTSThenContinue|VoiceTurnWithoutAudioStopStillGetsSTT|TextToTTSAudioEvents|EmojiEmotionText|DeviceControlSuccess)$' \
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

## 3. 固件单元测试与协议 E2E

devicesim 通过只证明平台链路可用，不能替代固件测试。修改语音状态机、关联规则、UDP
封包或表情映射后，在 `firmware/watcher` 执行：

```bash
python3 -m unittest \
  scripts.tests.test_ur_ai_contract \
  scripts.tests.test_ur_ai_runtime \
  scripts.tests.test_run_ur_ai_e2e -v
```

两类测试必须同时通过：

- `test_ur_ai_contract` 是静态合同检查，用于发现 Topic、方法名、队列和构建配置意外漂移；
  不能把它单独称为固件单元测试。
- `test_ur_ai_runtime` 会用主机 C++17 编译器直接编译固件共用的生产协议核心。`unit`
  套件运行 21 类表情、session/respId 关联、终态去重和 UDP 头边界；`e2e` 套件按
  devicesim 时序回放 UDP 预热、多轮、乱序、重复、打断尾帧和断线旧消息。

完整固件回归仍使用 `python3 -m unittest discover -s scripts/tests -v`，随后以固定 IDF 6
镜像构建目标板。主机协议 E2E 不包含真实 MQTT、云端模型、Opus 编解码或硬件；因此它必须
与上一节的真实 devicesim E2E、最终真机验收组合，任何一层都不能代替另外两层。

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

记忆验收区分同 session 历史、刚重建 session 的短时衔接和异步长期画像。
`TestVoicePreferenceRecall` 必须复述实际 STT 中的偏好；“没告诉我，比如……”即使含
关键词也必须失败，提示词不能预先泄露预期食材。长期整理保持后台执行；短时衔接仅查询
同企业/分身/Agent/归属用户最近两分钟的三条用户原文（合计 600 字符、150ms 查询预算），
排除当前、网页、归档、删除会话，不把原文写回画像。不承诺数据库落库前或超出短时窗口
的立即召回。用隔离与过期测试、后台阻塞测试和真实重连 E2E 分别证明。
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

## 6. 可重复的真机音频 E2E

该用例只用内置音频替代麦克风输入，其余全部走真实链路：主 MQTT、
`sessionCreate/audioStart/audioStop`、UDP AES-CTR、平台 ASR/LLM/TTS、真实 UDP 下行、
Opus 解码、播放队列和界面状态。不能用直接注入 STT 文字或伪造回复代替。

按以下固定用例分层执行和留证，不把任一层通过替代为全链路通过：

| 用例 | 输入与执行 | 必须断言 | 建议频率 |
|---|---|---|---|
| `VOICE-UNIT-001` 资源格式 | 运行 `go test ./tools/devicesim/cmd/opusfixture` | URAF 头、16 kHz/60 ms、帧边界和非法帧拒绝 | 每次修改生成器 |
| `VOICE-UNIT-002` 固件协议 | 运行 `test_ur_ai_contract`、`test_ur_ai_runtime` | token/respId、乱序、重复、过期、断线、UDP 和 21 类表情 | 每次修改语音固件 |
| `VOICE-RUNNER-001` 判定器 | 运行 `test_run_ur_ai_e2e` | 成功、乱序、缺阶段、崩溃标记和严格延迟门槛 | 每次修改 runner |
| `VOICE-SIM-001` 平台基线 | 运行 devicesim workflow（含两种打断路径） | ASR、LLM、TTS、UDP、MCP、多轮和打断 | 平台配置或服务变更后 |
| `VOICE-HW-001` 单轮闭环 | runner `--repeat 1`，固定语料“当前音量多少” | 全部阶段、STT 命中“音量”、七个 RESULT 标志、两项延迟和无致命标记 | 每次测试固件刷写后 |
| `VOICE-HW-002` 重复稳定性 | runner `--repeat 5` 或更高 | 每轮独立 session、全部通过、各轮原始日志和 JSON | MR 前至少五轮 |
| `VOICE-MANUAL-001` 硬件验收 | 真人唤醒、说话并观察/听取设备 | 麦克风、唤醒词、扬声器、字幕和表情实物效果 | 发版与现场验收 |

音频资源必须由 devicesim 的同一 MP3→Opus 编码器生成，不手工组帧：

```bash
cd backend/things
go run ./tools/devicesim/cmd/opusfixture \
  -input test/testdata/turn4_query_volume_howmuch.mp3 \
  -output ../../firmware/watcher/main/testdata/ur_ai_query_volume.opuspack \
  -sample-rate 16000 -frame-duration-ms 60 -tail-silence-frames 5
```

测试音频固件必须显式以 `ENABLE_UR_AI_E2E_TEST=1` 构建；生产固件默认为 `0`，
不带样本和串口测试命令。仅写当前应用分区，不覆盖 NVS、`nvsfactory` 或分区表。

固件上线并处于 idle 后，在串口直连主机上执行：

```bash
python3 firmware/watcher/scripts/run_ur_ai_e2e.py \
  --port <serial-port> --repeat 5 --timeout 75 \
  --log-dir <repo>/.temp/device-firmware/ur-ai-e2e
```

runner 每轮创建和关闭独立 session，并断言：

1. `sessionCreated → audioStarted → INPUT_DONE` 全部到达。
2. `respSttDone`、`respTextDone`、首个真实音频帧和 `respAudioDone` 全部到达；
   响应阶段可交错，不强制错误的固定顺序。
3. RESULT 中 STT/文本/音频七个标志均为 `1`，其中 STT 必须命中固定语料关键词
   “音量”，最终打印 `PASS`；日志不输出完整对话。
4. 日志不含 `Failed to resample output audio`、`assert failed`、`Guru Meditation`、
   `Backtrace`、复位或命令错误；重采样容量异常即失败，不能等堆断言。
5. 以 ESP 日志的统一 uptime 计算门禁：AudioStop→首帧小于 8 秒，
   STTDone→TextDone 小于 15 秒。不能相减计时原点不同的 `elapsed`。

每轮输出 JSON 摘要和单独原始日志；任一轮缺阶段、超时、崩溃或复位即非零退出。
循环 E2E 只证明云端到扬声器队列的可重复闭环；麦克风、唤醒词和实际扬声器响度
仍需最终人工真机验收。

## 7. 表情和界面

`respEmotion` 必须与当前 `respId` 匹配。允许的 emoji/emotion 映射以
`backend/things/tools/devicesim/emotion_test.go` 的 21 项白名单为准；未知、空值或过期
消息统一回退 `neutral`。STT 显示用户字幕，文本 delta 增量累积助手字幕并对终态去重。
tool 状态只能作界面提示，不能代替真实物模型控制。

## 8. 当前设备控制验收

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

## 9. 构建、OTA 与真机验收

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

## 10. 分层诊断

| 现象 | 先查 | 常见根因 |
|---|---|---|
| devicesim 失败 | 平台配置与服务日志 | Agent/模型/MCP 绑定错误，ASR/LLM/TTS 或 UDP 服务故障 |
| devicesim 通过、真机无 STT | MQTT 方法序列和 UDP 向量 | 未等 `audioStarted`、预热失败、nonce/序号错误、Opus 参数不一致 |
| 有文本无声音 | TTS 事件与 UDP 收包 | TTS 无帧、UDP 下行、Opus 解码或扬声器故障 |
| 一轮 PASS 后崩溃 | runner 原始日志、`assert failed` 前最后一条音频日志与 ELF 回溯 | 若看到重采样实际输出大于容量，属于返回前已发生的越界；保持 16 kHz 协议并改为 Opus 原生 24 kHz 解码，不能靠事后检查或把协议改成 24 kHz 掩盖 |
| 有声音但尾音被截 | 播放队列 | 收到 `respAudioDone` 后过早切换状态 |
| 有回复无控制 | MCP 与物模型日志 | AgentGroup/MCP 绑定、identifier、reply token 或目标上下文错误 |
| 多轮串话 | session/resp 短 ID | 旧 session 未清理、未校验 respId、重复终态未去重 |
| 打断无效 | VAD 与方法序列 | 未发 `respCancel`，或未重新等待 `audioStarted` |
| 断网后异常 | MQTT/UDP 生命周期 | 继续使用旧 UDP/session、clean session 重连后漏订阅，或普通断网误清 DeviceSecret |
| OTA 后无 MQTT | 启动串口的空闲/最小内部 RAM 与 SDK 返回码 | AI 大报文队列预占内部 RAM；队列改存指针、payload 放 PSRAM，并验证所有释放路径 |

排障顺序不可颠倒：devicesim 平台基线 → 报文/加密向量 → Watcher 固件 → 麦克风、
扬声器和显示硬件。
