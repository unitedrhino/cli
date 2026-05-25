// util.go — 旧命令共享的通用工具函数
package shared

import (
	"encoding/json"
	"fmt"
	"strings"
)

// parseBodyArg 解析 JSON body 参数
func parseBodyArg(raw string) (map[string]any, error) {
	if strings.TrimSpace(raw) == "" {
		return map[string]any{}, nil
	}
	var out map[string]any
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		return nil, fmt.Errorf("body must be JSON object: %w", err)
	}
	return out, nil
}
