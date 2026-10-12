# 设备语音：模型边界与工具关联测试

> 本文件是 [device-voice.md](device-voice.md) 的子文档。收录：模型边界合同测试、工具绑定隔离与工具响应关联/往返核验。

排查“识别到控制要求但模型只查询”时，增加模型边界合同测试，不直接清历史或修改提示词：

```bash
bash shell/remote-build.sh run --kind backend --scope backend --timeout 300 -- \
  bash -lc 'cd backend && go test -race ./core/service/aisvr/internal/domain/llm -run "^(TestTracedTransport.*|TestModelRequest.*|TestMediaGuard.*|TestVisionRouter.*)$" -count=10 -timeout 120s'
```

该测试使用真实模型工厂、门禁、路由和OpenAI适配器，将固定夹具发送到本地HTTP模拟
供应商，核验纯文本、查询历史、召回上下文和工具结果之后的消息顺序与工具定义。
供应商响应是模拟值，不能证明真实模型会选择控制工具，也不能替代失败的原始devicesim
用例。Runtime模型调用前的摘要仅能证明该边界的输入；不带会话/trace关联的摘要不能
单独用于跨并发定位。继续核对实际模型配置、MCP清单、工具选择和真实控制回执。
模型HTTP时序诊断不得预读、关闭或以旧GetBody覆盖当前请求体，不落盘请求JSON；
正文可能含对话、工具参数及内联图片。旧诊断文件需另行核实留存和删除授权，不清空共享目录。
旧provider测试若依赖硬编码外部凭据，不借其发起真实调用，也不把所选离线集合称为整包全绿。

`TestModelRequestToolBindingIsolation`从同一工厂衍生无工具、只读工具和写入工具模型，
经生产包装器并发发送重复请求，核验实际HTTP工具隔离、完整固定请求一致及模型/流式/
token上限/思考参数。供应商仍为本地固定SSE，不证明真实LLM正确选择工具。
时序观测也要覆盖竞态：连接复用可能让后台拨号回调晚于RoundTrip返回，回调写入与
日志读取需要同一同步边界；返回时复制快照后再写日志，不持锁等待网络或读取正文。
诊断竞态修复与模型选错工具是不同问题，前者转绿不能关闭原真实控制E2E失败。

工具响应的关联还需通过真实SDK和Runtime累积路径核验，而非只看出站请求：

```bash
bash shell/remote-build.sh run --kind backend --scope backend --timeout 300 -- \
  bash -lc 'cd backend && go test -race ./core/service/aisvr/internal/domain/agentruntime -run "^TestModelResponseToolCallCorrelation$" -count=10 -timeout 120s'
```

该用例以本地固定SSE覆盖共享模型并发查询/写入、同一响应两个工具的交错参数分片，
实际执行模型工厂、门禁、路由、SDK和`streamAndAccumulate`，逐项检查调用ID、工具名、
完整参数及文本增量。它不调用真实LLM，也不模拟设备回执；通过只证明这些固定响应
正确关联，不能据此宣布供应商工具选择正确或原语音控制E2E转绿。

还须覆盖工具执行后的下一步请求：同一资源入口把`-run`改为
`^TestModelResponseRuntimeToolRoundTrip$`，竞态重复十次。该用例执行真实工厂、SDK、
完整`Run`循环及MCP适配器，只在外部供应商/MCP边界使用固定合成读写响应，
核验读后写的实际调用、原参数、自动会话注入、下一步HTTP的tool结果/调用ID和单次终态，
同时覆盖同步返回与事件流；不能只调用累积函数就称完整Runtime已执行。
生产Run可能自动添加系统提示，夹具必须保留并核验，不能关掉装配分支凑预期消息数量。
这仍不是模型意图识别或设备回执验收；原真实控制用例失败时仍记未解决。
