# CLI 命令 — 数据权限

## 命令语法

### 项目权限
```
ur data project <subcommand> [选项]
```

### 区域权限
```
ur data area <subcommand> [选项]
```

## 子命令

### 项目权限

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询项目权限列表 |

### 区域权限

| 子命令 | 说明 |
|--------|------|
| `get-list` | 查询区域权限列表 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --page | | 否 | int | 页码，默认 1 |
| --size | | 否 | int | 每页数量，默认 20 |
| --json | -j | 否 | bool | 输出 JSON 格式 |

## 使用示例

### 示例 1：查询项目权限列表
```bash
ur data project get-list --page 1 --size 10
```

### 示例 2：查询区域权限列表
```bash
ur data area get-list --page 1 --size 10
```

## 对应 API

| 功能 | 子命令 | API 端点 |
|------|--------|----------|
| 项目权限 | get-list | `POST /api/v1/system/user/data/project/get-list` |
| 区域权限 | get-list | `POST /api/v1/system/user/data/area/get-list` |
