---
name: ur-ota
description: "Use when uploading and registering 联犀 OTA firmware, managing OTA modules, creating targeted or batch upgrade jobs, and checking device upgrade progress. Device-side MQTT download, rollback, and image confirmation use device-firmware."
---

# 联犀 OTA 平台管理

本技能负责平台侧“上传文件 → 创建固件记录 → 创建升级任务 → 查询结果”。设备端 Topic、
下载校验、分区切换、`reportInfo` 确认和 rollback 同时加载 `device-firmware` 的
`references/ota.md`。

## 操作前门禁

1. 运行 `ur check --json`，确认目标环境、企业、应用和 admin 权限。
2. 固定 ProductID、DeviceName、当前版本、目标版本和模块；先查询，不从列表选第一项。
3. 检查设备在线且已实现目标 OTA 协议；创建任务可能立即推送，必须在授权范围内。
4. 固件文件必须是应用分区单文件，先在本地核对大小、MD5、SHA-256 和目标版本。
5. 固件、下载 URL、平台 Token 和设备秘密不写入仓库或普通日志。

## 1. 上传固件文件

```bash
ur system upload-file --file <firmware.bin> --json
```

从真实响应提取平台返回的文件路径，保存为 `<uploaded-file-path>`；不要把本地路径、
猜测 URL 或旧示例中的 `url` 字段提交给固件创建接口。

## 2. 创建固件记录

当前创建合同使用 `filePaths`：

```bash
ur api /api/v1/things/ota/firmware/info/create --body '{
  "productID":"<productID>",
  "name":"<release-name>",
  "version":"<target-version>",
  "signMethod":"MD5",
  "isDiff":1,
  "srcVersion":"",
  "isNeedToVerify":2,
  "moduleCode":"<module-code>",
  "desc":"<release-summary>",
  "filePaths":["<uploaded-file-path>"]
}'
```

- `isDiff=1` 为整包，`2` 为差分包。
- `isNeedToVerify=1` 要求先验证固件，`2` 不要求预验证。
- 整包 `srcVersion` 传空字符串；差分包必须填写与目标版不同的源版本。
- 单文件设备只传一个文件路径；`moduleCode` 仅在产品定义模块时填写。

保存返回的 `<firmwareID>`，随后回读，核对产品、版本、包类型、算法和 `fileList`：

```bash
ur api /api/v1/things/ota/firmware/info/get-one \
  --body '{"id":"<firmwareID>"}'
```

## 3. 创建单设备验收任务

固件第一次实机验收使用静态、定向、主动推送、不需 App 确认的单设备任务：

```bash
ur api /api/v1/things/ota/firmware/job/create --body '{
  "firmwareID":"<firmwareID>",
  "type":2,
  "upgradeType":1,
  "retryInterval":5,
  "retryCount":1,
  "timeoutInMinutes":30,
  "maximumPerMinute":1,
  "isOverwriteMode":2,
  "isNeedPush":1,
  "isNeedConfirm":2,
  "targetSelection":2,
  "targetDeviceNames":["<deviceName>"]
}'
```

字段语义：`type=1` 是验证升级、`2` 是批量升级；`upgradeType=1` 静态、`2` 动态；
`targetSelection=2` 定向；`isNeedPush=1` 主动推送；`isNeedConfirm=2` 无需确认；
`isOverwriteMode=2` 不覆盖旧任务。重试和超时值应按产品风险调整，示例不是生产默认值。
定向任务不能同时传 `srcVersions`；全量或灰度任务才按平台合同填写待升级版本列表。

从响应保存 `<jobID>`。全量、灰度、分组或多设备任务会扩大影响面，只有用户明确授权
并完成单设备验收后才创建。

## 4. 查询任务和设备结果

```bash
ur api /api/v1/things/ota/firmware/job/get-one \
  --body '{"id":"<jobID>"}'
ur api /api/v1/things/ota/firmware/job/get-list \
  --body '{"firmwareID":"<firmwareID>","page":{"page":1,"size":20}}'
```

还要查询设备结果和原始进度日志；端点字段见
[设备升级结果](references/api/ur-ota-things-ota-firmware-device.md)。任务状态 `3` 表示
批次完成，但最终成功仍要求同时满足：设备运行目标版本、`reportInfo` 已更新、设备结果
成功且没有 rollback。

## 模块管理

只有产品确实有主控/模组等多模块时才创建 OTA 模块。先查询现有模块，避免重复 code：

```bash
ur api /api/v1/things/ota/module/info/get-list \
  --body '{"productID":"<productID>","page":{"page":1,"size":20}}'
```

完整字段见 [模块接口](references/api/ur-ota-things-ota-module-info.md)。单镜像产品不为形式
额外创建模块。

## 失败处理和收尾

- 上传成功但创建失败：保留响应和文件路径，不重复上传；修正已确认的字段后重试。
- 任务无下行：核对设备在线、主动推送、源版本、目标设备名和 OTA 订阅。
- 有进度但失败：按设备 `step/desc`、Hub 日志和设备串口定位，不能反复建任务掩盖根因。
- 禁止自动删除历史固件和任务；它们是验收证据。测试资源清理由用户明确授权后执行。
- 交付记录固件 ID、任务 ID、版本、文件摘要、目标设备、最终状态和任何人工确认项。

## 接口参考

- [固件记录接口](references/api/ur-ota-things-ota-firmware-info.md)
- [升级任务接口](references/api/ur-ota-things-ota-firmware-job.md)
- [设备升级结果](references/api/ur-ota-things-ota-firmware-device.md)
- [模块接口](references/api/ur-ota-things-ota-module-info.md)
