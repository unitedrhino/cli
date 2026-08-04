# 前端应用权限映射

| 应用 | 类型 | TENANT_CODE | 可调用 API authType |
|------|------|-------------|---------------------|
| platform-manage | 平台 | platform | platform, admin, all |
| platform-iot | 平台 | platform | platform, admin, all |
| org-manage | 组织 | 用户企业 | admin, all |
| org-iot | 组织 | 用户企业 | admin, all |
| org-energy | 组织 | 用户企业 | admin, all |
| console | 公共 | 用户企业 | all |

## 平台应用（platform-manage, platform-iot）

- **环境变量**: `VITE_GLOB_TENANT_CODE=platform`
- **请求头**: 固定发送 `tenant-code: platform`
- **可访问 API**: 所有 authType
- **典型场景**: 企业管理、应用配置、平台级数据查看

## 组织应用（org-manage, org-iot, org-energy）

- **环境变量**: 无 `VITE_GLOB_TENANT_CODE`
- **请求头**: 企业代码由登录用户决定
- **可访问 API**: admin、all
- **典型场景**: 设备管理、用户管理、企业级配置

## 公共应用（console）

- **环境变量**: `VITE_GLOB_APPID=77`
- **请求头**: 企业代码由登录用户决定
- **可访问 API**: all
- **典型场景**: 应用入口、企业切换、个人信息管理
