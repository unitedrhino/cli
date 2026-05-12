package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/config"
)

func runCheck(ctx context.Context, app config.CLIApp, args []string, stdout, stderr io.Writer) int {
	jsonMode := false
	for _, arg := range args {
		if arg == "--json" {
			jsonMode = true
		}
	}

	baseURL, err := config.GetBaseURL()
	if err != nil {
		outputCheckError(stderr, jsonMode, err)
		return 1
	}
	appID, err := config.GetAppID()
	if err != nil {
		outputCheckError(stderr, jsonMode, err)
		return 1
	}
	tenantCode, err := config.GetTenantCode()
	if err != nil {
		outputCheckError(stderr, jsonMode, err)
		return 1
	}

	if jsonMode {
		status := map[string]interface{}{
			"app_name":        app.DisplayName(),
			"binary_name":     app.BinaryName(),
			"base_url":        baseURL,
			"app_id":          appID,
			"tenant_code":     tenantCode,
			"allowed_auth":    app.AllowedAuthTypes(),
			"features":        app.Features(),
			"auth_status":     "checking",
			"auth_status_msg": "",
		}

		// 验证认证
		resp, err := client.DoAPI(ctx, client.APIRequest{
			Path: "/api/v1/system/user/self/get-one",
			Body: map[string]any{"withTenant": true},
		})
		if err != nil {
			status["auth_status"] = "error"
			status["auth_status_msg"] = err.Error()
			b, _ := json.Marshal(status)
			fmt.Fprintln(stdout, string(b))
			return 1
		}
		if resp.Code != 200 {
			status["auth_status"] = "failed"
			status["auth_status_msg"] = resp.Msg
			b, _ := json.Marshal(status)
			fmt.Fprintln(stdout, string(b))
			return 1
		}
		status["auth_status"] = "ok"
		b, _ := json.Marshal(status)
		fmt.Fprintln(stdout, string(b))
		return 0
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

func outputCheckError(stderr io.Writer, jsonMode bool, err error) {
	if jsonMode {
		b, _ := json.Marshal(map[string]interface{}{
			"auth_status":     "error",
			"auth_status_msg": err.Error(),
		})
		fmt.Fprintln(stderr, string(b))
	} else {
		fmt.Fprintln(stderr, err.Error())
	}
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
