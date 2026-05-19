package shared

import (
	"context"
	"fmt"
	"io"
)

func runAiTool(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAiToolHelp(stderr)
		return 2
	}

	subCmd := args[0]
	subArgs := args[1:]

	switch subCmd {
	case "artifact":
		return runAiToolArtifact(ctx, subArgs, stdout, stderr)
	case "validate":
		return runAiToolValidate(ctx, subArgs, stdout, stderr)
	case "run":
		return runAiToolRun(ctx, subArgs, stdout, stderr)
	case "edit":
		return runAiToolEdit(ctx, subArgs, stdout, stderr)
	case "render":
		return runAiToolRender(ctx, subArgs, stdout, stderr)
	default:
		fmt.Fprintf(stderr, "未知的 ai-tool 子命令: %s\n", subCmd)
		printAiToolHelp(stderr)
		return 2
	}
}

func printAiToolHelp(w io.Writer) {
	fmt.Fprintln(w, `ai-tool — AI 工具三件套管理

用法:
  ur ai-tool artifact get --id <id> [--output-dir <dir>]
  ur ai-tool artifact save --id <id> --dir <dir>
  ur ai-tool validate --id <id>
  ur ai-tool run --id <id> --inputs <json> [--timeout <seconds>]
  ur ai-tool edit --id <id> --instruction <text>
  ur ai-tool render --id <id> [--output <file>]

子命令:
  artifact get   从平台获取三件套（executor.js / document.md / manifest.json）
  artifact save  将本地三件套保存到平台
  validate       校验三件套一致性（JSON 格式、JS 安全、变量一致性、组件白名单）
  run            在沙箱中运行工具，轮询等待结果
  edit           AI 编辑三件套
  render         将 document.md 渲染为 HTML（含组件占位）`)
}
