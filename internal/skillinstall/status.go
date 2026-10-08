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
	"strings"
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
	// StatusDuplicate 表示 ur-api 本身正常，但目标 skills 目录内存在
	// ur-api 备份/旧版等同名前缀残留目录，会被 AI 客户端重复注册为技能。
	StatusDuplicate = "duplicate"
)

// duplicatePrefix 是残留目录识别前缀：目录名以 ur-api 开头但不是 ur-api
// 本身（如 ur-api.v0.4.1.bak、ur-api.ur-bak、ur-api-old）即视为残留。
const duplicatePrefix = "ur-api"

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
	// DuplicateDirs 是目标 skills 目录内检测到的 ur-api 残留目录完整路径，
	// 来源多为手工升级时的备份残留，会被 AI 客户端重复注册为技能。
	DuplicateDirs []string `json:"duplicateDirs,omitempty"`
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
	// 源侧按安装形态计算签名：子域 SKILL.md 转换为 GUIDE.md 后再比较，
	// install 写入目标的内容与此口径一致，status 才能判定 current。
	sourceFiles, err := treeSignatures(src, true)
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
		// 残留检测对所有状态生效：即使 ur-api 本身正常（甚至缺失），
		// 同级的备份目录也会被 AI 客户端注册成重复技能。
		status.DuplicateDirs = FindDuplicateDirs(target.Path)
		targetFiles, targetErr := treeSignatures(destination, false)
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
		if status.State == StatusCurrent && len(status.DuplicateDirs) > 0 {
			status.State = StatusDuplicate
		}
		result.Targets = append(result.Targets, status)
	}
	return result, nil
}

// FindDuplicateDirs 返回 targetPath 下所有 ur-api 前缀的残留目录完整路径。
// AI 客户端会注册 skills 目录下所有含 SKILL.md 的子目录（不识别 .bak 等后缀），
// 因此手工升级留下的 ur-api.v0.4.1.bak 之类备份会被识别为第二个 ur-api 技能。
func FindDuplicateDirs(targetPath string) []string {
	entries, err := os.ReadDir(targetPath)
	if err != nil {
		return nil
	}
	var duplicates []string
	for _, entry := range entries {
		name := entry.Name()
		if !entry.IsDir() || name == duplicatePrefix || !strings.HasPrefix(name, duplicatePrefix) {
			continue
		}
		duplicates = append(duplicates, filepath.Join(targetPath, name))
	}
	return duplicates
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
// applyAggregate 为 true 时先按安装形态转换（子域 SKILL.md → GUIDE.md 并
// 剥离 frontmatter），使源侧签名与 install 写入目标的内容可直接比较。
func treeSignatures(root string, applyAggregate bool) (map[string]string, error) {
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
		installedKey := filepath.ToSlash(relative)
		if applyAggregate && isAggregatable(installedKey) {
			data, readErr := os.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			transformedKey, content := transformAggregate(installedKey, data)
			sum := sha256.Sum256(content)
			signatures[transformedKey] = hex.EncodeToString(sum[:])
			return nil
		}
		signatures[installedKey] = hex.EncodeToString(hash.Sum(nil))
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
