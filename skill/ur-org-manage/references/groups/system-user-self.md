# system/user/self

> 该 group 共 30 个端点。

- `POST /api/v1/system/user/self/app/get-list` [all] 获取用户应用列表
- `POST /api/v1/system/user/self/app/get-one` [all] 获取用户应用详情
- `POST /api/v1/system/user/self/bind-account` [all] 绑定账号
- `POST /api/v1/system/user/self/cancel` [all] 注销用户
- `POST /api/v1/system/user/self/captcha` [public] 获取验证码
- `POST /api/v1/system/user/self/change-pwd` [all] 更新用户密码
- `POST /api/v1/system/user/self/forget-pwd` [public] 忘记密码
- `POST /api/v1/system/user/self/get-one` [all] 获取用户信息
- `POST /api/v1/system/user/self/login` [public] 用户登录
- `POST /api/v1/system/user/self/logout` [all] 用户登出
- `POST /api/v1/system/user/self/menu/get-list` [all] 获取用户菜单列表
- `POST /api/v1/system/user/self/message/get-list` [all] 用户消息列表
- `POST /api/v1/system/user/self/message/get-pending` [all] 用户待处理消息
- `POST /api/v1/system/user/self/message/handle` [all] 用户消息标记已处理
- `POST /api/v1/system/user/self/message/mark-all-read` [all] 用户消息全部已读
- `POST /api/v1/system/user/self/message/multi-delete` [all] 用户消息批量删除
- `POST /api/v1/system/user/self/message/multi-is-read` [all] 用户消息批量已读
- `POST /api/v1/system/user/self/message/statistics` [all] 用户消息统计
- `POST /api/v1/system/user/self/notify-preference/read` [all] 用户通知偏好读取
- `POST /api/v1/system/user/self/notify-preference/update` [all] 用户通知偏好更新
- `POST /api/v1/system/user/self/profile/get-list` [all] 获取用户配置列表
- `POST /api/v1/system/user/self/profile/get-one` [all] 获取用户配置详情
- `POST /api/v1/system/user/self/profile/update` [all] 更新用户配置
- `POST /api/v1/system/user/self/register` [public] 普通用户注册
- `POST /api/v1/system/user/self/resource/action/get-list` [all] 获取用户资源动作权限列表
- `POST /api/v1/system/user/self/third-auth/start` [public] 第三方登录授权起跳
- `POST /api/v1/system/user/self/third-login` [public] 第三方登录回调换平台登录态
- `POST /api/v1/system/user/self/third-register` [public] 第三方补全注册
- `POST /api/v1/system/user/self/update` [all] 更新用户基本数据
- `POST /api/v1/system/user/self/user/search` [all] 精准搜索用户
