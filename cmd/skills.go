package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var skillsCmd = &cobra.Command{
	Use:   "skills <subcommand>",
	Short: "Skill 管理",
	Long:  `管理 CLI Skill 文档：列出、更新、部署到 AI 工具、查看版本。`,
	// 关闭 Cobra flag 解析：--dry-run/--json/--dir 等由 cmd/shared 手动解析
	DisableFlagParsing: true,
	RunE:               wrapOldCommandNoCtx(shared.CobraBridge{}.RunSkills),
}

func init() {
	RootCmd.AddCommand(skillsCmd)
}
