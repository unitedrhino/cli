# ur-system system/app/menu

添加菜单 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/app/menu/create` | 添加菜单 | platform |
| POST | `/api/v1/system/app/menu/delete` | 删除菜单 | platform |
| POST | `/api/v1/system/app/menu/get-list` | 获取菜单列表 | platform |
| POST | `/api/v1/system/app/menu/update` | 更新菜单 | platform |

## 详细说明

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
| `role` | string | 否 |  角色类型 仅App.menuRole=all时有意义 all:两类都可以看 platform:平台管理员 tenant:企业管理员 |
| `tenantCode` | string | 否 |  跟着应用走, common的是通用应用,非common的是企业应用 |
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
| `role` | string | 否 |  角色类型 仅App.menuRole=all时有意义 all:两类都可以看 platform:平台管理员 tenant:企业管理员 |
| `tenantCode` | string | 否 |  跟着应用走, common的是通用应用,非common的是企业应用 |
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
