package things

import (
	"os"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
	"gitee.com/unitedrhino/cli/internal/config"
)

var schemaCmd = &cobra.Command{
	Use:   "schema <subcommand>",
	Short: "物模型管理",
	Long:  `物模型相关操作：浏览、查询、创建、更新、删除、导入 TSL。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app := resolveAppFromContext()
		code := shared.CobraBridge{}.RunSchema(app, args, cmd.OutOrStdout(), cmd.ErrOrStderr())
		if code != 0 {
			return &CLIError{Message: "", ExitCode: code}
		}
		return nil
	},
}

func init() {
	ThingsCmd.AddCommand(schemaCmd)
}

func resolveAppFromContext() config.CLIApp {
	appID := os.Getenv("UR_APP_ID")
	for _, a := range config.AllCLIApps() {
		if a.AppID() == appID {
			return a
		}
	}
	return config.AppOrgManage
}

