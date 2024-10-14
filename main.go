package main

import (
	_ "embed"
	"os"

	"gabe565.com/pwgen/cmd"
)

var version = "beta"

func main() {
	rootCmd := cmd.New(cmd.WithVersion(version))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
