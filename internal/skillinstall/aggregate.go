// aggregate.go — 聚合安装形态转换：把子域 SKILL.md 降级为无 frontmatter 的
// GUIDE.md，使 AI 工具（Claude Code / Codex / WorkBuddy 等）在 skills 目录里
// 只发现顶层 ur-api 一个技能入口，子域内容保留为普通参考文档随入口导航读取，
// 避免 20 个子域被注册成独立技能造成列表爆炸与触发条件互相干扰。
// 源目录（.gits/skills 同步的内置 skill/）保持 SKILL.md 标准格式不变，
// 转换只发生在 install / export 输出时，对源与生成链路零侵入。
package skillinstall

import (
	"bytes"
	"path/filepath"
	"strings"
)

// installedRelPath 返回源相对路径在安装产物中的形态：
// 顶层 SKILL.md 原样保留（唯一技能入口）；子目录内的 SKILL.md
// 改名为 GUIDE.md（降级为参考文档）。
func installedRelPath(rel string) string {
	normalized := filepath.ToSlash(rel)
	if normalized == "SKILL.md" || !strings.HasSuffix(normalized, "/SKILL.md") {
		return normalized
	}
	return strings.TrimSuffix(normalized, "SKILL.md") + "GUIDE.md"
}

// isAggregatable 判断相对路径是否为需要降级的子域 SKILL.md。
func isAggregatable(rel string) bool {
	normalized := filepath.ToSlash(rel)
	return normalized != "SKILL.md" && strings.HasSuffix(normalized, "/SKILL.md")
}

// stripYAMLFrontmatter 剥离 SKILL.md 开头的 YAML frontmatter 块
//（首行 --- 到下一个单独 --- 行）。无 frontmatter 或块未闭合时原样返回。
func stripYAMLFrontmatter(data []byte) []byte {
	content := data
	prefix := []byte("---")
	if !bytes.HasPrefix(content, prefix) {
		return data
	}
	// 首行必须恰为 "---"（允许 \r\n）
	rest := content[len(prefix):]
	rest = bytes.TrimPrefix(rest, []byte("\r\n"))
	if !bytes.HasPrefix(rest, []byte("\n")) {
		return data
	}
	rest = rest[1:]
	for {
		lineEnd := bytes.IndexByte(rest, '\n')
		if lineEnd < 0 {
			return data
		}
		line := bytes.TrimSuffix(rest[:lineEnd], []byte("\r"))
		if bytes.Equal(bytes.TrimSpace(line), prefix) {
			trimmed := rest[lineEnd+1:]
			// 保留正文内容，去掉 frontmatter 后紧随的一个空行
			if bytes.HasPrefix(trimmed, []byte("\r\n")) {
				trimmed = trimmed[2:]
			} else if bytes.HasPrefix(trimmed, []byte("\n")) {
				trimmed = trimmed[1:]
			}
			return trimmed
		}
		rest = rest[lineEnd+1:]
	}
}

// transformAggregate 把源文件转换为安装产物形态：子域 SKILL.md 降级为
// GUIDE.md 并剥离 frontmatter，其余文件原样返回。
func transformAggregate(rel string, data []byte) (string, []byte) {
	if !isAggregatable(rel) {
		return filepath.ToSlash(rel), data
	}
	return installedRelPath(rel), stripYAMLFrontmatter(data)
}
