// skillinstall — 把内置 ur-api skill 部署到各 AI 工具的 skills 目录
//
// 不同 AI 工具（Claude Code / Codex 等）存放 skills 的位置不同：
//   - Claude Code：~/.claude/skills/ 与项目 .claude/skills/
//   - Codex：~/.agents/skills/ 与项目 .agents/skills/（项目级常为软链目录）
//
// 部署单元是唯一的 ur-api 整个 skill（SKILL.md + 内部子域内容 + _meta.json），
// 整体拷贝覆盖到各目标的 ur-api/ 目录；只覆盖 ur-api，保留目标里其他 AI 自有 skill。
// Cursor（.cursor/rules）与 OpenCode 不在部署范围：ur-api 为大型 skill，不适用 rules 形态。
package skillinstall

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Target AI 工具的 skills 目标目录
type Target struct {
	// Path 目标 skills 根目录（ur-api 将被安装为 <Path>/ur-api/）
	Path string `json:"path"`
	// Scope 作用域：user（用户级）/ project（项目级）
	Scope string `json:"scope"`
	// Kind AI 工具类型：claude / codex
	Kind string `json:"kind"`
}

// TargetResult 单个目标的安装结果
type TargetResult struct {
	Path      string `json:"path"`
	Scope     string `json:"scope"`
	Kind      string `json:"kind"`
	Installed bool   `json:"installed"`
	Updated   bool   `json:"updated,omitempty"`
	Error     string `json:"error,omitempty"`
}

// Result 整体安装结果
type Result struct {
	// Source 内置 skills 源目录
	Source string        `json:"source"`
	Targets []TargetResult `json:"targets"`
}

// findRepoRoot 从 cwd 向上查找 git 仓库根目录（.git 所在目录）
func findRepoRoot(cwd string) string {
	dir := cwd
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

// DetectTargets 探测本机各 AI 工具的 skills 目标目录（仅返回实际存在的目录）
func DetectTargets(cwd string) ([]Target, error) {
	var targets []Target
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("无法获取用户目录: %w", err)
	}

	// 用户级
	if dirExists(filepath.Join(home, ".claude", "skills")) {
		targets = append(targets, Target{Path: filepath.Join(home, ".claude", "skills"), Scope: "user", Kind: "claude"})
	}
	if dirExists(filepath.Join(home, ".agents", "skills")) {
		targets = append(targets, Target{Path: filepath.Join(home, ".agents", "skills"), Scope: "user", Kind: "codex"})
	}

	// 项目级（当前 git 仓库根下）
	if root := findRepoRoot(cwd); root != "" {
		if dirExists(filepath.Join(root, ".claude", "skills")) {
			targets = append(targets, Target{Path: filepath.Join(root, ".claude", "skills"), Scope: "project", Kind: "claude"})
		}
		if dirExists(filepath.Join(root, ".agents", "skills")) {
			targets = append(targets, Target{Path: filepath.Join(root, ".agents", "skills"), Scope: "project", Kind: "codex"})
		}
	}
	return targets, nil
}

// dirExists 目录是否存在
func dirExists(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

// copyDir 递归复制目录（含空目录与文件权限）
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return os.MkdirAll(dst, 0o755)
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, info.Mode()|0o700)
		}
		// 只复制普通文件与软链（跳过其他特殊文件）
		if !info.Mode().IsRegular() && info.Mode()&os.ModeSymlink == 0 {
			return nil
		}
		if info.Mode()&os.ModeSymlink != 0 {
			link, err := os.Readlink(path)
			if err != nil {
				return err
			}
			return os.Symlink(link, target)
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, in)
		return err
	})
}

// Install 将内置 skills 源（整个 ur-api：SKILL.md + 子域 + _meta.json）整体拷贝覆盖
// 到各目标的 ur-api/ 目录。只覆盖 ur-api，保留目标内其他 AI 自有 skill；幂等。
// dryRun 为 true 时只模拟，不写盘。
func Install(src string, targets []Target, dryRun bool) (*Result, error) {
	if src == "" {
		return nil, fmt.Errorf("内置 skills 源目录为空：发布包需完整解压（应包含 skill/ 目录与 ur 二进制同级）；若只有二进制，请重新下载完整安装包，或运行 ur skills download 获取 skills 后自行拷贝到目标 AI 工具的 skills 目录")
	}
	if !dirExists(src) {
		return nil, fmt.Errorf("内置 skills 源目录不存在: %s：发布包需完整解压（应包含 skill/ 目录与 ur 二进制同级）；若只有二进制，请重新下载完整安装包，或运行 ur skills download 获取 skills 后自行拷贝到目标 AI 工具的 skills 目录", src)
	}

	result := &Result{Source: src}
	for _, target := range targets {
		tr := TargetResult{Path: target.Path, Scope: target.Scope, Kind: target.Kind}
		dest := filepath.Join(target.Path, "ur-api")
		exists := dirExists(dest)
		tr.Updated = exists
		if dryRun {
			tr.Installed = true
			result.Targets = append(result.Targets, tr)
			continue
		}
		// 覆盖安装：rename 旧目录到临时备份 → copy 新 → 成功删备份，失败回滚
		if err := os.MkdirAll(target.Path, 0o755); err != nil {
			tr.Error = fmt.Sprintf("创建目标目录失败: %v", err)
			result.Targets = append(result.Targets, tr)
			continue
		}
		backup := dest + ".ur-bak"
		os.RemoveAll(backup)
		if exists {
			if err := os.Rename(dest, backup); err != nil {
				tr.Error = fmt.Sprintf("备份旧 ur-api 失败: %v", err)
				result.Targets = append(result.Targets, tr)
				continue
			}
		}
		if err := copyDir(src, dest); err != nil {
			os.RemoveAll(dest)
			if exists {
				os.Rename(backup, dest)
			}
			tr.Error = fmt.Sprintf("安装失败（已回滚）: %v", err)
			result.Targets = append(result.Targets, tr)
			continue
		}
		os.RemoveAll(backup)
		tr.Installed = true
		result.Targets = append(result.Targets, tr)
	}
	return result, nil
}

// Summary 生成人类可读的安装结果摘要
func Summary(r *Result) string {
	if r == nil || len(r.Targets) == 0 {
		return "未检测到可部署的 AI skills 目录（~/.claude/skills、~/.agents/skills 或项目 .claude/skills/.agents/skills 均不存在）"
	}
	var lines []string
	for _, t := range r.Targets {
		status := "已安装"
		if t.Error != "" {
			status = "失败: " + t.Error
		} else if t.Updated {
			status = "已覆盖更新"
		}
		lines = append(lines, fmt.Sprintf("  %s [%s/%s] %s → %s", status, t.Scope, t.Kind, t.Path, filepath.Join(t.Path, "ur-api")))
	}
	return "ur-api 部署结果:\n" + strings.Join(lines, "\n")
}
