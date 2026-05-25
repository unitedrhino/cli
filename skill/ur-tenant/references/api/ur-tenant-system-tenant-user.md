# ur-tenant system/tenant/user

批量添加用户加入租户 等

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/user/batch-create` | 批量添加用户加入租户 | admin |
| POST | `/api/v1/system/tenant/user/delete` | 删除租户用户 | admin |
| POST | `/api/v1/system/tenant/user/get-list` | 获取租户用户列表 | admin |
| POST | `/api/v1/system/tenant/user/get-one` | 获取租户用户详情,会同时返回所拥有的角色列表 | admin |
| POST | `/api/v1/system/tenant/user/invite` | 邀请用户加入租户 | admin |
| POST | `/api/v1/system/tenant/user/invite-code/gen` | 生成租户用户邀请码 | admin |
| POST | `/api/v1/system/tenant/user/invite-code/get-one` | 获取当前有效的租户用户邀请码 | admin |
| POST | `/api/v1/system/tenant/user/invite-pending/delete` | 删除待处理邀请 | admin |
| POST | `/api/v1/system/tenant/user/invite-pending/get-list` | 获取待处理邀请列表 | admin |
| POST | `/api/v1/system/tenant/user/invite-send` | 发送邀请（支持邮件和手机号） | admin |
| POST | `/api/v1/system/tenant/user/update` | 更新租户用户 | admin |
