package things

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var sceneCmd = &cobra.Command{
	Use:   "scene <subcommand>",
	Short: "场景联动管理",
	Long:  `场景联动相关操作：场景信息查询、创建、更新、删除、触发，以及场景日志。`,
	RunE:  wrapOldCommandNoCtx(shared.CobraBridge{}.RunScene),
}

func init() {
	ThingsCmd.AddCommand(sceneCmd)
}
