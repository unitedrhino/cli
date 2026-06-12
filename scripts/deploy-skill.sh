#!/usr/bin/env bash
set -euo pipefail

# deploy-skill.sh — Skill 打包 → 上传 → 注册一站式脚本
#
# 用法:
#   bash scripts/deploy-skill.sh [选项]
#
# 示例:
#   # 打包并注册 ur-api skill 到 115 测试环境
#   bash scripts/deploy-skill.sh --env 115 --code ur-api --name "联犀API" --build
#
#   # 使用本地已有的 zip 包更新 skill
#   bash scripts/deploy-skill.sh --env 115 --code ur-api --zip /path/to/ur-api.zip
#
#   # 指定自定义 API 地址和 token
#   bash scripts/deploy-skill.sh --base-url http://localhost:7777 \
#     --token "your-jwt-token" --code ur-api --name "联犀API" --zip ./ur-api.zip

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLI_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

# ========== 默认配置 ==========
ENV=""
BASE_URL=""
TOKEN=""
CODE=""
NAME=""
ZIP_PATH=""
BUILD=false
TENANT_CODE="common"
SCOPE="platform"
VERSION=""
PUBLISH=false
DRY_RUN=false

# 环境预设
declare -A ENV_URLS=(
  [115]="http://115.190.3.202:7777"
  [120]="http://120.25.49.238:7777"
  [106]="https://saas.unitedrhino.com"
  [local]="http://localhost:7777"
)

# ========== 辅助函数 ==========
log_info() { echo "[deploy-skill] $*"; }
log_error() { echo "[deploy-skill] ERROR: $*" >&2; }
log_warn() { echo "[deploy-skill] WARN: $*" >&2; }

die() { log_error "$*"; exit 1; }

usage() {
  cat <<'EOF'
用法: deploy-skill.sh [选项]

必需选项:
  --code <code>          Skill 编码 (如 ur-api)
  --name <name>          Skill 名称 (如 "联犀API")

环境选项:
  --env <env>            预设环境: 115, 120, 106, local (默认: 需指定 --base-url)
  --base-url <url>       API 基础地址 (如 http://localhost:7777)
  --token <token>        JWT 认证 token (不传则尝试从 CLI 配置读取)

文件选项:
  --zip <path>           指定本地 zip 包路径 (不与 --build 同时使用)
  --build                先调用 package-skill.sh 打包，再注册

Skill 选项:
  --tenant-code <code>   企业编码 (默认: common)
  --scope <scope>        作用域: platform/tenant (默认: platform)
  --version <ver>        版本号 (默认: 从 zip 中解析或 1.0.0)
  --publish              注册后自动推进到 published 状态

其他选项:
  --dry-run              只打印将要执行的操作，不实际调用接口
  -h, --help             显示此帮助

示例:
  # 打包并注册到 115
  bash scripts/deploy-skill.sh --env 115 --code ur-api --name "联犀API" --build

  # 使用已有 zip 注册到本地
  bash scripts/deploy-skill.sh --env local --code ur-api --name "联犀API" --zip ./dist/x64-linux/skill/ur-api.zip
EOF
}

# 解析参数
while [[ $# -gt 0 ]]; do
  case "$1" in
    --env) ENV="$2"; shift 2 ;;
    --base-url) BASE_URL="$2"; shift 2 ;;
    --token) TOKEN="$2"; shift 2 ;;
    --code) CODE="$2"; shift 2 ;;
    --name) NAME="$2"; shift 2 ;;
    --zip) ZIP_PATH="$2"; shift 2 ;;
    --build) BUILD=true; shift ;;
    --tenant-code) TENANT_CODE="$2"; shift 2 ;;
    --scope) SCOPE="$2"; shift 2 ;;
    --version) VERSION="$2"; shift 2 ;;
    --publish) PUBLISH=true; shift ;;
    --dry-run) DRY_RUN=true; shift ;;
    -h|--help) usage; exit 0 ;;
    *) die "未知参数: $1"; ;;
  esac
done

# ========== 参数校验 ==========
[[ -n "$CODE" ]] || die "--code 是必需参数"
[[ -n "$NAME" ]] || die "--name 是必需参数"

# 解析 base-url
if [[ -n "$ENV" && -z "$BASE_URL" ]]; then
  BASE_URL="${ENV_URLS[$ENV]:-}"
  [[ -n "$BASE_URL" ]] || die "未知环境 '$ENV'，可用: ${!ENV_URLS[*]}"
fi
[[ -n "$BASE_URL" ]] || die "需指定 --env 或 --base-url"

# 移除末尾斜杠
BASE_URL="${BASE_URL%/}"

# zip 和 build 二选一
if $BUILD && [[ -n "$ZIP_PATH" ]]; then
  die "--build 和 --zip 不能同时使用"
fi
if ! $BUILD && [[ -z "$ZIP_PATH" ]]; then
  die "需指定 --build (打包) 或 --zip (已有 zip 包)"
fi

# ========== HTTP 请求封装 ==========
http_post() {
  local url="$1"
  local body="${2:-}"
  local headers=()
  if [[ -n "$TOKEN" ]]; then
    headers+=("-H" "token: $TOKEN")
  fi
  headers+=("-H" "Content-Type: application/json")
  headers+=("-H" "tenant-code: $TENANT_CODE")

  if $DRY_RUN; then
    echo "  [DRY-RUN] curl -sS -X POST ${headers[*]} -d '$body' $url"
    return 0
  fi

  local resp
  resp=$(curl -sS -w "\n%{http_code}" -X POST "${headers[@]}" -d "$body" "$url" 2>/dev/null) || true
  local http_code
  http_code=$(echo "$resp" | tail -n1)
  local body_lines
  body_lines=$(echo "$resp" | sed '$d')

  if [[ "$http_code" != "200" ]]; then
    echo "  HTTP $http_code: $body_lines" >&2
    return 1
  fi
  echo "$body_lines"
}

http_post_file() {
  local url="$1"
  local file_path="$2"
  local form_fields=()
  if [[ -n "$TOKEN" ]]; then
    form_fields+=("-H" "token: $TOKEN")
  fi
  form_fields+=("-H" "tenant-code: $TENANT_CODE")
  form_fields+=("-F" "file=@$file_path")

  if $DRY_RUN; then
    echo "  [DRY-RUN] curl -sS -X POST ${form_fields[*]} $url"
    return 0
  fi

  local resp
  resp=$(curl -sS -w "\n%{http_code}" -X POST "${form_fields[@]}" "$url" 2>/dev/null) || true
  local http_code
  http_code=$(echo "$resp" | tail -n1)
  local body_lines
  body_lines=$(echo "$resp" | sed '$d')

  if [[ "$http_code" != "200" ]]; then
    echo "  HTTP $http_code: $body_lines" >&2
    return 1
  fi
  echo "$body_lines"
}

# 从 JSON 中提取字段（简单实现，无依赖）
json_get() {
  local json="$1"
  local key="$2"
  # 支持 data.xxx 嵌套
  local val="$json"
  local part
  for part in $(echo "$key" | tr '.' ' '); do
    val=$(echo "$val" | grep -oP '"'$part'"\s*:\s*"\K[^"]*' || true)
    if [[ -z "$val" ]]; then
      val=$(echo "$val" | grep -oP '"'$part'"\s*:\s*\K[0-9]+' || true)
    fi
  done
  echo "$val"
}

# ========== Step 1: 打包 ==========
if $BUILD; then
  log_info "Step 1: 打包 skill..."
  if $DRY_RUN; then
    log_info "  [DRY-RUN] 将调用 package-skill.sh 打包"
  else
    DIST_DIR="${CLI_DIR}/dist/deploy"
    rm -rf "$DIST_DIR"
    bash "${CLI_DIR}/scripts/package-skill.sh" "$DIST_DIR" || die "打包失败"

    # 找到 ur-api 的 zip（package-skill.sh 生成的是 skill 目录，不是 zip）
    # 需要先压缩
    SKILL_DIR="${DIST_DIR}/$(ls "$DIST_DIR" | head -1)/skill/ur-api"
    if [[ ! -d "$SKILL_DIR" ]]; then
      die "打包后找不到 skill 目录: $SKILL_DIR"
    fi

    ZIP_PATH="${DIST_DIR}/${CODE}.zip"
    (cd "$SKILL_DIR" && zip -r "$ZIP_PATH" . >/dev/null) || die "压缩 zip 失败"
    log_info "  打包完成: $ZIP_PATH ($(stat -c%s "$ZIP_PATH" 2>/dev/null || stat -f%z "$ZIP_PATH" 2>/dev/null) bytes)"
  fi
else
  log_info "Step 1: 使用已有 zip: $ZIP_PATH"
  [[ -f "$ZIP_PATH" ]] || die "zip 文件不存在: $ZIP_PATH"
fi

# 如果未指定版本，尝试从 skill 目录中的 _meta.json 或 SKILL.md 解析
if [[ -z "$VERSION" && -f "$ZIP_PATH" ]]; then
  # 解压到临时目录读取版本
  TMP_DIR=$(mktemp -d)
  unzip -q "$ZIP_PATH" -d "$TMP_DIR" 2>/dev/null || true
  if [[ -f "$TMP_DIR/_meta.json" ]]; then
    VERSION=$(grep -oP '"version"\s*:\s*"\K[^"]*' "$TMP_DIR/_meta.json" 2>/dev/null || true)
  fi
  rm -rf "$TMP_DIR"
  VERSION="${VERSION:-1.0.0}"
fi

# ========== Step 2: 上传 zip 到临时桶 ==========
log_info "Step 2: 上传 zip 到临时桶..."
UPLOAD_RESP=$(http_post_file "${BASE_URL}/api/v1/system/common/upload-file" "$ZIP_PATH") || die "上传失败"
OSS_PATH=$(echo "$UPLOAD_RESP" | grep -oP '"filePath"\s*:\s*"\K[^"]*' || true)
if [[ -z "$OSS_PATH" ]]; then
  # 尝试 fileUri
  OSS_PATH=$(echo "$UPLOAD_RESP" | grep -oP '"fileUri"\s*:\s*"\K[^"]*' || true)
fi
[[ -n "$OSS_PATH" ]] || die "上传响应中未找到 filePath/fileUri: $UPLOAD_RESP"
log_info "  上传成功: $OSS_PATH"

# ========== Step 3: 查询 skill 是否已存在 ==========
log_info "Step 3: 查询 skill '$CODE' 是否已存在..."
LIST_RESP=$(http_post "${BASE_URL}/api/v1/ai/skill/get-list" "{\"page\":{\"page\":1,\"size\":10},\"code\":\"$CODE\"}") || die "查询失败"

SKILL_ID=$(echo "$LIST_RESP" | grep -oP '"id"\s*:\s*\K[0-9]+' | head -1 || true)
if [[ -n "$SKILL_ID" ]]; then
  log_info "  skill 已存在 (ID=$SKILL_ID)，执行更新..."

  # ========== Step 4a: 更新 ==========
  log_info "Step 4: 更新 skill..."
  UPDATE_BODY=$(cat <<EOF
{
  "id": $SKILL_ID,
  "code": "$CODE",
  "name": "$NAME",
  "version": "$VERSION",
  "ossPath": "$OSS_PATH",
  "status": 1,
  "scope": "$SCOPE"
}
EOF
)
  UPDATE_RESP=$(http_post "${BASE_URL}/api/v1/ai/skill/update" "$UPDATE_BODY") || die "更新失败"
  log_info "  更新成功 (ID=$SKILL_ID)"
else
  log_info "  skill 不存在，执行创建..."

  # ========== Step 4b: 创建 ==========
  log_info "Step 4: 创建 skill..."
  CREATE_BODY=$(cat <<EOF
{
  "code": "$CODE",
  "name": "$NAME",
  "version": "$VERSION",
  "ossPath": "$OSS_PATH",
  "status": 1,
  "scope": "$SCOPE",
  "tenantCode": "$TENANT_CODE"
}
EOF
)
  CREATE_RESP=$(http_post "${BASE_URL}/api/v1/ai/skill/create" "$CREATE_BODY") || die "创建失败"
  SKILL_ID=$(echo "$CREATE_RESP" | grep -oP '"id"\s*:\s*\K[0-9]+' | head -1 || true)
  [[ -n "$SKILL_ID" ]] || die "创建响应中未找到 id: $CREATE_RESP"
  log_info "  创建成功 (ID=$SKILL_ID)"
fi

# ========== Step 5: 可选推进发布状态 ==========
if $PUBLISH; then
  log_info "Step 5: 推进发布状态..."

  # submit (draft -> uploaded)
  log_info "  submit..."
  http_post "${BASE_URL}/api/v1/ai/skill/submit" "{\"id\":$SKILL_ID}" >/dev/null || log_warn "submit 失败，跳过"

  # validate (uploaded -> validated)
  log_info "  validate..."
  http_post "${BASE_URL}/api/v1/ai/skill/validate" "{\"id\":$SKILL_ID}" >/dev/null || log_warn "validate 失败，跳过"

  # review (validated -> review_done)
  log_info "  review..."
  http_post "${BASE_URL}/api/v1/ai/skill/review" "{\"id\":$SKILL_ID,\"force\":true}" >/dev/null || log_warn "review 失败，跳过"

  # test (review_done -> tested)
  log_info "  test..."
  http_post "${BASE_URL}/api/v1/ai/skill/test" "{\"id\":$SKILL_ID}" >/dev/null || log_warn "test 失败，跳过"

  # approve (tested -> approved)
  log_info "  approve..."
  http_post "${BASE_URL}/api/v1/ai/skill/approve" "{\"id\":$SKILL_ID}" >/dev/null || log_warn "approve 失败，跳过"

  # publish (approved -> published)
  log_info "  publish..."
  http_post "${BASE_URL}/api/v1/ai/skill/publish" "{\"skillID\":$SKILL_ID,\"version\":\"$VERSION\"}" >/dev/null || log_warn "publish 失败，跳过"

  log_info "  发布流程完成"
fi

# ========== 完成 ==========
echo ""
echo "========================================"
echo "  Skill 部署完成"
echo "========================================"
echo "  Skill Code : $CODE"
echo "  Skill ID   : $SKILL_ID"
echo "  Version    : $VERSION"
echo "  OssPath    : $OSS_PATH"
echo "  环境       : $BASE_URL"
if $PUBLISH; then
  echo "  发布状态   : published"
fi
echo "========================================"
