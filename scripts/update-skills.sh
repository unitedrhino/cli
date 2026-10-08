#!/usr/bin/env bash
set -euo pipefail

# update-skills.sh — 可选镜像上游 skills,并一键更新 + 同步到 skills 仓库
# 用法: bash scripts/update-skills.sh [--apply] [--mirror-only] [--skip-mirror] [--source <dir>]
#
# skills 流向(独立 skills 仓库为业务原型源，本仓 skill/ 为发布副本):
#   unitedrhino/skills --镜像--> cli skill/ --release 发版--> ur-api-skills-<版本>.zip
#                         └--生成 API 后同步公开内容回原型源仓
# 隔离工作区通过 UR_SKILLS_REPO 指定源仓，--source 可覆盖镜像源。
# 客户端(AI 工具)通过 ur upgrade / ur skills install 拿到发布版本,不直接感知任何源仓库。
#
# 流程:
#   1. 镜像原型源仓 skills(默认 dry-run 仅列差异,--apply 执行;--skip-mirror 跳过)
#   2. 从 backend/.swagger/ 读取最新 swagger
#   3. 运行 generate-api-lists.py 更新所有 skill 的 API 端点列表
#   4. 同步到 unitedrhino/skills 仓库
#   5. 提示提交信息

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLI_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"
SKILL_DIR="${CLI_DIR}/skill"

MIRROR_APPLY=0
MIRROR_ONLY=0
SKIP_MIRROR=0
SOURCE_OVERRIDE="${UR_SKILLS_SOURCE:-}"

while [ $# -gt 0 ]; do
  case "$1" in
    --apply) MIRROR_APPLY=1 ;;
    --mirror-only) MIRROR_ONLY=1 ;;
    --skip-mirror) SKIP_MIRROR=1 ;;
    --source)
      [ $# -ge 2 ] || { echo "错误: --source 需要路径参数" >&2; exit 1; }
      SOURCE_OVERRIDE="$2"; shift ;;
    *) echo "未知参数: $1(支持 --apply / --mirror-only / --skip-mirror / --source <dir>)" >&2; exit 1 ;;
  esac
  shift
done

# 查找 skills 仓库
SKILLS_REPO="${UR_SKILLS_REPO:-}"
for path in "${CLI_DIR}/../skills" "${CLI_DIR}/../../.gits/skills"; do
  if [ -z "$SKILLS_REPO" ] && [ -e "$path/.git" ]; then
    SKILLS_REPO="$(cd "$path" && pwd)"
    break
  fi
done

echo "========================================"
echo "  ur Skills 更新脚本"
echo "========================================"
echo ""

# Step 1: 从 skills 原型源仓镜像；显式 --source 可指定隔离编辑源。
if [ "$SKIP_MIRROR" -eq 1 ]; then
  echo "[1/5] 跳过镜像(--skip-mirror)"
else
  echo "[1/5] 镜像上游 skills(如已配置)..."
  SOURCE_DIR=""
  for candidate in "${SOURCE_OVERRIDE}" "${SKILLS_REPO}"; do
    # 默认候选必须位于 git 工作树内，避免把 ~/.ur 或用户级的安装产物误当编辑源
    if [ -n "$candidate" ] && [ -f "$candidate/SKILL.md" ] \
      && git -C "$candidate" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
      SOURCE_DIR="$(cd "$candidate" && pwd)"
      break
    fi
  done

  if [ -z "$SOURCE_DIR" ]; then
    echo "  未检测到上游 skills 源,跳过镜像(--source <dir> 或环境变量 UR_SKILLS_SOURCE 可指定)"
  elif ! command -v rsync >/dev/null 2>&1; then
    echo "  警告: 未安装 rsync,跳过镜像;请安装 rsync 后重试"
  else
    echo "  镜像源: $SOURCE_DIR"
    repo_root="$(git -C "$SOURCE_DIR" rev-parse --show-toplevel 2>/dev/null || true)"
    if [ -n "$repo_root" ] && [ -n "$(git -C "$repo_root" status --porcelain -- "${SOURCE_DIR#"$repo_root"/}" 2>/dev/null)" ]; then
      echo "  注意: 上游源存在未提交改动,镜像的是工作区当前状态"
    fi
    # -c 按内容校验和比较,避免仅 mtime 不同造成假差异
    RSYNC_ARGS=(-a -c --delete --exclude=_meta.json --exclude=.git --exclude=.worktrees
      --exclude=ur-iot/ur-iot-client --exclude=ur-iot/ur-iot-context --exclude=ur-iot/ur-iot-device
      --itemize-changes)
    # _meta.json 由 release.sh 打包时自动生成,不从主仓带入
    if [ "$MIRROR_APPLY" -eq 1 ]; then
      echo "  执行镜像(--apply):"
      python3 "${SCRIPT_DIR}/sync-skills.py" "${SOURCE_DIR}" "${SKILL_DIR}" --preserve-client-guides
      echo "  镜像完成"
    else
      echo "  预览(dry-run,加 --apply 执行):"
      rsync "${RSYNC_ARGS[@]}" -n "${SOURCE_DIR}/" "${SKILL_DIR}/" | sed 's/^/    /'
      echo "  注意: 镜像以源为准,保留三个内部客户端专题,删除其他过期文件;"
      echo "  若 skill/ 有尚未合入上游源的内容(如未合并 PR 的技能),先更新上游再 --apply"
    fi
  fi
fi

if [ "$MIRROR_ONLY" -eq 1 ]; then
  echo ""
  echo "仅镜像模式结束(--mirror-only),未执行 swagger 生成与 skills 仓库同步。"
  exit 0
fi

# Step 2: 检查 swagger
echo ""
echo "[2/5] 检查 swagger 文件..."
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

# Step 3: 生成 API 列表
echo ""
echo "[3/5] 生成 API 端点列表..."
cd "$CLI_DIR"
python3 scripts/generate-api-lists.py

# Step 4: 同步到 skills 仓库
echo ""
echo "[4/5] 同步到 skills 仓库..."
if [ -z "$SKILLS_REPO" ]; then
  echo "  警告: 找不到 unitedrhino/skills 仓库，跳过同步"
  echo "  期望路径: ${CLI_DIR}/../skills 或 ${CLI_DIR}/../../.gits/skills"
else
  echo "  skills 仓库: $SKILLS_REPO"

  # 同步五组完整目录和通用参考，内部客户端专题不进入公开源仓。
  python3 "${SCRIPT_DIR}/sync-skills.py" "${SKILL_DIR}" "${SKILLS_REPO}" --public

  echo "  同步完成"
fi

# Step 5: 检查变更
echo ""
echo "[5/5] 检查变更..."
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
echo "  git add skill/ && git commit -m 'chore(skill): 同步 skills 及 API 端点列表'"
echo "  git push"
if [ -n "$SKILLS_REPO" ]; then
  echo ""
  echo "  cd ${SKILLS_REPO}"
  echo "  git add -A && git commit -m 'chore(skill): 同步 API 端点列表'"
  echo "  git push origin main && git push gitee main"
fi
