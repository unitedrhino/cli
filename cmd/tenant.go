package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var tenantCmd = &cobra.Command{
	Use:   "tenant <subcommand>",
	Short: "企业管理",
	Long:  `企业相关操作：企业用户管理、角色管理、邀请管理等。`,
	RunE:  wrapOldCommand(shared.CobraBridge{}.RunTenant),
}

func init() {
	RootCmd.AddCommand(tenantCmd)
}
