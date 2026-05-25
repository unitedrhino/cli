package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var systemCmd = &cobra.Command{
	Use:   "system <subcommand>",
	Short: "系统管理",
	Long:  `系统相关操作：文件上传、批量聚合查询等。`,
	RunE:  wrapOldCommand(shared.CobraBridge{}.RunSystem),
}

func init() {
	RootCmd.AddCommand(systemCmd)
}
