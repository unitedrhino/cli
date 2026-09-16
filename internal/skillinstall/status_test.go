// status_test.go — 验证 Skills 版本与文件完整性诊断。
package skillinstall

import (
	"os"
	"path/filepath"
	"testing"
)

// TestInspectTargets 校验未安装、旧版本、残缺和当前版本四种状态。
func TestInspectTargets(t *testing.T) {
	source := makeFakeSource(t)
	root := t.TempDir()
	targets := []Target{
		{Name: "missing", Path: filepath.Join(root, "missing")},
		{Name: "outdated", Path: filepath.Join(root, "outdated")},
		{Name: "incomplete", Path: filepath.Join(root, "incomplete")},
		{Name: "current", Path: filepath.Join(root, "current")},
	}
	for _, target := range targets[1:] {
		if _, err := Install(source, []Target{target}, false); err != nil {
			t.Fatalf("Install %s: %v", target.Name, err)
		}
	}
	mustWrite(t, filepath.Join(targets[1].Path, "ur-api", "_meta.json"), `{"version":"v0.3.0"}`)
	if err := os.Remove(filepath.Join(targets[2].Path, "ur-api", "ur-view", "SKILL.md")); err != nil {
		t.Fatalf("remove incomplete fixture: %v", err)
	}

	result, err := InspectTargets(source, targets)
	if err != nil {
		t.Fatalf("InspectTargets: %v", err)
	}
	states := map[string]string{}
	for _, status := range result.Targets {
		states[status.Name] = status.State
	}
	want := map[string]string{
		"missing": StatusMissing, "outdated": StatusOutdated,
		"incomplete": StatusIncomplete, "current": StatusCurrent,
	}
	for name, state := range want {
		if states[name] != state {
			t.Errorf("state %s = %s, want %s", name, states[name], state)
		}
	}
}
