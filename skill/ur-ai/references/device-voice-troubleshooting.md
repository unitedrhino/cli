# 设备语音：分层排障与典型故障

> 本文件是 [device-voice.md](device-voice.md) 的子文档。收录：分层排障表与三类典型故障处理；基线用例命令见主文档。

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
