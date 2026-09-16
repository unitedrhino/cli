// upgrade_test.go — 验证 CLI 已是最新版时仍可刷新客户端中的 Skills 副本。
package shared

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"gitee.com/unitedrhino/cli/internal/notice"
	"gitee.com/unitedrhino/cli/internal/skillinstall"
	"gitee.com/unitedrhino/cli/internal/upgrade"
	"gitee.com/unitedrhino/cli/internal/version"
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

// TestUpgradeJSONPerformsUpgrade 验证 --json 只改变输出格式，不会退化成只检查。
func TestUpgradeJSONPerformsUpgrade(t *testing.T) {
	originalPerform := performUpgrade
	originalVersion := version.BuildVersion
	called := false
	performUpgrade = func(opts upgrade.Options) (*upgrade.Result, error) {
		called = true
		return &upgrade.Result{CurrentVersion: "v0.6.1", LatestVersion: "v0.6.2"}, nil
	}
	version.BuildVersion = "v0.6.1"
	t.Cleanup(func() {
		performUpgrade = originalPerform
		version.BuildVersion = originalVersion
	})
	notice.Reset()
	var stdout, stderr bytes.Buffer
	if code := runUpgrade([]string{"--json", "--no-skills"}, &stdout, &stderr); code != 0 {
		t.Fatalf("exit=%d stderr=%s", code, stderr.String())
	}
	if !called {
		t.Fatal("--json 应执行升级而不是只检查")
	}
	if !bytes.Contains(stdout.Bytes(), []byte(`"latestVersion": "v0.6.2"`)) {
		t.Fatalf("JSON 结果不完整: %s", stdout.String())
	}
}

// TestUpgradeForceForwarded 验证 --force 会传递到底层升级流程。
func TestUpgradeForceForwarded(t *testing.T) {
	originalPerform := performUpgrade
	originalVersion := version.BuildVersion
	force := false
	performUpgrade = func(opts upgrade.Options) (*upgrade.Result, error) {
		force = opts.Force
		return &upgrade.Result{CurrentVersion: "v0.6.2", LatestVersion: "v0.6.2"}, nil
	}
	version.BuildVersion = "v0.6.2"
	t.Cleanup(func() {
		performUpgrade = originalPerform
		version.BuildVersion = originalVersion
	})

	var stdout, stderr bytes.Buffer
	if code := runUpgrade([]string{"--force", "--json", "--no-skills"}, &stdout, &stderr); code != 0 {
		t.Fatalf("exit=%d stderr=%s", code, stderr.String())
	}
	if !force {
		t.Fatal("--force 应传递到底层升级流程")
	}
}

// TestUpgradeDeploysSkillsByDefault 验证统一升级入口默认同步客户端 Skills。
func TestUpgradeDeploysSkillsByDefault(t *testing.T) {
	originalPerform := performUpgrade
	originalDeploy := deployEmbeddedSkillsNow
	originalVersion := version.BuildVersion
	deployed := false
	performUpgrade = func(opts upgrade.Options) (*upgrade.Result, error) {
		return &upgrade.Result{CurrentVersion: "v0.6.2", LatestVersion: "v0.6.2", UpToDate: true}, nil
	}
	deployEmbeddedSkillsNow = func() (*skillinstall.Result, error) {
		deployed = true
		return &skillinstall.Result{Targets: []skillinstall.TargetResult{{Name: "codex-user", Installed: true}}}, nil
	}
	version.BuildVersion = "v0.6.2"
	t.Cleanup(func() {
		performUpgrade = originalPerform
		deployEmbeddedSkillsNow = originalDeploy
		version.BuildVersion = originalVersion
	})
	var stdout, stderr bytes.Buffer
	if code := runUpgrade([]string{"--json"}, &stdout, &stderr); code != 0 {
		t.Fatalf("exit=%d stderr=%s", code, stderr.String())
	}
	if !deployed {
		t.Fatal("ur upgrade 默认应同步客户端 Skills")
	}
	if !bytes.Contains(stdout.Bytes(), []byte(`"skillsInstalled": true`)) {
		t.Fatalf("JSON 未返回 Skills 同步结果: %s", stdout.String())
	}
}
