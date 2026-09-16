// status.go — 检查各客户端中的 ur-api 是否与 CLI 内置 Skills 完整一致。
package skillinstall

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

const (
	// StatusCurrent 表示目标内容与当前内置 Skills 完全一致。
	StatusCurrent = "current"
	// StatusMissing 表示目标尚未安装 ur-api。
	StatusMissing = "missing"
	// StatusOutdated 表示目标版本与内置 Skills 版本不同。
	StatusOutdated = "outdated"
	// StatusIncomplete 表示目标版本相同但文件缺失、变化或存在残留。
	StatusIncomplete = "incomplete"
)

// TargetStatus 描述一个 Skills 目标的版本和完整性状态。
type TargetStatus struct {
	// Name 是目标稳定名称。
	Name string `json:"name"`
	// Kind 是客户端类型。
	Kind string `json:"kind"`
	// Scope 是用户级、项目级或自定义作用域。
	Scope string `json:"scope"`
	// Origin 是目标来源。
	Origin string `json:"origin"`
	// Path 是客户端 Skills 根目录。
	Path string `json:"path"`
	// Destination 是 ur-api 的实际安装目录。
	Destination string `json:"destination"`
	// State 是 current、missing、outdated 或 incomplete。
	State string `json:"state"`
	// Version 是目标安装版本。
	Version string `json:"version,omitempty"`
	// ExpectedVersion 是 CLI 内置 Skills 版本。
	ExpectedVersion string `json:"expectedVersion,omitempty"`
	// FileCount 是目标文件数量。
	FileCount int `json:"fileCount"`
	// ExpectedFileCount 是内置 Skills 文件数量。
	ExpectedFileCount int `json:"expectedFileCount"`
	// MissingFiles 是目标缺失的文件数量。
	MissingFiles int `json:"missingFiles,omitempty"`
	// ChangedFiles 是内容与内置版本不同的文件数量。
	ChangedFiles int `json:"changedFiles,omitempty"`
	// ExtraFiles 是目标中多出的残留文件数量。
	ExtraFiles int `json:"extraFiles,omitempty"`
	// Error 是检查目标时遇到的错误。
	Error string `json:"error,omitempty"`
}

// StatusResult 汇总内置版本与所有目标检查结果。
type StatusResult struct {
	// Source 是 CLI 内置 Skills 路径。
	Source string `json:"source"`
	// Version 是 CLI 内置 Skills 版本。
	Version string `json:"version"`
	// Targets 是逐目标检查结果。
	Targets []TargetStatus `json:"targets"`
}

// InspectTargets 比较内置 Skills 与各目标中的 ur-api，定位旧版和不完整副本。
func InspectTargets(src string, targets []Target) (*StatusResult, error) {
	sourceFiles, err := treeSignatures(src)
	if err != nil {
		return nil, fmt.Errorf("读取内置 Skills 失败: %w", err)
	}
	result := &StatusResult{Source: src, Version: ReadVersion(src)}
	for _, target := range targets {
		destination := filepath.Join(target.Path, "ur-api")
		status := TargetStatus{
			Name:              target.Name,
			Kind:              target.Kind,
			Scope:             target.Scope,
			Origin:            target.Origin,
			Path:              target.Path,
			Destination:       destination,
			State:             StatusMissing,
			ExpectedVersion:   result.Version,
			ExpectedFileCount: len(sourceFiles),
		}
		targetFiles, targetErr := treeSignatures(destination)
		if os.IsNotExist(targetErr) {
			result.Targets = append(result.Targets, status)
			continue
		}
		if targetErr != nil {
			status.Error = targetErr.Error()
			result.Targets = append(result.Targets, status)
			continue
		}
		status.Version = ReadVersion(destination)
		status.FileCount = len(targetFiles)
		for path, signature := range sourceFiles {
			got, ok := targetFiles[path]
			if !ok {
				status.MissingFiles++
				continue
			}
			if got != signature {
				status.ChangedFiles++
			}
		}
		for path := range targetFiles {
			if _, ok := sourceFiles[path]; !ok {
				status.ExtraFiles++
			}
		}
		switch {
		case status.Version != result.Version:
			status.State = StatusOutdated
		case status.MissingFiles > 0 || status.ChangedFiles > 0 || status.ExtraFiles > 0:
			status.State = StatusIncomplete
		default:
			status.State = StatusCurrent
		}
		result.Targets = append(result.Targets, status)
	}
	return result, nil
}

// ReadVersion 读取 ur-api 根目录中的版本元数据。
func ReadVersion(root string) string {
	raw, err := os.ReadFile(filepath.Join(root, "_meta.json"))
	if err != nil {
		return "unknown"
	}
	var metadata struct {
		Version string `json:"version"`
	}
	if json.Unmarshal(raw, &metadata) != nil || metadata.Version == "" {
		return "unknown"
	}
	return metadata.Version
}

// readVersion 保留包内旧调用入口，统一委托给公开的 ReadVersion。
func readVersion(root string) string {
	return ReadVersion(root)
}

// treeSignatures 计算目录下每个普通文件的 SHA-256，用于完整性对比。
func treeSignatures(root string) (map[string]string, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s 不是目录", root)
	}
	signatures := make(map[string]string)
	err = filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("不支持的 Skills 文件类型: %s", path)
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		hash := sha256.New()
		_, copyErr := io.Copy(hash, file)
		closeErr := file.Close()
		if copyErr != nil {
			return copyErr
		}
		if closeErr != nil {
			return closeErr
		}
		signatures[filepath.ToSlash(relative)] = hex.EncodeToString(hash.Sum(nil))
		return nil
	})
	return signatures, err
}

// SortTargetStatuses 按名称和路径稳定排序状态结果，便于人类阅读和自动化比较。
func SortTargetStatuses(statuses []TargetStatus) {
	sort.Slice(statuses, func(i, j int) bool {
		if statuses[i].Name == statuses[j].Name {
			return statuses[i].Path < statuses[j].Path
		}
		return statuses[i].Name < statuses[j].Name
	})
}
