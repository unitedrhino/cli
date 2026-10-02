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

表情到达时设备应清除相机预览并唤醒屏幕，保持至本轮真实音频播放队列排空；恢复监听后
再回到 `neutral`。晚到表情需要最短可见期，不能被 listening/speaking 状态切换立即覆盖。

## 拍照识图与图片输入

视觉设备的 Agent capabilities 包含 `text`、`voice`、`image_input`。LLM 与 vision 配置
都指向真实支持 `text`、`image` 输入的模型；若图片流式生成存在已知编码问题，图片路由
使用同步生成，普通文字和语音保持流式。先通过模型配置测试接口用确定性图片证明真实
识图，再修改 Agent，不能回退纯文本模型冒充成功。

平台内置拍照工具由 fuzai MCP 提供。按环境 API 幂等查询或注册服务并回读真实 ID，不在
脚本或技能中硬编码历史 ID。fuzai 运行环境必须启用内置 MCP、Redis 和 DmRpc；Agent 在
保留已审核 IoT MCP 的同时显式绑定 fuzai。所有变更走现有 API，不直接写数据库。

服务范围、企业身份、空列表与工具发现的检查见
`device-firmware/references/photo-vision.md`的“平台与物模型”；平台私有ID可见不等于
设备企业运行时可加载。模拟模型优先级、无结果轮次及控制工具分层检查见
`device-firmware/references/voice-ai.md`，不要用默认模型对照未核验的真机配置。

`sessionCreated.supportedModalities` 包含 `image`，并返回当前 session 的短期 `uploadUrl`。
图片上传成功后有两条等价入口：

- 语音意图：`deviceTakePhoto` → 设备 `takePhoto` → `actionReply.data.fileUri` → 模型识图。
- 本地双击：设备上传后发送含 text 与 `image_url.imageUrl` 的 `inputSend`，并把
  `modalities` 设为 `text`、`audio`。

完整固件回执、上传安全、按键和真机步骤见
`device-firmware/references/photo-vision.md`。

设备业务不得进入 core：fuzai 返回标准 MCP `image(data,mimeType)` 内容块，
文件引用仅为文本元数据；core 的通用适配器传递结构化图文并复用多模态转换和视觉路由，
不解析工具文本中的私有图片字段、不判断工具名、不附加拍照提示词。
同一轮继续推理，内联图片不写历史/事件；core 与 fuzai 配套升级或回滚。
回归需同时覆盖与拍照无关的 MCP 图表工具（并发关联、跨轮隔离、无效图片）以及
真实模型识图（断言文字、颜色、形状），不能只检查 URL 或模型口头确认。

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
DEVICESIM_TEST_PATTERN='Test(MultiTurnVoiceChat|VoiceInterruptDuringTTSThenContinue|VoiceCancelDuringTTSThenContinue|VoiceTurnWithoutAudioStopStillGetsSTT|TextToTTSAudioEvents|EmojiEmotionText|EmojiNotSentForPlainQuestion|DeviceControlSuccess|TakePhotoEndToEnd|ImageInputEndToEnd)$' \
DEVICESIM_AUDIO_SAMPLE_RATE=16000 \
bash shell/devicesim-oneclick-test.sh
```

重点用例：

| 用例 | 验证目标 |
|---|---|
| `TestMultiTurnVoiceChat` | 同 session 多轮与上下文 |
| `TestVoiceInterruptDuringTTSThenContinue` | 直接 audioStart 打断播报并继续 |
| `TestVoiceCancelDuringTTSThenContinue` | 实体按键等价的 respCancel → audioStart，并继续多轮 |
| `TestVoiceTurnWithoutAudioStopStillGetsSTT` | 服务端异常恢复能力 |
| `TestTextToTTSAudioEvents` | TTS 真实音频事件 |
| `TestEmojiEmotionText` | 表情白名单和 respId |
| `TestEmojiNotSentForPlainQuestion` | 普通问题不误发情绪事件 |
| `TestDeviceControlSuccess` | MCP 到物模型回复闭环 |
| `TestTakePhotoEndToEnd` | MCP 拍照行为、可下载 fileUri 与真实识图 |
| `TestImageInputEndToEnd` | 双击等价的 image_url 多模态输入 |

保存脱敏的方法序列与阶段耗时。延迟门限建议为 AudioStop 到首个音频帧小于 8 秒、
STTDone 到 TextDone 小于 15 秒。

### 有声首音与可重复隔离验收

测“用户说完→回复首音”时运行`TestVoiceFirstTurnGreetingReplyHasAudio`，不能用
`respAudioStart`或首个静音UDP包替代。有声输入结束取最后一个超过RMS阈值的60ms帧，
回复首音要求连续两个有声解码帧，结果日志标明`endpoint=simulator_decoded_audio`。
`input_end_to_first_voiced_reply`包含输入尾部静音和audioStop等待；它不包含真实设备
麦克风、播放缓冲或扬声器声学延迟。现场首字需另用同时录下输入与设备回复的音频测量，
少量重复结果不报告为P95。

在已授权测试环境加载受限认证后，从SaaS仓库根目录复测；遵循仓库构建位置与资源门禁，
不把Token、供应商正文或完整MQTT参数打印出来：

```bash
DEVICESIM_AUDIO_SAMPLE_RATE=16000 bash shell/remote-build.sh run --kind backend --scope backend --timeout 900 -- \
  bash -lc 'cd backend && go test ./things/tools/devicesim -run "^TestVoiceFirstTurnGreetingReplyHasAudio$" -count=3 -timeout 720s'
bash shell/remote-build.sh run --kind backend --scope backend --timeout 180 -- \
  bash -lc 'cd backend && go test -race ./things/tools/devicesim -run "^Test(FirstVoice.*|LastSpeechFrameIndex|FrameRMS|OpusAudioReceiverRequiresConsecutiveSpeechFrames|RepeatedOpusSilenceRemainsSilent|SilenceAfterSpeechRemainsSilent)$" -count=10 -timeout 90s'
```

该首音用例使用随机产品/Agent前缀，退出时以独立30秒上下文执行`CleanupAll`。
当前devicesim的`CreatedDevice`是旧引导器的清理策略，自建产品时为false，不代表设备已存在；
完整清理以本次新建产品和Agent归属为保护条件，不修改公共复用用例或删除历史固定前缀资源。
清理失败同样算E2E失败。定位失败残留时先核对精确名称、创建窗口和关联关系，再仅回收
本次资源；空列表可能为null，读取时不能把验证器异常当成删除失败并盲目重复删除。

平台基线通过后，还要在 `firmware/watcher` 运行固件生产协议核心的单元测试与时序回放：

```bash
python3 -m unittest \
  scripts.tests.test_ur_ai_contract \
  scripts.tests.test_ur_ai_runtime \
  scripts.tests.test_run_ur_ai_e2e -v
```

其中静态合同检查只能发现源码合同漂移；`test_ur_ai_runtime` 才会编译运行固件共用的
表情、session/respId、终态去重和 UDP 帧头代码，并回放多轮、乱序、重复、打断和断线
旧消息。该协议 E2E 不连接云端，不能替代上面的 devicesim 真实 E2E 或最终真机测试；
三层结果应分别记录。

排查“识别到控制要求但模型只查询”时，增加模型边界合同测试，不直接清历史或修改提示词：

```bash
bash shell/remote-build.sh run --kind backend --scope backend --timeout 300 -- \
  bash -lc 'cd backend && go test ./core/service/aisvr/internal/domain/llm -run "^(TestTracedTransport.*|TestModelRequestPreservesCurrentInput|TestMediaGuard|TestVisionRouter.*)$" -count=10 -timeout 120s'
```

该测试使用真实模型工厂、门禁、路由和OpenAI适配器，将固定夹具发送到本地HTTP模拟
供应商，核验纯文本、查询历史、召回上下文和工具结果之后的消息顺序与工具定义。
供应商响应是模拟值，不能证明真实模型会选择控制工具，也不能替代失败的原始devicesim
用例。Runtime模型调用前的摘要仅能证明该边界的输入；不带会话/trace关联的摘要不能
单独用于跨并发定位。继续核对实际模型配置、MCP清单、工具选择和真实控制回执。
模型HTTP时序诊断不得预读、关闭或以旧GetBody覆盖当前请求体，不落盘请求JSON；
正文可能含对话、工具参数及内联图片。旧诊断文件需另行核实留存和删除授权，不清空共享目录。
旧provider测试若依赖硬编码外部凭据，不借其发起真实调用，也不把所选离线集合称为整包全绿。

进一步区分ASR措辞和输入通道时，可运行`TestVoiceVolumeRecognizedTextReplay`：
沿用共享模拟设备、原音频和提示词，取得真实STT后逐字经文字入口发送，并核验控制下行
和后续属性查询。使用上面的资源入口，把`-run`改为`^TestVoiceVolumeRecognizedTextReplay$`、
`-count=1`，包路径改为`./things/tools/devicesim`，测试预算改为240秒。
语音轮已经控制成功时该诊断明确跳过，跳过不计通过。测试会正常追加会话历史和模拟属性，
不清理共享历史；文字重放通过后仍须原样复测`TestVoiceVolumeWorkflow`，不能据此关闭
语音缺陷或认定某条历史记忆是根因。结果分开记录，未经关联的迟到控制不能算重放成功。

上述重放已经在共享Clone追加语音轮，不能单独归因于通道。更严格的对照使用
`TestVoiceVolumeTextControlAfterQuery`：仅在独占的新资源上取得真实STT，共享设备只做
查询→逐字文字控制→查询，并回收识别来源。再对照`TestVoiceControlSessionTransition`
的全新资源场景；若共享文字也失败而隔离通过，应继续检查有效上下文与模型选择，
不要归因于固件采音，也不能清共享历史、修改提示词或放宽控制门槛来获得通过。

若实际提示的稳定画像不同，可运行`TestVoiceControlStableProfileReplay`（同一资源入口，
`-count=1`、devicesim包、420秒预算）：只读固定共享模拟资源的画像，内存中经现有
记忆API写入本次独占分身，空画像与重放画像分别运行原控制流程并自动回收。
不读取真实用户画像、不输出正文、不改共享历史。标准API也创建来源记录，不能称为
仅替换系统提示词的纯变量实验；隔离通过只说明在这组新资源条件下未复现，仍须检查
共享动态召回/历史组合并原样复测失败用例，不能凭画像存在就修改生产过滤规则。
同一入口可改跑`TestVoiceControlRecallContextReplay`，在独占分身重放来源有效Dream摘要，
核对后台实际召回及控制下行。它不复制来源ID、历史时间、权重或访问计数，标准API也会
合并目标画像，所以只是组合对照，不是等价生产快照。记忆轨迹ID可能包含查询正文，
脱敏须取哈希，不能直接输出或只截短前缀。
需要覆盖冻结画像与新动态召回时，成对运行
`TestVoiceControlFrozenRecallContextReplay|TestVoiceControlFrozenProfileReplay`：
先用完整语音查询预热，再分别写入摘要或不写；核对实际控制轮画像块哈希，而非仅回读
持久化画像。预热增加短时原文，摘要API也改变中间文字查询画像，所以仍不属于纯变量
实验。摘要记录可能正文重复，实际格式化去重后的条数才是注入条数；红色复现必须保留，
不得用刷新缓存、清历史或隔离对照通过替代原共享失败的修复。
召回格式调整需先用`TestFormatPromptMemoryContext`覆盖背景不冒充系统指令、当前请求及
授权优先、原内容保留与空结果不注入；这类单测只验证格式，不证明模型工具选择。
候选发布后应分别原样复测共享控制（至少重复运行）、暖缓存摘要组合及画像对照，
并保留冷启动矩阵回归。共享用例重复通过而暖缓存组合仍失败时，只能记局部改善；
进程重启可能改变工具数组次序，须同时核验集合与有序指纹，不能据一次通过确认根因。

重放摘要时先通过现有`memoryKind=dreamSummary`、`status=1`服务端过滤，再校验总数及
分页ID唯一性。列表按更新时间排序时，无关记录的后台访问可能移动跨页边界；准备阶段的
重复ID失败不等于控制失败，也不得静默丢弃重复项后宣称读取完整。过滤后的来源内容须与
原目标集合一致，保留原失败记录后复测，不修改平台接口或共享摘要。

画像已成功写入但设备仍读旧值时，检查写入路径是否调用已有设备快照失效通知。
`TestMemoryCreateRefreshesDeviceSnapshot`经真实记忆创建logic及数据库覆盖预热→写入→
voice/MQTT回读，并保留HTTP同会话冻结及其他分身隔离断言（包为
`./core/service/aisvr/internal/logic/ai/clone/memory`）。只在成功写入后定向失效是缓存一致性
修复，不等于为通过测试手工清缓存；不得移除历史或放宽原控制断言，也不能以该单测
证明暖缓存组合的真实模型工具选择已恢复。

核对工具时注意：当前`/api/v1/ai/mcp/tools/get-tools`读取企业启用服务的缓存工具清单，
实现未按`sessionID`筛选，也不返回Runtime实际绑定的完整参数Schema。两会话该响应
相同不能证明装配相同；应结合Runtime的服务绑定、实时`tools/list`与实际工具调用取证。
启用无正文Runtime诊断时，用会话哈希关联`agentruntime.input/tools`：仅在
`fingerprint_valid=true`时比较数量及`set_hash`；集合摘要相同、有序摘要不同只说明
工具数组次序不同。指纹覆盖实际`WithTools`的名称、描述与参数定义（含前端工具），
不输出原文，但它仍不是供应商实际HTTP请求证明；诊断未部署时不能拿本地单测替代运行取证。
音频相同也不保证STT相同；跨场景对照先在内存比较真实STT并仅记录长度/哈希/相等性。

真机云端音频闭环应使用
`firmware/watcher/scripts/run_ur_ai_e2e.py --port <serial-port> --repeat 5 --timeout 75`。
它使用由 devicesim 编码器生成的 16 kHz/60 ms Opus 样本，只替换麦克风输入，仍连接真实
MQTT、UDP、ASR、LLM 和 TTS。每轮必须看到完整方法序列、播放队列排空与 PASS，
并使用 ESP 日志 uptime 验证 8 秒/15 秒门禁。生产固件不得启用该内置样本。
音频生成、测试构建和失败分层的完整步骤见
`device-firmware/references/voice-ai.md`。

## 分层排障

| 层级 | 判定方法 | 处理方向 |
|---|---|---|
| 平台 | devicesim 也失败 | Agent/模型/MCP 配置、ASR/LLM/TTS、UDP 服务 |
| 协议 | devicesim 通过，真机无 STT | MQTT token、UDP 预热、AES nonce、序号与 Opus 参数 |
| 播放 | 有文本无声音 | TTS 是否真实产帧、UDP 下行、Opus 解码、扬声器 |
| 堆损坏 | 完整回复后 assert、随后表现为黑屏或网络错误 | 核对是否错误使用 16→24 kHz 外部重采样；Watcher 保持 16 kHz 会话并由 Opus 解码器直接输出 24 kHz PCM |
| 控制 | 有回复无设备变化 | AgentGroup、MCP 绑定、物模型 identifier、reply 合同 |
| 生命周期 | 多轮或重连串话 | session/respId 关联、旧消息、断线后旧 UDP 未关闭 |
| 资源 | OTA 后平台离线且 MQTT 分配失败 | 检查音频初始化后的内部 RAM；大报文队列只存指针，payload 放 PSRAM并完整释放 |
| 视觉 | devicesim 识图失败 | 视觉模型输入模态、同步图片路由、fuzai/Redis/DmRpc 与 MCP 绑定 |
| 拍照 | 收到工具调用但无结果 | takePhoto 物模型、上传会话、actionReply fileUri 与服务端等待窗口 |
| 表情显示 | 表情一闪即逝 | UI 状态提前恢复 neutral，未等待真实播放队列排空或晚到事件无最短可见期 |

平台排障以 devicesim 为第一层 oracle；真机只负责验证固件时序和硬件链路。

### 取消后新轮无回复

先核对最新远端主线与部署二进制，再用上述两种打断测试分别复现。旧轮取消错误不等于
服务器 Bug；需要确认新轮的 UDP、ASR、MQTT 下行以及 loop 退出时间。
audioStop 后若识别结果仍在 hold、队列或记忆准备阶段，不能因为 LLM/TTS 尚未启动
就关闭父 context。回复工作从入队前计数，到消费结束释放；计数按 loop 隔离，
空 Final/关闭通道不得越过待处理回复直接退出，超时与用户关闭仍须能够释放资源。

### 前段语音已播，工具返回后一直等待

先关联同一 session/respId，区分“合成前模型等待五秒”和“首帧之后三秒帧间超时”。
已有真实音频后，工具执行与后续模型待输入可能暂时没有可合成文本，不能把它判为
供应商停滞。使用本轮独占等待状态及时通知播放循环；收到有效文本或输入结束后恢复
原帧间预算，等待仍受请求取消约束，不按工具名称、图片私有字段或固定提示词判断。
真正TTS失败不能伪造respAudioDone，应通过既有error关联当前respId及时结束设备等待，
固定提示不得包含供应商错误正文。旧轮取消不向新轮反馈错误。

从SaaS根目录运行确定性事件/音频回归（无真实模型和相机），再做平台和真机复测：

```bash
bash shell/remote-build.sh run --kind backend --scope backend --timeout 600 -- \
  bash -lc 'cd backend && go test -race ./core/service/aisvr/internal/domain/chat -run "^Test(TTSInputWaitTransitions|AudioPacingToolWaitCancel|ConsumeStreamEvents_(ToolWait.*|ReliableFailureExplicitError))$" -count=10 -timeout 180s'
```

失败用例必须先证明前段实际已发送、工具等待超过原预算、后段文字仍完整但音频丢失；
修复后同时断言前后音频、单次开始/正常完成。负例覆盖重复通知、跨轮隔离、取消、
输入结束后真实停滞与明确失败。随后原样重复TestVoiceTakePhotoEndToEnd及真机串口
拍照，要求真实有声帧与同轮终态；单测转绿、只有文字或首帧都不能关闭整机缺陷。

### ASR报错但设备一直等待

若是建立会话后先播放文字回复、再取消播报开麦，先比对ASR建连与audioStart时间。
豆包`45000081`表示等包超时，不能当作额度不足；仅建立会话/播放文字或图片回复时
不应提前消耗识别连接预算。显式监听保留首句预建连，尚未开麦则沿用有效音频首帧懒启动，
不得吞掉真正识别失败或提高供应商超时来掩盖生命周期错误。
`TestVoiceLoopASRStartsOnlyWhenListening`同步断言生产ASR句柄：未开麦时不启动、
开麦后真实Opus首帧启动并提交、显式监听仍预建连。与ASR失败/恢复和UDP退出清理用例
一起做竞态十次及完整chat包，再原样执行`TestVoiceCancelDuringTTSThenContinue`；
本地转绿不等于运行后端已更新，也不能拿普通多轮通过代替取消后继续验收。
若该修复后原用例仍红，继续检查是谁调用了初始OnListen：UDP路由预热也可能隐式
调用OnAudioStart，令循环误判已正式开麦。sessionCreate会话的预热只能建立路由，
不能产生监听意图；未声明显式会话的旧端才保留首包补监听，且不能覆盖先到的开麦参数。
`TestUDPPrewarmRequiresExplicitAudioStart`与`TestUDPPrewarmLegacyFallback`同步覆盖
新建/恢复、开麦与预热先后、旧端兼容和空目标；再跑chat/UDP整包及两种真实打断E2E。
不得删除静音预热、吞供应商错误或让测试忽略internalError来凑绿；过期实例后续恢复
不能抵消原轮失败。该边界不改变UDP加密报文、不新增HTTP API或数据库字段。

先关联短会话ID，核对UDP接收、解码、VAD送出与ASR送出帧数，以及供应商建连和流内错误。
送出音频而没有STT不能直接归因于麦克风或网络；供应商数字错误码也不能单独证明额度不足。
识别层的错误不能只关闭结果通道：应通过既有`error/internalError`反馈固定提示，屏蔽供应商
原始正文，并忽略取消上下文及过期识别实例。后续音频保留原识别重启和通道关闭所有权。

离线回归运行chat包的`TestASRFailureNotifiesDevice`、
`TestASRFailureIgnoresCancelledAndStaleRun`、`TestASRFailureRecoversOnNextAudio`，
覆盖建连失败、流内错误、取消/旧轮隔离和同会话新Opus帧恢复识别及文字回复；通过资源入口
分别重复十次、运行竞态检测及完整chat包。这些用例使用供应商/LLM/TTS替身，不代替真实
平台和设备稳定性验收。真实测试仍要求完整STT、文字、音频和播放终态，收到错误提示不能
算成功对话；保留原失败窗口，不因错误现在可见就放宽门禁或将新窗口与旧窗口拼接。
