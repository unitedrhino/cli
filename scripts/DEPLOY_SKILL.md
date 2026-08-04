# Skill 打包部署指南

## 概述

本文档说明如何将 Skill 从源码打包并注册到目标环境，涵盖：

1. **自动部署**（推荐）：使用 `deploy-skill.sh` 一键完成打包 → 上传 → 注册
2. **手动部署**：分步调用接口，适用于调试或自定义场景
3. **多环境同步**：将 Skill 同步到多个环境（115/106/120 等）

---

> **⚠️ 环境差异重要提示**
>
> 不同环境的对象存储和部署方式可能不同，操作前必须确认：
>
> | 环境 | 存储类型 | 部署方式 | 备注 |
> |------|---------|---------|------|
> | 115 测试 | RustFS | Docker Compose | 可用 `rustfs-import-public.sh` 导入 seed |
> | 106 生产 | 腾讯云 COS | systemd + Docker | **必须用 API 接口更新 skill**，不可直接写文件 |
> | 120 测试 | 视配置而定 | Docker Compose | 需查看 `.env` 中 `OssType` |
>
> **判断方法**：查看目标环境 `.env` 文件中的 `OSS_ENDPOINT`：
> - `127.0.0.1:9000` → RustFS/MinIO
> - `cos.*.myqcloud.com` → 腾讯云 COS
> - `oss-*-internal.aliyuncs.com` → 阿里云 OSS
>
> **核心原则**：skill 的元数据（`ai_skills` 表）必须通过 API 维护，S3 上的文件只是附件。

---

## 一、自动部署（deploy-skill.sh）

### 1.1 前置条件

- `curl` 已安装
- 目标环境可访问
- 拥有平台管理员 token（`tenant-code: platform`）

### 1.2 获取 Token

```bash
# 方式1：通过 CLI 登录获取（推荐）
cd .gits/cli
go run ./cmd/ur-platform-manage login --base-url https://saas.unitedrhino.com

# 查看 token
go run ./cmd/ur-platform-manage token --decode
```

```bash
# 方式2：手动 curl 登录
TOKEN=$(curl -sS -X POST http://115.190.3.202:7777/api/v1/system/user/self/login \
  -H "Content-Type: application/json" \
  -d '{"loginType":"pwd","account":"administrator","password":"SHA256密码","pwdType":1}' | \
  grep -oP '"accessToken":"\K[^"]*')
echo "TOKEN=$TOKEN"
```

### 1.3 基本用法

**场景 A：打包并注册到 115 测试环境**

```bash
cd /home/ubuntu/saas/.gits/cli

bash scripts/deploy-skill.sh \
  --env 115 \
  --token "$TOKEN" \
  --code ur-api \
  --name "联犀API" \
  --build \
  --publish
```

**场景 B：使用已有 zip 注册到本地**

```bash
bash scripts/deploy-skill.sh \
  --env local \
  --token "$TOKEN" \
  --code ur-api \
  --name "联犀API" \
  --zip ./dist/x64-linux/skill/ur-api.zip \
  --publish
```

**场景 C：指定完整参数**

```bash
bash scripts/deploy-skill.sh \
  --base-url http://localhost:7777 \
  --token "$TOKEN" \
  --code ur-iot \
  --name "物联网API" \
  --version "1.2.0" \
  --zip ./ur-iot.zip \
  --tenant-code common \
  --scope platform \
  --publish
```

### 1.4 参数说明

| 参数 | 必需 | 说明 |
|------|------|------|
| `--code` | ✅ | Skill 编码，如 `ur-api`、`ur-iot` |
| `--name` | ✅ | Skill 名称 |
| `--env` | ⚪ | 预设环境：`115`、`120`、`106`、`local` |
| `--base-url` | ⚪ | API 基础地址（与 `--env` 二选一） |
| `--token` | ⚪ | JWT token（不传则接口会返回 401） |
| `--build` | ⚪ | 先调用 `package-skill.sh` 打包 |
| `--zip` | ⚪ | 指定本地 zip 包路径（与 `--build` 二选一） |
| `--version` | ⚪ | 版本号（默认从 zip 的 `_meta.json` 读取，否则 `1.0.0`） |
| `--tenant-code` | ⚪ | 企业编码（默认 `common`） |
| `--scope` | ⚪ | 作用域：`platform` 或 `tenant`（默认 `platform`） |
| `--publish` | ⚪ | 注册后自动推进到 `published` 状态 |
| `--dry-run` | ⚪ | 只打印操作，不实际调用接口 |

### 1.5 脚本执行流程

```
输入参数（code, name, env, token, zip/build）
    ↓
Step 1: 打包（--build 时调用 package-skill.sh）
    ↓
Step 2: POST /api/v1/system/common/upload-file
        上传 zip → 返回临时桶路径
    ↓
Step 3: POST /api/v1/ai/skill/get-list
        查询 skill 是否已存在
    ↓
Step 4a: 已存在 → POST /api/v1/ai/skill/update
         更新 ossPath、version 等
    ↓
Step 4b: 不存在 → POST /api/v1/ai/skill/create
         创建新 skill
    ↓
Step 5（--publish 时）:
    submit → validate → review → test → approve → publish
    ↓
完成
```

### 1.6 环境预设

| 环境 | 地址 |
|------|------|
| `115` | `http://115.190.3.202:7777` |
| `120` | `http://120.25.49.238:7777` |
| `106` | `https://saas.unitedrhino.com` |
| `local` | `http://localhost:7777` |

---

## 二、手动部署（分步接口调用）

适用于调试、自定义参数或脚本无法满足的场景。

### 2.1 步骤总览

```
1. 打包 skill → zip 文件
2. 上传 zip → 获取临时桶路径
3. 查询/创建/更新 skill
4. 推进发布状态（可选）
```

### 2.2 详细步骤

#### Step 1: 打包

```bash
cd /home/ubuntu/saas/.gits/cli
bash scripts/package-skill.sh ./dist/manual

# 找到 skill 目录并压缩
SKILL_DIR="./dist/manual/x64-linux/skill/ur-api"
cd "$SKILL_DIR" && zip -r /tmp/ur-api.zip .
```

#### Step 2: 上传 zip

```bash
# 获取 token（见 1.2 节）
TOKEN="your-jwt-token"

# 上传文件到临时桶
UPLOAD_RESP=$(curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "tenant-code: platform" \
  -F "file=@/tmp/ur-api.zip" \
  http://115.190.3.202:7777/api/v1/system/common/upload-file)

echo "上传响应: $UPLOAD_RESP"
# 示例: {"code":200,"msg":"success","data":{"fileUri":"...","filePath":"temp/xxx.zip"}}

OSS_PATH=$(echo "$UPLOAD_RESP" | grep -oP '"filePath":"\K[^"]*')
echo "OssPath: $OSS_PATH"
```

#### Step 3: 查询 skill 是否存在

```bash
LIST_RESP=$(curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -H "tenant-code: common" \
  -d '{"page":{"page":1,"size":10},"code":"ur-api"}' \
  http://115.190.3.202:7777/api/v1/ai/skill/get-list)

echo "查询响应: $LIST_RESP"
SKILL_ID=$(echo "$LIST_RESP" | grep -oP '"id":\K[0-9]+' | head -1)
echo "Skill ID: ${SKILL_ID:-未找到}"
```

#### Step 4a: 更新已有 skill

```bash
if [[ -n "$SKILL_ID" ]]; then
  curl -sS -X POST \
    -H "token: $TOKEN" \
    -H "Content-Type: application/json" \
    -H "tenant-code: common" \
    -d "{
      \"id\": $SKILL_ID,
      \"code\": \"ur-api\",
      \"name\": \"联犀API\",
      \"version\": \"1.0.1\",
      \"ossPath\": \"$OSS_PATH\",
      \"status\": 1,
      \"scope\": \"platform\"
    }" \
    http://115.190.3.202:7777/api/v1/ai/skill/update
fi
```

#### Step 4b: 创建新 skill

```bash
if [[ -z "$SKILL_ID" ]]; then
  CREATE_RESP=$(curl -sS -X POST \
    -H "token: $TOKEN" \
    -H "Content-Type: application/json" \
    -H "tenant-code: common" \
    -d "{
      \"code\": \"ur-api\",
      \"name\": \"联犀API\",
      \"version\": \"1.0.0\",
      \"ossPath\": \"$OSS_PATH\",
      \"status\": 1,
      \"scope\": \"platform\",
      \"tenantCode\": \"common\"
    }" \
    http://115.190.3.202:7777/api/v1/ai/skill/create)

  echo "创建响应: $CREATE_RESP"
  SKILL_ID=$(echo "$CREATE_RESP" | grep -oP '"id":\K[0-9]+' | head -1)
fi
```

#### Step 5: 推进发布状态

```bash
# submit: draft → uploaded
curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"id\":$SKILL_ID}" \
  http://115.190.3.202:7777/api/v1/ai/skill/submit

# validate: uploaded → validated
curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"id\":$SKILL_ID}" \
  http://115.190.3.202:7777/api/v1/ai/skill/validate

# review: validated → review_done（force=true 跳过 AI 审阅）
curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"id\":$SKILL_ID,\"force\":true}" \
  http://115.190.3.202:7777/api/v1/ai/skill/review

# test: review_done → tested
curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"id\":$SKILL_ID}" \
  http://115.190.3.202:7777/api/v1/ai/skill/test

# approve: tested → approved
curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"id\":$SKILL_ID}" \
  http://115.190.3.202:7777/api/v1/ai/skill/approve

# publish: approved → published
curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"skillID\":$SKILL_ID,\"version\":\"1.0.0\"}" \
  http://115.190.3.202:7777/api/v1/ai/skill/publish
```

---

## 三、多环境同步

### 3.1 场景

Skill 在 115 验证通过后，需要同步到 106 生产环境。

### 3.2 方法一：脚本重复执行

```bash
TOKEN_115="115的token"
TOKEN_106="106的token"
ZIP="./dist/ur-api.zip"

# 部署到 115
bash scripts/deploy-skill.sh --env 115 --token "$TOKEN_115" \
  --code ur-api --name "联犀API" --zip "$ZIP" --publish

# 部署到 106
bash scripts/deploy-skill.sh --env 106 --token "$TOKEN_106" \
  --code ur-api --name "联犀API" --zip "$ZIP" --publish
```

### 3.3 方法二：通过 S3 + rustfs-seed 同步

适用于大批量环境初始化或 Docker 模板更新：

```bash
# 1. 打包并复制到 rustfs-seed
cd .gits/cli
bash scripts/seed-to-rustfs.sh --all

# 2. 发布时 rustfs-import-public.sh 会自动把 seed 导入 S3
# 3. 各环境 rustfs-sync-claw-skills.sh 从 S3 同步到宿主机
```

---

## 四、状态机说明

Skill 发布状态流转：

```
draft（草稿）
  → uploaded（已上传）      [submit]
  → validated（已验证）     [validate]
  → review_done（审阅完成） [review]
  → tested（测试完成）      [test]
  → approved（已审批）      [approve]
  → published（已发布）     [publish]
  → deprecated（已废弃）    [废弃]
```

**只有 `published` 状态的 skill 才能被 Agent 绑定和使用。**

---

## 五、常见问题

### Q1: 上传返回 401

检查 token 是否有效，以及是否正确使用 `token:` header（不是 `Authorization: Bearer`）。

### Q2: 创建返回 "code already exists"

`code` 在企业内唯一。如果要更新已有 skill，用 `--publish` 或直接调用 update 接口。

### Q3: skill 更新后未生效

- 检查 skill 状态是否为 `published`
- Agent 的 system prompt 缓存基于 `updated_at + skill_ids`，更新 skill 后需要等待缓存失效或重启服务
- 如需立即生效，可重启 `allinone` 容器

### Q4: zip 包应该包含什么

最小结构：
```
ur-api/
├── SKILL.md          # Skill 文档（必须）
├── _meta.json        # 元数据（脚本自动生成）
├── .skill-origin.json # 来源信息（脚本自动生成）
├── invoke.sh         # PicoClaw 包装脚本（脚本自动生成）
└── references/       # 参考文档（可选）
```

### Q5: 如何验证 skill 已正确注册

```bash
# 查询 skill 详情
curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"id": 123}' \
  http://115.190.3.202:7777/api/v1/ai/skill/get-one | jq

# 查询 Agent 可用 skill
curl -sS -X POST \
  -H "token: $TOKEN" \
  -H "Content-Type: application/json" \
  -H "tenant-code: common" \
  -d '{"page":{"page":1,"size":10}}' \
  http://115.190.3.202:7777/api/v1/ai/skill/get-list | jq
```

---

## 六、相关文件

| 文件 | 说明 |
|------|------|
| `scripts/deploy-skill.sh` | 自动部署脚本 |
| `scripts/package-skill.sh` | Skill 打包脚本 |
| `scripts/seed-to-rustfs.sh` | 同步到 rustfs-seed |
| `scripts/update-skills.sh` | 从 swagger 更新 skill API 列表 |
| `deploy/scripts/rustfs-sync-claw-skills.sh` | 环境初始化时从 S3 同步 skill |
| `backend/core/service/aisvr/internal/logic/ai/skill/createLogic.go` | 创建 skill 后端逻辑 |
| `backend/core/service/aisvr/internal/logic/ai/skill/updateLogic.go` | 更新 skill 后端逻辑 |
| `backend/core/service/aisvr/internal/logic/ai/skill/ossUtils.go` | OSS 文件处理逻辑 |
