package main

import (
	_ "embed"
	"os"

	"github.com/gabe565/pwgen-go/cmd"
)

func main() {
	rootCmd := cmd.New()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
