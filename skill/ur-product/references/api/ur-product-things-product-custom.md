# ur-product things/product/custom

获取产品自定义信息详情 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/product/custom/get-one` | 获取产品自定义信息详情 | admin |
| POST | `/api/v1/things/product/custom/update` | 更新产品自定义信息 | admin |

## 详细说明

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
