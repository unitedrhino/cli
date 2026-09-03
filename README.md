# ur CLI

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.23-blue.svg)](https://go.dev/)

联犀 SaaS + IoT 平台命令行工具 — **为 AI Agent 原生设计**，让人类和 AI Agent 都能在终端中操作联犀平台。统一单一入口 `ur`，通过 `--app` 参数切换应用上下文，覆盖平台管理、物联网、组织管理、能源管理等核心业务域，提供 AI Agent [Skills](./skill/)。

[安装](#安装与快速开始) · [Agent Skills](#agent-skills) · [认证](#认证) · [命令](#命令参考) · [进阶用法](#进阶用法)

> **仓库说明**：本项目已从 monorepo（`backend/cli/ur`）迁移为独立仓库。
> - 独立仓库地址：`https://gitee.com/unitedrhino/cli` / `https://github.com/unitedrhino/cli`
> - 原 monorepo 中的 `backend/cli/ur` 已废弃，不再维护

---

## 为什么选 ur CLI？

- **为 Agent 原生设计** — `generate-skills` 一键生成结构化 Skill 文档，AI Agent 无需额外适配即可调用联犀 API
- **统一入口** — 单个 `ur` 二进制，通过 `--app` 参数或 `UR_APP` 环境变量切换 5 大应用域
- **AI 友好调优** — Device Flow 认证流专为 AI 环境优化，`--no-wait` + `--setup-code` 分步授权，Agent 全程无需输入密码
- **全覆盖** — 平台管理、物联网、组织管理、能源管理、控制台 5 大应用域，Swagger 全量 API 自动解析
- **跨平台** — 支持 Linux/macOS/Windows 等主流平台（amd64 / arm64）
- **安全可控** — AccessKey/AccessSecret 本地存储，JWT 签名调用，无密码明文传输

---

## 功能

| 应用 | `--app` 值 | 覆盖域 | 典型能力 |
|------|-----------|--------|---------|
| 平台管理 | `platform-manage` | 企业、用户、应用、角色、权限 | 企业 CRUD、用户管理、应用配置、授权分配 |
| 物联网 | `iot` | 设备、产品、项目、场景、OTA | 设备管理、物模型、项目场景、OTA 升级、协议网关 |
| 组织管理 | `org-manage` | 组织用户、AI 智能体 | 企业内用户/角色/Agent 管理 |
| 能源管理 | `org-energy` | 能耗分析、电力集抄、预付费 | 用能概况、实时监控、预付费充值 |
| 控制台 | `console` | 个人信息、访问令牌 | 个人资料、访问令牌管理 |

---

## Agent Skills

安装 CLI 后，运行 `generate-skills` 生成 AI Agent 可用的 Skill 文档：

```bash
# 为物联网应用生成 Skills
ur --app iot generate-skills --output ./skills/ur-iot
```

生成的 Skill 可直接被 AI Agent 加载，实现零配置调用联犀 API。

| Skill | 说明 |
|-------|------|
| `ur-api` | 通用 API 调用指南、认证方式、角色权限说明（所有 Skill 的基础） |
| `ur-device` | 设备管理 — 设备 CRUD、属性控制、设备分享、物模型 |
| `ur-device-analytics` | 设备数据分析 — 属性历史查询、趋势分析、聚合统计、报表生成（物模型驱动） |
| `ur-device-debug` | 设备调试 — 日志查询（属性/事件/命令/上下线/异常/诊断/SDK）、实时调试（属性控制/行为调用/事件发送） |
| `ur-product` | 产品管理 — 产品定义、物模型、品类管理 |
| `ur-project` | 项目管理 — 项目 CRUD、区域管理、场景编辑 |
| `ur-system` | 系统管理 — 用户管理、角色权限、菜单资源、字典配置 |
| `ur-tenant` | 企业管理 — 企业 CRUD、应用绑定、配额管理 |
| `ur-user` | 用户管理 — 个人信息、企业成员、邀请码 |
| `ur-ai` | AI 管理 — Agent 配置、告警管理 |
| `scene-linkage` | 场景联动 — 规则模板生成与 JSON 校验（if/when/then） |
| `thing-model` | 物模型 — 模板生成、JSON 校验、affordance 定义 |
| `protocol-script` | 协议脚本 — yaegi 脚本模板、Go 代码校验 |

### AI 自助下载 skills（推荐，对所有 AI 工具通用）

AI 工具自己最清楚自身的 skills 目录在哪，CLI 只提供下载能力：先下载并解压 skills 包，再把 `ur-api` 目录整体拷贝到所用 AI 工具的 skills 目录下，重启 AI 工具生效。

```bash
# 下载最新 release 的 skills 包（ur-api-skills-<版本>.zip）并解压，默认输出到 ~/.ur/downloads/
ur skills download

# 指定下载目录 / 直接指定 zip 地址（私有化、离线场景）
ur skills download --output ~/skills-pkg
ur skills download --url "https://example.com/ur-api-skills-v0.4.1.zip"

# JSON 输出（AI 解析 localPath 后自行拷贝）
ur skills download --json
```

`--json` 输出示例：

```json
{"event":"skills_downloaded","downloadUrl":"https://github.com/unitedrhino/cli/releases/download/v0.4.1/ur-api-skills-v0.4.1.zip","localPath":"/home/user/.ur/downloads/ur-api","installHint":"请将上述 ur-api 目录整体拷贝到你所用 AI 工具的 skills 目录下（各 AI 工具的 skills 目录由 AI 自行确认，例如 Claude Code 为 ~/.claude/skills/），拷贝后重启 AI 工具生效"}
```

`ur skills install` 仍是 Claude Code / Codex 的便捷方式（自动探测本机 skills 目录并部署）；其他 AI 工具建议统一走 `ur skills download` 自助下载后自行拷贝。

---

## 安装与快速开始

### 环境要求

- **操作系统**：Linux / macOS / Windows 等
- **Go**：1.23+（仅从源码构建时需要）
- **网络**：可访问联犀 SaaS API 服务器

### 快速开始（人类用户）

> **AI Assistant 注意**：如果你是 AI Agent 且正在协助用户执行安装，请直接跳转到[快速开始（AI Agent）](#快速开始ai-agent)。

#### 安装 CLI

**方式一 — 下载预编译二进制（推荐）：**

```bash
# 1. 确定平台（注意使用 release 资产中的友好平台名）
PLATFORM="Linux-x86_64"    # Linux-x86_64 / Linux-aarch64 / macOS-x86_64 / macOS-arm64 / Windows-x86_64
VERSION="v0.4.1"

# 2. 下载
wget "https://github.com/unitedrhino/cli/releases/download/${VERSION}/ur-cli-${VERSION}-${PLATFORM}.tar.gz"

# 3. 解压完整发布包（包内结构为 <goos>-<goarch>/{ur, skill/}）
#    ur 二进制必须与 skill/ 目录保持同级，否则 ur skills 相关命令找不到内置 skills
mkdir -p ~/.local/lib/ur
tar -xzf "ur-cli-${VERSION}-${PLATFORM}.tar.gz" -C ~/.local/lib/ur --strip-components=1

# 4. 把 ur 暴露到 PATH（软链不影响 skill/ 的查找）
mkdir -p ~/.local/bin && ln -sf ~/.local/lib/ur/ur ~/.local/bin/ur

# Windows: 下载 .zip 包并完整解压，保持 ur.exe 与 skill/ 目录同级，
# 再将 ur.exe 所在目录加入系统 PATH
```

**方式二 — 从源码构建：**

```bash
git clone https://github.com/unitedrhino/cli.git
cd cli
go build -ldflags "-X main.version=$(git describe --tags)" -o dist/bin/ur .
```

#### 版本升级与 Skills

```bash
# 检查是否有新版本（--dry-run 只检查不安装）
ur upgrade --dry-run

# 升级到最新版本（升级成功后内置 skills 会随安装包自动同步；
# 每次运行命令时会自动低频检查（默认每天一次），发现新版本或版本过旧会在终端提醒）
ur upgrade

# 升级并自动把内置 skills 部署到本机各 AI 工具（Claude Code / Codex）
ur upgrade --install-skills

# 手动把内置 ur-api skill 部署到本机各 AI 工具的 skills 目录
# （ur-api 是一个统一 skill，整体部署，不拆分；部署后重启对应 AI 会话即可发现）
ur skills install
ur skills install --dry-run     # 预览目标，不实际写入
ur skills install --json        # JSON 输出
ur skills install --dir <path>  # 指定自定义目标目录（支持 ~ 路径展开）

# 仅下载 skills 包到本地（不部署）：AI 自助下载后自行拷贝到所用 AI 工具的 skills 目录，
# 对所有 AI 工具通用；详见「Agent Skills」章节
ur skills download
ur skills download --output ~/skills-pkg  # 指定下载目录（支持 ~ 路径展开）
ur skills download --url <zip地址>         # 直接指定 skills zip 地址（私有化/离线场景）

# 跳过自动版本检查（环境变量）
export UR_NO_UPDATE_CHECK=1
```

#### 配置与使用

```bash
# 1. 启动 Device Flow 认证（生成绑定码，在浏览器中完成授权）
ur setup

# 2. 验证认证状态
ur check

# 3. 开始使用（默认 org-manage 应用上下文）
ur api /api/v1/system/user/self/get-one

# 4. 切换应用上下文调用 IoT API
ur --app iot api /api/v1/things/device/info/get-list \
  --body '{"page":{"page":1,"size":10}}'

# 5. 生成 AI Agent Skills（可选，推荐）
ur generate-skills --output ./my-skills/
```

---

### 快速开始（AI Agent）

> AI Agent 应自动完成安装和认证，全程无需向用户索取密码。

**一键安装 + 认证（AI 自动执行）**

```bash
# 1. 下载对应平台的 release（AI 根据用户系统自动选择 PLATFORM）
VERSION="v0.4.1"
PLATFORM="Linux-x86_64"   # Linux-x86_64 / Linux-aarch64 / macOS-x86_64 / macOS-arm64 / Windows-x86_64

# Linux / macOS：解压完整发布包（包内为 <goos>-<goarch>/{ur, skill/}），
# 保持 ur 二进制与 skill/ 目录同级，再把 ur 暴露到 PATH
curl -L "https://github.com/unitedrhino/cli/releases/download/${VERSION}/ur-cli-${VERSION}-${PLATFORM}.tar.gz" -o /tmp/ur-cli.tar.gz
mkdir -p ~/.local/lib/ur && tar -xzf /tmp/ur-cli.tar.gz -C ~/.local/lib/ur --strip-components=1
mkdir -p ~/.local/bin && ln -sf ~/.local/lib/ur/ur ~/.local/bin/ur

# Windows (PowerShell)：完整解压 .zip 包，保持 ur.exe 与 skill/ 同级，再把所在目录加入 PATH
# Invoke-WebRequest -Uri "https://github.com/unitedrhino/cli/releases/download/${VERSION}/ur-cli-${VERSION}-${PLATFORM}.zip" -OutFile "ur.zip"; Expand-Archive "ur.zip" -DestinationPath "$env:USERPROFILE\.local\lib\ur"

# 2. 立即生成认证 URL（无需指定 --base-url，默认联犀 SaaS；私有化部署用 --base-url 或 UR_BASE_URL 覆盖）
ur login --no-wait --json
```

输出示例：
```json
{
  "status": "authorization_required",
  "verification_url": "https://saas.unitedrhino.com/#/user/settings?tab=access-tokens&setup=ABC123&redirect=thirdparty",
  "setup_code": "ABC123",
  "expires_in": 600,
  "next_command": "ur login --setup-code ABC123"
}
```

AI 解析 JSON，向用户发送：
> 请在浏览器中打开链接完成 CLI 授权：https://saas.unitedrhino.com/#/user/settings?tab=access-tokens&setup=ABC123&redirect=thirdparty

用户确认在浏览器中点击「完成 CLI 绑定」后，AI 自动执行：

```bash
ur login --setup-code ABC123 --json
```

输出示例：
```json
{
  "event": "authorization_complete",
  "tenant_code": "t1",
  "access_key": "ak_xxxx",
  "access_secret": "sk_xxxx",
  "user_id": "123"
}
```

**验证并生成 Skills**

```bash
ur check
ur generate-skills --output ./skills/
```

---

## 认证

ur CLI 使用 **Device Flow** 认证机制，支持两种模式：

| 命令 | 说明 |
|------|------|
| `setup` | 交互式认证（人类用户）— 阻塞等待浏览器授权完成 |
| `login` | Device Flow 认证（支持 `--no-wait` 非阻塞模式，适合 AI Agent） |
| `check` | 验证当前认证状态和 API 连通性 |
| `token --decode` | 查看并解码当前存储的访问令牌 |

```bash
# 人类模式：一键阻塞授权
ur setup

# Agent 模式：分步授权
ur login --no-wait --json       # 第 1 步：获取 URL 和绑定码
ur login --setup-code ABC123    # 第 2 步：用户确认后完成轮询

# 验证
ur check

# 查看当前 token
ur token --decode
```

---

## 命令参考

### 全局选项

```bash
ur --version                    # 查看 CLI 版本
ur --app iot <command>          # 切换应用上下文（iot / platform-manage / org-manage / org-energy / console）
UR_APP=iot ur <command>         # 通过环境变量切换
```

### API 调用

```bash
# 基本调用
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'

# 输出格式控制
ur api ... --format yaml        # json（默认）/ raw / yaml
ur api ... --transform data.list.0.name   # GJSON 路径提取
ur api ... --output result.json  # 保存到文件
ur api ... --debug              # HTTP 请求/响应详情（敏感头脱敏）
ur api ... --fields code,data.total      # 字段筛选
ur api ... --summarize          # 摘要模式（列表只保留前 5 条）

# 自定义请求头
ur api ... -H "X-Custom-Header: value"

# 从文件读取 body
ur api ... --body-file /tmp/payload.json
```

### 物模型命令

```bash
ur model template property --json      # 生成属性模板
ur model template event --yaml         # 生成事件模板
ur model template action --json        # 生成行为模板
ur model template full --yaml          # 生成完整物模型模板
ur model validate /tmp/model.json      # 校验物模型 JSON
ur model generate-script model.json --mode property --output script.go
```

### 场景联动命令

```bash
ur scene template auto                  # 自动触发场景模板
ur scene template manual               # 手动触发场景模板
ur scene validate /tmp/scene.json      # 校验场景联动 JSON
```

### 协议脚本命令

```bash
ur script template up-before           # 上行前处理模板
ur script template up-after            # 上行后处理模板
ur script template down-before         # 下行前处理模板
ur script template down-after          # 下行后处理模板
ur script validate /tmp/script.go      # 校验协议脚本
```

### Schema 与补全

```bash
ur schema                              # 查看 API schema
ur schema --json                       # JSON 格式输出
ur schema --auth-type admin            # 按权限过滤
ur schema /api/v1/things/device/info/create   # 查看指定接口

ur completion bash >> ~/.bashrc        # bash 补全
ur completion zsh >> ~/.zshrc          # zsh 补全
ur completion fish > ~/.config/fish/completions/ur.fish
```

### 配置管理

```bash
ur setup                               # 交互式配置
ur config --list                       # 列出所有配置
ur config --use prod                   # 切换配置
ur check                               # 验证配置和连通性
```

---

## 进阶用法

### 应用上下文切换

```bash
# 方式一：--app 参数
ur --app iot api /api/v1/things/device/info/get-list
ur --app platform-manage api /api/v1/system/tenant/info/get-list

# 方式二：UR_APP 环境变量
UR_APP=iot ur api /api/v1/things/device/info/get-list

# 方式三：临时覆盖（通过环境变量）
UR_BASE_URL=http://xxx UR_APP_ID=200 UR_TENANT_CODE=platform ur check
```

### 生成 Skills

```bash
# 为当前应用生成所有 Skill 文档
ur --app iot generate-skills

# 输出到指定目录
ur generate-skills --output ./my-skills/
```

生成的 Skill 文档可直接用于 AI Agent 调用联犀 API。

### 运行时环境变量

无需配置文件，直接通过环境变量认证：

```bash
export UR_BASE_URL="https://saas.unitedrhino.com"
export UR_APP_ID="200"
export UR_TENANT_CODE="platform"
export UR_ACCESS_KEY="ak_xxxx"
export UR_ACCESS_SECRET="sk_xxxx"

ur check
ur api /api/v1/things/device/info/get-list
```

---

## 目录结构

```
.
├── main.go                    # CLI 入口
├── cmd/
│   ├── shared/                # 共享命令逻辑（api / check / schema / login / setup / completion / model / scene / script / generate-skills）
│   └── ur/                    # 统一入口（--app 参数解析）
├── internal/
│   ├── config/                # CLIApp 类型 + Profile 配置
│   ├── auth/                  # Device Flow 认证逻辑 + Token 自动刷新
│   ├── client/                # HTTP 客户端（自动重试 + Debug 日志）
│   ├── response/              # 输出格式化（json / raw / yaml + GJSON 提取）
│   └── swagger/               # Swagger 解析
├── skill/                     # 预生成的 Skill 文档（供 AI Agent 使用）
├── shell/
│   ├── push.sh                # 推送当前分支到 origin + gitee
│   ├── pushm.sh               # 强制推送当前分支到两个远程
│   └── tag.sh                 # 打标签并推送到两个远程
├── scripts/
│   ├── release.sh             # 跨平台 Release 构建与发布（封装脚本）
│   ├── generate-api-lists.py  # 从 swagger 自动生成 skill API 端点列表
│   └── update-skills.sh       # 一键更新 skill 并同步到 skills 仓库
├── skill/                     # 预生成的 Skill 文档（供 AI Agent 使用）
├── SKILL_MAINTENANCE.md       # Skill 混合维护模式文档（手写骨架 + 自动生成端点）
└── references/                # 参考文档
```

---

## 开发

```bash
# 运行测试
go test ./...

# 单独测试某个包
go test ./internal/auth/...

# 查看测试覆盖率
go test -cover ./...
```

---

## 发布（维护者）

### 前置条件

需要 GitHub 和 Gitee 的 API token，写入项目根目录的 `.env` 文件（已被 `.gitignore` 忽略，不会提交）：

```bash
# .env 文件内容
GITHUB_TOKEN="ghp_xxxxxxxx"    # GitHub Personal Access Token（需要 repo 权限）
GITEE_TOKEN="xxxxxxxx"         # Gitee 私人令牌
```

`release.sh` 启动时会自动加载 `.env`，无需手动 export。

### 一键发布

```bash
# 语法：bash scripts/release.sh <版本号>
bash scripts/release.sh v0.3.7
```

脚本会自动完成：
1. 构建 39 个平台的二进制（Linux/macOS/Windows/FreeBSD/OpenBSD/NetBSD/Plan9/Solaris/Illumos/DragonFly/AIX × amd64/arm64/386/arm/mips/riscv64 等）
2. 复制 skill 资源到每个平台的发布目录，写入版本元数据 `_meta.json`
3. 打包 tar.gz（Unix）或 zip（Windows）
4. 生成 SHA256 校验和文件 `sha256sums.txt`
5. 创建 GitHub Release 并上传所有资产
6. 创建 Gitee Release 并上传所有资产

### 手动发布（仅某个平台）

如果只需要发布单个平台，可手动运行对应步骤：

```bash
# 临时设置 token
export GITHUB_TOKEN="ghp_xxxxxxxx"
export GITEE_TOKEN="xxxxxxxx"

# 仅构建和发布
bash scripts/release.sh v0.3.7
```

---

## 常见问题

### Q: `setup` 后提示「认证失败」？
A: 检查浏览器是否已完成绑定流程，或尝试 `ur setup --force` 重新认证。

### Q: API 返回「权限不足」？
A: 使用 `--auth-type` 参数切换权限类型，例如 `--auth-type admin`。

### Q: 如何切换应用上下文？
A: 使用 `--app` 参数：`ur --app iot api ...`，或通过 `UR_APP` 环境变量设置。

### Q: Token 过期了怎么办？
A: CLI 会自动刷新 token（使用保存的账号密码重新登录）。如自动刷新失败，运行 `ur login` 重新授权。
