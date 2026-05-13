package main

import (
	"context"
	"fmt"
	"os"

	"gitee.com/unitedrhino/cli/cmd/ur"
)

// version 由 release.sh 在构建时通过 -ldflags -X main.version=VERSION 注入
var version = "dev"

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-v" {
			fmt.Println(version)
			os.Exit(0)
		}
	}
	os.Exit(ur.Execute(context.Background(), version, os.Args[1:], os.Stdout, os.Stderr))
}
