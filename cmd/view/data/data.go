// data.go — 大屏组件元数据加载（go:embed components.json）
//
// 本包只负责加载与查询大屏组件 key 表及 IoT 支持矩阵，不依赖 cmd/shared 或
// cmd/view，因此 cmd/shared（validate/describe 实现）与 cmd/view（命令注册）
// 均可安全引用，不会形成循环 import。
// components.json 由前端组件目录盘点生成（来源见 snapshot 的 source 字段），
// 当前为占位版，后续会被完整盘点版本覆盖，schema 保持不变。
package data

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"
)

// componentsJSON 内嵌的组件元数据快照
//
//go:embed components.json
var componentsJSON []byte

// Component 描述一个大屏组件的元数据条目
type Component struct {
	// Key 组件唯一标识，如 BarCommon
	Key string `json:"key"`
	// ChartKey 图表态 key，固定为 "V"+Key
	ChartKey string `json:"chartKey"`
	// ConKey 容器态 key，固定为 "VC"+Key
	ConKey string `json:"conKey"`
	// Title 组件中文名
	Title string `json:"title"`
	// Category 组件分类，如 BAR / LINE / PIE
	Category string `json:"category"`
	// Package 所属包，如 CHARTS / INFORMATIONS / MEDIAS
	Package string `json:"package"`
	// ChartFrame 渲染框架，如 ECHARTS / COMMON
	ChartFrame string `json:"chartFrame"`
	// Image 组件缩略图文件名
	Image string `json:"image"`
	// IoTQueryTypes 该组件支持的 IoT 查询类型列表；空列表表示该组件不支持 IoT 数据绑定
	IoTQueryTypes []string `json:"iotQueryTypes"`
	// SingleValueSelfRender 是否为单值自渲染组件
	SingleValueSelfRender bool `json:"singleValueSelfRender"`
}

// Snapshot 表示 components.json 的完整快照结构
type Snapshot struct {
	// SnapshotDate 快照生成日期
	SnapshotDate string `json:"snapshotDate"`
	// Source 快照来源（前端组件目录路径）
	Source string `json:"source"`
	// Components 组件元数据列表
	Components []Component `json:"components"`
}

var (
	snapshotOnce sync.Once
	snapshot     *Snapshot
	snapshotErr  error
)

// Load 解析并缓存内嵌的组件元数据快照，重复调用返回同一实例
func Load() (*Snapshot, error) {
	snapshotOnce.Do(func() {
		var s Snapshot
		if err := json.Unmarshal(componentsJSON, &s); err != nil {
			snapshotErr = fmt.Errorf("解析内嵌 components.json 失败: %w", err)
			return
		}
		snapshot = &s
	})
	return snapshot, snapshotErr
}

// FindComponent 按组件 key 查找元数据，未找到返回 nil
func FindComponent(key string) *Component {
	s, err := Load()
	if err != nil {
		return nil
	}
	for i := range s.Components {
		if s.Components[i].Key == key {
			return &s.Components[i]
		}
	}
	return nil
}

// SupportsIoTQueryType 判断组件是否支持指定的 IoT 查询类型；
// 组件不存在或其 IoTQueryTypes 为空（不支持 IoT）时均返回 false
func SupportsIoTQueryType(key, queryType string) bool {
	c := FindComponent(key)
	if c == nil {
		return false
	}
	for _, qt := range c.IoTQueryTypes {
		if qt == queryType {
			return true
		}
	}
	return false
}
