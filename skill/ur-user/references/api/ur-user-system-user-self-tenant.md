# ur-user system/user/self/tenant

退出当前租户 等

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/self/tenant/delete` | 退出当前租户 | all |
| POST | `/api/v1/system/user/self/tenant/get-list` | 获取用户所处的租户列表 | all |
| POST | `/api/v1/system/user/self/tenant/get-one` | 获取当前用户在当前租户的详情 | all |
| POST | `/api/v1/system/user/self/tenant/join` | 用户加入租户（通过邀请码、邮件或手机邀请） | all |
| POST | `/api/v1/system/user/self/tenant/update` | 更新当前用户在当前租户的信息 | all |
