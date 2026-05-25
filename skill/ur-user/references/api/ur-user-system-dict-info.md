# ur-user system/dict/info

批量导出字典信息 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/dict/info/batch-export` | 批量导出字典信息 | platform |
| POST | `/api/v1/system/dict/info/batch-import` | 批量导入字典信息 | platform |
| POST | `/api/v1/system/dict/info/create` | 添加字典信息 | platform |
| POST | `/api/v1/system/dict/info/delete` | 删除字典信息 | platform |
| POST | `/api/v1/system/dict/info/get-list` | 获取字典信息列表 | all |
| POST | `/api/v1/system/dict/info/get-one` | 获取字典信息详情 | platform |
| POST | `/api/v1/system/dict/info/update` | 更新字典信息 | platform |

## 详细说明

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
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/dict/info/create \
  --body '{"code": "string", "desc": "string", "details": [{"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "expand": {}, "group": "string", "id": "string", "name": "示例名称", "structType": 1}'
```

### POST `/api/v1/system/dict/info/delete`

**说明**: 删除字典信息

**权限**: platform

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
ur api /api/v1/system/dict/info/delete \
  --body '{"id": "string"}'
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
  --body '{"code": "string", "desc": "string", "details": [{"children": [{"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [{"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": [], "desc": "string", "dictCode": "string", "expand": {}, "id": "string", "idPath": "string", "isLeaf": 1, "label": "string", "parent": {"children": "...", "desc": "...", "dictCode": "...", "expand": "...", "id": "...", "idPath": "...", "isLeaf": "...", "label": "...", "parent": "...", "parentID": "...", "sort": "...", "status": "...", "value": "..."}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}, "parentID": "string", "sort": 1, "status": 1, "value": "string"}], "expand": {}, "group": "string", "id": "string", "name": "示例名称", "structType": 1}'
```
