# CLI 命令 — CLI 绑定

## 命令语法
```
ur user self openclaw <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `setup-check` | 检查 CLI 绑定状态 |
| `setup-complete` | 完成 CLI 绑定 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body | | 是 | string | 请求体 JSON |
| --json | -j | 否 | bool | 输出 JSON 格式 |

## 使用示例

### 示例 1：检查 CLI 绑定状态
```bash
ur user self openclaw setup-check --body '{"deviceCode":"xxx"}'
```

### 示例 2：完成 CLI 绑定
```bash
ur user self openclaw setup-complete --body '{"deviceCode":"xxx","userCode":"yyy"}'
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| setup-check | `POST /api/v1/system/user/self/openclaw/setup-check` |
| setup-complete | `POST /api/v1/system/user/self/openclaw/setup-complete` |
