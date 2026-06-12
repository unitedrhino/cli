# ur-ai ai/skill

创建技能 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/ai/skill/create` | 创建技能 | admin |
| POST | `/api/v1/ai/skill/delete` | 删除技能 | admin |
| POST | `/api/v1/ai/skill/download` | 下载技能zip包 | admin |
| POST | `/api/v1/ai/skill/get-list` | 获取技能列表 | admin |
| POST | `/api/v1/ai/skill/get-one` | 获取技能详情 | admin |
| POST | `/api/v1/ai/skill/publish` | 发布新版本 | admin |
| POST | `/api/v1/ai/skill/update` | 更新技能 | admin |
| POST | `/api/v1/ai/skill/upload` | 上传技能zip包 | admin |

## 详细说明

### POST `/api/v1/ai/skill/create`

**说明**: 创建技能

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  技能编码 |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  技能ID (格式: int64) |
| `manifest` | string | 否 |  skill.yaml内容（JSON字符串） |
| `name` | string | 是 |  技能名称 |
| `ossPath` | string | 否 |  OSS存储路径 |
| `parentID` | integer | 否 |  父技能ID（0=顶级） (格式: int64) |
| `publishStatus` | string | 否 |  发布状态：draft/uploaded/validated/review_done/test_passed/approved/published/deprecated |
| `reviewResult` | string | 否 |  AI审阅结果（JSON） |
| `scope` | string | 否 |  范围：tenant=企业级 platform=平台级 |
| `sort` | integer | 否 |  排序 (格式: int64) |
| `status` | integer | 否 |  状态 1-启用 2-禁用 (格式: int64) |
| `tenantCode` | string | 否 |  企业编码，"common"=平台级 |
| `testResult` | string | 否 |  测试结果（JSON） |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |
| `version` | string | 否 |  版本号 |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": 1,
  "desc": "string",
  "id": 1,
  "manifest": "string",
  "name": "示例名称",
  "ossPath": "string",
  "parentID": 1,
  "publishStatus": "string",
  "reviewResult": "string",
  "scope": "string",
  "sort": 1,
  "status": 1,
  "tenantCode": "string",
  "testResult": "string",
  "updatedTime": 1,
  "version": "string"
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
ur api /api/v1/ai/skill/create \
  --body '{"code": "string", "createdTime": 1, "desc": "string", "id": 1, "manifest": "string", "name": "示例名称", "ossPath": "string", "parentID": 1, "publishStatus": "string", "reviewResult": "string", "scope": "string", "sort": 1, "status": 1, "tenantCode": "string", "testResult": "string", "updatedTime": 1, "version": "string"}'
```

### POST `/api/v1/ai/skill/delete`

**说明**: 删除技能

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
ur api /api/v1/ai/skill/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/skill/download`

**说明**: 下载技能zip包

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
  "data": {
    "ossPath": "string",
    "zipData": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/download \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/skill/get-list`

**说明**: 获取技能列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  名称模糊搜索 |
| `page` | object | 是 |  |
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | integer | 否 |  父技能ID过滤 (格式: int64) |
| `status` | integer | 否 |  状态过滤 (格式: int64) |
| `tenantCode` | string | 否 |  企业编码过滤 |

**请求示例**:
```json
{
  "name": "示例名称",
  "page": {
    "page": 1,
    "pageSize": 1
  },
  "parentID": 1,
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
        "code": "string",
        "createdTime": 1,
        "desc": "string",
        "id": 1,
        "manifest": "string",
        "name": "示例名称",
        "ossPath": "string",
        "parentID": 1,
        "publishStatus": "string",
        "reviewResult": "string",
        "scope": "string",
        "sort": 1,
        "status": 1,
        "tenantCode": "string",
        "testResult": "string",
        "updatedTime": 1,
        "version": "string"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/get-list \
  --body '{"name": "示例名称", "page": {"page": 1, "pageSize": 1}, "parentID": 1, "status": 1, "tenantCode": "string"}'
```

### POST `/api/v1/ai/skill/get-one`

**说明**: 获取技能详情

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
  "data": {
    "code": "string",
    "createdTime": 1,
    "desc": "string",
    "id": 1,
    "manifest": "string",
    "name": "示例名称",
    "ossPath": "string",
    "parentID": 1,
    "publishStatus": "string",
    "reviewResult": "string",
    "scope": "string",
    "sort": 1,
    "status": 1,
    "tenantCode": "string",
    "testResult": "string",
    "updatedTime": 1,
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/ai/skill/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/ai/skill/publish`

**说明**: 发布新版本

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `skillID` | integer | 是 |  技能ID (格式: int64) |
| `version` | string | 是 |  新版本号 |

**请求示例**:
```json
{
  "skillID": 1,
  "version": "string"
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
ur api /api/v1/ai/skill/publish \
  --body '{"skillID": 1, "version": "string"}'
```

### POST `/api/v1/ai/skill/update`

**说明**: 更新技能

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  技能编码 |
| `createdTime` | integer | 否 |  创建时间 (格式: int64) |
| `desc` | string | 否 |  描述 |
| `id` | integer | 否 |  技能ID (格式: int64) |
| `manifest` | string | 否 |  skill.yaml内容（JSON字符串） |
| `name` | string | 是 |  技能名称 |
| `ossPath` | string | 否 |  OSS存储路径 |
| `parentID` | integer | 否 |  父技能ID（0=顶级） (格式: int64) |
| `publishStatus` | string | 否 |  发布状态：draft/uploaded/validated/review_done/test_passed/approved/published/deprecated |
| `reviewResult` | string | 否 |  AI审阅结果（JSON） |
| `scope` | string | 否 |  范围：tenant=企业级 platform=平台级 |
| `sort` | integer | 否 |  排序 (格式: int64) |
| `status` | integer | 否 |  状态 1-启用 2-禁用 (格式: int64) |
| `tenantCode` | string | 否 |  企业编码，"common"=平台级 |
| `testResult` | string | 否 |  测试结果（JSON） |
| `updatedTime` | integer | 否 |  更新时间 (格式: int64) |
| `version` | string | 否 |  版本号 |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": 1,
  "desc": "string",
  "id": 1,
  "manifest": "string",
  "name": "示例名称",
  "ossPath": "string",
  "parentID": 1,
  "publishStatus": "string",
  "reviewResult": "string",
  "scope": "string",
  "sort": 1,
  "status": 1,
  "tenantCode": "string",
  "testResult": "string",
  "updatedTime": 1,
  "version": "string"
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
ur api /api/v1/ai/skill/update \
  --body '{"code": "string", "createdTime": 1, "desc": "string", "id": 1, "manifest": "string", "name": "示例名称", "ossPath": "string", "parentID": 1, "publishStatus": "string", "reviewResult": "string", "scope": "string", "sort": 1, "status": 1, "tenantCode": "string", "testResult": "string", "updatedTime": 1, "version": "string"}'
```

### POST `/api/v1/ai/skill/upload`

**说明**: 上传技能zip包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 是 |  技能编码 |
| `name` | string | 是 |  技能名称 |
| `parentID` | integer | 否 |  父技能ID (格式: int64) |
| `tenantCode` | string | 否 |  企业编码 |
| `zipData` | string | 是 |  Base64编码的zip包数据 |

**请求示例**:
```json
{
  "code": "string",
  "name": "示例名称",
  "parentID": 1,
  "tenantCode": "string",
  "zipData": "string"
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
ur api /api/v1/ai/skill/upload \
  --body '{"code": "string", "name": "示例名称", "parentID": 1, "tenantCode": "string", "zipData": "string"}'
```
