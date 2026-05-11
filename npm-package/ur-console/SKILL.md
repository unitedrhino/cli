---
name: ur-console
description: "ur-console — 联犀 SaaS 平台 控制台 CLI 工具"
metadata:
  hermes:
    tags: [console, profile, token, settings, personal]
---

# ur-console — 控制台

> **配置检查**：如果尚未配置联犀连接，请先运行 `ur-console login --no-wait`，按指引在浏览器中完成授权。`setup` 命令是终端交互式的，在 AI 聊天环境中无法使用。

## 应用信息

- **AppID**: 77
- **TenantCode**: platform
- **可调用权限**: all, public

## 功能概览

- **控制台**: 应用入口和租户切换
  API: `/api/v1/system/user/self/app/get-list`, `/api/v1/system/tenant/info/get-list`
- **个人信息**: 用户个人设置
  - **修改昵称**: 修改用户昵称
  - **修改密码**: 修改登录密码
  - **绑定账号**: 绑定第三方账号
  - **我的消息**: 查看站内消息
- **访问令牌**: API 访问令牌管理
  - **创建令牌**: 创建 AccessKey/Secret
  - **查看令牌**: 查看已有令牌
  - **删除令牌**: 删除令牌
- **续期管理**: 授权续期和充值
  API: `/api/v1/system/tenant/renewal`

## API 端点

共 89 个可调用端点（按 all/public 权限过滤）。

### system/app/core

- `POST /api/v1/system/app/core/get-one` — 无需登录获取应用信息

### system/config/core

- `POST /api/v1/system/common/sys-config/core/get-one` — 读取系统配置信息(无需登录)

### system/init

- `POST /api/v1/system/common/system/init` — 初始化系统

### system/ops/feedback

- `POST /api/v1/system/ops/feedback/create` — 添加帮助与反馈
- `POST /api/v1/system/ops/feedback/get-list` — 获取帮助与反馈
- `POST /api/v1/system/ops/feedback/update` — 更新帮助与反馈

### system/ops/workOrder

- `POST /api/v1/system/ops/work-order/create` — 添加工单
- `POST /api/v1/system/ops/work-order/get-list` — 获取工单列表
- `POST /api/v1/system/ops/work-order/update` — 更新工单

### system/tenant/app

- `POST /api/v1/system/tenant/app/create` — 绑定租户应用

### system/user/self/accessToken

- `POST /api/v1/system/user/self/access-token/create` — 创建访问令牌
- `POST /api/v1/system/user/self/access-token/delete` — 删除访问令牌
- `POST /api/v1/system/user/self/access-token/get-list` — 获取访问令牌列表
- `POST /api/v1/system/user/self/access-token/get-one` — 获取访问令牌详情
- `POST /api/v1/system/user/self/access-token/update` — 更新访问令牌

### system/user/self

- `POST /api/v1/system/user/self/app/get-list` — 获取用户应用列表
- `POST /api/v1/system/user/self/app/get-one` — 获取用户应用详情
- `POST /api/v1/system/user/self/bind-account` — 绑定账号
- `POST /api/v1/system/user/self/cancel` — 注销用户
- `POST /api/v1/system/user/self/captcha` — 获取验证码
- `POST /api/v1/system/user/self/change-pwd` — 更新用户密码
- `POST /api/v1/system/user/self/forget-pwd` — 忘记密码
- `POST /api/v1/system/user/self/get-one` — 获取用户信息
- `POST /api/v1/system/user/self/login` — 用户登录
- `POST /api/v1/system/user/self/logout` — 用户登出
- `POST /api/v1/system/user/self/menu/get-list` — 获取用户菜单列表
- `POST /api/v1/system/user/self/message/get-list` — 用户消息列表
- `POST /api/v1/system/user/self/message/get-pending` — 用户待处理消息
- `POST /api/v1/system/user/self/message/handle` — 用户消息标记已处理
- `POST /api/v1/system/user/self/message/mark-all-read` — 用户消息全部已读
- `POST /api/v1/system/user/self/message/multi-delete` — 用户消息批量删除
- `POST /api/v1/system/user/self/message/multi-is-read` — 用户消息批量已读
- `POST /api/v1/system/user/self/message/statistics` — 用户消息统计
- `POST /api/v1/system/user/self/notify-preference/read` — 用户通知偏好读取
- `POST /api/v1/system/user/self/notify-preference/update` — 用户通知偏好更新
- `POST /api/v1/system/user/self/profile/get-list` — 获取用户配置列表
- `POST /api/v1/system/user/self/profile/get-one` — 获取用户配置详情
- `POST /api/v1/system/user/self/profile/update` — 更新用户配置
- `POST /api/v1/system/user/self/register` — 普通用户注册
- `POST /api/v1/system/user/self/resource/action/get-list` — 获取用户资源动作权限列表
- `POST /api/v1/system/user/self/update` — 更新用户基本数据
- `POST /api/v1/system/user/self/user/search` — 精准搜索用户

### system/tenant/core

- `POST /api/v1/system/tenant/core/get-list` — 搜索租户信息
- `POST /api/v1/system/tenant/core/get-one` — 获取租户信息

### system/tenant/info

- `POST /api/v1/system/tenant/info/create` — 添加租户

### system/user/self/tenant

- `POST /api/v1/system/user/self/tenant/delete` — 退出当前租户
- `POST /api/v1/system/user/self/tenant/get-list` — 获取用户所处的租户列表
- `POST /api/v1/system/user/self/tenant/get-one` — 获取当前用户在当前租户的详情
- `POST /api/v1/system/user/self/tenant/join` — 用户加入租户（通过邀请码、邮件或手机邀请）
- `POST /api/v1/system/user/self/tenant/update` — 更新当前用户在当前租户的信息

### system/user/tenant

- `POST /api/v1/system/user/tenant/get-list` — 用户所处的租户列表

### things/user/device/collect

- `POST /api/v1/things/user/device/collect/batch-create` — 批量收藏设备
- `POST /api/v1/things/user/device/collect/batch-delete` — 批量取消收藏设备
- `POST /api/v1/things/user/device/collect/get-list` — 获取收藏设备列表

### things/user/device/share

- `POST /api/v1/things/user/device/share/batch-accept` — 接受批量分享设备
- `POST /api/v1/things/user/device/share/batch-create` — 生成批量分享设备二维码
- `POST /api/v1/things/user/device/share/batch-delete` — 批量取消分享设备
- `POST /api/v1/things/user/device/share/batch-get-list` — 获取批量分享的设备列表
- `POST /api/v1/things/user/device/share/create` — 分享设备
- `POST /api/v1/things/user/device/share/delete` — 取消分享设备
- `POST /api/v1/things/user/device/share/get-list` — 获取分享设备列表
- `POST /api/v1/things/user/device/share/get-one` — 获取分享设备详情
- `POST /api/v1/things/user/device/share/update` — 更新分享设备信息

### system/common

- `POST /api/v1/system/common/api/batch-agg` — 批量聚合接口请求
- `GET /api/v1/system/common/debug` — 调试接口GET
- `POST /api/v1/system/common/debug` — 调试接口POST
- `GET /api/v1/system/common/debug-tencent` — 腾讯云调试接口
- `GET /api/v1/system/common/download-file` — 下载本地文件
- `POST /api/v1/system/common/init-upload-file` — 初始化上传文件
- `POST /api/v1/system/common/ntp/get-one` — ntp时间同步
- `POST /api/v1/system/common/qr-code/get-one` — 获取小程序二维码
- `POST /api/v1/system/common/third/dept/get-list` — 获取第三方部门列表
- `POST /api/v1/system/common/third/dept/get-one` — 获取第三方部门详情
- `POST /api/v1/system/common/upload-file` — 文件直传
- `POST /api/v1/system/common/upload-url/create` — 获取文件上传地址
- `POST /api/v1/system/common/weather/get-one` — 获取天气情况
- `GET /api/v1/system/common/websocket/connect` — websocket连接

### system/dict/info

- `POST /api/v1/system/dict/info/get-list` — 获取字典信息列表

### things/ai/mcp

- `POST /api/v1/things/ai/mcp/message` — MCP消息发送
- `POST /api/v1/things/ai/mcp/run` — Stateless MCP HTTP
- `GET /api/v1/things/ai/mcp/sse` — SSE连接
- `POST /api/v1/things/ai/mcp/sse` — SSE连接（POST）

### things/hook

- `POST /api/v1/things/hook/` — Hook扩展统一入口

### things/project/info

- `POST /api/v1/things/project/info/get-list` — 获取项目列表
- `POST /api/v1/things/project/info/get-one` — 获取项目详情

### system/tenant/agreement

- `POST /api/v1/system/tenant/agreement/get-one` — 获取协议详情

### things/area/info

- `POST /api/v1/things/area/info/get-list` — 获取项目区域列表
- `POST /api/v1/things/area/info/get-one` — 获取项目区域详情

### things/user/area/apply

- `POST /api/v1/things/user/area/apply/create` — 申请用户区域权限

## 使用示例

```bash
# 配置
ur-console setup

# 验证连通性
ur-console check

# 调用 API
ur-console api /api/v1/system/user/self/get-one
```
