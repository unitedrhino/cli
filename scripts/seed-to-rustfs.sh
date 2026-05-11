#!/usr/bin/env bash
set -euo pipefail

# seed-to-rustfs.sh — 把 ur-api skill 文件复制到 rustfs-seed 目录，供 rustfs-import-public.sh 上传到 S3
#
# 用法:
#   bash scripts/seed-to-rustfs.sh                    # 同步 template/full + template/slim
#   bash scripts/seed-to-rustfs.sh --all              # 额外同步所有 production 实例
#   bash scripts/seed-to-rustfs.sh --dry-run          # 预览变更，不实际写入
#   bash scripts/seed-to-rustfs.sh --seed-dir DIR     # 只同步指定目录

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
REPO_ROOT="$(cd "${ROOT}/../.." && pwd)"

# 默认目标目录
declare -a TARGET_DIRS=()
ALL=false
DRY_RUN=false
CUSTOM_SEED_DIR=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --all) ALL=true; shift ;;
    --dry-run) DRY_RUN=true; shift ;;
    --seed-dir)
      if [[ -z "${2:-}" ]]; then
        echo "[seed-to-rustfs] error: --seed-dir requires a directory" >&2
        exit 1
      fi
      CUSTOM_SEED_DIR="$2"
      shift 2
      ;;
    -h|--help)
      echo "Usage: $(basename "$0") [--all] [--dry-run] [--seed-dir DIR]"
      echo ""
      echo "Options:"
      echo "  --all        Also sync production instances (106, 115, 120, 134)"
      echo "  --dry-run    Preview changes without writing"
      echo "  --seed-dir   Sync only the specified seed directory"
      exit 0
      ;;
    *)
      echo "[seed-to-rustfs] error: unknown arg: $1" >&2
      echo "Run with --help for usage." >&2
      exit 1
      ;;
  esac
done

# 构建目标目录列表
if [[ -n "$CUSTOM_SEED_DIR" ]]; then
  TARGET_DIRS=("$CUSTOM_SEED_DIR")
else
  TARGET_DIRS+=(
    "${REPO_ROOT}/deploy/docker/template/full/conf/rustfs-seed/public"
    "${REPO_ROOT}/deploy/docker/template/slim/conf/rustfs-seed/public"
  )
  if $ALL; then
    TARGET_DIRS+=(
      "${REPO_ROOT}/deploy/docker/production/106/conf/rustfs-seed/public"
      "${REPO_ROOT}/deploy/docker/production/115/conf/rustfs-seed/public"
      "${REPO_ROOT}/deploy/docker/production/120/conf/rustfs-seed/public"
      "${REPO_ROOT}/deploy/docker/production/134/conf/rustfs-seed/public"
    )
  fi
fi

# Step 1: 构建 + 生成 skill
DIST_DIR="${ROOT}/dist"
if ! $DRY_RUN; then
  echo "[seed-to-rustfs] building and generating skills..."
  bash "${ROOT}/scripts/package-skill.sh" "${DIST_DIR}"
fi

# Step 2: 复制到各 seed 目录
for seed_dir in "${TARGET_DIRS[@]}"; do
  target="${seed_dir}/skills/shared/ur-api"

  if [[ ! -d "${seed_dir}" ]]; then
    echo "[seed-to-rustfs] skip: seed dir not found: ${seed_dir}"
    continue
  fi

  if $DRY_RUN; then
    echo "[seed-to-rustfs] dry-run: would sync -> ${target}"
    if [[ -d "$target" ]]; then
      echo "  (existing files: $(find "$target" -type f | wc -l))"
    else
      echo "  (target does not exist yet)"
    fi
    continue
  fi

  echo "[seed-to-rustfs] syncing -> ${target}"
  rm -rf "${target}"
  mkdir -p "${target}"
  cp -R "${DIST_DIR}/skill/." "${target}/"
  echo "[seed-to-rustfs] done: ${target} ($(find "${target}" -type f | wc -l) files)"
done

echo "[seed-to-rustfs] all done."
