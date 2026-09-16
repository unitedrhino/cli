// upgrade_test.go — 验证 CLI 已是最新版时仍可刷新客户端中的 Skills 副本。
package shared

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"gitee.com/unitedrhino/cli/internal/skillinstall"
)

// TestInstallEmbeddedSkills 校验升级流程复用的部署步骤会更新已登记目标。
func TestInstallEmbeddedSkills(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_SKILL_TARGETS_FILE", filepath.Join(home, ".ur", "skill-targets.json"))
	source := filepath.Join(home, ".ur", "skills")
	writeSharedTestFile(t, filepath.Join(source, "SKILL.md"), "---\nname: ur-api\n---\n")
	writeSharedTestFile(t, filepath.Join(source, "_meta.json"), `{"version":"v1.0.0"}`)
	targetDir := filepath.Join(home, "client-skills")
	if _, err := skillinstall.UpsertConfiguredTarget(skillinstall.ConfiguredTarget{Name: "client", Type: skillinstall.TargetTypeFilesystem, Path: targetDir}); err != nil {
		t.Fatalf("UpsertConfiguredTarget: %v", err)
	}
	writeSharedTestFile(t, filepath.Join(targetDir, "ur-api", "stale.txt"), "stale")

	var stdout, stderr bytes.Buffer
	if code := installEmbeddedSkills(&stdout, &stderr); code != 0 {
		t.Fatalf("installEmbeddedSkills code=%d stderr=%s", code, stderr.String())
	}
	if _, err := os.Stat(filepath.Join(targetDir, "ur-api", "SKILL.md")); err != nil {
		t.Fatalf("installed skill missing: %v", err)
	}
	if _, err := os.Stat(filepath.Join(targetDir, "ur-api", "stale.txt")); !os.IsNotExist(err) {
		t.Fatalf("stale file should be removed, err=%v", err)
	}
}
