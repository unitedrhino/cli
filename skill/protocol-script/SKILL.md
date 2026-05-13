---
name: protocol-script-skill
description: "Use when implementing, reviewing, or debugging Yaegi-based protocol scripts for device uplink/downlink hooks, payload conversion, and联动 logic in the IoT platform."
---

# 协议脚本开发 Skill

协议脚本使用 **yaegi**（Go 解释器）引擎，允许在设备消息上行/下行的前后插入自定义 Go 脚本，实现协议转换、数据补充、联动触发等功能。

## 触发模式与函数签名

### 1. 上行前处理 (UpBefore) — TriggerDir=1, TriggerTimer=1

拦截设备上报消息，可修改或丢弃。返回 `nil` 表示丢弃消息。

```go
import "log"
import "context"
import "deviceMsg"

func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg {
    log.Printf("收到消息: %s", string(req.Payload))
    // 修改 req...
    return req // 返回 nil 则丢弃消息
}
```

### 2. 上行后处理 (UpAfter) — TriggerDir=1, TriggerTimer=2

设备消息处理完成后执行，用于联动、记录等副作用。无返回值。

```go
import "log"
import "context"
import "deviceMsg"

func Handle(ctx context.Context, req *deviceMsg.PublishMsg, resp *deviceMsg.PublishMsg) {
    log.Printf("请求: %s", string(req.Payload))
    log.Printf("响应: %s", string(resp.Payload))
    return
}
```

### 3. 下行前处理 (DownBefore) — TriggerDir=2, TriggerTimer=1

拦截平台下发指令，可修改或丢弃。签名同 UpBefore。

```go
func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg {
    // 修改下行指令...
    return req
}
```

### 4. 下行后处理 (DownAfter) — TriggerDir=2, TriggerTimer=2

指令下发后执行，用于记录、联动等。无返回值。

```go
func Handle(ctx context.Context, req *deviceMsg.PublishMsg) {
    log.Printf("指令已下发: %s", string(req.Payload))
    return
}
```

## PublishMsg 结构体

脚本的核心输入/输出数据结构：

```go
type PublishMsg struct {
    Handle       string // mqtt topic 第一段: thing, ota, config 等
    Type         string // 操作类型: property, event, action
    Payload      []byte // 消息内容（JSON 格式字节数组）
    Timestamp    int64  // 毫秒时间戳
    ProductID    string // 产品ID
    DeviceName   string // 设备名称
    Explain      string // 内部扩展字段
    ProtocolCode string // 协议网关 code
}
```

**重要**: `Payload` 是 `[]byte`，需要 `json.Unmarshal` 解析，修改后需 `json.Marshal` 回写。

## 可用包

| 包名 | 说明 | 常用函数 |
|------|------|---------|
| `log` | 日志（被引擎捕获） | `Print`, `PrintLn`, `Printf` |
| `context` | Go 标准库 | `Background()` |
| `strings` | Go 标准库 | `Contains`, `Split`, `Replace` |
| `json` | JSON 序列化 | `Marshal`, `Unmarshal` |
| `gjson` | 高性能 JSON 读写 | `Get`, `Set`, `Delete`, `SetRaw`, `GetMany`, `Parse` |
| `utils` | 类型转换 | `ToFloat64`, `ToInt64`, `ToString`, `ToBool` |
| `deviceMsg` | 消息类型 | `PublishMsg`, `CommonMsg`, `ThingReq` |
| `dm` | 设备管理 RPC | `ProductGet`, `DeviceGet`, `SchemaGet`, `DevPubMsg` |
| `schema` | 物模型 | `Model`, `ModelSimple` |

**注意**: `dm` 包函数需要真实服务运行，本地测试时不可用。

## 常用代码模式

### 模式 1: 解析和修改 Payload（json 方式）

```go
import "json"

var data map[string]any
json.Unmarshal(req.Payload, &data)
params := data["params"].(map[string]any)

// 修改属性
params["newField"] = 42

// 回写
req.Payload, _ = json.Marshal(data)
return req
```

### 模式 2: 读写 Payload（gjson 方式，无需 Unmarshal）

```go
import "gjson"

payload := string(req.Payload)
temp := gjson.Get(payload, "params.temperature").Float()
payload, _ = gjson.Set(payload, "params.celsius", (temp-32)*5/9)
payload, _ = gjson.Delete(payload, "params.rawData")
req.Payload = []byte(payload)
return req
```

### 模式 3: 类型转换

```go
import "utils"

val := params["temperature"]
f := utils.ToFloat64(val)   // any -> float64
i := utils.ToInt64(val)     // any -> int64
s := utils.ToString(val)    // any -> string
b := utils.ToBool(val)      // any -> bool
```

### 模式 4: 跨设备发送消息（需要真实服务）

```go
import "dm"
import "json"

payload, _ := json.Marshal(map[string]any{
    "method": "controlReply",
    "params": map[string]any{"switch": 1},
})
err := dm.DevPubMsg(ctx, &dm.DevPublish{
    ProductID:  "target-product",
    DeviceName: "target-device",
    Handle:     "thing",
    Type:       "property",
    Payload:    payload,
})
```

### 模式 5: 查询设备最新属性值（需要真实服务）

```go
import "dm"

ret, err := dm.DeviceMsg.PropertyLatestGetList(ctx, &dm.PropertyLatestGetListReq{
    ProductID:  req.ProductID,
    DeviceName: req.DeviceName,
    DataIDs:    []string{"Ep"},
})
if err == nil && len(ret.List) > 0 {
    oldVal := utils.ToFloat64(ret.List[0].Value)
    // ret.List[i].DataID = 属性标识符
    // ret.List[i].Value  = 属性值(string类型，需用utils转换)
}
```

### 模式 6: 累计值差值计算（需要真实服务）

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

### 模式 7: DTU 透传协议转换（packReport + SubDevices）

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

## 约束和注意事项

1. **必须定义 Handle 函数** — 入口固定为 `Handle`，签名必须与触发模式匹配
2. **import 使用短包名** — `import "dm"` 而非完整 Go 路径
3. **Payload 是 []byte** — 需要 `json.Unmarshal` 解析，修改后需 `json.Marshal` 回写
4. **log.Print 用于调试** — 日志会被引擎捕获，在调试接口中返回
5. **不要使用 goroutine** — yaegi 解释器不支持并发
6. **Before 脚本返回 nil 会丢弃消息** — 谨慎使用
7. **脚本 panic 会被自动恢复** — 不会导致服务崩溃，但该消息处理会失败
8. **缓存刷新周期 10 分钟** — 脚本修改后最多等 10 分钟生效
9. **dm 包函数需要真实服务** — 本地测试时仅 `gjson`/`json`/`utils`/`log`/`deviceMsg` 可用

## CLI 工具

### 生成脚本模板

```bash
ur script template up-before    # 上行前处理模板
ur script template up-after     # 上行后处理模板
ur script template down-before  # 下行前处理模板
ur script template down-after   # 下行后处理模板
```

### 校验脚本

```bash
ur script validate /tmp/script.go
```

### 从物模型生成脚本框架

```bash
ur model generate-script /tmp/model.json --mode property --output script.go
```

## 本地测试

写完脚本后，可通过本地测试 SDK 验证。在 `script_test.go` 中添加测试用例：

```go
func TestScript_MyScript(t *testing.T) {
    st := newTestScriptTrans()
    msg := &deviceMsg.PublishMsg{
        Handle:     "thing",
        Type:       "property",
        ProductID:  "my-product",
        DeviceName: "my-device",
        Payload:    []byte(`{"method":"report","params":{"temperature":98.6}}`),
    }

    script := `你的脚本代码`

    result, logs, err := st.PublishMsgRun(context.Background(), msg, script)
    // 验证结果...
}
```

运行测试：

```bash
cd backend/things/service/dmsvr/internal/domain/protocol
go test -v -run TestScript ./...
```

**测试文件位置**: `backend/things/service/dmsvr/internal/domain/protocol/script_test.go`

## 相关文档

- 脚本开发指南: `docs/物联网/协议脚本/脚本开发指南.md`
- API 参考: `docs/物联网/协议脚本/API参考.md`
- 示例脚本: `docs/物联网/协议脚本/示例脚本/`

## 关键源码文件

| 文件 | 用途 |
|------|------|
| `things/service/dmsvr/internal/domain/protocol/script.go` | yaegi 脚本引擎核心 |
| `things/service/dmsvr/internal/domain/protocol/def.go` | 常量定义 |
| `things/service/dmsvr/internal/startup/script.go` | 符号表注册 |
| `things/service/dmsvr/internal/logic/protocolmanage/script/*` | 脚本模板 |
| `things/service/dmsvr/internal/domain/protocol/script_test.go` | 本地测试 SDK |
