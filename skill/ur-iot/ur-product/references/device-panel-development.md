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
| `manifest.json` | 名称、版本、终端和物模型能力合同 | 按实际物模型声明 requirements，发布修改时更新 version |
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

命令最终状态为 succeeded、failed 或 timeout；超时表示结果未知，不自动重发。设备已确认但 `stateSynchronized: false` 时，应提示状态尚未同步并保留实际上报值，后续属性到达后再更新按钮。面板不得直接请求管理 API、读取父窗口 DOM 或保存账号令牌。

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
