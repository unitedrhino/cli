package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/internal/config"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "交互式初始化配置",
	Long:  `通过交互式向导配置 baseURL、appID、tenantCode、account、password。`,
	RunE:  runSetup,
}

func init() {
	RootCmd.AddCommand(setupCmd)
}

func runSetup(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(cmd.InOrStdin())

	// 默认使用当前解析的应用
	app := resolveAppFromContext()

	prompts := []struct {
		key   string
		label string
	}{
		{key: "baseURL", label: "baseURL"},
		{key: "appID", label: "appID"},
	}

	// 平台类应用 tenantCode 固定为 platform，组织类需要用户输入
	if app.DefaultTenantCode() == "" {
		prompts = append(prompts, struct{ key, label string }{key: "tenantCode", label: "tenantCode"})
	}

	prompts = append(prompts,
		struct{ key, label string }{key: "account", label: "account"},
		struct{ key, label string }{key: "password", label: "password"},
	)

	values := map[string]string{}
	values["appID"] = app.AppID()
	if tc := app.DefaultTenantCode(); tc != "" {
		values["tenantCode"] = tc
	}

	for _, prompt := range prompts {
		defaultVal := values[prompt.key]
		if defaultVal != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "%s [%s]: ", prompt.label, defaultVal)
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "%s: ", prompt.label)
		}
		line, err := reader.ReadString('\n')
		if err != nil && err != os.ErrInvalid {
			return &CLIError{Message: err.Error(), ExitCode: 1}
		}
		line = strings.TrimSpace(line)
		if line == "" && defaultVal != "" {
			continue
		}
		values[prompt.key] = line
	}

	tenantCode := values["tenantCode"]
	if tc := app.DefaultTenantCode(); tc != "" {
		tenantCode = tc
	}

	profile := config.Profile{
		BaseURL:    values["baseURL"],
		AppID:      values["appID"],
		TenantCode: tenantCode,
		Account:    values["account"],
		Password:   values["password"],
		Role:       string(app),
	}
	if err := config.SaveProfile(profile); err != nil {
		return &CLIError{Message: err.Error(), ExitCode: 1}
	}
	cmd.Printf("saved %s (app=%s)\n", config.ConfigPath(), app.DisplayName())
	return nil
}
