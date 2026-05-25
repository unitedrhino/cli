# ur-system system/checkIn

用户签到 等

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/check-in/do` | 用户签到 | admin |
| POST | `/api/v1/system/check-in/get-list` | 签到记录列表 | admin |
| POST | `/api/v1/system/check-in/point-balance/get` | 获取当前用户积分余额 | admin |
| POST | `/api/v1/system/check-in/point-log/adjust` | 管理员调整积分 | admin |
| POST | `/api/v1/system/check-in/point-log/get-list` | 积分流水列表 | admin |
