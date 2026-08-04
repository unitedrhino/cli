# ur-tenant system/tenant/app/menu

获取租户应用菜单列表 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/tenant/app/menu/get-list` | 获取租户应用菜单列表 | admin |
| POST | `/api/v1/system/tenant/app/menu/update` | 更新租户应用菜单 | admin |

## 详细说明

### POST `/api/v1/system/tenant/app/menu/get-list`

**说明**: 获取租户应用菜单列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 是 |  |
| `isRetTree` | boolean | 否 | 格式: boolean |
| `isStar` | integer | 否 |  是否收藏 1:是 2:否 (格式: int64) |
| `menuIDs` | array[string] | 否 |  |
| `tenantCode` | string | 否 | 指定租户,只有平台可以指定 |

**请求示例**:
```json
{
  "appID": "string",
  "isRetTree": true,
  "isStar": 1,
  "menuIDs": [
    "string"
  ],
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
ur api /api/v1/system/tenant/app/menu/get-list \
  --body '{"appID": "string", "isRetTree": true, "isStar": 1, "menuIDs": ["string"], "tenantCode": "string"}'
```

### POST `/api/v1/system/tenant/app/menu/update`

**说明**: 更新租户应用菜单

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appID` | string | 否 |  应用ID,新增必填,不可修改 |
| `children` | array[TenantMenu] | 否 |  |
| `icon` | string | 否 |  图标 |
| `id` | string | 否 |  编号 |
| `isHideInMenu` | integer | 否 |  是否隐藏菜单 1-是 2-否 (格式: int64) |
| `isStar` | integer | 否 | 是否收藏 1:是 2:否 (格式: int64) |
| `menu` | object | 否 |  |
| `menu.appID` | string | 否 |  模块编码 |
| `menu.body` | string | 否 |  前端自定义字段 |
| `menu.children` | array[MenuInfo] | 否 |  子菜单列表 |
| `menu.component` | string | 否 |  页面 |
| `menu.createdTime` | string | 否 |  创建时间 |
| `menu.defaultIcon` | string | 否 |  默认图标 |
| `menu.defaultIsHideInMenu` | integer | 否 |  默认是否隐藏菜单 1-是 2-否 (格式: int64) |
| `menu.defaultName` | string | 否 |  默认 菜单名称 |
| `menu.defaultSort` | integer | 否 |  默认左侧table排序序号 (格式: int64) |
| `menu.id` | string | 否 |  编号 |
| `menu.idPath` | string | 否 |  菜单ID路径 |
| `menu.parentID` | string | 否 |  父菜单ID，一级菜单为1 |
| `menu.path` | string | 否 |  系统的path |
| `menu.redirect` | string | 否 |  路由重定向 |
| `menu.role` | string | 否 |  角色类型 仅App.menuRole=all时有意义 all:两类都可以看 platform:平台管理员 tenant:租户管理员 |
| `menu.tenantCode` | string | 否 |  跟着应用走, common的是通用应用,非common的是租户应用 |
| `menu.type` | integer | 否 |  类型 1：菜单或者页面 2：iframe嵌入 3：外链跳转 (格式: int64) |
| `menuID` | string | 否 |  菜单ID,不可修改 |
| `name` | string | 否 |  菜单名称 |
| `parentMenuID` | string | 否 |  父菜单ID，一级菜单为1 |
| `sort` | integer | 否 |  左侧table排序序号 (格式: int64) |

**请求示例**:
```json
{
  "appID": "string",
  "children": [
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
ur api /api/v1/system/tenant/app/menu/update \
  --body '{"appID": "string", "children": [{"appID": "string", "children": [{"appID": "string", "children": [], "icon": "string", "id": "string", "isHideInMenu": 1, "isStar": 1, "menu": {"appID": "...", "body": "...", "children": "...", "component": "...", "createdTime": "...", "defaultIcon": "...", "defaultIsHideInMenu": "...", "defaultName": "...", "defaultSort": "...", "id": "...", "idPath": "...", "parentID": "...", "path": "...", "redirect": "...", "role": "...", "tenantCode": "...", "type": "..."}, "menuID": "string", "name": "示例名称", "parentMenuID": "string", "sort": 1}], "icon": "string", "id": "string", "isHideInMenu": 1, "isStar": 1, "menu": {"appID": "string", "body": "string", "children": [{"appID": "...", "body": "...", "children": "...", "component": "...", "createdTime": "...", "defaultIcon": "...", "defaultIsHideInMenu": "...", "defaultName": "...", "defaultSort": "...", "id": "...", "idPath": "...", "parentID": "...", "path": "...", "redirect": "...", "role": "...", "tenantCode": "...", "type": "..."}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}, "menuID": "string", "name": "示例名称", "parentMenuID": "string", "sort": 1}], "icon": "string", "id": "string", "isHideInMenu": 1, "isStar": 1, "menu": {"appID": "string", "body": "string", "children": [{"appID": "string", "body": "string", "children": [{"appID": "...", "body": "...", "children": "...", "component": "...", "createdTime": "...", "defaultIcon": "...", "defaultIsHideInMenu": "...", "defaultName": "...", "defaultSort": "...", "id": "...", "idPath": "...", "parentID": "...", "path": "...", "redirect": "...", "role": "...", "tenantCode": "...", "type": "..."}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}], "component": "string", "createdTime": "2026-01-01T00:00:00Z", "defaultIcon": "string", "defaultIsHideInMenu": 1, "defaultName": "示例名称", "defaultSort": 1, "id": "string", "idPath": "string", "parentID": "string", "path": "string", "redirect": "string", "role": "string", "tenantCode": "string", "type": 1}, "menuID": "string", "name": "示例名称", "parentMenuID": "string", "sort": 1}'
```
