# 设备语音：控制完成策略候选

> 本文件是 [device-voice.md](device-voice.md) 的子文档。收录：控制完成策略候选的验收、应用与配置助手人设测试。

需要对照助手“查询不等于执行”的业务规则时，使用独立具名用例
`TestVoiceControlCompletionPolicyFrozenReplay`，不要替换原共享用例的提示词：

```bash
DEVICESIM_AUDIO_SAMPLE_RATE=16000 bash shell/remote-build.sh run --kind backend --scope backend --timeout 600 -- \
  bash -lc 'cd backend && go test ./things/tools/devicesim -run "^TestVoiceControlCompletionPolicyFrozenReplay$" -count=3 -timeout 480s'
bash shell/converge-watcher-vision.sh --check --control-policy
```

隔离用例与配置脚本共用`testdata/control_completion_policy.txt`，仍保留原音频、
预热查询、45秒实际下行及后续属性查询；它是待验收的配置候选，不默认应用。
若复跑出现预热查询失败，不能凭后续设置成功或早期少量通过采纳候选。
先核对失败预热的真实STT长度/摘要及查询语义，再关联上传音频、会话和ASR链路；
识别输入偏离查询样本时不能直接归因于模型工具策略，不记录对话正文或重跑到全绿。
回放要求真实STT保留固定查询的对象及查询关键字，否则立即停止后续记忆写入与控制对照；
`TestWarmVolumeQuerySemanticGuard`覆盖这个输入门禁，但不模拟ASR正确率。
ASR一次提交可能包含多帧前导缓存，不能把提交批次与UDP帧数差直接当作丢包；
需要关联实际序号、解码帧数与ASR前PCM，再区分网络、接收队列、VAD和供应商层。
显式应用使用`--apply --control-policy`及`WATCHER_CONTROL_POLICY_BACKUP_FILE`：
备份须为受限目录中未占用的绝对路径，仅追加自有SystemPrompt，拒绝冻结继承提示词；
写前回读与写后配置核验不能替代API原子比较交换。恢复经既有Agent更新API提交备份，
先确认没有后续他人编辑。规则单测为`TestDeviceControlCompletionPolicy`，
脚本回归为`bash shell/test-converge-watcher-vision.sh`；Mock通过不等于实际语音通过。

测试助手人设通过后，仍需检查实际助手自有人设与候选的组合。使用当前助手所属组织的
认证，显式设置`DEVICESIM_CONTROL_POLICY_SOURCE_AGENT_ID=<来源助手ID>`及匹配的
`DEVICESIM_LLM_CONFIG_ID=<实际模型配置ID>`，运行
`TestVoiceControlConfiguredAgentPromptPolicy`。它只读来源身份、模型与自有人设，保留
原文追加同一规则，在独占模拟设备执行查询→语音设置→真实下行→上报后查询，并清理
新建集合；不更新来源，不复制分组、历史或全部MCP绑定，不代签真实设备控制。
未指定来源时跳过，不能把该跳过计作通过；`TestConfiguredAgentPolicyPrompt`覆盖身份、
模型、继承人设拒绝、逐字保留及重复规则边界。规则在测试环境应用后仍须回读配置，
再通过真实设备当前会话、实际状态与后续上报验证，不以候选人设的小样本通过关闭原失败。
