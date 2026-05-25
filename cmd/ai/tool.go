package ai

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var toolCmd = &cobra.Command{
	Use:   "tool <subcommand>",
	Short: "AI 工具管理",
	Long: `AI 工具相关操作：制品获取与保存、编辑、验证、渲染、运行。

对应旧命令: ur ai-tool`,
	RunE: wrapOldCommand(shared.CobraBridge{}.RunAiTool),
}

func init() {
	AICmd.AddCommand(toolCmd)
}
