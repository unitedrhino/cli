package shared

import (
	"encoding/json"
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/config"
)

func runConfig(args []string, stdout, stderr io.Writer) int {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	if len(args) == 0 || args[0] == "--list" {
		raw, _ := json.MarshalIndent(cfg, "", "  ")
		_, _ = fmt.Fprintln(stdout, string(raw))
		return 0
	}
	if args[0] == "--use" {
		if len(args) < 2 {
			fmt.Fprintln(stderr, "--use requires profile name")
			return 2
		}
		cfg.CurrentProfile = args[1]
		if err := config.WriteConfig(cfg); err != nil {
			fmt.Fprintln(stderr, err.Error())
			return 1
		}
		fmt.Fprintf(stdout, "current profile: %s\n", args[1])
		return 0
	}
	fmt.Fprintf(stderr, "unknown config option: %s\n", args[0])
	return 2
}
