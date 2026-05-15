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

## API 参考

<!-- API_LIST:ur-project -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/area/info/create` | 新增项目区域 | admin |
| POST | `/api/v1/things/area/info/delete` | 删除项目区域 | admin |
| POST | `/api/v1/things/area/info/get-list` | 获取项目区域列表 | all |
| POST | `/api/v1/things/area/info/get-one` | 获取项目区域详情 | all |
| POST | `/api/v1/things/area/info/update` | 更新项目区域 | admin |
| POST | `/api/v1/things/area/profile/get-list` | 获取区域配置列表 | admin |
| POST | `/api/v1/things/area/profile/get-one` | 获取区域配置详情 | admin |
| POST | `/api/v1/things/area/profile/update` | 更新区域配置 | admin |
| POST | `/api/v1/things/data/area/batch-delete` | 删除授权区域权限 | admin |
| POST | `/api/v1/things/data/area/batch-update` | 更新授权区域权限 | admin |
| POST | `/api/v1/things/data/area/get-list` | 获取区域权限列表 | admin |
| POST | `/api/v1/things/data/area/user/apply/deal` | 授权区域权限 | admin |
| POST | `/api/v1/things/data/area/user/apply/get-list` | 获取区域权限列表 | admin |
| POST | `/api/v1/things/data/project/batch-create` | 批量创建授权项目权限 | admin |
| POST | `/api/v1/things/data/project/batch-delete` | 批量删除授权项目权限 | admin |
| POST | `/api/v1/things/data/project/create` | 创建授权项目权限 | admin |
| POST | `/api/v1/things/data/project/delete` | 删除授权项目权限 | admin |
| POST | `/api/v1/things/data/project/get-list` | 获取项目权限列表 | admin |
| POST | `/api/v1/things/group/device/batch-create` | 添加分组设备 | admin |
| POST | `/api/v1/things/group/device/batch-delete` | 删除分组设备 | admin |
| POST | `/api/v1/things/group/device/batch-update` | 更新分组设备 | admin |
| POST | `/api/v1/things/group/info/create` | 创建分组 | admin |
| POST | `/api/v1/things/group/info/delete` | 删除分组 | admin |
| POST | `/api/v1/things/group/info/get-list` | 获取分组列表 | admin |
| POST | `/api/v1/things/group/info/get-one` | 获取分组详情信息 | admin |
| POST | `/api/v1/things/group/info/update` | 更新分组信息 | admin |
| POST | `/api/v1/things/project/crud/create` | 新增项目crud | admin |
| POST | `/api/v1/things/project/crud/delete` | 删除项目crud | admin |
| POST | `/api/v1/things/project/crud/get-list` | 获取项目crud列表 | admin |
| POST | `/api/v1/things/project/crud/get-one` | 获取项目crud详情 | admin |
| POST | `/api/v1/things/project/crud/update` | 更新项目crud | admin |
| POST | `/api/v1/things/project/info/create` | 新增项目 | admin |
| POST | `/api/v1/things/project/info/delete` | 删除项目 | admin |
| POST | `/api/v1/things/project/info/get-list` | 获取项目列表 | all |
| POST | `/api/v1/things/project/info/get-one` | 获取项目详情 | all |
| POST | `/api/v1/things/project/info/update` | 更新项目 | admin |
| POST | `/api/v1/things/project/profile/get-list` | 获取项目配置列表 | admin |
| POST | `/api/v1/things/project/profile/get-one` | 获取项目配置详情 | admin |
| POST | `/api/v1/things/project/profile/update` | 更新项目配置 | admin |

<!-- END_API_LIST -->

权限: 管理员

| 端点 | 说明 | 权限 |
|---|------|------|
| POST /api/v1/things/area/info/create | 新增项目区域 | 管理员 |
| POST /api/v1/things/area/info/delete | 删除项目区域 | 管理员 |
| POST /api/v1/things/area/info/get-list | 获取项目区域列表 | 所有用户 |
| POST /api/v1/things/area/info/get-one | 获取项目区域详情 | 所有用户 |
| POST /api/v1/things/area/info/update | 更新项目区域 | 管理员 |
| POST /api/v1/things/area/profile/get-list | 获取区域配置列表 | 管理员 |
| POST /api/v1/things/area/profile/get-one | 获取区域配置详情 | 管理员 |
| POST /api/v1/things/area/profile/update | 更新区域配置 | 管理员 |
| POST /api/v1/things/data/area/batch-delete | 删除授权区域权限 | 管理员 |
| POST /api/v1/things/data/area/batch-update | 更新授权区域权限 | 管理员 |
| POST /api/v1/things/data/area/get-list | 获取区域权限列表 | 管理员 |
| POST /api/v1/things/data/area/user/apply/deal | 授权区域权限 | 管理员 |
| POST /api/v1/things/data/area/user/apply/get-list | 获取区域权限列表 | 管理员 |
| POST /api/v1/things/data/project/batch-create | 批量创建授权项目权限 | 管理员 |
| POST /api/v1/things/data/project/batch-delete | 批量删除授权项目权限 | 管理员 |
| POST /api/v1/things/data/project/create | 创建授权项目权限 | 管理员 |
| POST /api/v1/things/data/project/delete | 删除授权项目权限 | 管理员 |
| POST /api/v1/things/data/project/get-list | 获取项目权限列表 | 管理员 |
| POST /api/v1/things/group/device/batch-create | 添加分组设备 | 管理员 |
| POST /api/v1/things/group/device/batch-delete | 删除分组设备 | 管理员 |
| POST /api/v1/things/group/device/batch-update | 更新分组设备 | 管理员 |
| POST /api/v1/things/group/info/create | 创建分组 | 管理员 |
| POST /api/v1/things/group/info/delete | 删除分组 | 管理员 |
| POST /api/v1/things/group/info/get-list | 获取分组列表 | 管理员 |
| POST /api/v1/things/group/info/get-one | 获取分组详情信息 | 管理员 |
| POST /api/v1/things/group/info/update | 更新分组信息 | 管理员 |
| POST /api/v1/things/hook/ | Hook扩展统一入口 | public |
| POST /api/v1/things/project/crud/create | 新增项目crud | 管理员 |
| POST /api/v1/things/project/crud/delete | 删除项目crud | 管理员 |
| POST /api/v1/things/project/crud/get-list | 获取项目crud列表 | 管理员 |
| POST /api/v1/things/project/crud/get-one | 获取项目crud详情 | 管理员 |
| POST /api/v1/things/project/crud/update | 更新项目crud | 管理员 |
| POST /api/v1/things/project/info/create | 新增项目 | 管理员 |
| POST /api/v1/things/project/info/delete | 删除项目 | 管理员 |
| POST /api/v1/things/project/info/get-list | 获取项目列表 | 所有用户 |
| POST /api/v1/things/project/info/get-one | 获取项目详情 | 所有用户 |
| POST /api/v1/things/project/info/update | 更新项目 | 管理员 |
| POST /api/v1/things/project/profile/get-list | 获取项目配置列表 | 管理员 |
| POST /api/v1/things/project/profile/get-one | 获取项目配置详情 | 管理员 |
| POST /api/v1/things/project/profile/update | 更新项目配置 | 管理员 |
| POST /api/v1/things/user/area/apply/create | 申请用户区域权限 | 所有用户 |


## 典型业务场景

### 项目与区域管理

**场景描述**：创建项目 / 划分区域 / 分配设备

**涉及 API**：
- `/api/v1/things/project/crud/create`
- `/api/v1/things/area/info/create`
- `/api/v1/things/area/profile/update`

**工作流**：
1. 创建项目
2. 创建区域（树形结构）
3. 配置区域属性
4. 设备归属区域

### 设备分组管理

**场景描述**：创建分组 / 添加设备到分组（跨区域归类）

**涉及 API**：
- `/api/v1/things/group/info/create`
- `/api/v1/things/group/device/batch-create`

**工作流**：
1. 创建分组
2. 添加设备到分组
3. 用于批量操作和权限控制

### 数据权限申请与审批

**场景描述**：申请区域权限 / 审批申请 / 权限等级：AuthAdmin(1)/AuthReadWrite(2)/AuthRead(3)

**涉及 API**：
- `/api/v1/things/user/area/apply/create`
- `/api/v1/things/data/area/user/apply/deal`
- `/api/v1/things/data/area/user/get-list`

**工作流**：
1. 用户提交区域权限申请
2. 项目/区域管理员查看待审批列表
3. 审批通过/拒绝
4. 系统校验权限等级约束（区域权限 ≤ 项目权限）
5. 授权生效


## 常用工作流

### 查询项目列表

```bash
ur api /api/v1/things/project/crud/get-list \
  --body '{"page":{"page":1,"size":10}}'
```


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
