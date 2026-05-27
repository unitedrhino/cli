# ur-user system/log

获取登录日志列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/log/login/get-list` | 获取登录日志列表 | admin |
| POST | `/api/v1/system/log/oper/get-list` | 获取操作日志列表 | admin |

## 详细说明

### POST `/api/v1/system/log/login/get-list`

**说明**: 获取登录日志列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  应用编号 |
| `code` | string | 否 |  登录状态（200成功 其它失败） |
| `dateRange` | object | 否 |  |
| `dateRange.end` | string | 否 | 结束时间 格式：yyyy-mm-dd |
| `dateRange.start` | string | 否 | 开始时间 格式：yyyy-mm-dd |
| `ipAddr` | string | 否 | 按ip地址查找 |
| `loginLocation` | string | 否 | 按登录地址查找 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `userID` | string | 否 |  用户id |
| `userName` | string | 否 |  登录账号 |

**请求示例**:
```json
{
  "appID": "string",
  "code": "string",
  "dateRange": {
    "end": "2026-01-01T00:00:00Z",
    "start": "2026-01-01T00:00:00Z"
  },
  "ipAddr": "string",
  "loginLocation": "string",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "userID": "string",
  "userName": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "appID": "string",
        "browser": "string",
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "ipAddr": "string",
        "loginLocation": "string",
        "msg": "string",
        "os": "string",
        "userID": "string",
        "userName": "string"
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
ur api /api/v1/system/log/login/get-list \
  --body '{"appID": "string", "code": "string", "dateRange": {"end": "2026-01-01T00:00:00Z", "start": "2026-01-01T00:00:00Z"}, "ipAddr": "string", "loginLocation": "string", "page": {"page": 1, "pageSize": 1}, "userID": "string", "userName": "string"}'
```

### POST `/api/v1/system/log/oper/get-list`

**说明**: 获取操作日志列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessKey` | string | 否 | open token 的 accessKey |
| `appID` | string | 否 |  应用编号 |
| `authType` | string | 否 | 认证方式（user/open） |
| `code` | string | 否 | 登录状态（200成功 其它失败） |
| `operName` | string | 否 | 按操作名称查找 |
| `operType` | string | 否 | 操作类型（create/update/delete/get/other） |
| `operUserID` | string | 否 |  操作用户id |
| `operUserName` | string | 否 | 按操作人员名称查找 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "accessKey": "string",
  "appID": "string",
  "authType": "string",
  "code": "string",
  "operName": "示例名称",
  "operType": "string",
  "operUserID": "string",
  "operUserName": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "accessKey": "string",
        "appID": "string",
        "authType": "string",
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "msg": "string",
        "operIpAddr": "string",
        "operLocation": "string",
        "operName": "示例名称",
        "operType": "string",
        "operUserName": "示例名称",
        "req": "string",
        "resp": "string",
        "uri": "string",
        "userID": "string"
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
ur api /api/v1/system/log/oper/get-list \
  --body '{"accessKey": "string", "appID": "string", "authType": "string", "code": "string", "operName": "示例名称", "operType": "string", "operUserID": "string", "operUserName": "示例名称", "page": {"page": 1, "pageSize": 1}}'
```
