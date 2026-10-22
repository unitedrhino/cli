# 智能空调面板

`ur things panel init ./air-conditioner --template air-conditioner` 创建样板；`ur things panel serve ./air-conditioner` 离线模拟运行。外部编辑器修改，使用 validate/pack 导入，不需要在线编辑器。

与默认开关面板共用 panel.css 和 panel.js，保持统一浅色、青绿主色、字体、间距与交互状态。桌面双栏，手机单栏；温度、模式、风速和电源均通过 window.urPanel 下发。关闭后先开机再调整其他设置，控件选中以属性上报为准，完成以命令回执为准。

示例物模型：power(bool/rw)、targetTemperature(int/rw/16–30/步长1/°C)、mode(enum/rw: 0制冷/1制热/2除湿/3自动)、fanSpeed(enum/rw: 0自动/1低风/2中风/3高风)；可选 currentTemperature(float/r/°C)。这些是演示字段，绑定实际产品前必须映射真实 identifier、单位和枚举键，不能直接假定所有空调相同。

模拟初始为关闭、16°C、制冷、自动风；这是模拟器按清单推导的状态，不能当成真实设备。室温为合成模拟值，不展示成实际测温。离线缓存仅查看，当前无蓝牙控制。

## 包结构与接入

`meta.json` 集中定义名称、备注、内容版本、支持终端 `platforms`、SDK 合同及能力 `requirements`。`readme`、`skills`、`thingModel`、`cover` 均为包内相对路径：

```text
meta.json
index.html / panel.css / panel.js / panel-sdk.js
README.md
skills/panel-development.md
.data/thing-model.json
.data/cover.png             # 网页首次导入时自动生成
```

支持终端值为 `web`、`h5`、`appAndroid`、`appIos`、`mini`；本模板只声明已实现的 Web、H5、Android、iOS 宿主，小程序需另验其宿主约束。meta.json必须位于ZIP根目录；不接受仅有manifest.json或HTML内嵌清单的包，也不要同时保留两个根清单。

配套模型使用联犀 JSON 格式。先在目标产品核对或导入 `.data/thing-model.json`，再导入面板，可选绑定设备、产品、品类。导入面板不会自动修改物模型。`.data/cover.png` 是隔离模拟截图，封面不代表实时状态。元数据、README 和技能文档均不能授予权限或指定资源归属。


只可导入ZIP，根目录meta.json为必需文件。字段的必需性、格式版本/内容版本/SDK版本区别、requirements各项含义及路径限制见随包development.html的“声明清单和能力合同”。打包时选择工程内全部文件，不能把外层工程目录一起压进去。导出包与下载模板都保留根meta.json及其引用的资料。


通用设备设置使用右上角齿轮和 `await window.urPanel.openSettings()`；先检查 `getContext().capabilities.settings.available`。调用不传设备ID或URL，宿主固定当前设备。Web打开当前设备的基础设置，移动端进入已有设备设置页；模拟仅展示模拟设备信息，不调用真实接口。无支持时隐藏入口，跳转失败留在原面板；返回后继续从宿主读取状态。定时、分享等通用能力仍由设备详情管理，不复制进HTML面板。
