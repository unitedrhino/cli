package shared

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
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

	// 当 allEndpoints=true 时，按多级目录结构拆分为引用文件：
	// Tier 2: 领域索引文件（只含 group 名称和数量）
	// Tier 3: group 详细文件（每个 group 的端点列表）
	if allEndpoints {
		refDir := filepath.Join(outputDir, "references")
		if err := os.MkdirAll(refDir, 0o755); err != nil {
			fmt.Fprintln(stderr, err.Error())
			return 1
		}
		groupsDir := filepath.Join(refDir, "groups")
		if err := os.MkdirAll(groupsDir, 0o755); err != nil {
			fmt.Fprintln(stderr, err.Error())
			return 1
		}

		// 按路径前缀和 group 分组
		prefixes := map[string]string{
			"system-index.md": "/api/v1/system/",
			"things-index.md": "/api/v1/things/",
			"ai-index.md":     "/api/v1/ai/",
		}

		// 先全局按 group 分组（用于生成 group 详细文件）
		allGroups := map[string][]swagger.Endpoint{}
		for _, item := range filtered {
			allGroups[item.Group] = append(allGroups[item.Group], item)
		}

		// 生成 Tier 3: 每个 group 一个详细文件
		groupFileCount := 0
		for group, eps := range allGroups {
			if len(eps) == 0 {
				continue
			}
			groupFile := strings.ReplaceAll(group, "/", "-") + ".md"
			if groupFile == "" || groupFile == "-.md" {
				groupFile = "uncategorized.md"
			}
			var gb strings.Builder
			gb.WriteString(fmt.Sprintf("# %s\n\n", group))
			gb.WriteString(fmt.Sprintf("> 该 group 共 %d 个端点。\n\n", len(eps)))
			for _, ep := range eps {
				summary := ep.Summary
				if summary == "" {
					summary = ep.Description
				}
				gb.WriteString(fmt.Sprintf("- `%s %s` [%s] %s\n", ep.Method, ep.Path, ep.AuthType, summary))
			}
			groupPath := filepath.Join(groupsDir, groupFile)
			if err := os.WriteFile(groupPath, []byte(gb.String()), 0o644); err != nil {
				fmt.Fprintln(stderr, err.Error())
				return 1
			}
			groupFileCount++
		}
		fmt.Fprintf(stdout, "generated %d group files in %s\n", groupFileCount, groupsDir)

		// 生成 Tier 2: 领域索引文件（只含 group 名称 + 数量 + 指引）
		for filename, prefix := range prefixes {
			groups := map[string][]swagger.Endpoint{}
			for _, item := range filtered {
				if !strings.HasPrefix(item.Path, prefix) {
					continue
				}
				groups[item.Group] = append(groups[item.Group], item)
			}
			if len(groups) == 0 {
				continue
			}

			var groupNames []string
			for g := range groups {
				groupNames = append(groupNames, g)
			}
			sort.Strings(groupNames)

			var idxBuilder strings.Builder
			domain := strings.TrimSuffix(filename, "-index.md")
			idxBuilder.WriteString(fmt.Sprintf("# %s 领域索引\n\n", domain))
			idxBuilder.WriteString(fmt.Sprintf("> 路径前缀 `%s` 下的所有 group。如需查看某个 group 的详细端点，调用 `skill_view(name=\"ur-api\", filePath=\"references/groups/{group文件名}.md\")`。\n\n", prefix))
			idxBuilder.WriteString("| Group | 端点数量 | 对应文件 |\n")
			idxBuilder.WriteString("|-------|---------|---------|\n")
			for _, group := range groupNames {
				eps := groups[group]
				groupFile := strings.ReplaceAll(group, "/", "-") + ".md"
				if groupFile == "" || groupFile == "-.md" {
					groupFile = "uncategorized.md"
				}
				idxBuilder.WriteString(fmt.Sprintf("| `%s` | %d | `references/groups/%s` |\n", group, len(eps), groupFile))
			}
			idxBuilder.WriteString("\n")

			idxPath := filepath.Join(refDir, filename)
			if err := os.WriteFile(idxPath, []byte(idxBuilder.String()), 0o644); err != nil {
				fmt.Fprintln(stderr, err.Error())
				return 1
			}
			fmt.Fprintf(stdout, "generated %s (%d groups)\n", idxPath, len(groups))
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
		b.WriteString("以下是最常用的查询场景及对应 API。**优先尝试这些接口，无需查阅子 skill**。\n\n")
		b.WriteString("| 查询场景 | 推荐 CLI | API 命令示例 | 权限 |\n")
		b.WriteString("|---------|---------|-------------|------|\n")
		b.WriteString("| 查询【我的】应用列表 | `ur-console` | `api /api/v1/system/user/self/app/get-list` | all（任何登录用户） |\n")
		b.WriteString("| 查询【我的】菜单列表 | `ur-console` | `api /api/v1/system/user/self/menu/get-list` | all（任何登录用户） |\n")
		b.WriteString("| 查询当前用户信息 | `ur-console` | `api /api/v1/system/user/self/get-one` | all（任何登录用户） |\n")
		b.WriteString("| 查询设备列表 | `ur-iot` | `api /api/v1/things/device/info/get-list` | admin/tenant |\n")
		b.WriteString("| 查询产品列表 | `ur-iot` | `api /api/v1/things/product/info/get-list` | admin/tenant |\n")
		b.WriteString("| 查询项目列表 | `ur-iot` | `api /api/v1/things/project/info/get-list` | admin/tenant |\n")
		b.WriteString("\n")
		b.WriteString("> **空结果说明**：如果 `user/self/app/get-list` 返回空列表（`list: []`），表示**当前用户没有任何应用权限**。\n")
		b.WriteString("> 此时应直接告知用户『您当前没有分配任何应用』，**不要**再去 `system/app/info/get-list`（platform 权限）查找。\n")
		b.WriteString("> `system/app/info/get-list` 是平台管理员专属接口，普通用户调用会返回权限不足。\n\n")
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
		b.WriteString("> 如需查看更多端点，按以下三级渐进式查阅（文档越小越优先）：\n")
		b.WriteString("> 1. 查阅领域索引（只含 group 名称和数量，约 2KB）：\n")
		b.WriteString(">    - `skill_view(name=\"ur-api\", filePath=\"references/system-index.md\")` — system 领域\n")
		b.WriteString(">    - `skill_view(name=\"ur-api\", filePath=\"references/things-index.md\")` — things 领域\n")
		b.WriteString(">    - `skill_view(name=\"ur-api\", filePath=\"references/ai-index.md\")` — AI 领域\n")
		b.WriteString("> 2. 从索引中找到目标 group，查阅 group 详细文件（约 0.5-3KB）：\n")
		b.WriteString(">    - `skill_view(name=\"ur-api\", filePath=\"references/groups/{group文件名}.md\")`\n")
		b.WriteString("> 3. 如需查看全部端点，调用 `skill_view(name=\"ur-api\", filePath=\"swagger-index.md\")`。\n\n")
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
