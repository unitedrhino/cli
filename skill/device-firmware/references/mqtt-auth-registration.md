# 动态注册与 MQTT 认证

本文件描述密钥认证设备的线协议。示例中的地址、ProductID、DeviceName 和秘密均为
占位符；实际值只从目标环境和产品详情取得。

## 连接前置条件

1. 设备已获得网络地址并完成 DNS 检查。
2. SNTP 校时成功；以下 `timestamp`、`expiry` 都是 Unix 秒，不是毫秒。
3. ProductID 和 DeviceName 稳定且符合产品配置。
4. DeviceSecret 为空时才走动态注册；已有有效 DeviceSecret 时直接正常连接。

MQTT 地址由目标环境提供，不写死在固件。常用端口为 `1883`（MQTT）、`8883`
（MQTTS）、`8083`（WS）和 `8084`（WSS）；生产固件优先使用 TLS，并验证服务端
证书、主机名和系统时间。

## 动态注册

### 主题和临时连接

```text
clientID: register&{productID}&{deviceName}
publish:  $ext/up/register/{productID}/{deviceName}
subscribe:$ext/down/register/{productID}/{deviceName}
```

临时注册连接先订阅下行并确认订阅成功，再发布请求。当前协议不依赖正式设备
用户名和密码，但固件仍应使用受支持的占位值，不把 ProductSecret 放到用户名、
密码或日志中。

### 签名

生成随机 `nonce` 和当前 Unix 秒 `timestamp`，签名原文严格为：

```text
deviceName={deviceName}&nonce={nonce}&productID={productID}&timestamp={timestamp}
```

使用 ProductSecret 的原始 ASCII 字节作为 HMAC-SHA1 key。将 20 字节摘要转为
40 个小写十六进制 ASCII 字符，再对该 ASCII 字符串做 Base64，得到
`signature`。ProductSecret 在签名时不能先 Base64 解码。

请求：

```json
{
  "msgToken": "<unique-token>",
  "payload": {
    "nonce": 2125656451,
    "timestamp": 1687525491,
    "signature": "<base64-of-lowercase-hmac-sha1-hex>",
    "retEnc": "aes128cbc"
  }
}
```

部分企业上下文要求 `payload.tenantCode`；只在目标平台明确要求时添加，不在固件中
猜测或硬编码。

### 响应解密和持久化

成功响应的密文优先读取 `data.payload`，兼容旧响应的顶层 `payload`：

1. Base64 解码密文。
2. AES-128-CBC 解密；key 为 ProductSecret 前 16 个 ASCII 字节，IV 为 16 个
   ASCII 字符 `0`。
3. 按实现使用的标准填充规则去除尾部填充，并解析 JSON。
4. 密钥认证响应校验 `encryptionType` 后读取 `psk` 作为 DeviceSecret。
5. 先写受保护持久化存储并回读，再更新运行时缓存；任何一步失败都不能标记注册成功。

不要打印完整请求、响应、密文解密结果、ProductSecret 或 DeviceSecret。动态注册
成功后断开临时客户端，再以正式身份连接。

## 正常 MQTT 鉴权

当前密钥认证推荐 HMAC-SHA256：

```text
clientID = {productID}&{deviceName}
username = {clientID};12010126;{connID};{expiry}
key      = Base64Decode(DeviceSecret)
digest   = lowercase_hex(HMAC-SHA256(key, username))
password = {digest};hmacsha256
```

- `connID` 使用每次连接新生成的短随机串；当前客户端常用 5 个字符。
- `expiry` 必须晚于服务器当前时间，单位为秒。
- ProductID、DeviceName、分号位置和算法后缀必须完全一致。
- 平台兼容的其他算法不应成为新固件默认值。

连接后先订阅属性、事件、行为、服务和 OTA 所需的下行主题，确认订阅完成，再发送
`reportInfo` 和首组属性。重连时必须重新订阅，不能假设 broker 保留会话。

## 重试与鉴权自愈

把错误分成两类：

- DNS、TCP、TLS、WiFi、超时等网络错误：保留 WiFi 和 DeviceSecret，只重连。
- SDK 明确返回 `-116` 或 `-117` 的鉴权错误：连续 3 次后只清 DeviceSecret，保留
  WiFi，重新动态注册。

注册重试使用 `30/60/120/240/300` 秒封顶退避；成功连接并稳定完成订阅后清零计数。
清密钥接口必须同时清持久化值和内存缓存，空值写入也必须生效。不同 SDK 的错误码
若不一致，先把 broker CONNACK/SDK 错误映射到“明确鉴权失败”，不能照搬数字。

## 最小诊断输出

允许 INFO 日志打印阶段和脱敏标识：网络就绪、时钟有效、是否需要注册、注册结果、
MQTT 建连结果、订阅数量、重连原因和退避时间。禁止打印 WiFi 密码、完整秘密、绑定
Token、完整动态注册响应或带签名的鉴权串。
