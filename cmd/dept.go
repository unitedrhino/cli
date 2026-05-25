package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var deptCmd = &cobra.Command{
	Use:   "dept <subcommand>",
	Short: "部门管理",
	Long:  `部门相关操作：部门用户查询、批量创建、批量删除。`,
	RunE:  wrapOldCommand(shared.CobraBridge{}.RunDept),
}

func init() {
	RootCmd.AddCommand(deptCmd)
}
