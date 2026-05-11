---
name: ur-iot-client
description: Use when comparing Web frontend, Win AI client, MQTT, and UDP client behavior for IoT AI tool execution under ur-api, especially during MCP and frontend-tools coexistence.
metadata:
  hermes:
    tags: [iot, client, mcp, frontend]
---


# ur-iot-client

## Overview

聚焦多客户端差异，避免把所有客户端都当成浏览器前端处理。

## Client Matrix

| 客户端 | 当前主路径 | 说明 |
|--------|-----------|------|
| Web 前端 | frontend tools | 负责消费 `tool_call_execute` 并回传结果 |
| Win AI 客户端 | MCP 优先 | 短期不强制切换 |
| MQTT 设备会话 | MCP 优先 | 设备侧链路不天然具备前端执行器 |
| UDP 语音链路 | MCP 优先 | 依赖现有设备会话上下文 |

## Rules

- 不能假设所有客户端都能消费 `tool_call_execute`
- Web 前端可以做设备列表、搜索、表单交互和工具执行
- Win AI / MQTT / UDP 需要先确认是否具备结果回传闭环，再谈替换 MCP
- 共存阶段必须保证同一工具在不同客户端下的语义尽量一致
