package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/config"
)

var checkOpts struct {
	jsonMode bool
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "检查当前配置和认证状态",
	Long:  `检查 CLI 的当前配置、认证状态及可用功能模块。`,
	RunE:  runCheck,
}

func init() {
	checkCmd.Flags().BoolVar(&checkOpts.jsonMode, "json", false, "JSON 格式输出")
	RootCmd.AddCommand(checkCmd)
}

func runCheck(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	jsonMode := checkOpts.jsonMode

	baseURL, err := config.GetBaseURL()
	if err != nil {
		return outputCheckError(cmd, jsonMode, err)
	}
	appID, err := config.GetAppID()
	if err != nil {
		return outputCheckError(cmd, jsonMode, err)
	}
	tenantCode, err := config.GetTenantCode()
	if err != nil {
		return outputCheckError(cmd, jsonMode, err)
	}

	app := resolveAppFromContext()

	if jsonMode {
		status := map[string]any{
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

		resp, err := client.DoAPI(ctx, client.APIRequest{
			Path: "/api/v1/system/user/self/get-one",
			Body: map[string]any{"withTenant": true},
		})
		if err != nil {
			status["auth_status"] = "error"
			status["auth_status_msg"] = err.Error()
			b, _ := json.Marshal(status)
			cmd.Println(string(b))
			return &CLIError{Message: err.Error(), ExitCode: 1}
		}
		if resp.Code != 200 {
			status["auth_status"] = "failed"
			status["auth_status_msg"] = resp.Msg
			b, _ := json.Marshal(status)
			cmd.Println(string(b))
			return &CLIError{Message: resp.Msg, ExitCode: 1}
		}
		status["auth_status"] = "ok"
		b, _ := json.Marshal(status)
		cmd.Println(string(b))
		return nil
	}

	cmd.Printf("[OK] 应用: %s (%s)\n", app.DisplayName(), app.BinaryName())
	cmd.Printf("[OK] 配置: baseURL=%s, appID=%s, tenantCode=%s\n", baseURL, appID, tenantCode)
	cmd.Printf("[OK] 可调用权限: %s\n", strings.Join(app.AllowedAuthTypes(), ", "))

	features := app.Features()
	if len(features) > 0 {
		cmd.Println("\n功能模块:")
		printFeatureTree(cmd.OutOrStdout(), features, "  ")
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/user/self/get-one",
		Body: map[string]any{"withTenant": true},
	})
	if err != nil {
		return &CLIError{Message: err.Error(), ExitCode: 1}
	}
	if resp.Code != 200 {
		return &CLIError{Message: fmt.Sprintf("[FAIL] 认证失败: %s", resp.Msg), ExitCode: 1}
	}
	cmd.Println("\n[OK] 认证通过")
	return nil
}

func outputCheckError(cmd *cobra.Command, jsonMode bool, err error) error {
	if jsonMode {
		b, _ := json.Marshal(map[string]any{
			"auth_status":     "error",
			"auth_status_msg": err.Error(),
		})
		cmd.PrintErrln(string(b))
	} else {
		cmd.PrintErrln(err.Error())
	}
	return &CLIError{Message: err.Error(), ExitCode: 1}
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

// resolveAppFromContext 从当前执行上下文解析 CLIApp
func resolveAppFromContext() config.CLIApp {
	appID := os.Getenv("UR_APP_ID")
	for _, a := range config.AllCLIApps() {
		if a.AppID() == appID {
			return a
		}
	}
	return config.AppOrgManage
}
