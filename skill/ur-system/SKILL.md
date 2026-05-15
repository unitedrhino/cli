---
name: ur-system
description: "系统管理：通用接口、应用管理、访问令牌、WebSocket推送、授权码、Hook能力。triggers: 系统管理, 文件上传, WebSocket, 批量接口, 应用管理, 菜单管理, Hook配置, 授权码, License"
metadata:
  hermes:
    tags: [system, websocket, file, license]
---


# ur-system — 系统管理

> **本文件由 `ur generate-skills` 自动生成，请勿手动编辑**
>
> 如需更新，修改 `scripts/lib/swagger.ts` 中的域名定义后重新生成。

## 核心概念

| 概念 | 说明 | 注意 |
|------|------|------|
| **AccessKey** | 令牌唯一标识（16位），用于查询和区分令牌 |  |
| **AccessSecret** | 令牌认证密钥（32位），仅创建时返回一次 | 需妥善保存！ |
| **Bearer JWT** | 用 AccessKey/Secret 自行构造 JWT，Authorization: Bearer 调用 API |  |
| **会话 Token** | 登录获取的 session token，token: header，短期有效 |  |
| **应用** | 系统应用（Web、App 等），tenantCode='common' 为全局应用 |  |
| **菜单** | 应用菜单配置，树形结构 |  |
| **WebSocket 频道** | namespace.type.seg1.seg2... 点分层级格式 |  |
| **频道通配符** | * 匹配单段，> 匹配多段（借鉴 NATS） |  |
| **Hook** | 扩展能力（Webhook） |  |
| **授权码** | 许可证管理 |  |

## 角色权限说明

本域 API 按调用者角色区分权限：

| 角色 | 权限范围 | 典型场景 |
|------|---------|----------|
| 平台管理员 | 创建全局应用、管理应用上下架、配置全局 Hook | 创建全局应用、管理应用上下架 |
| 租户管理员 | 订阅应用到租户、自定义菜单配置、使用通用接口 | 订阅应用到租户、文件上传 |
| 普通用户 | 文件上传、WebSocket 连接、批量接口调用 | 文件上传、批量接口调用 |

---

### 平台管理员视角

**权限范围**：创建全局应用、管理应用上下架、配置全局 Hook

#### 典型场景

**创建全局应用**

创建全局应用（tenantCode='common'），定义菜单树形结构

- 涉及 API: `/api/v1/system/app/info/create`, `/api/v1/system/app/menu/create`
- 工作流: 创建应用 → 定义菜单树 → 配置登录方式

**管理应用上下架**

控制应用的启用/禁用状态

- 涉及 API: `/api/v1/system/app/info/update`
- 工作流: 选择应用 → 更新状态


### 租户管理员视角

**权限范围**：订阅应用到租户、自定义菜单配置、使用通用接口

#### 典型场景

**订阅应用到租户**

为租户启用应用，自定义菜单配置

- 涉及 API: `/api/v1/system/tenant/app/create`, `/api/v1/system/tenant/app/menu/update`
- 工作流: 选择要启用的应用 → 自定义菜单（名称、图标、排序） → 配置登录方式

**文件上传**

上传文件到 OSS，获取访问路径

- 涉及 API: `/api/v1/system/common/upload-file`
- 工作流: 选择文件 → 上传 → 返回文件路径


### 普通用户视角

**权限范围**：文件上传、WebSocket 连接、批量接口调用

#### 典型场景

**文件上传**

上传文件 / 获取上传地址

- 涉及 API: `/api/v1/system/common/upload-file`, `/api/v1/system/common/upload-url/create`
- 工作流: 获取上传 URL 或直接上传 → 返回文件访问路径

**批量接口调用**

一次请求调用多个 API，减少网络开销

- 涉及 API: `/api/v1/system/common/api/batch-agg`
- 工作流: 构造批量请求体 → 一次请求获取多个接口数据


## API 参考

<!-- API_LIST:ur-system -->

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/agreement/create` | 创建协议 | platform |
| POST | `/api/v1/system/agreement/delete` | 删除协议 | platform |
| POST | `/api/v1/system/agreement/get-list` | 获取协议列表 | platform |
| POST | `/api/v1/system/agreement/get-one` | 获取协议详情 | platform |
| POST | `/api/v1/system/agreement/preview-render` | 预览渲染协议 | platform |
| POST | `/api/v1/system/agreement/update` | 更新协议 | platform |
| POST | `/api/v1/system/app/agreement/bind-batch-update` | 批量更新应用协议绑定 | platform |
| POST | `/api/v1/system/app/agreement/get-bind-list` | 获取应用协议绑定列表 | platform |
| POST | `/api/v1/system/app/core/get-one` | 无需登录获取应用信息 | public |
| POST | `/api/v1/system/app/info/create` | 添加应用 | platform |
| POST | `/api/v1/system/app/info/delete` | 删除应用 | platform |
| POST | `/api/v1/system/app/info/get-list` | 获取应用列表 | platform |
| POST | `/api/v1/system/app/info/get-one` | 获取应用详情 | platform |
| POST | `/api/v1/system/app/info/update` | 更新应用 | platform |
| POST | `/api/v1/system/app/menu/create` | 添加菜单 | platform |
| POST | `/api/v1/system/app/menu/delete` | 删除菜单 | platform |
| POST | `/api/v1/system/app/menu/get-list` | 获取菜单列表 | platform |
| POST | `/api/v1/system/app/menu/update` | 更新菜单 | platform |
| POST | `/api/v1/system/check-in/do` | 用户签到 | admin |
| POST | `/api/v1/system/check-in/get-list` | 签到记录列表 | admin |
| POST | `/api/v1/system/check-in/point-balance/get` | 获取当前用户积分余额 | admin |
| POST | `/api/v1/system/check-in/point-log/adjust` | 管理员调整积分 | admin |
| POST | `/api/v1/system/check-in/point-log/get-list` | 积分流水列表 | admin |
| POST | `/api/v1/system/common/api/batch-agg` | 批量聚合接口请求 | all |
| GET | `/api/v1/system/common/debug` | 调试接口GET | public |
| POST | `/api/v1/system/common/debug` | 调试接口POST | public |
| GET | `/api/v1/system/common/debug-tencent` | 腾讯云调试接口 | public |
| GET | `/api/v1/system/common/download-file` | 下载本地文件 | public |
| POST | `/api/v1/system/common/init-upload-file` | 初始化上传文件 | public |
| POST | `/api/v1/system/common/ntp/get-one` | ntp时间同步 | public |
| POST | `/api/v1/system/common/qr-code/get-one` | 获取小程序二维码 | all |
| POST | `/api/v1/system/common/sys-config/core/get-one` | 读取系统配置信息(无需登录) | public |
| POST | `/api/v1/system/common/sys-config/info/get-one` | 读取系统配置信息 | platform |
| POST | `/api/v1/system/common/sys-config/info/update` | 更新系统配置信息 | platform |
| POST | `/api/v1/system/common/system/init` | 初始化系统 | public |
| POST | `/api/v1/system/common/third/dept/get-list` | 获取第三方部门列表 | all |
| POST | `/api/v1/system/common/third/dept/get-one` | 获取第三方部门详情 | all |
| POST | `/api/v1/system/common/upload-file` | 文件直传 | all |
| POST | `/api/v1/system/common/upload-url/create` | 获取文件上传地址 | all |
| POST | `/api/v1/system/common/weather/get-one` | 获取天气情况 | all |
| GET | `/api/v1/system/common/websocket/connect` | websocket连接 | all |
| POST | `/api/v1/system/hook/capability/create` | 新增Hook能力 | platform |
| POST | `/api/v1/system/hook/capability/delete` | 删除Hook能力 | platform |
| POST | `/api/v1/system/hook/capability/get-list` | 获取Hook能力列表 | platform |
| POST | `/api/v1/system/hook/capability/update` | 更新Hook能力 | platform |
| POST | `/api/v1/system/hook/server/create` | 新增Hook服务 | platform |
| POST | `/api/v1/system/hook/server/delete` | 删除Hook服务 | platform |
| POST | `/api/v1/system/hook/server/get-list` | 获取Hook服务列表 | platform |
| POST | `/api/v1/system/hook/server/get-one` | 获取Hook服务详情 | platform |
| POST | `/api/v1/system/hook/server/update` | 更新Hook服务 | platform |
| POST | `/api/v1/system/job/task/cancel` | 取消执行任务 | platform |
| POST | `/api/v1/system/job/task/group/create` | 新建任务分组 | platform |
| POST | `/api/v1/system/job/task/group/delete` | 删除任务分组 | platform |
| POST | `/api/v1/system/job/task/group/get-list` | 获取任务分组列表 | platform |
| POST | `/api/v1/system/job/task/group/get-one` | 获取任务分组详情 | platform |
| POST | `/api/v1/system/job/task/group/update` | 更新任务分组 | platform |
| POST | `/api/v1/system/job/task/info/create` | 创建任务 | platform |
| POST | `/api/v1/system/job/task/info/delete` | 删除任务 | platform |
| POST | `/api/v1/system/job/task/info/get-list` | 获取任务列表 | platform |
| POST | `/api/v1/system/job/task/info/get-one` | 获取任务详情 | platform |
| POST | `/api/v1/system/job/task/info/start` | 启动任务 | platform |
| POST | `/api/v1/system/job/task/info/stop` | 停止任务 | platform |
| POST | `/api/v1/system/job/task/info/update` | 更新任务 | platform |
| POST | `/api/v1/system/job/task/send` | 发送延时请求 | platform |
| POST | `/api/v1/system/mall/license/batch-create` | 批量创建授权码 | platform |
| POST | `/api/v1/system/mall/license/create` | 创建授权码 | platform |
| POST | `/api/v1/system/mall/license/get-list` | 授权码列表 | admin |
| POST | `/api/v1/system/mall/license/get-one` | 授权码详情 | admin |
| POST | `/api/v1/system/mall/license/revoke` | 撤销授权码 | platform |
| POST | `/api/v1/system/mall/package/create` | 创建套餐 | platform |
| POST | `/api/v1/system/mall/package/delete` | 删除套餐 | platform |
| POST | `/api/v1/system/mall/package/get-list` | 获取套餐列表 | admin |
| POST | `/api/v1/system/mall/package/get-one` | 获取套餐详情 | admin |
| POST | `/api/v1/system/mall/package/update` | 更新套餐 | platform |
| POST | `/api/v1/system/mall/product/create` | 创建商品 | platform |
| POST | `/api/v1/system/mall/product/delete` | 删除商品 | platform |
| POST | `/api/v1/system/mall/product/get-list` | 获取商品列表 | admin |
| POST | `/api/v1/system/mall/product/get-one` | 获取商品详情 | admin |
| POST | `/api/v1/system/mall/product/update` | 更新商品 | platform |
| POST | `/api/v1/system/ops/feedback/create` | 添加帮助与反馈 | all |
| POST | `/api/v1/system/ops/feedback/get-list` | 获取帮助与反馈 | all |
| POST | `/api/v1/system/ops/feedback/update` | 更新帮助与反馈 | all |
| POST | `/api/v1/system/ops/work-order/create` | 添加工单 | all |
| POST | `/api/v1/system/ops/work-order/get-list` | 获取工单列表 | all |
| POST | `/api/v1/system/ops/work-order/update` | 更新工单 | all |
| POST | `/api/v1/system/resource/api/create` | 添加接口 | platform |
| POST | `/api/v1/system/resource/api/delete` | 删除接口 | platform |
| POST | `/api/v1/system/resource/api/get-list` | 获取接口列表 | platform |
| POST | `/api/v1/system/resource/api/update` | 更新接口 | platform |

<!-- END_API_LIST -->

权限: mixed

| 端点 | 说明 | 权限 |
|---|------|------|
| POST /api/v1/system/app/core/get-one | 无需登录获取应用信息 | 平台管理员 |
| POST /api/v1/system/agreement/create | 创建协议 | 平台管理员 |
| POST /api/v1/system/agreement/get-list | 获取协议列表 | 平台管理员 |
| POST /api/v1/system/agreement/get-one | 获取协议详情 | 平台管理员 |
| POST /api/v1/system/agreement/update | 更新协议 | 平台管理员 |
| POST /api/v1/system/agreement/delete | 删除协议 | 平台管理员 |
| POST /api/v1/system/agreement/preview-render | 预览渲染协议 | 平台管理员 |
| POST /api/v1/system/app/info/create | 添加应用 | 平台管理员 |
| POST /api/v1/system/app/info/delete | 删除应用 | 平台管理员 |
| POST /api/v1/system/app/info/get-list | 获取应用列表 | 平台管理员 |
| POST /api/v1/system/app/info/get-one | 获取应用详情 | 平台管理员 |
| POST /api/v1/system/app/info/update | 更新应用 | 平台管理员 |
| POST /api/v1/system/app/agreement/get-bind-list | 获取应用协议绑定列表 | 平台管理员 |
| POST /api/v1/system/app/agreement/bind-batch-update | 批量更新应用协议绑定 | 平台管理员 |
| POST /api/v1/system/app/menu/create | 添加菜单 | 平台管理员 |
| POST /api/v1/system/app/menu/delete | 删除菜单 | 平台管理员 |
| POST /api/v1/system/app/menu/get-list | 获取菜单列表 | 管理员 |
| POST /api/v1/system/app/menu/update | 更新菜单 | 平台管理员 |
| POST /api/v1/system/common/api/batch-agg | 批量聚合接口请求 | 所有用户 |
| POST /api/v1/system/common/debug | 调试接口POST | 所有用户 |
| POST /api/v1/system/common/debug-tencent | 腾讯云调试接口 | 所有用户 |
| POST /api/v1/system/common/download-file | 下载本地文件 | 所有用户 |
| POST /api/v1/system/common/init-upload-file | 初始化上传文件 | 所有用户 |
| POST /api/v1/system/common/ntp/get-one | ntp时间同步 | 所有用户 |
| POST /api/v1/system/common/qr-code/get-one | 获取小程序二维码 | 所有用户 |
| POST /api/v1/system/common/sys-config/core/get-one | 读取系统配置信息(无需登录) | 所有用户 |
| POST /api/v1/system/common/sys-config/info/get-one | 读取系统配置信息 | 平台管理员 |
| POST /api/v1/system/common/sys-config/info/update | 更新系统配置信息 | 平台管理员 |
| POST /api/v1/system/common/system/init | 初始化系统 | 所有用户 |
| POST /api/v1/system/common/third/dept/get-list | 获取第三方部门列表 | 所有用户 |
| POST /api/v1/system/common/third/dept/get-one | 获取第三方部门详情 | 所有用户 |
| POST /api/v1/system/common/upload-file | 文件直传 | 所有用户 |
| POST /api/v1/system/common/upload-url/create | 获取文件上传地址 | 所有用户 |
| POST /api/v1/system/common/weather/get-one | 获取天气情况 | 所有用户 |
| POST /api/v1/system/common/websocket/connect | websocket连接 | 所有用户 |
| POST /api/v1/system/hook/capability/create | 新增Hook能力 | 平台管理员 |
| POST /api/v1/system/hook/capability/delete | 删除Hook能力 | 平台管理员 |
| POST /api/v1/system/hook/capability/get-list | 获取Hook能力列表 | 平台管理员 |
| POST /api/v1/system/hook/capability/update | 更新Hook能力 | 平台管理员 |
| POST /api/v1/system/hook/server/create | 新增Hook服务 | 平台管理员 |
| POST /api/v1/system/hook/server/delete | 删除Hook服务 | 平台管理员 |

## 协议与登录页场景

### 协议资源分层

- `system/agreement`：协议本体（协议库 CRUD + 模板预览）
- `system/app/agreement`：应用与协议的绑定关系

### 登录页展示协议

登录页或未登录公开页需要展示协议时，走：

- `POST /api/v1/system/app/core/get-one`

并传：

```json
{
  "id": "100",
  "withAgreementCodes": ["user-agreement", "privacy-policy"]
}
```

从 `agreementsMap` 中按 code 读取协议内容。
| POST /api/v1/system/hook/server/get-list | 获取Hook服务列表 | 平台管理员 |
| POST /api/v1/system/hook/server/get-one | 获取Hook服务详情 | 平台管理员 |
| POST /api/v1/system/hook/server/update | 更新Hook服务 | 平台管理员 |
| POST /api/v1/system/license/info/batch-create | 批量创建授权码 | 平台管理员 |
| POST /api/v1/system/license/info/create | 创建授权码 | 平台管理员 |
| POST /api/v1/system/license/info/get-list | 授权码列表 | 管理员 |
| POST /api/v1/system/license/info/get-one | 授权码详情 | 管理员 |
| POST /api/v1/system/license/info/revoke | 撤销授权码 | 平台管理员 |
| POST /api/v1/system/license/info/update | 更新授权码 | 平台管理员 |
| POST /api/v1/system/license/info/use | 使用授权码 | 管理员 |
| POST /api/v1/system/license/info/verify | 验证授权码 | 管理员 |
| POST /api/v1/system/license/package/create | 创建套餐 | 平台管理员 |
| POST /api/v1/system/license/package/delete | 删除套餐 | 平台管理员 |
| POST /api/v1/system/license/package/get-list | 获取套餐列表 | 管理员 |
| POST /api/v1/system/license/package/get-one | 获取套餐详情 | 管理员 |
| POST /api/v1/system/license/package/update | 更新套餐 | 平台管理员 |
| POST /api/v1/system/license/record/get-list | 授权记录列表 | 管理员 |
| POST /api/v1/system/license/record/get-one | 授权记录详情 | 管理员 |
| POST /api/v1/system/tenant/core/get-list | 搜索租户信息 | public |
| POST /api/v1/system/tenant/core/get-one | 获取租户信息 | public |


## 典型业务场景

### 访问令牌管理

**场景描述**：创建 API 访问令牌用于第三方集成

**涉及 API**：
- `/api/v1/system/user/self/access-token/create`
- `/api/v1/system/user/self/access-token/get-list`

**工作流**：
1. 创建令牌获取 AccessKey/Secret
2. 构造 JWT: {userID, tenantCode, accessKey, exp}
3. 使用 Authorization: Bearer <jwt> 调用 API

### WebSocket 实时推送

**场景描述**：连接 WebSocket 获取实时数据推送

**涉及 API**：
- `/api/v1/system/common/websocket/connect`

**工作流**：
1. 携带 token 连接 WebSocket
2. 订阅频道: namespace.type.seg1.seg2
3. 使用通配符 * 或 > 匹配多个频道
4. 接收实时推送消息

### Hook 扩展配置

**场景描述**：配置 Webhook 处理平台事件

**涉及 API**：
- `/api/v1/system/hook/capability/create`
- `/api/v1/system/hook/server/create`

**工作流**：
1. 创建 Hook 能力
2. 配置回调地址
3. 处理事件回调


## 常用工作流

### 批量聚合接口

```bash
ur api /api/v1/system/common/api/batch-agg \
  --body '{"apis":[{"path":"/api/v1/system/user/self/get-one","body":{}}]}'
```


## 注意事项

- 所有接口使用 POST 方法
- 运行时通过 `UR_*` 环境变量注入认证；人工调试可通过 `ur setup` 初始化本地配置
- 请求头需要 `app-id`、`tenant-code` 和认证头
- 响应格式: `{code, msg, data}`，code=200 表示成功
- 分页格式: `{page: {page: 1, size: 10}, total: 0, list: [...]}`
