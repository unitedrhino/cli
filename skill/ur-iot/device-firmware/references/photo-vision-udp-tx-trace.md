# 7. 发送侧诊断构建与链路探针（分层排障）

本文件是 [photo-vision.md](photo-vision.md) 的子文档，收录其第 7 节「分层排障」中发送侧诊断构建部分：ENABLE_UR_UDP_TX_TRACE、UR_WIFI_TX 追踪、UR_NETIF_TX 交接包装与检查脚本。

需要补充发送侧证据时，可构建显式诊断包：同时设置
`ENABLE_UR_AI_E2E_TEST=1`和`ENABLE_UR_UDP_TX_TRACE=1`，后者默认0。
Watcher构建入口只允许已核验的固定IDF6镜像；构建前后运行
`check_wifi_tx_trace_owner.py`，应用/组件/SDK网络源码有其他setter引用、netif包装实现/
链接接管或必需目录
缺失就拒绝。以原始字节匹配ASCII标识符，兼容厂商非UTF8注释，但不忽略文件。
私有SDK回调仅把当前路由的有界数值事件入队，worker输出`[UR_WIFI_TX]`的
BEGIN/DONE/END及注册失败状态；不保存地址、nonce、帧或音频。
SDK私有头没有承诺Ethernet帧封装；回调触发次数不能代替当前路由的目标匹配。
零匹配时先看END的有界拒绝计数：空/短/超限帧、EtherType和IPv4结构拒绝、
最短/最长声明长度及开头IPv4/802.11数据类型候选；仅用于区分解析门禁，
不输出借用缓冲、扫描偏移或据候选启用新解析器。计数包含非音频流量，
不能把ignored或某个封装候选数视为目标丢包；每次BEGIN重置，仍需实际目标匹配。
进一步查看LAYOUT：标准帧控制位推导唯一MAC头尾，分别计24/26/30/32/36字节头、
QoS/四地址/Protected及完整性，再只在该头尾核验固定RFC1042/IPv4 SNAP与IP/UDP
声明边界。Protected位不能证明回调缓存载荷已加密或解密，不据此跳过安全头；
Null、分片和A-MSDU只计数、不展开。完整IP/UDP候选仍包含非语音流量，不读取
路由或创建目标事件。可对照[Linux主线头长定义](https://github.com/torvalds/linux/blob/master/net/wireless/util.c)，
独立实现标准位规则，不移植其许可受限代码；合成结构向量通过不是实际SDK封装证明。
先校准真实硬件的回调布局及匹配计数；布局不支持、零匹配、队列溢出或缺终态时
只能判证据不足，不能扫描载荷猜偏移。`status=1`只是驱动报告，不是空口ACK或
云端送达证明。`test_ur_udp_tx_trace.py`覆盖实际生产核心、SDK边界及构建保护，
其中21项C++用例各十次，包含超限先拒绝、标准头长/固定SNAP、逐字节截断、
错误偏移/安全头/分片/聚合拒绝及网络长度门禁，候选计数不产生目标事件；
替身测试和编译通过仍不替代硬件校准。
此观察不改变业务发送、序号、样本及验收门禁；测试后恢复关闭两个开关的正式包。

同一诊断开关另在网络栈向STA驱动交接处包装`esp_netif_transmit`及
`esp_netif_transmit_wrap`，不替换驱动、不修改SDK源码；可核对固定SDK的
[交接合同](https://github.com/espressif/esp-idf/blob/906c6ee6884d77e9ca6910ff0f8c8debd657dfdc/components/esp_netif/include/esp_netif_net_stack.h)
及对应`wlanif.c`调用路径。交接前只复制目标路由的代际/序号/长度，原函数只调用一次、
参数和返回值原样透传；交接返回后不重读可能已被消费的帧或pbuf。
worker输出独立`[UR_NETIF_TX] BEGIN/DONE/END`：DONE保留真实`result`返回码，
END保留calls/matched/dropped/ignored/pending；calls包含非音频流量，pending表示
匹配后尚未返回的交接。关闭/跨代际结果隔离，固定32项队列溢出明确计数。
`result=0`只表示交接API返回ESP_OK，不证明空口ACK、驱动完成或服务端收到，
不得用它替代`UR_WIFI_TX`或服务端入站证据。先在真机同轮核验实际匹配、零队列溢出、
零pending及完整序号/长度，不能由包装符号存在推导交接覆盖或原丢包已修复。
`test_ur_udp_netif_trace.py`执行生产包装器，八类C++各十次，覆盖两路径、原缓冲消费、
错误码、关闭/新代际、截断/范围、溢出和并发；另测占用与ELF门禁。
Watcher构建入口在诊断和正式两种路径都执行`check_udp_tx_trace_symbols.py`：
诊断必须链接实际STA交接入口，正式必须排除项目诊断入口/包装器；缺工具、无可识别
符号或仅未定义引用不能通过。未调用的普通transmit入口被链接回收不等于STA路径缺失。
