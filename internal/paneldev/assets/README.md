# 本地 HTML 设备面板

使用外部编辑器修改 HTML/CSS/JS；无需在线编辑器、账号或第三方 CDN。

```sh
ur things panel validate .
ur things panel serve .
ur things panel pack . --out ../panel.zip
ur things panel import ../panel.zip --scope current
```

`manifest.json` 只声明元数据及物模型合同。模板需要 `power` 布尔可写属性；导入后不会自动绑定。所有控制必须使用随包 `panel-sdk.js` 暴露的 `window.urPanel`；不要调用平台 HTTP、读取宿主凭据或依赖父窗口 DOM。readProperties 返回包含 value/timestamp/receivedTime/quality 的属性映射。writeProperties 返回 accepted，订阅 command 获得最终结果。

serve 仅模拟假设备；App 蓝牙、离线缓存和真机能力需要实际宿主验证。打包输出放在工程目录外，避免被当成运行资源。
