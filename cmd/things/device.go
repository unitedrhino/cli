package things

import (
	"gitee.com/unitedrhino/cli/cmd/shared"
	"github.com/spf13/cobra"
)

var deviceCmd = &cobra.Command{
	Use:   "device <subcommand>",
	Short: "设备调试与管理",
	Long: `设备相关操作：日志查询、属性控制、动作下发、模拟上报等。

对应 HTTP API 前缀: /api/v1/things/device/`,
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunDevice),
}

func init() {
	ThingsCmd.AddCommand(deviceCmd)
}
