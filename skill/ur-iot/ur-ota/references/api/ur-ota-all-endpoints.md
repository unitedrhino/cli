# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/ota/firmware/device/cancel` | 取消指定任务下的升级 | admin |
| POST | `/api/v1/things/ota/firmware/device/confirm` | 确认升级设备 | admin |
| POST | `/api/v1/things/ota/firmware/device/get-list` | 查询升级设备列表 | admin |
| POST | `/api/v1/things/ota/firmware/device/retry` | 重试设备升级 | admin |
| POST | `/api/v1/things/ota/firmware/info/create` | 添加升级包 | admin |
| POST | `/api/v1/things/ota/firmware/info/delete` | 删除升级包 | admin |
| POST | `/api/v1/things/ota/firmware/info/get-list` | 升级包列表 | admin |
| POST | `/api/v1/things/ota/firmware/info/get-one` | 查询升级包 | admin |
| POST | `/api/v1/things/ota/firmware/info/update` | 修改升级包 | admin |
| POST | `/api/v1/things/ota/firmware/job/create` | 创建升级任务 | admin |
| POST | `/api/v1/things/ota/firmware/job/get-list` | 获取升级包下的升级任务批次列表 | admin |
| POST | `/api/v1/things/ota/firmware/job/get-one` | 查询指定升级批次的详情 | admin |
| POST | `/api/v1/things/ota/firmware/job/update` | 更新升级批次 | admin |
| POST | `/api/v1/things/ota/module/info/create` | 创建模块 | admin |
| POST | `/api/v1/things/ota/module/info/delete` | 删除模块 | admin |
| POST | `/api/v1/things/ota/module/info/get-list` | 获取模块列表 | admin |
| POST | `/api/v1/things/ota/module/info/get-one` | 查询模块详情 | admin |
| POST | `/api/v1/things/ota/module/info/update` | 更新模块 | admin |

## 端点详情

### POST `/api/v1/things/ota/firmware/device/cancel`

**说明**: 取消指定任务下的升级

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceNames` | array[string] | 是 |  |
| `firmwareID` | string | 是 |  |
| `jobID` | string | 是 |  |

**请求示例**:
```json
{
  "deviceNames": [
    "string"
  ],
  "firmwareID": "string",
  "jobID": "string"
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
ur api /api/v1/things/ota/firmware/device/cancel \
  --body '{"deviceNames": ["string"], "firmwareID": "string", "jobID": "string"}'
```

### POST `/api/v1/things/ota/firmware/device/confirm`

**说明**: 确认升级设备

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 是 |  |
| `productID` | string | 是 |  |

**请求示例**:
```json
{
  "deviceName": "string",
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
ur api /api/v1/things/ota/firmware/device/confirm \
  --body '{"deviceName": "string", "productID": "string"}'
```

### POST `/api/v1/things/ota/firmware/device/get-list`

**说明**: 查询升级设备列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceName` | string | 否 |  设备名称 |
| `firmwareID` | string | 否 |  任务状态 |
| `isOnline` | integer | 否 | 设备是否在线 (格式: int64) |
| `jobID` | string | 否 |  作业ID |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `srcVersion` | string | 否 |  |
| `status` | integer | 否 | 设备升级作业状态。1：待确认。 2：待推送。 3：已推送。  4：升级中。 5:升级成功 6: 升级失败. 7:已取消 (格式: int64) |

**请求示例**:
```json
{
  "deviceName": "示例名称",
  "firmwareID": "string",
  "isOnline": 1,
  "jobID": "string",
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
  "srcVersion": "string",
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
        "createdTime": "string",
        "destVersion": "string",
        "detail": "string",
        "deviceName": "string",
        "firmwareID": "string",
        "id": "string",
        "isOnline": 1,
        "jobID": "string",
        "productID": "string",
        "productName": "string",
        "srcVersion": "string",
        "status": 1,
        "step": 1,
        "updatedTime": 1
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
ur api /api/v1/things/ota/firmware/device/get-list \
  --body '{"deviceName": "示例名称", "firmwareID": "string", "isOnline": 1, "jobID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "srcVersion": "string", "status": 1}'
```

### POST `/api/v1/things/ota/firmware/device/retry`

**说明**: 重试设备升级

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `deviceNames` | array[string] | 是 |  |
| `firmwareID` | string | 是 |  |
| `jobID` | string | 是 |  |

**请求示例**:
```json
{
  "deviceNames": [
    "string"
  ],
  "firmwareID": "string",
  "jobID": "string"
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
ur api /api/v1/things/ota/firmware/device/retry \
  --body '{"deviceNames": ["string"], "firmwareID": "string", "jobID": "string"}'
```

### POST `/api/v1/things/ota/firmware/info/create`

**说明**: 添加升级包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `extra` | string | 否 |  |
| `filePaths` | array[string] | 是 |  |
| `isDiff` | integer | 是 | 格式: int64 |
| `isNeedToVerify` | integer | 是 | 格式: int64 |
| `moduleCode` | string | 否 |  |
| `name` | string | 是 |  |
| `productID` | string | 是 |  |
| `signMethod` | string | 是 |  |
| `srcVersion` | string | 是 |  |
| `version` | string | 是 |  |

**请求示例**:
```json
{
  "desc": "string",
  "extra": "string",
  "filePaths": [
    "string"
  ],
  "isDiff": 1,
  "isNeedToVerify": 1,
  "moduleCode": "string",
  "name": "string",
  "productID": "string",
  "signMethod": "string",
  "srcVersion": "string",
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
ur api /api/v1/things/ota/firmware/info/create \
  --body '{"desc": "string", "extra": "string", "filePaths": ["string"], "isDiff": 1, "isNeedToVerify": 1, "moduleCode": "string", "name": "string", "productID": "string", "signMethod": "string", "srcVersion": "string", "version": "string"}'
```

### POST `/api/v1/things/ota/firmware/info/delete`

**说明**: 删除升级包

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
ur api /api/v1/things/ota/firmware/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/ota/firmware/info/get-list`

**说明**: 升级包列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 |  |

**请求示例**:
```json
{
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
        "createdTime": "2026-01-01T00:00:00Z",
        "desc": "string",
        "extra": "string",
        "fileList": [
          {
            "fileMd5": "string",
            "filePath": "string",
            "name": "string",
            "signature": "string"
          }
        ],
        "id": "string",
        "isDiff": 1,
        "isNeedToVerify": 1,
        "moduleCode": "string",
        "name": "示例名称",
        "productID": "string",
        "productName": "示例名称",
        "signMethod": "string",
        "status": 1,
        "version": "string"
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
ur api /api/v1/things/ota/firmware/info/get-list \
  --body '{"name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string"}'
```

### POST `/api/v1/things/ota/firmware/info/get-one`

**说明**: 查询升级包

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
  "data": {
    "createdTime": "2026-01-01T00:00:00Z",
    "desc": "string",
    "extra": "string",
    "fileList": [
      {
        "fileMd5": "string",
        "filePath": "string",
        "name": "string",
        "signature": "string"
      }
    ],
    "id": "string",
    "isDiff": 1,
    "isNeedToVerify": 1,
    "moduleCode": "string",
    "name": "示例名称",
    "productID": "string",
    "productName": "示例名称",
    "signMethod": "string",
    "status": 1,
    "version": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/ota/firmware/info/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/things/ota/firmware/info/update`

**说明**: 修改升级包

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  |
| `extra` | string | 否 |  |
| `id` | string | 是 |  |
| `name` | string | 否 |  |

**请求示例**:
```json
{
  "desc": "string",
  "extra": "string",
  "id": "string",
  "name": "string"
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
ur api /api/v1/things/ota/firmware/info/update \
  --body '{"desc": "string", "extra": "string", "id": "string", "name": "string"}'
```

### POST `/api/v1/things/ota/firmware/job/create`

**说明**: 创建升级任务

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 |  |
| `dynamicMode` | integer | 否 | 格式: int64 |
| `firmwareID` | string | 否 |  升级包ID，升级包的唯一标识符。 |
| `grayPercent` | integer | 否 | 灰度的范围,小数点后两位, 1.23%为 123 (格式: int64) |
| `id` | string | 否 |  |
| `isNeedConfirm` | integer | 否 | 格式: int64 |
| `isNeedPush` | integer | 否 |  物联网平台是否主动向设备推送升级任务。 (格式: int64) |
| `isOverwriteMode` | integer | 否 |  是否覆盖之前的升级任务。取值：2（不覆盖）、1（覆盖）。 (格式: int64) |
| `maximumPerMinute` | integer | 否 |  每分钟最多向多少个设备推送升级包下载URL。 动态升级 静态升级 (格式: int64) |
| `retryCount` | integer | 否 |  自动重试次数。1/2/5次 动态升级 静态升级 (格式: int64) |
| `retryInterval` | integer | 否 |  设备升级失败后，自动重试的时间间隔，单位为分钟。 动态升级 静态升级 (格式: int64) |
| `scheduleFinishTime` | integer | 否 |  指定结束升级的时间，单位为毫秒。 (格式: int64) |
| `scheduleTime` | integer | 否 | 格式: int64 |
| `srcVersions` | array[string] | 否 |  待升级版本号列表。最多可传入10个版本号。用逗号分隔多个版本号 |
| `status` | integer | 否 |  批次状态,计划中:1  执行中:2  已完成:3  已取消:4 (格式: int64) |
| `target` | string | 否 |  分组升级和区域升级填写对应的id |
| `targetDeviceNames` | array[string] | 否 |  定向升级的设备名称列表。最多可传入200个设备名称。以逗号分隔 |
| `targetSelection` | integer | 否 | 升级范围。 1：全量升级。 2：定向升级。 3：灰度升级。 4：分组升级(不做) 5: 区域升级(不做) (格式: int64) |
| `tenantCodes` | array[string] | 否 | 指定租户 |
| `timeoutInMinutes` | integer | 否 |  设备升级超时时间，单位为分钟。 动态升级 静态升级 (格式: int64) |
| `type` | integer | 否 |  升级包所属产品的JobType。 验证升级包:1  批量升级:2 (格式: int64) |
| `upgradeType` | integer | 否 |  升级策略，1-静态，2-动态 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": "string",
  "dynamicMode": 1,
  "firmwareID": "string",
  "grayPercent": 1,
  "id": "string",
  "isNeedConfirm": 1,
  "isNeedPush": 1,
  "isOverwriteMode": 1,
  "maximumPerMinute": 1,
  "retryCount": 1,
  "retryInterval": 1,
  "scheduleFinishTime": 1,
  "scheduleTime": 1,
  "srcVersions": [
    "string"
  ],
  "status": 1,
  "target": "string",
  "targetDeviceNames": [
    "string"
  ],
  "targetSelection": 1,
  "tenantCodes": [
    "string"
  ],
  "timeoutInMinutes": 1,
  "type": 1,
  "upgradeType": 1
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
ur api /api/v1/things/ota/firmware/job/create \
  --body '{"createdTime": "string", "dynamicMode": 1, "firmwareID": "string", "grayPercent": 1, "id": "string", "isNeedConfirm": 1, "isNeedPush": 1, "isOverwriteMode": 1, "maximumPerMinute": 1, "retryCount": 1, "retryInterval": 1, "scheduleFinishTime": 1, "scheduleTime": 1, "srcVersions": ["string"], "status": 1, "target": "string", "targetDeviceNames": ["string"], "targetSelection": 1, "tenantCodes": ["string"], "timeoutInMinutes": 1, "type": 1, "upgradeType": 1}'
```

### POST `/api/v1/things/ota/firmware/job/get-list`

**说明**: 获取升级包下的升级任务批次列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `firmwareID` | string | 是 |  固件ID |
| `page` | object | 是 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "firmwareID": "string",
  "page": {
    "orders": [
      {
        "field": "string",
        "sort": 1
      }
    ],
    "page": 1,
    "size": 1
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
        "createdTime": "string",
        "dynamicMode": 1,
        "firmwareID": "string",
        "grayPercent": 1,
        "id": "string",
        "isNeedConfirm": 1,
        "isNeedPush": 1,
        "isOverwriteMode": 1,
        "maximumPerMinute": 1,
        "retryCount": 1,
        "retryInterval": 1,
        "scheduleFinishTime": 1,
        "scheduleTime": 1,
        "srcVersions": [
          "string"
        ],
        "status": 1,
        "target": "string",
        "targetDeviceNames": [
          "string"
        ],
        "targetSelection": 1,
        "tenantCodes": [
          "string"
        ],
        "timeoutInMinutes": 1,
        "type": 1,
        "upgradeType": 1
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
ur api /api/v1/things/ota/firmware/job/get-list \
  --body '{"firmwareID": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}}'
```

### POST `/api/v1/things/ota/firmware/job/get-one`

**说明**: 查询指定升级批次的详情

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
  "data": {
    "createdTime": "string",
    "dynamicMode": 1,
    "firmwareID": "string",
    "grayPercent": 1,
    "id": "string",
    "isNeedConfirm": 1,
    "isNeedPush": 1,
    "isOverwriteMode": 1,
    "maximumPerMinute": 1,
    "retryCount": 1,
    "retryInterval": 1,
    "scheduleFinishTime": 1,
    "scheduleTime": 1,
    "srcVersions": [
      "string"
    ],
    "status": 1,
    "target": "string",
    "targetDeviceNames": [
      "string"
    ],
    "targetSelection": 1,
    "tenantCodes": [
      "string"
    ],
    "timeoutInMinutes": 1,
    "type": 1,
    "upgradeType": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/ota/firmware/job/get-one \
  --body '{"id": 1}'
```

### POST `/api/v1/things/ota/firmware/job/update`

**说明**: 更新升级批次

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTime` | string | 否 |  |
| `dynamicMode` | integer | 否 | 格式: int64 |
| `firmwareID` | string | 否 |  升级包ID，升级包的唯一标识符。 |
| `grayPercent` | integer | 否 | 灰度的范围,小数点后两位, 1.23%为 123 (格式: int64) |
| `id` | string | 否 |  |
| `isNeedConfirm` | integer | 否 | 格式: int64 |
| `isNeedPush` | integer | 否 |  物联网平台是否主动向设备推送升级任务。 (格式: int64) |
| `isOverwriteMode` | integer | 否 |  是否覆盖之前的升级任务。取值：2（不覆盖）、1（覆盖）。 (格式: int64) |
| `maximumPerMinute` | integer | 否 |  每分钟最多向多少个设备推送升级包下载URL。 动态升级 静态升级 (格式: int64) |
| `retryCount` | integer | 否 |  自动重试次数。1/2/5次 动态升级 静态升级 (格式: int64) |
| `retryInterval` | integer | 否 |  设备升级失败后，自动重试的时间间隔，单位为分钟。 动态升级 静态升级 (格式: int64) |
| `scheduleFinishTime` | integer | 否 |  指定结束升级的时间，单位为毫秒。 (格式: int64) |
| `scheduleTime` | integer | 否 | 格式: int64 |
| `srcVersions` | array[string] | 否 |  待升级版本号列表。最多可传入10个版本号。用逗号分隔多个版本号 |
| `status` | integer | 否 |  批次状态,计划中:1  执行中:2  已完成:3  已取消:4 (格式: int64) |
| `target` | string | 否 |  分组升级和区域升级填写对应的id |
| `targetDeviceNames` | array[string] | 否 |  定向升级的设备名称列表。最多可传入200个设备名称。以逗号分隔 |
| `targetSelection` | integer | 否 | 升级范围。 1：全量升级。 2：定向升级。 3：灰度升级。 4：分组升级(不做) 5: 区域升级(不做) (格式: int64) |
| `tenantCodes` | array[string] | 否 | 指定租户 |
| `timeoutInMinutes` | integer | 否 |  设备升级超时时间，单位为分钟。 动态升级 静态升级 (格式: int64) |
| `type` | integer | 否 |  升级包所属产品的JobType。 验证升级包:1  批量升级:2 (格式: int64) |
| `upgradeType` | integer | 否 |  升级策略，1-静态，2-动态 (格式: int64) |

**请求示例**:
```json
{
  "createdTime": "string",
  "dynamicMode": 1,
  "firmwareID": "string",
  "grayPercent": 1,
  "id": "string",
  "isNeedConfirm": 1,
  "isNeedPush": 1,
  "isOverwriteMode": 1,
  "maximumPerMinute": 1,
  "retryCount": 1,
  "retryInterval": 1,
  "scheduleFinishTime": 1,
  "scheduleTime": 1,
  "srcVersions": [
    "string"
  ],
  "status": 1,
  "target": "string",
  "targetDeviceNames": [
    "string"
  ],
  "targetSelection": 1,
  "tenantCodes": [
    "string"
  ],
  "timeoutInMinutes": 1,
  "type": 1,
  "upgradeType": 1
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
ur api /api/v1/things/ota/firmware/job/update \
  --body '{"createdTime": "string", "dynamicMode": 1, "firmwareID": "string", "grayPercent": 1, "id": "string", "isNeedConfirm": 1, "isNeedPush": 1, "isOverwriteMode": 1, "maximumPerMinute": 1, "retryCount": 1, "retryInterval": 1, "scheduleFinishTime": 1, "scheduleTime": 1, "srcVersions": ["string"], "status": 1, "target": "string", "targetDeviceNames": ["string"], "targetSelection": 1, "tenantCodes": ["string"], "timeoutInMinutes": 1, "type": 1, "upgradeType": 1}'
```

### POST `/api/v1/things/ota/module/info/create`

**说明**: 创建模块

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `productID` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "string",
  "desc": "string",
  "id": "string",
  "name": "string",
  "productID": "string"
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
ur api /api/v1/things/ota/module/info/create \
  --body '{"code": "string", "createdTime": "string", "desc": "string", "id": "string", "name": "string", "productID": "string"}'
```

### POST `/api/v1/things/ota/module/info/delete`

**说明**: 删除模块

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
ur api /api/v1/things/ota/module/info/delete \
  --body '{"id": 1}'
```

### POST `/api/v1/things/ota/module/info/get-list`

**说明**: 获取模块列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | 否 |  |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 |  |

**请求示例**:
```json
{
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
        "code": "string",
        "createdTime": "string",
        "desc": "string",
        "id": "string",
        "name": "string",
        "productID": "string"
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
ur api /api/v1/things/ota/module/info/get-list \
  --body '{"name": "string", "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string"}'
```

### POST `/api/v1/things/ota/module/info/get-one`

**说明**: 查询模块详情

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `id` | integer | 否 |  id (格式: int64) |

**请求示例**:
```json
{
  "code": "string",
  "id": 1
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "code": "string",
    "createdTime": "string",
    "desc": "string",
    "id": "string",
    "name": "string",
    "productID": "string"
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/ota/module/info/get-one \
  --body '{"code": "string", "id": 1}'
```

### POST `/api/v1/things/ota/module/info/update`

**说明**: 更新模块

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `code` | string | 否 |  |
| `createdTime` | string | 否 |  |
| `desc` | string | 否 |  |
| `id` | string | 否 |  |
| `name` | string | 否 |  |
| `productID` | string | 否 |  |

**请求示例**:
```json
{
  "code": "string",
  "createdTime": "string",
  "desc": "string",
  "id": "string",
  "name": "string",
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
ur api /api/v1/things/ota/module/info/update \
  --body '{"code": "string", "createdTime": "string", "desc": "string", "id": "string", "name": "string", "productID": "string"}'
```
