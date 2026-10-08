# CLI 命令 — 通知偏好

## 命令语法
```
ur user self notify-preference <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `read` | 读取通知偏好 |
| `update` | 更新通知偏好 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body | | 条件 | string | 请求体 JSON（update 必填） |
| --json | -j | 否 | bool | 输出 JSON 格式 |

## 使用示例

### 示例 1：读取通知偏好
```bash
ur user self notify-preference read
```

### 示例 2：更新通知偏好
```bash
ur user self notify-preference update --body '{"sms":true,"email":false}'
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| read | `POST /api/v1/system/user/self/notify-preference/read` |
| update | `POST /api/v1/system/user/self/notify-preference/update` |
