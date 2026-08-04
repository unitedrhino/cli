---
name: ur-api-quick-reference
description: "ur-api 快速参考：路径前缀→域映射、关键枚举值速查、CLI schema 命令。当需要快速确认端点归属或参数格式时加载。triggers: 路径前缀, 枚举值, API路径查询, schema, 端点归属"
metadata:
  hermes:
    tags: [reference, api, quick-reference]
---


# ur-api 快速参考

> 补充文档，与各子域 SKILL.md 配合使用。

## 路径前缀 → 域映射

| 路径前缀 | 所属域 | 典型操作 |
|---------|--------|---------|
| `/api/v1/things/device/info/` | ur-device | 设备 CRUD |
| `/api/v1/things/device/interact/` | ur-device | 属性控制、行为调用 |
| `/api/v1/things/device/msg/` | ur-device | 消息日志、属性快照 |
| `/api/v1/things/device/auth/` | ur-device | 设备认证（三元组） |
| `/api/v1/things/device/gateway/` | ur-device | 网关子设备拓扑 |
| `/api/v1/things/device/schema/` | ur-device | 设备物模型 |
| `/api/v1/things/user/device/` | ur-device | 用户-设备（分享/收藏） |
| `/api/v1/things/product/` | ur-product | 产品 CRUD |
| `/api/v1/things/schema/` | ur-product | 通用/产品物模型 |
| `/api/v1/things/protocol/` | ur-product | 协议脚本 |
| `/api/v1/things/ota/` | ur-product | OTA 固件管理 |
| `/api/v1/things/project/` | ur-project | 项目 CRUD |
| `/api/v1/things/area/` | ur-project | 区域管理 |
| `/api/v1/things/group/` | ur-project | 设备分组 |
| `/api/v1/things/data/` | ur-project | 数据权限管理 |
| `/api/v1/things/user/area/` | ur-project | 用户区域权限申请 |
| `/api/v1/things/ai/` | ur-ai | AI Agent/Clone/Session |
| `/api/v1/things/alarm/` | ur-ai | 告警规则和记录 |
| `/api/v1/things/scene/` | ur-ai | 场景联动 |
| `/api/v1/things/rule/` | ur-ai | 规则引擎 |
| `/api/v1/system/user/` | ur-user | 用户管理/登录/个人信息 |
| `/api/v1/system/role/` | ur-user | 角色权限 |
| `/api/v1/system/dept/` | ur-user | 部门管理 |
| `/api/v1/system/dict/` | ur-user | 字典管理 |
| `/api/v1/system/notify/` | ur-user | 通知配置/消息 |
| `/api/v1/system/log/` | ur-user | 操作日志 |
| `/api/v1/system/common/` | ur-system | 文件上传/WebSocket/批量接口 |
| `/api/v1/system/app/` | ur-system | 应用管理 |
| `/api/v1/system/hook/` | ur-system | Hook 扩展 |
| `/api/v1/system/tenant/info/` | ur-tenant | 企业 CRUD |
| `/api/v1/system/tenant/user/` | ur-tenant | 企业用户管理 |
| `/api/v1/system/tenant/app/` | ur-tenant + ur-system | 企业应用绑定 |
| `/api/v1/system/tenant/config/` | ur-tenant | 企业配置 |

## 关键枚举值速查

### 权限相关

| 字段 | 含义 | 可选值 |
|------|------|--------|
| x-auth-type（swagger 标注）| 接口所需权限等级 | `all`=所有登录用户, `admin`=管理员, `platform`=平台管理员 |
| authType（项目数据权限）| 用户对项目/区域的权限等级 | `1`=管理权限（可授权）, `2`=读写, `3`=只读, `4`=区域管理员 |

### 设备相关

| 字段 | 含义 | 可选值 |
|------|------|--------|
| deviceType | 设备类型 | `1`=直连设备, `2`=网关设备, `3`=子设备 |
| isOnline | 在线状态 | `1`=在线, `2`=离线 |
| protocolCode | 通信协议 | `urMqtt`（默认，原旧值 `iThings` 已废弃） |

### 物模型相关

| 字段 | 含义 | 可选值 |
|------|------|--------|
| schemaType | 物模型类型 | `property`=属性, `event`=事件, `action`=行为 |
| accessMode | 属性访问模式 | `r`=只读, `rw`=读写 |
| eventType | 事件类型 | `info`=信息, `alert`=告警, `fault`=故障 |
| dataType.type | 数据类型 | `bool`, `int`, `float`, `string`, `enum`, `timestamp`, `struct`, `array` |

### 通知相关

| 字段 | 含义 | 可选值 |
|------|------|--------|
| notifyType | 通知渠道 | `sms`=短信, `email`=邮件, `dingtalk`=钉钉, `wechat`=微信, `inner`=站内信 |
| loginType | 登录方式 | `pwd`=密码, `sms`=短信验证码 |

### OTA 进度码

| code | 含义 |
|------|------|
| 1-100 | 升级进度百分比 |
| -1 | 下载失败 |
| -2 | 校验失败 |
| -3 | 烧录失败 |
| -4 | 版本不匹配 |

## CLI Schema 自省命令

当不确定端点参数或路径时使用：

```bash
cd .claude/skills/ur-api/scripts

# 列出所有端点（含权限标注）
ur schema

# 按路径前缀筛选
ur schema /api/v1/things/device/

# 查看特定端点完整参数
ur schema /api/v1/things/device/info/create

# 只看管理员接口
ur schema --auth-type admin

# JSON 格式输出（便于程序化处理）
ur schema --json
```

## 分页通用格式

所有列表接口统一分页格式，**从 1 开始**（不是 0）：

```json
{
  "page": { "page": 1, "size": 10 },
  "total": 0,
  "list": []
}
```

## 常见字段陷阱

| 陷阱 | 正确做法 |
|------|---------|
| JWT 中 `userID` 类型 | 必须是**字符串**格式：`"userID": "12345"`，不是数字 |
| `deviceName` 误以为是名称 | `deviceName` 是设备唯一 ID，显示名称是 `deviceAlias` |
| 属性标识符大小写 | 必须与物模型完全一致，通常是**大驼峰**（`CurrentTemperature`） |
| 分页起始值 | `page.page=1`，不是 `0` |
| 控制离线设备 | 命令会写入影子设备（期望值），设备上线后自动同步 |
