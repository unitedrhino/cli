package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gitee.com/unitedrhino/cli/internal/skillinstall"
	"gitee.com/unitedrhino/cli/internal/upgrade"
	"gitee.com/unitedrhino/cli/internal/version"
)

func runSkills(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printSkillsHelp(stdout)
		return 0
	}

	switch args[0] {
	case "list", "ls":
		return runSkillsList(args[1:], stdout, stderr)
	case "update", "upgrade":
		return runSkillsUpdate(args[1:], stdout, stderr)
	case "install":
		return runSkillsInstall(args[1:], stdout, stderr)
	case "download":
		return runSkillsDownload(args[1:], stdout, stderr)
	case "version", "ver":
		return runSkillsVersion(args[1:], stdout, stderr)
	case "-h", "--help", "help":
		printSkillsHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "未知 skills 子命令: %s\n", args[0])
		printSkillsHelp(stderr)
		return 2
	}
}

// expandHomePath 展开路径开头表示主目录的 "~"（~、~/、~\）为当前用户主目录。
// Windows 原生 exe 不会自动展开 Git Bash 风格的 "~/xxx" 路径，直接使用会把字面 "~"
// 当成目录名写入错误位置，因此各命令的路径参数统一在此展开。
// 非 ~ 开头（含 ~user 形式，不支持展开）或无法获取主目录时原样返回。
func expandHomePath(path string) string {
	if path != "~" && !strings.HasPrefix(path, "~/") && !strings.HasPrefix(path, `~\`) {
		return path
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}
	if path == "~" {
		return home
	}
	// "~\xxx" 是 Windows 风格路径，统一按当前平台的路径分隔符处理
	rest := strings.ReplaceAll(path[2:], `\`, string(filepath.Separator))
	return filepath.Join(home, rest)
}

// runSkillsInstall 把内置 ur-api skill 整体拷贝部署到各 AI 工具的 skills 目录，
// 让对应 AI（Claude Code / Codex）重载后即可发现使用。
// 未指定 --dir 时自动探测本机常用目录；指定 --dir（可多次）时只安装到指定目录。
func runSkillsInstall(args []string, stdout, stderr io.Writer) int {
	dryRun := false
	jsonOutput := false
	customDirs := []string{}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dry-run":
			dryRun = true
		case "--json":
			jsonOutput = true
		case "--dir":
			if i+1 < len(args) {
				customDirs = append(customDirs, expandHomePath(args[i+1]))
				i++
			} else {
				fmt.Fprintln(stderr, "--dir 需要指定目标目录")
				return 2
			}
		case "-h", "--help":
			fmt.Fprintln(stdout, "用法: ur skills install [--dry-run] [--dir <目录>...] [--json]")
			fmt.Fprintln(stdout, "把内置 ur-api skill 整体部署到本机各 AI 工具（Claude Code / Codex 的用户级与项目级 skills 目录）")
			fmt.Fprintln(stdout, "不指定 --dir 时自动探测；指定 --dir（可多次）时只安装到这些目录，ur-api 会装到 <目录>/ur-api/")
			return 0
		default:
			fmt.Fprintf(stderr, "未知参数: %s\n", args[i])
			return 2
		}
	}

	src := upgrade.GetDefaultSkillsDir()

	// 目标目录：显式 --dir 优先（只装指定目录，可多次指定不同目录）；
	// 未指定时自动探测本机各 AI 工具的 skills 目录
	var targets []skillinstall.Target
	if len(customDirs) > 0 {
		for _, dir := range customDirs {
			targets = append(targets, skillinstall.Target{Path: dir, Scope: "custom", Kind: "custom"})
		}
	} else {
		detected, err := skillinstall.DetectTargets(cwdOr("."))
		if err != nil {
			fmt.Fprintf(stderr, "探测 AI skills 目录失败: %v\n", err)
			return 1
		}
		targets = detected
	}
	if len(targets) == 0 {
		fmt.Fprintln(stdout, "未检测到可部署的 AI skills 目录（~/.claude/skills、~/.agents/skills 或项目 .claude/skills/.agents/skills 均不存在；可用 --dir 指定目标目录）")
		return 0
	}

	result, err := skillinstall.Install(src, targets, dryRun)
	if err != nil {
		fmt.Fprintf(stderr, "安装失败: %v\n", err)
		return 1
	}
	if jsonOutput {
		output, _ := json.MarshalIndent(result, "", "  ")
		fmt.Fprintln(stdout, string(output))
		return 0
	}
	fmt.Fprintln(stdout, skillinstall.Summary(result))
	if dryRun {
		fmt.Fprintln(stdout, "（--dry-run 仅预览，未实际写入）")
	} else {
		fmt.Fprintln(stdout, "部署完成：重启对应 AI 工具会话后即可发现 ur-api skill")
	}
	return 0
}

// cwdOr 获取当前工作目录，失败时回退到给定默认值
func cwdOr(fallback string) string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	return fallback
}

// runSkillsDownload 从最新 release 下载 skills 独立包（ur-api-skills-<版本>.zip）并解压到本地，
// 供 AI 工具自助获取 ur-api skill 后自行拷贝到自己的 skills 目录。
// CLI 不再感知各 AI 工具的目录约定——skills 目录位置由 AI 根据自身工具确认。
// 参数说明：--output DIR 指定下载与解压目录（默认 ~/.ur/downloads/，支持 ~ 展开）；
// --url URL 直接指定 skills zip 地址（私有化/离线场景），跳过 release 查询；
// --json 以单行 JSON 事件输出结果。
func runSkillsDownload(args []string, stdout, stderr io.Writer) int {
	outputDir := ""
	zipURL := ""
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--output":
			if i+1 < len(args) {
				outputDir = expandHomePath(args[i+1])
				i++
			} else {
				fmt.Fprintln(stderr, "--output 需要指定目录")
				return 2
			}
		case "--url":
			if i+1 < len(args) {
				zipURL = args[i+1]
				i++
			} else {
				fmt.Fprintln(stderr, "--url 需要指定 zip 下载地址")
				return 2
			}
		case "--json":
			jsonOutput = true
		case "-h", "--help":
			fmt.Fprintln(stdout, "用法: ur skills download [--output <目录>] [--url <zip地址>] [--json]")
			fmt.Fprintln(stdout, "从最新 release 下载 skills 包（ur-api-skills-<版本>.zip）并解压到本地目录，")
			fmt.Fprintln(stdout, "之后把其中的 ur-api 目录整体拷贝到你所用 AI 工具的 skills 目录即可（对所有 AI 工具通用）")
			fmt.Fprintln(stdout, "--url 可直接指定 zip 地址（私有化/离线场景）")
			return 0
		default:
			fmt.Fprintf(stderr, "未知参数: %s\n", args[i])
			return 2
		}
	}

	if outputDir == "" {
		outputDir = expandHomePath("~/.ur/downloads")
	}

	result, err := upgrade.DownloadSkills(upgrade.SkillsDownloadOptions{OutputDir: outputDir, ZipURL: zipURL})
	if err != nil {
		fmt.Fprintf(stderr, "下载 skills 失败: %v\n", err)
		return 1
	}
	fmt.Fprintln(stdout, formatSkillsDownloadResult(result, jsonOutput))
	return 0
}

// skillsDownloadHint 下载成功后的拷贝指引：skills 目录位置由各 AI 工具自行确认，
// CLI 不再硬编码各工具目录（对所有 AI 工具通用）
const skillsDownloadHint = "请将上述 ur-api 目录整体拷贝到你所用 AI 工具的 skills 目录下（各 AI 工具的 skills 目录由 AI 自行确认，例如 Claude Code 为 ~/.claude/skills/），拷贝后重启 AI 工具生效"

// formatSkillsDownloadResult 构造 download 成功输出：
// JSON 模式为单行事件 {"event":"skills_downloaded",...}，否则为人类可读文本
func formatSkillsDownloadResult(res *upgrade.SkillsDownloadResult, jsonMode bool) string {
	if jsonMode {
		return fmt.Sprintf(`{"event":"skills_downloaded","downloadUrl":%q,"localPath":%q,"installHint":%q}`,
			res.DownloadURL, res.LocalPath, skillsDownloadHint)
	}
	return fmt.Sprintf("下载来源: %s\n本地解压路径: %s\n%s", res.DownloadURL, res.LocalPath, skillsDownloadHint)
}

func runSkillsList(args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for _, arg := range args {
		if arg == "--json" {
			jsonOutput = true
		}
	}

	skillsDir := upgrade.GetDefaultSkillsDir()
	result, err := upgrade.ListSkills(skillsDir)
	if err != nil {
		fmt.Fprintf(stderr, "列出 skills 失败: %v\n", err)
		return 1
	}

	if result.Skills == nil {
		fmt.Fprintln(stdout, "未安装 skills。请下载 CLI 完整安装包获取 skills。")
		return 0
	}

	if jsonOutput {
		output, _ := json.MarshalIndent(result, "", "  ")
		fmt.Fprintln(stdout, string(output))
		return 0
	}

	fmt.Fprintf(stdout, "Skills 版本: %s\n", result.Version)
	if result.Skills == nil {
		fmt.Fprintln(stdout, "  (未安装 skills)")
		return 0
	}

	fmt.Fprintf(stdout, "已安装 %d 个 skill:\n\n", len(result.Skills))
	for _, s := range result.Skills {
		if s.Description != "" {
			fmt.Fprintf(stdout, "  %-26s %-12s %s\n", s.Name, s.Version, s.Description)
		} else {
			fmt.Fprintf(stdout, "  %-26s %s\n", s.Name, s.Version)
		}
	}
	return 0
}

func runSkillsUpdate(args []string, stdout, stderr io.Writer) int {
	dryRun := false
	jsonOutput := false
	for _, arg := range args {
		switch arg {
		case "--dry-run":
			dryRun = true
		case "--json":
			jsonOutput = true
		}
	}

	if version.IsDev() {
		fmt.Fprintln(stdout, "当前为开发版本 (dev)，skills 随源码更新，无需单独升级。")
		return 0
	}

	skillsDir := upgrade.GetDefaultSkillsDir()
	if skillsDir == "" {
		fmt.Fprintln(stderr, "未找到 skills 目录，请先安装 CLI 完整包。")
		return 1
	}

	result, err := upgrade.UpdateSkills(skillsDir, dryRun)
	if err != nil {
		if jsonOutput {
			output, _ := json.MarshalIndent(result, "", "  ")
			fmt.Fprintln(stdout, string(output))
		} else {
			fmt.Fprintf(stderr, "升级 skills 失败: %v\n", err)
		}
		return 1
	}

	if jsonOutput {
		output, _ := json.MarshalIndent(result, "", "  ")
		fmt.Fprintln(stdout, string(output))
		return 0
	}

	if result.UpToDate {
		fmt.Fprintf(stdout, "Skills 已是最新版本 %s\n", result.CurrentVersion)
		return 0
	}

	if dryRun {
		fmt.Fprintf(stdout, "Skills 有新版本可用: %s → %s\n", result.CurrentVersion, result.LatestVersion)
		fmt.Fprintln(stdout, "运行 ur skills update 执行升级")
		return 0
	}

	fmt.Fprintf(stdout, "Skills 升级成功: %s → %s\n", result.CurrentVersion, result.LatestVersion)
	fmt.Fprintf(stdout, "已更新 %d 个 skill\n", result.UpdatedCount)
	return 0
}

func runSkillsVersion(args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	for _, arg := range args {
		if arg == "--json" {
			jsonOutput = true
		}
	}

	skillsDir := upgrade.GetDefaultSkillsDir()
	skillsVersion := upgrade.GetSkillsVersion(skillsDir)
	lastUpdated := upgrade.SkillsLastUpdated(skillsDir)

	if jsonOutput {
		output := fmt.Sprintf(`{"version": %q, "updatedAt": %q}`, skillsVersion, lastUpdated)
		fmt.Fprintln(stdout, output)
		return 0
	}

	if skillsDir == "" {
		fmt.Fprintln(stdout, "未找到 skills 目录")
		return 0
	}

	parts := []string{fmt.Sprintf("Skills 版本: %s", skillsVersion)}
	if lastUpdated != "" {
		parts = append(parts, fmt.Sprintf("最后更新: %s", lastUpdated))
	}
	parts = append(parts, fmt.Sprintf("路径: %s", skillsDir))
	fmt.Fprintln(stdout, strings.Join(parts, "\n"))
	return 0
}

func printSkillsHelp(w io.Writer) {
	fmt.Fprintln(w, "用法: ur skills <子命令> [选项]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "子命令:")
	fmt.Fprintln(w, "  list, ls        列出已安装的 skills")
	fmt.Fprintln(w, "  download        从最新 release 下载 skills 包并解压到本地（AI 自助获取后自行拷贝）")
	fmt.Fprintln(w, "  update, upgrade  升级 skills 到最新版本")
	fmt.Fprintln(w, "  install         把内置 ur-api skill 部署到本机各 AI 工具的 skills 目录")
	fmt.Fprintln(w, "  version, ver     查看 skills 版本信息")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "选项:")
	fmt.Fprintln(w, "  --json           以 JSON 格式输出")
	fmt.Fprintln(w, "  --dry-run        只检查更新，不安装（仅 update）")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "示例:")
	fmt.Fprintln(w, "  ur skills list              列出所有已安装的 skills")
	fmt.Fprintln(w, "  ur skills download          下载最新 skills 包到 ~/.ur/downloads/ 并解压")
	fmt.Fprintln(w, "  ur skills download --output ~/skills-pkg   指定下载目录")
	fmt.Fprintln(w, "  ur skills update --dry-run  检查 skills 是否有更新")
	fmt.Fprintln(w, "  ur skills update            升级 skills 到最新版本")
	fmt.Fprintln(w, "  ur skills install           部署 ur-api 到本机各 AI 工具")
	fmt.Fprintln(w, "  ur skills version           查看 skills 版本")
}
