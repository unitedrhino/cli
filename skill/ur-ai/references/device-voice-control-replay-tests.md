# 设备语音：控制重放与隔离对照

> 本文件是 [device-voice.md](device-voice.md) 的子文档。收录：语音控制失败排查中的重放与隔离对照用例。

进一步区分ASR措辞和输入通道时，可运行`TestVoiceVolumeRecognizedTextReplay`：
沿用共享模拟设备、原音频和提示词，取得真实STT后逐字经文字入口发送，并核验控制下行
和后续属性查询。使用上面的资源入口，把`-run`改为`^TestVoiceVolumeRecognizedTextReplay$`、
`-count=1`，包路径改为`./things/tools/devicesim`，测试预算改为240秒。
语音轮已经控制成功时该诊断明确跳过，跳过不计通过。测试会正常追加会话历史和模拟属性，
不清理共享历史；文字重放通过后仍须原样复测`TestVoiceVolumeWorkflow`，不能据此关闭
语音缺陷或认定某条历史记忆是根因。结果分开记录，未经关联的迟到控制不能算重放成功。

上述重放已经在共享Clone追加语音轮，不能单独归因于通道。更严格的对照使用
`TestVoiceVolumeTextControlAfterQuery`：仅在独占的新资源上取得真实STT，共享设备只做
查询→逐字文字控制→查询，并回收识别来源。再对照`TestVoiceControlSessionTransition`
的全新资源场景；若共享文字也失败而隔离通过，应继续检查有效上下文与模型选择，
不要归因于固件采音，也不能清共享历史、修改提示词或放宽控制门槛来获得通过。

若实际提示的稳定画像不同，可运行`TestVoiceControlStableProfileReplay`（同一资源入口，
`-count=1`、devicesim包、420秒预算）：只读固定共享模拟资源的画像，内存中经现有
记忆API写入本次独占分身，空画像与重放画像分别运行原控制流程并自动回收。
不读取真实用户画像、不输出正文、不改共享历史。标准API也创建来源记录，不能称为
仅替换系统提示词的纯变量实验；隔离通过只说明在这组新资源条件下未复现，仍须检查
共享动态召回/历史组合并原样复测失败用例，不能凭画像存在就修改生产过滤规则。
同一入口可改跑`TestVoiceControlRecallContextReplay`，在独占分身重放来源有效Dream摘要，
核对后台实际召回及控制下行。它不复制来源ID、历史时间、权重或访问计数，标准API也会
合并目标画像，所以只是组合对照，不是等价生产快照。记忆轨迹ID可能包含查询正文，
脱敏须取哈希，不能直接输出或只截短前缀。
需要覆盖冻结画像与新动态召回时，成对运行
`TestVoiceControlFrozenRecallContextReplay|TestVoiceControlFrozenProfileReplay`：
先用完整语音查询预热，再分别写入摘要或不写；核对实际控制轮画像块哈希，而非仅回读
持久化画像。预热增加短时原文，摘要API也改变中间文字查询画像，所以仍不属于纯变量
实验。摘要记录可能正文重复，实际格式化去重后的条数才是注入条数；红色复现必须保留，
不得用刷新缓存、清历史或隔离对照通过替代原共享失败的修复。
