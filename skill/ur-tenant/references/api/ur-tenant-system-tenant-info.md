# ur-tenant system/tenant/info

添加租户 等

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/info/create` | 添加租户 | all |
| POST | `/api/v1/system/tenant/info/delete` | 删除租户 | platform |
| POST | `/api/v1/system/tenant/info/get-list` | 获取租户列表 | platform |
| POST | `/api/v1/system/tenant/info/get-one` | 获取租户详情 | admin |
| POST | `/api/v1/system/tenant/info/transfer` | 转让租户（仅租户所有者） | admin |
| POST | `/api/v1/system/tenant/info/update` | 更新租户 | admin |
