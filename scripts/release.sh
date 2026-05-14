#!/usr/bin/env bash
set -euo pipefail

# ur CLI 跨平台 Release 构建与发布脚本
# 用法: bash scripts/release.sh [VERSION]
# 示例: bash scripts/release.sh v0.1.0
#
# 环境变量（推荐写入 .env 文件，脚本会自动加载）：
#   GITHUB_TOKEN  - GitHub Personal Access Token
#   GITEE_TOKEN   - Gitee 私人令牌
#   PARALLEL      - 并发构建数（默认 8）

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# 自动加载 .env 文件（如果存在）
if [[ -f "${ROOT}/.env" ]]; then
  # shellcheck source=/dev/null
  source "${ROOT}/.env"
fi

VERSION="${1:-v0.1.0}"
BUILD_DIR="${ROOT}/dist/release-${VERSION}"
RELEASE_DIR="${BUILD_DIR}/packages"
PARALLEL="${PARALLEL:-8}"

# 排除的平台（非原生或不需要）
EXCLUDE_PLATFORMS="js/wasm wasip1/wasm android/386 android/amd64 android/arm android/arm64 ios/amd64 ios/arm64"

# 平台 → 友好名称映射
platform_name() {
  local goos="$1" goarch="$2"
  case "${goos}/${goarch}" in
    linux/amd64)   echo "Linux-x86_64" ;;
    linux/arm64)   echo "Linux-aarch64" ;;
    linux/386)     echo "Linux-i386" ;;
    linux/arm)     echo "Linux-armv7" ;;
    linux/arm/v6)  echo "Linux-armv6" ;;
    darwin/amd64)  echo "macOS-x86_64" ;;
    darwin/arm64)  echo "macOS-arm64" ;;
    windows/amd64) echo "Windows-x86_64" ;;
    windows/arm64) echo "Windows-arm64" ;;
    windows/386)   echo "Windows-i386" ;;
    freebsd/amd64) echo "FreeBSD-x86_64" ;;
    freebsd/arm64) echo "FreeBSD-aarch64" ;;
    openbsd/amd64) echo "OpenBSD-x86_64" ;;
    *)             echo "${goos}-${goarch}" ;;
  esac
}

# 检查是否需要排除
is_excluded() {
  local platform="$1"
  for ex in $EXCLUDE_PLATFORMS; do
    if [[ "$platform" == "$ex" ]]; then
      return 0
    fi
  done
  return 1
}

echo "========================================"
echo "  ur CLI Release Builder"
echo "  Version: ${VERSION}"
echo "  Parallel: ${PARALLEL}"
echo "========================================"
echo ""

# 确保版本号以 v 开头
if [[ ! "$VERSION" =~ ^v ]]; then
  VERSION="v${VERSION}"
fi

# 创建目录
rm -rf "${BUILD_DIR}"
mkdir -p "${RELEASE_DIR}"

# 获取所有原生平台（排除 wasm/android/ios/js）
PLATFORMS=()
while IFS='/' read -r goos goarch; do
  platform="${goos}/${goarch}"
  if is_excluded "$platform"; then
    echo "[skip] ${platform} (excluded)"
    continue
  fi
  PLATFORMS+=("${goos}:${goarch}")
done < <(go tool dist list | sort)

echo ""
echo "准备构建 ${#PLATFORMS[@]} 个平台，并发数: ${PARALLEL}"
echo ""

# 构建单个平台的函数
build_one() {
  local goos="$1" goarch="$2"
  local platform="${goos}/${goarch}"
  local name="$(platform_name "$goos" "$goarch")"
  local platform_dir="${BUILD_DIR}/${goos}-${goarch}"

  local exe_suffix=""
  if [[ "$goos" == "windows" ]]; then
    exe_suffix=".exe"
  fi

  mkdir -p "${platform_dir}"

  local build_err=0
  if ! (cd "${ROOT}" && GOOS="$goos" GOARCH="$goarch" CGO_ENABLED=0 go build \
        -ldflags "-s -w -X main.version=${VERSION}" \
        -o "${platform_dir}/ur${exe_suffix}" . 2>/dev/null); then
    build_err=1
  fi

  if [[ $build_err -ne 0 ]]; then
    echo "FAILED:${platform}"
    rm -rf "${platform_dir}"
    return 1
  fi

  # 复制 skill 资源到发布包（排除 scripts/ 子目录）
  mkdir -p "${platform_dir}/skill"
  for item in "${ROOT}/skill"/*; do
    local skill_name=$(basename "$item")
    if [[ "$skill_name" == "scripts" ]]; then
      continue
    fi
    if [[ -d "$item" ]]; then
      # 子 skill 目录：复制整个目录但排除其中的 scripts/
      mkdir -p "${platform_dir}/skill/${skill_name}"
      for subitem in "$item"/*; do
        local sub_skill_name=$(basename "$subitem")
        if [[ "$sub_skill_name" == "scripts" ]]; then
          continue
        fi
        cp -R "$subitem" "${platform_dir}/skill/${skill_name}/"
      done
    else
      cp -R "$item" "${platform_dir}/skill/"
    fi
  done

  echo "OK:${platform}:${name}"
}
export -f build_one platform_name
export ROOT BUILD_DIR VERSION

# 并行构建
BUILD_OK=0
BUILD_FAIL=0
FAILED_PLATFORMS=()

results=$(printf '%s\n' "${PLATFORMS[@]}" | xargs -P"${PARALLEL}" -I{} bash -c 'IFS=: read -r goos goarch <<< "$1"; build_one "$goos" "$goarch"' _ {})

while IFS= read -r line; do
  if [[ "$line" == OK:* ]]; then
    platform="${line#OK:}"
    platform="${platform%%:*}"
    name="${line##*:}"
    goos="${platform%%/*}"
    goarch="${platform#*/}"

    pkg_name="ur-cli-${VERSION}-${name}"
    platform_dir="${BUILD_DIR}/${goos}-${goarch}"
    if [[ "$goos" == "windows" ]]; then
      (cd "${BUILD_DIR}" && zip -rq "${RELEASE_DIR}/${pkg_name}.zip" "${goos}-${goarch}")
    else
      (cd "${BUILD_DIR}" && tar -czf "${RELEASE_DIR}/${pkg_name}.tar.gz" "${goos}-${goarch}")
    fi
    echo "[build] ${platform} (${name}) ... OK"
    BUILD_OK=$((BUILD_OK + 1))
  elif [[ "$line" == FAILED:* ]]; then
    platform="${line#FAILED:}"
    echo "[build] ${platform} ... FAILED"
    BUILD_FAIL=$((BUILD_FAIL + 1))
    FAILED_PLATFORMS+=("$platform")
  fi
done <<< "$results"

echo ""
echo "========================================"
echo "  构建结果"
echo "========================================"
echo "  成功: ${BUILD_OK}"
echo "  失败: ${BUILD_FAIL}"
if [[ ${#FAILED_PLATFORMS[@]} -gt 0 ]]; then
  echo "  失败平台:"
  for fp in "${FAILED_PLATFORMS[@]}"; do
    echo "    - ${fp}"
  done
fi
echo ""

# 生成 checksums
echo "[checksum] 生成 SHA256 校验和..."
cd "${RELEASE_DIR}"
sha256sum * > "sha256sums.txt"
cd - >/dev/null

echo ""
echo "发布包已生成: ${RELEASE_DIR}"
ls -lh "${RELEASE_DIR}"
echo ""

# ========================================
# GitHub Release
# ========================================
release_github() {
  if [[ -z "${GITHUB_TOKEN:-}" ]]; then
    echo "[github] 未设置 GITHUB_TOKEN，跳过 GitHub Release"
    echo "         如需发布，请设置环境变量后重新运行:"
    echo "         export GITHUB_TOKEN=ghp_xxxxxxxx"
    return 1
  fi

  local repo="unitedrhino/cli"
  echo "[github] 创建 Release: ${VERSION} ..."

  local release_resp
  release_resp=$(curl -s -X POST \
    -H "Authorization: token ${GITHUB_TOKEN}" \
    -H "Accept: application/vnd.github.v3+json" \
    "https://api.github.com/repos/${repo}/releases" \
    -d "{\"tag_name\":\"${VERSION}\",\"name\":\"ur-cli ${VERSION}\",\"body\":\"ur CLI ${VERSION} 跨平台发布\"}" 2>/dev/null)

  local upload_url
  upload_url=$(echo "$release_resp" | grep -o '"upload_url": "[^"]*' | cut -d'"' -f4 | sed 's/{?name,label}//')

  if [[ -z "$upload_url" ]]; then
    echo "[github] 创建 Release 失败:"
    echo "$release_resp" | head -5
    return 1
  fi

  echo "[github] Release 创建成功，开始上传资产..."

  for asset in "${RELEASE_DIR}"/*; do
    local fname
    fname=$(basename "$asset")
    echo -n "[github] 上传 ${fname} ... "
    if curl -s -X POST \
      -H "Authorization: token ${GITHUB_TOKEN}" \
      -H "Accept: application/vnd.github.v3+json" \
      -H "Content-Type: application/octet-stream" \
      "${upload_url}?name=${fname}" \
      --data-binary "@$asset" >/dev/null 2>&1; then
      echo "OK"
    else
      echo "FAILED"
    fi
  done

  echo "[github] 发布完成: https://github.com/${repo}/releases/tag/${VERSION}"
}

# ========================================
# Gitee Release
# ========================================
release_gitee() {
  if [[ -z "${GITEE_TOKEN:-}" ]]; then
    echo "[gitee] 未设置 GITEE_TOKEN，跳过 Gitee Release"
    echo "        如需发布，请设置环境变量后重新运行:"
    echo "        export GITEE_TOKEN=xxxxxxxx"
    return 1
  fi

  local owner="unitedrhino"
  local repo="cli"
  echo "[gitee] 创建 Release: ${VERSION} ..."

  local release_resp
  release_resp=$(curl -s -X POST \
    -H "Content-Type: application/json" \
    "https://gitee.com/api/v5/repos/${owner}/${repo}/releases" \
    -d "{\"access_token\":\"${GITEE_TOKEN}\",\"tag_name\":\"${VERSION}\",\"target_commitish\":\"main\",\"name\":\"ur-cli ${VERSION}\",\"body\":\"ur CLI ${VERSION} 跨平台发布\"}" 2>/dev/null)

  local release_id
  release_id=$(echo "$release_resp" | grep -o '"id":[0-9]*' | head -1 | cut -d':' -f2)

  if [[ -z "$release_id" ]]; then
    echo "[gitee] 创建 Release 失败:"
    echo "$release_resp" | head -5
    return 1
  fi

  echo "[gitee] Release 创建成功 (id=${release_id})，开始上传资产..."

  for asset in "${RELEASE_DIR}"/*; do
    local fname
    fname=$(basename "$asset")
    echo -n "[gitee] 上传 ${fname} ... "
    if curl -s -X POST \
      "https://gitee.com/api/v5/repos/${owner}/${repo}/releases/${release_id}/attach_files" \
      -H "Content-Type: multipart/form-data" \
      -F "access_token=${GITEE_TOKEN}" \
      -F "file=@${asset}" >/dev/null 2>&1; then
      echo "OK"
    else
      echo "FAILED"
    fi
  done

  echo "[gitee] 发布完成: https://gitee.com/${owner}/${repo}/releases/tag/${VERSION}"
}

# 尝试发布
echo "========================================"
echo "  发布阶段"
echo "========================================"
echo ""

release_github || true
echo ""
release_gitee || true

echo ""
echo "========================================"
echo "  Done"
echo "========================================"
echo "构建产物: ${RELEASE_DIR}"
