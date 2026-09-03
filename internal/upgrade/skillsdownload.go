// skillsdownload.go — skills 独立安装包（ur-api-skills-<版本>.zip）的下载与解压。
//
// 供 `ur skills download` 使用：AI 工具自助下载后自行拷贝到自己的 skills 目录，
// CLI 不感知各 AI 工具的目录约定。复用本包已有的 release 查询（FetchLatestRelease）
// 与下载解压（downloadFile/extract）能力，不重复造轮子。
package upgrade

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// skillsAssetPrefix release 中 skills 独立资产的文件名前缀，完整名为 <前缀><版本>.zip
	skillsAssetPrefix = "ur-api-skills-"
	// GitHubReleaseBaseURL GitHub release 页面地址（下载失败时提示用户手动下载）
	GitHubReleaseBaseURL = "https://github.com/" + RepoOwner + "/" + RepoName + "/releases"
	// GiteeReleaseBaseURL Gitee release 页面地址（GitHub 不可达时的备选手动下载渠道）
	GiteeReleaseBaseURL = "https://gitee.com/" + RepoOwner + "/" + RepoName + "/releases"
)

// SkillsDownloadOptions skills 下载选项
type SkillsDownloadOptions struct {
	// OutputDir 下载与解压的目标目录（调用方应已把 ~ 展开为绝对路径）
	OutputDir string
	// ZipURL 直接指定 skills zip 下载地址（私有化/离线场景）；为空时自动查询最新 release
	ZipURL string
}

// SkillsDownloadResult skills 下载结果
type SkillsDownloadResult struct {
	// Version skills 包版本（--url 直接指定时从 zip 文件名反推，解析不出为 unknown）
	Version string `json:"version"`
	// DownloadURL 实际下载来源地址
	DownloadURL string `json:"downloadUrl"`
	// LocalPath 解压后 ur-api 目录的本地路径（可直接整体拷贝到 AI 工具的 skills 目录）
	LocalPath string `json:"localPath"`
}

// SkillsAssetName 返回指定版本的 skills 独立资产文件名（ur-api-skills-<版本>.zip），
// 与 scripts/release.sh 打包的资产名保持一致
func SkillsAssetName(version string) string {
	return skillsAssetPrefix + version + ".zip"
}

// FindSkillsAsset 在 Release 中查找 skills 独立资产（ur-api-skills-<tag>.zip）
func (r *Release) FindSkillsAsset() *Asset {
	// 精确匹配：资产名与 tag 一致
	expected := SkillsAssetName(r.TagName)
	for i := range r.Assets {
		if r.Assets[i].Name == expected {
			return &r.Assets[i]
		}
	}
	// 兜底：按前缀匹配（tag 与资产名版本存在细微差异时仍可下载）
	for i := range r.Assets {
		if strings.HasPrefix(r.Assets[i].Name, skillsAssetPrefix) && strings.HasSuffix(r.Assets[i].Name, ".zip") {
			return &r.Assets[i]
		}
	}
	return nil
}

// DownloadSkills 下载 skills 独立包并解压，返回解压后的 ur-api 目录路径。
// 流程：确定来源（--url 直接指定，或最新 release 中的 ur-api-skills-<版本>.zip 资产）→
// 下载 zip 到 OutputDir → 解压到 OutputDir → 定位含 SKILL.md 的 ur-api 目录。
// 下载或查询失败时，错误信息附带 GitHub 与 Gitee release 页面地址，便于手动下载兜底。
func DownloadSkills(opts SkillsDownloadOptions) (*SkillsDownloadResult, error) {
	downloadURL := opts.ZipURL
	version := "unknown"
	if downloadURL == "" {
		release, err := FetchLatestRelease()
		if err != nil {
			return nil, fmt.Errorf("查询最新 release 失败: %w（可访问 %s 或 %s 手动下载 skills 包，或用 --url 指定 zip 地址）",
				err, GitHubReleaseBaseURL, GiteeReleaseBaseURL)
		}
		asset := release.FindSkillsAsset()
		if asset == nil {
			return nil, fmt.Errorf("release %s 中未找到 skills 资产（%s），请访问 %s/tag/%s 或 %s/tag/%s 手动下载",
				release.TagName, SkillsAssetName(release.TagName),
				GitHubReleaseBaseURL, release.TagName, GiteeReleaseBaseURL, release.TagName)
		}
		downloadURL = asset.BrowserDownloadURL
		version = release.TagName
	}

	// zip 落盘文件名取 URL 路径最后一段（去掉 query/fragment）
	zipName := downloadURL
	if idx := strings.LastIndex(zipName, "/"); idx >= 0 {
		zipName = zipName[idx+1:]
	}
	if idx := strings.IndexAny(zipName, "?#"); idx >= 0 {
		zipName = zipName[:idx]
	}
	if zipName == "" {
		zipName = SkillsAssetName(version)
	}
	// --url 场景尝试从 zip 文件名反推版本（如 ur-api-skills-v0.4.1.zip → v0.4.1）
	if version == "unknown" && strings.HasPrefix(zipName, skillsAssetPrefix) && strings.HasSuffix(zipName, ".zip") {
		version = strings.TrimSuffix(strings.TrimPrefix(zipName, skillsAssetPrefix), ".zip")
	}

	if err := os.MkdirAll(opts.OutputDir, 0o755); err != nil {
		return nil, fmt.Errorf("创建下载目录失败: %w", err)
	}

	zipPath := filepath.Join(opts.OutputDir, zipName)
	if err := downloadFile(downloadURL, zipPath); err != nil {
		return nil, fmt.Errorf("下载 %s 失败: %w（可访问 %s 或 %s 手动下载 skills 包，或用 --url 指定 zip 地址）",
			downloadURL, err, GitHubReleaseBaseURL, GiteeReleaseBaseURL)
	}

	// 解压到 OutputDir（release 打包的 zip 内顶层为 ur-api/）
	if err := extract(zipPath, opts.OutputDir); err != nil {
		return nil, fmt.Errorf("解压 skills 包失败: %w", err)
	}

	// 定位含 SKILL.md 的 ur-api 目录（兼容 zip 无顶层目录的打包形态）
	localPath := findURAPIDir(opts.OutputDir)
	if localPath == "" {
		return nil, fmt.Errorf("解压后未找到含 SKILL.md 的 ur-api 目录，skills 包可能不完整: %s", downloadURL)
	}

	return &SkillsDownloadResult{Version: version, DownloadURL: downloadURL, LocalPath: localPath}, nil
}

// findURAPIDir 在目录中查找名为 ur-api 且含 SKILL.md 的目录；
// 先查根下（标准打包形态），未命中再全树遍历兜底
func findURAPIDir(dir string) string {
	direct := filepath.Join(dir, "ur-api")
	if fi, err := os.Stat(filepath.Join(direct, "SKILL.md")); err == nil && !fi.IsDir() {
		return direct
	}
	var found string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || found != "" {
			return nil
		}
		if info.IsDir() && info.Name() == "ur-api" {
			if fi, err := os.Stat(filepath.Join(path, "SKILL.md")); err == nil && !fi.IsDir() {
				found = path
				return filepath.SkipAll
			}
		}
		return nil
	})
	return found
}
