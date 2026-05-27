# ur-iot-config things/config/info

获取全局配置 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/config/info/get-one` | 获取全局配置 | all |
| POST | `/api/v1/things/config/info/update` | 更新全局配置 | platform |

## 详细说明

### POST `/api/v1/things/config/info/get-one`

**说明**: 获取全局配置

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "cfgType": "string",
    "configID": "string",
    "enabled": true,
    "id": 1,
    "inputModalities": "string",
    "isDefault": true,
    "jsonData": "string",
    "modelType": "string",
    "name": "示例名称",
    "outputModalities": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/config/info/get-one \
  --body '{}'
```

### POST `/api/v1/things/config/info/update`

**说明**: 更新全局配置

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cfgType` | string | 否 |  配置类型：llm/tts/asr/vad |
| `configID` | string | 否 |  配置唯一标识 |
| `enabled` | boolean | 否 |  是否启用 (格式: boolean) |
| `id` | integer | 否 |  配置ID (格式: int64) |
| `inputModalities` | string | 否 |  输入模态（JSON字符串） |
| `isDefault` | boolean | 否 |  是否默认 (格式: boolean) |
| `jsonData` | string | 否 |  配置参数（JSON字符串） |
| `modelType` | string | 否 |  模型级别：small/medium/large/xlarge |
| `name` | string | 否 |  配置名称 |
| `outputModalities` | string | 否 |  输出模态（JSON字符串） |
| `tenantCode` | string | 否 |  租户编码 |

**请求示例**:
```json
{
  "cfgType": "string",
  "configID": "string",
  "enabled": true,
  "id": 1,
  "inputModalities": "string",
  "isDefault": true,
  "jsonData": "string",
  "modelType": "string",
  "name": "示例名称",
  "outputModalities": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/config/info/update \
  --body '{"cfgType": "string", "configID": "string", "enabled": true, "id": 1, "inputModalities": "string", "isDefault": true, "jsonData": "string", "modelType": "string", "name": "示例名称", "outputModalities": "string", "tenantCode": "string"}'
```
