package things

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var scriptCmd = &cobra.Command{
	Use:   "script <subcommand>",
	Short: "协议脚本管理",
	Long:  `协议脚本相关操作：验证、模板生成。`,
	RunE:  wrapOldCommandNoCtx(shared.CobraBridge{}.RunScript),
}

func init() {
	ThingsCmd.AddCommand(scriptCmd)
}
