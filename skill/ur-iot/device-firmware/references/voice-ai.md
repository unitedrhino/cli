# 联犀设备语音 AI 从零接入

本指南用于把已经具备联犀 MQTT、物模型与 OTA 能力的设备接入语音对话、表情和
当前设备控制。协议行为以 `backend/things/tools/devicesim` 为唯一基准：先证明平台
可用，再验证固件协议，最后验证麦克风、扬声器和显示硬件。

设备还需要拍照识图、图片输入或表情可见期时，同时读取
[设备拍照识图](photo-vision.md)。

## 子文档导航

本指南原第 2–6 节的执行细节已按主题拆分到以下子文档，章节编号沿用原文：

- [测试基线](voice-ai-test-baselines.md)：第 2、3 节。devicesim 平台回归、模型与会话边界排查、固件单元测试与协议 E2E。
- [会话协议与音频](voice-ai-session-protocol.md)：第 4、5 节。MQTT 会话时序、打断与 VAD、队列与内存约束、停顿后恢复、UDP 封包与音频参数。
- [真机音频 E2E](voice-ai-hardware-e2e.md)：第 6 节。opusfixture 资源生成、分层用例表、runner 断言与延迟门禁、多轮参数。
- [拍照与控制复测](voice-ai-photo-control-e2e.md)：第 6 节专项复测。语音拍照与语音控制的真机自动复测入口与证据要求。

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

失败时按本轮session/resp关联实际工具：ASR已有设置意图却只调用
`get_device_properties`、随后文字和音频完成，属于未选择控制工具，不能误判为设备
下发超时。只有观察到`device_property_control`后，才继续定位平台下发、设备回复及上报。
保留原样本与失败断言，核对模型收到的当前输入、工具定义和有效配置；不要直接加固定
控制提示词、延长等待或以口头成功放行。`RunVoiceTurnSTTOnly`的FinalText为空是接口
行为，不证明服务器未回复；需另外核对同轮respTextDone/respAudioDone。

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
