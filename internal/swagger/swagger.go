package swagger

import (
	"errors"
	"os"
	"path/filepath"
)

var requiredFiles = []string{"core-api.json", "things-api.json", "core-ai.json", "things-ai.json"}

func ResolveFiles() ([]string, error) {
	roots := candidateRoots()
	for _, root := range roots {
		files := make([]string, 0, len(requiredFiles))
		ok := true
		for _, name := range requiredFiles {
			path := filepath.Join(root, name)
			if _, err := os.Stat(path); err != nil {
				ok = false
				break
			}
			files = append(files, path)
		}
		if ok {
			return files, nil
		}
	}
	return nil, errors.New("swagger files not found")
}

func candidateRoots() []string {
	var roots []string
	if env := os.Getenv("UR_SWAGGER_DIR"); env != "" {
		roots = append(roots, env)
	}
	roots = append(roots, "/opt/backend/.swagger")
	if wd, err := os.Getwd(); err == nil {
		dir := wd
		for {
			candidate := filepath.Join(dir, "backend", ".swagger")
			roots = append(roots, candidate)
			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
			dir = parent
		}
	}
	return dedupe(roots)
}

func dedupe(items []string) []string {
	seen := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	for _, item := range items {
		if item == "" {
			continue
		}
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}
