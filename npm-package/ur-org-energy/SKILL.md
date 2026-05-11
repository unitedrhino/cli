---
name: ur-org-energy
description: "ur-org-energy — 联犀 SaaS 平台 能源管理 CLI 工具"
metadata:
  hermes:
    tags: [energy, power, prepaid, device, consumption, automation]
---

# ur-org-energy — 能源管理

> **配置检查**：如果尚未配置联犀连接，请先运行 `ur-org-energy login --no-wait`，按指引在浏览器中完成授权。`setup` 命令是终端交互式的，在 AI 聊天环境中无法使用。

## 应用信息

- **AppID**: 1000
- **TenantCode**: 用户输入
- **可调用权限**: admin, all, public

## 功能概览

- **大屏**: 能源数据大屏展示
  API: `/api/v1/things/device/msg/property-latest/get-list`
- **设备空间**: 设备控制台、分组和区域管理
  - **控制台**: 设备控制台
  - **设备分组**: 设备分组管理
  - **区域管理**: 区域 CRUD
- **能耗分析**: 多维度能耗分析报表
  - **用能概况**: 总览用能数据
  - **同比环比分析**: 用能同比环比对比
  - **能耗趋势**: 用能趋势图表
  - **能耗排名**: 按区域/设备排名
  - **损耗分析**: 能源损耗分析
  - **用能报表**: 用能数据报表导出
  - **复费率报表**: 分时电价复费率统计
  - **集抄明细报表**: 电表集抄数据明细
  - **一次图**: 电力一次接线图展示
- **电力集抄**: 电力数据采集和报表
  - **数据监控**: 实时电力数据监控
  - **电力数据**: 电力历史数据查询
  - **极值报表**: 电力极值统计
  - **电力报表**: 电力数据报表
  - **配电图**: 电力一次接线图展示
- **场景联动**: 自动化和一键场景
  - **自动化**: 自动化场景管理
  - **一键场景**: 手动触发场景
- **预付费管理**: 房间、租客、充值和消费管理
  - **房间管理**: 房间 CRUD
  - **租客管理**: 租客 CRUD
  - **充值明细**: 充值记录查询
  - **消费明细**: 消费记录查询
  - **系统设置**: 预付费系统参数配置
- **系统管理**: 场景日志、权限区域、角色管理
  - **场景日志**: 场景执行日志
  - **权限区域**: 数据权限区域配置
  - **角色管理**: 租户角色管理

## API 端点

共 388 个可调用端点（按 admin/all/public 权限过滤）。

### things/device/version

- `POST /api/v1/things/device/version/get-list` — 获取设备模块版本列表
- `POST /api/v1/things/device/version/get-one` — 获取设备模块版本详情

### things/protocol/script

- `POST /api/v1/things/protocol/script/batch-export` — 批量导出协议脚本
- `POST /api/v1/things/protocol/script/batch-import` — 批量导入协议脚本
- `POST /api/v1/things/protocol/script/create` — 新增协议脚本
- `POST /api/v1/things/protocol/script/debug` — 协议脚本调试
- `POST /api/v1/things/protocol/script/delete` — 删除协议脚本
- `POST /api/v1/things/protocol/script/get-list` — 获取协议脚本列表
- `POST /api/v1/things/protocol/script/get-one` — 获取协议脚本详情
- `POST /api/v1/things/protocol/script/update` — 更新协议脚本

### system/dict/info

- `POST /api/v1/system/dict/info/get-list` — 获取字典信息列表

### system/ops/feedback

- `POST /api/v1/system/ops/feedback/create` — 添加帮助与反馈
- `POST /api/v1/system/ops/feedback/get-list` — 获取帮助与反馈
- `POST /api/v1/system/ops/feedback/update` — 更新帮助与反馈

### things/ai/mcp

- `POST /api/v1/things/ai/mcp/message` — MCP消息发送
- `POST /api/v1/things/ai/mcp/run` — Stateless MCP HTTP
- `GET /api/v1/things/ai/mcp/sse` — SSE连接
- `POST /api/v1/things/ai/mcp/sse` — SSE连接（POST）

### things/data/project

- `POST /api/v1/things/data/project/batch-create` — 批量创建授权项目权限
- `POST /api/v1/things/data/project/batch-delete` — 批量删除授权项目权限
- `POST /api/v1/things/data/project/create` — 创建授权项目权限
- `POST /api/v1/things/data/project/delete` — 删除授权项目权限
- `POST /api/v1/things/data/project/get-list` — 获取项目权限列表

### things/ota/firmware/job

- `POST /api/v1/things/ota/firmware/job/create` — 创建升级任务
- `POST /api/v1/things/ota/firmware/job/get-list` — 获取升级包下的升级任务批次列表
- `POST /api/v1/things/ota/firmware/job/get-one` — 查询指定升级批次的详情
- `POST /api/v1/things/ota/firmware/job/update` — 更新升级批次

### things/schema/common

- `POST /api/v1/things/schema/common/batch-export` — 批量导出通用物模型
- `POST /api/v1/things/schema/common/batch-import` — 批量导入通用物模型
- `POST /api/v1/things/schema/common/create` — 新增通用物模型
- `POST /api/v1/things/schema/common/delete` — 删除通用物模型
- `POST /api/v1/things/schema/common/get-list` — 获取通用物模型列表
- `POST /api/v1/things/schema/common/init` — 初始化通用物模型
- `POST /api/v1/things/schema/common/mock-gen` — 从参数生成物模型定义
- `POST /api/v1/things/schema/common/update` — 更新通用物模型

### system/checkIn

- `POST /api/v1/system/check-in/do` — 用户签到
- `POST /api/v1/system/check-in/get-list` — 签到记录列表
- `POST /api/v1/system/check-in/point-balance/get` — 获取当前用户积分余额
- `POST /api/v1/system/check-in/point-log/adjust` — 管理员调整积分
- `POST /api/v1/system/check-in/point-log/get-list` — 积分流水列表

### system/role/resource

- `POST /api/v1/system/role/resource/batch-update` — 批量更新角色资源动作权限
- `POST /api/v1/system/role/resource/get-list` — 获取角色资源动作权限列表

### things/ota/module/info

- `POST /api/v1/things/ota/module/info/create` — 创建模块
- `POST /api/v1/things/ota/module/info/delete` — 删除模块
- `POST /api/v1/things/ota/module/info/get-list` — 获取模块列表
- `POST /api/v1/things/ota/module/info/get-one` — 查询模块详情
- `POST /api/v1/things/ota/module/info/update` — 更新模块

### things/rule/alarm/record

- `POST /api/v1/things/rule/alarm/record/deal` — 处理告警(弃用)
- `POST /api/v1/things/rule/alarm/record/get-list` — 获取告警记录列表(弃用)

### system/ops/workOrder

- `POST /api/v1/system/ops/work-order/create` — 添加工单
- `POST /api/v1/system/ops/work-order/get-list` — 获取工单列表
- `POST /api/v1/system/ops/work-order/update` — 更新工单

### system/role/menu

- `POST /api/v1/system/role/menu/batch-update` — 更新角色对应菜单列表
- `POST /api/v1/system/role/menu/get-list` — 获取角色对应菜单列表

### things/device/schema

- `POST /api/v1/things/device/schema/batch-create` — 批量创建设备物模型
- `POST /api/v1/things/device/schema/batch-delete` — 批量删除设备物模型
- `POST /api/v1/things/device/schema/create` — 创建设备物模型
- `POST /api/v1/things/device/schema/get-list` — 获取设备物模型列表
- `POST /api/v1/things/device/schema/tsl-read` — 获取设备物模型tsl
- `POST /api/v1/things/device/schema/update` — 更新设备物模型

### things/product/custom

- `POST /api/v1/things/product/custom/get-one` — 获取产品自定义信息详情
- `POST /api/v1/things/product/custom/update` — 更新产品自定义信息

### things/rule/scene/log

- `POST /api/v1/things/rule/scene/log/get-list` — 获取场景日志列表(弃用)

### things/user/device/collect

- `POST /api/v1/things/user/device/collect/batch-create` — 批量收藏设备
- `POST /api/v1/things/user/device/collect/batch-delete` — 批量取消收藏设备
- `POST /api/v1/things/user/device/collect/get-list` — 获取收藏设备列表

### system/dept/syncJob

- `POST /api/v1/system/dept/sync-job/create` — 添加同步任务
- `POST /api/v1/system/dept/sync-job/delete` — 删除同步任务
- `POST /api/v1/system/dept/sync-job/execute` — 执行同步任务
- `POST /api/v1/system/dept/sync-job/get-list` — 获取同步任务列表
- `POST /api/v1/system/dept/sync-job/get-one` — 获取同步任务详情
- `POST /api/v1/system/dept/sync-job/update` — 更新同步任务

### system/mall/product

- `POST /api/v1/system/mall/product/get-list` — 获取商品列表
- `POST /api/v1/system/mall/product/get-one` — 获取商品详情

### system/tenant/app

- `POST /api/v1/system/tenant/app/create` — 绑定租户应用
- `POST /api/v1/system/tenant/app/get-list` — 获取租户应用列表

### things/alarm/info

- `POST /api/v1/things/alarm/info/create` — 新增告警
- `POST /api/v1/things/alarm/info/delete` — 删除告警
- `POST /api/v1/things/alarm/info/get-list` — 获取告警信息列表
- `POST /api/v1/things/alarm/info/get-one` — 获取告警信息
- `POST /api/v1/things/alarm/info/update` — 更新告警

### things/area/info

- `POST /api/v1/things/area/info/create` — 新增项目区域
- `POST /api/v1/things/area/info/delete` — 删除项目区域
- `POST /api/v1/things/area/info/get-list` — 获取项目区域列表
- `POST /api/v1/things/area/info/get-one` — 获取项目区域详情
- `POST /api/v1/things/area/info/update` — 更新项目区域

### system/tenant/renewal

- `POST /api/v1/system/tenant/renewal/get-list` — 获取续期列表
- `POST /api/v1/system/tenant/renewal/renew` — 续费

### system/user/dept

- `POST /api/v1/system/user/dept/batch-create` — 新增用户的部门列表
- `POST /api/v1/system/user/dept/batch-delete` — 删除用户的部门列表

### things/scene/log

- `POST /api/v1/things/scene/log/get-list` — 获取场景日志列表

### system/dept/user

- `POST /api/v1/system/dept/user/batch-create` — 批量授权部门用户
- `POST /api/v1/system/dept/user/batch-delete` — 批量取消授权部门用户
- `POST /api/v1/system/dept/user/get-list` — 获取部门授权列表

### system/tenant/info

- `POST /api/v1/system/tenant/info/create` — 添加租户
- `POST /api/v1/system/tenant/info/get-one` — 获取租户详情
- `POST /api/v1/system/tenant/info/transfer` — 转让租户（仅租户所有者）
- `POST /api/v1/system/tenant/info/update` — 更新租户

### things/group/device

- `POST /api/v1/things/group/device/batch-create` — 添加分组设备
- `POST /api/v1/things/group/device/batch-delete` — 删除分组设备
- `POST /api/v1/things/group/device/batch-update` — 更新分组设备

### things/group/info

- `POST /api/v1/things/group/info/create` — 创建分组
- `POST /api/v1/things/group/info/delete` — 删除分组
- `POST /api/v1/things/group/info/get-list` — 获取分组列表
- `POST /api/v1/things/group/info/get-one` — 获取分组详情信息
- `POST /api/v1/things/group/info/update` — 更新分组信息

### things/project/customData

- `POST /api/v1/things/project/crud/create` — 新增项目crud
- `POST /api/v1/things/project/crud/update` — 更新项目crud

### things/project/info

- `POST /api/v1/things/project/info/create` — 新增项目
- `POST /api/v1/things/project/info/delete` — 删除项目
- `POST /api/v1/things/project/info/get-list` — 获取项目列表
- `POST /api/v1/things/project/info/get-one` — 获取项目详情
- `POST /api/v1/things/project/info/update` — 更新项目

### things/protocol/config

- `POST /api/v1/things/protocol/config/create` — 创建协议配置
- `POST /api/v1/things/protocol/config/delete` — 删除协议配置
- `POST /api/v1/things/protocol/config/get-list` — 获取协议配置列表
- `POST /api/v1/things/protocol/config/get-one` — 获取协议配置详情
- `POST /api/v1/things/protocol/config/update` — 更新协议配置

### things/protocol/info

- `POST /api/v1/things/protocol/info/create` — 新增自定义协议
- `POST /api/v1/things/protocol/info/delete` — 删除自定义协议
- `POST /api/v1/things/protocol/info/get-list` — 获取自定义协议信息列表
- `POST /api/v1/things/protocol/info/get-one` — 获取自定义协议详情
- `POST /api/v1/things/protocol/info/update` — 更新自定义协议

### things/alarm/scene

- `POST /api/v1/things/alarm/scene/batch-create` — 更新告警和场景的关联
- `POST /api/v1/things/alarm/scene/delete` — 删除告警和场景的关联
- `POST /api/v1/things/alarm/scene/get-list` — 获取告警和场景的关联列表

### things/product/schema

- `POST /api/v1/things/product/schema/batch-create` — 批量创建产品物模型
- `POST /api/v1/things/product/schema/create` — 创建产品物模型
- `POST /api/v1/things/product/schema/delete` — 删除产品物模型
- `POST /api/v1/things/product/schema/get-list` — 获取产品物模型
- `POST /api/v1/things/product/schema/tsl-import` — 导入产品物模型tsl
- `POST /api/v1/things/product/schema/tsl-read` — 获取产品物模型tsl
- `POST /api/v1/things/product/schema/update` — 更新产品物模型

### things/protocol/image

- `POST /api/v1/things/protocol/image/pull` — 下载协议镜像

### things/protocol/script/device

- `POST /api/v1/things/protocol/script/device/create` — 新增协议脚本设备
- `POST /api/v1/things/protocol/script/device/delete` — 删除协议脚本设备
- `POST /api/v1/things/protocol/script/device/get-list` — 获取协议脚本设备列表
- `POST /api/v1/things/protocol/script/device/get-one` — 获取协议脚本设备详情
- `POST /api/v1/things/protocol/script/device/update` — 更新协议脚本设备

### system/tenant/agreement

- `POST /api/v1/system/tenant/agreement/create` — 添加协议
- `POST /api/v1/system/tenant/agreement/delete` — 删除协议
- `POST /api/v1/system/tenant/agreement/get-list` — 获取协议列表
- `POST /api/v1/system/tenant/agreement/get-one` — 获取协议详情
- `POST /api/v1/system/tenant/agreement/update` — 更新协议

### system/tenant/user/role

- `POST /api/v1/system/tenant/user/role/batch-update` — 更新租户用户的角色列表
- `POST /api/v1/system/tenant/user/role/get-list` — 获取租户用户角色列表

### things/data/area/user/apply

- `POST /api/v1/things/data/area/user/apply/deal` — 授权区域权限
- `POST /api/v1/things/data/area/user/apply/get-list` — 获取区域权限列表

### things/device/gateway

- `POST /api/v1/things/device/gateway/batch-create` — 添加网关子设备
- `POST /api/v1/things/device/gateway/batch-delete` — 解绑子设备
- `POST /api/v1/things/device/gateway/get-list` — 获取子设备列表

### things/ota/firmware/info

- `POST /api/v1/things/ota/firmware/info/create` — 添加升级包
- `POST /api/v1/things/ota/firmware/info/delete` — 删除升级包
- `POST /api/v1/things/ota/firmware/info/get-list` — 升级包列表
- `POST /api/v1/things/ota/firmware/info/get-one` — 查询升级包
- `POST /api/v1/things/ota/firmware/info/update` — 修改升级包

### things/product/config

- `POST /api/v1/things/product/config/update` — 更新配置

### things/project/profile

- `POST /api/v1/things/project/profile/get-list` — 获取项目配置列表
- `POST /api/v1/things/project/profile/get-one` — 获取项目配置详情
- `POST /api/v1/things/project/profile/update` — 更新项目配置

### things/rule/alarm/info

- `POST /api/v1/things/rule/alarm/info/create` — 新增告警(弃用)
- `POST /api/v1/things/rule/alarm/info/delete` — 删除告警(弃用)
- `POST /api/v1/things/rule/alarm/info/get-list` — 获取告警信息列表(弃用)
- `POST /api/v1/things/rule/alarm/info/get-one` — 获取告警信息(弃用)
- `POST /api/v1/things/rule/alarm/info/update` — 更新告警(弃用)

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

### system/config/core

- `POST /api/v1/system/common/sys-config/core/get-one` — 读取系统配置信息(无需登录)

### things/device/ota

- `POST /api/v1/things/device/info/batch-update` — 批量更新设备

### things/device/interact

- `POST /api/v1/things/device/interact/action-get-one` — 获取调用设备行为的结果
- `POST /api/v1/things/device/interact/action-resp` — 回复设备行为调用结果
- `POST /api/v1/things/device/interact/action-send` — 调用设备行为
- `POST /api/v1/things/device/interact/event-send` — 下行事件通知设备
- `POST /api/v1/things/device/interact/property-control-batch-send` — 批量调用设备属性
- `POST /api/v1/things/device/interact/property-control-get-one` — 获取调用设备属性的结果
- `POST /api/v1/things/device/interact/property-control-send` — 调用设备属性
- `POST /api/v1/things/device/interact/property-get-report-batch-send` — 批量请求设备获取设备最新属性
- `POST /api/v1/things/device/interact/property-get-report-send` — 请求设备获取设备最新属性
- `POST /api/v1/things/device/interact/schema-mock-gen` — 生成物模型模拟数据

### things/ota/firmware/device

- `POST /api/v1/things/ota/firmware/device/cancel` — 取消指定任务下的升级
- `POST /api/v1/things/ota/firmware/device/confirm` — 确认升级设备
- `POST /api/v1/things/ota/firmware/device/get-list` — 查询升级设备列表
- `POST /api/v1/things/ota/firmware/device/retry` — 重试设备升级

### things/product/info

- `POST /api/v1/things/product/info/batch-export` — 批量导出产品
- `POST /api/v1/things/product/info/batch-import` — 批量导入产品
- `POST /api/v1/things/product/info/create` — 新增产品
- `POST /api/v1/things/product/info/delete` — 删除产品
- `POST /api/v1/things/product/info/get-list` — 获取产品信息列表
- `POST /api/v1/things/product/info/get-one` — 获取产品详情
- `POST /api/v1/things/product/info/init` — 初始化产品
- `POST /api/v1/things/product/info/update` — 更新产品

### things/project/crud

- `POST /api/v1/things/project/crud/delete` — 删除项目crud
- `POST /api/v1/things/project/crud/get-list` — 获取项目crud列表
- `POST /api/v1/things/project/crud/get-one` — 获取项目crud详情

### things/protocol/sync

- `POST /api/v1/things/protocol/sync/device` — 设备同步(如果该协议不支持会返回不支持)
- `POST /api/v1/things/protocol/sync/product` — 产品同步(如果该协议不支持会返回不支持)

### system/app/core

- `POST /api/v1/system/app/core/get-one` — 无需登录获取应用信息

### system/mall/package

- `POST /api/v1/system/mall/package/get-list` — 获取套餐列表
- `POST /api/v1/system/mall/package/get-one` — 获取套餐详情

### system/tenant/config

- `POST /api/v1/system/tenant/config/get-one` — 获取租户配置
- `POST /api/v1/system/tenant/config/update` — 更新租户配置

### system/user/self/accessToken

- `POST /api/v1/system/user/self/access-token/create` — 创建访问令牌
- `POST /api/v1/system/user/self/access-token/delete` — 删除访问令牌
- `POST /api/v1/system/user/self/access-token/get-list` — 获取访问令牌列表
- `POST /api/v1/system/user/self/access-token/get-one` — 获取访问令牌详情
- `POST /api/v1/system/user/self/access-token/update` — 更新访问令牌

### things/device/profile

- `POST /api/v1/things/device/profile/delete` — 删除设备配置
- `POST /api/v1/things/device/profile/get-list` — 获取设备配置列表
- `POST /api/v1/things/device/profile/get-one` — 获取设备配置详情
- `POST /api/v1/things/device/profile/update` — 更新设备配置

### things/product/category

- `POST /api/v1/things/product/category/batch-export` — 批量导出产品品类
- `POST /api/v1/things/product/category/batch-import` — 批量导入产品品类
- `POST /api/v1/things/product/category/create` — 新增产品品类
- `POST /api/v1/things/product/category/delete` — 删除产品品类
- `POST /api/v1/things/product/category/get-list` — 获取产品品类列表
- `POST /api/v1/things/product/category/get-one` — 获取产品品类详情
- `POST /api/v1/things/product/category/schema/batch-create` — 批量新增产品品类物模型
- `POST /api/v1/things/product/category/schema/batch-delete` — 批量删除产品品类物模型
- `POST /api/v1/things/product/category/schema/batch-update` — 批量更新产品品类物模型
- `POST /api/v1/things/product/category/schema/get-list` — 获取产品品类物模型列表
- `POST /api/v1/things/product/category/update` — 更新产品品类

### things/protocol/container

- `POST /api/v1/things/protocol/container/install` — 安装协议容器
- `POST /api/v1/things/protocol/container/status` — 查询协议容器状态
- `POST /api/v1/things/protocol/container/stop` — 停止协议容器
- `POST /api/v1/things/protocol/container/uninstall` — 卸载协议容器
- `POST /api/v1/things/protocol/container/update` — 更新协议容器

### things/scene/info

- `POST /api/v1/things/scene/info/create` — 新增场景
- `POST /api/v1/things/scene/info/delete` — 删除场景
- `POST /api/v1/things/scene/info/get-list` — 获取场景信息列表
- `POST /api/v1/things/scene/info/get-one` — 获取场景信息详情
- `POST /api/v1/things/scene/info/manually-trigger` — 手动触发场景
- `POST /api/v1/things/scene/info/update` — 更新场景

### system/log

- `POST /api/v1/system/log/login/get-list` — 获取登录日志列表
- `POST /api/v1/system/log/oper/get-list` — 获取操作日志列表

### system/tenant/core

- `POST /api/v1/system/tenant/core/get-list` — 搜索租户信息
- `POST /api/v1/system/tenant/core/get-one` — 获取租户信息

### things/data/area

- `POST /api/v1/things/data/area/batch-delete` — 删除授权区域权限
- `POST /api/v1/things/data/area/batch-update` — 更新授权区域权限
- `POST /api/v1/things/data/area/get-list` — 获取区域权限列表

### things/device/group

- `POST /api/v1/things/device/group/batch-create` — 将设备加到多个分组中
- `POST /api/v1/things/device/group/batch-delete` — 删除设备所在分组
- `POST /api/v1/things/device/group/batch-update` — 更新设备所在分组

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

### system/dept/info

- `POST /api/v1/system/dept/info/create` — 添加部门详情
- `POST /api/v1/system/dept/info/delete` — 删除部门
- `POST /api/v1/system/dept/info/get-list` — 获取部门列表
- `POST /api/v1/system/dept/info/get-one` — 获取部门详情
- `POST /api/v1/system/dept/info/update` — 更新部门

### system/tenant/user

- `POST /api/v1/system/tenant/user/batch-create` — 批量添加用户加入租户
- `POST /api/v1/system/tenant/user/delete` — 删除租户用户
- `POST /api/v1/system/tenant/user/get-list` — 获取租户用户列表
- `POST /api/v1/system/tenant/user/get-one` — 获取租户用户详情,会同时返回所拥有的角色列表
- `POST /api/v1/system/tenant/user/invite` — 邀请用户加入租户
- `POST /api/v1/system/tenant/user/invite-code/gen` — 生成租户用户邀请码
- `POST /api/v1/system/tenant/user/invite-code/get-one` — 获取当前有效的租户用户邀请码
- `POST /api/v1/system/tenant/user/invite-pending/delete` — 删除待处理邀请
- `POST /api/v1/system/tenant/user/invite-pending/get-list` — 获取待处理邀请列表
- `POST /api/v1/system/tenant/user/invite-send` — 发送邀请（支持邮件和手机号）
- `POST /api/v1/system/tenant/user/update` — 更新租户用户

### things/device/bind-token

- `POST /api/v1/things/device/info/create` — 新增设备
- `POST /api/v1/things/device/info/update` — 更新设备

### things/hook

- `POST /api/v1/things/hook/` — Hook扩展统一入口

### things/protocol/service

- `POST /api/v1/things/protocol/service/delete` — 删除自定义协议服务器
- `POST /api/v1/things/protocol/service/get-list` — 获取自定义协议服务器信息列表

### things/rule/alarm/scene

- `POST /api/v1/things/rule/alarm/scene/batch-create` — 更新告警和场景的关联(弃用)
- `POST /api/v1/things/rule/alarm/scene/delete` — 删除告警和场景的关联(弃用)
- `POST /api/v1/things/rule/alarm/scene/get-list` — 获取告警和场景的关联列表(弃用)

### system/init

- `POST /api/v1/system/common/system/init` — 初始化系统

### system/mall/license

- `POST /api/v1/system/mall/license/get-list` — 授权码列表
- `POST /api/v1/system/mall/license/get-one` — 授权码详情

### system/role/info

- `POST /api/v1/system/role/info/create` — 添加角色
- `POST /api/v1/system/role/info/delete` — 删除角色
- `POST /api/v1/system/role/info/get-list` — 获取角色列表
- `POST /api/v1/system/role/info/update` — 更新角色

### system/user/info

- `POST /api/v1/system/user/info/create` — 创建用户信息
- `POST /api/v1/system/user/info/delete` — 刪除用户
- `POST /api/v1/system/user/info/get-list` — 查询用户信息列表
- `POST /api/v1/system/user/info/get-one` — 获取用户信息
- `POST /api/v1/system/user/info/update` — 更新用户基本数据

### system/user/tenant

- `POST /api/v1/system/user/tenant/get-list` — 用户所处的租户列表

### things/device/interact/gateway

- `POST /api/v1/things/device/interact/gateway-get-found-send` — 请求网关上报拓扑关系
- `POST /api/v1/things/device/interact/gateway-notify-bind-send` — 通知网关绑定子设备

### system/role/app

- `POST /api/v1/system/role/app/batch-update` — 更新APP权限
- `POST /api/v1/system/role/app/get-list` — 获取APP权限列表

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

### things/alarm/record

- `POST /api/v1/things/alarm/record/deal` — 处理告警
- `POST /api/v1/things/alarm/record/get-list` — 获取告警记录列表

### things/device/msg

- `POST /api/v1/things/device/msg/abnormal-log/get-list` — 获取设备异常日志
- `POST /api/v1/things/device/msg/event-log/get-list` — 获取事件历史记录
- `POST /api/v1/things/device/msg/gateway-can-bind/get-list` — 获取网关可以绑定的子设备列表
- `POST /api/v1/things/device/msg/hub-log/get-list` — 获取云端诊断日志
- `POST /api/v1/things/device/msg/property-agg/by-device/get-list` — 弃用
- `POST /api/v1/things/device/msg/property-agg/get-list` — 弃用
- `POST /api/v1/things/device/msg/property-latest-agg/get-list` — 聚合属性最新值
- `POST /api/v1/things/device/msg/property-latest/get-list` — 获取最新属性记录
- `POST /api/v1/things/device/msg/property-log-agg/by-device/get-list` — 聚合属性历史记录,设备维度
- `POST /api/v1/things/device/msg/property-log-agg/get-list` — 聚合属性历史记录
- `POST /api/v1/things/device/msg/property-log-latest/get-list` — 弃用
- `POST /api/v1/things/device/msg/property-log/batch-get-list` — 批量获取单个id属性历史记录
- `POST /api/v1/things/device/msg/property-log/get-list` — 获取单个id属性历史记录
- `POST /api/v1/things/device/msg/sdk-log/get-list` — 获取设备sdk日志
- `POST /api/v1/things/device/msg/send-log/get-list` — 获取设备命令日志
- `POST /api/v1/things/device/msg/shadow/get-list` — 获取设备影子列表
- `POST /api/v1/things/device/msg/status-log/get-list` — 获取设备状态日志

### things/rule/scene/info

- `POST /api/v1/things/rule/scene/info/create` — 新增场景(弃用)
- `POST /api/v1/things/rule/scene/info/delete` — 删除场景(弃用)
- `POST /api/v1/things/rule/scene/info/get-list` — 获取场景信息列表(弃用)
- `POST /api/v1/things/rule/scene/info/get-one` — 获取场景信息详情(弃用)
- `POST /api/v1/things/rule/scene/info/manually-trigger` — 手动触发场景(弃用)
- `POST /api/v1/things/rule/scene/info/update` — 更新场景(弃用)

### things/user/area/apply

- `POST /api/v1/things/user/area/apply/create` — 申请用户区域权限

### system/tenant/app/menu

- `POST /api/v1/system/tenant/app/menu/get-list` — 获取租户应用菜单列表
- `POST /api/v1/system/tenant/app/menu/update` — 更新租户应用菜单

### things/product/remoteConfig

- `POST /api/v1/things/product/remote-config/create` — 创建配置
- `POST /api/v1/things/product/remote-config/get-list` — 获取配置列表
- `POST /api/v1/things/product/remote-config/lastest-read` — 获取最新配置
- `POST /api/v1/things/product/remote-config/push-all` — 推送配置

### system/user/data

- `POST /api/v1/system/user/data/area/get-list` — 获取区域权限列表
- `POST /api/v1/system/user/data/project/get-list` — 获取项目权限列表

### system/user/self/tenant

- `POST /api/v1/system/user/self/tenant/delete` — 退出当前租户
- `POST /api/v1/system/user/self/tenant/get-list` — 获取用户所处的租户列表
- `POST /api/v1/system/user/self/tenant/get-one` — 获取当前用户在当前租户的详情
- `POST /api/v1/system/user/self/tenant/join` — 用户加入租户（通过邀请码、邮件或手机邀请）
- `POST /api/v1/system/user/self/tenant/update` — 更新当前用户在当前租户的信息

### things/area/profile

- `POST /api/v1/things/area/profile/get-list` — 获取区域配置列表
- `POST /api/v1/things/area/profile/get-one` — 获取区域配置详情
- `POST /api/v1/things/area/profile/update` — 更新区域配置

### things/device/info

- `POST /api/v1/things/device/info/batch-bind` — 批量绑定
- `POST /api/v1/things/device/info/batch-import` — 批量导入设备
- `POST /api/v1/things/device/info/batch-update-import` — 导入批量更新设备
- `POST /api/v1/things/device/info/bind` — 绑定
- `POST /api/v1/things/device/info/bind/token/create` — 创建绑定token
- `POST /api/v1/things/device/info/bind/token/get-one` — 绑定token状态查询
- `POST /api/v1/things/device/info/can-bind` — 是否可以绑定设备
- `POST /api/v1/things/device/info/count` — 设备统计详情
- `POST /api/v1/things/device/info/delete` — 删除设备
- `POST /api/v1/things/device/info/get-list` — 获取设备列表
- `POST /api/v1/things/device/info/get-one` — 获取设备详情
- `POST /api/v1/things/device/info/move` — 转移设备到新设备上
- `POST /api/v1/things/device/info/ota/upgrade` — 设备升级,获取升级包手动升级
- `POST /api/v1/things/device/info/transfer` — 转让设备
- `POST /api/v1/things/device/info/unbind` — 解绑设备

## 使用示例

```bash
# 配置
ur-org-energy setup

# 验证连通性
ur-org-energy check

# 调用 API
ur-org-energy api /api/v1/system/user/self/get-one
```
