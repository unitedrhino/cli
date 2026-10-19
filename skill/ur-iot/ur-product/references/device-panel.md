# HTML 设备面板开发

用于开发、校验和导入单设备 HTML 面板。面板管理只展示、模拟、导入、导出和删除；页面代码在本地编辑器或 AI 工具中维护。多设备大屏继续使用 ur-view。

## 从空目录开始

```sh
ur things panel init ./my-panel
# 空调模板与默认开关共用视觉和交互规范
ur things panel init ./air-conditioner --template air-conditioner
ur things panel validate ./my-panel
ur things panel serve ./my-panel
ur things panel pack ./my-panel --out ./my-panel.zip
```

这四条命令离线可用，无需账号、后端或 CDN。serve 只监听本机，界面明确标记模拟设备，不会调用真实设备。外部修改文件后重新运行页面加载最新校验资源。模板、SDK 和 mock 示例随 CLI 可执行文件分发；也可复制本技能 [assets/device-panel](../assets/device-panel) 后使用同样命令。ZIP 输出放工程目录外，现有文件不会被覆盖。

工程包含 manifest.json、index.html、panel-sdk.js 和包内 CSS/JS/图片/字体。仅单 HTML 时把 manifest 放入 `<script id="ur-panel-manifest" type="application/json">`，SDK 脚本内联。不能有外部 CDN、远程脚本、iframe、base、对象嵌入、符号链接、路径穿越或缺失资源。限制：压缩 10 MiB、展开 30 MiB、100 文件。导出归一化标准 ZIP，回导以内容复用资产。

## 视觉与交互样板（开发前必读）

新面板从 [开关样板](../assets/device-panel) 或 [空调样板](../assets/device-panel-air-conditioner) 改起，不重新凭空设计。两套样板共用 panel.css 和 panel.js，外层模拟工作台不属于面板资源。后台管理仍沿用所在仓库的既有设计体系，面板内部是可离线运行的独立 HTML，不引入后台 UI 库或 CDN。

| 约束 | 统一要求 |
| --- | --- |
| 颜色 | 从 panel.css 的 --background/card/foreground/muted-foreground/border/primary 等语义变量取值；默认浅色、青绿主色。业务控件不写裸色值。换品牌仅改变量，不逐控件另配色 |
| 字体 | 系统无衬线字体；标题 20–24px，正文/控件 12–14px，辅助 11–12px；温度可用 72–76px；数值等宽，不拉伸、不靠 emoji 装饰 |
| 间距 | 以 8px 为基本节奏，卡片间距 14–20px，内边距 20–28px；统一细边框、轻阴影和圆角，不混用玻璃、霓虹、厚阴影 |
| 结构 | 设备身份和连接状态 → 核心状态/控制 → 分组操作 → 固定反馈区 → 数据来源；桌面双栏、窄屏单栏，320px 不横向溢出 |
| 控件 | 触控主按钮至少 44px；模式使用同组等宽按钮，风速使用分段按钮；正常/选中/禁用/发送中/失败/未知均有明确表现；支持键盘焦点、aria-label、aria-pressed 和实时反馈 |
| 状态 | 不把关闭值、默认温度或缓存显示成实时状态；只读/离线/可选能力缺失禁用相应操作；accepted 与设备确认分开，等待期间锁定，失败/超时不自动重发 |
| 用户文案 | 不显示 SDK、原始 JSON、命令 ID、packageHash 或连接栈；技术详情放在宿主调试区域。模拟与缓存必须明确标识，不能暗示真实设备 |

可调整设备名称、图形、业务字段、模式文案、实际范围和控件数量；保留共享视觉变量、布局层级和状态语义。新增交互优先复用样板 class；用户明确指定品牌风格时可调整统一变量，但不要混进另一套组件外观。

空调样板使用 power、targetTemperature、mode、fanSpeed，及可选 currentTemperature；具体类型、单位、枚举见其 manifest.json。它们仅为演示合同，实际产品可能叫 PowerSwitch/SetTemperature 等，必须先查物模型再一致修改 HTML 控件映射、JS 和 requirements，不能为了套模板修改真实设备协议。

交付前至少用 serve 查看桌面及手机，检查320px宽度、开关、温度边界、模式/风速选中、等待回执及失败/离线状态；保留截图。不能仅验证 HTML 能打开就称为视觉验收通过。模拟工作台提供桌面/手机宽度切换和重新运行，不提供在线编辑器。

## 清单与物模型合同

模板使用 bool 可写属性 power；对接其他产品时先用物模型查询取得真实 identifier、类型、单位、枚举语义和读写权限，再调整 requirements，不能从界面文字臆造字段。清单只描述能力合同，不定义组件树。平台 enum 的 mapping 键必须为数字字符串；空调样板 mode=0制冷/1制热/2除湿/3自动，fanSpeed=0自动/1低风/2中风/3高风。数字类型 min/max/step 按平台合同使用数字字符串（例如 "16"/"30"/"1"），建模后回读确认范围未被默认值替换。不能把 cool/low 等文字作为枚举键；界面按数字键下发并兼容数字/数字字符串的回显。

- formatVersion 固定 1；entry 固定 index.html；sdkVersion 当前 1.0.0。
- platforms 为 web/h5/appAndroid/appIos/mini 的非空子集，声明不代表当前宿主已实现该能力。
- requirements 中 kind/identifier/required 必须明确。属性 access 为 read/write；define 至少包含 type，可约束 unit、mapping/values。write 需要真实模型 rw 且用户有控制权。
- action 只允许 dir=down；input 是以参数 identifier 为键、define 为值的映射。需要按实际类型编码，不要伪造行为。
- required=true 的能力缺失或不兼容会拒绝直接绑定；required=false 的差异以 unavailableOptional 返回，面板应隐藏或禁用对应操作。当前 SDK 支持 bool/int/float/string/enum/timestamp；struct/array 和物模型事件没有通用实现。
- 同一品类/产品可复用满足合同的面板，不要求整个物模型一致；单位、读写和枚举语义变化仍须校验。

## window.urPanel

SDK 通过本机宿主桥工作。不要使用 fetch 调后台、读取父窗口 DOM、访问主账号 token、拼设备密钥或指定任意 API 路径。

```js
const context = await urPanel.getContext();
const model = await urPanel.getThingModel();
const values = await urPanel.readProperties({ identifiers: ['power'] });
const unsubscribe = await urPanel.subscribe({ topics: ['property', 'connection', 'command'] }, (event) => {
  // property.data 是标识符映射；command.data.commandID 关联先前受理结果。
  console.log(event);
});
const accepted = await urPanel.writeProperties({ values: { power: true } });
// 受理不代表设备执行成功；根据 command 事件展示最终结果。
// 无参数行为也明确传 params: {}。
// await urPanel.invokeAction({ identifier: 'restart', params: {} });
// 不再使用时调用 unsubscribe()；页面卸载会销毁所有请求和订阅。
```

readProperties 返回 `{identifier:{value,timestamp,receivedTime,quality,source}}`。时间单位毫秒；null/UNKNOWN 不能展示成正常实时值。不同宿主能力读 context.capabilities；不可用时展示原因。命令状态为 accepted/succeeded/failed/timeout；Web 宿主在设备成功但属性尚未同步时可附带 stateSynchronized=false，此时显示“设备已确认，状态尚未同步”，不伪造上报值；超时表示执行未知，不得自动换到 BLE 再发一次。错误 code 为 INVALID_PARAMS/FORBIDDEN/UNSUPPORTED_CAPABILITY/TIMEOUT/SESSION_CLOSED。

Web 通过校验 iframe 身份后的 MessageChannel 运行，sandbox 只允许脚本；面板无法读取宿主登录态。App 缓存包和属性快照，离线仅供查看；控制需要联网。本轮暂不提供蓝牙控制，capabilities.ble.available=false。Android/iOS 包安装与离线冷启动需各自真机验证。小程序能力不足时保留原生面板，不绕开业务域名和通信约束。

## 导入和绑定

以下在线命令复用已配置 CLI 账号及应用上下文；首次导入/绑定须联网。

```sh
ur things panel import ./my-panel.zip --scope current
ur things panel list --json '{"page":{"page":1,"size":20}}'
ur things panel info --json '{"panelID":"<导入结果的ID>"}'
ur things panel binding --json '{"targetType":"product","targetID":"<产品标识>"}'
ur things panel bind --json '{"targetType":"product","targetID":"<产品标识>","mode":"panel","panelID":"<资产ID>","expectedRevision":0}'
ur things panel resolve --json '{"productID":"<产品标识>","deviceName":"<设备名>","platform":"web","sdkVersion":"1.0.0"}'
ur things panel export --json '{"panelID":"<资产ID>"}' --out ./exported.zip
```

普通企业使用 scope=current；common 公共资产及公共品类默认只由 platform 上下文的平台管理员维护。服务端以当前用户上传临时 filePath 导入，不能把任意 URL 或其他账号对象当面板源。导入不自动绑定，导出不包含绑定、账号、设备密钥和模拟状态。

解析顺序：设备 > 产品 > 当前品类 > 最近上级品类 > 系统默认。mode=inherit 才继续；panel/builtin 停止。高优先级资产不兼容时返回 builtin 和原因，不能默默选择低优先级的自定义面板。

更新需要先读取本层 revision。返回 applied=false/CONFIRMATION_REQUIRED 时先展示 impact，用户明确确认后用原请求加 acknowledgeFallback=true/impactToken 再提交；不要自动确认。模型或配置期间变化会使确认失效。删除有直接引用时失败，先逐项解绑；inherit 的移除和 builtin 的主动阻断含义不同。

## 验证

先本地 validate/serve/pack，再在隔离 KVM 中验证实际安装的 CLI、上传导入、导出回导、三层绑定、两台设备及 MQTT 回执。用户要求后台可控制、可获取状态时，必须导入后台资产并绑定产品/设备，从设备详情的“设备面板”交付实际运行入口；不能只给 serve 或管理详情里的模拟入口。测试设备可按用户授权以 MQTT 客户端实现，但下发、回执和属性上报必须经过实际平台链路，并明确设备为软件模拟。模拟运行不能代替真实设备链路。H5 通过后再编译小程序；Android/iOS 离线冷启动、损坏包恢复及重入需要真机证据；蓝牙不属于本轮验收。保持 packageHash、解析来源、实际终端能力与验收记录一致。
