package cmd

import (
	"github.com/spf13/cobra"
	"gitee.com/unitedrhino/cli/cmd/shared"
)

var generateSkillsOutputFlag string

var generateSkillsCmd = &cobra.Command{
	Use:   "generate-skills",
	Short: "从 swagger 生成 Skill 文档",
	Long:  `读取 swagger JSON 生成 Skill 文档，供 AI 工具使用。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app := resolveAppFromContext()
		// 将 cobra flag 重新组装为 RunGenerateSkills 期望的 args 格式
		if generateSkillsOutputFlag != "" {
			args = append(args, "--output", generateSkillsOutputFlag)
		}
		code := shared.CobraBridge{}.RunGenerateSkills(app, args, cmd.OutOrStdout(), cmd.ErrOrStderr())
		if code != 0 {
			return &CLIError{Message: "", ExitCode: code}
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(generateSkillsCmd)
	generateSkillsCmd.Flags().StringVarP(&generateSkillsOutputFlag, "output", "o", "", "指定输出目录")
}
