package things

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var deviceCmd = &cobra.Command{
	Use:   "device <subcommand>",
	Short: "设备调试与管理",
	Long: `设备相关操作：日志查询、属性控制、动作下发、模拟上报等。

对应 HTTP API 前缀: /api/v1/things/device/`,
	RunE: wrapOldCommand(shared.CobraBridge{}.RunDevice),
}

func init() {
	ThingsCmd.AddCommand(deviceCmd)
}
