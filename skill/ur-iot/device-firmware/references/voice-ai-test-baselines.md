# 语音 AI 测试基线（devicesim 与固件）

> 本文件是 [voice-ai.md](voice-ai.md) 的子文档，收录原指南第 2、3 节：devicesim 平台基线回归、模型与会话边界排查，以及固件单元测试与协议 E2E。

## 2. 先运行 devicesim 平台基线

对照真实设备前，显式设置`DEVICESIM_LLM_CONFIG_ID=<模型配置ID>`，并回读模拟Agent。
`BootstrapRequest.LlmConfigID`优先，环境变量只补缺；两者为空才沿用平台默认，
指定已有AgentID不会改写它。默认提供方额度失败不能推导真机同配置失败；切换后
仍须原样复跑失败音频，不用更简单的输入替代。模型选择可用
`go test ./things/tools/devicesim -run '^TestBootstrapLLMConfigEnvironment$' -count=10`
验证HTTP请求边界（在backend目录运行，不连接真实平台）。

控制异常需要区分共享测试设备的历史和会话切换。保留原失败用例，再用
`TestVoiceControlSessionTransition` 对照全新设备的首次语音、先文字查询再切语音，
以及完整复用原用例助手的`original_workflow`；均使用原音频和相同提示词，
仍要求真实控制下行目标40及上报后的查询。第三场景保持原STTOnly收集及等待门槛，
用于排除收集器差异，不能以新设备通过直接认定旧设备历史污染。
在已授权的测试环境、显式模型及现有认证下从backend目录执行：

```bash
go test ./things/tools/devicesim -run '^TestVoiceControlSessionTransition$' -count=1 -timeout 420s
```

新设备通过不代表原故障已修复，仍须同配置复测原失败设备并比较有效模型、Clone历史和工具。
怀疑历史覆盖当前要求时，用真实Runtime模型边界测试
`TestRecentDeviceInputControlModelBoundary`检查最新user消息、长历史压缩及跨轮隔离：
从backend目录运行`go test ./core/service/aisvr/internal/domain/agentruntime -run '^TestRecentDeviceInput' -count=10`。
它只替换外部模型，不能证明供应商工具选择或真实设备执行成功；不要据它清除旧历史。
该用例使用唯一prefix并清理本次新建的设备、产品和Agent，不删除共享MCP。
`NameSuffix`当前不参与资源名称，不能据它宣称隔离；`CreatedDevice`表示AC自动清理策略，
不是新建事实。非AC的`CleanupDevice`可能不删除；仅确认产品和Agent均为本次新建、
设备属于该新产品后才调用`CleanupAll`，不能据布尔字段清理已有共享资源。

在仓库根目录执行完整回归。凭据通过已有 profile 或环境变量注入，不写入命令、文档或日志：

`devicesim-oneclick-test.sh`自行以umask077新建0700目录/0600原日志，不依赖调用者权限；
同秒同路径或预置链接会在探测前失败。输出持有独占打开的描述符，后续路径替换也不
重定向证据；已有目录不批量chmod，应先核对归属和权限。原日志可能含对话/诊断，
受限保存不是脱敏，分享前仍须脱敏。离线复测运行
`python3 shell/test-devicesim-oneclick.py -v`：实际Bash入口七项文件系统检查，覆盖
权限、冲突、正常/悬空链接和运行中路径竞争，外部命令隔离，不作为平台E2E通过证据。

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

