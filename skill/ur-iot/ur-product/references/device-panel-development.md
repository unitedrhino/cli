# HTML 设备面板开发说明

面板管理页通过“接入说明”查看完整引导，也可在“导入面板”弹窗直接下载模板。接入顺序为下载模板 → 本地开发与模拟 → 导入绑定 → 设备实际验证。从“下载模板”选择开关或空调，解压后在本地编辑器或 AI 工具中开发。模板包含完整源码和本说明的离线版 `development.html`，也可直接导入后台查看效果。面板管理负责导入、模拟、导出、删除和绑定，源码在本地维护。

## 1. 先确认设备物模型

在产品或设备详情查看真实物模型，核对属性标识符、类型、读写权限、单位、范围和枚举含义。不同产品只要满足同一面板的能力要求，就能共用面板。

- 开关模板需要 `power`，类型 bool，可读写。
- 空调模板需要 `power`、`targetTemperature`、`mode`、`fanSpeed`；`currentTemperature` 是可选的室温。
- 样板字段只是起点。实际设备字段不同，应同时修改清单和 JS 中的映射，不能只改按钮文字或修改真实设备协议来凑样板。

## 2. 文件各自负责什么

| 文件 | 用途 | 开发时怎么改 |
| --- | --- | --- |
| `meta.json` | 名称、版本、终端和物模型能力合同 | 按实际物模型声明 requirements，发布修改时更新 version |
| `index.html` | 面板布局与控件 | 维护页面结构和控件的字段映射，保留 SDK 引用 |
| `panel.css` | 共享样式、响应式布局与状态外观 | 优先调整统一变量，保持桌面双栏、手机单栏及状态样式 |
| `panel.js` | 读取状态、响应操作和处理回执 | 通过 window.urPanel 调用宿主，按实际字段调整 |
| `panel-sdk.js` | 与后台、H5 或 App 宿主通信 | 随模板使用，不写入账号令牌或设备密钥 |
| `development.html` | 可离线阅读的开发说明 | 不参与设备控制，可保留在包内 |

单 HTML 导入需要把同一份清单内嵌为 `id="ur-panel-manifest"`、`type="application/json"` 的 script，并内联所需 SDK、样式和业务脚本；有独立资源时使用 ZIP。

模板无需 npm 安装，不依赖 CDN。需要使用框架时，先构建成静态 HTML/CSS/JS，再把构建产物按该目录结构打包；入口必须是根目录 `index.html`。图片、字体和脚本使用包内相对路径。

## 3. 本地运行、校验和打包

安装支持 `things panel` 命令的联犀 CLI 后，在解压目录的上一级运行：

```sh
# 下列 air-conditioner 是解压后的目录名，请按实际目录调整
ur things panel validate ./air-conditioner
ur things panel serve ./air-conditioner
# 外部编辑器修改后，重新运行模拟页面加载新内容
ur things panel pack ./air-conditioner --out ./air-conditioner.zip
```

`validate`、`serve`、`pack` 均可离线运行，无需登录；serve 的控制只改变模拟设备，不会发送到平台。直接双击 `index.html` 没有宿主桥，不能代替模拟运行。输出 ZIP 放在工程目录外，避免把旧包再次装入新包。CLI 不覆盖已有输出文件，重复打包时使用新文件名。

也可以让 CLI 创建工程：

```sh
ur things panel init ./my-switch
ur things panel init ./my-aircon --template air-conditioner
```

CLI 安装与升级方法见 [联犀 CLI 说明](https://gitee.com/unitedrhino/cli)。若当前版本提示未知命令，需要与此后台配套、包含设备面板命令的 CLI 构建包；正式发布版尚未包含该能力时，仅升级到旧发布版不能获得这些命令。未安装 CLI 时，仍可在本地编辑，把文件压缩为 ZIP 后导入后台模拟；压缩包根目录必须直接包含清单和入口，不能多套一层目录。

## 4. 声明清单和能力合同

下面是开关面板的最小清单；空调请从随包的完整清单修改。

```json
{
  "formatVersion": 1,
  "name": "我的开关面板",
  "version": "1.0.0",
  "sdkVersion": "1.0.0",
  "entry": "index.html",
  "platforms": ["web", "h5", "appAndroid", "appIos"],
  "requirements": [
    {"kind": "property", "identifier": "power", "access": "write", "required": true, "define": {"type": "bool"}}
  ]
}
```

清单是物模型合同，不是页面组件配置。`required: true` 的能力缺失或不兼容时不能绑定；可选能力设为 false，并在界面中隐藏或禁用不可用控件。属性写入要求真实模型可读写且当前用户有控制权限。

平台数值范围使用数字字符串，例如 `"min":"16","max":"30","step":"1"`。枚举 mapping 的键使用数字字符串，不能用 cool、low 等英文语义键。空调样板模式为 0制冷、1制热、2除湿、3自动；风速为 0自动、1低风、2中风、3高风。必须以实际物模型为准，不能假设所有空调都相同。

## 5. 通过 SDK 读取状态和控制设备

在 index.html 中先加载 `panel-sdk.js`，再加载业务 JS。SDK 在 iframe 内提供 `window.urPanel`；账号、设备身份和控制权限由宿主限定。

```js
const context = await window.urPanel.getContext();
const model = await window.urPanel.getThingModel();

// 先订阅，再读取初始快照，避免初始化期间漏掉新上报。
const unsubscribe = await window.urPanel.subscribe(
  { topics: ['property', 'connection', 'command'] },
  (event) => {
    if (event.topic === 'property') {
      // event.data 是 identifier -> 属性快照的映射，更新界面。
      console.log('属性更新', event.data);
    }
    if (event.topic === 'command') {
      // 用 event.data.commandID 关联受理结果，展示最终执行状态。
      console.log('执行结果', event.data);
    }
  },
);
const snapshot = await window.urPanel.readProperties({ identifiers: ['power'] });

// 请放在用户点击事件中，并先检查模型、当前权限和联网能力。
const accepted = await window.urPanel.writeProperties({ values: { power: true } });
// 有下行行为时，按实际行为模型调用：
// await window.urPanel.invokeAction({ identifier: 'restart', params: {} });
window.addEventListener('pagehide', () => unsubscribe(), { once: true });
```

| SDK 方法 | 参数与返回 | 界面应如何使用 |
| --- | --- | --- |
| `getContext()` | 返回当前设备上下文、mode、capabilities | 区分模拟、在线、离线，判断可用能力 |
| `getThingModel()` | 返回宿主限定的属性、行为及可用状态 | 按真实范围、单位、枚举和读写权限渲染控件 |
| `readProperties({identifiers})` | 返回 `{identifier:{value,timestamp,receivedTime,quality,source}}`，时间单位毫秒 | 未知或无效质量显示未知，缓存不标成实时数据 |
| `writeProperties({values})` | 返回 commandID、accepted | 只表示受理，先显示等待，不能立刻提示设备执行成功 |
| `invokeAction({identifier,params})` | 返回 commandID、accepted；无参数也传 `{}` | 参数遵循实际行为输入模型，不能虚构行为 |
| `subscribe({topics},callback)` | 返回取消订阅函数 | 监听属性、连接和命令；退出页面释放订阅 |

命令最终状态为 succeeded、failed 或 timeout；超时表示结果未知，不自动重发。设备已确认但 `stateSynchronized: false` 时，应提示状态尚未同步并保留实际上报值，后续属性到达后再更新按钮。平台接口可通过request调用，凭据由getAuth运行期读取；不读取父窗口DOM或保存账号令牌。

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

## 6. 在后台模拟、导入和绑定

1. 在“面板管理 → 导入面板”中填写名称、备注，拖拽或点击选择 HTML / ZIP，再确认导入。平台导入自动归为公共面板，其他企业导入自动归为当前企业，不需要选择归属。名称和备注写入导入清单并参与内容摘要，导出时保持一致。导入成功后进入可选绑定步骤，可以直接绑定设备、产品或品类，也可以点“稍后绑定”；绑定失败不会撤销导入。
2. 打开“查看 / 模拟”，测试控制、状态、失败反馈和 Web / 手机尺寸。这里使用模拟设备，不能作为真实设备控制验收。
3. 在导入成功后的可选步骤中选择目标并绑定，设备需先选所属产品；也可稍后到设备、产品或品类详情点击“绑定面板”。优先级是设备 > 产品 > 品类（含最近上级品类）；全部未绑定时后台显示空状态。
4. 从真实设备详情打开设备面板，核对初始状态、下发回执和后续属性上报。只看到页面或 accepted 都不算控制成功。
5. 修改后更新清单版本，重新打包并导入，再按需更换绑定。资产内容不可在线覆盖；列表中的“导出”导出选中资产，“导出模板”下载开发起点。

## 7. 交付前检查

- 沿用开关或空调样板的样式变量、字体、间距与反馈区，控制区不随意拼接另一套外观；320px 窄屏无横向溢出，主按钮适合触控。
- 覆盖加载、未知状态、关闭、只读、离线、等待回执、成功、失败、超时；模拟数据与缓存有明确标识。
- Web / 手机型号切换只是视口尺寸预览，Android/iOS App 仍需分别真机验证。
- 当前离线缓存用于查看，不支持蓝牙控制；微信小程序能力不足时使用原生设备页，不把 App 能力当成小程序能力。
- 常见导入失败：多套目录、缺清单、外部资源、字段合同不一致、范围/枚举错误；先 validate，再按具体错误修正。


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
