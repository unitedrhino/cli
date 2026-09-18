package upgrade

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	// GitHubAPI 默认的 GitHub API 地址
	GitHubAPI = "https://api.github.com"
	// RepoOwner CLI 仓库所有者
	RepoOwner = "unitedrhino"
	// RepoName CLI 仓库名
	RepoName = "cli"
)

// Release 表示 GitHub Release 信息
type Release struct {
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	PublishedAt time.Time `json:"published_at"`
	HTMLURL     string    `json:"html_url"`
	Assets      []Asset   `json:"assets"`
	// Source 是 Release 来源（gitee/github），json 不序列化；决定下载是否可走别名直链
	Source string `json:"-"`
}

// Asset 表示 Release 中的资源文件
type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Size               int64  `json:"size"`
}

// platformKey 返回当前平台的标识字符串，如 "linux-amd64"
func platformKey() string {
	return runtime.GOOS + "-" + runtime.GOARCH
}

// platformReleaseName 返回 release 包中平台对应的友好名称
// 与 release.sh 中的命名保持一致：ur-cli-${VERSION}-${name}.tar.gz
func platformReleaseName() string {
	switch platformKey() {
	case "linux-amd64":
		return "Linux-x86_64"
	case "linux-arm64":
		return "Linux-aarch64"
	case "darwin-amd64":
		return "macOS-x86_64"
	case "darwin-arm64":
		return "macOS-arm64"
	case "windows-amd64":
		return "Windows-x86_64"
	case "windows-arm64":
		return "Windows-arm64"
	default:
		return runtime.GOOS + "-" + runtime.GOARCH
	}
}

// GiteeAPI Gitee 开放 API 地址（国内直连可达，作为默认源；GitHub 不可达时兜底）
const GiteeAPI = "https://gitee.com/api/v5"

// GitHubLatestDownloadURL 版本无关永久直链前缀：releases/latest/download/<固定名>
// 永远指向最新 release 的同名资产（别名包由 scripts/release.sh 发布时上传）
const GitHubLatestDownloadURL = "https://github.com/" + RepoOwner + "/" + RepoName + "/releases/latest/download"

// SkillsAliasName 是 skills 独立包的版本无关别名资产名
const SkillsAliasName = "ur-api-skills.zip"

// aliasablePlatforms 是 release.sh 上传了版本无关别名包的平台集合；
// 仅这些平台可安全使用 latest 直链下载
var aliasablePlatforms = map[string]bool{
	"Linux-x86_64":   true,
	"Linux-aarch64":  true,
	"macOS-x86_64":   true,
	"macOS-arm64":    true,
	"Windows-x86_64": true,
}

// latestAliasDownload 返回当前平台在 GitHub latest 直链上的别名下载地址；
// 平台不在别名集合内时返回空串（调用方回退 API 资产地址）
func latestAliasDownload() (url, name string) {
	platform := platformReleaseName()
	if !aliasablePlatforms[platform] {
		return "", ""
	}
	suffix := ".tar.gz"
	if runtime.GOOS == "windows" {
		suffix = ".zip"
	}
	name = "ur-cli-" + platform + suffix
	return GitHubLatestDownloadURL + "/" + name, name
}

// FetchLatestRelease 获取最新 Release：默认优先 Gitee（国内网络可达），
// Gitee 失败时回退 GitHub。可用环境变量 UR_RELEASE_SOURCE=github 强制走 GitHub。
func FetchLatestRelease() (*Release, error) {
	if strings.TrimSpace(os.Getenv("UR_RELEASE_SOURCE")) == "github" {
		return fetchLatestReleaseFromGitHub()
	}
	release, giteeErr := fetchLatestReleaseFromGitee()
	if giteeErr == nil {
		return release, nil
	}
	release, ghErr := fetchLatestReleaseFromGitHub()
	if ghErr != nil {
		return nil, fmt.Errorf("Gitee 与 GitHub 均查询失败: Gitee: %v; GitHub: %w", giteeErr, ghErr)
	}
	return release, nil
}

// fetchLatestReleaseFromGitee 从 Gitee API 获取最新 Release（与 GitHub Release JSON 结构兼容）
func fetchLatestReleaseFromGitee() (*Release, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", GiteeAPI, RepoOwner, RepoName)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("无法连接 Gitee API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Gitee API 返回状态码 %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("解析 Gitee Release 信息失败: %w", err)
	}
	if release.TagName == "" {
		return nil, fmt.Errorf("Gitee Release 信息为空")
	}
	release.Source = "gitee"

	return &release, nil
}

// fetchLatestReleaseFromGitHub 从 GitHub API 获取最新 Release
func fetchLatestReleaseFromGitHub() (*Release, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", GitHubAPI, RepoOwner, RepoName)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("无法连接 GitHub API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API 返回状态码 %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("解析 Release 信息失败: %w", err)
	}
	release.Source = "github"

	return &release, nil
}

// FetchRelease 获取指定版本的 Release
func FetchRelease(version string) (*Release, error) {
	// 确保版本号以 v 开头
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	url := fmt.Sprintf("%s/repos/%s/%s/releases/tags/%s", GitHubAPI, RepoOwner, RepoName, version)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("无法连接 GitHub API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("版本 %s 不存在", version)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API 返回状态码 %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("解析 Release 信息失败: %w", err)
	}

	return &release, nil
}

// FindAsset 在 Release 中查找匹配当前平台的资源文件
func (r *Release) FindAsset() *Asset {
	platformName := platformReleaseName()
	suffix := ".tar.gz"
	if runtime.GOOS == "windows" {
		suffix = ".zip"
	}

	// 精确匹配优先：ur-cli-${TAG}-${PlatformName}${suffix}（release 若同时含
	// 版本无关别名包，Contains 会误命中，别名仅作直链下载用途）
	if r.TagName != "" {
		exact := "ur-cli-" + r.TagName + "-" + platformName + suffix
		for i := range r.Assets {
			if r.Assets[i].Name == exact {
				return &r.Assets[i]
			}
		}
	}

	// 匹配模式：ur-cli-${VERSION}-${PlatformName}.tar.gz
	expectedSuffix := platformName + suffix
	for i := range r.Assets {
		if strings.Contains(r.Assets[i].Name, expectedSuffix) {
			return &r.Assets[i]
		}
	}

	// 备选：按 GOOS-GOARCH 格式匹配
	fallbackSuffix := platformKey() + suffix
	for i := range r.Assets {
		if strings.Contains(r.Assets[i].Name, fallbackSuffix) {
			return &r.Assets[i]
		}
	}

	return nil
}

// IsNewer 按语义化版本判断 release 版本是否比当前版本新。
func IsNewer(current, latest string) bool {
	currentVersion, currentOK := parseVersion(current)
	latestVersion, latestOK := parseVersion(latest)
	if !currentOK || !latestOK {
		return false
	}
	return compareVersion(latestVersion, currentVersion) > 0
}

// semanticVersion 保存自动升级需要比较的语义化版本字段。
type semanticVersion struct {
	major      int
	minor      int
	patch      int
	prerelease string
}

// parseVersion 解析 vMAJOR.MINOR.PATCH[-PRERELEASE]，构建元数据不参与比较。
func parseVersion(value string) (semanticVersion, bool) {
	value = strings.TrimPrefix(strings.TrimSpace(value), "v")
	if value == "" || value == "dev" {
		return semanticVersion{}, false
	}
	value = strings.SplitN(value, "+", 2)[0]
	parts := strings.SplitN(value, "-", 2)
	numbers := strings.Split(parts[0], ".")
	if len(numbers) != 3 {
		return semanticVersion{}, false
	}
	parsed := make([]int, 3)
	for index, number := range numbers {
		value, err := strconv.Atoi(number)
		if err != nil || value < 0 {
			return semanticVersion{}, false
		}
		parsed[index] = value
	}
	result := semanticVersion{major: parsed[0], minor: parsed[1], patch: parsed[2]}
	if len(parts) == 2 {
		result.prerelease = parts[1]
	}
	return result, true
}

// compareVersion 返回 left 相对 right 的顺序，正数表示 left 更新。
func compareVersion(left, right semanticVersion) int {
	leftNumbers := []int{left.major, left.minor, left.patch}
	rightNumbers := []int{right.major, right.minor, right.patch}
	for index := range leftNumbers {
		if leftNumbers[index] > rightNumbers[index] {
			return 1
		}
		if leftNumbers[index] < rightNumbers[index] {
			return -1
		}
	}
	if left.prerelease == right.prerelease {
		return 0
	}
	if left.prerelease == "" {
		return 1
	}
	if right.prerelease == "" {
		return -1
	}
	return comparePrerelease(left.prerelease, right.prerelease)
}

// comparePrerelease 按 SemVer 规则比较点分隔的预发布标识。
func comparePrerelease(left, right string) int {
	leftParts := strings.Split(left, ".")
	rightParts := strings.Split(right, ".")
	limit := len(leftParts)
	if len(rightParts) < limit {
		limit = len(rightParts)
	}
	for index := 0; index < limit; index++ {
		leftNumber, leftErr := strconv.Atoi(leftParts[index])
		rightNumber, rightErr := strconv.Atoi(rightParts[index])
		switch {
		case leftErr == nil && rightErr == nil:
			if leftNumber > rightNumber {
				return 1
			}
			if leftNumber < rightNumber {
				return -1
			}
		case leftErr == nil:
			return -1
		case rightErr == nil:
			return 1
		default:
			if compared := strings.Compare(leftParts[index], rightParts[index]); compared != 0 {
				return compared
			}
		}
	}
	if len(leftParts) > len(rightParts) {
		return 1
	}
	if len(leftParts) < len(rightParts) {
		return -1
	}
	return 0
}
