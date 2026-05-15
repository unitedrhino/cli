package version

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Info 统一的版本信息结构
type Info struct {
	CLI      CLIVersion  `json:"cli"`
	Skills   SkillsInfo  `json:"skills,omitempty"`
	Latest   *LatestInfo `json:"latest,omitempty"`
}

// CLIVersion CLI 自身的版本信息
type CLIVersion struct {
	Version   string `json:"version"`
	Commit    string `json:"commit,omitempty"`
	BuildDate string `json:"buildDate,omitempty"`
	Platform  string `json:"platform"`
}

// SkillsInfo 本地 skills 的版本信息
type SkillsInfo struct {
	Version   string `json:"version"`
	Count     int    `json:"count"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// LatestInfo 远端最新版本信息
type LatestInfo struct {
	Version     string `json:"version"`
	PublishedAt string `json:"publishedAt"`
	URL         string `json:"url"`
}

// 构建时通过 ldflags 注入
var (
	BuildVersion = "dev"
	BuildCommit  = ""
	BuildDate    = ""
)

// GetCLIVersion 返回 CLI 版本信息
func GetCLIVersion() CLIVersion {
	return CLIVersion{
		Version:   BuildVersion,
		Commit:    BuildCommit,
		BuildDate: BuildDate,
		Platform:  runtime.GOOS + "-" + runtime.GOARCH,
	}
}

// GetSkillsInfo 扫描本地 skill 目录获取版本信息
func GetSkillsInfo(binaryPath string) SkillsInfo {
	skillsDir := findSkillsDir(binaryPath)
	entries, err := os.ReadDir(skillsDir)
	if err != nil {
		return SkillsInfo{Version: "unknown"}
	}

	count := 0
	version := "unknown"
	var latestMod time.Time

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		// 检查是否有 SKILL.md（是有效的 skill 目录）
		skillFile := filepath.Join(skillsDir, entry.Name(), "SKILL.md")
		if fi, err := os.Stat(skillFile); err == nil {
			count++
			if fi.ModTime().After(latestMod) {
				latestMod = fi.ModTime()
			}
		}
		// 读取 _meta.json 获取版本
		metaFile := filepath.Join(skillsDir, entry.Name(), "_meta.json")
		if data, err := os.ReadFile(metaFile); err == nil {
			var meta struct {
				Version string `json:"version"`
			}
			if json.Unmarshal(data, &meta) == nil && meta.Version != "" {
				version = meta.Version
			}
		}
	}

	info := SkillsInfo{
		Version: version,
		Count:   count,
	}
	if !latestMod.IsZero() {
		info.UpdatedAt = latestMod.Format(time.RFC3339)
	}
	return info
}

// IsDev 判断是否为开发版本
func IsDev() bool {
	return BuildVersion == "dev" || BuildVersion == ""
}

// findSkillsDir 查找 skills 目录（二进制同目录下的 skill/ 或 ~/.ur/skills/）
func findSkillsDir(binaryPath string) string {
	// 首选：二进制同目录下的 skill/
	if binaryPath != "" {
		dir := filepath.Join(filepath.Dir(binaryPath), "skill")
		if fi, err := os.Stat(dir); err == nil && fi.IsDir() {
			return dir
		}
	}
	// 备选：~/.ur/skills/
	home, err := os.UserHomeDir()
	if err == nil {
		dir := filepath.Join(home, ".ur", "skills")
		if fi, err := os.Stat(dir); err == nil && fi.IsDir() {
			return dir
		}
	}
	return ""
}

// FormatVersion 返回简洁版本字符串
func FormatVersion() string {
	return BuildVersion
}

// FormatVersionFull 返回完整版本信息字符串
func FormatVersionFull() string {
	var sb strings.Builder
	sb.WriteString(BuildVersion)
	if BuildCommit != "" {
		sb.WriteString(" (")
		sb.WriteString(BuildCommit)
		sb.WriteString(")")
	}
	sb.WriteString(" ")
	sb.WriteString(runtime.GOOS)
	sb.WriteString("/")
	sb.WriteString(runtime.GOARCH)
	if BuildDate != "" {
		sb.WriteString(" built ")
		sb.WriteString(BuildDate)
	}
	return sb.String()
}

// FormatVersionJSON 输出 JSON 格式的版本信息
func FormatVersionJSON(binaryPath string) string {
	info := Info{
		CLI:    GetCLIVersion(),
		Skills: GetSkillsInfo(binaryPath),
	}
	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": %q}`, err.Error())
	}
	return string(data)
}
