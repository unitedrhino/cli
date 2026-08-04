# ur-system im/moment

发表评论 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/system/im/moment/comment/create` | 发表评论 | all |
| POST | `/api/v1/system/im/moment/comment/delete` | 删除评论 | all |
| POST | `/api/v1/system/im/moment/create` | 发布动态 | all |
| POST | `/api/v1/system/im/moment/delete` | 删除动态 | all |
| POST | `/api/v1/system/im/moment/get` | 获取动态详情 | all |
| POST | `/api/v1/system/im/moment/like` | 点赞动态 | all |
| POST | `/api/v1/system/im/moment/list` | 查询动态列表 | all |
| POST | `/api/v1/system/im/moment/unlike` | 取消点赞动态 | all |

## 详细说明

### POST `/api/v1/system/im/moment/comment/create`

**说明**: 发表评论

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | object | 是 |  |
| `content.content` | string | 是 |  评论内容 |
| `content.mentionUserIDs` | array[integer] | 是 |  @ 的用户列表 |
| `id` | integer | 是 |  动态 ID (格式: int64) |

**请求示例**:
```json
{
  "content": {
    "content": "string",
    "mentionUserIDs": [
      1
    ]
  },
  "id": 1
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
ur api /api/v1/system/im/moment/comment/create \
  --body '{"content": {"content": "string", "mentionUserIDs": [1]}, "id": 1}'
```

### POST `/api/v1/system/im/moment/comment/delete`

**说明**: 删除评论

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `commentID` | integer | 是 |  评论 ID (格式: int64) |
| `momentID` | integer | 是 |  动态 ID (格式: int64) |

**请求示例**:
```json
{
  "commentID": 1,
  "momentID": 1
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
ur api /api/v1/system/im/moment/comment/delete \
  --body '{"commentID": 1, "momentID": 1}'
```

### POST `/api/v1/system/im/moment/create`

**说明**: 发布动态

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `content` | string | 是 |  文本内容 |
| `mediaList` | array[ImMomentMedia] | 是 |  媒体列表 |
| `mentionUserIDs` | array[integer] | 是 |  @ 的用户列表 |
| `projectIDs` | array[integer] | 是 |  可见项目列表（visibility=1 时生效） |
| `visibility` | integer | 是 |  可见范围：0 公开 / 1 项目 / 2 指定人 (格式: int64) |
| `visibleUserIDs` | array[integer] | 是 |  可见用户列表（visibility=2 时生效） |

**请求示例**:
```json
{
  "content": "string",
  "mediaList": [
    {
      "duration": 1,
      "name": "string",
      "size": 1,
      "thumb": "string",
      "type": "string",
      "url": "string"
    }
  ],
  "mentionUserIDs": [
    1
  ],
  "projectIDs": [
    1
  ],
  "visibility": 1,
  "visibleUserIDs": [
    1
  ]
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
ur api /api/v1/system/im/moment/create \
  --body '{"content": "string", "mediaList": [{"duration": 1, "name": "string", "size": 1, "thumb": "string", "type": "string", "url": "string"}], "mentionUserIDs": [1], "projectIDs": [1], "visibility": 1, "visibleUserIDs": [1]}'
```

### POST `/api/v1/system/im/moment/delete`

**说明**: 删除动态

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  动态 ID (格式: int64) |

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
ur api /api/v1/system/im/moment/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/system/im/moment/get`

**说明**: 获取动态详情

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  动态 ID (格式: int64) |

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
    "comments": [
      {
        "content": "string",
        "createdTime": 1,
        "id": 1,
        "mentionUserIDs": [
          1
        ],
        "momentID": 1,
        "status": 1,
        "updatedTime": 1,
        "userID": 1
      }
    ],
    "info": {
      "commentCount": 1,
      "content": "string",
      "createdTime": 1,
      "id": 1,
      "isLiked": true,
      "likeCount": 1,
      "mediaList": [
        {
          "duration": 1,
          "name": "string",
          "size": 1,
          "thumb": "string",
          "type": "string",
          "url": "string"
        }
      ],
      "mentionUserIDs": [
        1
      ],
      "projectIDs": [
        1
      ],
      "status": 1,
      "updatedTime": 1,
      "userID": 1,
      "visibility": 1,
      "visibleUserIDs": [
        1
      ]
    }
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/moment/get \
  --body '{"id": 1}'
```

### POST `/api/v1/system/im/moment/like`

**说明**: 点赞动态

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  动态 ID (格式: int64) |

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
    "likeCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/moment/like \
  --body '{"id": 1}'
```

### POST `/api/v1/system/im/moment/list`

**说明**: 查询动态列表

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `filterProjectID` | integer | 否 |  按项目过滤 (格式: int64) |
| `filterUserID` | integer | 否 |  按发布人过滤 (格式: int64) |
| `page` | integer | 是 |  页码 (格式: int64) |
| `pageSize` | integer | 是 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "filterProjectID": 1,
  "filterUserID": 1,
  "page": 1,
  "pageSize": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "commentCount": 1,
        "content": "string",
        "createdTime": 1,
        "id": 1,
        "isLiked": true,
        "likeCount": 1,
        "mediaList": [
          {
            "duration": 1,
            "name": "string",
            "size": 1,
            "thumb": "string",
            "type": "string",
            "url": "string"
          }
        ],
        "mentionUserIDs": [
          1
        ],
        "projectIDs": [
          1
        ],
        "status": 1,
        "updatedTime": 1,
        "userID": 1,
        "visibility": 1,
        "visibleUserIDs": [
          1
        ]
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/moment/list \
  --body '{"filterProjectID": 1, "filterUserID": 1, "page": 1, "pageSize": 1}'
```

### POST `/api/v1/system/im/moment/unlike`

**说明**: 取消点赞动态

**权限**: all

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | integer | 是 |  动态 ID (格式: int64) |

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
    "likeCount": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/system/im/moment/unlike \
  --body '{"id": 1}'
```
