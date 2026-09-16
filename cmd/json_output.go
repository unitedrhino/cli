// json_output.go 统一根命令中的 JSON 对象输出，并附加可供 AI 读取的提示。
package cmd

import (
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/notice"
)

// writeJSON 将值格式化后写入目标输出；对象会按需包含顶层 _notice。
func writeJSON(writer io.Writer, value any) error {
	raw, err := notice.MarshalJSON(value, false)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(writer, string(raw))
	return err
}
