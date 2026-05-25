# ur-system system/common

批量聚合接口请求 等

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/common/api/batch-agg` | 批量聚合接口请求 | all |
| GET | `/api/v1/system/common/debug` | 调试接口GET | public |
| POST | `/api/v1/system/common/debug` | 调试接口POST | public |
| GET | `/api/v1/system/common/debug-tencent` | 腾讯云调试接口 | public |
| GET | `/api/v1/system/common/download-file` | 下载本地文件 | public |
| POST | `/api/v1/system/common/init-upload-file` | 初始化上传文件 | public |
| POST | `/api/v1/system/common/ntp/get-one` | ntp时间同步 | public |
| POST | `/api/v1/system/common/qr-code/get-one` | 获取小程序二维码 | all |
| POST | `/api/v1/system/common/sys-config/info/get-one` | 读取系统配置信息 | platform |
| POST | `/api/v1/system/common/sys-config/info/update` | 更新系统配置信息 | platform |
| POST | `/api/v1/system/common/third/dept/get-list` | 获取第三方部门列表 | all |
| POST | `/api/v1/system/common/third/dept/get-one` | 获取第三方部门详情 | all |
| POST | `/api/v1/system/common/upload-file` | 文件直传 | all |
| POST | `/api/v1/system/common/upload-url/create` | 获取文件上传地址 | all |
| POST | `/api/v1/system/common/weather/get-one` | 获取天气情况 | all |
| GET | `/api/v1/system/common/websocket/connect` | websocket连接 | all |
