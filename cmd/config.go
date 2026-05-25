package cmd

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "管理 CLI 配置",
	Long:  `查看和管理 CLI 的配置文件、当前 profile 及企业号切换。`,
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出当前配置",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return &CLIError{Message: err.Error(), ExitCode: 1}
		}
		raw, _ := json.MarshalIndent(cfg, "", "  ")
		cmd.Println(string(raw))
		return nil
	},
}

var configUseCmd = &cobra.Command{
	Use:   "use <profile>",
	Short: "切换当前 profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return &CLIError{Message: err.Error(), ExitCode: 1}
		}
		cfg.CurrentProfile = args[0]
		if err := config.WriteConfig(cfg); err != nil {
			return &CLIError{Message: err.Error(), ExitCode: 1}
		}
		cmd.Printf("current profile: %s\n", args[0])
		return nil
	},
}

var configTenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "管理企业号",
	Long:  `查看当前企业号、列出所有 profile 的企业号、切换企业号。`,
	RunE:  runConfigTenant,
}

var configTenantListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有 profile 的企业号",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return &CLIError{Message: fmt.Sprintf("读取配置失败: %v", err), ExitCode: 1}
		}
		cmd.Println("Profile  →  企业号")
		cmd.Println("--------     ------")
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
			cmd.Printf("%s %-12s %s\n", marker, name, cfg.Profiles[name].TenantCode)
		}
		return nil
	},
}

var configTenantUseCmd = &cobra.Command{
	Use:   "use <企业号>",
	Short: "切换当前 profile 的企业号",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return &CLIError{Message: fmt.Sprintf("读取配置失败: %v", err), ExitCode: 1}
		}
		name := cfg.CurrentProfile
		if name == "" {
			name = "default"
			cfg.CurrentProfile = "default"
		}
		profile := cfg.Profiles[name]
		oldTenantCode := profile.TenantCode
		profile.TenantCode = args[0]
		cfg.Profiles[name] = profile
		if err := config.WriteConfig(cfg); err != nil {
			return &CLIError{Message: fmt.Sprintf("写入配置失败: %v", err), ExitCode: 1}
		}
		cmd.Printf("已切换企业: %s → %s\n", oldTenantCode, args[0])
		return nil
	},
}

func init() {
	configTenantCmd.AddCommand(configTenantListCmd)
	configTenantCmd.AddCommand(configTenantUseCmd)

	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configUseCmd)
	configCmd.AddCommand(configTenantCmd)
	RootCmd.AddCommand(configCmd)
}

func runConfigTenant(cmd *cobra.Command, args []string) error {
	// 无参数时查看当前企业号
	if len(args) == 0 {
		profile, err := config.CurrentProfile()
		if err != nil {
			return &CLIError{Message: fmt.Sprintf("读取当前 profile 失败: %v", err), ExitCode: 1}
		}
		cmd.Printf("当前企业: %s\n", profile.TenantCode)
		return nil
	}
	return cmd.Help()
}
