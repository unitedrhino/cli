# CLI 命令 — 部门用户管理

## 命令语法
```
ur dept user <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `batch-create` | 批量创建部门用户 |
| `batch-delete` | 批量删除部门用户 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body | | 是 | string | 请求体 JSON |
| --json | -j | 否 | bool | 输出 JSON 格式 |

## 使用示例

### 示例 1：批量创建部门用户
```bash
ur dept user batch-create --body '{"deptID":1,"userIDs":[1,2,3]}'
```

### 示例 2：批量删除部门用户
```bash
ur dept user batch-delete --body '{"deptID":1,"userIDs":[1,2,3]}'
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| batch-create | `POST /api/v1/system/user/dept/batch-create` |
| batch-delete | `POST /api/v1/system/user/dept/batch-delete` |
