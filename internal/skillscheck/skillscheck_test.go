// skillscheck_test.go 验证客户端 Skills 版本漂移会生成结构化提示。
package skillscheck

import (
	"os"
	"path/filepath"
	"testing"

	"gitee.com/unitedrhino/cli/internal/notice"
	"gitee.com/unitedrhino/cli/internal/skillinstall"
)

// TestLoadNotice 验证登记目标版本落后时生成提示。
func TestLoadNotice(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_SKILL_TARGETS_FILE", filepath.Join(home, ".ur", "targets.json"))
	source := filepath.Join(home, ".ur", "skills")
	writeFile(t, filepath.Join(source, "_meta.json"), `{"version":"v0.6.2"}`)
	targetRoot := filepath.Join(home, "client-skills")
	writeFile(t, filepath.Join(targetRoot, "ur-api", "_meta.json"), `{"version":"v0.6.1"}`)
	if _, err := skillinstall.UpsertConfiguredTarget(skillinstall.ConfiguredTarget{Name: "client", Type: skillinstall.TargetTypeFilesystem, Path: targetRoot, Enabled: true}); err != nil {
		t.Fatalf("登记目标: %v", err)
	}
	notice.Reset()
	LoadNotice(home)
	payload := notice.Snapshot()
	if payload.Skills == nil || payload.Skills.Target != "v0.6.2" || len(payload.Skills.Targets) != 1 {
		t.Fatalf("Skills 提示不完整: %#v", payload.Skills)
	}
}

// TestLoadNoticeIgnoresNewerTarget 验证旧 CLI 不会提示覆盖更新的客户端 Skills。
func TestLoadNoticeIgnoresNewerTarget(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_SKILL_TARGETS_FILE", filepath.Join(home, ".ur", "targets.json"))
	source := filepath.Join(home, ".ur", "skills")
	writeFile(t, filepath.Join(source, "_meta.json"), `{"version":"v0.6.1"}`)
	targetRoot := filepath.Join(home, "client-skills")
	writeFile(t, filepath.Join(targetRoot, "ur-api", "_meta.json"), `{"version":"v0.6.2"}`)
	if _, err := skillinstall.UpsertConfiguredTarget(skillinstall.ConfiguredTarget{Name: "client", Type: skillinstall.TargetTypeFilesystem, Path: targetRoot, Enabled: true}); err != nil {
		t.Fatalf("登记目标: %v", err)
	}
	notice.Reset()
	LoadNotice(home)
	if payload := notice.Snapshot(); payload.Skills != nil {
		t.Fatalf("不应提示用旧版覆盖新版 Skills: %#v", payload.Skills)
	}
}

// TestLoadNoticeIgnoresUnversionedTarget 验证无法判断版本的自维护目录不会被误报。
func TestLoadNoticeIgnoresUnversionedTarget(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_SKILL_TARGETS_FILE", filepath.Join(home, ".ur", "targets.json"))
	source := filepath.Join(home, ".ur", "skills")
	writeFile(t, filepath.Join(source, "_meta.json"), `{"version":"v0.6.2"}`)
	targetRoot := filepath.Join(home, "client-skills")
	writeFile(t, filepath.Join(targetRoot, "ur-api", "SKILL.md"), "# independently maintained\n")
	if _, err := skillinstall.UpsertConfiguredTarget(skillinstall.ConfiguredTarget{Name: "client", Type: skillinstall.TargetTypeFilesystem, Path: targetRoot, Enabled: true}); err != nil {
		t.Fatalf("登记目标: %v", err)
	}
	notice.Reset()
	LoadNotice(home)
	if payload := notice.Snapshot(); payload.Skills != nil {
		t.Fatalf("不应把无版本元数据的目录判定为落后: %#v", payload.Skills)
	}
}

// writeFile 创建测试文件及父目录。
func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("创建目录: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("写文件: %v", err)
	}
}
