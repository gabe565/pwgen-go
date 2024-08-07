package main

import (
	_ "embed"
	"os"

	"github.com/gabe565/pwgen-go/cmd"
)

var version = "beta"

func main() {
	rootCmd := cmd.New(cmd.WithVersion(version))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
