package things

import (
	"gitee.com/unitedrhino/cli/cmd/shared"
	"github.com/spf13/cobra"
)

var sceneCmd = &cobra.Command{
	Use:                "scene <subcommand>",
	Short:              "场景联动管理",
	Long:               `场景联动相关操作：场景信息查询、创建、更新、删除、触发，以及场景日志。`,
	DisableFlagParsing: true,
	RunE:               wrapOldCommandNoCtx(shared.CobraBridge{}.RunScene),
}

func init() {
	ThingsCmd.AddCommand(sceneCmd)
}
