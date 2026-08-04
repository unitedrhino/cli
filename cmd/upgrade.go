package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "CLI 版本升级",
	Long:  `检查并升级 CLI 到最新版本，支持指定版本和 dry-run 模式。`,
	// 关闭 Cobra flag 解析：--dry-run/--version 等由 cmd/shared 手动解析
	DisableFlagParsing: true,
	RunE:               wrapOldCommandNoCtx(shared.CobraBridge{}.RunUpgrade),
}

func init() {
	RootCmd.AddCommand(upgradeCmd)
}
