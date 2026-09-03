# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/dept/info/create` | 添加部门详情 | admin |
| POST | `/api/v1/system/dept/info/delete` | 删除部门 | admin |
| POST | `/api/v1/system/dept/info/get-list` | 获取部门列表 | admin |
| POST | `/api/v1/system/dept/info/get-one` | 获取部门详情 | admin |
| POST | `/api/v1/system/dept/info/update` | 更新部门 | admin |
| POST | `/api/v1/system/dept/role/batch-create` | 批量添加部门角色 | admin |
| POST | `/api/v1/system/dept/role/batch-delete` | 批量移除部门角色 | admin |
| POST | `/api/v1/system/dept/role/get-list` | 获取部门角色列表 | admin |
| POST | `/api/v1/system/dept/sync-job/create` | 添加同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/delete` | 删除同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/execute` | 执行同步任务 | admin |
| POST | `/api/v1/system/dept/sync-job/get-list` | 获取同步任务列表 | admin |
| POST | `/api/v1/system/dept/sync-job/get-one` | 获取同步任务详情 | admin |
| POST | `/api/v1/system/dept/sync-job/update` | 更新同步任务 | admin |
| POST | `/api/v1/system/dept/user/batch-create` | 批量授权部门用户 | admin |
| POST | `/api/v1/system/dept/user/batch-delete` | 批量取消授权部门用户 | admin |
| POST | `/api/v1/system/dept/user/get-list` | 获取部门授权列表 | admin |
| POST | `/api/v1/system/dict/detail/create` | 添加字典详情 | platform |
| POST | `/api/v1/system/dict/detail/delete` | 删除字典详情 | platform |
| POST | `/api/v1/system/dict/detail/get-list` | 获取字典详情列表 | platform |
| POST | `/api/v1/system/dict/detail/get-one` | 获取字典详情单个 | platform |
| POST | `/api/v1/system/dict/detail/update` | 更新字典详情 | platform |
| POST | `/api/v1/system/dict/info/batch-export` | 批量导出字典信息 | platform |
| POST | `/api/v1/system/dict/info/batch-import` | 批量导入字典信息 | platform |
| POST | `/api/v1/system/dict/info/create` | 添加字典信息 | platform |
| POST | `/api/v1/system/dict/info/delete` | 删除字典信息 | platform |
| POST | `/api/v1/system/dict/info/get-list` | 获取字典信息列表 | all |
| POST | `/api/v1/system/dict/info/get-one` | 获取字典信息详情 | platform |
| POST | `/api/v1/system/dict/info/update` | 更新字典信息 | platform |
| POST | `/api/v1/system/log/login/get-list` | 获取登录日志列表 | admin |
| POST | `/api/v1/system/log/oper/get-list` | 获取操作日志列表 | admin |
| POST | `/api/v1/system/notify/config/create` | 添加通知配置 | platform |
| POST | `/api/v1/system/notify/config/delete` | 删除通知配置 | platform |
| POST | `/api/v1/system/notify/config/get-list` | 获取通知列表 | platform |
| POST | `/api/v1/system/notify/config/get-one` | 获取通知详情 | platform |
| POST | `/api/v1/system/notify/config/send-test` | 测试发送通知 | platform |
| POST | `/api/v1/system/notify/config/tree` | 获取通知列表树 | platform |
| POST | `/api/v1/system/notify/config/update` | 更新通知配置 | platform |
| POST | `/api/v1/system/notify/config/template/delete` | 删除通知绑定模版 | platform |
| POST | `/api/v1/system/notify/config/template/get-list` | 获取通知绑定模版列表 | platform |
| POST | `/api/v1/system/notify/config/template/update` | 更新通知绑定模版 | platform |
| POST | `/api/v1/system/notify/message/info/delete` | 删除消息 | platform |
| POST | `/api/v1/system/notify/message/info/get-list` | 获取消息列表 | platform |
| POST | `/api/v1/system/notify/message/info/send` | 发送消息 | platform |
| POST | `/api/v1/system/notify/message/info/update` | 更新消息 | platform |
| POST | `/api/v1/system/notify/news/create` | 创建资讯 | platform |
| POST | `/api/v1/system/notify/news/delete` | 删除资讯 | platform |
| POST | `/api/v1/system/notify/news/info` | 获取资讯详情 | platform |
| POST | `/api/v1/system/notify/news/list` | 获取资讯列表 | platform |
| POST | `/api/v1/system/notify/news/update` | 更新资讯 | platform |
| POST | `/api/v1/system/notify/notification/copy` | 复制手动通知 | platform |
| POST | `/api/v1/system/notify/notification/create` | 创建手动通知 | platform |
| POST | `/api/v1/system/notify/notification/delete` | 删除手动通知 | platform |
| POST | `/api/v1/system/notify/notification/estimate-users` | 预估送达用户数 | platform |
| POST | `/api/v1/system/notify/notification/index` | 获取手动通知列表 | platform |
| POST | `/api/v1/system/notify/notification/read` | 获取手动通知详情 | platform |
| POST | `/api/v1/system/notify/notification/revoke` | 撤回手动通知 | platform |
| POST | `/api/v1/system/notify/notification/send` | 发送手动通知 | platform |
| POST | `/api/v1/system/notify/notification/update` | 更新手动通知 | platform |
| POST | `/api/v1/system/notify/template/create` | 添加通知模版 | platform |
| POST | `/api/v1/system/notify/template/delete` | 删除通知模版 | platform |
| POST | `/api/v1/system/notify/template/get-list` | 获取通知模版列表 | platform |
| POST | `/api/v1/system/notify/template/get-one` | 获取通知模版详情 | platform |
| POST | `/api/v1/system/notify/template/update` | 更新通知模版 | platform |
| POST | `/api/v1/system/role/app/batch-update` | 更新APP权限 | admin |
| POST | `/api/v1/system/role/app/get-list` | 获取APP权限列表 | admin |
| POST | `/api/v1/system/role/info/create` | 添加角色 | admin |
| POST | `/api/v1/system/role/info/delete` | 删除角色 | admin |
| POST | `/api/v1/system/role/info/get-list` | 获取角色列表 | admin |
| POST | `/api/v1/system/role/info/update` | 更新角色 | admin |
| POST | `/api/v1/system/role/menu/batch-update` | 更新角色对应菜单列表 | admin |
| POST | `/api/v1/system/role/menu/get-list` | 获取角色对应菜单列表 | admin |
| POST | `/api/v1/system/role/resource/batch-update` | 批量更新角色资源动作权限 | admin |
| POST | `/api/v1/system/role/resource/get-list` | 获取角色资源动作权限列表 | admin |
| POST | `/api/v1/system/user/data/area/get-list` | 获取区域权限列表 | admin |
| POST | `/api/v1/system/user/data/project/get-list` | 获取项目权限列表 | admin |
| POST | `/api/v1/system/user/dept/batch-create` | 新增用户的部门列表 | admin |
| POST | `/api/v1/system/user/dept/batch-delete` | 删除用户的部门列表 | admin |
| POST | `/api/v1/system/user/info/create` | 创建用户信息 | admin |
| POST | `/api/v1/system/user/info/delete` | 刪除用户 | admin |
| POST | `/api/v1/system/user/info/get-list` | 查询用户信息列表 | admin |
| POST | `/api/v1/system/user/info/get-one` | 获取用户信息 | admin |
| POST | `/api/v1/system/user/info/update` | 更新用户基本数据 | admin |
| POST | `/api/v1/system/user/self/access-token/create` | 创建访问令牌 | all |
| POST | `/api/v1/system/user/self/access-token/delete` | 删除访问令牌 | all |
| POST | `/api/v1/system/user/self/access-token/get-list` | 获取访问令牌列表 | all |
| POST | `/api/v1/system/user/self/access-token/get-one` | 获取访问令牌详情 | all |
| POST | `/api/v1/system/user/self/access-token/update` | 更新访问令牌 | all |
| POST | `/api/v1/system/user/self/app/get-list` | 获取用户应用列表 | all |
| POST | `/api/v1/system/user/self/app/get-one` | 获取用户应用详情 | all |
| POST | `/api/v1/system/user/self/bind-account` | 绑定账号 | all |
| POST | `/api/v1/system/user/self/cancel` | 注销用户 | all |
| POST | `/api/v1/system/user/self/captcha` | 获取验证码 | public |
| POST | `/api/v1/system/user/self/change-pwd` | 更新用户密码 | all |
| POST | `/api/v1/system/user/self/forget-pwd` | 忘记密码 | public |
| POST | `/api/v1/system/user/self/get-one` | 获取用户信息 | all |
| POST | `/api/v1/system/user/self/login` | 用户登录 | public |
| POST | `/api/v1/system/user/self/logout` | 用户登出 | all |
| POST | `/api/v1/system/user/self/menu/get-list` | 获取用户菜单列表 | all |
| POST | `/api/v1/system/user/self/message/get-list` | 用户消息列表 | all |
| POST | `/api/v1/system/user/self/message/get-pending` | 用户待处理消息 | all |
| POST | `/api/v1/system/user/self/message/handle` | 用户消息标记已处理 | all |
| POST | `/api/v1/system/user/self/message/mark-all-read` | 用户消息全部已读 | all |
| POST | `/api/v1/system/user/self/message/multi-delete` | 用户消息批量删除 | all |
| POST | `/api/v1/system/user/self/message/multi-is-read` | 用户消息批量已读 | all |
| POST | `/api/v1/system/user/self/message/statistics` | 用户消息统计 | all |
| POST | `/api/v1/system/user/self/notify-preference/read` | 用户通知偏好读取 | all |
| POST | `/api/v1/system/user/self/notify-preference/update` | 用户通知偏好更新 | all |
| POST | `/api/v1/system/user/self/profile/get-list` | 获取用户配置列表 | all |
| POST | `/api/v1/system/user/self/profile/get-one` | 获取用户配置详情 | all |
| POST | `/api/v1/system/user/self/profile/update` | 更新用户配置 | all |
| POST | `/api/v1/system/user/self/register` | 普通用户注册 | public |
| POST | `/api/v1/system/user/self/resource/action/get-list` | 获取用户资源动作权限列表 | all |
| POST | `/api/v1/system/user/self/third-auth/start` | 第三方登录授权起跳 | public |
| POST | `/api/v1/system/user/self/third-login` | 第三方登录回调换平台登录态 | public |
| POST | `/api/v1/system/user/self/third-register` | 第三方补全注册 | public |
| POST | `/api/v1/system/user/self/update` | 更新用户基本数据 | all |
| POST | `/api/v1/system/user/self/user/search` | 精准搜索用户 | all |
| POST | `/api/v1/system/user/self/thirdparty/setup-check` | 查询 CLI 绑定状态 | public |
| POST | `/api/v1/system/user/self/thirdparty/setup-complete` | 完成 CLI 绑定 | all |
| POST | `/api/v1/system/user/self/tenant/delete` | 退出当前租户 | all |
| POST | `/api/v1/system/user/self/tenant/get-list` | 获取用户所处的租户列表 | all |
| POST | `/api/v1/system/user/self/tenant/get-one` | 获取当前用户在当前租户的详情 | all |
| POST | `/api/v1/system/user/self/tenant/join` | 用户加入租户（通过邀请码、邮件、手机或客户端应用） | all |
| POST | `/api/v1/system/user/self/tenant/update` | 更新当前用户在当前租户的信息 | all |
| POST | `/api/v1/system/user/tenant/get-list` | 用户所处的租户列表 | all |

## 端点详情

### POST `/api/v1/system/dept/info/create`

**说明**: 添加部门详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `children` | array[DeptInfo] | 否 |  |
| `desc` | string | 否 |  描述 |
| `dingTalkID` | string | 否 | 钉钉的部门ID |
| `id` | string | 否 |  编号 |
| `idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `name` | string | 否 |  部门名称 |
| `parent` | object | 否 |  |
| `parent.children` | array[DeptInfo] | 否 |  |
| `parent.desc` | string | 否 |  描述 |
| `parent.dingTalkID` | string | 否 | 钉钉的部门ID |
| `parent.id` | string | 否 |  编号 |
| `parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.name` | string | 否 |  部门名称 |
| `parent.parent` | object | 否 |  |
| `parent.parent.children` | array[DeptInfo] | 否 |  |
| `parent.parent.desc` | string | 否 |  描述 |
| `parent.parent.dingTalkID` | string | 否 | 钉钉的部门ID |
| `parent.parent.id` | string | 否 |  编号 |
| `parent.parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.parent.name` | string | 否 |  部门名称 |
| `parent.parent.parent` | object | 否 |  |
| `parent.parent.parentID` | string | 否 | 父节点 |
| `parent.parent.sort` | integer | 否 |  排序标记 (格式: int64) |
| `parent.parent.status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `parent.parent.userCount` | integer | 否 | 只读,用户统计,包含下级部门的人数 (格式: int64) |
| `parent.parentID` | string | 否 | 父节点 |
| `parent.sort` | integer | 否 |  排序标记 (格式: int64) |
| `parent.status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `parent.userCount` | integer | 否 | 只读,用户统计,包含下级部门的人数 (格式: int64) |
| `parentID` | string | 否 | 父节点 |
| `sort` | integer | 否 |  排序标记 (格式: int64) |
| `status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `userCount` | integer | 否 | 只读,用户统计,包含下级部门的人数 (格式: int64) |

**请求示例**:
```json
{
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
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dept/info/create \
  --body '{"children": [{"children": [{"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [{"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [{"children": [{"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [{"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [{"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}'
```

### POST `/api/v1/system/dept/info/delete`

**说明**: 删除部门

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
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
ur api /api/v1/system/dept/info/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/dept/info/get-list`

**说明**: 获取部门列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dingTalkIDs` | array[string] | 否 | 钉钉的部门ID |
| `name` | string | 否 |  名称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 | 父节点 |
| `status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `tenantCode` | string | 否 |  |

**请求示例**:
```json
{
  "dingTalkIDs": [
    "string"
  ],
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
  "parentID": "string",
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
ur api /api/v1/system/dept/info/get-list \
  --body '{"dingTalkIDs": ["string"], "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentID": "string", "status": 1, "tenantCode": "string"}'
```

### POST `/api/v1/system/dept/info/get-one`

**说明**: 获取部门详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  编号 |
| `tenantCode` | string | 否 |  |
| `withChildren` | boolean | 否 | 是否返回子级 (格式: boolean) |
| `withFather` | boolean | 否 | 是否返回父级 (格式: boolean) |

**请求示例**:
```json
{
  "id": "string",
  "tenantCode": "string",
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
ur api /api/v1/system/dept/info/get-one \
  --body '{"id": "string", "tenantCode": "string", "withChildren": true, "withFather": true}'
```

### POST `/api/v1/system/dept/info/update`

**说明**: 更新部门

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `children` | array[DeptInfo] | 否 |  |
| `desc` | string | 否 |  描述 |
| `dingTalkID` | string | 否 | 钉钉的部门ID |
| `id` | string | 否 |  编号 |
| `idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `name` | string | 否 |  部门名称 |
| `parent` | object | 否 |  |
| `parent.children` | array[DeptInfo] | 否 |  |
| `parent.desc` | string | 否 |  描述 |
| `parent.dingTalkID` | string | 否 | 钉钉的部门ID |
| `parent.id` | string | 否 |  编号 |
| `parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.name` | string | 否 |  部门名称 |
| `parent.parent` | object | 否 |  |
| `parent.parent.children` | array[DeptInfo] | 否 |  |
| `parent.parent.desc` | string | 否 |  描述 |
| `parent.parent.dingTalkID` | string | 否 | 钉钉的部门ID |
| `parent.parent.id` | string | 否 |  编号 |
| `parent.parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.parent.name` | string | 否 |  部门名称 |
| `parent.parent.parent` | object | 否 |  |
| `parent.parent.parentID` | string | 否 | 父节点 |
| `parent.parent.sort` | integer | 否 |  排序标记 (格式: int64) |
| `parent.parent.status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `parent.parent.userCount` | integer | 否 | 只读,用户统计,包含下级部门的人数 (格式: int64) |
| `parent.parentID` | string | 否 | 父节点 |
| `parent.sort` | integer | 否 |  排序标记 (格式: int64) |
| `parent.status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `parent.userCount` | integer | 否 | 只读,用户统计,包含下级部门的人数 (格式: int64) |
| `parentID` | string | 否 | 父节点 |
| `sort` | integer | 否 |  排序标记 (格式: int64) |
| `status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `userCount` | integer | 否 | 只读,用户统计,包含下级部门的人数 (格式: int64) |

**请求示例**:
```json
{
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
ur api /api/v1/system/dept/info/update \
  --body '{"children": [{"children": [{"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [{"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [{"children": [{"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [{"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [{"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": [], "desc": "string", "dingTalkID": "string", "id": "string", "idPath": "string", "name": "示例名称", "parent": {"children": "...", "desc": "...", "dingTalkID": "...", "id": "...", "idPath": "...", "name": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "userCount": "..."}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}, "parentID": "string", "sort": 1, "status": 1, "userCount": 1}'
```

### POST `/api/v1/system/dept/role/batch-create`

**说明**: 批量添加部门角色

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 是 |  部门ID |
| `roleIDs` | array[string] | 是 |  角色ID列表 |

**请求示例**:
```json
{
  "deptID": "string",
  "roleIDs": [
    "string"
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
ur api /api/v1/system/dept/role/batch-create \
  --body '{"deptID": "string", "roleIDs": ["string"]}'
```

### POST `/api/v1/system/dept/role/batch-delete`

**说明**: 批量移除部门角色

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 是 |  部门ID |
| `roleIDs` | array[string] | 是 |  角色ID列表 |

**请求示例**:
```json
{
  "deptID": "string",
  "roleIDs": [
    "string"
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
ur api /api/v1/system/dept/role/batch-delete \
  --body '{"deptID": "string", "roleIDs": ["string"]}'
```

### POST `/api/v1/system/dept/role/get-list`

**说明**: 获取部门角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 是 |  部门ID |

**请求示例**:
```json
{
  "deptID": "string"
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
        "isFullAuth": 1,
        "name": "示例名称",
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
ur api /api/v1/system/dept/role/get-list \
  --body '{"deptID": "string"}'
```

### POST `/api/v1/system/dept/sync-job/create`

**说明**: 添加同步任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间 |
| `direction` | integer | 否 |  同步的方向,1上游同步到联犀(默认),2联犀同步到下游 (格式: int64) |
| `fieldMap` | object | 否 | 用户字段映射,左边是联犀的字段,右边是第三方的,不填写就是全量映射 |
| `id` | string | 否 |  编号 |
| `isAddSync` | integer | 否 | 新增人员自动同步,默认为1 (格式: int64) |
| `syncDeptIDs` | array[string] | 否 | 同步的第三方部门id列表,不填为同步全部 |
| `syncMode` | integer | 否 | 同步模式: 1:手动(默认) 2: 定时同步(半小时) 3: 实时同步 (格式: int64) |
| `thirdConfig` | object | 否 |  |
| `thirdConfig.appID` | string | 否 |  |
| `thirdConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `thirdConfig.appSecret` | string | 否 |  |
| `thirdType` | string | 否 | 同步的类型 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "direction": 1,
  "fieldMap": {},
  "id": "string",
  "isAddSync": 1,
  "syncDeptIDs": [
    "string"
  ],
  "syncMode": 1,
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
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dept/sync-job/create \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "direction": 1, "fieldMap": {}, "id": "string", "isAddSync": 1, "syncDeptIDs": ["string"], "syncMode": 1, "thirdConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "thirdType": "string"}'
```

### POST `/api/v1/system/dept/sync-job/delete`

**说明**: 删除同步任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |

**请求示例**:
```json
{
  "id": "string"
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
ur api /api/v1/system/dept/sync-job/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/dept/sync-job/execute`

**说明**: 执行同步任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `jobID` | string | 否 |  编号 |

**请求示例**:
```json
{
  "jobID": "string"
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
ur api /api/v1/system/dept/sync-job/execute \
  --body '{"jobID": "string"}'
```

### POST `/api/v1/system/dept/sync-job/get-list`

**说明**: 获取同步任务列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `direction` | integer | 否 |  同步的方向,1上游同步到联犀(默认),2联犀同步到下游 (格式: int64) |
| `name` | string | 否 |  名称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 | 父节点 |
| `status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |

**请求示例**:
```json
{
  "direction": 1,
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
  "parentID": "string",
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
        "createdTime": "2026-01-01T00:00:00Z",
        "direction": 1,
        "fieldMap": {},
        "id": "string",
        "isAddSync": 1,
        "syncDeptIDs": [
          "string"
        ],
        "syncMode": 1,
        "thirdConfig": {
          "appID": "string",
          "appKey": "string",
          "appSecret": "string"
        },
        "thirdType": "string"
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
ur api /api/v1/system/dept/sync-job/get-list \
  --body '{"direction": 1, "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentID": "string", "status": 1}'
```

### POST `/api/v1/system/dept/sync-job/get-one`

**说明**: 获取同步任务详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  编号 |
| `withChildren` | boolean | 否 | 是否返回子级 (格式: boolean) |
| `withFather` | boolean | 否 | 是否返回父级 (格式: boolean) |

**请求示例**:
```json
{
  "id": "string",
  "withChildren": true,
  "withFather": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdTime": "2026-01-01T00:00:00Z",
    "direction": 1,
    "fieldMap": {},
    "id": "string",
    "isAddSync": 1,
    "syncDeptIDs": [
      "string"
    ],
    "syncMode": 1,
    "thirdConfig": {
      "appID": "string",
      "appKey": "string",
      "appSecret": "string"
    },
    "thirdType": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dept/sync-job/get-one \
  --body '{"id": "string", "withChildren": true, "withFather": true}'
```

### POST `/api/v1/system/dept/sync-job/update`

**说明**: 更新同步任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间 |
| `direction` | integer | 否 |  同步的方向,1上游同步到联犀(默认),2联犀同步到下游 (格式: int64) |
| `fieldMap` | object | 否 | 用户字段映射,左边是联犀的字段,右边是第三方的,不填写就是全量映射 |
| `id` | string | 否 |  编号 |
| `isAddSync` | integer | 否 | 新增人员自动同步,默认为1 (格式: int64) |
| `syncDeptIDs` | array[string] | 否 | 同步的第三方部门id列表,不填为同步全部 |
| `syncMode` | integer | 否 | 同步模式: 1:手动(默认) 2: 定时同步(半小时) 3: 实时同步 (格式: int64) |
| `thirdConfig` | object | 否 |  |
| `thirdConfig.appID` | string | 否 |  |
| `thirdConfig.appKey` | string | 否 | 微信小程序无需填写 |
| `thirdConfig.appSecret` | string | 否 |  |
| `thirdType` | string | 否 | 同步的类型 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "direction": 1,
  "fieldMap": {},
  "id": "string",
  "isAddSync": 1,
  "syncDeptIDs": [
    "string"
  ],
  "syncMode": 1,
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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dept/sync-job/update \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "direction": 1, "fieldMap": {}, "id": "string", "isAddSync": 1, "syncDeptIDs": ["string"], "syncMode": 1, "thirdConfig": {"appID": "string", "appKey": "string", "appSecret": "string"}, "thirdType": "string"}'
```

### POST `/api/v1/system/dept/user/batch-create`

**说明**: 批量授权部门用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authType` | integer | 是 | 授权类型: 1 管理权限,可以修改该部门及修改部门人员的授权 2 普通权限 (格式: int64) |
| `deptID` | string | 否 |  |
| `isAuthChildren` | integer | 是 | 同时授权子节点 (格式: int64) |
| `userIDs` | array[string] | 是 |  |

**请求示例**:
```json
{
  "authType": 1,
  "deptID": "string",
  "isAuthChildren": 1,
  "userIDs": [
    "string"
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
ur api /api/v1/system/dept/user/batch-create \
  --body '{"authType": 1, "deptID": "string", "isAuthChildren": 1, "userIDs": ["string"]}'
```

### POST `/api/v1/system/dept/user/batch-delete`

**说明**: 批量取消授权部门用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 是 |  |
| `ids` | array[string] | 是 |  |

**请求示例**:
```json
{
  "deptID": "string",
  "ids": [
    "string"
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
ur api /api/v1/system/dept/user/batch-delete \
  --body '{"deptID": "string", "ids": ["string"]}'
```

### POST `/api/v1/system/dept/user/get-list`

**说明**: 获取部门授权列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptID` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "deptID": "string",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
        "authType": 1,
        "deptID": "string",
        "deptIdPath": "string",
        "id": "string",
        "isAuthChildren": 1,
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
        },
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
ur api /api/v1/system/dept/user/get-list \
  --body '{"deptID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/dict/detail/create`

**说明**: 添加字典详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `children` | array[DictDetail] | 否 |  |
| `desc` | string | 否 |  模块描述 |
| `dictCode` | string | 是 |  |
| `expand` | object | 否 |  自定义数据 |
| `i18nLabel` | string | 否 |  按当前请求语言翻译后的展示值 |
| `id` | string | 否 |  编号 |
| `idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `label` | string | 否 |  展示值（编辑场景使用） |
| `parent` | object | 否 |  |
| `parent.children` | array[DictDetail] | 否 |  |
| `parent.desc` | string | 否 |  模块描述 |
| `parent.dictCode` | string | 是 |  |
| `parent.expand` | object | 否 |  自定义数据 |
| `parent.i18nLabel` | string | 否 |  按当前请求语言翻译后的展示值 |
| `parent.id` | string | 否 |  编号 |
| `parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `parent.label` | string | 否 |  展示值（编辑场景使用） |
| `parent.parent` | object | 否 |  |
| `parent.parent.children` | array[DictDetail] | 否 |  |
| `parent.parent.desc` | string | 否 |  模块描述 |
| `parent.parent.dictCode` | string | 是 |  |
| `parent.parent.expand` | object | 否 |  自定义数据 |
| `parent.parent.i18nLabel` | string | 否 |  按当前请求语言翻译后的展示值 |
| `parent.parent.id` | string | 否 |  编号 |
| `parent.parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.parent.isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `parent.parent.label` | string | 否 |  展示值（编辑场景使用） |
| `parent.parent.parent` | object | 否 |  |
| `parent.parent.parentID` | string | 否 | 父节点 |
| `parent.parent.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `parent.parent.status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `parent.parent.value` | string | 否 |  字典值 |
| `parent.parentID` | string | 否 | 父节点 |
| `parent.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `parent.status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `parent.value` | string | 否 |  字典值 |
| `parentID` | string | 否 | 父节点 |
| `sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `value` | string | 否 |  字典值 |

**请求示例**:
```json
{
  "children": [
    {
      "children": [
        {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        }
      ],
      "desc": "string",
      "dictCode": "string",
      "expand": {},
      "i18nLabel": "string",
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "label": "string",
      "parent": {
        "children": [
          {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "value": "string"
    }
  ],
  "desc": "string",
  "dictCode": "string",
  "expand": {},
  "i18nLabel": "string",
  "id": "string",
  "idPath": "string",
  "isLeaf": 1,
  "label": "string",
  "parent": {
    "children": [
      {
        "children": [
          {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      }
    ],
    "desc": "string",
    "dictCode": "string",
    "expand": {},
    "i18nLabel": "string",
    "id": "string",
    "idPath": "string",
    "isLeaf": 1,
    "label": "string",
    "parent": {
      "children": [
        {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        }
      ],
      "desc": "string",
      "dictCode": "string",
      "expand": {},
      "i18nLabel": "string",
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "label": "string",
      "parent": {
        "children": [
          {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "value": "string"
    },
    "parentID": "string",
    "sort": 1,
    "status": 1,
    "value": "string"
  },
  "parentID": "string",
  "sort": 1,
  "status": 1,
  "value": "string"
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
ur api /api/v1/system/dict/detail/create \
  --body '{"children": [{"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}'
```

### POST `/api/v1/system/dict/detail/delete`

**说明**: 删除字典详情

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
ur api /api/v1/system/dict/detail/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/dict/detail/get-list`

**说明**: 获取字典详情列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dictCode` | string | 是 |  |
| `label` | string | 否 |  展示值 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 | 父节点 |
| `status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `value` | string | 否 |  字典值 |
| `values` | array[string] | 否 | 字典值过滤 |

**请求示例**:
```json
{
  "dictCode": "string",
  "label": "string",
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
  "status": 1,
  "value": "string",
  "values": [
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
        "children": [
          {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            }
          ],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
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
ur api /api/v1/system/dict/detail/get-list \
  --body '{"dictCode": "string", "label": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentID": "string", "status": 1, "value": "string", "values": ["string"]}'
```

### POST `/api/v1/system/dict/detail/get-one`

**说明**: 获取字典详情单个

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dictCode` | string | 是 |  |
| `id` | string | 否 |  编号 |
| `value` | string | 否 |  字典值 |
| `withChildren` | boolean | 否 | 是否返回子级 (格式: boolean) |
| `withFather` | boolean | 否 | 是否返回父级 (格式: boolean) |

**请求示例**:
```json
{
  "dictCode": "string",
  "id": "string",
  "value": "string",
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
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            }
          ],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      }
    ],
    "desc": "string",
    "dictCode": "string",
    "expand": {},
    "i18nLabel": "string",
    "id": "string",
    "idPath": "string",
    "isLeaf": 1,
    "label": "string",
    "parent": {
      "children": [
        {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            }
          ],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        }
      ],
      "desc": "string",
      "dictCode": "string",
      "expand": {},
      "i18nLabel": "string",
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "label": "string",
      "parent": {
        "children": [
          {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            }
          ],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "value": "string"
    },
    "parentID": "string",
    "sort": 1,
    "status": 1,
    "value": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dict/detail/get-one \
  --body '{"dictCode": "string", "id": "string", "value": "string", "withChildren": true, "withFather": true}'
```

### POST `/api/v1/system/dict/detail/update`

**说明**: 更新字典详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `children` | array[DictDetail] | 否 |  |
| `desc` | string | 否 |  模块描述 |
| `dictCode` | string | 是 |  |
| `expand` | object | 否 |  自定义数据 |
| `i18nLabel` | string | 否 |  按当前请求语言翻译后的展示值 |
| `id` | string | 否 |  编号 |
| `idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `label` | string | 否 |  展示值（编辑场景使用） |
| `parent` | object | 否 |  |
| `parent.children` | array[DictDetail] | 否 |  |
| `parent.desc` | string | 否 |  模块描述 |
| `parent.dictCode` | string | 是 |  |
| `parent.expand` | object | 否 |  自定义数据 |
| `parent.i18nLabel` | string | 否 |  按当前请求语言翻译后的展示值 |
| `parent.id` | string | 否 |  编号 |
| `parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `parent.label` | string | 否 |  展示值（编辑场景使用） |
| `parent.parent` | object | 否 |  |
| `parent.parent.children` | array[DictDetail] | 否 |  |
| `parent.parent.desc` | string | 否 |  模块描述 |
| `parent.parent.dictCode` | string | 是 |  |
| `parent.parent.expand` | object | 否 |  自定义数据 |
| `parent.parent.i18nLabel` | string | 否 |  按当前请求语言翻译后的展示值 |
| `parent.parent.id` | string | 否 |  编号 |
| `parent.parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.parent.isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `parent.parent.label` | string | 否 |  展示值（编辑场景使用） |
| `parent.parent.parent` | object | 否 |  |
| `parent.parent.parentID` | string | 否 | 父节点 |
| `parent.parent.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `parent.parent.status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `parent.parent.value` | string | 否 |  字典值 |
| `parent.parentID` | string | 否 | 父节点 |
| `parent.sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `parent.status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `parent.value` | string | 否 |  字典值 |
| `parentID` | string | 否 | 父节点 |
| `sort` | integer | 否 |  排序标记,默认为1 (格式: int64) |
| `status` | integer | 否 |  状态  1:启用,2:禁用 (格式: int64) |
| `value` | string | 否 |  字典值 |

**请求示例**:
```json
{
  "children": [
    {
      "children": [
        {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        }
      ],
      "desc": "string",
      "dictCode": "string",
      "expand": {},
      "i18nLabel": "string",
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "label": "string",
      "parent": {
        "children": [
          {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "value": "string"
    }
  ],
  "desc": "string",
  "dictCode": "string",
  "expand": {},
  "i18nLabel": "string",
  "id": "string",
  "idPath": "string",
  "isLeaf": 1,
  "label": "string",
  "parent": {
    "children": [
      {
        "children": [
          {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      }
    ],
    "desc": "string",
    "dictCode": "string",
    "expand": {},
    "i18nLabel": "string",
    "id": "string",
    "idPath": "string",
    "isLeaf": 1,
    "label": "string",
    "parent": {
      "children": [
        {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        }
      ],
      "desc": "string",
      "dictCode": "string",
      "expand": {},
      "i18nLabel": "string",
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "label": "string",
      "parent": {
        "children": [
          {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "value": "string"
    },
    "parentID": "string",
    "sort": 1,
    "status": 1,
    "value": "string"
  },
  "parentID": "string",
  "sort": 1,
  "status": 1,
  "value": "string"
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
ur api /api/v1/system/dict/detail/update \
  --body '{"children": [{"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}'
```

### POST `/api/v1/system/dict/info/batch-export`

**说明**: 批量导出字典信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dictCodes` | array[string] | 否 |  |

**请求示例**:
```json
{
  "dictCodes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "dicts": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dict/info/batch-export \
  --body '{"dictCodes": ["string"]}'
```

### POST `/api/v1/system/dict/info/batch-import`

**说明**: 批量导入字典信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `dicts` | string | 是 |  |

**请求示例**:
```json
{
  "dicts": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "errCount": 1,
    "ignoreCount": 1,
    "succCount": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dict/info/batch-import \
  --body '{"dicts": "string"}'
```

### POST `/api/v1/system/dict/info/create`

**说明**: 添加字典信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  编码 |
| `desc` | string | 否 |  页面 |
| `details` | array[DictDetail] | 否 |  withDetail 时返回 |
| `expand` | object | 否 | 前端自定义字段 |
| `group` | string | 否 |  分组 |
| `id` | string | 否 |  编号 |
| `name` | string | 否 |  名称 |
| `structType` | integer | 否 | 结构类型(不可修改) 1:列表(默认) 2:树型 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "details": [
    {
      "children": [
        {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        }
      ],
      "desc": "string",
      "dictCode": "string",
      "expand": {},
      "i18nLabel": "string",
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "label": "string",
      "parent": {
        "children": [
          {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "value": "string"
    }
  ],
  "expand": {},
  "group": "string",
  "id": "string",
  "name": "示例名称",
  "structType": 1
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
ur api /api/v1/system/dict/info/create \
  --body '{"code": "string", "desc": "string", "details": [{"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "expand": {}, "group": "string", "id": "string", "name": "示例名称", "structType": 1}'
```

### POST `/api/v1/system/dict/info/delete`

**说明**: 删除字典信息

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
ur api /api/v1/system/dict/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/dict/info/get-list`

**说明**: 获取字典信息列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `group` | string | 否 |  |
| `name` | string | 否 |  名称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 格式: int64 |
| `withDetail` | boolean | 否 |  true 时每条 DictInfo 附带 details（仅当前页字典分类） (格式: boolean) |

**请求示例**:
```json
{
  "group": "string",
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
  "withDetail": true
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
        "details": [
          {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          }
        ],
        "expand": {},
        "group": "string",
        "id": "string",
        "name": "示例名称",
        "structType": 1
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
ur api /api/v1/system/dict/info/get-list \
  --body '{"group": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "withDetail": true}'
```

### POST `/api/v1/system/dict/info/get-one`

**说明**: 获取字典信息详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  编码 |
| `id` | string | 是 |  编号 |

**请求示例**:
```json
{
  "code": "string",
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "desc": "string",
    "details": [
      {
        "children": [
          {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [
            {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            }
          ],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "i18nLabel": "string",
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
              "i18nLabel": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "label": "...",
              "parent": "...",
              "parentID": "...",
              "sort": "...",
              "status": "...",
              "value": "..."
            },
            "parentID": "string",
            "sort": 1,
            "status": 1,
            "value": "string"
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      }
    ],
    "expand": {},
    "group": "string",
    "id": "string",
    "name": "示例名称",
    "structType": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dict/info/get-one \
  --body '{"code": "string", "id": "string"}'
```

### POST `/api/v1/system/dict/info/update`

**说明**: 更新字典信息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  编码 |
| `desc` | string | 否 |  页面 |
| `details` | array[DictDetail] | 否 |  withDetail 时返回 |
| `expand` | object | 否 | 前端自定义字段 |
| `group` | string | 否 |  分组 |
| `id` | string | 否 |  编号 |
| `name` | string | 否 |  名称 |
| `structType` | integer | 否 | 结构类型(不可修改) 1:列表(默认) 2:树型 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "details": [
    {
      "children": [
        {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        }
      ],
      "desc": "string",
      "dictCode": "string",
      "expand": {},
      "i18nLabel": "string",
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "label": "string",
      "parent": {
        "children": [
          {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          }
        ],
        "desc": "string",
        "dictCode": "string",
        "expand": {},
        "i18nLabel": "string",
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "i18nLabel": "string",
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
            "i18nLabel": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "label": "...",
            "parent": "...",
            "parentID": "...",
            "sort": "...",
            "status": "...",
            "value": "..."
          },
          "parentID": "string",
          "sort": 1,
          "status": 1,
          "value": "string"
        },
        "parentID": "string",
        "sort": 1,
        "status": 1,
        "value": "string"
      },
      "parentID": "string",
      "sort": 1,
      "status": 1,
      "value": "string"
    }
  ],
  "expand": {},
  "group": "string",
  "id": "string",
  "name": "示例名称",
  "structType": 1
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
ur api /api/v1/system/dict/info/update \
  --body '{"code": "string", "desc": "string", "details": [{"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "i18nLabel": "string", "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "i18nLabel": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "expand": {}, "group": "string", "id": "string", "name": "示例名称", "structType": 1}'
```

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
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
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
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
  --body '{"appID": "string", "code": "string", "dateRange": {"end": "2026-01-01T00:00:00Z", "start": "2026-01-01T00:00:00Z"}, "ipAddr": "string", "loginLocation": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "userID": "string", "userName": "string"}'
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
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

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
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
  --body '{"accessKey": "string", "appID": "string", "authType": "string", "code": "string", "operName": "示例名称", "operType": "string", "operUserID": "string", "operUserName": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/notify/config/create`

**说明**: 添加通知配置

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  通知类型编码 |
| `desc` | string | 否 |  备注 |
| `group` | string | 否 | 分组 |
| `id` | string | 否 |  id编号 |
| `isEnabled` | integer | 否 | 是否启用 1:启用 2:禁用 (格式: int64) |
| `isRecord` | integer | 否 | 是否记录该消息,是的情况下会将消息存一份到消息中心 (格式: int64) |
| `name` | string | 否 | 通知的命名 |
| `params` | object | 否 | 变量属性 key是变量参数,value是变量描述 |
| `supportTypes` | array[string] | 否 | 支持的通知类型 sms(短信),email(邮箱),dingTalk(钉钉机器人),dingWebhook(钉钉webhook),wxEWebHook(企业微信webhook) |
| `templates` | array[NotifyConfigTemplate] | 否 | 传withTemplates的时候返回,只读 |
| `throttle` | integer | 否 | 防抖间隔秒数,0表示不防抖 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "group": "string",
  "id": "string",
  "isEnabled": 1,
  "isRecord": 1,
  "name": "string",
  "params": {},
  "supportTypes": [
    "string"
  ],
  "templates": [
    {
      "id": "string",
      "notifyCode": "string",
      "templateID": "string",
      "type": "string"
    }
  ],
  "throttle": 1
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
ur api /api/v1/system/notify/config/create \
  --body '{"code": "string", "desc": "string", "group": "string", "id": "string", "isEnabled": 1, "isRecord": 1, "name": "string", "params": {}, "supportTypes": ["string"], "templates": [{"id": "string", "notifyCode": "string", "templateID": "string", "type": "string"}], "throttle": 1}'
```

### POST `/api/v1/system/notify/config/delete`

**说明**: 删除通知配置

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
ur api /api/v1/system/notify/config/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/config/get-list`

**说明**: 获取通知列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  应用编号 |
| `group` | string | 否 | 分组 |
| `isEnabled` | integer | 否 | 是否启用 1:启用 2:禁用 (格式: int64) |
| `name` | string | 否 |  应用名称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `withTemplates` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "code": "string",
  "group": "string",
  "isEnabled": 1,
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
  "withTemplates": true
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
        "group": "string",
        "id": "string",
        "isEnabled": 1,
        "isRecord": 1,
        "name": "string",
        "params": {},
        "supportTypes": [
          "string"
        ],
        "templates": [
          {
            "id": "string",
            "notifyCode": "string",
            "templateID": "string",
            "type": "string"
          }
        ],
        "throttle": 1
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
ur api /api/v1/system/notify/config/get-list \
  --body '{"code": "string", "group": "string", "isEnabled": 1, "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "withTemplates": true}'
```

### POST `/api/v1/system/notify/config/get-one`

**说明**: 获取通知详情

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | integer | 否 |  id (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "desc": "string",
    "group": "string",
    "id": "string",
    "isEnabled": 1,
    "isRecord": 1,
    "name": "string",
    "params": {},
    "supportTypes": [
      "string"
    ],
    "templates": [
      {
        "id": "string",
        "notifyCode": "string",
        "templateID": "string",
        "type": "string"
      }
    ],
    "throttle": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/config/get-one \
  --body '{"code": "string", "id": 1}'
```

### POST `/api/v1/system/notify/config/send-test`

**说明**: 测试发送通知

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accounts` | array[string] | 否 | 指定账号列表 |
| `notifyCode` | string | 是 | 通知编码 |
| `params` | object | 否 | 模版变量 |
| `str1` | string | 否 |  |
| `str2` | string | 否 |  |
| `str3` | string | 否 |  |
| `templateID` | string | 否 | 指定模版ID |
| `type` | string | 是 | 渠道类型 sms\|email\|dingTalk\|message等 |

**请求示例**:
```json
{
  "accounts": [
    "string"
  ],
  "notifyCode": "string",
  "params": {},
  "str1": "string",
  "str2": "string",
  "str3": "string",
  "templateID": "string",
  "type": "string"
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
ur api /api/v1/system/notify/config/send-test \
  --body '{"accounts": ["string"], "notifyCode": "string", "params": {}, "str1": "string", "str2": "string", "str3": "string", "templateID": "string", "type": "string"}'
```

### POST `/api/v1/system/notify/config/tree`

**说明**: 获取通知列表树

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  应用编号 |
| `group` | string | 否 | 分组 |
| `isEnabled` | integer | 否 | 是否启用 1:启用 2:禁用 (格式: int64) |
| `name` | string | 否 |  应用名称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `withTemplates` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "code": "string",
  "group": "string",
  "isEnabled": 1,
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
  "withTemplates": true
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
            "code": "string",
            "desc": "string",
            "group": "string",
            "id": "string",
            "isEnabled": 1,
            "isRecord": 1,
            "name": "string",
            "params": {},
            "supportTypes": [],
            "templates": [],
            "throttle": 1
          }
        ],
        "code": "string",
        "id": "string",
        "name": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/config/tree \
  --body '{"code": "string", "group": "string", "isEnabled": 1, "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "withTemplates": true}'
```

### POST `/api/v1/system/notify/config/update`

**说明**: 更新通知配置

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  通知类型编码 |
| `desc` | string | 否 |  备注 |
| `group` | string | 否 | 分组 |
| `id` | string | 否 |  id编号 |
| `isEnabled` | integer | 否 | 是否启用 1:启用 2:禁用 (格式: int64) |
| `isRecord` | integer | 否 | 是否记录该消息,是的情况下会将消息存一份到消息中心 (格式: int64) |
| `name` | string | 否 | 通知的命名 |
| `params` | object | 否 | 变量属性 key是变量参数,value是变量描述 |
| `supportTypes` | array[string] | 否 | 支持的通知类型 sms(短信),email(邮箱),dingTalk(钉钉机器人),dingWebhook(钉钉webhook),wxEWebHook(企业微信webhook) |
| `templates` | array[NotifyConfigTemplate] | 否 | 传withTemplates的时候返回,只读 |
| `throttle` | integer | 否 | 防抖间隔秒数,0表示不防抖 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "desc": "string",
  "group": "string",
  "id": "string",
  "isEnabled": 1,
  "isRecord": 1,
  "name": "string",
  "params": {},
  "supportTypes": [
    "string"
  ],
  "templates": [
    {
      "id": "string",
      "notifyCode": "string",
      "templateID": "string",
      "type": "string"
    }
  ],
  "throttle": 1
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
ur api /api/v1/system/notify/config/update \
  --body '{"code": "string", "desc": "string", "group": "string", "id": "string", "isEnabled": 1, "isRecord": 1, "name": "string", "params": {}, "supportTypes": ["string"], "templates": [{"id": "string", "notifyCode": "string", "templateID": "string", "type": "string"}], "throttle": 1}'
```

### POST `/api/v1/system/notify/config/template/delete`

**说明**: 删除通知绑定模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `notifyCode` | string | 是 |  |
| `type` | string | 是 |  |

**请求示例**:
```json
{
  "notifyCode": "string",
  "type": "string"
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
ur api /api/v1/system/notify/config/template/delete \
  --body '{"notifyCode": "string", "type": "string"}'
```

### POST `/api/v1/system/notify/config/template/get-list`

**说明**: 获取通知绑定模版列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `notifyCode` | string | 是 |  |
| `type` | string | 否 |  |

**请求示例**:
```json
{
  "notifyCode": "string",
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
        "id": "string",
        "notifyCode": "string",
        "templateID": "string",
        "type": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/config/template/get-list \
  --body '{"notifyCode": "string", "type": "string"}'
```

### POST `/api/v1/system/notify/config/template/update`

**说明**: 更新通知绑定模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  |
| `notifyCode` | string | 否 |  |
| `templateID` | string | 否 |  |
| `type` | string | 否 |  |

**请求示例**:
```json
{
  "id": "string",
  "notifyCode": "string",
  "templateID": "string",
  "type": "string"
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
ur api /api/v1/system/notify/config/template/update \
  --body '{"id": "string", "notifyCode": "string", "templateID": "string", "type": "string"}'
```

### POST `/api/v1/system/notify/message/info/delete`

**说明**: 删除消息

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
ur api /api/v1/system/notify/message/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/message/info/get-list`

**说明**: 获取消息列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `group` | string | 否 | 分组 |
| `isGlobal` | integer | 否 | 全局消息 (格式: int64) |
| `notificationID` | string | 否 | 关联手动通知ID |
| `notifyCode` | string | 否 | 通知编码 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "group": "string",
  "isGlobal": 1,
  "notificationID": "string",
  "notifyCode": "string",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
        "actionText": "string",
        "actionURL": "string",
        "allowDismiss": "string",
        "body": "string",
        "coverImage": "string",
        "createdTime": "string",
        "emailFormat": "string",
        "expireTime": 1,
        "group": "string",
        "id": "string",
        "isGlobal": 1,
        "modalShowEndTime": "2026-01-01T00:00:00Z",
        "modalShowStartTime": "2026-01-01T00:00:00Z",
        "notificationID": "string",
        "notifyCode": "string",
        "notifyName": "示例名称",
        "notifyTime": 1,
        "priority": 1,
        "severity": "string",
        "showAsModal": "string",
        "str1": "string",
        "str2": "string",
        "str3": "string",
        "subject": "string",
        "type": "string",
        "videoUrl": "string"
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
ur api /api/v1/system/notify/message/info/get-list \
  --body '{"group": "string", "isGlobal": 1, "notificationID": "string", "notifyCode": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/notify/message/info/send`

**说明**: 发送消息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `body` | string | 是 |  |
| `isGlobal` | integer | 是 | 全局消息 (格式: int64) |
| `notifyCode` | string | 是 | 通知的code |
| `notifyTime` | integer | 否 | 通知时间,不填立刻发送 (格式: int64) |
| `str1` | string | 否 |  |
| `str2` | string | 否 |  |
| `str3` | string | 否 |  |
| `subject` | string | 是 | 通知主题 |
| `userIDs` | array[string] | 否 | 指定用户ID |

**请求示例**:
```json
{
  "body": "string",
  "isGlobal": 1,
  "notifyCode": "string",
  "notifyTime": 1,
  "str1": "string",
  "str2": "string",
  "str3": "string",
  "subject": "string",
  "userIDs": [
    "string"
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
ur api /api/v1/system/notify/message/info/send \
  --body '{"body": "string", "isGlobal": 1, "notifyCode": "string", "notifyTime": 1, "str1": "string", "str2": "string", "str3": "string", "subject": "string", "userIDs": ["string"]}'
```

### POST `/api/v1/system/notify/message/info/update`

**说明**: 更新消息

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionText` | string | 否 | 操作按钮文案 |
| `actionURL` | string | 否 | 操作跳转URL |
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `body` | string | 是 | 内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `createdTime` | string | 是 |  |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `expireTime` | integer | 否 | 有效期(毫秒时间戳) (格式: int64) |
| `group` | string | 是 | 分组 |
| `id` | string | 是 |  |
| `isGlobal` | integer | 是 | 格式: int64 |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `notificationID` | string | 否 | 关联手动通知ID |
| `notifyCode` | string | 是 |  |
| `notifyName` | string | 否 | 通知的名称,只读 |
| `notifyTime` | integer | 是 | 通知时间 (格式: int64) |
| `priority` | integer | 否 | 优先级 1:紧急 2:重要 3:普通 (格式: int64) |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `str1` | string | 是 |  |
| `str2` | string | 是 |  |
| `str3` | string | 是 |  |
| `subject` | string | 是 | 主题 |
| `type` | string | 否 | 渠道类型 sms\|email\|dingTalk\|message等 |
| `videoUrl` | string | 否 |  news专属：视频URL |

**请求示例**:
```json
{
  "actionText": "string",
  "actionURL": "string",
  "allowDismiss": "string",
  "body": "string",
  "coverImage": "string",
  "createdTime": "string",
  "emailFormat": "string",
  "expireTime": 1,
  "group": "string",
  "id": "string",
  "isGlobal": 1,
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "notificationID": "string",
  "notifyCode": "string",
  "notifyName": "示例名称",
  "notifyTime": 1,
  "priority": 1,
  "severity": "string",
  "showAsModal": "string",
  "str1": "string",
  "str2": "string",
  "str3": "string",
  "subject": "string",
  "type": "string",
  "videoUrl": "string"
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
ur api /api/v1/system/notify/message/info/update \
  --body '{"actionText": "string", "actionURL": "string", "allowDismiss": "string", "body": "string", "coverImage": "string", "createdTime": "string", "emailFormat": "string", "expireTime": 1, "group": "string", "id": "string", "isGlobal": 1, "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "notificationID": "string", "notifyCode": "string", "notifyName": "示例名称", "notifyTime": 1, "priority": 1, "severity": "string", "showAsModal": "string", "str1": "string", "str2": "string", "str3": "string", "subject": "string", "type": "string", "videoUrl": "string"}'
```

### POST `/api/v1/system/notify/news/create`

**说明**: 创建资讯

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionURL` | string | 否 |  跳转链接 |
| `body` | string | 是 |  正文内容（HTML格式） |
| `coverImage` | string | 否 |  封面图URL |
| `createdTime` | integer | 否 |  创建时间，只读 (格式: int64) |
| `expireTime` | string | 否 |  过期时间（毫秒时间戳） |
| `id` | string | 否 |  资讯ID |
| `notifyTime` | string | 否 |  发布时间（毫秒时间戳） |
| `status` | string | 否 |  状态 1:草稿 2:已发布 3:已下架 |
| `summary` | string | 否 |  摘要 |
| `title` | string | 是 |  标题 |

**请求示例**:
```json
{
  "actionURL": "string",
  "body": "string",
  "coverImage": "string",
  "createdTime": 1,
  "expireTime": "2026-01-01T00:00:00Z",
  "id": "string",
  "notifyTime": "2026-01-01T00:00:00Z",
  "status": "string",
  "summary": "string",
  "title": "string"
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
ur api /api/v1/system/notify/news/create \
  --body '{"actionURL": "string", "body": "string", "coverImage": "string", "createdTime": 1, "expireTime": "2026-01-01T00:00:00Z", "id": "string", "notifyTime": "2026-01-01T00:00:00Z", "status": "string", "summary": "string", "title": "string"}'
```

### POST `/api/v1/system/notify/news/delete`

**说明**: 删除资讯

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
ur api /api/v1/system/notify/news/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/news/info`

**说明**: 获取资讯详情

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
    "actionURL": "string",
    "body": "string",
    "coverImage": "string",
    "createdTime": 1,
    "expireTime": "2026-01-01T00:00:00Z",
    "id": "string",
    "notifyTime": "2026-01-01T00:00:00Z",
    "status": "string",
    "summary": "string",
    "title": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/news/info \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/news/list`

**说明**: 获取资讯列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `notifyTimeEnd` | integer | 否 |  发布时间结束 (格式: int64) |
| `notifyTimeStart` | integer | 否 |  发布时间开始 (格式: int64) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | string | 否 |  状态筛选 |
| `title` | string | 否 |  标题模糊搜索 |

**请求示例**:
```json
{
  "notifyTimeEnd": 1,
  "notifyTimeStart": 1,
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
  "status": "string",
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
        "actionURL": "string",
        "body": "string",
        "coverImage": "string",
        "createdTime": 1,
        "expireTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "notifyTime": "2026-01-01T00:00:00Z",
        "status": "string",
        "summary": "string",
        "title": "string"
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
ur api /api/v1/system/notify/news/list \
  --body '{"notifyTimeEnd": 1, "notifyTimeStart": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": "string", "title": "string"}'
```

### POST `/api/v1/system/notify/news/update`

**说明**: 更新资讯

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionURL` | string | 否 |  跳转链接 |
| `body` | string | 是 |  正文内容（HTML格式） |
| `coverImage` | string | 否 |  封面图URL |
| `createdTime` | integer | 否 |  创建时间，只读 (格式: int64) |
| `expireTime` | string | 否 |  过期时间（毫秒时间戳） |
| `id` | string | 否 |  资讯ID |
| `notifyTime` | string | 否 |  发布时间（毫秒时间戳） |
| `status` | string | 否 |  状态 1:草稿 2:已发布 3:已下架 |
| `summary` | string | 否 |  摘要 |
| `title` | string | 是 |  标题 |

**请求示例**:
```json
{
  "actionURL": "string",
  "body": "string",
  "coverImage": "string",
  "createdTime": 1,
  "expireTime": "2026-01-01T00:00:00Z",
  "id": "string",
  "notifyTime": "2026-01-01T00:00:00Z",
  "status": "string",
  "summary": "string",
  "title": "string"
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
ur api /api/v1/system/notify/news/update \
  --body '{"actionURL": "string", "body": "string", "coverImage": "string", "createdTime": 1, "expireTime": "2026-01-01T00:00:00Z", "id": "string", "notifyTime": "2026-01-01T00:00:00Z", "status": "string", "summary": "string", "title": "string"}'
```

### POST `/api/v1/system/notify/notification/copy`

**说明**: 复制手动通知

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
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/notification/copy \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/create`

**说明**: 创建手动通知

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionText` | string | 否 | 操作按钮文案 |
| `actionURL` | string | 否 | 操作跳转URL |
| `actualUsers` | integer | 否 | 实际送达数,只读 (格式: int64) |
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `audienceType` | integer | 否 | 受众类型 1:指定用户 2:全局 (格式: int64) |
| `channels` | array[string] | 否 | 发送渠道列表 |
| `content` | string | 是 | 内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `createdTime` | integer | 否 | 创建时间,只读 (格式: int64) |
| `creatorID` | string | 否 | 创建者ID,只读 |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `estimatedUsers` | integer | 否 | 预计送达数,只读 (格式: int64) |
| `expireTime` | integer | 否 | 有效期(毫秒时间戳) (格式: int64) |
| `group` | string | 否 | 分组 |
| `id` | string | 否 | 通知ID |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `priority` | integer | 否 | 优先级 1:紧急 2:重要 3:普通 (格式: int64) |
| `scheduleTime` | integer | 否 | 计划发送时间(毫秒时间戳) (格式: int64) |
| `sentTime` | integer | 否 | 实际发送时间(毫秒时间戳),只读 (格式: int64) |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `status` | integer | 否 | 状态 1:草稿 2:已计划 3:发送中 4:已发送 5:已撤回 (格式: int64) |
| `targetUserIDs` | array[string] | 否 | 目标用户ID列表 |
| `tenantCode` | string | 否 | 租户编码,空=平台级 |
| `title` | string | 是 | 标题 |
| `videoUrl` | string | 否 |  news专属：视频URL |

**请求示例**:
```json
{
  "actionText": "string",
  "actionURL": "string",
  "actualUsers": 1,
  "allowDismiss": "string",
  "audienceType": 1,
  "channels": [
    "string"
  ],
  "content": "string",
  "coverImage": "string",
  "createdTime": 1,
  "creatorID": "string",
  "emailFormat": "string",
  "estimatedUsers": 1,
  "expireTime": 1,
  "group": "string",
  "id": "string",
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "priority": 1,
  "scheduleTime": 1,
  "sentTime": 1,
  "severity": "string",
  "showAsModal": "string",
  "status": 1,
  "targetUserIDs": [
    "string"
  ],
  "tenantCode": "string",
  "title": "string",
  "videoUrl": "string"
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
ur api /api/v1/system/notify/notification/create \
  --body '{"actionText": "string", "actionURL": "string", "actualUsers": 1, "allowDismiss": "string", "audienceType": 1, "channels": ["string"], "content": "string", "coverImage": "string", "createdTime": 1, "creatorID": "string", "emailFormat": "string", "estimatedUsers": 1, "expireTime": 1, "group": "string", "id": "string", "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "priority": 1, "scheduleTime": 1, "sentTime": 1, "severity": "string", "showAsModal": "string", "status": 1, "targetUserIDs": ["string"], "tenantCode": "string", "title": "string", "videoUrl": "string"}'
```

### POST `/api/v1/system/notify/notification/delete`

**说明**: 删除手动通知

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
ur api /api/v1/system/notify/notification/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/estimate-users`

**说明**: 预估送达用户数

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
    "estimatedCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/notification/estimate-users \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/index`

**说明**: 获取手动通知列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `group` | string | 否 | 分组筛选 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `priority` | integer | 否 | 优先级筛选 (格式: int64) |
| `status` | integer | 否 | 状态筛选 (格式: int64) |
| `tenantCode` | string | 否 | 租户编码筛选 |
| `title` | string | 否 | 标题模糊搜索 |

**请求示例**:
```json
{
  "group": "string",
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
  "priority": 1,
  "status": 1,
  "tenantCode": "string",
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
        "actionText": "string",
        "actionURL": "string",
        "actualUsers": 1,
        "allowDismiss": "string",
        "audienceType": 1,
        "channels": [
          "string"
        ],
        "content": "string",
        "coverImage": "string",
        "createdTime": 1,
        "creatorID": "string",
        "emailFormat": "string",
        "estimatedUsers": 1,
        "expireTime": 1,
        "group": "string",
        "id": "string",
        "modalShowEndTime": "2026-01-01T00:00:00Z",
        "modalShowStartTime": "2026-01-01T00:00:00Z",
        "priority": 1,
        "scheduleTime": 1,
        "sentTime": 1,
        "severity": "string",
        "showAsModal": "string",
        "status": 1,
        "targetUserIDs": [
          "string"
        ],
        "tenantCode": "string",
        "title": "string",
        "videoUrl": "string"
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
ur api /api/v1/system/notify/notification/index \
  --body '{"group": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "priority": 1, "status": 1, "tenantCode": "string", "title": "string"}'
```

### POST `/api/v1/system/notify/notification/read`

**说明**: 获取手动通知详情

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
    "actionText": "string",
    "actionURL": "string",
    "actualUsers": 1,
    "allowDismiss": "string",
    "audienceType": 1,
    "channels": [
      "string"
    ],
    "content": "string",
    "coverImage": "string",
    "createdTime": 1,
    "creatorID": "string",
    "emailFormat": "string",
    "estimatedUsers": 1,
    "expireTime": 1,
    "group": "string",
    "id": "string",
    "modalShowEndTime": "2026-01-01T00:00:00Z",
    "modalShowStartTime": "2026-01-01T00:00:00Z",
    "priority": 1,
    "scheduleTime": 1,
    "sentTime": 1,
    "severity": "string",
    "showAsModal": "string",
    "status": 1,
    "targetUserIDs": [
      "string"
    ],
    "tenantCode": "string",
    "title": "string",
    "videoUrl": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/notification/read \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/revoke`

**说明**: 撤回手动通知

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
ur api /api/v1/system/notify/notification/revoke \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/notification/send`

**说明**: 发送手动通知

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 通知ID |
| `scheduleTime` | integer | 否 | 计划发送时间,不填立即发送 (格式: int64) |

**请求示例**:
```json
{
  "id": "string",
  "scheduleTime": 1
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
ur api /api/v1/system/notify/notification/send \
  --body '{"id": "string", "scheduleTime": 1}'
```

### POST `/api/v1/system/notify/notification/update`

**说明**: 更新手动通知

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `actionText` | string | 否 | 操作按钮文案 |
| `actionURL` | string | 否 | 操作跳转URL |
| `actualUsers` | integer | 否 | 实际送达数,只读 (格式: int64) |
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `audienceType` | integer | 否 | 受众类型 1:指定用户 2:全局 (格式: int64) |
| `channels` | array[string] | 否 | 发送渠道列表 |
| `content` | string | 是 | 内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `createdTime` | integer | 否 | 创建时间,只读 (格式: int64) |
| `creatorID` | string | 否 | 创建者ID,只读 |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `estimatedUsers` | integer | 否 | 预计送达数,只读 (格式: int64) |
| `expireTime` | integer | 否 | 有效期(毫秒时间戳) (格式: int64) |
| `group` | string | 否 | 分组 |
| `id` | string | 否 | 通知ID |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `priority` | integer | 否 | 优先级 1:紧急 2:重要 3:普通 (格式: int64) |
| `scheduleTime` | integer | 否 | 计划发送时间(毫秒时间戳) (格式: int64) |
| `sentTime` | integer | 否 | 实际发送时间(毫秒时间戳),只读 (格式: int64) |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `status` | integer | 否 | 状态 1:草稿 2:已计划 3:发送中 4:已发送 5:已撤回 (格式: int64) |
| `targetUserIDs` | array[string] | 否 | 目标用户ID列表 |
| `tenantCode` | string | 否 | 租户编码,空=平台级 |
| `title` | string | 是 | 标题 |
| `videoUrl` | string | 否 |  news专属：视频URL |

**请求示例**:
```json
{
  "actionText": "string",
  "actionURL": "string",
  "actualUsers": 1,
  "allowDismiss": "string",
  "audienceType": 1,
  "channels": [
    "string"
  ],
  "content": "string",
  "coverImage": "string",
  "createdTime": 1,
  "creatorID": "string",
  "emailFormat": "string",
  "estimatedUsers": 1,
  "expireTime": 1,
  "group": "string",
  "id": "string",
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "priority": 1,
  "scheduleTime": 1,
  "sentTime": 1,
  "severity": "string",
  "showAsModal": "string",
  "status": 1,
  "targetUserIDs": [
    "string"
  ],
  "tenantCode": "string",
  "title": "string",
  "videoUrl": "string"
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
ur api /api/v1/system/notify/notification/update \
  --body '{"actionText": "string", "actionURL": "string", "actualUsers": 1, "allowDismiss": "string", "audienceType": 1, "channels": ["string"], "content": "string", "coverImage": "string", "createdTime": 1, "creatorID": "string", "emailFormat": "string", "estimatedUsers": 1, "expireTime": 1, "group": "string", "id": "string", "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "priority": 1, "scheduleTime": 1, "sentTime": 1, "severity": "string", "showAsModal": "string", "status": 1, "targetUserIDs": ["string"], "tenantCode": "string", "title": "string", "videoUrl": "string"}'
```

### POST `/api/v1/system/notify/template/create`

**说明**: 添加通知模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `body` | string | 否 | 默认模版内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `desc` | string | 否 |  备注 |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `id` | string | 否 |  id编号 |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `name` | string | 否 | 通知的命名 |
| `notifyCode` | string | 否 | 对应的配置Code |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `subject` | string | 否 | 默认消息主题(邮箱需要填写) |
| `templateCode` | string | 否 |  短信通知模版编码(短信必填) |
| `type` | string | 否 | 对应的配置类型 sms:短信, email:邮箱, dingTalk:钉钉机器人, dingWebhook:钉钉webhook, wxMini:微信小程序推送, wxEWebHook:企业微信webhook, message:站内信通知, phoneCall:电话通知, banner:横幅通知, news:资讯通知, wxOfficialAccount:微信公众号 可选: ['sms', 'email', 'dingTalk', 'dingWebhook', 'wxMini', 'wxEWebHook', 'message', 'phoneCall', 'banner', 'news', 'wxOfficialAccount'] |
| `videoUrl` | string | 否 |  news专属：视频URL |
| `webhook` | string | 否 | 钉钉webhook及企微webhook |

**请求示例**:
```json
{
  "allowDismiss": "string",
  "body": "string",
  "coverImage": "string",
  "desc": "string",
  "emailFormat": "string",
  "id": "string",
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "name": "string",
  "notifyCode": "string",
  "severity": "string",
  "showAsModal": "string",
  "subject": "string",
  "templateCode": "string",
  "type": "sms",
  "videoUrl": "string",
  "webhook": "string"
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
ur api /api/v1/system/notify/template/create \
  --body '{"allowDismiss": "string", "body": "string", "coverImage": "string", "desc": "string", "emailFormat": "string", "id": "string", "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "name": "string", "notifyCode": "string", "severity": "string", "showAsModal": "string", "subject": "string", "templateCode": "string", "type": "sms", "videoUrl": "string", "webhook": "string"}'
```

### POST `/api/v1/system/notify/template/delete`

**说明**: 删除通知模版

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
ur api /api/v1/system/notify/template/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/template/get-list`

**说明**: 获取通知模版列表

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `notifyCode` | string | 否 |  应用编号 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `type` | string | 否 | 对应的配置类型 sms email |

**请求示例**:
```json
{
  "name": "string",
  "notifyCode": "string",
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
        "allowDismiss": "string",
        "body": "string",
        "coverImage": "string",
        "desc": "string",
        "emailFormat": "string",
        "id": "string",
        "modalShowEndTime": "2026-01-01T00:00:00Z",
        "modalShowStartTime": "2026-01-01T00:00:00Z",
        "name": "string",
        "notifyCode": "string",
        "severity": "string",
        "showAsModal": "string",
        "subject": "string",
        "templateCode": "string",
        "type": "sms",
        "videoUrl": "string",
        "webhook": "string"
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
ur api /api/v1/system/notify/template/get-list \
  --body '{"name": "string", "notifyCode": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "type": "string"}'
```

### POST `/api/v1/system/notify/template/get-one`

**说明**: 获取通知模版详情

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
    "allowDismiss": "string",
    "body": "string",
    "coverImage": "string",
    "desc": "string",
    "emailFormat": "string",
    "id": "string",
    "modalShowEndTime": "2026-01-01T00:00:00Z",
    "modalShowStartTime": "2026-01-01T00:00:00Z",
    "name": "string",
    "notifyCode": "string",
    "severity": "string",
    "showAsModal": "string",
    "subject": "string",
    "templateCode": "string",
    "type": "sms",
    "videoUrl": "string",
    "webhook": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/notify/template/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/notify/template/update`

**说明**: 更新通知模版

**权限**: platform

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `allowDismiss` | string | 否 |  banner专属：允许关闭 1=是 2=否 |
| `body` | string | 否 | 默认模版内容 |
| `coverImage` | string | 否 |  news专属：封面图URL |
| `desc` | string | 否 |  备注 |
| `emailFormat` | string | 否 |  email专属：邮件格式 richtext/html |
| `id` | string | 否 |  id编号 |
| `modalShowEndTime` | string | 否 |  弹窗结束时间(毫秒) |
| `modalShowStartTime` | string | 否 |  弹窗开始时间(毫秒) |
| `name` | string | 否 | 通知的命名 |
| `notifyCode` | string | 否 | 对应的配置Code |
| `severity` | string | 否 |  banner专属：严重等级 info/warning/error |
| `showAsModal` | string | 否 |  message/news专属：是否弹窗 1=是 2=否 |
| `subject` | string | 否 | 默认消息主题(邮箱需要填写) |
| `templateCode` | string | 否 |  短信通知模版编码(短信必填) |
| `type` | string | 否 | 对应的配置类型 sms:短信, email:邮箱, dingTalk:钉钉机器人, dingWebhook:钉钉webhook, wxMini:微信小程序推送, wxEWebHook:企业微信webhook, message:站内信通知, phoneCall:电话通知, banner:横幅通知, news:资讯通知, wxOfficialAccount:微信公众号 可选: ['sms', 'email', 'dingTalk', 'dingWebhook', 'wxMini', 'wxEWebHook', 'message', 'phoneCall', 'banner', 'news', 'wxOfficialAccount'] |
| `videoUrl` | string | 否 |  news专属：视频URL |
| `webhook` | string | 否 | 钉钉webhook及企微webhook |

**请求示例**:
```json
{
  "allowDismiss": "string",
  "body": "string",
  "coverImage": "string",
  "desc": "string",
  "emailFormat": "string",
  "id": "string",
  "modalShowEndTime": "2026-01-01T00:00:00Z",
  "modalShowStartTime": "2026-01-01T00:00:00Z",
  "name": "string",
  "notifyCode": "string",
  "severity": "string",
  "showAsModal": "string",
  "subject": "string",
  "templateCode": "string",
  "type": "sms",
  "videoUrl": "string",
  "webhook": "string"
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
ur api /api/v1/system/notify/template/update \
  --body '{"allowDismiss": "string", "body": "string", "coverImage": "string", "desc": "string", "emailFormat": "string", "id": "string", "modalShowEndTime": "2026-01-01T00:00:00Z", "modalShowStartTime": "2026-01-01T00:00:00Z", "name": "string", "notifyCode": "string", "severity": "string", "showAsModal": "string", "subject": "string", "templateCode": "string", "type": "sms", "videoUrl": "string", "webhook": "string"}'
```

### POST `/api/v1/system/role/app/batch-update`

**说明**: 更新APP权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appIDs` | array[string] | 是 | App列表数据 |
| `id` | string | 是 | 角色ID |

**请求示例**:
```json
{
  "appIDs": [
    "string"
  ],
  "id": "string"
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
ur api /api/v1/system/role/app/batch-update \
  --body '{"appIDs": ["string"], "id": "string"}'
```

### POST `/api/v1/system/role/app/get-list`

**说明**: 获取APP权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 角色ID |

**请求示例**:
```json
{
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "appIDs": [
      "string"
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/role/app/get-list \
  --body '{"id": "string"}'
```

### POST `/api/v1/system/role/info/create`

**说明**: 添加角色

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 | 角色编码 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  备注 |
| `id` | string | 否 |  编号 |
| `isFullAuth` | integer | 否 |  是否完全授权（1:是，2:否） (格式: int64) |
| `name` | string | 否 |  角色名称 |
| `status` | integer | 否 |  角色状态 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "isFullAuth": 1,
  "name": "示例名称",
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
ur api /api/v1/system/role/info/create \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "isFullAuth": 1, "name": "示例名称", "status": 1}'
```

### POST `/api/v1/system/role/info/delete`

**说明**: 删除角色

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
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/role/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/role/info/get-list`

**说明**: 获取角色列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `codes` | array[string] | 否 | 角色编码过滤 |
| `ids` | array[string] | 否 |  |
| `name` | string | 否 | 按名称查找角色 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 按状态查找角色 (格式: int64) |

**请求示例**:
```json
{
  "codes": [
    "string"
  ],
  "ids": [
    "string"
  ],
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
        "isFullAuth": 1,
        "name": "示例名称",
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
ur api /api/v1/system/role/info/get-list \
  --body '{"codes": ["string"], "ids": ["string"], "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1}'
```

### POST `/api/v1/system/role/info/update`

**说明**: 更新角色

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 | 角色编码 |
| `createdTime` | string | 否 |  创建时间 |
| `desc` | string | 否 |  备注 |
| `id` | string | 否 |  编号 |
| `isFullAuth` | integer | 否 |  是否完全授权（1:是，2:否） (格式: int64) |
| `name` | string | 否 |  角色名称 |
| `status` | integer | 否 |  角色状态 (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "id": "string",
  "isFullAuth": 1,
  "name": "示例名称",
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
ur api /api/v1/system/role/info/update \
  --body '{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "isFullAuth": 1, "name": "示例名称", "status": 1}'
```

### POST `/api/v1/system/role/menu/batch-update`

**说明**: 更新角色对应菜单列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  应用编号 |
| `id` | string | 是 | 角色编号 |
| `menuIDs` | array[string] | 是 | 菜单编号列表 |

**请求示例**:
```json
{
  "appID": "string",
  "id": "string",
  "menuIDs": [
    "string"
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
ur api /api/v1/system/role/menu/batch-update \
  --body '{"appID": "string", "id": "string", "menuIDs": ["string"]}'
```

### POST `/api/v1/system/role/menu/get-list`

**说明**: 获取角色对应菜单列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  应用编号 |
| `id` | string | 是 | 角色编号 |

**请求示例**:
```json
{
  "appID": "string",
  "id": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "menuIDs": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/role/menu/get-list \
  --body '{"appID": "string", "id": "string"}'
```

### POST `/api/v1/system/role/resource/batch-update`

**说明**: 批量更新角色资源动作权限

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `list` | array[RoleResourceActionItem] | 是 |  完整权限列表（全量替换） |
| `roleId` | string | 是 |  角色ID |

**请求示例**:
```json
{
  "list": [
    {
      "action": "string",
      "resource": "string"
    }
  ],
  "roleId": "string"
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
ur api /api/v1/system/role/resource/batch-update \
  --body '{"list": [{"action": "string", "resource": "string"}], "roleId": "string"}'
```

### POST `/api/v1/system/role/resource/get-list`

**说明**: 获取角色资源动作权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `resource` | string | 否 |  可选资源筛选 |
| `roleId` | string | 是 |  角色ID |

**请求示例**:
```json
{
  "resource": "string",
  "roleId": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "action": "string",
        "resource": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/role/resource/get-list \
  --body '{"resource": "string", "roleId": "string"}'
```

### POST `/api/v1/system/user/data/area/get-list`

**说明**: 获取区域权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `projectID` | string | 否 | 项目id |
| `userID` | string | 是 | 用户ID |

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
  "projectID": "string",
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
        "": {
          "areaID": "string",
          "areaIDPath": "string",
          "areaImg": "string",
          "areaName": "示例名称",
          "areaNamePath": "string",
          "children": [
            {
              "areaID": "...",
              "areaIDPath": "...",
              "areaImg": "...",
              "areaName": "...",
              "areaNamePath": "...",
              "children": "...",
              "configFile": "...",
              "createdTime": "...",
              "desc": "...",
              "deviceCount": "...",
              "devices": "...",
              "groupCount": "...",
              "isLeaf": "...",
              "isSysCreated": "...",
              "lowerLevelCount": "...",
              "parentAreaID": "...",
              "position": "...",
              "projectID": "...",
              "sort": "...",
              "tags": "...",
              "tenantCode": "...",
              "useBy": "..."
            }
          ],
          "configFile": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [
            {
              "adcode": "...",
              "address": "...",
              "area": "...",
              "areaID": "...",
              "categoryID": "...",
              "cert": "...",
              "cloneID": "...",
              "createdTime": "...",
              "desc": "...",
              "deviceAlias": "...",
              "deviceImg": "...",
              "deviceName": "...",
              "deviceType": "...",
              "distributor": "...",
              "expTime": "...",
              "file": "...",
              "firstBind": "...",
              "firstLogin": "...",
              "gateway": "...",
              "groupIDs": "...",
              "groupPurpose": "...",
              "groups": "...",
              "hardInfo": "...",
              "iccid": "...",
              "id": "...",
              "imei": "...",
              "isEnable": "...",
              "isOnline": "...",
              "lastBind": "...",
              "lastIp": "...",
              "lastLocalIp": "...",
              "lastLogin": "...",
              "lastOffline": "...",
              "logLevel": "...",
              "mac": "...",
              "mobileOperator": "...",
              "needConfirmVersion": "...",
              "netType": "...",
              "owner": "...",
              "phone": "...",
              "position": "...",
              "productID": "...",
              "productImg": "...",
              "productName": "...",
              "profiles": "...",
              "projectID": "...",
              "protocolConf": "...",
              "ratedPower": "...",
              "remoteAddr": "...",
              "rssi": "...",
              "schemaAlias": "...",
              "secret": "...",
              "softInfo": "...",
              "sort": "...",
              "status": "...",
              "subProtocolConf": "...",
              "tags": "...",
              "tenantCode": "...",
              "userID": "...",
              "version": "...",
              "withProperties": "..."
            }
          ],
          "groupCount": 1,
          "isLeaf": 1,
          "isSysCreated": 1,
          "lowerLevelCount": 1,
          "parentAreaID": "string",
          "position": {
            "latitude": 1,
            "longitude": 1
          },
          "projectID": "string",
          "sort": 1,
          "tags": {},
          "tenantCode": "string",
          "useBy": "string"
        },
        "authType": 1,
        "isAuthChildren": 1,
        "targetID": "string",
        "targetType": "string",
        "updatedTime": 1,
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
ur api /api/v1/system/user/data/area/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "projectID": "string", "userID": "string"}'
```

### POST `/api/v1/system/user/data/project/get-list`

**说明**: 获取项目权限列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `userID` | string | 是 | 用户ID |

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
        "authType": 1,
        "project": {
          "address": "string",
          "adminUser": {
            "avatar": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "email": "string",
            "lastIP": "string",
            "nickName": "string",
            "phone": "string",
            "regIP": "string",
            "userID": "string",
            "userName": "string"
          },
          "adminUserID": "string",
          "alarmStatus": 1,
          "area": 1,
          "areaCount": 1,
          "areas": [
            {
              "areaID": "...",
              "areaIDPath": "...",
              "areaImg": "...",
              "areaName": "...",
              "areaNamePath": "...",
              "children": "...",
              "configFile": "...",
              "createdTime": "...",
              "desc": "...",
              "deviceCount": "...",
              "devices": "...",
              "groupCount": "...",
              "isLeaf": "...",
              "isSysCreated": "...",
              "lowerLevelCount": "...",
              "parentAreaID": "...",
              "position": "...",
              "projectID": "...",
              "sort": "...",
              "tags": "...",
              "tenantCode": "...",
              "useBy": "..."
            }
          ],
          "attachments": [
            {
              "fileName": "...",
              "filePath": "...",
              "fileUrl": "...",
              "id": "...",
              "useBy": "..."
            }
          ],
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "deviceOnlineCount": 1,
          "isSysCreated": 1,
          "position": {
            "latitude": 1,
            "longitude": 1
          },
          "ppsm": 1,
          "projectID": "string",
          "projectImg": "string",
          "projectName": "示例名称",
          "sort": 1,
          "status": 1,
          "tags": {},
          "tenantCode": "string",
          "type": "string",
          "userCount": 1
        },
        "projectID": "string",
        "targetID": "string",
        "updatedTime": 1
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
ur api /api/v1/system/user/data/project/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "userID": "string"}'
```

### POST `/api/v1/system/user/dept/batch-create`

**说明**: 新增用户的部门列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptIDs` | array[string] | 是 | 角色编号列表 |
| `userID` | string | 是 |  用户id |

**请求示例**:
```json
{
  "deptIDs": [
    "string"
  ],
  "userID": "string"
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
ur api /api/v1/system/user/dept/batch-create \
  --body '{"deptIDs": ["string"], "userID": "string"}'
```

### POST `/api/v1/system/user/dept/batch-delete`

**说明**: 删除用户的部门列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deptIDs` | array[string] | 是 | 角色编号列表 |
| `userID` | string | 是 |  用户id |

**请求示例**:
```json
{
  "deptIDs": [
    "string"
  ],
  "userID": "string"
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
ur api /api/v1/system/user/dept/batch-delete \
  --body '{"deptIDs": ["string"], "userID": "string"}'
```

### POST `/api/v1/system/user/info/create`

**说明**: 创建用户信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `info` | object | 是 |  |
| `info.avatar` | string | 否 |  头像 |
| `info.email` | string | 否 |  邮箱 |
| `info.nickName` | string | 否 |  昵称 |
| `info.phone` | string | 否 |  手机号 |
| `info.userID` | string | 否 |  用户id |
| `info.userName` | string | 否 |  用户名 |
| `roleIDs` | array[string] | 否 | 平台的角色编号列表 |

**请求示例**:
```json
{
  "info": {
    "avatar": "string",
    "email": "string",
    "nickName": "string",
    "phone": "string",
    "userID": "string",
    "userName": "string"
  },
  "roleIDs": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/info/create \
  --body '{"info": {"avatar": "string", "email": "string", "nickName": "string", "phone": "string", "userID": "string", "userName": "string"}, "roleIDs": ["string"]}'
```

### POST `/api/v1/system/user/info/delete`

**说明**: 刪除用户

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `userID` | string | 否 |  用户id |

**请求示例**:
```json
{
  "userID": "string"
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
ur api /api/v1/system/user/info/delete \
  --body '{"userID": "string"}'
```

### POST `/api/v1/system/user/info/get-list`

**说明**: 查询用户信息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 |  |
| `email` | string | 否 |  邮箱 |
| `nickName` | string | 否 |  用户的昵称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `phone` | string | 否 |  手机号 |
| `userIDs` | array[string] | 否 |  |
| `userName` | string | 否 | 用户名(唯一) |

**请求示例**:
```json
{
  "account": "string",
  "email": "string",
  "nickName": "string",
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
  "phone": "string",
  "userIDs": [
    "string"
  ],
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
        "avatar": "string",
        "email": "string",
        "nickName": "string",
        "phone": "string",
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
ur api /api/v1/system/user/info/get-list \
  --body '{"account": "string", "email": "string", "nickName": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "phone": "string", "userIDs": ["string"], "userName": "string"}'
```

### POST `/api/v1/system/user/info/get-one`

**说明**: 获取用户信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `userID` | string | 否 |  用户id |
| `withTenant` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "userID": "string",
  "withTenant": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avatar": "string",
    "email": "string",
    "nickName": "string",
    "phone": "string",
    "userID": "string",
    "userName": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/info/get-one \
  --body '{"userID": "string", "withTenant": true}'
```

### POST `/api/v1/system/user/info/update`

**说明**: 更新用户基本数据

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  头像（传图片的filePath） |
| `email` | string | 否 |  邮箱,平台管理员可修改 |
| `nickName` | string | 否 |  用户昵称 |
| `phone` | string | 否 |  手机号,平台管理员可修改 |
| `sex` | integer | 否 |  性别（1:男，2:女，0:未知） (格式: int64) |
| `userID` | string | 是 |  用户ID（必填） |
| `userName` | string | 否 |  登录用户名,,平台管理员可修改 |

**请求示例**:
```json
{
  "avatar": "string",
  "email": "string",
  "nickName": "string",
  "phone": "string",
  "sex": 1,
  "userID": "string",
  "userName": "string"
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
ur api /api/v1/system/user/info/update \
  --body '{"avatar": "string", "email": "string", "nickName": "string", "phone": "string", "sex": 1, "userID": "string", "userName": "string"}'
```

### POST `/api/v1/system/user/self/access-token/create`

**说明**: 创建访问令牌

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  令牌描述 |
| `expTime` | string | 否 |  到期时间 |
| `tenantCodes` | array[string] | 否 |  可访问的租户列表 |

**请求示例**:
```json
{
  "desc": "string",
  "expTime": "2026-01-01T00:00:00Z",
  "tenantCodes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "accessKey": "string",
    "accessSecret": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "expTime": "2026-01-01T00:00:00Z",
    "id": "string",
    "tenantCodes": [
      "string"
    ],
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/access-token/create \
  --body '{"desc": "string", "expTime": "2026-01-01T00:00:00Z", "tenantCodes": ["string"]}'
```

### POST `/api/v1/system/user/self/access-token/delete`

**说明**: 删除访问令牌

**权限**: all

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
ur api /api/v1/system/user/self/access-token/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/user/self/access-token/get-list`

**说明**: 获取访问令牌列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessKey` | string | 否 |  按访问密钥标识过滤 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "accessKey": "string",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
        "accessSecret": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "expTime": "2026-01-01T00:00:00Z",
        "id": "string",
        "tenantCodes": [
          "string"
        ],
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
ur api /api/v1/system/user/self/access-token/get-list \
  --body '{"accessKey": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/user/self/access-token/get-one`

**说明**: 获取访问令牌详情

**权限**: all

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
    "accessKey": "string",
    "accessSecret": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "expTime": "2026-01-01T00:00:00Z",
    "id": "string",
    "tenantCodes": [
      "string"
    ],
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/access-token/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/system/user/self/access-token/update`

**说明**: 更新访问令牌

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  令牌描述 |
| `expTime` | string | 否 |  到期时间 |
| `id` | string | 是 |  令牌ID |
| `tenantCodes` | array[string] | 否 |  可访问的租户列表 |

**请求示例**:
```json
{
  "desc": "string",
  "expTime": "2026-01-01T00:00:00Z",
  "id": "string",
  "tenantCodes": [
    "string"
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
ur api /api/v1/system/user/self/access-token/update \
  --body '{"desc": "string", "expTime": "2026-01-01T00:00:00Z", "id": "string", "tenantCodes": ["string"]}'
```

### POST `/api/v1/system/user/self/app/get-list`

**说明**: 获取用户应用列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `notAppID` | string | 否 |  排除指定应用ID（如排除控制台大厅 77） |

**请求示例**:
```json
{
  "notAppID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
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
        "createdTime": "string",
        "expTime": 1,
        "id": "string",
        "sort": 1,
        "status": 1,
        "tenantCode": "string"
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
ur api /api/v1/system/user/self/app/get-list \
  --body '{"notAppID": "string"}'
```

### POST `/api/v1/system/user/self/app/get-one`

**说明**: 获取用户应用详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  |
| `subType` | string | 是 | wx:微信小程序  dingTalk:钉钉小程序 |
| `type` | string | 是 |   app:应用  mini:小程序 |

**请求示例**:
```json
{
  "appID": "string",
  "subType": "string",
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "appID": "string",
    "name": "示例名称",
    "tenantCodes": [
      "string"
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/app/get-one \
  --body '{"appID": "string", "subType": "string", "type": "string"}'
```

### POST `/api/v1/system/user/self/bind-account`

**说明**: 绑定账号

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 手机号注册时填写手机号 email填写邮箱 |
| `code` | string | 否 | 验证码    微信登录填code 账号密码登录时填写密码 |
| `codeID` | string | 否 | 验证码编号 微信登录填state |
| `type` | string | 是 | 验证方式:phone手机号注册 wxOpen 微信开放平台登录 wxIn 微信内 wxMiniP 微信小程序 可选: ['phone', 'email', 'wxOpen', 'wxIn', 'wxMiniP', 'dingApp'] |

**请求示例**:
```json
{
  "account": "string",
  "code": "******",
  "codeID": "string",
  "type": "phone"
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
ur api /api/v1/system/user/self/bind-account \
  --body '{"account": "string", "code": "******", "codeID": "string", "type": "phone"}'
```

### POST `/api/v1/system/user/self/cancel`

**说明**: 注销用户

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
ur api /api/v1/system/user/self/cancel \
  --body '{}'
```

### POST `/api/v1/system/user/self/captcha`

**说明**: 获取验证码

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 短信验证时填写手机号,邮箱验证时填写邮箱 |
| `code` | string | 否 | 验证码    微信邮箱验证登录填code |
| `codeID` | string | 否 | 验证码编号 微信邮箱验证登录填state |
| `type` | string | 是 | 验证方式:短信验证,图片验证码 可选: ['phone', 'image', 'email'] |
| `use` | string | 是 | 用途 可选: ['login', 'register', 'changePwd', 'forgetPwd', 'bindAccount'] |

**请求示例**:
```json
{
  "account": "string",
  "code": "string",
  "codeID": "string",
  "type": "phone",
  "use": "login"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "codeID": "string",
    "expire": 1,
    "url": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/captcha \
  --body '{"account": "string", "code": "string", "codeID": "string", "type": "phone", "use": "login"}'
```

### POST `/api/v1/system/user/self/change-pwd`

**说明**: 更新用户密码

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 | 验证码    微信登录填code 账号密码登录时填写密码 |
| `codeID` | string | 否 | 验证码编号 微信登录填state |
| `oldPassword` | string | 否 | pwd方式需要填写 如果原来没有设置过密码可以不填 输入原密码修改密码(sha256格式并需要输入图像验证码) |
| `password` | string | 否 | 密码 |
| `type` | string | 是 | 验证方式: phone手机号 email邮箱  pwd:原密码 可选: ['phone', 'email', 'pwd'] |

**请求示例**:
```json
{
  "code": "******",
  "codeID": "string",
  "oldPassword": "******",
  "password": "******",
  "type": "phone"
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
ur api /api/v1/system/user/self/change-pwd \
  --body '{"code": "******", "codeID": "string", "oldPassword": "******", "password": "******", "type": "phone"}'
```

### POST `/api/v1/system/user/self/forget-pwd`

**说明**: 忘记密码

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 手机号注册时填写手机号 账号密码注册时填写userName |
| `code` | string | 否 | 验证码    微信登录填code 账号密码登录时填写密码 |
| `codeID` | string | 否 | 验证码编号 微信登录填state |
| `password` | string | 是 | 密码 |
| `type` | string | 是 | 验证方式: phone手机号 email邮箱 可选: ['phone', 'email'] |

**请求示例**:
```json
{
  "account": "******",
  "code": "******",
  "codeID": "string",
  "password": "******",
  "type": "phone"
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
ur api /api/v1/system/user/self/forget-pwd \
  --body '{"account": "******", "code": "******", "codeID": "string", "password": "******", "type": "phone"}'
```

### POST `/api/v1/system/user/self/get-one`

**说明**: 获取用户信息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `withTenant` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "withTenant": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avatar": "string",
    "email": "string",
    "nickName": "string",
    "phone": "string",
    "userID": "string",
    "userName": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/get-one \
  --body '{"withTenant": true}'
```

### POST `/api/v1/system/user/self/login`

**说明**: 用户登录

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 登录账号(支持用户名,手机号,邮箱) 账号密码登录时需要填写 |
| `code` | string | 否 | 验证码    微信邮箱验证登录填code |
| `codeID` | string | 否 | 验证码编号 微信邮箱验证登录填state |
| `loginType` | string | 是 | 验证类型 phone 手机号 wxOpen 微信开放平台 wxIn 微信内 wxMiniP 微信小程序 pwd 账号密码 email 邮箱 可选: ['phone', 'wxOpen', 'wxIn', 'wxMiniP', 'dingApp', 'pwd', 'email'] |
| `password` | string | 否 | 密码，建议md5转换 密码登录时需要填写 |
| `pwdType` | integer | 否 | 账号密码登录时需要填写.0,无密码 1，明文 2，md5加密 (格式: int32) |

**请求示例**:
```json
{
  "account": "******",
  "code": "string",
  "codeID": "string",
  "loginType": "phone",
  "password": "******",
  "pwdType": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "info": {
      "avatar": "string",
      "email": "string",
      "nickName": "string",
      "phone": "string",
      "userID": "string",
      "userName": "string"
    },
    "token": {
      "accessExpire": "2026-01-01T00:00:00Z",
      "accessToken": "string",
      "refreshAfter": "2026-01-01T00:00:00Z"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/login \
  --body '{"account": "******", "code": "string", "codeID": "string", "loginType": "phone", "password": "******", "pwdType": 1}'
```

### POST `/api/v1/system/user/self/logout`

**说明**: 用户登出

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
ur api /api/v1/system/user/self/logout \
  --body '{}'
```

### POST `/api/v1/system/user/self/menu/get-list`

**说明**: 获取用户菜单列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  应用编号 |
| `isCommon` | integer | 否 |  是否常用菜单 1-是 2-否 (格式: int64) |

**请求示例**:
```json
{
  "appID": "string",
  "isCommon": 1
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
        "children": [
          {
            "appID": "string",
            "children": [],
            "icon": "string",
            "id": "string",
            "isHideInMenu": 1,
            "isStar": 1,
            "menu": {
              "appID": "...",
              "body": "...",
              "children": "...",
              "component": "...",
              "createdTime": "...",
              "defaultIcon": "...",
              "defaultIsHideInMenu": "...",
              "defaultName": "...",
              "defaultSort": "...",
              "id": "...",
              "idPath": "...",
              "parentID": "...",
              "path": "...",
              "redirect": "...",
              "role": "...",
              "tenantCode": "...",
              "type": "..."
            },
            "menuID": "string",
            "name": "示例名称",
            "parentMenuID": "string",
            "sort": 1
          }
        ],
        "icon": "string",
        "id": "string",
        "isHideInMenu": 1,
        "isStar": 1,
        "menu": {
          "appID": "string",
          "body": "string",
          "children": [
            {
              "appID": "...",
              "body": "...",
              "children": "...",
              "component": "...",
              "createdTime": "...",
              "defaultIcon": "...",
              "defaultIsHideInMenu": "...",
              "defaultName": "...",
              "defaultSort": "...",
              "id": "...",
              "idPath": "...",
              "parentID": "...",
              "path": "...",
              "redirect": "...",
              "role": "...",
              "tenantCode": "...",
              "type": "..."
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
        },
        "menuID": "string",
        "name": "示例名称",
        "parentMenuID": "string",
        "sort": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/menu/get-list \
  --body '{"appID": "string", "isCommon": 1}'
```

### POST `/api/v1/system/user/self/message/get-list`

**说明**: 用户消息列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 | 创建时间范围 |
| `group` | string | 否 | 分组 |
| `isHandled` | integer | 否 | 是否已处理 1:已处理 2:未处理 (格式: int64) |
| `isRead` | integer | 否 | 是否已读 1:已读 2:未读 (格式: int64) |
| `notifyCode` | string | 否 | 通知编码 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `str1` | string | 否 |  |
| `str2` | string | 否 |  |
| `str3` | string | 否 |  |
| `type` | string | 否 | 渠道类型 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "group": "string",
  "isHandled": 1,
  "isRead": 1,
  "notifyCode": "string",
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
  "str1": "string",
  "str2": "string",
  "str3": "string",
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
        "createdTime": "2026-01-01T00:00:00Z",
        "group": "string",
        "id": "string",
        "isHandled": 1,
        "isRead": 1,
        "message": {
          "actionText": "string",
          "actionURL": "string",
          "allowDismiss": "string",
          "body": "string",
          "coverImage": "string",
          "createdTime": "string",
          "emailFormat": "string",
          "expireTime": 1,
          "group": "string",
          "id": "string",
          "isGlobal": 1,
          "modalShowEndTime": "2026-01-01T00:00:00Z",
          "modalShowStartTime": "2026-01-01T00:00:00Z",
          "notificationID": "string",
          "notifyCode": "string",
          "notifyName": "示例名称",
          "notifyTime": 1,
          "priority": 1,
          "severity": "string",
          "showAsModal": "string",
          "str1": "string",
          "str2": "string",
          "str3": "string",
          "subject": "string",
          "type": "string",
          "videoUrl": "string"
        },
        "messageID": "string",
        "type": "string",
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
ur api /api/v1/system/user/self/message/get-list \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "group": "string", "isHandled": 1, "isRead": 1, "notifyCode": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "str1": "string", "str2": "string", "str3": "string", "type": "string"}'
```

### POST `/api/v1/system/user/self/message/get-pending`

**说明**: 用户待处理消息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

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
        "createdTime": "2026-01-01T00:00:00Z",
        "group": "string",
        "id": "string",
        "isHandled": 1,
        "isRead": 1,
        "message": {
          "actionText": "string",
          "actionURL": "string",
          "allowDismiss": "string",
          "body": "string",
          "coverImage": "string",
          "createdTime": "string",
          "emailFormat": "string",
          "expireTime": 1,
          "group": "string",
          "id": "string",
          "isGlobal": 1,
          "modalShowEndTime": "2026-01-01T00:00:00Z",
          "modalShowStartTime": "2026-01-01T00:00:00Z",
          "notificationID": "string",
          "notifyCode": "string",
          "notifyName": "示例名称",
          "notifyTime": 1,
          "priority": 1,
          "severity": "string",
          "showAsModal": "string",
          "str1": "string",
          "str2": "string",
          "str3": "string",
          "subject": "string",
          "type": "string",
          "videoUrl": "string"
        },
        "messageID": "string",
        "type": "string",
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
ur api /api/v1/system/user/self/message/get-pending \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/system/user/self/message/handle`

**说明**: 用户消息标记已处理

**权限**: all

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
ur api /api/v1/system/user/self/message/handle \
  --body '{"id": 1}'
```

### POST `/api/v1/system/user/self/message/mark-all-read`

**说明**: 用户消息全部已读

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
ur api /api/v1/system/user/self/message/mark-all-read \
  --body '{}'
```

### POST `/api/v1/system/user/self/message/multi-delete`

**说明**: 用户消息批量删除

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[integer] | 是 |  |

**请求示例**:
```json
{
  "ids": [
    1
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
ur api /api/v1/system/user/self/message/multi-delete \
  --body '{"ids": [1]}'
```

### POST `/api/v1/system/user/self/message/multi-is-read`

**说明**: 用户消息批量已读

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[integer] | 是 |  |

**请求示例**:
```json
{
  "ids": [
    1
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
ur api /api/v1/system/user/self/message/multi-is-read \
  --body '{"ids": [1]}'
```

### POST `/api/v1/system/user/self/message/statistics`

**说明**: 用户消息统计

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "totalPending": 1,
    "totalUnread": 1,
    "unreadCounts": {}
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/message/statistics \
  --body '{}'
```

### POST `/api/v1/system/user/self/notify-preference/read`

**说明**: 用户通知偏好读取

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "quietChannels": [
      "string"
    ],
    "quietEndTime": "2026-01-01T00:00:00Z",
    "quietStartTime": "2026-01-01T00:00:00Z",
    "subscriptions": [
      {
        "group": "string",
        "isEnabled": 1,
        "type": "string"
      }
    ],
    "tenantCode": "string",
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/notify-preference/read \
  --body '{}'
```

### POST `/api/v1/system/user/self/notify-preference/update`

**说明**: 用户通知偏好更新

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `quietChannels` | array[string] | 否 | 免打扰渠道列表 |
| `quietEndTime` | string | 否 | 免打扰结束时间 HH:MM |
| `quietStartTime` | string | 否 | 免打扰开始时间 HH:MM |
| `subscriptions` | array[UserNotifySubscription] | 否 | 订阅控制列表 |
| `tenantCode` | string | 否 | 租户编码,只读 |
| `userID` | string | 否 | 用户ID,只读 |

**请求示例**:
```json
{
  "quietChannels": [
    "string"
  ],
  "quietEndTime": "2026-01-01T00:00:00Z",
  "quietStartTime": "2026-01-01T00:00:00Z",
  "subscriptions": [
    {
      "group": "string",
      "isEnabled": 1,
      "type": "string"
    }
  ],
  "tenantCode": "string",
  "userID": "string"
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
ur api /api/v1/system/user/self/notify-preference/update \
  --body '{"quietChannels": ["string"], "quietEndTime": "2026-01-01T00:00:00Z", "quietStartTime": "2026-01-01T00:00:00Z", "subscriptions": [{"group": "string", "isEnabled": 1, "type": "string"}], "tenantCode": "string", "userID": "string"}'
```

### POST `/api/v1/system/user/self/profile/get-list`

**说明**: 获取用户配置列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `codes` | array[string] | 是 |  |

**请求示例**:
```json
{
  "codes": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "profiles": [
      {
        "code": "string",
        "params": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/profile/get-list \
  --body '{"codes": ["string"]}'
```

### POST `/api/v1/system/user/self/profile/get-one`

**说明**: 获取用户配置详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `withProjects` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "code": "string",
  "withProjects": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "params": "string",
    "projects": [
      {
        "address": "string",
        "adminUser": {
          "avatar": "string",
          "createdTime": "2026-01-01T00:00:00Z",
          "email": "string",
          "lastIP": "string",
          "nickName": "string",
          "phone": "string",
          "regIP": "string",
          "userID": "string",
          "userName": "string"
        },
        "adminUserID": "string",
        "alarmStatus": 1,
        "area": 1,
        "areaCount": 1,
        "areas": [
          {
            "areaID": "string",
            "areaIDPath": "string",
            "areaImg": "string",
            "areaName": "示例名称",
            "areaNamePath": "string",
            "children": [],
            "configFile": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "groupCount": 1,
            "isLeaf": 1,
            "isSysCreated": 1,
            "lowerLevelCount": 1,
            "parentAreaID": "string",
            "position": {
              "latitude": "...",
              "longitude": "..."
            },
            "projectID": "string",
            "sort": 1,
            "tags": {},
            "tenantCode": "string",
            "useBy": "string"
          }
        ],
        "attachments": [
          {
            "fileName": "string",
            "filePath": "string",
            "fileUrl": "string",
            "id": "string",
            "useBy": "string"
          }
        ],
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "deviceOnlineCount": 1,
        "isSysCreated": 1,
        "position": {
          "latitude": 1,
          "longitude": 1
        },
        "ppsm": 1,
        "projectID": "string",
        "projectImg": "string",
        "projectName": "示例名称",
        "sort": 1,
        "status": 1,
        "tags": {},
        "tenantCode": "string",
        "type": "string",
        "userCount": 1
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/profile/get-one \
  --body '{"code": "string", "withProjects": true}'
```

### POST `/api/v1/system/user/self/profile/update`

**说明**: 更新用户配置

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `params` | string | 是 |  |

**请求示例**:
```json
{
  "code": "string",
  "params": "string"
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
ur api /api/v1/system/user/self/profile/update \
  --body '{"code": "string", "params": "string"}'
```

### POST `/api/v1/system/user/self/register`

**说明**: 普通用户注册

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 否 | 手机号注册时填写手机号 账号密码注册时填写userName |
| `code` | string | 是 | 验证码    微信登录填code 账号密码登录时填写密码 |
| `codeID` | string | 否 | 验证码编号 微信登录填state |
| `expand` | object | 否 | 拓展, 微信登录方式 phoneCode:获取手机号code  手机号和邮箱注册: wxOpenCode:如果需要同时绑定微信则在这里填写开放平台的用户code |
| `isWithLogin` | boolean | 否 | 注册成功同时返回token,自动登录 (格式: boolean) |
| `nickName` | string | 否 |  用户的昵称 |
| `password` | string | 否 | 密码 |
| `regType` | string | 是 | 可选: ['phone', 'email', 'wxOpen', 'wxIn', 'wxMiniP', 'pwd', 'dingApp'] |

**请求示例**:
```json
{
  "account": "******",
  "code": "******",
  "codeID": "string",
  "expand": {},
  "isWithLogin": true,
  "nickName": "string",
  "password": "******",
  "regType": "phone"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "info": {
      "avatar": "string",
      "email": "string",
      "nickName": "string",
      "phone": "string",
      "userID": "string",
      "userName": "string"
    },
    "token": {
      "accessExpire": "2026-01-01T00:00:00Z",
      "accessToken": "string",
      "refreshAfter": "2026-01-01T00:00:00Z"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/register \
  --body '{"account": "******", "code": "******", "codeID": "string", "expand": {}, "isWithLogin": true, "nickName": "string", "password": "******", "regType": "phone"}'
```

### POST `/api/v1/system/user/self/resource/action/get-list`

**说明**: 获取用户资源动作权限列表

**权限**: all

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "action": "string",
        "resource": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/resource/action/get-list \
  --body '{}'
```

### POST `/api/v1/system/user/self/third-auth/start`

**说明**: 第三方登录授权起跳

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `platform` | string | 否 |  |
| `provider` | string | 是 | 可选: ['google', 'github', 'apple'] |

**请求示例**:
```json
{
  "platform": "string",
  "provider": "google"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "authUrl": "string",
    "state": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/third-auth/start \
  --body '{"platform": "string", "provider": "google"}'
```

### POST `/api/v1/system/user/self/third-login`

**说明**: 第三方登录回调换平台登录态

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  |
| `platform` | string | 否 |  |
| `provider` | string | 是 | 可选: ['google', 'github', 'apple'] |
| `state` | string | 是 |  |

**请求示例**:
```json
{
  "code": "string",
  "platform": "string",
  "provider": "google",
  "state": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "login": {
      "info": {
        "avatar": "string",
        "email": "string",
        "nickName": "string",
        "phone": "string",
        "userID": "string",
        "userName": "string"
      },
      "token": {
        "accessExpire": "2026-01-01T00:00:00Z",
        "accessToken": "string",
        "refreshAfter": "2026-01-01T00:00:00Z"
      }
    },
    "message": "string",
    "pendingToken": "string",
    "provider": "string",
    "status": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/third-login \
  --body '{"code": "string", "platform": "string", "provider": "google", "state": "string"}'
```

### POST `/api/v1/system/user/self/third-register`

**说明**: 第三方补全注册

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 是 |  |
| `code` | string | 是 |  |
| `codeID` | string | 否 |  |
| `nickName` | string | 否 |  |
| `password` | string | 否 |  |
| `pendingToken` | string | 是 |  |
| `regType` | string | 是 | 可选: ['phone', 'email'] |

**请求示例**:
```json
{
  "account": "string",
  "code": "string",
  "codeID": "string",
  "nickName": "string",
  "password": "string",
  "pendingToken": "string",
  "regType": "phone"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "info": {
      "avatar": "string",
      "email": "string",
      "nickName": "string",
      "phone": "string",
      "userID": "string",
      "userName": "string"
    },
    "token": {
      "accessExpire": "2026-01-01T00:00:00Z",
      "accessToken": "string",
      "refreshAfter": "2026-01-01T00:00:00Z"
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/third-register \
  --body '{"account": "string", "code": "string", "codeID": "string", "nickName": "string", "password": "string", "pendingToken": "string", "regType": "phone"}'
```

### POST `/api/v1/system/user/self/update`

**说明**: 更新用户基本数据

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  头像（传图片的filePath） |
| `nickName` | string | 否 |  用户昵称 |
| `sex` | integer | 否 |  性别（1:男，2:女，0:未知） (格式: int64) |

**请求示例**:
```json
{
  "avatar": "string",
  "nickName": "string",
  "sex": 1
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
ur api /api/v1/system/user/self/update \
  --body '{"avatar": "string", "nickName": "string", "sex": 1}'
```

### POST `/api/v1/system/user/self/user/search`

**说明**: 精准搜索用户

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `account` | string | 是 |  |

**请求示例**:
```json
{
  "account": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avatar": "string",
    "nickName": "string",
    "status": 1,
    "userID": "string",
    "userName": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/user/search \
  --body '{"account": "string"}'
```

### POST `/api/v1/system/user/self/thirdparty/setup-check`

**说明**: 查询 CLI 绑定状态

**权限**: public

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `setupCode` | string | 是 |  CLI 生成的绑定码 |

**请求示例**:
```json
{
  "setupCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "accessKey": "string",
    "accessSecret": "string",
    "status": "string",
    "tenantCode": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/thirdparty/setup-check \
  --body '{"setupCode": "string"}'
```

### POST `/api/v1/system/user/self/thirdparty/setup-complete`

**说明**: 完成 CLI 绑定

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `accessKey` | string | 是 |  访问密钥标识（前端传入，用于写入 Redis） |
| `accessSecret` | string | 是 |  访问密钥（前端传入，用于写入 Redis） |
| `accessTokenID` | string | 是 |  访问令牌 ID |
| `setupCode` | string | 是 |  CLI 生成的绑定码 |

**请求示例**:
```json
{
  "accessKey": "string",
  "accessSecret": "string",
  "accessTokenID": "string",
  "setupCode": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "success": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/thirdparty/setup-complete \
  --body '{"accessKey": "string", "accessSecret": "string", "accessTokenID": "string", "setupCode": "string"}'
```

### POST `/api/v1/system/user/self/tenant/delete`

**说明**: 退出当前租户

**权限**: all

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
ur api /api/v1/system/user/self/tenant/delete \
  --body '{}'
```

### POST `/api/v1/system/user/self/tenant/get-list`

**说明**: 获取用户所处的租户列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `withRole` | boolean | 否 |  同时返回角色信息 (格式: boolean) |

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
  "withRole": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "deptIDs": [
          "string"
        ],
        "deviceCount": 1,
        "isTenantOwner": 1,
        "pubTags": {},
        "roles": [
          {
            "code": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "id": "string",
            "isFullAuth": 1,
            "name": "示例名称",
            "status": 1
          }
        ],
        "status": 1,
        "tags": {},
        "tenant": {
          "appIDs": [
            "string"
          ],
          "avatar": "string",
          "code": "string",
          "desc": "string",
          "id": "string",
          "name": "示例名称",
          "resource": {
            "package": {
              "expTime": "...",
              "packageCode": "...",
              "packageName": "..."
            },
            "quota": {
              "aiTokenBalance": "...",
              "aiTokenMonthly": "...",
              "deviceMsgPerDay": "...",
              "deviceNum": "...",
              "diskSpaceGB": "...",
              "userNum": "..."
            },
            "usage": {
              "aiTokenBalanceUsed": "...",
              "aiTokenMonth": "...",
              "aiTokenUsed": "...",
              "deviceCount": "...",
              "userCount": "..."
            }
          }
        },
        "tenantCode": "string",
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
        },
        "userID": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/tenant/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "withRole": true}'
```

### POST `/api/v1/system/user/self/tenant/get-one`

**说明**: 获取当前用户在当前租户的详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `withRole` | boolean | 否 | 同时返回角色信息 (格式: boolean) |
| `withUser` | boolean | 否 | 同时返回用户信息 (格式: boolean) |

**请求示例**:
```json
{
  "withRole": true,
  "withUser": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "createdTime": "2026-01-01T00:00:00Z",
    "deptIDs": [
      "string"
    ],
    "deviceCount": 1,
    "isTenantOwner": 1,
    "pubTags": {},
    "roles": [
      {
        "code": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "id": "string",
        "isFullAuth": 1,
        "name": "示例名称",
        "status": 1
      }
    ],
    "status": 1,
    "tags": {},
    "tenant": {
      "appIDs": [
        "string"
      ],
      "avatar": "string",
      "code": "string",
      "desc": "string",
      "id": "string",
      "name": "示例名称",
      "resource": {
        "package": {
          "expTime": "2026-01-01T00:00:00Z",
          "packageCode": "string",
          "packageName": "示例名称"
        },
        "quota": {
          "aiTokenBalance": 1,
          "aiTokenMonthly": 1,
          "deviceMsgPerDay": 1,
          "deviceNum": 1,
          "diskSpaceGB": 1,
          "userNum": 1
        },
        "usage": {
          "aiTokenBalanceUsed": 1,
          "aiTokenMonth": "string",
          "aiTokenUsed": 1,
          "deviceCount": 1,
          "userCount": 1
        }
      }
    },
    "tenantCode": "string",
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
    },
    "userID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/self/tenant/get-one \
  --body '{"withRole": true, "withUser": true}'
```

### POST `/api/v1/system/user/self/tenant/join`

**说明**: 用户加入租户（通过邀请码、邮件、手机或客户端应用）

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  邀请码（code/email/phone 方式时填写，client 方式可为空） |
| `method` | string | 否 |  加入方法：code（邀请码）、email（邮件邀请）、phone（手机邀请）、client（客户端应用直接加入） |
| `tenantCode` | string | 否 |  租户编码（client 方式时必填，其他方式可选） |

**请求示例**:
```json
{
  "code": "string",
  "method": "string",
  "tenantCode": "string"
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
ur api /api/v1/system/user/self/tenant/join \
  --body '{"code": "string", "method": "string", "tenantCode": "string"}'
```

### POST `/api/v1/system/user/self/tenant/update`

**说明**: 更新当前用户在当前租户的信息

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 |  创建时间,只读 |
| `deptIDs` | array[string] | 否 | 部门ID列表,默认展示第一个 |
| `deviceCount` | integer | 否 | 用户所拥有的设备数量统计,只读 (格式: int64) |
| `isTenantOwner` | integer | 否 | 是否是租户管理员,只读 (格式: int64) |
| `pubTags` | object | 否 |  公共的标签,用户自己可以修改 |
| `roles` | array[RoleInfo] | 否 |  角色列表 |
| `status` | integer | 否 | 用户状态(管理员才可修改): 1启用 2禁用 (格式: int64) |
| `tags` | object | 否 |  管理员才可修改的标签 |
| `tenant` | object | 否 |  |
| `tenant.appIDs` | array[string] | 否 |  租户拥有的应用ID列表 |
| `tenant.avatar` | string | 否 |  头像,更新时传filePath |
| `tenant.code` | string | 否 |  租户编码,创建的时候自动生成 |
| `tenant.desc` | string | 否 |  应用描述 |
| `tenant.id` | string | 否 |  id编号 |
| `tenant.name` | string | 否 |  租户名称 |
| `tenant.resource` | object | 否 |  |
| `tenant.resource.package` | object | 否 |  |
| `tenant.resource.quota` | object | 否 |  |
| `tenant.resource.usage` | object | 否 |  |
| `tenantCode` | string | 否 |  租户编码,只读 |
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
| `userID` | string | 否 |  用户id,只读 |

**请求示例**:
```json
{
  "createdTime": "2026-01-01T00:00:00Z",
  "deptIDs": [
    "string"
  ],
  "deviceCount": 1,
  "isTenantOwner": 1,
  "pubTags": {},
  "roles": [
    {
      "code": "string",
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "id": "string",
      "isFullAuth": 1,
      "name": "示例名称",
      "status": 1
    }
  ],
  "status": 1,
  "tags": {},
  "tenant": {
    "appIDs": [
      "string"
    ],
    "avatar": "string",
    "code": "string",
    "desc": "string",
    "id": "string",
    "name": "示例名称",
    "resource": {
      "package": {
        "expTime": "2026-01-01T00:00:00Z",
        "packageCode": "string",
        "packageName": "示例名称"
      },
      "quota": {
        "aiTokenBalance": 1,
        "aiTokenMonthly": 1,
        "deviceMsgPerDay": 1,
        "deviceNum": 1,
        "diskSpaceGB": 1,
        "userNum": 1
      },
      "usage": {
        "aiTokenBalanceUsed": 1,
        "aiTokenMonth": "string",
        "aiTokenUsed": 1,
        "deviceCount": 1,
        "userCount": 1
      }
    }
  },
  "tenantCode": "string",
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
  },
  "userID": "string"
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
ur api /api/v1/system/user/self/tenant/update \
  --body '{"createdTime": "2026-01-01T00:00:00Z", "deptIDs": ["string"], "deviceCount": 1, "isTenantOwner": 1, "pubTags": {}, "roles": [{"code": "string", "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "id": "string", "isFullAuth": 1, "name": "示例名称", "status": 1}], "status": 1, "tags": {}, "tenant": {"appIDs": ["string"], "avatar": "string", "code": "string", "desc": "string", "id": "string", "name": "示例名称", "resource": {"package": {"expTime": "2026-01-01T00:00:00Z", "packageCode": "string", "packageName": "示例名称"}, "quota": {"aiTokenBalance": 1, "aiTokenMonthly": 1, "deviceMsgPerDay": 1, "deviceNum": 1, "diskSpaceGB": 1, "userNum": 1}, "usage": {"aiTokenBalanceUsed": 1, "aiTokenMonth": "string", "aiTokenUsed": 1, "deviceCount": 1, "userCount": 1}}}, "tenantCode": "string", "user": {"avatar": "string", "createdTime": "2026-01-01T00:00:00Z", "email": "string", "lastIP": "string", "nickName": "string", "phone": "string", "regIP": "string", "userID": "string", "userName": "string"}, "userID": "string"}'
```

### POST `/api/v1/system/user/tenant/get-list`

**说明**: 用户所处的租户列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 |  用户状态（1:启用，2:禁用） (格式: int64) |
| `userID` | string | 是 |  |
| `withRole` | boolean | 否 |  同时返回角色信息 (格式: boolean) |

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
  "status": 1,
  "userID": "string",
  "withRole": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "2026-01-01T00:00:00Z",
        "deptIDs": [
          "string"
        ],
        "deviceCount": 1,
        "isTenantOwner": 1,
        "pubTags": {},
        "roles": [
          {
            "code": "string",
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "id": "string",
            "isFullAuth": 1,
            "name": "示例名称",
            "status": 1
          }
        ],
        "status": 1,
        "tags": {},
        "tenant": {
          "appIDs": [
            "string"
          ],
          "avatar": "string",
          "code": "string",
          "desc": "string",
          "id": "string",
          "name": "示例名称",
          "resource": {
            "package": {
              "expTime": "...",
              "packageCode": "...",
              "packageName": "..."
            },
            "quota": {
              "aiTokenBalance": "...",
              "aiTokenMonthly": "...",
              "deviceMsgPerDay": "...",
              "deviceNum": "...",
              "diskSpaceGB": "...",
              "userNum": "..."
            },
            "usage": {
              "aiTokenBalanceUsed": "...",
              "aiTokenMonth": "...",
              "aiTokenUsed": "...",
              "deviceCount": "...",
              "userCount": "..."
            }
          }
        },
        "tenantCode": "string",
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
        },
        "userID": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/user/tenant/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "userID": "string", "withRole": true}'
```
