# 场景联动实战案例

> 本文件是 `scene-linkage/SKILL.md` 的子文档，收录常用场景联动完整 JSON 案例。温度超限案例见主文档 [SKILL.md](../SKILL.md#实战案例)。

## 案例2：定时开关灯（日出日落模式）

**需求**：每天日出后30分钟自动关闭路灯，日落后30分钟自动开启路灯。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "路灯定时开关",
  "desc": "根据日出日落时间自动控制路灯",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "timer",
        "order": 1,
        "timer": {
          "execType": "sunRises",
          "execAdd": 1800,
          "repeatType": "allDay"
        }
      },
      {
        "type": "timer",
        "order": 2,
        "timer": {
          "execType": "sunSet",
          "execAdd": 1800,
          "repeatType": "allDay"
        }
      }
    ]
  },
  "when": {
    "validRanges": [],
    "invalidRanges": [],
    "conditions": {
      "type": "and",
      "terms": []
    }
  },
  "then": {
    "actions": [
      {
        "order": 1,
        "type": "device",
        "device": {
          "productID": "streetLight01",
          "selectType": "all",
          "type": "propertyControl",
          "dataID": "PowerSwitch",
          "value": "0"
        }
      }
    ]
  }
}
```

**注意**：日出日落模式需要项目配置地理位置（经纬度），否则校验失败。

## 案例3：设备上线通知+延时执行

**需求**：当任意设备上线时，延时10秒后发送钉钉通知给指定账号。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "设备上线通知",
  "desc": "设备上线后延时通知",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "device",
        "order": 1,
        "device": {
          "productID": "smartDevice01",
          "selectType": "all",
          "type": "connected"
        }
      }
    ]
  },
  "when": {
    "validRanges": [],
    "invalidRanges": [],
    "conditions": {
      "type": "and",
      "terms": []
    }
  },
  "then": {
    "actions": [
      {
        "order": 1,
        "type": "delay",
        "delay": 10
      },
      {
        "order": 2,
        "type": "notify",
        "notify": {
          "type": "dingTalk",
          "notifyCode": "ruleScene",
          "userType": "account",
          "accounts": ["admin"],
          "params": {
            "body": "有设备上线"
          }
        }
      }
    ]
  }
}
```

## 案例4：结构体属性条件触发

**需求**：当设备的GPS位置纬度大于30且经度小于120时触发告警。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "GPS围栏告警",
  "desc": "设备超出指定地理围栏时触发",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "device",
        "order": 1,
        "device": {
          "productID": "gpsDevice01",
          "selectType": "fixed",
          "deviceName": "gps_001",
          "type": "propertyReport",
          "dataID": "Location",
          "terms": [
            { "column": "Location.Lat", "termType": "gt", "values": ["30.0"] },
            { "column": "Location.Lng", "termType": "lt", "values": ["120.0"] }
          ]
        }
      }
    ]
  },
  "when": {
    "validRanges": [],
    "invalidRanges": [],
    "conditions": {
      "type": "and",
      "terms": []
    }
  },
  "then": {
    "actions": [
      {
        "order": 1,
        "type": "alarm",
        "alarm": {
          "mode": "trigger"
        }
      }
    ]
  }
}
```

## 案例5：When 条件组合（工作时间 + 湿度条件）

**需求**：在工作日的 9:00-18:00，当湿度超过80%时开启除湿机。

```json
{
  "type": "auto",
  "deviceMode": "multi",
  "name": "高湿度自动除湿",
  "desc": "工作日高湿度自动开启除湿机",
  "status": 1,
  "if": {
    "triggers": [
      {
        "type": "device",
        "order": 1,
        "device": {
          "productID": "humiditySensor01",
          "selectType": "fixed",
          "deviceName": "sensor_001",
          "type": "propertyReport",
          "dataID": "Humidity",
          "termType": "gt",
          "values": ["80"]
        }
      }
    ]
  },
  "when": {
    "validRanges": [
      {
        "order": 1,
        "type": "date",
        "dateRange": {
          "type": "workday"
        }
      },
      {
        "order": 2,
        "type": "time",
        "timeRange": {
          "type": "customRange",
          "startTime": 32400,
          "endTime": 64800
        }
      }
    ],
    "invalidRanges": [],
    "conditions": {
      "type": "and",
      "terms": []
    }
  },
  "then": {
    "actions": [
      {
        "order": 1,
        "type": "device",
        "device": {
          "productID": "dehumidifier01",
          "selectType": "fixed",
          "deviceName": "dehu_001",
          "type": "propertyControl",
          "dataID": "PowerSwitch",
          "value": "1"
        }
      }
    ]
  }
}
```
