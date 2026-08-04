// updatecheck — CLI 自动版本检查
//
// 每次运行命令时低频（默认 24 小时一次）检查 GitHub 最新 Release，
// 在 stderr 输出升级提醒；"版本过老"（major 不同或 minor 落后 ≥2）给更强提示。
// 检查结果缓存到 ~/.ur/update-state.json，网络失败静默，不阻塞命令执行。
package updatecheck

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/upgrade"
	"gitee.com/unitedrhino/cli/internal/version"
)

// CheckInterval 两次自动检查之间的最小间隔
const CheckInterval = 24 * time.Hour

// stateFile 缓存文件（~/.ur/update-state.json）
func stateFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".ur", "update-state.json")
}

// State 缓存内容
type State struct {
	// LastChecked 上次成功检查时间（RFC3339）
	LastChecked string `json:"lastChecked"`
	// LatestVersion 上次检查到的最新版本
	LatestVersion string `json:"latestVersion,omitempty"`
}

// disabled 是否被环境变量关闭（UR_NO_UPDATE_CHECK=1/true）
func disabled() bool {
	v := strings.TrimSpace(os.Getenv("UR_NO_UPDATE_CHECK"))
	if v == "" {
		return false
	}
	if b, err := strconv.ParseBool(v); err == nil {
		return b
	}
	return v != "0"
}

// ShouldCheck 判断是否需要检查：state 文件不存在，或距上次检查已超过 CheckInterval
func ShouldCheck(state *State, now time.Time) bool {
	if state == nil || state.LastChecked == "" {
		return true
	}
	last, err := time.Parse(time.RFC3339, state.LastChecked)
	if err != nil {
		return true
	}
	return now.Sub(last) >= CheckInterval
}

// loadState 读取缓存 state；文件不存在或损坏返回 nil
func loadState() *State {
	path := stateFile()
	if path == "" {
		return nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil
	}
	return &state
}

// saveState 写入缓存 state（目录不存在则创建）
func saveState(state *State) error {
	path := stateFile()
	if path == "" {
		return fmt.Errorf("无法确定用户目录")
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// versionPart 解析 semver 的 major/minor；解析失败返回 -1
func versionPart(v string, idx int) int {
	v = strings.TrimPrefix(strings.TrimSpace(v), "v")
	parts := strings.SplitN(v, ".", 3)
	if len(parts) <= idx {
		return -1
	}
	n, err := strconv.Atoi(parts[idx])
	if err != nil {
		return -1
	}
	return n
}

// IsOld 判断当前版本是否"过老"：major 不同，或 minor 落后 ≥2
func IsOld(current, latest string) bool {
	curMajor, curMinor := versionPart(current, 0), versionPart(current, 1)
	latMajor, latMinor := versionPart(latest, 0), versionPart(latest, 1)
	if curMajor < 0 || latMajor < 0 {
		return false
	}
	if curMajor != latMajor {
		return true
	}
	if curMinor < 0 || latMinor < 0 {
		return false
	}
	return latMinor-curMinor >= 2
}

// Run 执行一次自动版本检查（应通过 go 启动，非阻塞）。
// 条件：非 dev 版本、未被 UR_NO_UPDATE_CHECK 关闭、距上次检查超 24h。
// 结果只写 stderr；所有失败静默。
func Run(ctx context.Context, stderr io.Writer) {
	if version.IsDev() || disabled() {
		return
	}
	state := loadState()
	if !ShouldCheck(state, time.Now()) {
		return
	}
	release, err := upgrade.FetchLatestRelease()
	if err != nil {
		return // 网络失败静默
	}
	latest := release.TagName
	if !upgrade.IsNewer(version.BuildVersion, latest) {
		// 已是最新：仍刷新检查时间，避免每次请求
		saveState(&State{LastChecked: time.Now().Format(time.RFC3339), LatestVersion: latest})
		return
	}
	// 写缓存后提醒
	saveState(&State{LastChecked: time.Now().Format(time.RFC3339), LatestVersion: latest})
	prefix := ""
	if IsOld(version.BuildVersion, latest) {
		prefix = "警告：当前版本过旧，建议尽快升级。"
	}
	fmt.Fprintf(stderr, "%s发现新版本 %s（当前 %s），运行 ur upgrade 升级\n",
		prefix, latest, version.BuildVersion)
}
