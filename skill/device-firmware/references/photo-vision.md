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

MCP查询、创建、更新和刷新工具须使用Agent所属企业的应用身份，不复用模型管理的
平台身份。按ID能看到平台私有服务，不代表企业会话按名称能加载它；遇到
`mcp service not found`先检查范围，再在授权企业幂等注册，不把旧私有服务改为公共。
列表为空可能是`data.list=null`，按空数组处理；注册/绑定后必须核验已发现
`deviceTakePhoto`并在新会话观察实际tool→action链路。对应维护脚本
`shell/converge-watcher-vision.sh`及其测试覆盖身份、空列表、工具缺失和重复创建边界。
本地双击的上传→inputSend通过不能证明语音MCP入口也通过。

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

## 6. 真机验收

1. 使用受控 OTA 从已确认旧版升级，核对镜像摘要、版本、`reportInfoReply` 和平台任务。
2. 说出明确拍照意图，验证 MCP、action 下行、预览、上传、actionReply、真实识图、字幕、
   语音和表情完整闭环。
3. 双击完成相同识图闭环，并回归单击、长按和播报中双击取消。
4. 连续拍照至少五次，记录每次耗时；无崩溃、持续堆下降或 MQTT 断线。
5. 自动回放全部 21 类表情，人工确认五类代表表情不会被状态切换过早清除。
6. 验证断网恢复、一次真实断电和不少于 15 分钟稳定运行。

真实按键、镜头画面、屏幕表情、扬声器和断电必须由现场观察；不可由协议回放代签。

支持串口测试命令的固件可执行 `run_ur_hw_e2e.py --port <串口> --expect <现场物品词>
--repeat 5 --log-dir <本次唯一日志目录>`；物品词只在设备内比较回复，不发给模型。
默认用例真实捕获/预览/上传，必须同时满足识物命中、文字、音频及完整播放。
链路齐全但 `text_match=0` 仍失败，应核对现场目标与画面，不能改成“有回复就通过”。

如用户明确只验硬件链路，可使用 `--flow-only --repeat 5`（不传 `--expect`）；
仍要求真实捕获、预览解码、上传、图片输入、文字完成、首帧、音频结束与播放排空，
结果明确标记 `recognitionChecked=false`。它不证明物品识别准确，也不能代替上述确定性
图片 E2E 或现场屏幕/扬声器确认。默认识物门禁保持不变。

同脚本的 `--reconnect` 在拍照前使用既有测试编译入口 `ur_hw_test wifi-reconnect`，
主动断开设备 Wi-Fi，约五秒后用原配置重连；不清 Wi-Fi、密钥或绑定，也不重启共享后端。
例如 `--reconnect --expect <现场物品词> --repeat 5 --log-dir <本次唯一日志目录>`。
runner 要求本次断开与订阅恢复，随后才拍照；还应独立核验恢复后的新属性回执及平台
身份/绑定/密钥摘要不变，旧在线状态或仅恢复订阅不能代替这些证据。
此入口仅在受控验收固件开放，结束后恢复正式包并核验云端确认；它不代替不同 AP 换网、
鉴权失败、真实断电或手机验收。已有日志目录可能被相机 runner 重写，每次必须新建唯一目录。

为兼容旧测试固件，仅在完整 `text_match=0` 链路之后，允许精确的
`FAIL reason=physical_photo_roundtrip` 及其唯一控制台 `0x1 (ERROR)` 返回；
其他原因、错误码、重复或先于终态的命令错误仍失败。采集终态必须等完整串口行，
防止分片截断错误原因；对应单测覆盖成功、缺阶段和上述拒绝分支。


同脚本 `--interrupt-only` 可省略 expect，**仍会真实拍照上传**，在实际首个音频帧后
经板级共享输入发送单击，要求 `respCancel → audioStarted`，再次单击退出监听。
此模式仅证明按键取消/恢复监听/退出，不证明识物或打断后的麦克风输入；需另做后续语音。
所有阶段出现崩溃、复位或命令错误都必须停止，不能因方法齐全或后续照片通过而忽略。
`test_run_ur_hw_e2e.py` 覆盖这类假通过及失败后不再触发后续拍照的门禁。

## 7. 分层排障

若已有文字及有声帧但没有`respAudioDone`，先核对服务端是否提前退出语音循环并取消TTS。
有效ASR Final与空Final是不同分支，两者都不能因ASR收尾期限截断已接收的活动回复。
用`TestFinalASRPreservesActiveTTS`验证有效Final分支，再用上述真实语音拍照E2E验证终态；
不要通过延长客户端等待或在异常取消后补发正常完成消息制造通过。

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

活动稳定性中若有完整文字与音频，但固定语音样本的 `stt_match=0`，仍判失败；
不要修改样本、关键词或重跑覆盖原窗口。关联该 session/轮次的 ASR 输入计数及既有 PCM，
在受限内存中对照原 Opus 解码，必要时枚举省略帧；仅保留长度、摘要和匹配结果。
精确匹配省略帧的 PCM 只能证明输入音频缺片段，不能定位丢失层。
接收计数可能在 ASR Final 时提前截取，不能单凭它断言 UDP 丢包；下一层需收发两端
仅含头部与序号的同轮证据。现场未采集的麦克风、声学和实屏结果继续单独标为未验收。

带发送取证的验收镜像在既有`CONFIG_UR_AI_E2E_TEST`下输出
`[UR_UDP_E2E] SEND seq=<序号> bytes=<数据报长度> sent=<套接字返回>`；
用`test_udp_send_evidence.py`核验实际发送函数、成功/失败/短返回及正式编译排除。
这些数值只说明设备提交结果，不说明服务端到达；不输出帧头、nonce、密钥或音频。
先固定源码、版本及镜像摘要，实际部署确认后才使用日志，不能把主机测试当作双端实测。
服务端包头观察应在新会话创建前就绪，只匹配该设备主题及当次 UDP 路由标识。
路由标识取 nonce 的零基偏移4–7字节（`nonce[4:8]`），在内存构造目标端口和 `udp[12:4]` 过滤；
不记录完整 nonce/密钥，也不宽泛采集共享端口。原始包仅在内存短暂解析，持久证据只留
会话短ID、入站时间、序号和长度；用有界观察器回收精确子进程，不保存 pcap/音频。
同时保存抓包器内核丢包计数；未取得该统计或存在丢包时，缺序号不能单独定位为网络丢包。
逐轮对照实际样本帧数、socket 返回与入站长度，同时比较序号集合和到达顺序，
因为乱序可能被服务端过期序号门禁丢弃。抓包晚启动、提前结束或解析不完整须标为
观察缺口，不能当作网络丢包；全部到达也只排除该轮入口之前的缺失，
仍需对照该轮解码、门控与 ASR 输入。新窗口通过不覆盖旧失败，测试结束恢复正式包。
即使关键词、文字与音频闭环全部通过，也可能缺失部分源帧；闭环与音频完整性分别记录。

主仓维护入口`firmware/watcher/scripts/ur_udp_evidence.py`只读取既有证据：
`--device-log <受限串口日志> --ingress-metadata <受限JSON> --session <八位短ID>
--expected-packets <实际总包数> --output <全新结果文件>`。
JSON外层为`sessions`数组；所选项包含`session`、`captureReady`、`exitCode`、
`incompleteBytes`、`packets:[{seq,bytes}]`和
`statistics:{captured,receivedByFilter,droppedByKernel}`。
用`parse_capture_statistics`读取真实英文tcpdump终态，缺失/重复不得补默认零；
读完stderr后再Wait，防止子进程管道被提前关闭而截断统计。
`receivedByFilter`不一定等于捕获数；要求捕获数等于解析数且内核丢包为零。
退出0为本轮完整，1为确定的发送/送达不一致，2为证据不足；首包未观测、迟启动、
统计缺失或解析不完整不能宣称网络丢包。`test_ur_udp_evidence.py`覆盖三态、
乱序/重复/四帧缺口与权限、旧文件/符号链接防覆盖，结果只保留数值，不覆盖旧失败。
隔离模拟双端校准通过只证明观测方法，不替代Watcher实际无线链路或旧失败复现。
窗口中止时保留原计划失败，不从SEND数量自动放宽预期；另由完整INPUT_DONE帧数
与已知预热数建立已完成输入前缀，可作额外对照，但前缀结果不能代替整组验收。
精确路由过滤的零丢包统计仍不能区分无线丢失与源端异常路由头；需更近发送侧证据，
不将抓包器零丢包推导为UDP提供可靠送达，或直接归咎ASR/LLM。

排除路由过滤盲点时，使用主仓`ur_udp_capture.py`在内存由匹配本轮路由的包学习
目标源IPv4/端口，再对该精确端点独立捕获；禁止放宽到共享端口。端点视角保留
错路由/非法头，仅输出`routeMatch/headerValid`等数值或布尔元数据，不落盘地址、
路由或原始包。完整性门禁拒绝假值，字段不完整只判证据不足；两种捕获各自记录
真实统计、就绪和退出，晚启动的端点视角不能补造早期包或拼接为完整窗口。
本机自有UDP校准可证明两种过滤不同，但不替代Watcher无线链路及真实语音E2E。

需要补充发送侧证据时，可构建显式诊断包：同时设置
`ENABLE_UR_AI_E2E_TEST=1`和`ENABLE_UR_UDP_TX_TRACE=1`，后者默认0。
Watcher构建入口只允许已核验的固定IDF6镜像；构建前后运行
`check_wifi_tx_trace_owner.py`，应用/组件/SDK网络源码有其他setter引用或必需目录
缺失就拒绝。以原始字节匹配ASCII标识符，兼容厂商非UTF8注释，但不忽略文件。
私有SDK回调仅把当前路由的有界数值事件入队，worker输出`[UR_WIFI_TX]`的
BEGIN/DONE/END及注册失败状态；不保存地址、nonce、帧或音频。
先校准真实硬件的回调布局及匹配计数；布局不支持、零匹配、队列溢出或缺终态时
只能判证据不足，不能扫描载荷猜偏移。`status=1`只是驱动报告，不是空口ACK或
云端送达证明。`test_ur_udp_tx_trace.py`覆盖实际生产核心、SDK边界及构建保护，
其中13项C++用例各十次；替身测试和编译通过仍不替代硬件校准。
此观察不改变业务发送、序号、样本及验收门禁；测试后恢复关闭两个开关的正式包。



若上传完成后在 HTTP 断线回调中看门狗重启，先用当前已部署镜像的原 ELF 解码，
不能用后来重建的 ELF 猜调用栈。检查被动断连是否仅因 `connected=false` 跳过接收线程
退出等待，以及 HTTP 成员锁是否先于传输对象销毁。连接状态不等于回调已结束；
用可控制放行时机的线程用例复现“迟到回调访问已销毁资源”，再验证析构顺序与 join。
受管理依赖采用有锚点校验的构建副本，不手改 vendor；单测通过后仍需连续真机复测。
