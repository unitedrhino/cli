#!/usr/bin/env bash
# ur CLI 一键安装脚本(Linux / macOS)。对齐 urops 安装体验:自动选平台、
# 自动查最新版本、SHA256 校验、国内默认 Gitee 源。
#
# 用法:
#   curl -fsSL <脚本地址> | bash
#   指定版本/来源: curl -fsSL <脚本地址> | bash -s -- --version v0.7.0 --source github
#   环境变量: UR_VERSION=v0.7.0 UR_SOURCE=gitee|github UR_INSTALL_DIR=~/.local/lib/ur
#
# 卸载: rm -rf ~/.local/lib/ur ~/.local/bin/ur
set -euo pipefail

VERSION="latest"
SOURCE="auto"
INSTALL_DIR="${UR_INSTALL_DIR:-$HOME/.local/lib/ur}"
BIN_DIR="$HOME/.local/bin"
REPO_GITHUB="unitedrhino/cli"
REPO_GITEE="unitedrhino/cli"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --version) VERSION="$2"; shift 2 ;;
    --source)  SOURCE="$2";  shift 2 ;;
    *) echo "未知参数: $1" >&2; exit 2 ;;
  esac
done

# ── 平台检测 ──────────────────────────────────────────────────────────────
os="$(uname -s)"
arch="$(uname -m)"
case "$os" in
  Linux) os_name="Linux" ;;
  Darwin) os_name="macOS" ;;
  *) echo "不支持的系统: $os(Windows 请用 install.ps1)" >&2; exit 1 ;;
esac
case "$arch" in
  x86_64|amd64) arch_name="x86_64" ;;
  aarch64|arm64) arch_name="aarch64" ;;
  armv7l|armv6l) arch_name="armv7" ;;
  *) echo "不支持的架构: $arch" >&2; exit 1 ;;
esac
PLATFORM="${os_name}-${arch_name}"

# ── 来源选择:默认国内走 Gitee ────────────────────────────────────────────
pick_source() {
  if [[ "$SOURCE" != "auto" ]]; then echo "$SOURCE"; return; fi
  # 可达性探测:Gitee 优先(国内),超时 3s 回退 GitHub
  if curl -fsSI --connect-timeout 3 -m 5 "https://gitee.com" >/dev/null 2>&1; then
    echo "gitee"
  else
    echo "github"
  fi
}
SRC="$(pick_source)"

# ── 版本解析:latest 时查最新 release tag ─────────────────────────────────
resolve_latest() {
  local src="$1"
  if [[ "$src" == "gitee" ]]; then
    curl -fsSL --connect-timeout 5 "https://gitee.com/api/v5/repos/${REPO_GITEE}/releases?per_page=1" \
      | grep -o '"tag_name": *"[^"]*"' | head -1 | cut -d'"' -f4
  else
    curl -fsSL --connect-timeout 5 "https://api.github.com/repos/${REPO_GITHUB}/releases/latest" \
      | grep -o '"tag_name": *"[^"]*"' | head -1 | cut -d'"' -f4
  fi
}
if [[ "$VERSION" == "latest" ]]; then
  VERSION="$(resolve_latest "$SRC")"
  [[ -n "$VERSION" ]] || { echo "查询最新版本失败,请用 --version 指定(如 v0.7.0)" >&2; exit 1; }
fi

ASSET="ur-cli-${VERSION}-${PLATFORM}.tar.gz"
if [[ "$SRC" == "gitee" ]]; then
  DL_URL="https://gitee.com/${REPO_GITEE}/releases/download/${VERSION}/${ASSET}"
  SUMS_URL="https://gitee.com/${REPO_GITEE}/releases/download/${VERSION}/sha256sums.txt"
else
  DL_URL="https://github.com/${REPO_GITHUB}/releases/download/${VERSION}/${ASSET}"
  SUMS_URL="https://github.com/${REPO_GITHUB}/releases/download/${VERSION}/sha256sums.txt"
fi

echo "[ur-install] 平台: ${PLATFORM}  版本: ${VERSION}  来源: ${SRC}"

# ── 下载 + SHA256 校验 ───────────────────────────────────────────────────
TMP="$(mktemp -d /tmp/ur-install-XXXXXX)"
trap 'rm -rf "$TMP"' EXIT
echo "[ur-install] 下载 ${ASSET} ..."
curl -fSL --retry 3 -o "$TMP/pkg.tar.gz" "$DL_URL"
if curl -fsSL --retry 2 -o "$TMP/sha256sums.txt" "$SUMS_URL" 2>/dev/null; then
  expected="$(grep " ${ASSET}\$" "$TMP/sha256sums.txt" | awk '{print $1}' | head -1)"
  if [[ -n "$expected" ]]; then
    actual="$(sha256sum "$TMP/pkg.tar.gz" | awk '{print $1}')"
    [[ "$actual" == "$expected" ]] || { echo "SHA256 校验失败" >&2; exit 1; }
    echo "[ur-install] SHA256 校验通过"
  fi
fi

# ── 安装:ur 与 skill/ 同级(README 约定),symlink 进 PATH ─────────────────
mkdir -p "$INSTALL_DIR" "$BIN_DIR"
tar -xzf "$TMP/pkg.tar.gz" -C "$INSTALL_DIR" --strip-components=1
ln -sf "$INSTALL_DIR/ur" "$BIN_DIR/ur"
chmod +x "$INSTALL_DIR/ur"

echo "[ur-install] 已安装到 ${INSTALL_DIR}"
if ! command -v ur >/dev/null 2>&1; then
  echo "[ur-install] 提示: 请把 ${BIN_DIR} 加入 PATH"
  echo "              echo 'export PATH=\"${BIN_DIR}:\$PATH\"' >> ~/.bashrc && source ~/.bashrc"
fi

"$BIN_DIR/ur" --version 2>/dev/null || true
echo "[ur-install] 完成。下一步: ur check --json(检查认证);AI 环境请优先复用 Sandbox 变量"
echo "[ur-install] 可选: ${BIN_DIR}/ur skills install 安装 ur-api skills 到本地 AI 工具目录"
