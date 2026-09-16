// project.go 统一 API 命令的项目上下文选择，避免 ID 精度损失与歧义请求。
package client

import (
	"fmt"
	"strings"
)

// ApplyProjectID 将项目写入请求头：显式参数/请求头优先于环境值；显式值冲突返回错误。
// projectSet 区分未传参数与显式空值，所有项目标识仅作为字符串处理。
func ApplyProjectID(headers map[string]string, projectID string, projectSet bool, envProjectID string) error {
	// selected 保存唯一显式项目值；空参数不可回退到环境值。
	selected := projectID
	if projectSet && (selected == "" || strings.TrimSpace(selected) != selected) {
		return fmt.Errorf("--project-id 必须是非空且不含首尾空白的字符串")
	}
	for key, value := range headers {
		if !strings.EqualFold(key, "project-id") {
			continue
		}
		if value == "" || strings.TrimSpace(value) != value {
			return fmt.Errorf("project-id 请求头不能为空或包含首尾空白")
		}
		if selected != "" && selected != value {
			return fmt.Errorf("项目上下文冲突：--project-id 与 project-id 请求头必须一致")
		}
		selected = value
	}
	if selected == "" {
		selected = envProjectID
		if strings.TrimSpace(selected) != selected {
			return fmt.Errorf("UR_PROJECT_ID 不得包含首尾空白")
		}
	}
	// 归一化大小写，避免 HTTP Header.Set 的 map 遍历顺序造成非确定覆盖。
	for key := range headers {
		if strings.EqualFold(key, "project-id") {
			delete(headers, key)
		}
	}
	if selected != "" {
		headers["project-id"] = selected
	}
	return nil
}
