# ur-iot-user things/user/area/apply

申请用户区域权限

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/user/area/apply/create` | 申请用户区域权限 | all |

## 详细说明

### POST `/api/v1/things/user/area/apply/create`

**说明**: 申请用户区域权限

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | integer | 是 |  区域ID (格式: int64) |
| `authType` | integer | 是 |  权限类型 1:读权限 4:管理权限 (格式: int64) |

**请求示例**:
```json
{
  "areaID": 1,
  "authType": 1
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
ur api /api/v1/things/user/area/apply/create \
  --body '{"areaID": 1, "authType": 1}'
```
