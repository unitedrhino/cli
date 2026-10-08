// aggregate_test.go — 聚合形态转换（子域 SKILL.md 降级 GUIDE.md）单元测试。
package skillinstall

import (
	"os"
	"path/filepath"
	"testing"
)

// TestInstalledRelPath 校验安装形态路径映射：顶层保留、子域改名。
func TestInstalledRelPath(t *testing.T) {
	cases := map[string]string{
		"SKILL.md":                      "SKILL.md",
		"_meta.json":                    "_meta.json",
		"ur-view/SKILL.md":              "ur-view/GUIDE.md",
		"references/api-conventions.md": "references/api-conventions.md",
		"ur-view/assets/app.js":         "ur-view/assets/app.js",
	}
	for input, want := range cases {
		if got := installedRelPath(input); got != want {
			t.Errorf("installedRelPath(%q) = %q, want %q", input, got, want)
		}
	}
}

// TestStripYAMLFrontmatter 校验 frontmatter 剥离与容错。
func TestStripYAMLFrontmatter(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"标准块", "---\nname: x\n---\n正文", "正文"},
		{"块后空行", "---\nname: x\n---\n\n正文", "正文"},
		{"无闭合原样", "---\nname: x\n正文", "---\nname: x\n正文"},
		{"无块原样", "# 标题\n正文", "# 标题\n正文"},
		{"正文含分隔线", "---\na: 1\n---\n---\n正文", "---\n正文"},
		{"首行非分隔", "---- \nname: x\n---\n正文", "---- \nname: x\n---\n正文"},
	}
	for _, c := range cases {
		if got := string(stripYAMLFrontmatter([]byte(c.input))); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}

// TestInstallAggregateDemotesSubskills 端到端校验安装产物聚合形态：
// 只保留顶层 SKILL.md 技能入口，子域降级为 GUIDE.md。
func TestInstallAggregateDemotesSubskills(t *testing.T) {
	src := t.TempDir()
	mustWrite(t, filepath.Join(src, "SKILL.md"), "---\nname: ur-api\ndescription: 入口\n---\n# ur-api\n")
	mustWrite(t, filepath.Join(src, "ur-device", "SKILL.md"), "---\nname: ur-device\ndescription: 子域\n---\n# 设备管理\n端点列表...")
	mustWrite(t, filepath.Join(src, "ur-device", "references", "api.md"), "# API 参考\n")
	mustWrite(t, filepath.Join(src, "_meta.json"), `{"version":"v0.5.0"}`)

	targetRoot := t.TempDir()
	if _, err := Install(src, []Target{{Name: "t", Path: targetRoot, Scope: "user", Kind: "claude"}}, false); err != nil {
		t.Fatalf("Install: %v", err)
	}
	dest := filepath.Join(targetRoot, "ur-api")

	// 顶层入口保留原内容
	top, err := os.ReadFile(filepath.Join(dest, "SKILL.md"))
	if err != nil {
		t.Fatalf("top SKILL.md: %v", err)
	}
	if top[0] != '-' {
		t.Errorf("top SKILL.md should keep frontmatter, got %q", top)
	}
	// 子域降级：SKILL.md 不存在，GUIDE.md 无 frontmatter
	if _, err := os.Stat(filepath.Join(dest, "ur-device", "SKILL.md")); err == nil {
		t.Error("ur-device/SKILL.md should not exist in aggregate install")
	}
	guide, err := os.ReadFile(filepath.Join(dest, "ur-device", "GUIDE.md"))
	if err != nil {
		t.Fatalf("ur-device/GUIDE.md: %v", err)
	}
	if guide[0] == '-' || len(guide) == 0 {
		t.Errorf("ur-device/GUIDE.md should strip frontmatter, got %q", guide)
	}
	// 参考文档原样
	if _, err := os.Stat(filepath.Join(dest, "ur-device", "references", "api.md")); err != nil {
		t.Errorf("references/api.md should exist as-is: %v", err)
	}
	// 聚合安装后 status 应判定 current（源签名按安装形态计算）
	statuses, err := InspectTargets(src, []Target{{Name: "t", Path: targetRoot, Scope: "user", Kind: "claude"}})
	if err != nil {
		t.Fatalf("InspectTargets: %v", err)
	}
	if statuses.Targets[0].State != StatusCurrent {
		t.Errorf("aggregate install should be current, got %s (%+v)", statuses.Targets[0].State, statuses.Targets[0])
	}
}
