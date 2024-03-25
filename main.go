package main

import (
	_ "embed"
	"os"

	"github.com/gabe565/pwgen-go/cmd"
)

//nolint:gochecknoglobals
var (
	version = "beta"
	commit  = ""
)

func main() {
	rootCmd := cmd.New(version, commit)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
