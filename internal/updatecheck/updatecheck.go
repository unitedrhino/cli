// Package updatecheck 为 CLI 提供低延迟、可缓存的版本检查与 AI 提示。
package updatecheck

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/notice"
	"gitee.com/unitedrhino/cli/internal/upgrade"
)

const (
	// CheckInterval 是成功检查后的缓存有效期。
	CheckInterval = 24 * time.Hour
	// FailureRetryInterval 是网络失败后的最小重试间隔。
	FailureRetryInterval = time.Hour
	// stateSchemaVersion 是缓存结构版本。
	stateSchemaVersion = 2
)

// State 保存最近一次检查结果；LastChecked 兼容 v0.6.1 及更早版本。
type State struct {
	// SchemaVersion 是缓存结构版本。
	SchemaVersion int `json:"schemaVersion,omitempty"`
	// LastChecked 兼容旧版缓存中的最近成功检查时间。
	LastChecked string `json:"lastChecked,omitempty"`
	// LastAttemptAt 是最近一次网络检查时间。
	LastAttemptAt string `json:"lastAttemptAt,omitempty"`
	// LastSuccessAt 是最近一次成功检查时间。
	LastSuccessAt string `json:"lastSuccessAt,omitempty"`
	// NextCheckAt 是下次允许网络检查的时间。
	NextCheckAt string `json:"nextCheckAt,omitempty"`
	// LatestVersion 是最近一次成功查询到的版本。
	LatestVersion string `json:"latestVersion,omitempty"`
	// ReleaseURL 是最近一次成功查询到的 Release 页面。
	ReleaseURL string `json:"releaseUrl,omitempty"`
}

// fetchLatestRelease 是可在测试中替换的 Release 查询入口。
var fetchLatestRelease = upgrade.FetchLatestRelease

// stateFile 返回缓存路径，测试和受管环境可通过 UR_UPDATE_STATE_FILE 覆盖。
func stateFile() string {
	if path := strings.TrimSpace(os.Getenv("UR_UPDATE_STATE_FILE")); path != "" {
		return path
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".ur", "update-state.json")
}

// disabled 判断是否完全关闭远端版本检查。
func disabled() bool {
	return envEnabled("UR_NO_UPDATE_CHECK")
}

// notifierDisabled 判断是否仅关闭 CLI 更新提示。
func notifierDisabled() bool {
	return envEnabled("UR_NO_UPDATE_NOTIFIER")
}

// envEnabled 解析兼容 1、true 及其他非零文本的开关变量。
func envEnabled(name string) bool {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return false
	}
	if enabled, err := strconv.ParseBool(value); err == nil {
		return enabled
	}
	return value != "0"
}

// ShouldCheck 判断缓存是否已经到达下一次检查时间。
func ShouldCheck(state *State, now time.Time) bool {
	if state == nil {
		return true
	}
	if state.NextCheckAt != "" {
		next, err := time.Parse(time.RFC3339, state.NextCheckAt)
		return err != nil || !now.Before(next)
	}
	last := state.LastSuccessAt
	if last == "" {
		last = state.LastChecked
	}
	if last == "" {
		return true
	}
	checkedAt, err := time.Parse(time.RFC3339, last)
	return err != nil || now.Sub(checkedAt) >= CheckInterval
}

// LoadCachedNotice 同步读取缓存并立即准备结构化提示，不访问网络。
func LoadCachedNotice(currentVersion string) {
	if !strings.HasPrefix(currentVersion, "v") || disabled() || notifierDisabled() {
		return
	}
	state := loadState()
	if state == nil || !upgrade.IsNewer(currentVersion, state.LatestVersion) {
		return
	}
	notice.SetUpdate(notice.Update{
		Current: currentVersion,
		Latest:  state.LatestVersion,
		Message: fmt.Sprintf("ur CLI %s → %s 可升级", currentVersion, state.LatestVersion),
		Command: "ur upgrade",
		URL:     state.ReleaseURL,
	})
}

// Refresh 在缓存过期时刷新远端版本；失败只更新退避时间，不影响业务命令。
func Refresh(currentVersion string) {
	if !strings.HasPrefix(currentVersion, "v") || disabled() {
		return
	}
	state := loadState()
	now := time.Now().UTC()
	if !ShouldCheck(state, now) {
		return
	}
	if state == nil {
		state = &State{}
	}
	state.SchemaVersion = stateSchemaVersion
	state.LastAttemptAt = now.Format(time.RFC3339)
	release, err := fetchLatestRelease()
	if err != nil {
		state.NextCheckAt = now.Add(FailureRetryInterval).Format(time.RFC3339)
		_ = saveState(state)
		return
	}
	state.LastChecked = now.Format(time.RFC3339)
	state.LastSuccessAt = state.LastChecked
	state.NextCheckAt = now.Add(CheckInterval).Format(time.RFC3339)
	state.LatestVersion = release.TagName
	state.ReleaseURL = release.HTMLURL
	_ = saveState(state)
	if notifierDisabled() || !upgrade.IsNewer(currentVersion, release.TagName) {
		return
	}
	notice.SetUpdate(notice.Update{
		Current: currentVersion,
		Latest:  release.TagName,
		Message: fmt.Sprintf("ur CLI %s → %s 可升级", currentVersion, release.TagName),
		Command: "ur upgrade",
		URL:     release.HTMLURL,
	})
}

// loadState 读取缓存；文件不存在或损坏时返回 nil。
func loadState() *State {
	path := stateFile()
	if path == "" {
		return nil
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var state State
	if json.Unmarshal(raw, &state) != nil {
		return nil
	}
	return &state
}

// saveState 通过临时文件和重命名原子更新缓存。
func saveState(state *State) error {
	path := stateFile()
	if path == "" {
		return fmt.Errorf("无法确定用户目录")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	temporary, err := os.CreateTemp(filepath.Dir(path), ".update-state-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	defer os.Remove(temporaryPath)
	if err := temporary.Chmod(0o644); err != nil {
		temporary.Close()
		return err
	}
	if _, err := temporary.Write(raw); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	return os.Rename(temporaryPath, path)
}

// IsOld 判断当前版本是否跨 major 或落后至少两个 minor，用于保留旧调用兼容性。
func IsOld(current, latest string) bool {
	currentParts := versionParts(current)
	latestParts := versionParts(latest)
	if currentParts == nil || latestParts == nil {
		return false
	}
	return currentParts[0] != latestParts[0] || latestParts[1]-currentParts[1] >= 2
}

// versionParts 提取 major 和 minor。
func versionParts(value string) []int {
	value = strings.TrimPrefix(strings.TrimSpace(value), "v")
	parts := strings.Split(value, ".")
	if len(parts) < 2 {
		return nil
	}
	major, majorErr := strconv.Atoi(parts[0])
	minor, minorErr := strconv.Atoi(parts[1])
	if majorErr != nil || minorErr != nil {
		return nil
	}
	return []int{major, minor}
}
