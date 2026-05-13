package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"gitee.com/unitedrhino/cli/internal/config"
)

// Execute 是所有 CLI 应用的统一入口
func Execute(app config.CLIApp, ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printHelp(app, stdout)
		return 0
	}
	switch args[0] {
	case "api":
		return runAPI(ctx, args[1:], stdout, stderr)
	case "check":
		return runCheck(ctx, app, args[1:], stdout, stderr)
	case "schema":
		return runSchema(app, args[1:], stdout, stderr)
	case "token":
		return runToken(ctx, args[1:], stdout, stderr)
	case "login":
		return runLogin(ctx, args[1:], stdout, stderr)
	case "setup":
		return runSetup(app, args[1:], stdout, stderr, os.Stdin)
	case "config":
		return runConfig(args[1:], stdout, stderr)
	case "generate-skills":
		return runGenerateSkills(app, args[1:], stdout, stderr)
	case "scene":
		return runScene(args[1:], stdout, stderr)
	case "script":
		return runScript(args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printHelp(app, stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n", args[0])
		printHelp(app, stderr)
		return 2
	}
}

// parseBodyArg 解析 JSON body 参数
func parseBodyArg(raw string) (map[string]any, error) {
	if strings.TrimSpace(raw) == "" {
		return map[string]any{}, nil
	}
	var out map[string]any
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		return nil, fmt.Errorf("body must be JSON object: %w", err)
	}
	return out, nil
}
