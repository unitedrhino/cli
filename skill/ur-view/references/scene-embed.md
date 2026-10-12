# 3D 嵌入场景组件（SceneEmbed）参考

> 本文件是 [ur-view/SKILL.md](../SKILL.md) 的子文档：自研 3D 场景页的嵌入组件说明（2026-09-09 SceneEmbed KVM 实测口径）。

## 组件定位

自研 3D 场景页经 iframe sandbox + postMessage 数据桥（ur-scene 协议 v1）接入，`option.nodeBindings`/`parsedNodes` 与 GlTFModel 同构，`nodePath` 对应场景页上报的锚点 path。画布 JSON 骨架：

```json
{
  "id": "comp-scene-001",
  "chartConfig": {
    "key": "SceneEmbed", "chartKey": "VSceneEmbed", "conKey": "VCSceneEmbed",
    "title": "3D嵌入场景", "category": "Three", "categoryName": "三维",
    "package": "Decorates", "chartFrame": "common", "image": "sceneEmbed.png"
  },
  "attr": { "x": 100, "y": 60, "w": 1200, "h": 800, "zIndex": 1 },
  "option": {
    "url": "<场景页地址，必填>",
    "parsedNodes": [],
    "selectedNodePath": "",
    "nodeBindings": [
      {
        "nodeID": "", "nodePath": "factory/line1/motor1", "nodeName": "1号电机",
        "productID": "<产品ID>", "productName": "", "deviceName": "<设备名>",
        "dataID": "<属性标识符>", "dataName": "", "dataUnit": "",
        "action": { "type": "color", "colorRules": [] }
      }
    ],
    "nodeDescs": {},
    "panelTreeWidth": 168
  },
  "request": {}
}
```

- `parsedNodes` 由场景页 anchors 上报自动回写（`{uuid:"", name, path, type:"Anchor"}`），手工编辑画布时留空数组即可
- `action` 仅面板占位：值下发（`ur-scene:data`）不依赖动作配置，场景页自行决定视觉应用
- 协议全文 / SDK / 场景页开发约束见 `docs/大屏/功能说明/3D嵌入场景组件/README.md`；API 转发仅 `/api/v1/` 前缀 POST，WS 订阅频道白名单 `prop`/`conn`
