// ai.go — AI 中台命名空间父命令
package ai

import "github.com/spf13/cobra"

// AICmd 是 AI 命名空间的根命令
var AICmd = &cobra.Command{
	Use:   "ai",
	Short: "AI 中台服务命令",
	Long:  `AI 相关操作：AI 工具、智能体、知识库等。`,
}
