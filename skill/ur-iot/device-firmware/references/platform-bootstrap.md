# 平台初始化：产品、物模型与设备

平台命令必须在目标环境中先执行 `ur check --json`。创建前查询同名资源，避免
重复产品和设备。平台写入需要管理员权限；本文件不授权删除现有资源。

## 1. 选择身份模式

| 场景 | 产品配置 | 设备侧初始秘密 | 平台设备创建 |
|---|---|---|---|
| 一机一密 | `authMode=2`、`autoRegister=1` | DeviceSecret | 先创建 |
| 一型一密、设备预创建 | `authMode=2`、`autoRegister=2` | ProductSecret | 先创建 |
| 一型一密、自动建设备 | `authMode=2`、`autoRegister=3` | ProductSecret | 注册时自动创建 |
| 自动建设备并绑定 | `authMode=2`、`autoRegister=4` | ProductSecret + 绑定上下文 | 注册/绑定时创建 |

WiFi+BLE 直连设备通常使用 `deviceType=1`、`netType=7`、
`protocolCode=urMqtt`。是否允许自动建设备必须由产品生命周期策略决定，不能为了
省一步操作擅自开启。

## 2. 创建或核对产品

先查询：

```bash
ur api /api/v1/things/product/info/get-list \
  --body '{"keyword":"<product-name>","page":{"page":1,"size":20}}'
```

CLI 暂无经过验证的产品创建 Layer 2 参数封装，创建使用明确的 API，不使用
`ur things product info create` 猜参数：

```bash
ur api /api/v1/things/product/info/create --body '{
  "productName":"<product-name>",
  "authMode":2,
  "deviceType":1,
  "deviceSchemaMode":1,
  "netType":7,
  "protocolCode":"urMqtt",
  "autoRegister":3,
  "onlineHandle":1,
  "status":1,
  "bindLevel":2,
  "compatMode":2
}'
```

按真实需求调整 `autoRegister` 和 `bindLevel`。响应中的 ProductID 用作后续主键。
随后回读详情；响应可能含 ProductSecret，输出重定向到权限 `0600` 的受控文件，
不要贴入对话或提交：

```bash
umask 077
ur api /api/v1/things/product/info/get-one \
  --body '{"productID":"<productID>"}' > <secure-dir>/product.json
```

核对 `authMode`、`autoRegister`、`protocolCode`、`status`、`bindLevel` 和
`compatMode`，不能只看到 `code=200` 就继续。

## 3. 定义物模型

先生成模板并本地校验：

```bash
ur things model template full --output <temp-dir>/model.json
ur things model validate <temp-dir>/model.json
```

设计约束：

- 属性表示持续状态；可控制属性使用 `mode=rw`，传感器值使用 `mode=r`。
- 事件表示瞬时信息，类型为 `info`、`alert` 或 `fault`。
- 行为表示需要执行并返回结果的动作，明确 `dir`、输入和输出。
- identifier 一旦交付就是协议字段；大小写敏感，避免后续重命名。
- 数值定义必须给出类型、上下限和步长；枚举值必须固定。

导入 TSL：

```bash
ur things schema tsl-import -p <productID> --tsl '<validated-tsl-json>' -j
ur things schema get-list -p <productID> -j
```

回读后把实际 identifier、类型、范围、读写模式、行为输入输出作为固件实现清单。
固件不能根据显示名称或历史示例猜 identifier。

## 4. 准备设备身份

### 静态设备

```bash
ur things device info create \
  -p <productID> -d <deviceName> --alias '<device-alias>' -j
ur things device info get-one -p <productID> -d <deviceName> -j
```

DeviceName 是唯一 ID，不是显示名称。若响应返回 DeviceSecret，同样只保存在受控
文件中。固件首次量产写入时必须使用安全存储或出厂分区。

### 动态注册设备

- `autoRegister=2`：先按上面创建设备，再允许它用 ProductSecret 注册。
- `autoRegister=3`：不预建设备，首次注册自动创建。
- 注册前查询设备列表，避免同一 DeviceName 已被其他实体占用。

## 5. 绑定 Token

需要 App/用户绑定时创建一次性 Token：

```bash
ur api /api/v1/things/device/info/bind/token/create --body '{}'
```

将返回 Token 放入受限临时文件，通过 BLE/SoftAP 配网链路传给设备。设备上线后在
`$thing/up/service/{productID}/{deviceName}` 发送 `appBindToken`。轮询状态：

```bash
ur api /api/v1/things/device/info/bind/token/get-one \
  --body '{"token":"<bind-token>"}'
```

`status=2` 表示设备已上报。Token 成功或过期后删除本地明文副本。

## 6. 平台侧冒烟测试

在真实固件接入前，验证物模型和平台链路：

```bash
ur things device mock -p <productID> -d <deviceName> -j
ur things device report -p <productID> -d <deviceName> \
  --params '{"<PropertyID>":<valid-value>}' -j
ur things device log property -p <productID> -d <deviceName>
```

模拟上报成功只能证明平台合同可用，不代表实体设备在线。后续仍需真实 MQTT、
设备回复和物理状态验证。
