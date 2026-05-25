package cmdutil

import (
	"io"
	"os"

	"gitee.com/unitedrhino/cli/internal/config"
)

// Factory 提供命令执行所需的共享依赖
type Factory struct {
	IOStreams
	Config *config.Config
	App    config.CLIApp
}

// IOStreams 封装输入输出，便于测试注入
type IOStreams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
}

// NewFactory 从环境构建 Factory
func NewFactory(app config.CLIApp) (*Factory, error) {
	cfg, err := config.ReadConfig()
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	return &Factory{
		IOStreams: IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
		Config: &cfg,
		App:    app,
	}, nil
}
