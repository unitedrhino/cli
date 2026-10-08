# CLI 命令 — 邀请管理

## 命令语法

### 邀请用户
```
ur tenant user invite [选项]
```

### 发送邀请
```
ur tenant user invite-send [选项]
```

### 邀请码管理
```
ur tenant user invite-code <subcommand> [选项]
```

### 待处理邀请
```
ur tenant user invite-pending <subcommand> [选项]
```

## 子命令

### 邀请码管理

| 子命令 | 说明 |
|--------|------|
| `gen` | 生成邀请码 |
| `get-one` | 查询邀请码详情 |

### 待处理邀请

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询待处理邀请列表 |
| `delete` | 删除待处理邀请 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body | | 条件 | string | 请求体 JSON（invite/invite-send/invite-code gen 必填） |
| --json | -j | 否 | bool | 输出 JSON 格式 |
| --page | | 否 | int | 页码，默认 1（invite-pending get-list 可用） |
| --size | | 否 | int | 每页数量，默认 20（invite-pending get-list 可用） |
| --id | | 条件 | string | 邀请码/邀请 ID（invite-code get-one/invite-pending delete 必填） |

## 使用示例

### 示例 1：邀请用户
```bash
ur tenant user invite --body '{"account":"newuser","roleID":1}'
```

### 示例 2：发送邀请
```bash
ur tenant user invite-send --body '{"userID":123,"roleID":1}'
```

### 示例 3：生成邀请码
```bash
ur tenant user invite-code gen --body '{"roleID":1,"expireTime":"2026-12-31"}'
```

### 示例 4：查询邀请码详情
```bash
ur tenant user invite-code get-one --id code123
```

### 示例 5：查询待处理邀请列表
```bash
ur tenant user invite-pending get-list --page 1 --size 10
```

### 示例 6：删除待处理邀请
```bash
ur tenant user invite-pending delete --id invite123
```

## 对应 API

| 功能 | 子命令 | API 端点 |
|------|--------|----------|
| 邀请用户 | invite | `POST /api/v1/system/tenant/user/invite` |
| 发送邀请 | invite-send | `POST /api/v1/system/tenant/user/invite-send` |
| 邀请码 | gen | `POST /api/v1/system/tenant/user/invite-code/gen` |
| 邀请码 | get-one | `POST /api/v1/system/tenant/user/invite-code/get-one` |
| 待处理邀请 | get-list | `POST /api/v1/system/tenant/user/invite-pending/get-list` |
| 待处理邀请 | delete | `POST /api/v1/system/tenant/user/invite-pending/delete` |
