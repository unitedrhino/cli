# 全部端点速查

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
| POST | `/api/v1/things/product/config/update` | 更新配置 | admin |
| POST | `/api/v1/things/product/custom/get-one` | 获取产品自定义信息详情 | admin |
| POST | `/api/v1/things/product/custom/update` | 更新产品自定义信息 | admin |
| POST | `/api/v1/things/product/info/batch-export` | 批量导出产品 | admin |
| POST | `/api/v1/things/product/info/batch-import` | 批量导入产品 | admin |
| POST | `/api/v1/things/product/info/create` | 新增产品 | admin |
| POST | `/api/v1/things/product/info/delete` | 删除产品 | admin |
| POST | `/api/v1/things/product/info/get-list` | 获取产品信息列表 | admin |
| POST | `/api/v1/things/product/info/get-one` | 获取产品详情 | admin |
| POST | `/api/v1/things/product/info/init` | 初始化产品 | admin |
| POST | `/api/v1/things/product/info/update` | 更新产品 | admin |
| POST | `/api/v1/things/product/remote-config/create` | 创建配置 | admin |
| POST | `/api/v1/things/product/remote-config/get-list` | 获取配置列表 | admin |
| POST | `/api/v1/things/product/remote-config/lastest-read` | 获取最新配置 | admin |
| POST | `/api/v1/things/product/remote-config/push-all` | 推送配置 | admin |
| POST | `/api/v1/things/product/schema/batch-create` | 批量创建产品物模型 | admin |
| POST | `/api/v1/things/product/schema/create` | 创建产品物模型 | admin |
| POST | `/api/v1/things/product/schema/delete` | 删除产品物模型 | admin |
| POST | `/api/v1/things/product/schema/get-list` | 获取产品物模型 | admin |
| POST | `/api/v1/things/product/schema/tsl-import` | 导入产品物模型tsl | admin |
| POST | `/api/v1/things/product/schema/tsl-read` | 获取产品物模型tsl | admin |
| POST | `/api/v1/things/product/schema/update` | 更新产品物模型 | admin |

## 端点详情

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
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
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
  --body '{"code": "string", "ids": [1], "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentID": "string", "projectID": "string"}'
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

### POST `/api/v1/things/product/config/update`

**说明**: 更新配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `devInit` | object | 否 |  |
| `devInit.areaID` | string | 否 | 设备创建默认分配给的区域ID |
| `devInit.deptID` | string | 否 | 设备创建默认分配给的部门 |
| `devInit.projectID` | string | 否 | 设备创建默认分配给的项目 |
| `devInit.tenantCode` | string | 否 | 设备创建默认分配给的租户 |
| `productID` | string | 是 |  |

**请求示例**:
```json
{
  "devInit": {
    "areaID": "string",
    "deptID": "string",
    "projectID": "string",
    "tenantCode": "string"
  },
  "productID": "string"
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
ur api /api/v1/things/product/config/update \
  --body '{"devInit": {"areaID": "string", "deptID": "string", "projectID": "string", "tenantCode": "string"}, "productID": "string"}'
```

### POST `/api/v1/things/product/custom/get-one`

**说明**: 获取产品自定义信息详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id 只读 |

**请求示例**:
```json
{
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "customTopics": [
      {
        "direction": 1,
        "topic": "string"
      }
    ],
    "loginAuthScript": "string",
    "productID": "string",
    "scriptLang": 1,
    "transformScript": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/custom/get-one \
  --body '{"productID": "string"}'
```

### POST `/api/v1/things/product/custom/update`

**说明**: 更新产品自定义信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `customTopics` | array[ProductCustomTopic] | 否 | 自定义topic数组 |
| `loginAuthScript` | string | 否 | 登录认证脚本 |
| `productID` | string | 是 |  |
| `scriptLang` | integer | 否 | 脚本语言类型（默认JavaScript） 1:JavaScript 2:lua 3:python (格式: int64) |
| `transformScript` | string | 否 | 协议转换脚本 |

**请求示例**:
```json
{
  "customTopics": [
    {
      "direction": 1,
      "topic": "string"
    }
  ],
  "loginAuthScript": "string",
  "productID": "string",
  "scriptLang": 1,
  "transformScript": "string"
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
ur api /api/v1/things/product/custom/update \
  --body '{"customTopics": [{"direction": 1, "topic": "string"}], "loginAuthScript": "string", "productID": "string", "scriptLang": 1, "transformScript": "string"}'
```

### POST `/api/v1/things/product/info/batch-export`

**说明**: 批量导出产品

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productIDs` | array[string] | 否 |  |

**请求示例**:
```json
{
  "productIDs": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "productsFile": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/info/batch-export \
  --body '{"productIDs": ["string"]}'
```

### POST `/api/v1/things/product/info/batch-import`

**说明**: 批量导入产品

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productsFile` | string | 是 | 产品数据文件的 OSS 地址（ZIP 压缩包） |

**请求示例**:
```json
{
  "productsFile": "string"
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
ur api /api/v1/things/product/info/batch-import \
  --body '{"productsFile": "string"}'
```

### POST `/api/v1/things/product/info/create`

**说明**: 新增产品

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authMode` | integer | 否 | 认证方式:1:账密认证,2:秘钥认证 (格式: int64) |
| `autoRegister` | integer | 否 | 动态注册:1:关闭,2:打开,3:打开并自动创建设备,4:在前面的基础上绑定没有也自动创建 (格式: int64) |
| `bindLevel` | integer | 否 | 绑定级别: 1:强绑定(默认,只有用户解绑之后才能绑定) 2:中绑定(可以通过token强制解绑设备) 3:弱绑定(app可以内部解绑被绑定的设备) (格式: int64) |
| `category` | object | 否 |  |
| `category.avatar` | string | 否 |  品类图片 |
| `category.children` | array[ProductCategory] | 否 |  |
| `category.code` | string | 否 | 品类编码,全局唯一,创建后不可修改,创建时必填 |
| `category.desc` | string | 否 | 描述 |
| `category.deviceCount` | integer | 否 | 设备数量 (格式: int64) |
| `category.id` | string | 否 |  |
| `category.idPath` | array[integer] | 否 | 只读 |
| `category.isLeaf` | integer | 否 | 格式: int64 |
| `category.name` | string | 否 | 品类名称 |
| `category.order` | integer | 否 | 排序,必须大于等于1,为0代表不修改 (格式: int64) |
| `category.parentID` | string | 否 |  |
| `categoryID` | string | 否 | 产品品类 |
| `compatMode` | integer | 否 | 协议兼容模式 1:旧productID(大写D) 2:新productId(小写d,默认) (格式: int64) |
| `config` | object | 否 |  |
| `config.devInit` | object | 否 |  |
| `config.devInit.areaID` | string | 否 | 设备创建默认分配给的区域ID |
| `config.devInit.deptID` | string | 否 | 设备创建默认分配给的部门 |
| `config.devInit.projectID` | string | 否 | 设备创建默认分配给的项目 |
| `config.devInit.tenantCode` | string | 否 | 设备创建默认分配给的租户 |
| `config.productID` | string | 是 |  |
| `createdTime` | string | 否 | 创建时间 只读 |
| `customUi` | object | 否 | 自定义ui,key是端的类型(web-client  mini-client) value是以下类型的对象{version:123(版本号,只读),isUpdateUi:bool(是否更新ui),path:string(前端路径,如果需要修改,需要将isUpdateUi置为true并在这个参数中传入压缩包的filePath)} |
| `defaultAgentId` | integer | 否 | 默认绑定的 Agent ID，0 表示未绑定 (格式: int64) |
| `desc` | string | 否 | 描述 |
| `deviceSchemaMode` | integer | 否 |  设备物模型模式:1:手动创建,2:设备自动创建 3: 设备自动创建及上报无定义自动创建 4: 设备自动创建及上报无定义自动创建(数字类型只使用浮点) (格式: int64) |
| `deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `onlineHandle` | integer | 否 | 在线处理:1: 自动 2: 永远在线 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `productImg` | string | 否 | 产品图片 |
| `productName` | string | 否 | 产品名称 |
| `protocol` | object | 否 |  |
| `protocol.code` | string | 否 |   iThings,iThings-thingsboard,wumei,aliyun,huaweiyun,tuya |
| `protocol.configFields` | array[ProtocolConfigField] | 否 | 配置字段列表,没有可以不传 |
| `protocol.containerPort` | integer | 否 |  容器端口 (格式: int64) |
| `protocol.deployStatus` | string | 否 |  部署状态: notInstalled/downloaded/running/stopped |
| `protocol.desc` | string | 否 |  |
| `protocol.deviceFields` | array[ProtocolConfigField] | 否 | 设备级的配置字段列表,没有可以不传 |
| `protocol.endpoints` | array[string] | 否 |  |
| `protocol.envVars` | object | 否 |  环境变量 |
| `protocol.etcdKey` | string | 否 |  |
| `protocol.id` | string | 否 |  |
| `protocol.image` | string | 否 |  镜像地址 |
| `protocol.isBuiltin` | integer | 否 |  是否内置 1:是 2:否 (格式: int64) |
| `protocol.isEnableSyncDevice` | integer | 是 | 是否支持主动同步设备信息 (格式: int64) |
| `protocol.isEnableSyncProduct` | integer | 是 | 是否支持主动同步产品信息 (格式: int64) |
| `protocol.name` | string | 否 |  |
| `protocol.productFields` | array[ProtocolConfigField] | 否 | 产品级的配置字段列表,没有可以不传 |
| `protocol.transProtocol` | string | 否 |  传输协议: mqtt,tcp,udp |
| `protocol.type` | string | 否 | 协议类型: 普通设备(默认):normal  音视频:media |
| `protocol.version` | string | 否 |  镜像版本 |
| `protocolCode` | string | 否 | 协议code,默认urMqtt,设备下发只会发送给主协议  urMqtt,urHttp,wumei,aliyun,huaweiyun,tuya |
| `protocolConf` | object | 否 | 协议配置 |
| `sceneMode` | string | 否 | 场景模式 读写类型: r(只读) rw(可读可写) none(不参与场景) |
| `secret` | string | 否 | 动态注册产品秘钥 只读 |
| `status` | integer | 否 | 产品状态 1:启用 2:禁用 3:开发中 (格式: int64) |
| `subProtocolCode` | string | 否 | 子协议,主协议和子协议传输类型必须不相同, 设备控制下发只会发送给主协议, 当设备是音视频设备但是控制协议需要单独走的时候就可以把主协议定义为普通协议,子协议定义为音视频协议,这样就能实现音视频走音视频协议,控制走子协议 |
| `subProtocolConf` | object | 否 | 子协议协议配置 |
| `tags` | object | 否 |  产品tag |
| `tenantCode` | string | 否 | 租户号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `trialTime` | string | 否 | 试用时间(单位为天,为0不限制) |

**请求示例**:
```json
{
  "authMode": 1,
  "autoRegister": 1,
  "bindLevel": 1,
  "category": {
    "avatar": "string",
    "children": [
      {
        "avatar": "string",
        "children": [
          {
            "avatar": "...",
            "children": "...",
            "code": "...",
            "desc": "...",
            "deviceCount": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "name": "...",
            "order": "...",
            "parentID": "..."
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
  "categoryID": "string",
  "compatMode": 1,
  "config": {
    "devInit": {
      "areaID": "string",
      "deptID": "string",
      "projectID": "string",
      "tenantCode": "string"
    },
    "productID": "string"
  },
  "createdTime": "2026-01-01T00:00:00Z",
  "customUi": {},
  "defaultAgentId": 1,
  "desc": "string",
  "deviceSchemaMode": 1,
  "deviceType": 1,
  "netType": 1,
  "onlineHandle": 1,
  "productID": "string",
  "productImg": "string",
  "productName": "示例名称",
  "protocol": {
    "code": "string",
    "configFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "containerPort": 1,
    "deployStatus": "string",
    "desc": "string",
    "deviceFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "endpoints": [
      "string"
    ],
    "envVars": {},
    "etcdKey": "string",
    "id": "string",
    "image": "string",
    "isBuiltin": 1,
    "isEnableSyncDevice": 1,
    "isEnableSyncProduct": 1,
    "name": "string",
    "productFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "transProtocol": "string",
    "type": "string",
    "version": "string"
  },
  "protocolCode": "string",
  "protocolConf": {},
  "sceneMode": "string",
  "secret": "string",
  "status": 1,
  "subProtocolCode": "string",
  "subProtocolConf": {},
  "tags": {},
  "tenantCode": "string",
  "trialTime": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "productID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/info/create \
  --body '{"authMode": 1, "autoRegister": 1, "bindLevel": 1, "category": {"avatar": "string", "children": [{"avatar": "string", "children": [{"avatar": "...", "children": "...", "code": "...", "desc": "...", "deviceCount": "...", "id": "...", "idPath": "...", "isLeaf": "...", "name": "...", "order": "...", "parentID": "..."}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}, "categoryID": "string", "compatMode": 1, "config": {"devInit": {"areaID": "string", "deptID": "string", "projectID": "string", "tenantCode": "string"}, "productID": "string"}, "createdTime": "2026-01-01T00:00:00Z", "customUi": {}, "defaultAgentId": 1, "desc": "string", "deviceSchemaMode": 1, "deviceType": 1, "netType": 1, "onlineHandle": 1, "productID": "string", "productImg": "string", "productName": "示例名称", "protocol": {"code": "string", "configFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "containerPort": 1, "deployStatus": "string", "desc": "string", "deviceFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "endpoints": ["string"], "envVars": {}, "etcdKey": "string", "id": "string", "image": "string", "isBuiltin": 1, "isEnableSyncDevice": 1, "isEnableSyncProduct": 1, "name": "string", "productFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "transProtocol": "string", "type": "string", "version": "string"}, "protocolCode": "string", "protocolConf": {}, "sceneMode": "string", "secret": "string", "status": 1, "subProtocolCode": "string", "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "trialTime": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/product/info/delete`

**说明**: 删除产品

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id 只读 |

**请求示例**:
```json
{
  "productID": "string"
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
ur api /api/v1/things/product/info/delete \
  --body '{"productID": "string"}'
```

### POST `/api/v1/things/product/info/get-list`

**说明**: 获取产品信息列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 过滤区域下有设备的产品 |
| `areaIDPath` | string | 否 | 过滤区域id路径下有设备的产品 |
| `categoryIDs` | array[integer] | 否 | 产品品类 |
| `deviceType` | integer | 否 | 过滤设备类型:0:全部,1:设备,2:网关,3:子设备 (格式: int64) |
| `deviceTypes` | array[integer] | 否 | 设备类型:1:设备,2:网关,3:子设备设备类型:1:设备,2:网关,3:子设备 |
| `netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productIDs` | array[string] | 否 | 过滤产品id列表 |
| `productName` | string | 否 | 过滤产品名称 |
| `projectID` | string | 否 | 项目id,只获取项目下有设备的 |
| `protocolCode` | string | 否 | 协议code |
| `protocolType` | string | 否 | 协议类型 |
| `sceneMode` | string | 否 | 场景模式 读写类型: r(只读) rw(可读可写) none(不参与场景) |
| `sceneModes` | array[string] | 否 | 场景模式 读写类型: r(只读) rw(可读可写) none(不参与场景) |
| `status` | integer | 否 | 产品状态 1:启用 2:禁用 3:开发中 (格式: int64) |
| `statuses` | array[integer] | 否 | 产品状态 1:启用 2:禁用 3:开发中 |
| `tags` | object | 否 |  key tag过滤查询,非模糊查询 为tag的名,value为tag对应的值 |
| `tenantCode` | string | 否 | 租户号过滤,common 为公共的产品,公共的大家都可以看 |
| `withCategory` | boolean | 否 | 同时返回品类详情 (格式: boolean) |
| `withProtocol` | boolean | 否 | 同时返回协议详情 (格式: boolean) |

**请求示例**:
```json
{
  "areaID": "string",
  "areaIDPath": "string",
  "categoryIDs": [
    1
  ],
  "deviceType": 1,
  "deviceTypes": [
    1
  ],
  "netType": 1,
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
  "productIDs": [
    "string"
  ],
  "productName": "示例名称",
  "projectID": "string",
  "protocolCode": "string",
  "protocolType": "string",
  "sceneMode": "string",
  "sceneModes": [
    "string"
  ],
  "status": 1,
  "statuses": [
    1
  ],
  "tags": {},
  "tenantCode": "string",
  "withCategory": true,
  "withProtocol": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "authMode": 1,
        "autoRegister": 1,
        "bindLevel": 1,
        "category": {
          "avatar": "string",
          "children": [
            {
              "avatar": "...",
              "children": "...",
              "code": "...",
              "desc": "...",
              "deviceCount": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "name": "...",
              "order": "...",
              "parentID": "..."
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
        "categoryID": "string",
        "compatMode": 1,
        "config": {
          "devInit": {
            "areaID": "string",
            "deptID": "string",
            "projectID": "string",
            "tenantCode": "string"
          },
          "productID": "string"
        },
        "createdTime": "2026-01-01T00:00:00Z",
        "customUi": {},
        "defaultAgentId": 1,
        "desc": "string",
        "deviceSchemaMode": 1,
        "deviceType": 1,
        "netType": 1,
        "onlineHandle": 1,
        "productID": "string",
        "productImg": "string",
        "productName": "示例名称",
        "protocol": {
          "code": "string",
          "configFields": [
            {
              "defaultValue": "...",
              "group": "...",
              "id": "...",
              "isRequired": "...",
              "key": "...",
              "label": "...",
              "mapping": "...",
              "max": "...",
              "min": "...",
              "placeholder": "...",
              "sort": "...",
              "type": "..."
            }
          ],
          "containerPort": 1,
          "deployStatus": "string",
          "desc": "string",
          "deviceFields": [
            {
              "defaultValue": "...",
              "group": "...",
              "id": "...",
              "isRequired": "...",
              "key": "...",
              "label": "...",
              "mapping": "...",
              "max": "...",
              "min": "...",
              "placeholder": "...",
              "sort": "...",
              "type": "..."
            }
          ],
          "endpoints": [
            "string"
          ],
          "envVars": {},
          "etcdKey": "string",
          "id": "string",
          "image": "string",
          "isBuiltin": 1,
          "isEnableSyncDevice": 1,
          "isEnableSyncProduct": 1,
          "name": "string",
          "productFields": [
            {
              "defaultValue": "...",
              "group": "...",
              "id": "...",
              "isRequired": "...",
              "key": "...",
              "label": "...",
              "mapping": "...",
              "max": "...",
              "min": "...",
              "placeholder": "...",
              "sort": "...",
              "type": "..."
            }
          ],
          "transProtocol": "string",
          "type": "string",
          "version": "string"
        },
        "protocolCode": "string",
        "protocolConf": {},
        "sceneMode": "string",
        "secret": "string",
        "status": 1,
        "subProtocolCode": "string",
        "subProtocolConf": {},
        "tags": {},
        "tenantCode": "string",
        "trialTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/product/info/get-list \
  --body '{"areaID": "string", "areaIDPath": "string", "categoryIDs": [1], "deviceType": 1, "deviceTypes": [1], "netType": 1, "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productIDs": ["string"], "productName": "示例名称", "projectID": "string", "protocolCode": "string", "protocolType": "string", "sceneMode": "string", "sceneModes": ["string"], "status": 1, "statuses": [1], "tags": {}, "tenantCode": "string", "withCategory": true, "withProtocol": true}'
```

### POST `/api/v1/things/product/info/get-one`

**说明**: 获取产品详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id |
| `withCategory` | boolean | 否 | 同时返回品类详情 (格式: boolean) |
| `withProtocol` | boolean | 否 | 同时返回协议详情 (格式: boolean) |

**请求示例**:
```json
{
  "productID": "string",
  "withCategory": true,
  "withProtocol": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "authMode": 1,
    "autoRegister": 1,
    "bindLevel": 1,
    "category": {
      "avatar": "string",
      "children": [
        {
          "avatar": "string",
          "children": [
            {
              "avatar": "...",
              "children": "...",
              "code": "...",
              "desc": "...",
              "deviceCount": "...",
              "id": "...",
              "idPath": "...",
              "isLeaf": "...",
              "name": "...",
              "order": "...",
              "parentID": "..."
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
    "categoryID": "string",
    "compatMode": 1,
    "config": {
      "devInit": {
        "areaID": "string",
        "deptID": "string",
        "projectID": "string",
        "tenantCode": "string"
      },
      "productID": "string"
    },
    "createdTime": "2026-01-01T00:00:00Z",
    "customUi": {},
    "defaultAgentId": 1,
    "desc": "string",
    "deviceSchemaMode": 1,
    "deviceType": 1,
    "netType": 1,
    "onlineHandle": 1,
    "productID": "string",
    "productImg": "string",
    "productName": "示例名称",
    "protocol": {
      "code": "string",
      "configFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "containerPort": 1,
      "deployStatus": "string",
      "desc": "string",
      "deviceFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "endpoints": [
        "string"
      ],
      "envVars": {},
      "etcdKey": "string",
      "id": "string",
      "image": "string",
      "isBuiltin": 1,
      "isEnableSyncDevice": 1,
      "isEnableSyncProduct": 1,
      "name": "string",
      "productFields": [
        {
          "defaultValue": "string",
          "group": "string",
          "id": "string",
          "isRequired": true,
          "key": "string",
          "label": "string",
          "mapping": {},
          "max": "string",
          "min": "string",
          "placeholder": "string",
          "sort": 1,
          "type": "string"
        }
      ],
      "transProtocol": "string",
      "type": "string",
      "version": "string"
    },
    "protocolCode": "string",
    "protocolConf": {},
    "sceneMode": "string",
    "secret": "string",
    "status": 1,
    "subProtocolCode": "string",
    "subProtocolConf": {},
    "tags": {},
    "tenantCode": "string",
    "trialTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/info/get-one \
  --body '{"productID": "string", "withCategory": true, "withProtocol": true}'
```

### POST `/api/v1/things/product/info/init`

**说明**: 初始化产品

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productIDs` | array[string] | 否 | 产品id 只读 |

**请求示例**:
```json
{
  "productIDs": [
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
ur api /api/v1/things/product/info/init \
  --body '{"productIDs": ["string"]}'
```

### POST `/api/v1/things/product/info/update`

**说明**: 更新产品

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `authMode` | integer | 否 | 认证方式:1:账密认证,2:秘钥认证 (格式: int64) |
| `autoRegister` | integer | 否 | 动态注册:1:关闭,2:打开,3:打开并自动创建设备,4:在前面的基础上绑定没有也自动创建 (格式: int64) |
| `bindLevel` | integer | 否 | 绑定级别: 1:强绑定(默认,只有用户解绑之后才能绑定) 2:中绑定(可以通过token强制解绑设备) 3:弱绑定(app可以内部解绑被绑定的设备) (格式: int64) |
| `category` | object | 否 |  |
| `category.avatar` | string | 否 |  品类图片 |
| `category.children` | array[ProductCategory] | 否 |  |
| `category.code` | string | 否 | 品类编码,全局唯一,创建后不可修改,创建时必填 |
| `category.desc` | string | 否 | 描述 |
| `category.deviceCount` | integer | 否 | 设备数量 (格式: int64) |
| `category.id` | string | 否 |  |
| `category.idPath` | array[integer] | 否 | 只读 |
| `category.isLeaf` | integer | 否 | 格式: int64 |
| `category.name` | string | 否 | 品类名称 |
| `category.order` | integer | 否 | 排序,必须大于等于1,为0代表不修改 (格式: int64) |
| `category.parentID` | string | 否 |  |
| `categoryID` | string | 否 | 产品品类 |
| `compatMode` | integer | 否 | 协议兼容模式 1:旧productID(大写D) 2:新productId(小写d,默认) (格式: int64) |
| `config` | object | 否 |  |
| `config.devInit` | object | 否 |  |
| `config.devInit.areaID` | string | 否 | 设备创建默认分配给的区域ID |
| `config.devInit.deptID` | string | 否 | 设备创建默认分配给的部门 |
| `config.devInit.projectID` | string | 否 | 设备创建默认分配给的项目 |
| `config.devInit.tenantCode` | string | 否 | 设备创建默认分配给的租户 |
| `config.productID` | string | 是 |  |
| `createdTime` | string | 否 | 创建时间 只读 |
| `customUi` | object | 否 | 自定义ui,key是端的类型(web-client  mini-client) value是以下类型的对象{version:123(版本号,只读),isUpdateUi:bool(是否更新ui),path:string(前端路径,如果需要修改,需要将isUpdateUi置为true并在这个参数中传入压缩包的filePath)} |
| `defaultAgentId` | integer | 否 | 默认绑定的 Agent ID，0 表示未绑定 (格式: int64) |
| `desc` | string | 否 | 描述 |
| `deviceSchemaMode` | integer | 否 |  设备物模型模式:1:手动创建,2:设备自动创建 3: 设备自动创建及上报无定义自动创建 4: 设备自动创建及上报无定义自动创建(数字类型只使用浮点) (格式: int64) |
| `deviceType` | integer | 否 | 设备类型:1:直连设备,2:网关,3:子设备,4:监控设备 (格式: int64) |
| `netType` | integer | 否 | 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN,7:wifi+ble,8:有线网,9:4G+BLE (格式: int64) |
| `onlineHandle` | integer | 否 | 在线处理:1: 自动 2: 永远在线 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `productImg` | string | 否 | 产品图片 |
| `productName` | string | 否 | 产品名称 |
| `protocol` | object | 否 |  |
| `protocol.code` | string | 否 |   iThings,iThings-thingsboard,wumei,aliyun,huaweiyun,tuya |
| `protocol.configFields` | array[ProtocolConfigField] | 否 | 配置字段列表,没有可以不传 |
| `protocol.containerPort` | integer | 否 |  容器端口 (格式: int64) |
| `protocol.deployStatus` | string | 否 |  部署状态: notInstalled/downloaded/running/stopped |
| `protocol.desc` | string | 否 |  |
| `protocol.deviceFields` | array[ProtocolConfigField] | 否 | 设备级的配置字段列表,没有可以不传 |
| `protocol.endpoints` | array[string] | 否 |  |
| `protocol.envVars` | object | 否 |  环境变量 |
| `protocol.etcdKey` | string | 否 |  |
| `protocol.id` | string | 否 |  |
| `protocol.image` | string | 否 |  镜像地址 |
| `protocol.isBuiltin` | integer | 否 |  是否内置 1:是 2:否 (格式: int64) |
| `protocol.isEnableSyncDevice` | integer | 是 | 是否支持主动同步设备信息 (格式: int64) |
| `protocol.isEnableSyncProduct` | integer | 是 | 是否支持主动同步产品信息 (格式: int64) |
| `protocol.name` | string | 否 |  |
| `protocol.productFields` | array[ProtocolConfigField] | 否 | 产品级的配置字段列表,没有可以不传 |
| `protocol.transProtocol` | string | 否 |  传输协议: mqtt,tcp,udp |
| `protocol.type` | string | 否 | 协议类型: 普通设备(默认):normal  音视频:media |
| `protocol.version` | string | 否 |  镜像版本 |
| `protocolCode` | string | 否 | 协议code,默认urMqtt,设备下发只会发送给主协议  urMqtt,urHttp,wumei,aliyun,huaweiyun,tuya |
| `protocolConf` | object | 否 | 协议配置 |
| `sceneMode` | string | 否 | 场景模式 读写类型: r(只读) rw(可读可写) none(不参与场景) |
| `secret` | string | 否 | 动态注册产品秘钥 只读 |
| `status` | integer | 否 | 产品状态 1:启用 2:禁用 3:开发中 (格式: int64) |
| `subProtocolCode` | string | 否 | 子协议,主协议和子协议传输类型必须不相同, 设备控制下发只会发送给主协议, 当设备是音视频设备但是控制协议需要单独走的时候就可以把主协议定义为普通协议,子协议定义为音视频协议,这样就能实现音视频走音视频协议,控制走子协议 |
| `subProtocolConf` | object | 否 | 子协议协议配置 |
| `tags` | object | 否 |  产品tag |
| `tenantCode` | string | 否 | 租户号,common 为公共的产品,公共的大家都可以看,但是只有default可以修改 |
| `trialTime` | string | 否 | 试用时间(单位为天,为0不限制) |

**请求示例**:
```json
{
  "authMode": 1,
  "autoRegister": 1,
  "bindLevel": 1,
  "category": {
    "avatar": "string",
    "children": [
      {
        "avatar": "string",
        "children": [
          {
            "avatar": "...",
            "children": "...",
            "code": "...",
            "desc": "...",
            "deviceCount": "...",
            "id": "...",
            "idPath": "...",
            "isLeaf": "...",
            "name": "...",
            "order": "...",
            "parentID": "..."
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
  "categoryID": "string",
  "compatMode": 1,
  "config": {
    "devInit": {
      "areaID": "string",
      "deptID": "string",
      "projectID": "string",
      "tenantCode": "string"
    },
    "productID": "string"
  },
  "createdTime": "2026-01-01T00:00:00Z",
  "customUi": {},
  "defaultAgentId": 1,
  "desc": "string",
  "deviceSchemaMode": 1,
  "deviceType": 1,
  "netType": 1,
  "onlineHandle": 1,
  "productID": "string",
  "productImg": "string",
  "productName": "示例名称",
  "protocol": {
    "code": "string",
    "configFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "containerPort": 1,
    "deployStatus": "string",
    "desc": "string",
    "deviceFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "endpoints": [
      "string"
    ],
    "envVars": {},
    "etcdKey": "string",
    "id": "string",
    "image": "string",
    "isBuiltin": 1,
    "isEnableSyncDevice": 1,
    "isEnableSyncProduct": 1,
    "name": "string",
    "productFields": [
      {
        "defaultValue": "string",
        "group": "string",
        "id": "string",
        "isRequired": true,
        "key": "string",
        "label": "string",
        "mapping": {},
        "max": "string",
        "min": "string",
        "placeholder": "string",
        "sort": 1,
        "type": "string"
      }
    ],
    "transProtocol": "string",
    "type": "string",
    "version": "string"
  },
  "protocolCode": "string",
  "protocolConf": {},
  "sceneMode": "string",
  "secret": "string",
  "status": 1,
  "subProtocolCode": "string",
  "subProtocolConf": {},
  "tags": {},
  "tenantCode": "string",
  "trialTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/product/info/update \
  --body '{"authMode": 1, "autoRegister": 1, "bindLevel": 1, "category": {"avatar": "string", "children": [{"avatar": "string", "children": [{"avatar": "...", "children": "...", "code": "...", "desc": "...", "deviceCount": "...", "id": "...", "idPath": "...", "isLeaf": "...", "name": "...", "order": "...", "parentID": "..."}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}], "code": "string", "desc": "string", "deviceCount": 1, "id": "string", "idPath": [1], "isLeaf": 1, "name": "示例名称", "order": 1, "parentID": "string"}, "categoryID": "string", "compatMode": 1, "config": {"devInit": {"areaID": "string", "deptID": "string", "projectID": "string", "tenantCode": "string"}, "productID": "string"}, "createdTime": "2026-01-01T00:00:00Z", "customUi": {}, "defaultAgentId": 1, "desc": "string", "deviceSchemaMode": 1, "deviceType": 1, "netType": 1, "onlineHandle": 1, "productID": "string", "productImg": "string", "productName": "示例名称", "protocol": {"code": "string", "configFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "containerPort": 1, "deployStatus": "string", "desc": "string", "deviceFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "endpoints": ["string"], "envVars": {}, "etcdKey": "string", "id": "string", "image": "string", "isBuiltin": 1, "isEnableSyncDevice": 1, "isEnableSyncProduct": 1, "name": "string", "productFields": [{"defaultValue": "string", "group": "string", "id": "string", "isRequired": true, "key": "string", "label": "string", "mapping": {}, "max": "string", "min": "string", "placeholder": "string", "sort": 1, "type": "string"}], "transProtocol": "string", "type": "string", "version": "string"}, "protocolCode": "string", "protocolConf": {}, "sceneMode": "string", "secret": "string", "status": 1, "subProtocolCode": "string", "subProtocolConf": {}, "tags": {}, "tenantCode": "string", "trialTime": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/things/product/remote-config/create`

**说明**: 创建配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 | 配置内容 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "content": "string",
  "productID": "string"
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
ur api /api/v1/things/product/remote-config/create \
  --body '{"content": "string", "productID": "string"}'
```

### POST `/api/v1/things/product/remote-config/get-list`

**说明**: 获取配置列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 是 | 产品id |

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
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "content": "string",
        "createTime": "2026-01-01T00:00:00Z",
        "id": "string"
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
ur api /api/v1/things/product/remote-config/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string"}'
```

### POST `/api/v1/things/product/remote-config/lastest-read`

**说明**: 获取最新配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "content": "string",
    "createTime": "2026-01-01T00:00:00Z",
    "id": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/remote-config/lastest-read \
  --body '{"productID": "string"}'
```

### POST `/api/v1/things/product/remote-config/push-all`

**说明**: 推送配置

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "productID": "string"
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
ur api /api/v1/things/product/remote-config/push-all \
  --body '{"productID": "string"}'
```

### POST `/api/v1/things/product/schema/batch-create`

**说明**: 批量创建产品物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `identifiers` | array[string] | 否 | 导入的标识符列表 |
| `list` | array[ProductSchemaInfo] | 否 | 如果没有填标识符列表取这里 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "identifiers": [
    "string"
  ],
  "list": [
    {
      "affordance": "string",
      "desc": "string",
      "extendConfig": "string",
      "funcGroup": 1,
      "identifier": "string",
      "isCanSceneLinkage": 1,
      "name": "示例名称",
      "order": 1,
      "productID": "string",
      "required": 1,
      "tag": 1,
      "type": 1
    }
  ],
  "productID": "string"
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
ur api /api/v1/things/product/schema/batch-create \
  --body '{"identifiers": ["string"], "list": [{"affordance": "string", "desc": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}], "productID": "string"}'
```

### POST `/api/v1/things/product/schema/create`

**说明**: 创建产品物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `affordance` | string | 否 |  |
| `desc` | string | 否 | 描述 |
| `extendConfig` | string | 否 | 拓展参数 |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifier` | string | 是 | 标识符 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 | 功能名称 |
| `order` | integer | 否 |  排序 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `required` | integer | 否 | 是否必须 1:是 2:否 (格式: int64) |
| `tag` | integer | 是 | 物模型标签 1:自定义 2:可选 3:必选  必选不可删除 (格式: int64) |
| `type` | integer | 是 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

**请求示例**:
```json
{
  "affordance": "string",
  "desc": "string",
  "extendConfig": "string",
  "funcGroup": 1,
  "identifier": "string",
  "isCanSceneLinkage": 1,
  "name": "示例名称",
  "order": 1,
  "productID": "string",
  "required": 1,
  "tag": 1,
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
ur api /api/v1/things/product/schema/create \
  --body '{"affordance": "string", "desc": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}'
```

### POST `/api/v1/things/product/schema/delete`

**说明**: 删除产品物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `identifier` | string | 是 | 标识符 |
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "identifier": "string",
  "productID": "string"
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
ur api /api/v1/things/product/schema/delete \
  --body '{"identifier": "string", "productID": "string"}'
```

### POST `/api/v1/things/product/schema/get-list`

**说明**: 获取产品物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `controlMode` | integer | 否 | 控制模式: 1: 可以群控,可以单控  2:只能单控 (格式: int64) |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifiers` | array[string] | 否 | 过滤标识符列表 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 是 | 产品id |
| `propertyMode` | string | 否 | 属性类型可用 读写类型: 1:r(只读) 2:rw(可读可写) |
| `tag` | integer | 否 | 过滤条件: 物模型标签 1:自定义 2:可选 3:必选 (格式: int64) |
| `type` | integer | 否 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |
| `types` | array[integer] | 否 |  |

**请求示例**:
```json
{
  "controlMode": 1,
  "funcGroup": 1,
  "identifiers": [
    "string"
  ],
  "isCanSceneLinkage": 1,
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
  "productID": "string",
  "propertyMode": "string",
  "tag": 1,
  "type": 1,
  "types": [
    1
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
        "affordance": "string",
        "desc": "string",
        "extendConfig": "string",
        "funcGroup": 1,
        "identifier": "string",
        "isCanSceneLinkage": 1,
        "name": "示例名称",
        "order": 1,
        "productID": "string",
        "required": 1,
        "tag": 1,
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
ur api /api/v1/things/product/schema/get-list \
  --body '{"controlMode": 1, "funcGroup": 1, "identifiers": ["string"], "isCanSceneLinkage": 1, "name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string", "propertyMode": "string", "tag": 1, "type": 1, "types": [1]}'
```

### POST `/api/v1/things/product/schema/tsl-import`

**说明**: 导入产品物模型tsl

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id 只读 |
| `tsl` | string | 是 | 物模型tsl |

**请求示例**:
```json
{
  "productID": "string",
  "tsl": "string"
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
ur api /api/v1/things/product/schema/tsl-import \
  --body '{"productID": "string", "tsl": "string"}'
```

### POST `/api/v1/things/product/schema/tsl-read`

**说明**: 获取产品物模型tsl

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `productID` | string | 是 | 产品id |

**请求示例**:
```json
{
  "productID": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "tsl": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/product/schema/tsl-read \
  --body '{"productID": "string"}'
```

### POST `/api/v1/things/product/schema/update`

**说明**: 更新产品物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `affordance` | string | 否 |  |
| `desc` | string | 否 | 描述 |
| `extendConfig` | string | 否 | 拓展参数 |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifier` | string | 是 | 标识符 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 | 功能名称 |
| `order` | integer | 否 |  排序 (格式: int64) |
| `productID` | string | 否 | 产品id 只读 |
| `required` | integer | 否 | 是否必须 1:是 2:否 (格式: int64) |
| `tag` | integer | 是 | 物模型标签 1:自定义 2:可选 3:必选  必选不可删除 (格式: int64) |
| `type` | integer | 是 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

**请求示例**:
```json
{
  "affordance": "string",
  "desc": "string",
  "extendConfig": "string",
  "funcGroup": 1,
  "identifier": "string",
  "isCanSceneLinkage": 1,
  "name": "示例名称",
  "order": 1,
  "productID": "string",
  "required": 1,
  "tag": 1,
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
ur api /api/v1/things/product/schema/update \
  --body '{"affordance": "string", "desc": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "productID": "string", "required": 1, "tag": 1, "type": 1}'
```
