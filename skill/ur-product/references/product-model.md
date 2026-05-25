# CLI 命令 — 物模型模板（model）

物模型模板生成、验证和协议脚本生成功能。

---

## model template — 生成物模型模板

### 命令语法

```
ur things model template [模板类型] [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| template-type |  | 否 | string | 模板类型：property/event/action/full（默认full） |
| --json |  | 否 | bool | 输出JSON格式（默认） |
| --yaml |  | 否 | bool | 输出YAML格式 |
| --output |  | 否 | string | 输出文件路径 |

### 模板类型说明

| 类型 | 说明 |
|------|------|
| property | 单属性模板（含 Define/Mode/Affordance 示例） |
| event | 单事件模板（含 Params/Dir/Type） |
| action | 单动作模板（含 Input/Output/Dir） |
| full | 完整 Model 模板（Properties + Events + Actions） |

### 使用示例

#### 示例1：生成完整物模型模板
```bash
ur things model template full
```

#### 示例2：生成单属性模板（YAML格式）
```bash
ur things model template property --yaml
```

#### 示例3：保存到文件
```bash
ur things model template full --output model.json
```

---

## model validate — 校验物模型 JSON

### 命令语法

```
ur things model validate <文件路径>
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| file |  | 是 | string | 物模型JSON文件路径，`-`表示从标准输入读取 |

### 校验规则

- 校验 properties/events/actions 结构
- 校验 identifier/name/define 等必填字段
- 校验数据类型：bool/int/float/string/struct/enum/array/timestamp/matrix
- 校验属性 mode：r/rw
- 校验事件 type：info/alert/fault
- 校验方向 dir：up/down

### 使用示例

#### 示例1：校验文件
```bash
ur things model validate model.json
```

#### 示例2：从标准输入校验
```bash
cat model.json | ur things model validate -
```

---

## model generate-script — 根据物模型生成协议脚本

### 命令语法

```
ur things model generate-script <物模型文件> [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| model-file |  | 是 | string | 物模型JSON文件路径 |
| --mode |  | 否 | string | 脚本模式：up-before/up-after/down-before/down-after（默认up-before） |
| --output |  | 否 | string | 输出文件路径 |

### 脚本模式说明

| 模式 | 说明 | Handle签名 |
|------|------|-----------|
| up-before | 上行前处理 | `func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg` |
| up-after | 上行后处理 | `func Handle(ctx context.Context, req *deviceMsg.PublishMsg, resp *deviceMsg.PublishMsg)` |
| down-before | 下行前处理 | `func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg` |
| down-after | 下行后处理 | `func Handle(ctx context.Context, req *deviceMsg.PublishMsg)` |

### 使用示例

#### 示例1：生成上行前处理脚本
```bash
ur things model generate-script model.json --mode up-before
```

#### 示例2：生成下行前处理脚本并保存
```bash
ur things model generate-script model.json --mode down-before --output script.go
```

### 说明

- 脚本会根据物模型中的 properties/events/actions 自动生成字段注释
- 生成的脚本包含基础的 JSON 解析和字段映射示例
- 返回 nil 表示丢弃消息（适用于 before 类型脚本）
