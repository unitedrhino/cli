package shared

import (
	"context"
	"fmt"
	"io"
)

// runAlarm 执行告警管理命令（新版告警模块，/api/v1/things/alarm/*）。
// 旧版规则引擎告警（/api/v1/things/rule/alarm/*）的后端接口已删除，对应
// info/record/scene 子命令已一并移除，避免调用不存在的 API。
func runAlarm(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printAlarmHelp(stdout)
		return 0
	}

	switch args[0] {
	case "rule":
		return runAlarmDmRule(ctx, args[1:], stdout, stderr)
	case "event":
		return runAlarmDmEvent(ctx, args[1:], stdout, stderr)
	case "notify-record":
		return runAlarmDmNotifyRecord(ctx, args[1:], stdout, stderr)
	case "notify-template":
		return runAlarmDmNotifyTemplate(ctx, args[1:], stdout, stderr)
	case "condition-template":
		return runAlarmDmConditionTemplate(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printAlarmHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown alarm subcommand: %s\n", args[0])
		printAlarmHelp(stderr)
		return 2
	}
}

// printAlarmHelp 打印告警管理帮助信息
func printAlarmHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur alarm <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Alarm management (/api/v1/things/alarm/*)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  rule                Alarm rule management (get-list, get-one, create, update, delete, status-update, evaluate-trigger)")
	fmt.Fprintln(w, "  event               Alarm event management (get-list, get-one, deal, false-alarm, stat)")
	fmt.Fprintln(w, "  notify-record       Alarm notify record management (get-list, resend)")
	fmt.Fprintln(w, "  notify-template     Alarm notify template management (get-list, get-one, create, update, delete, test-send)")
	fmt.Fprintln(w, "  condition-template  Alarm condition template management (get-list, get-one, create, update, delete)")
	fmt.Fprintln(w, "  help                Show this help message")
}
