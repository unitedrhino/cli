// Package skillscheck 以轻量方式检查各 AI 客户端中的 ur-api 版本。
package skillscheck

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gitee.com/unitedrhino/cli/internal/notice"
	"gitee.com/unitedrhino/cli/internal/skillinstall"
	"gitee.com/unitedrhino/cli/internal/upgrade"
)

// LoadNotice 读取目标版本元数据并为缺失或落后的客户端准备提示。
func LoadNotice(cwd string) {
	if envEnabled("UR_NO_SKILLS_NOTIFIER") {
		return
	}
	source := upgrade.GetDefaultSkillsDir()
	targetVersion := skillinstall.ReadVersion(source)
	if source == "" || targetVersion == "unknown" {
		return
	}
	targets, err := skillinstall.ResolveTargets(cwd)
	if err != nil {
		return
	}
	staleNames := make([]string, 0)
	currentVersions := make(map[string]struct{})
	for _, target := range targets {
		destination := filepath.Join(target.Path, "ur-api")
		current := skillinstall.ReadVersion(destination)
		if current == targetVersion {
			continue
		}
		if _, err := os.Stat(destination); os.IsNotExist(err) {
			current = "missing"
		} else if current == "unknown" {
			// 现有目录没有版本元数据时无法判断新旧，不把自维护目录误报为落后。
			continue
		} else if !upgrade.IsNewer(current, targetVersion) {
			// 客户端副本比当前 CLI 内置版本更新时不提示，避免旧 CLI 诱导降级。
			continue
		}
		staleNames = append(staleNames, target.Name)
		currentVersions[current] = struct{}{}
	}
	if len(staleNames) == 0 {
		return
	}
	sort.Strings(staleNames)
	notice.SetSkills(notice.Skills{
		Current: summarizeVersions(currentVersions),
		Target:  targetVersion,
		Targets: staleNames,
		Message: fmt.Sprintf("%d 个 AI Skills 目标需要更新到 %s", len(staleNames), targetVersion),
		Command: "ur upgrade",
	})
}

// summarizeVersions 把单一版本直接返回，多版本则返回 mixed，便于机器判断。
func summarizeVersions(versions map[string]struct{}) string {
	if len(versions) != 1 {
		return "mixed"
	}
	for version := range versions {
		return version
	}
	return "unknown"
}

// envEnabled 解析 Skills 通知关闭开关。
func envEnabled(name string) bool {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return false
	}
	if enabled, err := strconv.ParseBool(value); err == nil {
		return enabled
	}
	return value != "0"
}
