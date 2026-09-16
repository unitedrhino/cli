package shared

import (
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
	case "export":
		return runSkillsExport(args[1:], stdout, stderr)
	case "status", "doctor":
		return runSkillsStatus(args[1:], stdout, stderr)
	case "target", "targets":
		return runSkillsTarget(args[1:], stdout, stderr)
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

// runSkillsInstall 把内置 ur-api skill 整体部署到自动发现和用户登记的目标。
// 未指定 --dir 时合并已登记目标与本机客户端目录；指定 --dir 时只安装到显式目录。
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
		case "--all":
			// 默认已经覆盖所有自动发现与登记目标；保留 --all 作为明确表达和脚本兼容入口。
		case "--dir":
			if i+1 < len(args) {
				customDirs = append(customDirs, expandHomePath(args[i+1]))
				i++
			} else {
				fmt.Fprintln(stderr, "--dir 需要指定目标目录")
				return 2
			}
		case "-h", "--help":
			fmt.Fprintln(stdout, "用法: ur skills install [--all] [--dry-run] [--dir <目录>...] [--json]")
			fmt.Fprintln(stdout, "把内置 ur-api 整体部署到自动发现和用户登记的 AI Skills 目录")
			fmt.Fprintln(stdout, "不指定 --dir 时安装到全部目标；指定 --dir（可多次）时只安装到这些目录")
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
		for index, dir := range customDirs {
			targets = append(targets, skillinstall.Target{Name: fmt.Sprintf("command-line-%d", index+1), Path: dir, Scope: "custom", Kind: "custom", Origin: "command-line"})
		}
	} else {
		detected, err := skillinstall.ResolveTargets(cwdOr("."))
		if err != nil {
			fmt.Fprintf(stderr, "探测 AI skills 目录失败: %v\n", err)
			return 1
		}
		targets = detected
	}
	if len(targets) == 0 {
		fmt.Fprintln(stdout, "未发现可部署目标；可用 ur skills target add 登记目录，或用 --dir 临时指定")
		return 0
	}

	result, err := skillinstall.Install(src, targets, dryRun)
	if err != nil {
		fmt.Fprintf(stderr, "安装失败: %v\n", err)
		return 1
	}
	if jsonOutput {
		_ = writeJSON(stdout, result)
		if skillinstall.HasErrors(result) {
			return 1
		}
		return 0
	}
	fmt.Fprintln(stdout, skillinstall.Summary(result))
	if skillinstall.HasErrors(result) {
		return 1
	}
	if dryRun {
		fmt.Fprintln(stdout, "（--dry-run 仅预览，未实际写入）")
	} else {
		fmt.Fprintln(stdout, "部署完成：重启对应 AI 工具会话后即可发现 ur-api skill")
	}
	return 0
}

// runSkillsTarget 管理任意 AI 客户端的可复用 Skills 安装目标。
func runSkillsTarget(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		printSkillsTargetHelp(stdout)
		return 0
	}
	switch args[0] {
	case "detect":
		return runSkillsTargetDetect(args[1:], stdout, stderr)
	case "list", "ls":
		return runSkillsTargetList(args[1:], stdout, stderr)
	case "add", "set":
		return runSkillsTargetAdd(args[1:], stdout, stderr)
	case "remove", "rm", "delete":
		return runSkillsTargetRemove(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "未知 target 子命令: %s\n", args[0])
		printSkillsTargetHelp(stderr)
		return 2
	}
}

// runSkillsTargetDetect 输出当前环境自动识别到的客户端目录。
func runSkillsTargetDetect(args []string, stdout, stderr io.Writer) int {
	jsonOutput, ok := parseJSONOnlyArgs(args, stderr)
	if !ok {
		return 2
	}
	targets, err := skillinstall.DetectTargets(cwdOr("."))
	if err != nil {
		fmt.Fprintf(stderr, "探测 Skills 目标失败: %v\n", err)
		return 1
	}
	return printSkillsTargets(targets, jsonOutput, stdout)
}

// runSkillsTargetList 输出登记目标和自动发现目标的合并结果。
func runSkillsTargetList(args []string, stdout, stderr io.Writer) int {
	jsonOutput, ok := parseJSONOnlyArgs(args, stderr)
	if !ok {
		return 2
	}
	targets, err := skillinstall.ResolveTargets(cwdOr("."))
	if err != nil {
		fmt.Fprintf(stderr, "读取 Skills 目标失败: %v\n", err)
		return 1
	}
	return printSkillsTargets(targets, jsonOutput, stdout)
}

// runSkillsTargetAdd 新增或更新一个文件系统目标。
func runSkillsTargetAdd(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: ur skills target add <名称> --dir <目录> [--json]")
		return 2
	}
	name := args[0]
	directory := ""
	targetType := skillinstall.TargetTypeFilesystem
	jsonOutput := false
	for index := 1; index < len(args); index++ {
		switch args[index] {
		case "--dir":
			if index+1 >= len(args) {
				fmt.Fprintln(stderr, "--dir 需要指定目录")
				return 2
			}
			directory = expandHomePath(args[index+1])
			index++
		case "--type":
			if index+1 >= len(args) {
				fmt.Fprintln(stderr, "--type 需要指定类型")
				return 2
			}
			targetType = args[index+1]
			index++
		case "--json":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "未知参数: %s\n", args[index])
			return 2
		}
	}
	updated, err := skillinstall.UpsertConfiguredTarget(skillinstall.ConfiguredTarget{Name: name, Type: targetType, Path: directory, Enabled: true})
	if err != nil {
		fmt.Fprintf(stderr, "登记 Skills 目标失败: %v\n", err)
		return 1
	}
	result := map[string]any{"name": name, "type": targetType, "path": directory, "updated": updated}
	if jsonOutput {
		_ = writeJSON(stdout, result)
	} else if updated {
		fmt.Fprintf(stdout, "已更新 Skills 目标 %s → %s\n", name, directory)
	} else {
		fmt.Fprintf(stdout, "已登记 Skills 目标 %s → %s\n", name, directory)
	}
	return 0
}

// runSkillsTargetRemove 删除一个用户登记的目标。
func runSkillsTargetRemove(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: ur skills target remove <名称> [--json]")
		return 2
	}
	name := args[0]
	jsonOutput, ok := parseJSONOnlyArgs(args[1:], stderr)
	if !ok {
		return 2
	}
	removed, err := skillinstall.RemoveConfiguredTarget(name)
	if err != nil {
		fmt.Fprintf(stderr, "删除 Skills 目标失败: %v\n", err)
		return 1
	}
	if jsonOutput {
		_ = writeJSON(stdout, map[string]any{"name": name, "removed": removed})
	} else if removed {
		fmt.Fprintf(stdout, "已删除 Skills 目标 %s\n", name)
	} else {
		fmt.Fprintf(stdout, "Skills 目标 %s 不存在\n", name)
	}
	return 0
}

// parseJSONOnlyArgs 解析仅支持 --json 的简单子命令参数。
func parseJSONOnlyArgs(args []string, stderr io.Writer) (bool, bool) {
	jsonOutput := false
	for _, arg := range args {
		if arg == "--json" {
			jsonOutput = true
			continue
		}
		fmt.Fprintf(stderr, "未知参数: %s\n", arg)
		return false, false
	}
	return jsonOutput, true
}

// printSkillsTargets 以 JSON 或表格文本输出目标列表。
func printSkillsTargets(targets []skillinstall.Target, jsonOutput bool, stdout io.Writer) int {
	if jsonOutput {
		_ = writeJSON(stdout, map[string]any{"targets": targets})
		return 0
	}
	if len(targets) == 0 {
		fmt.Fprintln(stdout, "未发现 Skills 目标")
		return 0
	}
	for _, target := range targets {
		fmt.Fprintf(stdout, "%-22s %-10s %-10s %s\n", target.Name, target.Kind, target.Origin, target.Path)
	}
	return 0
}

// printSkillsTargetHelp 输出目标管理命令说明。
func printSkillsTargetHelp(w io.Writer) {
	fmt.Fprintln(w, "用法: ur skills target <子命令>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "子命令:")
	fmt.Fprintln(w, "  detect                    自动探测本机支持的 Skills 目录")
	fmt.Fprintln(w, "  list                      列出登记和自动发现的全部目标")
	fmt.Fprintln(w, "  add <名称> --dir <目录>   登记或更新任意客户端目录")
	fmt.Fprintln(w, "  remove <名称>             删除登记目标")
}

// runSkillsStatus 检查所有目标的版本和文件完整性。
func runSkillsStatus(args []string, stdout, stderr io.Writer) int {
	jsonOutput, ok := parseJSONOnlyArgs(args, stderr)
	if !ok {
		return 2
	}
	targets, err := skillinstall.ResolveTargets(cwdOr("."))
	if err != nil {
		fmt.Fprintf(stderr, "读取 Skills 目标失败: %v\n", err)
		return 1
	}
	source := upgrade.GetDefaultSkillsDir()
	result, err := skillinstall.InspectTargets(source, targets)
	if err != nil {
		fmt.Fprintf(stderr, "检查 Skills 状态失败: %v\n", err)
		return 1
	}
	skillinstall.SortTargetStatuses(result.Targets)
	if jsonOutput {
		_ = writeJSON(stdout, result)
	} else {
		fmt.Fprintf(stdout, "内置 Skills: %s (%s)\n", result.Version, result.Source)
		if len(result.Targets) == 0 {
			fmt.Fprintln(stdout, "未发现 Skills 目标")
		}
		for _, target := range result.Targets {
			fmt.Fprintf(stdout, "%-22s %-10s %-10s %s", target.Name, target.Kind, target.State, target.Path)
			if target.Version != "" {
				fmt.Fprintf(stdout, "  版本=%s", target.Version)
			}
			if target.MissingFiles+target.ChangedFiles+target.ExtraFiles > 0 {
				fmt.Fprintf(stdout, "  缺失=%d 变化=%d 多余=%d", target.MissingFiles, target.ChangedFiles, target.ExtraFiles)
			}
			fmt.Fprintln(stdout)
		}
	}
	for _, target := range result.Targets {
		if target.State != skillinstall.StatusCurrent || target.Error != "" {
			return 1
		}
	}
	return 0
}

// runSkillsExport 导出可由技能市场或没有固定目录的客户端导入的标准 ZIP。
func runSkillsExport(args []string, stdout, stderr io.Writer) int {
	format := "zip"
	outputPath := ""
	jsonOutput := false
	for index := 0; index < len(args); index++ {
		switch args[index] {
		case "--format":
			if index+1 >= len(args) {
				fmt.Fprintln(stderr, "--format 需要指定格式")
				return 2
			}
			format = args[index+1]
			index++
		case "--output", "-o":
			if index+1 >= len(args) {
				fmt.Fprintln(stderr, "--output 需要指定路径")
				return 2
			}
			outputPath = expandHomePath(args[index+1])
			index++
		case "--json":
			jsonOutput = true
		case "-h", "--help":
			fmt.Fprintln(stdout, "用法: ur skills export [--format zip] [--output <文件或目录>] [--json]")
			return 0
		default:
			fmt.Fprintf(stderr, "未知参数: %s\n", args[index])
			return 2
		}
	}
	if format != "zip" {
		fmt.Fprintf(stderr, "不支持的导出格式 %q，当前仅支持 zip\n", format)
		return 2
	}
	result, err := skillinstall.ExportZIP(upgrade.GetDefaultSkillsDir(), outputPath)
	if err != nil {
		fmt.Fprintf(stderr, "导出 Skills 失败: %v\n", err)
		return 1
	}
	if jsonOutput {
		_ = writeJSON(stdout, result)
	} else {
		fmt.Fprintf(stdout, "Skills ZIP 已导出: %s\n版本: %s\n文件: %d\n大小: %d 字节\n", result.Path, result.Version, result.Files, result.Bytes)
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
		_ = writeJSON(stdout, result)
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
			_ = writeJSON(stdout, result)
		} else {
			fmt.Fprintf(stderr, "升级 skills 失败: %v\n", err)
		}
		return 1
	}

	if jsonOutput {
		_ = writeJSON(stdout, result)
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
		_ = writeJSON(stdout, map[string]any{"version": skillsVersion, "updatedAt": lastUpdated})
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
	fmt.Fprintln(w, "  export          导出标准 ZIP，供扣子等平台导入")
	fmt.Fprintln(w, "  target          管理任意 AI 客户端的 Skills 安装目录")
	fmt.Fprintln(w, "  status, doctor  检查各目标的版本和文件完整性")
	fmt.Fprintln(w, "  update, upgrade  升级 skills 到最新版本")
	fmt.Fprintln(w, "  install         把内置 ur-api skill 部署到本机各 AI 工具的 skills 目录")
	fmt.Fprintln(w, "  version, ver     查看 skills 版本信息")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "选项:")
	fmt.Fprintln(w, "  --json           以 JSON 格式输出")
	fmt.Fprintln(w, "  --dry-run        预览更新或安装，不写入文件")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "示例:")
	fmt.Fprintln(w, "  ur skills list              列出所有已安装的 skills")
	fmt.Fprintln(w, "  ur skills download          下载最新 skills 包到 ~/.ur/downloads/ 并解压")
	fmt.Fprintln(w, "  ur skills download --output ~/skills-pkg   指定下载目录")
	fmt.Fprintln(w, "  ur skills update --dry-run  检查 skills 是否有更新")
	fmt.Fprintln(w, "  ur skills update            升级 skills 到最新版本")
	fmt.Fprintln(w, "  ur skills install           部署 ur-api 到本机各 AI 工具")
	fmt.Fprintln(w, "  ur skills target add workbuddy --dir ~/.codebuddy/skills")
	fmt.Fprintln(w, "  ur skills status            检查所有安装目标")
	fmt.Fprintln(w, "  ur skills export            导出标准技能 ZIP")
	fmt.Fprintln(w, "  ur skills version           查看 skills 版本")
}
