# 场景字段与物模型的对应关系

> 本文件是 `scene-linkage/SKILL.md` 的子文档，说明场景联动 `dataID`、比较值、`schemaAffordance` 与物模型的对应约束。主文档见 [SKILL.md](../SKILL.md)。

## 与物模型的关联

场景联动中的 `dataID` 对应物模型的 `identifier`。生成场景联动 JSON 前，需要了解目标产品的物模型定义：

| 场景联动字段 | 物模型对应 |
|-------------|-----------|
| dataID | Property.identifier / Event.identifier / Action.identifier |
| termType + values | 需匹配 Property.Define.type 的数据类型 |
| schemaAffordance | 只读，后端自动填充为物模型 affordance JSON |

**affordance 说明**：`schemaAffordance` 是后端自动填充的只读字段，包含物模型的完整定义（类型、范围、枚举映射等），用于前端展示和校验。其结构对应 `schema.PropertyAffordance` / `schema.EventAffordance` / `schema.ActionAffordance`。

生成场景联动时，AI 不需要填写 `schemaAffordance`，但需要确保：
1. `dataID` 存在于目标产品的物模型中
2. `value` 的数据类型与物模型定义一致
3. 结构体类型的 `value` 是合法的 JSON 对象字符串
