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
	if err := os.Remove(filepath.Join(targets[2].Path, "ur-api", "ur-view", "GUIDE.md")); err != nil {
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

// TestFindDuplicateDirs 校验仅识别 ur-api 前缀且非 ur-api 本身的目录为残留。
func TestFindDuplicateDirs(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, filepath.Join(root, "ur-api", "SKILL.md"), "# cur\n")
	mustWrite(t, filepath.Join(root, "ur-api.v0.4.1.bak", "SKILL.md"), "# bak\n")
	mustWrite(t, filepath.Join(root, "ur-api.ur-bak", "SKILL.md"), "# tmp\n")
	mustWrite(t, filepath.Join(root, "other-skill", "SKILL.md"), "# other\n")
	mustWrite(t, filepath.Join(root, "ur-api-note.md"), "# 非目录\n")

	got := FindDuplicateDirs(root)
	want := []string{
		filepath.Join(root, "ur-api.ur-bak"),
		filepath.Join(root, "ur-api.v0.4.1.bak"),
	}
	if len(got) != len(want) {
		t.Fatalf("FindDuplicateDirs = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("FindDuplicateDirs[%d] = %s, want %s", i, got[i], want[i])
		}
	}
}

// TestInspectTargets_Duplicate 校验 ur-api 本身正常但存在残留备份目录时报告 duplicate。
func TestInspectTargets_Duplicate(t *testing.T) {
	source := makeFakeSource(t)
	root := t.TempDir()
	target := Target{Name: "dup", Path: filepath.Join(root, "dup")}
	if _, err := Install(source, []Target{target}, false); err != nil {
		t.Fatalf("Install: %v", err)
	}
	// 模拟手工升级遗留的备份目录：留在 skills 扫描范围内
	mustWrite(t, filepath.Join(target.Path, "ur-api.v0.4.1.bak", "SKILL.md"), "---\nname: ur-api\n---\n")

	result, err := InspectTargets(source, []Target{target})
	if err != nil {
		t.Fatalf("InspectTargets: %v", err)
	}
	status := result.Targets[0]
	if status.State != StatusDuplicate {
		t.Errorf("state = %s, want %s", status.State, StatusDuplicate)
	}
	if len(status.DuplicateDirs) != 1 || status.DuplicateDirs[0] != filepath.Join(target.Path, "ur-api.v0.4.1.bak") {
		t.Errorf("duplicateDirs = %v", status.DuplicateDirs)
	}
}
