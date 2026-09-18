package swagger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveSwaggerFilesPrefersConfiguredRoot(t *testing.T) {
	root := t.TempDir()
	swaggerDir := filepath.Join(root, ".swagger")
	if err := os.MkdirAll(swaggerDir, 0o755); err != nil {
		t.Fatalf("mkdir swagger: %v", err)
	}
	// 与 requiredFiles 保持一致:清单扩充后测试同步写入全部必需文件。
	for _, name := range requiredFiles {
		if err := os.WriteFile(filepath.Join(swaggerDir, name), []byte(`{"openapi":"3.0.0","paths":{}}`), 0o644); err != nil {
			t.Fatalf("write swagger %s: %v", name, err)
		}
	}
	t.Setenv("UR_SWAGGER_DIR", swaggerDir)

	files, err := ResolveFiles()
	if err != nil {
		t.Fatalf("ResolveFiles error: %v", err)
	}
	if len(files) != len(requiredFiles) {
		t.Fatalf("file count = %d, want %d", len(files), len(requiredFiles))
	}
}
