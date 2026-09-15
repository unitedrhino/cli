# system/tenant/user

> 该 group 共 11 个端点。

- `POST /api/v1/system/tenant/user/batch-create` [admin] 批量添加用户加入租户
- `POST /api/v1/system/tenant/user/delete` [admin] 删除租户用户
- `POST /api/v1/system/tenant/user/get-list` [admin] 获取租户用户列表
- `POST /api/v1/system/tenant/user/get-one` [admin] 获取租户用户详情,会同时返回所拥有的角色列表
- `POST /api/v1/system/tenant/user/invite` [admin] 邀请用户加入租户
- `POST /api/v1/system/tenant/user/invite-code/gen` [admin] 生成租户用户邀请码
- `POST /api/v1/system/tenant/user/invite-code/get-one` [admin] 获取当前有效的租户用户邀请码
- `POST /api/v1/system/tenant/user/invite-pending/delete` [admin] 删除待处理邀请
- `POST /api/v1/system/tenant/user/invite-pending/get-list` [admin] 获取待处理邀请列表
- `POST /api/v1/system/tenant/user/invite-send` [admin] 发送邀请（支持邮件和手机号）
- `POST /api/v1/system/tenant/user/update` [admin] 更新租户用户
