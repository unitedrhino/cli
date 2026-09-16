// targets.go — 管理可复用的 Skills 安装目标，并合并自动探测与用户配置。
package skillinstall

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	// TargetTypeFilesystem 表示目标是本机可写的 Skills 根目录。
	TargetTypeFilesystem = "filesystem"
	// targetsConfigVersion 是目标配置文件格式版本。
	targetsConfigVersion = 1
)

// ConfiguredTarget 是用户登记的可复用 Skills 安装目标。
type ConfiguredTarget struct {
	// Name 是目标的唯一名称。
	Name string `json:"name"`
	// Type 是目标类型，当前支持 filesystem。
	Type string `json:"type"`
	// Path 是客户端的 Skills 根目录。
	Path string `json:"path"`
	// Enabled 控制目标是否参与批量安装和状态检查。
	Enabled bool `json:"enabled"`
}

// TargetsConfig 是 ~/.ur/skill-targets.json 的持久化结构。
type TargetsConfig struct {
	// Version 是配置文件格式版本。
	Version int `json:"version"`
	// Targets 是用户登记的目标列表。
	Targets []ConfiguredTarget `json:"targets"`
}

// TargetsConfigPath 返回 Skills 目标配置路径。
// UR_SKILL_TARGETS_FILE 主要供自动化环境和测试隔离默认配置。
func TargetsConfigPath() (string, error) {
	if path := strings.TrimSpace(os.Getenv("UR_SKILL_TARGETS_FILE")); path != "" {
		return normalizePath(path), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("无法获取用户目录: %w", err)
	}
	return filepath.Join(home, ".ur", "skill-targets.json"), nil
}

// LoadConfiguredTargets 读取用户登记的安装目标；配置尚未创建时返回空列表。
func LoadConfiguredTargets() ([]ConfiguredTarget, error) {
	path, err := TargetsConfigPath()
	if err != nil {
		return nil, err
	}
	raw, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("读取 Skills 目标配置失败: %w", err)
	}
	var config TargetsConfig
	if err := json.Unmarshal(raw, &config); err != nil {
		return nil, fmt.Errorf("解析 Skills 目标配置失败: %w", err)
	}
	if config.Version != targetsConfigVersion {
		return nil, fmt.Errorf("不支持的 Skills 目标配置版本 %d", config.Version)
	}
	return config.Targets, nil
}

// SaveConfiguredTargets 原子写入用户登记的安装目标。
func SaveConfiguredTargets(targets []ConfiguredTarget) error {
	path, err := TargetsConfigPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("创建 Skills 目标配置目录失败: %w", err)
	}
	normalized := append([]ConfiguredTarget(nil), targets...)
	sort.Slice(normalized, func(i, j int) bool { return normalized[i].Name < normalized[j].Name })
	raw, err := json.MarshalIndent(TargetsConfig{Version: targetsConfigVersion, Targets: normalized}, "", "  ")
	if err != nil {
		return fmt.Errorf("编码 Skills 目标配置失败: %w", err)
	}
	raw = append(raw, '\n')
	temp, err := os.CreateTemp(filepath.Dir(path), ".skill-targets-*.tmp")
	if err != nil {
		return fmt.Errorf("创建 Skills 目标临时配置失败: %w", err)
	}
	tempPath := temp.Name()
	defer os.Remove(tempPath)
	if err := temp.Chmod(0o600); err != nil {
		temp.Close()
		return fmt.Errorf("设置 Skills 目标配置权限失败: %w", err)
	}
	if _, err := temp.Write(raw); err != nil {
		temp.Close()
		return fmt.Errorf("写入 Skills 目标配置失败: %w", err)
	}
	if err := temp.Close(); err != nil {
		return fmt.Errorf("关闭 Skills 目标配置失败: %w", err)
	}
	if err := replaceFileAtomically(tempPath, path); err != nil {
		return fmt.Errorf("替换 Skills 目标配置失败: %w", err)
	}
	return nil
}

// replaceFileAtomically 通过同目录备份替换文件，并在替换失败时恢复旧文件。
func replaceFileAtomically(source, destination string) error {
	backup := destination + ".ur-bak"
	if err := os.Remove(backup); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	_, statErr := os.Stat(destination)
	exists := statErr == nil
	if statErr != nil && !errors.Is(statErr, os.ErrNotExist) {
		return statErr
	}
	if exists {
		if err := os.Rename(destination, backup); err != nil {
			return err
		}
	}
	if err := os.Rename(source, destination); err != nil {
		if exists {
			_ = os.Rename(backup, destination)
		}
		return err
	}
	if exists {
		_ = os.Remove(backup)
	}
	return nil
}

// UpsertConfiguredTarget 新增或更新一个文件系统目标，并返回是否覆盖已有名称。
func UpsertConfiguredTarget(target ConfiguredTarget) (bool, error) {
	target.Name = strings.TrimSpace(target.Name)
	if target.Name == "" {
		return false, errors.New("目标名称不能为空")
	}
	if strings.ContainsAny(target.Name, "/\\") {
		return false, errors.New("目标名称不能包含路径分隔符")
	}
	if target.Type == "" {
		target.Type = TargetTypeFilesystem
	}
	if target.Type != TargetTypeFilesystem {
		return false, fmt.Errorf("不支持的目标类型 %q，当前仅支持 filesystem", target.Type)
	}
	target.Path = normalizePath(target.Path)
	if target.Path == "" {
		return false, errors.New("目标目录不能为空")
	}
	target.Enabled = true
	targets, err := LoadConfiguredTargets()
	if err != nil {
		return false, err
	}
	updated := false
	for index := range targets {
		if targets[index].Name == target.Name {
			targets[index] = target
			updated = true
			break
		}
	}
	if !updated {
		targets = append(targets, target)
	}
	return updated, SaveConfiguredTargets(targets)
}

// RemoveConfiguredTarget 删除指定名称的用户目标，并返回目标是否存在。
func RemoveConfiguredTarget(name string) (bool, error) {
	targets, err := LoadConfiguredTargets()
	if err != nil {
		return false, err
	}
	filtered := targets[:0]
	found := false
	for _, target := range targets {
		if target.Name == name {
			found = true
			continue
		}
		filtered = append(filtered, target)
	}
	if !found {
		return false, nil
	}
	return true, SaveConfiguredTargets(filtered)
}

// ResolveTargets 合并用户登记目标、环境变量目标和已知客户端自动探测结果。
// 用户登记目标优先，以便同一路径保留用户指定的稳定名称。
func ResolveTargets(cwd string) ([]Target, error) {
	configured, err := LoadConfiguredTargets()
	if err != nil {
		return nil, err
	}
	resolved := make([]Target, 0, len(configured)+8)
	for _, target := range configured {
		if !target.Enabled || target.Type != TargetTypeFilesystem {
			continue
		}
		resolved = append(resolved, Target{
			Name:   target.Name,
			Path:   target.Path,
			Scope:  "custom",
			Kind:   "custom",
			Origin: "configured",
		})
	}
	detected, err := DetectTargets(cwd)
	if err != nil {
		return nil, err
	}
	return DedupeTargets(append(resolved, detected...)), nil
}
