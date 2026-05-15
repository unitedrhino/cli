package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"gitee.com/unitedrhino/cli/internal/config"
)

func runConfig(args []string, stdout, stderr io.Writer) int {
	if len(args) > 0 && args[0] == "tenant" {
		return runConfigTenant(args[1:], stdout, stderr)
	}

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

func runConfigTenant(args []string, stdout, stderr io.Writer) int {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Fprintf(stderr, "读取配置失败: %v\n", err)
		return 1
	}

	if len(args) == 0 {
		// 查看当前企业号
		profile, err := config.CurrentProfile()
		if err != nil {
			fmt.Fprintf(stderr, "读取当前 profile 失败: %v\n", err)
			return 1
		}
		fmt.Fprintf(stdout, "当前企业: %s\n", profile.TenantCode)
		return 0
	}

	switch args[0] {
	case "--list", "-l":
		// 列出所有 profile 的企业号
		fmt.Fprintln(stdout, "Profile  →  企业号")
		fmt.Fprintln(stdout, "--------     ------")
		names := make([]string, 0, len(cfg.Profiles))
		for name := range cfg.Profiles {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			marker := " "
			if name == cfg.CurrentProfile {
				marker = "*"
			}
			fmt.Fprintf(stdout, "%s %-12s %s\n", marker, name, cfg.Profiles[name].TenantCode)
		}
		return 0

	case "--use":
		if len(args) < 2 {
			fmt.Fprintln(stderr, "用法: ur config tenant --use <企业号>")
			return 2
		}
		newTenantCode := args[1]
		// 读取当前 profile
		name := cfg.CurrentProfile
		if name == "" {
			name = "default"
			cfg.CurrentProfile = "default"
		}
		profile := cfg.Profiles[name]
		oldTenantCode := profile.TenantCode
		profile.TenantCode = newTenantCode
		cfg.Profiles[name] = profile
		if err := config.WriteConfig(cfg); err != nil {
			fmt.Fprintf(stderr, "写入配置失败: %v\n", err)
			return 1
		}
		fmt.Fprintf(stdout, "已切换企业: %s → %s\n", oldTenantCode, newTenantCode)
		return 0

	case "-h", "--help":
		printConfigTenantHelp(stdout)
		return 0

	default:
		fmt.Fprintf(stderr, "未知参数: %s\n", args[0])
		printConfigTenantHelp(stderr)
		return 2
	}
}

func printConfigTenantHelp(w io.Writer) {
	fmt.Fprintln(w, "用法: ur config tenant [选项]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "选项:")
	fmt.Fprintln(w, "  (无参数)          查看当前企业号")
	fmt.Fprintln(w, "  --list, -l        列出所有 profile 的企业号")
	fmt.Fprintln(w, "  --use <企业号>     切换当前 profile 的企业号")
	fmt.Fprintln(w, "  -h, --help        显示帮助")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "示例:")
	fmt.Fprintln(w, "  ur config tenant              # 查看当前企业号")
	fmt.Fprintln(w, "  ur config tenant --list       # 列出所有企业号")
	fmt.Fprintln(w, "  ur config tenant --use t2     # 切换到企业 t2")
}
