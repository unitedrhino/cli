# ur-device things/device/edge

设备使用http协议用云端交互,需要在http头中带上mqtt的账号密码(basic auth) 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/device/edge/send/{handle}/{type}` | 设备使用http协议用云端交互,需要在http头中带上mqtt的账号密码(basic auth) | device |
| POST | `/api/v1/things/device/edge/upload-file` | 设备文件直传,需要在http头中带上mqtt的账号密码(basic auth) | device |

## 详细说明

### POST `/api/v1/things/device/edge/send/{handle}/{type}`

**说明**: 设备使用http协议用云端交互,需要在http头中带上mqtt的账号密码(basic auth)

**权限**: device

**路径参数**:

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| `handle` | path | string | 是 | 对应 mqtt topic的第一个 thing ota config 等等 |
| `type` | path | string | 是 | 操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为 |
| `productID` | header | string | 否 | 如果网关类型要操作子设备的topic,需要指定子设备的产品ID和设备ID,如果没有填写则默认是账号里的设备 |
| `deviceName` | header | string | 否 | 如果网关类型要操作子设备的topic,需要指定子设备的产品ID和设备ID,如果没有填写则默认是账号里的设备 |
| `authorization` | header | string | 是 | basic auth  账号密码使用mqtt格式生成的账号密码 |

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/edge/send/{handle}/{type} \
  --body '{}'
```

### POST `/api/v1/things/device/edge/upload-file`

**说明**: 设备文件直传,需要在http头中带上mqtt的账号密码(basic auth)

**权限**: device

**路径参数**:

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| `authorization` | header | string | 是 |  basic auth 账号密码使用mqtt格式生成的账号密码 |

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authorization` | string | 是 | // basic auth 账号密码使用mqtt格式生成的账号密码 |
| `file` | string | 是 | 格式: binary |

**请求示例**:
```json
{
  "authorization": "******",
  "file": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "filePath": "string",
    "fileUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/device/edge/upload-file \
  --body '{"authorization": "******", "file": "string"}'
```
