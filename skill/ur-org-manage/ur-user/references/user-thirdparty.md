# CLI 命令 — 第三方客户端绑定

## 命令语法
```
ur user self thirdparty <subcommand> [选项]
```

> 旧命令名 `ur user self openclaw` 仍可作为别名使用（行为完全一致），新脚本请统一使用 `thirdparty`。

## 子命令

| 子命令 | 说明 |
|--------|------|
| `setup-check` | 检查第三方客户端绑定状态 |
| `setup-complete` | 完成第三方客户端绑定 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body | | 是 | string | 请求体 JSON |
| --json | -j | 否 | bool | 输出 JSON 格式 |

## 使用示例

### 示例 1：检查第三方客户端绑定状态
```bash
ur user self thirdparty setup-check --body '{"deviceCode":"xxx"}'
```

### 示例 2：完成第三方客户端绑定
```bash
ur user self thirdparty setup-complete --body '{"deviceCode":"xxx","userCode":"yyy"}'
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| setup-check | `POST /api/v1/system/user/self/thirdparty/setup-check` |
| setup-complete | `POST /api/v1/system/user/self/thirdparty/setup-complete` |
