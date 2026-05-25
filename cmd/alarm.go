package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var alarmCmd = &cobra.Command{
	Use:   "alarm <subcommand>",
	Short: "告警管理",
	Long:  `告警相关操作：告警信息、告警记录、告警场景管理。`,
	RunE:  wrapOldCommand(shared.CobraBridge{}.RunAlarm),
}

func init() {
	RootCmd.AddCommand(alarmCmd)
}
