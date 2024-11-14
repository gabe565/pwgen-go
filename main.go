package main

import (
	_ "embed"
	"os"

	"gabe565.com/pwgen/cmd"
	"gabe565.com/utils/cobrax"
)

var version = "beta"

func main() {
	rootCmd := cmd.New(cobrax.WithVersion(version))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
