package ur

import (
	"context"
	"fmt"
	"io"
	"os"

	"gitee.com/unitedrhino/cli/cmd"
	"gitee.com/unitedrhino/cli/internal/config"
)

// Execute 是 ur 二进制入口
func Execute(ctx context.Context, version string, args []string, stdout, stderr io.Writer) int {
	app, filtered := resolveApp(args)
	// 顶层版本查询（ur --version / -v / ur --app iot --version）；
	// 只检查过滤后的首参数，避免子命令参数（如 ur upgrade --version v0.3.5）被误判
	if len(filtered) > 0 && (filtered[0] == "--version" || filtered[0] == "-v") {
		fmt.Fprintln(stdout, version)
		return 0
	}
	return cmd.Execute(app, version, filtered, stdout, stderr)
}

func resolveApp(args []string) (config.CLIApp, []string) {
	filtered := make([]string, 0, len(args))
	var app config.CLIApp
	for i := 0; i < len(args); i++ {
		if args[i] == "--app" && i+1 < len(args) {
			if parsed, err := config.ParseCLIApp(args[i+1]); err == nil {
				app = parsed
			}
			i++ // skip value
			continue
		}
		filtered = append(filtered, args[i])
	}
	if app == "" {
		if envApp := os.Getenv("UR_APP"); envApp != "" {
			if parsed, err := config.ParseCLIApp(envApp); err == nil {
				app = parsed
			}
		}
	}
	if app == "" {
		app = config.AppOrgManage
	}
	return app, filtered
}
