// layout_test.go 验证五组技能的安装升级、失败回滚、状态诊断与导出完整性。
package skillinstall

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// makeGroupedSource 构造包含各层元信息和嵌套资产的五组技能源，返回源目录。
func makeGroupedSource(t *testing.T) string {
	t.Helper()
	source := makeFakeSource(t)
	for _, group := range []string{"ur-iot", "ur-org-manage", "ur-ai", "ur-doc"} {
		mustWrite(t, filepath.Join(source, group, "SKILL.md"), "---\nname: "+group+"\n---\n# 业务组\n")
	}
	mustWrite(t, filepath.Join(source, "ur-iot", "ur-device", "SKILL.md"), "---\nname: ur-device\ndescription: 设备管理\n---\n# 设备技能\n")
	mustWrite(t, filepath.Join(source, "ur-iot", "ur-device", "references", "api", "control.md"), "# 控制接口\n")
	mustWrite(t, filepath.Join(source, "ur-iot", "device-firmware", "assets", "firmware.bin"), "firmware")
	return source
}

// TestGroupedInstallUpgrade 验证新装、旧平铺和旧 GUIDE 安装均收敛到源结构且不影响其他技能。
func TestGroupedInstallUpgrade(t *testing.T) {
	source := makeGroupedSource(t)
	for _, legacy := range []string{"", "ur-device/SKILL.md", "ur-device/GUIDE.md", "ur-iot/ur-device/GUIDE.md"} {
		t.Run(legacy, func(t *testing.T) {
			target := Target{Name: "client", Path: t.TempDir()}
			if legacy != "" {
				mustWrite(t, filepath.Join(target.Path, "ur-api", legacy), "旧内容")
			}
			mustWrite(t, filepath.Join(target.Path, "other-skill", "SKILL.md"), "其他技能")
			for range 2 {
				result, err := Install(source, []Target{target}, false)
				if err != nil || HasErrors(result) {
					t.Fatalf("install: %+v %v", result, err)
				}
				status, err := InspectTargets(source, []Target{target})
				if err != nil || status.Targets[0].State != StatusCurrent {
					t.Fatalf("status: %+v %v", status, err)
				}
				entry := filepath.Join(target.Path, "ur-api", "ur-iot", "ur-device", "SKILL.md")
				actual, err := os.ReadFile(entry)
				expected, _ := os.ReadFile(filepath.Join(source, "ur-iot", "ur-device", "SKILL.md"))
				if err != nil || string(actual) != string(expected) {
					t.Fatalf("技能元信息或正文改变: %v", err)
				}
				for _, stale := range []string{"ur-device", "ur-iot/ur-device/GUIDE.md"} {
					if _, err := os.Stat(filepath.Join(target.Path, "ur-api", stale)); !os.IsNotExist(err) {
						t.Fatalf("旧入口残留: %s", stale)
					}
				}
				if content, _ := os.ReadFile(filepath.Join(target.Path, "other-skill", "SKILL.md")); string(content) != "其他技能" {
					t.Fatal("其他技能被修改")
				}
			}
		})
	}
}

// TestGroupedInstallRollback 验证源在备份后不可读时，旧版目录及其入口会被恢复。
func TestGroupedInstallRollback(t *testing.T) {
	target := Target{Name: "client", Path: t.TempDir()}
	source := filepath.Join(target.Path, "ur-api", "embedded-source")
	mustWrite(t, filepath.Join(source, "SKILL.md"), "待安装源")
	old := filepath.Join(target.Path, "ur-api", "ur-device", "GUIDE.md")
	mustWrite(t, old, "旧技能")
	result, err := Install(source, []Target{target}, false)
	if err != nil || !HasErrors(result) {
		t.Fatalf("预期复制失败: %+v %v", result, err)
	}
	if content, err := os.ReadFile(old); err != nil || string(content) != "旧技能" {
		t.Fatalf("旧技能未恢复: %v", err)
	}
}

// TestGroupedExportZIP 验证各层 SKILL.md、元信息和深层资源按源字节导出。
func TestGroupedExportZIP(t *testing.T) {
	source := makeGroupedSource(t)
	output := filepath.Join(t.TempDir(), "skills.zip")
	if _, err := ExportZIP(source, output); err != nil {
		t.Fatal(err)
	}
	archive, err := zip.OpenReader(output)
	if err != nil {
		t.Fatal(err)
	}
	defer archive.Close()
	files := map[string]*zip.File{}
	for _, file := range archive.File {
		files[file.Name] = file
	}
	for _, relative := range []string{"SKILL.md", "ur-iot/SKILL.md", "ur-org-manage/SKILL.md", "ur-ai/SKILL.md", "ur-view/SKILL.md", "ur-doc/SKILL.md", "ur-iot/ur-device/SKILL.md", "ur-iot/ur-device/references/api/control.md", "ur-iot/device-firmware/assets/firmware.bin"} {
		file, ok := files["ur-api/"+relative]
		if !ok {
			t.Fatalf("导出漏文件: %s", relative)
		}
		reader, err := file.Open()
		if err != nil {
			t.Fatal(err)
		}
		actual, err := io.ReadAll(reader)
		reader.Close()
		expected, _ := os.ReadFile(filepath.Join(source, relative))
		if err != nil || string(actual) != string(expected) {
			t.Fatalf("导出内容改变: %s %v", relative, err)
		}
	}
}

// TestCopySourceRejectsOverlap 防止生成器向源内部递归复制或覆盖源目录。
func TestCopySourceRejectsOverlap(t *testing.T) {
	source := makeGroupedSource(t)
	for _, output := range []string{source, filepath.Join(source, "generated"), filepath.Dir(source)} {
		if err := CopySource(source, output); err == nil {
			t.Fatalf("接受重叠目录: %s", output)
		}
	}
}
