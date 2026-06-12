# 跨域工作流示例

## 工作流 1：全新 IoT 设备从零上线

**角色**：企业管理员 | **涉及域**：ur-product → ur-device → ur-ai（可选）

```bash
# 1. 创建产品（ur-product）
ur api /api/v1/things/product/info/create \
  --body '{"name":"智能开关","protocolCode":"urMqtt","deviceType":1}'
# → 获得 productID

# 2. 定义物模型属性（ur-product）
ur api /api/v1/things/schema/create \
  --body '{"productID":"<step1>","identifier":"Power","name":"开关","accessMode":"rw","dataType":{"type":"bool"}}'

# 3. 创建设备（ur-device）
ur api /api/v1/things/device/info/create \
  --body '{"productID":"<step1>","deviceName":"switch-001","deviceAlias":"一楼走廊开关"}'
# → 获得三元组（含 deviceSecret）

# 4. 验证属性（ur-device）
ur api /api/v1/things/device/msg/property-latest/get-list \
  --body '{"productID":"<step1>","deviceName":"switch-001"}'

# 5. 可选：创建场景联动（ur-ai）
ur api /api/v1/things/scene/info/create \
  --body '{"name":"自动关灯","triggers":[{"type":"timer","cron":"0 22 * * *"}]}'
```

**关键依赖**：步骤 3 依赖步骤 1 的 productID；设备控制依赖设备在线（isOnline=1）。

## 工作流 2：新企业完整初始化

**角色**：平台管理员 | **涉及域**：ur-tenant → ur-system → ur-tenant

```bash
# 1. 创建企业（ur-tenant）
ur api /api/v1/system/tenant/info/create \
  --body '{"name":"Acme Corp","tenantCode":"acme"}'

# 2. 为企业绑定应用（ur-system）
ur api /api/v1/system/tenant/app/create \
  --body '{"tenantCode":"acme","appID":77}'

# 3. 邀请企业管理员（ur-tenant）
ur api /api/v1/system/tenant/user/invite \
  --body '{"tenantCode":"acme","account":"admin@acme.com"}'
```

## 工作流 3：设备超阈值告警 + 通知

**角色**：企业管理员 | **涉及域**：ur-ai → ur-user

```bash
# 1. 创建告警规则（ur-ai）
ur api /api/v1/things/alarm/info/create \
  --body '{"name":"温度过高","productID":"<productID>","alarmLevel":1}'

# 2. 创建通知配置（ur-user）
ur api /api/v1/system/notify/config/create \
  --body '{"name":"设备告警通知","throttle":300}'

# 3. 创建通知模板（ur-user）
ur api /api/v1/system/notify/template/create \
  --body '{"configID":"<step2>","notifyType":"inner","template":"设备${deviceName}告警"}'
```
