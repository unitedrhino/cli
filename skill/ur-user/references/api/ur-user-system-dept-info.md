# ur-user system/dept/info

添加部门详情 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/dept/info/create` | 添加部门详情 | admin |
| POST | `/api/v1/system/dept/info/delete` | 删除部门 | admin |
| POST | `/api/v1/system/dept/info/get-list` | 获取部门列表 | admin |
| POST | `/api/v1/system/dept/info/get-one` | 获取部门详情 | admin |
| POST | `/api/v1/system/dept/info/update` | 更新部门 | admin |

## 详细说明

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
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
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
    "page": 1,
    "pageSize": 1
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
  --body '{"dingTalkIDs": ["string"], "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "parentID": "string", "status": 1, "tenantCode": "string"}'
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
