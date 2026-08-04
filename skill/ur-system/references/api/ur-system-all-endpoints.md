# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/agreement/create` | 创建协议 | platform |
| POST | `/api/v1/system/agreement/delete` | 删除协议 | platform |
| POST | `/api/v1/system/agreement/get-list` | 获取协议列表 | platform |
| POST | `/api/v1/system/agreement/get-one` | 获取协议详情 | platform |
| POST | `/api/v1/system/agreement/preview-render` | 预览渲染协议 | platform |
| POST | `/api/v1/system/agreement/update` | 更新协议 | platform |
| POST | `/api/v1/system/app/agreement/bind-batch-update` | 批量更新应用协议绑定 | platform |
| POST | `/api/v1/system/app/agreement/get-bind-list` | 获取应用协议绑定列表 | platform |
| POST | `/api/v1/system/app/core/get-one` | 无需登录获取应用信息 | public |
| POST | `/api/v1/system/app/info/create` | 添加应用 | platform |
| POST | `/api/v1/system/app/info/delete` | 删除应用 | platform |
| POST | `/api/v1/system/app/info/get-list` | 获取应用列表 | platform |
| POST | `/api/v1/system/app/info/get-one` | 获取应用详情 | platform |
| POST | `/api/v1/system/app/info/update` | 更新应用 | platform |
| POST | `/api/v1/system/app/menu/create` | 添加菜单 | platform |
| POST | `/api/v1/system/app/menu/delete` | 删除菜单 | platform |
| POST | `/api/v1/system/app/menu/get-list` | 获取菜单列表 | platform |
| POST | `/api/v1/system/app/menu/update` | 更新菜单 | platform |
| POST | `/api/v1/system/check-in/do` | 用户签到 | admin |
| POST | `/api/v1/system/check-in/get-list` | 签到记录列表 | admin |
| POST | `/api/v1/system/check-in/point-balance/get` | 获取当前用户积分余额 | admin |
| POST | `/api/v1/system/check-in/point-log/adjust` | 管理员调整积分 | admin |
| POST | `/api/v1/system/check-in/point-log/get-list` | 积分流水列表 | admin |
| POST | `/api/v1/system/common/api/batch-agg` | 批量聚合接口请求 | all |
| GET | `/api/v1/system/common/debug` | 调试接口GET | public |
| POST | `/api/v1/system/common/debug` | 调试接口POST | public |
| GET | `/api/v1/system/common/debug-tencent` | 腾讯云调试接口 | public |
| GET | `/api/v1/system/common/download-file` | 下载本地文件 | public |
| POST | `/api/v1/system/common/init-upload-file` | 初始化上传文件 | public |
| POST | `/api/v1/system/common/ntp/get-one` | ntp时间同步 | public |
| POST | `/api/v1/system/common/qr-code/get-one` | 获取小程序二维码 | all |
| POST | `/api/v1/system/common/sys-config/info/get-one` | 读取系统配置信息 | platform |
| POST | `/api/v1/system/common/sys-config/info/update` | 更新系统配置信息 | platform |
| POST | `/api/v1/system/common/third/dept/get-list` | 获取第三方部门列表 | all |
| POST | `/api/v1/system/common/third/dept/get-one` | 获取第三方部门详情 | all |
| POST | `/api/v1/system/common/upload-file` | 文件直传 | all |
| POST | `/api/v1/system/common/upload-url/create` | 获取文件上传地址 | all |
| POST | `/api/v1/system/common/weather/get-one` | 获取天气情况 | all |
| GET | `/api/v1/system/common/websocket/connect` | websocket连接 | all |
| POST | `/api/v1/system/common/sys-config/core/get-one` | 读取系统配置信息(无需登录) | public |
| POST | `/api/v1/system/common/system/init` | 初始化系统 | public |
| POST | `/api/v1/system/hook/capability/create` | 新增Hook能力 | platform |
| POST | `/api/v1/system/hook/capability/delete` | 删除Hook能力 | platform |
| POST | `/api/v1/system/hook/capability/get-list` | 获取Hook能力列表 | platform |
| POST | `/api/v1/system/hook/capability/update` | 更新Hook能力 | platform |
| POST | `/api/v1/system/hook/server/create` | 新增Hook服务 | platform |
| POST | `/api/v1/system/hook/server/delete` | 删除Hook服务 | platform |
| POST | `/api/v1/system/hook/server/get-list` | 获取Hook服务列表 | platform |
| POST | `/api/v1/system/hook/server/get-one` | 获取Hook服务详情 | platform |
| POST | `/api/v1/system/hook/server/update` | 更新Hook服务 | platform |
| GET | `/api/v1/system/im/call/active` | 查询群通话进行状态 | all |
| POST | `/api/v1/system/im/call/signal` | 通话信令统一入口 | all |
| POST | `/api/v1/system/im/channel/bind` | 绑定渠道端点 | all |
| POST | `/api/v1/system/im/channel/get-list` | 查询渠道绑定列表 | all |
| POST | `/api/v1/system/im/channel/set-mute` | 设置端点免打扰 | all |
| POST | `/api/v1/system/im/channel/unbind` | 解绑渠道端点 | all |
| POST | `/api/v1/system/im/chatlog/get-list` | 获取消息列表 | all |
| POST | `/api/v1/system/im/chatlog/mark-read` | 标记已读 | all |
| POST | `/api/v1/system/im/chatlog/search` | 搜索聊天记录 | all |
| POST | `/api/v1/system/im/message/recall` | 撤回消息 | all |
| POST | `/api/v1/system/im/message/send` | 发送消息 | all |
| POST | `/api/v1/system/im/conversation/add-members` | 群聊邀请成员 | all |
| POST | `/api/v1/system/im/conversation/create` | 创建会话 | all |
| POST | `/api/v1/system/im/conversation/create-ai` | 创建 AI 会话 | all |
| POST | `/api/v1/system/im/conversation/disband` | 群聊解散 | all |
| POST | `/api/v1/system/im/conversation/get` | 获取会话详情 | all |
| POST | `/api/v1/system/im/conversation/get-by-participants` | 按参与者查询单聊会话 | all |
| POST | `/api/v1/system/im/conversation/get-list` | 获取会话列表 | all |
| POST | `/api/v1/system/im/conversation/get-members` | 获取会话成员列表 | all |
| POST | `/api/v1/system/im/conversation/mute-member` | 群聊禁言/取消禁言成员 | all |
| POST | `/api/v1/system/im/conversation/remove-member` | 群聊移除/退出成员 | all |
| POST | `/api/v1/system/im/conversation/set-admin` | 群聊设置/取消管理员 | all |
| POST | `/api/v1/system/im/conversation/transfer-owner` | 群聊转让群主 | all |
| POST | `/api/v1/system/im/conversation/update-setting` | 更新会话设置 | all |
| POST | `/api/v1/system/im/friend/apply` | 发起好友申请 | all |
| POST | `/api/v1/system/im/friend/delete` | 删除好友 | all |
| POST | `/api/v1/system/im/friend/get-list` | 获取好友列表 | all |
| POST | `/api/v1/system/im/friend/handle` | 处理好友申请 | all |
| POST | `/api/v1/system/im/friend/request/get-list` | 获取好友申请列表 | all |
| POST | `/api/v1/system/im/moment/comment/create` | 发表评论 | all |
| POST | `/api/v1/system/im/moment/comment/delete` | 删除评论 | all |
| POST | `/api/v1/system/im/moment/create` | 发布动态 | all |
| POST | `/api/v1/system/im/moment/delete` | 删除动态 | all |
| POST | `/api/v1/system/im/moment/get` | 获取动态详情 | all |
| POST | `/api/v1/system/im/moment/like` | 点赞动态 | all |
| POST | `/api/v1/system/im/moment/list` | 查询动态列表 | all |
| POST | `/api/v1/system/im/moment/unlike` | 取消点赞动态 | all |
| POST | `/api/v1/system/job/task/cancel` | 取消执行任务 | platform |
| POST | `/api/v1/system/job/task/group/create` | 新建任务分组 | platform |
| POST | `/api/v1/system/job/task/group/delete` | 删除任务分组 | platform |
| POST | `/api/v1/system/job/task/group/get-list` | 获取任务分组列表 | platform |
| POST | `/api/v1/system/job/task/group/get-one` | 获取任务分组详情 | platform |
| POST | `/api/v1/system/job/task/group/update` | 更新任务分组 | platform |
| POST | `/api/v1/system/job/task/info/create` | 创建任务 | platform |
| POST | `/api/v1/system/job/task/info/delete` | 删除任务 | platform |
| POST | `/api/v1/system/job/task/info/get-list` | 获取任务列表 | platform |
| POST | `/api/v1/system/job/task/info/get-one` | 获取任务详情 | platform |
| POST | `/api/v1/system/job/task/info/start` | 启动任务 | platform |
| POST | `/api/v1/system/job/task/info/stop` | 停止任务 | platform |
| POST | `/api/v1/system/job/task/info/update` | 更新任务 | platform |
| POST | `/api/v1/system/job/task/send` | 发送延时请求 | platform |
| POST | `/api/v1/system/mall/license/batch-create` | 批量创建授权码 | platform |
| POST | `/api/v1/system/mall/license/create` | 创建授权码 | platform |
| POST | `/api/v1/system/mall/license/get-list` | 授权码列表 | admin |
| POST | `/api/v1/system/mall/license/get-one` | 授权码详情 | admin |
| POST | `/api/v1/system/mall/license/revoke` | 撤销授权码 | platform |
| POST | `/api/v1/system/mall/package/create` | 创建套餐 | platform |
| POST | `/api/v1/system/mall/package/delete` | 删除套餐 | platform |
| POST | `/api/v1/system/mall/package/get-list` | 获取套餐列表 | admin |
| POST | `/api/v1/system/mall/package/get-one` | 获取套餐详情 | admin |
| POST | `/api/v1/system/mall/package/update` | 更新套餐 | platform |
| POST | `/api/v1/system/mall/product/create` | 创建商品 | platform |
| POST | `/api/v1/system/mall/product/delete` | 删除商品 | platform |
| POST | `/api/v1/system/mall/product/get-list` | 获取商品列表 | admin |
| POST | `/api/v1/system/mall/product/get-one` | 获取商品详情 | admin |
| POST | `/api/v1/system/mall/product/update` | 更新商品 | platform |
| POST | `/api/v1/system/mcp/run` | Stateless MCP HTTP | all |
| POST | `/api/v1/system/ops/feedback/create` | 添加帮助与反馈 | all |
| POST | `/api/v1/system/ops/feedback/get-list` | 获取帮助与反馈 | all |
| POST | `/api/v1/system/ops/feedback/update` | 更新帮助与反馈 | all |
| POST | `/api/v1/system/ops/work-order/create` | 添加工单 | all |
| POST | `/api/v1/system/ops/work-order/get-list` | 获取工单列表 | all |
| POST | `/api/v1/system/ops/work-order/update` | 更新工单 | all |
| POST | `/api/v1/system/resource/api/create` | 添加接口 | platform |
| POST | `/api/v1/system/resource/api/delete` | 删除接口 | platform |
| POST | `/api/v1/system/resource/api/get-list` | 获取接口列表 | platform |
| POST | `/api/v1/system/resource/api/update` | 更新接口 | platform |

## 端点详情

### POST `/api/v1/system/agreement/create`

**说明**: 创建协议

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `content` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `remark` | string | 否 |  |
| `status` | integer | 否 | 格式: int64 |
| `title` | string | 否 |  |
| `updatedTime` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
  "content": "string",
  "createdTime": "string",
  "id": "string",
  "name": "string",
  "remark": "string",
  "status": 1,
  "title": "string",
  "updatedTime": "string"
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
ur api /api/v1/system/agreement/create \
  --body '{"code": "string", "content": "string", "createdTime": "string", "id": "string", "name": "string", "remark": "string", "status": 1, "title": "string", "updatedTime": "string"}'
```

### POST `/api/v1/system/agreement/delete`

**说明**: 删除协议

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/agreement/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/agreement/get-list`

**说明**: 获取协议列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  |
| `code` | string | 否 |  |
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int64 |
| `title` | string | 否 |  |

**请求示例**:
```json
{
  "appID": "string",
  "code": "string",
  "name": "string",
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
  "status": 1,
  "title": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": "string",
        "content": "string",
        "createdTime": "string",
        "id": "string",
        "name": "string",
        "remark": "string",
        "status": 1,
        "title": "string",
        "updatedTime": "string"
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
ur api /api/v1/system/agreement/get-list \
  --body '{"appID": "string", "code": "string", "name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "title": "string"}'
```

### POST `/api/v1/system/agreement/get-one`

**说明**: 获取协议详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
  "data": {
    "code": "string",
    "content": "string",
    "createdTime": "string",
    "id": "string",
    "name": "string",
    "remark": "string",
    "status": 1,
    "title": "string",
    "updatedTime": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/agreement/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/agreement/preview-render`

**说明**: 预览渲染协议

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  |
| `content` | string | 否 |  |
| `id` | string | 否 |  |
| `tenantCode` | string | 否 |  |

**请求示例**:
```json
{
  "appID": "string",
  "content": "string",
  "id": "string",
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "content": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/agreement/preview-render \
  --body '{"appID": "string", "content": "string", "id": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/agreement/update`

**说明**: 更新协议

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `content` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `remark` | string | 否 |  |
| `status` | integer | 否 | 格式: int64 |
| `title` | string | 否 |  |
| `updatedTime` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
  "content": "string",
  "createdTime": "string",
  "id": "string",
  "name": "string",
  "remark": "string",
  "status": 1,
  "title": "string",
  "updatedTime": "string"
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
ur api /api/v1/system/agreement/update \
  --body '{"code": "string", "content": "string", "createdTime": "string", "id": "string", "name": "string", "remark": "string", "status": 1, "title": "string", "updatedTime": "string"}'
```

### POST `/api/v1/system/app/agreement/bind-batch-update`

**说明**: 批量更新应用协议绑定

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agreementID` | string | 否 |  |
| `appID` | string | 否 |  |
| `list` | array[AppAgreementBind] | 是 |  |

**请求示例**:
```json
{
  "agreementID": "string",
  "appID": "string",
  "list": [
    {
      "agreement": {
        "code": "string",
        "content": "string",
        "createdTime": "string",
        "id": "string",
        "name": "string",
        "remark": "string",
        "status": 1,
        "title": "string",
        "updatedTime": "string"
      },
      "agreementID": "string",
      "app": {
        "agreementsMap": {},
        "appleConfig": {
          "appID": "string",
          "bundleID": "string",
          "keyID": "string",
          "privateKey": "string",
          "redirectURI": "string",
          "teamID": "string"
        },
        "avatar": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "dingConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "githubConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "googleConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "id": "string",
        "isUseMenu": 1,
        "isUseProxy": 1,
        "loginTypes": [
          "string"
        ],
        "menuRole": "string",
        "name": "示例名称",
        "nativeConfig": {
          "filePath": "string",
          "version": "string",
          "versionDesc": "string"
        },
        "proxy": "string",
        "sort": 1,
        "status": 1,
        "subType": "string",
        "tenantCode": "string",
        "tenantName": "示例名称",
        "trialTime": 1,
        "type": "string",
        "url": "string",
        "useBy": "string",
        "wxMiniConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "wxOpenConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        }
      },
      "appID": "string",
      "id": "string",
      "sort": 1,
      "status": 1,
      "type": "string"
    }
  ]
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
ur api /api/v1/system/app/agreement/bind-batch-update \
  --body '{"agreementID": "string", "appID": "string", "list": [{"agreement": {"code": "string", "content": "string", "createdTime": "string", "id": "string", "name": "string", "remark": "string", "status": 1, "title": "string", "updatedTime": "string"}, "agreementID": "string", "app": {"agreementsMap": {}, "appleConfig": {"appID": "string", "bundleID": "string", "keyID": "string", "privateKey": "string", "redirectURI": "string", "teamID": "string"}, "avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "dingConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "githubConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "googleConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "id": "string", "isUseMenu": 1, "isUseProxy": 1, "loginTypes": ["string"], "menuRole": "string", "name": "示例名称", "nativeConfig": {"filePath": "string", "version": "string", "versionDesc": "string"}, "proxy": "string", "sort": 1, "status": 1, "subType": "string", "tenantCode": "string", "tenantName": "示例名称", "trialTime": 1, "type": "string", "url": "string", "useBy": "string", "wxMiniConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "wxOpenConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}}, "appID": "string", "id": "string", "sort": 1, "status": 1, "type": "string"}]}'
```

### POST `/api/v1/system/app/agreement/get-bind-list`

**说明**: 获取应用协议绑定列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agreementID` | string | 否 |  |
| `appID` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int64 |
| `type` | string | 否 |  |
| `withAgreement` | boolean | 否 | 格式: boolean |
| `withApp` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "agreementID": "string",
  "appID": "string",
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
  "status": 1,
  "type": "string",
  "withAgreement": true,
  "withApp": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "agreement": {
          "code": "string",
          "content": "string",
          "createdTime": "string",
          "id": "string",
          "name": "string",
          "remark": "string",
          "status": 1,
          "title": "string",
          "updatedTime": "string"
        },
        "agreementID": "string",
        "app": {
          "agreementsMap": {},
          "appleConfig": {
            "appID": "string",
            "bundleID": "string",
            "keyID": "string",
            "privateKey": "string",
            "redirectURI": "string",
            "teamID": "string"
          },
          "avatar": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "dingConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          },
          "githubConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          },
          "googleConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          },
          "id": "string",
          "isUseMenu": 1,
          "isUseProxy": 1,
          "loginTypes": [
            "string"
          ],
          "menuRole": "string",
          "name": "示例名称",
          "nativeConfig": {
            "filePath": "string",
            "version": "string",
            "versionDesc": "string"
          },
          "proxy": "string",
          "sort": 1,
          "status": 1,
          "subType": "string",
          "tenantCode": "string",
          "tenantName": "示例名称",
          "trialTime": 1,
          "type": "string",
          "url": "string",
          "useBy": "string",
          "wxMiniConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          },
          "wxOpenConfig": {
            "appID": "string",
            "appKey": "string",
            "appSecret": "string"
          }
        },
        "appID": "string",
        "id": "string",
        "sort": 1,
        "status": 1,
        "type": "string"
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
ur api /api/v1/system/app/agreement/get-bind-list \
  --body '{"agreementID": "string", "appID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "type": "string", "withAgreement": true, "withApp": true}'
```

### POST `/api/v1/system/app/core/get-one`

**说明**: 无需登录获取应用信息

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  |
| `withAgreementCodes` | array[string] | 否 |  |

**请求示例**:
```json
{
  "id": "string",
  "withAgreementCodes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "agreementsMap": {},
    "appleConfig": {
      "appID": "string",
      "bundleID": "string",
      "keyID": "string",
      "privateKey": "string",
      "redirectURI": "string",
      "teamID": "string"
    },
    "avatar": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "dingConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "githubConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "googleConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "id": "string",
    "isUseMenu": 1,
    "isUseProxy": 1,
    "loginTypes": [
      "string"
    ],
    "menuRole": "string",
    "name": "示例名称",
    "nativeConfig": {
      "filePath": "string",
      "version": "string",
      "versionDesc": "string"
    },
    "proxy": "string",
    "sort": 1,
    "status": 1,
    "subType": "string",
    "tenantCode": "string",
    "tenantName": "示例名称",
    "trialTime": 1,
    "type": "string",
    "url": "string",
    "useBy": "string",
    "wxMiniConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "wxOpenConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/app/core/get-one \
  --body '{"id": "string", "withAgreementCodes": ["string"]}'
```

### POST `/api/v1/system/app/info/create`

**说明**: 添加应用

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agreementsMap` | object | 否 |  按请求协议 code 回填的协议映射 |
| `appleConfig` | object | 否 |  |
| `appleConfig.appID` | string | 否 |  |
| `appleConfig.bundleID` | string | 否 |  |
| `appleConfig.keyID` | string | 否 |  |
| `appleConfig.privateKey` | string | 否 |  |
| `appleConfig.redirectURI` | string | 否 |  |
| `appleConfig.teamID` | string | 否 |  |
| `avatar` | string | 否 |  头像 更新的时候传filePath即可 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  应用描述 |
| `dingConfig` | object | 否 |  |
| `dingConfig.appID` | string | 否 |  |
| `dingConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `dingConfig.appSecret` | string | 否 |  |
| `githubConfig` | object | 否 |  |
| `githubConfig.appID` | string | 否 |  |
| `githubConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `githubConfig.appSecret` | string | 否 |  |
| `googleConfig` | object | 否 |  |
| `googleConfig.appID` | string | 否 |  |
| `googleConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `googleConfig.appSecret` | string | 否 |  |
| `id` | string | 否 |  编号,创建时自动生成(默认),也可以指定 |
| `isUseMenu` | integer | 否 |  是否使用菜单 1-是 2-否 (格式: int64) |
| `isUseProxy` | integer | 否 |  是否使用代理 1-是 2-否 (格式: int64) |
| `loginTypes` | array[string] | 否 | 选项取决于 app里是否配置了 支持的登录类型(不填支持全部登录方式):    "email":邮箱(系统配置里配置了才能开启) "phone":手机号(系统配置里配置了才能开启)  "wxMiniP":微信小程序(app里配置了微信小程序才能开启)  "wxOpen": 微信开放平台登录(app里配置了微信开放才能开启)   "ding":钉钉应用(包含小程序,h5等方式,app里配置了钉钉才能开启)  "pwd":账号密码注册 |
| `menuRole` | string | 否 |  菜单角色类型 all-支持两类(菜单需选role) platform-平台管理员 tenant-租户管理员 |
| `name` | string | 否 |  应用名称 |
| `nativeConfig` | object | 否 |  |
| `nativeConfig.filePath` | string | 否 |  文件路径,拿来下载文件 |
| `nativeConfig.version` | string | 否 |  应用版本 |
| `nativeConfig.versionDesc` | string | 否 | 版本说明 |
| `proxy` | string | 否 |  代理配置(JSON字符串) |
| `sort` | integer | 否 | 排序,不能小于等于0 (格式: int64) |
| `status` | integer | 否 |  应用状态 1:上架, 2:下架(默认) (格式: int64) |
| `subType` | string | 否 |  子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `tenantCode` | string | 否 |  租户编码,common的是通用应用,非common的是租户应用 |
| `tenantName` | string | 否 |  租户名称 |
| `trialTime` | integer | 否 | 试用时间(单位为天,为0不限制,如果为0,是免费使用的,如果为-1 则不允许自己创建) (格式: int64) |
| `type` | string | 否 |  应用类型 web:web页面  native:原生  mini:小程序 |
| `url` | string | 否 |  应用地址 |
| `useBy` | string | 否 |  使用对象 client:登录直接授权该应用 admin:管理员使用,需要邀请 |
| `wxMiniConfig` | object | 否 |  |
| `wxMiniConfig.appID` | string | 否 |  |
| `wxMiniConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `wxMiniConfig.appSecret` | string | 否 |  |
| `wxOpenConfig` | object | 否 |  |
| `wxOpenConfig.appID` | string | 否 |  |
| `wxOpenConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `wxOpenConfig.appSecret` | string | 否 |  |

**请求示例**:
```json
{
  "agreementsMap": {},
  "appleConfig": {
    "appID": "string",
    "bundleID": "string",
    "keyID": "string",
    "privateKey": "string",
    "redirectURI": "string",
    "teamID": "string"
  },
  "avatar": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "dingConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "githubConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "googleConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "id": "string",
  "isUseMenu": 1,
  "isUseProxy": 1,
  "loginTypes": [
    "string"
  ],
  "menuRole": "string",
  "name": "示例名称",
  "nativeConfig": {
    "filePath": "string",
    "version": "string",
    "versionDesc": "string"
  },
  "proxy": "string",
  "sort": 1,
  "status": 1,
  "subType": "string",
  "tenantCode": "string",
  "tenantName": "示例名称",
  "trialTime": 1,
  "type": "string",
  "url": "string",
  "useBy": "string",
  "wxMiniConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "wxOpenConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
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
ur api /api/v1/system/app/info/create \
  --body '{"agreementsMap": {}, "appleConfig": {"appID": "string", "bundleID": "string", "keyID": "string", "privateKey": "string", "redirectURI": "string", "teamID": "string"}, "avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "dingConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "githubConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "googleConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "id": "string", "isUseMenu": 1, "isUseProxy": 1, "loginTypes": ["string"], "menuRole": "string", "name": "示例名称", "nativeConfig": {"filePath": "string", "version": "string", "versionDesc": "string"}, "proxy": "string", "sort": 1, "status": 1, "subType": "string", "tenantCode": "string", "tenantName": "示例名称", "trialTime": 1, "type": "string", "url": "string", "useBy": "string", "wxMiniConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "wxOpenConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}}'
```

### POST `/api/v1/system/app/info/delete`

**说明**: 删除应用

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/app/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/app/info/get-list`

**说明**: 获取应用列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[string] | 否 |  应用ID列表 |
| `name` | string | 否 |  应用名称 |
| `notID` | string | 否 |  排除某个应用ID |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 应用状态 1 上架 2 下架 (格式: int64) |
| `subType` | string | 否 |  子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `tenantIsBind` | integer | 否 |  过滤租户是否已绑定的应用 1-是 2-否,如果要过滤指定租户,通过http头里指定 (格式: int64) |
| `type` | string | 否 |  应用类型 web:web页面  native:原生  mini:小程序 |
| `useBy` | string | 否 |  使用对象 client:登录直接授权该应用 admin:管理员使用,需要邀请 |
| `withAgreementCodes` | array[string] | 否 |  非空时按 code 回填已绑定协议 |

**请求示例**:
```json
{
  "ids": [
    "string"
  ],
  "name": "示例名称",
  "notID": "string",
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
  "status": 1,
  "subType": "string",
  "tenantIsBind": 1,
  "type": "string",
  "useBy": "string",
  "withAgreementCodes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "agreementsMap": {},
        "appleConfig": {
          "appID": "string",
          "bundleID": "string",
          "keyID": "string",
          "privateKey": "string",
          "redirectURI": "string",
          "teamID": "string"
        },
        "avatar": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "dingConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "githubConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "googleConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "id": "string",
        "isUseMenu": 1,
        "isUseProxy": 1,
        "loginTypes": [
          "string"
        ],
        "menuRole": "string",
        "name": "示例名称",
        "nativeConfig": {
          "filePath": "string",
          "version": "string",
          "versionDesc": "string"
        },
        "proxy": "string",
        "sort": 1,
        "status": 1,
        "subType": "string",
        "tenantCode": "string",
        "tenantName": "示例名称",
        "trialTime": 1,
        "type": "string",
        "url": "string",
        "useBy": "string",
        "wxMiniConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "wxOpenConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
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
ur api /api/v1/system/app/info/get-list \
  --body '{"ids": ["string"], "name": "示例名称", "notID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "subType": "string", "tenantIsBind": 1, "type": "string", "useBy": "string", "withAgreementCodes": ["string"]}'
```

### POST `/api/v1/system/app/info/get-one`

**说明**: 获取应用详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  |
| `withAgreementCodes` | array[string] | 否 |  |

**请求示例**:
```json
{
  "id": "string",
  "withAgreementCodes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "agreementsMap": {},
    "appleConfig": {
      "appID": "string",
      "bundleID": "string",
      "keyID": "string",
      "privateKey": "string",
      "redirectURI": "string",
      "teamID": "string"
    },
    "avatar": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "dingConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "githubConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "googleConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "id": "string",
    "isUseMenu": 1,
    "isUseProxy": 1,
    "loginTypes": [
      "string"
    ],
    "menuRole": "string",
    "name": "示例名称",
    "nativeConfig": {
      "filePath": "string",
      "version": "string",
      "versionDesc": "string"
    },
    "proxy": "string",
    "sort": 1,
    "status": 1,
    "subType": "string",
    "tenantCode": "string",
    "tenantName": "示例名称",
    "trialTime": 1,
    "type": "string",
    "url": "string",
    "useBy": "string",
    "wxMiniConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "wxOpenConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/app/info/get-one \
  --body '{"id": "string", "withAgreementCodes": ["string"]}'
```

### POST `/api/v1/system/app/info/update`

**说明**: 更新应用

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `agreementsMap` | object | 否 |  按请求协议 code 回填的协议映射 |
| `appleConfig` | object | 否 |  |
| `appleConfig.appID` | string | 否 |  |
| `appleConfig.bundleID` | string | 否 |  |
| `appleConfig.keyID` | string | 否 |  |
| `appleConfig.privateKey` | string | 否 |  |
| `appleConfig.redirectURI` | string | 否 |  |
| `appleConfig.teamID` | string | 否 |  |
| `avatar` | string | 否 |  头像 更新的时候传filePath即可 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  应用描述 |
| `dingConfig` | object | 否 |  |
| `dingConfig.appID` | string | 否 |  |
| `dingConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `dingConfig.appSecret` | string | 否 |  |
| `githubConfig` | object | 否 |  |
| `githubConfig.appID` | string | 否 |  |
| `githubConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `githubConfig.appSecret` | string | 否 |  |
| `googleConfig` | object | 否 |  |
| `googleConfig.appID` | string | 否 |  |
| `googleConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `googleConfig.appSecret` | string | 否 |  |
| `id` | string | 否 |  编号,创建时自动生成(默认),也可以指定 |
| `isUseMenu` | integer | 否 |  是否使用菜单 1-是 2-否 (格式: int64) |
| `isUseProxy` | integer | 否 |  是否使用代理 1-是 2-否 (格式: int64) |
| `loginTypes` | array[string] | 否 | 选项取决于 app里是否配置了 支持的登录类型(不填支持全部登录方式):    "email":邮箱(系统配置里配置了才能开启) "phone":手机号(系统配置里配置了才能开启)  "wxMiniP":微信小程序(app里配置了微信小程序才能开启)  "wxOpen": 微信开放平台登录(app里配置了微信开放才能开启)   "ding":钉钉应用(包含小程序,h5等方式,app里配置了钉钉才能开启)  "pwd":账号密码注册 |
| `menuRole` | string | 否 |  菜单角色类型 all-支持两类(菜单需选role) platform-平台管理员 tenant-租户管理员 |
| `name` | string | 否 |  应用名称 |
| `nativeConfig` | object | 否 |  |
| `nativeConfig.filePath` | string | 否 |  文件路径,拿来下载文件 |
| `nativeConfig.version` | string | 否 |  应用版本 |
| `nativeConfig.versionDesc` | string | 否 | 版本说明 |
| `proxy` | string | 否 |  代理配置(JSON字符串) |
| `sort` | integer | 否 | 排序,不能小于等于0 (格式: int64) |
| `status` | integer | 否 |  应用状态 1:上架, 2:下架(默认) (格式: int64) |
| `subType` | string | 否 |  子类型 mini:(wx:微信小程序,ding:钉钉小程序)  native:安卓,ios,鸿蒙,win |
| `tenantCode` | string | 否 |  租户编码,common的是通用应用,非common的是租户应用 |
| `tenantName` | string | 否 |  租户名称 |
| `trialTime` | integer | 否 | 试用时间(单位为天,为0不限制,如果为0,是免费使用的,如果为-1 则不允许自己创建) (格式: int64) |
| `type` | string | 否 |  应用类型 web:web页面  native:原生  mini:小程序 |
| `url` | string | 否 |  应用地址 |
| `useBy` | string | 否 |  使用对象 client:登录直接授权该应用 admin:管理员使用,需要邀请 |
| `wxMiniConfig` | object | 否 |  |
| `wxMiniConfig.appID` | string | 否 |  |
| `wxMiniConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `wxMiniConfig.appSecret` | string | 否 |  |
| `wxOpenConfig` | object | 否 |  |
| `wxOpenConfig.appID` | string | 否 |  |
| `wxOpenConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `wxOpenConfig.appSecret` | string | 否 |  |

**请求示例**:
```json
{
  "agreementsMap": {},
  "appleConfig": {
    "appID": "string",
    "bundleID": "string",
    "keyID": "string",
    "privateKey": "string",
    "redirectURI": "string",
    "teamID": "string"
  },
  "avatar": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "dingConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "githubConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "googleConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "id": "string",
  "isUseMenu": 1,
  "isUseProxy": 1,
  "loginTypes": [
    "string"
  ],
  "menuRole": "string",
  "name": "示例名称",
  "nativeConfig": {
    "filePath": "string",
    "version": "string",
    "versionDesc": "string"
  },
  "proxy": "string",
  "sort": 1,
  "status": 1,
  "subType": "string",
  "tenantCode": "string",
  "tenantName": "示例名称",
  "trialTime": 1,
  "type": "string",
  "url": "string",
  "useBy": "string",
  "wxMiniConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "wxOpenConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  }
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
ur api /api/v1/system/app/info/update \
  --body '{"agreementsMap": {}, "appleConfig": {"appID": "string", "bundleID": "string", "keyID": "string", "privateKey": "string", "redirectURI": "string", "teamID": "string"}, "avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "dingConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "githubConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "googleConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "id": "string", "isUseMenu": 1, "isUseProxy": 1, "loginTypes": ["string"], "menuRole": "string", "name": "示例名称", "nativeConfig": {"filePath": "string", "version": "string", "versionDesc": "string"}, "proxy": "string", "sort": 1, "status": 1, "subType": "string", "tenantCode": "string", "tenantName": "示例名称", "trialTime": 1, "type": "string", "url": "string", "useBy": "string", "wxMiniConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "wxOpenConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}}'
```

### POST `/api/v1/system/app/menu/create`

**说明**: 添加菜单

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  模块编码 |
| `body` | string | 否 |  前端自定义字段 |
| `children` | array[MenuInfo] | 否 |  子菜单列表 |
| `component` | string | 否 |  页面 |
| `createdTime` | string | 否 |  创建时间 |
| `defaultIcon` | string | 否 |  默认图标 |
| `defaultIsHideInMenu` | integer | 否 |  默认是否隐藏菜单 1-是 2-否 (格式: int64) |
| `defaultName` | string | 否 |  默认 菜单名称 |
| `defaultSort` | integer | 否 |  默认左侧table排序序号 (格式: int64) |
| `id` | string | 否 |  编号 |
| `idPath` | string | 否 |  菜单ID路径 |
| `parentID` | string | 否 |  父菜单ID，一级菜单为1 |
| `path` | string | 否 |  系统的path |
| `redirect` | string | 否 |  路由重定向 |
| `role` | string | 否 |  角色类型 仅App.menuRole=all时有意义 all:两类都可以看 platform:平台管理员 tenant:租户管理员 |
| `tenantCode` | string | 否 |  跟着应用走, common的是通用应用,非common的是租户应用 |
| `type` | integer | 否 |  类型 1：菜单或者页面 2：iframe嵌入 3：外链跳转 (格式: int64) |

**请求示例**:
```json
{
  "appID": "string",
  "body": "string",
  "children": [
    {
      "appID": "string",
      "body": "string",
      "children": [
        {
          "appID": "string",
          "body": "string",
          "children": [],
          "component": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "defaultIcon": "string",
          "defaultIsHideInMenu": 1,
          "defaultName": "示例名称",
          "defaultSort": 1,
          "id": "string",
          "idPath": "string",
          "parentID": "string",
          "path": "string",
          "redirect": "string",
          "role": "string",
          "tenantCode": "string",
          "type": 1
        }
      ],
      "component": "string",
      "createdTime": "2026-01-01T00:00:00Z",
      "defaultIcon": "string",
      "defaultIsHideInMenu": 1,
      "defaultName": "示例名称",
      "defaultSort": 1,
      "id": "string",
      "idPath": "string",
      "parentID": "string",
      "path": "string",
      "redirect": "string",
      "role": "string",
      "tenantCode": "string",
      "type": 1
    }
  ],
  "component": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "defaultIcon": "string",
  "defaultIsHideInMenu": 1,
  "defaultName": "示例名称",
  "defaultSort": 1,
  "id": "string",
  "idPath": "string",
  "parentID": "string",
  "path": "string",
  "redirect": "string",
  "role": "string",
  "tenantCode": "string",
  "type": 1
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
ur api /api/v1/system/app/menu/create \
  --body '{"appID": "string", "body": "string", "children": [{"appID": "string", "body": "string", "children": [{"appID": "string", "body": "string", "children": [], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}'
```

### POST `/api/v1/system/app/menu/delete`

**说明**: 删除菜单

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/app/menu/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/app/menu/get-list`

**说明**: 获取菜单列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  应用编号 |
| `isRetTree` | boolean | 否 |  是否返回树形结构 (格式: boolean) |
| `isStar` | integer | 否 |  是否收藏菜单 1-是 2-否 (格式: int64) |
| `path` | string | 否 |  按菜单路径筛选 |

**请求示例**:
```json
{
  "appID": "string",
  "isRetTree": true,
  "isStar": 1,
  "path": "string"
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
        "body": "string",
        "children": [
          {
            "appID": "string",
            "body": "string",
            "children": [],
            "component": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "defaultIcon": "string",
            "defaultIsHideInMenu": 1,
            "defaultName": "示例名称",
            "defaultSort": 1,
            "id": "string",
            "idPath": "string",
            "parentID": "string",
            "path": "string",
            "redirect": "string",
            "role": "string",
            "tenantCode": "string",
            "type": 1
          }
        ],
        "component": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "defaultIcon": "string",
        "defaultIsHideInMenu": 1,
        "defaultName": "示例名称",
        "defaultSort": 1,
        "id": "string",
        "idPath": "string",
        "parentID": "string",
        "path": "string",
        "redirect": "string",
        "role": "string",
        "tenantCode": "string",
        "type": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/app/menu/get-list \
  --body '{"appID": "string", "isRetTree": true, "isStar": 1, "path": "string"}'
```

### POST `/api/v1/system/app/menu/update`

**说明**: 更新菜单

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  模块编码 |
| `body` | string | 否 |  前端自定义字段 |
| `children` | array[MenuInfo] | 否 |  子菜单列表 |
| `component` | string | 否 |  页面 |
| `createdTime` | string | 否 |  创建时间 |
| `defaultIcon` | string | 否 |  默认图标 |
| `defaultIsHideInMenu` | integer | 否 |  默认是否隐藏菜单 1-是 2-否 (格式: int64) |
| `defaultName` | string | 否 |  默认 菜单名称 |
| `defaultSort` | integer | 否 |  默认左侧table排序序号 (格式: int64) |
| `id` | string | 否 |  编号 |
| `idPath` | string | 否 |  菜单ID路径 |
| `parentID` | string | 否 |  父菜单ID，一级菜单为1 |
| `path` | string | 否 |  系统的path |
| `redirect` | string | 否 |  路由重定向 |
| `role` | string | 否 |  角色类型 仅App.menuRole=all时有意义 all:两类都可以看 platform:平台管理员 tenant:租户管理员 |
| `tenantCode` | string | 否 |  跟着应用走, common的是通用应用,非common的是租户应用 |
| `type` | integer | 否 |  类型 1：菜单或者页面 2：iframe嵌入 3：外链跳转 (格式: int64) |

**请求示例**:
```json
{
  "appID": "string",
  "body": "string",
  "children": [
    {
      "appID": "string",
      "body": "string",
      "children": [
        {
          "appID": "string",
          "body": "string",
          "children": [],
          "component": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "defaultIcon": "string",
          "defaultIsHideInMenu": 1,
          "defaultName": "示例名称",
          "defaultSort": 1,
          "id": "string",
          "idPath": "string",
          "parentID": "string",
          "path": "string",
          "redirect": "string",
          "role": "string",
          "tenantCode": "string",
          "type": 1
        }
      ],
      "component": "string",
      "createdTime": "2026-01-01T00:00:00Z",
      "defaultIcon": "string",
      "defaultIsHideInMenu": 1,
      "defaultName": "示例名称",
      "defaultSort": 1,
      "id": "string",
      "idPath": "string",
      "parentID": "string",
      "path": "string",
      "redirect": "string",
      "role": "string",
      "tenantCode": "string",
      "type": 1
    }
  ],
  "component": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "defaultIcon": "string",
  "defaultIsHideInMenu": 1,
  "defaultName": "示例名称",
  "defaultSort": 1,
  "id": "string",
  "idPath": "string",
  "parentID": "string",
  "path": "string",
  "redirect": "string",
  "role": "string",
  "tenantCode": "string",
  "type": 1
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
ur api /api/v1/system/app/menu/update \
  --body '{"appID": "string", "body": "string", "children": [{"appID": "string", "body": "string", "children": [{"appID": "string", "body": "string", "children": [], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}'
```

### POST `/api/v1/system/check-in/do`

**说明**: 用户签到

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "alreadyDone": true,
    "pointDelta": "string",
    "totalBalance": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/check-in/do \
  --body '{}'
```

### POST `/api/v1/system/check-in/get-list`

**说明**: 签到记录列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `endTime` | string | 否 |  结束时间(Unix毫秒) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `startTime` | string | 否 |  开始时间(Unix毫秒) |
| `userID` | string | 否 |  管理员可指定用户ID |

**请求示例**:
```json
{
  "endTime": "2026-01-01T00:00:00Z",
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
  "startTime": "2026-01-01T00:00:00Z",
  "userID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "checkDate": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "isContinuous": true,
        "pointDelta": "string",
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
ur api /api/v1/system/check-in/get-list \
  --body '{"endTime": "2026-01-01T00:00:00Z", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "startTime": "2026-01-01T00:00:00Z", "userID": "string"}'
```

### POST `/api/v1/system/check-in/point-balance/get`

**说明**: 获取当前用户积分余额

**权限**: admin

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "balance": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/check-in/point-balance/get \
  --body '{}'
```

### POST `/api/v1/system/check-in/point-log/adjust`

**说明**: 管理员调整积分

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `delta` | string | 是 |  积分变动量(正增负减,不为0) |
| `remark` | string | 是 |  调整备注(必填) |
| `targetUserID` | string | 是 |  目标用户ID |

**请求示例**:
```json
{
  "delta": "string",
  "remark": "string",
  "targetUserID": "string"
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
ur api /api/v1/system/check-in/point-log/adjust \
  --body '{"delta": "string", "remark": "string", "targetUserID": "string"}'
```

### POST `/api/v1/system/check-in/point-log/get-list`

**说明**: 积分流水列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `endTime` | string | 否 |  结束时间(Unix毫秒) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `source` | string | 否 |  来源过滤: checkIn/admin |
| `startTime` | string | 否 |  开始时间(Unix毫秒) |
| `userID` | string | 否 |  管理员可指定用户ID |

**请求示例**:
```json
{
  "endTime": "2026-01-01T00:00:00Z",
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
  "source": "string",
  "startTime": "2026-01-01T00:00:00Z",
  "userID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "balance": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "delta": "string",
        "id": "string",
        "remark": "string",
        "source": "string",
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
ur api /api/v1/system/check-in/point-log/get-list \
  --body '{"endTime": "2026-01-01T00:00:00Z", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "source": "string", "startTime": "2026-01-01T00:00:00Z", "userID": "string"}'
```

### POST `/api/v1/system/common/api/batch-agg`

**说明**: 批量聚合接口请求

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `reqs` | array[ApiBatchAggReqOne] | 是 |  |

**请求示例**:
```json
{
  "reqs": [
    {
      "body": [
        {}
      ],
      "uri": "string"
    }
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "lists": [
      [
        null
      ]
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/api/batch-agg \
  --body '{"reqs": [{"body": [{}], "uri": "string"}]}'
```

### GET `/api/v1/system/common/debug`

**说明**: 调试接口GET

**权限**: public

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "body": "string",
    "headers": {},
    "requestUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/debug \
  --body '{}'
```

### POST `/api/v1/system/common/debug`

**说明**: 调试接口POST

**权限**: public

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "body": "string",
    "headers": {},
    "requestUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/debug \
  --body '{}'
```

### GET `/api/v1/system/common/debug-tencent`

**说明**: 腾讯云调试接口

**权限**: public

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/debug-tencent \
  --body '{}'
```

### GET `/api/v1/system/common/download-file`

**说明**: 下载本地文件

**权限**: public

**路径参数**:

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| `filePath` | query | string | 是 | 文件的路径 |
| `sign` | query | string | 否 | 签名.如果是私有的需要填此参数 |

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/download-file \
  --body '{}'
```

### POST `/api/v1/system/common/init-upload-file`

**说明**: 初始化上传文件

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `file` | string | 是 | 格式: binary |

**请求示例**:
```json
{
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
ur api /api/v1/system/common/init-upload-file \
  --body '{"file": "string"}'
```

### POST `/api/v1/system/common/ntp/get-one`

**说明**: ntp时间同步

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceSendTime` | string | 是 |  |

**请求示例**:
```json
{
  "deviceSendTime": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "deviceSendTime": "string",
    "serverRecvTime": "string",
    "serverSendTime": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/ntp/get-one \
  --body '{"deviceSendTime": "string"}'
```

### POST `/api/v1/system/common/qr-code/get-one`

**说明**: 获取小程序二维码

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 | 应用ID |
| `envVersion` | string | 是 | 要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版。 |
| `page` | string | 是 | 默认是主页，页面 page，例如 pages/index/index，根路径前不要填加 /，不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面。scancode_time为系统保留参数，不允许配置 |
| `scene` | string | 是 |         	最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式） |

**请求示例**:
```json
{
  "appID": "string",
  "envVersion": "string",
  "page": "string",
  "scene": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "buffer": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/qr-code/get-one \
  --body '{"appID": "string", "envVersion": "string", "page": "string", "scene": "string"}'
```

### POST `/api/v1/system/common/sys-config/info/get-one`

**说明**: 读取系统配置信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "email": {
      "from": "string",
      "host": "string",
      "isEnable": 1,
      "isSsl": 1,
      "nickname": "string",
      "port": 1,
      "secret": "******"
    },
    "log": {
      "loginKeepDays": 1,
      "operKeepDays": 1
    },
    "map": {
      "accessKey": "string",
      "accessSecret": "string",
      "isEnable": 1,
      "type": "gaode"
    },
    "oem": {
      "avatar": "string",
      "footer": {
        "beian": "string",
        "beianLink": "string",
        "companyName": "示例名称",
        "companySiteLink": "string",
        "date": "string",
        "icp": "string",
        "icpLink": "string",
        "isEnable": 1
      },
      "loginFormDescription": "string",
      "loginFormTitle": "string",
      "loginPageDescription": "string",
      "loginPageTitle": "string",
      "sloganImage": "string",
      "sloganImageTemplate": "string",
      "sloganImageType": "template",
      "title": "string"
    },
    "sms": {
      "accessKeyID": "string",
      "accessKeySecret": "string",
      "appID": "string",
      "isEnable": 1,
      "signName": "string",
      "type": "ali"
    },
    "tenant": {
      "defaultLimit": {
        "aiTokenMonthly": 1,
        "deviceMsgPerDay": 1,
        "deviceNum": 1,
        "diskSpaceGB": 1,
        "userNum": 1
      }
    },
    "user": {
      "captchaLen": 1,
      "forceSetPassword": "1",
      "ownerUserID": "string",
      "passLevel": 1,
      "remindSetPassword": "1",
      "userOwnTenantLimit": 1
    },
    "weather": {
      "apiHost": "string",
      "apiKey": "string",
      "isEnable": 1
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/sys-config/info/get-one \
  --body '{}'
```

### POST `/api/v1/system/common/sys-config/info/update`

**说明**: 更新系统配置信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `email` | object | 否 |  |
| `email.from` | string | 否 |  发件人  你自己要发邮件的邮箱 |
| `email.host` | string | 否 |  服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议 |
| `email.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `email.isSsl` | integer | 否 |  是否SSL   是否开启SSL (格式: int64) |
| `email.nickname` | string | 否 |  昵称    发件人昵称 通常为自己的邮箱 |
| `email.port` | integer | 否 |  端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465 (格式: int64) |
| `email.secret` | string | 否 |  密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥 |
| `log` | object | 否 |  |
| `log.loginKeepDays` | integer | 是 |  登录日志保留天数 (格式: int64) |
| `log.operKeepDays` | integer | 是 |  操作日志保留天数 (格式: int64) |
| `map` | object | 否 |  |
| `map.accessKey` | string | 是 |  访问密钥 |
| `map.accessSecret` | string | 是 |  访问密钥 |
| `map.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `map.type` | string | 是 |  地图服务商模式 gaode: 高德 可选: ['gaode'] |
| `oem` | object | 否 |  |
| `oem.avatar` | string | 是 |  Logo图片地址 |
| `oem.footer` | object | 是 |  |
| `oem.footer.beian` | string | 否 |  公安备案号 |
| `oem.footer.beianLink` | string | 否 |  公安备案的跳转链接 |
| `oem.footer.companyName` | string | 是 |  公司名称 |
| `oem.footer.companySiteLink` | string | 否 |  公司网站链接 |
| `oem.footer.date` | string | 否 |  年份或日期 |
| `oem.footer.icp` | string | 否 |  ICP备案号 |
| `oem.footer.icpLink` | string | 否 |  ICP备案的跳转链接 |
| `oem.footer.isEnable` | integer | 是 |  是否启用(1:启用，2:禁用) (格式: int64) |
| `oem.loginFormDescription` | string | 否 |  登录表单描述 |
| `oem.loginFormTitle` | string | 否 |  登录表单标题 |
| `oem.loginPageDescription` | string | 否 |  登录页左侧描述 |
| `oem.loginPageTitle` | string | 否 |  登录页左侧标题 |
| `oem.sloganImage` | string | 否 |  上传的Slogan图片 |
| `oem.sloganImageTemplate` | string | 否 |  Slogan模板ID |
| `oem.sloganImageType` | string | 否 |  Slogan图片类型 可选: ['template', 'upload'] |
| `oem.title` | string | 是 |  系统标题 |
| `sms` | object | 否 |  |
| `sms.accessKeyID` | string | 否 |  AccessKey ID |
| `sms.accessKeySecret` | string | 否 |  AccessKey Secret |
| `sms.appID` | string | 否 |  应用ID 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666,可前往 [短信控制台](https:console.cloud.tencent.com/smsv2/app-manage) 查看 |
| `sms.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `sms.signName` | string | 否 |  签名 |
| `sms.type` | string | 否 |  短信类型, ali: 阿里云 tencent: 腾讯 可选: ['ali', 'tencent'] |
| `tenant` | object | 否 |  |
| `tenant.defaultLimit` | object | 是 |  |
| `tenant.defaultLimit.aiTokenMonthly` | integer | 否 |  默认 AI Token 月配额（单位：token） (格式: int64) |
| `tenant.defaultLimit.deviceMsgPerDay` | integer | 否 |  默认租户每天设备消息总量上限 (格式: int64) |
| `tenant.defaultLimit.deviceNum` | integer | 否 |  租户下的设备数量限制,0为不限制 (格式: int64) |
| `tenant.defaultLimit.diskSpaceGB` | integer | 否 |  默认磁盘空间 GB (格式: int64) |
| `tenant.defaultLimit.userNum` | integer | 否 |  租户下的用户数量限制,0为不限制 (格式: int64) |
| `user` | object | 否 |  |
| `user.captchaLen` | integer | 是 |  验证码长度 (格式: int32) |
| `user.forceSetPassword` | integer | 是 |  是否强制未设置密码的用户登录后设置密码(1:是，2:否) (格式: int64) 可选: ['1', '2'] |
| `user.ownerUserID` | string | 否 |  平台超级管理员ID 非平台超管不返回该参数,不可修改 |
| `user.passLevel` | integer | 是 |  用户密码强度级别 1:长度大于等于8位即可 2:包含数字,大写字母,小写字母,特殊字符起码两种  3:包含数字,大写字母,小写字母,特殊字符起码3种 4:包含数字,大写字母,小写字母,特殊字符   特殊字符: !@#~$%^&*()+\|_ (格式: int64) |
| `user.remindSetPassword` | integer | 是 |  是否提醒未设置密码的用户设置密码(1:是，2:否) (格式: int64) 可选: ['1', '2'] |
| `user.userOwnTenantLimit` | integer | 是 |  用户可直接创建免费租户数量，0表示不可直接创建免费租户，需提供授权码 (格式: int64) |
| `weather` | object | 否 |  |
| `weather.apiHost` | string | 是 |  天气API主机地址 |
| `weather.apiKey` | string | 是 |  天气API密钥 |
| `weather.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |

**请求示例**:
```json
{
  "email": {
    "from": "string",
    "host": "string",
    "isEnable": 1,
    "isSsl": 1,
    "nickname": "string",
    "port": 1,
    "secret": "******"
  },
  "log": {
    "loginKeepDays": 1,
    "operKeepDays": 1
  },
  "map": {
    "accessKey": "string",
    "accessSecret": "string",
    "isEnable": 1,
    "type": "gaode"
  },
  "oem": {
    "avatar": "string",
    "footer": {
      "beian": "string",
      "beianLink": "string",
      "companyName": "示例名称",
      "companySiteLink": "string",
      "date": "string",
      "icp": "string",
      "icpLink": "string",
      "isEnable": 1
    },
    "loginFormDescription": "string",
    "loginFormTitle": "string",
    "loginPageDescription": "string",
    "loginPageTitle": "string",
    "sloganImage": "string",
    "sloganImageTemplate": "string",
    "sloganImageType": "template",
    "title": "string"
  },
  "sms": {
    "accessKeyID": "string",
    "accessKeySecret": "string",
    "appID": "string",
    "isEnable": 1,
    "signName": "string",
    "type": "ali"
  },
  "tenant": {
    "defaultLimit": {
      "aiTokenMonthly": 1,
      "deviceMsgPerDay": 1,
      "deviceNum": 1,
      "diskSpaceGB": 1,
      "userNum": 1
    }
  },
  "user": {
    "captchaLen": 1,
    "forceSetPassword": "1",
    "ownerUserID": "string",
    "passLevel": 1,
    "remindSetPassword": "1",
    "userOwnTenantLimit": 1
  },
  "weather": {
    "apiHost": "string",
    "apiKey": "string",
    "isEnable": 1
  }
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
ur api /api/v1/system/common/sys-config/info/update \
  --body '{"email": {"from": "string", "host": "string", "isEnable": 1, "isSsl": 1, "nickname": "string", "port": 1, "secret": "******"}, "log": {"loginKeepDays": 1, "operKeepDays": 1}, "map": {"accessKey": "string", "accessSecret": "string", "isEnable": 1, "type": "gaode"}, "oem": {"avatar": "string", "footer": {"beian": "string", "beianLink": "string", "companyName": "示例名称", "companySiteLink": "string", "date": "string", "icp": "string", "icpLink": "string", "isEnable": 1}, "loginFormDescription": "string", "loginFormTitle": "string", "loginPageDescription": "string", "loginPageTitle": "string", "sloganImage": "string", "sloganImageTemplate": "string", "sloganImageType": "template", "title": "string"}, "sms": {"accessKeyID": "string", "accessKeySecret": "string", "appID": "string", "isEnable": 1, "signName": "string", "type": "ali"}, "tenant": {"defaultLimit": {"aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}}, "user": {"captchaLen": 1, "forceSetPassword": "1", "ownerUserID": "string", "passLevel": 1, "remindSetPassword": "1", "userOwnTenantLimit": 1}, "weather": {"apiHost": "string", "apiKey": "string", "isEnable": 1}}'
```

### POST `/api/v1/system/common/third/dept/get-list`

**说明**: 获取第三方部门列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 | 父节点 |
| `thirdConfig` | object | 是 |  |
| `thirdConfig.appID` | string | 否 |  |
| `thirdConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `thirdConfig.appSecret` | string | 否 |  |
| `thirdType` | string | 否 | 第三方的类型 ding:钉钉  wxE:企业微信 |

**请求示例**:
```json
{
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
  "parentID": "string",
  "thirdConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "thirdType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "children": [
          {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          }
        ],
        "desc": "string",
        "dingTalkID": "string",
        "id": "string",
        "idPath": "string",
        "name": "示例名称",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            }
          ],
          "desc": "string",
          "dingTalkID": "string",
          "id": "string",
          "idPath": "string",
          "name": "示例名称",
          "parent": {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "userCount": 1
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "userCount": 1
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
ur api /api/v1/system/common/third/dept/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentID": "string", "thirdConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "thirdType": "string"}'
```

### POST `/api/v1/system/common/third/dept/get-one`

**说明**: 获取第三方部门详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  编号 |
| `thirdConfig` | object | 是 |  |
| `thirdConfig.appID` | string | 否 |  |
| `thirdConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `thirdConfig.appSecret` | string | 否 |  |
| `thirdType` | string | 否 | 第三方的类型 ding:钉钉  wxE:企业微信 |
| `withChildren` | boolean | 否 | 是否返回子级 (格式: boolean) |
| `withFather` | boolean | 否 | 是否返回父级 (格式: boolean) |

**请求示例**:
```json
{
  "id": "string",
  "thirdConfig": {
    "appID": "string",
    "appKey": "string",
    "appSecret": "string"
  },
  "thirdType": "string",
  "withChildren": true,
  "withFather": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "children": [
      {
        "children": [
          {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          }
        ],
        "desc": "string",
        "dingTalkID": "string",
        "id": "string",
        "idPath": "string",
        "name": "示例名称",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            }
          ],
          "desc": "string",
          "dingTalkID": "string",
          "id": "string",
          "idPath": "string",
          "name": "示例名称",
          "parent": {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "userCount": 1
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "userCount": 1
      }
    ],
    "desc": "string",
    "dingTalkID": "string",
    "id": "string",
    "idPath": "string",
    "name": "示例名称",
    "parent": {
      "children": [
        {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            }
          ],
          "desc": "string",
          "dingTalkID": "string",
          "id": "string",
          "idPath": "string",
          "name": "示例名称",
          "parent": {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "userCount": 1
        }
      ],
      "desc": "string",
      "dingTalkID": "string",
      "id": "string",
      "idPath": "string",
      "name": "示例名称",
      "parent": {
        "children": [
          {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          }
        ],
        "desc": "string",
        "dingTalkID": "string",
        "id": "string",
        "idPath": "string",
        "name": "示例名称",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            }
          ],
          "desc": "string",
          "dingTalkID": "string",
          "id": "string",
          "idPath": "string",
          "name": "示例名称",
          "parent": {
            "children": [],
            "desc": "string",
            "dingTalkID": "string",
            "id": "string",
            "idPath": "string",
            "name": "示例名称",
            "parent": {
              "children": "...",
              "desc": "...",
              "dingTalkID": "...",
              "id": "...",
              "idPath": "...",
              "name": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "userCount": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "userCount": 1
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "userCount": 1
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "userCount": 1
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "userCount": 1
    },
    "parentID": "string",
    "sort": 1,
    "status": 1,
    "userCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/third/dept/get-one \
  --body '{"id": "string", "thirdConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "thirdType": "string", "withChildren": true, "withFather": true}'
```

### POST `/api/v1/system/common/upload-file`

**说明**: 文件直传

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `file` | string | 是 | 格式: binary |

**请求示例**:
```json
{
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
ur api /api/v1/system/common/upload-file \
  --body '{"file": "string"}'
```

### POST `/api/v1/system/common/upload-url/create`

**说明**: 获取文件上传地址

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `business` | string | 是 | 业务(如产品管理 productManage) |
| `filePath` | string | 是 | 文件路径(带文件名), 当前不支持.svg |
| `isPublic` | boolean | 否 | true时使用公开桶上传,返回永久fileUrl。用于通知富文本中的图片、封面图等需长期可访问的资源 (格式: boolean) |
| `rename` | boolean | 否 | true 文件重命名，false 不重命名(默认) (格式: boolean) |
| `scene` | string | 是 | 场景(业务定义 如产品图片 productImg) |
| `useBy` | string | 否 | 公开桶时必填。上传用途: user=用户级, tenant=租户级, platform=平台级(需supper权限) |

**请求示例**:
```json
{
  "business": "string",
  "filePath": "string",
  "isPublic": true,
  "rename": true,
  "scene": "string",
  "useBy": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "filePath": "string",
    "fileUrl": "string",
    "uploadUri": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/upload-url/create \
  --body '{"business": "string", "filePath": "string", "isPublic": true, "rename": true, "scene": "string", "useBy": "string"}'
```

### POST `/api/v1/system/common/weather/get-one`

**说明**: 获取天气情况

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `position` | object | 否 |  |
| `position.latitude` | number | 是 | 纬度 (格式: double) |
| `position.longitude` | number | 是 | 经度 (格式: double) |
| `projectID` | string | 否 |  |

**请求示例**:
```json
{
  "position": {
    "latitude": 1,
    "longitude": 1
  },
  "projectID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "air": {
      "aqi": "string",
      "category": "string",
      "co": "string",
      "level": "string",
      "no2": "string",
      "o3": "string",
      "pm10": "string",
      "pm2p5": "string",
      "primary": "string",
      "so2": "string"
    },
    "cloud": "string",
    "dew": "string",
    "feelsLike": "string",
    "humidity": "string",
    "icon": "string",
    "obsTime": "string",
    "precip": "string",
    "pressure": "string",
    "temp": "string",
    "text": "string",
    "vis": "string",
    "wind360": "string",
    "windDir": "string",
    "windScale": "string",
    "windSpeed": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/weather/get-one \
  --body '{"position": {"latitude": 1, "longitude": 1}, "projectID": "string"}'
```

### GET `/api/v1/system/common/websocket/connect`

**说明**: websocket连接

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/websocket/connect \
  --body '{}'
```

### POST `/api/v1/system/common/sys-config/core/get-one`

**说明**: 读取系统配置信息(无需登录)

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "email": {
      "from": "string",
      "host": "string",
      "isEnable": 1,
      "isSsl": 1,
      "nickname": "string",
      "port": 1,
      "secret": "******"
    },
    "log": {
      "loginKeepDays": 1,
      "operKeepDays": 1
    },
    "map": {
      "accessKey": "string",
      "accessSecret": "string",
      "isEnable": 1,
      "type": "gaode"
    },
    "oem": {
      "avatar": "string",
      "footer": {
        "beian": "string",
        "beianLink": "string",
        "companyName": "示例名称",
        "companySiteLink": "string",
        "date": "string",
        "icp": "string",
        "icpLink": "string",
        "isEnable": 1
      },
      "loginFormDescription": "string",
      "loginFormTitle": "string",
      "loginPageDescription": "string",
      "loginPageTitle": "string",
      "sloganImage": "string",
      "sloganImageTemplate": "string",
      "sloganImageType": "template",
      "title": "string"
    },
    "sms": {
      "accessKeyID": "string",
      "accessKeySecret": "string",
      "appID": "string",
      "isEnable": 1,
      "signName": "string",
      "type": "ali"
    },
    "tenant": {
      "defaultLimit": {
        "aiTokenMonthly": 1,
        "deviceMsgPerDay": 1,
        "deviceNum": 1,
        "diskSpaceGB": 1,
        "userNum": 1
      }
    },
    "user": {
      "captchaLen": 1,
      "forceSetPassword": "1",
      "ownerUserID": "string",
      "passLevel": 1,
      "remindSetPassword": "1",
      "userOwnTenantLimit": 1
    },
    "weather": {
      "apiHost": "string",
      "apiKey": "string",
      "isEnable": 1
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/common/sys-config/core/get-one \
  --body '{}'
```

### POST `/api/v1/system/common/system/init`

**说明**: 初始化系统

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `config` | object | 是 |  |
| `config.email` | object | 否 |  |
| `config.email.from` | string | 否 |  发件人  你自己要发邮件的邮箱 |
| `config.email.host` | string | 否 |  服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议 |
| `config.email.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `config.email.isSsl` | integer | 否 |  是否SSL   是否开启SSL (格式: int64) |
| `config.email.nickname` | string | 否 |  昵称    发件人昵称 通常为自己的邮箱 |
| `config.email.port` | integer | 否 |  端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465 (格式: int64) |
| `config.email.secret` | string | 否 |  密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥 |
| `config.log` | object | 否 |  |
| `config.log.loginKeepDays` | integer | 是 |  登录日志保留天数 (格式: int64) |
| `config.log.operKeepDays` | integer | 是 |  操作日志保留天数 (格式: int64) |
| `config.map` | object | 否 |  |
| `config.map.accessKey` | string | 是 |  访问密钥 |
| `config.map.accessSecret` | string | 是 |  访问密钥 |
| `config.map.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `config.map.type` | string | 是 |  地图服务商模式 gaode: 高德 可选: ['gaode'] |
| `config.oem` | object | 否 |  |
| `config.oem.avatar` | string | 是 |  Logo图片地址 |
| `config.oem.footer` | object | 是 |  |
| `config.oem.loginFormDescription` | string | 否 |  登录表单描述 |
| `config.oem.loginFormTitle` | string | 否 |  登录表单标题 |
| `config.oem.loginPageDescription` | string | 否 |  登录页左侧描述 |
| `config.oem.loginPageTitle` | string | 否 |  登录页左侧标题 |
| `config.oem.sloganImage` | string | 否 |  上传的Slogan图片 |
| `config.oem.sloganImageTemplate` | string | 否 |  Slogan模板ID |
| `config.oem.sloganImageType` | string | 否 |  Slogan图片类型 可选: ['template', 'upload'] |
| `config.oem.title` | string | 是 |  系统标题 |
| `config.sms` | object | 否 |  |
| `config.sms.accessKeyID` | string | 否 |  AccessKey ID |
| `config.sms.accessKeySecret` | string | 否 |  AccessKey Secret |
| `config.sms.appID` | string | 否 |  应用ID 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666,可前往 [短信控制台](https:console.cloud.tencent.com/smsv2/app-manage) 查看 |
| `config.sms.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `config.sms.signName` | string | 否 |  签名 |
| `config.sms.type` | string | 否 |  短信类型, ali: 阿里云 tencent: 腾讯 可选: ['ali', 'tencent'] |
| `config.tenant` | object | 否 |  |
| `config.tenant.defaultLimit` | object | 是 |  |
| `config.user` | object | 否 |  |
| `config.user.captchaLen` | integer | 是 |  验证码长度 (格式: int32) |
| `config.user.forceSetPassword` | integer | 是 |  是否强制未设置密码的用户登录后设置密码(1:是，2:否) (格式: int64) 可选: ['1', '2'] |
| `config.user.ownerUserID` | string | 否 |  平台超级管理员ID 非平台超管不返回该参数,不可修改 |
| `config.user.passLevel` | integer | 是 |  用户密码强度级别 1:长度大于等于8位即可 2:包含数字,大写字母,小写字母,特殊字符起码两种  3:包含数字,大写字母,小写字母,特殊字符起码3种 4:包含数字,大写字母,小写字母,特殊字符   特殊字符: !@#~$%^&*()+\|_ (格式: int64) |
| `config.user.remindSetPassword` | integer | 是 |  是否提醒未设置密码的用户设置密码(1:是，2:否) (格式: int64) 可选: ['1', '2'] |
| `config.user.userOwnTenantLimit` | integer | 是 |  用户可直接创建免费租户数量，0表示不可直接创建免费租户，需提供授权码 (格式: int64) |
| `config.weather` | object | 否 |  |
| `config.weather.apiHost` | string | 是 |  天气API主机地址 |
| `config.weather.apiKey` | string | 是 |  天气API密钥 |
| `config.weather.isEnable` | integer | 否 | 是否启用,1为是,2为否 (格式: int64) |
| `ownerUser` | object | 是 |  |
| `ownerUser.avatar` | string | 否 |  可选头像，创建的时候传filePath |
| `ownerUser.email` | string | 否 |  可选邮箱 |
| `ownerUser.nickName` | string | 否 |  可选昵称 |
| `ownerUser.password` | string | 是 |  密码原文 |
| `ownerUser.phone` | string | 否 |  可选手机号 |
| `ownerUser.userName` | string | 是 |  用户账号 |

**请求示例**:
```json
{
  "config": {
    "email": {
      "from": "string",
      "host": "string",
      "isEnable": 1,
      "isSsl": 1,
      "nickname": "string",
      "port": 1,
      "secret": "******"
    },
    "log": {
      "loginKeepDays": 1,
      "operKeepDays": 1
    },
    "map": {
      "accessKey": "string",
      "accessSecret": "string",
      "isEnable": 1,
      "type": "gaode"
    },
    "oem": {
      "avatar": "string",
      "footer": {
        "beian": "string",
        "beianLink": "string",
        "companyName": "示例名称",
        "companySiteLink": "string",
        "date": "string",
        "icp": "string",
        "icpLink": "string",
        "isEnable": 1
      },
      "loginFormDescription": "string",
      "loginFormTitle": "string",
      "loginPageDescription": "string",
      "loginPageTitle": "string",
      "sloganImage": "string",
      "sloganImageTemplate": "string",
      "sloganImageType": "template",
      "title": "string"
    },
    "sms": {
      "accessKeyID": "string",
      "accessKeySecret": "string",
      "appID": "string",
      "isEnable": 1,
      "signName": "string",
      "type": "ali"
    },
    "tenant": {
      "defaultLimit": {
        "aiTokenMonthly": 1,
        "deviceMsgPerDay": 1,
        "deviceNum": 1,
        "diskSpaceGB": 1,
        "userNum": 1
      }
    },
    "user": {
      "captchaLen": 1,
      "forceSetPassword": "1",
      "ownerUserID": "string",
      "passLevel": 1,
      "remindSetPassword": "1",
      "userOwnTenantLimit": 1
    },
    "weather": {
      "apiHost": "string",
      "apiKey": "string",
      "isEnable": 1
    }
  },
  "ownerUser": {
    "avatar": "string",
    "email": "string",
    "nickName": "string",
    "password": "******",
    "phone": "string",
    "userName": "string"
  }
}
```

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
ur api /api/v1/system/common/system/init \
  --body '{"config": {"email": {"from": "string", "host": "string", "isEnable": 1, "isSsl": 1, "nickname": "string", "port": 1, "secret": "******"}, "log": {"loginKeepDays": 1, "operKeepDays": 1}, "map": {"accessKey": "string", "accessSecret": "string", "isEnable": 1, "type": "gaode"}, "oem": {"avatar": "string", "footer": {"beian": "string", "beianLink": "string", "companyName": "示例名称", "companySiteLink": "string", "date": "string", "icp": "string", "icpLink": "string", "isEnable": 1}, "loginFormDescription": "string", "loginFormTitle": "string", "loginPageDescription": "string", "loginPageTitle": "string", "sloganImage": "string", "sloganImageTemplate": "string", "sloganImageType": "template", "title": "string"}, "sms": {"accessKeyID": "string", "accessKeySecret": "string", "appID": "string", "isEnable": 1, "signName": "string", "type": "ali"}, "tenant": {"defaultLimit": {"aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}}, "user": {"captchaLen": 1, "forceSetPassword": "1", "ownerUserID": "string", "passLevel": 1, "remindSetPassword": "1", "userOwnTenantLimit": 1}, "weather": {"apiHost": "string", "apiKey": "string", "isEnable": 1}}, "ownerUser": {"avatar": "string", "email": "string", "nickName": "string", "password": "******", "phone": "string", "userName": "string"}}'
```

### POST `/api/v1/system/hook/capability/create`

**说明**: 新增Hook能力

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  能力编码 |
| `desc` | string | 否 |  描述 |
| `id` | string | 否 |  能力ID |
| `kind` | string | 否 |  能力类型 sync=同步扩展点 async=异步事件通知（说明性元数据，默认 sync） |
| `serverID` | string | 否 |  关联服务ID |
| `subCode` | string | 否 |  子编码 |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "id": "string",
  "kind": "string",
  "serverID": "string",
  "subCode": "string"
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
ur api /api/v1/system/hook/capability/create \
  --body '{"code": "string", "desc": "string", "id": "string", "kind": "string", "serverID": "string", "subCode": "string"}'
```

### POST `/api/v1/system/hook/capability/delete`

**说明**: 删除Hook能力

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/hook/capability/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/hook/capability/get-list`

**说明**: 获取Hook能力列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  能力编码筛选 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `serverID` | integer | 否 |  服务ID筛选 (格式: int64) |
| `subCode` | string | 否 |  子编码筛选 |

**请求示例**:
```json
{
  "code": "string",
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
  "serverID": 1,
  "subCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": "string",
        "desc": "string",
        "id": "string",
        "kind": "string",
        "serverID": "string",
        "subCode": "string"
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
ur api /api/v1/system/hook/capability/get-list \
  --body '{"code": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "serverID": 1, "subCode": "string"}'
```

### POST `/api/v1/system/hook/capability/update`

**说明**: 更新Hook能力

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  能力编码 |
| `desc` | string | 否 |  描述 |
| `id` | string | 否 |  能力ID |
| `kind` | string | 否 |  能力类型 sync=同步扩展点 async=异步事件通知（说明性元数据，默认 sync） |
| `serverID` | string | 否 |  关联服务ID |
| `subCode` | string | 否 |  子编码 |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "id": "string",
  "kind": "string",
  "serverID": "string",
  "subCode": "string"
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
ur api /api/v1/system/hook/capability/update \
  --body '{"code": "string", "desc": "string", "id": "string", "kind": "string", "serverID": "string", "subCode": "string"}'
```

### POST `/api/v1/system/hook/server/create`

**说明**: 新增Hook服务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authConfig` | object | 否 |  鉴权配置，hmac 时为 {"secret":"xxx"} |
| `authType` | string | 否 |  鉴权类型 none/hmac/custom |
| `capabilities` | array[HookCapabilityInfo] | 否 |  关联的能力列表 |
| `desc` | string | 否 |  描述 |
| `endpoint` | string | 否 |  服务端点URL |
| `failPolicy` | string | 否 |  失败策略 fail/ignore |
| `headers` | object | 否 |  自定义请求头（authType=custom 时使用） |
| `id` | string | 否 |  服务ID |
| `maxRetry` | integer | 否 |  最大重试次数 (格式: int64) |
| `name` | string | 否 |  服务名称（租户内唯一） |
| `status` | integer | 否 |  状态 1:启用 2:禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码：common=全平台可用，其余=仅该租户（创建时默认 common） |
| `timeoutSec` | integer | 否 |  超时秒数 (格式: int64) |

**请求示例**:
```json
{
  "authConfig": {},
  "authType": "string",
  "capabilities": [
    {
      "code": "string",
      "desc": "string",
      "id": "string",
      "kind": "string",
      "serverID": "string",
      "subCode": "string"
    }
  ],
  "desc": "string",
  "endpoint": "string",
  "failPolicy": "string",
  "headers": {},
  "id": "string",
  "maxRetry": 1,
  "name": "示例名称",
  "status": 1,
  "tenantCode": "string",
  "timeoutSec": 1
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
ur api /api/v1/system/hook/server/create \
  --body '{"authConfig": {}, "authType": "string", "capabilities": [{"code": "string", "desc": "string", "id": "string", "kind": "string", "serverID": "string", "subCode": "string"}], "desc": "string", "endpoint": "string", "failPolicy": "string", "headers": {}, "id": "string", "maxRetry": 1, "name": "示例名称", "status": 1, "tenantCode": "string", "timeoutSec": 1}'
```

### POST `/api/v1/system/hook/server/delete`

**说明**: 删除Hook服务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/hook/server/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/hook/server/get-list`

**说明**: 获取Hook服务列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  服务名称筛选 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  状态筛选 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码筛选 |

**请求示例**:
```json
{
  "name": "示例名称",
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
  "status": 1,
  "tenantCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "authConfig": {},
        "authType": "string",
        "capabilities": [
          {
            "code": "string",
            "desc": "string",
            "id": "string",
            "kind": "string",
            "serverID": "string",
            "subCode": "string"
          }
        ],
        "desc": "string",
        "endpoint": "string",
        "failPolicy": "string",
        "headers": {},
        "id": "string",
        "maxRetry": 1,
        "name": "示例名称",
        "status": 1,
        "tenantCode": "string",
        "timeoutSec": 1
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
ur api /api/v1/system/hook/server/get-list \
  --body '{"name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "tenantCode": "string"}'
```

### POST `/api/v1/system/hook/server/get-one`

**说明**: 获取Hook服务详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
  "data": {
    "authConfig": {},
    "authType": "string",
    "capabilities": [
      {
        "code": "string",
        "desc": "string",
        "id": "string",
        "kind": "string",
        "serverID": "string",
        "subCode": "string"
      }
    ],
    "desc": "string",
    "endpoint": "string",
    "failPolicy": "string",
    "headers": {},
    "id": "string",
    "maxRetry": 1,
    "name": "示例名称",
    "status": 1,
    "tenantCode": "string",
    "timeoutSec": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/hook/server/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/hook/server/update`

**说明**: 更新Hook服务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authConfig` | object | 否 |  鉴权配置，hmac 时为 {"secret":"xxx"} |
| `authType` | string | 否 |  鉴权类型 none/hmac/custom |
| `capabilities` | array[HookCapabilityInfo] | 否 |  关联的能力列表 |
| `desc` | string | 否 |  描述 |
| `endpoint` | string | 否 |  服务端点URL |
| `failPolicy` | string | 否 |  失败策略 fail/ignore |
| `headers` | object | 否 |  自定义请求头（authType=custom 时使用） |
| `id` | string | 否 |  服务ID |
| `maxRetry` | integer | 否 |  最大重试次数 (格式: int64) |
| `name` | string | 否 |  服务名称（租户内唯一） |
| `status` | integer | 否 |  状态 1:启用 2:禁用 (格式: int64) |
| `tenantCode` | string | 否 |  租户编码：common=全平台可用，其余=仅该租户（创建时默认 common） |
| `timeoutSec` | integer | 否 |  超时秒数 (格式: int64) |

**请求示例**:
```json
{
  "authConfig": {},
  "authType": "string",
  "capabilities": [
    {
      "code": "string",
      "desc": "string",
      "id": "string",
      "kind": "string",
      "serverID": "string",
      "subCode": "string"
    }
  ],
  "desc": "string",
  "endpoint": "string",
  "failPolicy": "string",
  "headers": {},
  "id": "string",
  "maxRetry": 1,
  "name": "示例名称",
  "status": 1,
  "tenantCode": "string",
  "timeoutSec": 1
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
ur api /api/v1/system/hook/server/update \
  --body '{"authConfig": {}, "authType": "string", "capabilities": [{"code": "string", "desc": "string", "id": "string", "kind": "string", "serverID": "string", "subCode": "string"}], "desc": "string", "endpoint": "string", "failPolicy": "string", "headers": {}, "id": "string", "maxRetry": 1, "name": "示例名称", "status": 1, "tenantCode": "string", "timeoutSec": 1}'
```

### GET `/api/v1/system/im/call/active`

**说明**: 查询群通话进行状态

**权限**: all

**路径参数**:

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| `conversationId` | query | integer | 是 |  群聊会话 ID |

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "active": true,
    "callId": "string",
    "callType": "string",
    "participantIds": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/call/active \
  --body '{}'
```

### POST `/api/v1/system/im/call/signal`

**说明**: 通话信令统一入口

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `callId` | string | 否 |  通话 ID（invite 类信令为空，由服务端生成） |
| `callType` | string | 否 |  通话媒体类型：voice / video |
| `calleeId` | string | 否 |  1:1 被叫用户 ID |
| `conversationId` | integer | 否 |  关联的 IM 会话 ID (格式: int64) |
| `members` | array[string] | 否 |  群通话邀请成员用户 ID 列表 |
| `payload` | string | 否 |  扩展 JSON（media_state 的 mute/video 状态等） |
| `to` | string | 否 |  media_state 定向目标用户 ID（空=1:1 对端/群通话全员） |
| `type` | string | 是 |  信令类型：call_invite/call_accept/call_reject/call_cancel/call_end/call_heartbeat/group_invite/group_join/group_leave/media_state |

**请求示例**:
```json
{
  "callId": "string",
  "callType": "string",
  "calleeId": "string",
  "conversationId": 1,
  "members": [
    "string"
  ],
  "payload": "string",
  "to": "string",
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "callId": "string",
    "event": "string",
    "notifyUserIDs": [
      1
    ],
    "payload": "string",
    "state": "string",
    "zlmPushUrl": "string",
    "zlmServer": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/call/signal \
  --body '{"callId": "string", "callType": "string", "calleeId": "string", "conversationId": 1, "members": ["string"], "payload": "string", "to": "string", "type": "string"}'
```

### POST `/api/v1/system/im/channel/bind`

**说明**: 绑定渠道端点

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `channelId` | string | 是 |  渠道标识（device 为 {productID}:{deviceName}） |
| `channelName` | string | 否 |  渠道展示名（设备名称） |
| `channelType` | string | 是 |  渠道类型：device / ... |

**请求示例**:
```json
{
  "channelId": "string",
  "channelName": "示例名称",
  "channelType": "string"
}
```

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
ur api /api/v1/system/im/channel/bind \
  --body '{"channelId": "string", "channelName": "示例名称", "channelType": "string"}'
```

### POST `/api/v1/system/im/channel/get-list`

**说明**: 查询渠道绑定列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "channelId": "string",
        "channelName": "string",
        "channelType": "string",
        "createdTime": 1,
        "id": 1,
        "isMuted": 1,
        "userId": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/channel/get-list \
  --body '{}'
```

### POST `/api/v1/system/im/channel/set-mute`

**说明**: 设置端点免打扰

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  绑定记录 ID (格式: int64) |
| `isMuted` | integer | 是 |  0-下发 1-免打扰 (格式: int64) |

**请求示例**:
```json
{
  "id": 1,
  "isMuted": 1
}
```

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
ur api /api/v1/system/im/channel/set-mute \
  --body '{"id": 1, "isMuted": 1}'
```

### POST `/api/v1/system/im/channel/unbind`

**说明**: 解绑渠道端点

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  绑定记录 ID (格式: int64) |

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
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/channel/unbind \
  --body '{"id": 1}'
```

### POST `/api/v1/system/im/chatlog/get-list`

**说明**: 获取消息列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "clientMsgId": "string",
        "content": "string",
        "conversationId": 1,
        "extra": "string",
        "id": 1,
        "isAtAll": 1,
        "mentionUserIDs": [
          1
        ],
        "msgType": "string",
        "sendTime": 1,
        "senderId": "string",
        "senderType": "string",
        "seq": 1,
        "status": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/chatlog/get-list \
  --body '{"conversationId": 1, "page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/chatlog/mark-read`

**说明**: 标记已读

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `readSeq` | integer | 否 |  已读到的序号 (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "readSeq": 1
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
ur api /api/v1/system/im/chatlog/mark-read \
  --body '{"conversationId": 1, "readSeq": 1}'
```

### POST `/api/v1/system/im/chatlog/search`

**说明**: 搜索聊天记录

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 否 |  会话 ID（可选，限定单个会话） (格式: int64) |
| `keyword` | string | 是 |  搜索关键词 |
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "keyword": "string",
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "clientMsgId": "string",
        "content": "string",
        "conversationId": 1,
        "extra": "string",
        "id": 1,
        "isAtAll": 1,
        "mentionUserIDs": [
          1
        ],
        "msgType": "string",
        "sendTime": 1,
        "senderId": "string",
        "senderType": "string",
        "seq": 1,
        "status": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/chatlog/search \
  --body '{"conversationId": 1, "keyword": "string", "page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/message/recall`

**说明**: 撤回消息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `msgId` | integer | 是 |  消息 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "msgId": 1
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
ur api /api/v1/system/im/message/recall \
  --body '{"conversationId": 1, "msgId": 1}'
```

### POST `/api/v1/system/im/message/send`

**说明**: 发送消息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `clientMsgId` | string | 否 |  客户端消息 ID |
| `content` | string | 是 |  消息内容 |
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `extra` | string | 否 |  扩展字段 |
| `isAtAll` | integer | 否 |  是否 @ 全体成员 (格式: int64) |
| `mentionUserIDs` | array[integer] | 否 |  @ 的用户列表 |
| `msgType` | string | 是 |  消息类型 |

**请求示例**:
```json
{
  "clientMsgId": "string",
  "content": "string",
  "conversationId": 1,
  "extra": "string",
  "isAtAll": 1,
  "mentionUserIDs": [
    1
  ],
  "msgType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "sendTime": 1,
    "seq": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/message/send \
  --body '{"clientMsgId": "string", "content": "string", "conversationId": 1, "extra": "string", "isAtAll": 1, "mentionUserIDs": [1], "msgType": "string"}'
```

### POST `/api/v1/system/im/conversation/add-members`

**说明**: 群聊邀请成员

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `participants` | array[ImParticipant] | 是 |  被邀请的参与者列表 |

**请求示例**:
```json
{
  "conversationId": 1,
  "participants": [
    {
      "id": "string",
      "type": "string"
    }
  ]
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
ur api /api/v1/system/im/conversation/add-members \
  --body '{"conversationId": 1, "participants": [{"id": "string", "type": "string"}]}'
```

### POST `/api/v1/system/im/conversation/create`

**说明**: 创建会话

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  会话名称（群聊必填） |
| `participants` | array[ImParticipant] | 是 |  参与者列表 |
| `type` | string | 是 |  会话类型：single / group |

**请求示例**:
```json
{
  "name": "示例名称",
  "participants": [
    {
      "id": "string",
      "type": "string"
    }
  ],
  "type": "string"
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
ur api /api/v1/system/im/conversation/create \
  --body '{"name": "示例名称", "participants": [{"id": "string", "type": "string"}], "type": "string"}'
```

### POST `/api/v1/system/im/conversation/create-ai`

**说明**: 创建 AI 会话

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cloneId` | integer | 是 |  AI 分身（Clone）ID (格式: int64) |

**请求示例**:
```json
{
  "cloneId": 1
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
ur api /api/v1/system/im/conversation/create-ai \
  --body '{"cloneId": 1}'
```

### POST `/api/v1/system/im/conversation/disband`

**说明**: 群聊解散

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1
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
ur api /api/v1/system/im/conversation/disband \
  --body '{"conversationId": 1}'
```

### POST `/api/v1/system/im/conversation/get`

**说明**: 获取会话详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "info": {
      "avatar": "string",
      "createdTime": 1,
      "id": 1,
      "isMuted": 1,
      "isPinned": 1,
      "lastMsgContent": "string",
      "lastMsgID": 1,
      "lastMsgTime": 1,
      "lastReadMsgID": 1,
      "name": "示例名称",
      "ownerID": 1,
      "status": 1,
      "type": "string",
      "unreadCount": 1,
      "updatedTime": 1
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/conversation/get \
  --body '{"conversationId": 1}'
```

### POST `/api/v1/system/im/conversation/get-by-participants`

**说明**: 按参与者查询单聊会话

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `userID1` | integer | 是 |  用户 1 ID (格式: int64) |
| `userID2` | integer | 是 |  用户 2 ID (格式: int64) |

**请求示例**:
```json
{
  "userID1": 1,
  "userID2": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avatar": "string",
    "createdTime": 1,
    "id": 1,
    "isMuted": 1,
    "isPinned": 1,
    "lastMsgContent": "string",
    "lastMsgID": 1,
    "lastMsgTime": 1,
    "lastReadMsgID": 1,
    "name": "示例名称",
    "ownerID": 1,
    "status": 1,
    "type": "string",
    "unreadCount": 1,
    "updatedTime": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/conversation/get-by-participants \
  --body '{"userID1": 1, "userID2": 1}'
```

### POST `/api/v1/system/im/conversation/get-list`

**说明**: 获取会话列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "avatar": "string",
        "createdTime": 1,
        "id": 1,
        "isMuted": 1,
        "isPinned": 1,
        "lastMsgContent": "string",
        "lastMsgID": 1,
        "lastMsgTime": 1,
        "lastReadMsgID": 1,
        "name": "示例名称",
        "ownerID": 1,
        "status": 1,
        "type": "string",
        "unreadCount": 1,
        "updatedTime": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/conversation/get-list \
  --body '{"page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/conversation/get-members`

**说明**: 获取会话成员列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "avatar": "string",
        "groupNickname": "string",
        "isMuted": 1,
        "joinedTime": 1,
        "nickName": "string",
        "participantId": "string",
        "participantType": "string",
        "role": "string",
        "userName": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/conversation/get-members \
  --body '{"conversationId": 1}'
```

### POST `/api/v1/system/im/conversation/mute-member`

**说明**: 群聊禁言/取消禁言成员

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `isMuted` | integer | 是 |  是否禁言：1 禁言 / 0 取消禁言 (格式: int64) |
| `participantId` | string | 是 |  参与者 ID |
| `participantType` | string | 是 |  参与者类型 |

**请求示例**:
```json
{
  "conversationId": 1,
  "isMuted": 1,
  "participantId": "string",
  "participantType": "string"
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
ur api /api/v1/system/im/conversation/mute-member \
  --body '{"conversationId": 1, "isMuted": 1, "participantId": "string", "participantType": "string"}'
```

### POST `/api/v1/system/im/conversation/remove-member`

**说明**: 群聊移除/退出成员

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `participantId` | string | 是 |  参与者 ID |
| `participantType` | string | 是 |  参与者类型 |

**请求示例**:
```json
{
  "conversationId": 1,
  "participantId": "string",
  "participantType": "string"
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
ur api /api/v1/system/im/conversation/remove-member \
  --body '{"conversationId": 1, "participantId": "string", "participantType": "string"}'
```

### POST `/api/v1/system/im/conversation/set-admin`

**说明**: 群聊设置/取消管理员

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `isAdmin` | integer | 是 |  是否管理员：1 设置 / 0 取消 (格式: int64) |
| `participantId` | string | 是 |  参与者 ID |
| `participantType` | string | 是 |  参与者类型 |

**请求示例**:
```json
{
  "conversationId": 1,
  "isAdmin": 1,
  "participantId": "string",
  "participantType": "string"
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
ur api /api/v1/system/im/conversation/set-admin \
  --body '{"conversationId": 1, "isAdmin": 1, "participantId": "string", "participantType": "string"}'
```

### POST `/api/v1/system/im/conversation/transfer-owner`

**说明**: 群聊转让群主

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `newOwnerID` | integer | 是 |  新群主用户 ID (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "newOwnerID": 1
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
ur api /api/v1/system/im/conversation/transfer-owner \
  --body '{"conversationId": 1, "newOwnerID": 1}'
```

### POST `/api/v1/system/im/conversation/update-setting`

**说明**: 更新会话设置

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `conversationId` | integer | 是 |  会话 ID (格式: int64) |
| `isDeleted` | integer | 否 |  是否删除（退出会话列表） (格式: int64) |
| `isMuted` | integer | 否 |  是否免打扰 (格式: int64) |
| `isPinned` | integer | 否 |  是否置顶 (格式: int64) |

**请求示例**:
```json
{
  "conversationId": 1,
  "isDeleted": 1,
  "isMuted": 1,
  "isPinned": 1
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
ur api /api/v1/system/im/conversation/update-setting \
  --body '{"conversationId": 1, "isDeleted": 1, "isMuted": 1, "isPinned": 1}'
```

### POST `/api/v1/system/im/friend/apply`

**说明**: 发起好友申请

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `recvUserID` | integer | 是 |  被申请人用户 ID (格式: int64) |
| `reqMsg` | string | 否 |  申请附言 |

**请求示例**:
```json
{
  "recvUserID": 1,
  "reqMsg": "string"
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
ur api /api/v1/system/im/friend/apply \
  --body '{"recvUserID": 1, "reqMsg": "string"}'
```

### POST `/api/v1/system/im/friend/delete`

**说明**: 删除好友

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `friendUserID` | integer | 是 |  好友用户 ID (格式: int64) |

**请求示例**:
```json
{
  "friendUserID": 1
}
```

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
ur api /api/v1/system/im/friend/delete \
  --body '{"friendUserID": 1}'
```

### POST `/api/v1/system/im/friend/get-list`

**说明**: 获取好友列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "avatar": "string",
        "createdTime": 1,
        "friendUserID": 1,
        "id": 1,
        "nickName": "string",
        "remark": "string",
        "userID": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/friend/get-list \
  --body '{"page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/friend/handle`

**说明**: 处理好友申请

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `handleMsg` | string | 否 |  处理附言 |
| `handleResult` | integer | 是 |  处理结果：1 同意 / 2 拒绝 (格式: int64) |
| `requestID` | integer | 是 |  申请记录 ID (格式: int64) |

**请求示例**:
```json
{
  "handleMsg": "string",
  "handleResult": 1,
  "requestID": 1
}
```

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
ur api /api/v1/system/im/friend/handle \
  --body '{"handleMsg": "string", "handleResult": 1, "requestID": 1}'
```

### POST `/api/v1/system/im/friend/request/get-list`

**说明**: 获取好友申请列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `direction` | integer | 是 |  方向：0 我收到的 / 1 我发起的 (格式: int64) |
| `handleResult` | integer | 是 |  处理结果过滤：-1 全部 (格式: int64) |
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "direction": 1,
  "handleResult": 1,
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": 1,
        "handleMsg": "string",
        "handleResult": 1,
        "id": 1,
        "recvUserAvatar": "string",
        "recvUserID": 1,
        "recvUserNickName": "string",
        "reqMsg": "string",
        "reqUserAvatar": "string",
        "reqUserID": 1,
        "reqUserNickName": "string",
        "updatedTime": 1
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/friend/request/get-list \
  --body '{"direction": 1, "handleResult": 1, "page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/moment/comment/create`

**说明**: 发表评论

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | object | 是 |  |
| `content.content` | string | 是 |  评论内容 |
| `content.mentionUserIDs` | array[integer] | 是 |  @ 的用户列表 |
| `id` | integer | 是 |  动态 ID (格式: int64) |

**请求示例**:
```json
{
  "content": {
    "content": "string",
    "mentionUserIDs": [
      1
    ]
  },
  "id": 1
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
ur api /api/v1/system/im/moment/comment/create \
  --body '{"content": {"content": "string", "mentionUserIDs": [1]}, "id": 1}'
```

### POST `/api/v1/system/im/moment/comment/delete`

**说明**: 删除评论

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `commentID` | integer | 是 |  评论 ID (格式: int64) |
| `momentID` | integer | 是 |  动态 ID (格式: int64) |

**请求示例**:
```json
{
  "commentID": 1,
  "momentID": 1
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
ur api /api/v1/system/im/moment/comment/delete \
  --body '{"commentID": 1, "momentID": 1}'
```

### POST `/api/v1/system/im/moment/create`

**说明**: 发布动态

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 |  文本内容 |
| `mediaList` | array[ImMomentMedia] | 是 |  媒体列表 |
| `mentionUserIDs` | array[integer] | 是 |  @ 的用户列表 |
| `projectIDs` | array[integer] | 是 |  可见项目列表（visibility=1 时生效） |
| `visibility` | integer | 是 |  可见范围：0 公开 / 1 项目 / 2 指定人 (格式: int64) |
| `visibleUserIDs` | array[integer] | 是 |  可见用户列表（visibility=2 时生效） |

**请求示例**:
```json
{
  "content": "string",
  "mediaList": [
    {
      "duration": 1,
      "name": "string",
      "size": 1,
      "thumb": "string",
      "type": "string",
      "url": "string"
    }
  ],
  "mentionUserIDs": [
    1
  ],
  "projectIDs": [
    1
  ],
  "visibility": 1,
  "visibleUserIDs": [
    1
  ]
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
ur api /api/v1/system/im/moment/create \
  --body '{"content": "string", "mediaList": [{"duration": 1, "name": "string", "size": 1, "thumb": "string", "type": "string", "url": "string"}], "mentionUserIDs": [1], "projectIDs": [1], "visibility": 1, "visibleUserIDs": [1]}'
```

### POST `/api/v1/system/im/moment/delete`

**说明**: 删除动态

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  动态 ID (格式: int64) |

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
ur api /api/v1/system/im/moment/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/im/moment/get`

**说明**: 获取动态详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  动态 ID (格式: int64) |

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
  "data": {
    "comments": [
      {
        "content": "string",
        "createdTime": 1,
        "id": 1,
        "mentionUserIDs": [
          1
        ],
        "momentID": 1,
        "status": 1,
        "updatedTime": 1,
        "userID": 1
      }
    ],
    "info": {
      "commentCount": 1,
      "content": "string",
      "createdTime": 1,
      "id": 1,
      "isLiked": true,
      "likeCount": 1,
      "mediaList": [
        {
          "duration": 1,
          "name": "string",
          "size": 1,
          "thumb": "string",
          "type": "string",
          "url": "string"
        }
      ],
      "mentionUserIDs": [
        1
      ],
      "projectIDs": [
        1
      ],
      "status": 1,
      "updatedTime": 1,
      "userID": 1,
      "visibility": 1,
      "visibleUserIDs": [
        1
      ]
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/moment/get \
  --body '{"id": 1}'
```

### POST `/api/v1/system/im/moment/like`

**说明**: 点赞动态

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  动态 ID (格式: int64) |

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
  "data": {
    "likeCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/moment/like \
  --body '{"id": 1}'
```

### POST `/api/v1/system/im/moment/list`

**说明**: 查询动态列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `filterProjectID` | integer | 否 |  按项目过滤 (格式: int64) |
| `filterUserID` | integer | 否 |  按发布人过滤 (格式: int64) |
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "filterProjectID": 1,
  "filterUserID": 1,
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "commentCount": 1,
        "content": "string",
        "createdTime": 1,
        "id": 1,
        "isLiked": true,
        "likeCount": 1,
        "mediaList": [
          {
            "duration": 1,
            "name": "string",
            "size": 1,
            "thumb": "string",
            "type": "string",
            "url": "string"
          }
        ],
        "mentionUserIDs": [
          1
        ],
        "projectIDs": [
          1
        ],
        "status": 1,
        "updatedTime": 1,
        "userID": 1,
        "visibility": 1,
        "visibleUserIDs": [
          1
        ]
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/moment/list \
  --body '{"filterProjectID": 1, "filterUserID": 1, "page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/moment/unlike`

**说明**: 取消点赞动态

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  动态 ID (格式: int64) |

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
  "data": {
    "likeCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/moment/unlike \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/cancel`

**说明**: 取消执行任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `taskID` | string | 是 | 运行的taskID |

**请求示例**:
```json
{
  "taskID": "string"
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
ur api /api/v1/system/job/task/cancel \
  --body '{"taskID": "string"}'
```

### POST `/api/v1/system/job/task/group/create`

**说明**: 新建任务分组

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  任务组编码 |
| `createdTime` | string | 否 |  创建时间（Unix 毫秒） |
| `desc` | string | 否 |  描述 |
| `id` | string | 否 |  任务组ID |
| `name` | string | 否 |  组名 |
| `status` | integer | 否 |  状态 1:启用 2:禁用 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "name": "string",
  "status": 1
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
ur api /api/v1/system/job/task/group/create \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "name": "string", "status": 1}'
```

### POST `/api/v1/system/job/task/group/delete`

**说明**: 删除任务分组

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string"
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
ur api /api/v1/system/job/task/group/delete \
  --body '{"code": "string"}'
```

### POST `/api/v1/system/job/task/group/get-list`

**说明**: 获取任务分组列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  状态过滤 1:启用 2:禁用 (格式: int64) |

**请求示例**:
```json
{
  "name": "示例名称",
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
  "status": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "id": "string",
        "name": "string",
        "status": 1
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
ur api /api/v1/system/job/task/group/get-list \
  --body '{"name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1}'
```

### POST `/api/v1/system/job/task/group/get-one`

**说明**: 获取任务分组详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "name": "string",
    "status": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/job/task/group/get-one \
  --body '{"code": "string"}'
```

### POST `/api/v1/system/job/task/group/update`

**说明**: 更新任务分组

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  任务组编码 |
| `createdTime` | string | 否 |  创建时间（Unix 毫秒） |
| `desc` | string | 否 |  描述 |
| `id` | string | 否 |  任务组ID |
| `name` | string | 否 |  组名 |
| `status` | integer | 否 |  状态 1:启用 2:禁用 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "name": "string",
  "status": 1
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
ur api /api/v1/system/job/task/group/update \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "name": "string", "status": 1}'
```

### POST `/api/v1/system/job/task/info/create`

**说明**: 创建任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  任务编码 |
| `createdTime` | string | 否 |  创建时间（Unix 毫秒） |
| `cronExpr` | string | 否 |  Cron 表达式（type=1） |
| `desc` | string | 否 |  描述 |
| `executeAt` | string | 否 |  执行时间 Unix 毫秒（type=2） |
| `groupCode` | string | 是 |  关联分组编码 |
| `id` | string | 否 |  任务ID |
| `lastRunTime` | string | 否 |  上次执行时间 Unix 毫秒 |
| `name` | string | 否 |  任务名称 |
| `params` | string | 否 |  任务参数（JSON 字符串） |
| `priority` | integer | 否 |  优先级 1-10 (格式: int64) |
| `retriedCount` | integer | 否 |  已重试次数 (格式: int64) |
| `retryPolicy` | object | 否 |  |
| `retryPolicy.backoff` | string | 否 |  退避策略：fixed / exponential |
| `retryPolicy.maxRetry` | integer | 否 |  最大重试次数，默认 3 (格式: int64) |
| `retryPolicy.retryDelay` | integer | 否 |  重试间隔基数（秒），默认 10 (格式: int64) |
| `status` | integer | 否 |  状态 1:WaitRun 2:Running 3:WaitStop 4:Stopped 5:WaitDelete (格式: int64) |
| `type` | integer | 否 |  类型 1:定时(Cron) 2:延时(Delay) (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "cronExpr": "string",
  "desc": "string",
  "executeAt": "2026-01-01T00:00:00Z",
  "groupCode": "string",
  "id": "string",
  "lastRunTime": "2026-01-01T00:00:00Z",
  "name": "示例名称",
  "params": "string",
  "priority": 1,
  "retriedCount": 1,
  "retryPolicy": {
    "backoff": "string",
    "maxRetry": 1,
    "retryDelay": 1
  },
  "status": 1,
  "type": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "cronExpr": "string",
    "desc": "string",
    "executeAt": "2026-01-01T00:00:00Z",
    "groupCode": "string",
    "id": "string",
    "lastRunTime": "2026-01-01T00:00:00Z",
    "name": "示例名称",
    "params": "string",
    "priority": 1,
    "retriedCount": 1,
    "retryPolicy": {
      "backoff": "string",
      "maxRetry": 1,
      "retryDelay": 1
    },
    "status": 1,
    "type": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/job/task/info/create \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "cronExpr": "string", "desc": "string", "executeAt": "2026-01-01T00:00:00Z", "groupCode": "string", "id": "string", "lastRunTime": "2026-01-01T00:00:00Z", "name": "示例名称", "params": "string", "priority": 1, "retriedCount": 1, "retryPolicy": {"backoff": "string", "maxRetry": 1, "retryDelay": 1}, "status": 1, "type": 1}'
```

### POST `/api/v1/system/job/task/info/delete`

**说明**: 删除任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/job/task/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/info/get-list`

**说明**: 获取任务列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupCode` | string | 是 |  关联分组编码 |
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  按状态过滤 (格式: int64) |
| `type` | integer | 否 |  按类型过滤 1:定时 2:延时 (格式: int64) |

**请求示例**:
```json
{
  "groupCode": "string",
  "name": "示例名称",
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
  "status": 1,
  "type": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "cronExpr": "string",
        "desc": "string",
        "executeAt": "2026-01-01T00:00:00Z",
        "groupCode": "string",
        "id": "string",
        "lastRunTime": "2026-01-01T00:00:00Z",
        "name": "示例名称",
        "params": "string",
        "priority": 1,
        "retriedCount": 1,
        "retryPolicy": {
          "backoff": "string",
          "maxRetry": 1,
          "retryDelay": 1
        },
        "status": 1,
        "type": 1
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
ur api /api/v1/system/job/task/info/get-list \
  --body '{"groupCode": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "type": 1}'
```

### POST `/api/v1/system/job/task/info/get-one`

**说明**: 获取任务详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
  "data": {
    "code": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "cronExpr": "string",
    "desc": "string",
    "executeAt": "2026-01-01T00:00:00Z",
    "groupCode": "string",
    "id": "string",
    "lastRunTime": "2026-01-01T00:00:00Z",
    "name": "示例名称",
    "params": "string",
    "priority": 1,
    "retriedCount": 1,
    "retryPolicy": {
      "backoff": "string",
      "maxRetry": 1,
      "retryDelay": 1
    },
    "status": 1,
    "type": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/job/task/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/info/start`

**说明**: 启动任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/job/task/info/start \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/info/stop`

**说明**: 停止任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/job/task/info/stop \
  --body '{"id": 1}'
```

### POST `/api/v1/system/job/task/info/update`

**说明**: 更新任务

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  任务编码 |
| `createdTime` | string | 否 |  创建时间（Unix 毫秒） |
| `cronExpr` | string | 否 |  Cron 表达式（type=1） |
| `desc` | string | 否 |  描述 |
| `executeAt` | string | 否 |  执行时间 Unix 毫秒（type=2） |
| `groupCode` | string | 是 |  关联分组编码 |
| `id` | string | 否 |  任务ID |
| `lastRunTime` | string | 否 |  上次执行时间 Unix 毫秒 |
| `name` | string | 否 |  任务名称 |
| `params` | string | 否 |  任务参数（JSON 字符串） |
| `priority` | integer | 否 |  优先级 1-10 (格式: int64) |
| `retriedCount` | integer | 否 |  已重试次数 (格式: int64) |
| `retryPolicy` | object | 否 |  |
| `retryPolicy.backoff` | string | 否 |  退避策略：fixed / exponential |
| `retryPolicy.maxRetry` | integer | 否 |  最大重试次数，默认 3 (格式: int64) |
| `retryPolicy.retryDelay` | integer | 否 |  重试间隔基数（秒），默认 10 (格式: int64) |
| `status` | integer | 否 |  状态 1:WaitRun 2:Running 3:WaitStop 4:Stopped 5:WaitDelete (格式: int64) |
| `type` | integer | 否 |  类型 1:定时(Cron) 2:延时(Delay) (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "cronExpr": "string",
  "desc": "string",
  "executeAt": "2026-01-01T00:00:00Z",
  "groupCode": "string",
  "id": "string",
  "lastRunTime": "2026-01-01T00:00:00Z",
  "name": "示例名称",
  "params": "string",
  "priority": 1,
  "retriedCount": 1,
  "retryPolicy": {
    "backoff": "string",
    "maxRetry": 1,
    "retryDelay": 1
  },
  "status": 1,
  "type": 1
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
ur api /api/v1/system/job/task/info/update \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "cronExpr": "string", "desc": "string", "executeAt": "2026-01-01T00:00:00Z", "groupCode": "string", "id": "string", "lastRunTime": "2026-01-01T00:00:00Z", "name": "示例名称", "params": "string", "priority": 1, "retriedCount": 1, "retryPolicy": {"backoff": "string", "maxRetry": 1, "retryDelay": 1}, "status": 1, "type": 1}'
```

### POST `/api/v1/system/job/task/send`

**说明**: 发送延时请求

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 | 任务code |
| `groupCode` | string | 是 | 组需要提前创建好 |
| `option` | object | 否 |  |
| `option.deadline` | integer | 否 | 截止时间  秒时间戳 (格式: int64) |
| `option.priority` | integer | 否 | 优先级: 6:critical 最高优先级  3: default 普通优先级 1:low 低优先级以下两个参数优先使用ProcessIn (格式: int64) |
| `option.processAt` | integer | 否 |  固定时间发 秒时间戳 (格式: int64) |
| `option.processIn` | integer | 否 | 多久之后发 秒数 (格式: int64) |
| `option.taskID` | string | 否 | 指定运行的taskID |
| `option.timeout` | integer | 否 | 超时时间 优先使用 秒数 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "groupCode": "string",
  "option": {
    "deadline": 1,
    "priority": 1,
    "processAt": 1,
    "processIn": 1,
    "taskID": "string",
    "timeout": 1
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "taskID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/job/task/send \
  --body '{"code": "string", "groupCode": "string", "option": {"deadline": 1, "priority": 1, "processAt": 1, "processIn": 1, "taskID": "string", "timeout": 1}}'
```

### POST `/api/v1/system/mall/license/batch-create`

**说明**: 批量创建授权码

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `count` | integer | 否 | 格式: int32 |
| `desc` | string | 否 |  |
| `issuer` | string | 否 |  |
| `skuCode` | string | 否 |  |
| `sourceID` | string | 是 |  |
| `sourceType` | string | 是 |  |

**请求示例**:
```json
{
  "count": 1,
  "desc": "string",
  "issuer": "string",
  "skuCode": "string",
  "sourceID": "string",
  "sourceType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "count": 1,
    "licenseCodes": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/license/batch-create \
  --body '{"count": 1, "desc": "string", "issuer": "string", "skuCode": "string", "sourceID": "string", "sourceType": "string"}'
```

### POST `/api/v1/system/mall/license/create`

**说明**: 创建授权码

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `issuer` | string | 否 |  |
| `skuCode` | string | 否 |  |
| `sourceID` | string | 是 |  |
| `sourceType` | string | 是 |  |

**请求示例**:
```json
{
  "desc": "string",
  "issuer": "string",
  "skuCode": "string",
  "sourceID": "string",
  "sourceType": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "licenseCode": "string",
    "withSource": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/license/create \
  --body '{"desc": "string", "issuer": "string", "skuCode": "string", "sourceID": "string", "sourceType": "string"}'
```

### POST `/api/v1/system/mall/license/get-list`

**说明**: 授权码列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `boundTargetID` | string | 否 |  |
| `licenseCode` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `skuCode` | string | 否 |  |
| `sourceID` | string | 否 |  |
| `sourceType` | string | 否 |  |
| `status` | string | 否 |  |
| `withSource` | boolean | 否 |  是否返回来源对象详情（商品或套餐） (格式: boolean) |

**请求示例**:
```json
{
  "boundTargetID": "string",
  "licenseCode": "string",
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
  "skuCode": "string",
  "sourceID": "string",
  "sourceType": "string",
  "status": "string",
  "withSource": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "boundMeta": {},
        "boundTargetID": "string",
        "boundTime": "string",
        "createdTime": "string",
        "desc": "string",
        "grantResult": "string",
        "grantSnapshot": "string",
        "id": "string",
        "issuer": "string",
        "licenseCode": "string",
        "skuCode": "string",
        "sourceID": "string",
        "sourcePackage": {
          "category": "string",
          "createdTime": "string",
          "desc": "string",
          "id": "string",
          "items": [
            {
              "productID": "...",
              "skuCode": "...",
              "sort": "...",
              "status": "..."
            }
          ],
          "packageCode": "string",
          "packageName": "string",
          "status": 1,
          "verifyMethod": "string"
        },
        "sourceProduct": {
          "createdTime": "string",
          "desc": "string",
          "grantSchema": "string",
          "id": "string",
          "productCode": "string",
          "productName": "string",
          "productType": "string",
          "skus": [
            {
              "bindMode": "...",
              "configValues": "...",
              "durationDays": "...",
              "licenseMode": "...",
              "skuCode": "...",
              "skuName": "...",
              "status": "..."
            }
          ],
          "status": 1
        },
        "sourceType": "string",
        "status": "string",
        "validFrom": "string",
        "validUntil": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/license/get-list \
  --body '{"boundTargetID": "string", "licenseCode": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "skuCode": "string", "sourceID": "string", "sourceType": "string", "status": "string", "withSource": true}'
```

### POST `/api/v1/system/mall/license/get-one`

**说明**: 授权码详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `licenseCode` | string | 是 |  授权码 |
| `withSource` | boolean | 否 |  是否返回来源对象详情（商品或套餐） (格式: boolean) |

**请求示例**:
```json
{
  "licenseCode": "string",
  "withSource": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "boundMeta": {},
    "boundTargetID": "string",
    "boundTime": "string",
    "createdTime": "string",
    "desc": "string",
    "grantResult": "string",
    "grantSnapshot": "string",
    "id": "string",
    "issuer": "string",
    "licenseCode": "string",
    "skuCode": "string",
    "sourceID": "string",
    "sourcePackage": {
      "category": "string",
      "createdTime": "string",
      "desc": "string",
      "id": "string",
      "items": [
        {
          "productID": "string",
          "skuCode": "string",
          "sort": 1,
          "status": 1
        }
      ],
      "packageCode": "string",
      "packageName": "string",
      "status": 1,
      "verifyMethod": "string"
    },
    "sourceProduct": {
      "createdTime": "string",
      "desc": "string",
      "grantSchema": "string",
      "id": "string",
      "productCode": "string",
      "productName": "string",
      "productType": "string",
      "skus": [
        {
          "bindMode": "string",
          "configValues": {},
          "durationDays": 1,
          "licenseMode": "string",
          "skuCode": "string",
          "skuName": "string",
          "status": 1
        }
      ],
      "status": 1
    },
    "sourceType": "string",
    "status": "string",
    "validFrom": "string",
    "validUntil": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/license/get-one \
  --body '{"licenseCode": "string", "withSource": true}'
```

### POST `/api/v1/system/mall/license/revoke`

**说明**: 撤销授权码

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `licenseCode` | string | 是 |  授权码 |
| `withSource` | boolean | 否 |  是否返回来源对象详情（商品或套餐） (格式: boolean) |

**请求示例**:
```json
{
  "licenseCode": "string",
  "withSource": true
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
ur api /api/v1/system/mall/license/revoke \
  --body '{"licenseCode": "string", "withSource": true}'
```

### POST `/api/v1/system/mall/package/create`

**说明**: 创建套餐

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `category` | string | 否 |  |
| `desc` | string | 否 |  |
| `items` | array[MallPackageItem] | 否 |  |
| `packageCode` | string | 是 |  |
| `packageName` | string | 是 |  |
| `status` | integer | 否 | 格式: int32 |
| `verifyMethod` | string | 否 |  |

**请求示例**:
```json
{
  "category": "string",
  "desc": "string",
  "items": [
    {
      "productID": "string",
      "skuCode": "string",
      "sort": 1,
      "status": 1
    }
  ],
  "packageCode": "string",
  "packageName": "string",
  "status": 1,
  "verifyMethod": "string"
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
ur api /api/v1/system/mall/package/create \
  --body '{"category": "string", "desc": "string", "items": [{"productID": "string", "skuCode": "string", "sort": 1, "status": 1}], "packageCode": "string", "packageName": "string", "status": 1, "verifyMethod": "string"}'
```

### POST `/api/v1/system/mall/package/delete`

**说明**: 删除套餐

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/mall/package/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/mall/package/get-list`

**说明**: 获取套餐列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `category` | string | 否 |  |
| `packageCode` | string | 否 |  |
| `packageName` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int32 |

**请求示例**:
```json
{
  "category": "string",
  "packageCode": "string",
  "packageName": "string",
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
  "status": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "category": "string",
        "createdTime": "string",
        "desc": "string",
        "id": "string",
        "items": [
          {
            "productID": "string",
            "skuCode": "string",
            "sort": 1,
            "status": 1
          }
        ],
        "packageCode": "string",
        "packageName": "string",
        "status": 1,
        "verifyMethod": "string"
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/package/get-list \
  --body '{"category": "string", "packageCode": "string", "packageName": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1}'
```

### POST `/api/v1/system/mall/package/get-one`

**说明**: 获取套餐详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
  "data": {
    "category": "string",
    "createdTime": "string",
    "desc": "string",
    "id": "string",
    "items": [
      {
        "productID": "string",
        "skuCode": "string",
        "sort": 1,
        "status": 1
      }
    ],
    "packageCode": "string",
    "packageName": "string",
    "status": 1,
    "verifyMethod": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/package/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/mall/package/update`

**说明**: 更新套餐

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `category` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 是 |  |
| `items` | array[MallPackageItem] | 否 |  |
| `packageName` | string | 否 |  |
| `status` | integer | 否 | 格式: int32 |
| `verifyMethod` | string | 否 |  |

**请求示例**:
```json
{
  "category": "string",
  "desc": "string",
  "id": "string",
  "items": [
    {
      "productID": "string",
      "skuCode": "string",
      "sort": 1,
      "status": 1
    }
  ],
  "packageName": "string",
  "status": 1,
  "verifyMethod": "string"
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
ur api /api/v1/system/mall/package/update \
  --body '{"category": "string", "desc": "string", "id": "string", "items": [{"productID": "string", "skuCode": "string", "sort": 1, "status": 1}], "packageName": "string", "status": 1, "verifyMethod": "string"}'
```

### POST `/api/v1/system/mall/product/create`

**说明**: 创建商品

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `grantSchema` | string | 否 |  |
| `productCode` | string | 是 |  |
| `productName` | string | 是 |  |
| `productType` | string | 是 |  |
| `skus` | array[MallProductSku] | 否 |  |
| `status` | integer | 否 | 格式: int32 |

**请求示例**:
```json
{
  "desc": "string",
  "grantSchema": "string",
  "productCode": "string",
  "productName": "string",
  "productType": "string",
  "skus": [
    {
      "bindMode": "string",
      "configValues": {},
      "durationDays": 1,
      "licenseMode": "string",
      "skuCode": "string",
      "skuName": "string",
      "status": 1
    }
  ],
  "status": 1
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
ur api /api/v1/system/mall/product/create \
  --body '{"desc": "string", "grantSchema": "string", "productCode": "string", "productName": "string", "productType": "string", "skus": [{"bindMode": "string", "configValues": {}, "durationDays": 1, "licenseMode": "string", "skuCode": "string", "skuName": "string", "status": 1}], "status": 1}'
```

### POST `/api/v1/system/mall/product/delete`

**说明**: 删除商品

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
ur api /api/v1/system/mall/product/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/mall/product/get-list`

**说明**: 获取商品列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productCode` | string | 否 |  |
| `productName` | string | 否 |  |
| `productType` | string | 否 |  |
| `status` | integer | 否 | 格式: int32 |

**请求示例**:
```json
{
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
  "productCode": "string",
  "productName": "string",
  "productType": "string",
  "status": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "string",
        "desc": "string",
        "grantSchema": "string",
        "id": "string",
        "productCode": "string",
        "productName": "string",
        "productType": "string",
        "skus": [
          {
            "bindMode": "string",
            "configValues": {},
            "durationDays": 1,
            "licenseMode": "string",
            "skuCode": "string",
            "skuName": "string",
            "status": 1
          }
        ],
        "status": 1
      }
    ],
    "total": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/product/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productCode": "string", "productName": "string", "productType": "string", "status": 1}'
```

### POST `/api/v1/system/mall/product/get-one`

**说明**: 获取商品详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 否 |  id (格式: int64) |

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
  "data": {
    "createdTime": "string",
    "desc": "string",
    "grantSchema": "string",
    "id": "string",
    "productCode": "string",
    "productName": "string",
    "productType": "string",
    "skus": [
      {
        "bindMode": "string",
        "configValues": {},
        "durationDays": 1,
        "licenseMode": "string",
        "skuCode": "string",
        "skuName": "string",
        "status": 1
      }
    ],
    "status": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mall/product/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/mall/product/update`

**说明**: 更新商品

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `grantSchema` | string | 否 |  |
| `id` | string | 是 |  |
| `productName` | string | 否 |  |
| `productType` | string | 否 |  |
| `skus` | array[MallProductSku] | 否 |  |
| `status` | integer | 否 | 格式: int32 |

**请求示例**:
```json
{
  "desc": "string",
  "grantSchema": "string",
  "id": "string",
  "productName": "string",
  "productType": "string",
  "skus": [
    {
      "bindMode": "string",
      "configValues": {},
      "durationDays": 1,
      "licenseMode": "string",
      "skuCode": "string",
      "skuName": "string",
      "status": 1
    }
  ],
  "status": 1
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
ur api /api/v1/system/mall/product/update \
  --body '{"desc": "string", "grantSchema": "string", "id": "string", "productName": "string", "productType": "string", "skus": [{"bindMode": "string", "configValues": {}, "durationDays": 1, "licenseMode": "string", "skuCode": "string", "skuName": "string", "status": 1}], "status": 1}'
```

### POST `/api/v1/system/mcp/run`

**说明**: Stateless MCP HTTP

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/mcp/run \
  --body '{}'
```

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

### POST `/api/v1/system/ops/work-order/create`

**说明**: 添加工单

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id |
| `createdTime` | string | 否 | 创建时间 |
| `finishedTime` | integer | 否 | 处理完成时间 (格式: int64) |
| `handleTime` | integer | 否 | 处理时间 (格式: int64) |
| `id` | string | 否 |  |
| `issueDesc` | string | 否 | 问题描述 |
| `number` | string | 否 | 编号 |
| `params` | object | 否 | 参数 json格式 |
| `raiseUserID` | string | 否 | 问题提出的用户 |
| `status` | integer | 否 |  状态（1 待处理 2:处理中 3:已完成） (格式: int64) |
| `type` | string | 否 | 工单类型: deviceMaintenance:设备维修工单 可选: ['[deviceMaintenance]'] |

**请求示例**:
```json
{
  "areaID": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "finishedTime": 1,
  "handleTime": 1,
  "id": "string",
  "issueDesc": "string",
  "number": "string",
  "params": {},
  "raiseUserID": "string",
  "status": 1,
  "type": "[deviceMaintenance]"
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
ur api /api/v1/system/ops/work-order/create \
  --body '{"areaID": "string", "createdTime": "2026-01-01T00:00:00Z", "finishedTime": 1, "handleTime": 1, "id": "string", "issueDesc": "string", "number": "string", "params": {}, "raiseUserID": "string", "status": 1, "type": "[deviceMaintenance]"}'
```

### POST `/api/v1/system/ops/work-order/get-list`

**说明**: 获取工单列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 区域ID过滤 |
| `number` | string | 否 | 编号 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int64 |
| `type` | string | 否 | 工单类型: deviceMaintenance:设备维修工单 |

**请求示例**:
```json
{
  "areaID": "string",
  "number": "string",
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
  "status": 1,
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
        "areaID": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "finishedTime": 1,
        "handleTime": 1,
        "id": "string",
        "issueDesc": "string",
        "number": "string",
        "params": {},
        "raiseUserID": "string",
        "status": 1,
        "type": "[deviceMaintenance]"
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
ur api /api/v1/system/ops/work-order/get-list \
  --body '{"areaID": "string", "number": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "type": "string"}'
```

### POST `/api/v1/system/ops/work-order/update`

**说明**: 更新工单

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 项目区域id |
| `createdTime` | string | 否 | 创建时间 |
| `finishedTime` | integer | 否 | 处理完成时间 (格式: int64) |
| `handleTime` | integer | 否 | 处理时间 (格式: int64) |
| `id` | string | 否 |  |
| `issueDesc` | string | 否 | 问题描述 |
| `number` | string | 否 | 编号 |
| `params` | object | 否 | 参数 json格式 |
| `raiseUserID` | string | 否 | 问题提出的用户 |
| `status` | integer | 否 |  状态（1 待处理 2:处理中 3:已完成） (格式: int64) |
| `type` | string | 否 | 工单类型: deviceMaintenance:设备维修工单 可选: ['[deviceMaintenance]'] |

**请求示例**:
```json
{
  "areaID": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "finishedTime": 1,
  "handleTime": 1,
  "id": "string",
  "issueDesc": "string",
  "number": "string",
  "params": {},
  "raiseUserID": "string",
  "status": 1,
  "type": "[deviceMaintenance]"
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
ur api /api/v1/system/ops/work-order/update \
  --body '{"areaID": "string", "createdTime": "2026-01-01T00:00:00Z", "finishedTime": 1, "handleTime": 1, "id": "string", "issueDesc": "string", "number": "string", "params": {}, "raiseUserID": "string", "status": 1, "type": "[deviceMaintenance]"}'
```

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
| `id` | integer | 否 |  id (格式: int64) |

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
| `authType` | integer | 否 |   1(all) 全部人可以操作 2(admin) 默认授予租户管理员权限 3(superAdmin,supper) default租户才可以操作(超管是跨租户的) (格式: int64) |
| `authTypeCode` | string | 否 |  权限层级编码 |
| `groupCode` | string | 否 |  分组标识 |
| `method` | string | 否 |  接口请求方式: （1 GET 2 POST 3 HEAD 4 OPTIONS 5 PUT 6 DELETE 7 TRACE 8 CONNECT 9 其它） |
| `name` | string | 否 |  接口名称 |
| `operTypeCode` | string | 否 |  操作类型编码 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
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
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
  --body '{"accessCode": "string", "authType": 1, "authTypeCode": "string", "groupCode": "string", "method": "string", "name": "示例名称", "operTypeCode": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "resource": "string", "route": "string"}'
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
