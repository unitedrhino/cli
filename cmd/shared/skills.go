package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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
	case "view":
		return runSkillsView(args[1:], stdout, stderr)
	case "update", "upgrade":
		return runSkillsUpdate(args[1:], stdout, stderr)
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

func runSkillsView(args []string, stdout, stderr io.Writer) int {
	var filePath string
	var code string

	for i := 0; i < len(args); i++ {
		if args[i] == "--file" && i+1 < len(args) {
			filePath = args[i+1]
			i++
		} else if code == "" && !strings.HasPrefix(args[i], "-") {
			code = args[i]
		}
	}

	if code == "" {
		fmt.Fprintln(stderr, "用法: ur skills view <code> [--file <path>]")
		return 2
	}

	skillDir := findSkillDirByCode(code)
	if skillDir == "" {
		fmt.Fprintf(stderr, "未找到 skill: %s\n", code)
		return 1
	}

	targetFile := filepath.Join(skillDir, "SKILL.md")
	if filePath != "" {
		targetFile = filepath.Join(skillDir, filePath)
	}

	data, err := os.ReadFile(targetFile)
	if err != nil {
		fmt.Fprintf(stderr, "读取文件失败: %v\n", err)
		return 1
	}

	fmt.Fprintln(stdout, string(data))
	return 0
}

// findSkillDirByCode 按 skill 编码查找其目录路径
func findSkillDirByCode(code string) string {
	searchDirs := []string{
		upgrade.GetDefaultSkillsDir(),
		"/opt/skills-store",
	}

	if binaryPath, err := os.Executable(); err == nil {
		searchDirs = append(searchDirs, filepath.Join(filepath.Dir(binaryPath), "skill"))
	}

	if home, err := os.UserHomeDir(); err == nil {
		searchDirs = append(searchDirs, filepath.Join(home, ".ur", "skills"))
	}

	for _, dir := range searchDirs {
		if dir == "" {
			continue
		}

		directPath := filepath.Join(dir, code)
		if fi, err := os.Stat(directPath); err == nil && fi.IsDir() {
			if _, err := os.Stat(filepath.Join(directPath, "SKILL.md")); err == nil {
				return directPath
			}
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			nestedPath := filepath.Join(dir, entry.Name(), code)
			if fi, err := os.Stat(nestedPath); err == nil && fi.IsDir() {
				if _, err := os.Stat(filepath.Join(nestedPath, "SKILL.md")); err == nil {
					return nestedPath
				}
			}
		}
	}

	return ""
}

func printSkillsHelp(w io.Writer) {
	fmt.Fprintln(w, "用法: ur skills <子命令> [选项]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "子命令:")
	fmt.Fprintln(w, "  list, ls        列出已安装的 skills")
	fmt.Fprintln(w, "  view            查看指定 skill 的内容")
	fmt.Fprintln(w, "  update, upgrade  升级 skills 到最新版本")
	fmt.Fprintln(w, "  version, ver     查看 skills 版本信息")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "选项:")
	fmt.Fprintln(w, "  --json           以 JSON 格式输出")
	fmt.Fprintln(w, "  --dry-run        只检查更新，不安装（仅 update）")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "示例:")
	fmt.Fprintln(w, "  ur skills list                      列出所有已安装的 skills")
	fmt.Fprintln(w, "  ur skills view ur-api               查看 ur-api skill 的内容")
	fmt.Fprintln(w, "  ur skills view ur-api --file swagger-index.md  查看指定文件")
	fmt.Fprintln(w, "  ur skills update --dry-run          检查 skills 是否有更新")
	fmt.Fprintln(w, "  ur skills update                    升级 skills 到最新版本")
	fmt.Fprintln(w, "  ur skills version                   查看 skills 版本")
}
