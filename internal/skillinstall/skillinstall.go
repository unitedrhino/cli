// skillinstall — 把内置 ur-api skill 部署到各 AI 工具的 skills 目录。
//
// 不同 AI 工具（Claude Code / Codex / WorkBuddy 等）存放 skills 的位置不同：
//   - Claude Code：~/.claude/skills/ 与项目 .claude/skills/
//   - Codex：~/.agents/skills/ 与项目 .agents/skills/（项目级常为软链目录）
//   - WorkBuddy / CodeBuddy：~/.codebuddy/skills/、项目 .codebuddy/skills/，
//     或 CODEBUDDY_CONFIG_DIR 指定配置目录下的 skills/
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
	// Name 是目标的稳定名称，供配置、筛选和诊断输出使用。
	Name string `json:"name"`
	// Path 目标 skills 根目录（ur-api 将被安装为 <Path>/ur-api/）
	Path string `json:"path"`
	// Scope 作用域：user（用户级）/ project（项目级）
	Scope string `json:"scope"`
	// Kind AI 工具类型：claude / codex / workbuddy / custom。
	Kind string `json:"kind"`
	// Origin 表示目标来自自动探测、用户配置、环境变量或命令行参数。
	Origin string `json:"origin"`
}

// TargetResult 单个目标的安装结果
type TargetResult struct {
	// Name 是目标稳定名称。
	Name string `json:"name"`
	// Path 是客户端 Skills 根目录。
	Path string `json:"path"`
	// Scope 是用户级、项目级或自定义作用域。
	Scope string `json:"scope"`
	// Kind 是客户端类型。
	Kind string `json:"kind"`
	// Origin 是目标来源。
	Origin string `json:"origin"`
	// Installed 表示本次安装是否完成。
	Installed bool `json:"installed"`
	// Updated 表示目标原先已有 ur-api 并被覆盖更新。
	Updated bool `json:"updated,omitempty"`
	// Error 是单个目标的安装错误。
	Error string `json:"error,omitempty"`
}

// Result 整体安装结果
type Result struct {
	// Source 内置 skills 源目录
	Source string `json:"source"`
	// Targets 是逐目标安装结果。
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
	if dirExists(filepath.Join(home, ".claude")) {
		targets = append(targets, Target{Name: "claude-user", Path: filepath.Join(home, ".claude", "skills"), Scope: "user", Kind: "claude", Origin: "detected"})
	}
	if dirExists(filepath.Join(home, ".agents")) {
		targets = append(targets, Target{Name: "codex-user", Path: filepath.Join(home, ".agents", "skills"), Scope: "user", Kind: "codex", Origin: "detected"})
	}
	if dirExists(filepath.Join(home, ".codebuddy")) {
		targets = append(targets, Target{Name: "workbuddy-user", Path: filepath.Join(home, ".codebuddy", "skills"), Scope: "user", Kind: "workbuddy", Origin: "detected"})
	}
	if configDir := strings.TrimSpace(os.Getenv("CODEBUDDY_CONFIG_DIR")); configDir != "" {
		targets = append(targets, Target{Name: "workbuddy-config", Path: filepath.Join(configDir, "skills"), Scope: "user", Kind: "workbuddy", Origin: "environment"})
	}
	if envDirs := strings.TrimSpace(os.Getenv("UR_SKILLS_DIRS")); envDirs != "" {
		for index, dir := range filepath.SplitList(envDirs) {
			if strings.TrimSpace(dir) == "" {
				continue
			}
			targets = append(targets, Target{
				Name:   fmt.Sprintf("environment-%d", index+1),
				Path:   dir,
				Scope:  "custom",
				Kind:   "custom",
				Origin: "environment",
			})
		}
	}

	// 项目级（当前 git 仓库根下）
	if root := findRepoRoot(cwd); root != "" {
		if dirExists(filepath.Join(root, ".claude")) {
			targets = append(targets, Target{Name: "claude-project", Path: filepath.Join(root, ".claude", "skills"), Scope: "project", Kind: "claude", Origin: "detected"})
		}
		if dirExists(filepath.Join(root, ".agents")) {
			targets = append(targets, Target{Name: "codex-project", Path: filepath.Join(root, ".agents", "skills"), Scope: "project", Kind: "codex", Origin: "detected"})
		}
		if dirExists(filepath.Join(root, ".codebuddy")) {
			targets = append(targets, Target{Name: "workbuddy-project", Path: filepath.Join(root, ".codebuddy", "skills"), Scope: "project", Kind: "workbuddy", Origin: "detected"})
		}
	}
	return DedupeTargets(targets), nil
}

// DedupeTargets 按规范化绝对路径去重目标，保留同一路径最先出现的配置。
func DedupeTargets(targets []Target) []Target {
	seen := make(map[string]struct{}, len(targets))
	result := make([]Target, 0, len(targets))
	for _, target := range targets {
		target.Path = normalizePath(target.Path)
		if target.Path == "" {
			continue
		}
		key := filepath.Clean(target.Path)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, target)
	}
	return result
}

// normalizePath 将用户或环境变量提供的目标路径转换为干净的绝对路径。
func normalizePath(path string) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return ""
	}
	if path == "~" || strings.HasPrefix(path, "~/") || strings.HasPrefix(path, `~\`) {
		if home, err := os.UserHomeDir(); err == nil {
			if path == "~" {
				path = home
			} else {
				rest := strings.ReplaceAll(path[2:], `\`, string(filepath.Separator))
				path = filepath.Join(home, rest)
			}
		}
	}
	if abs, err := filepath.Abs(path); err == nil {
		return filepath.Clean(abs)
	}
	return filepath.Clean(path)
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
		out, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			in.Close()
			return err
		}
		_, copyErr := io.Copy(out, in)
		inCloseErr := in.Close()
		outCloseErr := out.Close()
		if copyErr != nil {
			return copyErr
		}
		if inCloseErr != nil {
			return inCloseErr
		}
		return outCloseErr
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
		target.Path = normalizePath(target.Path)
		tr := TargetResult{Name: target.Name, Path: target.Path, Scope: target.Scope, Kind: target.Kind, Origin: target.Origin}
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

// HasErrors 判断安装结果中是否有目标失败，便于命令行返回非零退出码。
func HasErrors(result *Result) bool {
	if result == nil {
		return false
	}
	for _, target := range result.Targets {
		if target.Error != "" {
			return true
		}
	}
	return false
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
		lines = append(lines, fmt.Sprintf("  %s %s [%s/%s] %s → %s", status, t.Name, t.Scope, t.Kind, t.Path, filepath.Join(t.Path, "ur-api")))
	}
	return "ur-api 部署结果:\n" + strings.Join(lines, "\n")
}
