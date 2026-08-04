# things/device/edge

> 该 group 共 2 个端点。

- `POST /api/v1/things/device/edge/send/{handle}/{type}` [device] 设备使用http协议用云端交互,需要在http头中带上mqtt的账号密码(basic auth)
- `POST /api/v1/things/device/edge/upload-file` [device] 设备文件直传,需要在http头中带上mqtt的账号密码(basic auth)
