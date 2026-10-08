// generate_skills_test.go 验证跨应用生成保留五组业务技能和元信息，且只在独立目录生成索引。
package shared

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"gitee.com/unitedrhino/cli/internal/config"
)

// TestGenerateAllGroupedSkills 使用真实内置技能与 Swagger 验证重复生成后的入口和资源完整性。
func TestGenerateAllGroupedSkills(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	t.Chdir(root)
	output := filepath.Join(t.TempDir(), "ur-api")
	for range 2 {
		var stdout, stderr bytes.Buffer
		if code := runGenerateSkills(config.AppIoT, []string{"--all", "--output", output}, &stdout, &stderr); code != 0 {
			t.Fatalf("generate: %d %s", code, stderr.String())
		}
		for _, relative := range []string{"SKILL.md", "ur-iot/SKILL.md", "ur-iot/ur-device/SKILL.md", "ur-iot/device-firmware/references/voice-ai.md", "ur-org-manage/SKILL.md", "ur-org-manage/ur-user/SKILL.md", "ur-ai/ai-tool/SKILL.md", "ur-view/SKILL.md", "ur-doc/SKILL.md"} {
			actual, err := os.ReadFile(filepath.Join(output, relative))
			expected, sourceErr := os.ReadFile(filepath.Join(root, "skill", relative))
			if err != nil || sourceErr != nil || !bytes.Equal(actual, expected) {
				t.Fatalf("技能生成发生丢失或改写: %s %v %v", relative, err, sourceErr)
			}
		}
		for _, relative := range []string{"swagger-index.md", "references/things-index.md", "references/system-index.md"} {
			if _, err := os.Stat(filepath.Join(output, relative)); err != nil {
				t.Fatalf("缺少生成索引: %s %v", relative, err)
			}
		}
		if _, err := os.Stat(filepath.Join(output, "ur-device")); !os.IsNotExist(err) {
			t.Fatal("生成器重建平铺目录")
		}
	}
}
