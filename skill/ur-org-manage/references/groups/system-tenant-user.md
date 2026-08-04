# system/tenant/user

> 该 group 共 11 个端点。

- `POST /api/v1/system/tenant/user/batch-create` [admin] 批量添加用户加入企业
- `POST /api/v1/system/tenant/user/delete` [admin] 删除企业用户
- `POST /api/v1/system/tenant/user/get-list` [admin] 获取企业用户列表
- `POST /api/v1/system/tenant/user/get-one` [admin] 获取企业用户详情,会同时返回所拥有的角色列表
- `POST /api/v1/system/tenant/user/invite` [admin] 邀请用户加入企业
- `POST /api/v1/system/tenant/user/invite-code/gen` [admin] 生成企业用户邀请码
- `POST /api/v1/system/tenant/user/invite-code/get-one` [admin] 获取当前有效的企业用户邀请码
- `POST /api/v1/system/tenant/user/invite-pending/delete` [admin] 删除待处理邀请
- `POST /api/v1/system/tenant/user/invite-pending/get-list` [admin] 获取待处理邀请列表
- `POST /api/v1/system/tenant/user/invite-send` [admin] 发送邀请（支持邮件和手机号）
- `POST /api/v1/system/tenant/user/update` [admin] 更新企业用户
