# 5. 可重复自动测试

本文件是 [photo-vision.md](photo-vision.md) 的子文档，收录其第 5 节「可重复自动测试」的完整内容：devicesim 视觉 E2E、Go 定向测试与固件单测命令。

先建立已有设备的保护基线：核对正常设备与目标设备的后端版本、模型配置和报文时序，
分别记录已验证的基础对话/控制与尚未验证的视觉能力；不同环境不能默认是同一版本。
DMA、采样率、AFE 句尾和按键状态机问题先在固件定位；后端修复应有独立失败用例，
不能用新设备异常或一次视觉 E2E 成功推导生产兼容性。

MCP 图片扩展只接管明确的 `image` 内容。无图片时旧 ChatSession 保持首块分派，
Runtime 保持原有文本块合并与无文本结果序列化；这两种入口不能互相替代。
必测 `TestMCPLegacyContentCompatibility`、`TestMCPRuntimeLegacyResults`，
覆盖音频/资源在前、文本在前、多文本、普通控制及无文本结果。
目前图片与音频/资源混合执行尚不支持，须显式报错，不能静默丢弃其他媒体。

先验证通用适配边界：从仓库 `backend` 目录执行以下定向测试。
图表工具用例与拍照无关，覆盖纯文本兼容、图文、多图片、非法/超限图片、
工具错误、并发 toolCallID 关联和跨轮媒体隔离；不得改成识别特定工具名才能通过。

```bash
go test -race ./core/service/aisvr/internal/domain/chat \
  ./core/service/aisvr/internal/domain/agentruntime \
  ./apps/fuzai/internal/domain/mcp \
  -run 'TestMCP|TestPhotoStandardMCPResult' -count=10
```

该测试不访问模型，不能冒充真实识图。后续必须运行下列平台 E2E，
并回归两种语音打断（见 [语音 AI](voice-ai.md)）。

先运行平台基线。测试凭据使用已有 profile、环境变量或权限不宽于 `0600` 的受限文件：

启动前以 `ur --app <应用> check --json` 核验实际环境、应用、组织与认证；需交给 Go 测试时，
通过同一应用的 `ur token --raw` 在内存取得 Token，再以环境变量传入，不打印 Token。
不要直接读取旧 `~/.ur/config.json` 推测当前 CLI 上下文：多应用配置可能已不同。
找不到 fuzai 或更新测试产品返回权限不足时，先核对上述上下文与资源归属，不能直接断言
后端缺功能、重新注册 MCP 或扩大权限。

视觉四项用例使用每次独立的随机前缀，并断言 `CleanupAll` 成功；已有 `CleanupDevice`
为固定前缀复用场景保留部分资源，不适用于一次性隔离资源。清理仅针对本次新建对象，
不得按 sim 前缀批量删除历史资源。还应运行 `TestImageInputDuringVoiceAudioStop`，
验证图片输入与停止收音交错时仍有完整语音，而不只是文本通过。

```bash
DEVICESIM_TEST_PATTERN='Test(TakePhotoEndToEnd|VoiceTakePhotoEndToEnd|ImageInputEndToEnd|EmojiEmotionText|EmojiNotSentForPlainQuestion)$' \
DEVICESIM_AUDIO_SAMPLE_RATE=16000 \
bash shell/devicesim-oneclick-test.sh
```

测试图片必须是仓库自有、内容固定的 JPEG，并断言模型实际识别其文字、颜色和形状；仅断言
“收到一段回复”不够。`TestTakePhotoEndToEnd` 验证工具调用、action、可下载 `fileUri` 和
识图回复，`TestImageInputEndToEnd` 验证双击等价的 `inputSend(image_url)` 路径。

文字拍照不能替代语音入口。`TestVoiceTakePhotoEndToEnd`使用可复生成的
`turn16_take_photo.mp3`，以16kHz Opus经真实MQTT/UDP完成ASR、工具与识图，
还要求解码到有声帧并在本用例45秒终态预算内收到相同respId的`respAudioDone`。
识图内容或首帧通过但缺少终态仍失败，保留方法序列继续排查，不延长等待掩盖缺事件。
它由模拟设备上传确定性图片，不证明真实相机；Watcher仍需独立真机语音拍照验收。

固件侧运行完整测试发现入口；关键纯逻辑用例至少连续 10 次：

```bash
cd firmware/watcher
python3 -m unittest discover -s scripts/tests -v
```

固件测试至少覆盖：URL 解析、同源约束、JPEG 校验、认证副本清零、平台响应、行为回执、
token 去重、相机互斥、队列拥塞、超时、断线、OTA 拒绝、图片信封、按键互斥、21 类表情
资源一致性、未知回退以及预览/表情切换。静态源码匹配不能替代编译执行生产共用核心。
