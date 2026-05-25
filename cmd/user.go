package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var userCmd = &cobra.Command{
	Use:   "user <subcommand>",
	Short: "用户管理",
	Long:  `用户相关操作：用户信息查询、创建、更新、删除，以及个人中心操作。`,
	RunE:  wrapOldCommand(shared.CobraBridge{}.RunUser),
}

func init() {
	RootCmd.AddCommand(userCmd)
}
