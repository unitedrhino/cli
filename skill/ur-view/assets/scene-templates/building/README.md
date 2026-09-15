# 建筑能耗真实接入模板

本目录可整体复制到任意路径使用。保留曲面屋顶、竖向幕墙、裙房和景观，以及分层展开、单击切层、区域拾取、间距调整和功率趋势。外形与 B1—8F 仅为中性空间示例，不代表任何现场层数、用途或供电边界。

所有遥测初始为空；没有模拟入口、随机读数或断网回退。整楼及楼层计量关系默认未配置。真实设备必须由大屏宿主绑定；单独打开只展示模型及未连接状态。

## 启动与打包

在本目录运行：

```bash
python3 -m http.server 8768 --bind 127.0.0.1
python3 package.py
# 需要单文件查看模型时（仍无数据桥）：
python3 package.py --preview .temp/building.html
```

访问本机服务根路径。默认 ZIP 为 `.temp/building.zip`，只打包 `index.html`、CSS、JS、本地依赖及许可证；测试、文档、截图不入运行包。预算 2 MB。离线依赖为 Three.js r128、OrbitControls、ur-scene SDK，不需要线上 CDN 或模型库。

## 部署配置

编辑 `js/settings.js`：

| 配置 | 含义 |
|---|---|
| `name` | 页面名称 |
| `telemetry.identifier` | 经物模型核实的功率属性，默认空，必须填写 |
| `telemetry.factor` | 默认空，核实后填写源值换算为 kW 的倍率；源单位为 W 时设 `0.001`，kW 时为 `1` |
| `telemetry.staleMs` / `refreshMs` | 记录有效期 / 恢复检查间隔，默认75分钟 / 1分钟，按上报周期核验 |
| `incoming` | 已核验进线的本地锚点ID；默认空，不与下级馈电相加 |
| `summableFloors` | 确认无总分表重叠后允许求和的楼层ID；默认空，逐回路展示 |
| `meters` | 添加或覆盖中性本地锚点 `{id,name,kind,telemetry?}`，kind 为 incoming 或 feeder，可覆盖各表属性、倍率与有效期 |
| `groups` | 添加或覆盖范围 `{id,name,meters:[本地ID],category}`，category 为 roof/outdoor/unassigned |
| `floorGroups` | 楼层ID到范围ID数组的映射；同一范围被多层引用即为共用计量 |

例如，现场确认某计量范围跨两个楼层，使用同一个 group ID 分别配置这两个楼层；不得复制读数、按面积分摊或重复累计。多锚点绑定同一物理设备时会按产品与设备去重，且属性/倍率/有效期必须一致，否则显示配置冲突并清空读数。

初始本地锚点包括 `main-a`、`main-b`、`level-0-a` 至 `level-8-b` 及三个额外范围。它们都是模型位置标识，与真实设备名无关。配置引用不存在的范围或锚点时立即报错，避免静默漏算。

## 宿主绑定与数据

1. 将 ZIP 上传为新场景资产，将返回的 `entryUrl` 配入 EmbedPage；组件尺寸与画布一致，位置从 `(0,0)` 开始。
2. 通过宿主锚点绑定界面，把 `meter/<本地ID>` 绑定到真实 `productID + deviceName`。产品与设备配置由宿主传入，页面没有硬编码身份或凭证。
3. 核验物模型属性、源单位与上报时间，再填写进线、区域和允许汇总的楼层配置；重新打包上传并发布。
4. SDK 接收初值和项目属性推送；快照与当前范围历史最多三路并发，全部请求限定绑定设备。重复握手保留状态，换绑清空旧值并拒绝旧在途响应。

历史按选定范围查询并缓存一分钟，在30分钟时间网格上取不晚于该点且未过期的记录；任一成员缺失保留断点。当前零值有效、负值蓝色、缺失和过期灰色。正功率四档按当前有效范围的最大正值划分，图例为 kW 区间，不表示容量负载率或过载告警。

## 验证

```bash
node data.test.cjs
# 使用已经安装 Playwright 的 Node 环境；不要将 node_modules 放进模板包。
NODE_PATH=<已有node_modules> node click.test.cjs
NODE_PATH=<已有node_modules> node verify.cjs --out .temp/verification
```

数据测试覆盖空绑定、单位换算、跨产品同名、重复设备、历史缺口、零/负/过期、乱序、失败恢复和换绑隔离。页面测试覆盖无绑定操作及显式测试桥接；测试夹具仅存在测试文件中，不随 ZIP 分发。

保持像素比上限1.5、单主要阴影光源、共享几何和材质；首帧先提交再启动数据连接。分别记录场景与宿主耗时，软件光栅帧率不能作为硬件加速性能证明。第三方许可见 `LICENSE.txt`。
