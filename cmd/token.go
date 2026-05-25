package cmd

import (
	"encoding/base64"
	"strings"

	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/internal/auth"
)

var tokenOpts struct {
	rawOutput bool
	decode    bool
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "获取当前认证 Token",
	Long:  `获取当前上下文的认证 Token，支持原始输出和解码查看。`,
	RunE:  runToken,
}

func init() {
	tokenCmd.Flags().BoolVarP(&tokenOpts.rawOutput, "raw", "r", false, "仅输出 token 字符串")
	tokenCmd.Flags().BoolVarP(&tokenOpts.decode, "decode", "d", false, "解码 JWT payload")
	RootCmd.AddCommand(tokenCmd)
}

func runToken(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	token, err := auth.ResolveToken(ctx)
	if err != nil {
		return &CLIError{Message: err.Error(), ExitCode: 1}
	}
	if tokenOpts.rawOutput {
		cmd.Println(token)
		return nil
	}
	if tokenOpts.decode {
		parts := strings.Split(token, ".")
		if len(parts) != 3 {
			return &CLIError{Message: "token is not a JWT", ExitCode: 1}
		}
		payload, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			return &CLIError{Message: err.Error(), ExitCode: 1}
		}
		cmd.Println(string(payload))
		return nil
	}
	cmd.Println(token)
	return nil
}
