# 语音拍照与控制的真机自动复测

> 本文件是 [voice-ai.md](voice-ai.md) 的子文档，收录原指南第 6 节的两个专项真机复测用例：语音拍照与语音控制（含资源、串口命令与证据要求）。

## 语音拍照的真机自动复测

音量查询样本不能证明拍照工具可用。新测试固件支持额外的真实语音样本：

```bash
python3 firmware/watcher/scripts/run_ur_ai_e2e.py \
  --port <serial-port> --fixture take_photo --repeat 5 --timeout 150 \
  --log-dir <repo>/.temp/device-firmware/voice-photo-e2e
```

对应命令为 `ur_ai_e2e <1-5> photo`；缺省仍查询音量。拍照资源使用
devicesim 的 `turn16_take_photo.mp3` 经同一 opusfixture 编码器生成，输出
`firmware/watcher/main/testdata/ur_ai_take_photo.opuspack`，16kHz/60ms/112帧。
只有测试开关启用时嵌入，正式固件不带资源/命令。修改生成器或样本后重新生成，
运行 `test_ur_ai_fixture` 编译实际固件解析器，验证三份真实包及损坏包拒绝。

除完整语音阶段、播放排空和原延迟门禁外，STT必须命中“拍照”，每轮必须按顺序
具备真实捕获、上传开始、校验成功与 `action_reply code=200 published=1`。
不能用模型口头确认、失败后一句回复、错样本或上一轮动作记录代替；runner单测覆盖
这些失败，串口参数与STT门禁测试编译实际生产函数。MQTT发布成功仅证明发送路径接受，
不是平台处理确认；完整模型回复仍须采集，严格语义另用确定性图片平台E2E核对。
该测试替代麦克风输入但使用真实摄像头；仍需现场验收麦克风、画面实物、物理显示和听感。

## 语音控制的真机自动复测

支持控制资源的新测试固件使用以下入口；正式固件不开启测试开关或嵌入样本。

```bash
python3 firmware/watcher/scripts/run_ur_ai_e2e.py \
  --port <serial-port> --fixture control_volume --turns 1 --repeat 1 \
  --log-dir <repo>/.temp/device-firmware/voice-control-e2e
```

对应串口`ur_ai_e2e 1 control`。原`turn5_set_volume_prefix_40.mp3`通过生产
opusfixture编码器生成`ur_ai_control_volume.opuspack`，16kHz/60ms/40帧；不注入
STT文字。每次必须先通过现有平台属性控制准备音量55，并核验实际上报。入口不偷偷
设置初值，初值已是40或不符55时拒绝；只有完整真实回复、播放排空后实际codec音量
变为40才通过。runner拒绝旧标记、无变化和错误初值；只允许单次单轮，重复执行须
重新准备55。75秒总预算及8秒/15秒延迟门禁不放宽。

设备值变化不是全部平台证据。须同时关联本轮MCP调用、属性下行、原token的
`controlReply code=200`与后续属性上报；只看到文字确认、设备原本就是目标值或
主机替身通过都不能判闭环成功。串口测试替代现场采音，不替代麦克风和扬声器验收。

MCP证据取自服务端实际调用日志，并以本轮会话及有时区的时间窗关联。
当前设备协议不下发工具调用生命周期，不要等待不存在的`toolCallStart/Result`
来判失败或新增接口；原取证失败应保留，用同轮只读日志补证，不重跑覆盖。

总回复预算不能等同于首帧延迟：拍照测试等待完整回复120秒，runner缺省150秒，
音量缺省75秒；显式 `--timeout` 可覆盖。长回复可能合成为一分钟音频，按60ms帧
实时下行时，固定60秒总等待会把生成耗时也算进去并提前关闭会话。先对齐TTS实际
音频时长、下行时序及设备关闭时间再归因；8秒首帧/15秒文字门禁仍严格保留。
这些预算仅影响测试入口，不改变生产活动超时；任何缺音频终态依然判失败。

关联串口文件名与服务端日志前先核对时区及时间格式。
`utils.DebugLog` 的 `ts` 只有当地时分秒，不含日期/偏移；不能直接与 runner 的
UTC文件名前缀比较。以环境文档、进程时区和同轮方法/短ID收敛窗口，空查询不证明
服务端没有事件；诊断摘要只输出方法、错误分类、帧数和耗时，不输出完整对话或图片地址。

