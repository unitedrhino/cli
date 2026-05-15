package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gitee.com/unitedrhino/cli/internal/config"
	"gitee.com/unitedrhino/cli/internal/version"
)

// Execute 是所有 CLI 应用的统一入口
func Execute(app config.CLIApp, ctx context.Context, ver string, args []string, stdout, stderr io.Writer) int {
	version.BuildVersion = ver
	if handleVersionFlag(args, stdout) {
		return 0
	}
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
	case "completion":
		return runCompletion(args[1:], stdout, stderr)
	case "scene":
		return runScene(args[1:], stdout, stderr)
	case "script":
		return runScript(args[1:], stdout, stderr)
	case "model":
		return runModel(args[1:], stdout, stderr)
	case "upgrade":
		return runUpgrade(args[1:], stdout, stderr)
	case "skills", "skill":
		return runSkills(args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printHelp(app, stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n", args[0])
		printHelp(app, stderr)
		return 2
	}
}

// handleVersionFlag 处理 --version / -v 标志，支持 --json 输出
func handleVersionFlag(args []string, stdout io.Writer) bool {
	hasVersion := false
	hasJSON := false
	for _, arg := range args {
		if arg == "--version" || arg == "-v" {
			hasVersion = true
		}
		if arg == "--json" {
			hasJSON = true
		}
	}
	if !hasVersion {
		return false
	}
	if hasJSON {
		binaryPath, _ := os.Executable()
		if binaryPath == "" {
			if cwd, err := os.Getwd(); err == nil {
				binaryPath = filepath.Join(cwd, "ur")
			}
		}
		fmt.Fprintln(stdout, version.FormatVersionJSON(binaryPath))
	} else {
		fmt.Fprintln(stdout, version.BuildVersion)
	}
	return true
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
