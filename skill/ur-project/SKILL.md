---
name: ur-project
description: "项目管理：项目CRUD、区域管理、分组管理、数据权限、用户权限申请。triggers: 项目管理, 区域管理, 设备分组, 数据权限, 项目列表, 区域授权, 用户权限申请, 项目区域"
metadata:
  hermes:
    tags: [project, area, group, permission]
---


# ur-project — 项目管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/lib/swagger.ts` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 |
|------|------|
| **项目** | 业务项目，包含区域、设备，租户级隔离 |
| **区域** | 项目下的地理位置/逻辑分区，树形结构 |
| **分组** | 设备分组，跨区域的设备归类 |
| **数据权限** | 用户对项目/区域的访问权限 |
| **权限分层** | platform租户管理员 > 租户管理员 > 项目管理员 > 区域管理员 > 普通用户 |
| **AuthType等级** | 1=管理权限（可授权），2=读写权限，3=只读权限，4=区域管理员 |
| **权限继承规则** | 用户必须先有项目权限才能申请区域权限，区域权限不能超过项目权限等级 |

## CLI 命令参考

| 功能组 | 说明 | 参考文档 |
|--------|------|---------|
| 项目管理 | 项目查询/创建/更新/删除 | [project-info.md](references/project-info.md) |
| 项目配置 | 项目配置查询/更新 | [project-profile.md](references/project-profile.md) |
| 项目分组 | 分组查询/创建/删除，设备批量添加 | [project-group.md](references/project-group.md) |
| 区域管理 | 区域查询/创建/更新/删除 | [area.md](references/project-area.md) |
| 区域配置 | 区域配置查询/更新 | [area.md](references/project-area.md) |

> 完整命令帮助：`ur things project help`

## 典型业务场景

### 项目与区域管理

**场景描述**：创建项目 / 划分区域 / 分配设备

**涉及 CLI**：
- `ur things project info create`
- `ur things area info create`
- `ur things area info get-list`

**工作流**：
1. 创建项目
2. 创建区域（树形结构）
3. 配置区域属性
4. 设备归属区域

### 设备分组管理

**场景描述**：创建分组 / 添加设备到分组（跨区域归类）

**涉及 CLI**：
- `ur things group info create`
- `ur things group info get-list`

**工作流**：
1. 创建分组
2. 添加设备到分组
3. 用于批量操作和权限控制

### 数据权限申请与审批

**场景描述**：申请区域权限 / 审批申请 / 权限等级：AuthAdmin(1)/AuthReadWrite(2)/AuthRead(3)

**涉及 CLI**：
- `ur things device info get-list`（查看设备数据权限）
- `ur things project info get-list`（查看项目列表）
- `ur things area info get-list`（查看区域列表）

**工作流**：
1. 用户提交区域权限申请
2. 项目/区域管理员查看待审批列表
3. 审批通过/拒绝
4. 系统校验权限等级约束（区域权限 ≤ 项目权限）
5. 授权生效


## 常用工作流

### 查询项目列表

```bash
ur things project info get-list
```


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
