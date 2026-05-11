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
	for _, name := range []string{"core-api.json", "things-api.json"} {
		if err := os.WriteFile(filepath.Join(swaggerDir, name), []byte(`{"openapi":"3.0.0","paths":{}}`), 0o644); err != nil {
			t.Fatalf("write swagger %s: %v", name, err)
		}
	}
	t.Setenv("UR_SWAGGER_DIR", swaggerDir)

	files, err := ResolveFiles()
	if err != nil {
		t.Fatalf("ResolveFiles error: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("file count = %d, want 2", len(files))
	}
}
