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
	// 再次检查 --version（处理 ur --app iot --version 场景）
	for _, arg := range args {
		if arg == "--version" || arg == "-v" {
			fmt.Fprintln(stdout, version)
			return 0
		}
	}

	app, filtered := resolveApp(args)
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
