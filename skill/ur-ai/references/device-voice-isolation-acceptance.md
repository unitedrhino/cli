# 设备语音：有声首音与隔离验收

> 本文件是 [device-voice.md](device-voice.md) 的子文档。收录：有声首音与可重复隔离验收、固件协议单测与真机云端音频闭环。

## 有声首音与可重复隔离验收

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

## 真机云端音频闭环

真机云端音频闭环应使用
`firmware/watcher/scripts/run_ur_ai_e2e.py --port <serial-port> --repeat 5 --timeout 75`。
它使用由 devicesim 编码器生成的 16 kHz/60 ms Opus 样本，只替换麦克风输入，仍连接真实
MQTT、UDP、ASR、LLM 和 TTS。每轮必须看到完整方法序列、播放队列排空与 PASS，
并使用 ESP 日志 uptime 验证 8 秒/15 秒门禁。生产固件不得启用该内置样本。
音频生成、测试构建和失败分层的完整步骤见
`ur-iot/device-firmware/references/voice-ai.md`。
