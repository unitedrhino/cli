# 本地 HTML 设备面板

外部编辑器维护 HTML/CSS/JS，使用 `ur things panel validate .`、`serve .`、`pack . --out ../panel.zip` 完成离线开发。默认模板需要 bool 可写属性 power。空调样板使用 `ur things panel init ./air-conditioner --template air-conditioner`。

样式修改从 panel.css 的语义变量开始，保留标题与连接状态、主控制区、分组操作、固定反馈区和页脚。只使用随包资源；不得在单个按钮中另造颜色、字体、圆角或阴影；完整视觉约束见开发技能的 device-panel.md。

panel.js 通过 window.urPanel 读取当前设备和模型，按实际权限禁用控件。accepted 仅表示受理，成功以 command 回执为准。未知值不伪装关闭；离线只展示缓存。meta.json 是能力合同，不是组件树；接入实际产品前核对真实物模型字段。

serve 使用假设备，不连接平台或蓝牙。mock.json 为扩展工具说明，默认模拟器按合同推导初始值，不读取该文件。打包文件放工程目录外。

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
