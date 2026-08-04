// asset.go — 大屏资源库管理子命令（ur view asset）
package view

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var assetCmd = &cobra.Command{
	Use:   "asset <subcommand>",
	Short: "大屏资源库管理",
	Long: `大屏资源库相关操作：资源查询、上传（先传文件再登记资源记录）、删除。

对应 HTTP API 前缀: /api/v1/view/asset/

Subcommands:
  get-list   查询资源列表 [--name <name>] [--type <image|video|audio|other>] [--group-id <id>] [--format <ext>] [--page <n>] [--size <n>]
  upload     上传资源 -f <file> [--name <name>] [--group-id <id>] [--tags <a,b>]
  delete     删除资源 --id <assetID>`,
	// 关闭 Cobra flag 解析：参数全部透传给 cmd/shared 的旧式 run 函数手动解析，
	// 避免 pflag 把 --id/-f 等业务参数当作未注册 flag 拒绝
	DisableFlagParsing: true,
	RunE:               wrapOldCommand(shared.CobraBridge{}.RunViewAsset),
}

func init() {
	ViewCmd.AddCommand(assetCmd)
}
