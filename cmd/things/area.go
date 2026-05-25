package things

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var areaCmd = &cobra.Command{
	Use:   "area <subcommand>",
	Short: "区域管理",
	Long:  `区域相关操作：区域信息、区域配置查询与管理。`,
	RunE:  wrapOldCommand(shared.CobraBridge{}.RunArea),
}

func init() {
	ThingsCmd.AddCommand(areaCmd)
}
