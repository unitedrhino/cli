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
	allEndpoints := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--output", "-o":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--output requires directory")
				return 2
			}
			outputDir = args[i+1]
			i++
		case "--all":
			allEndpoints = true
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

	var filtered []swagger.Endpoint
	if allEndpoints {
		filtered = endpoints
	} else {
		// 按应用的 AllowedAuthTypes 过滤
		filtered = swagger.FilterEndpointsByApp(endpoints, app.AllowedAuthTypes())
	}

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}

	// 生成 SKILL.md
	skillMD := generateSkillMD(app, filtered, allEndpoints)
	skillPath := filepath.Join(outputDir, "SKILL.md")
	if err := os.WriteFile(skillPath, []byte(skillMD), 0o644); err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	fmt.Fprintf(stdout, "generated %s\n", skillPath)

	// 生成 swagger-index.md（始终生成完整索引）
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

	// 当 allEndpoints=true 时，按路径前缀拆分为多个引用文件，避免单个文档过大
	if allEndpoints {
		refDir := filepath.Join(outputDir, "references")
		if err := os.MkdirAll(refDir, 0o755); err != nil {
			fmt.Fprintln(stderr, err.Error())
			return 1
		}
		prefixes := map[string]string{
			"system-apis.md":  "/api/v1/system/",
			"things-apis.md":  "/api/v1/things/",
			"ai-apis.md":      "/api/v1/ai/",
		}
		for filename, prefix := range prefixes {
			var refBuilder strings.Builder
			refBuilder.WriteString(fmt.Sprintf("# %s API Index\n\n", strings.TrimSuffix(filename, ".md")))
			refBuilder.WriteString(fmt.Sprintf("> 路径前缀 `%s` 的端点列表。\n\n", prefix))
			count := 0
			for _, item := range filtered {
				if !strings.HasPrefix(item.Path, prefix) {
					continue
				}
				summary := item.Summary
				if summary == "" {
					summary = item.Description
				}
				refBuilder.WriteString(fmt.Sprintf("- `%s %s` [%s] %s\n", item.Method, item.Path, item.AuthType, summary))
				count++
			}
			if count > 0 {
				refPath := filepath.Join(refDir, filename)
				if err := os.WriteFile(refPath, []byte(refBuilder.String()), 0o644); err != nil {
					fmt.Fprintln(stderr, err.Error())
					return 1
				}
				fmt.Fprintf(stdout, "generated %s (%d endpoints)\n", refPath, count)
			}
		}
	}

	return 0
}

// hermesTags 映射：CLI app → Hermes tags
var hermesTags = map[string]string{
	"ur-platform-manage": "[platform, admin, tenant, user, auth, system]",
	"ur-iot":             "[iot, device, product, project, ota, protocol, thing-model, rule, schema]",
	"ur-org-manage":      "[org, user, role, ai, agent, department]",
	"ur-org-energy":      "[energy, power, prepaid, device, consumption, automation]",
	"ur-console":         "[console, profile, token, settings, personal]",
	"ur-protocol":        "[protocol, gateway, script, container, image]",
	"ur-ota":             "[ota, firmware, upgrade, module]",
	"ur-rule":            "[rule, alarm, scene, automation]",
	"ur-schema":          "[schema, thing-model, tsl]",
	"ur-iot-user":        "[device-share, device-collect, user-device]",
	"ur-iot-config":      "[config, iot-settings]",
	"ur-iot-hook":        "[hook, webhook]",
}

func generateSkillMD(app config.CLIApp, endpoints []swagger.Endpoint, allEndpoints bool) string {
	var b strings.Builder

	name := app.BinaryName()
	display := app.DisplayName()
	if allEndpoints {
		name = "ur-api"
		display = "联犀 SaaS 平台"
	}

	b.WriteString(fmt.Sprintf("---\nname: %s\n", name))
	if allEndpoints {
		b.WriteString(fmt.Sprintf("description: \"%s — 联犀 SaaS 平台统一 API 工具（涵盖所有应用）\"\n", name))
		b.WriteString("metadata:\n  hermes:\n    tags: [api, cli, saas, iot, platform, org, energy, console]\n")
	} else {
		b.WriteString(fmt.Sprintf("description: \"%s — 联犀 SaaS 平台 %s CLI 工具\"\n", name, display))
		if tags, ok := hermesTags[name]; ok {
			b.WriteString("metadata:\n  hermes:\n    tags: " + tags + "\n")
		}
	}
	b.WriteString("---\n\n")
	b.WriteString(fmt.Sprintf("# %s — %s\n\n", name, display))

	// 配置状态提示
	b.WriteString("> **配置检查**：如果尚未配置联犀连接，请先运行 `" + name + " login --no-wait`，按指引在浏览器中完成授权。")
	b.WriteString("`setup` 命令是终端交互式的，在 AI 聊天环境中无法使用。\n\n")

	if allEndpoints {
		// 统一 skill：列出所有应用
		b.WriteString("## 应用选择\n\n")
		b.WriteString("根据当前操作的前端应用选择对应的 CLI：\n\n")
		b.WriteString("| CLI | 前端应用 | AppID | TenantCode | 可调用权限 |\n")
		b.WriteString("|-----|---------|-------|------------|-----------|\n")
		for _, a := range config.AllCLIApps() {
			tc := a.DefaultTenantCode()
			if tc == "" {
				tc = "用户输入"
			}
			b.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s | %s |\n", a.BinaryName(), a.DisplayName(), a.AppID(), tc, strings.Join(a.AllowedAuthTypes(), ", ")))
		}
		b.WriteString("\n")

		// 常见查询速查：当子 skill 不可用时，LLM 可直接从主文档获取常用 API
		b.WriteString("## 常见查询速查\n\n")
		b.WriteString("以下是最常用的查询场景及对应 API。优先尝试这些接口，无需查阅子 skill：\n\n")
		b.WriteString("| 查询场景 | 推荐 CLI | API 命令示例 |\n")
		b.WriteString("|---------|---------|-------------|\n")
		b.WriteString("| 查询【我的】应用列表 | `ur-console` | `api /api/v1/system/user/self/app/get-list` |\n")
		b.WriteString("| 查询【我的】菜单列表 | `ur-console` | `api /api/v1/system/user/self/menu/get-list` |\n")
		b.WriteString("| 查询当前用户信息 | `ur-console` | `api /api/v1/system/user/self/get-one` |\n")
		b.WriteString("| 查询设备列表 | `ur-iot` | `api /api/v1/things/device/info/get-list` |\n")
		b.WriteString("| 查询产品列表 | `ur-iot` | `api /api/v1/things/product/info/get-list` |\n")
		b.WriteString("| 查询项目列表 | `ur-iot` | `api /api/v1/things/project/info/get-list` |\n")
		b.WriteString("\n")
	} else {
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
	}

	// 按 group 分组（后续共用）
	groups := map[string][]swagger.Endpoint{}
	for _, ep := range endpoints {
		groups[ep.Group] = append(groups[ep.Group], ep)
	}

	// API 端点统计
	b.WriteString("## API 端点\n\n")
	if allEndpoints {
		b.WriteString(fmt.Sprintf("共 %d 个可调用端点（涵盖所有应用）。\n\n", len(endpoints)))
		b.WriteString("> **注意**：本 skill 为统一索引。上方「常见查询速查」已列出最常用 API，**优先尝试速查表中的接口**。\n")
		b.WriteString("> 如需查看更多端点，按需调用以下文件（比完整索引更小）：\n")
		b.WriteString("> - `skill_view(name=\"ur-api\", filePath=\"references/system-apis.md\")` — system 相关接口\n")
		b.WriteString("> - `skill_view(name=\"ur-api\", filePath=\"references/things-apis.md\")` — things 相关接口\n")
		b.WriteString("> - `skill_view(name=\"ur-api\", filePath=\"references/ai-apis.md\")` — AI 相关接口\n")
		b.WriteString("> 如需查看全部端点，调用 `skill_view(name=\"ur-api\", filePath=\"swagger-index.md\")`。\n\n")
		b.WriteString("### 分类索引\n\n")
		b.WriteString("| 分类 | 端点数量 |\n")
		b.WriteString("|------|---------|\n")
		for group, eps := range groups {
			b.WriteString(fmt.Sprintf("| %s | %d |\n", group, len(eps)))
		}
		b.WriteString("\n")
	} else {
		b.WriteString(fmt.Sprintf("共 %d 个可调用端点（按 %s 权限过滤）。\n\n", len(endpoints), strings.Join(app.AllowedAuthTypes(), "/")))
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
	}

	// 使用示例
	b.WriteString("## 使用示例\n\n")
	b.WriteString("```bash\n")
	if allEndpoints {
		b.WriteString("# 配置（以 iot 为例，其他应用同理）\nur-iot setup\n\n")
		b.WriteString("# 验证连通性\nur-iot check\n\n")
		b.WriteString("# 调用 API\nur-iot api /api/v1/system/user/self/get-one\n")
	} else {
		b.WriteString(fmt.Sprintf("# 配置\n%s setup\n\n", name))
		b.WriteString(fmt.Sprintf("# 验证连通性\n%s check\n\n", name))
		b.WriteString(fmt.Sprintf("# 调用 API\n%s api /api/v1/system/user/self/get-one\n", name))
	}
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
