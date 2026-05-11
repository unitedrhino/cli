package main

import (
	"context"
	"os"

	"gitee.com/unitedrhino/cli/cmd/shared"
	"gitee.com/unitedrhino/cli/internal/config"
)

func main() {
	os.Exit(shared.Execute(config.AppConsole, context.Background(), os.Args[1:], os.Stdout, os.Stderr))
}
