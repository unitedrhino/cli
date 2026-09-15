package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var flowCmd = &cobra.Command{
	Use:   "flow <subcommand>",
	Short: "流程审批中心",
	Long:  `流程审批中心：流程定义/表单模板/分类管理（仅管理员）、流程发起、审批任务办理、实例管控与抄送。`,
	// 旧命令桥接自行解析 --id/--task-id 等 flag，必须关闭 cobra 的 flag 解析，
	// 否则 cobra 会把未知 flag 当解析错误拦截（与 alarm 命令同款做法）
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunFlow),
}

func init() {
	RootCmd.AddCommand(flowCmd)
}
