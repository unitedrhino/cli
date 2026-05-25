// things.go — 物联网命名空间父命令
package things

import "github.com/spf13/cobra"

// ThingsCmd 是物联网命名空间的根命令
var ThingsCmd = &cobra.Command{
	Use:   "things",
	Short: "物联网服务命令",
	Long:  `物联网相关操作：设备、物模型、场景联动、协议脚本、OTA、聚合查询等。`,
}
