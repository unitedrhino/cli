# CLI 命令 — 属性控制与行为：生成 executor.js

本文件是 device-control.md 的子文档，收录生成 executor.js 的规则（规则引擎 AI 任务开发）。

- 脚本每次只执行一次业务操作，频率/次数交给任务调度，不写无限循环或内部定时器。
- 对于物模型允许的 20～30、步长 0.5，离散随机公式为 `20 + Math.floor(Math.random() * 21) * 0.5`；检查上下界，其他范围按真实物模型计算。
- 优先用上述 `ur things device` 领域命令；只有领域命令确实不覆盖需求时才使用 `ur api`。用 `Bun.spawn` 的参数数组调用 ur，继承环境，不拼接 shell 命令。并行读取 stdout/stderr，并以 `const exitCode = await proc.exited` 取得最终退出码；禁止读取可能仍为 `null` 的 `proc.exitCode`。
- 需要解析完整 `{code,data,msg}` 响应时，每条领域命令必须且只能传一个 `-j`（等价于 `--json`）；未传时可能只返回精简后的 `data`，不能混用两种输出合同。
- 项目输入若提供，必须是非空、无首尾空白且与执行上下文一致的字符串，禁止 Number/String 强转掩盖错误；创建者凭证缺失时明确失败。
- 失败必须抛错或 `console.error(...)` 后 `process.exit(1)`；禁止 catch 后在 stdout 输出 `{code:500}` 再正常退出，因为进程 0 会造成假成功。成功才在 stdout 最后一行输出结果 JSON。生成并实际保存 executor.js、manifest.json、skill.md 后才告知完成，不把聊天中的代码块当作已保存产物。终端临时文件不会自动同步到平台工具工作区；保存后必须重新读取平台源文件并真实运行核对。
