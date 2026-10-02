#!/usr/bin/env bash
# =============================================================================
# ur CLI 一键安装脚本（官网/文档站提供：curl -fsSL <url> | bash）
#
# 与 urops 安装同源同链路：查询 Harbor 公开制品最新版本 → 匿名 token →
# 下载当前平台的完整发布包（ur 二进制 + skill/ 目录）与 sha256 → 校验 →
# 解压安装 → 打印版本。无固定版本号：每次执行都取 Harbor 最新发布。
#
# 支持平台：Linux x86_64/aarch64、macOS Intel/Apple Silicon（自动检测）。
#
# 环境变量（可选）：
#   UR_PREFIX    安装前缀（默认 ~/.local；ur 落 PREFIX/bin，skill/ 落 PREFIX/lib/ur）
#   UR_TAG       指定安装版本（默认取 Harbor 最新 vX.Y.Z）
#   UR_PLATFORM  强制平台（默认自动检测，形如 Linux-x86_64/macOS-arm64）
# =============================================================================
set -euo pipefail

REGISTRY="docker.unitedrhino.com"
PROJECT="urops"
PREFIX="${UR_PREFIX:-$HOME/.local}"
LIB_DIR="$PREFIX/lib/ur"
BIN_DIR="$PREFIX/bin"

command -v tar >/dev/null || { echo "✗ 需要 tar" >&2; exit 1; }

# sha256 工具兼容：Linux 用 sha256sum，macOS 只有 shasum -a 256
if command -v sha256sum >/dev/null 2>&1; then
  sha_gen()  { sha256sum "$@"; }
  sha_check_dir() { (cd "$1" && sha256sum -c sha256sums.txt); }
else
  sha_gen()  { shasum -a 256 "$@"; }
  sha_check_dir() { (cd "$1" && shasum -a 256 -c sha256sums.txt); }
fi

# 平台检测：映射到发布资产命名（Linux-x86_64 / Linux-aarch64 / macOS-x86_64 / macOS-arm64）
if [[ -n "${UR_PLATFORM:-}" ]]; then
  PLATFORM="$UR_PLATFORM"
else
  case "$(uname -s)" in
    Linux)  GOOS="Linux";;
    Darwin) GOOS="macOS";;
    *) echo "✗ 不支持的平台: $(uname -s)（支持 Linux / macOS；Windows 用 install.ps1）" >&2; exit 1;;
  esac
  case "$(uname -m)" in
    x86_64|amd64)  GOARCH="amd64";;
    aarch64|arm64) GOARCH="arm64";;
    *) echo "✗ 不支持的架构: $(uname -m)" >&2; exit 1;;
  esac
  PLATFORM="${GOOS}-${GOARCH}"
fi
case "$PLATFORM" in
  Linux-amd64|Linux-x86_64)  REPO="ur-cli-linux-amd64";;
  Linux-arm64|Linux-aarch64) REPO="ur-cli-linux-arm64";;
  macOS-amd64|macOS-x86_64)  REPO="ur-cli-darwin-amd64";;
  macOS-arm64)               REPO="ur-cli-darwin-arm64";;
  *) echo "✗ 不支持的平台: $PLATFORM" >&2; exit 1;;
esac

# 1. 最新版本（独立 API 路径绕过 EdgeOne 历史缓存；按语义版本取最大值）
if [[ -z "${UR_TAG:-}" ]]; then
  UR_TAG=$(curl -fsSL --max-time 15 \
    "https://${REGISTRY}/api-edgeone-v1/v2.0/projects/${PROJECT}/repositories/${REPO}/artifacts?page_size=50&with_tag=true&sort=push_time:desc" \
    | grep -oE '"name":"v[0-9.]+"' | cut -d'"' -f4 | sort -V | tail -1 || true)
  [[ -n "${UR_TAG}" ]] || { echo "✗ 无法获取最新版本（检查到 ${REGISTRY} 的网络）" >&2; exit 1; }
fi
echo "==> 安装 ur CLI ${UR_TAG}（${PLATFORM}）"

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

# 2. 匿名 token（registry /v2/ 401 响应解析 token 服务）
T=$(curl -s -D - -o /dev/null "https://${REGISTRY}/v2/" | grep -i 'www-authenticate' | tr -d '\r' | sed -n 's/.*realm="\([^"]*\)".*/\1/p')
S=$(curl -s -D - -o /dev/null "https://${REGISTRY}/v2/" | grep -i 'www-authenticate' | tr -d '\r' | sed -n 's/.*service="\([^"]*\)".*/\1/p')
TOKEN=$(curl -fsSL --max-time 15 "${T}?service=${S}&scope=repository:${PROJECT}/${REPO}:pull" | sed -n 's/.*"token":"\([^"]*\)".*/\1/p')
[[ -n "${TOKEN}" ]] || { echo "✗ 获取匿名 token 失败" >&2; exit 1; }

# 3. manifest 层 digest（第 0 层=完整发布包 tar.gz，第 1 层=sha256 文件；
#    第 1 个 digest 为 config，跳过。兼容紧凑/带空格两种 JSON 序列化）
MANIFEST=$(curl -fsSL --max-time 15 -H "Authorization: Bearer ${TOKEN}" \
  -H "Accept: application/vnd.oci.image.manifest.v1+json, application/vnd.docker.distribution.manifest.v2+json" \
  "https://${REGISTRY}/v2/${PROJECT}/${REPO}/manifests/${UR_TAG}")
DIGESTS=$(echo "${MANIFEST}" | grep -oE '"digest": *"sha256:[0-9a-f]+"' | grep -oE 'sha256:[0-9a-f]+')
PKG_DIGEST=$(echo "${DIGESTS}" | sed -n '2p')
SHA_DIGEST=$(echo "${DIGESTS}" | sed -n '3p')
[[ -n "${PKG_DIGEST}" && -n "${SHA_DIGEST}" ]] || { echo "✗ manifest 解析失败" >&2; exit 1; }

# 4. 下载发布包与校验文件（docker 层 blob 是「tar.gz 包裹的文件」：
#    两层都需先解外层 tar 才得到真实的发布包与 sha256 清单）
echo "==> 下载发布包"
curl -fsSL --max-time 300 -H "Authorization: Bearer ${TOKEN}" \
  "https://${REGISTRY}/v2/${PROJECT}/${REPO}/blobs/${PKG_DIGEST}" -o "$WORK/pkg.blob"
curl -fsSL --max-time 60 -H "Authorization: Bearer ${TOKEN}" \
  "https://${REGISTRY}/v2/${PROJECT}/${REPO}/blobs/${SHA_DIGEST}" -o "$WORK/sha.blob"
tar -xzf "$WORK/sha.blob" -C "$WORK"
PKG_NAME=$(awk '{print $2}' "$WORK/sha256sums.txt" | head -1)
[[ -n "${PKG_NAME}" ]] || { echo "✗ 校验清单为空" >&2; exit 1; }
tar -xzf "$WORK/pkg.blob" -C "$WORK"
[[ -f "$WORK/${PKG_NAME}" ]] || { echo "✗ 发布包缺失: ${PKG_NAME}" >&2; exit 1; }

# 5. 校验
echo "==> 校验 sha256"
sha_check_dir "$WORK" || { echo "✗ 校验失败" >&2; exit 1; }

# 6. 安装：ur 与 skill/ 保持同级（ur skills 依赖该布局），并软链进 PATH
mkdir -p "$LIB_DIR" "$BIN_DIR"
tar -xzf "$WORK/${PKG_NAME}" -C "$LIB_DIR" --strip-components=1
chmod +x "$LIB_DIR/ur"
ln -sf "$LIB_DIR/ur" "$BIN_DIR/ur"

# 7. 版本
"$BIN_DIR/ur" --help 2>/dev/null | head -1 || true
echo "==> 完成：$BIN_DIR/ur（skill/ 位于 $LIB_DIR/skill）"
case ":$PATH:" in
  *":$BIN_DIR:"*) ;;
  *) echo "提示：$BIN_DIR 不在 PATH 中，请将其加入 PATH 后重新打开会话" ;;
esac
