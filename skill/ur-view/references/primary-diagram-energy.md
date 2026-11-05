# 一次图的数据绑定与能耗中心接入

先完成[绘图与样板验收](primary-diagram-case.md)。本指南按实际代码中的区域配置和大屏接口说明接入步骤，所有环境值均用占位符。

## 读取区域和电表

HTTP 请求带当前用户的认证及企业上下文，项目级查询带 `project-id`。大屏管理和图库接口使用 `app-id: 200`，能耗页面运行时使用所属应用的上下文。

```bash
# 获取项目及区域
ur things project info get-list
ur things area info get-list
# 所选区域下的设备，递归范围由实际区域树计算
ur things device info get-list --body '{"areaIDs":["<area-id>"],"page":{"page":1,"size":100}}'
# 查询该产品实际存在的属性标识和单位
ur api /api/v1/things/product/schema/get-list \
  --body '{"productID":"<product-id>","type":1,"page":{"page":1,"size":100}}'
# 核实某台设备的最新读数及上报时间
ur things device log property -p '<product-id>' -d '<device-name>'
```

按页取全设备和物模型。物模型 `affordance` 可能为 JSON 字符串，解析后从 `define.unit` 取单位。区域中存在的每台电表建立独立回路，以区域归属、设备用途和已有接线定义匹配。

维护下面的最小映射表，原始设备响应中的认证密钥无需进入画布或 AI 输入。

| 区域 | 回路用途 | 产品 ID | deviceName | 实际属性 |
| --- | --- | --- | --- | --- |
| `<floor-area>` | 照明 | `<meter-product>` | `<lighting-meter>` | Ia、Ib、Ic、P |
| `<floor-area>` | 插座 | `<meter-product>` | `<socket-meter>` | Ia、Ib、Ic、P |

同一台表在总览和所属明细中可展示同一身份。不同回路分别使用其对应电表，绑定前核对设备、属性和回路用途。

## 在组件中绑定实时数据

TextCommon 的 request 配置示例使用一台电表的 Ia 属性。

```json
{
  "requestDataType": 5,
  "requestIoTDeviceConfig": {
    "queryType": "property",
    "mode": "latest",
    "productID": "<meter-product>",
    "deviceNames": ["<lighting-meter>"],
    "dataIDs": ["Ia"],
    "chartMode": "",
    "aggFunc": "last"
  }
}
```

`deviceNames` 填设备唯一名称，`dataIDs` 填实际 identifier。当前一次图 latest 单值绑定的 `chartMode` 留空，采用项目属性推送。每条曲线、读数和仪表保留自己的 request，未匹配测点采用静态 `--`。

最新属性接口为 POST `/api/v1/things/device/msg/property-latest/get-list`。属性记录中的 `timestamp=0` 与默认零值需要按未上报处理。离线设备可显示真实历史值，并在回路附近注明最近上报时间。设备模拟来源写入案例和模型 notes。

WebSocket 握手携带 token、`app-id`、`tenant-code`，订阅 `project.prop.<project-id>` 与 `project.conn.<project-id>`。首载中文 dataName 与推送 identifier 通过 `_iotMeta.propertyColumns` 对应。

验收先等待所有待测读数成为非空值，再采样两次并观察目标设备消息。区域切换完成后重新观察推送与读数。连续切换后若读数停滞，可重新进入页面检查握手、订阅回应及目标设备消息，并记录异常。发生重连时检查握手、订阅回应和关闭原因。同账号连接上限也可能导致连接被替换，需要结合实际连接数和服务日志定位。

## 保存并发布大屏

创建大屏时将其 `projectID` 绑定到上述 IoT 项目。分别备份当前草稿 `content`、发布快照 `publishedContent` 和区域 profile，然后更新内容。

| 操作 | POST 接口 | 请求要点 |
| --- | --- | --- |
| 创建大屏 | `/api/v1/view/project/create` | `name`、`type: "screen"`、`projectID`、草稿状态 `status: 2` |
| 读取草稿与发布快照 | `/api/v1/view/project/detail/get-one` | `id`，核对草稿时 `forView: false`，核对发布态时 `forView: true` |
| 保存内容 | `/api/v1/view/project/detail/update` | `id`、字符串化的完整画布 `content` |
| 发布 | `/api/v1/view/project/update` | `id`、`status: 1` |

CLI 可使用 `ur view screen pull`、`validate`、`push --publish`。大画布通过文件或 HTTP body 传输。旧 CLI 版本可能尚未识别 ConnectLine，应结合当前组件注册和实际页面核对校验结果。

发布前比较当前快照与本轮备份，发现其他编辑时先合并。保存后重新发布，并回读两份内容。若仅修改画布内容，可直接通过现有接口更新。修改共享渲染代码时，同步构建发布物联网与能耗两个实际宿主应用。

## 绑定到能源配电图

### 页面操作

进入能耗中心的“电力集抄 → 配电图”，选择左侧“区域”，点击目标区域并绑定已创建的大屏。一个一级区域对应一张明细图，顶层区域可绑定总览。页面可能按已有配置默认显示“孪生”，使用“区域”查看按区域配置的图。

### API 操作

大屏 ID 与 IoT 项目 ID 分别填写。区域配置写入区域 profile，使用 code `energy_distribution_screen_v1`。

```json
{
  "areaID": "<area-id>",
  "code": "energy_distribution_screen_v1",
  "params": "{\"schemaVersion\":1,\"areaID\":\"<area-id>\",\"selectedScreenId\":\"<screen-id>\",\"selectedScreenName\":\"楼层电力一次图\",\"updatedAt\":\"<ISO-time>\"}"
}
```

POST `/api/v1/things/area/profile/update`，请求带目标 `project-id`。`params` 为字符串化 JSON，界面解析要求 `schemaVersion: 1` 和有效的 `selectedScreenId`。

回读 POST `/api/v1/things/area/profile/get-list`。

```json
{
  "areaID": "<area-id>",
  "codes": ["energy_distribution_screen_v1"]
}
```

从返回的 profiles 中找到相同 code，解析 params，对比 areaID、selectedScreenId 和大屏名称，再从实际页面切换到该区域核对。

项目级孪生使用项目 profile 的 `energy_twin_screen_v1`，其绑定粒度为整个项目。按区域接入时使用上述区域 profile。联犀 SaaS 前端的页面代码入口位于 `apps/web/apps/org-energy/src/views/energy/collection/electricity-map/`，字段定义见 `types.ts`。

## 上线验收与回滚

逐个区域检查名称、图元、设备身份、读数、缩放和重置。实际区域树重复点击已选节点可能取消选择，自动化应先检查选中状态。

物联网独立预览和能耗内嵌页分别检查。两个宿主编入同一共享包的各自产物，核对实际入口 JS 与服务器文件校验和，并检查图元透明着色、旋转和实时订阅。

保留更新前画布、发布快照和区域 profile。回滚时通过现有接口恢复对应草稿和发布内容，恢复本轮改动的区域配置，回读确认。发布记录保存本轮更新范围、源码提交和备份位置，环境实际值写入环境说明或任务记录。
