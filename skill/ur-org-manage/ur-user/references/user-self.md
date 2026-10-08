# CLI 命令 — 个人中心

## 命令语法
```
ur user self <subcommand> [选项]
```

## 子命令

| 子命令 | 说明 |
|--------|------|
| `login` | 用户登录 |
| `logout` | 用户登出 |
| `register` | 用户注册 |
| `get-one` | 获取当前用户信息 |
| `update` | 更新当前用户信息 |
| `change-pwd` | 修改密码 |
| `forget-pwd` | 忘记密码 |
| `captcha` | 获取验证码 |

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body | | 条件 | string | 请求体 JSON（login/register/update/change-pwd/forget-pwd 必填） |
| --json | -j | 否 | bool | 输出 JSON 格式 |

## 使用示例

### 示例 1：用户登录
```bash
ur user self login --body '{"account":"admin","password":"123456","loginType":"pwd"}'
```

### 示例 2：获取当前用户信息
```bash
ur user self get-one
```

### 示例 3：修改密码
```bash
ur user self change-pwd --body '{"oldPwd":"123456","newPwd":"abcdef"}'
```

### 示例 4：用户注册
```bash
ur user self register --body '{"account":"newuser","password":"123456"}'
```

### 示例 5：获取验证码
```bash
ur user self captcha --body '{"account":"admin","type":"sms"}'
```

### 示例 6：用户登出
```bash
ur user self logout
```

## 对应 API

| 子命令 | API 端点 |
|--------|----------|
| login | `POST /api/v1/system/user/self/login` |
| logout | `POST /api/v1/system/user/self/logout` |
| register | `POST /api/v1/system/user/self/register` |
| get-one | `POST /api/v1/system/user/self/get-one` |
| update | `POST /api/v1/system/user/self/update` |
| change-pwd | `POST /api/v1/system/user/self/change-pwd` |
| forget-pwd | `POST /api/v1/system/user/self/forget-pwd` |
| captcha | `POST /api/v1/system/user/self/captcha` |
