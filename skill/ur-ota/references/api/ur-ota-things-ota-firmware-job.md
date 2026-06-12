# ur-ota things/ota/firmware/job

创建升级任务 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/ota/firmware/job/create` | 创建升级任务 | admin |
| POST | `/api/v1/things/ota/firmware/job/get-list` | 获取升级包下的升级任务批次列表 | admin |
| POST | `/api/v1/things/ota/firmware/job/get-one` | 查询指定升级批次的详情 | admin |
| POST | `/api/v1/things/ota/firmware/job/update` | 更新升级批次 | admin |

## 详细说明

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
| `tenantCodes` | array[string] | 否 | 指定企业 |
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
| `page.page` | integer | 否 |  页码（从1开始） (格式: int64) |
| `page.pageSize` | integer | 否 |  每页大小 (格式: int64) |

**请求示例**:
```json
{
  "firmwareID": "string",
  "page": {
    "page": 1,
    "pageSize": 1
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
  --body '{"firmwareID": "string", "page": {"page": 1, "pageSize": 1}}'
```

### POST `/api/v1/things/ota/firmware/job/get-one`

**说明**: 查询指定升级批次的详情

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
| `tenantCodes` | array[string] | 否 | 指定企业 |
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
