// skillsdownload_test.go — skills 独立资产匹配与下载解压流程单元测试（不依赖外网）
package upgrade

import (
	"archive/zip"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// TestSkillsAssetName 校验资产命名与 release.sh 打包一致
func TestSkillsAssetName(t *testing.T) {
	if got := SkillsAssetName("v0.4.1"); got != "ur-api-skills-v0.4.1.zip" {
		t.Errorf("SkillsAssetName = %q, want ur-api-skills-v0.4.1.zip", got)
	}
}

// TestFindSkillsAsset 校验精确匹配、前缀兜底与未命中三种情况
func TestFindSkillsAsset(t *testing.T) {
	release := &Release{
		TagName: "v0.4.1",
		Assets: []Asset{
			{Name: "ur-cli-v0.4.1-Linux-x86_64.tar.gz", BrowserDownloadURL: "https://example.com/ur-cli-v0.4.1-Linux-x86_64.tar.gz"},
			{Name: "ur-api-skills-v0.4.1.zip", BrowserDownloadURL: "https://example.com/ur-api-skills-v0.4.1.zip"},
		},
	}
	asset := release.FindSkillsAsset()
	if asset == nil {
		t.Fatal("should find skills asset")
	}
	if asset.BrowserDownloadURL != "https://example.com/ur-api-skills-v0.4.1.zip" {
		t.Errorf("wrong asset url: %s", asset.BrowserDownloadURL)
	}

	// 前缀兜底：资产名版本与 tag 有细微差异时仍可命中
	release.TagName = "v0.4.2"
	asset = release.FindSkillsAsset()
	if asset == nil || asset.Name != "ur-api-skills-v0.4.1.zip" {
		t.Errorf("prefix fallback failed, got %+v", asset)
	}

	// 未命中
	empty := &Release{TagName: "v0.4.1"}
	if empty.FindSkillsAsset() != nil {
		t.Error("should return nil when no skills asset")
	}
}

// makeURSkillsZip 构造最小 skills zip（顶层 ur-api/ 含 SKILL.md），返回所在目录与文件名
func makeURSkillsZip(t *testing.T) (dir, name string) {
	t.Helper()
	dir = t.TempDir()
	name = "ur-api-skills-v0.4.1.zip"
	p := filepath.Join(dir, name)
	f, err := os.Create(p)
	if err != nil {
		t.Fatalf("create zip: %v", err)
	}
	defer f.Close()
	zw := zip.NewWriter(f)
	for _, entry := range []struct{ name, content string }{
		{"ur-api/SKILL.md", "---\nname: ur-api\n---\n"},
		{"ur-api/_meta.json", `{"version":"v0.4.1"}`},
	} {
		w, err := zw.Create(entry.name)
		if err != nil {
			t.Fatalf("zip create %s: %v", entry.name, err)
		}
		if _, err := w.Write([]byte(entry.content)); err != nil {
			t.Fatalf("zip write %s: %v", entry.name, err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatalf("zip close: %v", err)
	}
	return dir, name
}

// TestDownloadSkills_CustomURL 用 httptest 作为 --url 下载源，校验全流程不依赖外网
func TestDownloadSkills_CustomURL(t *testing.T) {
	zipDir, zipName := makeURSkillsZip(t)
	srv := httptest.NewServer(http.FileServer(http.Dir(zipDir)))
	defer srv.Close()

	outDir := filepath.Join(t.TempDir(), "downloads")
	res, err := DownloadSkills(SkillsDownloadOptions{
		OutputDir: outDir,
		ZipURL:    srv.URL + "/" + zipName,
	})
	if err != nil {
		t.Fatalf("DownloadSkills: %v", err)
	}
	// 版本应从 zip 文件名反推
	if res.Version != "v0.4.1" {
		t.Errorf("version = %q, want v0.4.1", res.Version)
	}
	if res.DownloadURL != srv.URL+"/"+zipName {
		t.Errorf("downloadUrl = %q", res.DownloadURL)
	}
	// 解压产物完整
	for _, f := range []string{"SKILL.md", "_meta.json"} {
		if _, err := os.Stat(filepath.Join(res.LocalPath, f)); err != nil {
			t.Errorf("extracted ur-api/%s should exist: %v", f, err)
		}
	}
}

// TestDownloadSkills_BadZip 缺少 ur-api/SKILL.md 的 zip 应报错且附带下载来源
func TestDownloadSkills_BadZip(t *testing.T) {
	zipDir := t.TempDir()
	p := filepath.Join(zipDir, "ur-api-skills-v0.4.1.zip")
	f, err := os.Create(p)
	if err != nil {
		t.Fatal(err)
	}
	zw := zip.NewWriter(f)
	w, _ := zw.Create("random.txt")
	w.Write([]byte("x"))
	zw.Close()
	f.Close()

	srv := httptest.NewServer(http.FileServer(http.Dir(zipDir)))
	defer srv.Close()

	_, err = DownloadSkills(SkillsDownloadOptions{
		OutputDir: filepath.Join(t.TempDir(), "dl"),
		ZipURL:    srv.URL + "/ur-api-skills-v0.4.1.zip",
	})
	if err == nil {
		t.Fatal("bad zip should error")
	}
	if !contains(err.Error(), "ur-api") {
		t.Errorf("error should mention ur-api dir, got: %v", err)
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
