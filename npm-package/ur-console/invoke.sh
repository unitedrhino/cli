#!/bin/bash
# PicoClaw Skill Wrapper — 联犀 SaaS
set -e
SKILL_DIR="$(cd "$(dirname "$0")" && pwd)"
SKILL_MD="$SKILL_DIR/SKILL.md"
MODE="full"; QUERY=""; SECTION=""
while [[ $# -gt 0 ]]; do case "$1" in --query) QUERY="$2"; MODE="query"; shift 2 ;; --section) SECTION="$2"; MODE="section"; shift 2 ;; *) shift ;; esac; done

# 辅助函数：提取章节，如果不存在则返回空
get_section() {
  local file="$1" start="$2" end="$3"
  awk -v s="$start" -v e="$end" 'BEGIN{p=0} $0 ~ s{p=1} p{print} p && $0 ~ e && $0 !~ s{exit}' "$file"
}

if [[ "$MODE" == "query" && -n "$QUERY" ]]; then
  if echo "$QUERY" | grep -qiE "登录|login|认证|auth|setup|配置"; then
    echo "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答，禁止添加文档外信息。"
    # 子 skill 用 ## 使用示例 替代 ## 快速开始
    sec1=$(get_section "$SKILL_MD" "^## 快速开始" "^##")
    sec2=$(get_section "$SKILL_MD" "^## 使用示例" "^##")
    sec3=$(get_section "$SKILL_MD" "^## 配置指引" "^##")
    if [[ -n "$sec1" || -n "$sec2" || -n "$sec3" ]]; then
      [[ -n "$sec1" ]] && echo "$sec1"
      [[ -n "$sec2" ]] && echo "---" && echo "$sec2"
      [[ -n "$sec3" ]] && echo "---" && echo "$sec3"
    else
      cat "$SKILL_MD"
    fi
  elif echo "$QUERY" | grep -qiE "api|调用|schema|接口|endpoint|设备|产品|项目"; then
    echo "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答。"
    # 子 skill 用 ## API 端点 替代 ## CLI 用法
    sec=$(get_section "$SKILL_MD" "^## API 端点" "^##")
    if [[ -n "$sec" ]]; then
      echo "$sec"
    else
      sec=$(get_section "$SKILL_MD" "^## CLI 用法" "^##")
      [[ -n "$sec" ]] && echo "$sec" || cat "$SKILL_MD"
    fi
  elif echo "$QUERY" | grep -qiE "错误|排查|401|403|404|troubleshoot|故障"; then
    echo "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答。"
    # 优先读取 references/troubleshooting.md
    if [[ -f "$SKILL_DIR/references/troubleshooting.md" ]]; then
      cat "$SKILL_DIR/references/troubleshooting.md"
    else
      sec=$(get_section "$SKILL_MD" "^## 常见错误排查" "^##")
      [[ -n "$sec" ]] && echo "$sec" || cat "$SKILL_MD"
    fi
  else cat "$SKILL_MD"; fi
else cat "$SKILL_MD"; fi
