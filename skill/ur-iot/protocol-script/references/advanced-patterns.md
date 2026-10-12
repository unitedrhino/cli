# 协议脚本进阶模式（累计值差值 / DTU 透传）

本文件是 `protocol-script` 技能主文档 `SKILL.md` 的子文档，收录模式 6、模式 7 两个较长的完整示例；
通用模式（Payload 解析/改写、类型转换、跨设备消息、属性查询）见主文档。

## 模式 6: 累计值差值计算（需要真实服务）

适用于电表、储能设备等需要从累计值计算增量的场景。

```go
import "log"
import "context"
import "deviceMsg"
import "json"
import "dm"
import "utils"

func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg {
    var data map[string]any
    err := json.Unmarshal(req.Payload, &data)
    if err != nil {
        return req
    }
    params, ok := data["params"].(map[string]any)
    if !ok {
        return req
    }

    // 定义差值计算字段映射: 源字段(累计值) -> 目标字段(变化量)
    fields := map[string]string{
        "Power_Discharge": "Power_Discharge_Chan",
        "Power_Charge":    "Power_Charge_Chan",
    }

    // 收集需要查询的 DataID
    var dataIDs []string
    for src := range fields {
        if _, exists := params[src]; exists {
            dataIDs = append(dataIDs, src)
        }
    }
    if len(dataIDs) == 0 {
        return req
    }

    // 查询设备最新属性值
    ret, err := dm.DeviceMsg.PropertyLatestGetList(ctx, &dm.PropertyLatestGetListReq{
        ProductID:  req.ProductID,
        DeviceName: req.DeviceName,
        DataIDs:    dataIDs,
    })
    if err != nil {
        log.Printf("查询历史属性失败: %v", err)
        return req
    }

    // 构建历史值映射
    oldValues := make(map[string]float64)
    for _, item := range ret.List {
        oldValues[item.DataID] = utils.ToFloat64(item.Value)
    }

    // 计算差值并写入
    for src, dst := range fields {
        if currVal, exists := params[src]; exists {
            curr := utils.ToFloat64(currVal)
            if old, has := oldValues[src]; has {
                diff := curr - old
                if diff < 0 {
                    diff = 0 // 累计值不应出现负差值(设备可能重置)
                }
                params[dst] = int64(diff)
            } else {
                params[dst] = 0
            }
        }
    }

    req.Payload, _ = json.Marshal(data)
    return req
}
```

**本地测试时**：由于 `dm` 包不可用，测试脚本中用硬编码的 `oldValues` 替代 `dm.DeviceMsg.PropertyLatestGetList` 查询。参考 `script_test.go` 中的 `TestScript_PowerDischargeDiff` 用例。

## 模式 7: DTU 透传协议转换（packReport + SubDevices）

适用于 DTU 透传设备接入，将自定义 JSON 格式转换为联犀标准物模型 `packReport` 格式。

**入参格式**（DTU 自定义协议）：
```json
{
    "params": {
        "dir": "up",
        "id": "02102925031500024611",
        "r_data": [
            {"name": "365.lux", "value": "792", "err": "0"}
        ]
    }
}
```

**出参格式**（标准 packReport）：
```json
{
    "method": "packReport",
    "msgToken": "02102925031500024611",
    "sys": {"noAsk": true},
    "subDevices": [
        {
            "productID": "01S",
            "deviceName": "365",
            "properties": [{"params": {"lux": "792"}}]
        }
    ]
}
```

**核心逻辑**：
```go
import "log"
import "context"
import "deviceMsg"
import "gjson"
import "json"
import "strings"

func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg {
    if !gjson.GetBytes(req.Payload, "params").Exists() {
        return req
    }

    var out deviceMsg.ThingReq
    out.Sys = &deviceMsg.SysConfig{NoAsk: true}
    out.Method = "packReport"
    out.MsgToken = gjson.GetBytes(req.Payload, "params.id").String()

    rData := gjson.GetBytes(req.Payload, "params.r_data")
    arr := rData.Array()

    // name 格式为 "设备ID.属性ID"，按设备分组
    var property = map[string]map[string]any{}
    for _, v := range arr {
        if v.Get("err").Int() != 0 {
            continue
        }
        name := v.Get("name").String()
        value := v.Get("value").String()
        dev, p, _ := strings.Cut(name, ".")
        if property[dev] == nil {
            property[dev] = make(map[string]any)
        }
        property[dev][p] = value
    }

    for k, v := range property {
        out.SubDevices = append(out.SubDevices, &deviceMsg.ThingSubDevice{
            ProductID:  "01S", // 替换为实际产品ID
            DeviceName: k,
            Properties: []*deviceMsg.TimeParams{{
                Params: v,
            }},
        })
    }

    req.Payload, _ = json.Marshal(out)
    return req
}
```

**关键点**：
- `deviceMsg.ThingReq` — packReport 消息结构体
- `deviceMsg.ThingSubDevice` — 子设备结构（ProductID + DeviceName + Properties）
- `deviceMsg.TimeParams` — 属性参数（Params map[string]any）
- `deviceMsg.SysConfig` — 系统配置（NoAsk: true 表示不需要设备回复）
- `strings.Cut` — 按分隔符拆分设备ID和属性ID
- 完整示例见 `示例脚本/09-有人云DTU上云协议转换.go`
