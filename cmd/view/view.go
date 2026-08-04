// view.go — 大屏可视化命名空间父命令
package view

import "github.com/spf13/cobra"

// ViewCmd 是大屏可视化命名空间的根命令
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "大屏可视化命令",
	Long:  `大屏可视化相关操作：大屏项目管理（screen）、资源库管理（asset）。`,
}
