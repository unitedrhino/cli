# 本地 HTML 设备面板

外部编辑器维护 HTML/CSS/JS，使用 `ur things panel validate .`、`serve .`、`pack . --out ../panel.zip` 完成离线开发。默认模板需要 bool 可写属性 power。空调样板使用 `ur things panel init ./air-conditioner --template air-conditioner`。

样式修改从 panel.css 的语义变量开始，保留标题与连接状态、主控制区、分组操作、固定反馈区和页脚。只使用随包资源；不得在单个按钮中另造颜色、字体、圆角或阴影；完整视觉约束见开发技能的 device-panel.md。

panel.js 通过 window.urPanel 读取当前设备和模型，按实际权限禁用控件。accepted 仅表示受理，成功以 command 回执为准。未知值不伪装关闭；离线只展示缓存。manifest 是能力合同，不是组件树；接入实际产品前核对真实物模型字段。

serve 使用假设备，不连接平台或蓝牙。mock.json 为扩展工具说明，默认模拟器按合同推导初始值，不读取该文件。打包文件放工程目录外。
