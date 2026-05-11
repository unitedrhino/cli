# ur CLI

联犀 SaaS 平台 API 的 Go 版命令行工具，按前端应用拆分为五个独立二进制。

> **仓库说明**：本项目已从 monorepo（`backend/cli/ur`）迁移为独立仓库。
> - 独立仓库地址：`https://gitee.com/unitedrhino/cli` / `https://github.com/unitedrhino/cli`
> - 原 monorepo 中的 `backend/cli/ur` 已废弃，不再维护

## 应用 CLI

| CLI 二进制 | 前端应用 | AppID | TenantCode | 可调用权限 |
|-----------|---------|-------|------------|-----------|
| `ur-platform-manage` | 平台管理 | 100 | platform | platform, admin, all |
| `ur-iot` | 物联网 | 200 | platform | platform, admin, all |
| `ur-org-manage` | 组织管理 | 300 | 用户输入 | admin, all |
| `ur-org-energy` | 能源管理 | 1000 | 用户输入 | admin, all |
| `ur-console` | 控制台 | 77 | platform | all |

## 目录结构

```
.
├── main.go                         # 向后兼容入口（默认 org-manage）
├── cmd/
│   ├── shared/                     # 共享命令逻辑
│   ├── ur-platform-manage/main.go
│   ├── ur-iot/main.go
│   ├── ur-org-manage/main.go
│   ├── ur-org-energy/main.go
│   ├── ur-console/main.go
│   └── ur/                         # 向后兼容包装
├── internal/
│   ├── config/
│   │   ├── app.go                  # CLIApp 类型 + 功能指引定义
│   │   └── config.go               # Profile、认证配置
│   ├── auth/                       # 认证逻辑
│   ├── client/                     # HTTP 客户端
│   └── swagger/                    # Swagger 解析
├── skill/                          # 生成的 Skill 文档
├── references/                     # 参考文档
└── scripts/
    ├── package-skill.sh            # 构建五个二进制
    └── seed-to-rustfs.sh           # 生产部署种子分发
```

## 常用命令

```bash
# 构建所有 CLI
bash scripts/package-skill.sh

# 或单独构建
go build -o /tmp/ur-iot ./cmd/ur-iot

# 使用
/tmp/ur-iot setup
/tmp/ur-iot check
/tmp/ur-iot api /api/v1/things/device/info/get-list --body '{"page":{"page":1,"size":10}}'
/tmp/ur-iot schema
/tmp/ur-iot schema --auth-type admin

# 测试
go test ./...
```

## 功能指引

功能指引数据定义在 `internal/config/app.go` 的 `Features()` 方法中。前端页面变更时，需同步更新对应应用的 `Features()` 方法，然后运行 `generate-skills` 重新生成 Skill 文档。

```bash
go run ./cmd/ur-iot generate-skills
```
