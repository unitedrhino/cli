# ur-system system/resource/api

添加接口 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/resource/api/create` | 添加接口 | platform |
| POST | `/api/v1/system/resource/api/delete` | 删除接口 | platform |
| POST | `/api/v1/system/resource/api/get-list` | 获取接口列表 | platform |
| POST | `/api/v1/system/resource/api/update` | 更新接口 | platform |

## 详细说明

### POST `/api/v1/system/resource/api/create`

**说明**: 添加接口

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessCode` | string | 是 |  模块编号 |
| `authTypeCode` | string | 否 |  权限层级编码 platform/admin/all/public/device |
| `desc` | string | 否 |  备注 |
| `id` | string | 否 |  接口编号 |
| `method` | string | 否 |  接口请求方式: （1 GET 2 POST 3 HEAD 4 OPTIONS 5 PUT 6 DELETE 7 TRACE 8 CONNECT 9 其它） |
| `name` | string | 否 |  接口名称 |
| `operType` | string | 否 |  操作类型（create/update/delete/get/other） |
| `recordLogMode` | integer | 否 |    1为自动模式(读取类型忽略,其他类型记录日志) 2全部记录 3不记录 (格式: int64) |
| `route` | string | 否 |  接口路由 |

**请求示例**:
```json
{
  "accessCode": "string",
  "authTypeCode": "string",
  "desc": "string",
  "id": "string",
  "method": "string",
  "name": "示例名称",
  "operType": "string",
  "recordLogMode": 1,
  "route": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/resource/api/create \
  --body '{"accessCode": "string", "authTypeCode": "string", "desc": "string", "id": "string", "method": "string", "name": "示例名称", "operType": "string", "recordLogMode": 1, "route": "string"}'
```

### POST `/api/v1/system/resource/api/delete`

**说明**: 删除接口

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  资源ID (格式: int64) |

**请求示例**:
```json
{
  "id": 1
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
ur api /api/v1/system/resource/api/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/resource/api/get-list`

**说明**: 获取接口列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessCode` | string | 否 |  应用编号 |
| `authType` | integer | 否 |   1(all) 全部人可以操作 2(admin) 默认授予企业管理员权限 3(superAdmin,supper) default企业才可以操作(超管是跨企业的) (格式: int64) |
| `authTypeCode` | string | 否 |  权限层级编码 |
| `groupCode` | string | 否 |  分组标识 |
| `method` | string | 否 |  接口请求方式: （1 GET 2 POST 3 HEAD 4 OPTIONS 5 PUT 6 DELETE 7 TRACE 8 CONNECT 9 其它） |
| `name` | string | 否 |  接口名称 |
| `operTypeCode` | string | 否 |  操作类型编码 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `resource` | string | 否 |  资源标识 |
| `route` | string | 否 |  接口路由 |

**请求示例**:
```json
{
  "accessCode": "string",
  "authType": 1,
  "authTypeCode": "string",
  "groupCode": "string",
  "method": "string",
  "name": "示例名称",
  "operTypeCode": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "resource": "string",
  "route": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "groups": [
      {
        "children": [
          {
            "children": [],
            "code": "string",
            "id": "string",
            "name": "示例名称"
          }
        ],
        "code": "string",
        "id": "string",
        "name": "示例名称"
      }
    ],
    "list": [
      {
        "accessCode": "string",
        "authTypeCode": "string",
        "desc": "string",
        "id": "string",
        "method": "string",
        "name": "示例名称",
        "operType": "string",
        "recordLogMode": 1,
        "route": "string"
      }
    ],
    "page": 1,
    "pageSize": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/resource/api/get-list \
  --body '{"accessCode": "string", "authType": 1, "authTypeCode": "string", "groupCode": "string", "method": "string", "name": "示例名称", "operTypeCode": "string", "page": {"page": 1, "pageSize": 1}, "resource": "string", "route": "string"}'
```

### POST `/api/v1/system/resource/api/update`

**说明**: 更新接口

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessCode` | string | 是 |  模块编号 |
| `authTypeCode` | string | 否 |  权限层级编码 platform/admin/all/public/device |
| `desc` | string | 否 |  备注 |
| `id` | string | 否 |  接口编号 |
| `method` | string | 否 |  接口请求方式: （1 GET 2 POST 3 HEAD 4 OPTIONS 5 PUT 6 DELETE 7 TRACE 8 CONNECT 9 其它） |
| `name` | string | 否 |  接口名称 |
| `operType` | string | 否 |  操作类型（create/update/delete/get/other） |
| `recordLogMode` | integer | 否 |    1为自动模式(读取类型忽略,其他类型记录日志) 2全部记录 3不记录 (格式: int64) |
| `route` | string | 否 |  接口路由 |

**请求示例**:
```json
{
  "accessCode": "string",
  "authTypeCode": "string",
  "desc": "string",
  "id": "string",
  "method": "string",
  "name": "示例名称",
  "operType": "string",
  "recordLogMode": 1,
  "route": "string"
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
ur api /api/v1/system/resource/api/update \
  --body '{"accessCode": "string", "authTypeCode": "string", "desc": "string", "id": "string", "method": "string", "name": "示例名称", "operType": "string", "recordLogMode": 1, "route": "string"}'
```
