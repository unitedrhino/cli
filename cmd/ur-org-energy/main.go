package main

import (
	"context"
	"fmt"
	"os"

	"gitee.com/unitedrhino/cli/cmd/shared"
	"gitee.com/unitedrhino/cli/internal/config"
)

var version = "dev"

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-v" {
			fmt.Println(version)
			os.Exit(0)
		}
	}
	os.Exit(shared.Execute(config.AppOrgEnergy, context.Background(), version, os.Args[1:], os.Stdout, os.Stderr))
}
