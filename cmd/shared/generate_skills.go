package shared

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gitee.com/unitedrhino/cli/internal/config"
	"gitee.com/unitedrhino/cli/internal/swagger"
)

func runGenerateSkills(app config.CLIApp, args []string, stdout, stderr io.Writer) int {
	outputDir := filepath.Join(".", "skill", app.BinaryName())
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--output", "-o":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--output requires directory")
				return 2
			}
			outputDir = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown generate-skills option: %s\n", args[i])
			return 2
		}
	}
	endpoints, err := swagger.LoadEndpoints()
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}

	// 按应用的 AllowedAuthTypes 过滤
	filtered := swagger.FilterEndpointsByApp(endpoints, app.AllowedAuthTypes())

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}

	// 生成 SKILL.md
	skillMD := generateSkillMD(app, filtered)
	skillPath := filepath.Join(outputDir, "SKILL.md")
	if err := os.WriteFile(skillPath, []byte(skillMD), 0o644); err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	fmt.Fprintf(stdout, "generated %s\n", skillPath)

	// 生成 swagger-index.md
	indexPath := filepath.Join(outputDir, "swagger-index.md")
	var builder strings.Builder
	builder.WriteString("# Swagger Index\n\n")
	builder.WriteString(fmt.Sprintf("> 本文件由 `%s generate-skills` 自动生成。\n\n", app.BinaryName()))
	for _, item := range filtered {
		summary := item.Summary
		if summary == "" {
			summary = item.Description
		}
		builder.WriteString(fmt.Sprintf("- `%s %s` [%s] %s\n", item.Method, item.Path, item.AuthType, summary))
	}
	if err := os.WriteFile(indexPath, []byte(builder.String()), 0o644); err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	fmt.Fprintf(stdout, "generated %s\n", indexPath)

	return 0
}

// hermesTags 映射：CLI app → Hermes tags
var hermesTags = map[string]string{
	"ur-platform-manage": "[platform, admin, tenant, user, auth, system]",
	"ur-iot":             "[iot, device, product, project, ota, protocol, thing-model]",
	"ur-org-manage":      "[org, user, role, ai, agent, department]",
	"ur-org-energy":      "[energy, power, prepaid, device, consumption, automation]",
	"ur-console":         "[console, profile, token, settings, personal]",
}

func generateSkillMD(app config.CLIApp, endpoints []swagger.Endpoint) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("---\nname: %s\n", app.BinaryName()))
	b.WriteString(fmt.Sprintf("description: \"%s — 联犀 SaaS 平台 %s CLI 工具\"\n", app.BinaryName(), app.DisplayName()))
	if tags, ok := hermesTags[app.BinaryName()]; ok {
		b.WriteString("metadata:\n  hermes:\n    tags: " + tags + "\n")
	}
	b.WriteString("---\n\n")
	b.WriteString(fmt.Sprintf("# %s — %s\n\n", app.BinaryName(), app.DisplayName()))

	// 配置状态提示
	b.WriteString("> **配置检查**：如果尚未配置联犀连接，请先运行 `" + app.BinaryName() + " login --no-wait`，按指引在浏览器中完成授权。")
	b.WriteString("`setup` 命令是终端交互式的，在 AI 聊天环境中无法使用。\n\n")

	// 应用信息
	b.WriteString("## 应用信息\n\n")
	b.WriteString(fmt.Sprintf("- **AppID**: %s\n", app.AppID()))
	tc := app.DefaultTenantCode()
	if tc == "" {
		tc = "用户输入"
	}
	b.WriteString(fmt.Sprintf("- **TenantCode**: %s\n", tc))
	b.WriteString(fmt.Sprintf("- **可调用权限**: %s\n\n", strings.Join(app.AllowedAuthTypes(), ", ")))

	// 功能概览
	features := app.Features()
	if len(features) > 0 {
		b.WriteString("## 功能概览\n\n")
		writeFeatures(&b, features, 0)
		b.WriteString("\n")
	}

	// API 端点统计
	b.WriteString("## API 端点\n\n")
	b.WriteString(fmt.Sprintf("共 %d 个可调用端点（按 %s 权限过滤）。\n\n", len(endpoints), strings.Join(app.AllowedAuthTypes(), "/")))

	// 按 group 分组
	groups := map[string][]swagger.Endpoint{}
	for _, ep := range endpoints {
		groups[ep.Group] = append(groups[ep.Group], ep)
	}
	for group, eps := range groups {
		b.WriteString(fmt.Sprintf("### %s\n\n", group))
		for _, ep := range eps {
			summary := ep.Summary
			if summary == "" {
				summary = ep.Description
			}
			b.WriteString(fmt.Sprintf("- `%s %s` — %s\n", ep.Method, ep.Path, summary))
		}
		b.WriteString("\n")
	}

	// 使用示例
	b.WriteString("## 使用示例\n\n")
	b.WriteString("```bash\n")
	b.WriteString(fmt.Sprintf("# 配置\n%s setup\n\n", app.BinaryName()))
	b.WriteString(fmt.Sprintf("# 验证连通性\n%s check\n\n", app.BinaryName()))
	b.WriteString(fmt.Sprintf("# 调用 API\n%s api /api/v1/system/user/self/get-one\n", app.BinaryName()))
	b.WriteString("```\n")

	return b.String()
}

func writeFeatures(b *strings.Builder, features []config.Feature, depth int) {
	indent := strings.Repeat("  ", depth)
	for _, f := range features {
		authNote := ""
		if len(f.Authority) > 0 {
			authNote = fmt.Sprintf(" `[仅%s]`", strings.Join(f.Authority, "/"))
		}
		b.WriteString(fmt.Sprintf("%s- **%s**: %s%s\n", indent, f.Name, f.Description, authNote))
		if len(f.APIs) > 0 && depth == 0 {
			b.WriteString(fmt.Sprintf("%s  API: `%s`\n", indent, strings.Join(f.APIs, "`, `")))
		}
		if len(f.SubFeatures) > 0 {
			writeFeatures(b, f.SubFeatures, depth+1)
		}
	}
}
