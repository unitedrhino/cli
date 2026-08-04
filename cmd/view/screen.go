// screen.go — 大屏项目管理子命令（ur view screen）
package view

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var screenCmd = &cobra.Command{
	Use:   "screen <subcommand>",
	Short: "大屏项目管理",
	Long: `大屏项目相关操作：项目 CRUD、发布/取消发布、画布拉取/推送、内容校验、组件摘要、页面截图。

对应 HTTP API 前缀: /api/v1/view/project/

Subcommands:
  get-list     查询大屏项目列表 [--body '<json>']
  get-one      查询大屏项目详情 --id <projectID>
  create       创建大屏项目 --body '<json>'
  update       更新大屏项目 --body '<json>'
  delete       删除大屏项目 --id <projectID>
  publish      发布大屏项目 --id <projectID>
  unpublish    取消发布 --id <projectID>
  pull         拉取画布到本地 --id <projectID> [-o <file>]
  push         推送本地画布到远端 -f <file> [--id <projectID>] [--publish] [--force]
  validate     校验画布内容 (-f <file> 或 --id <projectID>)
  describe     组件状态摘要 (-f <file> 或 --id <projectID>) [--json]
  screenshot   大屏页面截图 --id <projectID> -o <png> [--front-base <url>] [--wait <sec>] [--edit]`,
	// 关闭 Cobra flag 解析：参数全部透传给 cmd/shared 的旧式 run 函数手动解析，
	// 避免 pflag 把 --id/-f 等业务参数当作未注册 flag 拒绝
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunViewScreen),
}

func init() {
	ViewCmd.AddCommand(screenCmd)
}
