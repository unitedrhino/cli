# 一次图 + 实时数据绑定案例（多模态核对，配电一次图）

适用场景：配电一次图（DWG）铺满当底图，在回路/柜列位置叠加 TextCommon 实时显示
`Ua/Ia/P/TotalEnergy` 等属性。环境特定取值（租户/项目/屏/产品编号）一律用占位符，
落地时按目标环境替换。

### 取数命令速查（项目/区域/设备/物模型/实时值，均已实测）

前置：profile 配好目标租户（见下"环境与认证"），以下命令直接可用。

```bash
# 项目列表 → 拿 projectID（大屏/设备都挂在项目下）
ur things project info get-list
# 区域树（按用电类型分的区域，含 areaID/areaName/父子关系）
ur things area info get-list
# 设备列表（deviceName/deviceAlias/areaID/productID/isOnline；当前租户全量）
ur things device info get-list
# 按区域/产品过滤设备（body 透传 API 参数）
ur things device info get-list --body '{"areaID":"<areaID>","productID":"<productID>"}'
# 物模型属性（identifier/name/单位在 affordance.define.unit，affordance 是 JSON 字符串需二次解析；type=1 为属性）
ur api /api/v1/things/product/schema/get-list --body '{"productID":"<productID>","page":{"page":1,"size":100}}'
# 单设备最新属性值（核对绑定数据用）
ur things device log property -p <productID> -d <device-name>
# 兜底：任意 API（如跨项目查设备需 project-id 头时）
ur api /api/v1/things/device/info/get-list -H 'project-id: <projectID>' --body '{"page":{"page":1,"size":50}}'
```

字段对应关系（绑定画布要用）：

- `deviceName` → `requestIoTDeviceConfig.deviceNames`；`deviceAlias` 括号内文字 → 别名标签。
- `productID` → `requestIoTDeviceConfig.productID`。
- schema `identifier`（Ua/Ia/P/TotalEnergy…）→ `dataIDs`。
- 设备 `areaID` 只用于"按区域圈设备"，**与图上位置无直接关系**（区域按用电类型、图按柜列）。

### 环境与认证（坑已踩平）

```bash
# 服务器外网不可达时建隧道
ssh -f -N -L <local-port>:127.0.0.1:<port> <user>@<host>
# profile：直接写 ~/.ur/config.json 或 ur setup；baseURL 填隧道地址
# 关键坑：
#  1. 平台管理员对租户项目无权限(100012)，必须用租户管理员账号；
#  2. 带非 platform 的 tenant-code 登录可能报"参数重复"（用户-企业关系唯一约束）：
#     绕过 = 不带 tenant-code 登录拿 token，写入 profile.token，请求头 tenant-code 仍带目标租户；
#  3. ssh 隧道会断，命令 connection refused 时先重连隧道。
ur config list            # 看当前 profile
ur check                  # 验证认证
```

### 多模态核对流程（核心方法论，需 AI 看图判断）

1. **独立图片处理**：DWG 不能直接读，先渲染成 PNG。本地起 file-viewer 临时工程
   （`@file-viewer/renderer-cad`，`cad.canvasOptions.preserveDrawingBuffer=true`，dpr=2），
   合成 WebGL 层 + 文字标注层导出高清 PNG；必要时按区域裁剪局部图再读。
2. **AI 看图识别**：读 PNG 全貌与局部，识别柜列布局、进线/变压器间隔、回路标签文字。
3. **与平台数据交叉确认**（只读 SQL 或 `ur api`）：设备名前缀 ↔ 图纸分组
   （本例 `1Dx`↔1.dwg、`2Dx`↔2.dwg，图内"至LLT1-2#/T1D2#柜顶母排"文字佐证）；
   区域树维度（按用电类型）与图纸维度（按柜列）是否一致。
4. **定叠加坐标**（从零到画布坐标的完整操作）：
   - 看 PNG 全貌定电路区像素边界（左/右/上/下，如 x∈[<x0>,<x1>]、y∈[<y0>,<y1>]），
     换算画布坐标：`x_canvas = x_png/W*1920`、`y_canvas = y_png/H*1080`（Image `fit:'fill'` 铺满时）。
   - 数列：电路区被竖线分成 N 列，N 应等于配电柜组数（本例 6 列 ↔ 1D4-1D9）；列边界均分或看竖线。
   - 列内排序：设备按回路号**数值**序（1D4-2 在 1D4-10 前，别用字典序）自上而下均布，
     `row_h=(y1-y0)/n`，第 i 台 `y=y0+i*row_h`。
   - 别名标签文字取 `device_alias` 括号内（"变配电房/充电桩/备用"），与值标签同位叠加。
   - 图纸无回路编号时（模板图）以上均为"列序≈柜序"示意布局，须在结论里如实标注。
5. **诚实记录核对结论**：本例图纸为**通用模板图**——回路"编号/容量/用途"与柜号栏全空白，
   device_name 无法与图上回路按文本一一映射；叠加位置为"列序≈柜序"的示意绑定。
   正式项目应先在图上填写回路编号，或走图元级绑定路线（见 docs/大屏 TASK-029）。

### CLI 操作全流程

```bash
# 0a. 取数（项目/区域/设备/物模型/实时值，命令见"取数命令速查"）：
#     projectID → 建屏 body 与 project-id 头；deviceName/alias → 绑定与标签；identifier → dataIDs
# 0b. 环境：profile 配 baseURL/appID=200/tenantCode/account/password（ur setup 或直接写 ~/.ur/config.json）
#    服务器外网不可达时用 ssh 隧道：ssh -f -N -L <local-port>:127.0.0.1:<port> <user>@<host>，baseURL 填隧道地址
#    注意：平台管理员对租户项目无权限（100012），须用租户管理员账号
# 1. 上传底图（返回 url 含请求 Host，写画布时改成相对路径 /oss/... 以便任意入口访问）
ur view asset upload -f 1.png --name "一次图-1#配电室"
# 2. 建屏（body 带 projectID）
ur view screen create --body '{"projectName":"一次图案例-1#配电室","projectID":"<projectID>"}'
# 3. 本地编画布（Image 铺满 + TextCommon 叠加，结构见下）→ 校验 → 推送 → 发布
ur view screen validate -f screen.json
ur view screen push -f screen.json --id <screen-id>
ur view screen publish --id <screen-id>        # 预览页 /view/{id} 只认发布态
# 4. 验证：ur view screen describe --id <screen-id> 看绑定；浏览器/截图看实时值
```

### 画布绑定 JSON 要点

- 背景：`Image`（Informations/Mores/Image）`attr {0,0,1920,1080}`、`option.fit:'fill'`、
  `option.dataset:'/oss/...'`（相对路径）、`zIndex:-1`。
- 叠加：`TextCommon`，`option.format {enabled,prefix:'Ua ',suffix:' V',decimals:1}` 带单位，
  `backgroundColor:'#ffffffd9'` 白底压图。
- **全量节点布局**（用户要求所有回路都有节点时）：列序↔柜序（左起 6 列 ↔ 1D4-1D9/2D4-2D9），
  列内设备按回路号数值序自上而下均布（`row_h=(y1-y0)/n`）。每设备 5 个标签：
  首行**别名标签**（静态 TextCommon，`dataset="<设备名> <别名括号内用途>"`，蓝底左对齐，
  别名取 `device_alias` 括号内文字如"变配电房/充电桩/备用"），下面 2×2 值标签
  （`Ua Ia / P TotalEnergy`，fontSize 9，`ah=min(row_h*0.32,16)`、`gh=min((row_h-ah)/2,18)` 防重叠）。
  1# 331 组件 / 2# 301 组件，发布页渲染与 WS 推送均正常（组件多时 playwright 截图需加大 timeout）。
- 数据：`request.requestDataType=5` + `requestIoTDeviceConfig`
  `{queryType:'property', mode:'latest', productID, deviceNames:[设备名], dataIDs:[标识符]}`
  （**chartMode 留空**才走 WS 实时推送 `project.prop.{projectID}`，停轮询）。

### 列表封面回写

列表卡片缩略图读 `indexImage`，新建屏默认为空。回写流程：

```bash
# 1. 截发布页全图当封面（playwright/agent-browser 均可）
# 2. 上传（scene 必须带，返回临时桶签名 URL）
curl -F file=@cover.png -F scene=goView/projectIndexImage .../api/v1/system/common/upload-file
# 3. 取 fileUri 路径并剥掉 "/oss/temporary/" 前缀得到临时桶 key，传给 update：
#    后端会 CopyFromTempBucket 到公共桶并保存公共路径（/oss/public/...），永久有效
ur api /api/v1/view/project/update --body '{"id":<屏ID>,"indexImage":"<临时桶key>","type":"screen"}'
```

注意：传 `/oss/...` 相对路径或带签名的完整 URL 都不对——前者 key 不存在报错，
后者签名过期后封面失效；只有临时桶 key 会触发"拷到公共桶"分支。

### 验证技巧与坑

- `/view/{id}` 预览只认**发布态**；新建屏 `detail/get-one` 无默认 content（pull 报 missing content），画布需完整自建。
- playwright 同页 `goto` 切换 `#/view/27` → `#/view/28` 不重渲染（hash 路由同组件），须新开 page。
- 素材/封面 URL 由后端按请求 Host 生成：经隧道上传会得到隧道地址，画布 dataset 用相对路径 `/oss/...`。

### 案例实录与踩坑

- 属性标识符（配电表产品示例）：`Ua/Ub/Uc/Ia/Ib/Ic/P/PF/TotalEnergy/TotalEnergyChange`。
- 素材 url 由后端按请求 Host 生成：经隧道上传会得到隧道地址的 url，
  写画布必须改相对路径 `/oss/...`，否则用户浏览器打不开。
- 登录带非 platform 的 tenant-code 时后端可能报"参数重复"（用户-企业关系唯一约束，
  首登绑定后复登冲突）：绕过法——不带 tenant-code 登录拿 token 写入 profile.token，
  请求头 tenant-code 仍按目标租户带。
- 新建屏 `detail/get-one` 无默认 content，画布需完整自建（pull 会报 missing content）。
