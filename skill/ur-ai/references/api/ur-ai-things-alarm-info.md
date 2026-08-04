# ur-ai things/alarm/info

创建告警规则 等

## 端点概览

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/alarm/info/create` | 创建告警规则 | admin |
| POST | `/api/v1/things/alarm/info/delete` | 删除告警规则 | admin |
| POST | `/api/v1/things/alarm/info/evaluate-trigger` | 手动触发评估 | admin |
| POST | `/api/v1/things/alarm/info/get-list` | 获取告警规则列表 | admin |
| POST | `/api/v1/things/alarm/info/get-one` | 获取告警规则详情 | admin |
| POST | `/api/v1/things/alarm/info/status-update` | 更新告警规则状态 | admin |
| POST | `/api/v1/things/alarm/info/update` | 更新告警规则 | admin |

## 详细说明

### POST `/api/v1/things/alarm/info/create`

**说明**: 创建告警规则

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  描述 |
| `evalPeriod` | integer | 否 |  评估周期（分钟） (格式: int64) |
| `if` | object | 否 |  |
| `if.triggers` | array[AlarmTrigger] | 否 |  触发器列表 |
| `ladders` | array[AlarmNotifyLadder] | 否 |  通知阶梯 |
| `name` | string | 是 |  名称 |
| `target` | object | 否 |  |
| `target.areas` | array[integer] | 否 |  区域ID列表 |
| `target.categoryID` | string | 否 |  产品品类ID |
| `target.devices` | array[AlarmTargetDevice] | 否 |  指定设备列表 |
| `target.groups` | array[integer] | 否 |  分组ID列表 |
| `target.productID` | string | 否 |  产品ID |
| `target.selectType` | string | 否 |  all / fixed / area / areaWithChildren / group |

**请求示例**:
```json
{
  "desc": "string",
  "evalPeriod": 1,
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
  "ladders": [
    {
      "callbacks": [
        {
          "authToken": "string",
          "authType": "string",
          "bodyTemplate": "string",
          "headers": {},
          "method": "string",
          "name": "示例名称",
          "notifyCycle": "string",
          "notifyEndTime": "2026-01-01T00:00:00Z",
          "notifyStartTime": "2026-01-01T00:00:00Z",
          "retryTimes": 1,
          "timeoutSeconds": 1,
          "url": "string"
        }
      ],
      "channelTemplates": {},
      "channels": [
        "string"
      ],
      "delaySeconds": 1,
      "levels": [
        "string"
      ],
      "order": 1,
      "targets": {
        "groupIDs": [
          "string"
        ],
        "roleIDs": [
          "string"
        ],
        "userIDs": [
          "string"
        ]
      },
      "timing": "string"
    }
  ],
  "name": "示例名称",
  "target": {
    "areas": [
      1
    ],
    "categoryID": "string",
    "devices": [
      {
        "alias": "示例名称",
        "name": "string",
        "productID": "string"
      }
    ],
    "groups": [
      1
    ],
    "productID": "string",
    "selectType": "string"
  }
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
ur api /api/v1/things/alarm/info/create \
  --body '{"desc": "string", "evalPeriod": 1, "if": {"triggers": [{"criteria": [{"duration": "...", "frequency": "...", "id": "...", "kOfN": "...", "order": "...", "plain": "...", "rollingAggregate": "...", "type": "..."}], "id": "string", "level": "string", "templateID": "string", "templateName": "示例名称", "type": "string"}]}, "ladders": [{"callbacks": [{"authToken": "string", "authType": "string", "bodyTemplate": "string", "headers": {}, "method": "string", "name": "示例名称", "notifyCycle": "string", "notifyEndTime": "2026-01-01T00:00:00Z", "notifyStartTime": "2026-01-01T00:00:00Z", "retryTimes": 1, "timeoutSeconds": 1, "url": "string"}], "channelTemplates": {}, "channels": ["string"], "delaySeconds": 1, "levels": ["string"], "order": 1, "targets": {"groupIDs": ["string"], "roleIDs": ["string"], "userIDs": ["string"]}, "timing": "string"}], "name": "示例名称", "target": {"areas": [1], "categoryID": "string", "devices": [{"alias": "示例名称", "name": "string", "productID": "string"}], "groups": [1], "productID": "string", "selectType": "string"}}'
```

### POST `/api/v1/things/alarm/info/delete`

**说明**: 删除告警规则

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
ur api /api/v1/things/alarm/info/delete \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/info/evaluate-trigger`

**说明**: 手动触发评估

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `alarmIDs` | array[integer] | 否 |  告警ID列表，为空表示全量扫描 |

**请求示例**:
```json
{
  "alarmIDs": [
    1
  ]
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
ur api /api/v1/things/alarm/info/evaluate-trigger \
  --body '{"alarmIDs": [1]}'
```

### POST `/api/v1/things/alarm/info/get-list`

**说明**: 获取告警规则列表

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `createdTimeRange` | object | 否 |  |
| `createdTimeRange.end` | integer | 否 | 结束时间 unix时间戳 (格式: int64) |
| `createdTimeRange.start` | integer | 否 | 开始时间 unix时间戳 (格式: int64) |
| `keyword` | string | 否 |  名称模糊搜索 |
| `levels` | array[string] | 否 |  触发条件级别 |
| `page` | object | 否 |  |
| `page.orders` | array[OrderBy] | 否 | 排序 |
| `page.page` | integer | 否 |  页码 (格式: int64) |
| `page.size` | integer | 否 |  每页大小 (格式: int64) |
| `productID` | string | 否 |  产品ID |
| `status` | integer | 否 |  状态 (格式: int64) |

**请求示例**:
```json
{
  "createdTimeRange": {
    "end": 1,
    "start": 1
  },
  "keyword": "示例名称",
  "levels": [
    "string"
  ],
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
        "evalPeriod": 1,
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
        "ladders": [
          {
            "callbacks": [],
            "channelTemplates": {},
            "channels": [],
            "delaySeconds": 1,
            "levels": [],
            "order": 1,
            "targets": {
              "groupIDs": "...",
              "roleIDs": "...",
              "userIDs": "..."
            },
            "timing": "string"
          }
        ],
        "name": "示例名称",
        "projectID": "string",
        "status": 1,
        "target": {
          "areas": [
            1
          ],
          "categoryID": "string",
          "devices": [
            {
              "alias": "...",
              "name": "...",
              "productID": "..."
            }
          ],
          "groups": [
            1
          ],
          "productID": "string",
          "selectType": "string"
        },
        "targetName": "string",
        "updatedTime": "2026-01-01T00:00:00Z",
        "version": 1
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
ur api /api/v1/things/alarm/info/get-list \
  --body '{"createdTimeRange": {"end": 1, "start": 1}, "keyword": "示例名称", "levels": ["string"], "page": {"orders": [{"field": "string", "sort": 1}], "page": 1, "size": 1}, "productID": "string", "status": 1}'
```

### POST `/api/v1/things/alarm/info/get-one`

**说明**: 获取告警规则详情

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
    "evalPeriod": 1,
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
    "ladders": [
      {
        "callbacks": [
          {
            "authToken": "string",
            "authType": "string",
            "bodyTemplate": "string",
            "headers": {},
            "method": "string",
            "name": "示例名称",
            "notifyCycle": "string",
            "notifyEndTime": "2026-01-01T00:00:00Z",
            "notifyStartTime": "2026-01-01T00:00:00Z",
            "retryTimes": 1,
            "timeoutSeconds": 1,
            "url": "string"
          }
        ],
        "channelTemplates": {},
        "channels": [
          "string"
        ],
        "delaySeconds": 1,
        "levels": [
          "string"
        ],
        "order": 1,
        "targets": {
          "groupIDs": [
            "string"
          ],
          "roleIDs": [
            "string"
          ],
          "userIDs": [
            "string"
          ]
        },
        "timing": "string"
      }
    ],
    "name": "示例名称",
    "projectID": "string",
    "status": 1,
    "target": {
      "areas": [
        1
      ],
      "categoryID": "string",
      "devices": [
        {
          "alias": "示例名称",
          "name": "string",
          "productID": "string"
        }
      ],
      "groups": [
        1
      ],
      "productID": "string",
      "selectType": "string"
    },
    "targetName": "string",
    "updatedTime": "2026-01-01T00:00:00Z",
    "version": 1
  },
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/alarm/info/get-one \
  --body '{"id": "string"}'
```

### POST `/api/v1/things/alarm/info/status-update`

**说明**: 更新告警规则状态

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `id` | string | 是 |  ID |
| `status` | integer | 是 |  状态：1启用 2停用 (格式: int64) |

**请求示例**:
```json
{
  "id": "string",
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
ur api /api/v1/things/alarm/info/status-update \
  --body '{"id": "string", "status": 1}'
```

### POST `/api/v1/things/alarm/info/update`

**说明**: 更新告警规则

**权限**: admin

**请求体字段**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `desc` | string | 否 |  描述 |
| `evalPeriod` | integer | 否 |  评估周期（分钟） (格式: int64) |
| `id` | string | 是 |  ID |
| `if` | object | 否 |  |
| `if.triggers` | array[AlarmTrigger] | 否 |  触发器列表 |
| `ladders` | array[AlarmNotifyLadder] | 否 |  通知阶梯 |
| `name` | string | 否 |  名称 |
| `target` | object | 否 |  |
| `target.areas` | array[integer] | 否 |  区域ID列表 |
| `target.categoryID` | string | 否 |  产品品类ID |
| `target.devices` | array[AlarmTargetDevice] | 否 |  指定设备列表 |
| `target.groups` | array[integer] | 否 |  分组ID列表 |
| `target.productID` | string | 否 |  产品ID |
| `target.selectType` | string | 否 |  all / fixed / area / areaWithChildren / group |

**请求示例**:
```json
{
  "desc": "string",
  "evalPeriod": 1,
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
  "ladders": [
    {
      "callbacks": [
        {
          "authToken": "string",
          "authType": "string",
          "bodyTemplate": "string",
          "headers": {},
          "method": "string",
          "name": "示例名称",
          "notifyCycle": "string",
          "notifyEndTime": "2026-01-01T00:00:00Z",
          "notifyStartTime": "2026-01-01T00:00:00Z",
          "retryTimes": 1,
          "timeoutSeconds": 1,
          "url": "string"
        }
      ],
      "channelTemplates": {},
      "channels": [
        "string"
      ],
      "delaySeconds": 1,
      "levels": [
        "string"
      ],
      "order": 1,
      "targets": {
        "groupIDs": [
          "string"
        ],
        "roleIDs": [
          "string"
        ],
        "userIDs": [
          "string"
        ]
      },
      "timing": "string"
    }
  ],
  "name": "示例名称",
  "target": {
    "areas": [
      1
    ],
    "categoryID": "string",
    "devices": [
      {
        "alias": "示例名称",
        "name": "string",
        "productID": "string"
      }
    ],
    "groups": [
      1
    ],
    "productID": "string",
    "selectType": "string"
  }
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
ur api /api/v1/things/alarm/info/update \
  --body '{"desc": "string", "evalPeriod": 1, "id": "string", "if": {"triggers": [{"criteria": [{"duration": "...", "frequency": "...", "id": "...", "kOfN": "...", "order": "...", "plain": "...", "rollingAggregate": "...", "type": "..."}], "id": "string", "level": "string", "templateID": "string", "templateName": "示例名称", "type": "string"}]}, "ladders": [{"callbacks": [{"authToken": "string", "authType": "string", "bodyTemplate": "string", "headers": {}, "method": "string", "name": "示例名称", "notifyCycle": "string", "notifyEndTime": "2026-01-01T00:00:00Z", "notifyStartTime": "2026-01-01T00:00:00Z", "retryTimes": 1, "timeoutSeconds": 1, "url": "string"}], "channelTemplates": {}, "channels": ["string"], "delaySeconds": 1, "levels": ["string"], "order": 1, "targets": {"groupIDs": ["string"], "roleIDs": ["string"], "userIDs": ["string"]}, "timing": "string"}], "name": "示例名称", "target": {"areas": [1], "categoryID": "string", "devices": [{"alias": "示例名称", "name": "string", "productID": "string"}], "groups": [1], "productID": "string", "selectType": "string"}}'
```
