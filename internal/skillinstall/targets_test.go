// targets_test.go — 验证通用 Skills 目标配置的持久化、合并与删除行为。
package skillinstall

import (
	"os"
	"path/filepath"
	"testing"
)

// TestConfiguredTargetsLifecycle 校验目标新增、更新、读取和删除的完整生命周期。
func TestConfiguredTargetsLifecycle(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "skill-targets.json")
	t.Setenv("UR_SKILL_TARGETS_FILE", configPath)
	firstPath := filepath.Join(t.TempDir(), "first")
	secondPath := filepath.Join(t.TempDir(), "second")

	updated, err := UpsertConfiguredTarget(ConfiguredTarget{Name: "desktop-ai", Type: TargetTypeFilesystem, Path: firstPath})
	if err != nil || updated {
		t.Fatalf("first upsert = updated:%v err:%v", updated, err)
	}
	updated, err = UpsertConfiguredTarget(ConfiguredTarget{Name: "desktop-ai", Type: TargetTypeFilesystem, Path: secondPath})
	if err != nil || !updated {
		t.Fatalf("second upsert = updated:%v err:%v", updated, err)
	}
	targets, err := LoadConfiguredTargets()
	if err != nil {
		t.Fatalf("LoadConfiguredTargets: %v", err)
	}
	if len(targets) != 1 || targets[0].Path != secondPath || !targets[0].Enabled {
		t.Fatalf("unexpected configured targets: %+v", targets)
	}
	info, err := os.Stat(configPath)
	if err != nil {
		t.Fatalf("stat config: %v", err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Errorf("config permission = %o, want 600", info.Mode().Perm())
	}
	removed, err := RemoveConfiguredTarget("desktop-ai")
	if err != nil || !removed {
		t.Fatalf("remove = removed:%v err:%v", removed, err)
	}
	targets, err = LoadConfiguredTargets()
	if err != nil || len(targets) != 0 {
		t.Fatalf("targets after remove = %+v err:%v", targets, err)
	}
}

// TestResolveTargets_ConfiguredTargetWins 校验登记目标与自动发现路径重复时保留登记名称。
func TestResolveTargets_ConfiguredTargetWins(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_SKILL_TARGETS_FILE", filepath.Join(home, ".ur", "skill-targets.json"))
	workbuddy := filepath.Join(home, ".codebuddy", "skills")
	mustWrite(t, filepath.Join(workbuddy, "other", "SKILL.md"), "# other\n")
	if _, err := UpsertConfiguredTarget(ConfiguredTarget{Name: "team-workbuddy", Type: TargetTypeFilesystem, Path: workbuddy}); err != nil {
		t.Fatalf("UpsertConfiguredTarget: %v", err)
	}
	targets, err := ResolveTargets(t.TempDir())
	if err != nil {
		t.Fatalf("ResolveTargets: %v", err)
	}
	if len(targets) != 1 || targets[0].Name != "team-workbuddy" || targets[0].Origin != "configured" {
		t.Fatalf("unexpected resolved targets: %+v", targets)
	}
}

// TestConfiguredTargetValidation 校验不支持的类型和空目录会被拒绝。
func TestConfiguredTargetValidation(t *testing.T) {
	t.Setenv("UR_SKILL_TARGETS_FILE", filepath.Join(t.TempDir(), "skill-targets.json"))
	if _, err := UpsertConfiguredTarget(ConfiguredTarget{Name: "bad", Type: "marketplace", Path: t.TempDir()}); err == nil {
		t.Fatal("unsupported target type should fail")
	}
	if _, err := UpsertConfiguredTarget(ConfiguredTarget{Name: "empty", Type: TargetTypeFilesystem}); err == nil {
		t.Fatal("empty target path should fail")
	}
}
