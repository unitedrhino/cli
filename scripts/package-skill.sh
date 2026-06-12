#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SAAS_ROOT="$(cd "${ROOT}/../../.." && pwd)"
OUT_DIR="${SAAS_ROOT}/dist/resource/claw"
ARCHS=""
# SKILL_VERSION: 可通过环境变量或 --version 参数指定，默认 0.0.0-dev
SKILL_VERSION="${SKILL_VERSION:-0.0.0-dev}"

# GOOS-GOARCH → 分发名映射（如 windows-amd64 → x64-win）
goarch_to_distname() {
  local pair="$1"
  case "$pair" in
    linux-amd64)   echo "x64-linux" ;;
    linux-arm64)   echo "arm64-linux" ;;
    darwin-amd64)  echo "x64-mac" ;;
    darwin-arm64)  echo "arm64-mac" ;;
    windows-amd64) echo "x64-win" ;;
    windows-arm64) echo "arm64-win" ;;
    *) echo "$pair" ;;
  esac
}

# 分发名 → GOOS-GOARCH 反向映射
distname_to_goarch() {
  local name="$1"
  case "$name" in
    x64-linux)   echo "linux-amd64" ;;
    arm64-linux) echo "linux-arm64" ;;
    x64-mac)     echo "darwin-amd64" ;;
    arm64-mac)   echo "darwin-arm64" ;;
    x64-win)     echo "windows-amd64" ;;
    arm64-win)   echo "windows-arm64" ;;
    *) echo "$name" ;;
  esac
}

# 解析参数
while [[ $# -gt 0 ]]; do
  case "$1" in
    --arch)
      if [[ -z "${2:-}" ]]; then
        echo "[package-skill] error: --arch requires a comma-separated list" >&2
        exit 1
      fi
      ARCHS="$2"
      shift 2
      ;;
    --version)
      if [[ -z "${2:-}" ]]; then
        echo "[package-skill] error: --version requires a version" >&2
        exit 1
      fi
      SKILL_VERSION="$2"
      shift 2
      ;;
    -h|--help)
      echo "Usage: $(basename "$0") [OUT_DIR] [--arch ARCH_LIST] [--version VERSION]"
      echo ""
      echo "Arguments:"
      echo "  OUT_DIR            Output directory (default: ./dist)"
      echo ""
      echo "Options:"
      echo "  --arch LIST        Comma-separated list of GOOS-GOARCH pairs"
      echo "                     Examples: linux-amd64,linux-arm64,darwin-amd64,darwin-arm64"
      echo "  --version VERSION  Skills version (default: ${SKILL_VERSION})"
      echo "  -h, --help         Show this help"
      exit 0
      ;;
    -*)
      echo "[package-skill] error: unknown option: $1" >&2
      exit 1
      ;;
    *)
      OUT_DIR="$1"
      shift
      ;;
  esac
done

# 如果没有指定架构，只构建当前架构
if [[ -z "$ARCHS" ]]; then
  LOCAL_GOOS="$(go env GOOS)"
  LOCAL_GOARCH="$(go env GOARCH)"
  ARCHS="${LOCAL_GOOS}-${LOCAL_GOARCH}"
fi

# 解析架构列表
IFS=',' read -ra ARCH_LIST <<< "$ARCHS"

build_for_arch() {
  local arch_pair="$1"
  local goos="${arch_pair%-*}"
  local goarch="${arch_pair#*-}"
  local dist_name="$(goarch_to_distname "$arch_pair")"
  local arch_dir="${OUT_DIR}/${dist_name}"
  local bin_dir="${arch_dir}/bin"
  local skill_dir="${arch_dir}/skill"
  local plugin_dir="${arch_dir}/openclaw-plugin"

  mkdir -p "${bin_dir}" "${skill_dir}" "${plugin_dir}"

  echo "[package-skill] building for ${goos}/${goarch} -> ${arch_dir} (dist: ${dist_name})"

  # 构建五个应用 CLI 二进制
  local build_err=0
  local exe_suffix=""
  if [[ "${goos}" == "windows" ]]; then exe_suffix=".exe"; fi

  for app in platform-manage iot org-manage org-energy console; do
    if ! (cd "${ROOT}" && GOOS="${goos}" GOARCH="${goarch}" go build -o "${bin_dir}/ur-${app}${exe_suffix}" "./cmd/ur-${app}"); then
      echo "[package-skill] warning: failed to build ur-${app} for ${goos}/${goarch}" >&2
      build_err=1
    fi
  done

  # 向后兼容：保留旧的 ur 二进制
  if ! (cd "${ROOT}" && GOOS="${goos}" GOARCH="${goarch}" go build -o "${bin_dir}/ur${exe_suffix}" .); then
    echo "[package-skill] warning: failed to build ur for ${goos}/${goarch}" >&2
    build_err=1
  fi

  if [[ $build_err -ne 0 ]]; then
    echo "[package-skill] skipping skill generation for ${arch_pair} due to build errors"
    return
  fi

  # 生成统一的 ur-api skill 文档（包含所有应用的所有端点）
  local api_skill_dir="${skill_dir}/ur-api"
  mkdir -p "${api_skill_dir}"
  (cd "${ROOT}" && go run . generate-skills --all --output "${api_skill_dir}")
  cp -R "${ROOT}/references" "${api_skill_dir}/references" 2>/dev/null || true

  # 保留顶层 SKILL.md 作为向后兼容的入口（内容指向 ur-api）
  cat > "${skill_dir}/SKILL.md" <<'INDEX'
---
name: ur-api
description: "联犀 SaaS 平台 API 命令行工具 — 按前端应用拆分为 5 个独立 CLI，每个绑定固定的 app-id 和 tenant-code"
metadata:
  hermes:
    tags: [api, cli, saas, iot]
---

# ur-api — 联犀 SaaS 平台 API 工具（应用导向）

## ⚠️ 配置指引（首次使用必读）

**禁止在当前对话中向用户索取 AccessKey、AccessSecret、密码等敏感信息。**

### AI 环境中唯一可用的配置方式

`ur-xxx setup` 是终端交互式命令（需要逐行输入账号密码），**在当前对话环境中无法使用**。

**唯一可行的配置方式**是设备授权流（Device Flow），AI 分两步完成：

**第 1 步 — 获取授权 URL**：
```bash
ur-iot login --no-wait --json
```
输出 JSON 示例：
```json
{
  "verification_url": "https://console.unitedrhino.com/user/access-tokens?setup=ABC123",
  "setup_code": "ABC123",
  "expires_in": 600,
  "hint": "在浏览器中打开 verification_url 完成授权，然后执行: login --setup-code ABC123"
}
```
AI 解析 JSON，向用户展示 `verification_url` 和操作步骤。

**第 2 步 — 完成授权**（用户确认在浏览器中点击「完成 CLI 绑定」后）：
```bash
ur-iot login --setup-code ABC123 --json
```
输出 JSON 示例：
```json
{
  "event": "authorization_complete",
  "tenant_code": "t1",
  "access_key": "ak_xxxx",
  "access_secret": "sk_xxxx"
}
```

**全程无需在当前对话中输入任何密钥。**

### 人类终端模式（非 AI 环境）

如果用户在本地终端直接操作：
```bash
# 阻塞模式：生成 URL → 等待用户浏览器授权 → 自动保存配置
ur-iot login

# 或指定地址跳过交互选择
ur login --base-url https://saas.unitedrhino.com
```

### 认证方式说明

CLI 支持两种认证机制：
1. **Session Token**（`ur-iot login` 自动获取）：`token:` header
2. **AccessKey/JWT**（login 完成后自动保存）：AccessKey + AccessSecret → HS256 JWT → `Authorization: Bearer` header

所有接口均为 POST 方法。请求格式 `{code, msg, data}`。

---

## 应用选择

根据当前操作的前端应用选择对应的 CLI：

| CLI | 前端应用 | AppID | TenantCode | 可调用权限 |
|-----|---------|-------|------------|-----------|
| `ur-platform-manage` | 平台管理 | 100 | platform | platform, admin, all |
| `ur-iot` | 物联网 | 200 | platform | platform, admin, all |
| `ur-org-manage` | 组织管理 | 300 | 用户输入 | admin, all |
| `ur-org-energy` | 能源管理 | 1000 | 用户输入 | admin, all |
| `ur-console` | 控制台 | 77 | platform | all |

## 快速决策

| 用户意图 | 使用 CLI | 说明 |
|---------|---------|------|
| 管理企业、用户、应用、授权 | `ur-platform-manage` | 平台管理员操作，跨企业 |
| 管理设备、产品、项目、OTA、协议 | `ur-iot` | 物联网（**平台/企业共用**，tenant-code=platform） |
| 管理组织用户、角色、AI 智能体 | `ur-org-manage` | 组织管理（企业级，tenant-code 需输入） |
| 能耗分析、电力集抄、预付费 | `ur-org-energy` | 能源管理（企业级，tenant-code 需输入） |
| 个人信息、访问令牌、续期 | `ur-console` | 控制台（平台级，tenant-code=platform） |

### 物联网 vs 能源管理 — 能力重叠与分工

`ur-org-energy` **复用了物联网的核心设备能力**，但聚焦能源场景：

| 能力 | `ur-iot` | `ur-org-energy` | 说明 |
|------|---------|----------------|------|
| 设备管理 | ✅ 产品 + 设备 CRUD、物模型 | ✅ 设备控制台、分组、区域 | 能源管理不管理产品定义，只使用已创建的设备 |
| 场景联动 | ✅ 项目级场景编辑 | ✅ 自动化 + 一键场景 | 两者都支持，数据互通 |
| 区域管理 | ✅ 区域 CRUD | ✅ 区域 CRUD | 共用同一套区域数据 |
| OTA 升级 | ✅ 固件包、批量升级 | ❌ | 能源管理不涉及固件升级 |
| 数据流转 | ✅ 协议网关、协议脚本 | ❌ | 能源管理不涉及协议开发 |
| 能耗分析 | ❌ | ✅ 用能概况、同比环比、趋势、排名、损耗 | 能源特有 |
| 电力集抄 | ❌ | ✅ 实时监控、历史数据、极值报表 | 能源特有 |
| 预付费管理 | ❌ | ✅ 房间、租客、充值、消费 | 能源特有 |

### ur-iot 按角色可见菜单

`ur-iot` **不是平台管理员专属**，企业管理员也能使用大部分功能。

**平台管理员（platform）可见全部菜单**：
- 信息面板、设备地图
- 设备管理（产品管理、设备管理、**通用物模型**、**产品品类**）
- 项目管理（项目列表、项目详情、场景编辑）
- 区域管理
- OTA 升级（升级包列表、模块列表、批量升级）
- 数据流转（**协议网关**、协议脚本）

**企业管理员（admin）可见菜单**（不含 `authority: ['platform']` 标记的）：
- 信息面板、设备地图
- 设备管理（产品管理、设备管理）
- 项目管理（项目列表、项目详情、场景编辑）
- 区域管理
- OTA 升级（升级包列表、模块列表、批量升级）
- 数据流转（协议脚本）

> **平台专属菜单**（企业管理员不可见）：通用物模型、产品品类、协议网关
> 
> 前端路由单一事实来源：`apps/web/apps/iot/src/router/routes/modules/iot.ts`

**决策规则**：
1. 如果是**能源业务**（能耗分析、电力集抄、预付费）→ 用 `ur-org-energy`
2. 如果是**设备/产品/项目管理/OTA** → 平台管理员和企业管理员都用 `ur-iot`
3. 如果是**协议开发**（协议网关、通用物模型、产品品类）→ 必须用 `ur-iot`（且需平台管理员权限）

## 通用用法

```bash
# 验证连通性
ur-iot check

# 调用 API
ur-iot api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'

# 查看 API schema
ur-iot schema

# 查看 token
ur-iot token --decode
```
INDEX

  # 生成 Docker 场景 wrapper 脚本（直接调用 /usr/local/bin/ur-xxx）
  local api_skill_dir="${skill_dir}/ur-api"
  mkdir -p "${api_skill_dir}/scripts"
  for app in platform-manage iot org-manage org-energy console; do
    for cmd in api check schema token login config setup generate-skills; do
      cat > "${api_skill_dir}/scripts/${cmd}-${app}.sh" <<EOF
#!/usr/bin/env bash
set -euo pipefail
exec /usr/local/bin/ur-${app} ${cmd} "\$@"
EOF
      chmod +x "${api_skill_dir}/scripts/${cmd}-${app}.sh"
      # Windows 备用：生成 .cmd 包装脚本
      if [[ "${goos}" == "windows" ]]; then
        cat > "${api_skill_dir}/scripts/${cmd}-${app}.cmd" <<EOFCMD
@echo off
"C:\Program Files\ur\ur-${app}.exe" ${cmd} %*
EOFCMD
      fi
    done
  done

  # 为每个 skill 生成 PicoClaw 兼容的包装层（invoke.sh + 元数据）
  # OpenClaw 直接使用 SKILL.md，PicoClaw 需要 invoke.sh
  echo "[package-skill] generating PicoClaw wrappers for skills"
  for app_skill_dir in "${skill_dir}"/ur-*; do
    if [[ -d "$app_skill_dir" && -f "$app_skill_dir/SKILL.md" ]]; then
      local app_name="$(basename "$app_skill_dir")"
      # 生成 _meta.json
      cat > "${app_skill_dir}/_meta.json" << EOF
{
  "ownerId": "unitedrhino",
  "slug": "${app_name}",
  "version": "${SKILL_VERSION:-0.0.0-dev}",
  "publishedAt": $(date +%s)000
}
EOF
      # 生成 .skill-origin.json
      cat > "${app_skill_dir}/.skill-origin.json" << EOF
{
  "version": 1,
  "origin_kind": "third_party",
  "registry": "unitedrhino",
  "slug": "${app_name}",
  "registry_url": "https://github.com/unitedrhino/saas",
  "installed_version": "${SKILL_VERSION:-0.0.0-dev}",
  "installed_at": $(date +%s)000
}
EOF
      # 生成 invoke.sh (bash)
      cat > "${app_skill_dir}/invoke.sh" << 'PICOINVOKE'
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
PICOINVOKE
      chmod +x "${app_skill_dir}/invoke.sh"

      # 生成 invoke.ps1 (PowerShell) — Windows 备用
      cat > "${app_skill_dir}/invoke.ps1" << 'PICOINVOKE_PS'
# PicoClaw Skill Wrapper — 联犀 SaaS (PowerShell)
param(
    [string]$Query,
    [string]$Section
)

$SkillDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$SkillMd = Join-Path $SkillDir "SKILL.md"

function Get-SectionContent {
    param([string]$FilePath, [string]$StartPattern, [string]$EndPattern)
    $lines = Get-Content $FilePath
    $capture = $false
    $result = @()
    foreach ($line in $lines) {
        if ($line -match $StartPattern) { $capture = $true }
        if ($capture) { $result += $line }
        if ($capture -and $EndPattern -and ($line -match $EndPattern) -and ($line -notmatch $StartPattern)) { break }
    }
    return $result -join "`n"
}

if ($Query) {
    $q = $Query.ToLower()
    if ($q -match "登录|login|认证|auth|setup") {
        Write-Output "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答，禁止添加文档外信息。"
        $sec1 = Get-SectionContent $SkillMd "^## 快速开始" "^## CLI 用法"
        $sec2 = Get-SectionContent $SkillMd "^## 认证原理" "^## API 通用约定"
        if ($sec1) { Write-Output $sec1 } else { Write-Output "---"; Get-Content $SkillMd -Raw }
        if ($sec2) { Write-Output "---"; Write-Output $sec2 }
    } elseif ($q -match "api|调用|schema|接口|endpoint") {
        Write-Output "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答。"
        $sec = Get-SectionContent $SkillMd "^## CLI 用法" "^## API 通用约定"
        if ($sec) { Write-Output $sec } else { Get-Content $SkillMd -Raw }
    } elseif ($q -match "错误|排查|401|403|404|troubleshoot") {
        Write-Output "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答。"
        $sec = Get-SectionContent $SkillMd "^## 常见错误排查" "^## 各域 API 概览"
        if ($sec) { Write-Output $sec } else { Get-Content $SkillMd -Raw }
    } else {
        Get-Content $SkillMd -Raw
    }
} else {
    Get-Content $SkillMd -Raw
}
PICOINVOKE_PS
    fi
  done

  echo "[package-skill] done: ${arch_pair} -> ${dist_name}"
}

# 清理旧输出
rm -rf "${OUT_DIR}"
mkdir -p "${OUT_DIR}"

# 为每个架构构建
for arch in "${ARCH_LIST[@]}"; do
  build_for_arch "$arch"
done

# 生成架构感知 wrapper 脚本（裸机/宿主机场景）
ARCH_AWARE_WRAPPER="${OUT_DIR}/ur-wrapper.sh"
cat > "${ARCH_AWARE_WRAPPER}" <<'WRAPPER'
#!/usr/bin/env bash
set -euo pipefail

# 架构感知 wrapper — 根据当前机器架构选择对应二进制
# 用法: ur-wrapper.sh <app> <command> [args...]
# 示例: ur-wrapper.sh iot api /api/v1/system/user/self/get-one

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
APP="${1:-}"
shift || true

if [[ -z "$APP" ]]; then
  echo "usage: $(basename "$0") <app> <command> [args...]" >&2
  echo "  apps: platform-manage, iot, org-manage, org-energy, console" >&2
  exit 2
fi

ARCH_RAW="$(uname -m)"
OS_RAW="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "${OS_RAW}-${ARCH_RAW}" in
  linux-x86_64)  ARCH="x64-linux" ;;
  linux-aarch64) ARCH="arm64-linux" ;;
  darwin-x86_64) ARCH="x64-mac" ;;
  darwin-arm64)  ARCH="arm64-mac" ;;
  msys-x86_64|mingw-x86_64|cygwin-x86_64) ARCH="x64-win" ;;
  *)
    echo "unsupported architecture: ${OS_RAW}-${ARCH_RAW}" >&2
    exit 1
    ;;
esac

BINARY="${SCRIPT_DIR}/${ARCH}/bin/ur-${APP}"
if [[ ! -x "$BINARY" ]]; then
  echo "binary not found: $BINARY" >&2
  echo "available architectures:" >&2
  ls -1 "${SCRIPT_DIR}" | grep -E '^[a-z]+-[a-z0-9]+$' >&2 || true
  exit 1
fi

exec "$BINARY" "$@"
WRAPPER
chmod +x "${ARCH_AWARE_WRAPPER}"

# 同步 skill 文档到 npm-package/（取第一个架构的 skill 即可，文档与架构无关）
FIRST_ARCH_PAIR="${ARCH_LIST[0]}"
FIRST_DIST_NAME="$(goarch_to_distname "$FIRST_ARCH_PAIR")"
NPM_PKG_DIR="${ROOT}/npm-package"
if [[ -d "${OUT_DIR}/${FIRST_DIST_NAME}/skill" && -d "${NPM_PKG_DIR}" ]]; then
  echo "[package-skill] syncing skill docs to npm-package/"
  # 清理旧文件（保留 package.json, index.js, index.d.ts, scripts/）
  for entry in "${NPM_PKG_DIR}"/*; do
    name="$(basename "$entry")"
    if [[ "$name" == "package.json" || "$name" == "index.js" || "$name" == "index.d.ts" || "$name" == "scripts" ]]; then
      continue
    fi
    rm -rf "$entry"
  done
  cp -R "${OUT_DIR}/${FIRST_DIST_NAME}/skill/." "${NPM_PKG_DIR}/"
  echo "[package-skill] npm-package synced"
fi

echo "[package-skill] all done. output: ${OUT_DIR}"
echo "[package-skill] architectures:"
find "${OUT_DIR}" -maxdepth 1 -type d | grep -v "^${OUT_DIR}$" | sort | while read -r d; do
  dist_name="$(basename "$d")"
  go_pair="$(distname_to_goarch "$dist_name")"
  echo "  ${dist_name} (go: ${go_pair})"
done
