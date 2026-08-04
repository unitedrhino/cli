package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/ai"
	"gitee.com/unitedrhino/cli/cmd/generated"
	"gitee.com/unitedrhino/cli/cmd/things"
	"gitee.com/unitedrhino/cli/cmd/view"
	"gitee.com/unitedrhino/cli/internal/cmdutil"
	"gitee.com/unitedrhino/cli/internal/config"
)

// RootCmd 是 ur 的根命令
var RootCmd = &cobra.Command{
	Use:   "ur",
	Short: "联犀 SaaS 平台命令行工具",
	Long: `ur 是联犀 SaaS 平台的官方 CLI 工具，支持设备管理、物模型操作、
AI 工具开发、API 调用等功能。

通过访问令牌认证，支持多企业、多应用切换。`,
	SilenceUsage:  true,
	SilenceErrors: false,
}

var (
	appFlag string
	factory *cmdutil.Factory
)

func init() {
	RootCmd.PersistentFlags().StringVar(&appFlag, "app", "",
		"应用上下文 (platform-manage, iot, org-manage, org-energy, console)")

	// 注册命名空间父命令
	RootCmd.AddCommand(things.ThingsCmd)
	RootCmd.AddCommand(ai.AICmd)
	RootCmd.AddCommand(view.ViewCmd)

	// 注册 Layer 2 自动生成命令
	generated.RegisterSystemCommands(RootCmd)
	generated.RegisterThingsCommands(RootCmd)
}

// Execute 执行根命令，返回退出码
func Execute(app config.CLIApp, version string, args []string, stdout, stderr io.Writer) int {
	RootCmd.SetOut(stdout)
	RootCmd.SetErr(stderr)
	RootCmd.SetArgs(args)
	RootCmd.Version = version

	// 解析 --app 并注入环境变量
	resolveAppAndInject(app)

	if err := RootCmd.Execute(); err != nil {
		if exitErr, ok := err.(interface{ ExitCode() int }); ok {
			return exitErr.ExitCode()
		}
		return 1
	}
	return 0
}

func resolveAppAndInject(defaultApp config.CLIApp) {
	app := defaultApp
	if appFlag != "" {
		if parsed, err := config.ParseCLIApp(appFlag); err == nil {
			app = parsed
		}
	}
	// 仅在环境变量未设置时注入默认值，避免覆盖 sandbox 等外部注入的配置
	if appID := app.AppID(); appID != "" {
		if os.Getenv("UR_APP_ID") == "" {
			os.Setenv("UR_APP_ID", appID)
		}
	}
	if tc := app.DefaultTenantCode(); tc != "" {
		if os.Getenv("UR_TENANT_CODE") == "" {
			os.Setenv("UR_TENANT_CODE", tc)
		}
	}
}

// CLIError 提供可控退出码的错误类型
type CLIError struct {
	Message  string
	ExitCode int
}

func (e CLIError) Error() string { return e.Message }
