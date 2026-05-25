# ur-project things/group/info

创建分组 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/group/info/create` | 创建分组 | admin |
| POST | `/api/v1/things/group/info/delete` | 删除分组 | admin |
| POST | `/api/v1/things/group/info/get-list` | 获取分组列表 | admin |
| POST | `/api/v1/things/group/info/get-one` | 获取分组详情信息 | admin |
| POST | `/api/v1/things/group/info/update` | 更新分组信息 | admin |

## 详细说明

### POST `/api/v1/things/group/info/create`

**说明**: 创建分组

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 区域ID |
| `children` | array[GroupInfo] | 否 |  |
| `createdTime` | string | 否 | 创建时间 |
| `desc` | string | 否 | 分组描述 |
| `deviceCount` | integer | 否 | 设备数量统计 (格式: int64) |
| `devices` | array[DeviceCore] | 否 |  |
| `files` | object | 否 | key 是用途, value是文件的路径, |
| `id` | string | 否 | 分组ID |
| `idPath` | string | 否 | 分组ID路径 |
| `isLeaf` | integer | 否 | 格式: int64 |
| `name` | string | 否 | 分组名称 |
| `parentID` | string | 否 | 父组ID |
| `productID` | string | 否 | 产品ID |
| `productName` | string | 否 | 产品ID |
| `projectID` | string | 否 | 项目ID |
| `purpose` | string | 否 | 用途 不填默认为default |
| `tags` | object | 否 | 分组tag |

**请求示例**:
```json
{
  "areaID": "string",
  "children": [
    {
      "areaID": "string",
      "children": [
        {
          "areaID": "string",
          "children": [],
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [],
          "files": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "name": "示例名称",
          "parentID": "string",
          "productID": "string",
          "productName": "string",
          "projectID": "string",
          "purpose": "string",
          "tags": {}
        }
      ],
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "deviceCount": 1,
      "devices": [
        {
          "deviceName": "示例名称",
          "productID": "string",
          "productName": "string"
        }
      ],
      "files": {},
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "name": "示例名称",
      "parentID": "string",
      "productID": "string",
      "productName": "string",
      "projectID": "string",
      "purpose": "string",
      "tags": {}
    }
  ],
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "deviceCount": 1,
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
    }
  ],
  "files": {},
  "id": "string",
  "idPath": "string",
  "isLeaf": 1,
  "name": "示例名称",
  "parentID": "string",
  "productID": "string",
  "productName": "string",
  "projectID": "string",
  "purpose": "string",
  "tags": {}
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
ur api /api/v1/things/group/info/create \
  --body '{"areaID": "string", "children": [{"areaID": "string", "children": [{"areaID": "string", "children": [], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}'
```

### POST `/api/v1/things/group/info/delete`

**说明**: 删除分组

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
ur api /api/v1/things/group/info/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/group/info/get-list`

**说明**: 获取分组列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 区域ID |
| `name` | string | 否 | 分组名称 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `parentID` | string | 否 | 父组ID |
| `productID` | string | 否 | 产品ID |
| `purpose` | string | 否 | 用途 不填默认为default |
| `tags` | object | 否 | 分组tag |

**请求示例**:
```json
{
  "areaID": "string",
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
  "productID": "string",
  "purpose": "string",
  "tags": {}
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "areaID": "string",
        "children": [
          {
            "areaID": "string",
            "children": [],
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "files": {},
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "name": "示例名称",
            "parentID": "string",
            "productID": "string",
            "productName": "string",
            "projectID": "string",
            "purpose": "string",
            "tags": {}
          }
        ],
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "devices": [
          {
            "deviceName": "示例名称",
            "productID": "string",
            "productName": "string"
          }
        ],
        "files": {},
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "name": "示例名称",
        "parentID": "string",
        "productID": "string",
        "productName": "string",
        "projectID": "string",
        "purpose": "string",
        "tags": {}
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
ur api /api/v1/things/group/info/get-list \
  --body '{"areaID": "string", "name": "示例名称", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "parentID": "string", "productID": "string", "purpose": "string", "tags": {}}'
```

### POST `/api/v1/things/group/info/get-one`

**说明**: 获取分组详情信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  id |
| `purpose` | string | 否 | 用途 不填默认为default |
| `withChildren` | boolean | 否 | 格式: boolean |

**请求示例**:
```json
{
  "id": "string",
  "purpose": "string",
  "withChildren": true
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "areaID": "string",
    "children": [
      {
        "areaID": "string",
        "children": [
          {
            "areaID": "string",
            "children": [],
            "createdTime": "2026-01-01T00:00:00Z",
            "desc": "string",
            "deviceCount": 1,
            "devices": [],
            "files": {},
            "id": "string",
            "idPath": "string",
            "isLeaf": 1,
            "name": "示例名称",
            "parentID": "string",
            "productID": "string",
            "productName": "string",
            "projectID": "string",
            "purpose": "string",
            "tags": {}
          }
        ],
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "deviceCount": 1,
        "devices": [
          {
            "deviceName": "示例名称",
            "productID": "string",
            "productName": "string"
          }
        ],
        "files": {},
        "id": "string",
        "idPath": "string",
        "isLeaf": 1,
        "name": "示例名称",
        "parentID": "string",
        "productID": "string",
        "productName": "string",
        "projectID": "string",
        "purpose": "string",
        "tags": {}
      }
    ],
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "deviceCount": 1,
    "devices": [
      {
        "deviceName": "示例名称",
        "productID": "string",
        "productName": "string"
      }
    ],
    "files": {},
    "id": "string",
    "idPath": "string",
    "isLeaf": 1,
    "name": "示例名称",
    "parentID": "string",
    "productID": "string",
    "productName": "string",
    "projectID": "string",
    "purpose": "string",
    "tags": {}
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/group/info/get-one \
  --body '{"id": "string", "purpose": "string", "withChildren": true}'
```

### POST `/api/v1/things/group/info/update`

**说明**: 更新分组信息

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `areaID` | string | 否 | 区域ID |
| `children` | array[GroupInfo] | 否 |  |
| `createdTime` | string | 否 | 创建时间 |
| `desc` | string | 否 | 分组描述 |
| `deviceCount` | integer | 否 | 设备数量统计 (格式: int64) |
| `devices` | array[DeviceCore] | 否 |  |
| `files` | object | 否 | key 是用途, value是文件的路径, |
| `id` | string | 否 | 分组ID |
| `idPath` | string | 否 | 分组ID路径 |
| `isLeaf` | integer | 否 | 格式: int64 |
| `name` | string | 否 | 分组名称 |
| `parentID` | string | 否 | 父组ID |
| `productID` | string | 否 | 产品ID |
| `productName` | string | 否 | 产品ID |
| `projectID` | string | 否 | 项目ID |
| `purpose` | string | 否 | 用途 不填默认为default |
| `tags` | object | 否 | 分组tag |

**请求示例**:
```json
{
  "areaID": "string",
  "children": [
    {
      "areaID": "string",
      "children": [
        {
          "areaID": "string",
          "children": [],
          "createdTime": "2026-01-01T00:00:00Z",
          "desc": "string",
          "deviceCount": 1,
          "devices": [],
          "files": {},
          "id": "string",
          "idPath": "string",
          "isLeaf": 1,
          "name": "示例名称",
          "parentID": "string",
          "productID": "string",
          "productName": "string",
          "projectID": "string",
          "purpose": "string",
          "tags": {}
        }
      ],
      "createdTime": "2026-01-01T00:00:00Z",
      "desc": "string",
      "deviceCount": 1,
      "devices": [
        {
          "deviceName": "示例名称",
          "productID": "string",
          "productName": "string"
        }
      ],
      "files": {},
      "id": "string",
      "idPath": "string",
      "isLeaf": 1,
      "name": "示例名称",
      "parentID": "string",
      "productID": "string",
      "productName": "string",
      "projectID": "string",
      "purpose": "string",
      "tags": {}
    }
  ],
  "createdTime": "2026-01-01T00:00:00Z",
  "desc": "string",
  "deviceCount": 1,
  "devices": [
    {
      "deviceName": "示例名称",
      "productID": "string",
      "productName": "string"
    }
  ],
  "files": {},
  "id": "string",
  "idPath": "string",
  "isLeaf": 1,
  "name": "示例名称",
  "parentID": "string",
  "productID": "string",
  "productName": "string",
  "projectID": "string",
  "purpose": "string",
  "tags": {}
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
ur api /api/v1/things/group/info/update \
  --body '{"areaID": "string", "children": [{"areaID": "string", "children": [{"areaID": "string", "children": [], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}], "createdTime": "2026-01-01T00:00:00Z", "desc": "string", "deviceCount": 1, "devices": [{"deviceName": "示例名称", "productID": "string", "productName": "string"}], "files": {}, "id": "string", "idPath": "string", "isLeaf": 1, "name": "示例名称", "parentID": "string", "productID": "string", "productName": "string", "projectID": "string", "purpose": "string", "tags": {}}'
```
