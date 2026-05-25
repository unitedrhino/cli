# CLI 命令 — 场景联动（scene）

场景联动管理，支持场景的查询/创建/更新/删除/触发，以及场景日志查询和模板生成。

---

## scene info get-list — 查询场景列表

### 命令语法

```
ur scene info get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --name |  | 否 | string | 场景名称过滤 |
| --state |  | 否 | int | 场景状态过滤 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur scene info get-list
```

### 对应API

- `POST /api/v1/things/rule/scene/info/get-list`

---

## scene info get-one — 查询场景详情

### 命令语法

```
ur scene info get-one [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 场景ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur scene info get-one --id 12345
```

### 对应API

- `POST /api/v1/things/rule/scene/info/get-one`

---

## scene info create — 创建场景

### 命令语法

```
ur scene info create [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 场景联动JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur scene info create --body '{"type":"auto","deviceMode":"multi","name":"温度自动调节","status":1,"if":{"triggers":[{"type":"device","order":1,"device":{"productID":"p_smartswitch_001","selectType":"fixed","deviceName":"switch-001","type":"propertyReport","dataID":"Temperature","termType":"gt","values":["30"]}}]},"when":{"conditions":{"type":"and","terms":[]}},"then":{"actions":[{"order":1,"type":"device","device":{"productID":"p_smartswitch_001","selectType":"fixed","deviceName":"switch-001","type":"propertyControl","dataID":"PowerSwitch","value":"false"}}]}}'
```

### 对应API

- `POST /api/v1/things/rule/scene/info/create`

---

## scene info update — 更新场景

### 命令语法

```
ur scene info update [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --body |  | 是 | JSON | 场景联动JSON |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur scene info update --body '{"id":12345,"name":"温度自动调节v2","status":1}'
```

### 对应API

- `POST /api/v1/things/rule/scene/info/update`

---

## scene info delete — 删除场景

### 命令语法

```
ur scene info delete [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 场景ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur scene info delete --id 12345
```

### 对应API

- `POST /api/v1/things/rule/scene/info/delete`

---

## scene info trigger — 手动触发场景

### 命令语法

```
ur scene info trigger [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --id |  | 是 | string | 场景ID |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur scene info trigger --id 12345
```

### 对应API

- `POST /api/v1/things/rule/scene/info/manually-trigger`

---

## scene log get-list — 查询场景日志

### 命令语法

```
ur scene log get-list [选项]
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --scene-id |  | 否 | string | 场景ID过滤 |
| --result |  | 否 | string | 执行结果过滤 |
| --page |  | 否 | int | 页码（默认1） |
| --size |  | 否 | int | 每页条数（默认20） |
| --json | -j | 否 | bool | 输出JSON格式 |

### 使用示例

```bash
ur scene log get-list --scene-id 12345
```

### 对应API

- `POST /api/v1/things/rule/scene/log/get-list`

---

## scene validate — 校验场景联动 JSON

### 命令语法

```
ur scene validate <文件路径>
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| file |  | 是 | string | 场景联动JSON文件路径，`-`表示从标准输入读取 |

### 校验规则

- 校验 type/deviceMode/name 等必填字段
- 校验触发器类型：device/timer/weather
- 校验动作类型：device/delay/notify/alarm/scene
- 校验比较运算符：eq/not/btw/gt/gte/lt/lte/in/all
- 校验设备选择类型：all/fixed/area/areaWithChildren/group
- 自动触发场景必须包含 `if` 字段
- 所有场景必须包含 `then` 字段

### 使用示例

#### 示例1：校验文件
```bash
ur scene validate scene.json
```

#### 示例2：从标准输入校验
```bash
cat scene.json | ur scene validate -
```

---

## scene template auto — 生成自动触发场景模板

### 命令语法

```
ur scene template auto
```

### 使用示例

```bash
ur scene template auto
```

### 模板说明

生成自动触发场景模板，包含：
- `type`: auto
- `deviceMode`: multi
- `if.triggers`: 设备触发器示例
- `when`: 有效/无效时间范围和条件
- `then.actions`: 设备控制动作示例

---

## scene template manual — 生成手动触发场景模板

### 命令语法

```
ur scene template manual
```

### 使用示例

```bash
ur scene template manual
```

### 模板说明

生成手动触发场景模板，包含：
- `type`: manual
- `deviceMode`: multi
- `when`: 有效/无效时间范围和条件
- `then.actions`: 设备控制动作示例

### 注意事项

- 自动触发场景必须包含 `if`（触发器）字段
- 手动触发场景不需要 `if` 字段
- 场景类型：auto（自动）/ manual（手动）
- 设备模式：single（单设备）/ multi（多设备）
- 状态值：1（启用）/ 2（禁用）
