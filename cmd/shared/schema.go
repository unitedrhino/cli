package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"gitee.com/unitedrhino/cli/internal/config"
	"gitee.com/unitedrhino/cli/internal/swagger"
)

func runSchema(app config.CLIApp, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	authType := ""
	targetPath := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json", "-j":
			jsonOutput = true
		case "--auth-type":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--auth-type requires value")
				return 2
			}
			authType = args[i+1]
			i++
		default:
			if strings.HasPrefix(args[i], "-") {
				fmt.Fprintf(stderr, "unknown schema option: %s\n", args[i])
				return 2
			}
			targetPath = args[i]
		}
	}
	endpoints, err := swagger.LoadEndpoints()
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}

	// 按路径过滤
	if targetPath != "" {
		endpoints = swagger.FilterEndpoints(endpoints, targetPath, "")
	}

	// 按 authType 过滤：显式指定则用显式值，否则按应用默认
	if authType != "" {
		endpoints = swagger.FilterEndpoints(endpoints, "", authType)
	} else {
		endpoints = swagger.FilterEndpointsByApp(endpoints, app.AllowedAuthTypes())
	}

	if jsonOutput {
		raw, _ := json.MarshalIndent(endpoints, "", "  ")
		_, _ = fmt.Fprintln(stdout, string(raw))
		return 0
	}
	for _, item := range endpoints {
		summary := item.Summary
		if summary == "" {
			summary = item.Description
		}
		fmt.Fprintf(stdout, "%s %-6s [%s] %s\n", item.Path, item.Method, item.AuthType, summary)
	}
	return 0
}
