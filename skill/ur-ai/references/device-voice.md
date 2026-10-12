# 设备语音会话、表情与 MCP 控制

本指南描述联犀平台侧设备语音助手的配置与验证。设备固件实现、刷写、UDP 加密和
真机验收同时读取 `ur-iot/device-firmware/references/voice-ai.md`。

## 子文档导航

- [device-voice-isolation-acceptance.md](device-voice-isolation-acceptance.md)：有声首音、隔离复测、固件协议单测与真机云端音频闭环。
- [device-voice-model-boundary-tests.md](device-voice-model-boundary-tests.md)：模型边界合同测试与工具响应关联/往返核验。
- [device-voice-control-replay-tests.md](device-voice-control-replay-tests.md)：语音控制失败的重放与隔离对照用例。
- [device-voice-control-completion-policy.md](device-voice-control-completion-policy.md)：控制完成策略候选与配置助手人设测试。
- [device-voice-runtime-diagnostics.md](device-voice-runtime-diagnostics.md)：记忆召回、设备快照与运行时诊断取证。
- [device-voice-troubleshooting.md](device-voice-troubleshooting.md)：分层排障表与典型故障处理流程。

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
`ur-iot/device-firmware/references/photo-vision.md`的“平台与物模型”；平台私有ID可见不等于
设备企业运行时可加载。模拟模型优先级、无结果轮次及控制工具分层检查见
`ur-iot/device-firmware/references/voice-ai.md`，不要用默认模型对照未核验的真机配置。

`sessionCreated.supportedModalities` 包含 `image`，并返回当前 session 的短期 `uploadUrl`。
图片上传成功后有两条等价入口：

- 语音意图：`deviceTakePhoto` → 设备 `takePhoto` → `actionReply.data.fileUri` → 模型识图。
- 本地双击：设备上传后发送含 text 与 `image_url.imageUrl` 的 `inputSend`，并把
  `modalities` 设为 `text`、`audio`。

完整固件回执、上传安全、按键和真机步骤见
`ur-iot/device-firmware/references/photo-vision.md`。

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

## 物模型 AI 开放范围

只有物模型 `isCanSceneLinkage=1` 的能力开放给设备 AI；`2`（关闭）、`0`（未明确开启）
和未知值均不开放。先聚合设备与产品定义，再过滤，设备级关闭优先于产品级开启。
只读属性可以查询但不能写入，上行行为不可下发。普通 App 的完整物模型及现有管理权限
不因此删除或收紧。

排查“AI 不知道某个属性/行为”时，核对产品与设备覆盖项的开关、当前会话提示词、
`get_thing_model` 返回范围和实际工具执行回执。修改开关后重新开始会话以刷新提示词；
不能仅在提示词中隐藏，属性读取、属性控制和行为下发也要使用同一开放范围。
物模型查询失败必须拒绝执行；读取空列表只能查询已开放属性，空白名单不能变成全量查询。
fuzai `deviceTakePhoto` 也须校验聚合物模型中的 `takePhoto`，MCP 工具名注册可见不代表
当前设备行为已授权；双击拍照的本地用户操作不由 AI 工具开放开关决定。

修改过滤逻辑时先跑失败用例，再验证 `TestForAISceneLinkage`、
`TestForAIKeepsDeviceOverride`、`TestSchemaSceneLinkageRoundTrip`、
`TestTSLPromptSceneLinkage`、`TestMCPSceneLinkageIntegration` 和
`TestTakePhotoSceneLinkageIntegration`。MCP 集成用真实 HTTP 工具路由、内存 Redis 和
公共 RPC mock 覆盖开启/关闭/未配置、设备覆盖、只读、上行、查询失败及空白名单；
断言关闭项未进入模型结果、敏感值未泄漏、禁止项没有实际 RPC 下发。
这些模拟验证不能替代真机拍照、语音和控制回执验收。

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

进阶验收与诊断用例见文首「子文档导航」所列子文档。
