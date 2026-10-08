# CLI 命令 — 企业管理（个人中心）

## 命令语法
```
ur user self tenant <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询企业列表 |
| `get-one` | 查询企业详情 |
| `update` | 更新企业信息 |
| `delete` | 删除企业 |
| `join` | 加入企业 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page | | 否 | int | 页码，默认 1 |
| --size | | 否 | int | 每页数量，默认 20 |
| --json | -j | 否 | bool | 输出 JSON 格式 |
| --id | | 条件 | string | 企业 ID（get-one/delete 必填） |
| --body | | 条件 | string | 请求体 JSON（update/join 必填） |

## 使用示例

### 示例 1：查询企业列表
```bash
ur user self tenant get-list --page 1 --size 10
```

### 示例 2：查询企业详情
```bash
ur user self tenant get-one --id tenant123
```

### 示例 3：更新企业信息
```bash
ur user self tenant update --body '{"id":"tenant123","tenantName":"新名称"}'
```

### 示例 4：加入企业
```bash
ur user self tenant join --body '{"tenantCode":"abc123","inviteCode":"xxx"}'
```

### 示例 5：删除企业
```bash
ur user self tenant delete --id tenant123
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| get-list | `POST /api/v1/system/user/self/tenant/get-list` |
| get-one | `POST /api/v1/system/user/self/tenant/get-one` |
| update | `POST /api/v1/system/user/self/tenant/update` |
| delete | `POST /api/v1/system/user/self/tenant/delete` |
| join | `POST /api/v1/system/user/self/tenant/join` |
