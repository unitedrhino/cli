# Affordance 说明

本文件是 thing-model 技能主文档（SKILL.md）的子文档。

## 什么是 Affordance

Affordance 是物模型功能定义的**精简序列化格式**，用于场景联动、设备控制等场景快速传递物模型元数据。后端通过 `schema.DoToAffordanceStr()` 从完整的 Property/Event/Action 自动生成。

在场景联动中，`TriggerDevice.schemaAffordance`、`ActionDevice.schemaAffordance`、`TermProperty.schemaAffordance` 等字段均为**后端自动填充的只读字段**，AI 生成场景联动 JSON 时**不需要手动填写**。

## 三种 Affordance 结构

**PropertyAffordance**（属性功能定义）：
```json
{
  "isUseShadow": false,
  "define": { "type": "float", "min": "-20", "max": "55", "step": "0.1", "unit": "℃" },
  "mode": "rw",
  "controlMode": 1,
  "recordMode": 1,
  "isPassword": 2
}
```

**EventAffordance**（事件功能定义）：
```json
{
  "dir": "up",
  "type": "alert",
  "params": [
    { "identifier": "Voltage", "name": "电压", "define": { "type": "float", "min": "0", "max": "300", "step": "0.1", "unit": "V" } }
  ]
}
```

**ActionAffordance**（行为功能定义）：
```json
{
  "dir": "down",
  "input": [
    { "identifier": "Duration", "name": "持续时间", "define": { "type": "int", "min": "0", "max": "3600", "unit": "s" } }
  ],
  "output": [
    { "identifier": "Result", "name": "执行结果", "define": { "type": "bool", "mapping": {"0": "失败", "1": "成功"} } }
  ]
}
```

## 与场景联动的关系

| 场景联动字段 | 对应 Affordance | 说明 |
|-------------|----------------|------|
| `TriggerDevice.schemaAffordance` | PropertyAffordance / EventAffordance | 根据 `dataID` 和 `type` 自动填充 |
| `ActionDevice.schemaAffordance` | PropertyAffordance / ActionAffordance | 根据 `dataID` 和 `type` 自动填充 |
| `TermProperty.schemaAffordance` | PropertyAffordance | 根据 `dataID` 自动填充 |

生成场景联动时，AI 不需要填写 `schemaAffordance`，但需要理解：
1. `dataID` 必须存在于目标产品的物模型中
2. 属性控制的 `value` 必须符合 `PropertyAffordance.define` 的数据类型和范围
3. 事件触发的 `values` 比较必须匹配事件参数的类型
4. 行为调用的 `input` 参数必须符合 `ActionAffordance.input` 的定义
5. 结构体类型的 `value` 必须是合法的 JSON 对象字符串

## CommonSchema 中的 affordance 字段

CommonSchema 批量导入时，`affordance.value` 是 JSON 字符串，其内部结构就是上述 Affordance 格式（**不是**完整的 Property/Event/Action 结构）。

例如 Property 的 affordance：
```json
{"define":{"type":"float","min":"0","max":"500","step":"0.1","unit":"V"},"mode":"r","controlMode":1,"recordMode":1,"isPassword":2}
```

这与本 Skill 中完整 Property 结构不同——affordance 只包含控制相关的核心字段（define、mode、controlMode、recordMode、isPassword、isUseShadow），不包含 `identifier`、`name`、`tag` 等元数据。元数据在 CommonSchema 的同级字段（identifier、name、desc 等）中单独定义。

## 源码参考

| 文件 | 用途 |
|------|------|
| `backend/things/share/domain/schema/affordance.go` | Affordance 结构定义（PropertyAffordance、EventAffordance、ActionAffordance） |
| `backend/things/service/dmsvr/internal/repo/relationDB/assemble.go` | DO→Affordance 转换逻辑（`ToPropertyDo`、`ToEventDo`、`ToActionDo`） |
| `backend/things/share/domain/schema/model.go` | 完整 Property/Event/Action 定义 |
