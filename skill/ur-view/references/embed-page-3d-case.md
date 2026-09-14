# 3D 场景页接入案例（EmbedPage 内嵌页面，数字孪生）

适用场景：自研或 AI 生成的 three.js 单页 3D 场景（配电站/机房/园区），整包 zip 托管到平台，
嵌入大屏「内嵌页面」组件（chartKey `VEmbedPage`），经 ur-scene postMessage 桥接平台数据做数字孪生。
环境特定取值（租户/项目/屏/资产编号）一律用占位符，落地时按目标环境替换。

> 完整方法论（锚点规划、Provider 分层、趋势图三级数据来源、排障清单）与可复用案例包在
> **saas 主仓**：`docs/大屏/功能说明/内嵌页面组件/3D场景接入开发方案.md`、
> 同目录 `power-station-3d-realistic.zip`及`配电站场景/README.md`（可审阅源码、打包与回归工具）；
> `power-station-3d-urscene.zip`保留为早期接入基线。换现场须核对台账、绑定与字段别名，不只替换模型外观。
> 本文只固化 CLI/命令行实操序列与易漏步骤。

## 前置：场景包 zip 要求

- 必含 `index.html`（根目录或唯一一级子目录），包内引用一律相对路径
- **依赖全部本地化**（three.js 等打进 `libs/`）：场景页跑在 `sandbox="allow-scripts"` iframe 里，
  生产多为内网，公网 CDN 不可用
- `libs/scene-sdk.js` 用 v1.1+（才有 ready 每秒重发自愈握手）
- 上限：≤100MB / 解压 ≤300MB / ≤2000 文件 / 深度 ≤8，扩展名白名单（拒 php/sh/exe 等）
- 打包口径：在场景目录内 `zip -r ../scene.zip index.html css js libs`（只打内容项，不带外层目录）

## 1. 上传 / 覆盖场景包

CLI 暂无专用命令（`ur view asset upload` 是图片素材链路，不适用于场景 zip），用 curl multipart：

```bash
TOKEN=$(ur token --raw)   # 或 export UR_TOKEN=xxx
# 首次上传（新建资产，返回 data.id 即 assetId、data.entryUrl 即托管地址）
curl -X POST '<base>/api/v1/view/asset/upload-zip' \
  -H "token: $TOKEN" -H "tenant-code: <tenantCode>" -H "app-id: 200" \
  -F file=@scene.zip -F "screenId=<screenID>" -F "name=<assetName>"
# 覆盖上传（多传 assetId；托管 URL 不变，画布配置无需改动）
curl -X POST '<base>/api/v1/view/asset/upload-zip' \
  -H "token: $TOKEN" -H "tenant-code: <tenantCode>" -H "app-id: 200" \
  -F file=@scene.zip -F "screenId=<screenID>" -F "assetId=<assetID>" -F "name=<assetName>"
```

- 托管地址形态：`/api/v1/view/scene/{screenID}/{assetID}/index.html`，免登录静态代理，html 带 no-cache
- 删除语义：删资产清 `scene/{screenID}/{assetID}/`；删大屏清 `scene/{screenID}/` 整前缀；
  **模板包只读共享，删模板 → 实例化大屏内嵌页 404**

### 已发布场景的独立更新

需要保留回退点时上传新资产（不传assetId），先验证新entryUrl及运行文件，再替换画布URL。覆盖上传会直接改动旧URL所指内容，不能把旧URL当作可回退的备份。

分别用`/api/v1/view/project/detail/get-one`读取编辑态`{id:"<screenID>"}`与发布态`{id:"<screenID>",forView:true}`并备份；两者可能含不同位置或尚未发布的配置。只更新场景资源时，以发布态仅替换URL后保存并发布，再恢复原编辑态的其他字段（URL同步为新版）。写入前回读比较，发现并发变更停止覆盖；发布后分别核对两态及128条等实际绑定数。`project/update`使用`id`字段。

## 2. 画布接入与锚点批量绑定

用标准本地编辑工作流（`ur view screen pull` → 改 bigscreen.json → `validate` → `push --publish`）：

- 组件元素：`chartConfig.key = "EmbedPage"`（`chartKey="VEmbedPage"`、`conKey="VCEmbedPage"`），
  `option.url` 填托管 entryUrl（也可手填外部自部署 URL）
- 场景页 `reportAnchors` 后 `option.parsedNodes` 自动回写 `[{uuid:'',name,path,type:'Anchor'}]`，
  手工建屏时留空数组即可（首载自动填充）
- 绑定写在 `option.nodeBindings`：每设备一条 `{nodePath: 锚点path, productID, deviceName, dataID: 氛围字段}`；
  128 台这种量级不要手点——按设备清单脚本生成 JSON 片段塞进画布后整体 push
- **绑定只绑氛围字段**（每设备 1 个，如 P）：点亮在线状态/驱动氛围刷新；面板详情由场景页
  点击时 `sdk.callApi('property-latest/get-list')` 按需拉快照，绑定表不膨胀

### 场景数据与空状态

- SDK为全局单例`window.UrSceneSDK`，不要`new UrSceneSDK()`；先注册`onInit`、`onData`等回调再`ready()`，锚点规划/对象就绪后`reportAnchors`。握手只表示可以接收数据，不要求等待反射、五金或历史查询完成。
- 宿主在ready和iframe load均可能发送init。相同绑定（含乱序）应幂等，不清空累计值、趋势及在途快照引用；设备映射确实改变时再按业务重新初始化。
- 缺失指标先核对物模型定义，再查最新快照时间戳与近24小时历史。历史`list:null`不代表物模型不存在；过期快照不冒充当前值，零值是有效数据。
- 历史状态按电表和指标独立记录：加载、成功空结果、单点不足绘图、失败分别呈现。空结果也通知面板结束加载；失败允许重试，成功空结果保留节流，实时新点到达后恢复绘图。

## 3. 封面补录（必做，否则列表无缩略图）

全屏 EmbedPage 大屏的编辑器自动封面（html2canvas）**截不到 iframe 内场景内容**，程序化建屏更不经
编辑器保存，必须手动补：

```bash
# 1. agent-browser 打开发布页，等场景数据上屏后截图（对 iframe 元素截），缩放转 JPEG
# 2. 上传封面（注意 scene 是 goView/projectIndexImage，与素材的 goView/asset 不同）
curl -X POST '<base>/api/v1/system/common/upload-file' \
  -H "token: $TOKEN" -H "tenant-code: <tenantCode>" -H "app-id: 200" \
  -F file=@cover.jpg -F isPublic=true -F business=view \
  -F scene=goView/projectIndexImage -F useBy=user
# 3. 回写大屏 indexImage（取上一步返回的 data.fileUri）
ur api /api/v1/view/project/update --body '{"id":"<screenID>","indexImage":"<fileUri>"}'
```

- 截图前明确核验真实模式、有效绑定数与已出数设备数，并等待顶栏所需的进线快照。不要用文本`includes("128")`判断完成，`0/128`也会误命中；模拟数据封面须明确标注。
- 回写后检查`indexImage`及公开图片内容。接口可能将同源绝对URL规范为相对路径，先按站点基址归一化再比较；同时确认发布状态和场景引用未改变。

## 4. 验证与排障

- `ur view screen screenshot --id <screenID>` 截发布页核对（需先 agent-browser 登录保持会话）
- 数据对账：`ur things device log property -p <productID> -d <deviceName>` 比对面板读数
- **场景包更新后避免旧缓存误判**：覆盖同一资产时使用全新浏览器上下文或禁用缓存，核对js/css响应内容；独立新资产使用新的目录URL。index.html的no-cache不代表子资源也禁用缓存
- shell 有 `http_proxy` 时 Chromium 会继承导致 XHR 全挂（ERR_NETWORK），启动会话带 `no_proxy` 补目标域名；
  curl 直连一律 `--noproxy '*'`
- 场景页地址拼 `?debug=1` 开左上角调试浮层（握手/绑定条数/API 调用过程），现场排障首选
- WS 推送验证：探针窗口必须覆盖至少一个完整上报周期（现场电表可能 30 分钟一轮，120s 短探针会误判）

### 写实效果与性能验收

- 先优化柜体比例、倒角/门缝、材质色彩空间与接触阴影，再增加细节。重复零件共享几何/材质或实例化；小尺寸程序贴图与离线预滤波反射随ZIP提供，主体先显示、细节分批补齐，数据查询独立于模型加载。
- 每次对比保持镜头、视口与像素比一致；新旧版都等完整细节就绪后比较绘制数，不能拿旧版主体阶段对比新版完成阶段。场景程序提交首帧的Performance Mark须结合实际截图确认，不能独自证明屏幕已显示。
- 分别记录宿主导航、iframe导航至主体/细节、真实数据就绪耗时，以及ZIP体积、实际传输量、冷/热缓存、网络条件与GPU。宿主尚未创建iframe的失败单列记录；SwiftShader等软件GPU结果不替代普通电脑硬件加速帧率验收。
- 引擎升级或高斯泼溅不是默认提质步骤：先确认有适用的场景资产及渲染器/拾取兼容性，再测体积和加载成本；没有收益证据时保留已验证引擎。工程尺寸与母线走向以现场资料为准，展示留白不写成实测尺寸。
