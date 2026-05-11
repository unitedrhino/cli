# ur CLI

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.23-blue.svg)](https://go.dev/)

联犀 SaaS + IoT 平台命令行工具 — **为 AI Agent 原生设计**，让人类和 AI Agent 都能在终端中操作联犀平台。覆盖平台管理、物联网、组织管理、能源管理等核心业务域，提供 5 个应用 CLI 及 AI Agent [Skills](./skill/)。

[安装](#安装与快速开始) · [Agent Skills](#agent-skills) · [认证](#认证) · [命令](#三层命令调用) · [进阶用法](#进阶用法)

> **仓库说明**：本项目已从 monorepo（`backend/cli/ur`）迁移为独立仓库。
> - 独立仓库地址：`https://gitee.com/unitedrhino/cli` / `https://github.com/unitedrhino/cli`
> - 原 monorepo 中的 `backend/cli/ur` 已废弃，不再维护

---

## 为什么选 ur CLI？

- **为 Agent 原生设计** — `generate-skills` 一键生成结构化 Skill 文档，AI Agent 无需额外适配即可调用联犀 API
- **应用导向** — 按前端应用拆分为 5 个独立 CLI，自动绑定 AppID 和 TenantCode，告别繁琐配置
- **AI 友好调优** — Device Flow 认证流专为 AI 环境优化，`--no-wait` + `--setup-code` 分步授权，Agent 全程无需输入密码
- **全覆盖** — 平台管理、物联网、组织管理、能源管理、控制台 5 大应用域，Swagger 全量 API 自动解析
- **跨平台** — 支持 39 个原生平台（Linux/macOS/Windows/FreeBSD/OpenBSD/NetBSD/AIX 等）
- **安全可控** — AccessKey/AccessSecret 本地存储，JWT 签名调用，无密码明文传输

---

## 功能

| 应用 CLI | 前端应用 | 覆盖域 | 典型能力 |
|----------|---------|--------|---------|
| `ur-platform-manage` | 平台管理 | 租户、用户、应用、角色、权限 | 租户 CRUD、用户管理、应用配置、授权分配 |
| `ur-iot` | 物联网 | 设备、产品、项目、场景、OTA | 设备管理、物模型、项目场景、OTA 升级、协议网关 |
| `ur-org-manage` | 组织管理 | 组织用户、AI 智能体 | 租户内用户/角色/Agent 管理 |
| `ur-org-energy` | 能源管理 | 能耗分析、电力集抄、预付费 | 用能概况、实时监控、预付费充值 |
| `ur-console` | 控制台 | 个人信息、访问令牌 | 个人资料、访问令牌管理 |

---

## Agent Skills

安装 CLI 后，运行 `generate-skills` 生成 AI Agent 可用的 Skill 文档：

```bash
# 为物联网应用生成 Skills
ur-iot generate-skills --output ./skills/ur-iot
```

生成的 Skill 可直接被 AI Agent 加载，实现零配置调用联犀 API。

| Skill | 说明 |
|-------|------|
| `ur-api` | 通用 API 调用指南、认证方式、角色权限说明（所有 Skill 的基础） |
| `ur-iot` | 物联网核心 — 设备管理、产品管理、项目管理、场景联动、OTA、区域管理 |
| `ur-device` | 设备操作 — 设备 CRUD、属性控制、设备分享 |
| `ur-product` | 产品管理 — 产品定义、物模型、品类管理 |
| `ur-project` | 项目管理 — 项目 CRUD、区域管理、场景编辑 |
| `ur-system` | 系统管理 — 用户管理、角色权限、菜单资源、字典配置 |
| `ur-tenant` | 租户管理 — 租户 CRUD、应用绑定、配额管理 |
| `ur-user` | 用户管理 — 个人信息、租户成员、邀请码 |
| `ur-ai` | AI 管理 — Agent 配置、告警管理 |

---

## 安装与快速开始

### 环境要求

- **操作系统**：Linux / macOS / Windows / FreeBSD / OpenBSD / NetBSD 等
- **Go**：1.23+（仅从源码构建时需要）
- **网络**：可访问联犀 SaaS API 服务器

### 快速开始（人类用户）

> **AI Assistant 注意**：如果你是 AI Agent 且正在协助用户执行安装，请直接跳转到[快速开始（AI Agent）](#快速开始ai-agent)。

#### 安装 CLI

以下两种方式**任选其一**：

**方式一 — 下载预编译二进制（推荐）：**

```bash
# 1. 确定平台（示例：Linux x86_64）
PLATFORM="Linux-x86_64"
VERSION="v0.1.0"

# 2. 下载（优先 Gitee，国内更快）
curl -LO "https://gitee.com/unitedrhino/cli/releases/download/${VERSION}/ur-cli-${VERSION}-${PLATFORM}.tar.gz"

# 3. 解压并安装到 PATH
tar -xzf "ur-cli-${VERSION}-${PLATFORM}.tar.gz"
cd linux-amd64/bin
sudo cp ur-* /usr/local/bin/

# 4. 验证
ur-iot --help
```

> **Windows 用户**：下载 `.zip` 包，解压后将目录添加到系统 `PATH`。

**方式二 — 从源码构建：**

```bash
git clone https://gitee.com/unitedrhino/cli.git
cd cli

# 构建全部 CLI
bash scripts/package-skill.sh

# 安装到 PATH
sudo cp ./bin/* /usr/local/bin/
```

#### 配置与使用

```bash
# 1. 启动 Device Flow 认证（生成绑定码，在浏览器中完成授权）
ur-iot setup

# 2. 验证认证状态
ur-iot check

# 3. 开始使用
ur-iot api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'

# 4. 生成 AI Agent Skills（可选，推荐）
ur-iot generate-skills --output ./my-skills/
```

---

### 快速开始（AI Agent）

> 以下步骤面向 AI Agent，部分步骤需要用户在浏览器中配合完成。

**第 1 步 — 安装 CLI**

```bash
# 下载并安装（示例：Linux x86_64）
VERSION="v0.1.0"
PLATFORM="Linux-x86_64"
curl -LO "https://gitee.com/unitedrhino/cli/releases/download/${VERSION}/ur-cli-${VERSION}-${PLATFORM}.tar.gz"
tar -xzf "ur-cli-${VERSION}-${PLATFORM}.tar.gz"
cd linux-amd64/bin
sudo cp ur-* /usr/local/bin/
```

**第 2 步 — 启动 Device Flow 认证**

> 后台运行命令，它会输出一个授权 URL 和绑定码。提取 URL 和绑定码发送给用户，用户在浏览器中完成授权后，CLI 会自动轮询获取凭证。

```bash
ur-iot login --no-wait --json
```

输出示例：
```json
{
  "verification_url": "https://console.unitedrhino.com/#/user/settings?tab=access-tokens&setup=ABC123&redirect=openclaw",
  "setup_code": "ABC123",
  "expires_in": 600
}
```

将 `verification_url` 发送给用户，提示其在浏览器中打开链接并完成 CLI 绑定。

**第 3 步 — 完成授权**

用户确认在浏览器中点击「完成 CLI 绑定」后：

```bash
ur-iot login --setup-code ABC123 --json
```

**第 4 步 — 验证并生成 Skills**

```bash
# 验证认证状态
ur-iot check

# 生成 AI Agent Skills（最终目标）
ur-iot generate-skills --output ./skills/
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
ur-iot setup

# Agent 模式：分步授权
ur-iot login --no-wait --json       # 第 1 步：获取 URL 和绑定码
ur-iot login --setup-code ABC123    # 第 2 步：用户确认后完成轮询

# 验证
ur-iot check

# 查看当前 token
ur-iot token --decode
```

---

## 三层命令调用

### 1. API 快捷调用

```bash
# 查看设备列表
ur-iot api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'

# 获取当前用户信息
ur-iot api /api/v1/system/user/self/get-one

# 指定认证类型（admin / platform / all）
ur-iot api /api/v1/system/user/self/get --auth-type admin
```

### 2. Schema 浏览

```bash
# 列出当前应用所有可用 API
ur-iot schema

# 带过滤
ur-iot schema --auth-type admin
ur-iot schema | grep device
```

### 3. 通用 API 调用（直接指定完整路径）

```bash
# GET 风格（实际仍为 POST）
ur-iot api /api/v1/things/device/info/get-list

# POST 风格（带请求体）
ur-iot api /api/v1/things/device/info/get-list \
  --body '{"page":{"page":1,"size":10}}'
```

---

## 进阶用法

### 指定租户

`ur-org-manage` 和 `ur-org-energy` 需要显式指定租户：

```bash
ur-org-manage api /api/v1/org/department/get-list --tenant-code myorg
```

### 生成 Skills

```bash
# 为当前应用生成所有 Skill 文档
ur-iot generate-skills

# 输出到指定目录
ur-iot generate-skills --output ./my-skills/
```

生成的 Skill 文档可直接用于 AI Agent 调用联犀 API。

### 运行时环境变量

无需配置文件，直接通过环境变量认证：

```bash
export UR_BASE_URL="https://api.unitedrhino.com"
export UR_APP_ID="200"
export UR_TENANT_CODE="platform"
export UR_ACCESS_KEY="ak_xxxx"
export UR_ACCESS_SECRET="sk_xxxx"

ur-iot check
ur-iot api /api/v1/things/device/info/get-list
```

---

## 目录结构

```
.
├── main.go                    # 向后兼容入口（默认 org-manage）
├── cmd/
│   ├── shared/                # 共享命令逻辑
│   ├── ur-platform-manage/    # 平台管理 CLI
│   ├── ur-iot/                # 物联网 CLI
│   ├── ur-org-manage/         # 组织管理 CLI
│   ├── ur-org-energy/         # 能源管理 CLI
│   ├── ur-console/            # 控制台 CLI
│   └── ur/                    # 向后兼容包装
├── internal/
│   ├── config/                # CLIApp 类型 + Profile 配置
│   ├── auth/                  # Device Flow 认证逻辑
│   ├── client/                # HTTP 客户端
│   └── swagger/               # Swagger 解析
├── skill/                     # 预生成的 Skill 文档（供 AI Agent 使用）
├── scripts/
│   ├── package-skill.sh       # 构建五个二进制
│   ├── release.sh             # 跨平台 Release 构建与发布
│   └── seed-to-rustfs.sh      # 生产部署种子分发
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

如需执行 Release 发布，需配置环境变量：

```bash
# ~/.bashrc 中配置（已自动 source）
export GITHUB_TOKEN="ghp_xxxxxxxx"
export GITEE_TOKEN="xxxxxxxx"
```

然后运行：

```bash
bash scripts/release.sh v0.1.0
```

支持自动构建 39 个原生平台并发布到 GitHub + Gitee。

---

## 常见问题

### Q: `setup` 后提示「认证失败」？
A: 检查浏览器是否已完成绑定流程，或尝试 `ur-iot setup --force` 重新认证。

### Q: API 返回「权限不足」？
A: 使用 `--auth-type` 参数切换权限类型，例如 `--auth-type admin`。

### Q: org-manage / org-energy 提示缺少租户？
A: 这两个 CLI 需要显式指定 `--tenant-code`，例如 `ur-org-manage api ... --tenant-code myorg`。

---

## 应用 CLI 对照表

| CLI 二进制 | 前端应用 | AppID | 默认租户 | 可用权限 |
|-----------|---------|-------|---------|---------|
| `ur-platform-manage` | 平台管理 | 100 | platform | platform, admin, all |
| `ur-iot` | 物联网 | 200 | platform | platform, admin, all |
| `ur-org-manage` | 组织管理 | 300 | 用户输入 | admin, all |
| `ur-org-energy` | 能源管理 | 1000 | 用户输入 | admin, all |
| `ur-console` | 控制台 | 77 | platform | all |
