# ur CLI 安装与使用教程

联犀 SaaS 平台 API 的 Go 版命令行工具，按前端应用拆分为五个独立二进制。

> **仓库说明**：本项目已从 monorepo（`backend/cli/ur`）迁移为独立仓库。
> - 独立仓库地址：`https://gitee.com/unitedrhino/cli` / `https://github.com/unitedrhino/cli`
> - 原 monorepo 中的 `backend/cli/ur` 已废弃，不再维护

---

## 一、环境要求

- **Go**：1.23+（从源码构建时）
- **操作系统**：Linux / macOS / Windows
- **网络**：可访问联犀 SaaS API 服务器

---

## 发布环境配置（仅维护者）

如需执行 Release 发布，需配置以下环境变量：

```bash
# 1. 复制模板
cp .env.example .env

# 2. 编辑 .env，填入真实 token
# GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
# GITEE_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

| 变量 | 说明 | 获取方式 |
|------|------|---------|
| `GITHUB_TOKEN` | GitHub 个人访问令牌 | GitHub → Settings → Developer settings → Personal access tokens |
| `GITEE_TOKEN` | Gitee 私人令牌 | Gitee → 个人设置 → 私人令牌 → 生成新令牌 |

> `.env` 文件已被 `.gitignore` 忽略，不会被提交到仓库。

---

## 二、安装方式

### 方式一：直接下载预编译二进制（推荐）

从 GitHub 或 Gitee Release 下载对应平台的二进制包，解压后即可使用。**国内用户建议优先使用 Gitee**。

```bash
# 1. 确定平台和架构
# Linux x86_64:   Linux-x86_64
# Linux arm64:    Linux-aarch64
# macOS Intel:    macOS-x86_64
# macOS Apple Silicon: macOS-arm64
# Windows x86_64: Windows-x86_64
PLATFORM="Linux-x86_64"
VERSION="v0.1.0"

# 2. 下载（优先 Gitee，国内更快）
# Gitee 下载
curl -LO "https://gitee.com/unitedrhino/cli/releases/download/${VERSION}/ur-cli-${VERSION}-${PLATFORM}.tar.gz"

# 或 GitHub 下载（Gitee 不可用时备用）
# curl -LO "https://github.com/unitedrhino/cli/releases/download/${VERSION}/ur-cli-${VERSION}-${PLATFORM}.tar.gz"

# 3. 解压并安装
tar -xzf "ur-cli-${VERSION}-${PLATFORM}.tar.gz"
cd "linux-amd64/bin"  # 根据实际目录调整
sudo cp ur-* /usr/local/bin/

# 4. 验证
ur-iot --help
```

> **Windows 用户**：下载 `.zip` 包，解压后将二进制所在目录添加到系统 `PATH` 环境变量。
>
> **完整平台列表**：支持 39 个原生平台（Linux、macOS、Windows、FreeBSD、OpenBSD、NetBSD、AIX、DragonFly、Illumos、Plan9、Solaris 等），详见 Release 页面。

### 方式二：从源码构建（推荐开发者）

```bash
# 1. 克隆仓库
git clone https://gitee.com/unitedrhino/cli.git
cd cli

# 2. 构建全部 CLI 二进制
bash scripts/package-skill.sh

# 3. 构建产物位于 ./bin/ 目录
ls ./bin/
# ur-console  ur-iot  ur-org-energy  ur-org-manage  ur-platform-manage
```

### 方式三：单独构建某个 CLI

```bash
# 例如构建物联网 CLI
go build -o /usr/local/bin/ur-iot ./cmd/ur-iot

# 构建平台管理 CLI
go build -o /usr/local/bin/ur-platform-manage ./cmd/ur-platform-manage
```

### 方式四：安装到系统 PATH

```bash
# 构建并安装到 /usr/local/bin（需要 sudo 权限）
bash scripts/package-skill.sh
sudo cp ./bin/* /usr/local/bin/

# 验证安装
ur-iot --help
```

---

## 三、首次配置（Device Flow 认证）

所有 CLI 工具都需要先完成认证才能调用 API。

### 3.1 启动认证流程

```bash
# 以物联网 CLI 为例
ur-iot setup
```

输出示例：
```
请在浏览器中打开以下链接完成认证：
https://console.unitedrhino.com/#/user/settings?tab=access-tokens&setup=AbC123XyZ&redirect=openclaw

绑定码：AbC123XyZ
正在等待认证完成...
```

### 3.2 浏览器端操作

1. 复制链接到浏览器打开（需先登录联犀控制台）
2. 页面会自动进入 **CLI 绑定向导**
3. 点击「创建访问令牌」→ 输入令牌名称 → 确认创建
4. 系统会自动完成绑定并显示「绑定成功」
5. CLI 终端会同步显示认证成功信息

### 3.3 验证认证状态

```bash
ur-iot check
```

输出示例：
```
认证状态：已认证
用户：张三
租户：default
访问令牌：iot-dev-token（创建于 2026-05-07）
```

---

## 四、应用 CLI 对照表

| CLI 二进制 | 前端应用 | AppID | 默认租户 | 可用权限 |
|-----------|---------|-------|---------|---------|
| `ur-platform-manage` | 平台管理 | 100 | platform | platform, admin, all |
| `ur-iot` | 物联网 | 200 | platform | platform, admin, all |
| `ur-org-manage` | 组织管理 | 300 | 用户输入 | admin, all |
| `ur-org-energy` | 能源管理 | 1000 | 用户输入 | admin, all |
| `ur-console` | 控制台 | 77 | platform | all |

> **提示**：`org-manage` 和 `org-energy` 需要手动指定租户代码，其他 CLI 默认使用 `platform` 租户。

---

## 五、常用命令

### 5.1 基础命令

```bash
# 查看帮助
ur-iot --help

# 查看版本
ur-iot --version

# 检查认证状态
ur-iot check

# 重新认证（会清除旧配置）
ur-iot setup --force
```

### 5.2 API 调用

```bash
# GET 请求
ur-iot api /api/v1/things/device/info/get-list

# POST 请求（带请求体）
ur-iot api /api/v1/things/device/info/get-list \
  --body '{"page":{"page":1,"size":10}}'

# 指定认证类型（admin / platform / all）
ur-iot api /api/v1/system/user/self/get --auth-type admin

# 指定租户（仅 org-manage / org-energy 需要）
ur-org-manage api /api/v1/org/department/get-list --tenant-code myorg
```

### 5.3 查看 API 接口列表

```bash
# 列出当前应用所有可用接口
ur-iot schema

# 带过滤条件
ur-iot schema --auth-type admin
ur-iot schema | grep device
```

### 5.4 生成 Skill 文档

```bash
# 为当前应用生成 Skill 文档（Markdown 格式）
ur-iot generate-skills

# 输出到指定目录
ur-iot generate-skills --output ./my-skills/
```

---

## 六、配置说明

认证配置默认存储在用户目录下：

- **Linux/macOS**：`~/.config/ur-cli/profiles.json`
- **Windows**：`%APPDATA%\ur-cli\profiles.json`

配置文件结构：
```json
{
  "default": {
    "baseUrl": "https://api.unitedrhino.com",
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "tenantCode": "platform",
    "appID": 200
  }
}
```

---

## 七、目录结构

```
.
├── main.go                         # 向后兼容入口（默认 org-manage）
├── cmd/
│   ├── shared/                     # 共享命令逻辑
│   ├── ur-platform-manage/main.go
│   ├── ur-iot/main.go
│   ├── ur-org-manage/main.go
│   ├── ur-org-energy/main.go
│   ├── ur-console/main.go
│   └── ur/                         # 向后兼容包装
├── internal/
│   ├── config/                     # CLIApp 类型 + Profile 配置
│   ├── auth/                       # Device Flow 认证逻辑
│   ├── client/                     # HTTP 客户端
│   └── swagger/                    # Swagger 解析
├── skill/                          # 生成的 Skill 文档
├── references/                     # 参考文档
└── scripts/
    ├── package-skill.sh            # 构建五个二进制
    └── seed-to-rustfs.sh           # 生产部署种子分发
```

---

## 八、开发

```bash
# 运行测试
go test ./...

# 单独测试某个包
go test ./internal/auth/...

# 查看测试覆盖率
go test -cover ./...
```

---

## 九、常见问题

### Q: `setup` 后提示「认证失败」？
A: 检查浏览器是否已完成绑定流程，或尝试 `ur-iot setup --force` 重新认证。

### Q: API 返回「权限不足」？
A: 使用 `--auth-type` 参数切换权限类型，例如 `--auth-type admin`。

### Q: org-manage / org-energy 提示缺少租户？
A: 这两个 CLI 需要显式指定 `--tenant-code`，例如 `ur-org-manage api ... --tenant-code myorg`。

---

## 十、功能指引更新

功能指引数据定义在 `internal/config/app.go` 的 `Features()` 方法中。前端页面变更时，需同步更新对应应用的 `Features()` 方法，然后运行 `generate-skills` 重新生成 Skill 文档。

```bash
go run ./cmd/ur-iot generate-skills
```
