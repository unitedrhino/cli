# CLI 用法

## 全局选项

```bash
ur --version
ur -v

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
