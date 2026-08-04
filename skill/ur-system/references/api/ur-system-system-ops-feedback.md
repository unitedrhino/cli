# ur-system system/ops/feedback

添加帮助与反馈 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/ops/feedback/create` | 添加帮助与反馈 | all |
| POST | `/api/v1/system/ops/feedback/get-list` | 获取帮助与反馈 | all |
| POST | `/api/v1/system/ops/feedback/update` | 更新帮助与反馈 | all |

## 详细说明

### POST `/api/v1/system/ops/feedback/create`

**说明**: 添加帮助与反馈

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 | 所属应用ID |
| `contactInformation` | string | 否 | 联系方式 |
| `createdTime` | string | 否 | 创建时间 |
| `detail` | string | 否 | 详细信息（页面URL + AI对话历史） |
| `feedbackType` | string | 否 | 问题反馈类型: usageProblem/businessImpaired/businessUnavailable |
| `feedbackUserID` | string | 否 | 反馈用户ID |
| `id` | string | 否 |  |
| `issueDesc` | string | 否 | 反馈内容 |
| `projectID` | string | 否 |  |
| `raiseUserID` | string | 否 |  |
| `status` | integer | 否 |  状态（1 待处理 2:处理中 3:已完成） (格式: int64) |
| `tenantCode` | string | 否 |  |
| `type` | string | 否 | 反馈类型 设备问题:thingsDevice 智能场景:thingsScene 体验问题: experience 其他: other |
| `user` | object | 否 |  |
| `user.avatar` | string | 否 |  用户头像 |
| `user.createdTime` | string | 否 |  创建时间 |
| `user.email` | string | 否 |  邮箱 |
| `user.lastIP` | string | 否 |  最后登录ip |
| `user.nickName` | string | 否 |  用户的昵称 |
| `user.phone` | string | 否 |  手机号 |
| `user.regIP` | string | 否 |  注册ip |
| `user.userID` | string | 否 |  用户id |
| `user.userName` | string | 否 |  用户名(唯一) |

**请求示例**:
```json
{
  "appID": "string",
  "contactInformation": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "detail": "string",
  "feedbackType": "string",
  "feedbackUserID": "string",
  "id": "string",
  "issueDesc": "string",
  "projectID": "string",
  "raiseUserID": "string",
  "status": 1,
  "tenantCode": "string",
  "type": "string",
  "user": {
    "avatar": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "email": "string",
    "lastIP": "string",
    "nickName": "string",
    "phone": "string",
    "regIP": "string",
    "userID": "string",
    "userName": "string"
  }
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
ur api /api/v1/system/ops/feedback/create \
  --body '{"appID": "string", "contactInformation": "string", "createdTime": "2026-01-01T00:00:00Z", "detail": "string", "feedbackType": "string", "feedbackUserID": "string", "id": "string", "issueDesc": "string", "projectID": "string", "raiseUserID": "string", "status": 1, "tenantCode": "string", "type": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}}'
```

### POST `/api/v1/system/ops/feedback/get-list`

**说明**: 获取帮助与反馈

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 | 按应用过滤 |
| `feedbackType` | string | 否 | 按问题反馈类型过滤 |
| `isAllTenant` | boolean | 否 | default租户获取所有租户的信息 (格式: boolean) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `projectID` | string | 否 | 所属项目 |
| `status` | integer | 否 |  状态（1 待处理 2:处理中 3:已完成） (格式: int64) |
| `tenantCode` | string | 否 | 过滤所属租户,isAllTenant为true有效 |
| `type` | string | 否 | 反馈类型 设备问题:thingsDevice 智能场景:thingsScene 体验问题: experience 其他: other |

**请求示例**:
```json
{
  "appID": "string",
  "feedbackType": "string",
  "isAllTenant": true,
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
  },
  "projectID": "string",
  "status": 1,
  "tenantCode": "string",
  "type": "string"
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
        "contactInformation": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "detail": "string",
        "feedbackType": "string",
        "feedbackUserID": "string",
        "id": "string",
        "issueDesc": "string",
        "projectID": "string",
        "raiseUserID": "string",
        "status": 1,
        "tenantCode": "string",
        "type": "string",
        "user": {
          "avatar": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "email": "string",
          "lastIP": "string",
          "nickName": "string",
          "phone": "string",
          "regIP": "string",
          "userID": "string",
          "userName": "string"
        }
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
ur api /api/v1/system/ops/feedback/get-list \
  --body '{"appID": "string", "feedbackType": "string", "isAllTenant": true, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "projectID": "string", "status": 1, "tenantCode": "string", "type": "string"}'
```

### POST `/api/v1/system/ops/feedback/update`

**说明**: 更新帮助与反馈

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 | 所属应用ID |
| `contactInformation` | string | 否 | 联系方式 |
| `createdTime` | string | 否 | 创建时间 |
| `detail` | string | 否 | 详细信息（页面URL + AI对话历史） |
| `feedbackType` | string | 否 | 问题反馈类型: usageProblem/businessImpaired/businessUnavailable |
| `feedbackUserID` | string | 否 | 反馈用户ID |
| `id` | string | 否 |  |
| `issueDesc` | string | 否 | 反馈内容 |
| `projectID` | string | 否 |  |
| `raiseUserID` | string | 否 |  |
| `status` | integer | 否 |  状态（1 待处理 2:处理中 3:已完成） (格式: int64) |
| `tenantCode` | string | 否 |  |
| `type` | string | 否 | 反馈类型 设备问题:thingsDevice 智能场景:thingsScene 体验问题: experience 其他: other |
| `user` | object | 否 |  |
| `user.avatar` | string | 否 |  用户头像 |
| `user.createdTime` | string | 否 |  创建时间 |
| `user.email` | string | 否 |  邮箱 |
| `user.lastIP` | string | 否 |  最后登录ip |
| `user.nickName` | string | 否 |  用户的昵称 |
| `user.phone` | string | 否 |  手机号 |
| `user.regIP` | string | 否 |  注册ip |
| `user.userID` | string | 否 |  用户id |
| `user.userName` | string | 否 |  用户名(唯一) |

**请求示例**:
```json
{
  "appID": "string",
  "contactInformation": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "detail": "string",
  "feedbackType": "string",
  "feedbackUserID": "string",
  "id": "string",
  "issueDesc": "string",
  "projectID": "string",
  "raiseUserID": "string",
  "status": 1,
  "tenantCode": "string",
  "type": "string",
  "user": {
    "avatar": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "email": "string",
    "lastIP": "string",
    "nickName": "string",
    "phone": "string",
    "regIP": "string",
    "userID": "string",
    "userName": "string"
  }
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
ur api /api/v1/system/ops/feedback/update \
  --body '{"appID": "string", "contactInformation": "string", "createdTime": "2026-01-01T00:00:00Z", "detail": "string", "feedbackType": "string", "feedbackUserID": "string", "id": "string", "issueDesc": "string", "projectID": "string", "raiseUserID": "string", "status": 1, "tenantCode": "string", "type": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}}'
```
