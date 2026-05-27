# ur-product things/product/category

批量导出产品品类 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/product/category/batch-export` | 批量导出产品品类 | admin |
| POST | `/api/v1/things/product/category/batch-import` | 批量导入产品品类 | admin |
| POST | `/api/v1/things/product/category/create` | 新增产品品类 | admin |
| POST | `/api/v1/things/product/category/delete` | 删除产品品类 | admin |
| POST | `/api/v1/things/product/category/get-list` | 获取产品品类列表 | admin |
| POST | `/api/v1/things/product/category/get-one` | 获取产品品类详情 | admin |
| POST | `/api/v1/things/product/category/schema/batch-create` | 批量新增产品品类物模型 | admin |
| POST | `/api/v1/things/product/category/schema/batch-delete` | 批量删除产品品类物模型 | admin |
| POST | `/api/v1/things/product/category/schema/batch-update` | 批量更新产品品类物模型 | admin |
| POST | `/api/v1/things/product/category/schema/get-list` | 获取产品品类物模型列表 | admin |
| POST | `/api/v1/things/product/category/update` | 更新产品品类 | admin |

## 详细说明

### POST `/api/v1/things/product/category/batch-export`

**说明**: 批量导出产品品类

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `ids` | array[string] | 否 |  |

**请求示例**:
```json
{
  "ids": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "categoriesFile": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/category/batch-export \
  --body '{"ids": ["string"]}'
```

### POST `/api/v1/things/product/category/batch-import`

**说明**: 批量导入产品品类

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `categoriesFile` | string | 是 | 品类数据文件的 OSS 地址（ZIP 压缩包） |

**请求示例**:
```json
{
  "categoriesFile": "string"
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
ur api /api/v1/things/product/category/batch-import \
  --body '{"categoriesFile": "string"}'
```

### POST `/api/v1/things/product/category/create`

**说明**: 新增产品品类

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  品类图片 |
| `children` | array[ProductCategory] | 否 |  |
| `code` | string | 否 | 品类编码,全局唯一,创建后不可修改,创建时必填 |
| `desc` | string | 否 | 描述 |
| `deviceCount` | integer | 否 | 设备数量 (格式: int64) |
| `id` | string | 否 |  |
| `idPath` | array[integer] | 否 | 只读 |
| `isLeaf` | integer | 否 | 格式: int64 |
| `name` | string | 否 | 品类名称 |
| `order` | integer | 否 | 排序,必须大于等于1,为0代表不修改 (格式: int64) |
| `parentID` | string | 否 |  |

**请求示例**:
```json
{
  "avatar": "string",
  "children": [
    {
      "avatar": "string",
      "children": [
        {
          "avatar": "string",
          "children": [],
          "code": "string",
          "desc": "string",
          "deviceCount": 1,
          "id": "string",
          "idPath": [],
          "isLeaf": 1,
          "name": "示例名称",
          "order": 1,
          "parentID": "string"
        }
      ],
      "code": "string",
      "desc": "string",
      "deviceCount": 1,
      "id": "string",
      "idPath": [
        1
      ],
      "isLeaf": 1,
      "name": "示例名称",
      "order": 1,
      "parentID": "string"
    }
  ],
  "code": "string",
  "desc": "string",
  "deviceCount": 1,
  "id": "string",
  "idPath": [
    1
  ],
  "isLeaf": 1,
  "name": "示例名称",
  "order": 1,
  "parentID": "string"
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
ur api /api/v1/things/product/category/create \
  --body '{"avatar": "string", "children": [{"avatar": "string", "children": [{"avatar": "string", "children": [], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}'
```

### POST `/api/v1/things/product/category/delete`

**说明**: 删除产品品类

**权限**: admin

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
ur api /api/v1/things/product/category/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/product/category/get-list`

**说明**: 获取产品品类列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 | 按品类编码过滤 |
| `ids` | array[integer] | 否 | id过滤 |
| `name` | string | 否 | 过滤产品名称 |
| `page` | object | 否 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 |  |
| `projectID` | string | 否 | 项目id,只获取项目下有设备的 |

**请求示例**:
```json
{
  "code": "string",
  "ids": [
    1
  ],
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "parentID": "string",
  "projectID": "string"
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
        "children": [
          {
            "avatar": "string",
            "children": [],
            "code": "string",
            "desc": "string",
            "deviceCount": 1,
            "id": "string",
            "idPath": [],
            "isLeaf": 1,
            "name": "示例名称",
            "order": 1,
            "parentID": "string"
          }
        ],
        "code": "string",
        "desc": "string",
        "deviceCount": 1,
        "id": "string",
        "idPath": [
          1
        ],
        "isLeaf": 1,
        "name": "示例名称",
        "order": 1,
        "parentID": "string"
      }
    ],
    "num": 1,
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/category/get-list \
  --body '{"code": "string", "ids": [1], "name": "示例名称", "page": {"page": 1, "pageSize": 1}, "parentID": "string", "projectID": "string"}'
```

### POST `/api/v1/things/product/category/get-one`

**说明**: 获取产品品类详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  id |
| `withChildren` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "id": "string",
  "withChildren": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avatar": "string",
    "children": [
      {
        "avatar": "string",
        "children": [
          {
            "avatar": "string",
            "children": [],
            "code": "string",
            "desc": "string",
            "deviceCount": 1,
            "id": "string",
            "idPath": [],
            "isLeaf": 1,
            "name": "示例名称",
            "order": 1,
            "parentID": "string"
          }
        ],
        "code": "string",
        "desc": "string",
        "deviceCount": 1,
        "id": "string",
        "idPath": [
          1
        ],
        "isLeaf": 1,
        "name": "示例名称",
        "order": 1,
        "parentID": "string"
      }
    ],
    "code": "string",
    "desc": "string",
    "deviceCount": 1,
    "id": "string",
    "idPath": [
      1
    ],
    "isLeaf": 1,
    "name": "示例名称",
    "order": 1,
    "parentID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/category/get-one \
  --body '{"id": "string", "withChildren": true}'
```

### POST `/api/v1/things/product/category/schema/batch-create`

**说明**: 批量新增产品品类物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `identifiers` | array[string] | 是 |  |
| `productCategoryID` | string | 是 |  |

**请求示例**:
```json
{
  "identifiers": [
    "string"
  ],
  "productCategoryID": "string"
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
ur api /api/v1/things/product/category/schema/batch-create \
  --body '{"identifiers": ["string"], "productCategoryID": "string"}'
```

### POST `/api/v1/things/product/category/schema/batch-delete`

**说明**: 批量删除产品品类物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `identifiers` | array[string] | 是 |  |
| `productCategoryID` | string | 是 |  |

**请求示例**:
```json
{
  "identifiers": [
    "string"
  ],
  "productCategoryID": "string"
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
ur api /api/v1/things/product/category/schema/batch-delete \
  --body '{"identifiers": ["string"], "productCategoryID": "string"}'
```

### POST `/api/v1/things/product/category/schema/batch-update`

**说明**: 批量更新产品品类物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `identifiers` | array[string] | 是 |  |
| `productCategoryID` | string | 是 |  |

**请求示例**:
```json
{
  "identifiers": [
    "string"
  ],
  "productCategoryID": "string"
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
ur api /api/v1/things/product/category/schema/batch-update \
  --body '{"identifiers": ["string"], "productCategoryID": "string"}'
```

### POST `/api/v1/things/product/category/schema/get-list`

**说明**: 获取产品品类物模型列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productCategoryID` | string | 是 |  |
| `withFather` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "productCategoryID": "string",
  "withFather": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "identifiers": [
      "string"
    ],
    "list": [
      {
        "identifier": "string",
        "name": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/category/schema/get-list \
  --body '{"productCategoryID": "string", "withFather": true}'
```

### POST `/api/v1/things/product/category/update`

**说明**: 更新产品品类

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `avatar` | string | 否 |  品类图片 |
| `children` | array[ProductCategory] | 否 |  |
| `code` | string | 否 | 品类编码,全局唯一,创建后不可修改,创建时必填 |
| `desc` | string | 否 | 描述 |
| `deviceCount` | integer | 否 | 设备数量 (格式: int64) |
| `id` | string | 否 |  |
| `idPath` | array[integer] | 否 | 只读 |
| `isLeaf` | integer | 否 | 格式: int64 |
| `name` | string | 否 | 品类名称 |
| `order` | integer | 否 | 排序,必须大于等于1,为0代表不修改 (格式: int64) |
| `parentID` | string | 否 |  |

**请求示例**:
```json
{
  "avatar": "string",
  "children": [
    {
      "avatar": "string",
      "children": [
        {
          "avatar": "string",
          "children": [],
          "code": "string",
          "desc": "string",
          "deviceCount": 1,
          "id": "string",
          "idPath": [],
          "isLeaf": 1,
          "name": "示例名称",
          "order": 1,
          "parentID": "string"
        }
      ],
      "code": "string",
      "desc": "string",
      "deviceCount": 1,
      "id": "string",
      "idPath": [
        1
      ],
      "isLeaf": 1,
      "name": "示例名称",
      "order": 1,
      "parentID": "string"
    }
  ],
  "code": "string",
  "desc": "string",
  "deviceCount": 1,
  "id": "string",
  "idPath": [
    1
  ],
  "isLeaf": 1,
  "name": "示例名称",
  "order": 1,
  "parentID": "string"
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
ur api /api/v1/things/product/category/update \
  --body '{"avatar": "string", "children": [{"avatar": "string", "children": [{"avatar": "string", "children": [], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}'
```
