# HTML 设备面板开发

用于开发、校验和导入单设备 HTML 面板。面板管理展示、模拟、导入、更新、导出和删除；页面代码在本地编辑器或 AI 工具中维护。多设备大屏继续使用 ur-view。

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

工程包含 meta.json、index.html、panel-sdk.js 和包内 CSS/JS/图片/字体。仅单 HTML 时把 manifest 放入 `<script id="ur-panel-manifest" type="application/json">`，SDK 脚本内联。不能有外部 CDN、远程脚本、iframe、base、对象嵌入、符号链接、路径穿越或缺失资源。限制：压缩 10 MiB、展开 30 MiB、100 文件。导出归一化标准 ZIP，回导以内容复用资产。

## 视觉与交互样板（开发前必读）

新面板从 [开关样板](../assets/device-panel) 或 [空调样板](../assets/device-panel-air-conditioner) 改起，不重新凭空设计。两套样板共用 panel.css 和 panel.js，外层模拟工作台不属于面板资源。后台管理仍沿用所在仓库的既有设计体系，面板内部是可离线运行的独立 HTML，不引入后台 UI 库或 CDN。

| 约束 | 统一要求 |
| --- | --- |
| 颜色 | 从 panel.css 的 --background/card/foreground/muted-foreground/border/primary 等语义变量取值；默认浅色、青绿主色。业务控件不写裸色值。换品牌仅改变量，不逐控件另配色 |
| 字体 | 系统无衬线字体；标题 20–24px，正文/控件 12–14px，辅助 11–12px；温度可用 72–80px；数值等宽，不拉伸、不靠 emoji 装饰 |
| 间距 | 以 8px 为基本节奏，卡片间距 14–20px，内边距 20–28px；统一细边框、轻阴影和圆角，不混用玻璃、霓虹、厚阴影 |
| 结构 | 设备身份和连接状态 → 核心状态/控制 → 分组操作 → 固定反馈区 → 数据来源；桌面双栏、760px以下单栏，温控卡片使用浅青绿底色突出核心数值，其余控制卡保持白底；反馈紧随控制区，320px 不横向溢出 |
| 控件 | 触控主按钮至少 44px；模式使用同组等宽按钮，风速使用分段按钮；正常/选中/禁用/发送中/失败/未知均有明确表现；支持键盘焦点、aria-label、aria-pressed 和实时反馈 |
| 状态 | 不把关闭值、默认温度或缓存显示成实时状态；只读/离线/可选能力缺失禁用相应操作；accepted 与设备确认分开，等待期间锁定，失败/超时不自动重发 |
| 用户文案 | 不显示 SDK、原始 JSON、命令 ID、packageHash 或连接栈；技术详情放在宿主调试区域。模拟与缓存必须明确标识，不能暗示真实设备 |

可调整设备名称、图形、业务字段、模式文案、实际范围和控件数量；保留共享视觉变量、布局层级和状态语义。新增交互优先复用样板 class；用户明确指定品牌风格时可调整统一变量，但不要混进另一套组件外观。

空调样板使用 power、targetTemperature、mode、fanSpeed，及可选 currentTemperature；具体类型、单位、枚举见其 meta.json。它们仅为演示合同，实际产品可能叫 PowerSwitch/SetTemperature 等，必须先查物模型再一致修改 HTML 控件映射、JS 和 requirements，不能为了套模板修改真实设备协议。

交付前至少用 serve 查看桌面及手机，检查320px宽度、开关、温度边界、模式/风速选中、等待回执及失败/离线状态；保留截图。不能仅验证 HTML 能打开就称为视觉验收通过。模拟工作台提供桌面/手机宽度切换和重新运行，不提供在线编辑器。

后台面板管理提供“接入说明”和开关/空调“下载模板”，列表显示首次接入流程，导入弹窗也可直接下载两套模板。完整入门流程、文件职责、SDK 参数与发布验收见 [面板开发说明](device-panel-development.md)。模板下载包附同内容的 development.html，解压后可离线阅读，也可直接导入模拟。后台只有一个导入入口：填写名称和备注，拖拽或点击上传；平台导入自动公共，其他企业自动当前企业，不提供归属选择。导入成功后可选择绑定设备、产品或品类，也可稍后绑定。

## 清单与物模型合同

模板使用 bool 可写属性 power；对接其他产品时先用物模型查询取得真实 identifier、类型、单位、枚举语义和读写权限，再调整 requirements，不能从界面文字臆造字段。清单只描述能力合同，不定义组件树。平台 enum 的 mapping 键必须为数字字符串；空调样板 mode=0制冷/1制热/2除湿/3自动，fanSpeed=0自动/1低风/2中风/3高风。数字类型 min/max/step 按平台合同使用数字字符串（例如 "16"/"30"/"1"），建模后回读确认范围未被默认值替换。不能把 cool/low 等文字作为枚举键；界面按数字键下发并兼容数字/数字字符串的回显。

- formatVersion 固定 1；entry 固定 index.html；sdkVersion 当前 1.0.0。
- platforms 为 web/h5/appAndroid/appIos/mini 的非空子集，声明不代表当前宿主已实现该能力。
- requirements 中 kind/identifier/required 必须明确。属性 access 为 read/write；define 至少包含 type，可约束 unit、mapping/values。write 需要真实模型 rw 且用户有控制权。
- action 只允许 dir=down；input 是以参数 identifier 为键、define 为值的映射。需要按实际类型编码，不要伪造行为。
- required=true 的能力缺失或不兼容会拒绝直接绑定；required=false 的差异以 unavailableOptional 返回，面板应隐藏或禁用对应操作。当前 SDK 支持 bool/int/float/string/enum/timestamp；struct/array 和物模型事件没有通用实现。
- 同一品类/产品可复用满足合同的面板，不要求整个物模型一致；单位、读写和枚举语义变化仍须校验。

## window.urPanel

SDK 通过本机宿主桥工作；运行期可通过 getAuth 获取当前认证，通过 request 调现有平台接口。禁止读取父窗口 DOM、拼设备密钥或把凭据写入资源包。

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

Web 通过校验 iframe 身份后的 MessageChannel 运行，sandbox 只允许脚本；面板可显式调用 getAuth 读取运行期认证，不能读取宿主DOM或持久化凭据。App 缓存包和属性快照，离线仅供查看；控制需要联网。本轮暂不提供蓝牙控制，capabilities.ble.available=false。Android/iOS 包安装与离线冷启动需各自真机验证。小程序能力不足时保留原生面板，不绕开业务域名和通信约束。

## 平台接口与运行期认证

新控制面板按正式物模型和现有接口开发，不迁移旧小程序路由、字段别名或旧UI行为。定时、日志、分享、支付等属于设备详情通用能力，不放入控制面板。

- `await urPanel.getAuth()` 返回 `{baseURL,headers}`。`headers.token` 为当前登录token，同时包含宿主对应的应用、企业和项目请求头；不使用refresh token。每次调用读取实时会话，账号或企业改变后原面板失效。不得写入源码、ZIP、日志、URL或离线缓存。
- `await urPanel.request({path,data})` 调用当前平台已有 **POST** 接口，path为 `/api/v1/` 开头的纯相对路径；不传完整URL、query或自定义认证头。data为对象，结构遵循该接口，不受SDK标量属性声明格式约束。宿主自动携带当前认证上下文，后端执行原有鉴权；它不是设备限定令牌，接口权限与当前登录用户相同。
- 返回值是平台成功响应中的 `data`，不是 `{code,msg,data}` 外层包装；错误通过Promise拒绝，保留业务code/message。接口调用不自动跟踪设备回执；面板拿到msgToken后自行调用回执查询接口，或使用有现成跟踪的writeProperties/invokeAction。受理不等于执行成功。
- 请求经本机桥调用宿主请求客户端，因此Web/App不依赖iframe跨域fetch。包的CSP仍禁止直接联网；需要凭据时通过getAuth取得，平台调用优先使用request。不开放父窗口DOM或原生插件。
- 消息请求超时10秒，接口请求超时8秒；设备命令应使用异步接口参数，之后另查回执。超时不自动重发。请求消息上限64KiB，App原生请求消息上限32KiB；大查询应分页。
- 模拟getAuth返回空headers，不提供真实token；request仅执行宿主createMockAdapter的apiHandlers本地替身，未配置时返回UNSUPPORTED_CAPABILITY，禁止转向真实平台。默认下载样板仍使用属性语义方法；开发接口型模拟面板时须配对应本地替身。
- App离线不提供getAuth/request能力；账号token不会进入包或设备快照。实际App桥行为仍需Android/iOS真机验收。

```js
const context = await urPanel.getContext();
// 需要交给自有接口客户端时按调用读取；不要持久化或输出headers。
const auth = await urPanel.getAuth();
const { msgToken } = await urPanel.request({
  path: '/api/v1/things/device/interact/property-control-send',
  data: {
    productID: context.productID,
    deviceName: context.deviceName,
    data: JSON.stringify({ power: true }), // 改为正式模型的字段和值
    isAsync: true,
    shadowControl: 1, // 按接口合同选择，不从旧UI猜测
  },
});
// 后续查询 /api/v1/things/device/interact/property-control-get-one，
// data传同一productID/deviceName和msgToken；200设备回执才是成功。
```

复杂索引属性、设备profile与控制策略直接采用对应平台接口的正式参数；不再为每一条接口创建SDK方法，不添加旧面板兼容转换层。requirements仅描述现有语义方法可校验的能力；接口型能力不能伪填为受支持的数组声明，应在开发说明列出实际接口、模型校验与依赖，运行前读取正式模型检查。

## 导入和绑定

以下在线命令复用已配置 CLI 账号及应用上下文；首次导入/绑定须联网。

```sh
ur things panel import ./my-panel.zip
# 可选：覆盖导入清单中的名称及备注
ur things panel import ./my-panel.zip --name "客厅面板" --desc "空调设备控制"
ur things panel list --json '{"page":{"page":1,"size":20}}'
ur things panel info --json '{"panelID":"<导入结果的ID>"}'
ur things panel binding --json '{"targetType":"product","targetID":"<产品标识>"}'
ur things panel bind --json '{"targetType":"product","targetID":"<产品标识>","mode":"panel","panelID":"<资产ID>","expectedRevision":0}'
ur things panel resolve --json '{"productID":"<产品标识>","deviceName":"<设备名>","platform":"web","sdkVersion":"1.0.0"}'
ur things panel export --json '{"panelID":"<资产ID>"}' --out ./exported.zip
```

新导入由服务端按身份决定：平台为公共资产，其他企业为当前企业。ownerScope/--scope 仅兼容旧调用；平台传 current 也归为公共，普通企业显式请求 common 仍拒绝。公共品类仅平台管理员维护。名称覆盖后备注可为空以清除旧值；改写参与资产摘要，导出与列表信息一致。服务端以当前用户上传临时 filePath 导入，不能把任意 URL 或其他账号对象当面板源。导入不自动绑定，导出不包含绑定、账号、设备密钥和模拟状态。

解析顺序：设备 > 产品 > 当前品类 > 最近上级品类 > 系统默认。mode=inherit 才继续；panel/builtin 停止。高优先级资产不兼容时返回 builtin 和原因，不能默默选择低优先级的自定义面板。

更新需要先读取本层 revision。返回 applied=false/CONFIRMATION_REQUIRED 时先展示 impact，用户明确确认后用原请求加 acknowledgeFallback=true/impactToken 再提交；不要自动确认。模型或配置期间变化会使确认失效。删除有直接引用时失败，先逐项解绑；inherit 的移除和 builtin 的主动阻断含义不同。

## 验证

先本地 validate/serve/pack，再在隔离 KVM 中验证实际安装的 CLI、上传导入、导出回导、三层绑定、两台设备及 MQTT 回执。用户要求后台可控制、可获取状态时，必须导入后台资产并绑定产品/设备，从设备详情的“设备面板”交付实际运行入口；不能只给 serve 或管理详情里的模拟入口。测试设备可按用户授权以 MQTT 客户端实现，但下发、回执和属性上报必须经过实际平台链路，并明确设备为软件模拟。模拟运行不能代替真实设备链路。H5 通过后再编译小程序；Android/iOS 离线冷启动、损坏包恢复及重入需要真机证据；蓝牙不属于本轮验收。保持 packageHash、解析来源、实际终端能力与验收记录一致。


## 包元数据与资料目录

新模板以根目录 `meta.json` 为唯一清单。它包含 `formatVersion: 1`、`name`、`desc`、`version`、`entry: "index.html"`、`sdkVersion: "1.0.0"`、`platforms` 和 `requirements`。平台会用导入表单的名称和备注改写同一份清单，导出后与列表保持一致。

`platforms` 必须明确列出支持终端，只接受 `web`、`h5`、`appAndroid`、`appIos`、`mini`，不得重复或为空；声明不替代宿主和真机验证。现有 Web/H5/App 模板声明前四项，不因能加载 HTML 就自动增加小程序。

资料引用字段均为包内相对路径：`cover: ".data/cover.png"`、`thingModel: ".data/thing-model.json"`、`readme: "README.md"`、`skills: ["skills/panel-development.md"]`。除首次导入自动补齐的封面外，声明的文件必须存在；README/skills 只接受 Markdown，物模型只接受有效 JSON。真实模型还须用物模型校验命令及目标产品核对，不因面板包可导入就认为模型匹配。

`.data/` 允许保存截图和配套物模型，README 面向人、skills 面向开发 Agent。CLI init、validate、pack 和网页下载均保留这些资料；不扫描工程其他隐藏目录，不携带 node_modules。它们只作为资料，不自动导入物模型、不执行技能内容，不包含账号、密钥或固定设备身份。

网页首次导入在无凭据、禁止网络的临时模拟框中生成 1200×720 画面，再导出 640×384 PNG 到 `.data/cover.png`；封面超过1 MiB、损坏或超200万像素会拒绝。已有封面回导保留原图和内容摘要。旧包无图显示“暂无封面”，不借用其他资产图片。导出ZIP包含封面及开发资料。ZIP仍限制10 MiB压缩、30 MiB展开、100文件，不新增导入接口。

旧包的 manifest.json 或单HTML内嵌清单仍可读入，但新模板只维护 meta.json；两种根清单同时存在时拒绝导入，避免平台选错终端、版本或能力合同。离线/实际运行的 SDK 合同保持不变。


## 原位更新已有面板

后台面板卡片点击“更新”，填写名称和备注并选择新 HTML/ZIP，点击“确认更新”。更新保留 panelID 及设备、产品、品类的全部绑定；用户在设备面板点击“刷新”后加载新内容。开发仍在本地完成，没有在线编辑。

CLI 先查询并核对 `info --json '{"panelID":"<panel-id>"}'` 返回的 `data.panel.packageHash`，再显式提交原版本摘要：

```sh
ur things panel update ./my-panel --panel-id <panel-id> --expected-hash <原版本完整packageHash>
```

调用 `POST /api/v1/things/panel/package/update`：`panelID` 为保持不变的资产身份；`expectedPackageHash` 为查询时的64字符完整摘要；`filePath` 为当前登录账号上传的临时ZIP路径；可选 `name` 非空时同时替换 `desc`（空值清除备注），省略名称则使用新包名称备注。成功返回 `panelID/packageHash/manifest/reused`，同内容不改写。权限为 `things/panel/info:update`，普通企业不能修改公共面板。

任何直接引用（包括其他企业不可列出的绑定）存在时，更新要求 `requirements`、`platforms` 的声明及 `sdkVersion` 与原版本保持一致；需要改变依赖时另行导入并重新核对绑定。无引用可更换依赖。并发摘要不匹配、非法包或另一个面板已使用相同内容均拒绝，原包和绑定保持不变；不要自动读取新摘要重试覆盖。旧资源授权在五分钟有效期内读取原包快照，新解析返回新内容摘要；已打开页面不会强制中断，须刷新加载新版本。回退使用上个版本完整包和当前摘要再次更新。更新不执行物模型导入。
