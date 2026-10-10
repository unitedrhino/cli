# CLI 命令 — 消息管理

## 命令语法
```
ur user self message <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询消息列表 |
| `get-pending` | 查询待处理消息 |
| `handle` | 处理消息 |
| `mark-all-read` | 标记所有消息为已读 |
| `multi-delete` | 批量删除消息 |
| `multi-is-read` | 批量标记消息为已读 |
| `statistics` | 消息统计 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page | | 否 | int | 页码，默认 1 |
| --size | | 否 | int | 每页数量，默认 20 |
| --json | -j | 否 | bool | 输出 JSON 格式 |
| --body | | 条件 | string | 请求体 JSON（handle/multi-delete/multi-is-read 必填） |

## 使用示例

### 示例 1：查询消息列表
```bash
ur user self message get-list --page 1 --size 10
```

### 示例 2：查询待处理消息
```bash
ur user self message get-pending --page 1 --size 10
```

### 示例 3：处理消息
```bash
ur user self message handle --body '{"id":123,"action":"approve"}'
```

### 示例 4：标记所有消息已读
```bash
ur user self message mark-all-read
```

### 示例 5：批量删除消息
```bash
ur user self message multi-delete --body '{"ids":[1,2,3]}'
```

### 示例 6：消息统计
```bash
ur user self message statistics
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| get-list | `POST /api/v1/system/user/self/message/get-list` |
| get-pending | `POST /api/v1/system/user/self/message/get-pending` |
| handle | `POST /api/v1/system/user/self/message/handle` |
| mark-all-read | `POST /api/v1/system/user/self/message/mark-all-read` |
| multi-delete | `POST /api/v1/system/user/self/message/multi-delete` |
| multi-is-read | `POST /api/v1/system/user/self/message/multi-is-read` |
| statistics | `POST /api/v1/system/user/self/message/statistics` |
