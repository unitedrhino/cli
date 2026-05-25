package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "生成 shell 自动补全脚本",
	Long: `生成指定 shell 的自动补全脚本。

使用示例:
  # Bash
  source <(ur completion bash)

  # Zsh
  source <(ur completion zsh)

  # Fish
  ur completion fish | source

  # PowerShell
  ur completion powershell | Out-String | Invoke-Expression`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			return cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			return cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			return &CLIError{Message: "不支持的 shell: " + args[0], ExitCode: 2}
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
