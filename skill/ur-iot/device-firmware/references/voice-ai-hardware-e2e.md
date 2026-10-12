# 真机音频串口 E2E

> 本文件是 [voice-ai.md](voice-ai.md) 的子文档，收录原指南第 6 节：真机串口循环 E2E 的分层用例、音频资源生成、runner 断言与延迟门禁。

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
| `VOICE-SIM-IDLE` 长停顿 | `TestVoiceLongIdleRecovery`，建议 `-count=3` | 完整首轮结束后静置35秒，同session第二轮文本、有声帧、同respId音频起止；不重建会话 | 停顿后无回复或服务变更后 |
| `VOICE-HW-001` 单轮闭环 | runner `--repeat 1`，固定语料“当前音量多少” | 全部阶段、STT 命中“音量”、七个 RESULT 标志、两项延迟和无致命标记 | 每次测试固件刷写后 |
| `VOICE-HW-002` 重复稳定性 | runner `--repeat 5` 或更高 | 每轮独立 session、全部通过、各轮原始日志和 JSON | MR 前至少五轮 |
| `VOICE-MANUAL-001` 硬件验收 | 真人唤醒、说话并观察/听取设备 | 麦克风、唤醒词、扬声器、字幕和表情实物效果 | 发版与现场验收 |

长停顿用例已纳入默认 `shell/devicesim-oneclick-test.sh`。离线判定器
`TestIdleVoiceReplyEvidence` 用合成消息拒绝只有STT、静音/旧有声点、串轮音频起止及
过期终态，可竞态十次复测；它不复现实际服务故障。真实平台回归须检查解码有声帧，
不能用“识别了但没有回复”判通过；通过只定位到平台链路，不代签Watcher麦克风、
本地VAD、息屏及实际声学行为。环境认证和模型从环境文档注入，不为此改服务配置。

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
   非空文本与音频不能证明模型成功：平台可能将模型服务错误作为普通文本播报。
   测试专用固件须拒绝平台固定“AI 服务本次调用失败（”提示并输出
   `model_service_error`；拍照即使命中目标物品词也不能放行。
   覆盖“完整阶段＋错误提示”的失败用例及正常回复、普通错误码讨论、跨轮隔离。
   此判据仅识别已知平台错误，不替代语义正确性与现场物品识别验收。
4. 日志不含 `Failed to resample output audio`、`assert failed`、`Guru Meditation`、
   `Backtrace`、复位或命令错误；重采样容量异常即失败，不能等堆断言。
   `AES-CTR operation failed` 或 `esp-aes: Failed to allocate memory` 同样必须失败：
   丢帧后仍可能收到完整终态与 PASS，不能据此证明播放流畅。该门禁由单轮、多轮、
   拍照与按键 runner 共用，新增失败用例需验证“阶段齐全＋AES 错误”不会假通过。
5. 以 ESP 日志的统一 uptime 计算门禁：AudioStop→首帧小于 8 秒，
   STTDone→TextDone 小于 15 秒。不能相减计时原点不同的 `elapsed`。

每轮输出 JSON 摘要和单独原始日志；任一轮缺阶段、超时、崩溃或复位即非零退出。
语音runner自行限制新建目录0700/原日志0600，不依赖调用者umask。每轮在发命令前
排他创建日志，同秒同目录冲突会拒绝执行，不覆盖旧成功或失败。已有目录的权限
不会被自动修改，复用时先核验归属与权限，不批量收紧共享父目录。
循环 E2E 只证明云端到扬声器队列的可重复闭环；麦克风、唤醒词和实际扬声器响度
仍需最终人工真机验收。

TASK-143 新测试固件支持 `ur_ai_e2e 3`；对应 runner 增加 `--turns 3 --repeat 1`，
在一个 session 内完成三轮，而不是把三个独立会话冒充多轮。`--turns` 范围 1–5，
缺省 1 保持兼容；`--timeout` 为单轮预算，会话采集时限乘以 turns。
每轮等待正式播放排空及恢复监听路径，要求相同完整 session ID、新的成功 audioStart token，
独立校验 STT 关键词、文本、真实音频和延迟；首个失败即关闭会话，整组期间隔离现场麦克风。
测试日志只输出轮数与布尔判据，不输出完整 token/session、音频或对话。
`test_ur_ai_multiturn.py` 编译生产入口/快照，覆盖旧 token、会话重建、中途失败及清理；
主机通过仍须在支持该参数的新固件上实测，旧测试固件可能忽略参数，runner 会因缺轮拒绝通过。
同一句固定样本仅验证多轮音频/状态流程，不能替代不同问题的上下文记忆或打断验收。

