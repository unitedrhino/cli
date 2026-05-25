package things

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var projectCmd = &cobra.Command{
	Use:   "project <subcommand>",
	Short: "项目管理",
	Long:  `项目相关操作：项目信息、项目配置查询与管理。`,
	RunE:  wrapOldCommand(shared.CobraBridge{}.RunProject),
}

func init() {
	ThingsCmd.AddCommand(projectCmd)
}
