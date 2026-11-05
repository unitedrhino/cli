package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var systemCmd = &cobra.Command{
	Use:   "system <subcommand>",
	Short: "系统管理",
	Long:  `系统相关操作：文件上传、批量聚合查询等。`,
	// 子命令（如 upload-file）使用手动 flag 解析，必须禁用 cobra 预解析，
	// 否则 --file 等标志被 cobra 当作未知标志拦截（与 alarm/flow/skills 同款修复）。
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunSystem),
}

func init() {
	RootCmd.AddCommand(systemCmd)
}
