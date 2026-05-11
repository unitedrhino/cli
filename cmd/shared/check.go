package shared

import (
	"context"
	"fmt"
	"io"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/config"
)

func runCheck(ctx context.Context, app config.CLIApp, _ []string, stdout, stderr io.Writer) int {
	baseURL, err := config.GetBaseURL()
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	appID, err := config.GetAppID()
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	tenantCode, err := config.GetTenantCode()
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}

	fmt.Fprintf(stdout, "[OK] 应用: %s (%s)\n", app.DisplayName(), app.BinaryName())
	fmt.Fprintf(stdout, "[OK] 配置: baseURL=%s, appID=%s, tenantCode=%s\n", baseURL, appID, tenantCode)
	fmt.Fprintf(stdout, "[OK] 可调用权限: %s\n", strings.Join(app.AllowedAuthTypes(), ", "))

	// 输出功能概览
	features := app.Features()
	if len(features) > 0 {
		fmt.Fprintln(stdout, "\n功能模块:")
		printFeatureTree(stdout, features, "  ")
	}

	// 验证认证
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/get-one",
		Body: map[string]any{"withTenant": true},
	})
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "[FAIL] 认证失败: %s\n", resp.Msg)
		return 1
	}
	fmt.Fprintln(stdout, "\n[OK] 认证通过")
	return 0
}

func printFeatureTree(w io.Writer, features []config.Feature, indent string) {
	for _, f := range features {
		authNote := ""
		if len(f.Authority) > 0 {
			authNote = fmt.Sprintf(" [仅%s]", strings.Join(f.Authority, "/"))
		}
		fmt.Fprintf(w, "%s- %s: %s%s\n", indent, f.Name, f.Description, authNote)
		if len(f.SubFeatures) > 0 {
			printFeatureTree(w, f.SubFeatures, indent+"  ")
		}
	}
}
