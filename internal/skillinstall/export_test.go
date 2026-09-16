// export_test.go — 验证标准 Skills ZIP 的结构、内容与路径保护。
package skillinstall

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

// TestExportZIP 校验导出包只有 ur-api 根目录并保留嵌套资源。
func TestExportZIP(t *testing.T) {
	source := makeFakeSource(t)
	mustWrite(t, filepath.Join(source, "ur-view", "assets", "scene-templates", "building", "index.html"), "building")
	output := filepath.Join(t.TempDir(), "portable.zip")
	result, err := ExportZIP(source, output)
	if err != nil {
		t.Fatalf("ExportZIP: %v", err)
	}
	if result.Path != output || result.Files != 4 || result.Bytes == 0 {
		t.Fatalf("unexpected export result: %+v", result)
	}
	if _, err := ExportZIP(source, output); err != nil {
		t.Fatalf("ExportZIP overwrite: %v", err)
	}
	reader, err := zip.OpenReader(output)
	if err != nil {
		t.Fatalf("open zip: %v", err)
	}
	defer reader.Close()
	names := map[string]bool{}
	for _, file := range reader.File {
		names[file.Name] = true
	}
	for _, want := range []string{
		"ur-api/SKILL.md",
		"ur-api/_meta.json",
		"ur-api/ur-view/SKILL.md",
		"ur-api/ur-view/assets/scene-templates/building/index.html",
	} {
		if !names[want] {
			t.Errorf("missing zip entry %s", want)
		}
	}
}

// TestExportZIP_OutputDirectory 校验目录输出会自动生成带版本的文件名。
func TestExportZIP_OutputDirectory(t *testing.T) {
	source := makeFakeSource(t)
	outputDir := t.TempDir()
	result, err := ExportZIP(source, outputDir)
	if err != nil {
		t.Fatalf("ExportZIP: %v", err)
	}
	want := filepath.Join(outputDir, "ur-api-skills-v0.4.0.zip")
	if result.Path != want {
		t.Fatalf("output path = %s, want %s", result.Path, want)
	}
	if _, err := os.Stat(want); err != nil {
		t.Fatalf("output missing: %v", err)
	}
}

// TestExportZIP_RejectsOutputInsideSource 防止输出文件被递归收入自身。
func TestExportZIP_RejectsOutputInsideSource(t *testing.T) {
	source := makeFakeSource(t)
	if _, err := ExportZIP(source, filepath.Join(source, "nested.zip")); err == nil {
		t.Fatal("output inside source should fail")
	}
}
