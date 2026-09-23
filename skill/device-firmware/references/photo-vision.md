# 联犀设备拍照识图与表情闭环

本指南用于已接入联犀 MQTT、物模型和 AI 会话的带相机设备。目标是让语音工具调用和
设备本地按键都能进入同一图片理解链路，并以文字、语音和表情完成回复。语音会话基础
协议先读 [设备语音 AI](voice-ai.md)，平台模型与 MCP 配置同时读
`ur-ai/references/device-voice.md`。

## 完成标准

只有三层证据都通过才算完成：

1. devicesim 使用仓库自有确定性图片通过真实上传、视觉模型、MCP 和表情 E2E。
2. 固件共用生产逻辑的单元测试与协议回放覆盖上传、行为、图片输入、去重和生命周期。
3. 真机相机、预览、双击、屏幕、扬声器、断网、断电及稳定性验收通过。

模拟测试不能证明摄像头、按键和屏幕有效；单次 API `code=200` 也不能证明设备行为闭环。

## 1. 平台与物模型

产品至少声明一个下行行为，identifier、方向和输出必须与设备实现完全一致：

```json
{
  "identifier": "takePhoto",
  "dir": "down",
  "input": [],
  "output": [
    {
      "identifier": "fileUri",
      "name": "图片访问地址",
      "define": {"type": "string", "max": "2048"}
    }
  ]
}
```

视觉 Agent 应同时声明 `text`、`voice`、`image_input`，文字和语音仍可流式输出；图片输入
若当前模型的流式视觉编码不稳定，使用同步生成。LLM 与 vision 配置都必须指向真实支持
图片输入的同一配置，不能用纯文本模型返回猜测结果冒充识图成功。

设备控制 MCP 与拍照 MCP 分开授权：保留经过审核的 IoT MCP，并通过平台 API 幂等查询或
注册内置 fuzai MCP。不要硬编码环境中的历史 MCP ID。Agent 只绑定明确需要的服务，变更后
回读 Agent、模型、capabilities 和 MCP ID；不直接写数据库。

`sessionCreated.supportedModalities` 必须包含 `image`，并提供当前 session 的短期
`uploadUrl`。上传地址和认证信息属于会话凭据，不写入日志或持久化。

## 2. 两条拍照入口

### 语音工具入口

职责边界：fuzai 调用设备行为、等待上传结果并校验下载文件，返回标准 MCP
`image` 块（`data`、`mimeType`），`fileUri` 作为独立文本块的结果元数据。
禁止在文本 JSON 中夹带内联图片。core 不识别拍照工具名、私有字段或拍照提示词，
只传递通用图文内容，复用现有多模态转换、模型能力校验与视觉路由；
图片只用于当前请求，不写入历史或工具事件，也不另起对话轮次。
core 与 fuzai 配套升级/回滚，设备行为回执和本地 `inputSend(image_url)` 不变。

传输层也要回归：HTTP 响应必须完整且有界读取，超限、非 JSON 和失败状态只返回脱敏错误，
不能把截断后的正文作为普通工具文本保存。当前单结果图片合计上限 5 MiB，MCP 报文上限
8 MiB（包含 base64 膨胀与元数据）；用 `chatadapters/mcphttp` 的真实本地 HTTP 测试
覆盖大报文、SSE、截断、超限及失败正文不泄漏。

固定闭环为：

```text
用户说拍照意图 → 模型调用 deviceTakePhoto → 平台下发 takePhoto
→ 设备捕获并上传 → actionReply.fileUri → 模型识图
→ respText / respAudio / respEmotion
```

设备成功回执使用原始 `msgToken` 和 `actionID`：

```json
{
  "method": "actionReply",
  "actionID": "takePhoto",
  "msgToken": "<request-token>",
  "code": 200,
  "msg": "success",
  "data": {"fileUri": "<platform-file-uri>"}
}
```

相机繁忙、OTA、断网、捕获失败或上传失败必须返回稳定的非 200 结果；不得伪造 URL、清除
DeviceSecret 或阻塞 MQTT Yield。相同 token 只执行一次，一次只允许一个拍照任务。

### 本地按键入口

双击属于独立且互斥的按键事件，不能由两个单击拼成。双击时唤醒屏幕、建立或复用 AI
session、预览并上传，然后发送：

```json
{
  "method": "inputSend",
  "params": {
    "contents": [
      {"type": "text", "text": "请描述刚拍摄的画面"},
      {"type": "image_url", "imageUrl": "<fileUri>"}
    ],
    "modalities": ["text", "audio"]
  }
}
```

播报中双击先发送 `respCancel`。拍照繁忙或 OTA 中只提示用户，不破坏当前 session；单击
对话和长按动作必须做回归测试。

## 3. 捕获、上传与凭据边界

- 只在用户明确意图、平台 `takePhoto` 或本地双击时启动相机；禁止后台抓拍。
- 捕获 JPEG 后校验 SOI/EOI、非空和设备缓冲上限，预览只持有当前帧，不落盘。
- multipart 上传目标为平台 base URL 与 session `uploadUrl` 安全拼接后的同源地址。
- Basic Auth 使用 SDK 提供的短期 HTTP 认证副本。只允许受控复制，不暴露 DeviceSecret；
  上传成功、失败或 MQTT 断线后都立即清零副本。
- 只有 HTTP 200、平台 `code=200` 且 `data.fileUri` 非空时才成功；不记录认证头、图片内容
  或完整签名 URL。
- MQTT 回调只复制 token/actionID 并入队；相机 worker 捕获和上传，MQTT 任务统一发布回复。
- 动作总耗时必须小于服务端工具等待窗口，并为 DNS、HTTP 和捕获分别设置有界超时。

## 4. 表情显示生命周期

协议白名单和设备字体资源必须在构建时逐项校验，仍以 devicesim 的 21 类为基准。未知、
空值、旧 session 或错误 `respId` 回退或忽略为 `neutral`，不能把模型任意文本映射为资源名。

`respEmotion` 到达后：唤醒屏幕、清除相机预览、显示表情，并保持到当前轮真实音频播放队列
排空；随后恢复监听再回到 `neutral`。状态切换不能提前覆盖表情，晚到表情应保留一个可见
的最短时长。自动测试验证映射和状态机，人工真机至少确认 `happy`、`sad`、`angry`、
`thinking`、`neutral` 五类可见。

## 5. 可重复自动测试

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

```bash
DEVICESIM_TEST_PATTERN='Test(TakePhotoEndToEnd|ImageInputEndToEnd|EmojiEmotionText|EmojiNotSentForPlainQuestion)$' \
DEVICESIM_AUDIO_SAMPLE_RATE=16000 \
bash shell/devicesim-oneclick-test.sh
```

测试图片必须是仓库自有、内容固定的 JPEG，并断言模型实际识别其文字、颜色和形状；仅断言
“收到一段回复”不够。`TestTakePhotoEndToEnd` 验证工具调用、action、可下载 `fileUri` 和
识图回复，`TestImageInputEndToEnd` 验证双击等价的 `inputSend(image_url)` 路径。

固件侧运行完整测试发现入口；关键纯逻辑用例至少连续 10 次：

```bash
cd firmware/watcher
python3 -m unittest discover -s scripts/tests -v
```

固件测试至少覆盖：URL 解析、同源约束、JPEG 校验、认证副本清零、平台响应、行为回执、
token 去重、相机互斥、队列拥塞、超时、断线、OTA 拒绝、图片信封、按键互斥、21 类表情
资源一致性、未知回退以及预览/表情切换。静态源码匹配不能替代编译执行生产共用核心。

## 6. 真机验收

### 无需现场反复按键的串口回归

如已获准刷测试固件，启用 `ENABLE_UR_AI_E2E_TEST=1`，使用固件仓库
`scripts/run_ur_hw_e2e.py --port <串口> --expect <现场物品词> --repeat 5 --log-dir <受限目录>`。
脚本走真实相机捕获、上传、图片输入、模型识物与语音播放；预期物品词仅在设备内比较，
不发送给模型。追加 `--reconnect` 测 WiFi 断线后的 AI 恢复，追加 `--interrupt`
测真实播报中按键打断及后续拍照。两个选项可分别运行以隔离失败。

串口命令 `ur_hw_test photo <物品词>` 是完整用例；`photo-start`、`press`、
`state`、`wifi-reconnect` 是控制入口，ACK 不能当成 PASS。运行前确认该版本含这些
测试入口；结束后恢复 `ENABLE_UR_AI_E2E_TEST=0` 的正式固件并检查联网。
一个串口只保留一个读写进程，优先使用 runner 原始文件描述符；默认 DTR/RTS
翻转可能重启设备，不能将开串口造成的重启或此时的命令超时误判为业务故障。

MQTT 重连测试必须检查 AI 就绪，不仅检查属性上报：SDK 可能已发出重连事件，
但在同一次 Yield 内把返回码覆盖为成功。固件应保留待恢复事件，在任务中完成
重订阅后再通知 AI；用生产回调/恢复分支的编译执行测试覆盖该路径。
模拟事件不能证明物理按键触点、屏幕可见效果或真实听感；这些仍需现场确认。

1. 使用受控 OTA 从已确认旧版升级，核对镜像摘要、版本、`reportInfoReply` 和平台任务。
2. 说出明确拍照意图，验证 MCP、action 下行、预览、上传、actionReply、真实识图、字幕、
   语音和表情完整闭环。
3. 双击完成相同识图闭环，并回归单击、长按和播报中双击取消。
4. 连续拍照至少五次，记录每次耗时；无崩溃、持续堆下降或 MQTT 断线。
5. 自动回放全部 21 类表情，人工确认五类代表表情不会被状态切换过早清除。
6. 验证断网恢复、一次真实断电和不少于 15 分钟稳定运行。

真实按键、镜头画面、屏幕表情、扬声器和断电必须由现场观察；不可由协议回放代签。

## 7. 分层排障

| 现象 | 先查 | 常见根因 |
|---|---|---|
| devicesim 图片 E2E 失败 | 模型配置测试、Agent/MCP 回读 | 模型不支持 image、视觉流式编码、fuzai/Redis/DmRpc 或 MCP 绑定 |
| 未收到 `takePhoto` | toolCall 和物模型 | MCP 未授权、行为 identifier/dir 错误、设备上下文缺失 |
| 收到行为但无图片 | worker 与相机日志 | MQTT 回调做重活、相机互斥、捕获失败、OTA 拒绝 |
| 上传失败 | uploadUrl、HTTP 状态和脱敏错误码 | session 过期、URL 非同源、认证副本失效、JPEG 或大小不合法 |
| action 成功但模型看不到图 | actionReply 与文件下载 | token/actionID 错误、`data.fileUri` 结构错误、文件不可下载 |
| 双击无反应或触发单击 | 按键事件序列 | 未注册双击、单击/双击未互斥、任务仍繁忙 |
| 有文本无语音 | `modalities` 与 TTS/UDP | 图片输入漏 `audio`、TTS 失败或音频播放链路故障 |
| 表情一闪即逝 | respId、播放队列和 UI 状态 | 状态切换提前恢复 neutral、晚到表情无最短可见期 |
| 连续拍照后断线或崩溃 | 堆、帧所有权和凭据释放 | JPEG/JSON 未释放、跨任务悬空指针、认证副本未清零 |

排障顺序保持：平台模型配置测试 → devicesim 两条视觉 E2E → 固件单元/回放 → 真机相机、
网络、屏幕与按键。
