// skills_test.go — skills 子命令的 ~ 路径展开与 download 参数解析/输出构造单元测试
package shared

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"gitee.com/unitedrhino/cli/internal/upgrade"
)

// TestExpandHomePath 校验 ~、~/、~\ 前缀展开为用户主目录，其余路径原样返回
func TestExpandHomePath(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	cases := []struct {
		in   string
		want string
	}{
		{"~", home},
		{"~/skills", filepath.Join(home, "skills")},
		{`~\skills`, filepath.Join(home, "skills")},
		{`~\.workbuddy\skills`, filepath.Join(home, ".workbuddy", "skills")},
		{"/abs/path", "/abs/path"},
		{"relative/dir", "relative/dir"},
		{"~user/x", "~user/x"}, // ~user 形式不支持展开，原样返回
		{"", ""},
	}
	for _, tc := range cases {
		if got := expandHomePath(tc.in); got != tc.want {
			t.Errorf("expandHomePath(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// TestFormatSkillsDownloadResult 校验 JSON 与人类可读两种输出形态
func TestFormatSkillsDownloadResult(t *testing.T) {
	res := &upgrade.SkillsDownloadResult{
		Version:     "v0.4.1",
		DownloadURL: "https://example.com/ur-api-skills-v0.4.1.zip",
		LocalPath:   "/home/u/.ur/downloads/ur-api",
	}

	// JSON：单行事件，含关键字段
	var payload struct {
		Event       string `json:"event"`
		DownloadURL string `json:"downloadUrl"`
		LocalPath   string `json:"localPath"`
		InstallHint string `json:"installHint"`
	}
	line := formatSkillsDownloadResult(res, true)
	if err := json.Unmarshal([]byte(line), &payload); err != nil {
		t.Fatalf("json output should be valid single-line json: %v\n%s", err, line)
	}
	if payload.Event != "skills_downloaded" {
		t.Errorf("event = %q, want skills_downloaded", payload.Event)
	}
	if payload.DownloadURL != res.DownloadURL || payload.LocalPath != res.LocalPath {
		t.Errorf("downloadUrl/localPath mismatch: %+v", payload)
	}
	if payload.InstallHint == "" {
		t.Errorf("installHint should not be empty")
	}

	// 非 JSON：包含来源、路径与拷贝指引
	text := formatSkillsDownloadResult(res, false)
	for _, want := range []string{res.DownloadURL, res.LocalPath, "拷贝"} {
		if !bytes.Contains([]byte(text), []byte(want)) {
			t.Errorf("text output should contain %q, got:\n%s", want, text)
		}
	}
}

// makeSkillsZip 构造最小 skills zip（顶层 ur-api/ 含 SKILL.md 与 _meta.json），返回本地文件路径
func makeSkillsZip(t *testing.T, withSkill bool) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "ur-api-skills-v0.4.1.zip")
	f, err := os.Create(p)
	if err != nil {
		t.Fatalf("create zip: %v", err)
	}
	defer f.Close()
	zw := zip.NewWriter(f)
	files := map[string]string{}
	if withSkill {
		files["ur-api/SKILL.md"] = "---\nname: ur-api\n---\n"
		files["ur-api/_meta.json"] = `{"version":"v0.4.1"}`
	} else {
		files["ur-api/README.md"] = "no skill here"
	}
	for name, content := range files {
		w, err := zw.Create(name)
		if err != nil {
			t.Fatalf("zip create %s: %v", name, err)
		}
		if _, err := w.Write([]byte(content)); err != nil {
			t.Fatalf("zip write %s: %v", name, err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatalf("zip close: %v", err)
	}
	return p
}

// TestRunSkillsDownload 通过本地 httptest 模拟下载源，校验下载解压全流程与两种输出
func TestRunSkillsDownload(t *testing.T) {
	zipPath := makeSkillsZip(t, true)
	srv := httptest.NewServer(http.FileServer(http.Dir(filepath.Dir(zipPath))))
	defer srv.Close()

	outDir := filepath.Join(t.TempDir(), "downloads")
	var stdout, stderr bytes.Buffer
	code := runSkillsDownload([]string{
		"--url", srv.URL + "/ur-api-skills-v0.4.1.zip",
		"--output", outDir,
		"--json",
	}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit code = %d, stderr: %s", code, stderr.String())
	}

	var payload struct {
		Event       string `json:"event"`
		DownloadURL string `json:"downloadUrl"`
		LocalPath   string `json:"localPath"`
	}
	if err := json.Unmarshal(bytes.TrimSpace(stdout.Bytes()), &payload); err != nil {
		t.Fatalf("parse json output: %v\n%s", err, stdout.String())
	}
	if payload.Event != "skills_downloaded" {
		t.Errorf("event = %q, want skills_downloaded", payload.Event)
	}
	if payload.DownloadURL != srv.URL+"/ur-api-skills-v0.4.1.zip" {
		t.Errorf("downloadUrl = %q", payload.DownloadURL)
	}
	// 解压产物中 ur-api/SKILL.md 应存在
	if _, err := os.Stat(filepath.Join(payload.LocalPath, "SKILL.md")); err != nil {
		t.Errorf("extracted ur-api/SKILL.md should exist: %v", err)
	}

	// 非 JSON 输出：含下载来源与本地路径
	var stdout2, stderr2 bytes.Buffer
	code = runSkillsDownload([]string{
		"--url", srv.URL + "/ur-api-skills-v0.4.1.zip",
		"--output", outDir,
	}, &stdout2, &stderr2)
	if code != 0 {
		t.Fatalf("exit code = %d, stderr: %s", code, stderr2.String())
	}
	for _, want := range []string{"下载来源", "本地解压路径", "拷贝"} {
		if !bytes.Contains(stdout2.Bytes(), []byte(want)) {
			t.Errorf("output should contain %q, got:\n%s", want, stdout2.String())
		}
	}
}

// TestRunSkillsDownload_BadZip zip 内容缺少 ur-api/SKILL.md 时应报错退出
func TestRunSkillsDownload_BadZip(t *testing.T) {
	zipPath := makeSkillsZip(t, false)
	srv := httptest.NewServer(http.FileServer(http.Dir(filepath.Dir(zipPath))))
	defer srv.Close()

	var stdout, stderr bytes.Buffer
	code := runSkillsDownload([]string{
		"--url", srv.URL + "/ur-api-skills-v0.4.1.zip",
		"--output", filepath.Join(t.TempDir(), "dl"),
	}, &stdout, &stderr)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr: %s", code, stderr.String())
	}
}

// TestRunSkillsDownload_ArgErrors 参数缺失/未知参数返回错误码 2
func TestRunSkillsDownload_ArgErrors(t *testing.T) {
	cases := [][]string{
		{"--url"},
		{"--output"},
		{"--unknown"},
	}
	for _, args := range cases {
		var stdout, stderr bytes.Buffer
		if code := runSkillsDownload(args, &stdout, &stderr); code != 2 {
			t.Errorf("runSkillsDownload(%v) exit code = %d, want 2", args, code)
		}
	}
}

// writeSharedTestFile 创建命令层测试所需的文件及父目录。
func writeSharedTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

// TestRunSkillsTargetLifecycle 校验命令层的目标登记、列表和删除流程。
func TestRunSkillsTargetLifecycle(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_SKILL_TARGETS_FILE", filepath.Join(home, ".ur", "skill-targets.json"))
	targetDir := filepath.Join(home, "desktop-ai", "skills")

	var stdout, stderr bytes.Buffer
	if code := runSkillsTarget([]string{"add", "desktop-ai", "--dir", targetDir, "--json"}, &stdout, &stderr); code != 0 {
		t.Fatalf("target add code=%d stderr=%s", code, stderr.String())
	}
	stdout.Reset()
	stderr.Reset()
	if code := runSkillsTarget([]string{"list", "--json"}, &stdout, &stderr); code != 0 {
		t.Fatalf("target list code=%d stderr=%s", code, stderr.String())
	}
	var listed struct {
		Targets []struct {
			Name string `json:"name"`
			Path string `json:"path"`
		} `json:"targets"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &listed); err != nil {
		t.Fatalf("parse target list: %v\n%s", err, stdout.String())
	}
	if len(listed.Targets) != 1 || listed.Targets[0].Name != "desktop-ai" || listed.Targets[0].Path != targetDir {
		t.Fatalf("unexpected target list: %+v", listed.Targets)
	}
	stdout.Reset()
	stderr.Reset()
	if code := runSkillsTarget([]string{"remove", "desktop-ai", "--json"}, &stdout, &stderr); code != 0 {
		t.Fatalf("target remove code=%d stderr=%s", code, stderr.String())
	}
}

// TestRunSkillsInstallStatusAndExport 校验任意目录安装、完整性诊断和 ZIP 导出命令闭环。
func TestRunSkillsInstallStatusAndExport(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("UR_SKILL_TARGETS_FILE", filepath.Join(home, ".ur", "skill-targets.json"))
	source := filepath.Join(home, ".ur", "skills")
	writeSharedTestFile(t, filepath.Join(source, "SKILL.md"), "---\nname: ur-api\ndescription: test\n---\n")
	writeSharedTestFile(t, filepath.Join(source, "_meta.json"), `{"version":"v9.9.9"}`)
	writeSharedTestFile(t, filepath.Join(source, "ur-view", "SKILL.md"), "---\nname: ur-view\n---\n")
	targetDir := filepath.Join(home, "client", "skills")

	var stdout, stderr bytes.Buffer
	if code := runSkillsTarget([]string{"add", "client", "--dir", targetDir}, &stdout, &stderr); code != 0 {
		t.Fatalf("target add code=%d stderr=%s", code, stderr.String())
	}
	stdout.Reset()
	stderr.Reset()
	if code := runSkillsInstall([]string{"--all", "--json"}, &stdout, &stderr); code != 0 {
		t.Fatalf("install code=%d stderr=%s stdout=%s", code, stderr.String(), stdout.String())
	}
	stdout.Reset()
	stderr.Reset()
	if code := runSkillsStatus([]string{"--json"}, &stdout, &stderr); code != 0 {
		t.Fatalf("status code=%d stderr=%s stdout=%s", code, stderr.String(), stdout.String())
	}
	if !bytes.Contains(stdout.Bytes(), []byte(`"state": "current"`)) {
		t.Fatalf("status should be current: %s", stdout.String())
	}
	output := filepath.Join(t.TempDir(), "skills.zip")
	stdout.Reset()
	stderr.Reset()
	if code := runSkillsExport([]string{"--output", output, "--json"}, &stdout, &stderr); code != 0 {
		t.Fatalf("export code=%d stderr=%s stdout=%s", code, stderr.String(), stdout.String())
	}
	if _, err := os.Stat(output); err != nil {
		t.Fatalf("export output missing: %v", err)
	}
}
