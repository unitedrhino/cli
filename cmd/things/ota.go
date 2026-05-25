package things

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var otaCmd = &cobra.Command{
	Use:   "ota <subcommand>",
	Short: "OTA 升级管理",
	Long:  `OTA 相关操作：固件管理、升级任务、模块管理。`,
	RunE:  wrapOldCommand(shared.CobraBridge{}.RunOta),
}

func init() {
	ThingsCmd.AddCommand(otaCmd)
}
