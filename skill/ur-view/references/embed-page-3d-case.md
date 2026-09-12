# 3D 场景页接入案例（EmbedPage 内嵌页面，数字孪生）

适用场景：自研或 AI 生成的 three.js 单页 3D 场景（配电站/机房/园区），整包 zip 托管到平台，
嵌入大屏「内嵌页面」组件（chartKey `VEmbedPage`），经 ur-scene postMessage 桥接平台数据做数字孪生。
环境特定取值（租户/项目/屏/资产编号）一律用占位符，落地时按目标环境替换。

> 完整方法论（锚点规划、Provider 分层、趋势图三级数据来源、排障清单）与可复用案例包在
> **saas 主仓**：`docs/大屏/功能说明/内嵌页面组件/3D场景接入开发方案.md`、
> 同目录 `power-station-3d-urscene.zip`（128 台电表真实台账版，换现场只改包内 `js/data.js`）。
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
  -F file=@scene.zip -F screenId=<screenID> -F name=<assetName>
# 覆盖上传（多传 assetId；托管 URL 不变，画布配置无需改动）
curl -X POST '<base>/api/v1/view/asset/upload-zip' \
  -H "token: $TOKEN" -H "tenant-code: <tenantCode>" -H "app-id: 200" \
  -F file=@scene.zip -F screenId=<screenID> -F assetId=<assetID> -F name=<assetName>
```

- 托管地址形态：`/api/v1/view/scene/{screenID}/{assetID}/index.html`，免登录静态代理，html 带 no-cache
- 删除语义：删资产清 `scene/{screenID}/{assetID}/`；删大屏清 `scene/{screenID}/` 整前缀；
  **模板包只读共享，删模板 → 实例化大屏内嵌页 404**

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

## 4. 验证与排障

- `ur view screen screenshot --id <screenID>` 截发布页核对（需先 agent-browser 登录保持会话）
- 数据对账：`ur things device log property -p <productID> -d <deviceName>` 比对面板读数
- **场景包更新后验证必须 `agent-browser close` 重启会话**：index.html 有 no-cache 但 js/css 子资源
  走启发式缓存，reload 可能仍跑旧代码（network 日志中 js 请求不出现即命中缓存）
- shell 有 `http_proxy` 时 Chromium 会继承导致 XHR 全挂（ERR_NETWORK），启动会话带 `no_proxy` 补目标域名；
  curl 直连一律 `--noproxy '*'`
- 场景页地址拼 `?debug=1` 开左上角调试浮层（握手/绑定条数/API 调用过程），现场排障首选
- WS 推送验证：探针窗口必须覆盖至少一个完整上报周期（现场电表可能 30 分钟一轮，120s 短探针会误判）
