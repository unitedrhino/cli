# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/view/asset/delete` | 删除资源 | admin |
| POST | `/api/v1/view/asset/get-list` | 获取资源列表 | admin |
| POST | `/api/v1/view/asset/update` | 更新资源 | admin |
| POST | `/api/v1/view/asset/upload` | 上传资源 | admin |
| POST | `/api/v1/view/component/copy` | 复制组件 | admin |
| POST | `/api/v1/view/component/create` | 新增组件 | admin |
| POST | `/api/v1/view/component/delete` | 删除组件 | admin |
| POST | `/api/v1/view/component/design` | 保存组件设计内容 | admin |
| POST | `/api/v1/view/component/get-list` | 获取组件列表 | admin |
| POST | `/api/v1/view/component/get-one` | 获取组件详情 | admin |
| POST | `/api/v1/view/component/update` | 更新组件 | admin |
| POST | `/api/v1/view/dataset/create` | 新增数据集 | admin |
| POST | `/api/v1/view/dataset/delete` | 删除数据集 | admin |
| POST | `/api/v1/view/dataset/get-list` | 获取数据集列表 | admin |
| POST | `/api/v1/view/dataset/get-one` | 获取数据集详情 | admin |
| POST | `/api/v1/view/dataset/preview` | 预览数据集 | admin |
| POST | `/api/v1/view/dataset/update` | 更新数据集 | admin |
| POST | `/api/v1/view/datasource/create` | 新增数据源 | admin |
| POST | `/api/v1/view/datasource/delete` | 删除数据源 | admin |
| POST | `/api/v1/view/datasource/get-list` | 获取数据源列表 | admin |
| POST | `/api/v1/view/datasource/get-one` | 获取数据源详情 | admin |
| POST | `/api/v1/view/datasource/test` | 测试数据源连接 | admin |
| POST | `/api/v1/view/datasource/update` | 更新数据源 | admin |
| POST | `/api/v1/view/group/create` | 创建分组 | admin |
| POST | `/api/v1/view/group/delete` | 删除分组 | admin |
| POST | `/api/v1/view/group/get-list` | 获取分组列表 | admin |
| POST | `/api/v1/view/group/update` | 更新分组 | admin |
| POST | `/api/v1/view/project/create` | 新增项目 | admin |
| POST | `/api/v1/view/project/delete` | 删除项目 | admin |
| POST | `/api/v1/view/project/detail/get-one` | 获取项目详情内容 | admin |
| POST | `/api/v1/view/project/detail/update` | 更新项目详情 | admin |
| POST | `/api/v1/view/project/get-list` | 获取项目列表 | admin |
| POST | `/api/v1/view/project/get-one` | 获取项目详情 | admin |
| POST | `/api/v1/view/project/update` | 更新项目 | admin |
| POST | `/api/v1/view/system/energy/dashboard/device-status` | 获取能源大屏设备状态统计 | admin |
| POST | `/api/v1/view/system/energy/dashboard/overview` | 获取能源大屏顶部统计 | admin |
| POST | `/api/v1/view/system/uptime` | 获取平台运行时间 | admin |

## 端点详情

### POST `/api/v1/view/asset/delete`

**说明**: 删除资源

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 资源id |

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
ur api /api/v1/view/asset/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/asset/get-list`

**说明**: 获取资源列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `format` | string | 否 | 文件后缀筛选 |
| `groupId` | string | 否 | 分组筛选 |
| `name` | string | 否 | 名称模糊查询 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `type` | string | 否 | 类型筛选 |

**请求示例**:
```json
{
  "format": "string",
  "groupId": "string",
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
  "type": "string"
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
        "groupId": "string",
        "id": "string",
        "name": "示例名称",
        "size": 1,
        "tags": "string",
        "tenantCode": "string",
        "type": "string",
        "updatedTime": "2026-01-01T00:00:00Z",
        "url": "string"
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
ur api /api/v1/view/asset/get-list \
  --body '{"format": "string", "groupId": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "type": "string"}'
```

### POST `/api/v1/view/asset/update`

**说明**: 更新资源

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 资源id |
| `name` | string | 否 | 资源名称 |
| `tags` | string | 否 | 标签 |
| `type` | string | 否 | 类型 |

**请求示例**:
```json
{
  "id": "string",
  "name": "示例名称",
  "tags": "string",
  "type": "string"
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
ur api /api/v1/view/asset/update \
  --body '{"id": "string", "name": "示例名称", "tags": "string", "type": "string"}'
```

### POST `/api/v1/view/asset/upload`

**说明**: 上传资源

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupId` | string | 否 | 分组id |
| `name` | string | 否 | 资源名称 |
| `size` | integer | 否 | 文件大小(字节) (格式: int64) |
| `tags` | string | 否 | 标签,逗号分隔 |
| `type` | string | 否 | 类型: image/video/audio/other |
| `url` | string | 否 | 资源URL |

**请求示例**:
```json
{
  "groupId": "string",
  "name": "示例名称",
  "size": 1,
  "tags": "string",
  "type": "string",
  "url": "string"
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
ur api /api/v1/view/asset/upload \
  --body '{"groupId": "string", "name": "示例名称", "size": 1, "tags": "string", "type": "string", "url": "string"}'
```

### POST `/api/v1/view/component/copy`

**说明**: 复制组件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 组件id |

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
ur api /api/v1/view/component/copy \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/component/create`

**说明**: 新增组件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `categoryId` | string | 否 | 分类: chart/table/decorate/info/map/other |
| `content` | string | 否 | 组合模板JSON |
| `cover` | string | 否 | 预览图 |
| `description` | string | 否 | 说明 |
| `entry` | string | 否 | 组件入口 |
| `groupId` | string | 否 | 分组id |
| `name` | string | 是 | 组件名称 |
| `projectID` | string | 否 | 归属物联网项目ID |
| `version` | string | 否 | 版本 |

**请求示例**:
```json
{
  "categoryId": "string",
  "content": "string",
  "cover": "string",
  "description": "string",
  "entry": "string",
  "groupId": "string",
  "name": "示例名称",
  "projectID": "string",
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
ur api /api/v1/view/component/create \
  --body '{"categoryId": "string", "content": "string", "cover": "string", "description": "string", "entry": "string", "groupId": "string", "name": "示例名称", "projectID": "string", "version": "string"}'
```

### POST `/api/v1/view/component/delete`

**说明**: 删除组件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 组件id |

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
ur api /api/v1/view/component/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/component/design`

**说明**: 保存组件设计内容

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 | 组合模板JSON |
| `cover` | string | 否 | 预览图 |
| `id` | string | 是 | 组件id |

**请求示例**:
```json
{
  "content": "string",
  "cover": "string",
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
ur api /api/v1/view/component/design \
  --body '{"content": "string", "cover": "string", "id": "string"}'
```

### POST `/api/v1/view/component/get-list`

**说明**: 获取组件列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `categoryId` | string | 否 | 分类筛选 |
| `createdTime` | object | 否 |  |
| `createdTime.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `createdTime.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |
| `createdUserID` | string | 否 | 创建者ID筛选 |
| `groupId` | string | 否 | 分组筛选 |
| `name` | string | 否 | 名称模糊查询 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `projectID` | string | 否 | 归属物联网项目ID |
| `source` | string | 否 | 来源筛选 common/private |
| `status` | integer | 否 | 状态筛选 (格式: int64) |
| `updatedTime` | object | 否 |  |
| `updatedTime.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `updatedTime.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |

**请求示例**:
```json
{
  "categoryId": "string",
  "createdTime": {
    "end": 1,
    "start": 1
  },
  "createdUserID": "string",
  "groupId": "string",
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
  "projectID": "string",
  "source": "string",
  "status": 1,
  "updatedTime": {
    "end": 1,
    "start": 1
  }
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "categoryId": "string",
        "content": "string",
        "cover": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "createdUserID": "string",
        "description": "string",
        "entry": "string",
        "groupId": "string",
        "id": "string",
        "name": "示例名称",
        "projectID": "string",
        "source": "string",
        "status": 1,
        "tenantCode": "string",
        "updatedTime": "2026-01-01T00:00:00Z",
        "version": "string"
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
ur api /api/v1/view/component/get-list \
  --body '{"categoryId": "string", "createdTime": {"end": 1, "start": 1}, "createdUserID": "string", "groupId": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "projectID": "string", "source": "string", "status": 1, "updatedTime": {"end": 1, "start": 1}}'
```

### POST `/api/v1/view/component/get-one`

**说明**: 获取组件详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 组件id |

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
    "categoryId": "string",
    "content": "string",
    "cover": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "createdUserID": "string",
    "description": "string",
    "entry": "string",
    "groupId": "string",
    "id": "string",
    "name": "示例名称",
    "projectID": "string",
    "source": "string",
    "status": 1,
    "tenantCode": "string",
    "updatedTime": "2026-01-01T00:00:00Z",
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/component/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/component/update`

**说明**: 更新组件

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `categoryId` | string | 否 | 分类 |
| `content` | string | 否 | 组合模板JSON |
| `cover` | string | 否 | 预览图 |
| `description` | string | 否 | 说明 |
| `entry` | string | 否 | 组件入口 |
| `groupId` | string | 否 | 分组id |
| `id` | string | 是 | 组件id |
| `name` | string | 否 | 组件名称 |
| `projectID` | string | 否 | 归属物联网项目ID |
| `status` | integer | 否 | 状态 1:上线 2:下线 (格式: int64) |
| `version` | string | 否 | 版本 |

**请求示例**:
```json
{
  "categoryId": "string",
  "content": "string",
  "cover": "string",
  "description": "string",
  "entry": "string",
  "groupId": "string",
  "id": "string",
  "name": "示例名称",
  "projectID": "string",
  "status": 1,
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
ur api /api/v1/view/component/update \
  --body '{"categoryId": "string", "content": "string", "cover": "string", "description": "string", "entry": "string", "groupId": "string", "id": "string", "name": "示例名称", "projectID": "string", "status": 1, "version": "string"}'
```

### POST `/api/v1/view/dataset/create`

**说明**: 新增数据集

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cacheEnabled` | boolean | 否 | 是否启用缓存 (格式: boolean) |
| `cacheTTL` | integer | 否 | 缓存TTL(秒) (格式: int64) |
| `config` | string | 否 | 类型专属配置JSON |
| `desc` | string | 否 | 描述 |
| `dynamicParams` | string | 否 | 动态参数JSON |
| `groupId` | string | 否 | 分组id |
| `name` | string | 是 | 数据集名称 |
| `outputFields` | string | 否 | 输出字段JSON |
| `requestScript` | string | 否 | 请求前JS脚本 |
| `responseScript` | string | 否 | 响应后JS脚本 |
| `schema` | string | 否 | 字段定义JSON(兼容) |
| `sourceId` | string | 否 | 数据源id |
| `tags` | string | 否 | 标签,逗号分隔 |
| `type` | string | 否 | 类型: raw/sql/procedure/json/js/http/script/iot |

**请求示例**:
```json
{
  "cacheEnabled": true,
  "cacheTTL": 1,
  "config": "string",
  "desc": "string",
  "dynamicParams": "string",
  "groupId": "string",
  "name": "示例名称",
  "outputFields": "string",
  "requestScript": "string",
  "responseScript": "string",
  "schema": "string",
  "sourceId": "string",
  "tags": "string",
  "type": "string"
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
ur api /api/v1/view/dataset/create \
  --body '{"cacheEnabled": true, "cacheTTL": 1, "config": "string", "desc": "string", "dynamicParams": "string", "groupId": "string", "name": "示例名称", "outputFields": "string", "requestScript": "string", "responseScript": "string", "schema": "string", "sourceId": "string", "tags": "string", "type": "string"}'
```

### POST `/api/v1/view/dataset/delete`

**说明**: 删除数据集

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 数据集id |

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
ur api /api/v1/view/dataset/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/dataset/get-list`

**说明**: 获取数据集列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `groupId` | string | 否 | 分组筛选 |
| `name` | string | 否 | 名称模糊查询 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `scene` | string | 否 | 场景: manage/editor |
| `sourceId` | string | 否 | 数据源筛选 |
| `tags` | string | 否 | 标签筛选 |
| `type` | string | 否 | 类型筛选 |

**请求示例**:
```json
{
  "groupId": "string",
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
  "scene": "string",
  "sourceId": "string",
  "tags": "string",
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "cacheEnabled": true,
        "cacheTTL": 1,
        "config": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "displayName": "示例名称",
        "dynamicParams": "string",
        "groupId": "string",
        "id": "string",
        "name": "示例名称",
        "outputFields": "string",
        "requestScript": "string",
        "responseScript": "string",
        "schema": "string",
        "sourceId": "string",
        "sourceName": "示例名称",
        "tags": "string",
        "tenantCode": "string",
        "type": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/view/dataset/get-list \
  --body '{"groupId": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "scene": "string", "sourceId": "string", "tags": "string", "type": "string"}'
```

### POST `/api/v1/view/dataset/get-one`

**说明**: 获取数据集详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 数据集id |

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
    "cacheEnabled": true,
    "cacheTTL": 1,
    "config": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "displayName": "示例名称",
    "dynamicParams": "string",
    "groupId": "string",
    "id": "string",
    "name": "示例名称",
    "outputFields": "string",
    "requestScript": "string",
    "responseScript": "string",
    "schema": "string",
    "sourceId": "string",
    "sourceName": "示例名称",
    "tags": "string",
    "tenantCode": "string",
    "type": "string",
    "updatedTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/dataset/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/dataset/preview`

**说明**: 预览数据集

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 数据集id |

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
    "dataset": "string",
    "fields": "string",
    "rows": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/dataset/preview \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/dataset/update`

**说明**: 更新数据集

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `cacheEnabled` | boolean | 否 | 是否启用缓存 (格式: boolean) |
| `cacheTTL` | integer | 否 | 缓存TTL(秒) (格式: int64) |
| `config` | string | 否 | 类型专属配置JSON |
| `desc` | string | 否 | 描述 |
| `dynamicParams` | string | 否 | 动态参数JSON |
| `groupId` | string | 否 | 分组id |
| `id` | string | 是 | 数据集id |
| `name` | string | 否 | 数据集名称 |
| `outputFields` | string | 否 | 输出字段JSON |
| `requestScript` | string | 否 | 请求前JS脚本 |
| `responseScript` | string | 否 | 响应后JS脚本 |
| `schema` | string | 否 | 字段定义JSON(兼容) |
| `sourceId` | string | 否 | 数据源id |
| `tags` | string | 否 | 标签 |
| `type` | string | 否 | 类型 |

**请求示例**:
```json
{
  "cacheEnabled": true,
  "cacheTTL": 1,
  "config": "string",
  "desc": "string",
  "dynamicParams": "string",
  "groupId": "string",
  "id": "string",
  "name": "示例名称",
  "outputFields": "string",
  "requestScript": "string",
  "responseScript": "string",
  "schema": "string",
  "sourceId": "string",
  "tags": "string",
  "type": "string"
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
ur api /api/v1/view/dataset/update \
  --body '{"cacheEnabled": true, "cacheTTL": 1, "config": "string", "desc": "string", "dynamicParams": "string", "groupId": "string", "id": "string", "name": "示例名称", "outputFields": "string", "requestScript": "string", "responseScript": "string", "schema": "string", "sourceId": "string", "tags": "string", "type": "string"}'
```

### POST `/api/v1/view/datasource/create`

**说明**: 新增数据源

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `config` | string | 否 | 配置JSON |
| `desc` | string | 否 | 描述 |
| `name` | string | 是 | 数据源名称 |
| `type` | string | 是 | 类型: http/ws/mock/mysql/postgres |

**请求示例**:
```json
{
  "config": "string",
  "desc": "string",
  "name": "示例名称",
  "type": "string"
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
ur api /api/v1/view/datasource/create \
  --body '{"config": "string", "desc": "string", "name": "示例名称", "type": "string"}'
```

### POST `/api/v1/view/datasource/delete`

**说明**: 删除数据源

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 数据源id |

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
ur api /api/v1/view/datasource/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/datasource/get-list`

**说明**: 获取数据源列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 | 名称模糊查询 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `status` | integer | 否 | 状态筛选 (格式: int64) |
| `type` | string | 否 | 类型筛选 |

**请求示例**:
```json
{
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
  "type": "string"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "config": "string",
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "id": "string",
        "name": "示例名称",
        "status": 1,
        "tenantCode": "string",
        "type": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/view/datasource/get-list \
  --body '{"name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "status": 1, "type": "string"}'
```

### POST `/api/v1/view/datasource/get-one`

**说明**: 获取数据源详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 数据源id |

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
    "config": "string",
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "id": "string",
    "name": "示例名称",
    "status": 1,
    "tenantCode": "string",
    "type": "string",
    "updatedTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/datasource/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/datasource/test`

**说明**: 测试数据源连接

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 | 数据源id |

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
    "duration": 1,
    "message": "string",
    "success": true
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/datasource/test \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/datasource/update`

**说明**: 更新数据源

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `config` | string | 否 | 配置JSON |
| `desc` | string | 否 | 描述 |
| `id` | string | 是 | 数据源id |
| `name` | string | 否 | 数据源名称 |
| `type` | string | 否 | 类型 |

**请求示例**:
```json
{
  "config": "string",
  "desc": "string",
  "id": "string",
  "name": "示例名称",
  "type": "string"
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
ur api /api/v1/view/datasource/update \
  --body '{"config": "string", "desc": "string", "id": "string", "name": "示例名称", "type": "string"}'
```

### POST `/api/v1/view/group/create`

**说明**: 创建分组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 是 |  分组名称 |
| `sort` | integer | 否 |  排序 (格式: int64) |
| `type` | string | 是 |  分组类型 dataset/component/asset |

**请求示例**:
```json
{
  "name": "示例名称",
  "sort": 1,
  "type": "string"
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
ur api /api/v1/view/group/create \
  --body '{"name": "示例名称", "sort": 1, "type": "string"}'
```

### POST `/api/v1/view/group/delete`

**说明**: 删除分组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  分组id |

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
ur api /api/v1/view/group/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/group/get-list`

**说明**: 获取分组列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `scene` | string | 否 |  场景: manage/editor |
| `type` | string | 否 |  分组类型 dataset/component/asset |

**请求示例**:
```json
{
  "scene": "string",
  "type": "string"
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
        "id": "string",
        "name": "示例名称",
        "sort": 1,
        "tenantCode": "string",
        "type": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
      }
    ]
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/group/get-list \
  --body '{"scene": "string", "type": "string"}'
```

### POST `/api/v1/view/group/update`

**说明**: 更新分组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  分组id |
| `name` | string | 否 |  分组名称 |
| `sort` | integer | 否 |  排序 (格式: int64) |

**请求示例**:
```json
{
  "id": "string",
  "name": "示例名称",
  "sort": 1
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
ur api /api/v1/view/group/update \
  --body '{"id": "string", "name": "示例名称", "sort": 1}'
```

### POST `/api/v1/view/project/create`

**说明**: 新增项目

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  项目描述 |
| `indexImage` | string | 否 |  图片地址 |
| `isSystem` | boolean | 否 |  是否系统模板 (格式: boolean) |
| `name` | string | 否 |  项目名称 |
| `projectID` | string | 否 |  归属物联网项目ID |
| `status` | integer | 否 |  项目状态 1: 已发布 2: 未发布 (格式: int64) |
| `type` | string | 否 |  screen/template |

**请求示例**:
```json
{
  "desc": "string",
  "indexImage": "string",
  "isSystem": true,
  "name": "示例名称",
  "projectID": "string",
  "status": 1,
  "type": "string"
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
ur api /api/v1/view/project/create \
  --body '{"desc": "string", "indexImage": "string", "isSystem": true, "name": "示例名称", "projectID": "string", "status": 1, "type": "string"}'
```

### POST `/api/v1/view/project/delete`

**说明**: 删除项目

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  项目id |

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
ur api /api/v1/view/project/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/project/detail/get-one`

**说明**: 获取项目详情内容

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  项目id |

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
    "content": "string",
    "id": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/project/detail/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/project/detail/update`

**说明**: 更新项目详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 |  项目详情内容 |
| `id` | string | 否 |  项目id，只读 |

**请求示例**:
```json
{
  "content": "string",
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
ur api /api/v1/view/project/detail/update \
  --body '{"content": "string", "id": "string"}'
```

### POST `/api/v1/view/project/get-list`

**说明**: 获取项目列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | object | 否 |  |
| `createdTime.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `createdTime.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |
| `createdUserID` | string | 否 |  创建者ID |
| `name` | string | 否 |  名称模糊查询 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `projectID` | string | 否 |  归属物联网项目ID |
| `scope` | string | 否 |  all/private/system |
| `source` | string | 否 |  来源别名 all/custom/system |
| `status` | integer | 否 |  状态 (格式: int64) |
| `type` | string | 否 |  screen/template |
| `updatedTime` | object | 否 |  |
| `updatedTime.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `updatedTime.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": {
    "end": 1,
    "start": 1
  },
  "createdUserID": "string",
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
  "projectID": "string",
  "scope": "string",
  "source": "示例名称",
  "status": 1,
  "type": "string",
  "updatedTime": {
    "end": 1,
    "start": 1
  }
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
        "createdUserID": "string",
        "desc": "string",
        "id": "string",
        "indexImage": "string",
        "isSystem": true,
        "name": "示例名称",
        "projectID": "string",
        "resolution": "string",
        "source": "string",
        "status": 1,
        "tenantCode": "string",
        "type": "string",
        "updatedTime": "2026-01-01T00:00:00Z"
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
ur api /api/v1/view/project/get-list \
  --body '{"createdTime": {"end": 1, "start": 1}, "createdUserID": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "projectID": "string", "scope": "string", "source": "示例名称", "status": 1, "type": "string", "updatedTime": {"end": 1, "start": 1}}'
```

### POST `/api/v1/view/project/get-one`

**说明**: 获取项目详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  项目id |

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
    "createdUserID": "string",
    "desc": "string",
    "id": "string",
    "indexImage": "string",
    "isSystem": true,
    "name": "示例名称",
    "projectID": "string",
    "resolution": "string",
    "source": "string",
    "status": 1,
    "tenantCode": "string",
    "type": "string",
    "updatedTime": "2026-01-01T00:00:00Z"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/project/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/view/project/update`

**说明**: 更新项目

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  项目描述 |
| `id` | string | 是 |  项目id |
| `indexImage` | string | 否 |  图片地址 |
| `name` | string | 否 |  项目名称 |
| `projectID` | string | 否 |  归属物联网项目ID |
| `status` | integer | 否 |  项目状态 1: 已发布 2: 未发布 (格式: int64) |

**请求示例**:
```json
{
  "desc": "string",
  "id": "string",
  "indexImage": "string",
  "name": "示例名称",
  "projectID": "string",
  "status": 1
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
ur api /api/v1/view/project/update \
  --body '{"desc": "string", "id": "string", "indexImage": "string", "name": "示例名称", "projectID": "string", "status": 1}'
```

### POST `/api/v1/view/system/energy/dashboard/device-status`

**说明**: 获取能源大屏设备状态统计

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaIDs` | array[string] | 否 |  区域ID列表 |
| `limit` | integer | 否 |  返回条数限制 (格式: int64) |
| `projectID` | string | 否 |  项目ID |
| `timeEnd` | string | 否 |  自定义结束时间毫秒时间戳 |
| `timeRange` | string | 否 |  预设时间范围 |
| `timeStart` | string | 否 |  自定义开始时间毫秒时间戳 |

**请求示例**:
```json
{
  "areaIDs": [
    "string"
  ],
  "limit": 1,
  "projectID": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeRange": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "dataset": {
      "dimensions": [
        "string"
      ],
      "source": "string"
    },
    "summary": {
      "fault": 1,
      "inactive": 1,
      "offline": 1,
      "online": 1,
      "total": 1
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/system/energy/dashboard/device-status \
  --body '{"areaIDs": ["string"], "limit": 1, "projectID": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeRange": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/view/system/energy/dashboard/overview`

**说明**: 获取能源大屏顶部统计

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaIDs` | array[string] | 否 |  区域ID列表 |
| `limit` | integer | 否 |  返回条数限制 (格式: int64) |
| `projectID` | string | 否 |  项目ID |
| `timeEnd` | string | 否 |  自定义结束时间毫秒时间戳 |
| `timeRange` | string | 否 |  预设时间范围 |
| `timeStart` | string | 否 |  自定义开始时间毫秒时间戳 |

**请求示例**:
```json
{
  "areaIDs": [
    "string"
  ],
  "limit": 1,
  "projectID": "string",
  "timeEnd": "2026-01-01T00:00:00Z",
  "timeRange": "2026-01-01T00:00:00Z",
  "timeStart": "2026-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "avgPower": 1,
    "maxPower": 1,
    "minPower": 1,
    "todayEnergy": 1,
    "totalPower": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/system/energy/dashboard/overview \
  --body '{"areaIDs": ["string"], "limit": 1, "projectID": "string", "timeEnd": "2026-01-01T00:00:00Z", "timeRange": "2026-01-01T00:00:00Z", "timeStart": "2026-01-01T00:00:00Z"}'
```

### POST `/api/v1/view/system/uptime`

**说明**: 获取平台运行时间

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "days": 1,
    "hours": 1,
    "startTime": "2026-01-01T00:00:00Z",
    "uptimeSeconds": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/view/system/uptime \
  --body '{}'
```
