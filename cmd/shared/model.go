package shared

import (
	"fmt"
	"io"
)

// runModel 物模型相关命令
//   model template [property|event|action|full] [--json|--yaml] [--output file]
//   model validate <file>
//   model generate-script <model-file> [--mode up-before|up-after|down-before|down-after] [--output file]
func runModel(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printModelHelp(stderr)
		return 2
	}

	subCmd := args[0]
	subArgs := args[1:]

	switch subCmd {
	case "template":
		return runModelTemplate(subArgs, stdout, stderr)
	case "validate":
		return runModelValidate(subArgs, stdout, stderr)
	case "generate-script":
		return runModelGenerateScript(subArgs, stdout, stderr)
	case "help", "--help", "-h":
		printModelHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "未知的 model 子命令: %s\n", subCmd)
		printModelHelp(stderr)
		return 2
	}
}

func printModelHelp(w io.Writer) {
	fmt.Fprintln(w, "物模型命令:")
	fmt.Fprintln(w, "  model template [property|event|action|full] [--json|--yaml] [--output file]  生成物模型填写模板")
	fmt.Fprintln(w, "  model validate <file>                                                         校验物模型 JSON 结构")
	fmt.Fprintln(w, "  model generate-script <model-file> [--mode type] [--output file]              根据物模型生成协议脚本框架")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "模板类型:")
	fmt.Fprintln(w, "  property  单属性模板（含 Define/Mode/Affordance 示例）")
	fmt.Fprintln(w, "  event     单事件模板（含 Params/Dir/Type）")
	fmt.Fprintln(w, "  action    单动作模板（含 Input/Output/Dir）")
	fmt.Fprintln(w, "  full      完整 Model 模板（Properties + Events + Actions）")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "脚本模式:")
	fmt.Fprintln(w, "  up-before    上行前处理（返回 *PublishMsg）")
	fmt.Fprintln(w, "  up-after     上行后处理（无返回值）")
	fmt.Fprintln(w, "  down-before  下行前处理（返回 *PublishMsg）")
	fmt.Fprintln(w, "  down-after   下行后处理（无返回值）")
}
