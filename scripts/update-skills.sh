#!/usr/bin/env bash
set -euo pipefail

# update-skills.sh — 一键更新 CLI skills 并同步到 skills 仓库
# 用法: bash scripts/update-skills.sh
#
# 流程:
#   1. 从 backend/.swagger/ 读取最新 swagger
#   2. 运行 generate-api-lists.py 更新所有 skill 的 API 端点列表
#   3. 同步到 unitedrhino/skills 仓库
#   4. 提示提交信息

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLI_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"
SKILL_DIR="${CLI_DIR}/skill"

# 查找 skills 仓库
SKILLS_REPO=""
for path in "${CLI_DIR}/../skills" "${CLI_DIR}/../../.gits/skills"; do
  if [ -d "$path/.git" ]; then
    SKILLS_REPO="$(cd "$path" && pwd)"
    break
  fi
done

echo "========================================"
echo "  ur Skills 更新脚本"
echo "========================================"
echo ""

# Step 1: 检查 swagger
echo "[1/4] 检查 swagger 文件..."
SWAGGER_DIR=""
for candidate in "${UR_SWAGGER_DIR:-}" "${CLI_DIR}/../backend/.swagger" "${CLI_DIR}/../../backend/.swagger" "${CLI_DIR}/../../../backend/.swagger"; do
  if [ -n "$candidate" ] && [ -f "$candidate/core-api.json" ] && [ -f "$candidate/things-api.json" ]; then
    SWAGGER_DIR="$candidate"
    break
  fi
done

if [ -z "$SWAGGER_DIR" ]; then
  echo "错误: 找不到 swagger 文件 (core-api.json / things-api.json)"
  echo "请设置 UR_SWAGGER_DIR 环境变量，或在 backend/.swagger 附近运行"
  exit 1
fi
echo "  swagger 目录: $SWAGGER_DIR"

# Step 2: 生成 API 列表
echo ""
echo "[2/4] 生成 API 端点列表..."
cd "$CLI_DIR"
python3 scripts/generate-api-lists.py

# Step 3: 同步到 skills 仓库
echo ""
echo "[3/4] 同步到 skills 仓库..."
if [ -z "$SKILLS_REPO" ]; then
  echo "  警告: 找不到 unitedrhino/skills 仓库，跳过同步"
  echo "  期望路径: ${CLI_DIR}/../skills 或 ${CLI_DIR}/../../.gits/skills"
else
  echo "  skills 仓库: $SKILLS_REPO"

  # 复制主 skill
  cp "${SKILL_DIR}/SKILL.md" "${SKILLS_REPO}/SKILL.md"

  # 复制子 skill
  for sub in ai-tool ur-device ur-device-analytics ur-device-debug ur-product ur-project ur-user ur-system ur-tenant ur-ai ur-view scene-linkage thing-model protocol-script; do
    if [ -d "${SKILL_DIR}/${sub}" ]; then
      rm -rf "${SKILLS_REPO}/${sub}"
      cp -r "${SKILL_DIR}/${sub}" "${SKILLS_REPO}/${sub}"
      echo "  同步: $sub"
    fi
  done

  # 排除内部文档
  for ex in ur-iot-client ur-iot-context ur-iot-device; do
    rm -rf "${SKILLS_REPO}/${ex}"
  done

  echo "  同步完成"
fi

# Step 4: 检查变更
echo ""
echo "[4/4] 检查变更..."
cd "$CLI_DIR"
if git diff --quiet -- skill/ 2>/dev/null; then
  echo "  skill/ 目录无变更"
else
  echo "  skill/ 目录有变更:"
  git diff --stat -- skill/ | sed 's/^/    /'
fi

if [ -n "$SKILLS_REPO" ]; then
  cd "$SKILLS_REPO"
  if git diff --quiet 2>/dev/null; then
    echo "  skills 仓库无变更"
  else
    echo "  skills 仓库有变更:"
    git diff --stat | sed 's/^/    /'
  fi
fi

echo ""
echo "========================================"
echo "  完成"
echo "========================================"
echo ""
echo "下一步:"
echo "  cd ${CLI_DIR}"
echo "  git add skill/ && git commit -m 'chore(skill): 更新 API 端点列表'"
echo "  git push"
if [ -n "$SKILLS_REPO" ]; then
  echo ""
  echo "  cd ${SKILLS_REPO}"
  echo "  git add -A && git commit -m 'chore(skill): 同步 API 端点列表'"
  echo "  git push origin main && git push gitee main"
fi
