package ai

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

func wrapOldCommand(fn func(ctx context.Context, args []string, stdout, stderr io.Writer) int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		code := fn(cmd.Context(), args, cmd.OutOrStdout(), cmd.ErrOrStderr())
		if code != 0 {
			return &CLIError{Message: "", ExitCode: code}
		}
		return nil
	}
}
