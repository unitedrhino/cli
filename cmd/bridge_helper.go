// bridge_helper.go — 旧命令桥接到 Cobra 的通用工具
package cmd

import (
	"context"
	"io"

	"github.com/spf13/cobra"
)

// oldCommandFunc 不需要 app 的旧命令签名
type oldCommandFunc func(ctx context.Context, args []string, stdout, stderr io.Writer) int

// oldCommandFuncNoCtx 不需要 context 的旧命令签名
type oldCommandFuncNoCtx func(args []string, stdout, stderr io.Writer) int

// wrapOldCommand 将需要 ctx 的旧命令包装为 cobra RunE
func wrapOldCommand(fn oldCommandFunc) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		code := fn(cmd.Context(), args, cmd.OutOrStdout(), cmd.ErrOrStderr())
		if code != 0 {
			return &CLIError{Message: "", ExitCode: code}
		}
		return nil
	}
}

// wrapOldCommandNoCtx 将不需要 ctx 的旧命令包装为 cobra RunE
func wrapOldCommandNoCtx(fn oldCommandFuncNoCtx) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		code := fn(args, cmd.OutOrStdout(), cmd.ErrOrStderr())
		if code != 0 {
			return &CLIError{Message: "", ExitCode: code}
		}
		return nil
	}
}
