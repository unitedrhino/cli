package things

import (
	"gitee.com/unitedrhino/cli/cmd/shared"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:                "project <subcommand>",
	Short:              "项目管理",
	Long:               `项目相关操作：项目信息、项目配置查询与管理。`,
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunProject),
}

func init() {
	ThingsCmd.AddCommand(projectCmd)
}
