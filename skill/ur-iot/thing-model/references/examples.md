# 实战案例

本文件是 thing-model 技能主文档（SKILL.md）的子文档。

## 案例1：CJT188 水表

Excel 原始定义：
| 功能类型 | 标识符 | 名称 | 读写 | 数据类型 | 小数位 |
|----------|--------|------|------|----------|--------|
| property | PostiveFlux | 当前累积流量 | R | double | 2 |
| property | ValveStatus | 阀门状态 | R | int | 0 |
| property | ValveCtrl | 阀门控制 | W | int | 0 |

转换后的物模型 JSON：

```json
{
  "version": "1.0",
  "properties": [
    {
      "identifier": "PostiveFlux",
      "tag": 1,
      "name": "当前累积流量",
      "desc": "",
      "mode": "r",
      "define": { "type": "float", "min": "0", "max": "99999999", "step": "0.01", "unit": "m³" },
      "funcGroup": 1, "userPerm": 1, "controlMode": 1, "recordMode": 1, "isPassword": 2
    },
    {
      "identifier": "ValveStatus",
      "tag": 1,
      "name": "阀门状态",
      "desc": "",
      "mode": "r",
      "define": { "type": "enum", "mapping": {"0": "开", "1": "关", "2": "异常"} },
      "funcGroup": 1, "userPerm": 1, "controlMode": 1, "recordMode": 1, "isPassword": 2
    },
    {
      "identifier": "ValveCtrl",
      "tag": 1,
      "name": "阀门控制",
      "desc": "",
      "mode": "rw",
      "define": { "type": "enum", "mapping": {"0": "开阀", "1": "关阀"} },
      "funcGroup": 1, "userPerm": 3, "controlMode": 1, "recordMode": 1, "isPassword": 2
    }
  ],
  "events": [],
  "actions": [],
  "profile": { "productID": "" }
}
```

## 案例2：DLT645 三相电表（部分属性）

Excel 原始定义：电压、电流、功率、电能等 double 类型属性，以及跳合闸控制 uint32 属性。

转换后的物模型 JSON（精简示例）：

```json
{
  "version": "1.0",
  "properties": [
    {
      "identifier": "Ua", "tag": 1, "name": "A相电压", "desc": "", "mode": "r",
      "define": { "type": "float", "min": "0", "max": "500", "step": "0.1", "unit": "V" },
      "funcGroup": 1, "userPerm": 1, "controlMode": 1, "recordMode": 1, "isPassword": 2
    },
    {
      "identifier": "Ia", "tag": 1, "name": "A相电流", "desc": "", "mode": "r",
      "define": { "type": "float", "min": "0", "max": "1000", "step": "0.001", "unit": "A" },
      "funcGroup": 1, "userPerm": 1, "controlMode": 1, "recordMode": 1, "isPassword": 2
    },
    {
      "identifier": "EP", "tag": 1, "name": "正向有功总电能", "desc": "", "mode": "r",
      "define": { "type": "float", "min": "0", "max": "99999999", "step": "0.01", "unit": "kWh" },
      "funcGroup": 1, "userPerm": 1, "controlMode": 1, "recordMode": 1, "isPassword": 2
    },
    {
      "identifier": "P", "tag": 1, "name": "总有功功率", "desc": "", "mode": "r",
      "define": { "type": "float", "min": "-9999999", "max": "9999999", "step": "0.001", "unit": "kW" },
      "funcGroup": 1, "userPerm": 1, "controlMode": 1, "recordMode": 1, "isPassword": 2
    },
    {
      "identifier": "Relay", "tag": 1, "name": "跳合闸控制", "desc": "", "mode": "rw",
      "define": { "type": "enum", "mapping": {"0": "合闸", "1": "跳闸"} },
      "funcGroup": 1, "userPerm": 3, "controlMode": 1, "recordMode": 1, "isPassword": 2
    }
  ],
  "events": [
    {
      "identifier": "PowerAlarm", "tag": 1, "name": "电量告警", "desc": "电量异常告警",
      "type": "alert", "dir": "up",
      "params": [
        { "identifier": "AlarmType", "name": "告警类型", "define": { "type": "enum", "mapping": {"1": "过压", "2": "欠压", "3": "过流"} } }
      ]
    }
  ],
  "actions": [
    {
      "identifier": "ReadMeter", "tag": 1, "name": "抄表", "desc": "远程抄表",
      "dir": "down",
      "input": [],
      "output": [
        { "identifier": "EP", "name": "正向有功总电能", "define": { "type": "float", "min": "0", "max": "99999999", "step": "0.01", "unit": "kWh" } }
      ]
    }
  ],
  "profile": { "productID": "" }
}
```
