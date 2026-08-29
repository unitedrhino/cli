package cmd

import (
	"gitee.com/unitedrhino/cli/cmd/shared"
	"github.com/spf13/cobra"
)

var alarmCmd = &cobra.Command{
	Use:   "alarm <subcommand>",
	Short: "告警管理",
	Long: `告警相关操作（新版告警模块，/api/v1/things/alarm/*）：
rule（告警规则）、event（告警事件）、notify-record（通知记录）、
notify-template（通知模板）、condition-template（触发条件模板）。`,
	// 旧命令桥接自行解析 --id/--body/--page 等 flag，必须关闭 cobra 的 flag 解析，
	// 否则 cobra 会把未知 flag 当解析错误拦截（与 skills/upgrade 命令同款做法）
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunAlarm),
}

func init() {
	RootCmd.AddCommand(alarmCmd)
}
