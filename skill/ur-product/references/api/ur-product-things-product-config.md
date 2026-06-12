# ur-product things/product/config

更新配置

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/product/config/update` | 更新配置 | admin |

## 详细说明

### POST `/api/v1/things/product/config/update`

**说明**: 更新配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `devInit` | object | 否 |  |
| `devInit.areaID` | string | 否 | 设备创建默认分配给的区域ID |
| `devInit.deptID` | string | 否 | 设备创建默认分配给的部门 |
| `devInit.projectID` | string | 否 | 设备创建默认分配给的项目 |
| `devInit.tenantCode` | string | 否 | 设备创建默认分配给的企业 |
| `productID` | string | 是 |  |

**请求示例**:
```json
{
  "devInit": {
    "areaID": "string",
    "deptID": "string",
    "projectID": "string",
    "tenantCode": "string"
  },
  "productID": "string"
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
ur api /api/v1/things/product/config/update \
  --body '{"devInit": {"areaID": "string", "deptID": "string", "projectID": "string", "tenantCode": "string"}, "productID": "string"}'
```
