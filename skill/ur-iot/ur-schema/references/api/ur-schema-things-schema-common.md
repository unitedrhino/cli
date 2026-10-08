# ur-schema things/schema/common

批量导出通用物模型 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/schema/common/batch-export` | 批量导出通用物模型 | admin |
| POST | `/api/v1/things/schema/common/batch-import` | 批量导入通用物模型 | admin |
| POST | `/api/v1/things/schema/common/create` | 新增通用物模型 | admin |
| POST | `/api/v1/things/schema/common/delete` | 删除通用物模型 | admin |
| POST | `/api/v1/things/schema/common/get-list` | 获取通用物模型列表 | admin |
| POST | `/api/v1/things/schema/common/init` | 初始化通用物模型 | admin |
| POST | `/api/v1/things/schema/common/mock-gen` | 从参数生成物模型定义 | admin |
| POST | `/api/v1/things/schema/common/update` | 更新通用物模型 | admin |

## 详细说明

### POST `/api/v1/things/schema/common/batch-export`

**说明**: 批量导出通用物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `identifiers` | array[string] | 否 |  |

**请求示例**:
```json
{
  "identifiers": [
    "string"
  ]
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "schemasFile": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/schema/common/batch-export \
  --body '{"identifiers": ["string"]}'
```

### POST `/api/v1/things/schema/common/batch-import`

**说明**: 批量导入通用物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `schemasFile` | string | 是 | 物模型数据文件的 OSS 地址（ZIP 压缩包） |

**请求示例**:
```json
{
  "schemasFile": "string"
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
ur api /api/v1/things/schema/common/batch-import \
  --body '{"schemasFile": "string"}'
```

### POST `/api/v1/things/schema/common/create`

**说明**: 新增通用物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `affordance` | string | 否 |  |
| `desc` | string | 否 | 描述 |
| `extendConfig` | string | 否 | 拓展参数 |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifier` | string | 否 | 标识符 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 | 功能名称 |
| `order` | integer | 否 |  排序 (格式: int64) |
| `required` | integer | 否 | 是否必须 1:是 2:否 (格式: int64) |
| `type` | integer | 否 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

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
  "required": 1,
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
ur api /api/v1/things/schema/common/create \
  --body '{"affordance": "string", "desc": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "required": 1, "type": 1}'
```

### POST `/api/v1/things/schema/common/delete`

**说明**: 删除通用物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `identifier` | string | 否 | 标识符 |

**请求示例**:
```json
{
  "identifier": "string"
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
ur api /api/v1/things/schema/common/delete \
  --body '{"identifier": "string"}'
```

### POST `/api/v1/things/schema/common/get-list`

**说明**: 获取通用物模型列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 |  |
| `areaIDPath` | string | 否 | 过滤区域id路径下有设备的产品 |
| `controlMode` | integer | 否 | 控制模式: 1: 可以群控,可以单控  2:只能单控 (格式: int64) |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `groupID` | string | 否 |  |
| `groupIDs` | array[string] | 否 |  |
| `identifiers` | array[string] | 否 | 过滤标识符列表 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productCategoryCode` | string | 否 | 品类编码 |
| `productCategoryID` | string | 否 |  |
| `productCategoryWithFather` | boolean | 否 | 格式: boolean |
| `productIDs` | array[string] | 否 | 提取产品共有的物模型 |
| `productSceneMode` | string | 否 | 产品场景模式 读写类型: r(只读) rw(可读可写) none(不参与场景) |
| `projectID` | string | 否 | 过滤项目下共有的物模型 |
| `propertyMode` | string | 否 | 属性类型可用 读写类型: r(只读) rw(可读可写) |
| `type` | integer | 否 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |
| `types` | array[integer] | 否 |  |

**请求示例**:
```json
{
  "areaID": "string",
  "areaIDPath": "string",
  "controlMode": 1,
  "funcGroup": 1,
  "groupID": "string",
  "groupIDs": [
    "string"
  ],
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
  "productCategoryCode": "string",
  "productCategoryID": "string",
  "productCategoryWithFather": true,
  "productIDs": [
    "string"
  ],
  "productSceneMode": "string",
  "projectID": "string",
  "propertyMode": "string",
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
        "required": 1,
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
ur api /api/v1/things/schema/common/get-list \
  --body '{"areaID": "string", "areaIDPath": "string", "controlMode": 1, "funcGroup": 1, "groupID": "string", "groupIDs": ["string"], "identifiers": ["string"], "isCanSceneLinkage": 1, "name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productCategoryCode": "string", "productCategoryID": "string", "productCategoryWithFather": true, "productIDs": ["string"], "productSceneMode": "string", "projectID": "string", "propertyMode": "string", "type": 1, "types": [1]}'
```

### POST `/api/v1/things/schema/common/init`

**说明**: 初始化通用物模型

**权限**: admin

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/schema/common/init \
  --body '{}'
```

### POST `/api/v1/things/schema/common/mock-gen`

**说明**: 从参数生成物模型定义

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `inputParams` | string | 否 |  行为输入参数,json字符串 (for action) |
| `outputParams` | string | 否 |  行为输出参数,json字符串 (for action) |
| `params` | string | 否 |  属性参数或事件参数,json字符串 (for property/event) |
| `type` | integer | 是 |  物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

**请求示例**:
```json
{
  "inputParams": "string",
  "outputParams": "string",
  "params": "string",
  "type": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "schemas": [
      {
        "affordance": "string",
        "identifier": "string"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/schema/common/mock-gen \
  --body '{"inputParams": "string", "outputParams": "string", "params": "string", "type": 1}'
```

### POST `/api/v1/things/schema/common/update`

**说明**: 更新通用物模型

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `affordance` | string | 否 |  |
| `desc` | string | 否 | 描述 |
| `extendConfig` | string | 否 | 拓展参数 |
| `funcGroup` | integer | 否 |  功能分类: 1:普通功能 2:系统功能 (格式: int64) |
| `identifier` | string | 否 | 标识符 |
| `isCanSceneLinkage` | integer | 否 | 是否可以场景联动 (格式: int64) |
| `name` | string | 否 | 功能名称 |
| `order` | integer | 否 |  排序 (格式: int64) |
| `required` | integer | 否 | 是否必须 1:是 2:否 (格式: int64) |
| `type` | integer | 否 | 物模型类型 1:property属性 2:event事件 3:action行为 (格式: int64) |

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
  "required": 1,
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
ur api /api/v1/things/schema/common/update \
  --body '{"affordance": "string", "desc": "string", "extendConfig": "string", "funcGroup": 1, "identifier": "string", "isCanSceneLinkage": 1, "name": "示例名称", "order": 1, "required": 1, "type": 1}'
```
