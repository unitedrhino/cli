# ur-ai things/alarm/conditionTemplate

创建触发条件模板 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/condition-template/create` | 创建触发条件模板 | admin |
| POST | `/api/v1/things/alarm/condition-template/delete` | 删除触发条件模板 | admin |
| POST | `/api/v1/things/alarm/condition-template/get-list` | 获取触发条件模板列表 | admin |
| POST | `/api/v1/things/alarm/condition-template/get-one` | 获取触发条件模板详情 | admin |
| POST | `/api/v1/things/alarm/condition-template/update` | 更新触发条件模板 | admin |

## 详细说明

### POST `/api/v1/things/alarm/condition-template/create`

**说明**: 创建触发条件模板

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  描述 |
| `if` | object | 否 |  |
| `if.triggers` | array[AlarmTrigger] | 否 |  触发器列表 |
| `name` | string | 是 |  名称 |
| `productCategoryID` | string | 否 |  产品品类ID |
| `productID` | string | 否 |  产品ID |

**请求示例**:
```json
{
  "desc": "string",
  "if": {
    "triggers": [
      {
        "criteria": [
          {
            "duration": "...",
            "frequency": "...",
            "id": "...",
            "kOfN": "...",
            "order": "...",
            "plain": "...",
            "rollingAggregate": "...",
            "type": "..."
          }
        ],
        "id": "string",
        "level": "string",
        "templateID": "string",
        "templateName": "示例名称",
        "type": "string"
      }
    ]
  },
  "name": "示例名称",
  "productCategoryID": "string",
  "productID": "string"
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
ur api /api/v1/things/alarm/condition-template/create \
  --body '{"desc": "string", "if": {"triggers": [{"criteria": [{"duration": "...", "frequency": "...", "id": "...", "kOfN": "...", "order": "...", "plain": "...", "rollingAggregate": "...", "type": "..."}], "id": "string", "level": "string", "templateID": "string", "templateName": "示例名称", "type": "string"}]}, "name": "示例名称", "productCategoryID": "string", "productID": "string"}'
```

### POST `/api/v1/things/alarm/condition-template/delete`

**说明**: 删除触发条件模板

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  ID |

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
  "data": {},
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/condition-template/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/condition-template/get-list`

**说明**: 获取触发条件模板列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `keyword` | string | 否 |  名称模糊搜索 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productCategoryID` | string | 否 |  产品品类ID |
| `productID` | string | 否 |  产品ID |
| `status` | integer | 否 |  状态 (格式: int64) |

**请求示例**:
```json
{
  "keyword": "示例名称",
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
  "productCategoryID": "string",
  "productID": "string",
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
        "desc": "string",
        "id": "string",
        "if": {
          "triggers": [
            {
              "criteria": "...",
              "id": "...",
              "level": "...",
              "templateID": "...",
              "templateName": "...",
              "type": "..."
            }
          ]
        },
        "name": "示例名称",
        "productCategoryID": "string",
        "productID": "string",
        "refCount": 1,
        "status": 1,
        "updatedTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/things/alarm/condition-template/get-list \
  --body '{"keyword": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productCategoryID": "string", "productID": "string", "status": 1}'
```

### POST `/api/v1/things/alarm/condition-template/get-one`

**说明**: 获取触发条件模板详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 否 |  ID |

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
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "if": {
      "triggers": [
        {
          "criteria": [
            {
              "duration": "...",
              "frequency": "...",
              "id": "...",
              "kOfN": "...",
              "order": "...",
              "plain": "...",
              "rollingAggregate": "...",
              "type": "..."
            }
          ],
          "id": "string",
          "level": "string",
          "templateID": "string",
          "templateName": "示例名称",
          "type": "string"
        }
      ]
    },
    "name": "示例名称",
    "productCategoryID": "string",
    "productID": "string",
    "refCount": 1,
    "status": 1,
    "updatedTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/condition-template/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/condition-template/update`

**说明**: 更新触发条件模板

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  描述 |
| `id` | string | 是 |  ID |
| `if` | object | 否 |  |
| `if.triggers` | array[AlarmTrigger] | 否 |  触发器列表 |
| `name` | string | 否 |  名称 |
| `productCategoryID` | string | 否 |  产品品类ID |
| `productID` | string | 否 |  产品ID |
| `status` | integer | 否 |  状态 (格式: int64) |

**请求示例**:
```json
{
  "desc": "string",
  "id": "string",
  "if": {
    "triggers": [
      {
        "criteria": [
          {
            "duration": "...",
            "frequency": "...",
            "id": "...",
            "kOfN": "...",
            "order": "...",
            "plain": "...",
            "rollingAggregate": "...",
            "type": "..."
          }
        ],
        "id": "string",
        "level": "string",
        "templateID": "string",
        "templateName": "示例名称",
        "type": "string"
      }
    ]
  },
  "name": "示例名称",
  "productCategoryID": "string",
  "productID": "string",
  "status": 1
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
ur api /api/v1/things/alarm/condition-template/update \
  --body '{"desc": "string", "id": "string", "if": {"triggers": [{"criteria": [{"duration": "...", "frequency": "...", "id": "...", "kOfN": "...", "order": "...", "plain": "...", "rollingAggregate": "...", "type": "..."}], "id": "string", "level": "string", "templateID": "string", "templateName": "示例名称", "type": "string"}]}, "name": "示例名称", "productCategoryID": "string", "productID": "string", "status": 1}'
```
