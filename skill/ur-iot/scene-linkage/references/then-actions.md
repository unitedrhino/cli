# Then 执行动作字段详解

> 本文件是 `scene-linkage/SKILL.md` 的子文档，展开 `then` 字段下动作（Actions）的完整字段、子结构与示例。主文档见 [SKILL.md](../SKILL.md)。

## ActionDevice（设备动作）

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| productID | string | 是 | 产品ID |
| selectType | string | 是 | `all`/`fixed`/`area`/`areaWithChildren`/`group` |
| deviceName | string | fixed时必填 | 设备名 |
| areaID | string | area时必填 | 区域ID |
| groupID | string | group时必填 | 分组ID |
| type | string | 是 | `propertyControl`(属性控制) / `action`(调用行为) |
| dataID | string | 否 | 属性/行为标识符 |
| value | string | 单值时必填 | 发送的值（结构体类型传 JSON 对象字符串） |
| values | DeviceValue[] | 多值时必填 | 多属性控制列表 |

**DeviceValue**：
```json
{ "dataID": "Brightness", "value": "80" }
```

**结构体类型 value 示例**：
```json
{
  "dataID": "Color",
  "value": "{\"Red\":255,\"Green\":0,\"Blue\":0}"
}
```

## ActionNotify（通知动作）

```json
{
  "type": "message",
  "notifyCode": "ruleScene",
  "userType": "deviceProjectAdmin",
  "accounts": ["admin"],
  "userIDs": ["123"],
  "params": { "body": "告警内容" }
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| type | string | `sms`/`email`/`dingTalk`/`wxMini`/`message`/`phoneCall`/`wxEWebhook` |
| notifyCode | string | `ruleScene`(场景联动通知) / `ruleDeviceAlarm`(设备告警通知) |
| userType | string | `account`/`userID`/`deviceOwner`/`deviceProjectAdmin`/`deviceAreaAdmin`/`deviceProjectAll`/`deviceAreaAll` |
| accounts | string[] | userType=account时必填 |
| userIDs | string[] | userType=userID时必填 |
| params | map[string]string | 模板参数，如 `{"body":"告警内容"}` |

## ActionAlarm（告警动作）

```json
{ "mode": "trigger" }
```
- `mode`: `trigger`(触发告警) / `relieve`(解除告警)

## ActionScene（执行其他场景）

```json
{
  "sceneID": "123",
  "status": 1
}
```
- `sceneID`: 目标场景ID
- `status`: 自动化场景时修改的状态（1=启用, 2=禁用）
