package things

import (
	"gitee.com/unitedrhino/cli/cmd/shared"
	"github.com/spf13/cobra"
)

var aggCmd = &cobra.Command{
	Use:   "agg",
	Short: "设备属性聚合查询",
	Long: `查询设备属性的聚合值（平均值、最大值、最小值等）。

示例:
  ur things agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg
  ur things agg -p p_smartswitch_001 -i Temperature -f max,min`,
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunAgg),
}

func init() {
	ThingsCmd.AddCommand(aggCmd)
}
