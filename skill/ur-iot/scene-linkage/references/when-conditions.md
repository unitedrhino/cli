# When 生效条件字段详解

> 本文件是 `scene-linkage/SKILL.md` 的子文档，展开 `when` 字段下有效/无效时间段与附加条件的完整字段、子结构与示例。主文档见 [SKILL.md](../SKILL.md)。

**DateRange**：
```json
{ "type": "workday", "startDate": "2026-01-01", "endDate": "2026-12-31", "repeat": "1111100" }
```
- `type`: `allDay`/`workday`/`weekend`/`holiday`/`customRange`/`customWeek`
- `startDate`/`endDate`: 自定义范围时必填（格式 2006-01-02）
- `repeat`: customWeek时必填，7位二进制

**TimeRange**：
```json
{ "type": "customRange", "startTime": 28800, "endTime": 64800 }
```
- `type`: `allDay`/`light`/`night`/`customRange`
- `startTime`/`endTime`: 从0点起的秒数

## Conditions（附加条件）

| 字段 | 类型 | 说明 |
|------|------|------|
| type | string | `and`(全部满足) / `or`(任一满足) |
| terms | Term[] | 条件项列表 |

**Term**：
```json
{
  "order": 1,
  "columnType": "property",
  "property": { ... },
  "weather": { ... },
  "time": { ... }
}
```

- `columnType`: `property` / `event` / `time` / `weather`
- `property`: TermProperty（含 productID、deviceName、Compare）
- `weather`: TermWeather（同 TriggerWeather）
- `time`: TermTime（系统时间或日出日落偏移比较）

**TermTime**：
```json
{ "type": "sys", "termType": "btw", "values": ["28800", "64800"] }
```
- `type`: `sys`(系统时间) / `sunRises`(日出后) / `sunSet`(日落后)
- `termType`/`values`: 同比较逻辑
