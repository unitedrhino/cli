package things

import (
	"gitee.com/unitedrhino/cli/cmd/shared"
	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use:                "script <subcommand>",
	Short:              "协议脚本管理",
	Long:               `协议脚本相关操作：验证、模板生成。`,
	DisableFlagParsing: true,
	RunE:               wrapOldCommandNoCtx(shared.CobraBridge{}.RunScript),
}

func init() {
	ThingsCmd.AddCommand(scriptCmd)
}
