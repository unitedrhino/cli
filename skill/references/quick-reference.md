# 高频端点速查（Top 15）

| 操作 | 端点 | 最小请求体 |
|------|------|-----------|
| 设备列表 | POST /api/v1/things/device/info/get-list | `{"page":{"page":1,"size":10}}` |
| 当前用户信息 | POST /api/v1/system/user/self/get-one | `{}` |
| 控制设备属性 | POST /api/v1/things/device/interact/property-control-send | `{"productID":"x","deviceName":"x","data":{"Key":值}}` |
| 设备最新属性 | POST /api/v1/things/device/msg/property-latest/get-list | `{"productID":"x","deviceName":"x"}` |
| 产品列表 | POST /api/v1/things/product/info/get-list | `{"page":{"page":1,"size":10}}` |
| 产品物模型 | POST /api/v1/things/product/schema/get-list | `{"productID":"x"}` |
| 项目列表 | POST /api/v1/things/project/crud/get-list | `{"page":{"page":1,"size":10}}` |
| 创建访问令牌 | POST /api/v1/system/user/self/access-token/create | `{"name":"my-token","expTime":86400}` |
| 企业列表（平台管理员）| POST /api/v1/system/tenant/info/get-list | `{"page":{"page":1,"size":10}}` |
| 场景列表 | POST /api/v1/things/scene/info/get-list | `{"page":{"page":1,"size":10}}` |
| 手动触发场景 | POST /api/v1/things/scene/info/manually-trigger | `{"id":"<sceneID>"}` |
| 告警规则列表 | POST /api/v1/things/alarm/info/get-list | `{"page":{"page":1,"size":10}}` |
| 文件上传 | POST /api/v1/system/common/upload-file | multipart/form-data |
| 批量接口调用 | POST /api/v1/system/common/api/batch-agg | `{"apis":[{"path":"/api/v1/system/user/self/get-one","body":{}}]}` |
| 提交问题反馈 | POST /api/v1/system/ops/feedback/create | `{"feedbackType":"usageProblem","issueDesc":"...","contactInformation":"...","detail":"..."}` |
