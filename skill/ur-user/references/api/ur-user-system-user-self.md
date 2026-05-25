# ur-user system/user/self

获取用户应用列表 等

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/user/self/app/get-list` | 获取用户应用列表 | all |
| POST | `/api/v1/system/user/self/app/get-one` | 获取用户应用详情 | all |
| POST | `/api/v1/system/user/self/bind-account` | 绑定账号 | all |
| POST | `/api/v1/system/user/self/cancel` | 注销用户 | all |
| POST | `/api/v1/system/user/self/captcha` | 获取验证码 | public |
| POST | `/api/v1/system/user/self/change-pwd` | 更新用户密码 | all |
| POST | `/api/v1/system/user/self/forget-pwd` | 忘记密码 | public |
| POST | `/api/v1/system/user/self/get-one` | 获取用户信息 | all |
| POST | `/api/v1/system/user/self/login` | 用户登录 | public |
| POST | `/api/v1/system/user/self/logout` | 用户登出 | all |
| POST | `/api/v1/system/user/self/menu/get-list` | 获取用户菜单列表 | all |
| POST | `/api/v1/system/user/self/message/get-list` | 用户消息列表 | all |
| POST | `/api/v1/system/user/self/message/get-pending` | 用户待处理消息 | all |
| POST | `/api/v1/system/user/self/message/handle` | 用户消息标记已处理 | all |
| POST | `/api/v1/system/user/self/message/mark-all-read` | 用户消息全部已读 | all |
| POST | `/api/v1/system/user/self/message/multi-delete` | 用户消息批量删除 | all |
| POST | `/api/v1/system/user/self/message/multi-is-read` | 用户消息批量已读 | all |
| POST | `/api/v1/system/user/self/message/statistics` | 用户消息统计 | all |
| POST | `/api/v1/system/user/self/notify-preference/read` | 用户通知偏好读取 | all |
| POST | `/api/v1/system/user/self/notify-preference/update` | 用户通知偏好更新 | all |
| POST | `/api/v1/system/user/self/profile/get-list` | 获取用户配置列表 | all |
| POST | `/api/v1/system/user/self/profile/get-one` | 获取用户配置详情 | all |
| POST | `/api/v1/system/user/self/profile/update` | 更新用户配置 | all |
| POST | `/api/v1/system/user/self/register` | 普通用户注册 | public |
| POST | `/api/v1/system/user/self/resource/action/get-list` | 获取用户资源动作权限列表 | all |
| POST | `/api/v1/system/user/self/third-auth/start` | 第三方登录授权起跳 | public |
| POST | `/api/v1/system/user/self/third-login` | 第三方登录回调换平台登录态 | public |
| POST | `/api/v1/system/user/self/third-register` | 第三方补全注册 | public |
| POST | `/api/v1/system/user/self/update` | 更新用户基本数据 | all |
| POST | `/api/v1/system/user/self/user/search` | 精准搜索用户 | all |
