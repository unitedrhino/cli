// updatecheck_test.go — 自动版本检查逻辑单元测试
package updatecheck

import (
	"errors"
	"path/filepath"
	"testing"
	"time"

	"gitee.com/unitedrhino/cli/internal/notice"
	"gitee.com/unitedrhino/cli/internal/upgrade"
)

// TestShouldCheck 校验 24h 低频缓存逻辑
func TestShouldCheck(t *testing.T) {
	now := time.Now()
	cases := []struct {
		name  string
		state *State
		want  bool
	}{
		{"无缓存文件", nil, true},
		{"空状态", &State{}, true},
		{"刚检查过", &State{LastChecked: now.Add(-time.Hour).Format(time.RFC3339)}, false},
		{"23 小时前", &State{LastChecked: now.Add(-23 * time.Hour).Format(time.RFC3339)}, false},
		{"24 小时前", &State{LastChecked: now.Add(-24 * time.Hour).Format(time.RFC3339)}, true},
		{"损坏时间戳", &State{LastChecked: "not-a-time"}, true},
	}
	for _, tc := range cases {
		if got := ShouldCheck(tc.state, now); got != tc.want {
			t.Errorf("%s: ShouldCheck = %v, want %v", tc.name, got, tc.want)
		}
	}
}

// TestLoadCachedNotice 验证短命令无需等待网络即可读取已有升级提示。
func TestLoadCachedNotice(t *testing.T) {
	t.Setenv("UR_UPDATE_STATE_FILE", filepath.Join(t.TempDir(), "update-state.json"))
	if err := saveState(&State{LatestVersion: "v0.6.2", ReleaseURL: "https://example.test/v0.6.2"}); err != nil {
		t.Fatalf("保存缓存: %v", err)
	}
	notice.Reset()
	LoadCachedNotice("v0.6.1")
	payload := notice.Snapshot()
	if payload.Update == nil || payload.Update.Latest != "v0.6.2" || payload.Update.Command != "ur upgrade" {
		t.Fatalf("升级提示不完整: %#v", payload.Update)
	}
}

// TestRefreshFailureBackoff 验证网络失败会保留旧结果并设置退避时间。
func TestRefreshFailureBackoff(t *testing.T) {
	t.Setenv("UR_UPDATE_STATE_FILE", filepath.Join(t.TempDir(), "update-state.json"))
	originalFetch := fetchLatestRelease
	fetchLatestRelease = func() (*upgrade.Release, error) { return nil, errors.New("offline") }
	t.Cleanup(func() { fetchLatestRelease = originalFetch })
	if err := saveState(&State{LatestVersion: "v0.6.2"}); err != nil {
		t.Fatalf("保存缓存: %v", err)
	}
	Refresh("v0.6.1")
	state := loadState()
	if state == nil || state.LatestVersion != "v0.6.2" || state.NextCheckAt == "" {
		t.Fatalf("失败退避状态不完整: %#v", state)
	}
	next, err := time.Parse(time.RFC3339, state.NextCheckAt)
	if err != nil || time.Until(next) < 55*time.Minute {
		t.Fatalf("失败退避时间不正确: %q err=%v", state.NextCheckAt, err)
	}
}

// TestRefreshSuccess 验证刷新成功会更新缓存并准备当前进程提示。
func TestRefreshSuccess(t *testing.T) {
	t.Setenv("UR_UPDATE_STATE_FILE", filepath.Join(t.TempDir(), "update-state.json"))
	originalFetch := fetchLatestRelease
	fetchLatestRelease = func() (*upgrade.Release, error) {
		return &upgrade.Release{TagName: "v0.6.2", HTMLURL: "https://example.test/v0.6.2"}, nil
	}
	t.Cleanup(func() { fetchLatestRelease = originalFetch })
	notice.Reset()
	Refresh("v0.6.1")
	state := loadState()
	if state == nil || state.LatestVersion != "v0.6.2" || state.LastSuccessAt == "" {
		t.Fatalf("成功缓存不完整: %#v", state)
	}
	if notice.Snapshot().Update == nil {
		t.Fatal("刷新成功后应准备升级提示")
	}
}

// TestIsOld 校验"版本过老"判定
func TestIsOld(t *testing.T) {
	cases := []struct {
		name    string
		current string
		latest  string
		want    bool
	}{
		{"major 相同 minor 相同", "v0.3.8", "v0.3.8", false},
		{"minor 落后 1", "v0.3.8", "v0.4.0", false},
		{"minor 落后 2", "v0.2.8", "v0.4.0", true},
		{"minor 落后 3", "v0.1.0", "v0.4.0", true},
		{"major 不同", "v1.0.0", "v2.0.0", true},
		{"major 相同 minor 相同但 patch 落后", "v0.3.7", "v0.3.8", false},
		{"latest 比 current 旧", "v0.4.0", "v0.3.8", false},
		{"非法版本", "dev", "v0.4.0", false},
		{"无 v 前缀", "0.2.0", "0.4.0", true},
	}
	for _, tc := range cases {
		if got := IsOld(tc.current, tc.latest); got != tc.want {
			t.Errorf("%s: IsOld(%q, %q) = %v, want %v", tc.name, tc.current, tc.latest, got, tc.want)
		}
	}
}
