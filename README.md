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
| 平台管理 | `platform-manage` | 租户、用户、应用、角色、权限 | 租户 CRUD、用户管理、应用配置、授权分配 |
| 物联网 | `iot` | 设备、产品、项目、场景、OTA | 设备管理、物模型、项目场景、OTA 升级、协议网关 |
| 组织管理 | `org-manage` | 组织用户、AI 智能体 | 租户内用户/角色/Agent 管理 |
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
| `ur-tenant` | 租户管理 — 租户 CRUD、应用绑定、配额管理 |
| `ur-user` | 用户管理 — 个人信息、租户成员、邀请码 |
| `ur-ai` | AI 管理 — Agent 配置、告警管理 |
| `scene-linkage` | 场景联动 — 规则模板生成与 JSON 校验（if/when/then） |
| `thing-model` | 物模型 — 模板生成、JSON 校验、affordance 定义 |
| `protocol-script` | 协议脚本 — yaegi 脚本模板、Go 代码校验 |

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
# 1. 确定平台和架构
PLATFORM="linux-amd64"    # linux-amd64 / linux-arm64 / darwin-amd64 / darwin-arm64 / windows-amd64
VERSION="v0.3.0"

# 2. 下载
wget "https://github.com/unitedrhino/cli/releases/download/${VERSION}/ur-${VERSION}-${PLATFORM}.tar.gz"

# 3. 解压
mkdir -p ~/.local/bin
tar -xzf "ur-${VERSION}-${PLATFORM}.tar.gz" -C ~/.local/bin --strip-components=1 bin/ur

# Windows: 下载 .zip 包，解压后将 ur.exe 添加到系统 PATH
```

**方式二 — 从源码构建：**

```bash
git clone https://github.com/unitedrhino/cli.git
cd cli
go build -ldflags "-X main.version=$(git describe --tags)" -o dist/bin/ur .
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

> 以下步骤面向 AI Agent，部分步骤需要用户在浏览器中配合完成。

**第 1 步 — 安装 CLI**

```bash
# 下载并安装（示例：Linux x86_64）
VERSION="v0.3.0"
PLATFORM="linux-amd64"
wget "https://github.com/unitedrhino/cli/releases/download/${VERSION}/ur-${VERSION}-${PLATFORM}.tar.gz"
tar -xzf "ur-${VERSION}-${PLATFORM}.tar.gz" -C ~/.local/bin --strip-components=1 bin/ur
```

**第 2 步 — 启动 Device Flow 认证**

> 后台运行命令，它会输出一个授权 URL 和绑定码。提取 URL 和绑定码发送给用户，用户在浏览器中完成授权后，CLI 会自动轮询获取凭证。

```bash
ur login --no-wait --json
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
ur login --setup-code ABC123 --json
```

**第 4 步 — 验证并生成 Skills**

```bash
# 验证认证状态
ur check

# 生成 AI Agent Skills（最终目标）
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
export UR_BASE_URL="https://api.unitedrhino.com"
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
│   └── release.sh             # 跨平台 Release 构建与发布（封装脚本）
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

使用封装好的 release 脚本：

```bash
bash scripts/release.sh v0.3.0
```

脚本会自动完成：
1. 构建 5 个平台的二进制（Linux amd64/arm64、Windows amd64、macOS amd64/arm64）
2. 复制 skill 资源到 dist/
3. 打包 tar.gz / zip
4. 创建 GitHub Release
5. 上传构建产物

如需手动执行，需配置环境变量：

```bash
export GITHUB_TOKEN="ghp_xxxxxxxx"
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
