# 图片素材工作流

> 本文件是 [SKILL.md](../SKILL.md) 的子文档：素材上传、画布内 URL 引用与管理命令（自 SKILL.md 拆分整理，内容未改动）。

## 图片素材工作流

```bash
# 1. 上传（一条命令串两步：system/common/upload-file 拿 URL → view/asset/upload 登记）
ur view asset upload -f ./logo.png --name "企业LOGO" --tags "logo,brand"

# 2. 拿到返回的 url 后，在画布 JSON 中引用：
#    - Image 组件（Informations/Mores/Image）：option.dataset = "<url>"
#    - 画布背景图：editCanvasConfig.backgroundImage = "<url>"

# 3. 管理素材
ur view asset get-list --type image --page 1 --size 20
ur view asset get-list --name logo
ur view asset delete --id <assetID>
```

上传细节：
- 第一步走 `system/common/upload-file` 的 **multipart** 表单，与前端封面上传同口径：`isPublic=true` + `business=view` + `scene=goView/asset` + `useBy=user`，返回**公开桶永久 URL**（`public-*.cos.*`），可直接在画布 JSON 中长期引用
- `--tags` 传逗号分隔字符串（如 `"logo,brand"`），后端按字符串存储
- 上传时按扩展名自动推断类型：image（jpg/jpeg/png/gif/svg/webp/bmp/ico）、video（mp4/webm/mov）、audio（mp3/wav/ogg），其余归 `other`
