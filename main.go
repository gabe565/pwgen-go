package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"

	"gabe565.com/pwgen/cmd"
	"gabe565.com/pwgen/internal/config"
	"gabe565.com/utils/cobrax"
)

var version = "beta"

func main() {
	root := cmd.New(cobrax.WithVersion(version))
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintln(root.ErrOrStderr(), root.ErrPrefix(), err.Error())
		if errors.Is(err, config.ErrProfileNotFound) {
			_, _ = fmt.Fprintf(root.ErrOrStderr(),
				"  Hint: Run %q for a list of available profiles.\n",
				os.Args[0]+" profiles")
		}
		os.Exit(1)
	}
}
