// github_test.go 验证 Release 版本比较不会把旧版本误判为升级。
package upgrade

import "testing"

// TestIsNewer 覆盖补丁版本、回退版本和预发布版本。
func TestIsNewer(t *testing.T) {
	tests := []struct {
		current string
		latest  string
		want    bool
	}{
		{"v0.6.1", "v0.6.2", true},
		{"v0.6.2", "v0.7.0", true},
		{"v0.6.2", "v0.6.2", false},
		{"v0.6.2", "v0.6.1", false},
		{"v1.0.0-rc.1", "v1.0.0", true},
		{"v1.0.0", "v1.0.1-rc.1", true},
		{"v1.0.0-rc.2", "v1.0.0-rc.10", true},
		{"v1.0.0-rc.10", "v1.0.0-rc.2", false},
		{"dev", "v1.0.0", false},
		{"bad", "v1.0.0", false},
	}
	for _, test := range tests {
		if got := IsNewer(test.current, test.latest); got != test.want {
			t.Errorf("IsNewer(%q, %q)=%v, want %v", test.current, test.latest, got, test.want)
		}
	}
}
