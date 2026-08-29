package things

import (
	"gitee.com/unitedrhino/cli/cmd/shared"
	"github.com/spf13/cobra"
)

var otaCmd = &cobra.Command{
	Use:                "ota <subcommand>",
	Short:              "OTA 升级管理",
	Long:               `OTA 相关操作：固件管理、升级任务、模块管理。`,
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunOta),
}

func init() {
	ThingsCmd.AddCommand(otaCmd)
}
