# ur-project things/project/crud

删除项目crud 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/project/crud/delete` | 删除项目crud | admin |
| POST | `/api/v1/things/project/crud/get-list` | 获取项目crud列表 | admin |
| POST | `/api/v1/things/project/crud/get-one` | 获取项目crud详情 | admin |

## 详细说明

### POST `/api/v1/things/project/crud/delete`

**说明**: 删除项目crud

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
ur api /api/v1/things/project/crud/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/project/crud/get-list`

**说明**: 获取项目crud列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `params` | object | 否 | params过滤查询,key 是params的key, value是对象,说明: {"cmpType":"比较类型(如=、>、like)","value":"比较值(需与类型匹配)","castTo":"数据类型(int/float/string,选填，填则转换)"}，其中 like 仅支持 string 类型，value 需以字符串格式传入 示例: {"aaa":{"cmpType":">","castTo":"int","value":"100"} |
| `purpose` | string | 是 | 用途必填 |

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
  "params": {},
  "purpose": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "createdTime": "string",
        "id": "string",
        "params": "string",
        "purpose": "string",
        "sort": 1
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
ur api /api/v1/things/project/crud/get-list \
  --body '{"page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "params": {}, "purpose": "string"}'
```

### POST `/api/v1/things/project/crud/get-one`

**说明**: 获取项目crud详情

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
  "data": {
    "createdTime": "string",
    "id": "string",
    "params": "string",
    "purpose": "string",
    "sort": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/project/crud/get-one \
  --body '{"id": "string"}'
```
