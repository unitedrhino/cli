package things

import (
	"gitee.com/unitedrhino/cli/cmd/shared"
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:                "model <subcommand>",
	Short:              "物模型模板与验证",
	Long:               `物模型模板生成、验证、脚本生成。`,
	DisableFlagParsing: true,
	RunE:               wrapOldCommandNoCtx(shared.CobraBridge{}.RunModel),
}

func init() {
	ThingsCmd.AddCommand(modelCmd)
}
