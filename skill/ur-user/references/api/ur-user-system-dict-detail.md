# ur-user system/dict/detail

添加字典详情 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/dict/detail/create` | 添加字典详情 | platform |
| POST | `/api/v1/system/dict/detail/delete` | 删除字典详情 | platform |
| POST | `/api/v1/system/dict/detail/get-list` | 获取字典详情列表 | platform |
| POST | `/api/v1/system/dict/detail/get-one` | 获取字典详情单个 | platform |
| POST | `/api/v1/system/dict/detail/update` | 更新字典详情 | platform |

## 详细说明

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
| `id` | string | 否 |  编号 |
| `idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `label` | string | 否 |  展示值 |
| `parent` | object | 否 |  |
| `parent.children` | array[DictDetail] | 否 |  |
| `parent.desc` | string | 否 |  模块描述 |
| `parent.dictCode` | string | 是 |  |
| `parent.expand` | object | 否 |  自定义数据 |
| `parent.id` | string | 否 |  编号 |
| `parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `parent.label` | string | 否 |  展示值 |
| `parent.parent` | object | 否 |  |
| `parent.parent.children` | array[DictDetail] | 否 |  |
| `parent.parent.desc` | string | 否 |  模块描述 |
| `parent.parent.dictCode` | string | 是 |  |
| `parent.parent.expand` | object | 否 |  自定义数据 |
| `parent.parent.id` | string | 否 |  编号 |
| `parent.parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.parent.isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `parent.parent.label` | string | 否 |  展示值 |
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
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
  --body '{"children": [{"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}'
```

### POST `/api/v1/system/dict/detail/delete`

**说明**: 删除字典详情

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
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
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
    "page": 1,
    "pageSize": 1
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
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
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
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
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
  --body '{"dictCode": "string", "label": "string", "page": {"page": 1, "pageSize": 1}, "parentID": "string", "status": 1, "value": "string", "values": ["string"]}'
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
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
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
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
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
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
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
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
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
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": [],
            "desc": "string",
            "dictCode": "string",
            "expand": {},
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "label": "string",
            "parent": {
              "children": "...",
              "desc": "...",
              "dictCode": "...",
              "expand": "...",
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
| `id` | string | 否 |  编号 |
| `idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `label` | string | 否 |  展示值 |
| `parent` | object | 否 |  |
| `parent.children` | array[DictDetail] | 否 |  |
| `parent.desc` | string | 否 |  模块描述 |
| `parent.dictCode` | string | 是 |  |
| `parent.expand` | object | 否 |  自定义数据 |
| `parent.id` | string | 否 |  编号 |
| `parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `parent.label` | string | 否 |  展示值 |
| `parent.parent` | object | 否 |  |
| `parent.parent.children` | array[DictDetail] | 否 |  |
| `parent.parent.desc` | string | 否 |  模块描述 |
| `parent.parent.dictCode` | string | 是 |  |
| `parent.parent.expand` | object | 否 |  自定义数据 |
| `parent.parent.id` | string | 否 |  编号 |
| `parent.parent.idPath` | string | 否 | 1-2-3-的格式记录顶级区域到当前id的路径 |
| `parent.parent.isLeaf` | integer | 否 |  是否是叶子节点(不可修改) 1:是 2:否 (格式: int64) |
| `parent.parent.label` | string | 否 |  展示值 |
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
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "label": "string",
        "parent": {
          "children": [],
          "desc": "string",
          "dictCode": "string",
          "expand": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "label": "string",
          "parent": {
            "children": "...",
            "desc": "...",
            "dictCode": "...",
            "expand": "...",
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
  --body '{"children": [{"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}'
```
