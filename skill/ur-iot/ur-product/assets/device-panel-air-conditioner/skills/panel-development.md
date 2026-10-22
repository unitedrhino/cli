# 面板开发 Agent 约束

先读根 README.md、meta.json 和 .data/thing-model.json，再修改 index.html、panel.css、panel.js。本文只约束当前模板的开发，不授予平台操作权限。

- meta.json 是唯一清单，platforms 只列经测试的终端；不得同时新增 manifest.json。保持 entry=index.html、formatVersion=1、sdkVersion=1.0.0。
- 字段、类型、枚举、范围及读写以配套物模型为准；修改模型时同步 requirements 和界面。不要把元数据中的终端声明视为运行能力保证。
- 所有真实数据与控制使用 window.urPanel；禁止写入账号、密码、token、固定设备标识或在页面内保存登录凭据。
- 使用现有 CSS 变量、8px 间距节奏及至少44px触控区；桌面双栏，手机单栏，禁止另造主题。
- 已上报状态决定数值和选中项；发送时锁定控制，accepted 仅表示受理；成功回执不伪造属性更新，超时不自动重试。
- 图片和配套数据放 .data/；平台导入会在隔离模拟中生成 .data/cover.png。README、技能说明及物模型随 ZIP 分发，不在运行时执行或自动导入。
- 依次执行 ur things panel validate .、serve .、pack . --out ../panel.zip。模拟覆盖正常、失败、超时、离线；导入后台后再验真实状态、命令与回执。


通用设备设置使用右上角齿轮和 `await window.urPanel.openSettings()`；先检查 `getContext().capabilities.settings.available`。调用不传设备ID或URL，宿主固定当前设备。Web打开当前设备的基础设置，移动端进入已有设备设置页；模拟仅展示模拟设备信息，不调用真实接口。无支持时隐藏入口，跳转失败留在原面板；返回后继续从宿主读取状态。定时、分享等通用能力仍由设备详情管理，不复制进HTML面板。
