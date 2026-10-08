---
name: ur-iot-device
description: Use when designing or explaining IoT AI device tools under ur-api, including MCP tools, frontend tools, device list, thing model lookup, property control, action calls, real-time property reads, and semantic alignment between execution paths.
metadata:
  hermes:
    tags: [iot, device, mcp, tool-design, thing-model]
---


# ur-iot-device

## Overview

聚焦 IoT AI 设备工具本体，统一说明 `MCP` 与 `frontend tools` 在设备工具语义上的一致性要求。

## Tool Set

| 工具 | 作用 |
|------|------|
| `get_device_list` | 获取设备列表 |
| `get_thing_model` | 查看设备物模型 |
| `device_property_control` | 控制设备属性 |
| `device_action_send` | 调用设备行为 |
| `get_device_properties` | 查询实时属性 |

## Core Rules

- 控制前优先查看物模型，确认属性/行为是否存在
- 回答“当前状态”前优先查询实时属性，不要凭记忆回答
- 不向最终用户暴露 `productID`、`deviceName` 等内部标识
- frontend tools 与 MCP 的工具名、参数语义、错误语义应尽量一致
- `device_property_control` 需要保留只读属性过滤语义
- `device_action_send` 需要保留行为不存在、参数错误、设备离线等错误语义

## Current Direction

- Web 前端优先走 frontend tools
- 设备侧、Win AI、MQTT/UDP 短期继续保留 MCP
- 共存期必须保证同一工具在不同链路下语义一致
