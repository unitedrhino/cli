// updatecheck_test.go — 自动版本检查逻辑单元测试
package updatecheck

import (
	"testing"
	"time"
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
