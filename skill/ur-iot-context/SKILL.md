---
name: ur-iot-context
description: Use when deciding how IoT AI tools identify the target device in all-device-visible frontend experiences under ur-api, including list state, selected device state, session state, execution state, and compatibility with MCP session-bound context.
metadata:
  hermes:
    tags: [iot, context, mcp, device-selection]
---


# ur-iot-context

## Overview

聚焦“所有设备可见但前端不负责切换/绑定设备”场景下，AI 工具如何理解设备作用目标。

## State Model

- 列表态：前端当前能看到哪些设备
- 会话态：当前 AI 对话 session 中已知的辅助上下文
- 执行态：本次工具调用最终作用于哪台设备

## Rules

- 如果前端支持所有设备可见，前端本身不负责切换当前设备，也不负责绑定设备上下文
- Web 前端可以负责展示设备列表、搜索设备、执行工具，但不承担“当前目标设备管理器”职责
- MCP 兼容期允许继续使用 `_sessionID -> device` 的隐式上下文
- Web 新链路中，如果工具执行需要明确设备，应由后端上下文或工具参数给出，而不是要求前端先切换设备
- 当上下文不充分时，AI 应先查询设备列表、让用户补充设备信息，或调用后端上下文能力，而不是假设前端已经绑定目标设备

## Target Resolution Patterns

### 1. 单设备目标

用户通常会直接指定某个设备，例如：

- “关闭 3 号会议室空调”
- “把客厅主灯打开”

这种情况下，AI 应先解析出明确设备标识，再执行单设备工具。

### 2. 区域 + 类型批量目标

用户也会指定某个区域下的一类设备，例如：

- “把厕所的所有灯关闭”
- “打开二楼所有空调”

这种情况下，AI 不应假设只有一台设备，而应先解析：

- 区域：如厕所、客厅、二楼
- 设备类型：如灯、空调、风扇
- 操作范围：如全部、部分、单台

若后端暂未提供稳定的“区域 + 类型 -> 设备集合”能力，AI 应先查询设备列表，再决定是否批量调用。

## Resolution Priority

1. 显式单设备标识优先
2. 区域 + 类型批量条件其次
3. 会话中最近一次已确认目标再次之
4. `_sessionID` 隐式上下文仅在 MCP 兼容链路中兜底
5. 无法唯一确定时，先查询列表或向用户追问
