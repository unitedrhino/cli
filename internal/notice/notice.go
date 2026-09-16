// Package notice 管理 CLI 在业务结果之外提供给人类和 AI 的结构化提示。
package notice

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

// Update 描述 CLI 新版本提示。
type Update struct {
	// Current 是当前 CLI 版本。
	Current string `json:"current"`
	// Latest 是发现的最新 CLI 版本。
	Latest string `json:"latest"`
	// Message 是供人阅读的简短说明。
	Message string `json:"message"`
	// Command 是建议执行的统一升级命令。
	Command string `json:"command"`
	// URL 是对应 Release 页面。
	URL string `json:"url,omitempty"`
}

// Skills 描述客户端 Skills 与 CLI 内置版本不一致的提示。
type Skills struct {
	// Current 汇总客户端目标当前版本。
	Current string `json:"current"`
	// Target 是当前 CLI 内置的 Skills 版本。
	Target string `json:"target"`
	// Targets 列出缺失或版本落后的客户端目标。
	Targets []string `json:"targets,omitempty"`
	// Message 是供人阅读的简短说明。
	Message string `json:"message"`
	// Command 是建议执行的统一升级命令。
	Command string `json:"command"`
}

// Payload 是附加到 JSON 顶层的 _notice 内容。
type Payload struct {
	// Update 是可选的 CLI 版本提示。
	Update *Update `json:"update,omitempty"`
	// Skills 是可选的客户端 Skills 版本提示。
	Skills *Skills `json:"skills,omitempty"`
}

// pending 保存当前进程待返回的提示及是否已经输出。
var pending = struct {
	sync.RWMutex
	payload Payload
	emitted bool
}{}

// Reset 清理当前进程的提示状态，主要用于一次 CLI 调用和单元测试之间隔离。
func Reset() {
	pending.Lock()
	defer pending.Unlock()
	pending.payload = Payload{}
	pending.emitted = false
}

// SetUpdate 设置或替换 CLI 更新提示。
func SetUpdate(update Update) {
	pending.Lock()
	defer pending.Unlock()
	pending.payload.Update = &update
}

// SetSkills 设置或替换 Skills 更新提示。
func SetSkills(skills Skills) {
	pending.Lock()
	defer pending.Unlock()
	pending.payload.Skills = &skills
}

// Snapshot 返回当前提示的只读副本。
func Snapshot() Payload {
	pending.RLock()
	defer pending.RUnlock()
	result := pending.payload
	if result.Update != nil {
		update := *result.Update
		result.Update = &update
	}
	if result.Skills != nil {
		skills := *result.Skills
		skills.Targets = append([]string(nil), result.Skills.Targets...)
		result.Skills = &skills
	}
	return result
}

// Empty 判断提示内容是否为空。
func (payload Payload) Empty() bool {
	return payload.Update == nil && payload.Skills == nil
}

// MarshalJSON 把提示合并到 JSON 对象顶层；数组和标量保持原结构，由终端提示兜底。
func MarshalJSON(value any, indent bool) ([]byte, error) {
	payload := Snapshot()
	if payload.Empty() {
		if indent {
			return json.MarshalIndent(value, "", "  ")
		}
		return json.Marshal(value)
	}

	raw, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(raw, &object); err != nil || object == nil {
		if indent {
			return json.MarshalIndent(value, "", "  ")
		}
		return raw, nil
	}
	noticeRaw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	object["_notice"] = noticeRaw
	markEmitted()
	if indent {
		return json.MarshalIndent(object, "", "  ")
	}
	return json.Marshal(object)
}

// WriteHuman 在没有结构化输出提示时向 stderr 写一条简短说明。
func WriteHuman(writer io.Writer) {
	pending.Lock()
	defer pending.Unlock()
	if pending.emitted {
		return
	}
	if update := pending.payload.Update; update != nil {
		fmt.Fprintf(writer, "提示：发现 ur CLI 新版本 %s（当前 %s），运行 %s 升级\n", update.Latest, update.Current, update.Command)
	}
	if skills := pending.payload.Skills; skills != nil {
		fmt.Fprintf(writer, "提示：%d 个 AI Skills 目标需要更新到 %s，运行 %s 同步\n", len(skills.Targets), skills.Target, skills.Command)
	}
	if !pending.payload.Empty() {
		pending.emitted = true
	}
}

// markEmitted 标记提示已经随结构化输出返回，避免 stderr 重复输出。
func markEmitted() {
	pending.Lock()
	defer pending.Unlock()
	pending.emitted = true
}
