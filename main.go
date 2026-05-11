package main

import (
	"context"
	"os"

	"gitee.com/unitedrhino/cli/cmd/ur"
)

func main() {
	os.Exit(ur.Execute(context.Background(), os.Args[1:], os.Stdout, os.Stderr))
}
