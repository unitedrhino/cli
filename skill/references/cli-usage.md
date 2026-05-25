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
ur login                          # Device Auth 授权
ur login --no-wait --json         # AI 模式：获取 URL
ur login --setup-code ABC --json  # 完成授权
ur setup                          # 交互式配置
ur check                          # 验证连通性
ur config --list                  # 管理多环境配置
ur config --use prod
ur token --decode                 # 查看当前 token
ur token --raw
```

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

# 临时覆盖连接配置
UR_BASE_URL=http://host:7777 UR_APP_ID=77 UR_TENANT_CODE=platform ur check
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
