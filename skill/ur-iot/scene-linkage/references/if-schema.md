# If 触发器字段详解

> 本文件是 `scene-linkage/SKILL.md` 的子文档，展开 `if` 字段下触发器（Triggers）的完整字段、子结构与示例。主文档见 [SKILL.md](../SKILL.md)。

## TriggerDevice（设备触发）

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| productID | string | 是 | 产品ID |
| selectType | string | 是 | `all`/`fixed`/`area`/`areaWithChildren`/`group` |
| deviceName | string | fixed时必填 | 设备名 |
| areaID | string | area时必填 | 区域ID |
| groupID | string | group时必填 | 分组ID |
| mode | string | 否 | `edge`(边缘触发,默认) / `level`(水平触发) |
| type | string | 是 | `connected`/`disConnected`/`propertyReport`/`eventReport` |
| stateKeep | StateKeep | 否 | 状态保持配置 |
| dataID | string | property/event时必填 | 属性/事件标识符 |
| termType | string | 否 | 比较类型：`eq`/`not`/`btw`/`gt`/`gte`/`lt`/`lte`/`in`/`all` |
| values | string[] | 否 | 比较值列表 |
| terms | Cmp[] | 否 | 多条件比较（结构体类型用） |

**StateKeep 状态保持**：
```json
{ "type": "duration", "value": 60 }
```
- `type`: `duration`(持续秒数) / `repeat`(重复次数)
- `value`: 持续时间(秒) 或 重复次数

## TriggerTimer（定时触发）

```json
{
  "execType": "at",
  "execAt": 28800,
  "repeatType": "week",
  "execRepeat": "1111100"
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| execType | string | `at`(定点) / `sunRises`(日出) / `sunSet`(日落) / `loop`(循环) |
| execAdd | int | 日出日落前后偏移秒数（-10800 ~ 10800） |
| execAt | int | 执行时间点，从0点起的秒数（如 8:00 = 28800） |
| execLoopStart | int | 循环开始时间（秒） |
| execLoopEnd | int | 循环结束时间（秒） |
| execLoop | int | 循环间隔（秒） |
| repeatType | string | `once`/`week`/`mount`/`allDay`/`customRange` |
| execRepeat | string | 二进制重复规则：周(7位) / 月(31位) |
| execRepeatStart | string | 自定义范围开始日期（2006-01-02） |
| execRepeatEnd | string | 自定义范围结束日期（2006-01-02） |

**execRepeat 二进制规则**：
- 周重复：`"1111100"` = 周一到周五（从高位到低位：周日至周六）
- 月重复：`"1111111111111111111111111111111"` = 每天

## TriggerWeather（天气触发）

```json
{
  "type": "temp",
  "termType": "gt",
  "values": ["35"]
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| type | string | `temp`(温度) / `humidity`(湿度) |
| termType | string | 比较类型，同设备触发 |
| values | string[] | 比较值 |
