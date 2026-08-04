// skillinstall_test.go — skills 部署逻辑单元测试
package skillinstall

import (
	"os"
	"path/filepath"
	"testing"
)

// makeFakeSource 构造一个最小内置 skills 源（SKILL.md + 一个子域 + _meta.json）
func makeFakeSource(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	mustWrite(t, filepath.Join(dir, "SKILL.md"), "---\nname: ur-api\n---\n")
	mustWrite(t, filepath.Join(dir, "ur-view", "SKILL.md"), "---\nname: ur-view\n---\n")
	mustWrite(t, filepath.Join(dir, "_meta.json"), `{"version":"v0.4.0"}`)
	return dir
}

// mustWrite 写文件（自动创建父目录）
func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

// makeFakeHome 构造带 AI skills 目录的临时 HOME
func makeFakeHome(t *testing.T) string {
	t.Helper()
	home := t.TempDir()
	mustWrite(t, filepath.Join(home, ".claude", "skills", "other-skill", "SKILL.md"), "# other\n")
	mustWrite(t, filepath.Join(home, ".agents", "skills", "other2", "SKILL.md"), "# other2\n")
	return home
}

// TestDetectTargets 校验用户级与项目级目标探测（通过 HOME 环境变量注入）
func TestDetectTargets(t *testing.T) {
	home := makeFakeHome(t)
	t.Setenv("HOME", home)

	// 项目级：构造 git 根 + .claude/skills + .agents/skills
	proj := t.TempDir()
	mustWrite(t, filepath.Join(proj, ".git", "HEAD"), "ref: refs/heads/main\n")
	mustWrite(t, filepath.Join(proj, ".claude", "skills", "x", "SKILL.md"), "# x\n")
	mustWrite(t, filepath.Join(proj, ".agents", "skills", "y", "SKILL.md"), "# y\n")

	targets, err := DetectTargets(proj)
	if err != nil {
		t.Fatalf("DetectTargets: %v", err)
	}
	if len(targets) != 4 {
		t.Fatalf("expect 4 targets (user claude+codex, project claude+codex), got %d: %+v", len(targets), targets)
	}
	got := map[string]bool{}
	for _, tg := range targets {
		got[tg.Scope+"/"+tg.Kind] = true
	}
	for _, want := range []string{"user/claude", "user/codex", "project/claude", "project/codex"} {
		if !got[want] {
			t.Errorf("missing target %s in %+v", want, targets)
		}
	}
}

// TestDetectTargets_NoTargets 无任何 AI 目录时返回空
func TestDetectTargets_NoTargets(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	targets, err := DetectTargets(t.TempDir())
	if err != nil {
		t.Fatalf("DetectTargets: %v", err)
	}
	if len(targets) != 0 {
		t.Fatalf("expect no targets, got %+v", targets)
	}
}

// TestInstall_CopyAndPreserveOthers 校验整体拷贝安装、保留其他 skill、幂等覆盖
func TestInstall_CopyAndPreserveOthers(t *testing.T) {
	home := makeFakeHome(t)
	src := makeFakeSource(t)

	targets := []Target{{Path: filepath.Join(home, ".claude", "skills"), Scope: "user", Kind: "claude"}}
	result, err := Install(src, targets, false)
	if err != nil {
		t.Fatalf("Install: %v", err)
	}
	if len(result.Targets) != 1 || !result.Targets[0].Installed || result.Targets[0].Updated {
		t.Fatalf("first install should be fresh install, got %+v", result.Targets)
	}

	dest := filepath.Join(home, ".claude", "skills", "ur-api")
	for _, f := range []string{"SKILL.md", "ur-view/SKILL.md", "_meta.json"} {
		if _, err := os.Stat(filepath.Join(dest, f)); err != nil {
			t.Errorf("ur-api/%s should exist: %v", f, err)
		}
	}
	// 其他 AI 自有 skill 保留
	if _, err := os.Stat(filepath.Join(home, ".claude", "skills", "other-skill", "SKILL.md")); err != nil {
		t.Errorf("other-skill should be preserved: %v", err)
	}

	// 修改源后再次安装 → 覆盖更新（Updated=true）
	mustWrite(t, filepath.Join(src, "SKILL.md"), "---\nname: ur-api\n# v2\n---\n")
	result2, err := Install(src, targets, false)
	if err != nil {
		t.Fatalf("second Install: %v", err)
	}
	if !result2.Targets[0].Updated || !result2.Targets[0].Installed {
		t.Fatalf("second install should update, got %+v", result2.Targets)
	}
	data, _ := os.ReadFile(filepath.Join(dest, "SKILL.md"))
	if !contains(string(data), "# v2") {
		t.Errorf("ur-api SKILL.md should be updated")
	}
}

// TestInstall_DryRun 校验 --dry-run 不写盘
func TestInstall_DryRun(t *testing.T) {
	home := makeFakeHome(t)
	src := makeFakeSource(t)
	targets := []Target{{Path: filepath.Join(home, ".claude", "skills"), Scope: "user", Kind: "claude"}}
	result, err := Install(src, targets, true)
	if err != nil {
		t.Fatalf("Install dry-run: %v", err)
	}
	if len(result.Targets) != 1 || !result.Targets[0].Installed {
		t.Fatalf("dry-run should report installed, got %+v", result.Targets)
	}
	if _, err := os.Stat(filepath.Join(home, ".claude", "skills", "ur-api")); !os.IsNotExist(err) {
		t.Errorf("dry-run should not create ur-api dir")
	}
}

// TestInstall_EmptySource 源不存在时报错
func TestInstall_EmptySource(t *testing.T) {
	if _, err := Install("", nil, false); err == nil {
		t.Fatal("empty source should error")
	}
	if _, err := Install(filepath.Join(t.TempDir(), "nope"), nil, false); err == nil {
		t.Fatal("missing source dir should error")
	}
}

// contains 简单子串判断
func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(sub) == 0 || indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
