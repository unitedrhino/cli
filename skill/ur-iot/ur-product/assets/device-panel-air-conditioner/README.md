# 智能空调面板

`ur things panel init ./air-conditioner --template air-conditioner` 创建样板；`ur things panel serve ./air-conditioner` 离线模拟运行。外部编辑器修改，使用 validate/pack 导入，不需要在线编辑器。

与默认开关面板共用 panel.css 和 panel.js，保持统一浅色、青绿主色、字体、间距与交互状态。桌面双栏，手机单栏；温度、模式、风速和电源均通过 window.urPanel 下发。关闭后先开机再调整其他设置，控件选中以属性上报为准，完成以命令回执为准。

示例物模型：power(bool/rw)、targetTemperature(int/rw/16–30/步长1/°C)、mode(enum/rw: cool/heat/dry/auto)、fanSpeed(enum/rw: auto/low/medium/high)；可选 currentTemperature(float/r/°C)。这些是演示字段，绑定实际产品前必须映射真实 identifier、单位和枚举键，不能直接假定所有空调相同。

模拟初始为关闭、16°C、自动模式、自动风；这是模拟器按清单推导的状态，不能当成真实设备。室温为合成模拟值，不展示成实际测温。离线缓存仅查看，当前无蓝牙控制。
