// 文件说明：复现跨文件系统升级失败，并验证复制失败时旧二进制与安装目录保持完整。
package upgrade

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

// TestReplaceBinaryAcrossFilesystems 在独立 tmpfs 中准备下载文件，验证真实 EXDEV 场景。
func TestReplaceBinaryAcrossFilesystems(t *testing.T) {
	sourceDir, err := os.MkdirTemp("/dev/shm", "ur-upgrade-test-*")
	if err != nil {
		t.Skipf("独立 tmpfs 不可用: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(sourceDir) })
	targetDir := t.TempDir()
	var sourceStat, targetStat syscall.Stat_t
	if err := syscall.Stat(sourceDir, &sourceStat); err != nil {
		t.Fatal(err)
	}
	if err := syscall.Stat(targetDir, &targetStat); err != nil {
		t.Fatal(err)
	}
	if sourceStat.Dev == targetStat.Dev {
		t.Skip("测试环境目录处于同一文件系统")
	}
	source := filepath.Join(sourceDir, "ur")
	target := filepath.Join(targetDir, "ur")
	if err := os.WriteFile(source, []byte("new-version"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(target, []byte("old-version"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := replaceBinary(source, target); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(target)
	if err != nil || string(content) != "new-version" {
		t.Fatalf("替换内容错误: %q %v", content, err)
	}
	info, err := os.Stat(target)
	if err != nil || info.Mode().Perm() != 0o755 {
		t.Fatalf("新二进制执行权限错误: %v %v", info, err)
	}
	entries, err := os.ReadDir(targetDir)
	if err != nil || len(entries) != 1 {
		t.Fatalf("暂存文件未清理: %v %v", entries, err)
	}
}

// TestReplaceBinaryCopyFailure 在新文件不可读取时验证旧版和清理结果。
func TestReplaceBinaryCopyFailure(t *testing.T) {
	targetDir := t.TempDir()
	target := filepath.Join(targetDir, "ur")
	if err := os.WriteFile(target, []byte("old-version"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := replaceBinary(t.TempDir(), target); err == nil {
		t.Fatal("目录不能作为新二进制")
	}
	content, err := os.ReadFile(target)
	if err != nil || string(content) != "old-version" {
		t.Fatalf("旧版被破坏: %q %v", content, err)
	}
	entries, err := os.ReadDir(targetDir)
	if err != nil || len(entries) != 1 {
		t.Fatalf("失败后暂存文件未清理: %v %v", entries, err)
	}
}
