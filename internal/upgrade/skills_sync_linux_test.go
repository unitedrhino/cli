// 文件说明：验证从独立文件系统同步 Skills、重复更新及暂存失败时旧目录完整保留。
package upgrade

import (
	"os"
	"path/filepath"
	"testing"
)

// TestSyncSkillsAcrossFilesystems 覆盖解压目录位于独立 tmpfs、目标已存在与首次安装。
func TestSyncSkillsAcrossFilesystems(t *testing.T) {
	extractDir, err := os.MkdirTemp("/dev/shm", "ur-skills-test-*")
	if err != nil {
		t.Skipf("独立 tmpfs 不可用: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(extractDir) })
	source := filepath.Join(extractDir, "skill", "ur-view")
	if err := os.MkdirAll(source, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "SKILL.md"), []byte("new skill"), 0o644); err != nil {
		t.Fatal(err)
	}
	parent := t.TempDir()
	target := filepath.Join(parent, "skill")
	for run := 0; run < 3; run++ {
		count, err := SyncSkillsFromExtract(extractDir, target+string(os.PathSeparator))
		if err != nil || count != 1 {
			t.Fatalf("第 %d 次同步失败: %d %v", run, count, err)
		}
		content, err := os.ReadFile(filepath.Join(target, "ur-view", "SKILL.md"))
		if err != nil || string(content) != "new skill" {
			t.Fatalf("同步内容不符: %q %v", content, err)
		}
		entries, err := os.ReadDir(parent)
		if err != nil || len(entries) != 1 {
			t.Fatalf("残留暂存或备份: %v %v", entries, err)
		}
	}
}

// TestSyncSkillsCopyFailure 验证源文件损坏时不移动旧 Skills，且允许修复源后重试。
func TestSyncSkillsCopyFailure(t *testing.T) {
	extractDir := t.TempDir()
	source := filepath.Join(extractDir, "skill")
	if err := os.MkdirAll(source, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(source, "missing"), filepath.Join(source, "broken")); err != nil {
		t.Fatal(err)
	}
	parent := t.TempDir()
	target := filepath.Join(parent, "skill")
	if err := os.MkdirAll(target, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(target, "SKILL.md"), []byte("old skill"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := SyncSkillsFromExtract(extractDir, target); err == nil {
		t.Fatal("损坏来源必须失败")
	}
	content, err := os.ReadFile(filepath.Join(target, "SKILL.md"))
	if err != nil || string(content) != "old skill" {
		t.Fatalf("旧版未保留: %q %v", content, err)
	}
	entries, err := os.ReadDir(parent)
	if err != nil || len(entries) != 1 {
		t.Fatalf("失败后残留暂存: %v %v", entries, err)
	}
}
