// json_output.go 统一给命令 JSON 结果附加 AI 可解析提示。
package shared

import (
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/notice"
)

// writeJSON 序列化对象并在可用时合并顶层 _notice。
func writeJSON(writer io.Writer, value any) error {
	raw, err := notice.MarshalJSON(value, true)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(writer, string(raw))
	return err
}
