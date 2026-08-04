// helper.go — view 命令组共享的 Cobra 包装辅助函数
package view

import (
	"context"
	"io"

	"github.com/spf13/cobra"
)

// CLIError 提供可控退出码的错误类型
type CLIError struct {
	Message  string
	ExitCode int
}

func (e CLIError) Error() string { return e.Message }

// exitCodeCLIError 为 CLIError 补充 ExitCode() 方法（CLIError 字段名冲突无法直接加
// 同名方法，用普通字段持有而非嵌入），使 Execute 能按预期退出码返回
type exitCodeCLIError struct {
	err *CLIError
}

// Error 返回错误信息
func (e *exitCodeCLIError) Error() string { return e.err.Message }

// ExitCode 返回命令预期退出码
func (e *exitCodeCLIError) ExitCode() int { return e.err.ExitCode }

// wrapOldCommand 将 cmd/shared 中手动解析参数的旧式 run 函数包装为 Cobra RunE
func wrapOldCommand(fn func(ctx context.Context, args []string, stdout, stderr io.Writer) int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		code := fn(cmd.Context(), args, cmd.OutOrStdout(), cmd.ErrOrStderr())
		if code != 0 {
			return &exitCodeCLIError{err: &CLIError{Message: "", ExitCode: code}}
		}
		return nil
	}
}
