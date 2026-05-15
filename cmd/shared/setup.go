package shared

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"gitee.com/unitedrhino/cli/internal/config"
)

func runSetup(app config.CLIApp, _ []string, stdout, stderr io.Writer, in io.Reader) int {
	reader := bufio.NewReader(in)
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
	// 设置默认值
	values["appID"] = app.AppID()
	if tc := app.DefaultTenantCode(); tc != "" {
		values["tenantCode"] = tc
	}

	for _, prompt := range prompts {
		defaultVal := values[prompt.key]
		if defaultVal != "" {
			fmt.Fprintf(stdout, "%s [%s]: ", prompt.label, defaultVal)
		} else {
			fmt.Fprintf(stdout, "%s: ", prompt.label)
		}
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintln(stderr, err.Error())
			return 1
		}
		line = strings.TrimSpace(line)
		if line == "" && defaultVal != "" {
			continue // 使用默认值
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
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	fmt.Fprintf(stdout, "saved %s (app=%s)\n", config.ConfigPath(), app.DisplayName())
	return 0
}
