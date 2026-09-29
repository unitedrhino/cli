# CLI 用法

## 全局选项

```bash
ur --version
ur -v
ur --version --json --check-latest

# 切换应用上下文
ur --app iot api /api/v1/things/device/info/get-list
UR_APP=iot ur api /api/v1/things/device/info/get-list
```

## 认证与配置

```bash
# AI 固定先检查；成功时直接使用，不重新登录
ur check --json

# Device Flow（默认 method）
ur login --method device --no-wait --json
ur login --method device --setup-code ABC --json

# 账号密码：秘密优先从环境变量或 stdin 读取
UR_PASSWORD='<原始密码>' ur login --method password \
  --account '<账号>' --tenant-code '<企业编码>' --json
printf '%s' "$UR_PASSWORD" | ur login --method password \
  --account '<账号>' --tenant-code '<企业编码>' --password-stdin --json

# AK/SK：不要求 userID
UR_ACCESS_SECRET='<AccessSecret>' ur login --method aksk \
  --access-key '<AccessKey>' --tenant-code '<企业编码>' --json

ur setup                          # 人类终端兼容向导，不是 AI 首选
ur config --list                  # 管理多环境配置
ur config --use prod
ur token --decode                 # 查看当前 token
ur token --raw
```

`--password`、`--access-secret` 明文参数仅为兼容，可能暴露在 shell 历史或进程列表中。密码必须是原始密码，不要预先 SHA-256。

Sandbox 只需设置 `UR_BASE_URL`、`UR_APP_ID`、`UR_TENANT_CODE` 及以下完整认证组之一：`UR_TOKEN`；`UR_ACCESS_KEY` + `UR_ACCESS_SECRET`；`UR_ACCOUNT` + `UR_PASSWORD`。该模式不读取或改写磁盘 profile。

## CLI 更新提示与处理

业务命令的 JSON 对象可能包含顶层 `_notice`：

- `_notice.update`：CLI 有新版本，读取 `current`、`latest` 和 `command`。
- `_notice.skills`：一个或多个客户端中的 `ur-api` 缺失或版本落后，读取 `target`、`targets` 和 `command`。

看到通知时先完成用户当前请求，不要把 `_notice` 原样作为主要答案，也不要为了提示中断当前业务操作。用户要求升级时执行统一入口：

```bash
ur upgrade
```

该命令更新 CLI、内置 Skills，并把 Skills 部署到自动发现和用户登记的客户端；CLI 已是最新版时仍会刷新客户端副本。只检查使用 `ur upgrade --check --json`，恢复发布资源使用 `ur upgrade --force`，明确只升级 CLI 时使用 `ur upgrade --no-skills`。

自动检查优先读取本地缓存，过期后异步刷新；网络失败不会改变业务命令退出状态。需要纯净机器输出时按需关闭提示：

```bash
UR_NO_UPDATE_NOTIFIER=1 UR_NO_SKILLS_NOTIFIER=1 ur check --json
```

`UR_NO_UPDATE_CHECK=1` 会同时停止远端检查和 CLI 更新提示。

## Skills 多客户端安装与校验（v0.6.1+）

同一份 `ur-api` 可以部署到多个本地 AI 客户端，也可以导出为标准 ZIP，供没有固定本机目录的平台导入。

### 自动识别与安装

```bash
# 先确认 CLI 识别到了哪些目录
ur skills target detect
ur skills target list

# 安装到全部自动识别和已登记目标
ur skills install --all

# 按版本与 SHA256 文件清单检查安装结果
ur skills status
```

CLI 自动识别 Claude Code、Codex 与 WorkBuddy / CodeBuddy 的用户级目录；在 Git 仓库内执行时，还会从**当前工作目录向上**查找仓库根，并识别该项目现有的 `.claude/skills`、`.agents/skills`、`.codebuddy/skills`。执行批量安装前应先运行 `ur skills target detect`，确认没有选中不希望覆盖的项目级 `ur-api`。

`status` 的状态含义：

| 状态 | 含义 | 处理方式 |
|------|------|----------|
| `current` | 版本和文件内容均一致 | 无需处理 |
| `missing` | 目标中没有 `ur-api` | 重新安装 |
| `outdated` | 目标版本落后 | 重新安装或运行 `ur upgrade` |
| `incomplete` | 文件缺失、变化或存在多余文件 | 重新安装恢复完整副本 |

### 登记其他客户端

任何使用本地 Skills 目录的客户端都可以登记，不需要在 CLI 中写死客户端名称：

```bash
ur skills target add my-ai --dir /path/to/client/skills
ur skills target list
ur skills install --all

# 不再使用时只删除登记信息，不删除客户端中的文件
ur skills target remove my-ai
```

长期目标保存在 `~/.ur/skill-targets.json`。自动化环境可使用系统路径分隔符设置多个目录：

```bash
export UR_SKILLS_DIRS=/path/to/client-a/skills:/path/to/client-b/skills
```

WorkBuddy / CodeBuddy 的自定义配置根目录可通过 `CODEBUDDY_CONFIG_DIR` 指定。临时只安装一个或多个明确目录时使用 `ur skills install --dir <目录>`；一旦提供 `--dir`，本次不会写入自动识别和长期登记的其他目标。

### 导出 ZIP

对于不提供本地 Skills 目录、需要通过页面上传或转换能力包的客户端，导出标准 ZIP：

```bash
ur skills export --format zip --output ~/Downloads
```

ZIP 以 `ur-api/` 为根目录，不包含凭证。也可以直接下载 Release 中的 `ur-api-skills-<版本>.zip`。如果豆包、扣子等目标平台要求自己的清单或字段格式，应以该 ZIP 为统一输入再做平台适配，不能假定能够原样导入。完成目录安装或 ZIP 导入后，需要重启对应 AI 客户端，使其重新发现 Skill。

## API 调用

```bash
# 基本调用
ur api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'

# 输出格式控制
ur api ... --format yaml
ur api ... --format raw

# 字段筛选
ur api ... --fields code,data.total,data.list

# 摘要模式（列表只保留前 5 条）
ur api ... --summarize

# GJSON 路径提取
ur api ... --transform data.list.0.deviceName

# 保存输出到文件
ur api ... --output result.json

# 调试模式
ur api ... --debug

# 自定义请求头
ur api ... -H "X-Custom-Header: value"

# 从文件读取 body
ur api /api/v1/things/protocol/script/update --body-file /tmp/payload.json

# 临时 Sandbox 配置（示例使用 Session Token）
UR_BASE_URL='<平台地址>' UR_APP_ID='<应用ID>' \
UR_TENANT_CODE='<企业编码>' UR_TOKEN='<Session Token>' ur check --json
```

## 物模型命令

```bash
ur model template property --json
ur model template event --yaml --output event.yaml
ur model template action --json
ur model template full --yaml --output model.yaml
ur model validate /tmp/model.json
ur model generate-script /tmp/model.json --mode property --output script.go
```

## 场景联动命令

```bash
ur scene template auto
ur scene template manual
ur scene validate /tmp/scene.json
```

## 协议脚本命令

```bash
ur script template up-before
ur script template up-after
ur script template down-before
ur script template down-after
ur script validate /tmp/script.go
```

## Schema 与补全

```bash
ur schema
ur schema --json
ur schema --auth-type admin
ur schema /api/v1/things/device/info/create

# Shell 补全
ur completion bash >> ~/.bashrc
ur completion zsh >> ~/.zshrc
ur completion fish > ~/.config/fish/completions/ur.fish
```

## 输出选项

| 选项 | 说明 | 示例 |
|------|------|------|
| `--format json` | 美化 JSON（默认） | `ur api ... --format json` |
| `--format raw` | 单行 JSON | `ur api ... --format raw` |
| `--format yaml` | YAML 格式 | `ur api ... --format yaml` |
| `--transform PATH` | GJSON 路径提取 | `--transform data.list.0.name` |
| `--fields SELECTORS` | 字段筛选 | `--fields code,data.total` |
| `--summarize` | 摘要模式 | `ur api ... --summarize` |
| `--output FILE` | 保存到文件 | `--output result.json` |
| `--debug` | 打印 HTTP 详情 | `ur api ... --debug` |
| `--header, -H` | 自定义请求头 | `-H "X-Request-ID: abc"` |

> `--fields`、`--summarize`、`--transform` 三者互斥，只能同时用其中一个。
