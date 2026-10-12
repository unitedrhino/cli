# 通用物模型批量导入

本文件是 thing-model 技能主文档（SKILL.md）的子文档。

## 概述

通用物模型（CommonSchema）支持批量导入，用于快速创建大量物模型属性定义。导入流程：

1. **生成 JSON**：从 Excel 或其他数据源生成 `schema.json`
2. **打包 ZIP**：创建包含 `schema.json` 的 ZIP 文件
3. **上传 OSS**：通过 `/api/v1/system/common/upload-file` 上传
4. **调用导入**：通过 `/api/v1/things/schema/common/batch-import` 导入

## CommonSchemaInfo JSON 格式

```json
[
  {
    "type": 1,                    // 1=property, 2=event, 3=action
    "identifier": "Ua",           // 英文标识符（大驼峰）
    "name": {"value": "A相电压"},  // StringValue 包装
    "desc": {},                   // StringValue（可为空）
    "required": 2,                // 1=必填, 2=可选
    "isCanSceneLinkage": 1,       // 是否可场景联动
    "funcGroup": 1,               // 1=普通功能, 2=系统功能
    "userPerm": 1,                // 1=只读, 3=可读写
    "affordance": {"value": "{\"define\":{...},\"mode\":\"r\"}"},  // StringValue 包装的 JSON
    "extendConfig": "{}",
    "order": 1,
    "controlMode": 1
  }
]
```

**关键格式要点**：
- `name` 和 `affordance` 必须用 `{"value": "..."}` 包装（StringValue 格式）
- `affordance.value` 是 JSON 字符串，需要 `json.dumps`
- `identifier` 全局唯一，重复会导致导入错误

## 转换脚本

脚本位置：`.claude/skills/thing-model-skill/scripts/convert_gateway_schema.py`

支持的 Excel 模板：
- `docs/应用/能耗管理/网关导入模板/DLT645电表.xlsx` — 35 个电表属性
- `docs/应用/能耗管理/网关导入模板/CJT188水表.xlsx` — 3 个水表属性

```bash
cd /home/yl/code/saas
python3 .claude/skills/thing-model-skill/scripts/convert_gateway_schema.py

# 输出：
# /tmp/schema_import/schema.json — JSON 文件
# /tmp/schema_import/schema.zip   — ZIP 包（批量导入用）
```

## 批量导入脚本

脚本位置：`.claude/skills/thing-model-skill/scripts/batch_import_schema.sh`

```bash
# 自动登录并导入
./.claude/skills/thing-model-skill/scripts/batch_import_schema.sh http://115.190.3.202:7777

# 使用已有 token 导入
./.claude/skills/thing-model-skill/scripts/batch_import_schema.sh http://1.13.180.134:7777 <token>
```

## API 调用示例

**步骤 1：上传 ZIP 到 OSS**

```bash
ur system upload-file --file /tmp/schema_import/schema.zip

# 输出：
# fileUri:  /oss/ithings-temporary/xxx.zip
# filePath: xxx.zip
```

**步骤 2：调用批量导入**

```bash
curl -X POST http://115.190.3.202:7777/api/v1/things/schema/common/batch-import \
  -H "token: <jwt_token>" \
  -H "Content-Type: application/json" \
  -d '{"schemasFile": "/oss/ithings-temporary/xxx.zip"}'

# 响应：
# {"code":0,"data":{"succCount":32,"errCount":6,"ignoreCount":0,"errSchemas":[...]}}
```

**⚠️ 注意**：`schemasFile` 参数使用 `fileUri`（含 `/oss/` 前缀），不是 `filePath`。

## 导入结果解读

| 字段 | 说明 |
|------|------|
| `succCount` | 成功导入数量 |
| `errCount` | 导入错误数量（通常是 identifier 重复） |
| `ignoreCount` | 忽略数量 |
| `errSchemas` | 错误详情列表 |

## 常见错误

1. **identifier 重复**：系统中已存在同名标识符，返回错误
2. **格式错误**：`name`/`affordance` 未用 StringValue 包装
3. **无效文件地址**：使用了 `filePath` 而非 `fileUri`

## 批量导出（参考）

导出 API：`/api/v1/things/schema/common/batch-export`

导出逻辑源码：`backend/things/service/dmsvr/internal/logic/schemamanage/commonSchemaBatchExportLogic.go`
