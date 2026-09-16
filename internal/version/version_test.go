// version_test.go 验证 CLI 与随包 Skills 的版本信息读取。
package version

import (
	"os"
	"path/filepath"
	"testing"
)

// TestGetSkillsInfoReadsRootMetadata 校验发布包根目录中的 _meta.json 会作为 Skills 版本。
func TestGetSkillsInfoReadsRootMetadata(t *testing.T) {
	root := t.TempDir()
	binaryPath := filepath.Join(root, "ur")
	skillsDir := filepath.Join(root, "skill")
	if err := os.MkdirAll(filepath.Join(skillsDir, "ur-device"), 0o755); err != nil {
		t.Fatalf("创建 Skills 目录: %v", err)
	}
	if err := os.WriteFile(filepath.Join(skillsDir, "_meta.json"), []byte(`{"version":"v0.6.2"}`), 0o644); err != nil {
		t.Fatalf("写入版本元数据: %v", err)
	}
	if err := os.WriteFile(filepath.Join(skillsDir, "ur-device", "SKILL.md"), []byte("# skill\n"), 0o644); err != nil {
		t.Fatalf("写入 Skill: %v", err)
	}

	info := GetSkillsInfo(binaryPath)
	if info.Version != "v0.6.2" || info.Count != 1 {
		t.Fatalf("Skills 信息不正确: %#v", info)
	}
}
